package fetchInput

import (
	"fmt"
	"io"
	"net/http"
)

func GetPuzzleInput(year, day int64) ([]byte, error) {
	puzzleUrl := fmt.Sprintf("https://adventofcode.com/%d/day/1rst54", year)

	resp, err := http.Get(puzzleUrl)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf(
			"Expected Status 200, got %d\nfor URL %s",
			resp.StatusCode,
			puzzleUrl,
		)
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
