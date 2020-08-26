package posts

import (
	"errors"
	"log"

	"github.com/petegabriel/personalblog/infra"
)

/**
All posts.
 */
func All(page int) []infra.BlogPost {
	all := infra.All()
	if page == 1 {
		if len(all) > 10 {
			return all[:10]
		}else {
			return all
		}
	}

	i := page * 10
	if len(all) > i {
		return all[i:i+10]
	}else {
		return all
	}
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

/**
Get a post by its id.
 */
func GetById( id int) (*infra.BlogPost, error){
	p, err := infra.Get(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}