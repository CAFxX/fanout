package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int32 }
func genFn0(p0 int32, p1 int16) int {
    var v6 uint64 = 42
    var v7 bool = false
    v8 := GenS1{F0: 107, F1: int64(v6)}
    if v7 {
        var v9 int64 = -716
        _ = v9
    }
    var v10 bool = (v7 || (0 >= 100))
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 bool) uint8 {
    for fi11 := 0; fi11 < 2; fi11++ {
        switch 0 {
        case -9223372036854775808:
            v12 := GenS1{F0: 320, F1: 734}
            _ = v12
        }
        fmt.Println(fi11)
        if p0 {
            if p0 {
                v13 := []int{genFn0(0, 1000), (-(fi11)), fi11}
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: 0, F1: 1}
                    if false {
                        fmt.Println(fi11, p0)
                        var v16 uint32 = 2147483647
                        if (1e10 < (3.14159 / float64(1e10))) {
                            v16 -= v16
                            var v17 uint = 0
                            _ = v17
                        } else {
                            if true {
                                if (false && p0) {
                                    var v18 uint8 = uint8(v16)
                                    v19 := GenS2{F0: v18, F1: (fi14 % (((^(fi11))) | 1)), F2: v18}
                                    var v20 uint32 = (v16 | v16)
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                            }
                        }
                        _ = v16
                    }
                    v21 := GenS3{F0: -2147483648}
                    _ = fi14
                    _ = v15
                    _ = v21
                }
                _ = v13
            } else {
                v22 := GenS1{F0: 4294967295, F1: 462}
                v23 := GenS2{F0: 114, F1: fi11, F2: 2}
                v24 := []int{(fi11 + fi11), genFn0(215, -440), fi11, fi11}
                _ = v22
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                if (!false) {
                    var v26 float32 = (1e10 * 0.1)
                    var v27 float32 = (-2.25 / (1.5 / 1.0))
                    _ = v26
                    _ = v27
                }
                _ = fi25
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS3{F0: 1}
                v30 := []int{(fi28 + (fi11 * genFn0(-987, 32767))), fi28}
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            p0 = (0 != 42)
        }
        _ = fi11
    }
    p0 = true
    _ = p0
    return 2
}

func main() {
    v1 := []string{"a"}
    var v2 float64 = 1.5
    var v3 uint8 = 0
    var v4 uint64 = 0
    v5 := make(map[uint32]uint32)
    fmt.Println(v4)
    v31 := GenS2{F0: 140, F1: 9223372036854775807, F2: v3}
    if false {
        if (true || (693 == uint32(v3))) {
            fmt.Println(v2)
            v32 := GenS4{F0: 255, F1: -460}
            v3 += (^(v3))
            _ = v32
        }
    }
    _ = len(v1)
    var v33 int = 846
    var v34 uint8 = (v3 << ((42) & 63))
    v35 := GenS5{F0: (v33 >> ((255) & 63)), F1: v33}
    switch v33 {
    case 1:
        v36 := GenS4{F0: 1, F1: 630}
        _ = v36
    case 128:
        fmt.Println(v33, v3)
    case 42:
        v37 := GenS6{F0: 2147483647}
        _ = v37
    default:
        _ = len(v5)
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if true {
        fmt.Println(v3)
        var v38 uint16 = 0
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
