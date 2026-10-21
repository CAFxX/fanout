package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 int64; F2 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 uint32) uint32 {
    var v5 float32 = 1e10
    if (int32(p1) != -2147483648) {
        var v6 int32 = 1
        p1 = (^(p1))
        _ = v6
    } else {
        fmt.Println(p1, v5, p0)
        if false {
            if true {
                v7 := GenS1{F0: (p1 << ((42) & 63)), F1: 223, F2: 42}
                v8 := []int{}
                fmt.Println(v5, p1)
                _ = v7
                _ = v8
            } else {
                p0 += p0
                v9 := []int{32768, 893}
                for fi10 := 0; fi10 < 7; fi10++ {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        p1 -= (p1 / ((2) | 1))
                        if len(v9) > 0 {
                            v9[0] = fi11
                        }
                        var v12 uint32 = p1
                        _ = fi11
                        _ = v12
                    }
                    if (42 > 323) {
                        p0 = (p0 ^ -628)
                    }
                    _ = fi10
                }
                _ = v9
            }
            var v13 int = -9223372036854775808
            _ = v13
        } else {
            v14 := GenS1{F0: 586, F1: 2, F2: 2}
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS1{F0: (p1 ^ (^(p1))), F1: 812, F2: 32767}
                _ = fi15
                _ = v16
            }
            fmt.Println(v5, p1, p1)
            _ = v14
        }
    }
    var v17 bool = false
    if ((p1 / ((p1) | 1)) < p1) {
        for fi18 := 0; fi18 < 17; fi18++ {
            v17 = v17
            v19 := GenS2{F0: 42, F1: 32767, F2: 0}
            if v17 {
                for fi20 := 0; fi20 < 1; fi20++ {
                    p1 += 769
                    _ = fi20
                }
                p1 = (p1 << ((42) & 63))
            }
            _ = fi18
            _ = v19
        }
        var v21 uint8 = 115
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v5
    return (p1 % (((p1 << ((2) & 63))) | 1))
}
func genFn1(p0 int8, p1 uint8, p2 bool) int {
    var v22 float64 = 0.1
    fmt.Println(v22, v22, v22)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 417
}

func main() {
    var v1 uint = 0
    var v2 int64 = 1
    var v3 float32 = float32(float64(1.5))
    v4 := make(map[int]string)
    switch 630 {
    case -1:
        if true {
            var v23 int8 = 46
            var v24 int32 = 1
            var v25 int = genFn1(100, 128, false)
            _ = v23
            _ = v24
            _ = v25
        }
    }
    _ = len(v4)
    v26 := GenS1{F0: 65535, F1: (v1 % ((v1) | 1)), F2: 0}
    var v27 uint = (v1 << ((4294967295) & 63))
    v28 := GenS2{F0: 416, F1: 649, F2: 320}
    var v29 uint32 = 42
    var v30 int16 = -1
    if ((-2.25 / v3) < v3) {
        if true {
            v29 += (v29 / ((v29) | 1))
            fmt.Println(v30)
        }
        var v31 uint64 = 2
        var v32 int64 = (-(v2))
        _ = v31
        _ = v32
    }
    if true {
        if (false && true) {
            var v33 float32 = v3
            var v34 bool = false
            var v35 bool = (-32768 > 42)
            _ = v33
            _ = v34
            _ = v35
        }
    }
    var v36 float64 = (1e10 * 0.1)
    fmt.Println(v1)
    v37 := GenS3{F0: 4294967295, F1: 760, F2: ((v30 ^ (-(v30))) <= v30)}
    v38 := GenS4{F0: (v27 >> ((v27) & 63))}
    v39 := GenS1{F0: 2147483647, F1: 967, F2: 850}
    v40 := []int{363}
    var v41 int16 = v30
    if false {
        var v42 uint = 65
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
