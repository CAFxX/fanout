package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 2
    v2 := make(map[string]int)
    var v3 uint32 = 2
    var v4 int16 = 32767
    v5 := make(map[string]string)
    v6 := GenS1{F0: 2}
    v4 -= 1
    v3 += (v3 * v3)
    fmt.Println(v4, v4, v1)
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            v4 += v4
            fmt.Println(v4, fi7, fi7)
            fmt.Println(fi7)
        }
        var v8 uint = 756
        _ = fi7
        _ = v8
    }
    v9 := []int{-9223372036854775808, 722, -24}
    v1 -= (v1 - v1)
    for fi10 := 0; fi10 < 14; fi10++ {
        v4 = (v4 + (^(v4)))
        v11 := GenS1{F0: v3}
        _ = fi10
        _ = v11
    }
    if len(v9) > 0 {
        v9[0] = -657
    }
    v9 = append(v9, 4294967295)
    _ = len(v5)
    var v12 uint8 = 2
    var v13 uint32 = (v3 - v3)
    var v14 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
