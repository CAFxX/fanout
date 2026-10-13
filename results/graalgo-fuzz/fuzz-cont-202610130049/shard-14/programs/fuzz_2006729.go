package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 int32) int {
    var v6 uint32 = 0
    p0 -= 2147483647
    if false {
        switch 631 {
        case 0:
            for fi7 := 0; fi7 < 11; fi7++ {
                for fi8 := 0; fi8 < 7; fi8++ {
                    if false {
                        v9 := GenS1{F0: fi7, F1: (fi8 % (((fi8 ^ (^(fi7)))) | 1))}
                        var v10 uint8 = 0
                        _ = v10
                        _ = v9
                    } else {
                        var v11 int16 = 92
                        fmt.Println(fi8, v11, v6)
                        _ = v11
                    }
                    var v12 int16 = 0
                    var v13 uint64 = 2
                    _ = fi8
                    _ = v12
                    _ = v13
                }
                var v14 int16 = 750
                v15 := []int{(fi7 & int(v14)), (fi7 % ((128) | 1)), -1, 398}
                _ = fi7
                _ = v14
                _ = v15
            }
        case 2147483647:
            v16 := []int{-327, 128, -2147483648}
            _ = v16
        }
        v6 -= 1
        var v17 float64 = (float64(0.0) - 1e10)
        _ = v17
    }
    switch -722 {
    case 42:
        v18 := GenS2{F0: true, F1: false}
        _ = v18
    }
    _ = p0
    _ = v6
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 float64 = 3.14159
    v4 := []int{1, -534, 795}
    var v5 float32 = (3.14159 * 0.1)
    for fi19 := 0; fi19 < 7; fi19++ {
        v3 -= (3.14159 - v3)
        _ = fi19
    }
    v5 += v5
    fmt.Println(v3, v5)
    fmt.Println(v3)
    v20 := GenS3{F0: -128, F1: 2147483647}
    v1[42] = genFn0(1)
    fmt.Println(v3)
    fmt.Println(v3, v3)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
