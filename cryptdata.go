package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"
)

type Ziutcrypoto struct {
	tok int
}

func (z *Ziutcrypoto) GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (z *Ziutcrypoto) returnStr64() string {
	b, err := z.GenerateRandomBytes(z.tok)
	if err != nil {
		// log.Panic("Nope-> ", err)
		os.Exit(02)
	}
	c := base64.URLEncoding.EncodeToString(b)
	// log.Print("GENERATED ? NEW F NAME ?-> ", c)
	return c
}

func newZiutCrypto(s int) *Ziutcrypoto {
	return &Ziutcrypoto{tok: s}
}

func ziuted(ch chan string, f int) {
	c := newZiutCrypto(f)
	// log.Print("CRYPTO STARTET ->", f)

	go func() {
		// log.Print("WE ARE IN SOME GENERATOR HELP ME ! ")
		ch <- c.returnStr64()
	}()

}
