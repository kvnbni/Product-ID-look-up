package slackfunc

import(
	"net/http"
	product "kvnbni/Slack/productdetails"
	"kvnbni/Slack/redisfunc"
	"github.com/go-redis/redis"
	"encoding/json"
	"log"
	"io/ioutil"
	"bytes"
)

type slackReply struct{
	Ok bool `json:"ok"`
}

func PostMessage(messageDetails *product.ReqBody,productId string, redisHandler *redis.Client) bool {
	
	oauth_token,_:=redisfunc.RedisGet("token",redisHandler)
	token:="Bearer "+oauth_token //Slack authorization has format Bearer <token>
	
	//preparing the body of the http post 
	message:=make(map[string]string)
	message["text"]=productId
	message["channel"]=messageDetails.Event.Channel
	
	reply,err:=json.Marshal(&message)
	if err!=nil{
		log.Fatal(err)
	}
	
	//Initializing a http client
	httpClient:=&http.Client{}

	//http.NewRequest cause we need to customize the http header
	//NewBuffer returns *Buffer which executes both read and write methods making it of type io.Reader
	req, errPreparing:=http.NewRequest("POST","https://slack.com/api/chat.postMessage",bytes.NewBuffer(reply))
	if errPreparing!=nil{
		log.Fatal(errPreparing)
	}
	
	//Adding the custom headers for token and content type
	req.Header.Set("Authorization",token)
	req.Header.Set("Content-type","application/json")
	
	//Making the http request
	resp, errPosting:=httpClient.Do(req)
	if errPosting!=nil{
		log.Fatal(errPosting)
	}
	
	//Slack response for the message posted
	slackResponse, errReading := ioutil.ReadAll(resp.Body)
	if errReading!=nil{
		log.Fatal(errReading)
	}
	
	var sr slackReply
	err=json.Unmarshal(slackResponse, &sr)
	if err!=nil{
		log.Fatal(err)
	}
	if sr.Ok==true{
		return true
	}else{return false}
}

