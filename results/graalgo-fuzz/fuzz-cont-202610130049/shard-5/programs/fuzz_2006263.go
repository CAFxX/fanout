package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int) int {
    p0 += 9223372036854775807
    p0 -= (^(p0))
    fmt.Println(p0)
    v6 := GenS1{F0: 2147483648, F1: 9223372036854775807, F2: -794}
    _ = p0
    _ = v6
    return p0
}
func genFn1(p0 int16) uint {
    var v7 uint64 = 639
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{genFn0(fi9), (^(fi8)), -841}
            v11 := GenS2{F0: true, F1: 2, F2: 2}
            switch (fi8 ^ fi9) {
            case 4294967295:
                var v12 float64 = 1e10
                _ = v12
            case 270:
                v7 = 823
            }
            var v13 int64 = -53
            _ = fi9
            _ = v10
            _ = v11
            _ = v13
        }
        _ = fi8
    }
    v14 := []int{32767}
    v15 := GenS1{F0: 808, F1: 887, F2: 681}
    _ = p0
    _ = v14
    _ = v15
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 42
    v2 := []string{("\n\t" + "\n\t"), ((("a" + "") + "foo bar") + ("foo bar" + "")), ("" + ("a" + "\n\t"))}
    var v3 string = "a"
    var v4 uint32 = 337
    var v5 int = -514
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = (v5 + 0)
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            switch 0 {
            case 1:
                v3 = "hello"
            case -9223372036854775808:
                v1 = (v1 + v1)
                var v18 float32 = 0.0
                _ = v18
            }
            if false {
                var v19 bool = (v1 == 0)
                _ = v19
            } else {
                var v20 uint64 = 18446744073709551615
                _ = v20
            }
            var v21 float32 = -2.25
            _ = fi17
            _ = v21
        }
        if (0.1 == float64(-2.25)) {
            var v22 int16 = (v1 & v1)
            _ = v22
        }
        fmt.Println(v5, v3, v1)
        _ = fi16
    }
    if true {
        v5 -= v5
        fmt.Println(v3)
        fmt.Println(v3, v4, v5)
    }
    v23 := GenS3{F0: (v1 / (((v1 << ((291) & 63))) | 1)), F1: v1, F2: 1}
    for fi24 := 0; fi24 < 4; fi24++ {
        switch (-(fi24)) {
        case 42:
            if ((3.14159 / 1.5) != ((-2.25 / 1.0) * 0.1)) {
                v25 := GenS4{F0: -2147483648}
                if len(v2) > 0 {
                    v2[0] = ("" + v3)
                }
                if true {
                    var v26 int8 = -39
                    v27 := GenS4{F0: 2147483647}
                    _ = v26
                    _ = v27
                }
                _ = v25
            } else {
                if true {
                    v28 := GenS4{F0: 94}
                    for fi29 := 0; fi29 < 4; fi29++ {
                        if true {
                            if false {
                                fmt.Println(v4)
                                var v30 uint8 = 1
                                _ = v30
                            }
                        }
                        _ = fi29
                    }
                    _ = v28
                }
                if false {
                    var v31 int64 = -757
                    if false {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        if (0 >= 32767) {
                            if false {
                                v31 -= 935
                                v32 := GenS1{F0: (v4 % (((^(v4))) | 1)), F1: (-(v31)), F2: 68}
                                _ = len(v2)
                                _ = v32
                            }
                            var v33 uint8 = 21
                            _ = v33
                        } else {
                            switch genFn0((fi24 & (-(fi24)))) {
                            case 127:
                                var v34 int = fi24
                                _ = v34
                            case 230:
                                var v35 float32 = 0.0
                                _ = v35
                            case -2147483648:
                                var v36 uint32 = v4
                                _ = v36
                            default:
                                for fi37 := 0; fi37 < 19; fi37++ {
                                    if (true && (v5 > (fi37 - (v5 << ((4294967295) & 63))))) {
                                        _ = len(v2)
                                        if len(v2) > 0 {
                                            v2[0] = v3
                                        }
                                    } else {
                                        var v38 uint16 = 42
                                        if (3.14159 != 0.0) {
                                            var v39 uint32 = v4
                                            var v40 uint16 = v38
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v41 := GenS4{F0: -890}
                                            v42 := GenS4{F0: 42}
                                            v1 -= 1000
                                            _ = v41
                                            _ = v42
                                        }
                                        fmt.Println(fi37)
                                        _ = v38
                                    }
                                    _ = fi37
                                }
                            }
                        }
                        v1 -= 319
                    }
                    var v43 int = (-(v5))
                    _ = v31
                    _ = v43
                }
                var v44 float32 = 0.1
                _ = v44
            }
        }
        _ = fi24
    }
    switch genFn0(-656) {
    case 128:
        if true {
            var v45 uint32 = 2147483648
            v45 = v45
            _ = v45
        }
    case 2147483647:
        var v46 uint32 = 737
        _ = v46
    }
    if ((v5 + 32768) == v5) {
        if false {
            switch (v3 + "foo bar") {
            case "hello":
                fmt.Println(v5, v4, v3)
            case "a":
                _ = len(v2)
            default:
                if true {
                    for fi47 := 0; fi47 < 3; fi47++ {
                        if ((-32768 == (v1 >> ((135) & 63))) && false) {
                            v5 = genFn0(fi47)
                        }
                        for fi48 := 0; fi48 < 8; fi48++ {
                            v49 := GenS5{F0: int8(v4), F1: true}
                            v50 := GenS1{F0: (v4 % ((v4) | 1)), F1: 42, F2: 1}
                            if true {
                                var v51 int16 = (v1 % (((v1 & v1)) | 1))
                                if (false && (true && false)) {
                                    var v52 int8 = 70
                                    switch "foo bar" {
                                    case "":
                                        var v53 int = (-(v5))
                                        if false {
                                            v54 := GenS2{F0: (v3 <= "hello"), F1: 2, F2: 260}
                                            switch (v3 + ("\n\t" + v3)) {
                                            case "\n\t":
                                                v5 -= fi47
                                            case "a":
                                                v55 := GenS2{F0: false, F1: 18446744073709551615, F2: 42}
                                                _ = v55
                                            }
                                            v56 := GenS5{F0: (v52 >> ((42) & 63)), F1: (!(true || (v3 != v3)))}
                                            _ = v54
                                            _ = v56
                                        }
                                        if (0.1 < 0.1) {
                                            var v57 uint32 = 2
                                            fmt.Println(v57, v52)
                                            var v58 int32 = 1
                                            _ = v57
                                            _ = v58
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (!(-2147483648 == -1)) {
                                                var v59 uint32 = (v4 * 761)
                                                var v60 int = 128
                                                _ = v59
                                                _ = v60
                                            } else {
                                                var v61 uint64 = 254
                                                _ = v61
                                            }
                                        }
                                        _ = v53
                                    default:
                                        if ((!false) || false) {
                                            v1 = (v1 & (v51 % (((-(v1))) | 1)))
                                            var v62 int16 = v1
                                            _ = v62
                                        }
                                    }
                                    _ = v52
                                } else {
                                    _ = len(v2)
                                    _ = len(v2)
                                    v63 := GenS5{F0: -99, F1: (true || true)}
                                    if (2 == 0) {
                                        var v64 uint32 = 222
                                        _ = v64
                                    } else {
                                        fmt.Println(v3)
                                    }
                                    _ = v63
                                }
                                v65 := GenS5{F0: -29, F1: (18446744073709551615 == 42)}
                                _ = v51
                                _ = v65
                            }
                            if false {
                                v4 = (v4 / ((v4) | 1))
                                v1 += v1
                                v4 = (v4 / ((331) | 1))
                            } else {
                                v5 = genFn0(fi48)
                                var v66 string = "a"
                                v2 = append(v2, v66)
                                _ = v66
                            }
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = fi47
                    }
                    var v67 float32 = -2.25
                    _ = len(v2)
                    _ = v67
                } else {
                    v68 := GenS1{F0: 0, F1: 277, F2: 958}
                    _ = v68
                }
            }
            v69 := GenS6{F0: -104}
            _ = v69
        }
        for fi70 := 0; fi70 < 3; fi70++ {
            v1 = (v1 * v1)
            v71 := GenS7{F0: 2}
            if false {
                if (!true) {
                    var v72 int32 = -342
                    fmt.Println(v4)
                    _ = v72
                }
            } else {
                var v73 string = ("" + "\n\t")
                v2 = append(v2, "a")
                _ = v73
            }
            _ = fi70
            _ = v71
        }
        fmt.Println(v1, v1, v1)
    }
    var v74 bool = false
    v74 = v74
    v4 = (v4 << ((255) & 63))
    v1 = v1
    for fi75 := 0; fi75 < 5; fi75++ {
        for fi76 := 0; fi76 < 2; fi76++ {
            v1 -= v1
            v74 = v74
            _ = fi76
        }
        if (9223372036854775807 == fi75) {
            var v77 uint64 = 294
            if (!(uint16(v4) < 42)) {
                v5 = (^(v5))
            }
            var v78 uint16 = 32767
            _ = v77
            _ = v78
        } else {
            switch v3 {
            case "foo bar":
                for fi79 := 0; fi79 < 6; fi79++ {
                    var v80 int64 = 1
                    switch int(v1) {
                    case -9223372036854775808:
                        var v81 int8 = 100
                        _ = v81
                    default:
                        if len(v2) > 0 {
                            v2[0] = (v3 + v3)
                        }
                    }
                    _ = fi79
                    _ = v80
                }
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            case "":
                v5 = 42
            default:
                if (!(false || v74)) {
                    var v82 int = (fi75 / ((v5) | 1))
                    _ = v82
                }
            }
            v4 = v4
        }
        var v83 uint16 = 236
        _ = fi75
        _ = v83
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v74
}
