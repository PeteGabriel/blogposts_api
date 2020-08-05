package posts

import (
	"errors"
	"log"

	"github.com/petegabriel/personalblog/infra"
)




func All() []infra.BlogPost {
	return infra.All()
}

/**
Save a new post with given title and body.
*/
func Save(title, body string) (bool, error) {
	if title == "" {
		return false, errors.New("title must not be empty")
	}
	if body == "" {
		return false, errors.New("body must not be empty")
	}


	id, err := infra.Save(title, body)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return id > 0, nil
}
