package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string; F1 int; F2 uint }
func genFn0(p0 int8, p1 int32) uint32 {
    var v7 uint8 = 0
    var v8 int64 = 42
    p0 = (p0 % (((p0 % ((91) | 1))) | 1))
    var v9 uint = 65535
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint64, p1 uint64) int16 {
    if false {
        fmt.Println(p1)
        switch 128 {
        case 32768:
            if true {
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 uint16 = 65535
                    _ = fi10
                    _ = v11
                }
                if ((false || true) && (!(281 <= 1))) {
                    switch 127 {
                    case 2147483647:
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 uint = 0
                            for fi14 := 0; fi14 < 17; fi14++ {
                                var v15 int8 = -128
                                if (-636 <= 42) {
                                    p0 = (p0 >> ((351) & 63))
                                    v16 := []int{4294967295, fi12, -2147483648, (fi14 & fi14)}
                                    _ = v16
                                } else {
                                    if true {
                                        p1 += 2
                                    } else {
                                        v17 := []int{(fi14 << (((v13 & 629)) & 63)), (fi14 + fi12), fi14, (-(fi12))}
                                        fmt.Println(v13)
                                        var v18 int8 = v15
                                        _ = len(v17)
                                        _ = v17
                                        _ = v18
                                    }
                                }
                                fmt.Println(fi14, v13)
                                _ = fi14
                                _ = v15
                            }
                            fmt.Println(p1)
                            _ = fi12
                            _ = v13
                        }
                    }
                } else {
                    p1 = p0
                }
                v19 := GenS1{F0: 1, F1: 0, F2: "hello"}
                _ = v19
            }
        case 127:
            v20 := GenS2{F0: 603, F1: 519}
            _ = v20
        default:
            v21 := GenS3{F0: 1, F1: p0}
            _ = v21
        }
    }
    p1 += 335
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 998
    var v3 int16 = 32767
    var v4 bool = false
    var v5 float64 = 1e10
    var v6 uint32 = 42
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{fi22, fi22}
        _ = fi22
        _ = v23
    }
    var v24 string = ("a" + "a")
    fmt.Println(v5, v2)
    v25 := GenS4{F0: v24, F1: 128, F2: 255}
    if (((("" + "") + v24) < ("hello" + v24)) && (v4 || (v6 >= 42))) {
        var v26 uint64 = (v2 % (((^(v2))) | 1))
        _ = v26
    }
    v27 := []int{172, -1}
    for fi28 := 0; fi28 < 16; fi28++ {
        v2 += 494
        v27 = append(v27, 0)
        fmt.Println(v6, v24)
        _ = fi28
    }
    v6 = (v6 - (v6 + (v6 >> ((255) & 63))))
    for fi29 := 0; fi29 < 8; fi29++ {
        if len(v27) > 0 {
            v27[0] = -236
        }
        _ = fi29
    }
    v3 = (v3 % ((-526) | 1))
    v6 = genFn0(127, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
