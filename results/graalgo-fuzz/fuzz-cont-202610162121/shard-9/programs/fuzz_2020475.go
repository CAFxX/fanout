package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 633
    v2 := make(map[string]string)
    var v3 int16 = -426
    v4 := make(map[int]bool)
    var v5 float32 = 0.1
    var v6 uint64 = 2
    var v7 uint = 18446744073709551615
    var v8 int8 = -128
    v9 := GenS1{F0: v1}
    var v10 uint32 = 0
    fmt.Println(v8, v10, v5)
    v11 := GenS1{F0: 42}
    var v12 int8 = v8
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v7 = (v7 ^ 2)
            v15 := GenS1{F0: v1}
            _ = fi14
            _ = v15
        }
        var v16 int16 = (v3 ^ (v3 << ((18446744073709551615) & 63)))
        v16 = (v3 | -32768)
        _ = fi13
        _ = v16
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
