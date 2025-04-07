package models

type User struct {
    ID          int    `json:"id" gorm:"primaryKey"`
    Name        string `json:"name"`
    Email       string `json:"email" gorm:"unique"`
    Password    string `json:"password"`
    ProfilePic  string `json:"profile_pic"`
}