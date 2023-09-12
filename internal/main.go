package main

import (
	"encoding/json"
	"fmt"

	"github.com/darjun/ghtrending"
)

func main() {

	t := ghtrending.New(ghtrending.WithDaily())
	repos, err := t.FetchRepos()

	if err != nil {
		panic(err)
	}

	jsonRes, err := json.Marshal(repos)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonRes))

}
