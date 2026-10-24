package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]uint32)
    var v3 int16 = 37
    var v4 uint16 = v1
    var v5 int64 = 42
    switch 128 {
    case 42:
        v5 -= -9223372036854775808
    case 130:
        fmt.Println(v3, v1, v4)
    case 128:
        v4 = 2
    default:
        for fi6 := 0; fi6 < 16; fi6++ {
            switch -1 {
            case -294:
                var v7 int32 = 0
                _ = v7
            case 42:
                var v8 bool = false
                _ = v8
            case 127:
                v4 += (v4 / ((v1) | 1))
            }
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: "hello", F1: 441}
                v11 := []int{32767, (fi6 | fi6)}
                v12 := GenS2{F0: true}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := GenS3{F0: 186, F1: -927, F2: 42}
            _ = fi6
            _ = v13
        }
    }
    if (((0.1 / (((3.14159 * 1e10) + (3.14159 * 1.5)) + -2.25)) + -2.25) >= 1.5) {
        switch 4294967295 {
        case -1:
            fmt.Println(v4, v3, v1)
        case 0:
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(v5, v1)
                fmt.Println(v5, v5, fi14)
                _ = fi14
            }
        default:
            fmt.Println(v5, v1)
        }
        v15 := GenS3{F0: 19, F1: 511, F2: 37}
        _ = v15
    } else {
        fmt.Println(v5)
    }
    v3 += 42
    v4 = 2
    if false {
        if (!(true && true)) {
            for fi16 := 0; fi16 < 15; fi16++ {
                switch "foo bar" {
                case "hello":
                    if false {
                        fmt.Println(fi16, v1, v5)
                        var v17 int64 = (v5 << ((2) & 63))
                        _ = v17
                    } else {
                        v3 = (v3 >> ((2) & 63))
                        delete(v2, fi16)
                    }
                }
                _ = fi16
            }
            var v18 float64 = float64((-2.25 + 0.1))
            v19 := GenS4{F0: false, F1: 713, F2: 824}
            _ = v18
            _ = v19
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            v20 := GenS5{F0: 18446744073709551615, F1: ""}
            _ = v20
        case "foo bar":
            v21 := v2[-919]
            _ = v21
        }
    }
    v22 := GenS6{F0: 42}
    switch 2147483647 {
    case 127:
        v23 := GenS5{F0: 65535, F1: "\n\t"}
        _ = v23
    case 0:
        fmt.Println(v5)
    case 1:
        if false {
            fmt.Println(v1, v3)
            fmt.Println(v3)
        }
    }
    var v24 uint = 255
    v24 -= 450
    v1 -= (v1 - 428)
    var v25 float32 = (0.0 + -2.25)
    v25 = v25
    v26 := []int{-9223372036854775808, 32767, 2147483647}
    v1 += (v4 % ((v4) | 1))
    switch 128 {
    case -854:
        fmt.Println(v25, v4, v24)
    case -1:
        delete(v2, 211)
    }
    v27 := GenS7{F0: v24, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
