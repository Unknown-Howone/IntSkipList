package IntSkipList

import "testing"

func TestSkipList(t *testing.T) {
	sl := New()
	sl.Add(30)
	sl.Add(40)
	sl.Add(50)
	sl.Add(45)
	sl.Add(60)
	sl.Add(70)
	sl.Add(80)
	sl.Add(90)

	sl.Search(60)
	sl.Erase(60)
	sl.Search(60)
}
