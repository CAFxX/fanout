package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32) string {
    p0 -= p0
    var v6 uint64 = 1
    fmt.Println(p0)
    v7 := GenS1{F0: 4294967295, F1: (p0 | 1)}
    _ = p0
    _ = v6
    _ = v7
    return ""
}
func genFn1() uint16 {
    var v8 int = 915
    fmt.Println(v8)
    _ = v8
    return 65535
}

func main() {
    var v1 bool = true
    var v2 float32 = ((float32(3.14159) * (1e10 / 1.5)) / 1.0)
    v3 := []uint32{}
    var v4 int16 = -1
    v5 := make(map[int]bool)
    if (90 >= 42) {
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    var v9 uint64 = 0
    fmt.Println(v1, v9, v9)
    v10 := GenS1{F0: 0, F1: 2147483647}
    if (!v1) {
        var v11 int16 = -582
        v12 := v5[2147483647]
        _ = v11
        _ = v12
    } else {
        var v13 uint32 = 4294967295
        v14 := GenS1{F0: 1, F1: 2147483647}
        var v15 bool = (v2 == 1e10)
        _ = v13
        _ = v14
        _ = v15
    }
    _ = len(v3)
    v2 += ((v2 * v2) - (v2 - float32((0.0 + -2.25))))
    var v16 float64 = 3.14159
    if false {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(v16, v4, fi17)
                var v19 int8 = 1
                _ = fi18
                _ = v19
            }
            var v20 uint32 = 4294967295
            _ = fi17
            _ = v20
        }
        v21 := GenS2{F0: 2147483647}
        v16 = (v16 / v16)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
