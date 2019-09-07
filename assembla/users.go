package assembla

import (
	"net/http"
)

type UserService struct {
	client *Client
}

type User struct {
	ID string `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Picture string `json:"picture"`
	Email string `json:"email"`
	Organization string `json:"organization"`
	Phone string `json:"phone"`
}


//GET /v1/user
//Returns currently authenticated user
func (s *UserService) ListAuthenticatedUser() (User,*http.Response,error){
	req,err := s.client.NewRequest(http.MethodGet,"/user",nil)
	if err != nil{
		return User{},nil,err
	}

	var user User
	resp,err := s.client.Do(req,&user)
	if err != nil{
		return User{},resp,err
	}
	return user,resp,nil
}