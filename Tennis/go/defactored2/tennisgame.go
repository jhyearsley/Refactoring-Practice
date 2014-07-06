package tennisgame

type TennisGame struct {
	P1point, P2point         int
	P1res, P2res             string
	player1Name, player2Name string
}

func NewTennisGame(player1Name, player2Name string) *TennisGame {
	game := new(TennisGame)
	game.player1Name = player1Name
	game.player2Name = player2Name
	return game
}

func (game TennisGame) Score() string {
	score := ""

	if game.P1point == game.P2point && game.P1point < 4 {
		if game.P1point == 0 {
			score = "Love"
		}
		if game.P1point == 1 {
			score = "Fifteen"
		}
		if game.P1point == 2 {
			score = "Thirty"
		}
		if game.P1point == 3 {
			score = "Forty"
		}
		score += "-All"
	}

	if game.P1point == game.P2point && game.P1point > 3 {
		score = "Deuce"
	}

	if game.P1point > 0 && game.P2point == 0 {
		if game.P1point == 1 {
			game.P1res = "Fifteen"
		}
		if game.P1point == 2 {
			game.P1res = "Thirty"
		}
		if game.P1point == 3 {
			game.P1res = "Forty"
		}
		game.P2res = "Love"
		score = game.P1res + "-" + game.P2res
	}
	if game.P2point > 0 && game.P1point == 0 {
		if game.P2point == 1 {
			game.P2res = "Fifteen"
		}
		if game.P2point == 2 {
			game.P2res = "Thirty"
		}
		if game.P2point == 3 {
			game.P2res = "Forty"
		}
		game.P1res = "Love"
		score = game.P1res + "-" + game.P2res
	}

	if game.P1point > game.P2point && game.P1point < 4 {
		if game.P1point == 2 {
			game.P1res = "Thirty"
		}
		if game.P1point == 3 {
			game.P1res = "Forty"
		}
		if game.P2point == 1 {
			game.P2res = "Fifteen"
		}
		if game.P2point == 2 {
			game.P2res = "Thirty"
		}
		score = game.P1res + "-" + game.P2res
	}

	if game.P2point > game.P1point && game.P2point < 4 {
		if game.P2point == 2 {
			game.P2res = "Thirty"
		}
		if game.P2point == 3 {
			game.P2res = "Forty"
		}
		if game.P1point == 1 {
			game.P1res = "Fifteen"
		}
		if game.P1point == 2 {
			game.P1res = "Thirty"
		}
		score = game.P1res + "-" + game.P2res
	}

	if game.P1point > game.P2point && game.P2point >= 3 {
		score = "Advantage player1"
	}

	if game.P2point > game.P1point && game.P1point >= 3 {
		score = "Advantage player2"
	}

	if game.P1point >= 4 && game.P2point >= 0 && (game.P1point-game.P2point) >= 2 {
		score = "Win for player1"
	}

	if game.P2point >= 4 && game.P1point >= 0 && (game.P2point-game.P1point) >= 2 {
		score = "Win for player2"
	}

	return score
}

func (game *TennisGame) SetP1Score(number int) {
	for i := 0; i < number; i += 1 {
		game.P1Score()
	}
}

func (game *TennisGame) SetP2Score(number int) {
	for i := 0; i < number; i += 1 {
		game.P2Score()
	}
}

func (game *TennisGame) P1Score() {
	game.P1point += 1
}

func (game *TennisGame) P2Score() {
	game.P2point += 1
}

func (game *TennisGame) WonPoint(player string) {
	if player == "player1" {
		game.P1Score()
	} else {
		game.P2Score()
	}
}
