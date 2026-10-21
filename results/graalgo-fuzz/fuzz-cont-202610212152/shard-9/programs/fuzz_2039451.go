package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
func genFn0() int64 {
    var v5 int64 = -1
    v6 := []int{-1}
    for fi7 := 0; fi7 < 6; fi7++ {
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
        _ = fi7
    }
    if true {
        var v8 float64 = ((3.14159 - 1.5) * (0.1 - 1.5))
        v9 := GenS1{F0: 128, F1: 833}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 140, F1: 65535}
    _ = v10
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 bool) uint64 {
    if ((0 < 1) && p0) {
        fmt.Println(p0)
        p0 = p0
    } else {
        v11 := GenS1{F0: 128, F1: 2}
        _ = v11
    }
    v12 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 2
    var v3 uint32 = 1
    var v4 uint8 = 127
    fmt.Println(v1)
    if ((v4 >> ((18446744073709551615) & 63)) == (v4 - v4)) {
        v13 := GenS2{F0: -9223372036854775808}
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int16 = (v1 & v1)
            _ = fi14
            _ = v15
        }
        switch 0 {
        case 0:
            v16 := GenS3{F0: 42}
            _ = v16
        default:
            if true {
                v3 = v3
                v17 := []int{}
                _ = v17
            }
        }
        _ = v13
    }
    v4 = v4
    v18 := []int{}
    _ = len(v18)
    var v19 int8 = 1
    v4 -= 2
    var v20 uint32 = v3
    if true {
        var v21 string = "\n\t"
        _ = len(v18)
        v22 := GenS3{F0: 33}
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
