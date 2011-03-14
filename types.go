package main

import (
	"io"
	"os"
	"net"
	"log"
	"encoding/binary"
)

// This file handles various types utilized by the protocol.

// byte: 1
// short: 2
// int: 4
// long: 8
// float: 4
// double: 8
// modified utf-8 string: >= 2
//  First two bytes determine the length
// bool: 1

// Gets a string from the connection
func ReadString(c net.Conn) (string, os.Error) {
	buf1 := make([]byte, 2)
	buf1[0], _ = ReadByte(c)
	buf1[1], _ = ReadByte(c)
	length := int(buf1[0]*255)
	length += int(buf1[1])
	log8.Println("String length should be", length)
	buf := make([]byte, length)
	
	for i := 0; i < len(buf); i++ {
		char, err := ReadByte(c)
		if err != nil {
			return string(buf), err
		}
		buf[i] = char
	}
	return string(buf), nil
}

func makeString(orig string) (buf []byte) {
	buf = make([]byte, len(orig)+2)
	buf[1] = byte(len(orig))
	buf[0] = byte((len(orig))/255)
	for i, char := range orig {
		// WARNING: Only works with ASCII. Should fix eventually.
		buf[i+2] = byte(char)
	}
	return buf
}

func ReadByte(c io.Reader) (byte, os.Error) {
	buf := make([]byte, 1)
	_, err := c.Read(buf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//log.Println("Received", n, "bytes from server. (byte:", buf, "char:", string(buf))
	return buf[0], err
}

var bOrder = binary.BigEndian

func ReadShort(c net.Conn) (int16, os.Error) {
	var b int16
	binary.Read(c, bOrder, &b)
	return b, nil
}

func ReadInt(c net.Conn) (int32, os.Error) {
	var b int32
	binary.Read(c, bOrder, &b)
	return b, nil
}

func ReadLong(c net.Conn) (int64, os.Error) {
	var b int64
	binary.Read(c, bOrder, &b)
	return b, nil
}
// Floating point ones don't seem to work :/
func ReadFloat(c net.Conn) (float32, os.Error) {
	var b float32
	binary.Read(c, bOrder, &b)
	return b, nil
}

func ReadDouble(c net.Conn) (float64, os.Error) {
	var b float64
	binary.Read(c, bOrder, &b)
	return b, nil
}

func ReadDataStream(c net.Conn) ([]byte, os.Error) {
	buf := make([]byte, 1)
	for {
		b, _ := ReadByte(c)
		buf = append(buf, b)
		if b == 0x7F {
			return buf[1:], nil
		}
	}
	panic("Not reached!")
}

func WriteByte(c net.Conn, b byte) (os.Error) {
	n, err := c.Write([]byte{b})
	if err != nil {
		log0.Println(err)
	}
	log9.Println("Wrote", n, "bytes to server. (byte:", b, "char:", string(b))
	return err
}

func WriteShort(c net.Conn, b int16) {
	binary.Write(c, bOrder, b)
}

func WriteInt(c net.Conn, b int32) {
	binary.Write(c, bOrder, b)
}

func WriteLong(c net.Conn, b int64) {
	binary.Write(c, bOrder, b)
}

func WriteFloat(c net.Conn, b float32) {
	binary.Write(c, bOrder, b)
}

func WriteDouble(c net.Conn, b float64) {
	binary.Write(c, bOrder, b)
}

func WriteString(c net.Conn, s string) {
	buf := makeString(s)
	n, err := c.Write(buf)
	if err != nil {
		log.Println(err)
	}
	log8.Println("Wrote", n, "bytes to server. (bytes:", buf, "string:", s)
}