package scheduler

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sikab-biz-test/application/user"
	"sikab-biz-test/domain"
)

var workers = make(chan bool, 10)

func Process() {
	jsonFile, err := os.Open("users_data.json")
	if err != nil {
		log.Fatal("Couldn't open json file")
	}
	defer jsonFile.Close()

	bodyByte, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Could not read json body")
	}

	var users []domain.User
	if err = json.Unmarshal(bodyByte, &users); err != nil {
		log.Fatalf("Could not unmarshal json body into struct. Err is %s", err)
	}

	for _, userRecord := range users {
		go func(userRecord domain.User) {
			workers <- true
			result := make(chan domain.User)
			if err != nil {
				log.Fatal("Could not marshal user record")
			}
			go user.SaveToDB(result)
			result <- userRecord
			<-workers
		}(userRecord)
	}
}
