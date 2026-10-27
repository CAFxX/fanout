package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 uint) string {
    if (false && true) {
        v4 := GenS1{F0: false, F1: false, F2: 177}
        p1 -= p1
        _ = v4
    }
    p1 = p1
    var v5 int8 = 127
    _ = p0
    _ = p1
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint8, p1 int64, p2 int8) uint32 {
    v6 := GenS2{F0: 2147483647, F1: 510}
    var v7 uint8 = (p0 & p0)
    var v8 uint16 = 2
    var v9 int32 = -57
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 float64 = ((((float64((0.1 + 0.0)) * -2.25) / (1e10 * 1.5)) + (3.14159 * (1e10 / ((float64((((1.5 * float32(1e10)) + 0.0) / 0.1)) + 1.5) * -2.25)))) * float64(1e10))
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v10 uint64 = 4294967295
    v10 += (v10 + v10)
    if (1 > 2147483647) {
        v1 -= (0.1 - (v1 * -2.25))
        v11 := GenS1{F0: false, F1: true, F2: -770}
        _ = v11
    }
    fmt.Println(v10)
    v12 := GenS3{F0: 65535}
    var v13 uint64 = 166
    var v14 uint8 = 38
    fmt.Println(v10, v13, v10)
    var v15 int = 32767
    switch 137 {
    case -9223372036854775808:
        var v16 string = ""
        _ = v16
    default:
        fmt.Println(v13, v10, v10)
    }
    switch (v15 << ((1) & 63)) {
    case 4294967295:
        v13 += 4294967295
    case 0:
        fmt.Println(v15, v14, v1)
    }
    var v17 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
}
