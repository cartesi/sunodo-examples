package main

import (
	"encoding/json"
	"io/ioutil"
  "strconv"
  "fmt"
  "log"
  "os"
  "bytes"
  "strings"
  "net/http"
  "dapp/rollups"
)

var (
  infolog  = log.New(os.Stderr, "[ info ]  ", log.Lshortfile)
  errlog   = log.New(os.Stderr, "[ error ] ", log.Lshortfile)
)

func HandleAdvance(data *rollups.AdvanceResponse) error {
  fmt.Printf("Received advance request data %+v\n", data)
	fmt.Println("Adding notice")
  rollupServer := os.Getenv("ROLLUP_HTTP_SERVER_URL")
  if !strings.HasPrefix(rollupServer, "http://") && !strings.HasPrefix(rollupServer, "https://") {
		rollupServer = "http://" + rollupServer
	}

	notice := map[string]string{"payload": data.Payload}
	jsonData, err := json.Marshal(notice)
	if err != nil {
		return err
	}

	response, err := http.Post(rollupServer+"/notice", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Received notice status %d with body %s\n", response.StatusCode, body)
	return nil
}

func HandleInspect(data *rollups.InspectResponse) error {
	fmt.Printf("Received inspect request data %+v\n", data)
	fmt.Println("Adding report")
	rollupServer := os.Getenv("ROLLUP_HTTP_SERVER_URL")
  if !strings.HasPrefix(rollupServer, "http://") && !strings.HasPrefix(rollupServer, "https://") {
		rollupServer = "http://" + rollupServer
	}

	report := map[string]string{"payload": data.Payload}
	jsonData, err := json.Marshal(report)
	if err != nil {
		return err
	}

	response, err := http.Post(rollupServer+"/report", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Printf("Received report status %d\n", response.StatusCode)
	return nil
}

func Handler(response *rollups.FinishResponse) error {
  var err error

  switch response.Type {
  case "advance_state":
    data := new(rollups.AdvanceResponse)
    if err = json.Unmarshal(response.Data, data); err != nil {
      return fmt.Errorf("Handler: Error unmarshaling advance:", err)
    }
    err = HandleAdvance(data)
  case "inspect_state":
    data := new(rollups.InspectResponse)
    if err = json.Unmarshal(response.Data, data); err != nil {
      return fmt.Errorf("Handler: Error unmarshaling inspect:", err)
    }
    err = HandleInspect(data)
  }
  return err
}

func main() {
  finish := rollups.FinishRequest{"accept"}

  for true {
    infolog.Println("Sending finish")
    res, err := rollups.SendFinish(&finish)
    if err != nil {
      errlog.Panicln("Error: error making http request: ", err)
    }
    infolog.Println("Received finish status ", strconv.Itoa(res.StatusCode))
    
    if (res.StatusCode == 202){
      infolog.Println("No pending rollup request, trying again")
    } else {

      resBody, err := ioutil.ReadAll(res.Body)
      if err != nil {
        errlog.Panicln("Error: could not read response body: ", err)
      }
      
      var response rollups.FinishResponse
      err = json.Unmarshal(resBody, &response)
      if err != nil {
        errlog.Panicln("Error: unmarshaling body:", err)
      }

      finish.Status = "accept"
      err = Handler(&response)
      if err != nil {
        errlog.Println(err)
        finish.Status = "reject"
      }
    }
  }
}