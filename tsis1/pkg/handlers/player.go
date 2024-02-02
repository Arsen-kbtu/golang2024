package pkg

type Player struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Number    int    `json:"number"`
	Position  string `json:"position"`
}

var team = []Player{
	{FirstName: "Aaron", LastName: "Ramsdale", Age: 25, Number: 1, Position: "Goalkeeper"},
	{FirstName: "William", LastName: "Saliba", Age: 22, Number: 2, Position: "Defender"},
	{FirstName: "Ben", LastName: "White", Age: 26, Number: 4, Position: "Defender"},
	{FirstName: "Thomas", LastName: "Partey", Age: 30, Number: 5, Position: "Midfielder"},
	{FirstName: "Gabriel", LastName: "Magalhaes", Age: 26, Number: 6, Position: "Defender"},
	{FirstName: "Bukayo", LastName: "Saka", Age: 23, Number: 7, Position: "Winger"},
	{FirstName: "Martin", LastName: "Odegaard", Age: 25, Number: 8, Position: "Midfielder"},
	{FirstName: "Gabriel", LastName: "Jesus", Age: 26, Number: 9, Position: "Striker"},
	{FirstName: "Emil", LastName: "Smith-Rowe", Age: 23, Number: 10, Position: "Midfielder"},
	{FirstName: "Gabriel", LastName: "Martinelli", Age: 22, Number: 11, Position: "Winger"},
	{FirstName: "Jurrien", LastName: "Timber", Age: 22, Number: 12, Position: "Defender"},
	{FirstName: "Eddie", LastName: "Nketiah", Age: 24, Number: 14, Position: "Striker"},
	{FirstName: "Jakub", LastName: "Kiwior", Age: 23, Number: 15, Position: "Defender"},
	{FirstName: "Takehiro", LastName: "Tomiasu", Age: 25, Number: 18, Position: "Defender"},
	{FirstName: "Leandro", LastName: "Trossard", Age: 29, Number: 19, Position: "Winger"},
	{FirstName: "?", LastName: "Jorginho", Age: 32, Number: 20, Position: "Midfielder"},
	{FirstName: "Fabio", LastName: "Vieira", Age: 23, Number: 21, Position: "Midfielder"},
	{FirstName: "David", LastName: "Raya", Age: 28, Number: 22, Position: "Goalkeeper"},
	{FirstName: "Reiss", LastName: "Nelson", Age: 24, Number: 24, Position: "Winger"},
	{FirstName: "Kai", LastName: "Havertz", Age: 24, Number: 29, Position: "Midfielder"},
	{FirstName: "Olexandr", LastName: "Zinchenko", Age: 27, Number: 35, Position: "Defender"},
	{FirstName: "Declan", LastName: "Rice", Age: 25, Number: 41, Position: "Midfielder"},
}

func FindByNum(num int) (Player, bool) {
	for i := 0; i < len(team); i++ {
		if team[i].Number == num {
			return team[i], true
		}
	}
	return Player{}, false
}
func GetTeam() []Player {
	return team
}
