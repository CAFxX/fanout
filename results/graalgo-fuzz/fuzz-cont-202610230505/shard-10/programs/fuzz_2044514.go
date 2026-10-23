package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int64; F2 int64 }
func genFn0() uint64 {
    var v6 uint32 = 2
    var v7 int = 865
    v8 := GenS1{F0: 776, F1: -44, F2: 42}
    v7 += -181
    var v9 int32 = 994
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    var v3 uint16 = 47
    var v4 int8 = 88
    var v5 float64 = -2.25
    v10 := GenS1{F0: v1, F1: 122, F2: 1000}
    switch 0 {
    case 128:
        v11 := GenS2{F0: -616, F1: ("\n\t" + "a")}
        _ = v11
    case 4294967295:
        var v12 uint8 = 127
        _ = v12
    case 0:
        var v13 uint = 2
        _ = v13
    default:
        v14 := GenS3{F0: 42}
        _ = v14
    }
    var v15 int8 = (v4 * -5)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int8 = -1
        v18 := v2[1]
        var v19 uint32 = 65535
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := v2[65535]
        v22 := GenS3{F0: 128}
        for fi23 := 0; fi23 < 7; fi23++ {
            delete(v2, 0)
            var v24 uint64 = 880
            _ = fi23
            _ = v24
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v4 += (v15 - -128)
        for fi26 := 0; fi26 < 7; fi26++ {
            switch fi25 {
            case 52:
                if true {
                    var v27 uint8 = 59
                    fmt.Println(v3, v4, v1)
                    switch fi26 {
                    case 4294967295:
                        v4 += v15
                    default:
                        v2[273] = "a"
                    }
                    v28 := []int{fi25}
                    v15 = v15
                    _ = v27
                    _ = v28
                } else {
                    v5 = (v5 + v5)
                }
            case 32767:
                if (1 == -32768) {
                    fmt.Println(v3)
                }
            default:
                v15 -= (v15 % ((v15) | 1))
                v29 := GenS3{F0: 255}
                _ = v29
            }
            v30 := []int{fi25, (fi25 / (((fi26 - (fi26 ^ fi26))) | 1)), fi25}
            switch "\n\t" {
            case "hello":
                if false {
                    var v31 int8 = v4
                    v2[2147483648] = "a"
                    v31 += (v15 ^ -17)
                    _ = v31
                } else {
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    v32 := v2[114]
                    _ = v32
                }
            }
            _ = fi26
            _ = v30
        }
        for fi33 := 0; fi33 < 10; fi33++ {
            switch (^(fi25)) {
            case 9223372036854775807:
                fmt.Println(v3, fi25, fi33)
            case 42:
                var v34 int16 = 1000
                _ = v34
            default:
                v5 += (((v5 / v5) / v5) * v5)
            }
            v35 := GenS4{F0: ((-1 <= 6) || true)}
            switch fi25 {
            case 1:
                v36 := []int{32767, (fi25 ^ (fi33 / (((fi25 / ((fi25) | 1))) | 1))), 1}
                _ = v36
            case 158:
                v3 += (v3 / ((v3) | 1))
            case 4294967295:
                v3 += (^(v3))
            default:
                if (4294967295 != 0) {
                    var v37 float32 = ((1e10 - (float32(v5) + (((1.5 / (-2.25 / 1.5)) + (0.1 + 0.0)) + -2.25))) / 1e10)
                    v5 = (3.14159 + (v5 - v5))
                    _ = v37
                }
            }
            _ = fi33
            _ = v35
        }
        _ = fi25
    }
    if (false || false) {
        var v38 int = 128
        v1 = (v1 ^ uint(v3))
        v39 := v2[4294967295]
        _ = v38
        _ = v39
    }
    var v40 uint32 = 886
    var v41 uint8 = 116
    switch "\n\t" {
    case "hello":
        if true {
            v42 := GenS1{F0: 255, F1: -108, F2: 880}
            v15 -= 0
            _ = v42
        } else {
            var v43 uint = v1
            v5 -= (v5 + -2.25)
            _ = v43
        }
    case "a":
        delete(v2, 189)
    default:
        v44 := GenS5{F0: v1, F1: 42, F2: 1}
        _ = v44
    }
    switch "foo bar" {
    case "":
        var v45 string = "\n\t"
        _ = v45
    case "\n\t":
        v41 = v41
    }
    v5 = (v5 + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
