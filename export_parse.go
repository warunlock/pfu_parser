	package main


import (
    "fmt"
    "io/ioutil"
    "os"
    "encoding/xml"
	"time"
	"golang.org/x/text/encoding/charmap"
)
type 	Export_Table struct {
	XMLName xml.Name `xml:"Table"`
	Text    string   `xml:",chardata"`
	Rows    struct {
		Text string `xml:",chardata"`
		Row  []struct {
			Text                         string `xml:",chardata"`
			WICNUM                       string `xml:"WIC_NUM"`
			WICCASENUM                   string `xml:"WIC_CASE_NUM"`
			WICDTBEGIN                   string `xml:"WIC_DT_BEGIN"`
			WICDTEND                     string `xml:"WIC_DT_END"`
			WICSTATUS                    string `xml:"WIC_STATUS"`
			WICCD                        string `xml:"WIC_CD"`
			WICCDName                    string `xml:"WIC_CD_Name"`
			SIGNANLKNARKOTIKINTOXICATION string `xml:"SIGN_ANLK_NARKOTIK_INTOXICATION"`
			VIOLATIONEXTENSION           string `xml:"VIOLATION_EXTENSION"`
			NPSURNAME                    string `xml:"NP_SURNAME"`
			NPNAME                       string `xml:"NP_NAME"`
			NPPATRONYMIC                 string `xml:"NP_PATRONYMIC"`
			NPNUMIDENT                   string `xml:"NP_NUMIDENT"`
		} `xml:"Row"`
	} `xml:"Rows"`
} 
func main() {

var exptbl Export_Table

	encoder := charmap.Windows1251.NewEncoder()
    // Open our xmlFile
    xmlFile, err := os.Open("export.xml")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.xml")
    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()
    byteValue, _ := ioutil.ReadAll(xmlFile)
    xml.Unmarshal(byteValue, &exptbl)
	
//   fmt.Println(exptbl.Rows.Row[1])
//    fmt.Println(exptbl)
todaym := time.Now().Format("20060102")
fname := "result"+todaym+".csv"

tmplt := "номер е-лікарняного;номер випадку непрацездатності;дата початку непрацездатності;дата закінчення непрацездатності;статус е-лікарняного;код причини непрацездатності;назва причини непрацезданості;перебування у стані сп’яніння;відмітка про порушення режиму; прізвище;ім’я;по батькові;код РНОКПП\n"
tmplt, _ = encoder.String(tmplt)
f, _ := os.Create(fname)
f.Write([]byte(tmplt))
f.Sync()
defer f.Close()




    for _,r := range  exptbl.Rows.Row {
	tmplt = r.WICNUM+";"+r.WICCASENUM+";"+r.WICDTBEGIN[:10]+";"+r.WICDTEND[:10]+";"
	switch r. WICSTATUS  {
	case "A": 
		tmplt = tmplt + "закритий"+";"

	case "P": 
		tmplt = tmplt + "готовий до сплати"+";"
        }
	tmplt = tmplt +r.WICCD+";"+r.WICCDName+";"+r.SIGNANLKNARKOTIKINTOXICATION+";"+r.VIOLATIONEXTENSION+";"+r.NPSURNAME+";"+r.NPNAME+";"+r.NPPATRONYMIC+";"+r.NPNUMIDENT+"\n"

	tmplt, _ = encoder.String(tmplt)
	f.Write([]byte(tmplt))
	f.Sync()

		    
	}

}