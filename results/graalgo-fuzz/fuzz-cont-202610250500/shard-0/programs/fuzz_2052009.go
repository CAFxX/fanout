package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    var v7 int16 = -32768
    var v8 float32 = (1e10 + ((0.1 + 0.1) + -2.25))
    for fi9 := 0; fi9 < 10; fi9++ {
        v8 -= (v8 + (v8 - (v8 + (-2.25 - (((0.1 + v8) + v8) / (1.5 + (float32(float64(float32((1.5 / 0.1)))) - (float32(1.5) / v8))))))))
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = false
    var v3 uint8 = 0
    v4 := make(map[uint32]bool)
    var v5 int = -2147483648
    v6 := make(map[int]bool)
    v10 := GenS1{F0: -19, F1: uint32(v1), F2: (^(v1))}
    var v11 float64 = -2.25
    v11 = v11
    v5 -= 128
    v11 += 1.5
    _ = len(v6)
    v12 := GenS2{F0: 65535}
    v13 := GenS2{F0: 2147483647}
    fmt.Println(v5, v3, v3)
    v11 += float64((-2.25 + (1.5 * float32((1e10 + float64(1e10))))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
