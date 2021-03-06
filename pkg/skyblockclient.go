package pkg

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	hypixelURLBase = "https://api.hypixel.net"
)

type SkyblockClient struct {
	skyblockUrl string
	apiKey      string
	uuid        string
	client      http.Client
}

func NewSkyblockClient(uuid string, apiKey string) *SkyblockClient {
	sbc := SkyblockClient{}

	sbc.client = http.Client{Timeout: time.Duration(10) * time.Second}
	sbc.apiKey = apiKey
	sbc.uuid = uuid
	return &sbc
}

func (s SkyblockClient) doGet(url string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("unable to create request : %s", err.Error())
		return nil, err
	}

	req.Header.Add("API-Key", s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		log.Errorf("unable to perform query : %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("unable to read body response : %s", err.Error())
		return nil, err
	}

	return body, nil
}


func (s SkyblockClient) GetPlayer(playerUUID string) (*HypixelPlayerResponse, error) {

	url := fmt.Sprintf("%s/player?uuid=%s", hypixelURLBase, s.uuid)

	data, err := s.doGet(url)
	if err != nil {
		log.Errorf("unable to get player : %s", err.Error())
		return nil, err
	}

	var playerInfo HypixelPlayerResponse
	err = json.Unmarshal(data, &playerInfo)
	if err != nil {
		log.Errorf("unable to unmarshal player : %s", err.Error())
		return nil, err
	}

	return &playerInfo, nil
}

func (s SkyblockClient) GetFriends(playerUUID string) (*HypixelFriendsResponse, error) {

	url := fmt.Sprintf("%s/friends?uuid=%s", hypixelURLBase, s.uuid)

	data, err := s.doGet(url)
	if err != nil {
		log.Errorf("unable to get friends : %s", err.Error())
		return nil, err
	}

	var friends HypixelFriendsResponse
	err = json.Unmarshal(data, &friends)
	if err != nil {
		log.Errorf("unable to unmarshal friends : %s", err.Error())
		return nil, err
	}

	return &friends, nil
}


func (s SkyblockClient) GetStatus(playerUUID string) (*HypixelStatusResponse, error) {

	url := fmt.Sprintf("%s/status?uuid=%s", hypixelURLBase, playerUUID)

	data, err := s.doGet(url)
	if err != nil {
		log.Errorf("unable to get friends : %s", err.Error())
		return nil, err
	}

	var status  HypixelStatusResponse
	err = json.Unmarshal(data, &status)
	if err != nil {
		log.Errorf("unable to unmarshal status : %s", err.Error())
		return nil, err
	}

	return &status, nil
}

func (s SkyblockClient) GetSkyblockPlayer(playerUUID string) (*HypixelSkyblockResponse, error) {

	url := fmt.Sprintf("%s/skyblock/profiles?uuid=%s", hypixelURLBase, s.uuid)

	data, err := s.doGet(url)
	if err != nil {
		log.Errorf("unable to get player : %s", err.Error())
		return nil, err
	}

	fmt.Printf("skyblock %s\n", string(data))
	var playerInfo HypixelSkyblockResponse
	err = json.Unmarshal(data, &playerInfo)
	if err != nil {
		log.Errorf("unable to unmarshal skyblock profile : %s", err.Error())
		return nil, err
	}

	return &playerInfo, nil
}

// GetSkyblockProfile gets profile based on cutename
func (s SkyblockClient) GetSkyblockProfile(playerUUID string, cuteName string) (*HypixelSkyblockProfile, error) {

	sbPlayer, err := s.GetSkyblockPlayer(playerUUID)
	if err != nil {
		log.Errorf("unable to get skyblock player info : %s", err.Error())
		return nil, err
	}

	for _, profile := range sbPlayer.Profiles {
		if profile.CuteName == cuteName {
			return &profile, nil
		}
	}
	return nil, fmt.Errorf("unable to find profile for playerUUID %s : cuteName %s", playerUUID, cuteName)
}

// GetSkyblockProfileMember gets profile based on cutename and memberUUID
func (s SkyblockClient) GetSkyblockProfileMember(playerUUID string, cuteName string, memberUUID string) (*HypixelSkyblockMember, error) {

	profile, err := s.GetSkyblockProfile(playerUUID, cuteName)
	if err != nil {
		log.Errorf("unable to get skyblock profile info : %s", err.Error())
		return nil, err
	}

	if member,ok := profile.Members[memberUUID]; ok {
		return &member,nil
	}

	return nil, fmt.Errorf("unable to find member %s for playerUUID/cutename  %s/%s",memberUUID, playerUUID, cuteName)
}




