package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := []string{"\n\t"}
    var v4 uint32 = 2
    v2[-444] = 9223372036854775807
    switch "\n\t" {
    case "a":
        var v5 float32 = (3.14159 - ((-2.25 + (float32((0.0 - 1.5)) - (1.5 / 1.5))) / ((float32(((-2.25 * 0.1) + (3.14159 / (0.1 - float64(-2.25))))) - (0.0 + 1.5)) - -2.25)))
        _ = v5
    default:
        if true {
            v6 := GenS1{F0: 136}
            _ = v6
        } else {
            fmt.Println(v4)
        }
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    var v9 uint = 748
    var v10 int16 = -32768
    var v11 int8 = -128
    v10 -= (v10 / (((^(v10))) | 1))
    delete(v1, 2147483647)
    fmt.Println(v9, v10, v4)
    fmt.Println(v10)
    var v12 int64 = -9223372036854775808
    v13 := GenS1{F0: v12}
    var v14 bool = false
    if false {
        if v14 {
            var v15 int32 = -1
            v16 := GenS2{F0: v11, F1: (v15 << ((v9) & 63))}
            var v17 int32 = v15
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 uint = v9
        switch 2147483647 {
        case 42:
            v9 = 4294967295
        default:
            fmt.Println(v4, v10, v12)
        }
        _ = v18
    }
    if (!((v9 ^ 103) < 68)) {
        v11 += v11
        var v19 uint = 67
        _ = v19
    } else {
        v14 = false
        var v20 uint = 18446744073709551615
        _ = v20
    }
    fmt.Println(v14, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
