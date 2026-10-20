package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 uint) int8 {
    p0 += (^(p0))
    if true {
        for fi5 := 0; fi5 < 18; fi5++ {
            for fi6 := 0; fi6 < 13; fi6++ {
                p0 += p0
                var v7 string = "\n\t"
                var v8 int16 = 1000
                _ = fi6
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 686, F1: -2147483648, F2: 522}
            _ = fi5
            _ = v9
        }
        v10 := []int{84, 4294967295, 9223372036854775807}
        if false {
            var v11 float32 = (-2.25 - 0.1)
            fmt.Println(v11, v11)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: 42, F1: -992, F2: 1000}
                _ = fi12
                _ = v13
            }
        }
        _ = v10
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            p0 -= 0
            v15 := GenS2{F0: 35, F1: fi14, F2: 416}
            p0 = (^(p0))
            _ = fi14
            _ = v15
        }
        fmt.Println(p0)
        if (32767 != 0) {
            p0 += p0
            if (!(0 <= 42)) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 int16 = 868
                    _ = fi16
                    _ = v17
                }
                var v18 uint = 4294967295
                _ = v18
            } else {
                var v19 uint = (p0 + p0)
                _ = v19
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 bool = true
                v22 := []int{fi20, fi20}
                v23 := GenS1{F0: 9223372036854775807, F1: 0, F2: -32768}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        } else {
            var v24 float64 = (1e10 + (0.1 - -2.25))
            v25 := []int{}
            _ = v24
            _ = v25
        }
    }
    var v26 uint = p0
    p0 -= v26
    _ = p0
    _ = v26
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 255
    var v3 string = "a"
    v4 := []uint32{2147483647, 2147483648}
    var v27 uint16 = 0
    v28 := GenS2{F0: -1, F1: 0, F2: 2147483647}
    fmt.Println(v2, v2, v27)
    if true {
        for fi29 := 0; fi29 < 10; fi29++ {
            v1[99] = 762
            v27 = v27
            _ = fi29
        }
    }
    v2 += (v2 * v2)
    v1[1] = uint32(v2)
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 string = (v3 + ("" + v3))
        _ = fi30
        _ = v31
    }
    if true {
        var v32 int16 = 0
        var v33 uint16 = (v27 << ((v2) & 63))
        _ = v32
        _ = v33
    }
    v2 = (v2 / ((1) | 1))
    v27 += 255
    var v34 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
}
