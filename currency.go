package currency

import (
	"errors"
	"fmt"
	"strings"
)

type Currency struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	MinorUnits int    `json:"minor_units"`
}

var currencies = map[string]Currency{
	"AED": {"AED", "United Arab Emirates Dirham", "د.إ", 2},
	"AFN": {"AFN", "Afghan Afghani", "؋", 2},
	"ALL": {"ALL", "Albanian Lek", "L", 2},
	"AMD": {"AMD", "Armenian Dram", "֏", 2},
	"ANG": {"ANG", "Netherlands Antillean Guilder", "ƒ", 2},
	"AOA": {"AOA", "Angolan Kwanza", "Kz", 2},
	"ARS": {"ARS", "Argentine Peso", "$", 2},
	"AUD": {"AUD", "Australian Dollar", "A$", 2},
	"AWG": {"AWG", "Aruban Florin", "ƒ", 2},
	"AZN": {"AZN", "Azerbaijani Manat", "₼", 2},
	"BAM": {"BAM", "Bosnia and Herzegovina Convertible Mark", "KM", 2},
	"BBD": {"BBD", "Barbadian Dollar", "Bds$", 2},
	"BDT": {"BDT", "Bangladeshi Taka", "৳", 2},
	"BGN": {"BGN", "Bulgarian Lev", "лв", 2},
	"BHD": {"BHD", "Bahraini Dinar", ".د.ب", 3},
	"BIF": {"BIF", "Burundian Franc", "FBu", 0},
	"BMD": {"BMD", "Bermudian Dollar", "$", 2},
	"BND": {"BND", "Brunei Dollar", "B$", 2},
	"BOB": {"BOB", "Bolivian Boliviano", "Bs.", 2},
	"BRL": {"BRL", "Brazilian Real", "R$", 2},
	"BSD": {"BSD", "Bahamian Dollar", "B$", 2},
	"BTN": {"BTN", "Bhutanese Ngultrum", "Nu.", 2},
	"BWP": {"BWP", "Botswana Pula", "P", 2},
	"BYN": {"BYN", "Belarusian Ruble", "Br", 2},
	"BZD": {"BZD", "Belize Dollar", "BZ$", 2},
	"CAD": {"CAD", "Canadian Dollar", "C$", 2},
	"CDF": {"CDF", "Congolese Franc", "FC", 2},
	"CHF": {"CHF", "Swiss Franc", "Fr", 2},
	"CLP": {"CLP", "Chilean Peso", "$", 0},
	"CNY": {"CNY", "Chinese Yuan", "¥", 2},
	"COP": {"COP", "Colombian Peso", "$", 2},
	"CRC": {"CRC", "Costa Rican Colón", "₡", 2},
	"CUC": {"CUC", "Cuban Convertible Peso", "$", 2},
	"CUP": {"CUP", "Cuban Peso", "₱", 2},
	"CVE": {"CVE", "Cape Verdean Escudo", "$", 2},
	"CZK": {"CZK", "Czech Koruna", "Kč", 2},
	"DJF": {"DJF", "Djiboutian Franc", "Fdj", 0},
	"DKK": {"DKK", "Danish Krone", "kr", 2},
	"DOP": {"DOP", "Dominican Peso", "RD$", 2},
	"DZD": {"DZD", "Algerian Dinar", "د.ج", 2},
	"EGP": {"EGP", "Egyptian Pound", "£", 2},
	"ERN": {"ERN", "Eritrean Nakfa", "Nfk", 2},
	"ETB": {"ETB", "Ethiopian Birr", "Br", 2},
	"EUR": {"EUR", "Euro", "€", 2},
	"FJD": {"FJD", "Fijian Dollar", "FJ$", 2},
	"FKP": {"FKP", "Falkland Islands Pound", "£", 2},
	"FOK": {"FOK", "Faroese Króna", "kr", 2},
	"GBP": {"GBP", "British Pound Sterling", "£", 2},
	"GEL": {"GEL", "Georgian Lari", "₾", 2},
	"GGP": {"GGP", "Guernsey Pound", "£", 2},
	"GHS": {"GHS", "Ghanaian Cedi", "₵", 2},
	"GIP": {"GIP", "Gibraltar Pound", "£", 2},
	"GMD": {"GMD", "Gambian Dalasi", "D", 2},
	"GNF": {"GNF", "Guinean Franc", "FG", 0},
	"GTQ": {"GTQ", "Guatemalan Quetzal", "Q", 2},
	"GYD": {"GYD", "Guyanese Dollar", "G$", 2},
	"HKD": {"HKD", "Hong Kong Dollar", "HK$", 2},
	"HNL": {"HNL", "Honduran Lempira", "L", 2},
	"HRK": {"HRK", "Croatian Kuna", "kn", 2},
	"HTG": {"HTG", "Haitian Gourde", "G", 2},
	"HUF": {"HUF", "Hungarian Forint", "Ft", 2},
	"IDR": {"IDR", "Indonesian Rupiah", "Rp", 2},
	"ILS": {"ILS", "Israeli New Shekel", "₪", 2},
	"INR": {"INR", "Indian Rupee", "₹", 2},
	"IQD": {"IQD", "Iraqi Dinar", "ع.د", 3},
	"IRR": {"IRR", "Iranian Rial", "﷼", 2},
	"ISK": {"ISK", "Icelandic Króna", "kr", 0},
	"JMD": {"JMD", "Jamaican Dollar", "J$", 2},
	"JOD": {"JOD", "Jordanian Dinar", "د.ا", 3},
	"JPY": {"JPY", "Japanese Yen", "¥", 0},
	"KES": {"KES", "Kenyan Shilling", "KSh", 2},
	"KGS": {"KGS", "Kyrgyzstani Som", "с", 2},
	"KHR": {"KHR", "Cambodian Riel", "៛", 2},
	"KMF": {"KMF", "Comorian Franc", "CF", 0},
	"KPW": {"KPW", "North Korean Won", "₩", 2},
	"KRW": {"KRW", "South Korean Won", "₩", 0},
	"KWD": {"KWD", "Kuwaiti Dinar", "د.ك", 3},
	"KYD": {"KYD", "Cayman Islands Dollar", "CI$", 2},
	"KZT": {"KZT", "Kazakhstani Tenge", "₸", 2},
	"LAK": {"LAK", "Lao Kip", "₭", 2},
	"LBP": {"LBP", "Lebanese Pound", "ل.ل", 2},
	"LKR": {"LKR", "Sri Lankan Rupee", "Rs", 2},
	"LRD": {"LRD", "Liberian Dollar", "L$", 2},
	"LSL": {"LSL", "Lesotho Loti", "L", 2},
	"LYD": {"LYD", "Libyan Dinar", "ل.د", 3},
	"MAD": {"MAD", "Moroccan Dirham", "د.م.", 2},
	"MDL": {"MDL", "Moldovan Leu", "L", 2},
	"MGA": {"MGA", "Malagasy Ariary", "Ar", 2},
	"MKD": {"MKD", "Macedonian Denar", "ден", 2},
	"MMK": {"MMK", "Myanmar Kyat", "K", 2},
	"MNT": {"MNT", "Mongolian Tögrög", "₮", 2},
	"MOP": {"MOP", "Macanese Pataca", "P", 2},
	"MRU": {"MRU", "Mauritanian Ouguiya", "UM", 2},
	"MUR": {"MUR", "Mauritian Rupee", "₨", 2},
	"MVR": {"MVR", "Maldivian Rufiyaa", "Rf", 2},
	"MWK": {"MWK", "Malawian Kwacha", "MK", 2},
	"MXN": {"MXN", "Mexican Peso", "$", 2},
	"MYR": {"MYR", "Malaysian Ringgit", "RM", 2},
	"MZN": {"MZN", "Mozambican Metical", "MT", 2},
	"NAD": {"NAD", "Namibian Dollar", "N$", 2},
	"NGN": {"NGN", "Nigerian Naira", "₦", 2},
	"NIO": {"NIO", "Nicaraguan Córdoba", "C$", 2},
	"NOK": {"NOK", "Norwegian Krone", "kr", 2},
	"NPR": {"NPR", "Nepalese Rupee", "₨", 2},
	"NZD": {"NZD", "New Zealand Dollar", "NZ$", 2},
	"OMR": {"OMR", "Omani Rial", "ر.ع.", 3},
	"PAB": {"PAB", "Panamanian Balboa", "B/.", 2},
	"PEN": {"PEN", "Peruvian Sol", "S/", 2},
	"PGK": {"PGK", "Papua New Guinean Kina", "K", 2},
	"PHP": {"PHP", "Philippine Peso", "₱", 2},
	"PKR": {"PKR", "Pakistani Rupee", "₨", 2},
	"PLN": {"PLN", "Polish Złoty", "zł", 2},
	"PYG": {"PYG", "Paraguayan Guaraní", "₲", 0},
	"QAR": {"QAR", "Qatari Riyal", "ر.ق", 2},
	"RON": {"RON", "Romanian Leu", "lei", 2},
	"RSD": {"RSD", "Serbian Dinar", "дин", 2},
	"RUB": {"RUB", "Russian Ruble", "₽", 2},
	"RWF": {"RWF", "Rwandan Franc", "FRw", 0},
	"SAR": {"SAR", "Saudi Riyal", "ر.س", 2},
	"SBD": {"SBD", "Solomon Islands Dollar", "SI$", 2},
	"SCR": {"SCR", "Seychellois Rupee", "₨", 2},
	"SDG": {"SDG", "Sudanese Pound", "ج.س.", 2},
	"SEK": {"SEK", "Swedish Krona", "kr", 2},
	"SGD": {"SGD", "Singapore Dollar", "S$", 2},
	"SHP": {"SHP", "Saint Helena Pound", "£", 2},
	"SLE": {"SLE", "Sierra Leonean Leone", "Le", 2},
	"SLL": {"SLL", "Sierra Leonean Leone (old)", "Le", 2},
	"SOS": {"SOS", "Somali Shilling", "Sh", 2},
	"SRD": {"SRD", "Surinamese Dollar", "$", 2},
	"SSP": {"SSP", "South Sudanese Pound", "£", 2},
	"STN": {"STN", "São Tomé and Príncipe Dobra", "Db", 2},
	"SYP": {"SYP", "Syrian Pound", "£", 2},
	"SZL": {"SZL", "Eswatini Lilangeni", "L", 2},
	"THB": {"THB", "Thai Baht", "฿", 2},
	"TJS": {"TJS", "Tajikistani Somoni", "ЅМ", 2},
	"TMT": {"TMT", "Turkmenistani Manat", "m", 2},
	"TND": {"TND", "Tunisian Dinar", "د.ت", 3},
	"TOP": {"TOP", "Tongan Paʻanga", "T$", 2},
	"TRY": {"TRY", "Turkish Lira", "₺", 2},
	"TTD": {"TTD", "Trinidad and Tobago Dollar", "TT$", 2},
	"TVD": {"TVD", "Tuvaluan Dollar", "$", 2},
	"TWD": {"TWD", "New Taiwan Dollar", "NT$", 2},
	"TZS": {"TZS", "Tanzanian Shilling", "Sh", 2},
	"UAH": {"UAH", "Ukrainian Hryvnia", "₴", 2},
	"UGX": {"UGX", "Ugandan Shilling", "Sh", 0},
	"USD": {"USD", "United States Dollar", "$", 2},
	"UYU": {"UYU", "Uruguayan Peso", "$U", 2},
	"UZS": {"UZS", "Uzbekistani So'm", "сўм", 2},
	"VES": {"VES", "Venezuelan Bolívar Soberano", "Bs.", 2},
	"VND": {"VND", "Vietnamese Đồng", "₫", 0},
	"VUV": {"VUV", "Vanuatu Vatu", "VT", 0},
	"WST": {"WST", "Samoan Tālā", "T", 2},
	"XAF": {"XAF", "Central African CFA Franc", "FCFA", 0},
	"XCD": {"XCD", "East Caribbean Dollar", "$", 2},
	"XOF": {"XOF", "West African CFA Franc", "CFA", 0},
	"XPF": {"XPF", "CFP Franc", "₣", 0},
	"YER": {"YER", "Yemeni Rial", "﷼", 2},
	"ZAR": {"ZAR", "South African Rand", "R", 2},
	"ZMW": {"ZMW", "Zambian Kwacha", "K", 2},
	"ZWL": {"ZWL", "Zimbabwean Dollar", "Z$", 2},
}

