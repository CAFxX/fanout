package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0() uint16 {
    var v7 int16 = 42
    var v8 bool = true
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint64) uint32 {
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint = 536
        _ = fi9
        _ = v10
    }
    v11 := []int{1, 4294967295}
    var v12 int32 = -2147483648
    _ = p0
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []int{32768, -9223372036854775808, 9223372036854775807, -622}
    v2 := make(map[string]string)
    var v3 int = -1
    var v4 uint16 = 32767
    var v5 int64 = 1
    v6 := make(map[uint32]uint32)
    var v13 uint32 = 787
    if len(v1) > 0 {
        v1[0] = (v3 * -304)
    }
    fmt.Println(v4)
    v1 = append(v1, (v3 & v3))
    v4 = v4
    var v14 int32 = 2147483647
    var v15 float64 = 0.1
    var v16 uint32 = 539
    var v17 bool = (true && (0 != 42))
    v6[42] = (^(v13))
    var v18 uint = 4294967295
    delete(v6, v13)
    var v19 float64 = ((v15 / v15) * (3.14159 - 0.1))
    v20 := GenS1{F0: v5, F1: (v14 - int32(v5))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
