package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserInfo struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Location string `json: "location"`
}

type Result struct {
	English string `json: "english"`
	Maths   string `json: "maths"`
	Total   int    `json: "total"`
}

func login(userdata UserInfo) (Result, error) {

	fmt.Printf("%+v\n", userdata)

	var result Result
	// logic to check username and password is correct or not
	if userdata.Username == "admin" && userdata.Password == "admin" {

		result.English = "90"
		result.Maths = "80"
		result.Total = 170

		return result, nil

	}
	return result, fmt.Errorf("error: %s", "user is not admin")

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading body: ", err)
		return
	}
	fmt.Println("body data: ", string(body))

	// we want to store body data into userInfo
	var adminUser UserInfo
	err = json.Unmarshal(body, &adminUser)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	result, err := login(adminUser)
	if err != nil {

		fmt.Println("ERROR: ", err.Error())
		w.WriteHeader(400)
		fmt.Fprintln(w, err)
		return
	}

	// login successfull
	w.WriteHeader(200)
	// we want to send result in json format
	// marshalling
	resultJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Fprintln(w, string(resultJson))

}

func main() {
	fmt.Println("hello http-server")

	// registers
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
