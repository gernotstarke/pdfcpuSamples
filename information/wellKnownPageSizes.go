package information

type PageSize = struct {
	Width, Height float64
}

type SizeName string

var PageSizes = map[SizeName]*PageSize{

	"A4": {595,842},
	"A5": { 420, 595},
}


