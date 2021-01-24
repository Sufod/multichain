package bitcoinz_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBitcoinZ(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "bitcoinZ Suite")
}
