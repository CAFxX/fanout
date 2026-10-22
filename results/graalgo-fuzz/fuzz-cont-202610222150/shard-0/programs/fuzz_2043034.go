package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32) int8 {
    p0 -= p0
    switch -147 {
    case 4294967295:
        var v5 float64 = ((3.14159 - -2.25) + 1.5)
        _ = v5
    case 9223372036854775807:
        fmt.Println(p0)
    case 32768:
        var v6 int16 = 0
        _ = v6
    }
    _ = p0
    return 100
}
func genFn1() int32 {
    v7 := []int{2147483647, -506, 42, -9223372036854775808}
    var v8 int8 = 42
    _ = v7
    _ = v8
    return -532
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1e10
    var v3 float64 = -2.25
    v4 := []int{57, 715, -2147483648}
    v2 += v2
    var v9 uint = 411
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v10 uint8 = 149
    delete(v1, 2147483648)
    var v11 uint16 = 32767
    delete(v1, 4294967295)
    v12 := GenS1{F0: 691}
    if len(v4) > 0 {
        v4[0] = -536
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
