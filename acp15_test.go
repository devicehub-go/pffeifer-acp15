package pffeiferacp15_test

import (
	"fmt"
	"testing"

	pffeiferacp15 "github.com/devicehub-go/pffeifer-acp15"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestMain(t *testing.T) {
	pump := pffeiferacp15.New(unicomm.Options{
		Protocol: unicomm.Serial,
		TCP: unicommtcp.TCPOptions{
			Host: "10.0.3.55",
			Port: 4004,
		},
	})
	if err := pump.Connect(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(pump.GetIDN(0))
}
