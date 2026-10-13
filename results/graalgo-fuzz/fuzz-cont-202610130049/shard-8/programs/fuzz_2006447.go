package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int8; F1 string; F2 string }
type GenS3 struct { F0 uint; F1 int16; F2 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 1
    v3 := make(map[int]bool)
    var v4 uint = 898
    switch "hello" {
    case "\n\t":
        v4 += v4
    case "a":
        v5 := []int{42, 9223372036854775807}
        _ = v5
    case "hello":
        var v6 int64 = (v1 * v1)
        _ = v6
    default:
        v3[42] = true
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            if (-69 <= 1000) {
                fmt.Println(fi7)
                var v9 int = fi7
                v10 := GenS1{F0: 1, F1: 0}
                _ = v10
                _ = v9
            }
            v4 -= (v4 - (v4 - (v4 ^ (v4 % ((v4) | 1)))))
            v2 -= (v2 | (v2 << (((^(v4))) & 63)))
            _ = fi8
        }
        _ = fi7
    }
    var v11 int64 = 1
    switch 0 {
    case -107:
        v12 := GenS2{F0: -1, F1: "\n\t", F2: "hello"}
        _ = v12
    default:
        v13 := GenS2{F0: -101, F1: "foo bar", F2: ("foo bar" + ("foo bar" + "\n\t"))}
        _ = v13
    }
    var v14 bool = false
    for fi15 := 0; fi15 < 14; fi15++ {
        delete(v3, 703)
        _ = fi15
    }
    v14 = (!v14)
    switch "\n\t" {
    case "\n\t":
        var v16 float64 = -2.25
        _ = v16
    case "hello":
        v4 = 2
    }
    switch "foo bar" {
    case "a":
        switch "hello" {
        case "foo bar":
            v14 = false
        case "hello":
            v11 += (v1 / ((9223372036854775807) | 1))
        case "a":
            fmt.Println(v2, v2)
        default:
            if v14 {
                fmt.Println(v14, v4)
            } else {
                v14 = ("\n\t" == "hello")
                v11 += (v1 >> ((v4) & 63))
            }
        }
    default:
        var v17 uint8 = 177
        _ = v17
    }
    var v18 uint64 = v2
    v3[32768] = (!true)
    v18 = (v2 | 2)
    if (!v14) {
        if v14 {
            v18 = v18
            var v19 uint64 = (v18 | (^(v2)))
            var v20 uint16 = 255
            _ = v19
            _ = v20
        } else {
            var v21 float32 = 1.5
            v22 := GenS3{F0: 459, F1: -670, F2: -1}
            _ = v21
            _ = v22
        }
        fmt.Println(v1, v2)
    } else {
        if ((v18 * v2) != v18) {
            var v23 uint8 = 82
            _ = v23
        }
        if v14 {
            v24 := GenS3{F0: v4, F1: -622, F2: -2147483648}
            _ = v24
        }
    }
    fmt.Println(v14)
    fmt.Println(v14, v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
