package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 3.14159
    v3 := make(map[string]bool)
    v4 := []uint32{}
    _ = len(v1)
    delete(v1, 505)
    switch -9223372036854775808 {
    case -840:
        v2 += -2.25
    case -2147483648:
        switch 127 {
        case -9223372036854775808:
            v5 := GenS1{F0: 2, F1: 157, F2: -50}
            _ = v5
        case 32767:
            fmt.Println(v2)
        case -1:
            v6 := GenS1{F0: 737, F1: 255, F2: -451}
            _ = v6
        }
    case 65:
        for fi7 := 0; fi7 < 14; fi7++ {
            v2 = (v2 / v2)
            v8 := GenS1{F0: 145, F1: 241, F2: 1}
            _ = fi7
            _ = v8
        }
    }
    switch "hello" {
    case "\n\t":
        var v9 int32 = 707
        _ = v9
    case "":
        var v10 uint32 = 42
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: -128}
        _ = fi11
        _ = v12
    }
    var v13 int64 = -878
    v2 = (float32(-2.25) - v2)
    var v14 int16 = 0
    switch 42 {
    case 147:
        v15 := GenS1{F0: 459, F1: 2, F2: v13}
        _ = v15
    }
    var v16 float32 = (((v2 * (0.1 + v2)) + 0.0) - v2)
    switch "a" {
    case "foo bar":
        if true {
            switch "a" {
            case "hello":
                var v17 uint8 = 1
                _ = v17
            case "foo bar":
                v4 = append(v4, 2)
            }
            var v18 uint8 = 0
            _ = v18
        } else {
            if false {
                var v19 float32 = v16
                v19 += v19
                _ = v19
            } else {
                v20 := GenS1{F0: 2147483647, F1: 114, F2: v13}
                if (0 <= 2147483648) {
                    v14 += (v14 | (v14 >> ((0) & 63)))
                    var v21 uint8 = 0
                    _ = v21
                }
                _ = v20
            }
        }
    case "":
        var v22 int8 = -94
        _ = v22
    case "a":
        switch 2147483647 {
        case 127:
            for fi23 := 0; fi23 < 4; fi23++ {
                v14 -= (v14 << ((65535) & 63))
                if true {
                    v2 = v2
                    var v24 uint8 = 2
                    _ = v24
                }
                switch "hello" {
                case "\n\t":
                    var v25 int32 = 42
                    _ = v25
                case "a":
                    v13 = (v13 << ((2) & 63))
                default:
                    var v26 string = "foo bar"
                    _ = v26
                }
                _ = fi23
            }
        case -518:
            v4 = append(v4, 848)
        case -1:
            v3["a"] = false
        }
    default:
        switch "hello" {
        case "foo bar":
            v16 -= v16
        case "hello":
            var v27 uint8 = 34
            _ = v27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
