// Copyright 2022 LitmusChaos Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package tokens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestValidateTokenExpiryDuration validates the negative and positive test cases for token expiry duration string input
func TestValidateTokenExpiryDuration(t *testing.T) {

	positiveTestCases := []string{"1d", "1D", "30d", "30D", "15d", "15D", "1h", "1H", "24h", "24H", "12h", "12H", "1m", "1M", "60m", "60M", "30m", "30M"}
	negativeTestCases := []string{"0123", "abcdf", "./^&*@$-as12", " ", "-31D", "100d", "500H", "-21h", "375m", "-9M", "abcM", "123em", "563sh", "ac45m"}

	for _, testCase := range positiveTestCases {

		dayHourMinuteChar, dayHourMinuteValue, err := validateTokenExpiryDuration(testCase)

		switch testCase {
		case "1d":
			assert.Equal(t, dayHourMinuteChar, 'd', "Expected %c but got %c", 'd', dayHourMinuteChar)
			assert.Equal(t, dayHourMinuteValue, 1, "Expected %d but got %d", 1, dayHourMinuteValue)
		case "30D":
			assert.Equal(t, dayHourMinuteChar, 'D', "Expected %c but got %c", 'D', dayHourMinuteChar)
			assert.Equal(t, dayHourMinuteValue, 30, "Expected %d but got %d", 30, dayHourMinuteValue)
		case "1h":
			assert.Equal(t, dayHourMinuteChar, 'h', "Expected %c but got %c", 'h', dayHourMinuteChar)
			assert.Equal(t, dayHourMinuteValue, 1, "Expected %d but got %d", 1, dayHourMinuteValue)
		case "12H":
			assert.Equal(t, dayHourMinuteChar, 'H', "Expected %c but got %c", 'H', dayHourMinuteChar)
			assert.Equal(t, dayHourMinuteValue, 12, "Expected %d but got %d", 12, dayHourMinuteValue)
		case "60m":
			assert.Equal(t, dayHourMinuteChar, 'm', "Expected %c but got %c", 'm', dayHourMinuteChar)
			assert.Equal(t, dayHourMinuteValue, 60, "Expected %d but got %d", 60, dayHourMinuteValue)
		case "30M":
			assert.Equal(t, dayHourMinuteChar, 'M', "Expected %c but got %c", 'M', dayHourMinuteChar)
			assert.Equal(t, dayHourMinuteValue, 30, "Expected %d but got %d", 30, dayHourMinuteValue)
		}

		assert.Nil(t, err, "Unxpected error for testcase: %s", testCase)
	}

	for _, testCase := range negativeTestCases {

		_, _, err := validateTokenExpiryDuration(testCase)

		assert.NotNil(t, err, "Expected error for testcase: %s", testCase)
	}
}
