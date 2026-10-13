package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 int, p1 string, p2 uint64) bool {
    var v6 int16 = 0
    fmt.Println(v6, v6, p1)
    var v7 uint16 = 2
    v8 := []int{(p0 % ((p0) | 1)), (p0 % ((p0) | 1)), p0}
    v8 = append(v8, -690)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 65535
    v3 := make(map[int]int)
    var v4 int64 = 42
    v5 := make(map[int]int)
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := []int{fi9}
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(v4)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v12 int = (fi9 - (fi11 ^ (fi11 / (((fi9 >> ((1) & 63))) | 1))))
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if false {
        if false {
            delete(v3, -9223372036854775808)
        } else {
            v13 := GenS1{F0: ((!(!false)) && (true && (float32((0.0 * ((float64(((1e10 * 3.14159) - 0.0)) / 1.0) + (1.5 / -2.25)))) != ((0.1 - ((1e10 + 0.0) - 3.14159)) - (-2.25 * 0.0))))), F1: 674}
            v14 := GenS1{F0: ((-2.25 * 1.5) == 0.1), F1: 531}
            v15 := GenS1{F0: true, F1: -1}
            _ = v13
            _ = v14
            _ = v15
        }
    } else {
        var v16 float32 = 0.1
        for fi17 := 0; fi17 < 3; fi17++ {
            v2 -= (v2 + v2)
            v4 = (v4 * 42)
            v4 = v4
            _ = fi17
        }
        _ = v16
    }
    fmt.Println(v1, v1)
    fmt.Println(v2)
    var v18 uint32 = v2
    var v19 float32 = -2.25
    var v20 uint = 422
    v21 := []int{-703, 452, 9223372036854775807}
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            if ((true && (!false)) && ((-2.25 - ((0.0 - v19) / -2.25)) <= v19)) {
                v18 += (^(v18))
            } else {
                v2 += v18
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v21 = append(v21, (fi24 + (fi24 << ((v20) & 63))))
        _ = fi24
    }
    var v25 float64 = float64(v19)
    if false {
        var v26 uint32 = v18
        var v27 bool = true
        _ = v26
        _ = v27
    } else {
        v25 += v25
    }
    v28 := GenS1{F0: (25 >= 127), F1: 2147483647}
    v18 -= (v18 % ((uint32(v20)) | 1))
    v2 = (^(v18))
    delete(v5, -535)
    v29 := GenS1{F0: false, F1: 0}
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int32 = -1
        for fi32 := 0; fi32 < 14; fi32++ {
            if (!true) {
                _ = len(v5)
                v20 += v20
                v33 := GenS1{F0: (255 <= 207), F1: -1}
                _ = v33
            }
            _ = fi32
        }
        _ = len(v21)
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
