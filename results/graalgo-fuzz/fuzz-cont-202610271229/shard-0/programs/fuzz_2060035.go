package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 bool; F1 int; F2 bool }
func genFn0(p0 uint64) int16 {
    v5 := GenS1{F0: 140, F1: -128, F2: 815}
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 0
        var v8 uint8 = 78
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return -450
}

func main() {
    var v1 uint = 535
    v2 := make(map[int]string)
    var v3 int32 = -846
    v4 := []int{2147483647}
    for fi9 := 0; fi9 < 5; fi9++ {
        v1 += (v1 - v1)
        fmt.Println(v3, v1, fi9)
        _ = fi9
    }
    var v10 uint8 = 255
    var v11 uint16 = 42
    var v12 uint8 = v10
    fmt.Println(v11, v1)
    v10 += v12
    v13 := GenS2{F0: (128 > 0), F1: 1}
    var v14 uint64 = 1
    v15 := GenS1{F0: v10, F1: 100, F2: (v3 - v3)}
    if ((1.5 - 1.5) >= (0.0 / 3.14159)) {
        v16 := GenS3{F0: 56, F1: v10, F2: v3}
        v17 := GenS2{F0: (!(!false)), F1: 42}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 float64 = ((float64((((((0.0 / float32(0.1)) * (-2.25 * (3.14159 / 1.0))) / 1.5) - 1e10) - 0.0)) * 3.14159) - 1.5)
        v11 -= v11
        _ = fi18
        _ = v19
    }
    switch -2147483648 {
    case -755:
        v1 = v1
    case 2147483647:
        v10 += v10
    case -986:
        v20 := GenS4{F0: 44, F1: int32(v1)}
        _ = v20
    }
    var v21 float32 = 3.14159
    if true {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 9; fi23++ {
                if len(v4) > 0 {
                    v4[0] = (fi22 + fi23)
                }
                var v24 uint16 = (v11 | v11)
                var v25 int32 = (v3 ^ v3)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = len(v4)
            if false {
                v1 = (^(v1))
                v12 += v10
                fmt.Println(fi22, v10)
            }
            _ = fi22
        }
        v11 -= 627
        v14 -= v14
    }
    var v26 uint16 = (v11 % ((2) | 1))
    if false {
        if len(v4) > 0 {
            v4[0] = 127
        }
        v3 = v3
        if ((^(v1)) > v1) {
            fmt.Println(v1)
        } else {
            v27 := GenS2{F0: true, F1: -1}
            _ = v27
        }
    } else {
        v10 += (v10 << ((v1) & 63))
        var v28 float32 = (v21 * (v21 * v21))
        if (v10 != (v10 % ((v12) | 1))) {
            v21 = 1.5
            for fi29 := 0; fi29 < 8; fi29++ {
                if true {
                    v14 -= v14
                    _ = len(v2)
                } else {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        var v31 uint8 = (v10 & 127)
                        if (0.0 >= (1e10 * 3.14159)) {
                            v32 := GenS5{F0: (549 == 268), F1: (^(fi29)), F2: ((!true) || true)}
                            _ = v32
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                v1 += v1
                var v33 int32 = (v3 - (v3 + -1))
                _ = fi29
                _ = v33
            }
            v11 += v26
        }
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
}
