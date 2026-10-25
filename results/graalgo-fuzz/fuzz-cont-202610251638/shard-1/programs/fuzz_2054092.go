package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 int32, p1 bool) uint64 {
    v4 := []int{}
    v5 := GenS1{F0: 1, F1: 2}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint16) int8 {
    var v6 int32 = -649
    v6 = (v6 ^ (^(v6)))
    if (!(0.0 <= (0.0 * (1.5 / 1e10)))) {
        fmt.Println(v6, v6)
        for fi7 := 0; fi7 < 6; fi7++ {
            if (p0 < (p0 / ((p0) | 1))) {
                var v8 uint16 = p0
                _ = v8
            } else {
                var v9 uint64 = 4294967295
                v10 := GenS1{F0: 42, F1: v9}
                _ = v10
                _ = v9
            }
            var v11 string = "a"
            _ = fi7
            _ = v11
        }
        var v12 uint16 = 65535
        _ = v12
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    var v13 uint16 = 255
    v2 -= v2
    var v14 float32 = 0.1
    var v15 float32 = v14
    v16 := GenS1{F0: 9223372036854775807, F1: 614}
    _ = len(v3)
    v17 := GenS1{F0: -9223372036854775808, F1: 2}
    fmt.Println(v15, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
