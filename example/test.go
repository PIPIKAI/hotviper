package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/hotviper"
	"github.com/pipikai/hotviper/example/common/result"
)

func main() {
	config, err := hotviper.NewHotViper("config", "yaml", "config")
	if err != nil {
		panic(err)
	}
	g := gin.Default()

	//yaml.Unmarshal会根据yaml标签的字段进行赋值
	g.GET("/", func(ctx *gin.Context) {
		result.OkWithData(ctx, config.GetConfig())
	})
	g.GET("/edit", func(ctx *gin.Context) {
		edited := ctx.Query("value")
		log.Printf("value : %s", edited)
		err := config.SetConfig(edited)
		if err != nil {
			log.Println(err)
			result.FailWithMsg(ctx, result.ResultError, err.Error())
		}
	})
	// 将 config 设置为 默认
	g.GET("/reset", func(ctx *gin.Context) {
		err := config.SetDefault()
		if err != nil {
			log.Println(err)
			result.FailWithMsg(ctx, result.ResultError, err.Error())
		} else {
			result.Ok(ctx)
		}

	})

	// 将 config roll back
	g.GET("/rollback", func(ctx *gin.Context) {
		err := config.RollBack()
		if err != nil {
			log.Println(err)
			result.FailWithMsg(ctx, result.ResultError, err.Error())
		} else {
			result.Ok(ctx)
		}
	})
	g.Run(":3222")
}
