package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint64; F1 string }

func main() {
    var v1 uint64 = 79
    v2 := make(map[int]bool)
    var v3 int = 42
    v4 := make(map[uint32]uint32)
    var v5 uint = 409
    var v6 uint32 = 77
    v7 := []int{(v3 << ((v5) & 63)), (v3 & v3)}
    for fi8 := 0; fi8 < 10; fi8++ {
        if (true || (true && (0 <= (v6 % ((v6) | 1))))) {
            _ = len(v4)
        } else {
            v3 += 709
            v9 := v4[v6]
            v6 = (^(v6))
            _ = v9
        }
        v2[v3] = false
        var v10 float32 = (((-2.25 / 1.0) / 1.5) * 1.5)
        _ = fi8
        _ = v10
    }
    v1 += 837
    v11 := GenS1{F0: false, F1: ("" + "foo bar")}
    v12 := GenS2{F0: 255}
    switch "" {
    case "a":
        if (2147483648 != uint32(v5)) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS3{F0: 42, F1: (^(v1)), F2: (v1 - v1)}
                if true {
                    v6 -= (v6 ^ (v6 / (((v6 ^ v6)) | 1)))
                    var v15 uint32 = 89
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 float64 = 0.1
                if ((1e10 / 3.14159) >= 0.1) {
                    v6 = v6
                    var v18 uint8 = 239
                    fmt.Println(fi16)
                    _ = v18
                } else {
                    if (true || ((42 >= 9223372036854775807) || false)) {
                        fmt.Println(v17)
                        var v19 int16 = 1
                        var v20 float64 = 1e10
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            v21 := GenS2{F0: 223}
                            var v22 uint16 = 380
                            fmt.Println(v3)
                            _ = v21
                            _ = v22
                        }
                    }
                }
                fmt.Println(fi16, v3)
                _ = fi16
                _ = v17
            }
            if (42 != -128) {
                var v23 int = (v3 >> ((uint(v6)) & 63))
                v24 := GenS4{F0: 1, F1: -109}
                _ = v23
                _ = v24
            } else {
                for fi25 := 0; fi25 < 14; fi25++ {
                    if len(v7) > 0 {
                        v7[0] = 515
                    }
                    _ = fi25
                }
                var v26 uint32 = (v6 * v6)
                _ = v26
            }
        }
    default:
        if (((0.0 * (((0.0 / 1.5) - float64(1.5)) * ((float64(0.1) * 1.5) + 0.0))) != 0.1) && true) {
            var v27 uint = (v5 << ((v5) & 63))
            fmt.Println(v6)
            var v28 uint32 = v6
            _ = v27
            _ = v28
        } else {
            if (((v5 >> (((v5 * uint(v1))) & 63)) < v5) || false) {
                v5 -= v5
                var v29 uint8 = 42
                var v30 string = "a"
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                switch v3 {
                case 42:
                    switch "" {
                    case "\n\t":
                        v32 := v2[(v3 + (fi31 ^ fi31))]
                        _ = v32
                    case "":
                        var v33 int = v3
                        _ = v33
                    case "a":
                        if (-2147483648 != -210) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                    default:
                        if true {
                            if ((66 >= uint8(v5)) || false) {
                                fmt.Println(fi31, v5)
                                v34 := GenS5{F0: 1, F1: ("foo bar" + ("hello" + "hello"))}
                                v1 = v1
                                _ = v34
                            } else {
                                if false {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    if (1 <= -9223372036854775808) {
                                        var v35 float32 = -2.25
                                        v36 := GenS3{F0: 2, F1: (v1 ^ 18446744073709551615), F2: v1}
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    v6 -= 2
                                    for fi37 := 0; fi37 < 9; fi37++ {
                                        v6 += v6
                                        _ = fi37
                                    }
                                    var v38 uint16 = 32767
                                    _ = v38
                                }
                            }
                        }
                    }
                case 1:
                    var v39 int64 = 0
                    _ = v39
                }
                v3 = (^(v3))
                var v40 uint8 = 1
                _ = fi31
                _ = v40
            }
            v6 = uint32(v1)
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        v7 = append(v7, 32767)
        v5 = 2
        _ = fi41
    }
    var v42 uint16 = uint16(v5)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
