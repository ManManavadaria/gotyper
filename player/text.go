package player

import (
	"encoding/json"
	"io"
	"log"
	"math/rand/v2"
	"os"
)

func GenerateText() (string, error) {
	// words := []string{
	// 	// "ａｐｐｌｅ", "ｂａｎａｎａ", "ｃｈｅｒｒｙ", "ｄｏｇ", "ｅｌｅｐｈａｎｔ", "ｆｌｏｗｅｒ", "ｇｕｉｔａｒ", "ｈｏｕｓｅ", "ｉｓｌａｎｄ", "ｊｕｎｇｌｅ",
	// 	"kangaroo", "lemon", "mountain",
	// 	"notebook", "orange", "pencil", "queen", "robot", "sunshine", "tiger",
	// 	"umbrella", "violin", "waterfall", "xylophone", "yogurt", "zebra", "airplane", "butterfly", "candle", "diamond",
	// 	"engine", "forest", "giraffe", "hammer", "iceberg", "jacket", "key", "lantern", "motorcycle", "necklace",
	// 	"ocean", "parrot", "quartz", "rainbow", "squirrel", "treasure", "unicorn", "volcano", "whistle", "x-ray",
	// 	"yacht", "zeppelin", "archery", "breeze", "chocolate", "dolphin", "emerald", "firefly", "galaxy", "harmony",
	// 	"illusion", "jigsaw", "kaleidoscope", "lighthouse", "marathon", "nostalgia", "octopus", "penguin", "quiver", "rhapsody",
	// 	"symphony", "telescope", "underwater", "vortex", "whirlpool", "xenon", "yearning", "zeppelin", "adventure", "balloon",
	// 	"courage", "destiny", "enigma", "fantasy", "glacier", "horizon", "infinity", "journey", "kindness", "lullaby",
	// 	"miracle", "nebula", "opulence", "paradise", "quirky", "reverie", "serendipity", "tranquility", "utopia", "voyage",
	// }
	// return strings.Join(words, " "), nil

	f, err := os.Open("./player/data/context_en.json")
	if err != nil {
		log.Fatal(err)
	}

	var jsonData []struct {
		Id   int    `json:"id"`
		Para string `json:"para"`
	}

	// if err := json.NewDecoder(f).Decode(&jsonData); err != nil {
	// 	log.Fatal("Failed to load the writing text", err)
	// }

	b, _ := io.ReadAll(f)

	if err := json.Unmarshal(b, &jsonData); err != nil {
		log.Fatal(err)
	}

	// log.Print(jsonData)
	// log.Fatal(".")

	wordString := jsonData[rand.IntN(len(jsonData))]

	return wordString.Para, nil
}
