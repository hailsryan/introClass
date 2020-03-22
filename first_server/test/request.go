package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	rand.Seed(time.Now().Unix())
	res, err := http.Get("http://localhost:8080/wipe")
	if err != nil {
		panic("Failed to wipe" + err.Error())
	}
	if res.StatusCode != 200 {
		panic("Wrong status code" + strconv.Itoa(res.StatusCode))
	}

	reqs := rand.Intn(20)
	count := 0

	rl := rate.NewLimiter(1, 1)
	ctx, can := context.WithTimeout(context.Background(), 1*time.Second)
	defer can()
	for i := 0; i < reqs; i++ {
		rl.Wait(ctx)
		payload := rand.Intn(1000)
		count += payload
		fmt.Println("Sending: ", payload)
		res, err = http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer([]byte(strconv.Itoa(payload))))
		if err != nil {
			panic("Error posting to add" + err.Error())
		}
	}

	res, err = http.Get("http://localhost:8080/total")
	if err != nil {
		panic("Request to total failed" + err.Error())
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Could not read body" + err.Error())
	}
	defer res.Body.Close()
	bodyString := string(b)
	bodyInt, err := strconv.Atoi(bodyString)
	if err != nil {
		panic("Conversion to int from string failed" + err.Error())
	}
	if bodyInt != count {
		panic(fmt.Sprintf("Got %d expected %d from total", bodyInt, count))
	}
	fmt.Println(fmt.Sprintf("Success! Got %d expected %d from total", bodyInt, count))
}
