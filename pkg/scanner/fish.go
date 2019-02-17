package scanner

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type FishRow struct {
	Name            string
	Ticker          string
	Sector          string
	Industry        string
	NumYrs          int
	CCCSeq          int
	Price           float32
	DividendYield   float32
	CurrentDividend float32
	PayoutPerYear   int
	Annualized      float32
	// 	Sch
	// 	PreviousPayout
	// 	Ex-Div
	// 	Pay
	MRInc  float32
	DGR1Y  float32
	DGR3Y  float32
	DGR5Y  float32
	DGR10Y float32
	// 	A/D*
	DEG        float32
	EPSPay     float32
	TTMPE      float32
	DebtEquity float32
	Tweed      float32
	Chowder    float32
	Graham     float32
}

type CandidateStocks []*FishRow

func Fish() (ret CandidateStocks) {
	ret = make(CandidateStocks, 0, 800)
	var xlsx *excelize.File
	var err error
	if _, err = os.Stat("/tmp/fish.xlsx"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		url := "https://bitly.com/USDividendChampions"
		xlsx, err = readXLXSFromUrl(url)
	} else {
		xlsx, err = excelize.OpenFile("/tmp/fish.xlsx")
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("All CCC", "A3")
	fmt.Printf("reading informatio correct %s\n", cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("All CCC")
	for _, row := range rows[6:] {

		ny, _ := strconv.Atoi(row[4])
		cccseq, _ := strconv.Atoi(row[5])
		price, _ := strconv.ParseFloat(row[8], 32)
		divYeild, _ := strconv.ParseFloat(row[9], 32)
		curdiv, _ := strconv.ParseFloat(row[10], 32)
		ppy, _ := strconv.Atoi(row[11])
		ann, _ := strconv.ParseFloat(row[12], 32)
		mri, _ := strconv.ParseFloat(row[17], 32)
		dgi1, _ := strconv.ParseFloat(row[18], 32)
		dgi3, _ := strconv.ParseFloat(row[19], 32)
		dgi5, _ := strconv.ParseFloat(row[20], 32)
		dgi10, _ := strconv.ParseFloat(row[21], 32)
		deg, _ := strconv.ParseFloat(row[23], 32)
		epspay, _ := strconv.ParseFloat(row[25], 32)
		ttmpe, _ := strconv.ParseFloat(row[26], 32)
		dte, _ := strconv.ParseFloat(row[39], 32)
		tweed, _ := strconv.ParseFloat(row[40], 32)
		chow, _ := strconv.ParseFloat(row[41], 32)
		graham, _ := strconv.ParseFloat(row[42], 32)
		r := &FishRow{
			Name:            row[0],
			Ticker:          row[1],
			Sector:          row[2],
			Industry:        row[3],
			NumYrs:          ny,
			CCCSeq:          cccseq,
			Price:           float32(price),
			DividendYield:   float32(divYeild),
			CurrentDividend: float32(curdiv),
			PayoutPerYear:   ppy,
			Annualized:      float32(ann),
			MRInc:           float32(mri),
			DGR1Y:           float32(dgi1),
			DGR3Y:           float32(dgi3),
			DGR5Y:           float32(dgi5),
			DGR10Y:          float32(dgi10),
			DEG:             float32(deg),
			EPSPay:          float32(epspay),
			TTMPE:           float32(ttmpe),
			DebtEquity:      float32(dte),
			Tweed:           float32(tweed),
			Chowder:         float32(chow),
			Graham:          float32(graham),
		}
		ret = append(ret, r)
	}
	return ret
}

func readXLXSFromUrl(url string) (*excelize.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	xlsx, err := excelize.OpenReader(resp.Body)

	err = xlsx.SaveAs("/tmp/fish.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	return xlsx, err
}
