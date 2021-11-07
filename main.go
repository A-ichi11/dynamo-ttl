package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type User struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name,range"`
	Age    int    `dynamo:"Age"`
	Text   string `dynamo:"Text"`
	TTL    int64
}

// 本来はenvから取得した方が良い
const AWS_REGION = "ap-northeast-1"

// const DYNAMO_ENDPOINT = "http://localhost:8000"

// func getAWSSession() *session.Session {
// 	return session.Must(session.NewSession())
// }

func main() {
	// クライアントの設定
	// util.LoadEnv()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
		// DynamoDB-localを利用する場合は以下が必要

		// Endpoint:    aws.String(DYNAMO_ENDPOINT),
		// Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)

	// テーブル作成をする為に、一度テーブルを削除します
	// db.Table("UserTable").DeleteTable().Run()

	// テーブル作成
	// err = db.CreateTable("UserTable", User{}).Run()
	// if err != nil {
	// 	fmt.Println("err1")
	// 	panic(err)
	// }
	// テーブルの指定
	table := db.Table("UserTable")

	// User構造体をuser変数に定義
	var user User

	// DBにPutします
	// err = table.Put(&User{
	// 	UserID: "1234",
	// 	Name:   "太郎",
	// 	Age:    20,
	// }).Run()
	// if err != nil {
	// 	fmt.Println("err2")
	// 	panic(err)
	// }

	// DBからGetします

	err = table.Get("UserID", "1234").Range("Name", dynamo.Equal, "太郎").One(&user)
	if err != nil {
		fmt.Println("err3")
		panic(err)
	}
	fmt.Printf("GetDB%+v\n", user)

	// DBのデータをUpdateします
	// text := "新しいtextです"
	now := time.Now()
	setTTL := now.Add(1 * time.Hour).Unix()

	err = table.Update("UserID", "1234").
		Range("Name", "太郎").
		Set("TTL", setTTL).
		Value(&user)
	if err != nil {
		fmt.Println("err4")
		panic(err)
	}
	fmt.Printf("UpdateDB%+v\n", user)

	// // DBのデータをDeleteします
	// err = table.Delete("UserID", "1").Range("Name", "Test1").Run()
	// if err != nil {
	// 	panic(err)
	// }

	// // Delete出来ているか確認
	// err = table.Get("UserID", "1").Range("Name", dynamo.Equal, "Test1").One(&user)
	// if err != nil {
	// 	// Delete出来ていれば、dynamo: no item found のエラーとなる
	// 	fmt.Println("getError:", err)
	// }
}
