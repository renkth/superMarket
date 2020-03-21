package main

import (
	"log"
	"net/http"
	"superMarket/controllers"
)

func main() {
	goodsController := &controllers.GoodsController{}
	http.HandleFunc("/goods/create", goodsController.Create)
	http.HandleFunc("/goods/update", goodsController.Update)
	http.HandleFunc("/goods/delete", goodsController.Delete)
	http.HandleFunc("/goods/selectById", goodsController.SelectById)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}