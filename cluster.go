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
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

type ClusterHealth struct {
	ClusterName         string `json:"cluster_name"`
	Status              string `json:"status"`
	TimedOut            bool   `json:"timed_out"`
	NumberOfNodes       int    `json:"number_of_nodes"`
	NumberOfDataNodes   int    `json:"number_of_data_nodes"`
	ActivePrimaryShards int    `json:"active_primary_shards"`
	ActiveShards        int    `json:"active_shards"`
	RelocatingShards    int    `json:"relocating_shards"`
	InitializingShards  int    `json:"initializing_shards"`
	UnassignedShards    int    `json:"unassigned_shards"`
}

func (s *Server) GetClusterHealth(indices ...string) (ClusterHealth, error) {
	var ch ClusterHealth
	var path = ""
	if len(indices) == 0 {
		path = getPath("_cluster", "health")
	} else {
		path = getPath("_cluster", "health", strings.Join(indices, ","))
	}
	err := Get(s.getUrl(path, ""), s.hr).GetResponse(&ch)
	return ch, err
}

type ClusterState struct {
	ClusterName  string                      `json:"cluster_name"`
	MasterNode   string                      `json:"master_node"`
	Version      string                      `json:"version"`
	Nodes        map[string]ClusterStateNode `json:"nodes"`
	Metadata     simplejson.Json             `json:"metadata"`
	RoutingTable simplejson.Json             `json:"routing_table"` // TODO: Routing Table
	RoutingNodes simplejson.Json             `json:"routing_nodes"` // TODO: Routing Nodes
	Allocations  simplejson.Json             `json:"allocations"`   // TODO: Allocations
}

type ClusterStateNode struct {
	Name             string          `json:"name"`
	TransportAddress string          `json:"transport_address"`
	Attributes       simplejson.Json `json:"attributes"` // TODO: Attributes
}

func (s *Server) GetClusterState() (ClusterState, error) {
	var cs ClusterState
	var path = getPath("_cluster", "state")
	err := Get(s.getUrl(path, ""), s.hr).GetResponse(&cs)
	return cs, err
}

type ClusterStats struct {
	ClusterName string          `json:"cluster_name"`
	Timestamp   int64           `json:"timestamp"`
	Status      string          `json:"status"`
	Indices     simplejson.Json `json:"indices"`
	Nodes       simplejson.Json `json:"nodes"`
}

func (s *Server) GetClusterStats() (ClusterStats, error) {
	var cs ClusterStats
	var path = getPath("_cluster", "stats")
	err := Get(s.getUrl(path, ""), s.hr).GetResponse(&cs)
	return cs, err
}

type Task struct {
	InsertOrder       int    `json:"insert_order"`
	Priority          string `json:"priority"`
	Source            string `json:"source"`
	TimeInQueueMillis int    `json:"time_in_queue_millis"`
	TimeInQueue       string `json:"time_in_queue"`
}
type PendingTasks struct {
	Tasks []Task `json:"tasks"`
}

func (s *Server) GetClusterPendingTasks() (PendingTasks, error) {
	var pt PendingTasks
	var path = getPath("_cluster", "pending_tasks")
	err := Get(s.getUrl(path, ""), s.hr).GetResponse(&pt)
	return pt, err
}

//usage
//  s := plasticine.NewServer("localhost", "9200")
//	str, err := s.ClusterReroute(true,
//		plasticine.Commands(
//			plasticine.Command(plasticine.Move().Index("test1").Shard(1).FromNode("node1").ToNode("node2")),
//			plasticine.Command(plasticine.Allocate().Index("test2").Shard(0).Node("ad").AllowPrimary(false)),
//		))
//	fmt.Println(str, err)
func (s *Server) ClusterReroute(dry_run bool, v interface{}) (string, error) {
	var url = ""
	if dry_run {
		url = s.getUrl(getPath("_cluster", "reroute"), "dry_run")
	} else {
		url = s.getUrl(getPath("_cluster", "reroute"), "")
	}
	return Post(url, s.hr).setData(v).String()
}

type ClusterSettings struct {
	Transient  map[string]int `json:"transient"`
	Persistent map[string]int `json:"persistent"`
}

