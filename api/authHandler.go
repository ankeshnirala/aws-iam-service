package api

import (
	"encoding/json"
	"net/http"

	"github.com/ankeshnirala/go/aws-iam-service/types"
	"github.com/ankeshnirala/go/aws-iam-service/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleSignup(ctx *gin.Context) {

	// sync request body data with SignupRequest
	req := new(types.SignupRequest)
	if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// check if user already exist with email
	isUserExist := s.mongoStore.GetUserByEmail(req.Email)
	if isUserExist != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, "email already exist")
		return
	}

	// create new user if everythig is fine
	user, err := types.NewRootUser(req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// insert created user in db
	insertRes, err := s.mongoStore.CreateUser(user)
	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// send the response
	util.WriteJSON(ctx, http.StatusOK, insertRes)
}

func (s *Server) handleMemberSignup(ctx *gin.Context) {

	createdBy, isTrue := ctx.Get("userID")

	if !isTrue {
		util.WriteJSON(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// sync request body data with SignupRequest
	req := new(types.SignupRequest)
	if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// check if user already exist with email
	isUserExist := s.mongoStore.GetUserByEmail(req.Email)
	if isUserExist != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, "email already exist")
		return
	}

	// create new user if everythig is fine
	user, err := types.NewUser(req.FirstName, req.LastName, req.Email, req.Password, createdBy.(string))
	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// insert created user in db
	insertRes, err := s.mongoStore.CreateUser(user)
	if err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// send the response
	util.WriteJSON(ctx, http.StatusOK, insertRes)
}

func (s *Server) handleLogin(ctx *gin.Context) {

	// validate login requuest params
	req := new(types.LoginRequest)
	if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
		util.WriteJSON(ctx, http.StatusBadRequest, err)
		return
	}

	// check if user does not exist with email
	// entered password is wring
	user := s.mongoStore.GetUserByEmail(req.Email)
	if user == nil || !user.ValidPassword(req.Password) {
		util.WriteJSON(ctx, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// now we need to create JWT for the user
	token, err := util.CreateJWT(user)
	if err != nil {
		util.WriteJSON(ctx, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Send JWT token as response
	// ctx.SetCookie("jwt-api-key", token, 3600, "/", "localhost", false, true)
	// Set the token in the response header
	ctx.Header("Authorization", "Bearer "+token)

	loginRes := types.LoginResponse{
		Email: req.Email,
		Token: token,
	}

	util.WriteJSON(ctx, http.StatusOK, loginRes)
}
