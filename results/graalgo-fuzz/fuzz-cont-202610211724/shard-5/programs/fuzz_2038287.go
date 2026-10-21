package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 int64; F2 int64 }
func genFn0(p0 uint8, p1 int32, p2 int8) uint {
    fmt.Println(p0)
    v6 := GenS1{F0: (p0 % ((p0) | 1)), F1: 1, F2: 18446744073709551615}
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 uint = 383
    var v2 bool = (!true)
    var v3 uint = 2
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    var v7 float64 = 1.5
    if (v3 < v3) {
        v8 := GenS1{F0: 24, F1: 42, F2: 18446744073709551615}
        _ = v8
    }
    var v9 float32 = -2.25
    var v10 float32 = v9
    v11 := v5[("hello" + ("" + "foo bar"))]
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint32 = 4294967295
        v13 = (v13 | 2147483648)
        for fi14 := 0; fi14 < 9; fi14++ {
            v3 -= (v1 | v3)
            if v2 {
                v15 := GenS2{F0: 0, F1: (v11 + (v11 + "a")), F2: 4294967295}
                _ = v15
            } else {
                if (-9223372036854775808 != 42) {
                    v9 = v9
                    var v16 bool = ((^(v1)) > 724)
                    v17 := []int{(fi14 << ((18446744073709551615) & 63)), (-(fi12)), fi12}
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS3{F0: 0, F1: 1, F2: fi12}
                    v7 = ((1.5 + v7) * v7)
                    _ = v18
                }
                v19 := GenS3{F0: 787, F1: 1, F2: (fi12 / (((fi12 + (fi12 + (fi14 ^ (fi14 >> (((^(v1))) & 63)))))) | 1))}
                _ = v19
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    if v2 {
        v20 := GenS1{F0: 42, F1: 0, F2: 1}
        v21 := GenS3{F0: 42, F1: 18446744073709551615, F2: 2147483647}
        v22 := GenS4{F0: -224, F1: -1, F2: 1}
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
