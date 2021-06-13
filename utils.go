package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/eidolon/wordwrap"
)

func md5hash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func printReview(r *review) {
	fmt.Printf("\033[1m%s\t\t%s\t\t%s\033[0m\n", time.Unix(int64(r.Date/1000), 0).Format("2006-01-02"), fmt.Sprintf("%d/5", r.Rating), r.User)
	fmt.Println("-------------------------------------------------------------------------------")
	wrapper := wordwrap.Wrapper(80, false)
	fmt.Printf("\033[1m%s\033[0m\n%s\n\n\n", wrapper(r.Title), wrapper(r.Text))
}

func save(file *reviewFile) {
	f, err := os.Create("./" + md5hash(file.URL) + ".json")
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.Encode(file)
}

func load(url string) *reviewFile {
	bytes, err := ioutil.ReadFile(md5hash(url) + ".json")
	if err != nil {
		panic(err)
	}

	file := &reviewFile{}
	if err := json.Unmarshal(bytes, file); err != nil {
		panic(err)
	}

	return file
}

func fileExists(url string) bool {
	if _, err := os.Stat(md5hash(url) + ".json"); os.IsNotExist(err) {
		return false
	}
	return true
}
