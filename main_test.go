package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestFunctionSum(t *testing.T) {
	data := []int{1, 1, 1, 3}
	expected := 6
	require.Equal(t, expected, Sum(data))
}

func TestFunctionRemove(t *testing.T) {
	data := []int{1, 1, 1, 3}
	expected := []int{1, 1, 3}
	require.Equal(t, expected, Remove(data, 2))
}

func TestFunctionFindIndexPosibility(t *testing.T) {
	data := []int{1, 1, 1, 3}
	expected := 2
	require.Equal(t, expected, FindIndexPosibility(data, 3))
}

func TestMaxPossibility(t *testing.T) {
	tests := []struct {
		name        string
		max_value   int
		possibility int
		expectedLog []string
		expected    int
	}{
		{
			name:        "Test 1",
			max_value:   6,
			possibility: 3,
			expectedLog: []string{
				"1,1,1,1,1,1 = 6",
				"1,1,1,1,2 = 6",
				"1,1,1,3 = 6",
				"1,2,3 = 6",
				"3,3 = 6",
			},
			expected: 5,
		},
		{
			name:        "Test 2",
			max_value:   10,
			possibility: 5,
			expectedLog: []string{
				"1,1,1,1,1,1,1,1,1,1 = 10",
				"1,1,1,1,1,1,1,1,2 = 10",
				"1,1,1,1,1,1,1,3 = 10",
				"1,1,1,1,1,1,4 = 10",
				"1,1,1,1,1,5 = 10",
				"1,1,1,2,5 = 10",
				"1,1,3,5 = 10",
				"1,4,5 = 10",
				"5,5 = 10",
			},
			expected: 9,
		},
		{
			name:        "Test 3",
			max_value:   21,
			possibility: 4,
			expectedLog: []string{
				"1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,3 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,4 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2,4 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,1,1,1,3,4 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,1,1,4,4 = 21",
				"1,1,1,1,1,1,1,1,1,1,1,2,4,4 = 21",
				"1,1,1,1,1,1,1,1,1,1,3,4,4 = 21",
				"1,1,1,1,1,1,1,1,1,4,4,4 = 21",
				"1,1,1,1,1,1,1,2,4,4,4 = 21",
				"1,1,1,1,1,1,3,4,4,4 = 21",
				"1,1,1,1,1,4,4,4,4 = 21",
				"1,1,1,2,4,4,4,4 = 21",
				"1,1,3,4,4,4,4 = 21",
				"1,4,4,4,4,4 = 21",
			},
			expected: 16,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			listData := []int{}
			index := test.max_value - 1
			var output int = 0
			for i := 0; i < test.max_value; i++ {
				listData = append(listData, 1)
			}
			for i := 0; i < test.max_value; i++ {
				if i != 0 {
					if index != 0 {
						if listData[index] < test.possibility {
							listData[index] = listData[index] + listData[index-1]
							listData = Remove(listData, index-1)
							index = FindIndexPosibility(listData, test.possibility)
						} else {
							index = index - 1
							listData[index] = listData[index] + listData[index-1]
							listData = Remove(listData, index-1)
							index = FindIndexPosibility(listData, test.possibility)
						}
					}
				}

				var printData string = ""
				for index, logData := range listData {
					s := strconv.Itoa(logData)
					if index == 0 {
						printData = printData + s
					} else {
						printData = printData + "," + s
					}
				}

				logPrint := fmt.Sprintf("%s = %d", printData, Sum(listData))
				require.Equal(t, test.expectedLog[i], logPrint)
				output = output + 1

				if index == 0 {
					break
				}
			}

			require.Equal(t, test.expected, output)
		})
	}
}
