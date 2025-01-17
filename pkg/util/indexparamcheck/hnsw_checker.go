package indexparamcheck

type hnswChecker struct {
	floatVectorBaseChecker
}

func (c hnswChecker) StaticCheck(params map[string]string) error {
	if !CheckIntByRange(params, EFConstruction, HNSWMinEfConstruction, HNSWMaxEfConstruction) {
		return errOutOfRange(EFConstruction, HNSWMinEfConstruction, HNSWMaxEfConstruction)
	}

	if !CheckIntByRange(params, HNSWM, HNSWMinM, HNSWMaxM) {
		return errOutOfRange(HNSWM, HNSWMinM, HNSWMaxM)
	}

	return c.floatVectorBaseChecker.staticCheck(params)
}

func (c hnswChecker) CheckTrain(params map[string]string) error {
	if err := c.StaticCheck(params); err != nil {
		return err
	}
	return c.floatVectorBaseChecker.CheckTrain(params)
}

func newHnswChecker() IndexChecker {
	return &hnswChecker{}
}
