package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// ※接続クライアントの生成は①、②いずれかの方法で行う

// // ① ~/.aws/credentialsの情報を使ってクライアントを生成（TODO：docker内からやるにはどうするか別途確認）
// func GetClient() (*sns.SNS, error) {
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))

// 	return sns.New(sess), nil
// }

// ② アクセスキーを明示的に指定して生成
func GetClient(acsId string, secId string, reg string) (*sns.SNS, error) {
	creds := credentials.NewStaticCredentials(acsId, secId, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(reg),
	})

	if err != nil {
		return nil, err
	}

	return sns.New(sess), nil
}

// 送信メッセージ、送信対象の電話番号を引数にとり、PublishInputのインスタンスを返却する
func CreateInputMessage(msg string, phoneNum string) *sns.PublishInput {
	pin := &sns.PublishInput{}
	pin.SetMessage(msg)
	pin.SetPhoneNumber(phoneNum)
	return pin
}

func main() {
	// まずは接続クライアントの生成
	client, err := GetClient(
		"AWS_ACCESS_KEY_IDを設定",        // 使用するAWSアクセスキーを設定
		"AWS_SECRET_ACCESS_KEY_IDを設定", // 使用するAWSシークレットキーを設定
		"REGIONを設定",                   // 使用するAWSのリージョンを設定
	)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	// ポイントは電話番号入力時に国コードを入力することです。今回は日本を想定して、[+81]を設定します。
	msgIn := CreateInputMessage(
		"TestMessage",    // 送信するメッセージを設定
		"+81XXXXXXXXXXX", // 送信先の電話番号を設定
	)

	result, err := client.Publish(msgIn)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("Result: %s", result.String())
}
