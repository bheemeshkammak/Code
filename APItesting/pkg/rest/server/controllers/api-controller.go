package controllers

import (
	"github.com/bheemeshkammak/Code/apitesting/pkg/rest/server/models"
	"github.com/bheemeshkammak/Code/apitesting/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ApiController struct {
	apiService *services.ApiService
}

func NewApiController() (*ApiController, error) {
	apiService, err := services.NewApiService()
	if err != nil {
		return nil, err
	}
	return &ApiController{
		apiService: apiService,
	}, nil
}

func (apiController *ApiController) CreateApi(context *gin.Context) {
	// validate input
	var input models.Api
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger api creation
	if _, err := apiController.apiService.CreateApi(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Api created successfully"})
}

func (apiController *ApiController) UpdateApi(context *gin.Context) {
	// validate input
	var input models.Api
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger api update
	if _, err := apiController.apiService.UpdateApi(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Api updated successfully"})
}

func (apiController *ApiController) FetchApi(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger api fetching
	api, err := apiController.apiService.GetApi(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, api)
}

func (apiController *ApiController) DeleteApi(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger api deletion
	if err := apiController.apiService.DeleteApi(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Api deleted successfully",
	})
}

func (apiController *ApiController) ListApis(context *gin.Context) {
	// trigger all apis fetching
	apis, err := apiController.apiService.ListApis()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, apis)
}

func (*ApiController) PatchApi(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*ApiController) OptionsApi(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*ApiController) HeadApi(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
