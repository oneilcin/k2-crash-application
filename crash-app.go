/*
Copyright 2017 K2 authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func sendLogs(filename string) error {
	logBuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("Error reading file in sendLogs function: %v", err)
		return err
	}

	log := string(logBuffer)
	// Get specific ansible task that failed if one exists in logs.
	failedTask := getFailedTask(log)

	// Format data for POST request and send to ElasticSearch.
	url := "https://krakencrashreporter.kubeme.io/krakencrashreporter/krakencrashes"
	t := time.Now().UTC()
	ts := t.Format("2006-01-02T15:04:05.000Z")
	body := map[string]string{"k2_log": log, "failed_task": failedTask, "date": ts}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("Error preparing data in json format for POST request: %v", err)
		return err
	}

	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		err = fmt.Errorf("Error creating POST request in http.NewRequest: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	// Error handling for http.DefaultClient.Do(req).
	if err != nil {
		err = fmt.Errorf("Error sending POST request: %v", err)
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		response := fmt.Sprintf("POST request returned a non-200 response code: %v", resp)
		return errors.New(response)
	}

	return nil
}

func getFailedTask(logs string) string {
	// Parse tasks into array of arrays.
	var rgx = regexp.MustCompile(`TASK \[(.+?)\]`)
	failedTasks := rgx.FindAllStringSubmatch(logs, -1)

	// Check that regex found "TASK [...]", if not we assume there were no ansible
	// tasks in the logs string parameter but still collect this data and send to ES.
	if failedTasks == nil {
		failedTask := "Log does not contain any ansible tasks"
		return failedTask
	}

	// Assume the last task is the task that failed.
	failedTask := failedTasks[len(failedTasks)-1][1]
	return failedTask
}

func main() {
	logFile := os.Args[1]

	err := sendLogs(logFile)
	if err != nil {
		fmt.Printf("[CRASHAPP] %v\n", err)
	}
}
