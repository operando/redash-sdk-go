package redash

import (
	"io/ioutil"
	"strconv"
)

type GetUserListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetUserList() *GetUserListOutput {
	path := "/api/users"

	resp, err := c.get(path)
	if err != nil {
		return &GetUserListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetUserListOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetUserInput struct {
	UserId int
}

type GetUserOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetUser(input *GetUserInput) *GetUserOutput {
	path := "/api/users/" + strconv.Itoa(input.UserId)

	resp, err := c.get(path)
	if err != nil {
		return &GetUserOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetUserOutput{Body: string(b), StatusCode: resp.StatusCode}
}
