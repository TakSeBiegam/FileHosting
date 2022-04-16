package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

/////////////////////////
//                     //
/* SERVER SIDE ACTIONS */
//                     //
/////////////////////////

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	r.ParseMultipartForm(10 << 20) //10MB max file size -> http://orokepal.pl/tabele/Matematyka/potegi2.html
	file, handler, err := r.FormFile("uploadingFile")
	if err != nil {
		fmt.Println("Error while uploading\n")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("AFTER STEP ONE\n")
	var pathDirection string = "Go/temp-images" //path where files should be stored

	//tip for tempFile.Name()
	tempFile, err := ioutil.TempFile(pathDirection, "upload-*"+handler.Filename) //Creating temp file with name upload-*filename for ex. upload-RANDOMNUMBERS-FILENAME
	specialID := strings.Trim(tempFile.Name(), handler.Filename)                 //Removing filename from specialID
	specialID = strings.Trim(specialID, "Go/temp-images\\")                      //Removing path from specialID
	specialID = cutString(specialID)
	//Here some magic because we are removing upload- from specialID which is using very dangerous function CutString to remove "upload-" from specialID

	//WARNING//
	//IF YOU EDITING SOMETHING ABOVE YOU NEED TO EDIT cutString() FUNCTION TOO !!!

	// <<< TURN BUG INTO FUTURE BELOW >>>
	fmt.Fprintf(w, "Saved on server with name: %+v \n", handler.Filename)
	fmt.Fprintf(w, "Saved on server with special ID: %+v \n", specialID)

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
	//Path where you want to search files
	pathDirection := "C:\\IT\\programming\\Go\\temp-images\\"
	Openfile, err := os.Open(pathDirection + "upload-" + r.FormValue("fileName"))
	if err != nil {
		fmt.Println(err)                 //COMENT FOR SERVER
		fmt.Fprintf(w, "File not found") //COMMENT FOR USER ON BROWSER
		return
	}
	fmt.Printf("%+v\n", Openfile.Name())
	io.Copy(w, Openfile)
}

///////////////////////
//                   //
/*  MAIN CONTROLLER  */
//                   //
///////////////////////

func setupRoots() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/search", searchFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Printf("File Hosting created By Arkadiusz Oskar Kurylo\n")
	setupRoots()
}
