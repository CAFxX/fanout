package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint16; F1 string; F2 int64 }
type GenS7 struct { F0 bool; F1 uint16; F2 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0() uint8 {
    v6 := []int{42, 202, 32768, 32767}
    var v7 uint = 872
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint64 = 2
        _ = fi8
        _ = v9
    }
    if (false || ((v7 < (^(v7))) || (false && true))) {
        v10 := GenS1{F0: 1}
        _ = v10
    }
    if ((1 == -101) && false) {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint8 = 1
            v13 := GenS2{F0: 127, F1: ("a" + "hello"), F2: int32(v12)}
            _ = fi11
            _ = v12
            _ = v13
        }
        v7 -= v7
        fmt.Println(v7)
    }
    _ = v6
    _ = v7
    return 58
}
func genFn1(p0 int16) uint {
    p0 = (p0 + p0)
    switch 42 {
    case -875:
        for fi14 := 0; fi14 < 9; fi14++ {
            if false {
                v15 := []int{(fi14 >> ((65535) & 63)), (fi14 - (fi14 % ((169) | 1))), (^(fi14))}
                _ = v15
            } else {
                v16 := GenS3{F0: false, F1: 127}
                if (((0.1 + ((float64(0.0) - 1.5) - 3.14159)) * 1.5) != (float64(0.1) - ((0.1 / (0.1 + 0.0)) * 1e10))) {
                    var v17 uint = 1
                    switch "a" {
                    case "hello":
                        p0 = -651
                    }
                    fmt.Println(fi14, p0)
                    _ = v17
                }
                v18 := []int{(fi14 & 2147483647)}
                _ = v16
                _ = v18
            }
            if true {
                var v19 int32 = -1
                v20 := GenS3{F0: (1 > 42), F1: -874}
                p0 -= p0
                _ = v19
                _ = v20
            }
            _ = fi14
        }
    case 127:
        p0 -= (^(p0))
    case -132:
        var v21 int32 = 0
        _ = v21
    default:
        for fi22 := 0; fi22 < 9; fi22++ {
            p0 -= (p0 | (p0 | -437))
            _ = fi22
        }
    }
    var v23 uint32 = 2147483647
    fmt.Println(v23)
    p0 += (-(p0))
    _ = p0
    _ = v23
    return 255
}

func main() {
    v1 := []int{1, -2147483648}
    var v2 int32 = -2147483648
    var v3 uint32 = 828
    var v4 int8 = -1
    v5 := make(map[uint32]string)
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS4{F0: v4, F1: fi24}
        _ = fi24
        _ = v25
    }
    v26 := GenS5{F0: -533, F1: 1}
    v3 -= (v3 % ((v3) | 1))
    var v27 uint16 = 255
    if len(v1) > 0 {
        v1[0] = 415
    }
    var v28 uint64 = uint64(v3)
    if false {
        v28 += (v28 % ((v28) | 1))
        fmt.Println(v28, v27)
    }
    v27 = (v27 * v27)
    v2 = (v2 % ((v2) | 1))
    v29 := GenS6{F0: 2, F1: "foo bar", F2: 1}
    v27 = v27
    v3 += (v3 % ((uint32(v28)) | 1))
    v30 := GenS7{F0: (v3 >= v3), F1: uint16(v3), F2: (196 != 0)}
    var v31 float64 = 1e10
    if (!(238 > 127)) {
        if true {
            fmt.Println(v31)
            var v32 string = (("\n\t" + "") + "\n\t")
            _ = v32
        } else {
            v27 += (v27 - (v27 / (((^(v27))) | 1)))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "" {
            case "foo bar":
                fmt.Println(v4)
            case "hello":
                v27 = v27
            default:
                fmt.Println(v3, v2, v4)
            }
        }
        var v33 uint64 = 380
        var v34 int8 = (v4 | -128)
        _ = v33
        _ = v34
    } else {
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 int = fi35
            v36 = fi35
            var v37 bool = true
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    v27 = 1
    if false {
        v38 := GenS8{F0: 0}
        v39 := GenS9{F0: v27, F1: 776, F2: 32767}
        _ = v38
        _ = v39
    } else {
        if len(v1) > 0 {
            v1[0] = -1
        }
    }
    v40 := v5[400]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v5
}
