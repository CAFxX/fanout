package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    v3 := []string{}
    v4 := []int{0, -982, 389, 455}
    v1[""] = 115
    v5 := GenS1{F0: -110}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6 := GenS1{F0: 100}
    v7 := GenS2{F0: 0}
    fmt.Println(v2)
    var v8 uint32 = uint32(v2)
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float64 = 1e10
        v11 -= ((-2.25 * ((0.1 * (((v11 - 3.14159) - (((v11 * (v11 / ((0.0 * v11) / (v11 / 1.0)))) + (3.14159 / (v11 + (3.14159 / ((float64(float32(0.1)) + (-2.25 + ((float64(3.14159) / (1e10 / float64(float32(((v11 * float64(-2.25)) - 1.5))))) - v11))) + (0.0 * v11)))))) + (0.0 * 1e10))) + 1e10)) / v11)) * ((1.5 * v11) + 0.0))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int16 = -1
        if true {
            var v14 int = (fi12 / (((fi12 / ((fi12) | 1))) | 1))
            for fi15 := 0; fi15 < 14; fi15++ {
                v9 = v9
                var v16 int8 = 100
                v13 += (v13 + 336)
                v17 := v1[v9]
                v3 = append(v3, (("" + (v9 + v9)) + "a"))
                v18 := GenS3{F0: -1, F1: 9223372036854775807}
                var v19 int64 = 9223372036854775807
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: (v8 * (v8 >> (((v2 | (v2 / ((834) | 1)))) & 63)))}
            _ = v14
            _ = v20
        } else {
            var v21 bool = true
            v22 := GenS2{F0: (^(v8))}
            _ = v21
            _ = v22
        }
        _ = fi12
        _ = v13
    }
    v8 = (v8 - 42)
    for fi23 := 0; fi23 < 13; fi23++ {
        v9 = v9
        _ = fi23
    }
    switch ("\n\t" + v9) {
    case "\n\t":
        v2 -= 65535
    case "foo bar":
        v9 = ""
    default:
        v4 = append(v4, 128)
    }
    if true {
        var v24 string = (("" + v9) + v9)
        fmt.Println(v24, v8, v2)
        v25 := GenS1{F0: 0}
        _ = v24
        _ = v25
    }
    fmt.Println(v9)
    var v26 int8 = 39
    v1[v9] = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
