package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0(p0 int8) string {
    p0 = 21
    v7 := GenS1{F0: 4294967295, F1: 42, F2: "foo bar"}
    _ = p0
    _ = v7
    return "foo bar"
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]bool)
    var v3 int64 = 1
    var v4 uint16 = 32767
    v5 := []uint32{673, 4294967295}
    v6 := []int{337, int(v3), -9223372036854775808, 42}
    v4 += (v4 & 139)
    var v8 float64 = (float64(float32((1.5 - float64(0.0)))) * v1)
    fmt.Println(v3, v1)
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    fmt.Println(v4, v4)
    v3 = v3
    v8 = -2.25
    fmt.Println(v3, v8)
    v8 = 1e10
    if true {
        var v9 int8 = int8(v3)
        _ = v9
    }
    var v10 bool = true
    var v11 uint8 = 128
    var v12 bool = v10
    v13 := GenS2{F0: -2147483648, F1: 127, F2: 1}
    v1 = ((v1 + v8) + v8)
    v1 -= v8
    var v14 uint16 = (v4 | (v4 ^ (v4 << ((65535) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
