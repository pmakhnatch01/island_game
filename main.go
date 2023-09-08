package main

import (
    "fmt"
    "game"
    "islands"
    "userinterface"
)

func main() {
    startingIsland, endingIsland, allIslands, gameMap := islands.GenerateIslands()

    player1 := &game.Player{CurrentIsland: startingIsland, ChosenTransportation: game.Car}
    player2 := &game.Player{CurrentIsland: startingIsland, ChosenTransportation: game.HotAirBalloon}

    for {
        if player1.CurrentIsland == endingIsland || player2.CurrentIsland == endingIsland {
            break
        }

        player1Move := userinterface.GetUserMove(player1, allIslands)
        player2Move := userinterface.GetUserMove(player2, allIslands)

        pathType, exists := gameMap.Connections[player1.CurrentIsland][player1Move]
        if exists && (pathType == game.Bridge && player1.ChosenTransportation == game.Car || pathType == game.FlightPath && player1.ChosenTransportation == game.HotAirBalloon) {
            player1.CurrentIsland = player1Move
        }

        pathType, exists = gameMap.Connections[player2.CurrentIsland][player2Move]
        if exists && (pathType == game.Bridge && player2.ChosenTransportation == game.Car || pathType == game.FlightPath && player2.ChosenTransportation == game.HotAirBalloon) {
            player2.CurrentIsland = player2Move
        }
    }

    if player1.CurrentIsland == endingIsland {
        fmt.Println("Player 1 wins!")
    } else {
        fmt.Println("Player 2 wins!")
    }
}
