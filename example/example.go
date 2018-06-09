/**
 *     SpecterGO  Copyright (C) 2018  SpecterTeam
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
