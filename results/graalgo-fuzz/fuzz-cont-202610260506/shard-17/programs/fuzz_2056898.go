package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 int) int8 {
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{p0}
            p0 = (p0 | p0)
            _ = fi7
            _ = v8
        }
        var v9 float32 = (0.0 * ((-2.25 - -2.25) / (float32(0.1) / float32(float64(1e10)))))
        _ = v9
    }
    fmt.Println(p0)
    v10 := GenS1{F0: 917}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            if false {
                p0 -= p0
            }
            v13 := []int{p0, 9223372036854775807}
            if (-447 < 0) {
                v13 = append(v13, (fi12 + (fi12 & p0)))
            } else {
                var v14 int = (fi11 >> ((2) & 63))
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v15 bool = false
    _ = p0
    _ = v10
    _ = v15
    return 1
}
func genFn1(p0 uint64, p1 uint, p2 uint16) int64 {
    v16 := GenS2{F0: 859, F1: -128, F2: 0}
    if (uint16(p1) > p2) {
        p2 -= p2
    } else {
        fmt.Println(p0)
    }
    v17 := []int{128, 788, 42, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = float32(3.14159)
    v3 := make(map[int]bool)
    var v4 int16 = 42
    v5 := make(map[string]int)
    v6 := make(map[uint32]bool)
    v2 -= 1.5
    v4 = 1
    var v18 uint = 42
    var v19 int32 = 1
    var v20 bool = false
    v18 = (v18 / ((2) | 1))
    if true {
        delete(v3, 32768)
        v21 := GenS3{F0: 2, F1: "a"}
        var v22 string = "a"
        _ = v21
        _ = v22
    } else {
        var v23 int32 = -2147483648
        fmt.Println(v2, v2)
        v24 := GenS2{F0: 0, F1: 123, F2: v4}
        _ = v23
        _ = v24
    }
    var v25 uint = v18
    var v26 bool = false
    v26 = (false || (v20 || (-9223372036854775808 > 9223372036854775807)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
