package crontab

import "log"

type CrontabDemo struct {
}

func newCrontabDemo() CrontabDemo {
	return CrontabDemo{}
}

func (c CrontabDemo) getSpec() string {
	return "@every 5m"
}

func (c CrontabDemo) Run() {
	log.Println("*******-crontab-test-******")
}
