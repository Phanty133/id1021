package trains

import (
	"errors"
	"fmt"

	"github.com/phanty133/id1021/13-dijkstra/pkg/pqueue"
)

type PathNode struct {
	city *City
	prev *City
	dist int
	qIdx int
}

func (n *NetworkGraph) DijkstraFind(from, to string) (int, []string, int, error) {
	if from == to {
		return 0, []string{from}, 0, nil
	}

	fromCity := n.LookupCity(from)
	toCity := n.LookupCity(to)

	if fromCity == nil {
		return 0, nil, 0, fmt.Errorf("city %s not found", from)
	}

	if toCity == nil {
		return 0, nil, 0, fmt.Errorf("city %s not found", to)
	}

	pq := pqueue.NewArrHeap[*PathNode](n.NumCities)
	cityNodes := make([]*PathNode, n.NumCities)

	pq.Add(&PathNode{fromCity, nil, 0, 0}, 0)

	for !pq.Empty() {
		cur, _ := pq.Remove()

		if cur.city == toCity {
			// Assemble the shortest past and return
			path := make([]string, 0)
			city := cur.city

			for city != nil {
				path = append(path, city.Name)
				city = cityNodes[city.idx].prev
			}

			cityNodesNum := 0

			for _, n := range cityNodes {
				if n != nil {
					cityNodesNum++
				}
			}

			return cur.dist, path, cityNodesNum, nil
		}

		cityNodes[cur.city.idx] = cur

		for _, n := range cur.city.Neighbors {
			node := cityNodes[n.City.idx]

			if node == nil {
				node = &PathNode{n.City, cur.city, n.Dist + cur.dist, 0}
				cityNodes[n.City.idx] = node
				pq.Add(node, node.dist)
			}

			if node.dist > n.Dist+cur.dist {
				node.dist = n.Dist + cur.dist
				node.prev = cur.city
				pq.Update(node, node.dist)
			}
		}
	}

	return 0, nil, 0, errors.New("no path found")
}
