package models;
import ("gorm.io/gorm")
type Blog struct{
    gorm.Model
    Title string `json: "title"`
    Content string `json: "author"`
    Slug string `json: "slug"`
}
