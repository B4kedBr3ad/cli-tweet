package main

import(
	"fmt"
	"github.com/ChimeraCoder/anaconda"
        "flag"
)

func GetTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey("API Key")
    anaconda.SetConsumerSecret("API Secret")
    api := anaconda.NewTwitterApi("Access Token", "Access Token Secret")
    return api
}

func main(){
var text string
flag.StringVar(&text, "t", "", "tweet text")
flag.Parse()
api := GetTwitterApi()


tweet, err := api.PostTweet(text, nil)
if err != nil {
    panic(err)
    }
fmt.Print("tweeted:" + tweet.Text)

}
