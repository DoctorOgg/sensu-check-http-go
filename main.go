package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Url         string
	CheckString string
	HttpTimeout int
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-check-http-go",
			Short:    "A simple replacement for the ruby based http check for sensu",
			Keyspace: "sensu.io/plugins/sensu-check-http-go/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "url",
			Env:       "CHECK_URL",
			Argument:  "url",
			Shorthand: "u",
			Default:   "",
			Usage:     "URL to check",
			Value:     &plugin.Url,
		},
		&sensu.PluginConfigOption{
			Path:      "checkstring",
			Env:       "CHECKSTRING",
			Argument:  "checkstring",
			Shorthand: "c",
			Default:   "",
			Usage:     "String to Match",
			Value:     &plugin.CheckString,
		},
		&sensu.PluginConfigOption{
			Path:      "httptimeout",
			Env:       "HTTPTIMEOUT",
			Argument:  "timeout",
			Shorthand: "t",
			Default:   10,
			Usage:     "Timeout value",
			Value:     &plugin.HttpTimeout,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if len(plugin.Url) == 0 || len(plugin.CheckString) == 0 {
		return sensu.CheckStateWarning, fmt.Errorf("Please specify an URL ( -u https://example.com ) and a check string ( -c \"farts\" ) for this check to run")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {

	c := http.Client{Timeout: time.Duration(plugin.HttpTimeout) * time.Second}
	resp, err := c.Get(plugin.Url)
	if err != nil {
		fmt.Printf("URL: %s, Error %s", plugin.Url, err)
		return sensu.CheckStateCritical, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("ERROR: %s, status: %v", plugin.Url, resp.StatusCode)
		return sensu.CheckStateCritical, nil
	} else if resp.StatusCode == 200 && !strings.Contains(string(body), plugin.CheckString) {
		fmt.Printf("ERROR: %s, status: %v, String not found: %s", plugin.Url, resp.StatusCode, plugin.CheckString)
		return sensu.CheckStateCritical, nil
	}
	fmt.Printf("OK: %s, status: %v, String found: %s", plugin.Url, resp.StatusCode, plugin.CheckString)
	return sensu.CheckStateOK, nil

}
