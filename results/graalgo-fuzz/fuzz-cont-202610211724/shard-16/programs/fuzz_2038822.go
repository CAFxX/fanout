package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
func genFn0(p0 int64, p1 uint, p2 uint16) int64 {
    if (-503 != 1) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 bool = (true && true)
            switch "foo bar" {
            case "foo bar":
                switch (fi6 / ((4294967295) | 1)) {
                case 1:
                    v7 = (71 < 0)
                case 32768:
                    p0 = p0
                default:
                    fmt.Println(p0)
                }
            case "a":
                v8 := GenS1{F0: 0}
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        v9 := []int{-2147483648, 9223372036854775807, 127}
        _ = v9
    }
    p1 = (^(p1))
    if (1e10 > -2.25) {
        var v10 bool = true
        p2 = p2
        v11 := []int{1}
        _ = v10
        _ = v11
    } else {
        if true {
            var v12 int8 = -82
            var v13 uint64 = 0
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: p2}
            fmt.Println(p2)
            _ = v14
        }
        p0 = (p0 ^ (p0 / (((p0 * 0)) | 1)))
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS2{F0: 0, F1: -1, F2: 2}
        var v17 int64 = (p0 & -743)
        var v18 uint32 = 2147483647
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []int{-9223372036854775808, 4294967295, -348}
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 bool = (((1e10 / 0.1) * -2.25) < ((float64(1e10) / 3.14159) + (1.5 / 3.14159)))
    v5 := []string{"foo bar"}
    var v19 int16 = 220
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v19 += (v19 >> ((255) & 63))
    switch v3 {
    case "hello":
        v20 := GenS1{F0: 646}
        _ = v20
    case "foo bar":
        switch 42 {
        case 1:
            fmt.Println(v19)
        case -1:
            v19 += (^(v19))
        }
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v3)
    case "\n\t":
        v19 -= v19
    }
    var v21 bool = v4
    if (2147483647 <= -1) {
        if ((false && v4) && (!false)) {
            if false {
                fmt.Println(v3, v3)
            }
            var v22 uint64 = 0
            if v21 {
                v21 = v4
            }
            _ = v22
        }
        switch v3 {
        case "hello":
            v21 = v21
        case "foo bar":
            var v23 int16 = -586
            _ = v23
        default:
            v1 = append(v1, -9223372036854775808)
        }
        v24 := GenS3{F0: 128, F1: 0, F2: -75}
        _ = v24
    }
    var v25 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
