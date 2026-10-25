package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
func genFn0(p0 bool, p1 int32, p2 uint) uint32 {
    p0 = true
    var v7 int64 = -855
    switch 9223372036854775807 {
    case 2147483647:
        var v8 uint = 1
        _ = v8
    }
    var v9 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 500
}
func genFn1(p0 uint) uint {
    p0 -= (p0 << (((p0 * p0)) & 63))
    p0 = 4294967295
    _ = p0
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 498
    var v3 uint16 = 405
    v4 := make(map[int]string)
    var v5 float32 = (float32(((((3.14159 - (-2.25 - 1e10)) - (-2.25 + ((0.1 - (float64(3.14159) / 1.0)) / 0.1))) + 0.1) - -2.25)) / -2.25)
    var v6 uint32 = 1
    v5 = v5
    for fi10 := 0; fi10 < 17; fi10++ {
        v2 -= v2
        _ = fi10
    }
    fmt.Println(v5)
    v11 := []int{32768, 124, -1, -321}
    var v12 int64 = 396
    var v13 uint32 = 2
    v13 = v6
    var v14 string = ""
    v3 -= (v3 + (v3 - (^(v3))))
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: 32767, F1: 42, F2: (v14 + v14)}
            _ = fi16
            _ = v17
        }
        v13 = (v13 * 1)
        v13 = v13
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
