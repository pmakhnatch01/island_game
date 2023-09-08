package islands

import "game"

func GenerateIslands() (*game.Island, *game.Island, []*game.Island, game.GameMap) {
    islandA := &game.Island{Name: "A"}
    islandB := &game.Island{Name: "B"}
    islandC := &game.Island{Name: "C"}
    islandD := &game.Island{Name: "D"}

    allIslands := []*game.Island{islandA, islandB, islandC, islandD}

    gameMap := game.GameMap{
        Connections: make(map[*game.Island]map[*game.Island]game.PathType),
    }

    gameMap.Connections[islandA] = map[*game.Island]game.PathType{
        islandB: game.Bridge,
        islandC: game.FlightPath,
    }
    gameMap.Connections[islandB] = map[*game.Island]game.PathType{
        islandD: game.FlightPath,
    }
    gameMap.Connections[islandC] = map[*game.Island]game.PathType{
        islandD: game.Bridge,
    }

    return islandA, islandD, allIslands, gameMap
}
