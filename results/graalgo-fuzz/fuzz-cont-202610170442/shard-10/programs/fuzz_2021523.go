package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 int8 }
type GenS3 struct { F0 int32; F1 uint; F2 string }
func genFn0(p0 uint64, p1 int32, p2 int64) int8 {
    var v7 int16 = 1
    var v8 float32 = float32(3.14159)
    v9 := []int{4294967295, 280, -504, 222}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 int64) uint {
    if (!false) {
        v10 := []int{-1}
        if (42 < 65535) {
            v11 := GenS1{F0: 0, F1: 2, F2: true}
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(p1)
                _ = fi12
            }
            _ = v11
        }
        if (true && false) {
            if len(v10) > 0 {
                v10[0] = 32767
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                if true {
                    if (p0 != (p0 % ((-969) | 1))) {
                        p1 = (p1 >> ((421) & 63))
                        p1 = (p1 ^ (p1 ^ p1))
                    } else {
                        v10 = append(v10, fi13)
                        var v14 int64 = (p1 / ((p1) | 1))
                        var v15 uint = 1
                        _ = v14
                        _ = v15
                    }
                    if (100 != 1) {
                        p0 += p0
                        var v16 float64 = 3.14159
                        _ = v16
                    }
                    p0 = (p0 * (p0 >> ((240) & 63)))
                } else {
                    var v17 uint8 = 219
                    switch "hello" {
                    case "\n\t":
                        v18 := GenS1{F0: v17, F1: 742, F2: false}
                        _ = v18
                    }
                    _ = v17
                }
                p1 += p1
                _ = fi13
            }
        }
        _ = v10
    } else {
        p1 += (p1 | p1)
        fmt.Println(p1, p1)
    }
    if false {
        p0 = -32768
        v19 := []int{}
        var v20 uint16 = 557
        _ = v19
        _ = v20
    } else {
        switch 32768 {
        case -9223372036854775808:
            for fi21 := 0; fi21 < 19; fi21++ {
                for fi22 := 0; fi22 < 15; fi22++ {
                    p1 = p1
                    p1 = p1
                    _ = fi22
                }
                _ = fi21
            }
        case 128:
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := []int{fi23, (^(fi23))}
                _ = fi23
                _ = v24
            }
        case 32767:
            for fi25 := 0; fi25 < 18; fi25++ {
                v26 := GenS2{F0: 42, F1: 0, F2: 127}
                _ = fi25
                _ = v26
            }
        default:
            var v27 float64 = (1e10 + 0.0)
            _ = v27
        }
        p1 = 0
    }
    _ = p0
    _ = p1
    return 139
}

func main() {
    v1 := []uint32{0}
    var v2 uint32 = 1
    var v3 int32 = -177
    v4 := []int{4294967295, 568}
    v5 := []uint32{}
    var v6 uint = 2
    for fi28 := 0; fi28 < 3; fi28++ {
        v29 := GenS3{F0: (v3 * 231), F1: 255, F2: ("hello" + "foo bar")}
        for fi30 := 0; fi30 < 18; fi30++ {
            fmt.Println(v3, v6)
            var v31 int64 = 0
            v31 += v31
            _ = fi30
            _ = v31
        }
        v32 := GenS2{F0: 2, F1: (-(v3)), F2: -1}
        _ = fi28
        _ = v29
        _ = v32
    }
    v2 = (v2 >> (((v6 % ((genFn1(1000, 972)) | 1))) & 63))
    _ = len(v5)
    v33 := GenS2{F0: v2, F1: (v3 + (v3 << (((v6 >> ((v6) & 63))) & 63))), F2: genFn0(756, 0, -248)}
    v6 -= (v6 & 2)
    var v34 uint8 = 42
    fmt.Println(v3, v6, v6)
    if (true || false) {
        if false {
            for fi35 := 0; fi35 < 3; fi35++ {
                if len(v1) > 0 {
                    v1[0] = (v2 * v2)
                }
                v3 -= (v3 + v3)
                _ = fi35
            }
            var v36 uint32 = v2
            _ = v36
        }
        _ = len(v1)
        for fi37 := 0; fi37 < 9; fi37++ {
            var v38 int8 = -1
            v34 = 209
            if len(v1) > 0 {
                v1[0] = 0
            }
            _ = fi37
            _ = v38
        }
    }
    v39 := GenS3{F0: v3, F1: (v6 << (((v6 | (v6 * 1))) & 63)), F2: "hello"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
