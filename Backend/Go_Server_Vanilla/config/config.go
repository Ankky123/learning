package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	//this will contain the database connection, which currently is a map
	//I will also define various methods on this struct to make all of them available with this struct
	Users map[string]User
}

func Init() *Config {
	//I will initialize the Config variable here
	app := &Config{
		Users: map[string]User{
			"DefaultUser": {
				Name:  "DefaultUser",
				Email: "d@g.com",
				Age:   0,
			},
		},
	}

	return app
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

//Lets make an index response
func (app *Config) Greetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Greetings!!!"))

}

func (app *Config) AddUser(w http.ResponseWriter, r *http.Request) {

	//How to add the user, first make sure same user doesn't exist
	//will need to read user name from query string
	//userName := r.URL.Query().Get("Name")
	//I am receiving the data in the json
	// I am supposedly working with JSON, so I am going to decode the query values to a user object
	// the data is in the body
	var newUser User
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Could not read the body")
		w.Header().Add("Content-Type", "error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not read the body"))
		return
	}
	log.Print("The userName is: ", newUser.Name)

	//will check in the map if any user of same name exists
	//if it exists, I will return an error that bad request
	_, user_exists := app.Users[newUser.Name]
	if user_exists {
		log.Print("The user already exist and cannot be added again")
		w.Header().Add("Content-Type", "error")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exists"))
	} else {

		err = json.Unmarshal(body, &newUser)
		if err != nil {
			log.Print("Could not unmarshal jsonn", err)
			w.Header().Add("Content-Type", "error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error in unmarshaling"))
			return
		}
		app.Users[newUser.Name] = User{Name: newUser.Name, Email: newUser.Email, Age: newUser.Age}
		w.Header().Add("Content-Type", "text")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("User added to the database"))

	}
}

func (app *Config) GetUsers(w http.ResponseWriter, r *http.Request) {
	//here i will return all the users
	//i will need to encode the data into the json
	var resp []byte
	//looping the map
	for _, userObj := range app.Users {

		//I need to add teh userObj to the json response
		respCurr, err := json.Marshal(userObj)
		if err != nil {
			log.Print("error getting the users")
			w.Header().Add("Content-Type", "error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not fetch users"))
			return
		} else {
			resp = append(resp, respCurr...)
		}

	}

	w.Header().Add("Content-Type", "applicatio-json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}
