package scanner

import (
	"testing"
)

func TestOpenPort(t *testing.T){
	portsToScan := []int{10, 22, 80, 122, 555}
	expectedOpen := 2

    actualOpen, _ := PortScanner("scanme.nmap.org", portsToScan) 

    if expectedOpen != actualOpen {
        t.Errorf("expected %d, got %d", expectedOpen, actualOpen)
    }
}

func TestClosedPort(t *testing.T){
	portsToScan := []int{10, 22, 80, 122, 555}
	expectedClosed := 3

    _, actualClosed := PortScanner("scanme.nmap.org", portsToScan) 

    if expectedClosed != actualClosed {
        t.Errorf("expected %d, got %d", expectedClosed, actualClosed)
    }
}

func TestTotalPortsScanned(t *testing.T){
	portsToScan := []int{10, 22, 80, 122, 555}
	expectedPorts := 5

    actualOpen, actualClosed := PortScanner("scanme.nmap.org", portsToScan)
	actualPorts := actualOpen + actualClosed

    if expectedPorts != actualPorts {
        t.Errorf("expected %d, got %d", actualPorts, expectedPorts)
    }
}
