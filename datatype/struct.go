package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("%s\n", data)
}

// ssh-keygen -t rsa -C '1349153326@qq.com' -f ~/.ssh/gitee_id_rsa
// ssh-keygen -t rsa -C 'lixiaoyu1760@gmail.com' -f ~/.ssh/github_id_rsa
