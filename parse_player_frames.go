package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strconv"

	dem "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

type PlayerFrame struct {
	tick int `parquet:"name=tick, type=INT32"`
	steamID uint64 `parquet:"name=steamID, type=UINT_64`
	gameUserID int `parquet:"name=gameUserID, type=INT32"`
	playerName string `parquet:"name=playerName, type=UTF8"`
	side string `parquet:"name=side, type=UTF8"`
	teamName string `parquet:"name=teamName, type=UTF8"`
	currentWeapon string `parquet:"name=currentWeapon, type=UTF8"`
	hasDefuser bool `parquet:"name=hasDefuser, type=BOOLEAN"`
	hasHelmet bool `parquet:"name=hasHelmet, type=BOOLEAN"`
	hp int `parquet:"name=hp, type=INT32"`
	armor int `parquet:"name=armor, type=INT32"`
	equipmentValue int `parquet:"name=equipmentValue, type=INT32"`
	isAirborne bool `parquet:"name=isAirborne, type=BOOLEAN"`
	isAlive bool `parquet:"name=isAlive, type=BOOLEAN"`
	isBlinded bool `parquet:"name=isBlinded, type=BOOLEAN"`
	isDucking bool `parquet:"name=isDucking, type=BOOLEAN"`
	isDuckingInProgress bool `parquet:"name=isDuckingInProgress, type=BOOLEAN"`
	isInBombZone bool `parquet:"name=isInBombZone, type=BOOLEAN"`
	isInBuyZone bool `parquet:"name=isInBuyZone, type=BOOLEAN"`
	isStanding bool `parquet:"name=isStanding, type=BOOLEAN"`
	isUnDuckingInProgress bool `parquet:"name=isUnDuckingInProgress, type=BOOLEAN"`
	isWalking bool `parquet:"name=isWalking, type=BOOLEAN"`
	lastPlaceName string `parquet:"name=lastPlaceName, type=UTF8"`
	ping int `parquet:"name=ping, type=INT32"`
	x float64 `parquet:"name=x, type=FLOAT"`
	y float64 `parquet:"name=y, type=FLOAT"`
	z float64 `parquet:"name=z, type=FLOAT"`
	velX float64 `parquet:"name=velX, type=FLOAT"`
	velY float64 `parquet:"name=velY, type=FLOAT"`
	velZ float64 `parquet:"name=velZ, type=FLOAT"`
	eyeX float64 `parquet:"name=eyeX, type=FLOAT"`
	eyeY float64 `parquet:"name=eyeY, type=FLOAT"`
	eyeZ float64 `parquet:"name=eyeZ, type=FLOAT"`
	viewX float64 `parquet:"name=viewX, type=FLOAT"`
	viewY float64 `parquet:"name=viewY, type=FLOAT"`
}

