package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 int32) int64 {
    v4 := GenS1{F0: 2, F1: (p1 % ((p1) | 1))}
    var v5 uint64 = 578
    p1 += -551
    if true {
        fmt.Println(v5, p1, p1)
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{2147483647, (fi6 & -71), 453, p0}
            _ = fi6
            _ = v7
        }
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []uint32{4294967295, 4294967295}
    var v2 uint64 = 0
    v3 := make(map[uint32]bool)
    if true {
        v8 := GenS2{F0: -2147483648, F1: -9223372036854775808, F2: (-2147483648 < 2147483647)}
        v1 = append(v1, 120)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: 2147483647, F1: -2147483648}
            var v11 uint32 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int8 = 100
    v13 := GenS2{F0: 1, F1: -551, F2: false}
    for fi14 := 0; fi14 < 5; fi14++ {
        if true {
            var v15 int64 = -296
            _ = v15
        } else {
            var v16 uint16 = 42
            _ = v16
        }
        _ = fi14
    }
    delete(v3, 514)
    var v17 uint16 = 621
    v18 := GenS3{F0: 219}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int32 = 42
        _ = fi19
        _ = v20
    }
    if (false && true) {
        var v21 bool = (1 < 9223372036854775807)
        _ = v21
    } else {
        for fi22 := 0; fi22 < 12; fi22++ {
            delete(v3, 2)
            for fi23 := 0; fi23 < 19; fi23++ {
                v12 = -100
                _ = fi23
            }
            var v24 int8 = v12
            _ = fi22
            _ = v24
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v25 bool = (255 != 255)
        _ = v25
    }
    var v26 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
}
