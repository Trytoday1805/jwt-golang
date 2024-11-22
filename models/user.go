package models

type Role string

const (
    AdminRole      Role = "admin"
    EditorRole     Role = "editor" 
    ReaderRole     Role = "reader"
)

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Username string `json:"username" bson:"username"`
    Password string `json:"password" bson:"password"`
    Role     Role   `json:"role" bson:"role"`
}