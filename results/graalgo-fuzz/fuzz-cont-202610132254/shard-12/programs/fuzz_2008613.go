package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    v1 := []int{}
    var v2 int16 = 354
    var v3 float32 = ((1e10 / 1.5) * ((1e10 * 0.1) - -2.25))
    v4 := make(map[string]int)
    var v5 int64 = 1
    v6 := GenS1{F0: false, F1: 0}
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: (255 > 18446744073709551615), F1: 2147483648}
        _ = fi7
        _ = v8
    }
    v2 += v2
    v1 = append(v1, 0)
    v9 := GenS2{F0: 1, F1: v2}
    switch -278 {
    case 759:
        v4["\n\t"] = 529
    case 1:
        v5 = v5
    }
    var v10 float64 = (((3.14159 / 1.0) + float64(v3)) + 1.5)
    var v11 int64 = (v5 | v5)
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(fi12, v3)
        var v13 int8 = 42
        v11 = -1
        _ = fi12
        _ = v13
    }
    if true {
        var v14 uint64 = 145
        v15 := GenS2{F0: 42, F1: -1}
        _ = v14
        _ = v15
    } else {
        v5 = 0
        v16 := GenS1{F0: (163 == 1), F1: 0}
        _ = v16
    }
    switch -9223372036854775808 {
    case -1:
        var v17 uint8 = 0
        _ = v17
    case -565:
        switch "\n\t" {
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                var v18 float32 = 0.1
                _ = v18
            case "":
                var v19 uint8 = 98
                _ = v19
            default:
                v20 := GenS3{F0: (v11 / ((v5) | 1)), F1: -595}
                _ = v20
            }
        case "hello":
            v11 += v11
        }
    }
    fmt.Println(v5, v2)
    v11 = (-(v11))
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(v11, v11)
        _ = fi21
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
