package main

import "github.com/rscampos3/GoLang/tree/main/go-live/model"

func main() {
	produto1 := model.NewProduct()
	produto1.Name = "Carrinho"

	produto2 := model.NewProduct()
	produto2.Name = "Boneca"

	products := model.Products{}
	products.Add(produto1)
	products.Add(produto2)
}
