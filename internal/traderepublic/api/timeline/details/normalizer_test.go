package details_test

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/details"
	"github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes"
	_ "github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes/transaction"
)

func TestItNormalizesSupportTransactionDetails(t *testing.T) {
	t.Parallel()

	logger := log.New()
	logger.Out = io.Discard
	normalizer := details.NewTransactionResponseNormalizer(logger)

	for testCaseName, testCase := range fakes.TransactionTestCasesSupported {
		if !testCase.Enabled {
			continue
		}

		var response details.Response

		err := json.Unmarshal([]byte(testCase.Source.JSON), &response)

		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		actual, err := normalizer.Normalize(response)

		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))
		assert.Equal(t, testCase.Result, actual, fmt.Sprintf("case '%s'", testCaseName))
	}
}
