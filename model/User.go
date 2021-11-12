package model

type User struct {
	ID            int
	UserName      string
	Password      string
	Email         string
	Phone         string
	Birthday      string
	UserDiarys    []Diary
	UserPhotos    []Photo
	UserDocuments []Document
	UserMoments   []Moment
	UserFriends   []Friend
}
