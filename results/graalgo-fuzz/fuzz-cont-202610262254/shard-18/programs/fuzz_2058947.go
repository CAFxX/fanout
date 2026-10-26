package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := []int{fi6, (fi6 + fi6), (fi6 * 2147483647), -2147483648}
            v8 := GenS1{F0: -1, F1: 9223372036854775807, F2: -119}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint8 = 199
        _ = fi5
        _ = v9
    }
    var v10 uint = 2
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: -270, F1: -9223372036854775808, F2: -1}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := []int{72, 0, 127}
    v10 -= 280
    _ = v10
    _ = v14
    return 127
}
func genFn1(p0 uint16, p1 int64) uint32 {
    v15 := []int{2147483647, 316, 0, 631}
    var v16 uint8 = 128
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]int)
    var v3 string = ""
    v4 := make(map[int]uint32)
    var v17 uint32 = 381
    if (false && (genFn1(889, 9223372036854775807) <= genFn1(1, 576))) {
        v1 = 42
    }
    v18 := GenS1{F0: -1, F1: 9223372036854775807, F2: 1}
    fmt.Println(v3, v17, v17)
    _ = len(v4)
    v19 := GenS1{F0: -2147483648, F1: 42, F2: -255}
    v20 := GenS1{F0: 473, F1: int64(v17), F2: 0}
    delete(v4, 42)
    if false {
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := []int{}
            v1 = 100
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(v3, v17)
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 string = ((v3 + v3) + ((v3 + v3) + v3))
        v26 := GenS2{F0: 660, F1: -32768}
        var v27 int8 = (v1 / ((75) | 1))
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := []int{}
    v17 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
}
