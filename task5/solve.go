package main

import "fmt"
import "net/http"
import "strconv"
import "strings"
import "io/ioutil"

func main() {
    dict := make(map[string]string)
    num := 0;
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path[1:]
        if path == "" && r.Method == "POST" {
            num += 1
			defer r.Body.Close()
			bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
            dict[strconv.Itoa(num)] = strings.Split(bodyString, "\"")[3]
            
            w.Header().Set("Content-Type", "application/json")
            fmt.Fprintf(w, "{\"key\": \"%d\"}", num)
        } else if path != "" {
			url, ok := dict[path]
			if ok {
				http.Redirect(w, r, url, http.StatusMovedPermanently)
			}
        }
    })
    
    http.ListenAndServe(":8082", nil)
}