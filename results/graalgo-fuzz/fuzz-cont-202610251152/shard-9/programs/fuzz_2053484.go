package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32 }
func genFn0() int {
    if (32767 <= 7) {
        switch "a" {
        case "hello":
            v6 := []int{32768}
            _ = v6
        }
        var v7 int8 = 0
        v7 += v7
        _ = v7
    } else {
        var v8 string = "\n\t"
        if (!(1 >= 0)) {
            v9 := GenS1{F0: -1, F1: -412, F2: -1}
            switch v8 {
            case "\n\t":
                v10 := []int{}
                _ = v10
            case "hello":
                if (("hello" + ("\n\t" + v8)) > ((v8 + v8) + v8)) {
                    if false {
                        v11 := GenS1{F0: 319, F1: -821, F2: -121}
                        _ = v11
                    }
                }
            }
            _ = v9
        } else {
            if false {
                if false {
                    v8 = v8
                }
                for fi12 := 0; fi12 < 7; fi12++ {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 int32 = -94
                        if ((1.5 / (-2.25 + -2.25)) <= 1e10) {
                            v14 -= v14
                        }
                        var v15 uint = 0
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    var v16 int16 = 32767
                    _ = fi12
                    _ = v16
                }
                var v17 uint64 = 863
                _ = v17
            }
        }
        _ = v8
    }
    var v18 int8 = -76
    v19 := []int{-1, 194, 128, 267}
    _ = v18
    _ = v19
    return 127
}

func main() {
    var v1 uint64 = 2
    var v2 int8 = -128
    var v3 int16 = -32768
    var v4 uint = 1
    v5 := make(map[int]int)
    var v20 int = 42
    if (0.1 >= (1e10 - (float32(((0.0 / 3.14159) * 0.0)) - 1e10))) {
        v4 = v4
    }
    if (false || false) {
        if (!true) {
            v4 -= v4
            var v21 float64 = 1e10
            for fi22 := 0; fi22 < 7; fi22++ {
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 int32 = 42
                    v25 := GenS2{F0: 1, F1: (v24 - (v24 << (((v4 ^ 18446744073709551615)) & 63)))}
                    v26 := GenS2{F0: 9223372036854775807, F1: (v24 & (v24 - v24))}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                fmt.Println(v21, v1, v4)
                _ = fi22
            }
            _ = v21
        }
        switch (v20 | v20) {
        case 989:
            switch "\n\t" {
            case "\n\t":
                for fi27 := 0; fi27 < 9; fi27++ {
                    v5[fi27] = 2147483647
                    _ = fi27
                }
            case "foo bar":
                v3 -= (^(v3))
            default:
                var v28 uint = 894
                _ = v28
            }
        case 128:
            var v29 int8 = v2
            _ = v29
        case 9223372036854775807:
            v5[(v20 + (v20 % (((v20 % ((int(v3)) | 1))) | 1)))] = v20
        default:
            var v30 int64 = -1
            _ = v30
        }
        v31 := []int{v20}
        _ = v31
    }
    var v32 uint64 = v1
    if (255 == 880) {
        v33 := GenS2{F0: 1, F1: 42}
        for fi34 := 0; fi34 < 7; fi34++ {
            v2 += (v2 | (v2 >> ((v4) & 63)))
            switch 697 {
            case 465:
                fmt.Println(v20, v32)
            case 1:
                fmt.Println(v3)
            }
            var v35 int8 = -128
            _ = fi34
            _ = v35
        }
        v36 := v5[genFn0()]
        _ = v33
        _ = v36
    } else {
        var v37 uint64 = (v32 / (((^(v32))) | 1))
        _ = v37
    }
    var v38 bool = (v32 <= 42)
    if v38 {
        fmt.Println(v2)
    } else {
        delete(v5, genFn0())
        v39 := GenS3{F0: 2}
        _ = v39
    }
    var v40 float32 = 3.14159
    var v41 uint16 = 179
    fmt.Println(v3)
    v1 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
