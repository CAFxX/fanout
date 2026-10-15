package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }
func genFn0(p0 int16, p1 uint8) uint32 {
    p1 += (p1 - (p1 + (p1 & p1)))
    p0 += p0
    v7 := GenS1{F0: false, F1: 645}
    v8 := GenS1{F0: ((!(2147483647 == -2147483648)) || (false || false)), F1: 255}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int8) bool {
    p0 += (p0 % ((p0) | 1))
    var v9 uint8 = 42
    var v10 int = 9223372036854775807
    _ = p0
    _ = v10
    _ = v9
    return (!true)
}

func main() {
    var v1 int32 = 42
    var v2 int32 = v1
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v6 float32 = 0.1
    var v11 bool = true
    v1 = v2
    v12 := []int{0, 1, -1}
    if v11 {
        switch 9223372036854775807 {
        case 32767:
            var v13 int16 = 32767
            _ = v13
        case 42:
            v1 = v1
        case 294:
            v2 += v1
        }
    } else {
        v14 := GenS1{F0: false, F1: 4294967295}
        var v15 float32 = ((v6 * (float32((((3.14159 + (3.14159 / -2.25)) / 3.14159) / float64(v6))) * 1e10)) * v6)
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int64 = -1
        _ = fi16
        _ = v17
    }
    v11 = (float64(v6) < float64(v6))
    switch 2147483647 {
    case -2147483648:
        var v18 bool = (1000 < 32767)
        _ = v18
    case -9223372036854775808:
        switch -1 {
        case -810:
            if len(v12) > 0 {
                v12[0] = 1
            }
        case 128:
            _ = len(v12)
        }
    }
    v6 += 0.0
    fmt.Println(v1, v2, v6)
    var v19 uint64 = 947
    if ((!(127 < 127)) && (((!(1 <= -1)) && v11) || v11)) {
        fmt.Println(v2)
    } else {
        v20 := GenS2{F0: (v2 - v1), F1: "\n\t", F2: 0}
        v3[2147483648] = 127
        _ = v20
    }
    if len(v12) > 0 {
        v12[0] = 2147483647
    }
    var v21 uint = 42
    fmt.Println(v19, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
