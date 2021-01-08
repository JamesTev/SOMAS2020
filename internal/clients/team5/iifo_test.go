package team5

import (
	"testing"

	"github.com/SOMAS2020/SOMAS2020/internal/common/disasters"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

var c = initClient()

func TestAnalyseDisasterPeriod(t *testing.T) {
	// can use same spatial and mag info because we're only assessing period
	dInfo := disasterInfo{report: disasters.DisasterReport{X: 0, Y: 0, Magnitude: 1}}
	dh1 := disasterHistory{}
	dh2 := disasterHistory{8: dInfo}
	dh3 := disasterHistory{3: dInfo, 5: dInfo, 7: dInfo, 9: dInfo}

	var tests = []struct {
		name       string
		dh         disasterHistory
		wantPeriod uint // output tier
		wantConf   float64
	}{
		{"no past disasters", dh1, 0, 0},
		{"1 past disaster", dh2, 7, 50},
		{"many periodic disasters", dh3, 2, 100},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c.disasterHistory = tc.dh
			ansPeriod, ansConf := c.estimateDisasterPeriod()
			if ansPeriod != tc.wantPeriod {
				t.Errorf("period: got %d, want %d", ansPeriod, tc.wantPeriod)
			}
			if ansConf != tc.wantConf {
				t.Errorf("conf: got %.3f, want %.3f", ansConf, tc.wantConf)
			}
		})
	}
}

func TestUpdateForecastingReputations(t *testing.T) {
	receivedPreds := shared.ReceivedDisasterPredictionsDict{
		shared.Team1: shared.ReceivedDisasterPredictionInfo{
			PredictionMade: shared.DisasterPrediction{
				Confidence: 60,
			},
			SharedFrom: shared.Team1,
		},
		shared.Team2: shared.ReceivedDisasterPredictionInfo{
			PredictionMade: shared.DisasterPrediction{
				Confidence: 20,
			},
			SharedFrom: shared.Team2,
		},
		shared.Team3: shared.ReceivedDisasterPredictionInfo{
			PredictionMade: shared.DisasterPrediction{
				Confidence: 100,
			},
			SharedFrom: shared.Team3,
		},
	}
	c.opinions = opinionMap{
		shared.Team1: &wrappedOpininon{opinion{forecastReputation: 0.0}},
		shared.Team2: &wrappedOpininon{opinion{forecastReputation: 0.0}},
		shared.Team3: &wrappedOpininon{opinion{forecastReputation: 0.0}},
	}
	c.disasterHistory = disasterHistory{} // no disasters recorded
	c.updateForecastingReputations(receivedPreds)
	if c.opinions[shared.Team1].getForecastingRep() >= 0 {
		t.Error("Received prediction with confidence > 50 percent with no disasters")
	}

	if c.opinions[shared.Team2].getForecastingRep() < 0 {
		t.Error("Expected no negative change to reputation after sensible prediction")
	}
	c.disasterHistory = disasterHistory{1: disasterInfo{}, 5: disasterInfo{}} // no disasters recorded
	c.updateForecastingReputations(receivedPreds)

	if c.opinions[shared.Team3].getForecastingRep() >= 0 {
		t.Error("Received perfectly confident prediction. Expected rep. to decrease.")
	}

}

func initClient() *client {
	c := createClient()
	return c
}