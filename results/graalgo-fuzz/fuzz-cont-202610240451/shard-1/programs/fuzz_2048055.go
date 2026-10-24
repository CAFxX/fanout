package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint) uint8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 255, F1: 1000, F2: 248}
        v7 := []int{(fi5 % ((128) | 1))}
        var v8 int64 = 9223372036854775807
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if (682 >= -148) {
        var v9 int64 = 42
        v9 = -9223372036854775808
        _ = v9
    } else {
        if (2147483647 != int32(p0)) {
            v10 := []int{135, -761}
            p0 -= 42
            _ = v10
        }
        v11 := []int{-191}
        v12 := GenS2{F0: -148}
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "hello":
        var v13 int64 = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(fi14, p0)
        _ = fi14
    }
    _ = p0
    return 128
}
func genFn1(p0 int16, p1 uint64) uint32 {
    v15 := GenS1{F0: 1, F1: p0, F2: 65535}
    v16 := GenS3{F0: 492, F1: false}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 987
}

func main() {
    var v1 uint32 = 212
    v2 := []string{"a", "\n\t", "foo bar"}
    var v3 uint8 = 1
    var v4 uint16 = 255
    for fi17 := 0; fi17 < 10; fi17++ {
        _ = len(v2)
        fmt.Println(v1)
        v18 := GenS2{F0: -1}
        _ = fi17
        _ = v18
    }
    var v19 int32 = -535
    var v20 float64 = -2.25
    if (1e10 < -2.25) {
        switch 885 {
        case -9223372036854775808:
            fmt.Println(v3, v3, v19)
        default:
            var v21 uint16 = (^(v4))
            _ = v21
        }
        switch "foo bar" {
        case "\n\t":
            v22 := GenS4{F0: v1, F1: genFn0(18446744073709551615)}
            _ = v22
        default:
            v4 += (v4 * (v4 ^ v4))
        }
    }
    if (((-1 < -55) || true) || true) {
        v23 := GenS5{F0: false}
        switch 32767 {
        case 32768:
            if (282 > 1000) {
                fmt.Println(v4, v20, v1)
            }
        case -1:
            for fi24 := 0; fi24 < 19; fi24++ {
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi25 := 0; fi25 < 4; fi25++ {
                        v3 = (v3 + (v3 | v3))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v26 uint32 = genFn1(1, uint64(v4))
                        _ = fi25
                        _ = v26
                    }
                    if (0 >= v4) {
                        switch 4294967295 {
                        case 32767:
                            v1 = (v1 ^ v1)
                        default:
                            var v27 bool = ((genFn1(1, uint64(v1)) != v1) && (35 < -1))
                            _ = v27
                        }
                        var v28 int = fi24
                        if true {
                            switch -160 {
                            case 4294967295:
                                var v29 bool = ((1 == -1) && true)
                                _ = v29
                            default:
                                v28 = v28
                            }
                        } else {
                            if (false && true) {
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v31 float32 = (3.14159 - -2.25)
                                        _ = v31
                                    case "":
                                        v1 += 65535
                                    default:
                                        var v32 uint32 = v1
                                        _ = v32
                                    }
                                    _ = fi30
                                }
                                var v33 uint8 = genFn0(1)
                                v2 = append(v2, "")
                                _ = v33
                            }
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "")
                            }
                            v4 = v4
                        }
                        _ = v28
                    }
                }
                _ = fi24
            }
        case 759:
            for fi34 := 0; fi34 < 7; fi34++ {
                for fi35 := 0; fi35 < 1; fi35++ {
                    v36 := GenS1{F0: (v4 >> ((255) & 63)), F1: -1, F2: (^(v4))}
                    fmt.Println(v20)
                    var v37 float64 = v20
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = fi34
            }
        }
        _ = v23
    }
    v38 := GenS2{F0: 351}
    if (false || false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
        v4 -= (^(v4))
    } else {
        v2 = append(v2, "a")
        var v39 int64 = -9223372036854775808
        _ = v39
    }
    v40 := GenS2{F0: int16(v4)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v40
}
