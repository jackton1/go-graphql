package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

type Album struct {
	ID     string `json:"id,omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Type   string `json:"type"`
}

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

var albums []Album = []Album{
	Album{
		ID:     "ts-fearless",
		Artist: "1",
		Title:  "Fearless",
		Year:   "2008",
		Type:   "album",
	},
}

var artists []Artist = []Artist{
	Artist{
		ID:   "1",
		Name: "Taylor Swift",
		Type: "artist",
	},
}

var songs []Song = []Song{
	Song{
		ID:       "1",
		Album:    "ts-fearless",
		Title:    "Fearless",
		Duration: "4:01",
		Type:     "song",
	},
	Song{
		ID:       "2",
		Album:    "ts-fearless",
		Title:    "Fifteen",
		Duration: "4:54",
		Type:     "song",
	},
}

func main() {

	songType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"album": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"duration": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	artistType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Artist",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})


	albumType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Album",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"artist": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.String,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.ObjectConfig(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"songs": &graphql.Field{
				Type: graphql.NewList(songType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Simulate long running query.
					fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

					time.Sleep(1 * time.Minute)

					fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
					return songs, nil
				},
			},
			"albums": &graphql.Field{
				Type: graphql.NewList(albumType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Simulate long running query.
					fmt.Printf("Current Unix Time: %v\n", time.Now())

					time.Sleep(1 * time.Minute)

					fmt.Printf("Current Unix Time: %v\n", time.Now())
					return albums, nil
				},
			},
			"artists": &graphql.Field{
				Type: graphql.NewList(artistType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Simulate long running query.
					fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

					time.Sleep(1 * time.Minute)

					fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
					return artists, nil
				},
			},
		},
	})

	rootMutation := graphql.ObjectConfig(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createSong": &graphql.Field{
				Type: songType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"album": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var song Song
					song.ID = params.Args["id"].(string)
					song.Album = params.Args["album"].(string)
					song.Title = params.Args["title"].(string)
					song.Duration = params.Args["duration"].(string)
					songs = append(songs, song)
					return song, nil
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("URL: %s", r.URL)
		log.Printf("Query: %s", r.URL.Query())
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Fatalf("Error encoding result, error: %v", err)
		}
	})
	err = http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Fatalf("Error running server, error: %v", err)
	}
}