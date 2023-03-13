package adSrv

import (
	"fmt"
	realip "github.com/ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
	"log"
	"net"
)

type Server struct {
	geoip *geoip2.Reader
}

func NewServer(geoip *geoip2.Reader) *Server {
	return &Server{geoip: geoip}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", s.handleHttp)
}

func (s *Server) handleHttp(ctx *fasthttp.RequestCtx) {
	ip := realip.FromRequest(ctx)
	ua := string(ctx.Request.Header.UserAgent())
	parsed := user_agent.New(ua)
	browserName, browserVersion := parsed.Browser()

	country, err := s.geoip.Country(net.ParseIP(ip))
	if err != nil {
		log.Fatalf("country parsing in geolite db ERROR: %s", err)
	}

	city, err := s.geoip.City(net.ParseIP(ip))
	if err != nil {
		log.Fatalf("city parsing in geolite db ERROR: %s", err)
	}

	ctx.WriteString(fmt.Sprintf("User-Agent: %s\n", ua))
	ctx.WriteString(fmt.Sprintf("IP: %s\n", ip))
	ctx.WriteString(fmt.Sprintf("Browser name: %s\n Browser verison: %s\n", browserName, browserVersion))
	ctx.WriteString(fmt.Sprintf("Country: %s\n", country))
	ctx.WriteString(fmt.Sprintf("City: %s\n", city))

}
