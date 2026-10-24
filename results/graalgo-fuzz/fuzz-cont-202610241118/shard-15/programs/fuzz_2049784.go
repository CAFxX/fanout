package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int64, p1 uint, p2 int) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := []int{(fi6 - (fi6 >> ((0) & 63))), (p2 + (fi7 << ((42) & 63))), (fi7 - p2), fi7}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: p2}
    if false {
        fmt.Println(p2)
        v10 := []int{p2, p2, p2}
        _ = v10
    } else {
        p0 = p0
        v11 := GenS2{F0: 4294967295, F1: 2147483647}
        _ = v11
    }
    var v12 int8 = 123
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 933
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{0, 1}
    v3 := make(map[uint32]string)
    v4 := []uint32{1, 2147483648}
    var v5 int32 = -17
    for fi13 := 0; fi13 < 15; fi13++ {
        v5 = (^(v5))
        _ = fi13
    }
    delete(v1, 1)
    var v14 int8 = 127
    v5 -= (^(v5))
    if true {
        fmt.Println(v14, v14)
        for fi15 := 0; fi15 < 14; fi15++ {
            v14 -= (v14 & (v14 & 100))
            _ = fi15
        }
    } else {
        v16 := GenS3{F0: 0, F1: 992}
        v14 -= (v14 >> ((761) & 63))
        for fi17 := 0; fi17 < 6; fi17++ {
            delete(v1, fi17)
            _ = fi17
        }
        _ = v16
    }
    var v18 uint8 = 64
    if (!(!true)) {
        var v19 uint = 4294967295
        v19 = (v19 & (v19 >> ((v19) & 63)))
        _ = v19
    } else {
        v1[-638] = (-1 < -474)
    }
    if (!(!(false || true))) {
        var v20 int16 = -32768
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
