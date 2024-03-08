package pjstealth_go

import (
	"pjstealth-go/utils/random"
)

var cssInfo = KwArgs{
	"activeborder":      random.Choice([]string{"rgb(118, 118, 118)", "rgb(128, 128, 128)", "rgb(109, 109, 109)"}),
	"activetext":        random.Choice([]string{"rgb(255, 102, 0)", "rgb(100, 102, 204)"}),
	"graytext":          random.Choice([]string{"rgb(128, 104, 128)", "rgb(244, 109, 109)"}),
	"highlight":         random.Choice([]string{"rgb(185, 213, 255)", "rgb(120, 109, 255)"}),
	"highlighttext":     random.Choice([]string{"rgb(0, 0, 0)", "rgb(220, 220, 220)"}),
	"linktext":          random.Choice([]string{"rgb(0, 230, 230)", "rgb(0, 120, 120)"}),
	"threeddarkshadow":  random.Choice([]string{"rgb(118, 118, 118)", "rgb(35, 35, 35)"}),
	"threedface":        random.Choice([]string{"rgb(210, 210, 210)", "rgb(240, 240, 240)"}),
	"threedhighlight":   random.Choice([]string{"rgb(55, 55, 55)", "rgb(108, 108, 118)"}),
	"threedlightshadow": random.Choice([]string{"rgb(118, 118, 118)", "rgb(80, 80, 80)"}),
	"threedshadow":      random.Choice([]string{"rgb(118, 118, 118)", "rgb(49, 49, 49)"}),
	"visitedtext":       random.Choice([]string{"rgb(85, 55, 139)", "rgb(0, 120, 204)"}),
	"windowframe":       random.Choice([]string{"rgb(118, 118, 118)", "rgb(55, 55, 55)"}),
}

var fontNames = []string{
	"Andale Mono",
	"Arial",
	"Arial Black",
	"Arial Hebrew",
	"Arial MT",
	"Arial Narrow",
	"Arial Rounded MT Bold",
	"Arial Unicode MS",
	"Bitstream Vera Sans Mono",
	"Book Antiqua",
	"Bookman Old Style",
	"Calibri",
	"Cambria",
	"Cambria Math",
	"Century",
	"Century Gothic",
	"Century Schoolbook",
	"Comic Sans",
	"Comic Sans MS",
	"Consolas",
	"Courier",
	"Courier New",
	"Geneva",
	"Georgia",
	"Helvetica",
	"Helvetica Neue",
	"Impact",
	"Lucida Bright",
	"Lucida Calligraphy",
	"Lucida Console",
	"Lucida Fax",
	"LUCIDA GRANDE",
	"Lucida Handwriting",
	"Lucida Sans",
	"Lucida Sans Typewriter",
	"Lucida Sans Unicode",
	"Microsoft Sans Serif",
	"Monaco",
	"Monotype Corsiva",
	"MS Gothic",
	"MS Outlook",
	"MS PGothic",
	"MS Reference Sans Serif",
	"MS Sans Serif",
	"MS Serif",
	"MYRIAD",
	"MYRIAD PRO",
	"Palatino",
	"Palatino Linotype",
	"Segoe Print",
	"Segoe Script",
	"Segoe UI",
	"Segoe UI Light",
	"Segoe UI Semibold",
	"Segoe UI Symbol",
	"Tahoma",
	"Times",
	"Times New Roman",
	"Times New Roman PS",
	"Trebuchet MS",
	"Verdana",
	"Wingdings",
	"Wingdings 2",
	"Wingdings 3",
	"Bradley Hand",
	"Bradley Hand ITC",
	"Bremen Bd BT",
	"Britannic Bold",
	"Broadway",
	"Browallia New",
	"BrowalliaUPC",
	"Brush Script MT",
	"Californian FB",
	"Calisto MT",
	"Calligrapher",
	"Candara",
	"CaslonOpnface BT",
	"Castellar",
	"Centaur",
	"Cezanne",
	"CG Omega",
	"CG Times",
	"Chalkboard",
	"Chalkboard SE",
	"Chalkduster",
	"Charlesworth",
	"Charter Bd BT",
	"Charter BT",
	"Chaucer",
	"ChelthmITC Bk BT",
	"Chiller",
	"Clarendon",
	"Clarendon Condensed",
	"CloisterBlack BT",
	"Cochin",
	"Colonna MT",
	"Constantia",
	"Cooper Black",
	"Copperplate",
	"Copperplate Gothic",
	"Copperplate Gothic Bold",
	"Copperplate Gothic Light",
	"CopperplGoth Bd BT",
	"Corbel",
	"Cordia New",
	"CordiaUPC",
	"Cornerstone",
	"Coronet",
	"Cuckoo",
	"Curlz MT",
	"DaunPenh",
	"Dauphin",
	"David",
	"DB LCD Temp",
	"DELICIOUS",
	"Denmark",
	"DFKai-SB",
	"Didot",
	"DilleniaUPC",
	"DIN",
	"DokChampa",
	"Dotum",
	"DotumChe",
	"Ebrima",
	"Edwardian Script ITC",
	"Elephant",
	"English 111 Vivace BT",
	"Engravers MT",
	"EngraversGothic BT",
	"Eras Bold ITC",
	"Eras Demi ITC",
	"Eras Light ITC",
	"Eras Medium ITC",
	"EucrosiaUPC",
	"Futura Md BT",
	"Futura ZBlk BT",
	"FuturaBlack BT",
	"Gabriola",
	"Galliard BT",
	"Gautami",
	"Geeza Pro",
	"Geometr231 BT",
	"Geometr231 Hv BT",
	"Geometr231 Lt BT",
	"GeoSlab 703 Lt BT",
	"GeoSlab 703 XBd BT",
	"Gigi",
	"Gill Sans",
	"Gill Sans MT",
	"Gill Sans MT Condensed",
	"Gill Sans MT Ext Condensed Bold",
	"Gill Sans Ultra Bold",
	"Gill Sans Ultra Bold Condensed",
	"Gisha",
	"Harlow Solid Italic",
	"Harrington",
	"Heather",
	"Heiti SC",
	"Heiti TC",
	"HELV",
	"Herald",
	"High Tower Text",
	"Hiragino Kaku Gothic ProN",
	"Hiragino Mincho ProN",
	"Hoefler Text",
	"Humanst 521 Cn BT",
	"Humanst521 BT",
	"Humanst521 Lt BT",
	"Imprint MT Shadow",
	"Incised901 Bd BT",
	"Incised901 BT",
	"Incised901 Lt BT",
	"INCONSOLATA",
	"Informal Roman",
	"Informal011 BT",
	"INTERSTATE",
	"IrisUPC",
	"Iskoola Pota",
	"JasmineUPC",
	"Jazz LET",
	"Jenson",
	"Jester",
	"Jokerman",
	"Juice ITC",
	"Kabel Bk BT",
	"Kabel Ult BT",
	"Kailasa",
	"Microsoft PhagsPa",
	"Microsoft Tai Le",
	"Microsoft Uighur",
	"Microsoft YaHei",
	"Microsoft Yi Baiti",
	"MingLiU",
}

