package api

// Base API server instance description
type API struct {
}

// API constructor: build base API instance
func New() *API {
	return &API{}
}

// Start http server, configure loggers, router, database connection and etc...
func (api *API) Start() error {
	return nil
}
