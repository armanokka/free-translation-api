package entity

var SupportedLanguageCodes = map[string]string{
	"hmn":   "Hmong",
	"iw":    "Hebrew (Old)",
	"zh-tw": "Chinese (Taiwan)",
	"zh-cn": "Chinese (PRC)",
	"il":    "Iloko",
	"zh-hk": "Chinese (Hong Kong)",
	"zh-sg": "Chinese (Singapore)",
	"zh-mo": "Chinese (Macau)",
	"om":    "Oromo",
	"iu":    "Inuktitut",
	"ab":    "Abkhaz",
	"sq":    "Albanian",
	"gn":    "Guaraní",
	"bh":    "Bihari",
	"fj":    "Fijian",
	"na":    "Nauru",
	"sg":    "Sango",
	"ie":    "Interlingue",
	"tn":    "Tswana",
	"sa":    "Sanskrit",
	"ln":    "Lingala",
	"ba":    "Bashkir",
	"fo":    "Faroese",
	"vo":    "Volapük",
	"ks":    "Kashmiri",
	"ik":    "Inupiaq",
	"os":    "Ossetian",
	"as":    "Assamese",
	"oc":    "Occitan",
	"ti":    "Tigrinya",
	"lg":    "Ganda",
	"ee":    "Ewe",
	"bi":    "Bislama",
	"ts":    "Tsonga",
	"to":    "Tonga",
	"rm":    "Romansh",
	"ia":    "Interlingua",
	"qu":    "Quechua",
	"bo":    "Tibetan Standard",
	"br":    "Breton",
	"ve":    "Venda",
	"gv":    "Manx",
	"wo":    "Wolof",
	"ay":    "Aymara",
	"ss":    "Swati",
	"kl":    "Kalaallisut",
	"rn":    "Kirundi",
	"za":    "Zhuang",
	"dv":    "Divehi",
	"yo":    "Yoruba",
	"nl":    "Dutch",
	"ja":    "Japanese",
	"mi":    "Māori",
	"cr":    "Cree",
	"ml":    "Malayalam",
	"wa":    "Walloon",
	"gd":    "Scottish Gaelic",
	"si":    "Sinhala",
	"ty":    "Tahitian",
	"am":    "Amharic",
	"hz":    "Herero",
	"se":    "Northern Sami",
	"uk":    "Ukrainian",
	"cy":    "Welsh",
	"gl":    "Galician",
	"mt":    "Maltese",
	"nd":    "Northern Ndebele",
	"kv":    "Komi",
	"nb":    "Norwegian Bokmål",
	"st":    "Southern Sotho",
	"tr":    "Turkish",
	"bn":    "Bengali",
	"eo":    "Esperanto",
	"is":    "Icelandic",
	"li":    "Limburgish",
	"sm":    "Samoan",
	"th":    "Thai",
	"ar":    "Arabic",
	"rw":    "Kinyarwanda",
	"la":    "Latin",
	"tg":    "Tajik",
	"ur":    "Urdu",
	"hr":    "Croatian",
	"kn":    "Kannada",
	"ps":    "Pashto",
	"ug":    "Uyghur",
	"lv":    "Latvian",
	"sw":    "Swahili",
	"tk":    "Turkmen",
	"lt":    "Lithuanian",
	"mr":    "Marathi",
	"an":    "Aragonese",
	"cs":    "Czech",
	"ka":    "Georgian",
	"de":    "German",
	"id":    "Indonesian",
	"fa":    "Persian",
	"te":    "Telugu",
	"av":    "Avaric",
	"ig":    "Igbo",
	"sl":    "Slovene",
	"lu":    "Luba-Katanga",
	"ta":    "Tamil",
	"cv":    "Chuvash",
	"lb":    "Luxembourgish",
	"pi":    "Pāli",
	"mg":    "Malagasy",
	"su":    "Sundanese",
	"tl":    "Tagalog",
	"az":    "Azerbaijani",
	"ca":    "Catalan",
	"lo":    "Lao",
	"nr":    "Southern Ndebele",
	"es":    "Spanish",
	"hy":    "Armenian",
	"kj":    "Kwanyama",
	"sk":    "Slovak",
	"ro":    "Romanian",
	"ru":    "Russian",
	"ch":    "Chamorro",
	"ha":    "Hausa",
	"ho":    "Hiri Motu",
	"km":    "Khmer",
	"nv":    "Navajo",
	"cu":    "Old Church Slavonic",
	"ny":    "Chichewa",
	"ht":    "Haitian",
	"kg":    "Kongo",
	"ky":    "Kyrgyz",
	"xh":    "Xhosa",
	"mh":    "Marshallese",
	"so":    "Somali",
	"bg":    "Bulgarian",
	"he":    "Hebrew",
	"it":    "Italian",
	"no":    "Norwegian",
	"fi":    "Finnish",
	"gu":    "Gujarati",
	"ki":    "Kikuyu",
	"nn":    "Norwegian Nynorsk",
	"pt":    "Portuguese",
	"da":    "Danish",
	"hu":    "Hungarian",
	"ku":    "Kurdish",
	"uz":    "Uzbek",
	"yi":    "Yiddish",
	"fr":    "French",
	"pl":    "Polish",
	"tt":    "Tatar",
	"mk":    "Macedonian",
	"sd":    "Sindhi",
	"sv":    "Swedish",
	"ng":    "Ndonga",
	"sr":    "Serbian",
	"ce":    "Chechen",
	"ff":    "Fula",
	"ms":    "Malay",
	"ii":    "Nuosu",
	"et":    "Estonian",
	"hi":    "Hindi",
	"kr":    "Kanuri",
	"or":    "Oriya",
	"zh":    "Chinese",
	"co":    "Corsican",
	"ko":    "Korean",
	"ga":    "Irish",
	"mn":    "Mongolian",
	"fy":    "Western Frisian",
	"be":    "Belarusian",
	"en":    "English",
	"el":    "Greek",
	"my":    "Burmese",
	"pa":    "Panjabi",
	"vi":    "Vietnamese",
	"io":    "Ido",
	"kk":    "Kazakh",
	"oj":    "Ojibwe",
	"bs":    "Bosnian",
	"kw":    "Cornish",
	"jv":    "Javanese",
	"ne":    "Nepali",
	"sc":    "Sardinian",
	"af":    "Afrikaans",
	"ae":    "Avestan",
	"eu":    "Basque",
	"sn":    "Shona",
}

