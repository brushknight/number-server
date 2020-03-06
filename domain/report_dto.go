package domain

type ReportDTO struct {
	UniqNumbersTotal               uint64
	UniqNumbersFromLastReport      uint64
	DuplicateNumbersFromLastReport uint64
	AllNumbersTotal                uint64
	Rps                            uint64
}

func CreateReportDTO(uniqNumbersTotal uint64, uniqNumbersFromLastReport uint64, duplicateNumbersFromLastReport uint64, allNumbersTotal uint64, rps uint64) ReportDTO {
	return ReportDTO{
		UniqNumbersTotal:               uniqNumbersTotal,
		UniqNumbersFromLastReport:      uniqNumbersFromLastReport,
		DuplicateNumbersFromLastReport: duplicateNumbersFromLastReport,
		AllNumbersTotal:                allNumbersTotal,
		Rps:                            rps}
}
