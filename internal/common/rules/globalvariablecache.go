package rules

import "github.com/pkg/errors"

type VariableValuePair struct {
	VariableName VariableFieldName
	Values       []float64
}

var VariableMap = map[VariableFieldName]VariableValuePair{}

// RegisterNewVariable Registers the provided variable in the global variable cache
func RegisterNewVariable(pair VariableValuePair) error {
	return registerNewVariableInternal(pair, VariableMap)
}

// registerNewVariableInternal provides primal register logic for any variable cache
func registerNewVariableInternal(pair VariableValuePair, variableStore map[VariableFieldName]VariableValuePair) error {
	if _, ok := variableStore[pair.VariableName]; ok {
		return errors.Errorf("attempted to re-register a variable that had already been registered")
	}
	variableStore[pair.VariableName] = pair
	return nil
}

// UpdateVariable Updates variable in global cache with new value
func UpdateVariable(variableName VariableFieldName, newValue VariableValuePair) error {
	return updateVariableInternal(variableName, newValue, VariableMap)
}

// updateVariableInternal provides primal update logic for any variable cache
func updateVariableInternal(variableName VariableFieldName, newValue VariableValuePair, variableStore map[VariableFieldName]VariableValuePair) error {
	if _, ok := variableStore[variableName]; ok {
		variableStore[variableName] = newValue
		return nil
	}
	return errors.Errorf("attempted to modify a variable has not been defined")
}

type VariableFieldName int

const (
	NumberOfIslandsContributingToCommonPool VariableFieldName = iota
	NumberOfFailedForages
	NumberOfBrokenAgreements
	MaxSeverityOfSanctions
	NumberOfIslandsAlive
	NumberOfBallotsCast
	NumberOfAllocationsSent
	IslandsAlive
	SpeakerSalary
	JudgeSalary
	PresidentSalary
	ExpectedTaxContribution
	ExpectedAllocation
	IslandTaxContribution
	IslandAllocation
	TestVariable
)
