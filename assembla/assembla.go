package assembla

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.assembla.com"
	apiVersion = "/v1/"
)

type authType int


const (
	apiKeySecret authType = iota
)

type Client struct {
	client *http.Client
	baseUrl *url.URL
	apiKey, apiSecret string
	authType authType
	UserService *UserService
}

func NewClient(httpClient *http.Client, apiKey string, apiSecret string) *Client{
	client := newClient(httpClient)
	err := client.SetBaseURL(defaultBaseURL + apiVersion)
	if err != nil{
		panic(err)
	}
	client.apiKey = apiKey
	client.apiSecret = apiSecret
	return client
}

func newClient(httpClient *http.Client) *Client {
	if httpClient == nil{
		httpClient = http.DefaultClient
	}
	c := &Client{client:httpClient}
	c.UserService = &UserService{client:c}
	return c
}

func (c *Client) SetBaseURL(urlStr string) error {
	baseUrl,err := url.Parse(urlStr)
	if err != nil{
		return err
	}
	c.baseUrl = baseUrl
	return nil
}

func (c *Client) NewRequest(method,path string, opt interface{})(*http.Request,error){
	u := *c.baseUrl
	unescaped,err := url.PathUnescape(path)
	if err != nil{
		return nil,err
	}

	u.RawPath = c.baseUrl.Path + path
	u.Path = c.baseUrl.Path + unescaped

	req := &http.Request{
		Method:method,
		URL: &u,
		Proto: "HTTP/1.1",
		ProtoMajor:1,
		ProtoMinor:1,
		Header:make(http.Header),
		Host:u.Host,
	}

	if method == "POST" || method == "PUT"{
		bodyBytes,err := json.Marshal(opt)
		if err != nil{
			return nil,err
		}
		bodyReader := bytes.NewReader(bodyBytes)

		u.RawQuery=""
		req.Body = ioutil.NopCloser(bodyReader)
		req.GetBody = func()(io.ReadCloser,error){
			return ioutil.NopCloser(bodyReader),nil
		}
		req.ContentLength = int64(bodyReader.Len())
		req.Header.Set("Content-Type","application/json")
	}

	req.Header.Set("Accept","application/json")

	switch c.authType {
	case apiKeySecret:
		req.Header.Set("X-Api-Key",c.apiKey)
		req.Header.Set("X-Api-Secret",c.apiSecret)
	}

	return req,nil
}

func (c *Client) Do(req *http.Request,v interface{})(*http.Response,error){
	resp,err := c.client.Do(req)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest{
		return nil,fmt.Errorf("error occurred executing request : %s",resp.Status)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	if err != nil{
		return nil,err
	}

	return resp,nil
}