package colorama

const (
	CSI = "\033["
	OSC = "\033]"
	BEL = "\007"
	M   = "m"
)

type Colors struct {
	BLACK, RED, GREEN, YELLOW, BLUE, MAGENTA, CYAN, WHITE, RESET string
}

var Fore = Colors{
	BLACK:   CSI + "30" + M,
	RED:     CSI + "31" + M,
	GREEN:   CSI + "32" + M,
	YELLOW:  CSI + "33" + M,
	BLUE:    CSI + "34" + M,
	MAGENTA: CSI + "35" + M,
	CYAN:    CSI + "36" + M,
	WHITE:   CSI + "37" + M,
	RESET:   CSI + "39" + M,
}

var Back = Colors{
	BLACK:   CSI + "40" + M,
	RED:     CSI + "41" + M,
	GREEN:   CSI + "42" + M,
	YELLOW:  CSI + "43" + M,
	BLUE:    CSI + "44" + M,
	MAGENTA: CSI + "45" + M,
	CYAN:    CSI + "46" + M,
	WHITE:   CSI + "47" + M,
	RESET:   CSI + "49" + M,
}

var Style = struct {
	RESET, BRIGHT, DIM, NORMAL string
}{
	RESET:  CSI + "0" + M,
	BRIGHT: CSI + "1" + M,
	DIM:    CSI + "2" + M,
	NORMAL: CSI + "22" + M,
}

func Reset() string { return Fore.RESET + Back.RESET + Style.RESET }
