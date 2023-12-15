package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngenohkevin/go-auth/db/store"
	"github.com/ngenohkevin/go-auth/token"
	"github.com/ngenohkevin/go-auth/utils"
)

type Server struct {
	config     utils.Config
	store      *store.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store *store.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
