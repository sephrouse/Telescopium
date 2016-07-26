package controllers

import (
	"Telescopium/app/routes"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) NotFound() revel.Result {
	return c.Render()
}

func (c App) StockIndex() revel.Result {
	return c.Render()
}

func (c App) CustomStock() revel.Result {
	return c.Render()
}

func (c App) BuyStock() revel.Result {
	return c.Render()
}

func (c App) SellStock() revel.Result {
	return c.Render()
}

func (c App) MyAsset() revel.Result {
	return c.Render()
}

func (c App) Profile() revel.Result {
	return c.Render()
}

func (c App) Logout() revel.Result {
	return c.Redirect(routes.App.Login())
}
