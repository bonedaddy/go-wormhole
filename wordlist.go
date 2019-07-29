package main

//WordList contains the mapping of PGP words
//to be used as parts of the Wormhole Code
var WordList = map[string][]string{
	"00": []string{"aardvark", "adroitness"},
	"01": []string{"absurd", "adviser"},
	"02": []string{"accrue", "aftermath"},
	"03": []string{"acme", "aggregate"},
	"04": []string{"adrift", "alkali"},
	"05": []string{"adult", "almighty"},
	"06": []string{"afflict", "amulet"},
	"07": []string{"ahead", "amusement"},
	"08": []string{"aimless", "antenna"},
	"09": []string{"Algol", "applicant"},
	"0A": []string{"allow", "Apollo"},
	"0B": []string{"alone", "armistice"},
	"0C": []string{"ammo", "article"},
	"0D": []string{"ancient", "asteroid"},
	"0E": []string{"apple", "Atlantic"},
	"0F": []string{"artist", "atmosphere"},
	"10": []string{"assume", "autopsy"},
	"11": []string{"Athens", "Babylon"},
	"12": []string{"atlas", "backwater"},
	"13": []string{"Aztec", "barbecue"},
	"14": []string{"baboon", "belowground"},
	"15": []string{"backfield", "bifocals"},
	"16": []string{"backward", "bodyguard"},
	"17": []string{"banjo", "bookseller"},
	"18": []string{"beaming", "borderline"},
	"19": []string{"bedlamp", "bottomless"},
	"1A": []string{"beehive", "Bradbury"},
	"1B": []string{"beeswax", "bravado"},
	"1C": []string{"befriend", "Brazilian"},
	"1D": []string{"Belfast", "breakaway"},
	"1E": []string{"berserk", "Burlington"},
	"1F": []string{"billiard", "businessman"},
	"20": []string{"bison", "butterfat"},
	"21": []string{"blackjack", "Camelot"},
	"22": []string{"blockade", "candidate"},
	"23": []string{"blowtorch", "cannonball"},
	"24": []string{"bluebird", "Capricorn"},
	"25": []string{"bombast", "caravan"},
	"26": []string{"bookshelf", "caretaker"},
	"27": []string{"brackish", "celebrate"},
	"28": []string{"breadline", "cellulose"},
	"29": []string{"breakup", "certify"},
	"2A": []string{"brickyard", "chambermaid"},
	"2B": []string{"briefcase", "Cherokee"},
	"2C": []string{"Burbank", "Chicago"},
	"2D": []string{"button", "clergyman"},
	"2E": []string{"buzzard", "coherence"},
	"2F": []string{"cement", "combustion"},
	"30": []string{"chairlift", "commando"},
	"31": []string{"chatter", "company"},
	"32": []string{"checkup", "component"},
	"33": []string{"chisel", "concurrent"},
	"34": []string{"choking", "confidence"},
	"35": []string{"chopper", "conformist"},
	"36": []string{"Christmas", "congregate"},
	"37": []string{"clamshell", "consensus"},
	"38": []string{"classic", "consulting"},
	"39": []string{"classroom", "corporate"},
	"3A": []string{"cleanup", "corrosion"},
	"3B": []string{"clockwork", "councilman"},
	"3C": []string{"cobra", "crossover"},
	"3D": []string{"commence", "crucifix"},
	"3E": []string{"concert", "cumbersome"},
	"3F": []string{"cowbell", "customer"},
	"40": []string{"crackdown", "Dakota"},
	"41": []string{"cranky", "decadence"},
	"42": []string{"crowfoot", "December"},
	"43": []string{"crucial", "decimal"},
	"44": []string{"crumpled", "designing"},
	"45": []string{"crusade", "detector"},
	"46": []string{"cubic", "detergent"},
	"47": []string{"dashboard", "determine"},
	"48": []string{"deadbolt", "dictator"},
	"49": []string{"deckhand", "dinosaur"},
	"4A": []string{"dogsled", "direction"},
	"4B": []string{"dragnet", "disable"},
	"4C": []string{"drainage", "disbelief"},
	"4D": []string{"dreadful", "disruptive"},
	"4E": []string{"drifter", "distortion"},
	"4F": []string{"dropper", "document"},
	"50": []string{"drumbeat", "embezzle"},
	"51": []string{"drunken", "enchanting"},
	"52": []string{"Dupont", "enrollment"},
	"53": []string{"dwelling", "enterprise"},
	"54": []string{"eating", "equation"},
	"55": []string{"edict", "equipment"},
	"56": []string{"egghead", "escapade"},
	"57": []string{"eightball", "Eskimo"},
	"58": []string{"endorse", "everyday"},
	"59": []string{"endow", "examine"},
	"5A": []string{"enlist", "existence"},
	"5B": []string{"erase", "exodus"},
	"5C": []string{"escape", "fascinate"},
	"5D": []string{"exceed", "filament"},
	"5E": []string{"eyeglass", "finicky"},
	"5F": []string{"eyetooth", "forever"},
	"60": []string{"facial", "fortitude"},
	"61": []string{"fallout", "frequency"},
	"62": []string{"flagpole", "gadgetry"},
	"63": []string{"flatfoot", "Galveston"},
	"64": []string{"flytrap", "getaway"},
	"65": []string{"fracture", "glossary"},
	"66": []string{"framework", "gossamer"},
	"67": []string{"freedom", "graduate"},
	"68": []string{"frighten", "gravity"},
	"69": []string{"gazelle", "guitarist"},
	"6A": []string{"Geiger", "hamburger"},
	"6B": []string{"glitter", "Hamilton"},
	"6C": []string{"glucose", "handiwork"},
	"6D": []string{"goggles", "hazardous"},
	"6E": []string{"goldfish", "headwaters"},
	"6F": []string{"gremlin", "hemisphere"},
	"70": []string{"guidance", "hesitate"},
	"71": []string{"hamlet", "hideaway"},
	"72": []string{"highchair", "holiness"},
	"73": []string{"hockey", "hurricane"},
	"74": []string{"indoors", "hydraulic"},
	"75": []string{"indulge", "impartial"},
	"76": []string{"inverse", "impetus"},
	"77": []string{"involve", "inception"},
	"78": []string{"island", "indigo"},
	"79": []string{"jawbone", "inertia"},
	"7A": []string{"keyboard", "infancy"},
	"7B": []string{"kickoff", "inferno"},
	"7C": []string{"kiwi", "informant"},
	"7D": []string{"klaxon", "insincere"},
	"7E": []string{"locale", "insurgent"},
	"7F": []string{"lockup", "integrate"},
	"80": []string{"merit", "intention"},
	"81": []string{"minnow", "inventive"},
	"82": []string{"miser", "Istanbul"},
	"83": []string{"Mohawk", "Jamaica"},
	"84": []string{"mural", "Jupiter"},
	"85": []string{"music", "leprosy"},
	"86": []string{"necklace", "letterhead"},
	"87": []string{"Neptune", "liberty"},
	"88": []string{"newborn", "maritime"},
	"89": []string{"nightbird", "matchmaker"},
	"8A": []string{"Oakland", "maverick"},
	"8B": []string{"obtuse", "Medusa"},
	"8C": []string{"offload", "megaton"},
	"8D": []string{"optic", "microscope"},
	"8E": []string{"orca", "microwave"},
	"8F": []string{"payday", "midsummer"},
	"90": []string{"peachy", "millionaire"},
	"91": []string{"pheasant", "miracle"},
	"92": []string{"physique", "misnomer"},
	"93": []string{"playhouse", "molasses"},
	"94": []string{"Pluto", "molecule"},
	"95": []string{"preclude", "Montana"},
	"96": []string{"prefer", "monument"},
	"97": []string{"preshrunk", "mosquito"},
	"98": []string{"printer", "narrative"},
	"99": []string{"prowler", "nebula"},
	"9A": []string{"pupil", "newsletter"},
	"9B": []string{"puppy", "Norwegian"},
	"9C": []string{"python", "October"},
	"9D": []string{"quadrant", "Ohio"},
	"9E": []string{"quiver", "onlooker"},
	"9F": []string{"quota", "opulent"},
	"A0": []string{"ragtime", "Orlando"},
	"A1": []string{"ratchet", "outfielder"},
	"A2": []string{"rebirth", "Pacific"},
	"A3": []string{"reform", "pandemic"},
	"A4": []string{"regain", "Pandora"},
	"A5": []string{"reindeer", "paperweight"},
	"A6": []string{"rematch", "paragon"},
	"A7": []string{"repay", "paragraph"},
	"A8": []string{"retouch", "paramount"},
	"A9": []string{"revenge", "passenger"},
	"AA": []string{"reward", "pedigree"},
	"AB": []string{"rhythm", "Pegasus"},
	"AC": []string{"ribcage", "penetrate"},
	"AD": []string{"ringbolt", "perceptive"},
	"AE": []string{"robust", "performance"},
	"AF": []string{"rocker", "pharmacy"},
	"B0": []string{"ruffled", "phonetic"},
	"B1": []string{"sailboat", "photograph"},
	"B2": []string{"sawdust", "pioneer"},
	"B3": []string{"scallion", "pocketful"},
	"B4": []string{"scenic", "politeness"},
	"B5": []string{"scorecard", "positive"},
	"B6": []string{"Scotland", "potato"},
	"B7": []string{"seabird", "processor"},
	"B8": []string{"select", "provincial"},
	"B9": []string{"sentence", "proximate"},
	"BA": []string{"shadow", "puberty"},
	"BB": []string{"shamrock", "publisher"},
	"BC": []string{"showgirl", "pyramid"},
	"BD": []string{"skullcap", "quantity"},
	"BE": []string{"skydive", "racketeer"},
	"BF": []string{"slingshot", "rebellion"},
	"C0": []string{"slowdown", "recipe"},
	"C1": []string{"snapline", "recover"},
	"C2": []string{"snapshot", "repellent"},
	"C3": []string{"snowcap", "replica"},
	"C4": []string{"snowslide", "reproduce"},
	"C5": []string{"solo", "resistor"},
	"C6": []string{"southward", "responsive"},
	"C7": []string{"soybean", "retraction"},
	"C8": []string{"spaniel", "retrieval"},
	"C9": []string{"spearhead", "retrospect"},
	"CA": []string{"spellbind", "revenue"},
	"CB": []string{"spheroid", "revival"},
	"CC": []string{"spigot", "revolver"},
	"CD": []string{"spindle", "sandalwood"},
	"CE": []string{"spyglass", "sardonic"},
	"CF": []string{"stagehand", "Saturday"},
	"D0": []string{"stagnate", "savagery"},
	"D1": []string{"stairway", "scavenger"},
	"D2": []string{"standard", "sensation"},
	"D3": []string{"stapler", "sociable"},
	"D4": []string{"steamship", "souvenir"},
	"D5": []string{"sterling", "specialist"},
	"D6": []string{"stockman", "speculate"},
	"D7": []string{"stopwatch", "stethoscope"},
	"D8": []string{"stormy", "stupendous"},
	"D9": []string{"sugar", "supportive"},
	"DA": []string{"surmount", "surrender"},
	"DB": []string{"suspense", "suspicious"},
	"DC": []string{"sweatband", "sympathy"},
	"DD": []string{"swelter", "tambourine"},
	"DE": []string{"tactics", "telephone"},
	"DF": []string{"talon", "therapist"},
	"E0": []string{"tapeworm", "tobacco"},
	"E1": []string{"tempest", "tolerance"},
	"E2": []string{"tiger", "tomorrow"},
	"E3": []string{"tissue", "torpedo"},
	"E4": []string{"tonic", "tradition"},
	"E5": []string{"topmost", "travesty"},
	"E6": []string{"tracker", "trombonist"},
	"E7": []string{"transit", "truncated"},
	"E8": []string{"trauma", "typewriter"},
	"E9": []string{"treadmill", "ultimate"},
	"EA": []string{"Trojan", "undaunted"},
	"EB": []string{"trouble", "underfoot"},
	"EC": []string{"tumor", "unicorn"},
	"ED": []string{"tunnel", "unify"},
	"EE": []string{"tycoon", "universe"},
	"EF": []string{"uncut", "unravel"},
	"F0": []string{"unearth", "upcoming"},
	"F1": []string{"unwind", "vacancy"},
	"F2": []string{"uproot", "vagabond"},
	"F3": []string{"upset", "vertigo"},
	"F4": []string{"upshot", "Virginia"},
	"F5": []string{"vapor", "visitor"},
	"F6": []string{"village", "vocalist"},
	"F7": []string{"virus", "voyager"},
	"F8": []string{"Vulcan", "warranty"},
	"F9": []string{"waffle", "Waterloo"},
	"FA": []string{"wallet", "whimsical"},
	"FB": []string{"watchword", "Wichita"},
	"FC": []string{"wayside", "Wilmington"},
	"FD": []string{"willow", "Wyoming"},
	"FE": []string{"woodlark", "yesteryear"},
	"FF": []string{"Zulu", "Yucatan"},
}
