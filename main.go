package main 

import (
	"net/http"
	"fmt"
	//"io/ioutil"
	"encoding/json"
	"encoding/csv"
	"os"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting server")
}

func main() {
	http.HandleFunc("/record", writeToFile)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

type confusionData struct {
	URL string `json: "url"`
	Time string `json: "time"`
}

func writeToFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in writeToFile")
	switch r.Method {
	case "POST":
		fmt.Println("POST method")

		var data confusionData
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&data)

		fmt.Println(data)

		records, err := os.OpenFile("records.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		check(err)
		defer records.Close()

		writer := csv.NewWriter(records)

		var csv_data []string
		csv_data = append(csv_data, data.URL)
		csv_data = append(csv_data, data.Time)

		fmt.Println("csv before write: ", csv_data)

		write_err := writer.Write(csv_data)
		check(write_err)
		writer.Flush()

		fmt.Println("success writing to file!")

	default:
		fmt.Println("unsupported Method")
	

	}
}

