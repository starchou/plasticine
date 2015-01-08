// Copyright 2015 star Chou, All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package plasticine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/bitly/go-hostpool"
)

type context struct {
	url          string
	method       string
	params       map[string]interface{}
	data         interface{}
	req          *http.Request
	resp         *http.Response
	hostpoolresp hostpool.HostPoolResponse
	body         []byte
}

func newContext(url, method string, hr hostpool.HostPoolResponse) *context {
	var resp http.Response
	req := http.Request{
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	return &context{url, method, map[string]interface{}{}, nil, &req, &resp, hr, nil}
}

// Get returns *context with GET method.
func Get(url string, hr hostpool.HostPoolResponse) *context {
	return newContext(url, "GET", hr)
}

// Post returns *context with POST method.
func Post(url string, hr hostpool.HostPoolResponse) *context {
	return newContext(url, "POST", hr)
}

// Put returns *context with PUT method.
func Put(url string, hr hostpool.HostPoolResponse) *context {
	return newContext(url, "PUT", hr)
}

// Delete returns *context DELETE method.
func Delete(url string, hr hostpool.HostPoolResponse) *context {
	return newContext(url, "DELETE", hr)
}

// Head returns *context Head method.
func Head(url string, hr hostpool.HostPoolResponse) *context {
	return newContext(url, "HEAD", hr)
}

// setData returns *context set es data
func (c *context) setHostPoolResponse(hr hostpool.HostPoolResponse) *context {
	c.hostpoolresp = hr
	return c
}

// setData returns *context set es data
func (c *context) setMethod(method string) *context {
	c.method = method
	return c
}

// setData returns *context set es data
func (c *context) setData(data interface{}) *context {
	c.data = data
	return c
}

// setParams returns *context set url paramater
func (c *context) setParams(args map[string]interface{}) *context {
	c.params = args
	return c
}

// Body adds request raw body.
// it supports string and []byte.
func (c *context) Body(data interface{}) *context {
	switch t := data.(type) {
	case string:
		c.setBodyString(t)
	case []byte:
		c.setBodyBytes(t)
	default:
		c.setBodyJson(data)
	}
	return c
}

func (c *context) setBodyString(body string) {
	bf := bytes.NewBufferString(body)
	c.req.Body = ioutil.NopCloser(bf)
	c.req.ContentLength = int64(len(body))
}

func (c *context) setBodyBytes(body []byte) {
	bf := bytes.NewBuffer(body)
	c.req.Body = ioutil.NopCloser(bf)
	c.req.ContentLength = int64(len(body))
}

func (c *context) setBodyJson(data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if Debug {
		println("data:" + string(body))
	}
	c.setBodyBytes(body)
	c.req.Header.Set("Content-Type", "application/json")
	return nil
}

func (c *context) excute() (*http.Response, error) {
	if c.resp.StatusCode != 0 {
		return c.resp, nil
	}
	var paramBody string
	if len(c.params) > 0 {
		var buf bytes.Buffer
		for k, v := range c.params {
			switch reflect.TypeOf(v).String() {
			case "bool":
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(strconv.FormatBool(v.(bool))))
				buf.WriteByte('&')
			case "string":
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(v.(string)))
				buf.WriteByte('&')
			case "slice":
				for _, val := range v.([]string) {
					buf.WriteString(url.QueryEscape(k))
					buf.WriteByte('=')
					buf.WriteString(url.QueryEscape(val))
					buf.WriteByte('&')
				}
			default:
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", v)))
				buf.WriteByte('&')
			}

		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	if c.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Index(c.url, "?") != -1 {
			c.url += "&" + paramBody
		} else {
			c.url = c.url + "?" + paramBody
		}
	} else if c.req.Method == "POST" && c.req.Body == nil {

	}
	if c.data != nil {
		c.Body(c.data)
	}

	c.req.Header.Add("Accept", "application/json")
	url, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	c.req.URL = url
	if Debug {
		println("url:" + string(c.url))
	}
	client := &http.Client{}
	resp, err := client.Do(c.req)
	c.hostpoolresp.Mark(err)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (c *context) GetStatus() (string, error) {
	resp, err := c.excute()
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}

// String returns the body string in response.
// it calls Response inner.
func (c *context) String() (string, error) {
	data, err := c.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (c *context) Bytes() ([]byte, error) {
	if c.body != nil {
		return c.body, nil
	}
	resp, err := c.excute()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.body = data
	return data, nil
}

func (c *context) GetResponse(v interface{}) error {
	body, err := c.Bytes()
	if err != nil {
		return err
	}
	if Debug {
		println("body:" + string(body))
	}
	err = json.Unmarshal(body, v)
	return err
}

func (c *context) GetBaseResponse() (*BaseResponse, error) {
	body, err := c.Bytes()
	if err != nil {
		return nil, err
	}
	if Debug {
		println("body:" + string(body))
	}
	baseResp := new(BaseResponse)
	jsonErr := json.Unmarshal(body, baseResp)
	if jsonErr != nil {
		return nil, jsonErr
	} else {
		return baseResp, nil
	}
}

func (c *context) GetSearch() (*SearchResult, error) {
	body, err := c.Bytes()
	if err != nil {
		return nil, err
	}

	search := new(SearchResult)
	jsonErr := json.Unmarshal(body, search)
	if jsonErr != nil {
		return nil, jsonErr
	} else {
		return search, nil
	}
}
