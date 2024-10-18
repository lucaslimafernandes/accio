package readfiles

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestReadHostsFile(t *testing.T) {

	var hostsToml *Hosts

	// Get the dir path where is the code file is located
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	projectPath := filepath.Join(dir, "../../", "hosts.example.toml")

	hostsToml, err := ReadHostsFile(&projectPath)
	if err != nil {
		t.Fatalf("Failed to read 'hosts.example.toml', expected: nil, got: %v", err)
	}

	// 2 Nodes
	if len(hostsToml.Nodes) != 2 {
		t.Fatalf("Expected 2 nodes, got: %v", len(hostsToml.Nodes))
	}

	// Test user node 1
	if hostsToml.Nodes[0].User != "ubuntu" {
		t.Fatalf("Expected user: ubuntu, got: %v", hostsToml.Nodes[0].User)
	}

}
