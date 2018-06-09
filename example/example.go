package example

import (
	"bufio"
	"github.com/beito123/binary"
	id "github.com/satori/go.uuid"
	"github.com/SpecterTeam/go-raknet"
	"github.com/SpecterTeam/go-raknet/identifier"
	"github.com/SpecterTeam/go-raknet/server"
	"github.com/SpecterTeam/SpecterGO/utils"
	"os"
)
//Start an MCPE Server.
func MCPEServer() {
	//Create a logger.
	logger := utils.NewLogger()

	//Generate a uuid.
	uuid, _ := id.NewV4()

	//Load the Minecraft data.
	mc := identifier.Minecraft{
		Connection:        raknet.ConnectionGoRaknet,
		ServerName:        "SpecterGO-Server",
		ServerProtocol:    raknet.NetworkProtocol,
		VersionTag:        "1.0.0",
		OnlinePlayersCount: 0,
		MaxPlayersCount:    10,
		GUID:              binary.ReadLong(uuid.Bytes()[0:8]),
		WorldName:         "world",
		Gamemode:          "0",
		Legacy:            false,
	}

	//Load the server with the minecraft data.
	ser := &server.Server{
		Logger:              &logger,
		MaxConnections:      10,
		MTU:                 1472,
		Identifier:          mc,
		UUID:                uuid,
		BroadcastingEnabled: true,
	}

	logger.Info("Starting the server...")

	//Start the server with the default port and local ip.
	ser.Start("0.0.0.0", 19132)

	logger.Info("Server started! Press enter to stop the server.")
	//Wait until anything is input
	bufio.NewScanner(os.Stdin).Scan()

	//Shutdown the server and leave.
	ser.Shutdown()

	logger.Info("Stopping the server...")
}