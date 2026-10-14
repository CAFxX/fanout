package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint) int8 {
    p0 = (p0 >> ((p0) & 63))
    v5 := []int{32767, 42, 127}
    fmt.Println(p0)
    var v6 float32 = ((-2.25 / (0.1 * 1.5)) - float32(((1.5 * float64(float32(0.0))) - float64((3.14159 / (1.5 + 1e10))))))
    _ = p0
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 int8, p1 int8) uint64 {
    fmt.Println(p1)
    p0 += genFn0(0)
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 103
    var v2 int32 = -572
    v3 := []uint32{2, 2, 65535}
    var v4 uint = 1
    v1 -= 1
    v1 = (v1 | 255)
    var v7 float64 = 3.14159
    v4 += (^(v4))
    if (1 <= 9) {
        var v8 uint = v4
        fmt.Println(v7, v7)
        _ = v8
    }
    v7 -= v7
    var v9 float32 = 1.5
    var v10 float32 = v9
    var v11 uint16 = 42
    if false {
        if (!(-1 <= -614)) {
            fmt.Println(v10, v1)
        }
        var v12 uint32 = 42
        _ = v12
    }
    v13 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
