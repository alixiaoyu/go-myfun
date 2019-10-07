package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("---->", "SUCCESS")
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---->", string(valueGet))
	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])

}
