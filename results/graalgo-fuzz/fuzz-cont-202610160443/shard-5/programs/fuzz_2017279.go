package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 int, p2 int) uint8 {
    var v7 int64 = p0
    for fi8 := 0; fi8 < 17; fi8++ {
        if false {
            if (!true) {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v10 := []int{(p2 * (fi9 | (fi8 % ((-9223372036854775808) | 1)))), (fi9 - (^(p1))), (p2 % ((fi9) | 1)), fi9}
                    v11 := GenS1{F0: 9223372036854775807, F1: 0, F2: fi9}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            var v12 uint16 = 42
            fmt.Println(p1, p1)
            _ = v12
        }
        _ = fi8
    }
    fmt.Println(p2, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 108
}
func genFn1() int32 {
    var v13 int64 = -9223372036854775808
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint64 = 0
        v15 = v15
        _ = fi14
        _ = v15
    }
    v13 = v13
    var v16 uint64 = 4294967295
    _ = v13
    _ = v16
    return 624
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    var v3 int8 = 1
    v4 := []uint32{2147483647, 42}
    var v5 string = "hello"
    v6 := []uint32{uint32(v1), 959, 42, 42}
    v17 := GenS1{F0: -9223372036854775808, F1: 100, F2: 128}
    var v18 bool = (((!(0 != 840)) && true) && true)
    v1 -= (^(v1))
    var v19 int16 = 0
    var v20 uint32 = 2147483647
    if (!v18) {
        fmt.Println(v19)
    }
    v21 := GenS2{F0: false, F1: ((v19 ^ v19) != -32768), F2: -2147483648}
    var v22 int32 = -1
    var v23 float64 = 1e10
    v20 += v20
    var v24 int64 = 1
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
