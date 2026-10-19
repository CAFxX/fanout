package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 uint64) bool {
    fmt.Println(p0)
    v6 := []int{}
    v7 := GenS1{F0: 32767, F1: -937, F2: false}
    _ = p0
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int64) uint {
    var v8 uint32 = 42
    fmt.Println(v8, p0)
    p0 -= p0
    var v9 uint64 = 0
    fmt.Println(p0)
    _ = p0
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint = 948
    v2 := make(map[uint32]uint32)
    var v3 uint = v1
    var v4 float32 = 0.1
    var v5 int16 = 1
    var v10 int = -964
    if true {
        if true {
            var v11 int16 = v5
            _ = v11
        }
        var v12 uint64 = 4294967295
        var v13 uint32 = 2147483647
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: (v10 ^ v10), F1: 1, F2: (65535 != 254)}
        _ = v14
    }
    v10 -= (v10 ^ v10)
    var v15 float32 = v4
    var v16 uint64 = 18446744073709551615
    v3 -= (v1 | (v3 - 684))
    v15 = float32((-2.25 / 1e10))
    v3 += v3
    _ = len(v2)
    fmt.Println(v16, v15)
    if ((1 <= -1) || (!(32767 >= uint16(v16)))) {
        v17 := GenS2{F0: 42, F1: (true && (((!(false && (genFn1(-9223372036854775808) <= genFn1(-1)))) || (("foo bar" + ("" + "\n\t")) >= "hello")) || (((!((v3 + (v1 & genFn1(-9223372036854775808))) <= genFn1(523))) && false) || (42 <= 299))))}
        if (float64(float32((0.0 + 1e10))) > float64(v15)) {
            if true {
                if false {
                    v1 -= (v3 ^ (^(v3)))
                }
                v18 := GenS3{F0: (true || false)}
                _ = v18
            }
        }
        _ = v17
    }
    if false {
        if (("a" != "") && false) {
            v5 += -32768
            for fi19 := 0; fi19 < 17; fi19++ {
                fmt.Println(v16, v1)
                var v20 uint32 = 42
                if (false || true) {
                    var v21 uint = v1
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            if false {
                if (v10 != (v10 / (((v10 / ((42) | 1))) | 1))) {
                    v22 := GenS3{F0: true}
                    if (float32((-2.25 + float64(0.1))) == ((float32(3.14159) - (v4 - 1.5)) * (v4 * v4))) {
                        if true {
                            v16 = v16
                        }
                    } else {
                        v23 := GenS3{F0: (!false)}
                        if ("" <= (("a" + "foo bar") + ("" + ("\n\t" + ("foo bar" + "a"))))) {
                            fmt.Println(v15, v4, v3)
                            for fi24 := 0; fi24 < 16; fi24++ {
                                var v25 uint16 = 42
                                _ = fi24
                                _ = v25
                            }
                            v5 += v5
                        }
                        if false {
                            for fi26 := 0; fi26 < 17; fi26++ {
                                v3 = 630
                                var v27 uint64 = 351
                                _ = fi26
                                _ = v27
                            }
                            var v28 float32 = 1.5
                            switch (v10 % ((4294967295) | 1)) {
                            case -675:
                                v15 -= v4
                            case 1:
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    var v30 int8 = -1
                                    v30 += 0
                                    _ = fi29
                                    _ = v30
                                }
                            case 127:
                                v5 = v5
                            default:
                                v31 := []int{v10, 32768, (v10 << ((520) & 63)), (-(v10))}
                                _ = v31
                            }
                            _ = v28
                        }
                        _ = v23
                    }
                    _ = v22
                } else {
                    delete(v2, 65535)
                    if true {
                        delete(v2, 42)
                        if false {
                            v32 := v2[2]
                            for fi33 := 0; fi33 < 13; fi33++ {
                                v3 += 4294967295
                                var v34 bool = false
                                v4 += v4
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        } else {
                            if false {
                                for fi35 := 0; fi35 < 18; fi35++ {
                                    if false {
                                        fmt.Println(v10, v10)
                                        if (((v10 % ((v10) | 1)) <= 127) && (!false)) {
                                            var v36 float64 = ((((0.1 + 1e10) + -2.25) - 1e10) * 1e10)
                                            var v37 int64 = 101
                                            _ = v36
                                            _ = v37
                                        }
                                        for fi38 := 0; fi38 < 16; fi38++ {
                                            var v39 int = (v10 & v10)
                                            _ = len(v2)
                                            if (4294967295 > (v16 & 18446744073709551615)) {
                                                v40 := GenS4{F0: "\n\t", F1: 255}
                                                _ = v40
                                            }
                                            _ = fi38
                                            _ = v39
                                        }
                                    }
                                    _ = fi35
                                }
                                v2[155] = 2147483648
                            } else {
                                for fi41 := 0; fi41 < 17; fi41++ {
                                    v42 := GenS5{F0: 0, F1: false, F2: 2}
                                    var v43 uint8 = 1
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                                v1 -= 557
                            }
                        }
                        var v44 int64 = 1
                        _ = v44
                    }
                }
            }
        }
        v45 := []int{}
        _ = v45
    }
    _ = len(v2)
    if false {
        v15 = 0.1
        v5 = 473
        var v46 int8 = 100
        _ = v46
    } else {
        switch (v10 % (((v10 & v10)) | 1)) {
        case 42:
            for fi47 := 0; fi47 < 16; fi47++ {
                if (42 == -2147483648) {
                    v1 += (v3 >> (((v3 << (((v3 * v1)) & 63))) & 63))
                    v15 = v15
                    if false {
                        v48 := GenS6{F0: 516}
                        _ = v48
                    }
                }
                v16 = v16
                v49 := v2[863]
                _ = fi47
                _ = v49
            }
        case 127:
            delete(v2, 2147483647)
        case 9223372036854775807:
            if true {
                for fi50 := 0; fi50 < 16; fi50++ {
                    for fi51 := 0; fi51 < 11; fi51++ {
                        fmt.Println(v16, v15, v4)
                        _ = fi51
                    }
                    _ = fi50
                }
                v15 += v4
            }
        default:
            v3 += v1
        }
    }
    var v52 int8 = 0
    if false {
        v1 += (v3 % (((v1 << (((v1 * 65535)) & 63))) | 1))
        if (false || ((int64(v5) >= 42) && true)) {
            if false {
                fmt.Println(v3, v52)
                for fi53 := 0; fi53 < 6; fi53++ {
                    if true {
                        var v54 int = v10
                        for fi55 := 0; fi55 < 11; fi55++ {
                            v56 := v2[2]
                            v52 = (v52 / ((-128) | 1))
                            _ = fi55
                            _ = v56
                        }
                        var v57 float64 = -2.25
                        _ = v54
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 17; fi58++ {
                        var v59 uint64 = v16
                        var v60 int16 = v5
                        var v61 float32 = float32(0.1)
                        _ = fi58
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    v62 := GenS5{F0: 42, F1: false, F2: uint8(v16)}
                    _ = fi53
                    _ = v62
                }
                v63 := GenS7{F0: v5, F1: 1}
                _ = v63
            } else {
                v64 := []int{v10, (v10 % ((42) | 1))}
                if (!(9223372036854775807 != -1)) {
                    v65 := GenS7{F0: v5, F1: 2}
                    v2[2147483648] = 31
                    var v66 int64 = -1
                    _ = v65
                    _ = v66
                }
                v4 = float32((3.14159 - -2.25))
                _ = v64
            }
            v67 := []int{(-(v10)), v10, v10, 2147483647}
            var v68 int64 = -214
            _ = v67
            _ = v68
        }
    } else {
        switch "hello" {
        case "":
            for fi69 := 0; fi69 < 12; fi69++ {
                v70 := GenS2{F0: 0, F1: true}
                v71 := GenS8{F0: (v5 / (((^(v5))) | 1))}
                switch 32767 {
                case 127:
                    if false {
                        if false {
                            v5 -= 1000
                            v72 := []int{(-(v10)), (fi69 + (v10 % (((v10 & fi69)) | 1))), fi69}
                            _ = v72
                        } else {
                            v73 := []int{int(v16)}
                            v74 := GenS9{F0: 1}
                            _ = v73
                            _ = v74
                        }
                        var v75 uint = genFn1(42)
                        v5 -= v5
                        _ = v75
                    }
                case -1:
                    v1 = (v1 << (((v3 / ((genFn1(9223372036854775807)) | 1))) & 63))
                }
                _ = fi69
                _ = v70
                _ = v71
            }
        case "hello":
            v52 = 1
        case "\n\t":
            v15 -= (v4 * v4)
        default:
            v15 = 1e10
        }
        for fi76 := 0; fi76 < 19; fi76++ {
            if (!false) {
                var v77 bool = true
                _ = v77
            } else {
                v5 += -218
                v15 = (v4 - v15)
            }
            v78 := GenS6{F0: -1}
            v3 = 4294967295
            _ = fi76
            _ = v78
        }
    }
    var v79 uint64 = (v16 << (((v1 | (^(v3)))) & 63))
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v79
}
