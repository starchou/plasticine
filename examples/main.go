package main

import (
	"fmt"
	"strconv"

	es "github.com/starchou/plasticine"
)

type User struct {
	Id   int
	Name string
}

func main() {
	s := es.NewServer("localhost", "9200")
	//s.Debug(true)
	user := &User{1, "star"}
	//Insert
	s.Index("plasticine", "user", strconv.Itoa(user.Id), nil, user)
	//Update
	user.Name = "chou"
	s.Index("plasticine", "user", strconv.Itoa(user.Id), nil, user)

	//Get
	res, _ := s.Get("plasticine", "user", "1", nil)
	fmt.Println(res.Found)

	//MGet
	// args := make(map[string]interface{})
	// args["_index"] = "plasticine"
	// args["_type"] = "user"
	// args["_id"] = "1"
	// args2 := make(map[string]interface{})
	// args2["_index"] = "plasticine"
	// args2["_type"] = "user"
	// args2["_id"] = "2"
	v := make(map[string]interface{})
	v["docs"] = es.Array(es.Doc().Index("plasticine").Type("user").Id("1"),
		es.Doc().Index("plasticine").Type("user").Id("2")) //es.Array(args, args2)
	result, _ := s.MultiGet("", "", "", v)
	fmt.Println(result)

	//Delete
	s.Delete("plasticine", "user", "1", nil)

	//Search DSL
	// es.Search().Index("plasticine").Type("News").From(1).Size(10).Query(
	// 	es.QueryString().Query("what is plasticine").Fields("Title", "Content"),
	// ).Filter(
	// 	es.And().And(
	// 		es.Range("Created").From("2012-12-10T15:00:00-08:00").To("2013-5-10T15:00:00-08:00"),
	// 		es.Terms().Terms("IsDelete", "0"),
	// 	),
	// ).Result(s)

	//Search DSL (facet)
	// es.Search().Query(es.MatchAll(nil)).Facets(
	// 	es.Facets().Terms("my_facet").Field("tag").Size(10).Order("term").Facet(),
	// ).Result(s)
}
