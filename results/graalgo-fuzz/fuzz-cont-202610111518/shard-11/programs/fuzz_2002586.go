package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int32, p1 int) int64 {
    v7 := GenS1{F0: 0, F1: 0}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch "" {
            case "foo bar":
                v10 := GenS2{F0: 0, F1: 2147483647, F2: 4294967295}
                _ = v10
            default:
                var v11 int8 = 31
                _ = v11
            }
            _ = fi9
        }
        _ = fi8
    }
    p0 = (p0 >> ((2) & 63))
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 int8 = 100
    var v6 int = -9223372036854775808
    fmt.Println(v6)
    v12 := GenS3{F0: 819, F1: 95}
    fmt.Println(v5, v6)
    var v13 uint64 = 18446744073709551615
    if (42 > 0) {
        v5 -= -47
        delete(v4, "hello")
        fmt.Println(v13)
    } else {
        _ = len(v2)
        fmt.Println(v6)
        var v14 float64 = (-2.25 / 0.1)
        _ = v14
    }
    switch (^(v6)) {
    case 39:
        fmt.Println(v6, v13)
    }
    var v15 int8 = (v5 & (^(v5)))
    v16 := GenS4{F0: 0}
    var v17 uint16 = 42
    v18 := GenS5{F0: 307, F1: (v13 << ((255) & 63)), F2: uint8(v17)}
    if (!(!(4294967295 >= 18446744073709551615))) {
        if true {
            v17 += (v17 * 704)
            var v19 float32 = 3.14159
            v20 := GenS3{F0: 1, F1: (v13 * v13)}
            _ = v19
            _ = v20
        }
        var v21 uint = 1
        _ = v21
    }
    v22 := GenS6{F0: 9223372036854775807}
    switch (v6 % (((v6 << ((1) & 63))) | 1)) {
    case -1:
        switch "\n\t" {
        case "foo bar":
            switch 1 {
            case 32767:
                if len(v1) > 0 {
                    v1[0] = (v6 % (((v6 * v6)) | 1))
                }
            default:
                v13 -= (^(v13))
            }
        case "hello":
            fmt.Println(v5)
        case "a":
            var v23 int64 = -9223372036854775808
            _ = v23
        default:
            v13 = (v13 ^ (v13 << ((2) & 63)))
        }
    default:
        var v24 uint = 18446744073709551615
        _ = v24
    }
    var v25 float64 = (1e10 * 0.0)
    var v26 uint64 = v13
    for fi27 := 0; fi27 < 12; fi27++ {
        v13 = 18446744073709551615
        v17 -= v17
        for fi28 := 0; fi28 < 14; fi28++ {
            v1 = append(v1, fi28)
            switch "foo bar" {
            case "a":
                switch "hello" {
                case "hello":
                    fmt.Println(v6, v13)
                case "foo bar":
                    var v29 bool = false
                    _ = v29
                case "\n\t":
                    fmt.Println(v26, v26, v25)
                default:
                    var v30 float64 = 0.1
                    _ = v30
                }
            default:
                fmt.Println(v5)
            }
            var v31 int8 = (v15 ^ v15)
            switch int(v26) {
            case 743:
                v32 := GenS5{F0: uint32(v13), F1: 18446744073709551615, F2: 145}
                _ = v32
            }
            _ = fi28
            _ = v31
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
