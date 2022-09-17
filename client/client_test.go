package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAllServers(t *testing.T) {
	// Start all servers
	cmd := exec.Command("bash", "-c", "python3 ../run_servers.py")
	stdout, err := cmd.CombinedOutput()

	t.Logf(string(stdout))
	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// somewhat frequent pattern
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=Ec --pattern=\"Windows N[A-Za-z][A-Za-z0-9]*\" --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// Check if total match is consistent
	wanted_total_match := "676160"
	t.Logf(string(stdout))
	splitted := strings.Fields(string(stdout))
	total_match := splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// frequent
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=c --pattern=http --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	wanted_total_match = "2709263"
	t.Logf(string(stdout))
	splitted = strings.Fields(string(stdout))
	total_match = splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// Rare case
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=c --pattern=alexander.com --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	wanted_total_match = "2843"
	t.Logf(string(stdout))
	splitted = strings.Fields(string(stdout))
	total_match = splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// turn down all servers
	cmd = exec.Command("bash", "-c", "python3 ../stop_servers.py")
	stdout, err = cmd.CombinedOutput()
	t.Logf(string(stdout))
}

func TestAllFaultTolerance(t *testing.T) {
	// Start only four servers
	cmd := exec.Command("bash", "-c", "python3 ../run_servers.py --servers 1 2 3 4")
	stdout, err := cmd.CombinedOutput()

	t.Logf(string(stdout))
	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// Frequent Pattern
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=Ec --pattern=http --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// Check if total match is consistent
	wanted_total_match := "1091212"
	t.Logf(string(stdout))
	splitted := strings.Fields(string(stdout))
	total_match := splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// rare pattern
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=c --pattern=alexander.com --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// Check if total match is consistent
	wanted_total_match = "1080"
	t.Logf(string(stdout))
	splitted = strings.Fields(string(stdout))
	total_match = splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// turn down servers
	cmd = exec.Command("bash", "-c", "python3 ../stop_servers.py --servers 1 2 3 4")
	stdout, err = cmd.CombinedOutput()
	t.Logf(string(stdout))
}

func TestPartialExist(t *testing.T) {
	// Start only four servers
	cmd := exec.Command("bash", "-c", "python3 ../run_servers.py --servers 1 2 3 4")
	stdout, err := cmd.CombinedOutput()

	t.Logf(string(stdout))
	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// pattern only appear in vm1
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=c --pattern=05/Feb/2024:14 --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// Check if total match is consistent
	wanted_total_match := "19"
	t.Logf(string(stdout))
	splitted := strings.Fields(string(stdout))
	total_match := splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// pattern appear in vm 2 3 4
	cmd = exec.Command("bash", "-c", "cd ..; go run client/client.go --flags=c --pattern=www.mckay.net --prefix=vm")
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		t.Errorf("CMD %s\n", cmd.String())
		t.Errorf(string(stdout))
		t.Errorf(err.Error())
	}

	// Check if total match is consistent
	wanted_total_match = "8"
	t.Logf(string(stdout))
	splitted = strings.Fields(string(stdout))
	total_match = splitted[len(splitted)-1]

	if wanted_total_match != total_match {
		t.Errorf("want %v, but got %v", wanted_total_match, total_match)
	}

	// turn down servers
	cmd = exec.Command("bash", "-c", "python3 ../stop_servers.py --servers 1 2 3 4")
	stdout, err = cmd.CombinedOutput()
	t.Logf(string(stdout))
}
