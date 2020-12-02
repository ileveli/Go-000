package week02

import ( 
	"database/sql" 
	"fmt" 
  
	"github.com/pkg/errors" 
 ) 
  
 func Dao() error { 
	return sql.ErrNoRows 
 } 
  
 func LoginService() error { 
	return errors.Wrap(dao(), "user is not exists") 
 } 

 func Api() error {
	 return LoginService()
 }
  
 func main() { 
	err := Api() 
	if err != nil { 
	   fmt.Println(err)
	   return 
	} 
	fmt.Println("ok!")
 }
