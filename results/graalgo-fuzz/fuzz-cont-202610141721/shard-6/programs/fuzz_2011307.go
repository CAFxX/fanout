package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8; F1 string; F2 int32 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint32; F1 string; F2 uint16 }

func main() {
    var v1 uint8 = 222
    v2 := make(map[uint32]int)
    var v3 uint = 42
    if true {
        v1 = (v1 << ((v3) & 63))
        v4 := GenS1{F0: 805, F1: -15, F2: 0}
        delete(v2, 65535)
        _ = v4
    }
    v5 := []int{}
    var v6 int = -491
    if (1 < -1) {
        var v7 int8 = -128
        v8 := GenS2{F0: 461, F1: 65535, F2: v7}
        if (((-2.25 / 1.5) + ((0.0 - 0.1) - 1e10)) >= 0.1) {
            var v9 int64 = 0
            for fi10 := 0; fi10 < 12; fi10++ {
                switch "\n\t" {
                case "hello":
                    v11 := GenS2{F0: 0, F1: 42, F2: 127}
                    _ = v11
                case "\n\t":
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v1 = v1
                        var v13 int = fi12
                        _ = fi12
                        _ = v13
                    }
                case "a":
                    if true {
                        var v14 int64 = v9
                        fmt.Println(v3, v9)
                        var v15 string = "\n\t"
                        _ = v14
                        _ = v15
                    } else {
                        v6 = (v6 - (v6 | 32768))
                    }
                default:
                    v16 := GenS3{F0: -1, F1: 0, F2: 389}
                    _ = v16
                }
                if (2 < 42) {
                    _ = len(v5)
                    var v17 uint8 = (v1 | v1)
                    var v18 string = (("hello" + ("\n\t" + "foo bar")) + ("foo bar" + ""))
                    _ = v17
                    _ = v18
                } else {
                    if true {
                        delete(v2, 42)
                        v3 -= v3
                        v19 := GenS3{F0: -1, F1: 32767, F2: 255}
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 6; fi20++ {
                            fmt.Println(v3)
                            fmt.Println(v9, v1, v9)
                            _ = fi20
                        }
                    }
                }
                v21 := GenS3{F0: -549, F1: 1000, F2: 0}
                _ = fi10
                _ = v21
            }
            _ = v9
        } else {
            var v22 bool = false
            _ = v22
        }
        _ = v7
        _ = v8
    } else {
        switch 127 {
        case 32767:
            var v23 float64 = (-2.25 - 1e10)
            _ = v23
        }
    }
    var v24 bool = true
    v25 := GenS4{F0: 892, F1: (v3 << (((v3 + v3)) & 63))}
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int16 = 233
        v27 -= 32767
        v28 := GenS3{F0: (v27 - 32767), F1: 1, F2: 937}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int64 = int64(v1)
    if len(v5) > 0 {
        v5[0] = (-(v6))
    }
    v3 = (v3 ^ (v3 - (v3 >> ((0) & 63))))
    switch v6 {
    case -918:
        v5 = append(v5, v6)
    case 4294967295:
        v24 = (32767 < 1)
    case -2147483648:
        for fi30 := 0; fi30 < 17; fi30++ {
            switch fi30 {
            case 0:
                v31 := GenS5{F0: v1, F1: ("a" + ("a" + ("hello" + ("" + ("\n\t" + "hello"))))), F2: 2147483647}
                _ = v31
            }
            _ = fi30
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v2[2] = (-(v6))
    v32 := GenS2{F0: 593, F1: 83, F2: -128}
    for fi33 := 0; fi33 < 16; fi33++ {
        v34 := GenS6{F0: false, F1: (v6 % (((^(fi33))) | 1))}
        fmt.Println(v6, v3)
        v35 := GenS7{F0: 610, F1: "", F2: 2}
        _ = fi33
        _ = v34
        _ = v35
    }
    if (442 > 335) {
        _ = len(v5)
        switch v6 {
        case -2147483648:
            v29 = v29
        case -9223372036854775808:
            for fi36 := 0; fi36 < 14; fi36++ {
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int8 = 1
                    _ = fi37
                    _ = v38
                }
                var v39 int16 = int16(v1)
                fmt.Println(v3)
                _ = fi36
                _ = v39
            }
        case 0:
            var v40 uint32 = 42
            _ = v40
        default:
            var v41 bool = (float64(1e10) > 1e10)
            _ = v41
        }
    } else {
        v3 = 763
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v5
    _ = v6
}
