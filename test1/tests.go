package main

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// Test1 принимает на вход список строк и выводит их в одной строке через пробел, возможно, в случайном порядке.
// Смысл этого кода не в конкатенации слов через пробел, а в понимании работы многопоточности.
//
// Вам необходимо исправить код так, чтобы он работал, не удаляя ничего из него, а только внося исправления.
//
// Доп. задание: напишите тест на эту функцию.
//
// Пример вызова функции:
//     data1 := []string{"hello", "world", "test", "data", "code"}
//     r1 := Test1(data1)
//     fmt.Println(r1)
func Test1(list []string) string {
	var result string
	var mu sync.Mutex

	for _, l := range list {
		go func() {
			mu.Lock()
			result += l + " "
			mu.Unlock()
		}()
	}

	return result
}

// Test2 создаёт файл во временной директории и записывает в него строку из data.
//
// Вам необходимо исправить его так, чтобы он работал верно под любой ОС.
//
// Доп. задание: Можно ли упростить данный код?
//
// Пример вызова функции:
//     err := Test2("test2.txt", "hello test2")
func Test2(filename, data string) error {
	tmp, exist := os.LookupEnv("TMP")
	if !exist {
		tmp, exist = os.LookupEnv("TEMP")
	}
	if exist {
		f, err := os.Create(tmp + "/" + filename)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write([]byte(data))
		if err != nil {
			return err
		}
	} else {
		return errors.New("TMP folder is not defined")
	}

	return nil
}

// Test3 отправляет запрос через прокси и возвращает внешний IP.
//
// Вам необходимо реализовать поддержку прокси.
//
// Пример вызова функции:
//     ip, err := Test3("http://user:pass@127.0.0.1:3128/")
func Test3(proxyURL string) (string, error) {
	os.Setenv("HTTP_PROXY", proxyURL)
	os.Setenv("HTTPS_PROXY", proxyURL)
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var v struct {
		Origin string
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return "", err
	}

	return v.Origin, nil
}

// Test4 оборачивает соединение, для подсчёта принятых/переданных данных.
//
// Вам необходимо реализовать данный функционал с поддержкой многопоточного доступа к данным.
// По желанию, можете написать тест на данную функцию.
//
// Пример вызова функции:
//     conn, err := net.Dial(...)
//     ...
//     countConn := Test4(conn)
//     _, err = conn.Write([]byte("send test data"))
//     ...
//     fmt.Printf("Write: %d\n", countConn.WriteByteCount())
//
//     ...
//     _, err := conn.Read(buf)
//     ...
//     fmt.Printf("Read: %d\n", countConn.ReadByteCount())
//
//     err = conn.Close()
//     ...
//     fmt.Printf("Read: %d\n", countConn.ReadByteCount())
//     fmt.Printf("Write: %d\n", countConn.WriteByteCount())

type Test4Conn interface {
	net.Conn
	ReadByteCount() uint64
	WriteByteCount() uint64
}

type connWithCounter struct {
	conn       net.Conn
	bytesRead  int
	bytesWrite int
	mu         sync.Mutex
}

func (c *connWithCounter) Read(b []byte) (n int, err error) {
	n, err = c.conn.Read(b)
	c.mu.Lock()
	c.bytesRead = c.bytesRead + n
	c.mu.Unlock()
	return
}

func (c *connWithCounter) Write(b []byte) (n int, err error) {
	n, err = c.conn.Write(b)
	c.mu.Lock()
	c.bytesWrite = c.bytesWrite + n
	c.mu.Unlock()
	return
}

func (c *connWithCounter) ReadByteCount() uint64 {
	return uint64(c.bytesRead)
}

func (c *connWithCounter) WriteByteCount() uint64 {
	return uint64(c.bytesWrite)
}

// Чтобы структура connWithCounter реализовывала интерфейс net.Conn, нужно сделать proxy-реализацию всех недостающих методов
// Пример для метода Close:
func (c *connWithCounter) Close() error {
	return c.conn.Close()
}

func (c *connWithCounter) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *connWithCounter) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c *connWithCounter) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *connWithCounter) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *connWithCounter) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}
