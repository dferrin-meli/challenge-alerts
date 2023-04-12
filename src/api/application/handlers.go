package application

func (r *Server) AddHandlers() *Server {
	r.GET("/ping", ping)
	return r
}
