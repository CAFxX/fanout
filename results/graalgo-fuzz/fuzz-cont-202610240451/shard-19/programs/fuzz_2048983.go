package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint8) int64 {
    v7 := GenS1{F0: -1, F1: 591, F2: -434}
    p0 = (p0 ^ (p0 * (p0 << ((0) & 63))))
    v8 := GenS2{F0: 42, F1: "\n\t"}
    v9 := []int{42, 127, -9223372036854775808, 966}
    var v10 uint8 = (p0 / ((p0) | 1))
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -93
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 1
    v3 := []int{128, 127, -1, 32767}
    var v4 uint = 4294967295
    v5 := []string{("a" + (("\n\t" + (("a" + "") + "")) + "hello"))}
    v6 := make(map[string]bool)
    if (!((1000 < -1) || true)) {
        v11 := GenS3{F0: 2147483647, F1: 998, F2: (^(v4))}
        v6[("foo bar" + "hello")] = true
        v4 = v4
        _ = v11
    } else {
        var v12 int8 = 100
        _ = v12
    }
    v4 += v4
    v13 := GenS4{F0: 4294967295}
    fmt.Println(v2)
    fmt.Println(v2, v4)
    fmt.Println(v4)
    var v14 uint = (^(v4))
    var v15 uint16 = 313
    v14 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