func (s *Server) UpdateClusterSettings(transient, persistent map[string]int) (ClusterSettings, error) {
	var cs ClusterSettings
	var j = simplejson.New()
	if transient != nil {
		j.Set("transient", transient)
	}
	if persistent != nil {
		j.Set("persistent", persistent)
	}
	data, err := j.Encode()
	if err != nil {
		return cs, err
	}
	var url = s.getUrl(getPath("_cluster", "settings"), "")
	err = Put(url, s.hr).setData(data).GetResponse(&cs)
	return cs, err
}

type NodeStats struct {
	ClusterName string          `json:"cluster_name"`
	Nodes       simplejson.Json `json:"nodes"`
}

func (s *Server) GetNodeStats(para []string, nodes ...string) (NodeStats, error) {
	var ns NodeStats
	var path = ""
	if len(nodes) == 0 {
		path = getPath("_nodes", "stats")
	} else {
		path = getPath("_nodes", strings.Join(nodes, ","), "stats", strings.Join(para, ","))
	}
	err := Get(s.getUrl(path, ""), s.hr).GetResponse(&ns)
	return ns, err
}

func (s *Server) GetNodeInfo(para []string, nodes ...string) (NodeStats, error) {
	var ns NodeStats
	var path = ""
	if len(nodes) == 0 && len(para) == 0 {
		path = getPath("_nodes", "_all", "_all")
	} else {
		path = getPath("_nodes", strings.Join(nodes, ","), strings.Join(para, ","))
	}
	err := Get(s.getUrl(path, ""), s.hr).GetResponse(&ns)
	return ns, err
}

func (s *Server) ShutDownNode(delay int, nodes ...string) (string, error) {
	var path = ""
	if len(nodes) == 0 {
		path = getPath("_cluster", "nodes", "_all", "_shutdown")
	} else {
		path = getPath("_cluster", "_nodes", strings.Join(nodes, ","), "_shutdown")
	}
	return Post(s.getUrl(path, getQuery(map[string]string{"delay": strconv.Itoa(delay)})), s.hr).String()
}
func Command(v ...interface{}) *simplejson.Json {
	j := simplejson.New()
	if len(v) > 0 {
		for i := 0; i < len(v); i++ {
			j.Set("command", v[0])
		}
	}
	return j
}

func Commands(v ...*simplejson.Json) []byte {
	length := len(v)
	j := simplejson.New()
	tmp := []interface{}{}
	for i := 0; i < length; i++ {
		tmp = append(tmp, v[i].Get("command"))
	}
	j.Set("commands", tmp)
	b, err := j.Encode()
	if err != nil {
		return nil
	}
	return b
}

type MoveCommand struct {
	Json *simplejson.Json `json:"move"`
}

func Move() *MoveCommand {
	return &MoveCommand{simplejson.New()}
}
func (c *MoveCommand) Index(val string) *MoveCommand {
	c.Json.Set("index", val)
	return c
}
func (c *MoveCommand) Shard(val int) *MoveCommand {
	c.Json.Set("shard", val)
	return c
}
func (c *MoveCommand) FromNode(val string) *MoveCommand {
	c.Json.Set("from_node", val)
	return c
}
func (c *MoveCommand) ToNode(val string) *MoveCommand {
	c.Json.Set("to_node", val)
	return c
}

type CancelCommand struct {
	Json *simplejson.Json `json:"cancel"`
}

func Cancel() *CancelCommand {
	return &CancelCommand{simplejson.New()}
}
func (c *CancelCommand) Index(val string) *CancelCommand {
	c.Json.Set("index", val)
	return c
}
func (c *CancelCommand) Shard(val int) *CancelCommand {
	c.Json.Set("shard", val)
	return c
}
func (c *CancelCommand) Node(val string) *CancelCommand {
	c.Json.Set("node", val)
	return c
}
func (c *CancelCommand) AllowPrimary(val bool) *CancelCommand {
	c.Json.Set("allow_primary", val)
	return c
}

type AllocateCommand struct {
	Json *simplejson.Json `json:"allocate"`
}

func Allocate() *AllocateCommand {
	return &AllocateCommand{simplejson.New()}
}
func (c *AllocateCommand) Index(val string) *AllocateCommand {
	c.Json.Set("index", val)
	return c
}
func (c *AllocateCommand) Shard(val int) *AllocateCommand {
	c.Json.Set("shard", val)
	return c
}
func (c *AllocateCommand) Node(val string) *AllocateCommand {
	c.Json.Set("node", val)
	return c
}
func (c *AllocateCommand) AllowPrimary(val bool) *AllocateCommand {
	c.Json.Set("allow_primary", val)
	return c
}
