package main

import (
	"fmt"
	"github.com/kpfaulkner/skyblockexplorer/pkg"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	fmt.Printf("So it begins...\n")

	apiKey := os.Getenv("HYPIXEL_API_KEY")

	//joshUUID := "69d3b6504a284bad8e144445d9f76f40"
	willUUID := "add5b864e50c4d728a3fe74e94ee6977"

	//unknownUUID := "dcb254fa10f54007a7d88d06185ea919"

	sbc := pkg.NewSkyblockClient(willUUID, apiKey)



	sbPlayer, err := sbc.GetSkyblockPlayer(willUUID)
	if err != nil {
		log.Errorf("player query failed %s", err.Error())
		return
	}

	fmt.Printf("sbplayer %+v\n", sbPlayer)


	sbPlayer2, err := sbc.GetSkyblockProfileMember(willUUID, "Grapes", "dcb254fa10f54007a7d88d06185ea919")
	if err != nil {
		log.Errorf("player query failed %s", err.Error())
		return
	}

	fmt.Printf("sbplayer2 %+v\n", sbPlayer2)



}
