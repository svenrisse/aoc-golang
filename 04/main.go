package day4

func solveA(g Grid) int {
	count := 0

	for _, p := range g.allPos('X') {
		count += g.countMAS(p)
	}

	return count
}

type Coord struct {
	l, r int
}

func (c Coord) add(dir Coord) Coord {
	return Coord{c.l + dir.l, c.r + dir.r}
}

type Grid []string

func (g Grid) allPos(lookingFor rune) []Coord {
	result := make([]Coord, 0)
	for l, line := range g {
		for r, rune := range line {
			if rune == lookingFor {
				result = append(result, Coord{l, r})
			}
		}
	}

	return result
}

func (g Grid) countMAS(pos Coord) int {
	allDirections := []Coord{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	result := 0
	for _, dir := range allDirections {
		mas := g.wordFromPos(pos, dir, 3)
		if mas == "MAS" {
			result += 1
		}
	}
	return result
}

func (g Grid) wordFromPos(pos Coord, dir Coord, length int) string {
	res := ""
	for range make([]int, length) {
		pos = pos.add(dir)
		rune := g.runeAtPos(pos)
		res += string(rune)
	}
	return res
}

func (g Grid) runeAtPos(pos Coord) rune {
	if pos.l < 0 || pos.r < 0 || pos.l >= len(g) || pos.r >= len(g[0]) {
		return ' '
	}
	return []rune(g[pos.l])[pos.r]
}
