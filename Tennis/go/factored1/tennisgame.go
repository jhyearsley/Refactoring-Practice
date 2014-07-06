package tennisgame

var POINTS = []string{"Love", "Fifteen", "Thirty", "Forty"}

type TennisGame struct {
	player1Score, player2Score int
	player1Name, player2Name   string
}

func NewTennisGame(player1Name, player2Name string) *TennisGame {
	game := new(TennisGame)
	game.player1Name, game.player2Name = player1Name, player2Name
	return game
}

func (game TennisGame) Score() string {
	if game.someoneHasWon() {
		return "Win for " + game.winningPlayerName()
	}

	if game.isEndgame() {
		if game.pointsAreEven() {
			return "Deuce"
		} else {
			return "Advantage " + game.winningPlayerName()
		}
	} else {
		if game.pointsAreEven() {
			return POINTS[game.player1Score] + "-All"
		} else {
			return POINTS[game.player1Score] + "-" + POINTS[game.player2Score]
		}
	}
}

func (game TennisGame) someoneHasWon() bool {
	return game.isEndgame() && (game.player1Score-game.player2Score >= 2 || game.player2Score-game.player1Score >= 2)
}

func (game TennisGame) pointsAreEven() bool {
	return game.player1Score == game.player2Score
}

func (game TennisGame) isEndgame() bool {
	return game.player1Score > 3 || game.player2Score > 3
}

func (game TennisGame) winningPlayerName() string {
	if game.player1Score > game.player2Score {
		return game.player1Name
	} else {
		return game.player2Name
	}
}

func (game *TennisGame) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.player1Score += 1
	} else {
		game.player2Score += 1
	}
}
