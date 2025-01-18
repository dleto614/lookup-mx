package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	var file *string
	var output *string
	var domain *string

	var data []string

	var results []string
	var mx []*net.MX

	file = flag.String("i", "", "Specify input file with a list of domains.")
	domain = flag.String("d", "", "Specify domain.")
	output = flag.String("o", "", "Specify the output file to write results into.")

	flag.Parse()

	if ChkFlag("i") {
		if ChkFlag("o") {
			fmt.Printf("[*] Reading file: %s \n", *file)
		}

		data = ReadFile(file)
	} else if ChkFlag("d") {
		data = append(data, *domain)
	} else {
		usage()
		os.Exit(1)
	}

	for _, DataString := range data {

		var err error

		mx, err = net.LookupMX(DataString)

		// Check if nil or not for error after looking up MX records.
		if err != nil {
			// Print error and go back to the for loop without continuing the code in the for loop.

			// This will skip this domain since it doesn't have a valid MX record so we don't need
			// to do anythng else and can just iterate to the next domain in our array.
			log.Println(err)
			continue
		}
		
		JsonResult := ConvertJson(DataString, results)

		// Write to stdin if no output file arg was passed
		if !ChkFlag("o") {
			//fmt.Printf("MX records for '%s' is: \n", DataString)
			/*for _, result := range results {
				fmt.Println(result)
			}*/
			fmt.Println(string(JsonResult))
		} else {
			// Write to file
			FileWrite(string(JsonResult), *output)
		}
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: \n", os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
}

// Convert the domain and mx records to json

func ConvertJson(domain string, results []string) []byte {

	type MxRecords struct {
		Domain string   `json:"domain"`
		MX     []string `json:"mx"`
	}

	RecordsData := MxRecords{
		Domain: domain,
		MX:     results,
	}

	JsonResult, _ := json.Marshal(RecordsData)

	return JsonResult
}

// Extract the host in the MX struct and append to []string that is returned.

func ParseMX(mx []*net.MX) []string {

	var results []string

	for _, record := range mx {
		//fmt.Println(string(record.Host))
		results = append(results, record.Host)
	}

	return results
}

func ChkFlag(name string) bool {
	found := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}

func ReadFile(file *string) []string {

	var data []string

	fd, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	fd.Close()
	return data
}

func FileWrite(data string, data_file string) {

	file, err := os.OpenFile(data_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Create and open file to append into

	if err != nil {
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil { // Write data to file
		return
	}

}
