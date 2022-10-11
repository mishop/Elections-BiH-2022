package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"time"
)

type Response []struct {
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	TotalVotes int     `json:"totalVotes"`
	Percentage float64 `json:"percentage"`
}
type JsonData struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	TotalVotes int     `json:"totalVotes"`
	Percentage float64 `json:"percentage"`
}
type Data []struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	TotalVotes int     `json:"totalVotes"`
	Percentage float64 `json:"percentage"`
}

func main() {
	start := time.Now()
	var newdata Data
	fmt.Println("Data:", newdata)
	// Get request
	for i := 1; i <= 2241; i++ {
		resp, err := http.Get("https://www.izbori.ba/api_2018/race6_pollingstationspartyresult/%22WebResult_2022GENT1_2022_4_20_14_10_43%22/" + strconv.Itoa(i) + "/3")
		if err != nil {
			fmt.Println("No response from request")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body) // response body is []byte
		//	fmt.Println(string(body))
		var result Response
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		//fmt.Printf("a: length: %d, capacity: %d, data: %v\n", len(result), cap(result), result)

		// fmt.Println(PrettyPrint(result))
		for j, _ := range result {
			p := JsonData{Id: i, Name: result[j].Name, TotalVotes: result[j].TotalVotes, Percentage: result[j].Percentage}
			newdata = append(newdata, p)
		}
		// Loop through the data node for the FirstName
	}
	fmt.Println(newdata)
	file, _ := json.MarshalIndent(newdata, "", " ")
	_ = ioutil.WriteFile("party.json", file, 0644)
	//fmt.Printf("%s\n", len(result))
	fmt.Println(time.Since(start))
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
