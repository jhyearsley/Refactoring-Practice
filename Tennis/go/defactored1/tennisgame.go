package tennisgame

type TennisGame struct {
	score1, score2           int
	player1Name, player2Name string
}

func NewTennisGame(player1Name, player2Name string) *TennisGame {
	game := &TennisGame{
		player1Name: player1Name,
		player2Name: player2Name,
	}

	return game
}

func (game *TennisGame) WonPoint(playerName string) {
	if playerName == "player1" {
		game.score1 += 1
	} else {
		game.score2 += 1
	}
}

func (game TennisGame) Score() string {
	score := ""
	tempScore := 0
	if game.score1 == game.score2 {
		switch game.score1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		case 3:
			score = "Forty-All"
		default:
			score = "Deuce"
		}
	} else if game.score1 >= 4 || game.score2 >= 4 {
		minusResult := game.score1 - game.score2
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.score1
			} else {
				score += "-"
				tempScore = game.score2
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