var fontsInfo = KwArgs{}

func init() {
	for _, tmpFontName := range fontNames {
		fontsInfo[tmpFontName] = random.Choice(fontNames)
	}
}

var windowsWebgl = [][]string{
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Direct3D11 vs_5_0 ps_5_0, D3D11)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 3090 Direct3D11 vs_5_0 ps_5_0, D3D11)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 3070 Ti Direct3D12 vs_5_0 ps_5_0, D3D12)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Ti Direct3D12 vs_5_0 ps_5_0, D3D12)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Direct3D12 vs_5_0 ps_5_0, D3D12)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 2080 Ti Direct3D12 vs_5_0 ps_5_0, D3D12)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 2070 Super Direct3D12 vs_5_0 ps_5_0, D3D12)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 2080 Super Direct3D12 vs_5_0 ps_5_0, D3D12)"},
	{"Google Inc.(NVIDIA)", "ANGLE (NVIDIA, NVIDIA GeForce RTX 2060 Super Direct3D12 vs_5_0 ps_5_0, D3D12)"},
}

var macWebgl = [][]string{
	{"Google Inc. (Apple)", "ANGLE (Apple, Apple M1, OpenGL 4.1)"},
	{"Google Inc. (ATI Technologies Inc.)", "ANGLE (ATI Technologies Inc., AMD Radeon Pro 5300M OpenGL Engine, OpenGL 4.1)"},
	{"Google Inc. (Intel Inc.)", "ANGLE (Intel Inc., Intel(R) Iris(TM) Plus Graphics 655, OpenGL 4.1)"},
	{"Google Inc. (Apple)", "ANGLE (Apple, Apple M2, OpenGL 4.1)"},
}

var webrtc = true
var headlessCheck = true

var canvasFeature = KwArgs{
	"height": random.RandInt(-5, 5),
	"width":  random.RandInt(-5, 5),
	"r":      random.RandInt(-5, 5),
	"g":      random.RandInt(-5, 5),
	"b":      random.RandInt(-5, 5),
	"a":      random.RandInt(-5, 5),
}

var videoFeature = KwArgs{
	"start_index":  265,
	"random_value": random.Random(),
}

var clientRectFeature = false

var navigatorHardwareConcurrency = random.Choice([]int{8, 4, 16})

var envData = KwArgs{
	"Win32": KwArgs{
		"webgl_infos":  random.Choice(windowsWebgl),
		"sys_platform": "Windows",
	},
	"MacIntel": KwArgs{
		"webgl_infos":  macWebgl,
		"sys_platform": "macOS",
	},
	"Linux x86_64": KwArgs{
		"webgl_infos":  random.Choice(windowsWebgl),
		"sys_platform": "Windows",
	},
	"webrtc":                         webrtc,
	"headless_check":                 headlessCheck,
	"canvasfeature":                  canvasFeature,
	"videofeature":                   videoFeature,
	"clientrectfeature":              clientRectFeature,
	"languages":                      []string{"en-US", "en"},
	"language":                       "en-US",
	"navigator_hardware_concurrency": navigatorHardwareConcurrency,
	"device_memory":                  navigatorHardwareConcurrency,
	"is_mobile":                      false,
	"fontsfeature":                   fontsInfo,
	"cssfeature":                     cssInfo,
	"screen_color_depth":             random.Choice([]int{16, 24, 30}),
}