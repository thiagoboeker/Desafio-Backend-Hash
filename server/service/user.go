package service

import (
	protopkg "../proto"
	"fmt"
	redis "github.com/garyburd/redigo/redis"
	context "golang.org/x/net/context"
)

// RegisterBirthdate - Regista o usuario na base do Redis
func (s *Server) RegisterBirthdate(ctx context.Context, in *protopkg.User) (*protopkg.Done, error) {
	conn, err := redis.Dial("tcp", redisConn)
	if err != nil {
		return nil, fmt.Errorf("Erro on conection")
	}
	defer conn.Close()
	resp, err := redis.Bool(conn.Do("HSETNX", fmt.Sprintf("user:%v", in.ID), "birthdate", in.Birthdate))
	if err != nil {
		return nil, fmt.Errorf("Error on redis %v", err)
	}
	return &protopkg.Done{
		Status: resp,
	}, nil
}
