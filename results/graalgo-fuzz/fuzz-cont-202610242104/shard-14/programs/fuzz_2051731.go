package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 uint8; F2 int32 }
type GenS10 struct { F0 int8; F1 uint8 }
func genFn0(p0 int, p1 uint32, p2 uint16) uint16 {
    switch (-(p0)) {
    case 128:
        fmt.Println(p2, p0, p2)
    case -870:
        if true {
            var v4 int16 = 1000
            if (v4 < int16(p0)) {
                p2 = (p2 << ((65535) & 63))
            } else {
                for fi5 := 0; fi5 < 13; fi5++ {
                    switch ("" + "") {
                    case "a":
                        for fi6 := 0; fi6 < 16; fi6++ {
                            p2 = (p2 | (p2 + p2))
                            fmt.Println(p2)
                            _ = fi6
                        }
                    case "hello":
                        var v7 float32 = 0.0
                        _ = v7
                    case "\n\t":
                        for fi8 := 0; fi8 < 19; fi8++ {
                            var v9 int64 = -68
                            _ = fi8
                            _ = v9
                        }
                    }
                    v10 := GenS1{F0: p1, F1: 2}
                    var v11 uint64 = 4294967295
                    _ = fi5
                    _ = v10
                    _ = v11
                }
            }
            p1 += p1
            _ = v4
        } else {
            var v12 uint = 1
            _ = v12
        }
    }
    p0 += (-(p0))
    _ = p0
    _ = p1
    _ = p2
    return (p2 / ((p2) | 1))
}

func main() {
    var v1 int8 = -128
    var v2 int16 = -868
    v3 := make(map[uint32]string)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint8 = 255
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: (v1 + int8(v2))}
    v16 := GenS1{F0: 130, F1: 2}
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint = 255
        var v19 int16 = v2
        _ = fi17
        _ = v18
        _ = v19
    }
    if (2147483647 > 65535) {
        if (1 != -1) {
            v20 := GenS3{F0: 376}
            _ = v20
        } else {
            v21 := GenS4{F0: 658, F1: genFn0(865, 2147483647, 65535)}
            v22 := GenS5{F0: v1, F1: 1}
            v2 = (v2 << ((255) & 63))
            _ = v21
            _ = v22
        }
    } else {
        v1 += v1
        if (false || false) {
            for fi23 := 0; fi23 < 8; fi23++ {
                v24 := GenS6{F0: 4294967295}
                for fi25 := 0; fi25 < 13; fi25++ {
                    v2 = (v2 ^ -1)
                    v1 -= v1
                    v26 := GenS7{F0: 0, F1: 128, F2: (-(v1))}
                    _ = fi25
                    _ = v26
                }
                v2 = 42
                _ = fi23
                _ = v24
            }
        }
        var v27 int8 = (v1 / ((0) | 1))
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 float32 = (((0.1 / float32(0.1)) + -2.25) * (float32(float64(((-2.25 / (1.5 + 1e10)) - (0.1 / (3.14159 - float32(0.1)))))) + 0.1))
        fmt.Println(v29, fi28, fi28)
        _ = fi28
        _ = v29
    }
    v30 := []int{127, 9223372036854775807}
    v2 = (^(v2))
    switch -9223372036854775808 {
    case 994:
        v31 := GenS8{F0: 2147483647}
        _ = v31
    case -2147483648:
        var v32 uint = 785
        _ = v32
    case 127:
        v33 := GenS9{F0: (2147483647 > -1), F1: 2, F2: 1}
        _ = v33
    default:
        if false {
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 float32 = 3.14159
                var v36 float32 = v35
                _ = fi34
                _ = v35
                _ = v36
            }
            if (536 < -9223372036854775808) {
                v2 = (v2 / ((42) | 1))
                var v37 uint32 = 586
                _ = v37
            } else {
                var v38 int64 = -1
                v39 := GenS10{F0: v1, F1: 128}
                var v40 int64 = v38
                _ = v38
                _ = v39
                _ = v40
            }
        } else {
            v41 := GenS8{F0: 1}
            v1 = (v1 % ((v1) | 1))
            _ = v41
        }
    }
    v1 = (v1 * v1)
    switch 1 {
    case 103:
        v42 := GenS4{F0: 42, F1: 1}
        _ = v42
    case 963:
        var v43 int8 = (v1 | v1)
        _ = v43
    default:
        v2 = (v2 | (v2 & int16(v1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
}
