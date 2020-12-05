package v1

const CONTROLLER = `
package controller

/*
date:{{ .now }}
*/

import (
  "github.com/gin-gonic/gin"
  "🌴🌴/db"
  "🌴🌴/model"
)

// Add{{ .ModelName }} add
func Add{{ .ModelName }}(c *gin.Context) interface{}{
   	var req = dbaccess.{{ .ModelName }}{}
   	err := c.ShouldBind(&req)
   	if err != nil {
       return Response{Code: -1, Msg: err.Error()}
   	}
   	data,err := model.Add{{ .ModelName }}(&req)
   	if err != nil{
       return Response{Code: -1, Msg: err.Error()}
   	}
	return Response{Code: 0, Msg: "添加成功", Data: data}
}

// Update{{ .ModelName }} update
func Update{{ .ModelName }}(c *gin.Context) interface{}{
   	var req = dbaccess.{{ .ModelName }}{}
   	err := c.ShouldBind(&req)
   	if err != nil {
       return Response{Code: -1, Msg: err.Error()}
   	}
   	data,err := model.Update{{ .ModelName }}(&req)
    if err != nil{
       return Response{Code: -1, Msg: err.Error()}
    }
    return Response{Code: 0, Msg: "更新成功", Data: data}
}

// Get{{ .ModelName }}ByID  get xxx by id
func Get{{ .ModelName }}ByID(c *gin.Context) interface{}{
    _id := c.Param("oid")
    id, err := strconv.Atoi(_id)
    if err != nil {
      return Response{Code: -1, Msg: err.Error()}
    }
     data,err := model.Get{{ .ModelName }}ByID(id)
     if err != nil{
       return Response{Code: -1, Msg: err.Error()}
     }
     return Response{Code: 0, Msg: "ok", Data: data}
}

// List{{ .ModelName }} // list by page condition
func List{{ .ModelName }}(c *gin.Context) interface{}{
    var req = dbaccess.{{ .ModelName }}{}
    err := c.ShouldBind(&req)
    if err != nil {
       return Response{Code: -1, Msg: err.Error()}
    }
    data,err := model.List{{ .ModelName }}(&req)
    if err != nil{
       return Response{Code: -1, Msg: err.Error()}
     }
     return Response{Code: 0, Msg: "ok", Data: data}
}

// Delete{{ .ModelName }} Delete
func Delete{{ .ModelName }}(c *gin.Context) interface{}{
    var req = dbaccess.{{ .ModelName }}{}
    err := c.ShouldBind(&req)
    if err != nil {
       return Response{Code: -1, Msg: err.Error()}
    }
    err = model.Delete{{ .ModelName }}(req.ID)
    if err != nil{
       return Response{Code: -1, Msg: err.Error()}
    }
    return Response{Code: 0, Msg: "删除成功"}
}

/*
	路由

	你可以走rest风格
	🐲🐲 := r.Group(prefix + "/👌👌👌")
	{ 
		🐲🐲.POST("/add", route(controller.Add{{ .ModelName }}))
		🐲🐲.POST("/update", route(controller.Update{{ .ModelName }}))
		🐲🐲.GET("/list/:oid", route(controller.Get{{ .ModelName }}ByID))
		🐲🐲.GET("/list", route(controller.List{{ .ModelName }}))
		🐲🐲.GET("/delete", route(controller.Delete{{ .ModelName }}))
	}

*/

func route(f func(ctx *gin.Context)interface{})gin.HandlerFunc{
	return func(context *gin.Context) {
		context.JSON(200, f(context))
	}
}

`
