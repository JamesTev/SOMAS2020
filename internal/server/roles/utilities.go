package roles

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common/rules"
	"github.com/SOMAS2020/SOMAS2020/internal/server"
	"github.com/pkg/errors"
)

var serverReference *server.SOMASServer

func setServerReference(serverRef *server.SOMASServer) {
	serverReference = serverRef
}

// PickUpRulesByVariable returns a list of rule_id's which are affected by certain variables
func PickUpRulesByVariable(variableName string, ruleStore map[string]rules.RuleMatrix) ([]string, error) {
	var Rules []string
	if _, ok := rules.VariableMap[variableName]; ok {
		for k, v := range ruleStore {
			_, err := searchForStringInArray(variableName, v.RequiredVariables)
			if err != nil {
				Rules = append(Rules, k)
			}
		}
		return Rules, nil
	} else {
		return []string{}, errors.Errorf("Variable name '%v' was not found in the variable cache", variableName)
	}
}

func searchForStringInArray(val string, array []string) (int, error) {
	for i, v := range array {
		if v == val {
			return i, nil
		}
	}
	return 0, errors.Errorf("Not found")
}

type DataPacket struct {
	integerData int
	textData    string
	booleanData bool
}

type Communication struct {
	recipient int
	sender    int
	data      map[int]DataPacket
}

func broadcastToAllIslands(sender int, data map[int]DataPacket) {
	islandsAlive := rules.VariableMap["islands_alive"]
	for _, v := range islandsAlive.Values {
		communicateWithIslands(int(v), sender, data)
	}
}

func communicateWithIslands(recipient int, sender int, data map[int]DataPacket) {
	communication := Communication{
		recipient: recipient,
		sender:    sender,
		data:      data,
	}
	//Send to islands
	print(communication) //// Get rid of this
}

const (
	BallotID                 = iota
	PresidentAllocationCheck = iota
	SpeakerID                = iota
	RoleConducted            = iota
	ResAllocID               = iota
	SpeakerBallotCheck       = iota
	PresidentID              = iota
	RuleName                 = iota
	RuleVoteResult           = iota
)
