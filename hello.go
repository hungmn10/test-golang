package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")
	log :=fmt.Println



	log("================================")
	log("==================Test with time")
	log("================================")
	log("The time is", time.Now())
	log("The time is", time.Now())
	log("The time is", time.Now().Unix())
	log("The time is", time.Now().UnixNano())
	log("The time is", time.Now().Nanosecond())
	log("The time is", time.Now().Second())
	log("The time is", time.Now().Location())


	log("====================================")
	log("==================Test with timezone")
	log("====================================")
	americanLocation, americaErr := time.LoadLocation("America/New_York")
	vietnamLocation, vietnamErr := time.LoadLocation("Asia/Ho_Chi_Minh")
	if (americaErr == nil || vietnamErr==nil){
		log("System error", "LoadLocation")
	}
	timeAmerica := time.Date(2018,12,30,23,55,0,0, americanLocation)
	timeVietnam := time.Date(2018,12,30,23,55,0,0, vietnamLocation)

	log("The american time is",timeAmerica,timeAmerica.Unix())
	log("The vietnam time is",timeVietnam,timeVietnam.Unix())
}
