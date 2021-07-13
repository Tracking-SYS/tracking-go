package services

import (
	"github.com/google/wire"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ProductProvider,
	wire.Bind(new(ProductServiceInterface), new(*ProductService)),

	TaskProvider,
	wire.Bind(new(TaskServiceInterface), new(*TaskService)),
)
