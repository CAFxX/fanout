package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 bool }
func genFn0(p0 int16, p1 int16, p2 int16) bool {
    p1 = p1
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: (p1 << ((345) & 63)), F1: (((^(p0)) >= (p0 << ((1) & 63))) || ("a" != "")), F2: 0}
            v9 := GenS1{F0: (p2 + (p2 ^ (p0 | -882))), F1: (449 >= -1), F2: 255}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{(^(fi6)), (fi6 / ((fi6) | 1))}
        _ = fi6
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return true
}
func genFn1(p0 bool) int64 {
    if (65535 != 42) {
        fmt.Println(p0)
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            v13 := []int{}
            var v14 int8 = 90
            v15 := GenS1{F0: -1, F1: p0, F2: 2}
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 uint16 = 2
        v16 = v16
        _ = fi12
        _ = v16
    }
    _ = p0
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"hello", "\n\t", "", ""}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    var v17 float32 = 3.14159
    v17 += -2.25
    var v18 uint32 = 65535
    v19 := GenS2{F0: false, F1: 984, F2: v18}
    v18 = 4294967295
    if true {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v17, fi20)
            var v21 float64 = 0.0
            _ = fi20
            _ = v21
        }
    } else {
        fmt.Println(v17, v17)
    }
    var v22 int8 = -80
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS3{F0: 983, F1: -9223372036854775808, F2: (false && false)}
        if (42 == 2) {
            if ((v17 != v17) && (!(2 != 18446744073709551615))) {
                var v25 int8 = (v22 & v22)
                var v26 int64 = -9223372036854775808
                if ((!(!((^(fi23)) > fi23))) && (fi23 != fi23)) {
                    var v27 int8 = -87
                    _ = v27
                }
                _ = v25
                _ = v26
            }
            _ = len(v2)
        }
        _ = fi23
        _ = v24
    }
    var v28 bool = true
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            var v31 uint16 = 1
            if ((v18 | v18) <= 2147483648) {
                fmt.Println(v31, v18, v28)
            }
            _ = fi30
            _ = v31
        }
        v28 = v28
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
