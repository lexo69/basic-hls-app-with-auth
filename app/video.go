package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetVideo(c *fiber.Ctx) error {
	//Get parameters
	id := c.Params("id")
	filename := c.Params("filename")
	token := c.Params("token")
	//Check if user is in access list by token
	user, ok := List[token]
	if !ok {
		//return 401 if not
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Unauthorized",
		)
	}
	log.Println("id " + id + " filename: " + filename + " token: " + token + "user: " + user.Name)
	//send th requested file
	return c.SendFile("./video/" + id + "/" + filename)

}

func GeneratePlaylist(c *fiber.Ctx) error {
	//Struct to hold the request data
	var request_data struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	//Parse request data and check for errors
	if err := c.BodyParser(&request_data); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}
	//Check if login and password are not empty
	if request_data.Login == "" || request_data.Password == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Unauthorized",
		)
	}
	//Find a user in AppUsers struct
	u, err := AppUsers.FindUser(request_data.Login)
	if err != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Unauthorized",
		)
	}
	//Compare passwords
	if u.Password != request_data.Password {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Unauthorized",
		)
	}
	//Generate a token (random string in this case)
	token := RandomString(128)
	//Add token and reference to user to Access List map
	List[token] = &u

	//Create a playlist file
	fo, err := os.Create("./playlists/" + token + ".m3u8")
	//Close a playlist file
	defer func() {
		if err := fo.Close(); err != nil {
			log.Println("Error closing file")
		}
	}()
	//Check for errors
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Server error",
		)
	}

	//Write media data to a playlist (Video 1 in this case)
	fo.WriteString("#EXTM3U\n")
	fo.WriteString("#EXTINF\n")
	//Generate an url for video, inserting a token value
	fo.WriteString("http://" + HOST + "/video/" + token + "/1/1.m3u8\n")
	//Send playlist
	return c.SendFile("./playlists/" + token + ".m3u8")

}

func InitVideoRoutes() {
	HTTP.Get("/video/:token/:id/:filename", GetVideo)
	HTTP.Post("/playlist", GeneratePlaylist)
}
