package main

func main() {
	// 创建一场游戏，有两个玩家
	n := 2
	game := NewGame(n)  // 玩家ID有1，2
	// 游戏开始，掷骰子50次
	game.Start(game, 50)

}

