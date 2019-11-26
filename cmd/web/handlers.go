package main

import (
    "html/template"
    "log"           
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        "./ui/html/home.page.tmpl",
        "./ui/html/base.layout.tmpl",
    }

    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // We then use the Execute() method on the template set to write the template
    // content as the response body. The last parameter to Execute() represents any
    // dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func grid(w http.ResponseWriter, r *http.Request) {
    var cookie, err = r.Cookie("callsign")
    if r.URL.Path != "/grid" {
        http.NotFound(w, r)
        return
    }

    callsign := cookie.Value
    log.Println(callsign)

    data := struct {
        Callsign string
    }{
        Callsign: "Welcome "+callsign+"!",
    }

    grid_files := []string{
        "./ui/html/grid.page.tmpl",
        "./ui/html/base.layout.tmpl",
    }

    ts, err := template.ParseFiles(grid_files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, data)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

// func players(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet && r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodGet)
// 		w.Header().Set("Allow", http.MethodPost)
// 		http.Error(w, "Method Not Allowed", 405)
// 		return
// 	}
// 	w.Write([]byte("Display a player snippet."))
// }

func playersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil{
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
        callsign := r.Form.Get("callsign")
        cookie := http.Cookie {
            Name: "callsign",
            Value: callsign,
            Path: "/",
        }
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/grid", http.StatusSeeOther)
    }
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet."))
}
