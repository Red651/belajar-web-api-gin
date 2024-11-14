package main

import (
	siswacontroller "github.com/Red651/belajar-web-api/controller/siswaController"
	"github.com/Red651/belajar-web-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("api/show", siswacontroller.Index)
	r.GET("api/show/:id", siswacontroller.ShowById)
	r.GET("api/show/nama/:nama", siswacontroller.ShowByNama)
	r.POST("api/insert", siswacontroller.Insert)
	r.PUT("api/Update/:id", siswacontroller.Update)
	r.DELETE("api/delete/:id", siswacontroller.Delete)

	r.Run()
}
