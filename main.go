package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type WeatherStruct struct {
	ID   int64
	Main string
	Desc string
	Icon string
}

type JSONResponse struct {
	Coord   interface{}
	Weather []WeatherStruct
}

func getEmojiLookup(code string) (string, error) {
	file, _ := ioutil.ReadFile("weather-moji.json")

	var result map[string]interface{}

	if err := json.Unmarshal([]byte(file), &result); err != nil {
		panic(err)
	}

	value, ok := result[code]

	if !ok {
		return "", errors.New("code not found")
	}

	return value.(string), nil
	//If key exists for code, return that, else return nil
}

//http://venkat.io/posts/twitter-api-auth-golang
func updateTwitter(message string) {
	ConsumerKey := os.Getenv("ConsumerKey")
	ConsumerSecret := os.Getenv("ConsumerSecret")
	AccessToken := os.Getenv("AccessToken")
	AccessTokenSecret := os.Getenv("AccessTokenSecret")

	consumer := oauth.NewConsumer(ConsumerKey,
		ConsumerSecret,
		oauth.ServiceProvider{})

	//NOTE: remove this line or turn off Debug if you don't
	//want to see what the headers look like
	consumer.Debug(true)

	t := oauth.AccessToken{
		Token:  AccessToken,
		Secret: AccessTokenSecret,
	}

	client, err := consumer.MakeHttpClient(&t)

	if err != nil {
		log.Fatal(err)
	}

	/* To get user info */
	//userData, err := client.Get("https://api.twitter.com/1.1/account/verify_credentials.json")
	//
	//if err != nil {
	//	log.Fatal(err, userData)
	//}
	//
	//respBody, err := ioutil.ReadAll(userData.Body)
	//fmt.Println(string(respBody))

	response, err := client.PostForm(
		"https://api.twitter.com/1.1/account/update_profile.json",
		url.Values{"name": []string{message}})

	if err != nil {
		log.Fatal(err, response)
	}

	_ = response.Body.Close()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env-example file")
	}

	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", os.Getenv("City"), os.Getenv("AppId")))

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result JSONResponse

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return
	}

	rawCode := result.Weather[0].ID

	code := strconv.FormatInt(rawCode, 10) // Convert to string, base 10 for decimal integers

	emoji, emojiErr := getEmojiLookup(code)

	if emojiErr != nil {
		log.Fatal(emojiErr)
	}

	updateTwitter(fmt.Sprintf("%s %s", os.Getenv("Prefix"), emoji))
}
