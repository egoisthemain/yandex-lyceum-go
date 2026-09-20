package main

import (
	"sort"
	"strings"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + (float64(p.Assists) / 2)
	} else {
		p.Rating = (float64(p.Goals) + (float64(p.Assists) / 2.0)) / float64(p.Misses)
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	newPlayer := Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: 0}
	newPlayer.calculateRating()
	return newPlayer
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			iName := strings.ToLower(players[i].Name)
			jName := strings.ToLower(players[j].Name)
			return iName < jName
		} else {
			return players[i].Goals > players[j].Goals
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			iName := strings.ToLower(players[i].Name)
			jName := strings.ToLower(players[j].Name)
			return iName < jName
		} else {
			return players[i].Rating > players[j].Rating
		}
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		var iGM, jGM float64
		if players[i].Misses == 0 {
			iGM = 999999.0
		} else {
			iGM = float64(players[i].Goals / players[i].Misses)
		}

		if players[j].Misses == 0 {
			jGM = 999999.0
		} else {
			jGM = float64(players[j].Goals / players[j].Misses)
		}
		if iGM == jGM {
			iName := strings.ToLower(players[i].Name)
			jName := strings.ToLower(players[j].Name)
			return iName < jName
		} else {
			return iGM > jGM
		}
	})
	return players
}
