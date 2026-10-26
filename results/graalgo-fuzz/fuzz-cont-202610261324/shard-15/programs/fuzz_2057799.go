package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 int32, p2 uint16) uint {
    var v4 uint8 = (p0 ^ p0)
    v5 := []int{4294967295}
    if true {
        var v6 string = "foo bar"
        p0 -= p0
        var v7 int32 = p1
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 uint64) uint8 {
    p1 = (^(p1))
    if true {
        v8 := GenS1{F0: 42, F1: true, F2: 0}
        var v9 uint64 = 1
        p1 -= 564
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = 505
    var v3 uint = 255
    v10 := GenS1{F0: -2147483648, F1: (!(128 != 85)), F2: (v2 - -692)}
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v2, fi11)
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS2{F0: (v3 * genFn0(2, 212, 32767)), F1: 1, F2: -778}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v1 += v1
    if true {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint8 = genFn1(840, 42)
            v1 = 1e10
            var v16 int8 = -1
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int8 = 1
        v3 += (v3 % ((0) | 1))
        _ = v17
    }
    switch 128 {
    case 4294967295:
        if false {
            var v18 int64 = 9223372036854775807
            _ = v18
        }
    case 42:
        v3 -= genFn0(42, 1, 703)
    }
    v19 := []int{0, 971}
    v2 += v2
    v2 -= (v2 % ((v2) | 1))
    _ = len(v19)
    v3 = v3
    v20 := GenS1{F0: 0, F1: true, F2: v2}
    v19 = append(v19, -9223372036854775808)
    v2 = 32767
    v1 -= v1
    if len(v19) > 0 {
        v19[0] = 0
    }
    v21 := GenS3{F0: 255, F1: 32767}
    fmt.Println(v3, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
