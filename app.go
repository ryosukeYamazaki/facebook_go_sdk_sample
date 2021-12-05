package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	fb "github.com/huandu/facebook/v2"
)

func main() {
	permalink := "https://freee.workplace.com/groups/1779121192119476/permalink/1779125452119050/"
	postId := PostID(permalink)

	access_token := os.Getenv("ACCESS_TOKEN")
	res, _ := fb.Get(postId, fb.Params{
		"fields":       "message",
		"access_token": access_token,
	})
	fmt.Println("Here is my Facebook first name:", res["message"])
}

func PostID(workplaceURL string) string {
	parsedUrl, err := url.Parse(workplaceURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsedUrl.Path)
	splitedPath := strings.Split(parsedUrl.Path, "/")
	fmt.Println(splitedPath[4])
	// TODO error 処理
	return splitedPath[4]
}
