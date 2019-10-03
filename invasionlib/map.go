package invasion

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	numTurns       = 10000
	CityNameLength = 16
	cityNameLength = 16
)

// NewMap creates a new map initialized with a source of randomness
func NewMap() *Map {
	return &Map{rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// NewMapFromFile returns a new instance of map based on the file
func NewMapFromFile(file string) (*Map, error) {
	// Instantiate the map
	m := NewMap()

	// Open the file and return any errors
	f, err := os.Open(file)
	if err != nil {
		return m, err
	}
	defer f.Close()

	// Scan the file by lines
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		city := &City{Roads: make([]*Road, 0)}
		// Split the lines on spaces
		l := strings.Split(strings.TrimSpace(scan.Text()), " ")
		// If the line is empty or too long ignore it
		if len(l) > 0 && len(l) < 6 && !strings.Contains(l[0], "=") {
			city.Name = l[0]
			l = l[1:]
			for _, rd := range l {
				// Split the roads on "="
				kv := strings.Split(rd, "=")
				city.Roads = append(city.Roads, &Road{Direction: kv[0], City: kv[1]})
			}
			m.Cities = append(m.Cities, city)
			continue
		}
	}
	return m, nil
}

// Map represents the game map. It contains an array of cities and some aliens
type Map struct {
	Cities []*City
	Aliens []*Alien

	rand *rand.Rand
}

// NewCities generates a random map to play this game on
func (m *Map) NewCities(n int) {
	cityNames := m.randStrings(n, cityNameLength)
	m.Cities = make([]*City, 0)
	for _, c := range cityNames {
		m.Cities = append(m.Cities, m.NewCity(c, cityNames))
	}
}

// Marshall returns the map in string format
func (m *Map) String() string {
	outBytes := make([]byte, 0, cityNameLength*5*len(m.Cities))
	for _, c := range m.Cities {
		outBytes = append(outBytes, []byte(c.String())...)
	}
	return string(outBytes)
}