var SupportedLanguageCodesEn = map[string]string{
	"Hmong":               "hmn",
	"Hebrew (Old)":        "iw",
	"Chinese (Taiwan)":    "zh-tw",
	"Iloko":               "il",
	"Chinese (PRC)":       "zh-cn",
	"Chinese (Hong Kong)": "zh-hk",
	"Chinese (Singapore)": "zh-sg",
	"Chinese (Macau)":     "zh-mo",
	"Belarusian":          "be",
	"Bosnian":             "bs",
	"Japanese":            "ja",
	"Romanian":            "ro",
	"Marathi":             "mr",
	"Russian":             "ru",
	"Hiri Motu":           "ho",
	"Irish":               "ga",
	"Khmer":               "km",
	"Kongo":               "kg",
	"Spanish":             "es",
	"Thai":                "th",
	"Latvian":             "lv",
	"Navajo":              "nv",
	"Northern Ndebele":    "nd",
	"Slovene":             "sl",
	"Afrikaans":           "af",
	"Galician":            "gl",
	"Sundanese":           "su",
	"Greek":               "el",
	"Marshallese":         "mh",
	"Pashto":              "ps",
	"Western Frisian":     "fy",
	"Kyrgyz":              "ky",
	"Latin":               "la",
	"Mongolian":           "mn",
	"Telugu":              "te",
	"Esperanto":           "eo",
	"German":              "de",
	"Oriya":               "or",
	"Tajik":               "tg",
	"Lithuanian":          "lt",
	"Sardinian":           "sc",
	"Sinhala":             "si",
	"Welsh":               "cy",
	"Hausa":               "ha",
	"Hindi":               "hi",
	"Korean":              "ko",
	"Limburgish":          "li",
	"Malagasy":            "mg",
	"Chinese":             "zh",
	"French":              "fr",
	"Haitian":             "ht",
	"Italian":             "it",
	"Hebrew":              "he",
	"Pāli":                "pi",
	"Yiddish":             "yi",
	"Kurdish":             "ku",
	"Samoan":              "sm",
	"Somali":              "so",
	"Uyghur":              "ug",
	"Burmese":             "my",
	"Norwegian":           "no",
	"Persian":             "fa",
	"Scottish Gaelic":     "gd",
	"Azerbaijani":         "az",
	"Corsican":            "co",
	"Luba-Katanga":        "lu",
	"Swedish":             "sv",
	"Gujarati":            "gu",
	"Hungarian":           "hu",
	"Indonesian":          "id",
	"Javanese":            "jv",
	"Catalan":             "ca",
	"Czech":               "cs",
	"Dutch":               "nl",
	"Estonian":            "et",
	"Nepali":              "ne",
	"Polish":              "pl",
	"Swahili":             "sw",
	"Cornish":             "kw",
	"Danish":              "da",
	"Kinyarwanda":         "rw",
	"Urdu":                "ur",
	"Yoruba":              "yo",
	"Arabic":              "ar",
	"Kwanyama":            "kj",
	"Sindhi":              "sd",
	"Ukrainian":           "uk",
	"Chamorro":            "ch",
	"Malayalam":           "ml",
	"Portuguese":          "pt",
	"Tagalog":             "tl",
	"Lao":                 "lo",
	"Nuosu":               "ii",
	"Aragonese":           "an",
	"Chuvash":             "cv",
	"Croatian":            "hr",
	"Icelandic":           "is",
	"Malay":               "ms",
	"Māori":               "mi",
	"Southern Ndebele":    "nr",
	"Tatar":               "tt",
	"Turkish":             "tr",
	"Uzbek":               "uz",
	"Norwegian Nynorsk":   "nn",
	"Finnish":             "fi",
	"Basque":              "eu",
	"Chichewa":            "ny",
	"Serbian":             "sr",
	"Southern Sotho":      "st",
	"Komi":                "kv",
	"Luxembourgish":       "lb",
	"Norwegian Bokmål":    "nb",
	"Walloon":             "wa",
	"Old Church Slavonic": "cu",
	"Slovak":              "sk",
	"Armenian":            "hy",
	"Cree":                "cr",
	"Ndonga":              "ng",
	"Ojibwe":              "oj",
	"Xhosa":               "xh",
	"Bulgarian":           "bg",
	"Herero":              "hz",
	"Ido":                 "io",
	"Vietnamese":          "vi",
	"Bengali":             "bn",
	"Kikuyu":              "ki",
	"Shona":               "sn",
	"Turkmen":             "tk",
	"Amharic":             "am",
	"Avaric":              "av",
	"Kanuri":              "kr",
	"Maltese":             "mt",
	"Northern Sami":       "se",
	"Tamil":               "ta",
	"Kannada":             "kn",
	"Tahitian":            "ty",
	"Panjabi":             "pa",
	"Avestan":             "ae",
	"Georgian":            "ka",
	"Kazakh":              "kk",
	"Macedonian":          "mk",
	"Chechen":             "ce",
	"English":             "en",
	"Fula":                "ff",
	"Igbo":                "ig",
}

