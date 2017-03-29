package tab

type Tab map[string][]rune

var validVoices = []string{
	"bassdrum",
	"closedhihat",
	"openhihat",
	"snaredrum",
}

func isValidVoice(name string) bool {
	for _, v := range validVoices {
		if v == name {
			return true
		}
	}
	return false
}
