package main

type gpath struct {
	BoundDepth int
	Bound      map[int]int
	Tail       map[int]int
	LastIndx   map[int]int
	LastRec    map[int]int
	Bin        []map[int]int
}

func new_gpath(bound []int) {
	return
}

func (gp gpath) add(new int) bool {
	for i, bound := range gp.Bound {
		if new == gp.Tail[i] && gp.LastRec[i] >= 0 {
			new = gp.LastIndx[i] - gp.LastRec[i]
			if _, old := gp.Bin[i][new]; !old {
				if len(gp.Bin[i]) == bound {
					return false
				}
				gp.Bin[i][new] = 0
			}
			gp.LastRec[i] = gp.LastIndx[i]
			gp.LastIndx[i] += 1
			break
		}
	}
	return true
}
