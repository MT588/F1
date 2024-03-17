package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ernestosuarez/itertools"
)

type Driver struct {
	Name  string
	Price float64
	Race1 int
	Race2 int
}

type Constructor struct {
	Name  string
	Price float64
	Race1 int
	Race2 int
}

type Team struct {
	Driver1      Driver
	Driver2      Driver
	Driver3      Driver
	Driver4      Driver
	Driver5      Driver
	Constructor1 Constructor
	Constructor2 Constructor
	Price        float64
	Score        int
}

func (d *Driver) String() (res string) {
	res += d.Name
	return
}

func (c *Constructor) String() (res string) {
	res += c.Name
	return
}

func (t *Team) String() (res string) {
	res += "The drivers in the best Team are: " + t.Driver1.String() + ", " + t.Driver2.String() + ", " + t.Driver3.String() + ", " + t.Driver4.String() + ", " + t.Driver5.String() + "\n"
	res += "The constructors in the best Team are: " + t.Constructor1.Name + " and " + t.Constructor2.Name + "\n"
	res += fmt.Sprintf("Score a total of %d points", t.Score) + "\n"
	res += fmt.Sprintf("Costing a total of %.2f dollars", t.Price)
	return
}

var DriverData = map[string]Driver{
	"HAM": {Name: "HAM", Price: 19.5, Race1: 12, Race2: 6},
	"NOR": {Name: "NOR", Price: 23.2, Race1: 16, Race2: 8},
	"PIA": {Name: "PIA", Price: 19.4, Race1: 10, Race2: 23},
	"ALO": {Name: "ALO", Price: 16.2, Race1: 7, Race2: 16},
	"STR": {Name: "STR", Price: 11.0, Race1: 8, Race2: -17},
	"VER": {Name: "VER", Price: 30.4, Race1: 45, Race2: 36},
	"ALB": {Name: "ALB", Price: 7.3, Race1: 0, Race2: 6},
	"SAR": {Name: "SAR", Price: 5.9, Race1: 3, Race2: 7},
	"RIC": {Name: "RIC", Price: 8.7, Race1: 5, Race2: 0},
	"TSU": {Name: "TSU", Price: 7.6, Race1: -1, Race2: -1},
	"BOT": {Name: "BOT", Price: 6.0, Race1: 0, Race2: -1},
	"ZHO": {Name: "ZHO", Price: 6.9, Race1: 11, Race2: -2},
	"HUL": {Name: "HUL", Price: 6.7, Race1: -3, Race2: 9},
	"MAG": {Name: "MAG", Price: 7.2, Race1: 7, Race2: 7},
	"OCO": {Name: "OCO", Price: 8.8, Race1: 7, Race2: 8},
	"GAS": {Name: "GAS", Price: 7.5, Race1: 6, Race2: -20},
	"PER": {Name: "PER", Price: 22.1, Race1: 31, Race2: 31},
	"LEC": {Name: "LEC", Price: 20.4, Race1: 22, Race2: 37},
	"SAI": {Name: "SAI", Price: 18.8, Race1: 36, Race2: 0},
	"RUS": {Name: "RUS", Price: 19.2, Race1: 20, Race2: 15},
}
var ConstructorData = map[string]Constructor{
	"Red Bull Racing": {Name: "Red Bull Racing", Price: 28.2, Race1: 89, Race2: 90},
	"Ferrari":         {Name: "Ferrari", Price: 19.9, Race1: 73, Race2: 58},
	"Mercedes":        {Name: "Mercedes", Price: 20.3, Race1: 42, Race2: 36},
	"Mclaren":         {Name: "Mclaren", Price: 23.6, Race1: 36, Race2: 41},
	"Aston Martin":    {Name: "Aston Martin", Price: 14, Race1: 20, Race2: 9},
	"Haas F1 Team":    {Name: "Haas F1 Team", Price: 6.7, Race1: 9, Race2: 19},
	"Williams":        {Name: "Williams", Price: 6.7, Race1: 4, Race2: 14},
	"RB":              {Name: "RB", Price: 8.3, Race1: 7, Race2: 4},
	"Kick Sauber":     {Name: "Kick Sauber", Price: 6.3, Race1: 10, Race2: -4},
	"Alpine":          {Name: "Alpine", Price: 8.1, Race1: 12, Race2: -13},
}

