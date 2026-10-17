package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int32) int32 {
    var v6 string = ""
    p0 = (p0 % (((p0 | 18446744073709551615)) | 1))
    fmt.Println(p1, p1, v6)
    v6 = ("foo bar" + v6)
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1() int64 {
    v7 := []int{2147483647}
    switch -1 {
    case 0:
        var v8 int = -906
        _ = v8
    }
    switch "hello" {
    case "":
        var v9 uint32 = 4294967295
        _ = v9
    default:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 363
    var v2 float32 = 1e10
    var v3 uint8 = 128
    v4 := []uint32{}
    var v5 bool = true
    var v10 int64 = -173
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v10 += (v10 & int64(v3))
    if ((v5 || (((!v5) || (false && (!(((1000 < 32767) && true) || (2 > 2))))) && (v5 && (("hello" + "") < ("\n\t" + "a"))))) || v5) {
        var v11 string = "hello"
        v12 := GenS1{F0: 2147483647, F1: 42}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v11
        _ = v12
    }
    fmt.Println(v5, v1)
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = -1
        fmt.Println(v5, v14, v14)
        for fi15 := 0; fi15 < 16; fi15++ {
            if v5 {
                switch ("hello" + "a") {
                case "foo bar":
                    v16 := GenS2{F0: (v3 - v3), F1: ((("a" + ("foo bar" + ("\n\t" + "hello"))) + ("" + (("hello" + ("\n\t" + "foo bar")) + "foo bar"))) > "\n\t")}
                    _ = v16
                case "\n\t":
                    v4 = append(v4, v1)
                }
                v3 -= v3
                v14 += v14
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v5 = ((1.5 <= 3.14159) && false)
    v17 := GenS3{F0: -32768}
    v1 += 42
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := GenS4{F0: 127}
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 uint16 = 65535
            if (-518 != 885) {
                switch "foo bar" {
                case "":
                    var v22 int = fi20
                    _ = v22
                case "\n\t":
                    var v23 int64 = (v10 / ((v10) | 1))
                    _ = v23
                }
                _ = len(v4)
            } else {
                v24 := GenS3{F0: 0}
                fmt.Println(v3)
                v4 = append(v4, 1)
                if len(v4) > 0 {
                    v4[0] = (v1 & 65535)
                }
                var v25 uint8 = v3
                _ = v24
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        v10 += int64(v3)
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
