package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func main() {
	// Read URLs from file or standard input
	urls, err := readURLs(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create empty list of parameters
	parameters := make(url.Values)

	// Loop through URLs and extract parameters
	for _, u := range urls {
		// Parse URL to extract parameters
		u, err := url.Parse(u)
		if err != nil {
			fmt.Printf("Error parsing URL %s: %s\n", u, err)
			continue
		}
		parameters = mergeValues(parameters, u.Query())
	}

	// Output list of parameter names
	for k := range parameters {
		fmt.Println(k)
	}
}

// Read URLs from file or standard input
func readURLs(r *os.File) ([]string, error) {
	urls := make([]string, 0)

	if stat, _ := r.Stat(); stat != nil && stat.Mode()&os.ModeNamedPipe == 0 {
		// Read URLs from file
		f, err := os.Open(r.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	} else {
		// Read URLs from standard input
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return urls, nil
}

// Merge two url.Values maps
func mergeValues(dest, src url.Values) url.Values {
	for k, v := range src {
		dest[k] = append(dest[k], v...)
	}
	return dest
}
