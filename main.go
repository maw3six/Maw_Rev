package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

const banner = `

███╗   ███╗ █████╗ ██╗    ██╗██████╗ ███████╗██╗   ██╗
████╗ ████║██╔══██╗██║    ██║██╔══██╗██╔════╝██║   ██║
██╔████╔██║███████║██║ █╗ ██║██████╔╝█████╗  ██║   ██║
██║╚██╔╝██║██╔══██║██║███╗██║██╔══██╗██╔══╝  ╚██╗ ██╔╝
██║ ╚═╝ ██║██║  ██║╚███╔███╔╝██║  ██║███████╗ ╚████╔╝ 
╚═╝     ╚═╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝  ╚═══╝  
                                                      
`

type ResponseItem struct {
	Domain string `json:"domain"`
}

type ResponseData []ResponseItem

func reverseIP(ip string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("https://api.webscan.cc/?action=query&ip=%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("\u001b[31m[ERROR]\u001b[0m Failed to fetch data for %s: %s\n", ip, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\u001b[31m[ERROR]\u001b[0m Failed to read response for %s: %s\n", ip, err)
		return
	}

	if len(body) == 0 {
		fmt.Printf("\u001b[31m[ERROR]\u001b[0m Empty response for %s\n", ip)
		return
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("\u001b[31m[ERROR]\u001b[0m Failed to parse JSON for %s: %s\n", ip, err)
		return
	}

	fmt.Printf("\u001b[32m[%s] Get [%d]\u001b[0m\n", ip, len(data))
	for _, item := range data {
		results <- item.Domain
	}
}

func main() {
	fmt.Println(banner)
	fmt.Printf("\u001b[33mDate : %s\u001b[0m\n", time.Now().Format("2006-01-02"))
	fmt.Println("\u001b[34mCoded By : @Maw3six\u001b[0m")

	var inputFile, outputFile string
	var threadCount int

	fmt.Print("\u001b[35mlist : \u001b[0m")
	fmt.Scanln(&inputFile)

	fmt.Print("\u001b[35mThread : \u001b[0m")
	fmt.Scanln(&threadCount)

	fmt.Print("\u001b[35mSave as : \u001b[0m")
	fmt.Scanln(&outputFile)

	fmt.Println("\n\u001b[36mSTART ..\u001b[0m\n")

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("\u001b[31m[ERROR]\u001b[0m Failed to open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup
	results := make(chan string, 100)

	go func() {
		output, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("\u001b[31m[ERROR]\u001b[0m Failed to create output file: %s\n", err)
			os.Exit(1)
		}
		defer output.Close()
		writer := bufio.NewWriter(output)
		for result := range results {
			writer.WriteString(result + "\n")
			writer.Flush()
		}
	}()

	scanner := bufio.NewScanner(file)
	semaphore := make(chan struct{}, threadCount)
	for scanner.Scan() {
		ip := scanner.Text()
		wg.Add(1)
		semaphore <- struct{}{}
		go func(ip string) {
			reverseIP(ip, results, &wg)
			<-semaphore
		}(ip)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("\u001b[31m[ERROR]\u001b[0m Failed to read input file: %s\n", err)
	}

	wg.Wait()
	close(results)
}
