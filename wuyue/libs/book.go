package libs

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/PuerkitoBio/goquery"

	"work-codes/wuspider/wuyue/common"
	"work-codes/wuspider/wuyue/config"
)

func GetBook(bookNo string) {
	bookUrl := config.BaseUrl + "/book/" + bookNo
	doc := common.Curl(bookUrl)
	name := doc.Find("#info h1").First().Text()
	fmt.Println(name)
	intro := ""
	doc.Find("#intro p").Each(func(i int, s *goquery.Selection) {
		intro += s.Text() + "\n"
	})
	author := doc.Find("#info p").First().Text()
	authorArr := strings.Split(author, "：")
	fmt.Println("author: ", authorArr[1])

	serialize := strings.Split(strings.Split(doc.Find("#info p").Eq(1).Text(), "：")[1], ",")[0]
	fmt.Println("serialize: ", serialize)
	serializeStatus := 0
	if serialize == "完结" {
		serializeStatus = 1
	}

	lastUpdatedAt := strings.Split(doc.Find("#info p").Eq(2).Text(), "：")[1]
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", lastUpdatedAt)
	fmt.Println("最后更新时间: ", updatedAt.Unix())

	lastChapter := strings.Split(doc.Find("#info p").Eq(3).Text(), "：")[1]
	fmt.Println("最新章节: ", lastChapter)

	image, _ := doc.Find("#sidebar #fmimg img").Attr("src")
	fmt.Println("image: ", image)

	recommendBooks := []string{}
	doc.Find("#listtj a").Each(func(i int, s *goquery.Selection) {
		recommendBooks = append(recommendBooks, s.Text())
	})
	fmt.Println("推荐阅读：", recommendBooks)

	// 字符串转rune
	bookType := doc.Find(".con_top a").Eq(1).Text()
	fmt.Println("类型：", bookType[0:6])

	// 插入作者数据
	accountParams := bson.M{
		"_id":       bson.NewObjectId(),
		"name":      authorArr[1],
		"trueName":  authorArr[1],
		"isAuther":  1,
		"status":    1,
		"createdAt": time.Now().Unix(),
		"updatedAt": time.Now().Unix(),
	}
	err := AccountSave(accountParams)
	if err != nil {
		fmt.Println("AccountSave: ", err)
	}
	account, err := AccountFindOne(bson.M{"name": authorArr[1]})

	bookParams := bson.M{
		"_id":         bson.NewObjectId(),
		"bookNo":      bookNo,
		"name":        name,
		"accountId":   (*account).Id,
		"type":        bookType[0:6],
		"intro":       intro,
		"image":       image,
		"rate":        5.0,
		"wordCount":   100.5,
		"serialize":   serializeStatus,
		"loveCount":   1000,
		"likes":       recommendBooks,
		"url":         bookUrl,
		"lastChapter": lastChapter,
		"status":      1,
		"createdAt":   time.Now().Unix(),
		"updatedAt":   updatedAt.Unix(),
	}
	err = BookSave(bookParams)
	if err != nil {
		fmt.Println("BookSave: ", err)
	}

	// 下载章节列表
	chapList := []string{}
	doc.Find("#list dd").Each(func(i int, s *goquery.Selection) {
		chapList = append(chapList, s.Find("a").Text())
		url, _ := s.Find("a").Attr("href")
		GetChapter(url)
	})
	// fmt.Println("章节列表： ", chapList)
}
