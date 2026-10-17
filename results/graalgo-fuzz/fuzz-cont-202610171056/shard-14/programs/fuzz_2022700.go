package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint; F2 int32 }
func genFn0(p0 bool, p1 uint) int16 {
    v4 := GenS1{F0: 0, F1: -882, F2: 628}
    p0 = p0
    p1 = (p1 ^ (p1 ^ p1))
    _ = p0
    _ = p1
    _ = v4
    return -32768
}
func genFn1(p0 uint16, p1 uint32, p2 uint8) int64 {
    fmt.Println(p0, p2, p0)
    fmt.Println(p0, p1)
    v5 := []int{42}
    var v6 uint = 188
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint = 2
    var v2 uint64 = uint64(v1)
    v3 := make(map[uint32]bool)
    v7 := []int{127, -9223372036854775808, 427, 407}
    v1 = (v1 + v1)
    v2 += (v2 % ((4294967295) | 1))
    for fi8 := 0; fi8 < 7; fi8++ {
        if true {
            v1 -= v1
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            if (2 <= 127) {
                var v10 uint8 = 128
                var v11 float64 = 0.0
                var v12 float64 = v11
                _ = v10
                _ = v11
                _ = v12
            } else {
                v2 = (v2 << (((^(v1))) & 63))
            }
            var v13 int8 = 93
            _ = fi9
            _ = v13
        }
        v14 := GenS2{F0: 42, F1: fi8}
        _ = fi8
        _ = v14
    }
    var v15 float32 = (1e10 / (0.0 + 0.1))
    v1 += v1
    v1 = 2
    if (32767 == 548) {
        var v16 int64 = -320
        var v17 uint = 225
        v1 += 255
        _ = v16
        _ = v17
    } else {
        fmt.Println(v15, v2)
        _ = len(v3)
        if (1e10 != ((((0.1 + (-2.25 + (((1e10 * 0.0) + 1e10) + 1.5))) * (3.14159 - 0.1)) - 1.5) * (-2.25 / (((float64(v15) - 1.5) - 1.5) - (((float64(v15) - 1e10) / float64(v15)) - ((float64((v15 * float32(0.1))) + 1e10) * 3.14159)))))) {
            var v18 uint8 = 95
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int8 = -81
        var v21 uint = v1
        _ = fi19
        _ = v20
        _ = v21
    }
    v15 -= (v15 / v15)
    var v22 int = -452
    v15 += v15
    if false {
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = 1
            _ = fi23
            _ = v24
        }
    }
    if (!true) {
        var v25 uint16 = 433
        var v26 int8 = 30
        _ = v25
        _ = v26
    } else {
        fmt.Println(v2)
    }
    v27 := GenS3{F0: "\n\t"}
    for fi28 := 0; fi28 < 8; fi28++ {
        switch "foo bar" {
        case "foo bar":
            for fi29 := 0; fi29 < 14; fi29++ {
                switch (-(fi29)) {
                case 4294967295:
                    if true {
                        delete(v3, 2)
                        v30 := GenS3{F0: "\n\t"}
                        if (!true) {
                            var v31 uint32 = 2147483647
                            v15 = 0.0
                            _ = v31
                        } else {
                            if false {
                                fmt.Println(v1, fi28)
                                if ((true || (!(false && (254 == 154)))) || (((65535 == 679) && false) || false)) {
                                    v32 := GenS2{F0: 0, F1: fi28}
                                    _ = v32
                                }
                            }
                        }
                        _ = v30
                    } else {
                        var v33 uint64 = v2
                        _ = v33
                    }
                default:
                    v15 = (v15 * ((v15 / v15) - v15))
                }
                _ = fi29
            }
        default:
            v34 := GenS4{F0: 154, F1: 0, F2: -320}
            _ = v34
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v7
}
