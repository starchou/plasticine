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

type CatApi struct {
	server *Server
	v      string
}

func (s *Server) Cat() *CatApi {
	return &CatApi{s, ""}
}
func (c *CatApi) V(v bool) {
	if v {
		c.v = c.v
	}
}
func (c *CatApi) Aliases() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "aliases"), c.v), c.server.hr).String()
}
func (c *CatApi) Allocation() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "allocation"), c.v), c.server.hr).String()
}

func (c *CatApi) Count(s ...string) (string, error) {
	if len(s) == 0 {
		return Get(c.server.getUrl(getPath("_cat", "count"), c.v), c.server.hr).String()
	} else {
		return Get(c.server.getUrl(getPath("_cat", "count", s[0]), c.v), c.server.hr).String()
	}
}

func (c *CatApi) Fielddata(q ...string) (string, error) {
	if len(q) == 0 {
		return Get(c.server.getUrl(getPath("_cat", "fielddata"), c.v), c.server.hr).String()
	} else {
		return Get(c.server.getUrl(getPath("_cat", "fielddata", strings.Join(q, ",")), c.v), c.server.hr).String()
	}
}

func (c *CatApi) Health() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "health"), c.v), c.server.hr).String()
}

func (c *CatApi) Indices(q ...string) (string, error) {
	if len(q) == 0 {
		return Get(c.server.getUrl(getPath("_cat", "indices"), c.v), c.server.hr).String()
	} else {
		return Get(c.server.getUrl(getPath("_cat", "indices", strings.Join(q, ",")), c.v), c.server.hr).String()
	}
}

func (c *CatApi) Master() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "master"), c.v), c.server.hr).String()
}

func (c *CatApi) Nodes() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "nodes"), c.v), c.server.hr).String()
}

func (c *CatApi) PendingTasks() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "pending_tasks"), c.v), c.server.hr).String()
}

func (c *CatApi) Plugins() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "plugins"), c.v), c.server.hr).String()
}

func (c *CatApi) Recovery() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "recovery"), c.v), c.server.hr).String()
}

func (c *CatApi) ThreadPool() (string, error) {
	return Get(c.server.getUrl(getPath("_cat", "thread_pool"), c.v), c.server.hr).String()
}

func (c *CatApi) Shards(q ...string) (string, error) {
	if len(q) == 0 {
		return Get(c.server.getUrl(getPath("_cat", "shards"), c.v), c.server.hr).String()
	} else {
		return Get(c.server.getUrl(getPath("_cat", "shards", strings.Join(q, ",")), c.v), c.server.hr).String()
	}
}
