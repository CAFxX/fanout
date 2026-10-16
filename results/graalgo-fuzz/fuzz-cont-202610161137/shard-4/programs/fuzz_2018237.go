package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 float64 = ((((1.5 / (1e10 + 0.0)) - ((0.0 * ((((1.5 * 0.1) * 0.1) - (-2.25 * (((0.0 * 0.0) - 3.14159) + (0.1 - float64(3.14159))))) / 3.14159)) + (1.5 + (0.1 * 0.1)))) * 1.5) - 1e10)
    v2 := make(map[uint32]int)
    var v3 int16 = 1
    switch "hello" {
    case "hello":
        var v4 uint64 = 1
        _ = v4
    case "\n\t":
        v1 = v1
    case "":
        v3 += (v3 - v3)
    default:
        var v5 uint = 255
        _ = v5
    }
    v1 -= float64(0.1)
    switch 1 {
    case 32767:
        var v6 string = ""
        _ = v6
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            if (306 <= -2147483648) {
                v1 = v1
                var v8 float64 = (float64((((0.0 + 0.1) / (-2.25 - float32((((v1 - v1) - float64((float32((0.1 / float64(3.14159))) * (float32(3.14159) - (1.5 - 3.14159))))) / (0.0 - v1))))) - 1e10)) / v1)
                _ = v8
            } else {
                var v9 uint32 = 2147483648
                _ = v9
            }
            v3 = (-(v3))
            _ = fi7
        }
    }
    var v10 uint64 = 0
    v11 := []int{-893, 32768, 32767}
    var v12 int = 888
    if (true && (2147483647 > 1)) {
        switch "a" {
        case "\n\t":
            if true {
                v10 = 4294967295
            }
        case "foo bar":
            var v13 float64 = v1
            _ = v13
        case "":
            v10 = 42
        default:
            v14 := GenS1{F0: (v10 / (((v10 | v10)) | 1)), F1: 61}
            _ = v14
        }
        var v15 bool = (122 != 42)
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v12, v12, v10)
            v15 = true
            fmt.Println(fi16, v3, v1)
            _ = fi16
        }
        _ = v15
    }
    v10 += (v10 ^ (v10 + (v10 / ((42) | 1))))
    v10 -= (v10 & 1)
    v17 := GenS1{F0: v10, F1: 240}
    if (311 >= 447) {
        _ = len(v2)
    } else {
        v1 = 0.1
        switch "" {
        case "":
            fmt.Println(v1, v10)
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
}
