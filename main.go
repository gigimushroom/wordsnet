package main

import (
	"time"
	"strings"
	"io/ioutil"
	"fmt"
	"log"
	"os/exec"
	"github.com/boltdb/bolt"
)

func callOxfordAPI(word string) string {
	if len(word) > 0 {
		out, err := exec.Command("./ox", word).Output()
		if err != nil {
			log.Println(err)
		}
		// Get rid of the first word
		return strings.TrimPrefix(string(out), word)
	}	
	return ""
}

func saveWordToDB(db *bolt.DB, k string, v string) {

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dic"))
		err := b.Put([]byte(k), []byte(v))
		return err
	})
/*
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dic"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
*/
}
func main() {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("words.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("dic"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// Read words from Dict.cn txt
	content, err := ioutil.ReadFile("wordbook.txt")
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	count := 0
	for _, words := range lines {
		w := strings.Fields(words)
		if len(w) > 4 {
			result := callOxfordAPI(w[0])
			saveWordToDB(db, w[0], w[1] + " " + w[4] + result)
			count++
			fmt.Println("Processed", count, "words", ". Current word", w)
			time.Sleep(time.Second * 1/2)
		}		
	}
	fmt.Println("Successfully stored total", count, "words")
	/*
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("dic"))
	
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	*/
}