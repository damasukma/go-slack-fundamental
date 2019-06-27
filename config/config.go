package config

import(

	_"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)


type ConfigSetup struct{
	Token string
	Channel string
}

func Setup() (*ConfigSetup, error){
	
	var configSetup ConfigSetup
	
	ctx, err := os.Open("config.json")

	defer ctx.Close()

	if err != nil {
		return nil, err
	}
	
	byteValue, err := ioutil.ReadAll(ctx)  

	if err != nil {
		return nil, err
	}
	
	json.Unmarshal(byteValue, &configSetup)

	return &configSetup, nil
}
