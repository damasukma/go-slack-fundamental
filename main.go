package main 

import(
	"fmt"
	"github.com/nlopes/slack"
	"log"
	_"flag"
	"os"
	"os/signal"
	_"syscall"
	"time"
	_"reflect"
	"encoding/json"
	"slack-info/config"
)



func main(){
	defer catch()


	envConfig, err := config.Setup()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("TOKEN[%s] CHANNEL[%s]", envConfig.Token, envConfig.Channel)

	graceSignal := make(chan os.Signal, 1)

	slackApi := slack.New(
		envConfig.Token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)
	


	CheckSignal(graceSignal)
	


	rtm := slackApi.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		log.Println("Event Received:")

		switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
			//
			case *slack.ConnectedEvent:
				log.Println("Info:", ev.Info)

				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", envConfig.Channel))


			case *slack.MessageEvent:
				data := make(map[string]interface{})

				res, _ := json.Marshal(ev.Msg)
				json.Unmarshal([]byte(res), &data)
				
				if data["text"] != "test"{
					log.Println("Cannot Found Keyword %s", data["text"])
					break
				}

				rtm.SendMessage(rtm.NewOutgoingMessage("test 123", envConfig.Channel))


			case *slack.RTMError:
				log.Fatal("Error $s \n", ev.Error())	

			default:
					
		}
	}


	

}


func CheckSignal(grace chan os.Signal){
	signal.Notify(grace, os.Interrupt)
	go func(){
		defer close(grace)
		
		<-grace
		fmt.Println("Stop")
		fmt.Println("Wait ...")
		time.Sleep(time.Second * 3)
		fmt.Println("Finish")
		os.Exit(0)


	}()
}

func catch(){
	if r := recover(); r != nil{
		log.Fatal(r)
	}
}

