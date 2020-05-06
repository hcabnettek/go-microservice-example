package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Plot     string             `json:"plot,omitempty" bson:"plot,omitempty"`
	Genres   []string           `json:"genres,omitempty" bson:"genres,omitempty"`
	Runtime  int32              `json:"runtime,omitempty" bson:"runtime,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Tomatoes Tomatoes           `json:"tomatoes,omitempty" bson:"tomatoes,omitempty"`
}

type Tomatoes struct {
	Dvd         int                `bson:"dvd,omitempty" json:"dvd,omitempty"`
	LastUpdated primitive.DateTime `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Viewer      TomatoViewer       `bson:"viewer,omitempty" json:"viewer,omitempty"`
}

type TomatoViewer struct {
	Rating     int32 `bson:"rating,omitempty" json:"rating,omitempty"`
	NumReviews int32 `bson:"numReviews,omitempty" json:"numReviews,omitempty"`
	Meter      int32 `bson:"meter,omitempty" json:"meter,omitempty"`
}
