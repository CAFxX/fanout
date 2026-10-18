package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 string; F1 int8; F2 int64 }

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = (v1 | 2147483647)
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v1 -= (v1 % (((v1 | (v2 ^ (v1 | v2)))) | 1))
    var v5 float64 = -2.25
    var v6 float64 = v5
    var v7 float64 = (1e10 / (v6 - v6))
    var v8 uint16 = 800
    fmt.Println(v7)
    fmt.Println(v5)
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 2, F1: (1e10 <= 0.0), F2: 32767}
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                v2 += -1
                v12 := GenS2{F0: (^(v8)), F1: 793}
                _ = v12
            } else {
                var v13 int = fi9
                _ = v13
            }
            v6 = (1e10 / v5)
            v1 += (v1 & (^(v2)))
            _ = fi11
        }
        switch fi9 {
        case 4294967295:
            v14 := []int{-2147483648}
            _ = v14
        case -888:
            for fi15 := 0; fi15 < 16; fi15++ {
                v1 = v2
                _ = fi15
            }
        default:
            v7 -= 1.5
        }
        _ = fi9
        _ = v10
    }
    switch -213 {
    case -2147483648:
        switch ("hello" + "hello") {
        case "a":
            v6 = 1e10
        case "hello":
            fmt.Println(v2, v5)
        case "\n\t":
            v8 = 782
        }
    case 127:
        if false {
            switch 127 {
            case 608:
                v16 := GenS1{F0: 42, F1: (190 <= 996), F2: -609}
                _ = v16
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                if true {
                    v1 = (-(v2))
                    v18 := GenS2{F0: v8, F1: -1}
                    _ = v18
                } else {
                    v19 := []int{}
                    _ = v19
                }
                fmt.Println(v1, v6, v6)
                _ = fi17
            }
        }
    case 80:
        if false {
            for fi20 := 0; fi20 < 3; fi20++ {
                v8 += v8
                var v21 float64 = float64(0.1)
                _ = fi20
                _ = v21
            }
            fmt.Println(v7)
        } else {
            v22 := GenS3{F0: "\n\t", F1: 42, F2: 0}
            v2 -= -1
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 12; fi23++ {
            v8 += v8
            fmt.Println(v6, v6)
            var v24 int = fi23
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS2{F0: (v8 | (^(v8))), F1: -208}
    switch "" {
    case "foo bar":
        v26 := GenS3{F0: "foo bar", F1: 100, F2: 1}
        _ = v26
    case "\n\t":
        v1 = 0
    case "a":
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float64 = v6
            var v29 bool = true
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    var v30 uint8 = 2
    v31 := v3[42]
    v32 := []int{-9223372036854775808, 0, -975, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
