package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint64; F1 uint64; F2 int64 }

func main() {
    var v1 uint8 = 172
    v2 := []int{-785}
    var v3 int64 = -476
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    v6 := make(map[string]string)
    var v7 float64 = 1.5
    for fi8 := 0; fi8 < 19; fi8++ {
        switch "foo bar" {
        case "hello":
            for fi9 := 0; fi9 < 14; fi9++ {
                switch "a" {
                case "foo bar":
                    v10 := GenS1{F0: -128, F1: (-382 >= v3), F2: (true || (false && true))}
                    _ = v10
                case "hello":
                    if true {
                        v1 -= v1
                        if len(v2) > 0 {
                            v2[0] = (fi9 & fi8)
                        }
                        v11 := GenS1{F0: -128, F1: (262 <= (fi8 << ((73) & 63))), F2: true}
                        _ = v11
                    } else {
                        var v12 int8 = -51
                        _ = v12
                    }
                case "a":
                    fmt.Println(v1, fi9)
                }
                var v13 uint64 = 0
                v1 = (^(v1))
                fmt.Println(fi9, v3)
                _ = fi9
                _ = v13
            }
        case "foo bar":
            switch "" {
            case "":
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v1, fi8, fi8)
                default:
                    v1 = v1
                }
            case "hello":
                fmt.Println(v1, v3)
            case "\n\t":
                var v14 int = fi8
                _ = v14
            }
        case "\n\t":
            var v15 int64 = v3
            _ = v15
        default:
            var v16 int64 = (v3 >> ((502) & 63))
            _ = v16
        }
        v1 = 1
        _ = fi8
    }
    v7 += ((3.14159 * (v7 / v7)) / 1e10)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if false {
        v17 := GenS1{F0: -63, F1: (-1 <= -918), F2: true}
        fmt.Println(v3, v7)
        _ = v17
    } else {
        v1 += (^(v1))
    }
    var v18 uint = 255
    if false {
        v7 += 3.14159
        var v19 float32 = ((float32(1.5) / 1.0) - (3.14159 / 1e10))
        v19 -= v19
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    if true {
        if (2 == 65535) {
            if false {
                v20 := GenS2{F0: -2147483648, F1: -1, F2: 2147483648}
                if (v7 < 0.0) {
                    v21 := GenS3{F0: 887, F1: (^(v18))}
                    v2 = append(v2, 399)
                    _ = v21
                }
                if false {
                    v22 := GenS4{F0: 2, F1: 2, F2: (v3 * (v3 ^ v3))}
                    var v23 uint = (^(v18))
                    _ = v22
                    _ = v23
                }
                _ = v20
            } else {
                for fi24 := 0; fi24 < 20; fi24++ {
                    switch (fi24 ^ fi24) {
                    case -9223372036854775808:
                        var v25 int32 = 1
                        _ = v25
                    case 9223372036854775807:
                        v26 := GenS1{F0: -78, F1: (0 >= -1), F2: true}
                        _ = v26
                    case 355:
                        v3 = v3
                    }
                    var v27 uint16 = 65535
                    _ = fi24
                    _ = v27
                }
            }
            v28 := GenS4{F0: 18446744073709551615, F1: 42, F2: (v3 + v3)}
            _ = v28
        }
        var v29 int16 = int16(v3)
        _ = v29
    }
    if true {
        v18 = v18
    }
    delete(v5, ("hello" + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
