package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8) uint16 {
    p0 = p0
    if true {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int8 = (p0 | p0)
            _ = fi6
            _ = v7
        }
        var v8 int32 = 0
        _ = v8
    } else {
        var v9 int32 = 42
        p0 = (^(p0))
        _ = v9
    }
    switch -246 {
    case 32768:
        p0 -= p0
    case 128:
        var v10 uint64 = 0
        _ = v10
    case 32767:
        fmt.Println(p0)
    }
    v11 := []int{-596}
    _ = p0
    _ = v11
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 float64 = 0.0
    var v3 uint8 = 99
    var v4 int64 = 1
    var v5 int = 2147483647
    if (128 == (v3 * v3)) {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint16 = 867
                v15 := GenS1{F0: 127, F1: 0, F2: v5}
                var v16 float32 = ((1.5 * 1.5) * ((0.0 - 1.5) + (1e10 * -2.25)))
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := []int{}
                v19 := GenS1{F0: 100, F1: (v5 | (v5 + fi12)), F2: -489}
                if false {
                    v20 := GenS2{F0: 0, F1: 42}
                    _ = len(v18)
                    _ = v20
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            v1 = v1
            _ = fi12
        }
        fmt.Println(v4)
        var v21 int16 = -1
        _ = v21
    }
    v2 = (0.0 + v2)
    v22 := GenS2{F0: 654, F1: (v5 * 32767)}
    v23 := GenS3{F0: (v4 ^ (v4 + v4))}
    v24 := []int{int(v3)}
    _ = len(v24)
    fmt.Println(v3)
    v3 = v3
    if ((0 > 127) && (!((!true) && (!((false && false) || ((v2 * (((((0.0 * v2) * (v2 - 1.5)) / (float64(((((3.14159 + -2.25) - float32(3.14159)) / 1.0) + -2.25)) - (float64(3.14159) * (float64((float32(v2) / (3.14159 + float32(float64(1e10))))) + 0.0)))) - v2) + ((-2.25 * (v2 * 3.14159)) * 0.1))) != v2)))))) {
        fmt.Println(v2)
    } else {
        v2 += 0.1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
