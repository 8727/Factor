package unit

import (
    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
    "fmt"
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
//  fmt.Println(url)
    return
}

type GetVersion struct {Version string}
type GetUnit struct {FactoryNumber string}
type GetCertificate struct {
	SerialNumber string
	Title string
	Number string
	IssueDate string
	ValidUpTo string
}
type GetUnitInfo struct {
    Unit GetUnit
    Certificate GetCertificate
}
type GetLocation struct {
    InstallationPlace string
    DirectionTo string
    DirectionFrom string
}
type GetGnns struct {
    Latitude float64
    Longitude float64
}
type GetCam struct {
    Model string
    Focus_length_mm float32
}
type GetLens struct {ModelName string}
type GetAnalyzer struct {Mode string}
type GetRadar struct {Enabled bool}
type GetNc4 struct {Enabled bool}

func Version() (response string){
    data := GetJson("http://127.0.0.1/updater/installed-factor-version")
    var version GetVersion
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &version)
        response = version.Version
        return
    }else{
        response = "ERROR"
        return
    }
}
func Factory() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/unitinfo")
    var info GetUnitInfo
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &info)
        response = info.Unit.FactoryNumber
        return
    }else{
        response = "ERROR"
        return
    }
}
func Serial() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/unitinfo")
    var info GetUnitInfo
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &info)
        response = info.Certificate.SerialNumber
        return
    }else{
        response = "ERROR"
        return
    }
}

func Name() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/unitinfo")
    var info GetUnitInfo
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &info)
        response = info.Certificate.Title
        return
    }else{
        response = "ERROR"
        return
    }
}
func SeriesAndNumber() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/unitinfo")
    var info GetUnitInfo
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &info)
        response = info.Certificate.Number
        return
    }else{
        response = "ERROR"
        return
    }
}
func IssuanceDate() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/unitinfo")
    var info GetUnitInfo
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &info)
        response = info.Certificate.IssueDate
        return
    }else{
        response = "ERROR"
        return
    }
}
func ValidUntil() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/unitinfo")
    var info GetUnitInfo
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &info)
        response = info.Certificate.ValidUpTo
        return
    }else{
        response = "ERROR"
        return
    }
}

func Location() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/Location")
    var location GetLocation
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &location)
        response = location.InstallationPlace
        return
    }else{
        response = "ERROR"
        return
    }
}
func DirectionTo() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/Location")
    var location GetLocation
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &location)
        response = location.DirectionTo
        return
    }else{
        response = "ERROR"
        return
    }
}
func DirectionFrom() (response string){
    data := GetJson("http://127.0.0.1/unitinfo/api/Location")
    var location GetLocation
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &location)
        response = location.DirectionFrom
        return
    }else{
        response = "ERROR"
        return
    }
}
func Cam() (response string){
    data := GetJson("http://127.0.0.1:13001/configure/recognition_camera")
    var cam GetCam
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &cam)
        response = cam.Model
        return
    }else{
        response = "ERROR"
        return
    }
}
func Lens() (response string){
    data := GetJson("http://127.0.0.1/videomodule/api/VideoModule/recognition/lens")
    var lens GetLens
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &lens)
        response = lens.ModelName
        return
    }else{
        response = "ERROR"
        return
    }
}
func LensFocus() (response string){
    data := GetJson("http://127.0.0.1:13001/configure/recognition_camera")
    var cam GetCam
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &cam)
        response = fmt.Sprint(cam.Focus_length_mm)
        return
    }else{
        response = "ERROR"
        return
    }
}
func Analyzer() (response string){
    data := GetJson("http://127.0.0.1:13001/configure/analyzer/current")
    var analyzer GetAnalyzer
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &analyzer)
        response = analyzer.Mode
        return
    }else{
        response = "ERROR"
        return
    }
}
func Radar() (response string){
    data := GetJson("http://127.0.0.1:13001/configure/radar")
    var radar GetRadar
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &radar)
        if radar.Enabled == true{
            response = "True"
        }else{
            response = "False"
        }
        return
    }else{
        response = "ERROR"
        return
    }
}
func T3() (response string){
    data := GetJson("http://127.0.0.1:13001/configure/nc4_server")
    var nc4 GetNc4
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &nc4)
        if nc4.Enabled == true{
            response = "True"
        }else{
            response = "False"
        }
        return
    }else{
        response = "ERROR"
        return
    }
}

func Request(key string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "version":
        return Version()
    case "factory":
        return Factory()
    case "serial":
        return Serial()

    case "name":
        return Name()		
    case "number":
        return SeriesAndNumber()		
    case "issuedate":
        return IssuanceDate()		
    case "validupto":
        return ValidUntil()		
		
    case "location":
        return Location()
    case "directionto":
        return DirectionTo()
    case "directionfrom":
        return DirectionFrom()
    case "cam":
        return Cam()
    case "lens":
        return Lens()
    case "lensfocus":
        return LensFocus()
    case "analyzer":
        return Analyzer()
    case "radar":
        return Radar()
    case "t3":
        return T3()
    default:
      return "ERROR Unsupported Metric"
    }
}