package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 uint8, p2 int8) int8 {
    v7 := GenS1{F0: -173, F1: 828}
    p0 = int64(p2)
    if true {
        p1 += (p1 << ((18446744073709551615) & 63))
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 ^ p2)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 uint8 = 128
    var v5 int8 = 42
    var v6 int64 = -292
    v8 := GenS1{F0: 42, F1: 0}
    var v9 float32 = float32(((1.5 - (-2.25 / 1.0)) + 1e10))
    var v10 uint64 = 18446744073709551615
    v3["a"] = (true && (255 != (v2 >> ((2) & 63))))
    v11 := GenS2{F0: -2147483648}
    v3["hello"] = false
    v2 = (v2 >> ((v2) & 63))
    delete(v3, ("foo bar" + "hello"))
    v9 += float32(0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
