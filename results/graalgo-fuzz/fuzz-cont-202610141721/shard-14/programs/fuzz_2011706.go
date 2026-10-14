package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int32 }
func genFn0(p0 bool, p1 uint64, p2 bool) int8 {
    p0 = (192 >= 2147483647)
    if true {
        v7 := []int{}
        fmt.Println(p0, p2, p0)
        _ = v7
    }
    var v8 uint64 = p1
    v9 := GenS1{F0: -608, F1: 949}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 int32 = 1
    v3 := make(map[int]string)
    var v4 uint32 = 2147483647
    v5 := make(map[int]uint32)
    var v6 string = "a"
    if (-689 < v2) {
        v10 := []int{0}
        _ = v10
    }
    v11 := []int{128}
    var v12 int32 = 42
    var v13 string = (v6 + "a")
    v14 := GenS2{F0: 255, F1: v6}
    v15 := GenS2{F0: 4294967295, F1: (v6 + v6)}
    if false {
        var v16 string = (v13 + v13)
        if true {
            if (v4 == v4) {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 float32 = -2.25
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v19 := GenS1{F0: 9223372036854775807, F1: -1}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        } else {
            fmt.Println(v2, v13, v16)
        }
        v13 = "hello"
        _ = v16
    } else {
        fmt.Println(v2, v13)
    }
    v20 := v3[-2147483648]
    _ = len(v11)
    v13 = v13
    v4 = (v4 << ((4294967295) & 63))
    var v21 float32 = 3.14159
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            if ((2 <= 65535) || false) {
                v13 = v6
                var v24 bool = (-32768 > 1000)
                _ = v24
            }
            _ = fi23
        }
        v25 := GenS1{F0: 42, F1: -1}
        v26 := GenS2{F0: 679, F1: (v20 + v20)}
        _ = fi22
        _ = v25
        _ = v26
    }
    v27 := GenS3{F0: 130, F1: 2}
    v28 := GenS4{F0: 175, F1: (v12 % (((v12 % ((v12) | 1))) | 1))}
    if (!(42 < 42)) {
        if false {
            var v29 int16 = -490
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 bool = true
                _ = fi30
                _ = v31
            }
            if false {
                var v32 uint16 = 2
                for fi33 := 0; fi33 < 13; fi33++ {
                    var v34 float64 = 1e10
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 float32 = float32(((3.14159 / v34) / float64((v21 * 0.1))))
                        _ = fi35
                        _ = v36
                    }
                    _ = fi33
                    _ = v34
                }
                v37 := GenS1{F0: -9223372036854775808, F1: (v2 & v2)}
                _ = v32
                _ = v37
            }
            _ = v29
        } else {
            _ = len(v3)
            var v38 int8 = (v1 | 70)
            _ = v38
        }
        var v39 uint32 = 65535
        _ = v39
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        switch fi40 {
        case 9223372036854775807:
            var v41 int8 = (v1 - 0)
            _ = v41
        case -9223372036854775808:
            if len(v11) > 0 {
                v11[0] = -1
            }
        default:
            fmt.Println(fi40, v20)
        }
        if (v21 > 3.14159) {
            for fi42 := 0; fi42 < 17; fi42++ {
                _ = len(v11)
                _ = fi42
            }
            v43 := GenS5{F0: -1}
            _ = v43
        }
        _ = fi40
    }
    if false {
        v13 = ("a" + v13)
        v44 := GenS6{F0: 42, F1: v2}
        var v45 bool = false
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
