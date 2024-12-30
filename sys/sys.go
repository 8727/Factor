package sys

import (
	"factor/archive"
    "fmt"
    "io/ioutil"
    "os/exec"
    "strings"
)

func NetSpeed(message string) (response string){
    message = strings.ToLower(message)
    if message == "eth0" || message == "eth1" {
        fileSpeed, err := ioutil.ReadFile("/sys/class/net/" + message + "/speed")
        if err != nil {
            response = "ERROR"
        }
        response = string(fileSpeed[:(len(fileSpeed)-1)])
        return
    }else{
        response = "ERROR Unsupported Metric"
        return
    }
}

func License(message string) (response string){
    message = strings.ToLower(message)
	switch message {
    case "reset":
        cmd := exec.Command("sh", "-c", "rm -f /opt/vlabs/vsdk/data/license.sdt")
        err := cmd.Start()
        if err != nil {
            response = fmt.Sprintf("ERROR Cannot execute command: %s", err)
            return
        }
        go cmd.Wait()
        cmd = exec.Command("sh", "-c", "systemctl restart factor-appearance.service")
        err = cmd.Start()
        if err != nil {
            response = fmt.Sprintf("ERROR Cannot execute command: %s", err)
            return
        }
        go cmd.Wait()
        cmd = exec.Command("sh", "-c", "systemctl restart factor-enricher.service")
        err = cmd.Start()
        if err != nil {
            response = fmt.Sprintf("ERROR Cannot execute command: %s", err)
            return
        }
        go cmd.Wait()
		response = "License Reset"
        return
	case "data":
		out, err := exec.Command("sh", "-c", "stat -c %y /opt/vlabs/vsdk/data/license.sdt").Output()
        if err != nil {
            response = "ERROR No license file"
			return
        }
		response = string(out)[:strings.IndexAny(string(out), ".")]
		return
    default:
		response = "ERROR Unsupported Metric"
		return
    }
}

func Fsync(message string) (response string){
    message = strings.ToLower(message)
	switch message {
	case "reset":
		cmd := exec.Command("sh", "-c", "systemctl restart fsync-service.service")
        err := cmd.Start()
        if err != nil {
            response = fmt.Sprintf("ERROR Cannot execute command: %s", err)
            return
        }
        go cmd.Wait()
        response = "Fsync Reset"
        return
	case "status":
		get := archive.Lastintervalm("60")
		if get == "0" || get == "ERROR" {
			cmd := exec.Command("sh", "-c", "systemctl restart fsync-service.service")
			err := cmd.Start()
			if err != nil {
				response = fmt.Sprintf("ERROR Cannot execute command: %s", err)
				return
			}
			go cmd.Wait()
			response = "1"
			return
		}else{
			response = "0"
			return
		}
    default:
		response = "ERROR Unsupported Metric"
		return
    }
}

func Vision(message string) (response string){
    message = strings.ToLower(message)
    if message == "reset" {
        cmd := exec.Command("sh", "-c", "systemctl restart vision.service")
        err := cmd.Start()
        if err != nil {
            response = fmt.Sprintf("ERROR Cannot execute command: %s", err)
            return
        }
        go cmd.Wait()
        response = "Vision Reset"
        return
    }else{
        response = "ERROR Unsupported Metric"
        return
    }
}

func Cli(message string) (response string){
	out, err := exec.Command("sh", "-c", message).Output()
    if err != nil {
		response = "ERROR Cli"
		return
    }
	response = string(out)
	return
}

func Request(key string, message string) (response string){
    key = strings.ToLower(key)
    switch key {
    case "netspeed":
        return NetSpeed(message)
    case "license":
        return License(message)
    case "fsync":
        return Fsync(message)
	case "vision":
        return Vision(message)
	case "cli":
        return Cli(message)
    default:
      return "ERROR Unsupported Metric"
    }
}