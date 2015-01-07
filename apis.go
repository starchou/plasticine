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

import "encoding/json"

//default put method
func (s *Server) Index(index string, _type string, id string, args map[string]interface{}, data interface{}) (*BaseResponse, error) {
	return s.index("PUT", index, _type, id, args, data)
}

func (s *Server) IndexOfPost(index string, _type string, args map[string]interface{}, data interface{}) (*BaseResponse, error) {
	return s.index("POST", index, _type, "", args, data)
}
func (s *Server) index(method, index string, _type string, id string, args map[string]interface{}, data interface{}) (*BaseResponse, error) {
	ctx := newContext(s.getUrl(getPath(index, _type, id), ""), method, s.hr)
	ctx.setParams(args)
	ctx.setData(data)
	return ctx.GetBaseResponse()
}

// GetSource retrieves the document by id and converts it to provided interface
func (s *Server) GetSource(index string, _type string, id string, args map[string]interface{}, source interface{}) error {
	ctx := Get(s.getUrl(getPath(index, _type, id, "_source"), ""), s.hr)
	ctx.setParams(args)
	body, err := ctx.Bytes()
	if err == nil {
		err = json.Unmarshal(body, &source)
	}
	return err
}

//
func (s *Server) Delete(index string, _type string, id string, args map[string]interface{}) (*BaseResponse, error) {
	ctx := Delete(s.getUrl(getPath(index, _type, id), ""), s.hr)
	ctx.setParams(args)
	return ctx.GetBaseResponse()
}
func (s *Server) UpdateDocument(index string, _type string, id string, args map[string]interface{}, data interface{}) (*BaseResponse, error) {
	ctx := newContext(s.getUrl(getPath(index, _type, id, "_update"), "pretty"), "POST", s.hr)
	ctx.setParams(args)
	d := make(map[string]interface{})
	d["doc"] = data
	ctx.setData(d)
	return ctx.GetBaseResponse()
}
func (s *Server) UpdateScript(index string, _type string, id string, args map[string]interface{}, data interface{}) (*BaseResponse, error) {
	ctx := newContext(s.getUrl(getPath(index, _type, id, "_update"), "pretty"), "POST", s.hr)
	ctx.setParams(args)
	d := make(map[string]interface{})
	d["script"] = data
	ctx.setData(d)
	return ctx.GetBaseResponse()
}
func (s *Server) DeleteByQuery(index string, _type string, id string, args map[string]interface{}) (*BaseResponse, error) {
	ctx := Delete(s.getUrl(getPath(index, _type, "_query"), "pretty"), s.hr)
	ctx.setParams(args)
	return ctx.GetBaseResponse()
}

func (s *Server) Search(index string, _type string, args map[string]interface{}, data interface{}) (*SearchResult, error) {
	ctx := Post(s.getUrl(getPath(index, _type, "_search"), "pretty"), s.hr)
	ctx.setParams(args)
	ctx.setData(data)
	return ctx.GetSearch()
}

func (s *Server) SearchByJson(index string, _type string, json Jsoner) (*SearchResult, error) {
	ctx := Post(s.getUrl(getPath(index, _type, "_search"), "pretty"), s.hr)
	ctx.setData(json.Encode())
	println(string(json.Encode()))
	return ctx.GetSearch()
}
