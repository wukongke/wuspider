package main

import (
	"fmt"
	"time"
	"work-codes/wuspider/wuyue/common"
	"work-codes/wuspider/wuyue/config"
	"work-codes/wuspider/wuyue/libs"
	"work-codes/wuspider/wuyue/models"

	"gopkg.in/mgo.v2/bson"
)

func init() {
	types := config.Types
	for _, v := range types {
		filter := map[string]interface{}{
			"typeNo": v["typeNo"],
		}
		var row models.Type
		err := models.TypeVO.Find(filter).One(&row)
		if err != nil {
			v["_id"] = bson.NewObjectId()
			v["status"] = 1
			v["createdAt"] = time.Now().Unix()
			v["updatedAt"] = time.Now().Unix()
			_ = models.TypeVO.Insert(v)
		}
	}
}

func main() {
	defer common.MgoClose()
	now := time.Now()
	for _, v := range config.Books {
		libs.GetBook(v)
	}
	duration := time.Since(now)
	fmt.Println("耗时:", duration)
}
