package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32; F1 int64; F2 bool }
func genFn0() uint64 {
    if (18446744073709551615 >= 18446744073709551615) {
        for fi4 := 0; fi4 < 15; fi4++ {
            v5 := []int{}
            var v6 int32 = -52
            v7 := GenS1{F0: -2147483648, F1: 42}
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
        v8 := GenS2{F0: -895, F1: 1}
        v9 := []int{42, 4294967295, 455}
        _ = v8
        _ = v9
    }
    v10 := []int{-9223372036854775808, 866, -2147483648, 42}
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 float32 = 0.1
        _ = fi11
        _ = v12
    }
    _ = len(v10)
    _ = v10
    return 4294967295
}
func genFn1(p0 uint64) bool {
    v13 := GenS2{F0: 283, F1: 65535}
    if (!(!true)) {
        v14 := []int{-1, 32768, 9223372036854775807}
        _ = v14
    } else {
        p0 += (p0 / (((p0 * p0)) | 1))
        v15 := GenS2{F0: -32768, F1: 332}
        _ = v15
    }
    fmt.Println(p0)
    var v16 uint32 = 196
    _ = p0
    _ = v13
    _ = v16
    return ((((1e10 + 1.5) - -2.25) / (1e10 * 1.5)) > (1.5 * 1.5))
}

func main() {
    v1 := []string{}
    v2 := []int{2147483647, 32767}
    var v3 uint64 = 1
    fmt.Println(v3)
    v3 += (v3 - (v3 & v3))
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 float64 = 1e10
        var v19 uint64 = 18446744073709551615
        _ = fi17
        _ = v18
        _ = v19
    }
    v3 = v3
    v20 := GenS3{F0: -956, F1: -9223372036854775808, F2: ((1.5 > float32(float64(0.1))) && false)}
    v3 = (v3 + v3)
    v2 = append(v2, -9223372036854775808)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
