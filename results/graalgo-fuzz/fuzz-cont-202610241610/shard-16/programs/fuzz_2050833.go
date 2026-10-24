package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 bool, p2 uint32) uint16 {
    var v4 uint8 = 42
    var v5 uint = 18446744073709551615
    if ((3.14159 + 0.0) == (0.0 - (3.14159 * 1e10))) {
        p1 = false
    } else {
        v6 := GenS1{F0: 81}
        _ = v6
    }
    p2 -= p2
    v7 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v7
    return 65535
}

func main() {
    var v1 int32 = -511
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    var v8 int16 = -32768
    var v9 int = -1
    fmt.Println(v9, v8)
    fmt.Println(v9, v1, v3)
    for fi10 := 0; fi10 < 18; fi10++ {
        v8 += v8
        v1 += (v1 << ((919) & 63))
        _ = fi10
    }
    v3 += (1e10 / v3)
    v8 -= -32768
    v8 -= v8
    v3 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
