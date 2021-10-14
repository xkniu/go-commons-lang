package commutil

import "testing"

func TestGetLocalIpv4(t *testing.T) {
	result1 := NetUtils.GetLocalIpv4()
	result2 := NetUtils.GetLocalIpv4()

	t.Logf("GetLocalIpv4: %v", result1)
	if result1 == "" || result1 != result2 {
		t.Errorf("case %d: expected: %v, got: %v", 0, result1, result2)
	}
}
