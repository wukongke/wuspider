package libs

import (
	"work-codes/wuspider/wuyue/models"

	"gopkg.in/mgo.v2/bson"
)

func AccountFindOne(filter interface{}) (*models.Account, error) {
	var row models.Account
	err := models.AccountVO.Find(filter).One(&row)
	return &row, err
}

func AccountSave(data bson.M) error {
	// var account models.Book
	// err := models.AccountVO.Find(bson.M{"name": data["name"]}).One(&account)
	// if err != nil {
	// 	err = models.AccountVO.Insert(data)
	// } else {
	// 	_, err = models.AccountVO.UpdateAll(bson.M{"name": data["name"]}, bson.M{"$set": data})
	// }
	_, err := models.AccountVO.Upsert(bson.M{"name": data["name"]}, bson.M{"$set": data})
	return err
}

func BookSave(data bson.M) error {
	_, err := models.BookVO.Upsert(bson.M{"bookNo": data["bookNo"]}, bson.M{"$set": data})
	return err
}

func ChapterSave(data bson.M) error {
	_, err := models.ChapterVO.Upsert(bson.M{"chapterNo": data["chapterNo"]}, bson.M{"$set": data})
	return err
}
