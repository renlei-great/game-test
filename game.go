package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

const totalStep = 99  // 地图总长

type User struct {
	Id int
	//Advance bool  // 是否向前进
	Location int  // 所处位置
}

type Game struct {
	UndoneUser map[int]*User  // 未到达终点的用户
	DoneUser map[int]*User  // 已完成的用户
}

// NewGame n：玩家数
func NewGame(n int) Game {
	undoneUser := make(map[int]*User)
	for i := 1; i <= n; i++{
		user := User{Id: i, Location: 1}
		undoneUser[i] = &user // 初始化每一个用户都默认在第一步
	}
	game := Game{}
	game.UndoneUser = undoneUser
	game.DoneUser = make(map[int]*User)
	return game
}

// Dice 掷骰子
func (g *Game) dice() int {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	rInt := int(n.Int64()) % 6 + 1
	return rInt
}

// UserMove 用户移动
//    输入：userId: 用户id
//    输出：*int 走了几步，可以是负数，如果是负数表示回撤了几步
func (g *Game) UserMove(userId int) (*int, error) {
	step := g.dice()
	user, ok  := g.UndoneUser[userId]
	if !ok{
		if _, ok  := g.DoneUser[userId]; ok{
			step = 0
			return &step, errors.New(fmt.Sprintf("UserError: %d The user has reached the destination.", userId))
		}
		return nil, errors.New("UserError: The user does not exist.")
	}

	moveStep := step + user.Location
	if moveStep == totalStep{
		g.DoneUser[userId] = user  // 移动到终点，加入到完成集合
		delete(g.UndoneUser, userId)  // 移动到终点，从未完成集合去除
		return &step, nil
	} else if moveStep > totalStep {
		user.Location = user.Location - step
		temp, _ :=strconv.Atoi("-"+strconv.Itoa(step))
		return &temp, nil
	}
	user.Location = moveStep
	return &step, nil
}

func (g *Game) Start(game Game, frequency int){
	for i := 0; i <= frequency; i++{
		for j := 1; j <= 2; j++{
			step, err := game.UserMove(j)
			if err != nil{
				fmt.Println(err)
			}
			if step != nil{
				fmt.Println("玩家",j, "，移动了：", *step)
			}
		}
	}
}

