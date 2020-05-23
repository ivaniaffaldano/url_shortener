package helpers

type ErrorRequest struct {
	Error        	string    `json:"error"`
}

type DeletedResponse struct {
	Deleted        	bool    `json:"deleted"`
}