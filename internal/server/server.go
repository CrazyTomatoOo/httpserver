package server

import "HttpServer/pkg/database"

type Server struct {
	DB *database.Database
}

func NewServer() *Server {
	return &Server{
		DB: &database.Database{},
	}
}

func (s *Server) Init() {
	s.DB.Init()
}
