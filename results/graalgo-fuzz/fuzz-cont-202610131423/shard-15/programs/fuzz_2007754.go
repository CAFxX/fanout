package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 uint32, p1 string, p2 string) uint32 {
    if (1 >= 127) {
        v4 := []int{1, 32767, -9223372036854775808}
        fmt.Println(p2)
        _ = v4
    } else {
        var v5 int16 = -32768
        var v6 string = "hello"
        v7 := GenS1{F0: uint16(p0), F1: 128, F2: true}
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 string = p1
    v9 := GenS2{F0: 2}
    var v10 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 int64, p1 uint, p2 uint16) uint32 {
    p2 += 65535
    v11 := GenS2{F0: 1}
    var v12 uint64 = 2
    v13 := []int{-2147483648, 32768}
    p1 = (p1 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    v3 := []int{0, 32767}
    var v14 uint = 608
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 float64 = 1e10
            _ = fi16
            _ = v17
        }
        v14 -= v14
        _ = fi15
    }
    switch "hello" {
    case "\n\t":
        for fi18 := 0; fi18 < 3; fi18++ {
            v14 = (v14 / ((v14) | 1))
            v14 -= (v14 + (v14 + v14))
            _ = fi18
        }
    case "hello":
        var v19 float32 = 1e10
        _ = v19
    default:
        var v20 int64 = 0
        _ = v20
    }
    v14 = (v14 % ((65535) | 1))
    if (227 >= 42) {
        var v21 uint8 = 1
        _ = v21
    }
    var v22 int32 = -2147483648
    var v23 uint8 = 255
    if ((2 <= 0) && false) {
        v3 = append(v3, 32767)
        switch 127 {
        case 9223372036854775807:
            var v24 int16 = -32768
            _ = v24
        case 4294967295:
            for fi25 := 0; fi25 < 4; fi25++ {
                for fi26 := 0; fi26 < 14; fi26++ {
                    v27 := GenS3{F0: (v23 - (v23 / (((v23 << (((v14 << (((^(v14))) & 63))) & 63))) | 1))), F1: -648}
                    _ = fi26
                    _ = v27
                }
                v3 = append(v3, 620)
                if false {
                    var v28 int32 = v22
                    if false {
                        v29 := GenS4{F0: 253}
                        _ = v29
                    }
                    _ = v28
                } else {
                    if (true || true) {
                        v30 := GenS2{F0: v23}
                        var v31 uint32 = 2
                        v32 := GenS5{F0: (v14 * 2), F1: 83, F2: (v23 + (^(v23)))}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v2[127] = true
                }
                _ = fi25
            }
        case 1:
            if true {
                fmt.Println(v22)
            } else {
                var v33 uint8 = (^(v23))
                v34 := GenS6{F0: v22, F1: (-(v22))}
                _ = v33
                _ = v34
            }
        }
    }
    v14 = v14
    fmt.Println(v22, v14, v23)
    v35 := GenS1{F0: 0, F1: v23, F2: false}
    v14 += (v14 - v14)
    for fi36 := 0; fi36 < 12; fi36++ {
        v22 = v22
        for fi37 := 0; fi37 < 5; fi37++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch (fi36 * (^(fi37))) {
            case 32768:
                switch "a" {
                case "a":
                    fmt.Println(v22)
                default:
                    v14 = uint(v23)
                }
            case 0:
                _ = len(v3)
            default:
                var v38 string = "hello"
                _ = v38
            }
            _ = fi37
        }
        _ = fi36
    }
    if (uint16(v23) == 42) {
        switch "\n\t" {
        case "":
            if true {
                if (true || (!false)) {
                    v1[42] = 1
                    switch (("\n\t" + "\n\t") + (("\n\t" + "\n\t") + "hello")) {
                    case "hello":
                        var v39 int16 = int16(v22)
                        _ = v39
                    }
                    v23 = (v23 >> (((v14 & v14)) & 63))
                }
            } else {
                fmt.Println(v22, v22, v14)
            }
        }
        v2[4294967295] = (-400 < 2147483647)
        for fi40 := 0; fi40 < 13; fi40++ {
            _ = len(v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v41 := GenS7{F0: (v14 & v14), F1: 42, F2: 500}
            _ = fi40
            _ = v41
        }
    }
    fmt.Println(v23)
    for fi42 := 0; fi42 < 3; fi42++ {
        v22 = v22
        v22 = v22
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
}
