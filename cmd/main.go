package main

import "demochain/core"

func main()  {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jack")
	bc.SendData("Send 1 EOS to Jacky")
	bc.PrintInfo()
}