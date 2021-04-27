package main

import (
	"context"

	"github.com/google/go-github/v34/github"
)

func main() {

	client := github.NewClient(nil)
	client.Organizations.List(context.Background(), "willnorris", nil)
}
