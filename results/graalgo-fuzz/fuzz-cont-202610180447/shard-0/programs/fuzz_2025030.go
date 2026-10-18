package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 uint, p2 uint64) string {
    switch (p0 ^ (-(p0))) {
    case 127:
        var v4 int32 = -2147483648
        _ = v4
    case 423:
        switch "" {
        case "":
            var v5 uint16 = 0
            _ = v5
        case "foo bar":
            fmt.Println(p1)
        case "a":
            var v6 float64 = 1e10
            _ = v6
        }
    }
    v7 := GenS1{F0: 1, F1: 65535}
    var v8 int64 = 42
    var v9 int32 = 0
    switch (p0 + int(p1)) {
    case -1:
        if false {
            fmt.Println(v8, p1)
            var v10 uint64 = (p2 >> ((p1) & 63))
            var v11 int8 = -91
            _ = v10
            _ = v11
        }
    case -149:
        v12 := []int{}
        _ = v12
    case 590:
        v13 := GenS1{F0: -1, F1: 1}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int = -735
    v2 := []int{v1, v1, -2147483648}
    var v3 uint8 = 211
    if len(v2) > 0 {
        v2[0] = int(v3)
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    v14 := GenS2{F0: ("" + "\n\t"), F1: 9223372036854775807}
    if (v1 != v1) {
        v3 += (v3 ^ (v3 >> ((621) & 63)))
        var v15 float64 = (1e10 + 0.0)
        v1 += (v1 | (v1 * v1))
        _ = v15
    }
    if (481 < 0) {
        var v16 float32 = -2.25
        v1 -= v1
        v1 = v1
        _ = v16
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            if len(v2) > 0 {
                v2[0] = (v1 ^ (v1 ^ v1))
            }
            _ = fi17
        }
        v2 = append(v2, (v1 >> ((598) & 63)))
    }
    switch v1 {
    case -2147483648:
        v1 -= v1
    case -1:
        var v18 int32 = 1
        _ = v18
    default:
        v1 -= (v1 | 128)
    }
    var v19 int = 127
    if false {
        v20 := GenS3{F0: int16(v3)}
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v1, v1, v3)
            var v22 uint = 2
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
        switch "a" {
        case "\n\t":
            v23 := GenS2{F0: ("foo bar" + "foo bar"), F1: 42}
            _ = v23
        case "a":
            v2 = append(v2, (^(v1)))
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            v19 = v19
            fmt.Println(fi24, v1)
            _ = fi24
        }
    }
    v19 = (v1 + 0)
    var v25 uint32 = 2147483647
    v25 = (v25 * (v25 >> ((255) & 63)))
    if true {
        fmt.Println(v1)
        v25 = (v25 - v25)
    } else {
        v1 = (v19 << ((42) & 63))
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
