package services

import (
	"github.com/bheemeshkammak/Code/apitesting/pkg/rest/server/daos"
	"github.com/bheemeshkammak/Code/apitesting/pkg/rest/server/models"
)

type ApiService struct {
	apiDao *daos.ApiDao
}

func NewApiService() (*ApiService, error) {
	apiDao, err := daos.NewApiDao()
	if err != nil {
		return nil, err
	}
	return &ApiService{
		apiDao: apiDao,
	}, nil
}

func (apiService *ApiService) CreateApi(api *models.Api) (*models.Api, error) {
	return apiService.apiDao.CreateApi(api)
}

func (apiService *ApiService) UpdateApi(id int64, api *models.Api) (*models.Api, error) {
	return apiService.apiDao.UpdateApi(id, api)
}

func (apiService *ApiService) DeleteApi(id int64) error {
	return apiService.apiDao.DeleteApi(id)
}

func (apiService *ApiService) ListApis() ([]*models.Api, error) {
	return apiService.apiDao.ListApis()
}

func (apiService *ApiService) GetApi(id int64) (*models.Api, error) {
	return apiService.apiDao.GetApi(id)
}
