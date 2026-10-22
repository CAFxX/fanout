package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
func genFn0() int8 {
    var v5 uint32 = 2147483647
    fmt.Println(v5)
    v5 = 65535
    _ = v5
    return -128
}
func genFn1(p0 uint32, p1 int32) int8 {
    var v6 float64 = (3.14159 - 0.0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    var v3 bool = true
    v4 := []uint32{65535, 2147483647, 2}
    if (!v3) {
        _ = len(v4)
    }
    if ((!(-2.25 <= (-2.25 * 3.14159))) && (-2147483648 < -429)) {
        var v7 uint64 = (v2 << ((4294967295) & 63))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            for fi8 := 0; fi8 < 14; fi8++ {
                v2 -= v2
                _ = fi8
            }
            v3 = (702 < 816)
        } else {
            v9 := GenS1{F0: 42, F1: 127}
            _ = v9
        }
        _ = v7
    }
    v10 := GenS1{F0: 1, F1: 127}
    fmt.Println(v3, v3)
    if v3 {
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    var v12 int8 = 0
    v2 = v2
    v3 = v3
    _ = len(v4)
    v13 := GenS1{F0: 1, F1: genFn1(65535, 42)}
    var v14 float64 = float64((3.14159 - 1e10))
    if true {
        switch "" {
        case "a":
            v15 := GenS1{F0: 293, F1: v12}
            _ = v15
        default:
            var v16 uint8 = 1
            _ = v16
        }
        if v3 {
            var v17 uint = 255
            var v18 uint8 = 255
            var v19 int8 = (v12 >> ((v17) & 63))
            _ = v17
            _ = v18
            _ = v19
        }
    }
    var v20 uint32 = 2
    v21 := GenS2{F0: false, F1: 127, F2: -962}
    var v22 float32 = (float32(-2.25) * 0.0)
    v22 = v22
    v12 -= -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
