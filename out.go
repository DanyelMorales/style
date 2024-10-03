package style

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func ReportVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, Documentation, format, a...)
}

func ReportV(w io.Writer, format string) {
	LogViewF(w, Documentation, format)
}

func NoticeVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, Notice, format, a...)
}

func NoticeV(w io.Writer, format string) {
	LogViewF(w, Notice, format)
}

func TitleV(w io.Writer, title string) {
	LogViewF(w, AddonEnable, title)
}

func TitleVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, AddonEnable, format, a...)
}

func ActionV(w io.Writer, format string) {
	LogViewF(w, Running, format)
}
func SuccessfulVf(w io.Writer, format string) {
	LogViewF(w, ThumbsUp, format)
}
func FailedActionV(w io.Writer, format string) {
	LogViewF(w, Failure, format)
}
func ExitActionV(w io.Writer, format string) {
	LogViewF(w, Deleted, format)
}
func ExitActionVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, Deleted, format, a...)
}
func ActionVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, Running, format, a...)
}
func SuccessfulActionVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, ThumbsUp, format, a...)
}
func FailedActionVF(w io.Writer, format string, a ...interface{}) {
	LogViewF(w, Failure, format, a...)
}
func ActLogCVF(w io.Writer, print bool, format string, a ...interface{}) {
	if !print {
		return
	}
	LogViewF(w, Running, format, a...)
}

func OkLogVCF(w io.Writer, print bool, format string, a ...interface{}) {
	if !print {
		return
	}
	LogViewF(w, ThumbsUp, format, a...)
}

func ErrorLogVCF(w io.Writer, print bool, format string, a ...interface{}) {
	if !print {
		return
	}
	LogViewF(w, Failure, format, a...)
}

func OnReportF(format string, a ...interface{}) {
	LogF(Documentation, format, a...)
}

func OnReport(format string) {
	Log(Documentation, format)
}

func OnNoticeF(format string, a ...interface{}) {
	LogF(Notice, format, a...)
}

func OnNotice(format string) {
	Log(Notice, format)
}

func OnTitle(title string) {
	Log(AddonEnable, title)
}

func OnTitleF(format string, a ...interface{}) {
	LogF(AddonEnable, format, a...)
}

func OnAction(format string) {
	Log(Running, format)
}
func SuccessfulAction(format string) {
	Log(ThumbsUp, format)
}
func FailedAction(format string) {
	Log(Failure, format)
}
func ExitAction(format string) {
	Log(Deleted, format)
}
func ExitActionF(format string, a ...interface{}) {
	LogF(Deleted, format, a...)
}
func OnActionF(format string, a ...interface{}) {
	LogF(Running, format, a...)
}
func SuccessfulActionF(format string, a ...interface{}) {
	LogF(ThumbsUp, format, a...)
}
func FailedActionF(format string, a ...interface{}) {
	LogF(Failure, format, a...)
}
func ActLogCF(print bool, format string, a ...interface{}) {
	if !print {
		return
	}
	LogF(Running, format, a...)
}
func OkLogCF(print bool, format string, a ...interface{}) {
	if !print {
		return
	}
	LogF(ThumbsUp, format, a...)
}
func ErrorLogCF(print bool, format string, a ...interface{}) {
	if !print {
		return
	}
	LogF(Failure, format, a...)
}

func PrintF(st Enum, format string, a ...interface{}) {
	msg := StringF(st, format, a...)
	fmt.Print(msg)
}

func Print(st Enum, format string) {
	msg := String(st, format)
	fmt.Print(msg)
}
func LogViewF(w io.Writer, st Enum, format string, a ...interface{}) {
	msg := StringF(st, format, a...)
	fmt.Fprintln(w, msg)
}
func LogF(st Enum, format string, a ...interface{}) {
	msg := StringF(st, format, a...)
	log.Print(msg)
}

func Log(st Enum, format string) {
	msg := String(st, format)
	log.Print(msg)
}

// StringF writes a basic formatted string to stdout
func String(st Enum, format string) string {
	prefix, ok := Config[st]
	if ok {
		var color = string(prefix.Color)
		str := color + prefix.Prefix + format + "\n"
		return strings.ToLower(str)
	}
	return ""
}

func StringF(st Enum, format string, a ...interface{}) string {
	str := String(st, format)
	return fmt.Sprintf(str, a...)
}
