package tpl

const SERVICE = `
package service

/*
date:{{ .now }}
*/

import (
  "🌴🌴/dbaccess"
)

// Add{{ .ModelName }} add
func Add{{ .ModelName }}(b *dbaccess.{{ .ModelName }})(*{{ .ModelName }},error){
       return  dbaccess.Add{{ .ModelName }}(b)
}

// Update{{ .ModelName }} update
func Update{{ .ModelName }}(b *dbaccess.{{ .ModelName }})(*{{ .ModelName }},error){
     return  dbaccess.Update{{ .ModelName }}(b)
}

// Get{{ .ModelName }}ByID get by id
func Get{{ .ModelName }}ByID(id int)(*dbaccess.{{ .ModelName }},error){
     return dbaccess.Get{{ .ModelName }}ByID(id)
}

// List{{ .ModelName }}  page by condition
func List{{ .ModelName }}(b *dbaccess.{{ .ModelName }})(*dbaccess.DataStore,error){
     list,err := dbaccess.List{{ .ModelName }}(b)
     if err != nil{
        return nil,err
     }
     total,err := dbaccess.Count{{ .ModelName }}(b)
     if err != nil{
             return nil,err
      }
    return &dbaccess.DataStore{Total:total,Data:list,TotalPage:(int(total)+b.PageSize-1)/b.PageSize} ,nil
}

// Delete{{ .ModelName }} delete
func Delete{{ .ModelName }}(id int) error {
   return dbaccess.Delete{{ .ModelName }}(id)
}

`
