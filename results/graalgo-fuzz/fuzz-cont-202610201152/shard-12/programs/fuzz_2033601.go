package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint16 = 2
    v2 := []string{"a", "\n\t"}
    var v3 int64 = -1
    var v4 uint64 = 4294967295
    var v5 uint32 = 373
    v5 = (v5 >> ((755) & 63))
    _ = len(v2)
    fmt.Println(v1, v4)
    var v6 int8 = int8(v3)
    var v7 float64 = 0.0
    v7 -= float64((3.14159 + 0.0))
    v8 := GenS1{F0: (v4 - v4)}
    var v9 int8 = (-(v6))
    for fi10 := 0; fi10 < 1; fi10++ {
        v6 = v6
        fmt.Println(v3)
        v11 := GenS1{F0: (v4 >> ((42) & 63))}
        _ = fi10
        _ = v11
    }
    v4 -= (^(v4))
    var v12 uint32 = v5
    v2 = append(v2, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
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
