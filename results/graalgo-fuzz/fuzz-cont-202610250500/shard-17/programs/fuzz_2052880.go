package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 1000
    var v3 uint = 42
    if false {
        for fi4 := 0; fi4 < 5; fi4++ {
            fmt.Println(v2, fi4)
            var v5 bool = false
            var v6 uint64 = 0
            _ = fi4
            _ = v5
            _ = v6
        }
        v3 -= (v3 & (v3 ^ 42))
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v3 = 627
            _ = fi7
        }
    }
    var v8 uint8 = 128
    v9 := []int{2147483647, -805}
    v3 -= v3
    v2 += (v2 % ((1000) | 1))
    if true {
        v2 = v2
        switch "a" {
        case "":
            var v10 float64 = (-2.25 - float64(-2.25))
            _ = v10
        case "foo bar":
            v11 := GenS1{F0: -1, F1: (v8 < (v8 + 42)), F2: v3}
            _ = v11
        default:
            v12 := GenS1{F0: 862, F1: false, F2: (v3 * 4294967295)}
            _ = v12
        }
        v13 := GenS1{F0: 1, F1: ((v8 ^ (v8 | v8)) <= v8), F2: v3}
        _ = v13
    }
    var v14 bool = true
    if (!(!false)) {
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS1{F0: -1, F1: (!(42 >= -394)), F2: v3}
                v3 = v3
                _ = fi16
                _ = v17
            }
            v18 := v1[42]
            v19 := GenS1{F0: -9223372036854775808, F1: v14, F2: 18446744073709551615}
            _ = fi15
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 int = fi20
            if (120 >= 42) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    if v14 {
                        fmt.Println(fi20)
                        var v23 float32 = 1.5
                        _ = v23
                    } else {
                        var v24 float32 = 0.0
                        fmt.Println(fi20)
                        var v25 int8 = 42
                        _ = v24
                        _ = v25
                    }
                    switch fi22 {
                    case -730:
                        v26 := GenS1{F0: -475, F1: v14, F2: (^(v3))}
                        _ = v26
                    case 1:
                        v21 -= fi20
                    case 39:
                        v3 = v3
                    }
                    if v14 {
                        delete(v1, 684)
                    }
                    _ = fi22
                }
            } else {
                fmt.Println(v21)
            }
            fmt.Println(fi20)
            _ = fi20
            _ = v21
        }
        var v27 uint8 = (v8 ^ v8)
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
