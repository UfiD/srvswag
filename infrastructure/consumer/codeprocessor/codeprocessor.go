package codeprocessor

import "time"

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Do() string {
	time.Sleep(time.Minute)
	return "DONE"
}
