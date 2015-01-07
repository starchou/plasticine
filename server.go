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
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/bitly/go-hostpool"
)

const (
	Version         = "0.1"
	DefaultProtocol = "http"
	DefaultDomain   = "localhost"
	DefaultPort     = "9200"
	// A decay duration of zero results in the default behaviour
	DefaultDecayDuration = 0
)

var (
	Debug = false
)

type Server struct {
	Hosts         []string
	Domain        string
	Port          string
	Protocol      string
	hp            hostpool.HostPool
	hr            hostpool.HostPoolResponse
	once          sync.Once
	DecayDuration time.Duration
}

func NewServer(host, port string) *Server {
	h := []string{host}
	return &Server{
		Hosts:         h,
		Port:          port,
		Domain:        DefaultDomain,
		Protocol:      DefaultProtocol,
		DecayDuration: time.Duration(DefaultDecayDuration * time.Second),
	}
}
func (s *Server) Debug(flag bool) {
	Debug = flag
}
func (s *Server) SetHosts(hosts []string) {
	s.Hosts = hosts
}
func (s *Server) initHostPool() {
	if len(s.Hosts) == 0 {
		s.Hosts = append(s.Hosts, fmt.Sprintf("%s:%s", s.Domain, s.Port))
	}
	s.hp = hostpool.NewEpsilonGreedy(s.Hosts, s.DecayDuration, &hostpool.LinearEpsilonValueCalculator{})
}
func (s *Server) getUrl(path, query string) string {
	host, port := s.getHostAndPort()
	if len(query) > 0 {
		return fmt.Sprintf("%s://%s:%s%s?%s", s.Protocol, host, port, path, query)
	} else {
		return fmt.Sprintf("%s://%s:%s%s", s.Protocol, host, port, path)
	}
}

func getPath(paths ...string) string {
	path := ""
	for _, p := range paths {
		if p == "" {
			continue
		}
		path += "/" + p
	}
	return path
}

func getQuery(query map[string]string) string {
	var param string
	if len(query) > 0 {
		var buf bytes.Buffer
		for k, v := range query {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
		param = buf.String()
		param = param[0 : len(param)-1]
	}
	return param
}
func (s *Server) getHostAndPort() (string, string) {
	s.once.Do(s.initHostPool)
	hr := s.hp.Get()
	s.hr = hr
	h := strings.Split(hr.Host(), ":")
	if len(h) == 2 {
		return h[0], h[1]
	}
	return h[0], s.Port
}
