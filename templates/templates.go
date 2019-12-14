package templates

import "strings"

func New(template string) [][]bool {
	switch strings.ToLower(template) {
	case "block":
		return block()
	case "beehive":
		return behive()
	case "loaf":
		return loaf()
	case "boat":
		return boat()
	case "tub":
		return tub()
	case "blinker":
		return blinker()
	case "toad":
		return toad()
	case "pulsar":
		return pulsar()
	default:
		return pulsar()
	}
}

func block() [][]bool {
	return [][]bool{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}
}

func behive() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false},
		{false, false, true, true, false, false},
		{false, true, false, false, true, false},
		{false, false, true, true, false, false},
		{false, false, false, false, false, false},
	}
}

func loaf() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false},
		{false, false, true, true, false, false},
		{false, true, false, false, true, false},
		{false, false, true, false, true, false},
		{false, false, false, true, false, false},
		{false, false, false, false, false, false},
	}
}

func boat() [][]bool {
	return [][]bool{
		{false, false, false, false, false},
		{false, true, true, false, false},
		{false, true, false, true, true},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}
}

func tub() [][]bool {
	return [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, true, false, true, true},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}
}

func blinker() [][]bool {
	return [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
}

func toad() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, true, true, true, false},
		{false, true, true, true, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
	}
}

func beacon() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false},
		{false, true, true, false, false, false},
		{false, true, true, false, false, false},
		{false, false, false, true, true, false},
		{false, false, false, true, true, false},
		{false, false, false, false, false, false},
	}
}

func pulsar() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, true, false, false, false, true, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, true, true, true, false, false, true, true, false, true, true, false, false, true, true, true, false},
		{false, false, false, true, false, true, false, true, false, true, false, true, false, true, false, false, false},
		{false, false, false, false, false, true, true, false, false, false, true, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, true, false, false, false, true, true, false, false, false, false, false},
		{false, false, false, true, false, true, false, true, false, true, false, true, false, true, false, false, false},
		{false, true, true, true, false, false, true, true, false, true, true, false, false, true, true, true, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, true, false, false, false, true, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	}
}
