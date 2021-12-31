package main

import "github.com/jinzhu/gorm"

var neighborhoodsMap = map[string]map[string]interface{}{
	"'s-Lands Werf": {
		"key": "s_lands_werf",
	},
	"1e Loosterweg": {
		"key": "1e_loosterweg",
	},
	"A.J. Koejemansstraat": {
		"key": "aj_koejemansstraat",
	},
	"Wattbaan": {
		"key": "wattbaan",
	},
	"Aaf Bouberstraat": {
		"key": "aaf_bouberstraat",
	},
	"Aart van der Leeuwkade": {
		"key": "aart_van_der_leeuwkade",
	},
	"Abel Tasmanstraat": {
		"key": "abel_tasmanstraat",
	},
	"Adelaarstraat": {
		"key": "adelaarstraat",
	},
	"Adriaen van Ostadelaan": {
		"key": "adriaen_van_ostadelaan",
	},
	"Albatrosstraat": {
		"key": "albatrosstraat",
	},
	"Allard Piersonlaan": {
		"key": "allard_piersonlaan",
	},
	"Amsterdamsestraatweg": {
		"key": "amsterdamsestraatweg",
	},
	"Anton de Haenstraat": {
		"key": "anton_de_haenstraat",
	},
	"Archimedesstraat": {
		"key": "archimedesstraat",
	},
	"Arienslaan": {
		"key": "arienslaan",
	},
	"Assumburgweg": {
		"key": "assumburgweg",
	},
	"Bagijnhof": {
		"key": "bagijnhof",
	},
	"Beierlandsestraat": {
		"key": "beierlandsestraat",
	},
	"Benjamin Willem ter Kuilestraat": {
		"key": "benjamin_willem_ter_kuilestraat",
	},
	"Bickersgracht": {
		"key": "bickersgracht",
	},
	"Bijlmerdreef": {
		"key": "bijlmerdreef",
	},
	"Binckhorstlaan": {
		"key": "binckhorstlaan",
	},
	"Bloklandstraat": {
		"key": "bloklandstraat",
	},
	"Boekweitstraat": {
		"key": "boekweitstraat",
	},
	"Boompjes": {
		"key": "boompjes",
	},
	"Borgerstraat": {
		"key": "borgerstraat",
	},
	"Boterstraat": {
		"key": "boterstraat",
	},
	"Botterstraat": {
		"key": "botterstraat",
	},
	"Buizerdhof": {
		"key": "buizerdhof",
	},
	"Calandkade": {
		"key": "calandkade",
	},
	"Catharinaland": {
		"key": "catharinaland",
	},
	"Crynssenstraat": {
		"key": "crynssenstraat",
	},
	"Daendelsstraat": {
		"key": "daendelsstraat",
	},
	"Damasstraat": {
		"key": "damasstraat",
	},
	"Dautzenbergstraat": {
		"key": "dautzenbergstraat",
	},
	"De Haasstraat": {
		"key": "de_haasstraat",
	},
	"Dedemsvaartweg": {
		"key": "dedemsvaartweg",
	},
	"Derde Kostverlorenkade": {
		"key": "derde_kostverlorenkade",
	},
	"Dikninge": {
		"key": "dikninge",
	},
	"Ditlaar": {
		"key": "ditlaar",
	},
	"Dordtselaan": {
		"key": "dordtselaan",
	},
	"Dorpsstraat": {
		"key": "dorpsstraat",
	},
	"Drebbelstraat": {
		"key": "drebbelstraat",
	},
	"Droogbak": {
		"key": "droogbak",
	},
	"Eisenhowerlaan": {
		"key": "eisenhowerlaan",
	},
	"Elandsgracht": {
		"key": "elandsgracht",
	},
	"Emmalaan": {
		"key": "emmalaan",
	},
	"Franselaan": {
		"key": "franselaan",
	},
	"Gerard Doustraat": {
		"key": "gerard_doustraat",
	},
	"Goeverneurlaan": {
		"key": "goeverneurlaan",
	},
	"Graaf Balderikstraat": {
		"key": "graaf_balderikstraat",
	},
	"Grote Bickersstraat": {
		"key": "grote_bickersstraat",
	},
	"Grote Visserijstraat": {
		"key": "grote_visserijstraat",
	},
	"H.J. Schimmelplein": {
		"key": "hj_schimmelplein",
	},
	"Halleystraat": {
		"key": "halleystraat",
	},
	"Hanoistraat": {
		"key": "hanoistraat",
	},
	"Haparandaweg": {
		"key": "haparandaweg",
	},
	"Harstenhoekweg": {
		"key": "harstenhoekweg",
	},
	"Heathrowstraat": {
		"key": "heathrowstraat",
	},
	"Heeswijkstraat": {
		"key": "heeswijkstraat",
	},
	"Hekbootstraat": {
		"key": "hekbootstraat",
	},
	"Helperzoom": {
		"key": "helperzoom",
	},
	"Hendrikje Stoffelsstraat": {
		"key": "hendrikje_stoffelsstraat",
	},
	"Herenweg": {
		"key": "herenweg",
	},
	"Herman Gorterhof": {
		"key": "herman_gorterhof",
	},
	"Het Hout": {
		"key": "het_hout",
	},
	"Hoofdstraat": {
		"key": "hoofdstraat",
	},
	"Hoofdweg": {
		"key": "hoofdweg",
	},
	"Hooglandstraat": {
		"key": "hooglandstraat",
	},
	"Hoogoord": {
		"key": "hoogoord",
	},
	"Huis te Landelaan": {
		"key": "huis_te_landelaan",
	},
	"Isa C Hubertstraat": {
		"key": "isa_c_hubertstraat",
	},
	"J.J. van Deinselaan": {
		"key": "jj_van_deinselaan",
	},
	"Jaap Edendreef": {
		"key": "jaap_edendreef",
	},
	"Jan van Galenstraat": {
		"key": "jan_van_galenstraat",
	},
	"Jan van Scorelstraat": {
		"key": "jan_van_scorelstraat",
	},
	"Jisperveldstraat": {
		"key": "jisperveldstraat",
	},
	"Johan Wagenaarlaan": {
		"key": "johan_wagenaarlaan",
	},
	"Johan van Hasseltkade": {
		"key": "johan_van_hasseltkade",
	},
	"John M. Keynesplein": {
		"key": "john_m_keynesplein",
	},
	"Jonkvrouw Sanderijndreef": {
		"key": "jonkvrouw_sanderijndreef",
	},
	"Juliana van Stolberglaan": {
		"key": "juliana_van_stolberglaan",
	},
	"Kalverstraat": {
		"key": "kalverstraat",
	},
	"Kamelenspoor": {
		"key": "kamelenspoor",
	},
	"Karwijzaaderf": {
		"key": "karwijzaaderf",
	},
	"Koekoekweg": {
		"key": "koekoekweg",
	},
	"Koesteeg": {
		"key": "koesteeg",
	},
	"Koningin Emmaplein": {
		"key": "koningin_emmaplein",
	},
	"Koninginnegracht": {
		"key": "koninginnegracht",
	},
	"Koningsweg": {
		"key": "koningsweg",
	},
	"Kritostraat": {
		"key": "kritostraat",
	},
	"Kruidenhof": {
		"key": "kruidenhof",
	},
	"Laan van Nieuw Guinea": {
		"key": "laan_van_nieuw_guinea",
	},
	"Langestraat": {
		"key": "langestraat",
	},
	"Lindenweg": {
		"key": "lindenweg",
	},
	"Looiershof": {
		"key": "looiershof",
	},
	"Maaswijkstraat": {
		"key": "maaswijkstraat",
	},
	"Marco Pololaan": {
		"key": "marco_pololaan",
	},
	"Marijkelaan": {
		"key": "marijkelaan",
	},
	"Markelerbergpad": {
		"key": "markelerbergpad",
	},
	"Marktweg": {
		"key": "marktweg",
	},
	"Marnixstraat": {
		"key": "marnixstraat",
	},
	"Media Park Blvd": {
		"key": "media_park_blvd",
	},
	"Mijnsherenlaan": {
		"key": "mijnsherenlaan",
	},
	"NDSM-Pier": {
		"key": "ndsm_pier",
	},
	"Nieuwe Binnenweg": {
		"key": "nieuwe_binnenweg",
	},
	"Nieuwezijds Voorburgwal": {
		"key": "nieuwezijds_voorburgwal",
	},
	"Noordeinde": {
		"key": "noordeinde",
	},
	"Noordermarkt": {
		"key": "noordermarkt",
	},
	"Oldenzaalsestraat": {
		"key": "oldenzaalsestraat",
	},
	"Oosteinderlaan": {
		"key": "oosteinderlaan",
	},
	"Oosterschelde": {
		"key": "oosterschelde",
	},
	"Oostmaaslaan": {
		"key": "oostmaaslaan",
	},
	"Oppenheimstraat": {
		"key": "oppenheimstraat",
	},
	"Oranjeboomstraat": {
		"key": "oranjeboomstraat",
	},
	"Osdorper Ban": {
		"key": "osdorper_ban",
	},
	"Ostadestraat": {
		"key": "ostadestraat",
	},
	"Oudegracht": {
		"key": "oudegracht",
	},
	"Oudenoord": {
		"key": "oudenoord",
	},
	"Oudezijds Achterburgwal": {
		"key": "oudezijds_achterburgwal",
	},
	"Oudezijds Voorburgwal": {
		"key": "oudezijds_voorburgwal",
	},
	"Padangstraat": {
		"key": "padangstraat",
	},
	"Palmstraat": {
		"key": "palmstraat",
	},
	"Peppelweg": {
		"key": "peppelweg",
	},
	"Pizarrolaan": {
		"key": "pizarrolaan",
	},
	"Pleinweg": {
		"key": "pleinweg",
	},
	"Poolseweg": {
		"key": "poolseweg",
	},
	"Poolsterplein": {
		"key": "poolsterplein",
	},
	"Postjeskade": {
		"key": "postjeskade",
	},
	"Potvisstraat": {
		"key": "potvisstraat",
	},
	"Pretoriusstraat": {
		"key": "pretoriusstraat",
	},
	"Prins Hendrikkade": {
		"key": "prins_hendrikkade",
	},
	"Prinsengracht": {
		"key": "prinsengracht",
	},
	"Prinsenlaan": {
		"key": "prinsenlaan",
	},
	"Professor J.H. Bavincklaan": {
		"key": "professor_jh_bavincklaan",
	},
	"Reguliersteeg": {
		"key": "reguliersteeg",
	},
	"Reyer Anslostraat": {
		"key": "reyer_anslostraat",
	},
	"Rijndijk": {
		"key": "rijndijk",
	},
	"Rijndijkstraat": {
		"key": "rijndijkstraat",
	},
	"Robijnlaan": {
		"key": "robijnlaan",
	},
	"Rotterdamsestraat": {
		"key": "rotterdamsestraat",
	},
	"Saffraanweg": {
		"key": "saffraanweg",
	},
	"Schalkeroord": {
		"key": "schalkeroord",
	},
	"Schieweg": {
		"key": "schieweg",
	},
	"Schinnenbaan": {
		"key": "schinnenbaan",
	},
	"Schrijverspark": {
		"key": "schrijverspark",
	},
	"Schuitenweg": {
		"key": "schuitenweg",
	},
	"Shetlands": {
		"key": "shetlands",
	},
	"Sierplein": {
		"key": "sierplein",
	},
	"Sir Winston Churchillweg": {
		"key": "sir_winston_churchillweg",
	},
	"Slotlaan": {
		"key": "slotlaan",
	},
	"Smidswater": {
		"key": "smidswater",
	},
	"Spaarndammerdijk": {
		"key": "spaarndammerdijk",
	},
	"Stationssingel": {
		"key": "stationssingel",
	},
	"Strevelsweg": {
		"key": "strevelsweg",
	},
	"Teakhout": {
		"key": "teakhout",
	},
	"Terschellingsestraat": {
		"key": "terschellingsestraat",
	},
	"Teylersplein": {
		"key": "teylersplein",
	},
	"Theeroos": {
		"key": "theeroos",
	},
	"Tolstraat": {
		"key": "tolstraat",
	},
	"Traaij": {
		"key": "traaij",
	},
	"Troelstrakade": {
		"key": "troelstrakade",
	},
	"Truffautstraat": {
		"key": "truffautstraat",
	},
	"Twijnstraat": {
		"key": "twijnstraat",
	},
	"Utrechtseweg": {
		"key": "utrechtseweg",
	},
	"Valkenburgerstraat": {
		"key": "valkenburgerstraat",
	},
	"Van 't Hoffplein": {
		"key": "van_t_hoffplein",
	},
	"Van Alkemadelaan": {
		"key": "van_alkemadelaan",
	},
	"Van Alkemadestraat": {
		"key": "van_alkemadestraat",
	},
	"Van Hogendorpstraat": {
		"key": "van_hogendorpstraat",
	},
	"Van Kootenstraat": {
		"key": "van_kootenstraat",
	},
	"Van Naeltwijckstraat": {
		"key": "van_naeltwijckstraat",
	},
	"Van der Lelijstraat": {
		"key": "van_der_lelijstraat",
	},
	"Van der Meydestraat": {
		"key": "van_der_meydestraat",
	},
	"Veerpolder": {
		"key": "veerpolder",
	},
	"Vleutenseweg": {
		"key": "vleutenseweg",
	},
	"Volmerlaan": {
		"key": "volmerlaan",
	},
	"Von Geusaustraat": {
		"key": "von_geusaustraat",
	},
	"Voorstraat": {
		"key": "voorstraat",
	},
	"Vredenburg": {
		"key": "vredenburg",
	},
	"Wamelstraat": {
		"key": "wamelstraat",
	},
	"Watertorenweg": {
		"key": "watertorenweg",
	},
	"Weissenbruchstraat": {
		"key": "weissenbruchstraat",
	},
	"Werfkade": {
		"key": "werfkade",
	},
	"Westersingel": {
		"key": "westersingel",
	},
	"Willem Augustinstraat": {
		"key": "willem_augustinstraat",
	},
	"Willem Nakkenstraat": {
		"key": "willem_nakkenstraat",
	},
	"Willem de Zwijgerlaan": {
		"key": "willem_de_zwijgerlaan",
	},
	"Windvorst": {
		"key": "windvorst",
	},
	"Witte De Withstraat": {
		"key": "witte_de_withstraat",
	},
	"Yokohamadreef": {
		"key": "yokohamadreef",
	},
	"Zeussingel": {
		"key": "zeussingel",
	},
	"Zilverberg": {
		"key": "zilverberg",
	},
	"Zwaansteeg": {
		"key": "zwaansteeg",
	},
	"Zwijnsbergenwe": {
		"key": "zwijnsbergenweg",
	},
	"Dotterlei": {
		"key": "dotterlei",
	},
	"Zijlstraat": {
		"key": "zijlstraat",
	},
	"Electrablauw": {
		"key": "electrablauw",
	},
	"Veldweg": {
		"key": "veldweg",
	},
	"Malpertuisstraat": {
		"key": "malpertuisstraat",
	},
	"Domselaerstraat": {
		"key": "domselaerstraat",
	},
	"Eerste Weteringdwarsstraat": {
		"key": "eerste_weteringdwarsstraat",
	},
	"Rustenburgerstraat": {
		"key": "rustenburgerstraat",
	},
	"oude Spiegelstraat": {
		"key": "oude_spiegelstraat",
	},
	"Weesperzijde": {
		"key": "weesperzijde",
	},
	"Van der Boechorststraat": {
		"key": "van_der_boechorststraat",
	},
	"De Genestetlaan": {
		"key": "de_genestetlaan",
	},
	"Donkerslootstraat": {
		"key": "donkerslootstraat",
	},
	"Hofsingel": {
		"key": "hofsingel",
	},
	"Nieuwstraat": {
		"key": "nieuwstraat",
	},
	"Over de Vesten": {
		"key": "over_de_vesten",
	},
	"Keizersgracht": {
		"key": "keizersgracht",
	},
	"Maria Snelplantsoen": {
		"key": "maria_snelplantsoen",
	},
	"Gerrit Mannourystraat": {
		"key": "gerrit_mannourystraat",
	},
	"Tenerifestraat": {
		"key": "tenerifestraat",
	},
	"Boterdiep": {
		"key": "boterdiep",
	},
	"Grote Beerstraat": {
		"key": "grote_beerstraat",
	},
	"Generaal van Geenplein": {
		"key": "generaal_van_geenplein",
	},

	"raadhuisstraat 26": {
		"key": "raadhuisstraat",
	},
	"jvan wassenaerstraat 11": {
		"key": "jvan_wassenaerstraat",
	},
	"arendstraat 15": {
		"key": "arendstraat",
	},
	"buntgras 51": {
		"key": "buntgras",
	},
	"dorpsstraat 600": {
		"key": "dorpsstraat-koop",
	},
	"prins frederikdreef 4": {
		"key": "prins_frederikdreef",
	},
	"hoge markt 28": {
		"key": "hoge_markt",
	},
	"dries 5": {
		"key": "dries",
	},
	"rontgenlaan 18": {
		"key": "rontgenlaan",
	},
	"kerkstraat 13": {
		"key": "kerkstraat",
	},
	"volantestraat 41": {
		"key": "volantestraat",
	},
	"de malthezer 51": {
		"key": "de_malthezer",
	},
	"rietpolderlaan 60": {
		"key": "rietpolderlaan",
	},
	"oudendijk 51": {
		"key": "oudendijk",
	},
	"trambaan 6": {
		"key": "trambaan",
	},
	"graaf janstraat 55": {
		"key": "graaf_janstraat",
	},
	"rustenborchdreef 106": {
		"key": "rustenborchdreef",
	},
	"werfstraat 41": {
		"key": "werfstraat",
	},
	"windbaan 58": {
		"key": "windbaan",
	},
	"stopperknoop 44": {
		"key": "stopperknoop",
	},
	"paramariboplein 56 3": {
		"key": "paramariboplein",
	},
	"admiralengracht 183 3": {
		"key": "admiralengracht",
	},
	"admiralengracht 183 i": {
		"key": "admiralengracht",
	},
	"kanaalweg oostzijde 22": {
		"key": "kanaalweg_oostzijde",
	},
	"veilingweg 38": {
		"key": "veilingweg",
	},
	"dorpsstraat 15 b": {
		"key": "dorpsstraat-koop",
	},
	"klaproosstraat 14": {
		"key": "klaproosstraat",
	},
	"osdorper ban 703 pp": {
		"key": "osdorper_ban-koop",
	},
	"kroonkruid 165": {
		"key": "kroonkruid",
	},
	"in de camp 18": {
		"key": "in_de_camp",
	},
	"dr kuyperkade 5 rd": {
		"key": "dr_kuyperkade",
	},
	"commelinstraat 10 3": {
		"key": "commelinstraat",
	},
}

func setupNeighborhoodsMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range neighborhoodsMap {
		DB.Raw("SELECT * FROM `neighbourhoods`  WHERE (`key` = ?) ORDER BY `neighbourhoods`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		neighborhoodsMap[name] = data
	}
}

type Neighbourhood struct {
	gorm.Model
	Key    string `json:"key" gorm:"type:varchar(50);index;unique"`
	Name   string `json:"name" gorm:"type:varchar(50);index;unique"`
	CityId int    `json:"city_id" gorm:"type:integer;index"`
}
