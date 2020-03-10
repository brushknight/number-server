package reporter

import (
	"fmt"
	"io"
)

func ProcessReportsChannel(reportsQueue chan ReportDTO, writer io.Writer, env string) {
	for report := range reportsQueue {
		writer.Write([]byte(buildReportString(report, env)))
	}
}

func buildReportString(report ReportDTO, env string) string {

	if env == "prd" {
		return fmt.Sprintf("Received %d unique numbers, %d duplicates. Unique total: %d\n", report.UniqNumbersFromLastReport, report.DuplicateNumbersFromLastReport, report.UniqNumbersTotal)
	} else {
		return fmt.Sprintf("Received %d unique numbers, %d duplicates. Total %d Uniq: %d. Rps %d\n", report.UniqNumbersFromLastReport, report.DuplicateNumbersFromLastReport, report.UniqNumbersTotal, report.AllNumbersTotal, report.Rps)
	}
}
