package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }
func genFn0() bool {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: -1, F1: -9223372036854775808}
        if (682 >= 42) {
            var v8 uint16 = 956
            v9 := []int{-527}
            _ = v8
            _ = v9
        }
        var v10 string = ""
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 float64 = 0.0
    var v12 int16 = 679
    v11 += (float64(((0.1 + 0.1) - 1.5)) - (3.14159 * v11))
    var v13 int = 649
    _ = v11
    _ = v12
    _ = v13
    return true
}

func main() {
    var v1 float64 = ((0.1 + 0.0) + (1e10 - (0.1 * float64(float32(1.5)))))
    v2 := make(map[string]string)
    var v3 uint64 = 748
    var v4 int32 = 2147483647
    v5 := make(map[string]string)
    fmt.Println(v1, v3)
    v3 = 18446744073709551615
    v3 -= (v3 >> ((0) & 63))
    v4 = (v4 - v4)
    if (!true) {
        if true {
            v3 -= v3
            var v14 uint32 = 65535
            _ = v14
        }
        var v15 float32 = 1.5
        var v16 int = 128
        _ = v15
        _ = v16
    } else {
        v1 += (v1 - 0.0)
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := []int{fi17, fi17}
        _ = fi17
        _ = v18
    }
    v3 = v3
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v1, v1)
        fmt.Println(v3, v3, v4)
        var v20 float32 = -2.25
        _ = fi19
        _ = v20
    }
    if (127 != 117) {
        v21 := GenS2{F0: -1, F1: 4294967295, F2: 0}
        _ = v21
    }
    var v22 float32 = 1e10
    v23 := GenS2{F0: 100, F1: 296, F2: 255}
    fmt.Println(v4, v22)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
