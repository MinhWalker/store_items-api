package services

import (
	"github.com/MinhWalker/store_items-api/src/domain/items"
	"github.com/MinhWalker/store_utils-go/rest_errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("emplement me!", http.StatusNotImplemented, "Not_emplement", nil)
}
