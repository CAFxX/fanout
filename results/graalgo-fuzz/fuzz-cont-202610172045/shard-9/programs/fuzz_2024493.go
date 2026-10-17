package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
func genFn0(p0 int8) int64 {
    v4 := GenS1{F0: 94}
    fmt.Println(p0)
    fmt.Println(p0)
    p0 -= (p0 / (((p0 ^ p0)) | 1))
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 int16, p1 int, p2 int16) int8 {
    if false {
        p1 -= 32768
        var v5 int32 = -2147483648
        _ = v5
    }
    var v6 int64 = genFn0(1)
    if (42 > 542) {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint = 42
            if true {
                p0 = (p0 | p0)
                p2 -= 296
                v9 := GenS1{F0: 92}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        v10 := GenS1{F0: 1}
        v11 := []int{}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return int8(p1)
}

func main() {
    v1 := []uint32{2147483647, 480, 587, 0}
    var v2 uint = 18446744073709551615
    var v3 int8 = 0
    var v12 uint8 = 127
    v1 = append(v1, 757)
    v2 = (v2 - v2)
    fmt.Println(v3, v3, v3)
    v12 += v12
    v2 = v2
    var v13 int16 = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v14 := GenS1{F0: 18446744073709551615}
    v15 := GenS2{F0: int32(v13), F1: 642, F2: 0}
    v13 += -593
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
