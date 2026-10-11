package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8) int32 {
    v6 := GenS1{F0: 1000}
    for fi7 := 0; fi7 < 3; fi7++ {
        if ((-1 < 1000) && (1 != 2)) {
            p0 -= (p0 + p0)
            var v8 uint64 = 299
            v8 -= v8
            _ = v8
        } else {
            p0 = 100
            var v9 uint64 = 2
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{9223372036854775807}
    _ = p0
    _ = v10
    _ = v6
    return 745
}
func genFn1(p0 uint) int64 {
    fmt.Println(p0)
    v11 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 42
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    v4 := make(map[int]string)
    var v5 int64 = 1
    if (-2147483648 < genFn0(70)) {
        fmt.Println(v5)
        for fi12 := 0; fi12 < 19; fi12++ {
            v5 = (v5 & genFn1(4294967295))
            var v13 uint16 = 582
            fmt.Println(v5, fi12, v1)
            _ = fi12
            _ = v13
        }
        if true {
            v5 -= v5
            switch 1 {
            case -444:
                v5 -= v5
            case 42:
                v14 := GenS1{F0: 32767}
                _ = v14
            case -1:
                fmt.Println(v1, v5)
            }
            v5 -= genFn1(903)
        }
    }
    switch -2147483648 {
    case 2147483647:
        var v15 int = 42
        _ = v15
    default:
        v5 = (v5 % (((v5 ^ (v5 << ((1) & 63)))) | 1))
    }
    var v16 int = -464
    var v17 uint64 = 321
    fmt.Println(v5)
    var v18 float32 = 1e10
    var v19 int8 = 42
    fmt.Println(v1)
    var v20 int64 = v5
    v21 := GenS2{F0: 1}
    _ = len(v4)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
