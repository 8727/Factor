package exporter

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "time"
    "encoding/json"
    "strings"
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
    return
}

type GetExporter struct {
    UsedBytes uint64
    TotalBytes uint64
    UsedPercent uint8
    CountOfItems uint32
    LastExportedItemTime string
    LastExportItemDeletedTime string
    FirstNotExportedItemTime string
    RemovedItemsCount uint64
}

func UsedPercent(message string) (response string){
    data := GetJson("http://127.0.0.1/exporter/Endpoints/" + message + "/destination/queue")
    var exporter GetExporter
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &exporter)
        response = fmt.Sprintf("%d", exporter.UsedPercent)
        return
    }else{
        response = "ERROR"
        return
    }
}
func CountOfItems(message string) (response string){
    data := GetJson("http://127.0.0.1/exporter/Endpoints/" + message + "/destination/queue")
    var exporter GetExporter
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &exporter)
        response = fmt.Sprintf("%d", exporter.CountOfItems)
        return
    }else{
        response = "ERROR"
        return
    }
}

func Request(key string, message string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "usedpercent":
        return UsedPercent(message)
    case "countofitems":
        return CountOfItems(message)
    default:
      return "ERROR Unsupported Metric"
    }
}

/*
AisMtp
stream_camea
monitoring_doris
stream_duplo2_n1
stream_duplo2_n2
stream_ftp_n1
stream_ftp_n2
stream_local
violation_duplo2_n1
violation_duplo2_n2
violation_ftp_n1
violation_ftp_n2
violation_local
*/