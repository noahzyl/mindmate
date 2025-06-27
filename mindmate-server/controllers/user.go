/*
 * Handle HTTP requests about users
 */

package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/mindmate/models"
	"github.com/noahzyl/mindmate/pkg/utils"
	"strconv"
)

type UserController struct{}

// UserRequest struct is used to save parameters from the frontend which are sent in JSON
type UserRequest struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// User struct sent to the frontend
type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u *UserController) Register(ctx *gin.Context) {
	user := &UserRequest{}
	err := ctx.BindJSON(user) // Parse parameters in the request
	if err != nil {
		ReturnErrorJson(ctx, 2001, gin.H{"Parameters parsing error": err.Error()})
		return
	}

	if user.Username == "" || user.Password == "" || user.ConfirmPassword == "" {
		ReturnErrorJson(ctx, 2001, "Invalid empty data")
		return
	}

	if user.Password != user.ConfirmPassword {
		ReturnErrorJson(ctx, 2001, "Password and Password confirmation mismatch")
		return
	}

	userQuery, _ := models.GetUserInfoByUsername(user.Username)
	if userQuery.Id != 0 {
		ReturnErrorJson(ctx, 4001, "Username exists")
		return
	}

	// Hash password and add new user to database
	hashedPassword, cryptErr := utils.HashStringWithBcrypt(user.Password)
	if cryptErr != nil {
		ReturnErrorJson(ctx, 5001, gin.H{"Failed to add user": cryptErr.Error()})
		return
	}
	userId, dbErr := models.AddUser(user.Username, hashedPassword)
	if dbErr != nil {
		ReturnErrorJson(ctx, 5001, gin.H{"Failed to add user": dbErr.Error()})
		return
	}

	data := UserResponse{
		Id:       userId,
		Username: user.Username,
	}
	ReturnSuccessJson(ctx, 0, "Registration succeeded", data, 1)
}

func (u *UserController) Login(ctx *gin.Context) {
	user := &UserRequest{}
	err := ctx.BindJSON(user) // Parse parameters in the request
	if err != nil {
		ReturnErrorJson(ctx, 2001, gin.H{"Parameters parsing error": err.Error()})
		return
	}

	if user.Username == "" || user.Password == "" {
		ReturnErrorJson(ctx, 2001, "Invalid empty data")
		return
	}

	// Get user's data from database
	userInfo, _ := models.GetUserInfoByUsername(user.Username)
	if userInfo.Id == 0 {
		ReturnErrorJson(ctx, 2001, "Username or Password is incorrect")
		return
	}
	if !utils.CheckBcryptHash(user.Password, userInfo.Password) {
		ReturnErrorJson(ctx, 2001, "Username or Password is incorrect")
		return
	}

	// Set session
	session := sessions.Default(ctx)
	session.Set("user_id_"+strconv.Itoa(userInfo.Id), userInfo.Id)
	session.Save() // Save the session in redis

	// Login succeed
	data := UserResponse{
		Id:       userInfo.Id,
		Username: userInfo.Username,
	}
	ReturnSuccessJson(ctx, 0, "Login succeed", data, 1)
}
