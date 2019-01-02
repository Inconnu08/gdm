package gdm

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func Download(DestinationFolder string) {
	var wg sync.WaitGroup
	wg.Add(len(ImageUrls))

	for i, url := range ImageUrls {
		go func(i int, url string) {
			defer wg.Done()
			fmt.Printf("\ndownloading: %s\n", url)

			response, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			file, err := os.Create(DestinationFolder + strconv.Itoa(i+1) + ".jpg")
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}

			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
			if err := response.Body.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Success!")
		}(i, url)
	}
	wg.Wait()
}

func Elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Function took %v\n", time.Since(start))
	}
}
