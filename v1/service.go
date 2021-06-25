package v1

const SERVICE = `
package service

/*
date:{{ .now }}
*/

import (
  "🌴🌴/dbaccess"
)

// Add{{ .ModelName }} add
func Add{{ .ModelName }}(req *dbaccess.{{ .ModelName }})(*dbaccess.{{ .ModelName }},error){
       return  dbaccess.Add{{ .ModelName }}(req)
}

// Update{{ .ModelName }} update
func Update{{ .ModelName }}(req *dbaccess.{{ .ModelName }})(*dbaccess.{{ .ModelName }},error){
     return  dbaccess.Update{{ .ModelName }}(req)
}

// Get{{ .ModelName }}ByID get by id
func Get{{ .ModelName }}ByID(id int)(*dbaccess.{{ .ModelName }},error){
     return dbaccess.Get{{ .ModelName }}ByID(id)
}

// List{{ .ModelName }}  page by condition
// func List{{ .ModelName }}(req *dbaccess.{{ .ModelName }})(*dbaccess.DataStore,error){
func List{{ .ModelName }}(req *db.CommonRequestParams) (interface{}, error) {
	return db.ListDeviceMobile(req)
    //  list,err := dbaccess.List{{ .ModelName }}(req)
    //  if err != nil{
    //     return nil,err
    //  }
    //  total,err := dbaccess.Count{{ .ModelName }}(req)
    //  if err != nil{
    //          return nil,err
    //   }
    // return &dbaccess.DataStore{Total:total,Data:list,TotalPage:(int(total)+req.PageSize-1)/req.PageSize} ,nil
}

// Delete{{ .ModelName }} delete
func Delete{{ .ModelName }}(id int) error {
   return dbaccess.Delete{{ .ModelName }}(id)
}

`
