//You are given a m x n 2D grid initialized with these three possible values.
//
//-1 - A wall or an obstacle.
//0 - A gate.
//INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.
//Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.
//
//Example:
//
//Given the 2D grid:
//
//INF  -1   0  INF
//INF INF  INF  -1
//INF  -1  INF  -1
//0    -1  INF INF
//After running your function, the 2D grid should be:
//
//3  -1   0   1
//2   2   1  -1
//1  -1   2  -1
//0  -1   3   4

package queue

type Coordinate struct {
	x int
	y int
}

const INF = 2147483647

func isValid(i, j, x, y int) bool {
	if i < 0 || j < 0 || i >= x || j >= y {
		return false
	}
	return true
}

func wallsAndGates(rooms [][]int) {

	gates := []Coordinate{}
	x := len(rooms)

	if x == 0 {
		return
	}

	y := len(rooms[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if rooms[i][j] == 0 {
				gates = append(gates, Coordinate{x: i, y: j})
			}
		}
	}

	if len(gates) == 0 {
		return
	}

	q := []Coordinate{}
	incr := []Coordinate{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for i := 0; i < len(gates); i++ {
		gate := gates[i]
		q = append(q, gate)
		level := 0
		visited := make(map[Coordinate]bool)
		for len(q) > 0 {
			//fmt.Printf("Gates %v", q)
			l := len(q)
			level += 1
			for k := 0; k < l; k += 1 {
				item := q[0]
				q = q[1:]
				for j := range incr {
					newX := item.x + incr[j].x
					newY := item.y + incr[j].y
					if !isValid(newX, newY, x, y) {
						continue
					}
					if rooms[newX][newY] <= 0 {
						continue
					}
					newCoordinate := Coordinate{newX, newY}
					if _, ok := visited[newCoordinate]; ok {
						continue
					}
					visited[newCoordinate] = true

					if rooms[newX][newY] < level {
						continue
					}
					rooms[newX][newY] = level
					q = append(q, newCoordinate)
				}
			}
		}
		//fmt.Printf("gate: %v, distance: %v\n", gate, rooms)

	}
	return
}
