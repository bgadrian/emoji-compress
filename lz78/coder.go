package lz78

import (
	"fmt"
	"strconv"
)

//coder handles LZ78 result => output file and output file => LZ78
type coder struct {
	result string
	tmp    map[string]string
}

//output adds the new WK to the result
func (e *coder) output(ws, k string) error {

	//lazy init
	if e.tmp == nil {
		e.tmp = make(map[string]string, 200)
	}

	//TODO replace with a real algorithm
	uniqueAddress := 0
	if len(ws) > 0 {
		uniqueAddress, ok := e.tmp[ws]
		if ok == false {
			uniqueAddress = strconv.Itoa(len(e.tmp) + 1)
			e.tmp[ws] = uniqueAddress
		}
	}

	e.result = e.result + fmt.Sprintf("|%04d%s|", uniqueAddress, k)

	return nil
}

//decode from ...
func (e *coder) decode() (dictionary, error) {

	return dictionary{}, nil
}