func findBestTeam() Team {
	bestTeam := Team{}
	bestTeamList := make([]interface{}, 0)
	driverNames := make(itertools.List, 0, len(DriverData))
	for name := range DriverData {
		driverNames = append(driverNames, name)
	}
	constructorNames := make(itertools.List, 0, len(ConstructorData))
	for name := range ConstructorData {
		constructorNames = append(constructorNames, name)
	}

	for combo := range itertools.CombinationsList(driverNames, 5) {
		for combo2 := range itertools.CombinationsList(constructorNames, 2) {
			currentTeamPrice := 0.0
			currentTeamPoints := 0
			biggestDriver1 := 0
			biggestDriver2 := 0

			for _, driver := range combo {
				if DriverData[driver.(string)].Race1 > biggestDriver1 {
					biggestDriver1 = DriverData[driver.(string)].Race1
				}
				if DriverData[driver.(string)].Race2 > biggestDriver2 {
					biggestDriver2 = DriverData[driver.(string)].Race2
				}
				currentTeamPrice += DriverData[driver.(string)].Price
				currentTeamPoints += DriverData[driver.(string)].Race1 + DriverData[driver.(string)].Race2
			}

			currentTeamPoints += biggestDriver1 + biggestDriver2

			for _, constructor := range combo2 {
				currentTeamPrice += ConstructorData[constructor.(string)].Price
				currentTeamPoints += ConstructorData[constructor.(string)].Race1 + ConstructorData[constructor.(string)].Race2
			}

			if currentTeamPoints > bestTeam.Score && currentTeamPrice <= 100 {
				bestTeamList = append(combo, combo2)
				bestTeam.Price = currentTeamPrice
				bestTeam.Score = currentTeamPoints
				//fmt.Println(bestTeamList, bestScore, bestTeamPrice)
			}
		}
	}
	bestTeam.Driver1 = DriverData[bestTeamList[0].(string)]
	bestTeam.Driver2 = DriverData[bestTeamList[1].(string)]
	bestTeam.Driver3 = DriverData[bestTeamList[2].(string)]
	bestTeam.Driver4 = DriverData[bestTeamList[3].(string)]
	bestTeam.Driver5 = DriverData[bestTeamList[4].(string)]
	bestTeam.Constructor1 = ConstructorData[bestTeamList[5].(itertools.List)[0].(string)]
	bestTeam.Constructor2 = ConstructorData[bestTeamList[5].(itertools.List)[1].(string)]
	return bestTeam
}

func main() {
	best := findBestTeam()
	fmt.Println(best.String())

	url := "https://api.formula1.com/6657193977244c13?d=account.formula1.com"
	method := "POST"
  
	payload := strings.NewReader(`{"solution":{"interrogation":{"st":162229509,"sr":1959639815,"cr":78830557},"version":"stable"},"error":null,"performance":{"interrogation":185}}`)
  
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(string(body))
	
	url2 := "https://api.formula1.com/v2/account/subscriber/authenticate/by-password"
	method2 := "POST"
  
	payload2 := strings.NewReader(`{
	  "DistributionChannel": "d861e38f-05ea-4063-8776-a7e2b6d885a4",
	  "Login": "Maarten",
	  "Password": "Mvandun588!"
  }`)
  
	client2 := &http.Client {
	}
	req2, err2 := http.NewRequest(method2, url2, payload2)
  
	if err2 != nil {
	  fmt.Println(err2)
	  return
	}
	req2.Header.Add("apiKey", "fCUCjWrKPu9ylJwRAv8BpGLEgiAuThx7")
	req2.Header.Add("Cookie", "reese84={{reese84}}; login={\"event\":\"login\",\"componentId\":\"component_login_page\",\"actionType\":\"failed\"}")
  
	res2, err2 := client2.Do(req2)
	if err2 != nil {
	  fmt.Println(err2)
	  return
	}
	defer res2.Body.Close()
  
	body2, err2 := ioutil.ReadAll(res2.Body)
	if err2 != nil {
	  fmt.Println(err2)
	  return
	}
	fmt.Println(string(body2))
}
