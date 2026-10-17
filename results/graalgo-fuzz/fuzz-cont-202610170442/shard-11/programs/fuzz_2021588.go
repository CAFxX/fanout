package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint64, p2 int8) uint64 {
    var v6 uint32 = 0
    var v7 uint8 = 255
    var v8 float32 = (0.1 - float32(float64((3.14159 + 0.0))))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = -62
    var v2 float64 = (float64(1.5) - 3.14159)
    var v3 uint = 0
    var v4 int16 = -703
    var v5 uint64 = 18446744073709551615
    v9 := GenS1{F0: (v5 * v5), F1: int32(v3)}
    switch 128 {
    case 629:
        var v10 bool = false
        _ = v10
    default:
        v11 := GenS2{F0: 760}
        _ = v11
    }
    var v12 int64 = 970
    v13 := []int{42, 1}
    v1 -= 192
    var v14 uint64 = (^(v5))
    v15 := GenS2{F0: 2}
    v14 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
