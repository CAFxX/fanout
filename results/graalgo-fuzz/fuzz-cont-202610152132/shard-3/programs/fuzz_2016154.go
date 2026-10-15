package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float32 = float32(1.5)
    var v2 uint32 = 14
    var v3 int8 = 42
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 int8 = (v3 % ((v3) | 1))
    v7 := GenS1{F0: 1}
    var v8 int32 = -852
    var v9 uint8 = 7
    v2 = (v2 * v2)
    for fi10 := 0; fi10 < 9; fi10++ {
        v6 += (v6 % (((v6 % ((v6) | 1))) | 1))
        v3 = v6
        _ = fi10
    }
    v11 := []int{-399, 127, 554, -197}
    v8 = -538
    v12 := GenS1{F0: 902}
    v8 += -25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
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
