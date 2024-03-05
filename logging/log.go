package logging

import "log"

func Info(s string, params ...any) {
	writeLog(s, "(~)", params...)
}

func Error(s string, params ...any) {
	writeLog(s, "(-)", params...)
}

func Ok(s string, params ...any) {
	writeLog(s, "(+)", params...)
}

func writeLog(s string, status string, params ...any) {
	s = status + " " + s
	if len(params) == 0 {
		log.Print(s)
		return
	}
	log.Printf(s, params...)
}
