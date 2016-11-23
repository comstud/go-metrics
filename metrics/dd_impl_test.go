package metrics

import "testing"

func testMetricsAddr(t *testing.T, addr string, expected string) {
	cli, err := NewDDClient(addr)
	if err != nil {
		t.Fatalf("Failed to init with '%s': %s", addr, err)
	}
	res := cli.GetAddr()
	if res != expected {
		t.Fatalf("GetAddr mismatch for '%s': '%s' != '%s'", addr, res, expected)
	}
}

func TestMetricsAddr(t *testing.T) {
	testMetricsAddr(t, "dogstatsd:8125", "dogstatsd:8125")
	testMetricsAddr(t, "dogstatsd", "dogstatsd:8125")
	testMetricsAddr(t, "udp://dogstatsd:8125", "dogstatsd:8125")
	testMetricsAddr(t, "udp://dogstatsd", "dogstatsd:8125")
	testMetricsAddr(t, "", "127.0.0.1:8125")
	testMetricsAddr(t, "127.0.0.2:8126", "127.0.0.2:8126")
}
