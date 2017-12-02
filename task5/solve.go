package main

import "fmt"
import "net/http"
import "strconv"

func main() {
    dict := make(map[string]string)
    num := 0;
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path[1:]
        if path == "" && r.Method == "POST" {
            num += 1
            dict[strconv.Itoa(num)] = r.FormValue("url")
            
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