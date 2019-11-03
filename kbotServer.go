package main

import(
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	product "kvnbni/Slack/productdetails"
	"github.com/go-redis/redis"
	"kvnbni/Slack/redisfunc"
	"kvnbni/Slack/slackfunc"
)


//This function handles the incoming requests from Slack 
func slackServer(w http.ResponseWriter, r *http.Request){
	requestBody, err:= ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Printf("Encountered %v while reading the http response body",err)
		http.Error(w, "Couldn't read the response body", http.StatusBadRequest)
	}
	
	var request product.ReqBody
	err=json.Unmarshal(requestBody, &request)
	if err!=nil{
		log.Fatal(err)
	}
	
	if request.Type=="url_verification"{
		handleChallenge(request.Challenge,w)
	}else if request.Type=="event_callback"{
		w.WriteHeader(http.StatusOK) //Responding to Slack with status OK 200
		handleEventCallBack(&request)
	}
}

//This function handles the challenge forwarded by Slack
func handleChallenge(challenge string, w http.ResponseWriter){
	fmt.Fprintf(w,challenge) 
}

//This function handles the product details extraction
func handleEventCallBack(request *product.ReqBody){
	textReceived:=request.Event.Text
	
	productName:=product.ProductName(textReceived) //Custom func to return clean text
	
	//Redis Client initiation
	redisHandler:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//Getting product id
	productId,redisReadError:=redisfunc.RedisGet(productName,redisHandler)
	if redisReadError!=nil{
		log.Fatal(redisReadError)
	}

	//Posting message back to Slack
	slackfunc.PostMessage(request, productId, redisHandler) 
}

func main(){
	http.HandleFunc("/slack/",slackServer) 
	log.Fatal(http.ListenAndServe(":8000",nil))

}
