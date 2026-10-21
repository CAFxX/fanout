package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 uint32, p1 string, p2 int8) int8 {
    p2 = p2
    fmt.Println(p1, p2)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 int) int32 {
    v6 := GenS1{F0: 2, F1: ("hello" < "a"), F2: (-(p0))}
    fmt.Println(p0)
    var v7 int = (^(p0))
    _ = p0
    _ = v6
    _ = v7
    return -575
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]int)
    var v3 float32 = 3.14159
    var v4 uint = 912
    var v5 bool = ("" != "a")
    fmt.Println(v3, v1)
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi9, v5)
            if (70 <= 18446744073709551615) {
                var v10 uint = (^(v4))
                v11 := GenS1{F0: (v4 | v4), F1: v5, F2: v1}
                _ = v10
                _ = v11
            }
            switch "\n\t" {
            case "a":
                var v12 uint = v4
                _ = v12
            case "foo bar":
                var v13 uint = v4
                _ = v13
            case "hello":
                v5 = (-32768 == -32768)
            }
            _ = fi9
        }
        _ = fi8
    }
    v14 := GenS2{F0: 65535}
    v3 = v3
    var v15 bool = true
    var v16 uint16 = 42
    switch "\n\t" {
    case "hello":
        fmt.Println(v15)
    case "\n\t":
        fmt.Println(v3, v1, v5)
    case "a":
        var v17 string = "foo bar"
        _ = v17
    }
    fmt.Println(v16, v4)
    if (-32768 == int16(v1)) {
        v1 = (v1 << (((v4 % ((v4) | 1))) & 63))
    } else {
        v18 := []int{4294967295, v1, (v1 / ((v1) | 1))}
        _ = v18
    }
    var v19 int = 2147483647
    for fi20 := 0; fi20 < 9; fi20++ {
        if v5 {
            v3 = (float32((((1e10 + float64(v3)) * 3.14159) / ((3.14159 + (((((3.14159 + ((0.0 + (1.5 * 0.1)) / 1e10)) + (3.14159 - -2.25)) / -2.25) * 0.0) - 3.14159)) + float64(v3)))) - 1e10)
        }
        _ = fi20
    }
    switch "foo bar" {
    case "hello":
        v21 := GenS3{F0: 2147483647, F1: 759, F2: 864}
        _ = v21
    default:
        var v22 uint8 = 2
        _ = v22
    }
    fmt.Println(v19, v1, v5)
    switch v1 {
    case -977:
        v23 := GenS3{F0: 0, F1: 79, F2: 494}
        _ = v23
    default:
        var v24 bool = (!((v4 < (v4 >> ((v4) & 63))) || ((v16 * (v16 & v16)) == (v16 - (v16 >> ((v4) & 63))))))
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
