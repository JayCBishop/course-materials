package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		var hash string
		if len(sourceHash) == 32 {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				return password
			}
		} else if len(sourceHash) == 64 {
			hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				return password
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return ""
}

func GenHashMaps(filename string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		GenSHAHashMaps(filename)
		wg.Done()
	}()
	go func() {
		GenMD5HashMaps(filename)
		wg.Done()
	}()
	wg.Wait()
}

func GenSHAHashMaps(filename string) {
	shalookup = make(map[string]string)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()
		sha256 := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		shalookup[password] = sha256
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GenMD5HashMaps(filename string) {
	md5lookup = make(map[string]string)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()
		md5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		md5lookup[password] = md5
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil
	} else {
		return "", errors.New("password does not exist")
	}
}

func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil
	} else {
		return "", errors.New("not implemented")
	}
}
