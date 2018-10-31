package onlinediscover

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

var (
	Host string
)

func HttpPOST(url string, parm string) string{
	client := &http.Client{}
	url = "http://" + Host + url
	request, err := http.NewRequest("POST", url , strings.NewReader(parm))
	if err != nil {
		fmt.Print("error occur during create post request")
		return ""
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Print("error occur during post request")
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("error occur during reading body")
		return ""
	}
	return string(body)
}

