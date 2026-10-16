package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int16; F2 uint32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint }

func main() {
    var v1 bool = false
    v2 := []uint32{975, 0, 2}
    var v3 uint8 = 42
    var v4 uint64 = 0
    v5 := make(map[string]bool)
    var v6 int = 32767
    v7 := GenS1{F0: 2147483647, F1: -1}
    if v1 {
        switch v6 {
        case 4294967295:
            if false {
                for fi8 := 0; fi8 < 16; fi8++ {
                    v4 = v4
                    switch "a" {
                    case "\n\t":
                        var v9 int = (fi8 + v6)
                        _ = v9
                    default:
                        switch ("a" + (("a" + "") + "foo bar")) {
                        case "":
                            v4 -= v4
                        }
                    }
                    v10 := GenS1{F0: 787, F1: 202}
                    _ = fi8
                    _ = v10
                }
                _ = len(v5)
            }
        case -2147483648:
            switch (-(v6)) {
            case -403:
                switch v6 {
                case 32767:
                    var v11 bool = (v1 || v1)
                    _ = v11
                case 783:
                    var v12 int32 = 2147483647
                    _ = v12
                case 32768:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        if v1 {
                            switch "foo bar" {
                            case "hello":
                                v1 = (42 != 4294967295)
                            case "":
                                var v14 int = int(v4)
                                _ = v14
                            case "\n\t":
                                if v1 {
                                    v3 += v3
                                } else {
                                    v4 -= (v4 % (((^(v4))) | 1))
                                }
                            default:
                                var v15 uint8 = (v3 >> ((2) & 63))
                                _ = v15
                            }
                            v16 := GenS2{F0: -1}
                            if (-2.25 >= 1.5) {
                                for fi17 := 0; fi17 < 14; fi17++ {
                                    var v18 uint8 = 255
                                    v1 = (((float32(1.5) + (1.5 * float32(-2.25))) + ((((((-2.25 * 0.0) + (1e10 / -2.25)) / (3.14159 + 1.5)) / -2.25) + 1.5) * 0.0)) == 0.0)
                                    var v19 float64 = 0.1
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                v3 -= (^(v3))
                                v2 = append(v2, 0)
                            }
                            _ = v16
                        }
                        _ = fi13
                    }
                }
            default:
                var v20 int8 = 127
                _ = v20
            }
        default:
            v4 -= (v4 | (v4 % (((v4 | 18446744073709551615)) | 1)))
        }
        switch (v6 + v6) {
        case -9223372036854775808:
            var v21 uint16 = 255
            _ = v21
        }
        var v22 int8 = 0
        _ = v22
    } else {
        fmt.Println(v6, v3)
        v23 := GenS1{F0: 65535, F1: int64(v6)}
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS3{F0: (uint(v3) < 171), F1: 477, F2: 65535}
            var v26 uint32 = 2147483647
            if len(v2) > 0 {
                v2[0] = uint32(v3)
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS1{F0: 2, F1: 989}
        if (!(0 <= 4294967295)) {
            fmt.Println(v1, fi27, fi27)
        } else {
            v29 := GenS4{F0: 1, F1: "hello"}
            _ = v29
        }
        if false {
            v30 := GenS2{F0: 0}
            _ = v30
        } else {
            v31 := GenS2{F0: -1}
            fmt.Println(v3, v3)
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    v32 := GenS5{F0: uint32(v4)}
    var v33 uint = 1
    switch "\n\t" {
    case "\n\t":
        var v34 float32 = -2.25
        _ = v34
    case "hello":
        v35 := GenS6{F0: 0, F1: (^(v33))}
        _ = v35
    }
    var v36 int8 = 0
    v4 = (v4 >> (((v33 & (v33 + v33))) & 63))
    var v37 uint8 = v3
    var v38 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
