// Package auth contains authentication related structs ValidUsersList and TokenList as well as the related functions.
package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

// Login represents a username and password combination.
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// A list of username+password combos which comprise the kitchen backend admins
type ValidUsersList struct {
	ValidUsers []ValidUser `json:"admins"`
}

// Same as Login (but used as a temporary holding struct for when a login request comes in to differentiate from Login since this is not necessarily a valid user)
type ValidUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth_list *ValidUsersList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/auth/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/auth/list.json")
	}
	err2 := json.Unmarshal(raw, &auth_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the auth_list")
	}
}

func (auth_list *ValidUsersList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(auth_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal auth_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/auth/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/auth/list.json file")
	}
}

// Validate will return true if a username/password combination exists in a given ValidUsersList
func (auth_list *ValidUsersList) Validate(username, password string) bool {
	for _, user := range auth_list.ValidUsers {
		// fmt.Println(user.Username)
		// fmt.Println(user.Password)
		if username == user.Username && password == user.Password {
			return true
		}
	}
	return false
}

func (elem *ValidUsersList) Update(arg *ValidUsersList) {

	*elem = *arg
	elem.WriteFile()

}
