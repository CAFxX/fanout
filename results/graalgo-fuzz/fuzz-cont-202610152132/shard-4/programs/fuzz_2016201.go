package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 uint32, p1 uint64) int {
    fmt.Println(p1)
    var v6 int64 = 9223372036854775807
    p0 += (p0 & 183)
    _ = p0
    _ = p1
    _ = v6
    return -1
}
func genFn1() uint8 {
    v7 := GenS1{F0: -610, F1: 830}
    switch "foo bar" {
    case "a":
        if true {
            var v8 uint64 = 0
            fmt.Println(v8)
            for fi9 := 0; fi9 < 8; fi9++ {
                v8 = (v8 % (((v8 % (((v8 - (v8 ^ 18446744073709551615))) | 1))) | 1))
                var v10 uint32 = 2147483647
                v11 := []int{(fi9 * (fi9 | fi9)), (fi9 + (fi9 - (fi9 ^ 9223372036854775807))), (fi9 & genFn0(411, (v8 >> ((65535) & 63)))), (fi9 >> ((431) & 63))}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
    case "\n\t":
        v12 := []int{-2147483648, -2147483648}
        _ = v12
    case "":
        var v13 uint8 = 127
        _ = v13
    default:
        var v14 uint32 = 0
        _ = v14
    }
    _ = v7
    return 128
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]string)
    var v3 int16 = 0
    var v4 float64 = 1.5
    v5 := make(map[int]uint32)
    v4 = v4
    v15 := GenS2{F0: 128, F1: v1}
    if false {
        var v16 uint8 = 92
        v3 = -304
        v1 = v1
        _ = v16
    } else {
        fmt.Println(v4, v4)
        switch 42 {
        case -1:
            var v17 int64 = 9223372036854775807
            _ = v17
        case -9223372036854775808:
            if ((v3 / (((v3 ^ (v3 + 1))) | 1)) <= v3) {
                var v18 uint16 = 603
                _ = v18
            } else {
                v19 := GenS2{F0: 42, F1: v1}
                var v20 int32 = -537
                v21 := GenS3{F0: genFn0(1, 0), F1: v3, F2: (("a" + ("a" + ("" + "a"))) + "foo bar")}
                _ = v19
                _ = v20
                _ = v21
            }
        default:
            if false {
                v1 = 1
                switch -1 {
                case -9223372036854775808:
                    var v22 int8 = 100
                    _ = v22
                }
            }
        }
        _ = len(v2)
    }
    var v23 string = ""
    fmt.Println(v4, v1, v1)
    switch "foo bar" {
    case "foo bar":
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                v3 = int16(fi24)
                _ = fi25
            }
            _ = fi24
        }
    default:
        v4 -= 1.5
    }
    delete(v2, "foo bar")
    var v26 bool = false
    switch v23 {
    case "\n\t":
        v27 := GenS1{F0: 42, F1: -2147483648}
        _ = v27
    }
    var v28 string = "\n\t"
    v29 := GenS3{F0: -1, F1: -1, F2: v28}
    var v30 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
