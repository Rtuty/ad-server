package adSrv

import (
	"fmt"
	realip "github.com/ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	dao "modules/internal/db/mongo"
	"net"
)

type Server struct {
	geoip *geoip2.Reader
	dao   *dao.CustomerMetric
}

func NewServer(geoip *geoip2.Reader, dao *dao.CustomerMetric) *Server {
	return &Server{
		geoip: geoip,
		dao:   dao,
	}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", s.handleHttp)
}

func (s *Server) handleHttp(ctx *fasthttp.RequestCtx) {
	// Создаем клиент MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb:127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("mongo db client ERROR: %s", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("mongo db client defer ERROR: %s", err)
		}
	}()

	// Собираем метрики, подключаемся к Geolite
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
	ctx.WriteString(fmt.Sprintf("Browser name: %s Browser verison: %s\n", browserName, browserVersion))
	ctx.WriteString(fmt.Sprintf("Country name: %s Code: %s\n", country.Country.Names, country.Country.IsoCode))
	ctx.WriteString(fmt.Sprintf("City name: %s Code: %s\n", city.City.Names, city.City.GeoNameID))

	//Создаем DAO структуру
	CustomerMetric, err := dao.NewAdDAO(ctx, client)
	if err != nil {
		log.Fatalf("DAO client ERROR: %s", err)
	}

}
