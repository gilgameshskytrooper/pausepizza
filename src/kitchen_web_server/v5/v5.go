// Package V5 defines the Version 5 routing rules for the Kitchen Management App
// It makes heavy use of Gorilla Mux slugs to parse out slugs from links (e.g. from the route site.com/slug1/slug2/slug3, the work we did in the main package will be accessible using vars["slug1"], vars["slug2"], vars["slug2"])
package v5

import (
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/appetizers"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/auth"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/desserts"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/drinks"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/ingredients"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/landing"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/orders"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/pizza"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/sides"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/websocket"
)

// ObjectStore is a large struct which contains pointers to the struct representations of all API endpoints for the Client Side Ordering App.
// If you are confused about which structs pertain to which API object, compare the JSON tags on the struct field declarations relevant structs.
// For example, in pizza/pizza.go, we have the declaration.
//		List []SpecialtyItem `json:"list"`
// This pertains to the "list" in the assets/v5/pizza/list.json file.
//		Title                  string              `json:"title"`
// corresponds to the element
//	"list":{"title":....}
// in the same JSON file.
//
// The vast majority of these structs share 4 functions:
//	Initialize()
// Initialize will read in the JSON file specified by the struct, and then deserialize (in go terms "unmarshal") them into the struct
//	WriteFile()
// WriteFile which is a struct function that copies the values stored in its struct instance into the specified list.json file.
//	Update()
// Update takes a struct pointer arg, and sets the operating structs pointer to the arg
//	FindFilenames()
// FindFilenames will take either one or two string arguments representing the Title and Parameter (or just the Title for image links that do not associate with a parameter such as "Large", or "Pita"),
type ObjectStore struct {
	Pizza_List           *pizza.PizzaList
	Pizza_Specialty_List *pizza.SpecialtyList
	Pizza_Build_List     *pizza.BuildList
	Desserts_List        *desserts.DessertsList
	Milkshake_List       *desserts.MilkshakeList
	Icecream_List        *desserts.IceCreamList
	Cookie_List          *desserts.CookieList
	Appetizers_List      *appetizers.AppetizersList
	Quesadilla_List      *appetizers.QuesadillaList
	Cheesybread_List     *appetizers.CheesybreadList
	Chickenfingers_List  *appetizers.ChickenfingersList
	Drinks_List          *drinks.DrinksList
	Bottled_Drinks_List  *drinks.BottledDrinksList
	Fountain_Drinks_List *drinks.FountainDrinksList
	Sides_List           *sides.SidesList
	Chips_List           *sides.ChipsList
	Sauces_List          *sides.SaucesList
	Ingredients_List     *ingredients.IngredientsList
	Landing_List         *landing.Landing
	Times                *landing.TimesItems
	Tokens               *auth.TokenList
	Admins               *auth.ValidUsersList
	SuperAdmin           *auth.Super
	Orders_List          *orders.OrderList
	WebSocketHub         *websocket.Hub
}

// Initializes the in-program store of all objects that is kept track of by the Pause Kitchen application.
// Since they are of type pointer, we first generate pointers to empty structs and assign them to all fields of an ObjectStore obj through the
//	obj.Object = &package.Struct{}
// pointer generation and assignment operation.
//
// Next, we use each object's
//	Initialize()
// method to populate each object pointer using the values stored in JSON files.
func (obj *ObjectStore) Initialize() {

	obj.Pizza_List = &pizza.PizzaList{}
	obj.Pizza_Specialty_List = &pizza.SpecialtyList{}
	obj.Pizza_Build_List = &pizza.BuildList{}
	obj.Desserts_List = &desserts.DessertsList{}
	obj.Milkshake_List = &desserts.MilkshakeList{}
	obj.Icecream_List = &desserts.IceCreamList{}
	obj.Cookie_List = &desserts.CookieList{}
	obj.Appetizers_List = &appetizers.AppetizersList{}
	obj.Quesadilla_List = &appetizers.QuesadillaList{}
	obj.Cheesybread_List = &appetizers.CheesybreadList{}
	obj.Chickenfingers_List = &appetizers.ChickenfingersList{}
	obj.Drinks_List = &drinks.DrinksList{}
	obj.Bottled_Drinks_List = &drinks.BottledDrinksList{}
	obj.Fountain_Drinks_List = &drinks.FountainDrinksList{}
	obj.Sides_List = &sides.SidesList{}
	obj.Chips_List = &sides.ChipsList{}
	obj.Sauces_List = &sides.SaucesList{}
	obj.Ingredients_List = &ingredients.IngredientsList{}
	obj.Landing_List = &landing.Landing{}
	obj.Times = &landing.TimesItems{}
	obj.Tokens = &auth.TokenList{}
	obj.Admins = &auth.ValidUsersList{}
	obj.SuperAdmin = &auth.Super{}
	obj.Orders_List = &orders.OrderList{}
	obj.WebSocketHub = &websocket.Hub{}

	obj.Pizza_List.Initialize()
	obj.Pizza_Specialty_List.Initialize()
	obj.Pizza_Build_List.Initialize()
	obj.Desserts_List.Initialize()
	obj.Milkshake_List.Initialize()
	obj.Icecream_List.Initialize()
	obj.Cookie_List.Initialize()
	obj.Appetizers_List.Initialize()
	obj.Quesadilla_List.Initialize()
	obj.Cheesybread_List.Initialize()
	obj.Chickenfingers_List.Initialize()
	obj.Drinks_List.Initialize()
	obj.Bottled_Drinks_List.Initialize()
	obj.Fountain_Drinks_List.Initialize()
	obj.Sides_List.Initialize()
	obj.Chips_List.Initialize()
	obj.Sauces_List.Initialize()
	obj.Ingredients_List.Initialize()
	obj.Landing_List.Initialize()
	obj.Times.Initialize()
	obj.Tokens.Initialize()
	obj.Admins.Initialize()
	obj.SuperAdmin.Initialize()
	obj.Orders_List.Initialize()
	obj.WebSocketHub.Initialize()
}
