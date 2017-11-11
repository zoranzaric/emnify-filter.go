package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
)

type event_type struct {
	Id int
}
type event struct {
	Id        string
	EventType event_type `json:"event_type"`
}

var addr = ":6666"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		// rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		decoder := json.NewDecoder(rdr1)
		var t event
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		if t.EventType.Id == 6 {
			_, err := http.Post("http://127.0.0.1:5000/", "application/json", bytes.NewBuffer(buf))
			
					if err != nil {
						panic(err)
					}
		}
	})

	http.ListenAndServe(addr, nil)
}
