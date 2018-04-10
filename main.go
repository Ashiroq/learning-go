package learninggo

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/doge", func(res http.ResponseWriter, req *http.Request) {
		// res.Header().Set("Content-Type", "image/jpeg")
		res.Write([]byte("<html><head></head><body>"))
		for i := 0; i < 15; i++ {
			doge, err := getDoge()
			if err == nil {
				res.Write([]byte("<img src=\"" + doge.Url + "\" alt=\"\" width=\"150\" />"))
				println(doge.Url)
			} else {
				println(err)
			}
		}
		res.Write([]byte("</body></html>"))
	})
	println("helloworld!")
	http.ListenAndServe("localhost:8000", nil)
}

type DogeData struct {
	Url string `json:"message"`
}

func getDoge() (DogeData, error) {
	res, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return DogeData{}, err
	}

	defer res.Body.Close()

	var data DogeData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return DogeData{}, err
	}
	// println(data.Url)
	return data, nil
}
