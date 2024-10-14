package main

type AlertService struct {
	chat          DicscordChat
	chatBandwidth chatBandwidth
}

func NewAlertService(chat DicscordChat, chatBandwidth chatBandwidth) *AlertService {
	return &AlertService{
		chat: chat,

		chatBandwidth: chatBandwidth,
	}
}

func (a *AlertService) Start() {
	a.chat.SendMessage("message")
}

type DicscordChat struct {
	url      string
	username string
	password string
}

func (d DicscordChat) SendMessage(message string) {
	// send message to discord

}

type chatBandwidth struct {
	url      string
	username string
	password string
}

func (c chatBandwidth) SendMessage(message string) {
	// send message to chat bandwidth
}
