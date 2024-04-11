package main

import (
	"github.com/anjush-bhargavan/go_trade_notification/pkg/config"
	"github.com/anjush-bhargavan/go_trade_notification/pkg/rabbitmq"
)

func main() {

	cfg := config.LoadConfig()
	rabbitmq.ConsumeNotificationMessages(cfg)

}