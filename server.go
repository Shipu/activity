package main

import (
	_app "github.com/shipu/tracker/bootstrap"
	_api "github.com/shipu/tracker/routes/api"
)

func main() {

	_app.Init()

	_api.Load()

	_app.Start()
}
