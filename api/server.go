package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngenohkevin/go-auth/db/store"
	"github.com/ngenohkevin/go-auth/utils"
)

type Server struct {
	config utils.Config
	store  *store.Store
	router *gin.Engine
}

//func NewServer(store *store.Store) *Server {
//	server := &Server{
//		store: store,
//	}
//	router := gin.Default()
//
//}
