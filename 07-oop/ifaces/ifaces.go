package main

import (
	"errors"
	"log"
)

// Logger - журнал.
type Logger interface {
	Log(string) error
}

// DBLogger - журнал в БД.
type DBLogger struct {
	connString string // строка для подключения к БД
}

// Log записывает сообщение в журнал в БД.
func (dbl *DBLogger) Log(msg string) error {
	if dbl.connString == "" {
		return errors.New("БД недоступна")
	}
	// выполняем запись сообщения в БД
	return nil
}

// MemLogger - заглушка журнала в памяти для тестов.
type MemLogger int

// Log ничего не делает. Обратите внимание на отсутствие имён у получателя и аргумента.
func (*MemLogger) Log(string) error {
	return nil
}

func main() {
	l := new(DBLogger)
	err := logMsg(l, "сообщение")
	if err != nil {
		log.Println(err)
		return
	}
}

// logMsg записывает сообщение в журнал
func logMsg(l Logger, msg string) error {
	err := l.Log(msg)
	return err
}
