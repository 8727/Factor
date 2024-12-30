package archive
 
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"net/http"
	"io/ioutil"
	"time"
)

type GetCount struct {Count uint32}
type GetLastintervals struct {TracksCount uint32}

type GetArchive struct {
    MinAllowedDiskFreeSpaceGBytes float32
    MaxAllowedTrackAgeDays uint32
    ArchiveSizeCurrentGBytes float32
    ArchiveSizeMaxGBytes float32
    TracksCountCurrent uint32
    TracksCountMax uint32
    AvailableDiskSpaceGBytes float32
    TotalDiskSizeGBytes float32
    OldestTrackAgeDays uint32
}
//******************************************************************************************
func GetJson(url string) (json string){
    http.DefaultClient.Timeout = 25 * time.Second
    resp, err := http.Get(url)
    if err != nil {
        return "ERROR"
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "ERROR"
    }
    return string(body)
}
//******************************************************************************************
func CountCSV(url string, threshoMin float64, threshoMax float64) (count int, countmin int, countmax int){
	countmin = 0
	countmax = 0
	count = 0
	http.DefaultClient.Timeout = 25 * time.Second
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	record, e := reader.Read()
	for {
		record, e = reader.Read()
		if e != nil {
			break
		}
		value, _ := strconv.ParseFloat(record[1], 64)
		if value > threshoMin {
           countmax = countmax + 1
		}
		if value < threshoMax {
           countmin = countmin + 1
		}
		count = count + 1
	}
	return
}
//******************************************************************************************
func ReadJson(data string) (response string){
	var count GetCount
	if data != "ERROR"{
		json.Unmarshal([]byte(data), &count)
		return fmt.Sprintf("%d", count.Count)
	}else{
		return "ERROR"
	}
}
//******************************************************************************************
func PercentToString(countmax int, count int) (percentString string){
	if countmax > 0 {
		return fmt.Sprintf("%.2f", (100 * float64(count) / float64(countmax)))
	}else{
		return "0.0001"
	}
}
//******************************************************************************************
func StringToInt(countStrint string) (countInt int){
	i, _ := strconv.Atoi(countStrint)
	countInt = i
	return
}
//******************************************************************************************
func GetQuery(key string, timeRange string) (response string){
	key = strings.ToLower(key)
    switch key {
    case "cars":
		return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange))
		
    case "allviolations":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=Speeding20%3BSpeeding40%3BSpeeding60%3BSpeeding80%3BRoadside%3BWrongWay%3BSeatBelt%3BLights%3BBusLane%3BStopping%3BParking%3BRedLightSingleFrontier%3BRedLightDoubleFrontier%3BStopLineSingleFrontier%3BStopLineDoubleFrontier%3BPhoneInHand%3BProhibitedManeuver"))
    case "allspeed":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=Speeding20%3BSpeeding40%3BSpeeding60%3BSpeeding80"))
    case "roadside":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=Roadside"))
    case "wrongway":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=WrongWay"))
    case "seatbelt":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=SeatBelt"))
    case "lights":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=Lights"))
    case "buslane":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=BusLane"))
    case "stopping":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=Stopping"))
    case "parking":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=Parking"))
    case "redlight":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=RedLightSingleFrontier%3BRedLightDoubleFrontier"))
    case "stopline":
		return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=StopLineSingleFrontier%3BStopLineDoubleFrontier"))
    case "phoneinhand":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=PhoneInHand"))
    case "prohibitedmaneuver":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=ProhibitedManeuver"))
	case "pedestrians":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&violationTypes=NoYieldToPedestrian"))
	
	case "nospeed":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&WithoutSpeedOnly=true"))
	case "percentnospeed":
		m := StringToInt(ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange)))
		c := StringToInt(ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&WithoutSpeedOnly=true")))
        return PercentToString(m, c)
		
	case "zerospeed":
		return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&maxSpeed=1"))
	case "percentzerospeed":
		m := StringToInt(ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange)))
		c := StringToInt(ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&maxSpeed=1")))
        return PercentToString(m, c)
	
	case "percentradar":
		count, _, countmax := CountCSV("http://127.0.0.1/archive/tracks/csv?" + timeRange + "&columns=trackId&columns=radarSpeed&maxRowCount=10000", 0.1, 1000)
		return PercentToString(count, countmax)
		
	case "shorttrack":
		_, countmin, _ := CountCSV("http://127.0.0.1/archive/tracks/csv?" + timeRange + "&columns=trackId&columns=trackPointsCount&maxRowCount=10000", 80, 10)
        return fmt.Sprintf("%d", countmin)
	case "percentshorttrack":
        count, countmin, _ := CountCSV("http://127.0.0.1/archive/tracks/csv?" + timeRange + "&columns=trackId&columns=trackPointsCount&maxRowCount=10000", 80, 10)
		return PercentToString(count, countmin)
		
    case "longtrack":
        _, _, countmax := CountCSV("http://127.0.0.1/archive/tracks/csv?" + timeRange + "&columns=trackId&columns=trackPointsCount&maxRowCount=10000", 80, 10)
        return fmt.Sprintf("%d", countmax)
    case "percentlongtrack":
        count, _, countmax := CountCSV("http://127.0.0.1/archive/tracks/csv?" + timeRange + "&columns=trackId&columns=trackPointsCount&maxRowCount=10000", 80, 10)
		return PercentToString(count, countmax)
	
	case "noplate":
        return ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&regId=NOPLATE*"))
    case "percentnoplate":
		m := StringToInt(ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange)))
		c := StringToInt(ReadJson(GetJson("http://127.0.0.1/archive/tracks/count?" + timeRange + "&regId=NOPLATE*")))
        return PercentToString(m, c)

    default:
		return "ERROR Unsupported Metric"
    }
}
//******************************************************************************************
func GetTimeZona() (timeZona string){
    dateTime := time.Now()
    sourceZona := string(dateTime.Format("2006-01-02T15:04:05-0700"))[len(string(dateTime.Format("2006-01-02T15:04:05"))):]
    if sourceZona[0] == '+'{
        timeZona = "%2B"
    }else{
        timeZona = "%2D"
    }
    sourceZona = sourceZona[1:]
    timeZona = timeZona + sourceZona
    sourceZona = sourceZona[2:]
    timeZona = timeZona[:5] + "%3A" + sourceZona
    return
}
//******************************************************************************************
func Yesterday(key string) (response string){
    dateTime := time.Now()
    timeZona := GetTimeZona()
    dateEnd := string(dateTime.Format("2006-01-02"))
    dateTime = dateTime.Add(-24 * time.Hour)
    dateStart := string(dateTime.Format("2006-01-02"))
	return GetQuery(key, "dateStart=" + dateStart + "T00%3A00%3A00" + timeZona + "&dateEnd=" + dateEnd + "T00%3A00%3A00" + timeZona)
}
//******************************************************************************************
func PreviousHour(key string) (response string){
    dateTime := time.Now()
    timeZona := GetTimeZona()
    dateEnd := string(dateTime.Format("2006-01-02T15:04:05"))
    dateEnd = dateEnd[:strings.IndexAny(dateEnd, ":")]
    dateTime = dateTime.Add(-1 * time.Hour)
    dateStart := string(dateTime.Format("2006-01-02T15:04:05"))
    dateStart = dateStart[:strings.IndexAny(dateStart, ":")]
	return GetQuery(key, "dateStart=" + dateStart + "%3A00%3A00" + timeZona + "&dateEnd=" + dateEnd + "%3A00%3A00" + timeZona)
}
//******************************************************************************************
func LastHours(key string, hours string) (response string){
    dateTime := time.Now()
    timeZona := GetTimeZona()
    dateEnd := string(dateTime.Format("2006-01-02T15:04:05"))
    dateEnd = strings.Replace(dateEnd, ":", "%3A", -1)
    dateTime = dateTime.Add(-(time.Duration(StringToInt(hours))) * time.Hour)
    dateStart := string(dateTime.Format("2006-01-02T15:04:05"))
    dateStart = strings.Replace(dateStart, ":", "%3A", -1)
	return GetQuery(key, "dateStart=" + dateStart + timeZona + "&dateEnd=" + dateEnd + timeZona)
}
//******************************************************************************************
func Currentday(key string) (response string){
    dateTime := time.Now()
    timeZona := GetTimeZona()
    dateStart := string(dateTime.Format("2006-01-02"))
    dateTime = dateTime.Add(+24 * time.Hour)
    dateEnd := string(dateTime.Format("2006-01-02"))
	return GetQuery(key, "dateStart=" + dateStart + "T00%3A00%3A00" + timeZona + "&dateEnd=" + dateEnd + "T00%3A00%3A00" + timeZona)
}
//******************************************************************************************
func ArchiveSizeCurrentGBytes() (response string){
    data := GetJson("http://127.0.0.1/archive/trackscleaner/volumesinfo")
    var archive GetArchive
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &archive)
        response = fmt.Sprintf("%f", (archive.ArchiveSizeCurrentGBytes))
        return
    }else{
        response = "ERROR"
        return
    }
}
func TracksCountCurrent() (response string){
    data := GetJson("http://127.0.0.1/archive/trackscleaner/volumesinfo")
    var archive GetArchive
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &archive)
        response = fmt.Sprintf("%d", archive.TracksCountCurrent)
        return
    }else{
        response = "ERROR"
        return
    }
}
func AvailableDiskSpaceGBytes() (response string){
    data := GetJson("http://127.0.0.1/archive/trackscleaner/volumesinfo")
    var archive GetArchive
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &archive)
        response = fmt.Sprintf("%f", archive.AvailableDiskSpaceGBytes)
        return
    }else{
        response = "ERROR"
        return
    }
}
func TotalDiskSizeGBytes() (response string){
    data := GetJson("http://127.0.0.1/archive/trackscleaner/volumesinfo")
    var archive GetArchive
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &archive)
        response = fmt.Sprintf("%f", archive.TotalDiskSizeGBytes)
        return
    }else{
        response = "ERROR"
        return
    }
}
func OldestTrackAgeDays() (response string){
    data := GetJson("http://127.0.0.1/archive/trackscleaner/volumesinfo")
    var archive GetArchive
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &archive)
        response = fmt.Sprintf("%d", archive.OldestTrackAgeDays)
        return
    }else{
        response = "ERROR"
        return
    }
}
//******************************************************************************************
func Archive(key string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "archivesizecurrentgbytes":
        return ArchiveSizeCurrentGBytes()
    case "trackscountcurrent":
        return TracksCountCurrent()
    case "availablediskspacegbytes":
        return AvailableDiskSpaceGBytes()
    case "totaldisksizegbytes":
        return TotalDiskSizeGBytes()
    case "oldesttrackagedays":
        return OldestTrackAgeDays()
    default:
      return "ERROR Unsupported Metric"
    }
}

//******************************************************************************************
func Lastintervalm(seconds string) (response string){
    value, err := strconv.ParseInt(seconds, 10, 64)
    if err != nil {
        response = "ERROR"
        return
    }
	data := GetJson("http://127.0.0.1/archive/statistics?timeIntervalS=" + fmt.Sprintf("%d", value * 60))
    var cars GetLastintervals
    if data != "ERROR"{
        json.Unmarshal([]byte(data), &cars)
        response = fmt.Sprintf("%d", cars.TracksCount)
        return
    }else{
        response = "ERROR"
        return
    }
}
