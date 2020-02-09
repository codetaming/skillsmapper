package pipeline

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type checkDetailsData struct {
	Email string
}

func Check(jsonStr []byte, url string) {
	data := checkDetailsData{}
	json.Unmarshal(jsonStr, &data)
	_, err := http.Get(url)
	if err != nil {
		print("error")
	}
	fmt.Println(data.Email)
}
