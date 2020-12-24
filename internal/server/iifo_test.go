package server

import (
	"reflect"
	"testing"

	"github.com/SOMAS2020/SOMAS2020/internal/common/baseclient"
	"github.com/SOMAS2020/SOMAS2020/internal/common/gamestate"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
	"github.com/SOMAS2020/SOMAS2020/pkg/testutils"
)

type mockClientIIFO struct {
	baseclient.Client
	foragingValues  shared.ShareForageInfomation
	otherIslandInfo []shared.ShareForageInfomation
}

func (c *mockClientIIFO) MakeForageInfo() shared.ShareForageInfomation {
	return c.foragingValues
}

func (c *mockClientIIFO) ReceiveForageInfo(otherIslandInfo []shared.ShareForageInfomation) {
	c.otherIslandInfo = otherIslandInfo
}

func (c *mockClientIIFO) getOtherIslandInfo() []shared.ShareForageInfomation {
	return c.otherIslandInfo
}

func makeForagingInfo(contribution shared.Resources, resources shared.Resources, shareTo []shared.ClientID) shared.ShareForageInfomation {
	if len(shareTo) > 0 {
		return shared.ShareForageInfomation{
			DecisionMade:     shared.ForageDecision{Type: shared.DeerForageType, Contribution: contribution},
			ResourceObtained: resources,
			ShareTo:          shareTo,
		}
	}
	// People can be selfish and choose not to share their foraging information
	return shared.ShareForageInfomation{
		DecisionMade:     shared.ForageDecision{Type: shared.DeerForageType, Contribution: contribution},
		ResourceObtained: resources,
	}
}

func receiveForagingInfo(contribution shared.Resources, resources shared.Resources, sharedFrom shared.ClientID) shared.ShareForageInfomation {
	return shared.ShareForageInfomation{
		DecisionMade:     shared.ForageDecision{Type: shared.DeerForageType, Contribution: contribution},
		ResourceObtained: resources,
		SharedFrom:       sharedFrom,
	}
}

func TestGetForageSharingWorks(t *testing.T) {
	clientInfos := map[shared.ClientID]gamestate.ClientInfo{
		shared.Team1: {
			LifeStatus: shared.Alive,
		},
		shared.Team2: {
			LifeStatus: shared.Critical,
		},
		shared.Team3: {
			LifeStatus: shared.Dead,
		},
	}

	clientMap := map[shared.ClientID]baseclient.Client{
		shared.Team1: &mockClientIIFO{
			foragingValues: makeForagingInfo(52.7, 64, []shared.ClientID{shared.Team2, shared.Team3}),
		},
		shared.Team2: &mockClientIIFO{
			foragingValues: makeForagingInfo(22.2, 22.3, []shared.ClientID{}),
		},
		shared.Team3: &mockClientIIFO{
			foragingValues: makeForagingInfo(33.2, 233.3, []shared.ClientID{shared.Team2}),
		},
	}

	want := shared.MakeForagingDict{
		shared.Team1: makeForagingInfo(52.7, 64, []shared.ClientID{shared.Team2, shared.Team3}),
		shared.Team2: makeForagingInfo(22.2, 22.3, []shared.ClientID{}),
	}

	server := &SOMASServer{
		gameState: gamestate.GameState{
			ClientInfos: clientInfos,
		},
		clientMap: clientMap,
	}

	got, err := server.getForageSharing()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want '%#v' got '%#v'", want, got)
	}

	testutils.CompareTestErrors(nil, err, t)
}

func TestDistributeForageSharing(t *testing.T) {
	clientInfos := map[shared.ClientID]gamestate.ClientInfo{
		shared.Team1: {
			LifeStatus: shared.Alive,
		},
		shared.Team2: {
			LifeStatus: shared.Critical,
		},
		shared.Team3: {
			LifeStatus: shared.Dead,
		},
	}

	mockClient := map[shared.ClientID]*mockClientIIFO{
		shared.Team1: {},
		shared.Team2: {},
		shared.Team3: {},
	}

	clientMap := map[shared.ClientID]baseclient.Client{
		shared.Team1: mockClient[shared.Team1],
		shared.Team2: mockClient[shared.Team2],
		shared.Team3: mockClient[shared.Team3],
	}

	input := shared.MakeForagingDict{
		shared.Team1: makeForagingInfo(52.7, 64, []shared.ClientID{shared.Team2, shared.Team3}),
		shared.Team2: makeForagingInfo(22.2, 22.3, []shared.ClientID{}),
	}

	want := shared.ReceiveForagingDict{
		shared.Team1: []shared.ShareForageInfomation(nil),
		shared.Team2: []shared.ShareForageInfomation{receiveForagingInfo(52.7, 64, shared.Team1)},
	}

	server := &SOMASServer{
		gameState: gamestate.GameState{
			ClientInfos: clientInfos,
		},
		clientMap: clientMap,
	}

	//	getOtherIslandInfo
	server.distributeForageSharing(input)
	for id := range mockClient {
		got := mockClient[id].getOtherIslandInfo()
		w := want[id]
		if !reflect.DeepEqual(w, got) {
			t.Errorf("want '%#v' got '%#v' for %#v", w, got, id)
		}
	}
}
