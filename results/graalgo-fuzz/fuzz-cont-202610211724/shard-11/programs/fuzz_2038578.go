package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 string = "\n\t"
    var v2 float64 = float64((1e10 + (3.14159 / (0.1 / ((0.0 + 1.5) - -2.25)))))
    v3 := []uint32{2147483647, 4294967295}
    v4 := make(map[uint32]string)
    v2 = 0.0
    var v5 int = 127
    v5 -= (v5 * v5)
    fmt.Println(v2, v1, v2)
    var v6 int64 = 1
    var v7 int = (v5 / (((v5 ^ v5)) | 1))
    var v8 int = (v7 & -1)
    var v9 float32 = 3.14159
    for fi10 := 0; fi10 < 14; fi10++ {
        v9 = (v9 * 1e10)
        for fi11 := 0; fi11 < 17; fi11++ {
            v4[120] = v1
            v3 = append(v3, 0)
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(v6)
    v12 := GenS1{F0: 42}
    var v13 int32 = -1
    v14 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
