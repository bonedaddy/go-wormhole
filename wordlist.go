package main

import "crypto/rand"

//WordList contains the mapping of PGP words
//to be used as parts of the Wormhole Code
var WordList = [][]string{
	[]string{"aardvark", "adroitness"},
	[]string{"absurd", "adviser"},
	[]string{"accrue", "aftermath"},
	[]string{"acme", "aggregate"},
	[]string{"adrift", "alkali"},
	[]string{"adult", "almighty"},
	[]string{"afflict", "amulet"},
	[]string{"ahead", "amusement"},
	[]string{"aimless", "antenna"},
	[]string{"Algol", "applicant"},
	[]string{"allow", "Apollo"},
	[]string{"alone", "armistice"},
	[]string{"ammo", "article"},
	[]string{"ancient", "asteroid"},
	[]string{"apple", "Atlantic"},
	[]string{"artist", "atmosphere"},
	[]string{"assume", "autopsy"},
	[]string{"Athens", "Babylon"},
	[]string{"atlas", "backwater"},
	[]string{"Aztec", "barbecue"},
	[]string{"baboon", "belowground"},
	[]string{"backfield", "bifocals"},
	[]string{"backward", "bodyguard"},
	[]string{"banjo", "bookseller"},
	[]string{"beaming", "borderline"},
	[]string{"bedlamp", "bottomless"},
	[]string{"beehive", "Bradbury"},
	[]string{"beeswax", "bravado"},
	[]string{"befriend", "Brazilian"},
	[]string{"Belfast", "breakaway"},
	[]string{"berserk", "Burlington"},
	[]string{"billiard", "businessman"},
	[]string{"bison", "butterfat"},
	[]string{"blackjack", "Camelot"},
	[]string{"blockade", "candidate"},
	[]string{"blowtorch", "cannonball"},
	[]string{"bluebird", "Capricorn"},
	[]string{"bombast", "caravan"},
	[]string{"bookshelf", "caretaker"},
	[]string{"brackish", "celebrate"},
	[]string{"breadline", "cellulose"},
	[]string{"breakup", "certify"},
	[]string{"brickyard", "chambermaid"},
	[]string{"briefcase", "Cherokee"},
	[]string{"Burbank", "Chicago"},
	[]string{"button", "clergyman"},
	[]string{"buzzard", "coherence"},
	[]string{"cement", "combustion"},
	[]string{"chairlift", "commando"},
	[]string{"chatter", "company"},
	[]string{"checkup", "component"},
	[]string{"chisel", "concurrent"},
	[]string{"choking", "confidence"},
	[]string{"chopper", "conformist"},
	[]string{"Christmas", "congregate"},
	[]string{"clamshell", "consensus"},
	[]string{"classic", "consulting"},
	[]string{"classroom", "corporate"},
	[]string{"cleanup", "corrosion"},
	[]string{"clockwork", "councilman"},
	[]string{"cobra", "crossover"},
	[]string{"commence", "crucifix"},
	[]string{"concert", "cumbersome"},
	[]string{"cowbell", "customer"},
	[]string{"crackdown", "Dakota"},
	[]string{"cranky", "decadence"},
	[]string{"crowfoot", "December"},
	[]string{"crucial", "decimal"},
	[]string{"crumpled", "designing"},
	[]string{"crusade", "detector"},
	[]string{"cubic", "detergent"},
	[]string{"dashboard", "determine"},
	[]string{"deadbolt", "dictator"},
	[]string{"deckhand", "dinosaur"},
	[]string{"dogsled", "direction"},
	[]string{"dragnet", "disable"},
	[]string{"drainage", "disbelief"},
	[]string{"dreadful", "disruptive"},
	[]string{"drifter", "distortion"},
	[]string{"dropper", "document"},
	[]string{"drumbeat", "embezzle"},
	[]string{"drunken", "enchanting"},
	[]string{"Dupont", "enrollment"},
	[]string{"dwelling", "enterprise"},
	[]string{"eating", "equation"},
	[]string{"edict", "equipment"},
	[]string{"egghead", "escapade"},
	[]string{"eightball", "Eskimo"},
	[]string{"endorse", "everyday"},
	[]string{"endow", "examine"},
	[]string{"enlist", "existence"},
	[]string{"erase", "exodus"},
	[]string{"escape", "fascinate"},
	[]string{"exceed", "filament"},
	[]string{"eyeglass", "finicky"},
	[]string{"eyetooth", "forever"},
	[]string{"facial", "fortitude"},
	[]string{"fallout", "frequency"},
	[]string{"flagpole", "gadgetry"},
	[]string{"flatfoot", "Galveston"},
	[]string{"flytrap", "getaway"},
	[]string{"fracture", "glossary"},
	[]string{"framework", "gossamer"},
	[]string{"freedom", "graduate"},
	[]string{"frighten", "gravity"},
	[]string{"gazelle", "guitarist"},
	[]string{"Geiger", "hamburger"},
	[]string{"glitter", "Hamilton"},
	[]string{"glucose", "handiwork"},
	[]string{"goggles", "hazardous"},
	[]string{"goldfish", "headwaters"},
	[]string{"gremlin", "hemisphere"},
	[]string{"guidance", "hesitate"},
	[]string{"hamlet", "hideaway"},
	[]string{"highchair", "holiness"},
	[]string{"hockey", "hurricane"},
	[]string{"indoors", "hydraulic"},
	[]string{"indulge", "impartial"},
	[]string{"inverse", "impetus"},
	[]string{"involve", "inception"},
	[]string{"island", "indigo"},
	[]string{"jawbone", "inertia"},
	[]string{"keyboard", "infancy"},
	[]string{"kickoff", "inferno"},
	[]string{"kiwi", "informant"},
	[]string{"klaxon", "insincere"},
	[]string{"locale", "insurgent"},
	[]string{"lockup", "integrate"},
	[]string{"merit", "intention"},
	[]string{"minnow", "inventive"},
	[]string{"miser", "Istanbul"},
	[]string{"Mohawk", "Jamaica"},
	[]string{"mural", "Jupiter"},
	[]string{"music", "leprosy"},
	[]string{"necklace", "letterhead"},
	[]string{"Neptune", "liberty"},
	[]string{"newborn", "maritime"},
	[]string{"nightbird", "matchmaker"},
	[]string{"Oakland", "maverick"},
	[]string{"obtuse", "Medusa"},
	[]string{"offload", "megaton"},
	[]string{"optic", "microscope"},
	[]string{"orca", "microwave"},
	[]string{"payday", "midsummer"},
	[]string{"peachy", "millionaire"},
	[]string{"pheasant", "miracle"},
	[]string{"physique", "misnomer"},
	[]string{"playhouse", "molasses"},
	[]string{"Pluto", "molecule"},
	[]string{"preclude", "Montana"},
	[]string{"prefer", "monument"},
	[]string{"preshrunk", "mosquito"},
	[]string{"printer", "narrative"},
	[]string{"prowler", "nebula"},
	[]string{"pupil", "newsletter"},
	[]string{"puppy", "Norwegian"},
	[]string{"python", "October"},
	[]string{"quadrant", "Ohio"},
	[]string{"quiver", "onlooker"},
	[]string{"quota", "opulent"},
	[]string{"ragtime", "Orlando"},
	[]string{"ratchet", "outfielder"},
	[]string{"rebirth", "Pacific"},
	[]string{"reform", "pandemic"},
	[]string{"regain", "Pandora"},
	[]string{"reindeer", "paperweight"},
	[]string{"rematch", "paragon"},
	[]string{"repay", "paragraph"},
	[]string{"retouch", "paramount"},
	[]string{"revenge", "passenger"},
	[]string{"reward", "pedigree"},
	[]string{"rhythm", "Pegasus"},
	[]string{"ribcage", "penetrate"},
	[]string{"ringbolt", "perceptive"},
	[]string{"robust", "performance"},
	[]string{"rocker", "pharmacy"},
	[]string{"ruffled", "phonetic"},
	[]string{"sailboat", "photograph"},
	[]string{"sawdust", "pioneer"},
	[]string{"scallion", "pocketful"},
	[]string{"scenic", "politeness"},
	[]string{"scorecard", "positive"},
	[]string{"Scotland", "potato"},
	[]string{"seabird", "processor"},
	[]string{"select", "provincial"},
	[]string{"sentence", "proximate"},
	[]string{"shadow", "puberty"},
	[]string{"shamrock", "publisher"},
	[]string{"showgirl", "pyramid"},
	[]string{"skullcap", "quantity"},
	[]string{"skydive", "racketeer"},
	[]string{"slingshot", "rebellion"},
	[]string{"slowdown", "recipe"},
	[]string{"snapline", "recover"},
	[]string{"snapshot", "repellent"},
	[]string{"snowcap", "replica"},
	[]string{"snowslide", "reproduce"},
	[]string{"solo", "resistor"},
	[]string{"southward", "responsive"},
	[]string{"soybean", "retraction"},
	[]string{"spaniel", "retrieval"},
	[]string{"spearhead", "retrospect"},
	[]string{"spellbind", "revenue"},
	[]string{"spheroid", "revival"},
	[]string{"spigot", "revolver"},
	[]string{"spindle", "sandalwood"},
	[]string{"spyglass", "sardonic"},
	[]string{"stagehand", "Saturday"},
	[]string{"stagnate", "savagery"},
	[]string{"stairway", "scavenger"},
	[]string{"standard", "sensation"},
	[]string{"stapler", "sociable"},
	[]string{"steamship", "souvenir"},
	[]string{"sterling", "specialist"},
	[]string{"stockman", "speculate"},
	[]string{"stopwatch", "stethoscope"},
	[]string{"stormy", "stupendous"},
	[]string{"sugar", "supportive"},
	[]string{"surmount", "surrender"},
	[]string{"suspense", "suspicious"},
	[]string{"sweatband", "sympathy"},
	[]string{"swelter", "tambourine"},
	[]string{"tactics", "telephone"},
	[]string{"talon", "therapist"},
	[]string{"tapeworm", "tobacco"},
	[]string{"tempest", "tolerance"},
	[]string{"tiger", "tomorrow"},
	[]string{"tissue", "torpedo"},
	[]string{"tonic", "tradition"},
	[]string{"topmost", "travesty"},
	[]string{"tracker", "trombonist"},
	[]string{"transit", "truncated"},
	[]string{"trauma", "typewriter"},
	[]string{"treadmill", "ultimate"},
	[]string{"Trojan", "undaunted"},
	[]string{"trouble", "underfoot"},
	[]string{"tumor", "unicorn"},
	[]string{"tunnel", "unify"},
	[]string{"tycoon", "universe"},
	[]string{"uncut", "unravel"},
	[]string{"unearth", "upcoming"},
	[]string{"unwind", "vacancy"},
	[]string{"uproot", "vagabond"},
	[]string{"upset", "vertigo"},
	[]string{"upshot", "Virginia"},
	[]string{"vapor", "visitor"},
	[]string{"village", "vocalist"},
	[]string{"virus", "voyager"},
	[]string{"Vulcan", "warranty"},
	[]string{"waffle", "Waterloo"},
	[]string{"wallet", "whimsical"},
	[]string{"watchword", "Wichita"},
	[]string{"wayside", "Wilmington"},
	[]string{"willow", "Wyoming"},
	[]string{"woodlark", "yesteryear"},
	[]string{"Zulu", "Yucatan"},
}

