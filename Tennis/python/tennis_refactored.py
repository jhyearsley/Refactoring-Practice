class TennisGameRefactored:
    score_names = {0: "Love", 1: "Fifteen", 2: "Thirty", 3: "Forty"}

    def __init__(self, player1Name, player2Name):
        self.player1Name = player1Name
        self.player2Name = player2Name
        self.p1points = 0
        self.p2points = 0
        
    def won_point(self, playerName):
        if playerName == self.player1Name:
            self.p1points += 1
        else:
            self.p2points += 1
    
    def print_score(self):
        score_is_tied = self.check_for_tie()
        someone_won = self.check_for_winner()
        someone_has_advantage = self.check_for_advantage()
        
        if score_is_tied:
            result = self.is_tied()
            
        else:
            if someone_won:
                result = self.winner()
        
            elif someone_has_advantage:
                result = self.advantage()
            
            else:
                result = self.different_scores()   
                    
        return result
    
    def check_for_tie(self):
        if self.p1points == self.p2points:
            return True
        else:
            return False
    
    def is_tied(self):
        if 0 <= self.p1points and self.p2points < 4:
            return self.score_names[self.p1points] + '-All'    
        else:
            return "Deuce"
        
    def different_scores(self):
        p1_score = self.score_names[self.p1points]
        p2_score = self.score_names[self.p2points]
                
        return "%s-%s" % (p1_score, p2_score)
        
    def check_for_advantage(self):
        min_score = min(self.p1points, self.p2points)
        abs_diff_score = abs(self.p1points - self.p2points)
        
        if min_score >= 3 and abs_diff_score == 1:
            return True
        else:
            return False
        
    def advantage(self):
        if self.p1points > self.p2points:
            return "Advantage %s" % self.player1Name
        else:
            return "Advantage %s" % self.player2Name
            
    def check_for_winner(self):
        max_score = max(self.p1points, self.p2points)
        abs_diff_score = abs(self.p1points - self.p2points)
        
        if max_score > 3 and abs_diff_score != 1:
            return True
        else:
            return False
        
    def winner(self):
        if self.p1points > self.p2points:
            champ = self.player1Name                    
        else:
            champ = self.player2Name
            
        return 'Win for %s' % champ
            
TennisGame = TennisGameRefactored