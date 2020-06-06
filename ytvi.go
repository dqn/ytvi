package ytvi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func fetch(videoID string) (url.Values, error) {
	u := "https://www.youtube.com/get_video_info?video_id=" + videoID
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	v, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}

	return v, nil
}

func GetVideoInfo(videoID string) (*PlayerResponse, error) {
	v, err := fetch(videoID)
	if err != nil {
		return nil, err
	}

	if status, ok := v["status"]; !ok {
		err = fmt.Errorf("failed to fetch video info")
		return nil, err
	} else if status[0] != "ok" {
		reason, ok := v["reason"]

		if ok {
			err = fmt.Errorf(reason[0])
			return nil, err
		}

		err = fmt.Errorf("unknown error")
		return nil, err
	}

	j, ok := v["player_response"]
	if !ok {
		err = fmt.Errorf("player_response not found")
		return nil, err
	}

	var p PlayerResponse
	if err := json.Unmarshal([]byte(j[0]), &p); err != nil {
		return nil, err
	}

	return &p, nil
}
