package tennisgame

type TennisGame struct {
	p1, p2   int
	p1N, p2N string
}

func NewTennisGame(p1N, p2N string) *TennisGame {
	game := &TennisGame{}
	game.p1N, game.p2N = p1N, p2N
	return game
}

func (game TennisGame) Score() string {
	s := ""
	if game.p1 < 4 && game.p2 < 4 {
		p := []string{"Love", "Fifteen", "Thirty", "Forty"}
		s = p[game.p1]
		if game.p1 == game.p2 {
			return s + "-All"
		} else {
			return s + "-" + p[game.p2]
		}
	} else {
		if game.p1 == game.p2 {
			return "Deuce"
		}
		if game.p1 > game.p2 {
			s = game.p1N
		} else {
			s = game.p2N
		}
		if (game.p1-game.p2)*(game.p1-game.p2) == 1 {
			return "Advantage " + s
		} else {
			return "Win for " + s
		}
	}
}

func (game *TennisGame) WonPoint(playerName string) {
	if playerName == "player1" {
		game.p1 += 1
	} else {
		game.p2 += 1
	}
}
