package libs

import (
	"fmt"
	"strings"
	"time"
	"work-codes/wuspider/wuyue/common"
	"work-codes/wuspider/wuyue/config"

	"gopkg.in/mgo.v2/bson"
)

func GetChapter(url string) {
	chapUrl := config.BaseUrl + url
	doc := common.Curl(chapUrl)
	chapTitle := doc.Find(".bookname h1").Text()
	fmt.Println("chapTitle: ", chapTitle)

	// bookNoString, _ := doc.Find(".footer_cont a").Attr("href")
	// bookNo := strings.Split(bookNoString, "/")
	// // bookName := doc.Find(".footer_cont a").Text()
	// fmt.Println("bookNo: ", bookNo[2])

	chapterNo := strings.Split(url, "/")
	chapterArr := strings.Split(chapterNo[3], ".")

	contentHtml, _ := doc.Find("#content").Html()
	content := ""
	contentArr := strings.Split(contentHtml, "<br/><br/>")
	for _, v := range contentArr {
		content += v + "\n"
	}
	chapData := bson.M{
		"_id":       bson.NewObjectId(),
		"chapterNo": chapterArr[0],
		"title":     chapTitle,
		"bookNo":    chapterNo[2],
		"content":   content,
		"url":       chapUrl,
		"status":    1,
		"createdAt": time.Now().Unix(),
		"updatedAt": time.Now().Unix(),
	}
	err := ChapterSave(chapData)
	if err != nil {
		fmt.Println("ChapterSave err: ", err)
	}
}
