package rest_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallRestApi() {
	response, err := http.Get("https://randomuser.me/api/?results=2")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
