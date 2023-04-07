package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Dt struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	var st = Dt{}
	for true {
		data := map[string]interface{}{
			"water": rand.Intn(100),
			"wind":  rand.Intn(100),
		}
		reqJSON, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}
		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(body, &st)
		// fmt.Println(string(body))
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Println(st.Water, st.Wind)
		// ss, err := json.Marshal(st)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// sf := fmt.Sprintf(`
		// {
		// "water":%d,
		// "wind" :%d
		// }
		// `, st.Water, st.Wind)
		fmt.Printf("{\n")
		fmt.Printf(`  "water":%d`+"\n", st.Water)
		fmt.Printf(`  "wind" :%d`+"\n", st.Wind)
		fmt.Printf("}\n")
		if st.Water <= 5 {
			fmt.Println("status water : aman")
		} else if st.Water >= 6 && st.Water <= 8 {
			fmt.Println("status water : siaga")
		} else {
			fmt.Println("status water : bahaya")
		}
		if st.Wind <= 6 {
			fmt.Println("status Wind : aman")
		} else if st.Wind >= 7 && st.Wind <= 15 {
			fmt.Println("status Wind : siaga")
		} else {
			fmt.Println("status Wind : bahaya")
		}
		time.Sleep(15 * time.Second)
	}
}
