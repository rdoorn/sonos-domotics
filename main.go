package main

// https://www.domoticz.com/forum/viewtopic.php?t=1785 // virtual device?
// https://www.domoticz.com/forum/viewtopic.php?t=10940 // sonos
// https://github.com/jishi/node-sonos-http-api // sonos api
// https://www.domoticz.com/forum/viewtopic.php?t=11577 // update virtual device
// https://github.com/dhleong/ps4-waker/issues/14 // ps4 waker -> netflix

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type sonosConfig struct {
	urlStr      *string
	loginStr    *string
	passwordStr *string
	listenerStr *string
}

var config sonosConfig

func init() {
	config.listenerStr = flag.String("listen", os.Getenv("LISTENER"), "listner address:port")
	config.urlStr = flag.String("url", os.Getenv("SONOS_URL"), "sonos http://address:port")
	config.loginStr = flag.String("login", os.Getenv("SONOS_LOGIN"), "sonos login")
	config.passwordStr = flag.String("password", os.Getenv("SONOS_PASSWORD"), "sonos password")
	flag.Parse()
	if *config.urlStr == "" {
		flag.Usage()
		log.Fatal("url required!")
	}
	if *config.loginStr == "" {
		flag.Usage()
		log.Fatal("login required!")
	}
	if *config.passwordStr == "" {
		flag.Usage()
		log.Fatal("password required!")
	}
	if *config.listenerStr == "" {
		l := "127.0.0.1:5025"
		config.listenerStr = &l
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	groupZones()
	fmt.Printf("query: %s\n", r.URL.Query())
	if volume, ok := r.URL.Query()["volume"]; ok {
		room := "Living Room bar"
		path := strings.Split(r.URL.Path, "/")
		fmt.Printf("len path: %d", len(path))
		if len(path) > 2 {
			room = path[1]

		}
		fmt.Printf("Setting volume for %s to %s\n", room, volume)
		//setVolume(r.URL.Path[1])
	}
	result, err := getApi(r.URL.String())
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Fprintf(w, string(result))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*config.listenerStr, nil))
}

func groupZones() {
	zones, err := getZones()
	if err != nil {
		log.Fatalf("Failed to get zone: %s\n", err)
	}
	//fmt.Printf("Zone data: %+v\n", zones)

	for _, zone := range *zones {
		fmt.Printf("Got zone: %s\n", zone.Coordinator.RoomName)
		if zone.Coordinator.RoomName == "Living Room bar" {
			if len(zone.Members) == 3 {
				fmt.Printf("Living room is grouped!\n")
			} else {
				fmt.Printf("Expected 3 members, but got: %d\n", len(zone.Members))
				expectedMembers := map[string]bool{
					"Living Room2": false,
					"Kitchen":      false,
				}
				for _, member := range zone.Members {
					if _, ok := expectedMembers[member.RoomName]; ok {
						fmt.Printf("member: %s is part of %s\n", member.RoomName, zone.Coordinator.RoomName)
						expectedMembers[member.RoomName] = true
					}
				}
				for name, member := range expectedMembers {
					if member == false {
						fmt.Printf("Need to join %s to Bar\n", name)
						err := setGroup(name, "Living Room bar")
						if err != nil {
							fmt.Printf("Failed to join %s to %s: %s", name, "Living Room bar", err)
						}
					}
				}
				fmt.Printf("Living room is grouped!\n")
			}
		}
	}
}

func getState() (*State, error) {
	data, err := getApi("/state")
	if err != nil {
		return nil, err
	}
	var state State
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func getZones() (*Zones, error) {
	data, err := getApi("/zones")
	if err != nil {
		return nil, err
	}
	var zones Zones
	err = json.Unmarshal(data, &zones)
	if err != nil {
		return nil, err
	}
	return &zones, nil
}

/*
func PostPathData(path string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	url := fmt.Sprintf("%s%s", *domotics.urlStr, path)
	fmt.Printf("GET on: [%s]\n", url)
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(*domotics.loginStr, *domotics.passwordStr)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Output: %s", bodyText)
	return nil
}
*/

func setGroup(room, group string) error {
	path := fmt.Sprintf("/%s/join/%s", UrlEncoded(room), UrlEncoded(group))
	result, err := getApi(path)
	if err != nil {
		return err
	}
	var r Result
	err = json.Unmarshal(result, &r)
	if err != nil {
		return err
	}
	/*
		if r.status != "success" {
			return fmt.Errorf(r.status)
		}*/
	return nil
}

func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}

func getApi(path string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	url := fmt.Sprintf("%s%s", *config.urlStr, path)
	fmt.Printf("GET on: [%s]\n", url)
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(*config.loginStr, *config.passwordStr)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Output: %s", bodyText)
	return bodyText, nil

}
