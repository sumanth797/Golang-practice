package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Fullname for Author for RealWorld Application(addItems)
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"eMail"`
}

//RealWorld Application (addItems)
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// type Articles []Article

var post_data []Post

// Creating Item struct for values
// for example_2
type Item struct {
	Data      string `json:"data"` //UpperCase for exporting the fieldnames
	OtherData int    `json:"OtherData"`
}

var struct_data []Item = []Item{} // example_2
//for example_1
var data []string = []string{}

//A simple api is created that returns simple of articles based on a GET request made to a particular URL
func addItems(w http.ResponseWriter, r *http.Request) {
	// articles := Articles{
	// 	Article{Title: "Test title", Description: "Test description", Content: "Hello world"},
	// }
	var newpost Post
	json.NewDecoder(r.Body).Decode(&newpost)
	post_data = append(post_data, newpost)
	// fmt.Println("Endpoint Hit:All Articles")
	w.Header().Set("Content-type", "application/json") // for getting json format in postman instead of text
	// json.NewEncoder(w).Encode(articles)
	json.NewEncoder(w).Encode(post_data)

}
func testArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST tester")
	//w.Write([]byte("This is an Example"))
}

//create a simple server which can handle http request
func main() {
	fmt.Println("RestfulApi")
	//using postman to make http requests
	handleRequests()
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage hit")
}
func handleRequests() {
	//Creating the router
	myRouter := mux.NewRouter().StrictSlash(true)
	//Path in the function like "/" describes the action but method is doing the work we want
	myRouter.HandleFunc("/", homePage)
	/*To send data into http request is by putting the data(dynamic variable("data")) inside request path(route handler function) Ex: "/additem/{data}"
	and access that data using variable name and using "mux" library we are able to do this*/
	//If we dont specify any http method "mux" will accept any method that can access '/additem or /root and soon'
	//In postman if we call any method without specifying any http method it will call and add data the previous ones for every method. Check screeshots on 14/5/2021 at 11:48pm
	//So to prevent this we call http methods
	//'Methods' takes variadic parameter so it can take many methods like - myRouter.HandleFunc("/additem/{data}", addItem).Methods("GET","DELETE","PUT"...)
	myRouter.HandleFunc("/add/{data}", example_1).Methods("POST") // after calling any methods other than this it will return error without displaying anything.
	//Route path is describing the resource and not the action
	myRouter.HandleFunc("/adds", example_2).Methods("POST")
	myRouter.HandleFunc("/posts", addItems).Methods("POST")
	myRouter.HandleFunc("/posts", getAllItems).Methods("GET")            // Get all posts
	myRouter.HandleFunc("/posts/{id}", getSingleItem).Methods("GET")     //Getting individual Posts
	myRouter.HandleFunc("/posts/{id}", updateItem).Methods("PUT")        // used for updating the resources of all of the fields.
	myRouter.HandleFunc("/posts/{id}", patchupdateItem).Methods("PATCH") // Updating particular field in post
	myRouter.HandleFunc("/posts/{id}", deleteItem).Methods("DELETE")
	myRouter.HandleFunc("/articles", testArticles).Methods("POST")
	//start web server
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

//Route Handler function
//Defining router handler using this function
func example_1(w http.ResponseWriter, r *http.Request) {
	//Accessing that dynamic variable and "mux.Vars" returns the route variable for current request and returns a map[string]string keys are variable identifiers.
	// In this case data is the identifier so we can get access to data variable by accessing the item in this map that has key of data.
	routeVariable := mux.Vars(r)["data"] // accessing the data by typing data in brackets
	data = append(data, routeVariable)   //Add data to slice
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
	//When we restart the application we loose all of the state and slice of data stored is lost
}

func example_2(w http.ResponseWriter, r *http.Request) {
	//get Item value from json body and we can do this using json decoder with request body which is stored in request value.
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem) //Decode takes the pointer where the value is stored and we can get it by adding the address.
	//In postman after running it go to body and change to raw and type to json and enter struct data value and displays ["","",""]
	struct_data = append(struct_data, newItem)
	// To set key value wit lower case not possible in struct because cant export but can be possible by using "json tags" which is present in encoding/json package.
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(struct_data)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Data", "Application/json")
	json.NewEncoder(w).Encode(post_data)
}
func getSingleItem(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the post from the route parameter
	var idParameter = mux.Vars(r)["id"] //getting access in this map whic has the key of id
	//next step to convert from string to integer
	id, err := strconv.Atoi(idParameter)
	if err != nil {
		w.WriteHeader(400) //status code
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	//error checking
	//In API id is the index, this is a bad idea but using databse use to generate unique id for best practice
	if id >= len(post_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}
	posts := post_data[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}

func updateItem(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the post to update which is in route parameters.
	var idParameter = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParameter)
	if err != nil {
		w.WriteHeader(400) //status code
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if id >= len(post_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}
	//Get Value from JSON Body
	//create a variable to hold the decoded data from JSON body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost) //place a pointer where we wanted to store the decoded data get the address of that variable.
	post_data[id] = updatedPost
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)

}

func patchupdateItem(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the post to update which is in route parameters.
	var idParameter = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParameter)
	if err != nil {
		w.WriteHeader(400) //status code
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if id >= len(post_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}
	//Get the current value
	// posts := post_data[id]
	posts := &post_data[id]
	//json.NewDecoder(r.Body).Decode(&posts)
	json.NewDecoder(r.Body).Decode(posts)
	//post_data[id] = posts // To reflect changes everywhere rather than id
	//(or) use pointer to update this

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the post to update which is in route parameters.
	var idParameter = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParameter)
	if err != nil {
		w.WriteHeader(400) //status code
		w.Write([]byte("ID couldn't be converted to integer"))
		return
	}
	if id >= len(post_data) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found with specified Id"))
		return
	}
	//Delete the post from the slice
	post_data = append(post_data[:id], post_data[id+1:]...)
	w.WriteHeader(200)
}
