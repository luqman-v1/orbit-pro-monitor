package repo

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/luqman-v1/orbit-pro-monitor/model"
	log "github.com/sirupsen/logrus"
)

type OrbitConfig struct {
	URL      string
	Username string
	Password string
}

type IOrbit interface {
	Auth(ctx context.Context) (string, error)
	SetInfo(ctx context.Context, request model.RequestParam, cookie string) ([]byte, error)
}

func NewOrbit(cfg OrbitConfig) IOrbit {
	return &OrbitConfig{
		URL:      cfg.URL,
		Username: cfg.Username,
		Password: cfg.Password,
	}
}

func (o *OrbitConfig) Auth(ctx context.Context) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/login.cgi", o.URL), nil)
	if err != nil {
		log.Error("http.NewRequestWithContext", err)
		return "", err
	}
	q := req.URL.Query()
	q.Add("Action", "Digest")
	q.Add("username", "admin")
	q.Add("realm", "Highwmg")
	q.Add("nonce", "1000")
	q.Add("response", "e9f706b6a66d1b28e5b890c2b65f9628")
	q.Add("qop", "auth")
	q.Add("cnonce", "bc5b26a1fa47e45e")
	q.Add("nc", "00000005")
	q.Add("temp", "marvell")
	q.Add("_", "1687745737873")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Error("client.Do", err)
		return "", err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	var cookie string
	for _, c := range res.Cookies() {
		if c.Name == "CGISID" {
			cookie = fmt.Sprintf("%v=%v", c.Name, c.Value)
			break
		}
	}
	return cookie, nil
}

func (o *OrbitConfig) SetInfo(ctx context.Context, request model.RequestParam, cookie string) ([]byte, error) {
	url := "http://192.168.8.1/xml_action.cgi?method=set"

	b, _ := xml.Marshal(request)

	payload := strings.NewReader(fmt.Sprintf(`<?xml version="1.0" encoding="US-ASCII"?>%v`, string(b)))

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		log.Error("client.Do", err)
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("io.ReadAll", err)
		return nil, err
	}
	return body, nil
}
