package profile

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func (request Request) getSeasons() (<-chan Seasons, <-chan error) {
	seasonIn := make(chan string)
	go func() {
		fmt.Println("API URL ", fmt.Sprintf(seasonsURL, request.Name, request.Platform))
		seasonIn <- fmt.Sprintf(seasonsURL, request.Name, request.Platform)
		close(seasonIn)
	}()

	return getSeasonsHttp(seasonIn)
}

func getSeasonsHttp(channel <-chan string) (<-chan Seasons, <-chan error) {

	out := make(chan Seasons)
	err := make(chan error)

	go func() {
		for url := range channel {
			client := http.Client{
				Timeout: httpTimeout,
			}

			resp, httpError := client.Get(url)
			err <- httpError

			defer resp.Body.Close()

			seasons := Seasons{}

			json.NewDecoder(resp.Body).Decode(&seasons)

			out <- seasons
		}

		close(err)
		close(out)
	}()

	return out, err
}


type Seasons struct {
	Seasons struct {
		Num4 struct {
			Ncsa struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"ncsa"`
			Emea struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"emea"`
		} `json:"4"`
		Num5 struct {
			Ncsa struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"ncsa"`
			Emea struct {
				Wins int `json:"wins"`
				Losses int `json:"losses"`
				Abandons int `json:"abandons"`
				Season int `json:"season"`
				Region string `json:"region"`
				Ranking struct {
					Rating float64 `json:"rating"`
					NextRating int `json:"next_rating"`
					PrevRating int `json:"prev_rating"`
					Mean float64 `json:"mean"`
					Stdev int `json:"stdev"`
					Rank int `json:"rank"`
				} `json:"ranking"`
			} `json:"emea"`
		} `json:"5"`
	} `json:"seasons"`
}