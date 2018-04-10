package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

type Dic struct {
	db    *bolt.DB
	words []string
}

var dic Dic

func (d *Dic) getRandomWords() string {
	fmt.Println("Service is getRandomWords.")
	var shuffle []string
	r := rand.New(rand.NewSource(time.Now().Unix()))
	count := 0
	for _, i := range r.Perm(len(d.words)) {
		val := d.words[i]
		shuffle = append(shuffle, val)
		count++
		if count > 0 {
			break
		}
	}

	fmt.Println(shuffle)
	var result []byte
	for _, v := range shuffle {
		d.db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("dic"))
			detail := b.Get([]byte(v))
			s := fmt.Sprintf("Word:%s \nDetails:\n%s\n", v, detail)
			result = append(result, s...)
			return nil
		})
	}
	return string(result)
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(dic.getRandomWords()))
}

func main() {

	fmt.Println("Service is up.")
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("words.db", 0666, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var words []string
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("dic"))

		b.ForEach(func(k, v []byte) error {
			//fmt.Printf("key=%s, value=%s\n", k, v)
			words = append(words, string(k))
			return nil
		})
		return nil
	})

	dic.db = db
	dic.words = words

	router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	router.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
