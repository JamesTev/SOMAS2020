// Package team5 contains code for team 5's client implementation
package team5

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common/baseclient"
	"github.com/SOMAS2020/SOMAS2020/internal/common/gamestate"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

const id = shared.Team5

func init() {
	baseclient.RegisterClient(
		id,
		&client{
			BaseClient:          baseclient.NewClient(id),
			cpRequestHistory:    CPRequestHistory{},
			cpAllocationHistory: CPAllocationHistory{},
			taxAmount:           0,
			allocation:          0,
			config: clientConfig{
				JBThreshold:       100,
				MiddleThreshold:   60,
				ImperialThreshold: 30,
			},
		},
	)
}

type client struct {
	*baseclient.BaseClient

	cpRequestHistory    CPRequestHistory
	cpAllocationHistory CPAllocationHistory

	taxAmount shared.Resources

	// allocation is the president's response to your last common pool resource request
	allocation shared.Resources

	config clientConfig
}

func (c client) wealth() wealthTier {
	cData := c.gameState().ClientInfo
	switch {
	case cData.LifeStatus == shared.Critical:
		return dying
	case cData.Resources > c.config.ImperialThreshold && cData.Resources < c.config.MiddleThreshold:
		return imperialStudent
	case cData.Resources > c.config.JBThreshold:
		return jeffBezos
	default:
		return middleClass
	}
}

func (c *client) StartOfTurn() {
	c.Logf("Wealth state: %v", c.wealth())
	c.Logf("Resources: %v", c.gameState().ClientInfo.Resources)

	for clientID, status := range c.gameState().ClientLifeStatuses {
		if status != shared.Dead && clientID != c.GetID() {
			return
		}
	}
	c.Logf("I'm all alone :c")

}

func (c *client) gameState() gamestate.ClientGameState {
	return c.BaseClient.ServerReadHandle.GetGameState()
}

//------------------------------------Comunication--------------------------------------------------------//
// to get information on minimum tax amount and cp allocation
func (c *client) ReceiveCommunication(
	sender shared.ClientID,
	data map[shared.CommunicationFieldName]shared.CommunicationContent,
) {
	for field, content := range data {
		switch field {
		case shared.TaxAmount:
			c.taxAmount = shared.Resources(content.IntegerData)
		case shared.AllocationAmount:
			c.allocation = shared.Resources(content.IntegerData)
		}
	}
}
