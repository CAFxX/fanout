package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 uint16, p1 uint8) uint64 {
    p0 = uint16(p1)
    v6 := GenS1{F0: -1, F1: 2, F2: -1}
    v7 := GenS1{F0: 9223372036854775807, F1: 931, F2: 100}
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 int32, p2 uint8) int8 {
    p1 -= p1
    switch 32768 {
    case 1:
        var v8 int16 = 0
        _ = v8
    case -940:
        v9 := GenS1{F0: 127, F1: 41, F2: 100}
        _ = v9
    case -1:
        v10 := GenS1{F0: 4294967295, F1: 0, F2: 42}
        _ = v10
    default:
        p1 += (p1 / ((p1) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    return 20
}

func main() {
    var v1 int64 = -480
    var v2 uint64 = 790
    var v3 int32 = 0
    var v4 int8 = 1
    var v5 float32 = 0.0
    v3 -= (v3 / ((42) | 1))
    v4 = v4
    v11 := []int{2147483647}
    var v12 string = "hello"
    var v13 int8 = genFn1(42, (v3 | v3), 0)
    var v14 uint32 = 42
    v5 = float32((((0.0 * ((((3.14159 - ((1.5 - float64((float32(float64(v5)) / v5))) * 3.14159)) / (-2.25 / ((float64(((v5 - (((1.5 + ((3.14159 + (v5 + float32((0.1 * -2.25)))) - 1e10)) + (-2.25 - v5)) - v5)) * v5)) * 0.1) - 3.14159))) * 1.5) - (0.0 * 1.5))) / (3.14159 - (1e10 + -2.25))) / 1.5))
    var v15 int8 = (v13 / (((v4 & v4)) | 1))
    fmt.Println(v1, v13)
    for fi16 := 0; fi16 < 20; fi16++ {
        switch (fi16 / ((fi16) | 1)) {
        case -9223372036854775808:
            if false {
                var v17 uint16 = 208
                var v18 uint32 = v14
                var v19 uint64 = (v2 % (((v2 ^ genFn0((v17 * (v17 | v17)), 0))) | 1))
                _ = v17
                _ = v18
                _ = v19
            } else {
                v3 = v3
                v4 -= (v4 / (((v15 + (v13 % (((v15 * int8(v14))) | 1)))) | 1))
            }
        default:
            var v20 int32 = (v3 >> ((469) & 63))
            _ = v20
        }
        var v21 uint64 = v2
        _ = fi16
        _ = v21
    }
    if ((((v1 | (v1 + int64(v3))) == v1) && (uint64(v14) < v2)) || true) {
        if true {
            if true {
                v22 := GenS1{F0: 255, F1: v2, F2: v15}
                switch 1 {
                case 683:
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if ((v12 != v12) && (true || ((1e10 + (v5 * 1.5)) > v5))) {
                            v24 := GenS1{F0: fi23, F1: genFn0(337, 117), F2: (v15 - v13)}
                            _ = v24
                        }
                        fmt.Println(v3)
                        _ = fi23
                    }
                case -480:
                    v4 = (v15 / (((v4 | genFn1(42, int32(v15), 128))) | 1))
                }
                switch -432 {
                case 32768:
                    v25 := GenS1{F0: -1, F1: v2, F2: (-(v15))}
                    _ = v25
                case 128:
                    var v26 int64 = (-(v1))
                    _ = v26
                case -587:
                    v14 = (v14 + v14)
                }
                _ = v22
            } else {
                fmt.Println(v1, v14, v2)
            }
        } else {
            if (811 <= 1) {
                switch v12 {
                case "":
                    _ = len(v11)
                case "foo bar":
                    fmt.Println(v12, v5, v12)
                case "hello":
                    if false {
                        for fi27 := 0; fi27 < 9; fi27++ {
                            var v28 float64 = 0.1
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            _ = fi27
                            _ = v28
                        }
                    }
                }
                v4 -= int8(v14)
                var v29 float32 = ((v5 + v5) - ((v5 / 0.1) * v5))
                _ = v29
            }
            v30 := GenS1{F0: -9223372036854775808, F1: (^(v2)), F2: (v4 >> ((255) & 63))}
            _ = v30
        }
        if ((v13 << ((0) & 63)) > (v13 * (^(v15)))) {
            _ = len(v11)
        }
    }
    v31 := GenS1{F0: 0, F1: v2, F2: v13}
    var v32 string = v12
    if (false && false) {
        var v33 string = v32
        v34 := GenS1{F0: -2147483648, F1: 0, F2: v13}
        _ = v33
        _ = v34
    } else {
        var v35 int16 = -67
        var v36 int8 = (^(v4))
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
