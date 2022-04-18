package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//////////////////////
//                  //
/*   GLOBAL PATHS   */
//                  //
//////////////////////

var urlToSavedFilesFolder string = "http://127.0.0.1:5500/Go/temp-images/"
var pathToFolderOfDatabase string = "Go/temp-images" //path where files should be stored

///////////////////////
//                   //
/* HELPFUL FUNCTIONS */
//                   //
///////////////////////
func cutString(str string) string {
	//REMOVING upload- if you didn't edit code below you shouldnt need this to edit
	//BUT IF SOMETHING IS WRONG WITH SPECIAL ID YOU NEED TO EDIT THIS
	//READ func uploadFile() !!!
	return str[7:]
}

func addPortToCors(w *http.ResponseWriter, port string) { //variable w is obviously | port should be a set with numbers for ex. 8080 or 2137
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:"+port)
}

/////////////////////////
//                     //
/* SERVER SIDE ACTIONS */
//                     //
/////////////////////////

func uploadFile(w http.ResponseWriter, r *http.Request) {
	addPortToCors(&w, "*")

	fmt.Fprintf(w, "Uploading File\n")

	r.ParseMultipartForm(10 << 20) //10MB max file size -> http://orokepal.pl/tabele/Matematyka/potegi2.html
	file, handler, err := r.FormFile("uploadingFile")
	if err != nil {
		fmt.Println("Error while uploading\n")
		fmt.Fprintf(w, "Error while uploading - Contact with administrator\n")
		fmt.Println(err)
		return
	}

	defer file.Close()
	///var pathDirection string = "Go/temp-images" //path where files should be stored

	//tip for tempFile.Name()
	tempFile, err := ioutil.TempFile(pathToFolderOfDatabase, "upload-*"+handler.Filename) //Creating temp file with name upload-*filename for ex. upload-RANDOMNUMBERS-FILENAME
	specialID := strings.Trim(tempFile.Name(), handler.Filename)                          //Removing filename from specialID
	specialID = strings.Trim(specialID, "Go/temp-images\\")                               //Removing path from specialID
	specialID = cutString(specialID)
	//Here some magic because we are removing upload- from specialID by using very dangerous function CutString()

	//WARNING//
	//IF YOU EDITING SOMETHING ABOVE YOU NEED TO EDIT cutString() FUNCTION TOO !!!

	// <<< TURN BUG INTO FUTURE BELOW >>>
	fmt.Fprintf(w, "Your File Name: %+v \n", handler.Filename)
	fmt.Fprintf(w, "Saved on server with special ID: %+v \n", specialID)
	fmt.Fprintf(w, "Save your file name and special ID if you wanna download your file later: !\n")
	fmt.Fprintf(w, "Or use this special url: "+"http://127.0.0.1:5500/"+pathToFolderOfDatabase+"/upload-"+specialID+handler.Filename+"\n")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func searchFile(w http.ResponseWriter, r *http.Request) {
	///urlPathDirection := "http://127.0.0.1:5500/Go/temp-images/"
	http.Redirect(w, r, urlToSavedFilesFolder+"upload-"+r.FormValue("fileName"), http.StatusFound)
}

///////////////////////
//                   //
/*  MAIN CONTROLLER  */
//                   //
///////////////////////

func setupRoots() {
	//subpages to be served
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/search", searchFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Printf("File Hosting created By Arkadiusz Oskar Kurylo\n")
	setupRoots()
}
