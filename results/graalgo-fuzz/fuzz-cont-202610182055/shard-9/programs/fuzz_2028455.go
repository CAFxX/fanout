package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64 }

func main() {
    var v1 int16 = 42
    var v2 float64 = float64(0.0)
    var v3 uint64 = 1
    var v4 uint16 = 162
    v5 := GenS1{F0: 1}
    v4 = v4
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 1; fi6++ {
            switch fi6 {
            case 282:
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 float64 = v2
                    var v9 float32 = ((1.5 / 1.0) / 1.5)
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 42:
                var v10 uint8 = 1
                _ = v10
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int64 = -1
                var v13 float32 = 1e10
                _ = fi11
                _ = v12
                _ = v13
            }
            switch ("hello" + "foo bar") {
            case "":
                v14 := []int{}
                _ = v14
            case "\n\t":
                var v15 bool = (false && ((((!(!false)) || (((0.0 * 3.14159) - 1.5) < (float64((3.14159 - (1.5 / float32(v2)))) - v2))) || ((v4 % ((v4) | 1)) == 849)) || false))
                _ = v15
            case "foo bar":
                fmt.Println(v2)
            }
            _ = fi6
        }
    case "a":
        switch "hello" {
        case "a":
            v16 := GenS1{F0: 628}
            _ = v16
        case "\n\t":
            var v17 float64 = (-2.25 + -2.25)
            _ = v17
        case "":
            v18 := []int{32767, 42, 9223372036854775807}
            _ = v18
        }
    case "foo bar":
        if false {
            v19 := GenS2{F0: 42, F1: 772}
            fmt.Println(v3)
            _ = v19
        }
    default:
        v1 = v1
    }
    v20 := []int{-2147483648, 9223372036854775807, 32768}
    fmt.Println(v4)
    v21 := GenS2{F0: -1, F1: int64(v4)}
    var v22 int = -9223372036854775808
    v22 -= v22
    v23 := GenS2{F0: -86, F1: -9223372036854775808}
    var v24 int64 = 0
    switch "\n\t" {
    case "hello":
        _ = len(v20)
    }
    if (265 >= (^(v22))) {
        v25 := GenS2{F0: 0, F1: v24}
        v20 = append(v20, -336)
        _ = v25
    } else {
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(v1, fi26, v1)
            if (-128 == -105) {
                v20 = append(v20, v22)
            }
            _ = fi26
        }
    }
    var v27 float64 = (((v2 - (float64((3.14159 + 3.14159)) / v2)) / (v2 + v2)) * (v2 - v2))
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS1{F0: -242}
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    fmt.Println(v2, v2)
    var v31 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
