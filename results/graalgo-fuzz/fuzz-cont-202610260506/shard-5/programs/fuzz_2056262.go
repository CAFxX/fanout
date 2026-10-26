package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint8 = 211
    v2 := []uint32{2147483648}
    var v3 int32 = -2147483648
    v3 = -2147483648
    fmt.Println(v1, v3)
    v4 := GenS1{F0: 0}
    _ = len(v2)
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 5; fi6++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v7 int8 = -56
        var v8 int = (-(fi6))
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = (-(v3))
    v9 := GenS1{F0: 250}
    fmt.Println(v1)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
