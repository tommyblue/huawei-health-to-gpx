package ghht

import (
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type heartRate struct {
	k int
	v int
}

type position struct {
	lat float32
	lon float32
	k   int
	alt float32
	t   float32
}

type TrackLine map[string]string
type TimedLine map[string]TrackLine
type HuaweiTrack map[int]TimedLine

/*
ParseTrackDump gets a dump as a test, loops over lines and, for each line, identifies the
type and populates a map
*/
func ParseTrackDump(trackDump string) *HuaweiTrack {
	track := HuaweiTrack{}
	scanner := bufio.NewScanner(strings.NewReader(trackDump))
	for scanner.Scan() {
		timestamp, recordType, payload, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := track[timestamp]; !ok {
			track[timestamp] = TimedLine{}
		}
		track[timestamp][recordType] = payload
	}

	/* [
		s-r (k, v sempre 0) cadence
		rs (k, v)
		lbs (alt, t, lon, lat, k) location
		p-m (k, v) ritmo medio al km (in secondi)
		b-p-m (k, v)
		h-r (k, v) heart-rate
	] */
	return &track
}

func parseLine(line string) (int, string, TrackLine, error) {
	payload := TrackLine{}
	var tp string
	var timestamp int
	var err error
	for _, value := range strings.Split(line, ";") {
		if value != "" {
			r := strings.Split(value, "=")
			if r[0] == "tp" {
				tp = r[1]
			}
			v := r[1]
			if isTimestamp(r[0], tp) {
				v = fixTimestamp(v)
				timestamp, err = strconv.Atoi(v)
				if err != nil {
					return 0, "", TrackLine{}, err
				}
			}
			payload[r[0]] = v
			// fmt.Println(r[0])
		}
	}

	if timestamp == 0 {
		return 0, "", TrackLine{}, err
	}
	return timestamp, tp, payload, nil
}

func isTimestamp(value, lineType string) bool {
	return (lineType == "h-r" && value == "k") || (lineType == "lbs" && value == "t")
}

// All timestamps must have 9 digits
func fixTimestamp(timestampStr string) string {
	f, err := strconv.ParseFloat(timestampStr, 64)

	if err != nil {
		log.Fatal(err)
	}

	t := int(f)
	oom := int(math.Log10(float64(t)))
	divisor := 1
	if oom > 9 {
		divisor = int(math.Pow(10, float64(oom-9)))
	} else if oom < 9 {
		divisor = int(math.Pow(0.1, float64(9-oom)))
	}
	t = int(t / divisor)

	return strconv.Itoa(t)
}

func parseTimestamp(timestamp string) time.Time {
	t, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return time.Unix(t, 0)
}