var SupportedLanguageCodesRu = map[string]string{
	"Хмонг":                  "hmn",
	"Иврит (Старый)":         "iw",
	"Идиш":                   "yi",
	"Илоко":                  "il",
	"Люксембургский":         "lb",
	"Пали":                   "pi",
	"Греческий":              "el",
	"Персидский":             "fa",
	"Португальский":          "pt",
	"Чаморро":                "ch",
	"Галисийский":            "gl",
	"Грузинский":             "ka",
	"Итальянский":            "it",
	"Каталонская":            "ca",
	"Румынский":              "ro",
	"Шона":                   "sn",
	"Южный Ндебеле":          "nr",
	"Македонский":            "mk",
	"Малаялам":               "ml",
	"Африкаанс":              "af",
	"Люба-Катанга":           "lu",
	"Ндонга":                 "ng",
	"Немецкий":               "de",
	"Норвежский Букмольский": "nb",
	"Норвежский Нюнорскский": "nn",
	"Самоанский":             "sm",
	"Туркменский":            "tk",
	"Боснийский":             "bs",
	"Гереро":                 "hz",
	"Йоруба":                 "yo",
	"Конго":                  "kg",
	"Польский":               "pl",
	"Суахили":                "sw",
	"Маори":                  "mi",
	"Северные Саамы":         "se",
	"Синдский":               "sd",
	"Узбекский":              "uz",
	"Валлонский":             "wa",
	"Фула":                   "ff",
	"Кри":                    "cr",
	"Курдский":               "ku",
	"Оджибве":                "oj",
	"Сингальский":            "si",
	"Армянский":              "hy",
	"Западно-Фризский":       "fy",
	"Коми":                   "kv",
	"Корсиканский":           "co",
	"Навахо":                 "nv",
	"Непальцы":               "ne",
	"Панджаби":               "pa",
	"Бенгальский":            "bn",
	"Бирманская":             "my",
	"Латинский":              "la",
	"Мальтийский":            "mt",
	"Французский":            "fr",
	"Хири Моту":              "ho",
	"Голландский":            "nl",
	"Лаосский":               "lo",
	"Маршалльский":           "mh",
	"Аварский":               "av",
	"Арабский":               "ar",
	"Валлийский":             "cy",
	"Сунданский":             "su",
	"Тагальский":             "tl",
	"Южный Сото":             "st",
	"Болгарский":             "bg",
	"Гаитянский":             "ht",
	"Эсперанто":              "eo",
	"Кикуйю":                 "ki",
	"Киньяруанда":            "rw",
	"Киргизский":             "ky",
	"Коса":                   "xh",
	"Уйгурский":              "ug",
	"Хауса":                  "ha",
	"Урду":                   "ur",
	"Шотландский Гэльский":   "gd",
	"Гуджарати":              "gu",
	"Игбо":                   "ig",
	"Китайский Гонконгский":  "zh-hk",
	"Китайский Тайваньский":  "zh-tw",
	"Нуосу":                  "ii",
	"Словацкий":              "sk",
	"Эстонский":              "et",
	"Китайский КНР":          "zh-cn",
	"Лимбургский":            "li",
	"Маратхи":                "mr",
	"Ория":                   "or",
	"Старославянский":        "cu",
	"Турецкий":               "tr",
	"Канури":                 "kr",
	"Малайский":              "ms",
	"Тайский":                "th",
	"Корнуоллский":           "kw",
	"Сербский":               "sr",
	"Авестийский":            "ae",
	"Английский":             "en",
	"Кваньяма":               "kj",
	"Малагасийский":          "mg",
	"Сардинский":             "sc",
	"Украинский":             "uk",
	"Сомалийский":            "so",
	"Арагонский":             "an",
	"Ирландский":             "ga",
	"Пушту":                  "ps",
	"Таитянский":             "ty",
	"Телугу":                 "te",
	"Чичева":                 "ny",
	"Датский":                "da",
	"Индонезийский":          "id",
	"Исландский":             "is",
	"Китайский":              "zh",
	"Монгольский":            "mn",
	"Норвежский":             "no",
	"Словенский":             "sl",
	"Хорватский":             "hr",
	"Чеченская":              "ce",
	"Чувашская":              "cv",
	"Японский":               "ja",
	"Латышский":              "lv",
	"Литовский":              "lt",
	"Азербайджанский":        "az",
	"Казахский":              "kk",
	"Яванский":               "jv",
	"Вьетнамский":            "vi",
	"Каннада":                "kn",
	"Тамильский":             "ta",
	"Финский":                "fi",
	"Чешский":                "cs",
	"Амхарский":              "am",
	"Белорусский":            "be",
	"Русский":                "ru",
	"Баскский":               "eu",
	"Венгерский":             "hu",
	"Иврит":                  "he",
	"Испанский":              "es",
	"Корейский":              "ko",
	"Таджикский":             "tg",
	"Шведский":               "sv",
	"Кхмерский":              "km",
	"Северные Ндебеле":       "nd",
	"Хинди":                  "hi",
	"Идо":                    "io",
	"Китайский Сингапурский": "zh-sg",
	"Татарский":              "tt",
}

