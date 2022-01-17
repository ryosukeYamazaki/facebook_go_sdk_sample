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
	// permalink := "https://freee.workplace.com/groups/1779121192119476/permalink/3594996737198570/"
	workplacePath := "/groups/1779121192119476/permalink/3594996737198570/"
	// postId := PostID(permalink)
	postIdFromPath := PostIDFrom(workplacePath)

	access_token := os.Getenv("ACCESS_TOKEN")
	app_id := os.Getenv("APP_ID")
	app_secrete := os.Getenv("APP_SECRET")
	globalApp := fb.New(app_id, app_secrete)
	globalApp.EnableAppsecretProof = true
	session := globalApp.Session(access_token)
	res, _ := session.Get(postIdFromPath, fb.Params{"fields": "message"})
	fmt.Println("Here is my Facebook first name:", res["message"])
}

func PostIDFrom(workplacePath string) string {
	splitedPath := strings.Split(workplacePath, "/")
	// TODO error 処理
	return splitedPath[4]
}

func PostID(workplaceURL string) string {
	parsedUrl, err := url.Parse(workplaceURL)
	if err != nil {
		log.Fatal(err)
	}
	splitedPath := strings.Split(parsedUrl.Path, "/")
	// TODO error 処理
	return splitedPath[4]
}
