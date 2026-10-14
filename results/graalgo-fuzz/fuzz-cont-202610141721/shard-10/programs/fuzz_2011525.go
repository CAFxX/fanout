package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int64) uint16 {
    p1 -= 0
    v6 := []int{-642, 9223372036854775807, 0, 565}
    v6 = append(v6, -1)
    v7 := GenS1{F0: 14}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    var v1 float32 = (((1.5 - 3.14159) - 1.5) + (0.1 - 0.0))
    var v2 bool = false
    var v3 uint64 = 2
    v4 := []int{32767, 487, -2147483648, 404}
    var v5 int64 = 42
    var v8 int32 = -2147483648
    v9 := GenS1{F0: 42}
    v5 = (v5 & v5)
    if false {
        fmt.Println(v5, v1)
        fmt.Println(v5, v5)
    }
    if len(v4) > 0 {
        v4[0] = 32768
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v1, v3, v2)
        var v11 uint32 = 65535
        v4 = append(v4, fi10)
        _ = fi10
        _ = v11
    }
    v3 = v3
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