var uniqueSymbolCurrencies = map[string]Currency{
	"د.ج":  {"DZD", "Algerian Dinar", "د.ج", 2},
	"с":    {"KGS", "Kyrgyzstani Som", "с", 2},
	"RD$":  {"DOP", "Dominican Peso", "RD$", 2},
	"CF":   {"KMF", "Comorian Franc", "CF", 0},
	"₺":    {"TRY", "Turkish Lira", "₺", 2},
	"ل.ل":  {"LBP", "Lebanese Pound", "ل.ل", 2},
	"₴":    {"UAH", "Ukrainian Hryvnia", "₴", 2},
	"Rf":   {"MVR", "Maldivian Rufiyaa", "Rf", 2},
	"UM":   {"MRU", "Mauritanian Ouguiya", "UM", 2},
	"ج.س.": {"SDG", "Sudanese Pound", "ج.س.", 2},
	"ل.د":  {"LYD", "Libyan Dinar", "ل.د", 3},
	"CFA":  {"XOF", "West African CFA Franc", "CFA", 0},
	"N$":   {"NAD", "Namibian Dollar", "N$", 2},
	"L$":   {"LRD", "Liberian Dollar", "L$", 2},
	"Rs":   {"LKR", "Sri Lankan Rupee", "Rs", 2},
	"ع.د":  {"IQD", "Iraqi Dinar", "ع.د", 3},
	"Fdj":  {"DJF", "Djiboutian Franc", "Fdj", 0},
	"R$":   {"BRL", "Brazilian Real", "R$", 2},
	"MK":   {"MWK", "Malawian Kwacha", "MK", 2},
	"lei":  {"RON", "Romanian Leu", "lei", 2},
	"֏":    {"AMD", "Armenian Dram", "֏", 2},
	"₭":    {"LAK", "Lao Kip", "₭", 2},
	"د.إ":  {"AED", "United Arab Emirates Dirham", "د.إ", 2},
	"₲":    {"PYG", "Paraguayan Guaraní", "₲", 0},
	"ЅМ":   {"TJS", "Tajikistani Somoni", "ЅМ", 2},
	"zł":   {"PLN", "Polish Złoty", "zł", 2},
	"лв":   {"BGN", "Bulgarian Lev", "лв", 2},
	"NZ$":  {"NZD", "New Zealand Dollar", "NZ$", 2},
	"₽":    {"RUB", "Russian Ruble", "₽", 2},
	"฿":    {"THB", "Thai Baht", "฿", 2},
	"VT":   {"VUV", "Vanuatu Vatu", "VT", 0},
	"MT":   {"MZN", "Mozambican Metical", "MT", 2},
	"Nfk":  {"ERN", "Eritrean Nakfa", "Nfk", 2},
	"NT$":  {"TWD", "New Taiwan Dollar", "NT$", 2},
	"Ft":   {"HUF", "Hungarian Forint", "Ft", 2},
	"FJ$":  {"FJD", "Fijian Dollar", "FJ$", 2},
	"BZ$":  {"BZD", "Belize Dollar", "BZ$", 2},
	"ден":  {"MKD", "Macedonian Denar", "ден", 2},
	"FCFA": {"XAF", "Central African CFA Franc", "FCFA", 0},
	"KSh":  {"KES", "Kenyan Shilling", "KSh", 2},
	"FRw":  {"RWF", "Rwandan Franc", "FRw", 0},
	"T$":   {"TOP", "Tongan Paʻanga", "T$", 2},
	"₪":    {"ILS", "Israeli New Shekel", "₪", 2},
	"৳":    {"BDT", "Bangladeshi Taka", "৳", 2},
	"₮":    {"MNT", "Mongolian Tögrög", "₮", 2},
	"дин":  {"RSD", "Serbian Dinar", "дин", 2},
	"CI$":  {"KYD", "Cayman Islands Dollar", "CI$", 2},
	"S/":   {"PEN", "Peruvian Sol", "S/", 2},
	"Kz":   {"AOA", "Angolan Kwanza", "Kz", 2},
	"₾":    {"GEL", "Georgian Lari", "₾", 2},
	"Q":    {"GTQ", "Guatemalan Quetzal", "Q", 2},
	"R":    {"ZAR", "South African Rand", "R", 2},
	"G":    {"HTG", "Haitian Gourde", "G", 2},
	"RM":   {"MYR", "Malaysian Ringgit", "RM", 2},
	"TT$":  {"TTD", "Trinidad and Tobago Dollar", "TT$", 2},
	"T":    {"WST", "Samoan Tālā", "T", 2},
	"m":    {"TMT", "Turkmenistani Manat", "m", 2},
	"Ar":   {"MGA", "Malagasy Ariary", "Ar", 2},
	"₦":    {"NGN", "Nigerian Naira", "₦", 2},
	"Kč":   {"CZK", "Czech Koruna", "Kč", 2},
	"₹":    {"INR", "Indian Rupee", "₹", 2},
	"៛":    {"KHR", "Cambodian Riel", "៛", 2},
	"د.ت":  {"TND", "Tunisian Dinar", "د.ت", 3},
	"G$":   {"GYD", "Guyanese Dollar", "G$", 2},
	"₡":    {"CRC", "Costa Rican Colón", "₡", 2},
	"J$":   {"JMD", "Jamaican Dollar", "J$", 2},
	"$U":   {"UYU", "Uruguayan Peso", "$U", 2},
	"FC":   {"CDF", "Congolese Franc", "FC", 2},
	"kn":   {"HRK", "Croatian Kuna", "kn", 2},
	"د.م.": {"MAD", "Moroccan Dirham", "د.م.", 2},
	"ر.س":  {"SAR", "Saudi Riyal", "ر.س", 2},
	"Db":   {"STN", "São Tomé and Príncipe Dobra", "Db", 2},
	"Z$":   {"ZWL", "Zimbabwean Dollar", "Z$", 2},
	"FBu":  {"BIF", "Burundian Franc", "FBu", 0},
	"S$":   {"SGD", "Singapore Dollar", "S$", 2},
	"A$":   {"AUD", "Australian Dollar", "A$", 2},
	"₫":    {"VND", "Vietnamese Đồng", "₫", 0},
	"Fr":   {"CHF", "Swiss Franc", "Fr", 2},
	"ر.ع.": {"OMR", "Omani Rial", "ر.ع.", 3},
	"Bds$": {"BBD", "Barbadian Dollar", "Bds$", 2},
	".د.ب": {"BHD", "Bahraini Dinar", ".د.ب", 3},
	"€":    {"EUR", "Euro", "€", 2},
	"KM":   {"BAM", "Bosnia and Herzegovina Convertible Mark", "KM", 2},
	"د.ا":  {"JOD", "Jordanian Dinar", "د.ا", 3},
	"Rp":   {"IDR", "Indonesian Rupiah", "Rp", 2},
	"₸":    {"KZT", "Kazakhstani Tenge", "₸", 2},
	"؋":    {"AFN", "Afghan Afghani", "؋", 2},
	"₵":    {"GHS", "Ghanaian Cedi", "₵", 2},
	"сўм":  {"UZS", "Uzbekistani So'm", "сўм", 2},
	"د.ك":  {"KWD", "Kuwaiti Dinar", "د.ك", 3},
	"₼":    {"AZN", "Azerbaijani Manat", "₼", 2},
	"B/.":  {"PAB", "Panamanian Balboa", "B/.", 2},
	"₣":    {"XPF", "CFP Franc", "₣", 0},
	"ر.ق":  {"QAR", "Qatari Riyal", "ر.ق", 2},
	"D":    {"GMD", "Gambian Dalasi", "D", 2},
	"FG":   {"GNF", "Guinean Franc", "FG", 0},
	"SI$":  {"SBD", "Solomon Islands Dollar", "SI$", 2},
	"Nu.":  {"BTN", "Bhutanese Ngultrum", "Nu.", 2},
	"HK$":  {"HKD", "Hong Kong Dollar", "HK$", 2},
}

