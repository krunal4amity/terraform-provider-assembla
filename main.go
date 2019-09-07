package main

import (
	"fmt"
	"github.com/krunal4amity/terraform-provider-assembla/assembla"
)

func main(){
	c := assembla.NewClient(nil,"06aa938f884f26db76e9","4c9a0ff184ead06bb5837fdac9357a5c88d5a5c8")
	u,_,err := c.Users.ListAuthenticatedUser()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(u)
	}

}