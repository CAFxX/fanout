package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
func genFn0(p0 string) int32 {
    var v6 uint32 = 2
    p0 = (p0 + p0)
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 uint64, p1 int32, p2 int32) int16 {
    var v7 float64 = (1.5 * 0.0)
    var v8 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 896
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint = 65535
        v10 += 417
        if true {
            if (!(!(true && false))) {
                v2 = v2
                if (1 < 32767) {
                    if (389 == (v10 & v10)) {
                        v11 := []int{(-(fi9)), fi9, -9223372036854775808, fi9}
                        v10 = 65535
                        _ = v11
                    } else {
                        var v12 int16 = 1000
                        _ = v12
                    }
                    fmt.Println(v10, fi9)
                } else {
                    var v13 uint8 = 0
                    _ = v13
                }
            }
            v14 := []int{(fi9 >> ((57) & 63)), fi9}
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    v2 = (v2 - (v2 * (v2 - v2)))
    v4[""] = true
    v2 = (v2 % ((v2) | 1))
    v15 := GenS1{F0: 128, F1: 0, F2: -239}
    v16 := GenS1{F0: 2147483647, F1: 566, F2: 42}
    var v17 int64 = -328
    switch 127 {
    case 919:
        v18 := GenS1{F0: 0, F1: 1, F2: -390}
        _ = v18
    case 32767:
        v2 -= v2
    }
    v17 = (v17 | 165)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
