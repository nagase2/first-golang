package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"

	log "github.com/sirupsen/logrus"
	//"gopkg.in/natefinch/lumberjack.v2"
)

type AuthConfig struct {
	Authentication []struct {
		ChannelId    string   `yaml:"channel_id"`
		SignalingKey string   `yaml:"signaling_key"`
		DeviceId     []string `yaml:"device_id"`
	}
}

// yamlファイルを読み込む
func readAuthYaml() AuthConfig {
	var configs AuthConfig

	source, err := ioutil.ReadFile("auth.yaml")
	if err != nil {
		fmt.Printf("failed reading config file: %v\n", err)
	}
	err = yaml.Unmarshal(source, &configs)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return configs
}

// Logger設定を行う
func setupLogger() {
	// TODO: 環境変数（環境）
	log.SetOutput(&lumberjack.Logger{
		Filename:   "log/app.log", // ファイル名
		MaxSize:    100,           // ローテーションするファイルサイズ(megabytes)
		MaxBackups: 100,           // 保持する古いログの最大ファイル数
		MaxAge:     30,            // 古いログを保持する日数
		LocalTime:  true,          // バックアップファイルの時刻フォーマットをサーバローカル時間指定
		Compress:   false,         // ローテーションされたファイルのgzip圧縮
	})
}
func str2time(t string) time.Time {
	// YYYY-MM-DDTHH:MM:SSZZZZの形式で渡される文字列tをtime.Time型に変換して返す
	parsedTime, _ := time.Parse("2006-01-02T15:04:05Z", t)
	return parsedTime
}

func logWriter() {
	//yamlを読み込むサンプル
	authConfig := readAuthYaml()
	fmt.Print(authConfig)
	dt := time.Now()
	fmt.Println(dt.Format("2006-01-02 15:04:05"))

	st := "2021-06-08T07:51:32.593704Z"
	fmt.Println(str2time(st))

	setupLogger()
	// ログファイル
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 最後にファイル閉じるよ
	// defer file.Close()

	// ログが出力された箇所をログに出力
	//	log.SetReportCaller(true)

	log.SetLevel(log.TraceLevel)
	log.Print("Logging in Go!!!!!")
	log.Trace("this is trace level of log")
	log.Debug("this is debug level of log")
	log.Info("this is info level of log")
	log.Warn("this is warn level of log")
	log.Error("this is error level of log")

	// log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{})

	log.WithFields(log.Fields{
		"record_ts":           "ｓsss",
		"timestamp":           "dddd",
		"channel_id":          "cccccc",
		"client_id":           "33333",
		"role":                "xxxx",
		"channel_connections": "dddd",
		"device_id":           "🈲",
	}).Info("A group of walrus emerges from the ocean")

	// log.WithFields(log.Fields{
	// 	"omg":    true,
	// 	"number": 100,
	// }).Fatal("The ice breaks!")

	//log.SetOutput(os.Stdout)
	log.Info("🐮🐮🐮🐮🐮🐮🐮🐮🐮🐮🐮🐮")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("🐩🐩🐩🐩🐩🐩🐩🐩")
}



// Controller example
type DeliveryService struct {
	Name string
	Num  int
}

// 受信
func (t *DeliveryService) Receive() string {
	fmt.Println("aaaccccc")
	return "received"
}

// 受信
func (t *DeliveryService) Send() string {
	fmt.Println("send")
	return "sent"
}
func (t *DeliveryService) addAndGet(headString string) int {
	t.Num = t.Num + 1
	fmt.Println(headString, t.Num)
	return t.Num
}

// これがインスタンスを生成するメソッドの役割をする
func NewThing(someParameter string, startNum int) *DeliveryService {
	p := new(DeliveryService)
	p.Name = someParameter
	p.Num = startNum
	return p
}

func main() {
	fmt.Println("Start!!!!")
	thing := NewThing("waw", 100)
	fmt.Println("🐔", thing.Name)
	thing.addAndGet("🐸")
	thing.addAndGet("🐸")

	thingPointer := *thing
	thingPointer.addAndGet("🐮")
	thingPointer.addAndGet("🐮")

	thing.addAndGet("🐸")
	thing.addAndGet("🐸")

	thing2 := NewThing("🍑waw", 1000)
	fmt.Println("🍑", thing2.Name)
	thing2.addAndGet("🍑")
	thing2.addAndGet("🍑")

	var x DeliveryService
	fmt.Println("🏠", x.Name)
	x.Receive()
	x.Send()
	x.addAndGet("🐵")
	x.addAndGet("🐵")
	x.addAndGet("🐵")

}
