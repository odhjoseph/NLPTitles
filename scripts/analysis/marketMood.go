package main 

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
	"io/ioutil"
)


func main() {
	var (
		marketTitles, marketSummary = createDictionary()
	)

	fmt.Println(marketTitles, marketSummary)

}

func createDictionary() (string, string) {
	var( 
		path = "/Users/josephodhiambo/Python/NLPTitles/scripts/jsonFeeds/"
		articles map[string]interface{}
		mTitles string
		mSummary string
	)

	files, err := ioutil.ReadDir(path)
	if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
		jsonFile, err := os.Open(path + f.Name())
		if err != nil {
			log.Println("This shouldn't be possible, unless empty directory", err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal([]byte(byteValue), &articles)
	}

	for k := range articles {
		mTitles += k
		if articles[k] != nil {
			mSummary += articles[k].(string)
		}
	}

	return mTitles, mSummary

}