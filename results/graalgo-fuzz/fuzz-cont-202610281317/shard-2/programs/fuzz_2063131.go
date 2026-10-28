package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int16 = 345
    v2 := []uint32{}
    var v3 uint32 = 4294967295
    v4 := []int{}
    var v5 int64 = 0
    var v6 int64 = -113
    v6 = (^(v6))
    v5 -= 42
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 += v6
        _ = fi7
    }
    fmt.Println(v5, v3)
    var v8 int = 4294967295
    v4 = append(v4, 127)
    v8 += (v8 & (v8 + (v8 - (v8 & -2147483648))))
    var v9 int8 = -1
    v10 := GenS1{F0: v3}
    fmt.Println(v5, v1)
    fmt.Println(v6, v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
