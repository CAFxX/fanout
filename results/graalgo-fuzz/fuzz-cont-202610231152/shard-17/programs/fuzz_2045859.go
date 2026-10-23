package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0() int32 {
    var v5 uint16 = 421
    if true {
        if ((!("\n\t" != "a")) || true) {
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 uint = 0
                v8 := GenS1{F0: (v5 << ((v7) & 63)), F1: "\n\t", F2: (^(v7))}
                v9 := []int{(fi6 * 1), (fi6 | (-(fi6))), (-(fi6))}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 uint = 42
            v11 := []int{128}
            _ = v10
            _ = v11
        } else {
            v5 -= v5
            v12 := GenS1{F0: v5, F1: ("hello" + "\n\t"), F2: 65535}
            _ = v12
        }
        v13 := []int{-9223372036854775808, 395}
        _ = v13
    }
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 uint) int8 {
    var v14 uint32 = 65535
    var v15 int64 = 957
    var v16 int64 = (v15 / ((285) | 1))
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            v15 -= v15
            var v18 uint64 = 211
            switch ("a" + "hello") {
            case "a":
                fmt.Println(p1, v16, v15)
            }
            _ = v18
        }
        v19 := GenS2{F0: true, F1: 65535}
        if false {
            var v20 uint32 = v14
            _ = v20
        } else {
            v21 := GenS1{F0: 32767, F1: "foo bar", F2: (p0 * uint(v14))}
            v22 := GenS1{F0: 1, F1: (("foo bar" + ((("" + "") + "hello") + ((("a" + (("foo bar" + "foo bar") + "a")) + ("hello" + ("a" + "hello"))) + "\n\t"))) + "\n\t"), F2: p0}
            for fi23 := 0; fi23 < 4; fi23++ {
                if (p1 == (^(p1))) {
                    v24 := GenS2{F0: false, F1: (v14 + (^(v14)))}
                    if (!true) {
                        var v25 bool = true
                        var v26 int = (-(fi17))
                        if false {
                            if (float32(1e10) > 1.5) {
                                v27 := GenS2{F0: v25, F1: v14}
                                _ = v27
                            } else {
                                if true {
                                    if true {
                                        if (v25 || v25) {
                                            var v28 uint32 = 931
                                            _ = v28
                                        } else {
                                            if (!true) {
                                                v15 += v16
                                                v15 = (v16 ^ (v15 - int64(fi23)))
                                                v16 = v16
                                            }
                                            if (!(1 != -240)) {
                                                p0 = (p0 | p0)
                                            } else {
                                                v14 -= (v14 >> (((p1 << (((p1 & p0)) & 63))) & 63))
                                                v15 = -9223372036854775808
                                            }
                                            var v29 int64 = v15
                                            _ = v29
                                        }
                                    } else {
                                        if v25 {
                                            v30 := []int{(fi23 ^ -9223372036854775808), (fi23 * fi23), (fi17 | 32767), 4294967295}
                                            _ = v30
                                        }
                                        switch -1 {
                                        case -9223372036854775808:
                                            v31 := GenS3{F0: 42}
                                            _ = v31
                                        default:
                                            switch fi17 {
                                            case 877:
                                                var v32 uint8 = 0
                                                _ = v32
                                            case -265:
                                                v33 := []int{fi17, int(p0)}
                                                _ = v33
                                            default:
                                                v34 := []int{}
                                                _ = v34
                                            }
                                        }
                                    }
                                    p0 -= p1
                                    v16 -= (^(v16))
                                } else {
                                    v35 := GenS1{F0: 343, F1: "a", F2: (p1 & (p1 * p0))}
                                    _ = v35
                                }
                                fmt.Println(v16)
                                v16 = (v16 + (v15 & (v15 % ((v15) | 1))))
                            }
                            v25 = (!true)
                            v36 := GenS2{F0: false, F1: (v14 + v14)}
                            _ = v36
                        }
                        _ = v25
                        _ = v26
                    } else {
                        switch -9223372036854775808 {
                        case 9223372036854775807:
                            fmt.Println(v15)
                        }
                        if (p0 >= 42) {
                            v15 = (v16 / ((0) | 1))
                        } else {
                            if false {
                                var v37 int64 = v16
                                v37 = v16
                                _ = v37
                            }
                            fmt.Println(fi23, p1, v14)
                        }
                    }
                    var v38 int16 = -1
                    _ = v24
                    _ = v38
                }
                _ = fi23
            }
            _ = v21
            _ = v22
        }
        _ = fi17
        _ = v19
    }
    v14 += (v14 * v14)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 float64 = ((-2.25 / 0.1) + float64((0.0 / (0.0 + 1.5))))
    v2 := make(map[int]int)
    v3 := []int{32768, -191, -2147483648}
    v4 := make(map[string]uint32)
    switch ("\n\t" + ("hello" + "")) {
    case "a":
        if true {
            v1 += (float64((((0.0 / 3.14159) * float32((1e10 / 3.14159))) - 0.0)) - 1e10)
        } else {
            switch ("hello" + "hello") {
            case "foo bar":
                switch 32767 {
                case 32767:
                    v1 = v1
                case 348:
                    v1 += -2.25
                }
            }
            for fi39 := 0; fi39 < 7; fi39++ {
                for fi40 := 0; fi40 < 13; fi40++ {
                    fmt.Println(fi39)
                    var v41 uint = 65535
                    v41 = (v41 << ((255) & 63))
                    _ = fi40
                    _ = v41
                }
                fmt.Println(v1)
                fmt.Println(fi39)
                _ = fi39
            }
            fmt.Println(v1)
        }
    case "":
        if len(v3) > 0 {
            v3[0] = -433
        }
    default:
        if false {
            delete(v2, 127)
            for fi42 := 0; fi42 < 11; fi42++ {
                if (18446744073709551615 <= 2) {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        fmt.Println(fi43, fi43)
                        v1 = float64(3.14159)
                        _ = fi43
                    }
                    v44 := GenS3{F0: 4294967295}
                    var v45 int8 = 1
                    _ = v44
                    _ = v45
                } else {
                    v1 += v1
                }
                for fi46 := 0; fi46 < 5; fi46++ {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 32768
                        }
                        v1 = v1
                    } else {
                        var v47 uint32 = 1
                        v48 := GenS1{F0: 255, F1: ("" + "\n\t"), F2: 255}
                        _ = v47
                        _ = v48
                    }
                    v1 += v1
                    _ = fi46
                }
                if false {
                    v49 := GenS4{F0: 127, F1: 50, F2: 2}
                    _ = v49
                }
                _ = fi42
            }
        } else {
            var v50 uint16 = 0
            _ = v50
        }
    }
    if ("foo bar" >= "\n\t") {
        v1 -= ((v1 * v1) - (v1 / v1))
    } else {
        for fi51 := 0; fi51 < 20; fi51++ {
            fmt.Println(v1, v1)
            _ = fi51
        }
        if false {
            var v52 uint16 = 753
            var v53 int = 9223372036854775807
            if true {
                for fi54 := 0; fi54 < 12; fi54++ {
                    _ = len(v4)
                    switch "\n\t" {
                    case "\n\t":
                        v1 = v1
                    case "foo bar":
                        fmt.Println(v1, v52)
                    case "":
                        if false {
                            v53 = v53
                            fmt.Println(v53, fi54, v1)
                        }
                    }
                    v53 = 42
                    _ = fi54
                }
                var v55 bool = true
                if (!false) {
                    fmt.Println(v52)
                    var v56 int64 = int64(v52)
                    v2[v53] = (-(v53))
                    _ = v56
                }
                _ = v55
            } else {
                v57 := GenS1{F0: v52, F1: "hello", F2: 18446744073709551615}
                _ = v57
            }
            _ = v52
            _ = v53
        }
    }
    switch 32768 {
    case -994:
        for fi58 := 0; fi58 < 15; fi58++ {
            v1 = 0.1
            _ = fi58
        }
    case -1:
        var v59 int64 = -1
        _ = v59
    case -9223372036854775808:
        v1 -= ((float64(0.0) * v1) - v1)
    }
    fmt.Println(v1)
    if true {
        if (!false) {
            switch 449 {
            case 896:
                if ((true || false) || false) {
                    v60 := GenS4{F0: 112, F1: 1, F2: 7}
                    if ((true && (42 < 395)) && (0 <= -1)) {
                        var v61 int32 = -1
                        var v62 float64 = (v1 * v1)
                        _ = len(v3)
                        _ = v61
                        _ = v62
                    }
                    _ = v60
                } else {
                    v1 = v1
                    v1 = (float64(1.5) / v1)
                }
            default:
                var v63 uint8 = 127
                _ = v63
            }
        } else {
            v64 := GenS4{F0: 42, F1: 213, F2: 2}
            _ = v64
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if (65535 <= 2147483647) {
            var v65 float32 = (3.14159 + float32(1.5))
            _ = v65
        }
        v1 = 3.14159
        switch -1 {
        case 0:
            var v66 int16 = 1
            _ = v66
        }
    }
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    delete(v2, 675)
    fmt.Println(v1)
    _ = len(v3)
    var v67 uint8 = 2
    v68 := GenS5{F0: 18446744073709551615}
    v1 -= v1
    v67 -= 160
    v69 := GenS1{F0: 259, F1: "a", F2: uint(v67)}
    if false {
        if false {
            var v70 int64 = 9223372036854775807
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v70 += (v70 | v70)
            _ = v70
        } else {
            var v71 uint32 = 904
            _ = v71
        }
        v67 += 1
        v72 := GenS6{F0: v67}
        _ = v72
    }
    var v73 uint = 255
    v73 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v67:", v67)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v69
    _ = v73
}