//GetRandomWordSet returns a random entry in the word set
//using crypto/rand
func GetRandomWordSet() []string {
	key := make([]byte, 1)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return WordList[int(key[0])]
}

//GetEvenWord returns the word in the even index at
//the specified index. Will panic if out of range
func GetEvenWord(ind int) string {
	if ind < 0 || ind >= len(WordList) {
		panic("index if out of bounds for word list")
	}

	return WordList[ind][0]
}

//GetRandomEvenWord returns an even word using the
//crypto/rand stream
func GetRandomEvenWord() string {
	return GetRandomWordSet()[0]
}

//GetOddWord returns the word in the odd index at
//the specified index. Will panic if out of range
func GetOddWord(ind int) string {
	if ind < 0 || ind >= len(WordList) {
		panic("index is out of bounds for word list")
	}

	return WordList[ind][1]
}

//GetRandomOddWord returns an odd word using the
//crypto/rand stream
func GetRandomOddWord() string {
	return GetRandomWordSet()[1]
}

//GetRandomWords returns a slice of strings
//up to the size n filled with random words
//from the list, alternating between odd/even
//choices
func GetRandomWords(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			strs[i] = GetRandomOddWord()
		} else {
			strs[i] = GetRandomEvenWord()
		}
	}
	return strs
}
