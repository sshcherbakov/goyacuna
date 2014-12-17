package goyacuna

import (
	"testing"
)

func TestToMap(t *testing.T) {

	dcr := &DealCountRequest{WalletAccountId: "xxx"}
	mapss := *toStringMap(dcr)

	if len(mapss) != 1 {
		t.Log(mapss)
		t.Fail()
	}

}
