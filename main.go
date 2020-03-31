package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var baseURL string
var sysChain string
var authToken string

type testRunner struct {
	failed    int
	succeeded int
}

func (t *testRunner) test(testFunc func() error) {
	if err := testFunc(); err != nil {
		t.failed++
	} else {
		t.succeeded++
	}
}

func runAllTests() error {
	runner := testRunner{}

	runner.test(knownHeadsOfChain)
	runner.test(headBlockDetails)
	runner.test(chainIdentifier)
	runner.test(currentCheckpoint)
	runner.test(rpcDocs)
	runner.test(listP2PConnections)
	runner.test(listPeers)
	runner.test(listPoolConnectionPoints)
	runner.test(selfPeerID)
	runner.test(nodeBandwidthStats)
	runner.test(supportedNetworkVersion)
	runner.test(listProtocols)
	runner.test(garbageCollectorStats)
	runner.test(memoryUsageStats)
	runner.test(blockValidatorWorkerState)
	runner.test(listChainValidators)
	runner.test(chainValidatorWorkerState)
	runner.test(workerDDBState)
	runner.test(listValidatorWorkers)
	runner.test(listPrevalidators)
	runner.test(stateOfPrevalidator)

	if runner.failed > 0 {
		return fmt.Errorf("%d out of %d tests failed", runner.failed, runner.failed+runner.succeeded)
	}

	return nil
}

func knownHeadsOfChain() error {
	params := "/chains/" + sysChain + "/blocks"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func headBlockDetails() error {
	params := "/chains/" + sysChain + "/blocks/head"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func chainIdentifier() error {
	params := "/chains/" + sysChain + "/chain_id"
	statusCode, data, err := tezosGetString(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func currentCheckpoint() error {
	params := "/chains/" + sysChain + "/checkpoint"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""
	level := data["block"].(map[string]interface{})["level"].(float64)

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if level == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d expected level greater than 0 got %d", params, statusCode, int(level))
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func rpcDocs() error {
	params := "/describe?recurse=true"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listP2PConnections() error {
	params := "/network/connections"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listPeers() error {
	params := "/network/peers"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listPoolConnectionPoints() error {
	params := "/network/points"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func selfPeerID() error {
	params := "/network/self"
	statusCode, data, err := tezosGetString(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func nodeBandwidthStats() error {
	params := "/network/stat"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""
	totalSent, _ := strconv.ParseFloat(data["total_sent"].(string), 32)
	totalRecv, _ := strconv.ParseFloat(data["total_recv"].(string), 32)

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if totalSent == 0 || totalRecv == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d total data sent %d total data received %d", params, statusCode, int(totalSent), int(totalRecv))
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func supportedNetworkVersion() error {
	params := "/network/version"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listProtocols() error {
	params := "/protocols"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func garbageCollectorStats() error {
	params := "/stats/gc"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func memoryUsageStats() error {
	params := "/stats/memory"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func blockValidatorWorkerState() error {
	params := "/workers/block_validator"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""
	phase := data["status"].(map[string]interface{})["phase"]

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if phase != "running" {
		logMessage = fmt.Sprintf("FAILED: %s phase: '%s' but expected 'running'", params, phase)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listChainValidators() error {
	params := "/workers/chain_validators"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func chainValidatorWorkerState() error {
	params := "/workers/chain_validators/" + sysChain
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""
	phase := data["status"].(map[string]interface{})["phase"]

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if phase != "running" {
		logMessage = fmt.Sprintf("FAILED: %s phase: '%s' but expected 'running'", params, phase)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func workerDDBState() error {
	params := "/workers/chain_validators/" + sysChain + "/ddb"
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""
	activeChains := data["active_chains"].(float64)
	activeConnections := data["active_connections"].(float64)

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if activeChains == 0 || activeConnections == 0 {
		logMessage = fmt.Sprintf("FAILED: %s active_chains: %d active_connections: %d", params, int(activeChains), int(activeConnections))
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listValidatorWorkers() error {
	params := "/workers/chain_validators/" + sysChain + "/peers_validators"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func listPrevalidators() error {
	params := "/workers/prevalidators"
	statusCode, data, err := tezosGetSlice(params)
	logMessage := ""
	prevalidator := data[0].(map[string]interface{})
	prevalidatorStatus := prevalidator["status"].(map[string]interface{})
	prevalidatorPhase := prevalidatorStatus["phase"]

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if prevalidatorPhase != "running" {
		logMessage = fmt.Sprintf("FAILED: %s phase: '%s' but expected 'running'", params, prevalidatorPhase)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func stateOfPrevalidator() error {
	params := "/workers/prevalidators/" + sysChain
	statusCode, data, err := tezosGetMap(params)
	logMessage := ""
	prevalidatorStatus := data["status"].(map[string]interface{})
	prevalidatorPhase := prevalidatorStatus["phase"]

	if err != nil {
		logMessage = fmt.Sprintf("FAILED: %s %s", params, err)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode != 200 || len(data) == 0 {
		logMessage = fmt.Sprintf("FAILED: %s %d %s", params, statusCode, data)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if prevalidatorPhase != "running" {
		logMessage = fmt.Sprintf("FAILED: %s phase: '%s' but expected 'running'", params, prevalidatorPhase)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage = fmt.Sprintf("PASSED: %s %d", params, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage = fmt.Sprintf("FAILED: %s unknown failure", params)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func getURL(params string) string {
	url := baseURL + params
	if authToken != "" {
		url = url + "?auth=" + authToken
	}

	return url
}

func tezosGetSlice(params string) (int, []interface{}, error) {
	resp, err := http.Get(getURL(params))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var data []interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)

	return resp.StatusCode, data, err
}

func tezosGetMap(params string) (int, map[string]interface{}, error) {
	resp, err := http.Get(getURL(params))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)

	return resp.StatusCode, data, err
}

func tezosGetString(params string) (int, string, error) {
	resp, err := http.Get(getURL(params))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var data string
	err = json.NewDecoder(resp.Body).Decode(&data)

	return resp.StatusCode, data, err
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("USAGE: %s <base-url> <sys-chain> [auth-token]\n", os.Args[0])
		os.Exit(1)
	}

	baseURL = os.Args[1]
	sysChain = os.Args[2]
	if len(os.Args) > 3 {
		authToken = os.Args[3]
	}

	if err := runAllTests(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}
}
