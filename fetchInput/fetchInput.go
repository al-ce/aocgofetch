package fetchInput

import (
	"fmt"
	"io"
	"net/http"
)

func GetPuzzleInput(year, day int64, sessionCookie string) (string, error) {
	puzzleUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", puzzleUrl, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)

	if resp.StatusCode != 200 {
		err = fmt.Errorf(
			"Expected Status 200, got %d\nfor URL %s",
			resp.StatusCode,
			puzzleUrl,
		)
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body[:]), nil
}
