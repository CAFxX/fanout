package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]uint32)
    var v3 string = "a"
    var v4 string = ((v3 + v3) + v3)
    var v5 uint32 = 2147483648
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v4 = ("\n\t" + v3)
            v1 -= ((((0.1 + v1) - v1) / (0.1 * v1)) / 1e10)
            _ = fi7
        }
        _ = fi6
    }
    v5 = v5
    var v8 float32 = (v1 / (1e10 + v1))
    fmt.Println(v1, v3, v5)
    v9 := GenS1{F0: 0, F1: 344}
    var v10 int16 = 1
    var v11 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
