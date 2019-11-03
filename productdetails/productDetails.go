package productdetails

import (
	"strings"
)

type ReqBody struct{
	Token string `json:"token"`
	Challenge string `json:"challenge"`
	Type string `json:"type"`
	Team_id string `json:"team_id"`
	Event event `json:"event"`
}

type event struct{
	Text string `json:"text"`
	User string `json:"user"`
	Channel string `json:channel`
}

func ProductName(text string) string{
	if strings.Contains(text,".amazon.com"){
		productNameRaw:=strings.Split(text,"|") //Splits the string into []string with two parts. One before the delimiter and one after.
		productName:= productNameRaw[1]
		productName=strings.TrimSuffix(productName,">")
		return productName
	}else{
		return "Invalid"
	}

}
