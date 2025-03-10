package player

import "strings"

func GenerateText() (string, error) {
	words := []string{
		"apple", "banana", "cherry", "dog", "elephant", "flower", "guitar", "house", "island", "jungle",
		// "kangaroo", "lemon", "mountain", "notebook", "orange", "pencil", "queen", "robot", "sunshine", "tiger",
		// "umbrella", "violin", "waterfall", "xylophone", "yogurt", "zebra", "airplane", "butterfly", "candle", "diamond",
		// "engine", "forest", "giraffe", "hammer", "iceberg", "jacket", "key", "lantern", "motorcycle", "necklace",
		// "ocean", "parrot", "quartz", "rainbow", "squirrel", "treasure", "unicorn", "volcano", "whistle", "x-ray",
		// "yacht", "zeppelin", "archery", "breeze", "chocolate", "dolphin", "emerald", "firefly", "galaxy", "harmony",
		// "illusion", "jigsaw", "kaleidoscope", "lighthouse", "marathon", "nostalgia", "octopus", "penguin", "quiver", "rhapsody",
		// "symphony", "telescope", "underwater", "vortex", "whirlpool", "xenon", "yearning", "zeppelin", "adventure", "balloon",
		// "courage", "destiny", "enigma", "fantasy", "glacier", "horizon", "infinity", "journey", "kindness", "lullaby",
		// "miracle", "nebula", "opulence", "paradise", "quirky", "reverie", "serendipity", "tranquility", "utopia", "voyage",
	}
	return strings.Join(words, " "), nil
}
