package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int8 = 42
    var v2 uint16 = 2
    var v3 uint8 = 103
    v4 := []int{-9223372036854775808, 42, 2147483647}
    _ = len(v4)
    if true {
        v1 += (v1 >> ((18446744073709551615) & 63))
    }
    switch 2147483647 {
    case -1:
        switch "\n\t" {
        case "\n\t":
            var v5 float64 = 1e10
            _ = v5
        case "a":
            v3 = v3
        case "":
            if true {
                fmt.Println(v2, v3)
                var v6 uint = 4294967295
                v7 := GenS1{F0: -60, F1: 42}
                _ = v6
                _ = v7
            }
        }
    case -577:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    default:
        var v8 int8 = v1
        _ = v8
    }
    switch "" {
    case "\n\t":
        switch 128 {
        case 42:
            var v9 int32 = 0
            _ = v9
        case -9223372036854775808:
            var v10 string = ""
            _ = v10
        case 32767:
            v11 := GenS2{F0: 933}
            _ = v11
        }
    case "hello":
        var v12 uint32 = 303
        _ = v12
    default:
        switch "" {
        case "\n\t":
            var v13 float64 = 1e10
            _ = v13
        case "hello":
            var v14 string = "\n\t"
            _ = v14
        case "foo bar":
            v15 := GenS1{F0: -1, F1: 0}
            _ = v15
        default:
            v1 += 1
        }
    }
    if (!false) {
        _ = len(v4)
        switch "\n\t" {
        case "a":
            fmt.Println(v2, v1, v2)
        default:
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 string = ""
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v19 := GenS2{F0: -9223372036854775808}
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS3{F0: "\n\t", F1: 1, F2: 54}
    var v21 int8 = -8
    v22 := GenS4{F0: 2147483648}
    fmt.Println(v3, v1, v1)
    _ = len(v4)
    v1 = (v21 & (v1 & -1))
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint64 = 0
        _ = fi23
        _ = v24
    }
    if (1 < 42) {
        var v25 float64 = (0.1 / float64(0.1))
        _ = v25
    }
    var v26 int = 0
    if len(v4) > 0 {
        v4[0] = (v26 | 127)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
