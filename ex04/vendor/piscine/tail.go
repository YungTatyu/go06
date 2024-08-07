package piscine

const (
	StateStart = iota
	StateHyphen
	StateCOption
	StateSpaceBeforeBytes
	StateBytes
	StateBytesDone
	StateFiles
)

const (
	UintMax = ^uint(0)
	IntMax  = int(UintMax >> 1)
)

type Tail struct {
	bytes int
	files []string
}

func (t *Tail) Parse(args []string) bool {
	state := StateStart
	for _, s := range args {
		if state == StateBytesDone || state == StateFiles {
			state = StateFiles
			t.files = append(t.files, s)
			continue
		}
		if state == StateHyphen {
			PrintErrorMsg("invalid option: -")
			return false
		}
		var i uint = 0
		var bytes string
		sLen := StrLen(s)
		for {
			if i >= sLen {
				break
			}
			var r rune = rune(s[i])
			switch state {
			case StateStart:
				if r != '-' {
					PrintErrorMsg("unexpected value: " + s)
					return false
				}
				state = StateHyphen
				i++
			case StateHyphen:
				if r != 'c' {
					PrintErrorMsg("invalid option: " + string(r))
					return false
				}
				state = StateCOption
				i++
			case StateCOption:
				if IsSpace(r) {
					state = StateSpaceBeforeBytes
					i++
					break
				}
				if !IsDigit(r) {
					PrintErrorMsg("unexpected value: " + s)
					return false
				}
				state = StateBytes
			case StateSpaceBeforeBytes:
				if !IsSpace(r) {
					if !IsDigit(r) {
						PrintErrorMsg("unexpected value: " + s)
						return false
					}
					state = StateBytes
					break
				}
				i++
			case StateBytes:
				if !IsDigit(r) {
					PrintErrorMsg("invalid number of bytes: ‘" + s + "‘")
					return false
				}
				bytes += string(r)
				i++
			}
		}
		if state == StateBytes {
			state = StateBytesDone
			if !t.ParseBytes(bytes) {
				PrintErrorMsg("invalid number of bytes: ‘" + bytes + "‘")
				return false
			}
		}
	}
	if state != StateFiles {
		PrintErrorMsg("invalid arguments")
		return false
	}
	return true
}

func (t *Tail) ParseBytes(s string) bool {
	var sum int = 0
	for _, v := range s {
		if !IsDigit(v) {
			return false
		}
		if sum > IntMax/10 {
			return false
		}
		sum = sum*10 + int(v-'0')
	}
	t.bytes = sum
	return true
}

func StrLen(s string) uint {
	var i uint
	for range s {
		i++
	}
	return i
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
