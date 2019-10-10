package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CommentList struct {
	Comments []Comment `json:"comments"`
	Time     string    `json:"time"`
	Abc      string    `json:"abc"`
}

func main() {
	coms := CommentList{Time: "abc", Comments: []Comment{{Name: "def", Age: 12}, {Name: "dev", Age: 3}}}
	encode(coms)
	str := `{"comments":[{"name":"def","age":12},{"name":"dev","age":3}],"time":"abc","abc":"def"}`
	coms1 := &CommentList{}
	err := json.Unmarshal([]byte(str), coms1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(coms1.Abc)

}

func encode(coms CommentList) {
	s, err := json.Marshal(coms)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", s)
	fmt.Println()
}
