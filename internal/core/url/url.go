package url

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/oudrag/server/internal/core/app"
)

type Subdomain []string

type URL struct {
	*url.URL
	Subdomain   Subdomain
	Domain      string
	TLD         string
	Port        int
	ParsedQuery url.Values
}

func NewAppUrl() *URL {
	u, err := Parse(app.GetEnv(app.BaseUrl))
	if err != nil {
		panic(err)
	}

	return u
}

func Parse(rawURL string) (*URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	cu := &URL{URL: u}

	cu.parseHost()
	err = cu.parseQuery()

	return cu, err
}

func (u *URL) AddURI(uri string) *URL {
	u.Path = uri

	return u
}

func (u *URL) parseHost() {
	parts := strings.Split(u.Host, ".")

	if len(parts) > 1 {
		tldAndPort := parts[len(parts)-1]
		ps := strings.Split(tldAndPort, ":")
		if len(ps) == 2 {
			port, err := strconv.Atoi(ps[1])
			if err == nil {
				u.Port = port
			}
		}
		u.TLD = ps[0]

		u.Domain = parts[len(parts)-2] + "." + u.TLD
		u.Subdomain = parts[:len(parts)-2]
	}
}

func (u *URL) parseQuery() error {
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err
	}

	u.ParsedQuery = q
	return nil
}
