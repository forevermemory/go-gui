package tpl

const CONTROLLER = `
package controller

/*
date:{{ .now }}
*/

import (
  "github.com/gin-gonic/gin"
  "🌴🌴/dbaccess"
  "🌴🌴/service"
)

// Add{{ .ModelName }} add
func Add{{ .ModelName }}(c *gin.Context) interface{}{
   	var u = dbaccess.{{ .ModelName }}{}
   	err := c.ShouldBind(&u)
   	if err != nil {
       return Response{Code: -1, Msg: err.Error()}
   	}
   	data,err := service.Add{{ .ModelName }}(&u)
   	if err != nil{
       return Response{Code: -1, Msg: err.Error()}
   	}
	  return Response{Code: 0, Msg: "ok", Data: data}
}

// Update{{ .ModelName }} update
func Update{{ .ModelName }}(c *gin.Context) interface{}{
   	var u = dbaccess.{{ .ModelName }}{}
   	err := c.ShouldBind(&u)
   	if err != nil {
       return Response{Code: -1, Msg: err.Error()}
   	}
   	data,err := service.Update{{ .ModelName }}(&u)
    if err != nil{
       return Response{Code: -1, Msg: err.Error()}
    }
    return Response{Code: 0, Msg: "ok", Data: data}
}

// Get{{ .ModelName }}ByID  get xxx by id
func Get{{ .ModelName }}ByID(c *gin.Context) interface{}{
    _id := c.Param("oid")
    id, err := strconv.Atoi(_id)
    if err != nil {
      return Response{Code: -1, Msg: err.Error()}
    }
     data,err := service.Get{{ .ModelName }}ByID(id)
     if err != nil{
       return Response{Code: -1, Msg: err.Error()}
     }
     return Response{Code: 0, Msg: "ok", Data: data}
}

// List{{ .ModelName }} // list by page condition
func List{{ .ModelName }}(c *gin.Context) interface{}{
    var u = dbaccess.{{ .ModelName }}{PageSize:10,PageNo:1}
    err := c.ShouldBind(&u)
    if err != nil {
       return Response{Code: -1, Msg: err.Error()}
    }
    data,err := service.List{{ .ModelName }}(&u)
    if err != nil{
       return Response{Code: -1, Msg: err.Error()}
     }
     return Response{Code: 0, Msg: "ok", Data: data}
}

// Delete{{ .ModelName }} Delete
func Delete{{ .ModelName }}(c *gin.Context) interface{}{
    var u = dbaccess.{{ .ModelName }}{}
    err := c.ShouldBind(&u)
    if err != nil {
       return Response{Code: -1, Msg: err.Error()}
    }
    err = service.Delete{{ .ModelName }}(u.ID)
    if err != nil{
       return Response{Code: -1, Msg: err.Error()}
    }
    return Response{Code: 0, Msg: "ok"}
}

/*
路由
  🐲🐲 := r.Group(prefix + "/👌👌👌")
  { 
    🐲🐲.POST("/add", route(controller.Add{{ .ModelName }}))
    🐲🐲.POST("/update", route(controller.Update{{ .ModelName }}))
    🐲🐲.GET("/list/:oid", route(controller.Get{{ .ModelName }}ById))
    🐲🐲.GET("/list", route(controller.List{{ .ModelName }}))
    🐲🐲.GET("/delete", route(controller.Delete{{ .ModelName }}))
  }

*/


`
