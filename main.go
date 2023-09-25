package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis"
)

type Subdomains struct {
	XMLName    xml.Name `xml:"subdomains",json:"subdomains"`
	Subdomains []string `xml:"subdomain",json:"subdomain"`
}

type Cookies struct {
	XMLName xml.Name `xml:"cookies",json:"cookies"`
	Cookies []Cookie `xml:"cookie",json:"cookie"`
}

type Cookie struct {
	XMLName xml.Name `xml:"cookie",json:"cookie"`
	Name    string   `xml:"name,attr",json:"name"`
	Host    string   `xml:"host,attr",json:"host"`
	Value   string   `xml:",chardata",json:"value"`
}

type Config struct {
	XMLName    xml.Name   `xml:"config",json:"config"`
	Subdomains Subdomains `xml:"subdomains",json:"subdomains"`
	Cookies    Cookies    `xml:"cookies",json:"cookies"`
}

func ReadXML(ctx context.Context, fileContents []byte) Config {
	var config Config
	xml.Unmarshal(fileContents, &config)
	for _, v := range config.Cookies.Cookies {
		fmt.Println(v.Name, v.Host, v.Value)
	}
	return config
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func AddToRedis(ctx context.Context, config Config) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func main() {
	ctx := context.TODO()
	xmlFile, err := os.Open("a.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	config := ReadXML(ctx, byteValue)
	AddToRedis(ctx, config)

}
