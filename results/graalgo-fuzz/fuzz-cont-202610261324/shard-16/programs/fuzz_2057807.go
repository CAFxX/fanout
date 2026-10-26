package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int16 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    v3 := []int{42, 42}
    var v4 float64 = -2.25
    var v5 bool = true
    var v6 uint = 2
    for fi7 := 0; fi7 < 13; fi7++ {
        v6 = v6
        _ = fi7
    }
    v8 := GenS1{F0: 1, F1: 2, F2: 32767}
    var v9 string = ""
    var v10 uint16 = 227
    v6 += v6
    if (v6 > (^(v6))) {
        fmt.Println(v1, v1)
    }
    var v11 bool = v5
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
