package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 uint, p2 bool) uint64 {
    switch 32768 {
    case 32768:
        var v6 uint64 = 408
        _ = v6
    case 116:
        var v7 uint32 = 42
        _ = v7
    case 78:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 string = "foo bar"
            var v10 float32 = ((1e10 / 1.0) - (3.14159 - (0.1 - (0.0 / (1.5 / (float32(0.1) + 1.5))))))
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 uint64 = 42
    var v12 float32 = (3.14159 - ((0.0 / (0.1 - (1e10 - 1.5))) + 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return v11
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1.5
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    var v13 uint16 = 255
    switch 127 {
    case 128:
        var v14 uint = 65535
        _ = v14
    default:
        fmt.Println(v13)
    }
    v2 = -2.25
    var v15 int64 = 42
    var v16 float64 = float64((((float32(float64(((v2 - float32((1.5 / 1.5))) - (float32(3.14159) * v2)))) / 1.0) + v2) / float32(3.14159)))
    v2 = -2.25
    v17 := GenS1{F0: 807, F1: 1000}
    if (!(42 <= 127)) {
        v18 := GenS1{F0: 1, F1: 42}
        v19 := GenS1{F0: 202, F1: -1}
        var v20 int8 = int8(v15)
        _ = v18
        _ = v19
        _ = v20
    } else {
        v21 := GenS1{F0: 0, F1: 241}
        v22 := v5[196]
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v15, v13)
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 uint16 = 0
                var v26 bool = false
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
        _ = v21
        _ = v22
    }
    v13 += 356
    v27 := v1[127]
    var v28 int8 = -128
    fmt.Println(v28, v13)
    var v29 int16 = 1000
    v30 := []int{4294967295}
    for fi31 := 0; fi31 < 19; fi31++ {
        _ = len(v30)
        _ = fi31
    }
    v32 := GenS1{F0: 122, F1: v29}
    v2 = 0.0
    if (65535 == 323) {
        _ = len(v3)
        switch 127 {
        case -9223372036854775808:
            v30 = append(v30, 9223372036854775807)
        case 897:
            if ((v13 | (v13 & (v13 & v13))) == v13) {
                switch "a" {
                case "":
                    if true {
                        for fi33 := 0; fi33 < 16; fi33++ {
                            v34 := GenS2{F0: 1}
                            v28 += v28
                            v28 = (v28 / ((v28) | 1))
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        fmt.Println(v2, v28)
                    }
                }
                switch -1 {
                case 127:
                    v16 += (v16 * v16)
                case 42:
                    if true {
                        v1[0] = (v27 & v27)
                        for fi35 := 0; fi35 < 18; fi35++ {
                            fmt.Println(v15, v13)
                            switch "a" {
                            case "\n\t":
                                var v36 uint = 1
                                _ = v36
                            case "":
                                v37 := GenS1{F0: 249, F1: v29}
                                _ = v37
                            case "hello":
                                v29 = v29
                            default:
                                var v38 uint32 = v27
                                _ = v38
                            }
                            _ = fi35
                        }
                    } else {
                        fmt.Println(v28, v27, v29)
                        v39 := GenS3{F0: -166, F1: "foo bar"}
                        _ = v39
                    }
                }
            } else {
                v2 = v2
                v40 := GenS4{F0: -2147483648, F1: (v15 | v15)}
                _ = v40
            }
        case 128:
            if len(v30) > 0 {
                v30[0] = 148
            }
        default:
            v13 = v13
        }
    } else {
        v41 := GenS5{F0: 739, F1: true}
        for fi42 := 0; fi42 < 7; fi42++ {
            if false {
                var v43 uint = 42
                _ = v43
            }
            for fi44 := 0; fi44 < 15; fi44++ {
                fmt.Println(v29, fi42)
                _ = fi44
            }
            var v45 float32 = (0.1 / v2)
            _ = fi42
            _ = v45
        }
        v46 := GenS6{F0: int8(v15), F1: 1, F2: (v27 & 4294967295)}
        _ = v41
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
