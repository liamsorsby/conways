package utils

func Evaluate(dataSet [][]bool) [][]bool {
	var dataset [][]bool

	for dataIndex, b := range dataSet {
		var set []bool
		var aliveCount int

		for index, state := range b {
			firstLength := len(dataSet)
			secondLength := len(dataSet[0])
			aliveCount = checkState(dataSet, dataIndex-1, index-1, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex-1, index, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex-1, index+1, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex, index-1, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex, index+1, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex+1, index-1, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex+1, index, firstLength, secondLength, aliveCount)
			aliveCount = checkState(dataSet, dataIndex+1, index+1, firstLength, secondLength, aliveCount)

			if state {
				if aliveCount < 2 || aliveCount > 3 {
					set = append(set, false)
				} else {
					set = append(set, true)
				}
			} else {
				if aliveCount == 3 {
					set = append(set, true)
				} else {
					set = append(set, false)
				}
			}
		}

		aliveCount = 0

		dataset = append(dataset, set)
	}

	return dataset
}

func checkState(dataset [][]bool, firstIndex int, secondIndex int, firstLength, secondLength, aliveCount int) int {
	if firstIndex >= firstLength || firstIndex < 0 || secondIndex >= secondLength || secondIndex < 0 {
		return aliveCount
	}

	if dataset[firstIndex][secondIndex] {
		return aliveCount + 1
	}

	return aliveCount
}
