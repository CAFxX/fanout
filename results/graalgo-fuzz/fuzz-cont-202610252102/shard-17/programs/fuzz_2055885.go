package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    v3 -= v3
    fmt.Println(v3)
    v3 -= v3
    fmt.Println(v3)
    for fi4 := 0; fi4 < 4; fi4++ {
        v3 += v3
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := GenS1{F0: 42, F1: 780, F2: 255}
            v7 := v1[(fi4 + (fi5 ^ (fi4 | (fi5 | 2147483647))))]
            v3 -= (v3 * (v3 % ((v3) | 1)))
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    if (-2147483648 != 42) {
        var v8 float32 = 0.0
        if ((32767 >= 32767) || true) {
            v9 := GenS1{F0: 2147483648, F1: 18446744073709551615, F2: 42}
            v3 += v3
            if (253 > 42) {
                v8 += -2.25
            }
            _ = v9
        }
        var v10 int16 = -294
        _ = v10
        _ = v8
    }
    fmt.Println(v3)
    v3 += (v3 % (((-(v3))) | 1))
    if true {
        v3 -= v3
        fmt.Println(v3)
        v11 := GenS2{F0: 32767}
        _ = v11
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v12 := []int{4294967295, -1, 331}
    for fi13 := 0; fi13 < 19; fi13++ {
        v12 = append(v12, (fi13 | fi13))
        _ = fi13
    }
    v12 = append(v12, 32768)
    fmt.Println(v3)
    switch ("" + "\n\t") {
    case "":
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                v12 = append(v12, fi14)
                _ = fi15
            }
            _ = fi14
        }
    case "a":
        v16 := GenS2{F0: 0}
        _ = v16
    case "foo bar":
        switch 127 {
        case -2147483648:
            v17 := GenS1{F0: 1, F1: 18446744073709551615, F2: 1}
            _ = v17
        case 127:
            var v18 int16 = -266
            _ = v18
        case -9223372036854775808:
            v3 -= v3
        default:
            fmt.Println(v3)
        }
    default:
        switch "\n\t" {
        case "hello":
            var v19 int8 = v3
            _ = v19
        case "\n\t":
            var v20 uint = 255
            _ = v20
        case "a":
            if ((127 <= 85) || (377 >= 733)) {
                v3 -= (v3 | -128)
            }
        default:
            _ = len(v2)
        }
    }
    _ = len(v12)
    for fi21 := 0; fi21 < 16; fi21++ {
        if ((v3 / ((v3) | 1)) != (v3 + 127)) {
            var v22 int8 = (^(v3))
            _ = v22
        }
        if (42 != 9223372036854775807) {
            var v23 uint16 = 255
            _ = len(v2)
            var v24 int8 = (v3 * (-(v3)))
            _ = v23
            _ = v24
        } else {
            if ((!(166 >= 65535)) || true) {
                var v25 string = "a"
                var v26 int64 = -9223372036854775808
                _ = v25
                _ = v26
            } else {
                v27 := GenS1{F0: 2147483648, F1: 679, F2: 42}
                var v28 int8 = -97
                var v29 uint = 0
                _ = v27
                _ = v28
                _ = v29
            }
            var v30 uint16 = 255
            _ = v30
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 int = fi21
            v33 := GenS3{F0: "hello", F1: 128}
            var v34 uint64 = 42
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
