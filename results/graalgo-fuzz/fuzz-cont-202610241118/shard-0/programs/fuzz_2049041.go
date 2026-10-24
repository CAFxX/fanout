package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int64 {
    var v6 int16 = 582
    v7 := GenS1{F0: 2}
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{2, 65535}
    v2 := []int{}
    var v3 int32 = 42
    var v4 int8 = 0
    var v5 uint32 = 0
    _ = len(v2)
    var v8 float32 = 1.5
    v8 = (0.1 / (v8 - float32(0.1)))
    fmt.Println(v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(v3, v4)
        fmt.Println(fi10, v5)
        _ = fi10
    }
    var v11 uint64 = 18446744073709551615
    v1 = append(v1, (v5 >> ((2) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
