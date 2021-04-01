package controllers

import (
	"fmt"
	"github.com/MinhWalker/store_items-api/src/domain/items"
	"github.com/MinhWalker/store_items-api/src/services"
	"github.com/MinhWalker/store_oauth-go/oauth"
	"net/http"
)

var(
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request)  {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error to the user.
		return
	}

	item := items.Item{
		Seller:            oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: Return error json to user.
	}

	fmt.Println(result)
	//TODO: Return created item with HTTP status 201 - Created.
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request)  {
	
}