package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 int, p1 uint16) uint32 {
    v6 := GenS1{F0: (42 != 42), F1: -128}
    var v7 int64 = -710
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 bool = (0.1 != 0.0)
    var v2 uint32 = 42
    v3 := make(map[int]bool)
    var v4 uint8 = 127
    var v5 bool = v1
    v8 := []int{2147483647, -2147483648}
    if (2 < 1) {
        var v9 int32 = 42
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := v3[fi10]
            _ = fi10
            _ = v11
        }
        v3[42] = v1
        _ = v9
    }
    v5 = v5
    delete(v3, 128)
    delete(v3, -1)
    if len(v8) > 0 {
        v8[0] = 589
    }
    v12 := GenS1{F0: (-854 > -1), F1: 127}
    delete(v3, -510)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
