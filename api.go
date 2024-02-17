package afreecachat

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	dataUrl  = "https://live.afreecatv.com/afreeca/player_live_api.php?bj_id=%s"
	loginUrl = "https://login.afreecatv.com/app/LoginAction.php"
)

func (c *Client) setSocketData() error {
	data := url.Values{}
	data.Set("bid", c.Token.BJID)
	data.Set("player_type", "html5")

	req, err := http.NewRequest("POST", fmt.Sprintf(dataUrl, c.Token.BJID), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if c.Token.pdBoxTicket != "" {
		c.httpClient.Jar.SetCookies(req.URL, []*http.Cookie{
			{
				Name:  "PdboxTicket",
				Value: c.Token.pdBoxTicket,
			},
		})
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	jsonResult := gjson.GetBytes(body, "CHANNEL")
	result := jsonResult.Get("RESULT").Int()
	switch result {
	case -6:
		return errors.New("login required")
	}

	domain := jsonResult.Get("CHDOMAIN").String()
	port := jsonResult.Get("CHPT").Int() + 1

	c.socketAddress = fmt.Sprintf("wss://%s:%d/Websocket", domain, port)
	c.Token.chatRoom = jsonResult.Get("CHATNO").String()

	return nil
}

func (c *Client) login() error {
	data := url.Values{}
	data.Set("szWork", "login")
	data.Set("szType", "json")
	data.Set("szUid", c.Token.Identifier.ID)
	data.Set("szPassword", c.Token.Identifier.Password)

	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	// Header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:122.0) Gecko/20100101 Firefox/122.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	result := gjson.GetBytes(body, "RESULT").Bool()
	if result {
		cookies := resp.Header["Set-Cookie"]
		for _, cookie := range cookies {
			ck := strings.Split(cookie, "=")
			if ck[0] == "PdboxTicket" {
				// TODO -- 에러 처리
				ticket := strings.Split(ck[1], ";")[0]
				c.Token.pdBoxTicket = ticket
				break
			}
		}
	}

	return nil
}
