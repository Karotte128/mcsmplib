package mcsmplib

type Request struct {
	Method string
	Params []any
}

func RPCDiscover() Request {
	r := Request{
		Method: "rpc.discover",
		Params: []any{},
	}

	return r
}

func AllowlistAdd(players []Player) Request {
	playersParam := make([]any, len(players))
	for i, player := range players {
		playersParam[i] = player
	}

	r := Request{
		Method: "allowlist/add",
		Params: []any{playersParam},
	}

	return r
}
