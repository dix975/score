package match

import "dix975.com/score/team"

type Set struct {
	HomeTeamScore int
	AwayTeamScore int
}

type Match struct {

	HomeTeam  team.Team
	AwayTeam team.Team
	Sets Set

}