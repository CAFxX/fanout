package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int8 = -1
    var v2 int = 32767
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    var v5 string = ""
    var v6 uint64 = 42
    v7 := []int{-403, -2147483648, v2, (v2 % ((-2147483648) | 1))}
    var v8 uint = 1
    _ = len(v7)
    v7 = append(v7, (v2 / ((v2) | 1)))
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 float64 = 1e10
        switch ("foo bar" + "\n\t") {
        case "hello":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            var v11 float32 = 1.5
            _ = v11
        }
        v2 -= -9223372036854775808
        _ = fi9
        _ = v10
    }
    var v12 float64 = ((0.1 - -2.25) * (((float64(float32(0.0)) + (0.1 + 0.0)) * (1e10 - 1e10)) * 3.14159))
    if true {
        _ = len(v3)
        var v13 uint8 = 255
        fmt.Println(v2, v2, v5)
        _ = v13
    } else {
        v6 -= 18446744073709551615
        v14 := v3[int(v6)]
        var v15 int16 = -32768
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 2147483648, F1: 1}
    var v17 uint32 = uint32(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
