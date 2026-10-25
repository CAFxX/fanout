package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0() int16 {
    var v5 uint8 = 250
    var v6 uint = 2
    v7 := GenS1{F0: 1, F1: (v6 - uint(v5)), F2: (v6 & 0)}
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 2147483648
    var v3 float32 = (1e10 + 0.1)
    var v4 int32 = 42
    v8 := []int{32768, 127, 0, 826}
    v3 -= v3
    var v9 int32 = (v4 << ((323) & 63))
    v4 = (v4 / ((v4) | 1))
    v8 = append(v8, -908)
    if (1000 <= 1000) {
        v9 = (v4 / (((v4 ^ v4)) | 1))
        v1 = v1
    }
    v3 = (float32(3.14159) * float32(-2.25))
    if len(v8) > 0 {
        v8[0] = 32767
    }
    v10 := GenS2{F0: (^(v2)), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
