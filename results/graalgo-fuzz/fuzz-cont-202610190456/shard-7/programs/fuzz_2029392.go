package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }
func genFn0(p0 int8) int32 {
    v5 := GenS1{F0: (-(p0)), F1: 18446744073709551615}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return -289
}

func main() {
    var v1 bool = ("hello" < "a")
    var v2 uint32 = 0
    v3 := []uint32{42, 0, (v2 / (((^(v2))) | 1)), (v2 | (v2 % ((65535) | 1)))}
    var v4 uint16 = 65535
    if (!true) {
        var v6 uint32 = (v2 % ((768) | 1))
        _ = v6
    } else {
        var v7 string = ("\n\t" + "\n\t")
        _ = v7
    }
    _ = len(v3)
    fmt.Println(v2, v2, v2)
    v2 += (v2 & 4294967295)
    fmt.Println(v4, v2, v1)
    if (!(0 > 785)) {
        v2 += (^(v2))
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (1 == 170) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    var v8 float32 = float32(3.14159)
    v9 := GenS1{F0: 37, F1: 427}
    _ = len(v3)
    v10 := GenS2{F0: 62}
    v4 = 2
    switch "\n\t" {
    case "\n\t":
        var v11 uint8 = 1
        _ = v11
    case "foo bar":
        if (!v1) {
            v12 := GenS3{F0: 0, F1: "hello", F2: 2}
            _ = v12
        }
    case "a":
        var v13 int16 = -1
        _ = v13
    default:
        v1 = v1
    }
    var v14 int8 = -72
    var v15 uint8 = 123
    var v16 int16 = 716
    switch "a" {
    case "a":
        var v17 uint8 = uint8(v4)
        _ = v17
    case "\n\t":
        v15 -= 218
    case "":
        fmt.Println(v2)
    }
    var v18 float64 = 0.1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
