package server

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: 8090,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	s.server.Run(":" + s.port)
}