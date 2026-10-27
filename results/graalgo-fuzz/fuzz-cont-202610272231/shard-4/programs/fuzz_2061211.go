package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8, p1 int, p2 int8) uint32 {
    p2 -= p2
    var v4 uint16 = 796
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{4294967295}
    var v3 bool = true
    fmt.Println(v3)
    v1 += (v1 + 1e10)
    var v5 float32 = v1
    switch -1 {
    case 50:
        switch -9223372036854775808 {
        case 32767:
            var v6 uint8 = 127
            _ = v6
        }
    case 127:
        v7 := GenS1{F0: -238}
        _ = v7
    }
    var v8 bool = (false && (true || false))
    fmt.Println(v5, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 int = 252
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
