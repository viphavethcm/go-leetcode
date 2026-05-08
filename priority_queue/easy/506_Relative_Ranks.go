package easy

import (
	"strconv"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

// 10,3,8,9,4
// 10,9,8,4,3
func FindRelativeRanks(scores []int) []string {
	type person struct {
		score int
		index int
	}
	rs := make([]string, len(scores))
	rank := 1
	maxHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		personA := a.(person)
		personB := b.(person)
		return utils.IntComparator(personB.score, personA.score)
	})

	for idx, score := range scores {
		current := person{score: score, index: idx}
		maxHeap.Enqueue(current)
	}
	for !maxHeap.Empty() {
		currentPerson, _ := maxHeap.Dequeue()
		switch rank {
		case 1:
			rs[currentPerson.(person).index] = "Gold Medal"
		case 2:
			rs[currentPerson.(person).index] = "Silver Medal"
		case 3:
			rs[currentPerson.(person).index] = "Bronze Medal"
		default:
			rs[currentPerson.(person).index] = strconv.Itoa(rank)
		}
		rank++
	}
	return rs
}
