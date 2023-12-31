package main

import (
	"fmt"
)

type FilmIndusrty interface {
	Register(movie Movie)
	Unregister(movie Movie)
	NotifyAll()
}

type CinemaNetwork struct {
	filialsList []*Cinema
	moviesList  []Movie
}

func (n *CinemaNetwork) Register(cinema *Cinema) {
	n.filialsList = append(n.filialsList, cinema)
}

func (n *CinemaNetwork) Unregister(cinema *Cinema) {
	for i, t := range n.filialsList {
		if t == cinema {
			n.filialsList = append(n.filialsList[:i], n.filialsList[i+1:]...)
			break
		}
	}
}

type Movie struct {
	Title string
}

func (n *CinemaNetwork) AddMovie(movie Movie) {
	n.moviesList = append(n.moviesList, movie)
	n.NotifyAll(n.moviesList)
}

func (n *CinemaNetwork) RemoveMovie(movie Movie) {
	for i, m := range n.moviesList {
		if m == movie {
			n.moviesList = append(n.moviesList[:i], n.moviesList[i+1:]...)
			break
		}
	}
	n.NotifyAll(n.moviesList)
}

func (n *CinemaNetwork) NotifyAll(movie []Movie) {
	for _, theater := range n.filialsList {
		theater.Update(movie)
	}
}

type Cinemas interface {
	Update(movie []Movie)
}

type Cinema struct {
	Name    string
	Movies  []Movie
	Network *CinemaNetwork
}

func (t *Cinema) RegisterCinema() {
	t.Network.Register(t)
}

func (t *Cinema) UnregisterCinema() {
	t.Network.Unregister(t)
}

func (t *Cinema) Update(movie []Movie) {
	t.Movies = movie
	fmt.Printf("Popular movies in %s: ", t.Name)
	for i, movie := range movie {
		fmt.Printf(movie.Title)
		if i < len(t.Movies)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}

func main() {
	Kinopark := &CinemaNetwork{}

	theater1 := Cinema{Name: "KinoparkSaryarqa", Network: Kinopark}
	theater2 := Cinema{Name: "KinoparkKeruen", Network: Kinopark}
	theater3 := Cinema{Name: "KinoparkKeruencity", Network: Kinopark}

	theater1.RegisterCinema()
	theater2.RegisterCinema()
	theater3.RegisterCinema()

	movie1 := Movie{Title: "Гран Туризмо"}
	movie2 := Movie{Title: "Астрал. 13-й этаж"}

	Kinopark.AddMovie(movie1)
	fmt.Println()
	Kinopark.AddMovie(movie2)
	fmt.Println()

	Kinopark.RemoveMovie(movie1)
}
