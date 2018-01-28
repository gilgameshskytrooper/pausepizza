// Kitchen Management API Backend
// 2018 Andrew Lee

// This contains the backend code to serve the endpoints for the Pause Kitchen Management API.
// Using this server, Pause Kitchen workers can manage the data controlled by the client ordering app with ease. Tasks such as adding new items, removing them entirely, disabling them until the next shipment, and changing information such as the image is all handled by the Kitchen Management App.
// This server automatically updates the endpoints for the client ordering app whenever a change is registered on the Kitchen Management side.
// Since this server has write access to the key data stores (unlike the client app), this  server is protected by a token based endpoint system.
//
// Note, some aspects of the app such as the email package and the auth and superauth JSON files in assets/v5/ have been left out of the repo due to the fact that they contain sensitive information.
// They are instead replaced by files titled
//	assets/v5/auth/example.json  assets/v5/superadmin/example.json
// and
//	src/kitchen_web_server/auth/example.go
// which should give you a good idea of where to get started if you need to implement something similar
//
// This server does not implement a database of any kind. Although the argument can be made that using a database will ease the stress of expansion and allow for the accumulation of a ton of new items. However, for this specific project, there are no requirements for either. Unlike a typical website that hosts content created by users, the objects that need to be stored in this server will be fairly small in number and variety. Furthermore, while fast, using a database to find values before sending to the client will always be slower than statically hosting the JSON files.
// Furthermore, comparing the number the stateful data operations only need to occur when the Kitchen Management App needs to change somemthing versus the amount of people trying to access that data on the client ordering app (the ordering app will access those static file serving endpoints 100's of 1000's of times more often than the Kitchen side will modify items), it seemed like a logical decision to having the JSON files that are being served at the GET endpoints serve as the stateful database of some sorts, and to have this server write to those files directly.
//
// Authentication
//
// To authenticate with the kitchen management server, the user will initially send a POST to the endpoint,
//	v5/login
// with the information of username and password.
// If the username and password is a valid registered application user, a token consisting of a randomly generated 40 character string will be generated by the server, and sent back to the client to use as the first slug.
// This generated token is time sensitive and will expire at the duration passed into the GenerateNewToken() function in auth/token.go. The web server also runs a cron job at the beginning of every minute to delete expired tokens.
//
// Authentication Examples
//
// The following is an example using curl to receive a valid token to use.
//	curl -H "Content-Type: application/json" -X POST -d '{"username":"[a valid username]","[a valid password]":"abcd"}' [server address]/v5/login
// which should return
//	{"associatedUser":"leeas","value":"[a valid token]",
// From this point, the user uses the "value" object in the JSON response to GET and to POST to all endpoints.
// To GET the v5/pizza endpoint after generating a valid access token:
//	curl [server address]/v5/[generated token]/pizza
// To POST to the pizza endpoint with new data:
//	curl -H "Content-Type: application/json" -X POST -d '[INSERT RAW JSON CONTENT]' [server address]/[generated token]/pizza
// Note, for more specific information about the workings of the GET and POST endpoints, please look at the documentation below.
// Finally, the server also keeps a fail safe superadmin username and password account which (using the same algorithm generating the token, generates a log in for a super log in (which cannot be modified by interacting with the endpoint.
// Instead, the randomly generated username and password combinations will be emailed to the list of a target list defined in the v5/superadmin/list.json file.
//
// GET Endpoints
//
// Here are all the endpoints for the Kitchen Management App.
//	v5/landing
// This is the first reached by the management app. The pauseinfo object should not be modified directly if possible since this app runs a cron job using the data in v5/landing/set in turn these booleans true of false.
// Also, this contains the 5 primary menu types of Pizza, Desserts, Appetizers, Drinks, and Sides.
//	v5/landing/set
// landing/set stores the time parameters defining when various services at the Pause Kitchen are avaiable
//	v5/pizza
//	v5/desserts
//	v5/drinks
//	v5/appetizers
//	v5/sides
// The Kitchen app also has a few endpoints that does not exist at all for the Client Ordering app.
//	v5/auth
// auth is the endpoint storing all valid username and passwords
//	v5/tokens
// tokens is the endpoint that serves the JSON file with the list of currently valid tokens
//	v5/orders
// orders is the endpoint that implements notifications websockets.
// **Implement when done**
//
// GET Examples
//
// GET hosts all files in assets/ directly as JSON files (similar to the Client Ordering app).
// Here are some examples of how to use curl to test these endpoints.
//	curl [server address]/v5/[generated token]/pizza
// Pizza Endpoint
//	curl [server address]/v5/[generated token]/landing/set
// Set endpoint
//
// POST Endpoints
//
// POSTing to an endpoint involves sending the modified ENDPOINT JSON object as a POST request.
// For example, in order to change the Pizza variables in curl, you would something like:
//	curl -H "Content-Type: application/json" -X POST -d '{"list":[{"title":"Poop","available":true,"deliverable":true,"image":{"normal":"files/pizzabuild.jpg","monochrome":"files/pizzabuild.mono.jpg"}},{"title":"Specialty Pizza","available":true,"deliverable":true,"image":{"normal":"files/pizzaspecialty.jpg","monochrome":"files/pizzaspecialty.mono.jpg"}}]}' [server address]/v5/[generated token]/pizza
//
// Web Sockets
//
// **To be writen once the websockets part is done**
package main