// Parse player frames by tick
func main() {
	fl := new(flag.FlagSet)
	demoPathPtr := fl.String("demo", "", "Demo file `path`")
	filenamePtr := fl.String("filename", "", "Filename to write")

	err := fl.Parse(os.Args[1:])
	checkError(err)

	demPath := *demoPathPtr
	filename := *filenamePtr

	// Read in demofile
	f, err := os.Open(demPath)
	defer f.Close()
	checkError(err)

	// Create new demoparser
	p := dem.NewParser(f)
	defer p.Close()

	// CSV writer
	file, err := os.Create(filename)
	checkError(err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	playerFrames := []PlayerFrame{}

	// Define column headers
	headers := []string{
		"tick",
		"steamID",
		"gameUserID",
		"playerName",
		"side",
		"teamName",
		"currentWeapon",
		"hasDefuser",
		"hasHelmet",
		"hp",
		"armor",
		"equipmentValue",
		"isAirborne",
		"isAlive",
		"isBlinded",
		"isDucking",
		"isDuckingInProgress",
		"isInBombZone",
		"isInBuyZone",
		"isStanding",
		"isUnDuckingInProgress",
		"isWalking",
		"lastPlaceName",
		"ping",
		"x",
		"y",
		"z",
		"velX",
		"velY",
		"velZ",
		"eyeX",
		"eyeY",
		"eyeZ",
		"viewX",
		"viewY",
	}

	// Write column headers
	writer.Write(headers)

	// Parse a demo frame
	p.RegisterEventHandler(func(e events.FrameDone) {
		// tick steamid gameuserid playername side team currentweapon primary secondary flashes smokes fires he decoy defuser hashelmet hp armor 
		// isairborne isalive isblinded isducking isduckinginprogress isinbombzone isinbuyzone isstanding isunduckinginprogress iswalking 
		// lastplacename ping x y z velx vely velz eyex eyey eyez viewx viewy	
		gs := p.GameState()

		// Parse T players
		for _, p := range gs.TeamTerrorists().Members() {
			if p != nil {
				playerFrameRow := PlayerFrame{}
				playerFrameRow.tick = gs.IngameTick()
				playerFrameRow.steamID = p.SteamID64
				playerFrameRow.gameUserID = p.UserID
				playerFrameRow.playerName = p.Name
				playerFrameRow.side = "T"
				playerFrameRow.teamName = p.TeamState.ClanName()
				if p.ActiveWeapon() != nil {
					playerFrameRow.currentWeapon = p.ActiveWeapon().String()
				} else {
					playerFrameRow.currentWeapon = ""
				}
				playerFrameRow.hasDefuser = p.HasDefuseKit()
				playerFrameRow.hasHelmet = p.HasHelmet()
				playerFrameRow.hp = p.Health()
				playerFrameRow.armor = p.Armor()
				playerFrameRow.equipmentValue = p.EquipmentValueCurrent()
				playerFrameRow.isAirborne = p.IsAirborne()
				playerFrameRow.isAlive = p.IsAlive()
				playerFrameRow.isBlinded = p.IsBlinded()
				playerFrameRow.isDucking = p.IsDucking()
				playerFrameRow.isDuckingInProgress = p.IsDuckingInProgress()
				playerFrameRow.isInBombZone = p.IsInBombZone()
				playerFrameRow.isInBuyZone = p.IsInBuyZone()
				playerFrameRow.isStanding = p.IsStanding()
				playerFrameRow.isUnDuckingInProgress = p.IsUnDuckingInProgress()
				playerFrameRow.isWalking = p.IsWalking()
				playerFrameRow.lastPlaceName = p.LastPlaceName()
				playerFrameRow.ping = p.Ping()
				playerFrameRow.x = p.Position().X
				playerFrameRow.y = p.Position().Y
				playerFrameRow.z = p.Position().Z
				playerFrameRow.velX = p.Velocity().X
				playerFrameRow.velY = p.Velocity().Y
				playerFrameRow.velZ = p.Velocity().Z
				playerFrameRow.eyeX = p.PositionEyes().X
				playerFrameRow.eyeY = p.PositionEyes().Y
				playerFrameRow.eyeZ = p.PositionEyes().Z
				playerFrameRow.viewX = float64(p.ViewDirectionX())
				playerFrameRow.viewY = float64(p.ViewDirectionY())

				playerFrames = append(playerFrames, playerFrameRow)
			}
		}

		// Parse CT players
		for _, p := range gs.TeamCounterTerrorists().Members() {
			if p != nil {
				if p != nil {
					playerFrameRow := PlayerFrame{}
					playerFrameRow.tick = gs.IngameTick()
					playerFrameRow.steamID = p.SteamID64
					playerFrameRow.gameUserID = p.UserID
					playerFrameRow.playerName = p.Name
					playerFrameRow.side = "T"
					playerFrameRow.teamName = p.TeamState.ClanName()
					if p.ActiveWeapon() != nil {
						playerFrameRow.currentWeapon = p.ActiveWeapon().String()
					} else {
						playerFrameRow.currentWeapon = ""
					}
					playerFrameRow.hasDefuser = p.HasDefuseKit()
					playerFrameRow.hasHelmet = p.HasHelmet()
					playerFrameRow.hp = p.Health()
					playerFrameRow.armor = p.Armor()
					playerFrameRow.equipmentValue = p.EquipmentValueCurrent()
					playerFrameRow.isAirborne = p.IsAirborne()
					playerFrameRow.isAlive = p.IsAlive()
					playerFrameRow.isBlinded = p.IsBlinded()
					playerFrameRow.isDucking = p.IsDucking()
					playerFrameRow.isDuckingInProgress = p.IsDuckingInProgress()
					playerFrameRow.isInBombZone = p.IsInBombZone()
					playerFrameRow.isInBuyZone = p.IsInBuyZone()
					playerFrameRow.isStanding = p.IsStanding()
					playerFrameRow.isUnDuckingInProgress = p.IsUnDuckingInProgress()
					playerFrameRow.isWalking = p.IsWalking()
					playerFrameRow.lastPlaceName = p.LastPlaceName()
					playerFrameRow.ping = p.Ping()
					playerFrameRow.x = p.Position().X
					playerFrameRow.y = p.Position().Y
					playerFrameRow.z = p.Position().Z
					playerFrameRow.velX = p.Velocity().X
					playerFrameRow.velY = p.Velocity().Y
					playerFrameRow.velZ = p.Velocity().Z
					playerFrameRow.eyeX = p.PositionEyes().X
					playerFrameRow.eyeY = p.PositionEyes().Y
					playerFrameRow.eyeZ = p.PositionEyes().Z
					playerFrameRow.viewX = float64(p.ViewDirectionX())
					playerFrameRow.viewY = float64(p.ViewDirectionY())
	
					playerFrames = append(playerFrames, playerFrameRow)
				}
			}
		}
	})

	// Parse demofile to end
	err = p.ParseToEnd()

	// Write all results
	for _, p := range playerFrames {
		r := make([]string, 0, 1+len(headers))
		
		r = append(
			r,
			strconv.Itoa(p.tick),
			strconv.FormatUint(p.steamID, 32),
			strconv.Itoa(p.gameUserID),
			p.playerName,
			p.side,
			p.teamName,
			p.currentWeapon,
			strconv.FormatBool(p.hasDefuser),
			strconv.FormatBool(p.hasHelmet),
			strconv.Itoa(p.hp),
			strconv.Itoa(p.armor),
			strconv.Itoa(p.equipmentValue),
			strconv.FormatBool(p.isAirborne),
			strconv.FormatBool(p.isAlive),
			strconv.FormatBool(p.isBlinded),
			strconv.FormatBool(p.isDucking),
			strconv.FormatBool(p.isDuckingInProgress),
			strconv.FormatBool(p.isInBombZone),
			strconv.FormatBool(p.isInBuyZone),
			strconv.FormatBool(p.isStanding),
			strconv.FormatBool(p.isUnDuckingInProgress),
			strconv.FormatBool(p.isWalking),
			p.lastPlaceName,
			strconv.Itoa(p.ping),
			strconv.FormatFloat(p.x, 'f', 5, 64),
			strconv.FormatFloat(p.y, 'f', 5, 64),
			strconv.FormatFloat(p.z, 'f', 5, 64),
			strconv.FormatFloat(p.velX, 'f', 5, 64),
			strconv.FormatFloat(p.velY, 'f', 5, 64),
			strconv.FormatFloat(p.velZ, 'f', 5, 64),
			strconv.FormatFloat(p.eyeX, 'f', 5, 64),
			strconv.FormatFloat(p.eyeY, 'f', 5, 64),
			strconv.FormatFloat(p.eyeZ, 'f', 5, 64),
			strconv.FormatFloat(p.viewX, 'f', 5, 64),
			strconv.FormatFloat(p.viewY, 'f', 5, 64),
		)

		writer.Write(r)
	}

	// Check error
	checkError(err)
}

// Function to handle errors
func checkError(err error) {
	if (err != nil) {
		panic(err)
	}
}
