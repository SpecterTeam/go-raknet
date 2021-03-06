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
package identifier

import (
	"strconv"

	"github.com/SpecterTeam/go-raknet"
)

const (
	MinecraftHeader      = "MCPE"
	MinecraftSeparator   = ";"
	MinecraftCountLegacy = 6
	MinecraftCount       = 9
)

//var MinecraftVersionTagAlphabet = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}

type Minecraft struct {
	Connection *raknet.ConnectionType

	ServerName        string
	ServerProtocol    int
	VersionTag        string
	OnlinePlayersCount int
	MaxPlayersCount    int
	GUID              int64
	WorldName         string
	Gamemode          string
	Legacy            bool
}

func (id Minecraft) ConnectionType() *raknet.ConnectionType {
	return id.Connection
}

func (id Minecraft) Build() string {
	if id.Legacy {
		return MinecraftHeader + MinecraftSeparator +
			id.ServerName + MinecraftSeparator +
			strconv.Itoa(id.ServerProtocol) + MinecraftSeparator +
			id.VersionTag + MinecraftSeparator +
			strconv.Itoa(id.OnlinePlayersCount) + MinecraftSeparator +
			strconv.Itoa(id.MaxPlayersCount)
	}

	return MinecraftHeader + MinecraftSeparator +
		id.ServerName + MinecraftSeparator +
		strconv.Itoa(id.ServerProtocol) + MinecraftSeparator +
		id.VersionTag + MinecraftSeparator +
		strconv.Itoa(id.OnlinePlayersCount) + MinecraftSeparator +
		strconv.Itoa(id.MaxPlayersCount) + MinecraftSeparator +
		strconv.FormatInt(id.GUID, 10) + MinecraftSeparator +
		id.WorldName + MinecraftSeparator +
		id.Gamemode
}

func (id *Minecraft) SetServerName(name string) {
	id.ServerName = name
}

func (id *Minecraft) SetServerProtocol(protocol int) {
	id.ServerProtocol = protocol
}

func (id *Minecraft) SetVersionTag(version string) {
	id.VersionTag = version
}

func (id *Minecraft) SetOnlinePlayersCount(count int) {
	id.OnlinePlayersCount = count
}

func (id *Minecraft) SetMaxPlayersCount(max int) {
	id.MaxPlayersCount = max
}

func (id *Minecraft) SetGUID(guid int64) {
	id.GUID = guid
}

func (id *Minecraft) SetWorldName(name string) {
	id.WorldName = name
}

func (id *Minecraft) SetGameMode(gamemode string) {
	id.Gamemode = gamemode
}

func (id *Minecraft) SetLegacy(l bool) {
	id.Legacy = l
}