// List return list of all currencies
func List() map[string]Currency {
	return currencies
}

// UniqueSymbolList return list of unique currencies symbols
func UniqueSymbolList() map[string]Currency {
	return uniqueSymbolCurrencies
}

// FindCurrency search of all currency codes and returns Currency struct
//
//	curr := FindCurrency("USD")
func FindCurrency(code string) (*Currency, error) {
	upperCode := strings.ToUpper(code)

	if currency, exists := currencies[upperCode]; exists {
		return &currency, nil
	}

	return nil, fmt.Errorf("currency with code %s not found", code)
}

// ValidateCurrency validate currency code
//
//	valid := ValidateCurrency("USD")
func ValidateCurrency(code string) bool {
	_, err := FindCurrency(code)
	return err == nil
}

// FindCurrenciesBySymbol return list of Currency by symbol
//
// currencies := FindCurrenciesBySymbol("€")
func FindCurrenciesBySymbol(symbol string) []Currency {
	var result []Currency
	for _, currency := range currencies {
		if currency.Symbol == symbol {
			result = append(result, currency)
		}
	}

	return result
}

// FindUniqueSymbolCurrency returns the Currency struct for the provided currency symbol, if it is unique for that currency
//
//	curr := FindUniqueSymbolCurrency("€")
func FindUniqueSymbolCurrency(symbol string) (*Currency, error) {
	if currency, exists := uniqueSymbolCurrencies[symbol]; exists {
		return &currency, nil
	}

	return nil, errors.New(fmt.Sprintf("currency with unique symbol %s not found", symbol))
}

// HasUniqueSymbol check if Currency has unique symbol
//
// isUnique := currency.HasUniqueSymbol()
func (currency *Currency) HasUniqueSymbol() bool {
	_, err := FindUniqueSymbolCurrency(currency.Symbol)

	return err == nil
}
