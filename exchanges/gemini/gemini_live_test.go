//+build mock_test_off

// This will build if build tag mock_test_off is parsed and will do live testing
// using all tests in (exchange)_test.go
package gemini

import (
	"log"
	"os"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/exchanges/sharedtestvalues"
)

var mockTests = false

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	err := cfg.LoadConfig("../../testdata/configtest.json", true)
	if err != nil {
		log.Fatal("Test Failed - Gemini load config error", err)
	}
	geminiConfig, err := cfg.GetExchangeConfig("Gemini")
	if err != nil {
		log.Fatal("Test Failed - Gemini Setup() init error", err)
	}
	geminiConfig.API.AuthenticatedSupport = true
	geminiConfig.API.Credentials.Key = apiKey
	geminiConfig.API.Credentials.Secret = apiSecret
	g.SetDefaults()
	err = g.Setup(geminiConfig)
	if err != nil {
		log.Fatal("Test Failed - Gemini setup error", err)
	}
	g.API.Endpoints.URL = geminiSandboxAPIURL
	log.Printf(sharedtestvalues.LiveTesting, g.GetName(), g.API.Endpoints.URL)
	os.Exit(m.Run())
}
