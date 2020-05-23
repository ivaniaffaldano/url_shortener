package route

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"url_shortener/app/controllers"
)

func GetRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// healthCheck
	router.Get("/", healthCheck) // return 204 to show service is up and running

	// GET route
	router.Get("/api/get/{shortUrl}", controllers.GetUrl)
	router.Get("/{shortUrl}", controllers.RedirectShortUrl)

	// POST route
	router.Post("/api/create", controllers.CreateUrl)

	// DELETE route
	router.Delete("/api/delete/{id}", controllers.DeleteUrl)

	// swagger Docs (only Local environment)
	appEnv := os.Getenv("app_env")
	if appEnv == "local" {
		router.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/docs/swagger.json"), //The url pointing to API definition"
		))
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, "docs"))
		FileServer(router, "/docs", filesDir) // expose swagger docs
	}

	return router;
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
