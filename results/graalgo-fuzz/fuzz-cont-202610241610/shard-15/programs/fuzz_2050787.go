package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("hello" + "\n\t"), "", "hello", "hello"}
    var v3 int16 = 42
    var v4 bool = false
    var v5 uint8 = 127
    v1[0] = 209
    v4 = (v4 || false)
    var v6 uint16 = 1
    var v7 float64 = ((1e10 / ((-2.25 / 3.14159) + 0.1)) * float64((-2.25 - ((float32(float64(0.0)) * 3.14159) / (((0.1 * float32(0.0)) / ((0.1 - float32((float64(1e10) + (3.14159 - 0.0)))) - 0.0)) - (-2.25 / 1.0))))))
    var v8 int8 = 0
    var v9 uint8 = v5
    v6 = v6
    v3 = (v3 & -1)
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((v3 >> ((37) & 63)) != v3) {
            fmt.Println(v9)
            var v11 uint16 = 0
            v11 -= (v6 / ((v11) | 1))
            _ = v11
        } else {
            v12 := v1[42]
            v13 := GenS1{F0: (v6 << ((uint(v12)) & 63)), F1: uint64(v12)}
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            v2 = append(v2, "\n\t")
            v15 := GenS1{F0: v6, F1: 0}
            _ = fi14
            _ = v15
        }
        v5 = (v5 & (^(v5)))
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
