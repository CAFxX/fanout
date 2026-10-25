package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }

func main() {
    var v1 uint16 = 42
    var v2 uint = 2
    var v3 uint8 = 0
    var v4 float64 = -2.25
    v5 := []int{0, 1, 0, -1}
    var v6 bool = (-560 < 2147483647)
    var v7 uint16 = v1
    var v8 float64 = v4
    v7 = (v7 ^ uint16(v2))
    v9 := GenS1{F0: v2, F1: "a", F2: -32768}
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v10 := GenS1{F0: (v2 ^ (^(v2))), F1: "\n\t", F2: -917}
    for fi11 := 0; fi11 < 2; fi11++ {
        v5 = append(v5, int(v7))
        var v12 uint64 = 4294967295
        _ = fi11
        _ = v12
    }
    if v6 {
        var v13 uint32 = uint32(v3)
        v6 = false
        _ = v13
    }
    v6 = true
    var v14 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
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
