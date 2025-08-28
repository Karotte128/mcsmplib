package mcsmplib

type Request struct {
	Method string
	Params []any
}

type RPCDiscover struct{}

func (m RPCDiscover) Build() Request {
	r := Request{
		Method: "rpc.discover",
		Params: []any{},
	}

	return r
}

type AllowlistAdd struct {
	Player
}

func (m AllowlistAdd) Build(players []Player) Request {
	playersParam := make([]any, len(players))
	for i, player := range players {
		playersParam[i] = player
	}

	r := Request{
		Method: "rpc.discover",
		Params: []any{playersParam},
	}

	return r
}
