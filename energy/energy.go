package energy

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "strings"
)

type GetPower struct {
    Channel1voltage string `json:"Channel1 voltage(V)"`
    Channel1current string `json:"Channel1 current(A)"`
	Channel2voltage string `json:"Channel2 voltage(V)"`
    Channel2current string `json:"Channel2 current(A)"`
    Channel3voltage string `json:"Channel3 voltage(V)"`
    Channel3current string `json:"Channel3 current(A)"`
	
    Inputvoltage string `json:"Input voltage(V)"`
    Inputcurrent string `json:"Input current(A)"`
}

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
//******************************************************************************************

func Voltage() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Inputvoltage
        return
    }else{
        response = "ERROR"
        return
    }
}
func Current() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Inputcurrent
        return
    }else{
        response = "ERROR"
        return
    }
}
func Power() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        voltage, err := strconv.ParseFloat(power.Inputvoltage, 32)
        if err != nil {
            voltage = 0
        }
        current, err := strconv.ParseFloat(power.Inputcurrent, 32)
        if err != nil {
            current = 0
        }
        response =  fmt.Sprintf("%f", (voltage * current))
        return
    }else{
        response = "ERROR"
        return
    }
}
//******************************************************************************************
func Ch1voltage() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Channel1voltage
        return
    }else{
        response = "ERROR"
        return
    }
}
func Ch1current() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Channel1current
        return
    }else{
        response = "ERROR"
        return
    }
}

//******************************************************************************************
func Ch2voltage() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Channel2voltage
        return
    }else{
        response = "ERROR"
        return
    }
}
func Ch2current() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Channel2current
        return
    }else{
        response = "ERROR"
        return
    }
}

//******************************************************************************************
func Ch3voltage() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Channel3voltage
        return
    }else{
        response = "ERROR"
        return
    }
}
func Ch3current() (response string){
    data := GetJson("http://127.0.0.1:13029/params")
    var power GetPower
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &power)
        response = power.Channel3current
        return
    }else{
        response = "ERROR"
        return
    }
}


//******************************************************************************************

func Request(key string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "voltage":
        return Voltage()
    case "current":
        return Current()
    case "power":
        return Power()	
	case "voltagech1":
        return Ch1voltage()
	case "currentch1":
        return Ch1current()
	case "voltagech2":
        return Ch2voltage()
	case "currentch2":
        return Ch2current()	
	case "voltagech3":
        return Ch3voltage()
	case "currentch3":
        return Ch3current()		
    default:
      return "ERROR Unsupported Metric"
    }
}
