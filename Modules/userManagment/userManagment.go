package userManagment

import "github.com/gorilla/sessions"


type User struct  {

}

type Users struct {

} 

func isLoggedIn() bool { 
   session, _ := store.Get(r, "session-name")
    if username, ok :=session.Value["username"]; ok {
    	return true;
    } 

    return false;
  
}

func logIn(username string, password string) bool {
	session, _ := store.Get(r, "session-name")
    if !isLoggedIn() {
    		session.Values["username"] = username;
    		    session.Save(r, w)

    } else {
    		return true;
    }
}