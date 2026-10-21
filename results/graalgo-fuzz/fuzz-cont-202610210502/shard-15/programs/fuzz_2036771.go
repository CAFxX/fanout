package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int8) int8 {
    v6 := GenS1{F0: 130}
    p0 -= (p0 & (p0 << ((620) & 63)))
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 4294967295}
        v9 := []int{fi7, (fi7 / (((fi7 ^ -330)) | 1)), fi7, fi7}
        v9 = append(v9, fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{1, 1, 32767}
    _ = p0
    _ = v10
    _ = v6
    return (p0 + p0)
}
func genFn1(p0 uint, p1 string, p2 string) bool {
    switch -895 {
    case 127:
        switch 42 {
        case 42:
            var v11 int8 = 1
            _ = v11
        case 101:
            var v12 uint = p0
            _ = v12
        case 1:
            var v13 uint8 = 42
            _ = v13
        default:
            for fi14 := 0; fi14 < 8; fi14++ {
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := GenS2{F0: -948, F1: 42, F2: 2}
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := []int{84, fi17, int(p0), 874}
                    _ = fi17
                    _ = v18
                }
                _ = fi14
            }
        }
    default:
        var v19 int64 = -61
        _ = v19
    }
    v20 := []int{326, 184}
    v21 := GenS2{F0: -1, F1: 1, F2: 905}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return (0 >= 2147483647)
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[uint32]uint32)
    var v3 bool = (!false)
    var v4 uint32 = 65535
    var v5 string = "hello"
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 float64 = 1.5
        _ = fi22
        _ = v23
    }
    switch (v5 + v5) {
    case "a":
        v3 = ((v4 << ((4294967295) & 63)) > (v4 + v4))
    default:
        v3 = v3
    }
    if true {
        v4 -= v4
    }
    var v24 int16 = -32768
    switch "hello" {
    case "hello":
        v5 = v5
    case "a":
        fmt.Println(v24)
    default:
        v25 := GenS3{F0: 545, F1: -18, F2: 189}
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v3 = (0 > -9223372036854775808)
        var v27 int = -2147483648
        var v28 uint64 = 431
        _ = fi26
        _ = v27
        _ = v28
    }
    if (714 > 0) {
        var v29 uint32 = (v4 / ((2147483648) | 1))
        v29 = (v4 | (v29 - (v4 << ((4294967295) & 63))))
        _ = len(v1)
        _ = v29
    } else {
        v4 = 411
    }
    switch -165 {
    case 808:
        switch -9223372036854775808 {
        case 32768:
            v30 := GenS3{F0: 42, F1: genFn0(127), F2: 42}
            _ = v30
        case -162:
            fmt.Println(v3)
        }
    case 490:
        fmt.Println(v4, v5, v4)
    default:
        if v3 {
            var v31 float64 = 1.5
            _ = v31
        } else {
            if (659 < (-(v24))) {
                v32 := GenS2{F0: -838, F1: -1, F2: 603}
                if false {
                    v5 = v5
                    var v33 int = 4294967295
                    _ = v33
                } else {
                    v5 = "foo bar"
                    v34 := v2[(v4 ^ (v4 + (v4 << ((4294967295) & 63))))]
                    _ = v34
                }
                _ = v32
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                if v3 {
                    var v36 int16 = (v24 + v24)
                    _ = v36
                } else {
                    v4 = (v4 + v4)
                    for fi37 := 0; fi37 < 18; fi37++ {
                        if (!(1 <= -1)) {
                            v38 := GenS2{F0: int32(v4), F1: -791, F2: 798}
                            v4 = 682
                            _ = v38
                        }
                        var v39 uint32 = 1
                        _ = fi37
                        _ = v39
                    }
                    v40 := GenS3{F0: 1, F1: -67, F2: 437}
                    _ = v40
                }
                _ = fi35
            }
        }
    }
    if v3 {
        _ = len(v2)
        var v41 int16 = (v24 * v24)
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
