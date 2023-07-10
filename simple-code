package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	ImdbID      string `json:"imdb_id"`
	Title       string `json:"title"`
	Year        string `json:"year"`
	Ratting     string `json:"ratting"`
	IsSuperHero bool   `json:"is_super_hero"`
}

var movies = []Movie{}

func getAllMovies(c *fiber.Ctx) error {
	return c.JSON(movies)
}

func createMovie(c *fiber.Ctx) error {
	movie := new(Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	movies = append(movies, *movie)
	return c.JSON(movie)
}

func getMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, movie := range movies {
		if movie.ImdbID == id {
			return c.JSON(movie)
		}
	}
	errMsg := map[string]string{
		"message": "Movie id " + id + " not found",
	}

	return c.Status(http.StatusBadRequest).JSON(errMsg)
}

func deleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, movie := range movies {
		if movie.ImdbID == id {
			movies = append(movies[:i], movies[i+1:]...)
			return c.Status(204).SendString("")
		}
	}
	errMsg := map[string]string{
		"message": "Movie id " + id + " not found",
	}

	return c.Status(http.StatusBadRequest).JSON(errMsg)
}

func main() {
	app := fiber.New()

	app.Get("/api/movies/:id", getMovie)
	app.Get("/api/movies", getAllMovies)
	app.Post("/api/movies", createMovie)
	app.Delete("/api/movies/:id", deleteMovie)

	app.Listen(":8080")
}
