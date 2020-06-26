package Proxy

import (
	"testing"
)

func TestAgentTask_RentHouse(t *testing.T) {
	agent :=NewAgentTask()
	agent.RentHouse("beijing",8000)
}