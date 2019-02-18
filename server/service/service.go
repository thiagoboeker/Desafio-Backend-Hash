package service

import (
	protopkg "../proto"
	"fmt"
	redis "github.com/garyburd/redigo/redis"
	context "golang.org/x/net/context"
	"time"
)

var bdayDiscount float32 = 0.05
var bFDiscount float32 = 0.1
var bFDAY time.Time
var redisConn = "redis-ss-0.redis-svc.default.svc.cluster.local:6379"

// Checa o desconto baseado nas regras, prioriza o dia da BlackFriday
func checkForDiscount(bday time.Time) float32 {
	today := time.Now()
	bFDAY, _ = time.Parse("2006-01-02", fmt.Sprintf("%v-%v-%v", today.Year(), 11, 25))
	if today.YearDay() == bFDAY.YearDay() {
		return bFDiscount
	}
	if today.YearDay() == bday.YearDay() {
		return bdayDiscount
	}
	return 0
}

// GetDiscount - Retorna o disconto do produto
func (s *Server) GetDiscount(context context.Context, in *protopkg.RequestDiscount) (*protopkg.Response, error) {
	conn, err := redis.Dial("tcp", redisConn)
	if err != nil {
		return nil, fmt.Errorf("Erro on conection")
	}
	defer conn.Close()

	// Pega o preço
	price, err := redis.Int64(conn.Do("HGET", fmt.Sprintf("product:%v", in.ProductID), "price"))
	if err != nil {
		return nil, fmt.Errorf("product not found")
	}

	// Pega a data de nascimento
	birthdate, err := redis.String(conn.Do("HGET", fmt.Sprintf("user:%v", in.UserID), "birthdate"))
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Transforma a data em um objeto
	bdtotime, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return nil, fmt.Errorf("error on date value %v", err)
	}
	//Checa o desconto
	discount := checkForDiscount(bdtotime)
	realPrice := int64(discount * float32(price))
	return &protopkg.Response{
		Pct:          discount,
		ValueInCents: int64(realPrice),
	}, nil
}

// Server - The main Object
type Server struct{}
