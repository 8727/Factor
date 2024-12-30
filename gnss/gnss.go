package gnss

import (
    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
    "strings"
    "fmt"
)

func GetJson(url string) (json string){
    http.DefaultClient.Timeout = 2 * time.Second
    resp, err := http.Get(url)
    if err != nil {
        json = "ERROR"
        return
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        json = "ERROR"
        return
    }
    json = string(body)
//  fmt.Println(url)
    return
}

type GetGnns struct {
    State string
    Latitude float64
    Longitude float64
}

func Status() (response string){
    data := GetJson("http://127.0.0.1/gnss/coords")
    var gnns GetGnns
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &gnns)
        response = gnns.State
        return
    }else{
        response = "ERROR"
        return
    }
}
func Latitude() (response string){
    data := GetJson("http://127.0.0.1/gnss/coords")
    var gnns GetGnns
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &gnns)
        response = fmt.Sprint(gnns.Latitude)
        return
    }else{
        response = "ERROR"
        return
    }
}
func Longitude() (response string){
    data := GetJson("http://127.0.0.1/gnss/coords")
    var gnns GetGnns
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &gnns)
        response = fmt.Sprint(gnns.Longitude)
        return
    }else{
        response = "ERROR"
        return
    }
}

func Request(key string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "status":
        return Status()
    case "latitude":
        return Latitude()
    case "longitude":
        return Longitude()
    default:
      return "ERROR Unsupported Metric"
    }
}
