package controller

import (
    "net/http"
    "QN/utils"
    "github.com/gorilla/sessions"
    "QN/service"
    "QN/entity"
    "github.com/google/uuid"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))
var sessionStore = make(map[string]entity.SessionData)
var sessionIDKey = "sessionID"
var sessionMaxAge = 1800

func login(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Check user credentials
    username := r.PostFormValue("username")
    password := r.FormValue("password")

    // Verify user credentials here
    if service.CheckValidUser(username, password) {
        sessionID := generateSessionID()

        //local set
        sessionData := entity.SessionData{
            Username: username,
            // other information...
        }
        sessionStore[sessionID] = sessionData
        http.SetCookie(w, &http.Cookie{
            Name:     sessionIDKey,
            Value:    sessionID,
            Path:     "/",
            MaxAge:   sessionMaxAge, // 30 minutes, in seconds
            HttpOnly: true,
        })


        // // Set session options
        // sessionOptions := sessions.Options{
        //     Path:     "/",
        //     MaxAge:   1800, // 30 minutes, in seconds
        //     HttpOnly: true,
        // }
        // // Create a new session with the session options
        // session, _ := store.New(r, sessionID)
        // session.Options = &sessionOptions
        // // Store session data
        // session.Values[sessionIDKey] = sessionID
        // // Save the session
        // session.Save(r, w)

        utils.SendJSONResponse(w, entity.ResponseData{Mess: "LOGIN SUCCESS!", Status: http.StatusOK})
    } else {
        utils.SendJSONResponse(w, entity.ResponseData{Mess: "Invalid credentials!", Status: http.StatusUnauthorized})
    }
}

func withSessionCheck(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // session, _ := store.Get(r, "session-name")
        // authenticated, ok := session.Values["authenticated"].(bool)
        // if !ok || !authenticated {
        //     utils.SendJSONResponse(w, entity.ResponseData{Mess: "Access denied. Please log in.", Status: http.StatusForbidden})
        //     return
        // }

        sessionCookie, err := r.Cookie(sessionIDKey)
        if err != nil{
            utils.SendJSONResponse(w, entity.ResponseData{Mess: err.Error(), Status: http.StatusUnauthorized})
            return
        }
        sessionID := sessionCookie.Value
        _, ok := sessionStore[sessionID]
        if !ok {
            utils.SendJSONResponse(w, entity.ResponseData{Mess: "Access denied. Unauthorized access.", Status: http.StatusForbidden})
            return
        }
        next(w, r)
    }
}

func generateSessionID() string {
    sessionID := uuid.New().String()
    return sessionID
}