package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Word struct {
	Word   string `json:"Ord"`
	Gender string `json:"Genus"`
}

func main() {
	words, length := getWords()
	words = words[:length]

	wordsJson, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./words.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(wordsJson)
	jsonFile.Close()
	fmt.Println("JSON data was written to " + jsonFile.Name())
}

func getWords() ([]Word, int) {
	resp, err := http.Get("https://svn.spraakdata.gu.se/sb/fnplusplus/pub/dalin_saldo.html?fbclid=IwAR1UC2j2j8vNEftUdIK_fwcsqiiKzDeF9SwiwP8sigufUEHenA7zSDPvhdw")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	pageContent := string(html)

	htmlWords := strings.Split(pageContent, "<tr>")[2:]
	words := make([]Word, len(htmlWords)-2)

	inputIndex := 0
	for _, html := range htmlWords {
		if isNoun(html) {
			fmt.Println(html)
			word := Word{Word: extractModernSpelling(html), Gender: extractGender(html)}
			words[inputIndex] = word
			inputIndex++
		}
	}

	return words, inputIndex
}

func extractOldSpelling(html string) string {
	return strings.Split(strings.Split(html, `http://litteraturbanken.se/query/dalin.xql?word=`)[1], `&limit=10">`)[0]
}

func extractModernSpelling(html string) string {
	return strings.Split(strings.Split(html, extractOldSpelling(html)+"</a></td><td>")[1], "</td>")[0]
}

func isNoun(html string) bool {
	return len(strings.Split(html, extractModernSpelling(html)+"</td><td>nn")) > 1
}

func extractGender(html string) string {
	if htmlSplit := strings.Split(html, `http://spraakbanken.gu.se/ws/dalin-ws/gen/html/nn_`); len(htmlSplit) > 1 {
		return strings.Trim(strings.Split(htmlSplit[1], "_")[0], "0123456789")
	} else {
		return extractGenderSpecialCase(html)
	}
}

func extractGenderSpecialCase(html string) string {
	if htmlSplit := strings.Split(html, "<td>"+extractModernSpelling(html)+"</td><td>nn</td><td>-</td><td>"); len(htmlSplit) > 1 {
		if gender := strings.Split(htmlSplit[1], "</td>")[0][:1]; gender == "f" || gender == "m" || gender == "n" {
			return gender
		}
	}
	return "okänt"
}
