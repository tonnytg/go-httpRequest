package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// FetchHost get argv return to chek url then return nil or err
func FetchHost(valor *[]string) error {
	for _, url := range os.Args[1:] {
		fmt.Printf("Url Fetch: %v\n", valor)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return fmt.Errorf("%s", err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			return fmt.Errorf("%s", err)
		}
		fmt.Printf("%s", b)
	}
	return nil
}

func main() {
	valor := os.Args[1:]
	fmt.Printf("%v", valor)
	err := FetchHost(&valor)
	if err != nil {
		fmt.Println("erro")
	}
}
