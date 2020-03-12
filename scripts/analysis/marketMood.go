package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"fmt"

	"gopkg.in/jdkato/prose.v2"
)

//MarketFile returns file with part of it cleaned for python analysis
type MarketFile struct {
	marketTitles  string
	marketSummary string
	cleanedTitles string
}

func main() {
	var (
		marketTitles, marketSummary = createDictionary()
	)

	titlesDoc, err := prose.NewDocument(marketTitles)
	if err != nil {
		log.Println("Failed to create titlesDoc", err)
	}

	summaryDocs, err := prose.NewDocument(marketSummary)
	if err != nil {
		log.Println("Failed to create summaryDocs", err)
	}

	fmt.Println(titlesDoc, summaryDocs)
	// for _, tok := range summaryDocs.Tokens() {
	// 	fmt.Println(tok.Text, tok.Tag, tok.Label)
	// }

}

func createDictionary() (string, string) {
	var (
		path     = "/Users/josephodhiambo/Python/NLPTitles/scripts/jsonFeeds/"
		articles map[string]interface{}
		mTitles  string
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
