package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
func genFn0(p0 int32) int8 {
    switch "a" {
    case "foo bar":
        p0 -= p0
    case "a":
        v6 := []int{4294967295, 2147483647}
        _ = v6
    }
    v7 := []int{323, 128, 128}
    for fi8 := 0; fi8 < 16; fi8++ {
        if (!false) {
            fmt.Println(fi8)
        }
        _ = fi8
    }
    var v9 uint8 = 127
    _ = p0
    _ = v7
    _ = v9
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    v5 := []int{-1, 1}
    var v10 int8 = -128
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 uint16 = 1
    if (9223372036854775807 == 9223372036854775807) {
        v12 := GenS1{F0: -368, F1: 78, F2: 4294967295}
        v13 := GenS1{F0: 502, F1: 238, F2: 816}
        _ = v12
        _ = v13
    }
    var v14 uint64 = 241
    v11 += (v11 & v11)
    fmt.Println(v11, v14, v14)
    fmt.Println(v10, v11)
    if (18446744073709551615 >= (v14 | (^(v14)))) {
        fmt.Println(v11, v11)
        fmt.Println(v11)
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        if (42 > 1) {
            v14 -= (v14 ^ (v14 / (((^(v14))) | 1)))
            switch ("" + "foo bar") {
            case "\n\t":
                fmt.Println(v10, v11)
            case "hello":
                v11 += (v11 % (((v11 << ((uint(v14)) & 63))) | 1))
            case "foo bar":
                delete(v1, 2147483648)
            }
            v1[2147483648] = 42
        }
        v11 = (^(v11))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi15
    }
    if true {
        _ = len(v2)
        v5 = append(v5, 4294967295)
    } else {
        if (true && (65535 <= 4294967295)) {
            for fi16 := 0; fi16 < 16; fi16++ {
                v2[32767] = fi16
                _ = fi16
            }
            if true {
                var v17 int32 = -566
                _ = v17
            } else {
                v5 = append(v5, 4294967295)
                var v18 int16 = 1
                _ = v18
            }
            v14 = v14
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int = fi19
                var v21 bool = (true && true)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        if true {
            var v22 int32 = 1
            switch ("" + "") {
            case "foo bar":
                var v23 int = -1
                _ = v23
            default:
                v24 := GenS2{F0: ("\n\t" + "\n\t"), F1: false, F2: (v22 % (((^(v22))) | 1))}
                _ = v24
            }
            switch int(v22) {
            case 378:
                var v25 uint64 = 2
                _ = v25
            case -661:
                fmt.Println(v11, v10, v14)
            }
            _ = v22
        }
    }
    var v26 uint16 = (v11 / (((v11 + v11)) | 1))
    var v27 uint16 = 2
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
