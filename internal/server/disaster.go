package server

import "github.com/SOMAS2020/SOMAS2020/internal/common/disasters"

// probeDisaster checks if a disaster occurs this turn
func (s *SOMASServer) probeDisaster() (disasters.Environment, error) {
	s.logf("start probeDisaster")
	defer s.logf("finish probeDisaster")

	e := s.gameState.Environment
	e = e.SampleForDisaster()                                                                // update env instance with sampled disaster info
	effects := e.ComputeDisasterEffects(s.gameState.CommonPool, s.gameConfig.DisasterConfig) // get disaster effects - absolute, proportional and CP-mitigated

	disasterReport := e.DisplayReport(s.gameState.CommonPool, s.gameConfig.DisasterConfig) // displays disaster info and effects
	s.logf(disasterReport)
	s.islandDeplete(effects.CommonPoolMitigated) //island's resource will be depleted by disaster only when disaster happens and cp cannot fully mitigate

	return e, nil
}
