# ytvi

Fetch YouTube Video Info without authentication.

## Installation

```bash
$ go get github.com/dqn/ytvi
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/dqn/ytvi"
)

func main() {
	resp, err := ytvi.GetVideoInfo("VIDEO_ID")
	if err != nil {
		// handle error
	}

	r := resp.Microformat.PlayerMicroformatRenderer
	fmt.Println("Title:", r.Title.SimpleText)
	fmt.Println("Description:", r.Description.SimpleText)
	fmt.Println("Publish Date:", r.PublishDate)
	fmt.Println("Caregory:", r.Category)
	fmt.Println("Length Seconds:", r.LengthSeconds)
	fmt.Println("View Count:", r.ViewCount)
	fmt.Println("Channel Name:", r.OwnerChannelName)
	fmt.Println("Live Started At:", r.LiveBroadcastDetails.StartTimestamp)
	fmt.Println("Live Ended At:", r.LiveBroadcastDetails.EndTimestamp)
	fmt.Println("Living Now:", r.LiveBroadcastDetails.IsLiveNow)

	d := resp.VideoDetails
	fmt.Println("Thumbnail URL:", d.Thumbnail.Thumbnails[0].URL)
	fmt.Println("Average Rating:", d.AverageRating)
	fmt.Println("Live Content:", d.IsLiveContent)
}

```

## License

MIT
