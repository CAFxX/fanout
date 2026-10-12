package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0() int16 {
    var v5 int32 = 42
    v6 := GenS1{F0: 42, F1: 286}
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 2, F1: 368}
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return -703
}

func main() {
    var v1 uint32 = 42
    v2 := []int{4294967295, 32768, 42, -157}
    var v3 uint = 0
    var v4 int = -294
    v4 += v4
    fmt.Println(v3, v1, v3)
    if (v3 > (v3 * (v3 >> ((v3) & 63)))) {
        for fi9 := 0; fi9 < 16; fi9++ {
            v2 = append(v2, 2147483647)
            fmt.Println(v1, fi9, fi9)
            _ = fi9
        }
    }
    var v10 int32 = 296
    v3 += v3
    v11 := GenS2{F0: 812, F1: 161}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint = v3
        fmt.Println(v10, v13)
        v13 = 65535
        _ = fi12
        _ = v13
    }
    switch (v4 % ((-9223372036854775808) | 1)) {
    case 32767:
        var v14 uint8 = 2
        _ = v14
    }
    v15 := GenS3{F0: (v1 << (((v3 & (v3 + v3))) & 63)), F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
