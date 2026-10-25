package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    var v1 bool = false
    var v2 float32 = float32(0.1)
    v3 := make(map[int]string)
    var v4 uint8 = 127
    var v5 int16 = 1
    v6 := []int{42}
    var v7 bool = true
    v8 := GenS1{F0: v5, F1: 2}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 float64 = ((float64(0.1) / float64(v2)) - (0.1 / 3.14159))
        fmt.Println(v7, v1, v5)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        switch "foo bar" {
        case "":
            v3[(fi11 - fi11)] = "a"
        case "\n\t":
            v12 := GenS1{F0: -32768, F1: 693}
            _ = v12
        case "hello":
            var v13 int = (fi11 + fi11)
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS1{F0: (v5 ^ (v5 & (-(v5)))), F1: 327}
    var v15 int16 = (v5 >> ((0) & 63))
    v2 += ((1.5 * (0.1 / 1.0)) / (v2 - (((0.1 - ((v2 - (((((v2 / (v2 * 0.0)) / ((v2 - 1.5) / (v2 * float32((1e10 / 1e10))))) * v2) - v2) - 3.14159)) / 1.5)) * v2) + 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
