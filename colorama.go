package colorama

import "fmt"

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

func ANSISequence(letter, default_ string, numbers []int) string {
	if len(numbers) > 0 {
		default_ = fmt.Sprint(numbers[0])
	}
	return CSI + default_ + letter
}

/*
	Moves the cursor n (default 1) cells in the given direction.
	If the cursor is already at the edge of the screen, this has no effect.
*/
// CUU – Cursor Up
func CUU(n ...int) string { return ANSISequence("A", "1", n) }

// CUD – Cursor Down
func CUD(n ...int) string { return ANSISequence("B", "1", n) }

// CUF – Cursor Forward
func CUF(n ...int) string { return ANSISequence("C", "1", n) }

// CUB – Cursor Back
func CUB(n ...int) string { return ANSISequence("D", "1", n) }

// CNL – Cursor Next Line
func CNL(n ...int) string {
	/*
		Moves cursor to beginning of the line n (default 1) lines down.
		(not ANSI.SYS)
	*/
	return ANSISequence("E", "1", n)
}

// CPL – Cursor Previous Line
func CPL(n ...int) string {
	/*
		Moves cursor to beginning of the line n (default 1) lines up.
		(not ANSI.SYS)
	*/
	return ANSISequence("F", "1", n)
}

// CHA – Cursor Horizontal Absolute
func CHA(n ...int) string {
	/*
		Moves the cursor to column n.
		(not ANSI.SYS)
	*/
	return ANSISequence("G", "1", n)
}

// ED – Erase Display
func ED(n ...int) string {
	/*
		Clears part of the screen. If n is 0 (or missing),
		clear from cursor to end of screen. If n is 1,
		clear from cursor to beginning of the screen.
		If n is 2, clear entire screen
		(and moves cursor to upper left on DOS ANSI.SYS).
	*/
	return ANSISequence("J", "0", n)
}

// EL – Erase in Line
func EL(n ...int) string {
	/*
		Erases part of the line. If n is zero (or missing),
		clear from cursor to the end of the line.
		If n is one, clear from cursor to beginning of the line.
		If n is two, clear entire line. Cursor position does not change.
	*/
	return ANSISequence("K", "0", n)
}

// SU – Scroll Up
func SU(n ...int) string {
	/*
		Scroll whole page up by n (default 1) lines.
		New lines are added at the bottom.
		(not ANSI.SYS)
	*/
	return ANSISequence("S", "1", n)
}

// SD – Scroll Down
func SD(n ...int) string {
	/*
		Scroll whole page down by n (default 1) lines.
		New lines are added at the top.
		(not ANSI.SYS)
	*/
	return ANSISequence("T", "1", n)
}

// CUP – Cursor Position
func CUP(numbers ...int) string {
	/*
		Moves the cursor to row n, column m. The values are 1-based,
		and default to 1 (top left corner) if omitted.
		A sequence such as CSI ;5H is a synonym for CSI 1;5H as well
		as CSI 17;H is the same as CSI 17H and CSI 17;1H
	*/
	var n, m string
	if len(numbers) > 1 {
		n = fmt.Sprint(numbers[0])
		m = fmt.Sprint(numbers[1])
	} else if len(numbers) == 1 {
		m = fmt.Sprint(numbers[0])
	}
	return CSI + n + ";" + m + "H"
}

// SCP – Save Cursor Position
func SCP() string {
	/* Saves the cursor position. */
	return CSI + "s"
}

// RCP – Restore Cursor Position
func RCP() string {
	/* Restores the cursor position. */
	return CSI + "u"
}
