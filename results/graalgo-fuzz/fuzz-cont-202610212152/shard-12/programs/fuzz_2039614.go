package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
func genFn0(p0 uint64, p1 uint8) bool {
    p1 += (p1 % ((255) | 1))
    var v4 int32 = 612
    v5 := []int{-267}
    p0 -= (p0 | 42)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (-1 == 42)
}
func genFn1(p0 int, p1 uint8) string {
    var v6 uint8 = (p1 >> ((2) & 63))
    p0 += (p0 & (p0 ^ p0))
    v6 -= 0
    _ = p0
    _ = p1
    _ = v6
    return "a"
}

func main() {
    var v1 uint16 = 778
    var v2 int16 = -81
    var v3 float64 = float64(1e10)
    switch -1 {
    case -451:
        v7 := GenS1{F0: 42}
        _ = v7
    case 127:
        v8 := []int{-17, 924, 4294967295}
        _ = v8
    }
    var v9 uint = 0
    if false {
        var v10 uint = v9
        v11 := GenS2{F0: 0}
        for fi12 := 0; fi12 < 3; fi12++ {
            if false {
                fmt.Println(fi12, fi12)
                v2 -= (v2 - -1)
                v9 = (v10 | v10)
            }
            v13 := []int{(-(fi12)), (fi12 % (((-(fi12))) | 1)), (-(fi12)), (fi12 >> ((0) & 63))}
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    } else {
        if false {
            var v14 uint32 = 2147483647
            var v15 int = 128
            v16 := []int{v15}
            _ = v14
            _ = v15
            _ = v16
        } else {
            var v17 int64 = 1
            v18 := GenS2{F0: 127}
            _ = v17
            _ = v18
        }
        v19 := GenS3{F0: 42, F1: 18446744073709551615, F2: 42}
        v20 := []int{}
        _ = v19
        _ = v20
    }
    v2 = v2
    v21 := GenS2{F0: 173}
    if false {
        if false {
            if true {
                var v22 uint64 = 1
                _ = v22
            }
            var v23 int = int(v9)
            _ = v23
        } else {
            v9 = (v9 ^ 0)
            v9 = (v9 ^ 820)
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 int16 = v2
        var v26 int64 = -771
        v9 = (v9 ^ 18446744073709551615)
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int = 1
    v28 := []int{-833, v27, (v27 & (v27 | v27)), 42}
    v1 = (v1 + (v1 ^ (v1 + v1)))
    v27 -= -1
    _ = len(v28)
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
