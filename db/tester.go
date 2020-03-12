package  

func main() {
	var paths = []string{"/Users/josephodhiambo/Python/NLPTitles/scripts/jsonFeeds/",
	"/Users/josephodhiambo/Python/NLPTitles/scripts/titleDisplay/"}
	
	var articles map[string]interface{}

	jsonFile, err := os.Open(paths[0] + "search2020-03-05.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &articles)

	for k := range articles {
		mTitles += k
		if articles[k] != nil {
			mSummary += articles[k].(string)
		}
	}


}