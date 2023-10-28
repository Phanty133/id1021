package trains

import "errors"

const MOD_VAL int = 541

type City struct {
	Name      string
	Neighbors []Connection
}

type Connection struct {
	City *City
	Dist int // In minutes
}

type NetworkGraph struct {
	Cities    [][]*City
	NumCities int
}

func NewNetworkGraph() NetworkGraph {
	return NetworkGraph{make([][]*City, MOD_VAL), 0}
}

func (n *NetworkGraph) GetCities() []string {
	cities := make([]string, 0)

	for _, bucket := range n.Cities {
		if bucket == nil {
			continue
		}

		for _, c := range bucket {
			cities = append(cities, c.Name)
		}
	}

	return cities
}

func (c *City) Connect(nxt *City, dst int) {
	c.Neighbors = append(c.Neighbors, Connection{nxt, dst})
}

func getCityHash(name string) int {
	hash := 0

	for _, c := range name {
		hash = (hash * 31 % MOD_VAL) + int(c)
	}

	return hash % MOD_VAL
}

func (n *NetworkGraph) AddCity(c *City) {
	cityIdx := getCityHash(c.Name)
	bucket := n.Cities[cityIdx]

	if bucket == nil {
		n.Cities[cityIdx] = make([]*City, 0)
	}

	n.Cities[cityIdx] = append(n.Cities[cityIdx], c)
	n.NumCities++
}

func (n *NetworkGraph) LookupCity(name string) *City {
	cityIdx := getCityHash(name)
	bucket := n.Cities[cityIdx]

	if bucket == nil {
		return nil
	}

	for _, c := range bucket {
		if c.Name == name {
			return c
		}
	}

	return nil
}

func (c *City) IsConnected(to *City) bool {
	for _, n := range c.Neighbors {
		if n.City == to {
			return true
		}
	}

	return false
}

func (n *NetworkGraph) FillData(from []string, to []string, dist []int) {
	for i := 0; i < len(from); i++ {
		fromCity := n.LookupCity(from[i])
		toCity := n.LookupCity(to[i])

		if fromCity == nil {
			fromCity = &City{from[i], nil}
			n.AddCity(fromCity)
		}

		if toCity == nil {
			toCity = &City{to[i], nil}
			n.AddCity(toCity)
		}

		// Skip duplicate connections just in case
		if fromCity.IsConnected(toCity) {
			continue
		}

		fromCity.Connect(toCity, dist[i])
		toCity.Connect(fromCity, dist[i])
	}
}

func (n *NetworkGraph) CountBucketCollisions() (int, map[int]int) {
	collisions := 0
	bucketSizes := make(map[int]int, 0)

	for i, bucket := range n.Cities {
		if bucket == nil {
			continue
		}

		if len(bucket) > 1 {
			collisions++
		}

		bucketSizes[i] = len(bucket)
	}

	return collisions, bucketSizes
}

func (n *NetworkGraph) DepthFirstDistance(from string, to string, maxDistance int) (int, error) {
	fromCity := n.LookupCity(from)
	toCity := n.LookupCity(to)

	if fromCity == nil || toCity == nil {
		return 0, errors.New("city not found")
	}

	return fromCity.shortestDistance(toCity, maxDistance)
}

func (c *City) shortestDistance(to *City, maxDistance int) (int, error) {
	if maxDistance < 0 {
		return 0, errors.New("max distance exceeded")
	}

	if c == to {
		return 0, nil
	}

	minDist := -1

	for _, con := range c.Neighbors {
		dist, err := con.City.shortestDistance(to, maxDistance-con.Dist)

		if err != nil {
			continue
		}

		if minDist == -1 || dist+con.Dist < minDist {
			minDist = dist + con.Dist
		}
	}

	if minDist == -1 {
		return 0, errors.New("no path found")
	}

	return minDist, nil
}

func (n *NetworkGraph) DepthFirstDistanceNoLoop(from string, to string) (int, error) {
	fromCity := n.LookupCity(from)
	toCity := n.LookupCity(to)

	if fromCity == nil || toCity == nil {
		return 0, errors.New("city not found")
	}

	visited := make([]string, n.NumCities)

	return fromCity.shortestDistanceNoLoop(toCity, visited)
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func (c *City) shortestDistanceNoLoop(to *City, visited []string) (int, error) {
	if c == to {
		return 0, nil
	}

	minDist := -1

	for _, con := range c.Neighbors {
		if contains(visited, con.City.Name) {
			continue
		}

		dist, err := con.City.shortestDistanceNoLoop(to, append(visited, c.Name))

		if err != nil {
			continue
		}

		if minDist == -1 || dist+con.Dist < minDist {
			minDist = dist + con.Dist
		}
	}

	if minDist == -1 {
		return 0, errors.New("no path found")
	}

	return minDist, nil
}

func (n *NetworkGraph) DepthFirstDistanceNoLoopMaxed(from string, to string) (int, []string, error) {
	fromCity := n.LookupCity(from)
	toCity := n.LookupCity(to)

	if fromCity == nil || toCity == nil {
		return 0, nil, errors.New("city not found")
	}

	path := make([]string, 0, n.NumCities)

	return fromCity.shortestDistanceNoLoopMaxed(toCity, path, -1, -1)
}

func (c *City) shortestDistanceNoLoopMaxed(
	to *City,
	path []string,
	closestDistFound int,
	distLeft int,
) (int, []string, error) {
	newPath := make([]string, len(path)+1)
	copy(newPath, path)
	newPath[len(path)] = c.Name

	if closestDistFound != -1 && distLeft <= 0 {
		return 0, nil, errors.New("max distance exceeded")
	}

	if c == to {
		return 0, newPath, nil
	}

	minDist := -1
	var minPath []string = nil

	for _, con := range c.Neighbors {
		if contains(path, con.City.Name) {
			continue
		}

		dist, path, err := con.City.shortestDistanceNoLoopMaxed(
			to,
			newPath,
			closestDistFound,
			distLeft-con.Dist,
		)

		if err != nil {
			continue
		}

		if minDist == -1 || dist+con.Dist < minDist {
			minDist = dist + con.Dist
			minPath = path

			if closestDistFound == -1 || minDist < closestDistFound {
				closestDistFound = minDist
				distLeft = closestDistFound
			}
		}
	}

	if minDist == -1 {
		return 0, nil, errors.New("no path found")
	}

	return minDist, minPath, nil
}
