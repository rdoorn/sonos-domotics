package main

// https://www.domoticz.com/forum/viewtopic.php?t=1785 // virtual device?
// https://www.domoticz.com/forum/viewtopic.php?t=10940 // sonos
// https://github.com/jishi/node-sonos-http-api // sonos api
// https://www.domoticz.com/forum/viewtopic.php?t=11577 // update virtual device
// https://github.com/dhleong/ps4-waker/issues/14 // ps4 waker -> netflix

import (
	"flag"
	"log"
	"os"
)

type sonosConfig struct {
	addrStr     *string
	loginStr    *string
	passwordStr *string
}

var config sonosConfig

func init() {
	config.addrStr = flag.String("addr", os.Getenv("IMAP_ADDR"), "imap address:port")
	config.loginStr = flag.String("login", os.Getenv("IMAP_LOGIN"), "imap login")
	config.passwordStr = flag.String("password", os.Getenv("IMAP_PASSWORD"), "imap password")
	flag.Parse()
	if *config.addrStr == "" {
		flag.Usage()
		log.Fatal("address required!")
	}
	if *config.loginStr == "" {
		flag.Usage()
		log.Fatal("login required!")
	}
	if *config.passwordStr == "" {
		flag.Usage()
		log.Fatal("password required!")
	}
}

func main() {
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

func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}
*/
