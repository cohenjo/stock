package scanner

type FilterFunc = func(CandidateStocks, float32) CandidateStocks
type FishFilterFunc = func(*FishRow, float32) bool

type Filter struct {
	FF FilterFunc
	FV float32
}

type FilterRegistery struct {
	filters map[string]Filter
}

var FR = NewRegistry()

func NewRegistry() *FilterRegistery {
	return &FilterRegistery{
		filters: make(map[string]Filter),
	}
}

func (fr *FilterRegistery) RegisterFilter(name string, f FilterFunc, v float32) {
	fr.filters[name] = Filter{FF: f, FV: v}
}

func (fr *FilterRegistery) RunFilters(cs CandidateStocks) CandidateStocks {
	for _, filter := range fr.filters {
		cs = filter.FF(cs, filter.FV)
	}
	return cs
}

// RegisterFilters - based off "The K.I.S.S. System"
// by: https://seekingalpha.com/author/the-part-time-investor
func RegisterFilters() {
	FR.RegisterFilter("chowder", ChowderFilter, 13)
	FR.RegisterFilter("debtEquite", DebtToEquityFilter, 1)
	FR.RegisterFilter("divYield", DividendYieldFilter, 2.5)
	FR.RegisterFilter("numYears", NumYearFilter, 10)
	// FR.RegisterFilter("dgr1", DGR1Filter, 5)
	// FR.RegisterFilter("dgr3", DGR3Filter, 7)
	FR.RegisterFilter("dgr5", DGR5Filter, 5)
	// FR.RegisterFilter("dgr10", DGR10Filter, 7)
	// FR.RegisterFilter("payoutRatio", EPSPayRatioFilter, 60)
	// FR.RegisterFilter("priceEarning", PriceEarningFilter, 15)
	// FR.RegisterFilter("graham", GrahamFilter, 0)
}

func RegisterHighFilters() {
	FR.RegisterFilter("chowder", ChowderFilter, 8)
	// FR.RegisterFilter("debtEquite", DebtToEquityFilter, 1)
	FR.RegisterFilter("divYield", DividendYieldFilter, 4)
	// FR.RegisterFilter("numYears", NumYearFilter, 10)
	FR.RegisterFilter("dgr1", DGR1Filter, 4)
	FR.RegisterFilter("dgr3", DGR3Filter, 4)
	FR.RegisterFilter("dgr5", DGR5Filter, 4)
	FR.RegisterFilter("dgr10", DGR10Filter, 4)
	FR.RegisterFilter("payoutRatio", EPSPayRatioFilter, 60)
	FR.RegisterFilter("priceEarning", PriceEarningFilter, 15)
	// FR.RegisterFilter("graham", GrahamFilter, 0)
}

// FILTERS start here

func ChowderFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.Chowder > val) }
	return GenericFilter(cs, f, val)
}

func DebtToEquityFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.DebtEquity < val) }
	return GenericFilter(cs, f, val)
}

func DividendYieldFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.DividendYield > val) }
	return GenericFilter(cs, f, val)
}

func NumYearFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.NumYrs > int(val)) }
	return GenericFilter(cs, f, val)
}

func DGR3Filter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.DGR3Y > val) }
	return GenericFilter(cs, f, val)
}
func DGR1Filter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.DGR1Y > val) }
	return GenericFilter(cs, f, val)
}

func DGR10Filter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.DGR10Y > val) }
	return GenericFilter(cs, f, val)
}

func DGR5Filter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.DGR5Y > val) }
	return GenericFilter(cs, f, val)
}

func EPSPayRatioFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.EPSPay < val) }
	return GenericFilter(cs, f, val)
}

func PriceEarningFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.TTMPE < val) }
	return GenericFilter(cs, f, val)
}

func GrahamFilter(cs CandidateStocks, val float32) CandidateStocks {
	f := func(s *FishRow, v float32) bool { return (s.Graham < val) }
	return GenericFilter(cs, f, val)
}

func GenericFilter(cs CandidateStocks, f FishFilterFunc, val float32) CandidateStocks {
	ret := cs[:0]
	for _, stock := range cs {
		if f(stock, val) {
			ret = append(ret, stock)
		}
	}
	return ret
}
