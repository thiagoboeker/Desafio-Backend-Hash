package service

import (
	protopkg "../proto"
	"fmt"
	redis "github.com/garyburd/redigo/redis"
	context "golang.org/x/net/context"
)

// RegisterProduct - Registra o produto na base do redis
func (s *Server) RegisterProduct(context context.Context, in *protopkg.Product) (*protopkg.Done, error) {
	conn, err := redis.Dial("tcp", redisConn)
	if err != nil {
		return nil, fmt.Errorf("Error on redis")
	}
	defer conn.Close()
	resp, err := redis.Bool(conn.Do("HSET", fmt.Sprintf("product:%v", in.ID), "price", in.Price))
	return &protopkg.Done{
		Status: resp,
	}, nil
}
