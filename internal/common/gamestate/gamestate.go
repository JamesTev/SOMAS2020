// Package gamestate contains information about the current game state.
package gamestate

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common/disasters"
	"github.com/SOMAS2020/SOMAS2020/internal/common/foraging"
	"github.com/SOMAS2020/SOMAS2020/internal/common/roles"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

type IIGOBaseRoles struct {
	BasePresident roles.President
	BaseSpeaker   roles.Speaker
	BaseJudge     roles.Judge
}

// GameState represents the game's state.
type GameState struct {
	// Season represents the current (1-index) season of the game.
	Season uint

	// Turn represents the current (1-index) Turn of the game.
	Turn uint

	// CommonPool represents the amount of resources in the common pool.
	CommonPool shared.Resources

	// ClientInfos map from the shared.ClientID to ClientInfo.
	ClientInfos    map[shared.ClientID]ClientInfo
	Environment    disasters.Environment
	DeerPopulation foraging.DeerPopulationModel

	// IIGOInfo contains states of the base versions of all roles
	IIGOInfo IIGOBaseRoles

	// [INFRA] add more details regarding state of game here
	// REMEMBER TO EDIT `Copy` IF YOU ADD ANY REFERENCE TYPES (maps, slices, channels, functions etc.)
}

// Copy returns a deep copy of the GameState.
func (g GameState) Copy() GameState {
	ret := g
	ret.ClientInfos = copyClientInfos(g.ClientInfos)
	ret.Environment = g.Environment.Copy()
	ret.DeerPopulation = g.DeerPopulation.Copy()
	return ret
}

// GetClientGameStateCopy returns the ClientGameState for the client having the id.
func (g *GameState) GetClientGameStateCopy(id shared.ClientID) ClientGameState {
	return ClientGameState{
		Season:     g.Season,
		Turn:       g.Turn,
		ClientInfo: g.ClientInfos[id].Copy(),
	}
}

func copyClientInfos(m map[shared.ClientID]ClientInfo) map[shared.ClientID]ClientInfo {
	ret := make(map[shared.ClientID]ClientInfo, len(m))
	for k, v := range m {
		ret[k] = v.Copy()
	}
	return ret
}

// ClientInfo contains the client struct as well as the client's attributes
type ClientInfo struct {
	// Resources contains the amount of resources owned by the client.
	Resources shared.Resources // nice

	LifeStatus shared.ClientLifeStatus
	// CriticalConsecutiveTurnsCounter is the number of consecutive turns the client is in critical state.
	// Client will die if this Counter reaches config.MaxCriticalConsecutiveTurns
	CriticalConsecutiveTurnsCounter uint

	// [INFRA] add more client information here
	// REMEMBER TO EDIT `Copy` IF YOU ADD ANY REFERENCE TYPES (maps, slices, channels, functions etc.)
}

// Copy returns a deep copy of the ClientInfo.
func (c ClientInfo) Copy() ClientInfo {
	ret := c
	return ret
}
