package ghht

import (
	"bufio"
	"fmt"
	"strings"
)

type TCX struct {
	k   uint
	lat float32
	lon float32
	alt float32
	t   float32
	v   uint
}

func FromDump(dump string) {
	r := map[string][]map[string]string{}
	scanner := bufio.NewScanner(strings.NewReader(dump))
	for scanner.Scan() {
		tp, m := parseLine(scanner.Text())
		_, ok := r[tp]
		if ok {
			r[tp] = append(r[tp], m)
		} else {
			r[tp] = []map[string]string{m}
		}
	}

	keys := make([]string, len(r))
	i := 0
	for k := range r {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	/* [
		s-r (k, v sempre 0) cadence
		rs (k, v)
		lbs (alt, t, lon, lat, k) location
		p-m (k, v)
		b-p-m (k, v)
		h-r (k, v) heart-rate
	] */
	fmt.Println(r["h-r"])
}

func parseLine(line string) (string, map[string]string) {
	m := map[string]string{}
	var tp string
	for _, s := range strings.Split(line, ";") {
		if s != "" {
			r := strings.Split(s, "=")
			if r[0] == "tp" {
				tp = r[1]
			} else {
				m[r[0]] = r[1]
			}
		}
	}

	return tp, m
}
