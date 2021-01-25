// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package service

import (
	"github.com/livekit/livekit-server/pkg/auth"
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/pkg/rtc"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config, keyProvider auth.KeyProvider, roomStore RoomStore, router routing.Router, currentNode routing.LocalNode, selector routing.NodeSelector) (*LivekitServer, error) {
	rtcConfig := rtc.RTCConfigFromConfig(conf)
	externalIP := externalIpFromNode(currentNode)
	webRTCConfig, err := rtc.NewWebRTCConfig(rtcConfig, externalIP)
	if err != nil {
		return nil, err
	}
	roomManager := NewRoomManager(roomStore, router, currentNode, selector, webRTCConfig)
	roomService, err := NewRoomService(roomManager)
	if err != nil {
		return nil, err
	}
	rtcService := NewRTCService(conf, roomStore, roomManager, router, currentNode)
	livekitServer, err := NewLivekitServer(conf, roomService, rtcService, keyProvider, router, roomManager, currentNode)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil
}
