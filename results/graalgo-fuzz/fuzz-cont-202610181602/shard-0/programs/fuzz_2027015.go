package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint = 603
    var v2 uint32 = 65535
    v3 := make(map[string]uint32)
    var v4 int = -297
    var v5 int = v4
    v6 := make(map[int]bool)
    if (!true) {
        if (true || true) {
            delete(v6, 127)
            _ = len(v6)
        }
        v7 := GenS1{F0: 1, F1: 0, F2: (2147483648 <= v2)}
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                v4 = fi8
                _ = fi9
            }
            switch "foo bar" {
            case "foo bar":
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := []int{(^(v4)), fi10, int(v2)}
                    var v12 int16 = 913
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                v13 := GenS2{F0: 139, F1: 42}
                _ = v13
            }
            _ = fi8
        }
        _ = v7
    } else {
        fmt.Println(v5, v4)
    }
    var v14 int8 = 26
    v2 = v2
    for fi15 := 0; fi15 < 4; fi15++ {
        v1 += v1
        _ = fi15
    }
    v14 = (v14 & -128)
    v5 = (v5 ^ (v5 % ((524) | 1)))
    v16 := []int{32768}
    v17 := GenS3{F0: 107, F1: 1}
    fmt.Println(v5)
    if (!false) {
        v5 -= v4
        _ = len(v6)
        var v18 string = (("hello" + ("a" + (("\n\t" + "") + ""))) + (((("hello" + ("\n\t" + "hello")) + "foo bar") + "hello") + "hello"))
        _ = v18
    } else {
        fmt.Println(v4)
        v19 := GenS2{F0: 2, F1: -9223372036854775808}
        v5 += (v4 / ((v4) | 1))
        _ = v19
    }
    if len(v16) > 0 {
        v16[0] = v4
    }
    switch "\n\t" {
    case "\n\t":
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(fi20, v1)
            v14 = -1
            _ = fi20
        }
    default:
        v1 = v1
    }
    v21 := GenS2{F0: 76, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
