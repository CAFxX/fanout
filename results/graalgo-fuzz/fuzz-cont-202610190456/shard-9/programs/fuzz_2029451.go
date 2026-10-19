package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = -1
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    var v4 int8 = 127
    var v5 float32 = ((((0.1 - float32((((1e10 - float64(1e10)) - 3.14159) * -2.25))) * 0.1) + float32(3.14159)) - -2.25)
    var v6 int8 = (v4 & (v4 | 127))
    v3[943] = 4294967295
    fmt.Println(v1)
    v7 := GenS1{F0: 4294967295}
    v1 = v1
    switch "hello" {
    case "\n\t":
        var v8 uint32 = 2
        _ = v8
    }
    var v9 int16 = (v1 & (^(v1)))
    v1 -= v9
    var v10 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