var SupportedLanguageCodesZh = map[string]string{
	"哼哼":        "hmn",
	"希伯来语（旧）":   "iw",
	"Hiri Motu": "ho",
	"伊洛科":       "il",
	"乌克兰语":      "uk",
	"基库尤语":      "ki",
	"恩东加人":      "ng",
	"查莫罗语":      "ch",
	"格鲁吉亚语":     "ka",
	"法语":        "fr",
	"爱沙尼亚语":     "et",
	"英语":        "en",
	"印度尼西亚语":    "id",
	"古吉拉特语":     "gu",
	"爱尔兰语":      "ga",
	"诺苏语":       "ii",
	"乌尔都语":      "ur",
	"南非荷兰语":     "af",
	"哈萨克语":      "kk",
	"塞尔维亚语":     "sr",
	"意大利语":      "it",
	"蒙古人":       "mn",
	"乌兹别克语":     "uz",
	"南恩德贝莱":     "nr",
	"爪哇语":       "jv",
	"老挝语":       "lo",
	"刚果语":       "kg",
	"巴斯克语":      "eu",
	"索马里语":      "so",
	"约鲁巴语":      "yo",
	"维吾尔语":      "ug",
	"马拉雅拉姆语":    "ml",
	"Herero":    "hz",
	"世界语":       "eo",
	"伊博语":       "ig",
	"南索托语":      "st",
	"奥里亚语":      "or",
	"希腊语":       "el",
	"拉脱维亚语":     "lv",
	"白俄罗斯语":     "be",
	"中國香港":      "zh-hk",
	"捷克语":       "cs",
	"瑞典语":       "sv",
	"阿拉贡语":      "an",
	"马耳他语":      "mt",
	"希伯来语":      "he",
	"泰米尔语":      "ta",
	"马达加斯加语":    "mg",
	"他加禄语":      "tl",
	"关山语":       "kj",
	"意第绪语":      "yi",
	"拉丁语":       "la",
	"罗马尼亚语":     "ro",
	"荷兰语":       "nl",
	"伊多语":       "io",
	"塔吉克语":      "tg",
	"奇切瓦语":      "ny",
	"斯洛文尼亚语":    "sl",
	"芬兰语":       "fi",
	"西弗里斯兰语":    "fy",
	"土库曼语":      "tk",
	"大溪地语":      "ty",
	"撒丁语":       "sc",
	"新加坡汉语":     "zh-sg",
	"阿拉伯语":      "ar",
	"信德语":       "sd",
	"威尔士语":      "cy",
	"科西嘉语":      "co",
	"马绍尔人":      "mh",
	"中國中華人民共和國": "zh-cn",
	"僧伽罗语":      "si",
	"巽他语":       "su",
	"挪威博克马尔语":   "nb",
	"斯瓦希里语":     "sw",
	"毛利人":       "mi",
	"波兰语":       "pl",
	"冰岛语":       "is",
	"匈牙利语":      "hu",
	"海地语":       "ht",
	"立陶宛语":      "lt",
	"印地语":       "hi",
	"普什图语":      "ps",
	"缅甸语":       "my",
	"马来语":       "ms",
	"丹麦语":       "da",
	"克里语":       "cr",
	"加泰罗尼亚语":    "ca",
	"土耳其语":      "tr",
	"西班牙语":      "es",
	"北恩德贝莱人":    "nd",
	"卡纳达语":      "kn",
	"孟加拉语":      "bn",
	"古教堂斯拉夫语":   "cu",
	"越南语":       "vi",
	"亚美尼亚语":     "hy",
	"卢巴-加丹加语":   "lu",
	"科萨语":       "xh",
	"豪萨语":       "ha",
	"中國澳門":      "zh-mo",
	"俄语":        "ru",
	"克罗地亚语":     "hr",
	"北萨米人":      "se",
	"卡努里语":      "kr",
	"康沃尔语":      "kw",
	"马其顿语":      "mk",
	"中國台灣":      "zh-tw",
	"旁遮普语":      "pa",
	"吉尔吉斯语":     "ky",
	"林堡语":       "li",
	"汉语":        "zh",
	"波斯语":       "fa",
	"车臣语":       "ce",
	"加利西亚语":     "gl",
	"富拉语":       "ff",
	"科米语":       "kv",
	"泰语":        "th",
	"韩语":        "ko",
	"尼泊尔人":      "ne",
	"巴利语":       "pi",
	"库尔德语":      "ku",
	"挪威尼诺斯克语":   "nn",
	"斯洛伐克语":     "sk",
	"波斯尼亚语":     "bs",
	"绍纳语":       "sn",
	"阿塞拜疆语":     "az",
	"鞑靼语":       "tt",
	"纳瓦霍人":      "nv",
	"挪威人":       "no",
	"楚瓦什语":      "cv",
	"奥吉布威语":     "oj",
	"德语":        "de",
	"葡萄牙语":      "pt",
	"阿维斯坦语":     "ae",
	"高棉语":       "km",
	"保加利亚语":     "bg",
	"基尼亚卢旺达语":   "rw",
	"泰卢固语":      "te",
	"瓦隆语":       "wa",
	"萨摩亚语":      "sm",
	"阿姆哈拉语":     "am",
	"马拉地语":      "mr",
	"卢森堡语":      "lb",
	"日语":        "ja",
	"苏格兰盖尔语":    "gd",
	"阿瓦里语":      "av",
}
