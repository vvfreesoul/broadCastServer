package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type HostConfig struct {
	IP       string
	CPU      string
	Mem      string
	Storage  string
	Hostname string
}

func FindRemoteInsConfigByIP(ip string) (*HostConfig, error) {
	// Run SSH command to get the hostname
	sshHostnameCmd := exec.Command("ssh", ip, "hostname")
	hostnameBytes, err := sshHostnameCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute SSH command to get hostname: %v", err)
	}
	hostname := string(hostnameBytes)

	// Run SSH command to retrieve CPU information
	sshCPUCmd := exec.Command("ssh", ip, "cat /proc/cpuinfo")
	cpuOut, err := sshCPUCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve CPU information: %v", err)
	}

	// Parse CPU information
	cpuRegex := regexp.MustCompile(`model name\s+:\s+(.*)\n`)
	cpuMatch := cpuRegex.FindStringSubmatch(string(cpuOut))
	cpu := ""
	if len(cpuMatch) > 1 {
		cpu = cpuMatch[1]
	}

	// Run SSH command to retrieve memory information
	sshMemCmd := exec.Command("ssh", ip, "cat /proc/meminfo")
	memOut, err := sshMemCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve memory information: %v", err)
	}

	// Parse memory information
	memRegex := regexp.MustCompile(`MemTotal:\s+(\d+) kB\n`)
	memMatch := memRegex.FindStringSubmatch(string(memOut))
	mem := ""
	if len(memMatch) > 1 {
		memBytes := 1024 * 1024 * atoi(memMatch[1])
		mem = fmt.Sprintf("%d GB", memBytes/(1024*1024*1024))
	}

	// Run SSH command to retrieve storage information
	sshStorageCmd := exec.Command("ssh", ip, "df -h /")
	storageOut, err := sshStorageCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve storage information: %v", err)
	}

	// Parse storage information
	storageRegex := regexp.MustCompile(`/dev/.*?\s+(\d+G)\s+(\d+G)\s+(\d+G)\s+(\d+)%\s+/`)
	storageMatch := storageRegex.FindStringSubmatch(string(storageOut))
	storage := ""
	if len(storageMatch) > 4 {
		storage = fmt.Sprintf("%s used (%s/%s total)", storageMatch[4], storageMatch[2], storageMatch[1])
	}

	// Construct HostConfig object
	config := &HostConfig{
		IP:       ip,
		CPU:      cpu,
		Mem:      mem,
		Storage:  storage,
		Hostname: hostname,
	}

	return config, nil
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	// Test the function
	ip := "192.168.177.8" // Replace with the actual IP address
	config, err := FindRemoteInsConfigByIP(ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("IP: %s\nHostname: %s\nCPU: %s\nMem: %s\nStorage: %s\n", config.IP, config.Hostname, config.CPU, config.Mem, config.Storage)
}
