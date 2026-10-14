package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int32; F2 bool }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int, p1 uint32, p2 uint64) uint64 {
    var v4 uint16 = 2
    v5 := GenS1{F0: 42, F1: 1, F2: (v4 * (v4 ^ (v4 / ((148) | 1))))}
    var v6 int64 = -9223372036854775808
    switch p0 {
    case -1:
        switch "" {
        case "a":
            v7 := []int{(p0 << ((18446744073709551615) & 63)), 1, (p0 >> ((909) & 63))}
            _ = v7
        }
    case -713:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 bool = ((^(p1)) >= (^(p1)))
            _ = fi8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            if (!true) {
                var v11 bool = ((((!(false || (p1 != p1))) || (42 <= 2)) && false) || true)
                _ = v11
            } else {
                p1 -= uint32(p2)
            }
            if true {
                p0 = fi10
            }
            v6 = v6
            _ = fi10
        }
    }
    v6 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 int16 = 32767
    var v2 bool = ((0.1 / (float32((((0.1 / float64(float32(-2.25))) + 1e10) + 1e10)) + 3.14159)) <= (-2.25 - 3.14159))
    v3 := make(map[int]uint32)
    switch 9223372036854775807 {
    case 21:
        v12 := GenS1{F0: 526, F1: 1, F2: 255}
        _ = v12
    case -1:
        v1 = 1
    case 32767:
        switch 1 {
        case -907:
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 uint = 2
                v1 += v1
                var v15 bool = (((((float32(0.0) * float32(((-2.25 + 1e10) * 3.14159))) * (-2.25 - 0.1)) != (-2.25 - (-2.25 * (1e10 * 0.0)))) || v2) || v2)
                _ = fi13
                _ = v14
                _ = v15
            }
        case 0:
            for fi16 := 0; fi16 < 6; fi16++ {
                switch "" {
                case "hello":
                    v1 = (v1 ^ v1)
                case "":
                    v17 := GenS2{F0: ""}
                    _ = v17
                case "foo bar":
                    if v2 {
                        var v18 bool = (!true)
                        v19 := GenS3{F0: 0, F1: -949, F2: v2}
                        _ = v18
                        _ = v19
                    } else {
                        v3[(fi16 + (fi16 << ((65535) & 63)))] = 42
                    }
                default:
                    v20 := GenS4{F0: v2, F1: (^(fi16))}
                    _ = v20
                }
                _ = fi16
            }
        case 9223372036854775807:
            if (0.0 == 1.5) {
                v21 := []int{128, -2147483648, 42, -1}
                var v22 float64 = 0.0
                var v23 float32 = float32(v22)
                _ = v21
                _ = v22
                _ = v23
            } else {
                delete(v3, 2147483647)
                for fi24 := 0; fi24 < 19; fi24++ {
                    v25 := []int{}
                    var v26 uint = 446
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                if (1 != 42) {
                    var v27 int16 = -1000
                    if ((-1 < 268) || v2) {
                        v27 = 913
                        for fi28 := 0; fi28 < 1; fi28++ {
                            if true {
                                for fi29 := 0; fi29 < 1; fi29++ {
                                    v27 = (v27 / (((-(v27))) | 1))
                                    var v30 uint8 = 1
                                    if v2 {
                                        var v31 uint = 42
                                        _ = v31
                                    } else {
                                        v1 = v27
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                                if (42 < 255) {
                                    if (!((^(fi28)) <= 128)) {
                                        v32 := v3[(fi28 / ((32768) | 1))]
                                        v1 -= v1
                                        switch ("" + "\n\t") {
                                        case "a":
                                            var v33 uint32 = 2
                                            _ = v33
                                        case "hello":
                                            fmt.Println(fi28, v2, fi28)
                                        }
                                        _ = v32
                                    }
                                    v34 := GenS2{F0: ""}
                                    v35 := []int{(fi28 / ((fi28) | 1)), 4294967295}
                                    _ = v34
                                    _ = v35
                                }
                            }
                            v36 := GenS2{F0: "foo bar"}
                            _ = fi28
                            _ = v36
                        }
                    } else {
                        if (!(356 <= 4294967295)) {
                            fmt.Println(v1, v1)
                            for fi37 := 0; fi37 < 15; fi37++ {
                                v38 := GenS2{F0: ""}
                                var v39 float32 = -2.25
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                        v1 -= v27
                    }
                    for fi40 := 0; fi40 < 5; fi40++ {
                        if (3.14159 != 1.5) {
                            var v41 float32 = (1e10 - (float32(3.14159) + -2.25))
                            _ = v41
                        }
                        _ = fi40
                    }
                    _ = v27
                }
            }
        default:
            v2 = (3.14159 <= ((0.1 / -2.25) / (0.1 + (0.0 * 0.0))))
        }
    }
    v42 := []int{9223372036854775807, -488, -2147483648}
    delete(v3, 127)
    switch -2147483648 {
    case -30:
        v1 = -184
    case -176:
        if (48 == 0) {
            v2 = v2
            v43 := GenS5{F0: 0}
            v2 = v2
            _ = v43
        } else {
            for fi44 := 0; fi44 < 15; fi44++ {
                v3[fi44] = 4294967295
                var v45 bool = v2
                _ = fi44
                _ = v45
            }
            v46 := GenS5{F0: 119}
            switch 992 {
            case 9223372036854775807:
                for fi47 := 0; fi47 < 17; fi47++ {
                    delete(v3, fi47)
                    v2 = v2
                    v3[(fi47 / ((1) | 1))] = 2147483648
                    _ = fi47
                }
            default:
                var v48 uint16 = 255
                _ = v48
            }
            _ = v46
        }
    default:
        fmt.Println(v2, v1)
    }
    if (!true) {
        _ = len(v42)
        for fi49 := 0; fi49 < 8; fi49++ {
            var v50 bool = false
            _ = fi49
            _ = v50
        }
    }
    if (-1 >= -87) {
        v2 = ((255 <= 2) && (((65535 < 42) && true) || v2))
        _ = len(v3)
    }
    var v51 uint = 264
    v52 := GenS3{F0: 42, F1: -1, F2: (((305 != 632) || (-2147483648 != 9223372036854775807)) || (false && (!v2)))}
    if (-9223372036854775808 <= -719) {
        var v53 uint64 = 42
        _ = v53
    } else {
        var v54 uint = (v51 - (v51 ^ v51))
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v51
    _ = v52
}
