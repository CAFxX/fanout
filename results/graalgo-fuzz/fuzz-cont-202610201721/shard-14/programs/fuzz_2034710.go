package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8, p1 int64, p2 int16) int8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(p1, p0, fi6)
        _ = fi6
    }
    v7 := GenS1{F0: 42, F1: -448, F2: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 string, p1 string) int8 {
    v8 := []int{4294967295, 862}
    _ = len(v8)
    v8 = append(v8, 86)
    _ = p0
    _ = p1
    _ = v8
    return 127
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 uint16 = 1
    var v4 int = -9223372036854775808
    var v5 int16 = -508
    for fi9 := 0; fi9 < 10; fi9++ {
        if (1e10 != (3.14159 - float32(float64(1.5)))) {
            v3 = (v3 % ((v3) | 1))
            v5 -= -1
            var v10 bool = true
            _ = v10
        }
        _ = fi9
    }
    if (4294967295 < 65535) {
        v11 := GenS1{F0: 772, F1: 32768, F2: 42}
        _ = v11
    }
    var v12 uint8 = (v2 | (v2 % ((128) | 1)))
    v2 -= v2
    var v13 int32 = -1
    switch v4 {
    case 638:
        var v14 int64 = 0
        _ = v14
    default:
        var v15 uint = 65535
        _ = v15
    }
    switch "hello" {
    case "\n\t":
        v4 += 780
    case "hello":
        switch "a" {
        case "foo bar":
            var v16 int = v4
            _ = v16
        case "\n\t":
            v17 := GenS2{F0: 1, F1: 800, F2: v4}
            _ = v17
        }
    case "foo bar":
        var v18 bool = (1.5 <= 0.0)
        _ = v18
    default:
        v1[v4] = ((("foo bar" + "hello") + (("a" + "a") + "a")) + "\n\t")
    }
    if false {
        var v19 int16 = v5
        var v20 uint16 = (v3 & 32767)
        v3 = 32767
        _ = v19
        _ = v20
    }
    v3 = 1
    switch "hello" {
    case "hello":
        v4 = v4
    case "\n\t":
        var v21 uint32 = 2147483647
        _ = v21
    }
    v22 := GenS3{F0: 9223372036854775807, F1: v3}
    var v23 string = ""
    var v24 string = v23
    v25 := GenS2{F0: 65535, F1: (v5 ^ 514), F2: (v4 + 0)}
    var v26 uint16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
