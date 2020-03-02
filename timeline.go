package wetsuit

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/linkage-community/wetsuit/entity"
)

type TimelineOption func(c *url.Values)

func Limit(l int) TimelineOption {
	return func(q *url.Values) {
		q.Set("count", strconv.Itoa(l))
	}
}
func SinceID(sid int) TimelineOption {
	return func(q *url.Values) {
		q.Set("sinceId", strconv.Itoa(sid))
	}
}
func MaxID(mid int) TimelineOption {
	return func(q *url.Values) {
		q.Set("maxId", strconv.Itoa(mid))
	}
}
func Search(target string) TimelineOption {
	return func(q *url.Values) {
		q.Set("search", target)
	}
}

func (c *Client) GetTimeline(key string, options ...TimelineOption) (*[]entity.Post, error) {
	q := url.Values{}
	for _, o := range options {
		o(&q)
	}
	qs := q.Encode()
	if len(qs) != 0 {
		qs = "?" + qs
	}

	bytes, err := c.Get("/v1/timelines/" + key + qs)
	if err != nil {
		return nil, err
	}
	response := []entity.Post{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
