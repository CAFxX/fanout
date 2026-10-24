package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint8, p1 int8) int32 {
    switch "a" {
    case "foo bar":
        switch "hello" {
        case "hello":
            v7 := GenS1{F0: 426}
            _ = v7
        case "\n\t":
            fmt.Println(p0, p0)
        case "a":
            var v8 int16 = 1000
            _ = v8
        }
    }
    var v9 bool = false
    v10 := GenS2{F0: 1}
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v11 := []int{42, -79, 127}
        _ = v11
    case "\n\t":
        var v12 int8 = (p1 ^ (-(p1)))
        _ = v12
    default:
        var v13 uint32 = 2147483647
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 int = -817
    v2 := []uint32{127, 360, 683}
    var v3 float32 = (1.5 - (1e10 / -2.25))
    var v4 int64 = -370
    v5 := []uint32{0}
    v6 := []uint32{2147483648, 503, 65535}
    v6 = append(v6, 2147483647)
    v1 = v1
    switch "hello" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v14 int8 = -128
    switch v1 {
    case -382:
        v3 = -2.25
    case 32768:
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(v3, v1)
            _ = fi15
        }
    case -309:
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 float32 = (3.14159 + float32((1e10 / (0.1 * 1e10))))
            _ = fi16
            _ = v17
        }
    default:
        var v18 float64 = (-2.25 / 0.1)
        _ = v18
    }
    v19 := GenS2{F0: 42}
    v20 := GenS2{F0: 2}
    fmt.Println(v4, v3)
    v21 := GenS1{F0: -2147483648}
    v1 += 128
    fmt.Println(v14)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint64 = 4294967295
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS3{F0: "\n\t"}
            fmt.Println(v14, v23, v14)
            switch fi22 {
            case 128:
                if (564 < 255) {
                    v26 := GenS4{F0: 176, F1: (v23 % ((42) | 1))}
                    _ = v26
                }
            }
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v4 -= -505
    var v27 uint = 18446744073709551615
    for fi28 := 0; fi28 < 15; fi28++ {
        _ = len(v6)
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 bool = (false && (65535 < 1))
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 float64 = (0.1 - (0.0 + 0.0))
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 bool = false
        v34 := GenS5{F0: 28, F1: 2147483647}
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
