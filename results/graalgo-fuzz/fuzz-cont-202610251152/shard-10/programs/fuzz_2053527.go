package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }

func main() {
    var v1 float32 = (-2.25 + float32(float64(1e10)))
    v2 := []string{"foo bar", "foo bar", "\n\t", "\n\t"}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 531
    var v5 float64 = 0.1
    _ = len(v2)
    v4 = (v4 % ((2) | 1))
    v6 := GenS1{F0: 543, F1: 977, F2: 9223372036854775807}
    v7 := v3[65535]
    v7 = v7
    v8 := GenS1{F0: (v4 / ((v4) | 1)), F1: 999, F2: 771}
    if (!v7) {
        _ = len(v2)
        v1 -= (0.0 / v1)
        var v9 float32 = (v1 + float32((v5 - (1e10 * -2.25))))
        _ = v9
    }
    var v10 int = -215
    v7 = v7
    v10 = -2147483648
    v11 := GenS1{F0: (v4 / ((2) | 1)), F1: 106, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
