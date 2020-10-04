package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInMemoryCache(t *testing.T) {
	testCaseTable := []struct {
		inputData      string
		expectedResult interface{}
		expectedFound  bool
	}{
		{
			inputData:      "nonExistingKey",
			expectedResult: nil,
			expectedFound:  false,
		},
		{
			inputData:      "testKey",
			expectedResult: "testValue",
			expectedFound:  true,
		},
	}
	cacheClient := NewCache(5*time.Second, 5*time.Second)
	cacheClient.Set("testKey", "testValue")

	for _, testCase := range testCaseTable {
		data, found := cacheClient.Get(testCase.inputData)
		assert.Equal(t, testCase.expectedResult, data, "Actual data is different than expected one")
		assert.Equal(t, testCase.expectedFound, found, "Actual found boolean is different than expected one")
	}
}

func TestInMemoryCacheExpirationTime(t *testing.T) {
	testKey := "testKey"
	testValue := "testValue"
	cacheClient := NewCache(2*time.Second, 2*time.Second)
	cacheClient.Set(testKey, testValue)
	time.Sleep(3 * time.Second)
	data, found := cacheClient.Get(testKey)
	var expectedResult error // error is equal to nil by default
	expectedFound := false
	assert.Equal(t, expectedResult, data, "Actual data is different than expected one")
	assert.Equal(t, expectedFound, found, "Actual found boolean is different than expected one")
}
