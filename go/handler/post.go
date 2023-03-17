package handler

import (
	"fmt"
	"strconv"

	"github.com/MuShaf-NMS/SV-BE-Test/dto"
	"github.com/MuShaf-NMS/SV-BE-Test/helper"
	"github.com/MuShaf-NMS/SV-BE-Test/service"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	GetAll(ctx *gin.Context)
	GetAllFilter(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type postHandler struct {
	service service.PostService
}

func (ph *postHandler) GetAll(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Limit must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	offset, err := strconv.Atoi(ctx.Param("offset"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Offset must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	posts, err := ph.service.GetAll(limit, offset, "")
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Message, e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	ctx.JSON(200, posts)
}

func (ph *postHandler) GetAllFilter(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Limit must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	offset, err := strconv.Atoi(ctx.Param("offset"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Offset must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	status := ctx.Param("status")
	posts, err := ph.service.GetAll(limit, offset, status)
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Message, e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	ctx.JSON(200, posts)
}

func (ph *postHandler) Create(ctx *gin.Context) {
	var post dto.PostDTO
	ctx.BindJSON(&post)
	if err := helper.Validate(post); err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder("Validation error", errs)
		ctx.JSON(400, res)
		return
	}
	err := ph.service.Create(post)
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Message, e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	ctx.JSON(200, gin.H{})
}

func (ph *postHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Id must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	post, err := ph.service.GetOne(uint(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Message, e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	ctx.JSON(200, post)
}

func (ph *postHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Offset must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	var post dto.PostDTO
	ctx.BindJSON(&post)
	if err := helper.Validate(post); err != nil {
		errs := helper.ValidationError(err)
		fmt.Println(errs)
		res := helper.ErrorResponseBuilder("Validation error", errs)
		ctx.JSON(400, res)
		return
	}
	err = ph.service.Update(post, uint(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Message, e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	ctx.JSON(200, gin.H{})
}

func (ph *postHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Offset must be integer", nil)
		ctx.JSON(400, res)
		return
	}
	err = ph.service.Delete(uint(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Message, e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	ctx.JSON(200, gin.H{})
}

func NewPostHandler(service service.PostService) PostHandler {
	return &postHandler{
		service: service,
	}
}
