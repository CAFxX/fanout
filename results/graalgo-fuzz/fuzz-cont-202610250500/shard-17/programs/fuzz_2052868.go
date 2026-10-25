package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint64, p1 uint64) bool {
    var v7 int64 = -266
    v8 := []int{0, -756}
    var v9 int8 = 5
    v10 := GenS1{F0: false, F1: 128, F2: (v9 > v9)}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint64 = 42
    var v2 uint64 = 42
    var v3 uint16 = 65535
    v4 := []uint32{2147483647}
    var v5 bool = (!true)
    var v6 uint64 = (v1 | (^(v2)))
    var v11 uint = 110
    fmt.Println(v1)
    v12 := GenS2{F0: 765, F1: 4294967295}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "":
        var v13 bool = (42 <= 72)
        _ = v13
    case "foo bar":
        var v14 uint8 = 127
        _ = v14
    case "a":
        if (-963 == int32(v2)) {
            v15 := GenS1{F0: (!v5), F1: 2, F2: v5}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 15; fi16++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = fi16
            }
            var v17 float64 = 1e10
            _ = v17
        }
    }
    v11 += 65535
    v3 = (v3 * (v3 % (((^(v3))) | 1)))
    v18 := GenS1{F0: ((!false) || (-71 == -128)), F1: 127, F2: (2 < uint32(v2))}
    switch "\n\t" {
    case "hello":
        v1 -= (v2 / (((v6 - 42)) | 1))
    case "\n\t":
        var v19 int32 = 929
        _ = v19
    default:
        switch "hello" {
        case "foo bar":
            v20 := GenS3{F0: "hello"}
            _ = v20
        case "":
            v11 = v11
        case "\n\t":
            v6 += (v1 & 18446744073709551615)
        }
    }
    var v21 string = ("foo bar" + "\n\t")
    var v22 int8 = 1
    v5 = false
    var v23 int16 = -524
    if false {
        fmt.Println(v1, v2)
        v24 := GenS3{F0: v21}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
