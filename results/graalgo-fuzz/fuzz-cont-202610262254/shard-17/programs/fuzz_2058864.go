package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint8, p1 bool, p2 uint32) int64 {
    v6 := []int{}
    var v7 int8 = 31
    v8 := GenS1{F0: -1, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint16, p1 string, p2 uint16) uint {
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 float64 = 0.1
        _ = fi9
        _ = v10
    }
    fmt.Println(p1)
    for fi11 := 0; fi11 < 19; fi11++ {
        p2 -= p0
        _ = fi11
    }
    v12 := GenS1{F0: -9223372036854775808, F1: 1}
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 0
    var v3 int64 = 42
    var v4 int32 = -2147483648
    var v5 int32 = (v4 >> ((2) & 63))
    v13 := GenS2{F0: -1}
    v5 += (v4 + (v4 << ((42) & 63)))
    v14 := GenS2{F0: 32767}
    v5 += 1
    v15 := GenS3{F0: (v5 / (((v4 * (v5 + v4))) | 1)), F1: 2, F2: 152}
    if (65535 <= 42) {
        fmt.Println(v3, v3, v4)
    }
    v16 := v1[42]
    v4 += v5
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS3{F0: 1, F1: 255, F2: 0}
        _ = fi17
        _ = v18
    }
    v19 := GenS4{F0: v16}
    if (((v5 - (v5 & v4)) > (v4 | v4)) && (2 <= 0)) {
        v20 := GenS5{F0: int32(v16), F1: v3}
        var v21 uint = 520
        var v22 int = -850
        _ = v20
        _ = v21
        _ = v22
    } else {
        v23 := []int{1}
        _ = v23
    }
    switch "hello" {
    case "":
        var v24 int64 = (v3 << ((65535) & 63))
        _ = v24
    case "foo bar":
        switch 1 {
        case 127:
            v4 = -283
        case 4294967295:
            for fi25 := 0; fi25 < 14; fi25++ {
                v4 += v4
                _ = fi25
            }
        default:
            v26 := GenS2{F0: 1000}
            _ = v26
        }
    case "hello":
        v4 = v5
    default:
        v3 += (-(v3))
    }
    switch 1 {
    case -2147483648:
        var v27 int16 = 1
        _ = v27
    case 251:
        v28 := []int{2147483647}
        _ = v28
    case -567:
        for fi29 := 0; fi29 < 2; fi29++ {
            for fi30 := 0; fi30 < 19; fi30++ {
                var v31 uint8 = 127
                _ = fi30
                _ = v31
            }
            v3 = (v3 - (v3 / ((268) | 1)))
            _ = fi29
        }
    default:
        fmt.Println(v2, v16, v2)
    }
    v32 := []int{785, 32767, 127, 2147483647}
    switch "a" {
    case "\n\t":
        if (!true) {
            var v33 bool = (134 >= 180)
            _ = v33
        }
    case "hello":
        for fi34 := 0; fi34 < 3; fi34++ {
            fmt.Println(v3, v5, fi34)
            for fi35 := 0; fi35 < 7; fi35++ {
                if false {
                    switch "" {
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            var v36 bool = true
                            _ = v36
                        case "a":
                            _ = len(v1)
                        case "foo bar":
                            v37 := GenS6{F0: (v5 >> ((42) & 63)), F1: true}
                            _ = v37
                        default:
                            var v38 int8 = (v2 ^ (v2 * (v2 << ((18446744073709551615) & 63))))
                            _ = v38
                        }
                    }
                    if (true && (!(!((v2 & 100) < v2)))) {
                        v39 := GenS7{F0: v16}
                        v1[1] = v16
                        _ = v39
                    }
                    var v40 float32 = float32(0.0)
                    _ = v40
                }
                var v41 uint8 = 1
                _ = fi35
                _ = v41
            }
            _ = fi34
        }
    default:
        var v42 uint64 = 2
        _ = v42
    }
    var v43 int8 = -123
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
}
