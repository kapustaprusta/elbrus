package elbrus

import "net/http"

// Start ...
func Start(config *Config) error {
	srv := newServer()
	srv.logger.Infof("started listening on %s\n", config.BindAddr)

	return http.ListenAndServe(config.BindAddr, srv)
}
