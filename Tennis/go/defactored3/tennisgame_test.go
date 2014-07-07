package tennisgame

import (
	"math"
	"testing"
)

var testAllScoresData = []struct {
	player1Score  int
	player2Score  int
	expectedScore string
}{
	{0, 0, "Love-All"},
	{1, 1, "Fifteen-All"},
	{2, 2, "Thirty-All"},
	{3, 3, "Forty-All"},
	{4, 4, "Deuce"},

	{1, 0, "Fifteen-Love"},
	{0, 1, "Love-Fifteen"},
	{2, 0, "Thirty-Love"},
	{0, 2, "Love-Thirty"},
	{3, 0, "Forty-Love"},
	{0, 3, "Love-Forty"},
	{4, 0, "Win for player1"},
	{0, 4, "Win for player2"},

	{2, 1, "Thirty-Fifteen"},
	{1, 2, "Fifteen-Thirty"},
	{3, 1, "Forty-Fifteen"},
	{1, 3, "Fifteen-Forty"},
	{4, 1, "Win for player1"},
	{1, 4, "Win for player2"},

	{3, 2, "Forty-Thirty"},
	{2, 3, "Thirty-Forty"},
	{4, 2, "Win for player1"},
	{2, 4, "Win for player2"},

	{4, 3, "Advantage player1"},
	{3, 4, "Advantage player2"},
	{5, 4, "Advantage player1"},
	{4, 5, "Advantage player2"},
	{15, 14, "Advantage player1"},
	{14, 15, "Advantage player2"},

	{6, 4, "Win for player1"},
	{4, 6, "Win for player2"},
	{16, 14, "Win for player1"},
	{14, 16, "Win for player2"},
}

func TestAllScores(t *testing.T) {

	for _, test := range testAllScoresData {
		game := NewTennisGame("player1", "player2")

		highestScore := int(math.Max(float64(test.player1Score), float64(test.player2Score)))

		for i := 0; i < highestScore; i += 1 {
			if i < test.player1Score {
				game.WonPoint("player1")
			}
			if i < test.player2Score {
				game.WonPoint("player2")
			}
		}

		if test.expectedScore != game.Score() {
			t.Fatalf("For [%v:%v], expected score to be %v, but got %v\n", test.player1Score, test.player2Score, test.expectedScore, game.Score())
		}
	}

}

func TestRealisticGame(t *testing.T) {
	game := NewTennisGame("player1", "player2")
	points := []string{"player1", "player1", "player2", "player2", "player1", "player1"}
	expectedScores := []string{"Fifteen-Love", "Thirty-Love", "Thirty-Fifteen", "Thirty-All", "Forty-Thirty", "Win for player1"}
	for i := 0; i < 6; i += 1 {
		game.WonPoint(points[i])
		if expectedScores[i] != game.Score() {
			t.Fatalf("For iteration %v, expected score to be %v, but got %v\n", i, expectedScores[i], game.Score())
		}
	}
}
