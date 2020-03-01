package domain

type ReportDTO struct {
	uniqNumbersTotal               uint64
	uniqNumbersFromLastReport      uint64
	duplicateNumbersFromLastReport uint64
	allNumbersTotal                uint64
}

func CreateReportDTO(uniqNumbersTotal uint64, uniqNumbersFromLastReport uint64, duplicateNumbersFromLastReport uint64, allNumbersTotal uint64) ReportDTO {
	return ReportDTO{
		uniqNumbersTotal:               uniqNumbersTotal,
		uniqNumbersFromLastReport:      uniqNumbersFromLastReport,
		duplicateNumbersFromLastReport: duplicateNumbersFromLastReport,
		allNumbersTotal:                allNumbersTotal}
}
