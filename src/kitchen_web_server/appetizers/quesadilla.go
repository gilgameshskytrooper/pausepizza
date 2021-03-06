package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

// The most parent struct of quesadilla
// Compare at v5/appetizers/quesadillalist.json with JSON tags at the end of each field declaration line for a better idea on how everything works.
type QuesadillaList struct {
	Quesadillas []Quesadilla `json:"list"`
}

type Quesadilla struct {
	Title                  string        `json:"title"`
	FreeIngredients        []string      `json:"freeIngredients"`
	AddableIngredients     []string      `json:"addableIngredients"`
	AddableIngredientTypes []string      `json:"addableIngredientTypes"`
	Price                  []PriceStruct `json:"price"`
	PricePerExtra          float32       `json:"pricePerExtra"`
	Image                  []ImgStruct   `json:"image"`
	Api                    string        `json:"api"`
}

// Initialize() will initialize the values for an existing QuesadillaList object by getting data from the respective endpoint list.json file and unmarshaling them into the struct.
func (ques_list *QuesadillaList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/appetizers/quesadilla/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/appetizers/quesadilla/list.json")
	}
	err2 := json.Unmarshal(raw, &ques_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the ques_list")
	}
}

// WriteFile() will write the current values of the QuesadillaList instance that this function is operating on into the relevant list.json file
func (ques_list *QuesadillaList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(ques_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal ques_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/appetizers/quesadilla/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/appetizers/quesadilla/list.json file")
	}
}

// Update() will reassign the pointer to the QuesadillaList that this function will operate on to a new pointer passed as an argument.
// Furthermore, it will write these changes back to the JSON files.
func (elem *QuesadillaList) Update(arg *QuesadillaList) {

	*elem = *arg
	elem.WriteFile()

}

// Function FindFilenames() will look for a given title in its list of objects and return two strings:
//	1) The link to the normal colored image for an item
//	1) The link to the monochromatic image for an item (used to represent items that cannot be clicked on the menu in the Client Ordering App.
func (elem *QuesadillaList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range elem.Quesadillas {
		// fmt.Println(item.Title)
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
