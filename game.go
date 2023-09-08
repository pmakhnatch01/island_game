package game

type PathType int

const (
	Bridge PathType = iota
	FlightPath
)

type TransportationType int

const (
	Car TransportationType = iota
	HotAirBalloon
)

type Island struct {
	Name string
}

type GameMap struct {
	Connections map[*Island]map[*Island]PathType
}

type Player struct {
	CurrentIsland        *Island
	ChosenTransportation TransportationType
}
