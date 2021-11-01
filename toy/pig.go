package toy

import (
    "fmt"
    "math/rand"
)

const (
    win            = 100
    gamesPerSeries = 10
)

type score struct {
    player, oppoent, thisTurn int
}

type action func(current score) (result score, turnIsOver bool)

func roll(s score) (score, bool) {
    outcome := rand.Intn(6) + 1
    if outcome == 1 {
        return score{s.oppoent, s.player, 0}, false
    }

    return score{s.player, s.oppoent, s.thisTurn + outcome}, true
}

func stay(s score) (score, bool) {
    return score{s.oppoent, s.player + s.thisTurn, 0}, true
}

type stragety func(score) action

func stayAtK(k int) stragety {
    return func(s score) action {
        if s.thisTurn > k {
            return stay
        }
        return roll
    }
}

func play(stragety0, stragety1 stragety) int {
    strageties := []stragety{stragety0, stragety1}
    var s score
    var turnIsOver bool
    currentPlayer := rand.Intn(2)
    for s.player+s.thisTurn < win {
        action := strageties[currentPlayer](s)
        s, turnIsOver = action(s)
        if turnIsOver {
            currentPlayer = (currentPlayer + 1) % 2
        }
    }

    return currentPlayer
}

func roundRobin(strageties []stragety) ([]int, int) {
    wins := make([]int, len(strageties))
    for i := 0; i < len(strageties); i++ {
        for j := i + 1; j < len(strageties); j++ {
            for k := 0; k < gamesPerSeries; k++ {
                winner := play(strageties[i], strageties[j])
                if winner == 0 {
                    wins[i]++
                } else {
                    wins[j]++
                }
            }
        }
    }
    gamesPerStragety := gamesPerSeries * (len(strageties) - 1)

    return wins, gamesPerStragety
}

func ratioString(vals ...int) string {
    total := 0
    for _, val := range vals {
        total += val
    }

    var s string

    for _, val := range vals {
        if s != "" {
            s += ", "
        }
        pct := 100 * float64(val) / float64(total)
        s += fmt.Sprintf("%d / %d (%0.1f %%)", val, total, pct)
    }

    return s
}

func StartPigGame() {
    strageties := make([]stragety, win)
    for k := range strageties {
        strageties[k] = stayAtK(k + 1)
    }
    wins, games := roundRobin(strageties)
    for k:= range strageties {
        fmt.Printf("Wins, losses staying at k = % 4d: %s\n", k, ratioString(wins[k], games - wins[k]))
    }
}

