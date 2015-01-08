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

import "strings"

type IndicesApi struct {
	server *Server
}

func (s *Server) Indices() *IndicesApi {
	return &IndicesApi{s}
}

//delete
func (s *IndicesApi) Delete(index ...string) (string, error) {
	var path = getPath(strings.Join(index, ","))
	return Delete(s.server.getUrl(path, ""), s.server.hr).String()
}

//delete
func (s *IndicesApi) Exists(index, _type string) (bool, error) {
	var path = getPath(index, _type)
	status, err := Head(s.server.getUrl(path, ""), s.server.hr).GetStatus()
	if status == "200" {
		return true, err
	} else {
		return false, err
	}
}

//status
func (s *IndicesApi) Status(index ...string) (string, error) {
	var path = getPath(strings.Join(index, ","), "_status")
	return Post(s.server.getUrl(path, ""), s.server.hr).String()
}

//clean Cache
func (s *IndicesApi) CleanCache(index ...string) (string, error) {
	var path = getPath(strings.Join(index, ","), "_cache", "clear")
	return Post(s.server.getUrl(path, ""), s.server.hr).String()
}

//flush
func (s *IndicesApi) Flush(index ...string) (string, error) {
	var path = getPath(strings.Join(index, ","), "_flush")
	return Post(s.server.getUrl(path, ""), s.server.hr).String()
}

//refresh
func (s *IndicesApi) Refresh(index ...string) (string, error) {
	var path = getPath(strings.Join(index, ","), "_refresh")
	return Post(s.server.getUrl(path, ""), s.server.hr).String()
}

//optimize
func (s *IndicesApi) Optimize(index ...string) (string, error) {
	var path = getPath(strings.Join(index, ","), "_optimize")
	return Post(s.server.getUrl(path, ""), s.server.hr).String()
}
