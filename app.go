package main
 
import (
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    "github.com/elbuo8/4square-venues"
)
 
func main() {
    mux := mux.NewRouter()
    mux.HandleFunc("/", IndexHandler).Methods("GET")
    mux.HandleFunc("/venues/{query}", FoursquareHandler).Methods("GET")
    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":3000")
}
 
 
func IndexHandler(w http.ResponseWriter, r *http.Request){
	p := "Page"
	fmt.Fprintf(w, "Home %s\n", p)
}
 
 
func FoursquareHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    category := vars["query"]
    fs := fsvenues.NewFSVenuesClient("FOURSQUARE_ID", "FOURSQUARE_SECRET")
    params := make(map[string]string)
    params["ll"] = "32.7,-96.8"
    params["limit"] = "5"
    params["query"] = category
 
    if v, e := fs.GetVenues(params); e == nil {
        data, _ := json.Marshal(v)
        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    } else {
        fmt.Println(e)
    }
}