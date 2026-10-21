package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 uint64) bool {
    fmt.Println(p0)
    var v7 float32 = 3.14159
    _ = p0
    _ = v7
    return true
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{2147483647, 65535, 65535, 2147483647}
    var v3 uint8 = v1
    v4 := make(map[int]int)
    var v5 int16 = 336
    var v6 bool = false
    v3 += 53
    v5 += v5
    v1 -= v3
    v6 = (!(1.5 == 0.1))
    delete(v4, 42)
    v8 := GenS1{F0: -632, F1: 2147483647, F2: 0}
    v3 = (v3 / ((v1) | 1))
    switch -2147483648 {
    case 128:
        fmt.Println(v6, v5, v1)
    }
    fmt.Println(v6)
    v1 = (v3 >> ((1) & 63))
    if (127 != (^(v1))) {
        var v9 bool = true
        var v10 float32 = (3.14159 / ((3.14159 - 0.1) + 0.1))
        _ = v10
        _ = v9
    } else {
        if v6 {
            for fi11 := 0; fi11 < 10; fi11++ {
                v3 = (v1 - v1)
                _ = fi11
            }
            var v12 int8 = 127
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
