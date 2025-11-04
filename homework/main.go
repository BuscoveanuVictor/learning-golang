package main

import (
	"fmt"
)
func main(){
	var a int
	fmt.Scan(&a)
	if(a==1){
		server()
	}else{
		client()
	}

}