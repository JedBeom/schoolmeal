package schoolmeal

import (
	"errors"
	"net/http"
	"time"
)

func (s *School) reloadSession() error {
	req, err := http.NewRequest("GET", makeURL(s.Zone, linkMainPage), nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	s.sess = nil
	for _, c := range res.Cookies() {
		if c.Name == sessionName {
			s.sess = c
			s.sess.Expires = time.Now().Add(1000 * 60 * 30)
			return nil
		}
	}

	if s.sess == nil {
		return errors.New("schoolmeal: cannot get new session")
	}

	return nil
}
