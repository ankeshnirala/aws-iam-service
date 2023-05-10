package api

import (
	"context"
	"net/http"

	"github.com/ankeshnirala/go/aws-iam-service/types"
	"github.com/ankeshnirala/go/aws-iam-service/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetUsers(ctx *gin.Context) {
	var data []*types.User

	cursor, err := s.mongoStore.Find()
	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	err = cursor.All(context.TODO(), &data)
	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	util.WriteJSON(ctx, http.StatusOK, data)
}

func (s *Server) handleGetUserProfile(ctx *gin.Context) {
	userID, isTrue := ctx.Get("userID")

	if !isTrue {
		util.WriteJSON(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var data *types.User
	result, err := s.mongoStore.FindByID(userID.(string))

	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	result.Decode(&data)
	util.WriteJSON(ctx, http.StatusOK, data)
}
