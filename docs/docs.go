// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ivan Iaffaldano",
            "email": "i.iaffaldano@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/create": {
            "post": {
                "description": "get short URL from dest URL, params example: {\"destination_url\":\"example\"}",
                "summary": "Create Short URL",
                "parameters": [
                    {
                        "description": "destination URL",
                        "name": "destination_url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorRequest"
                        }
                    }
                }
            }
        },
        "/api/delete": {
            "delete": {
                "description": "delete URL from Id, params example: http://localhost:8080/api/delete/2SHcWFg",
                "summary": "Delete Url",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.DeletedResponse"
                        }
                    }
                }
            }
        },
        "/api/get/{shortUrl}": {
            "get": {
                "description": "get Destination URL from Short Url, params example: http://localhost:8080/api/get/2SHcWFg",
                "summary": "Get Destination URL from Short Url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short URL",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Url"
                        }
                    }
                }
            }
        },
        "/{shortUrl}": {
            "get": {
                "description": "Redirect to Destination URL from Short Url, params example: http://localhost:8080/2SHcWFg",
                "summary": "Redirect Url to DestinationUrl",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short URL",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorRequest"
                        }
                    },
                    "301": {}
                }
            }
        }
    },
    "definitions": {
        "helpers.DeletedResponse": {
            "type": "object",
            "properties": {
                "deleted": {
                    "type": "boolean"
                }
            }
        },
        "helpers.ErrorRequest": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.Url": {
            "type": "object",
            "properties": {
                "destination_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "short_url": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "URL Shortener API",
	Description: "This is a test.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
