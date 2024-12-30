package main

import (
    "fmt"
    
    "factor/archive"
    "factor/energy"
    "factor/exporter"
    "factor/gnss"
    "factor/sys"
    "factor/unit"
    
    "git.zabbix.com/ap/plugin-support/plugin"
    "git.zabbix.com/ap/plugin-support/plugin/container"
)

const pluginName = "Factor"

type Plugin struct {
    plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
    p.Infof("received request to handle %s key with %d parameters", key, len(params))
    if len(params) == 0 || params[0] == "" {
        return nil, fmt.Errorf("Invalid first parameter.")
    }
    
    if len(params) > 2 {
        return nil, fmt.Errorf("Too many parameters.")
    }
    
    switch key {
      
    case "archive":
        return archive.Archive(params[0]), nil
      
    case "currentday":
        return archive.Currentday(params[0]), nil
      
    case "energy":
        return energy.Request(params[0]), nil
      
    case "exporter":
        if len(params) > 1 {
            return exporter.Request(params[0], params[1]), nil
        } else {
            return nil, fmt.Errorf("Invalid second parameter.")
        }
      
    case "gnss":
        return gnss.Request(params[0]), nil
      
    case "lasthours":
        if len(params) > 1 {
            return archive.LastHours(params[0], params[1]), nil
        } else {
            return nil, fmt.Errorf("Invalid second parameter.")
        }
    
	case "lastintervalm":
        return archive.Lastintervalm(params[0]), nil
      
    case "previoushour":
        return archive.PreviousHour(params[0]), nil
      
    case "sys":
        if len(params) > 1 {
            return sys.Request(params[0], params[1]), nil
        } else {
            return nil, fmt.Errorf("Invalid second parameter.")
        }
      
    case "unit":
        return unit.Request(params[0]), nil
      
    case "yesterday":
        return archive.Yesterday(params[0]), nil
      
    default:
        return nil, plugin.UnsupportedMetricError
    }
}

func init() {
    plugin.RegisterMetrics(&impl, pluginName, //err := plugin.RegisterMetrics(&impl, pluginName,
        "archive", "Status.",
        "currentday", "Currentday.",
        "energy", "Energy.",
        "exporter", "Exporter.",
        "gnss", "Gnss.",
        "lasthours", "LastHours.",
		"lastintervalm", "LastIntervalM.",
        "previoushour", "PreviousHour.",
        "sys", "Sys.",
        "unit", "Unit.",
        "yesterday", "Yesterday.")
//    if err != nil {
//        panic(errs.Wrap(err, "failed to register metrics"))
//    }
}

func main() {
    h, err := container.NewHandler(impl.Name())
    if err != nil {
        panic(fmt.Sprintf("failed to create plugin handler %s", err.Error()))
    }
    impl.Logger = &h
    
err = h.Execute()
    if err != nil {
        panic(fmt.Sprintf("failed to execute plugin handler %s", err.Error()))
    }
}