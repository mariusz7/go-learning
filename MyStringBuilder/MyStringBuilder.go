package mystringbuilder


type MyStringBuilder []string


func (sb *MyStringBuilder) Append(s string) {

	*sb = append(*sb, s)
}

func (sb *MyStringBuilder) String() string {

	if len(*sb) <= 0 {

		return ""

	} else if len(*sb) == 1 {

		//Return either cached or the only one value appended
		return (*sb)[0]

	} else { //Many strings to concatenate

		var reqCap int //0 by default

		for _, str := range (*sb) {
			reqCap += len(str)
		}

		var conc []byte = make([]byte, 0, reqCap)

		for _, str := range (*sb) {
			//Should not reallocate here
			conc = append(conc, str...)
		}

		//TODO Fix as there will be nil dereference below
		*sb = nil //Free previous content
		*sb = append(*sb, string(conc[:]))

		return (*sb)[0]
	}
}

