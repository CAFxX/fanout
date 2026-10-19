package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 string; F1 string; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int64; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS8 struct { F0 int; F1 uint64; F2 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS11 struct { F0 uint8; F1 uint32 }
type GenS12 struct { F0 uint32; F1 int8; F2 int32 }
type GenS13 struct { F0 uint8; F1 uint8 }
type GenS14 struct { F0 uint32; F1 int8 }
func genFn0(p0 int8, p1 uint, p2 int) uint16 {
    p1 -= p1
    v5 := GenS1{F0: (p0 ^ int8(p2)), F1: 2147483647}
    var v6 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1() int {
    var v7 int8 = -1
    var v8 int8 = 1
    _ = v7
    _ = v8
    return 325
}

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[uint32]bool)
    var v3 int8 = 1
    v4 := make(map[uint32]bool)
    v9 := GenS1{F0: (v3 & v3), F1: 4294967295}
    var v10 float64 = (1e10 + float64(-2.25))
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS2{F0: 1}
        _ = fi11
        _ = v12
    }
    v13 := []int{9223372036854775807, 32767, 42, 42}
    v3 = (v3 & 70)
    for fi14 := 0; fi14 < 12; fi14++ {
        fmt.Println(fi14, v10)
        if (928 < -9223372036854775808) {
            if true {
                var v15 uint16 = 42
                if true {
                    v15 += v15
                    for fi16 := 0; fi16 < 20; fi16++ {
                        if ((false || true) && false) {
                            switch fi14 {
                            case 32:
                                v17 := GenS3{F0: 0, F1: (fi14 ^ fi14)}
                                _ = v17
                            case 32768:
                                v18 := GenS2{F0: v15}
                                _ = v18
                            case -1:
                                v19 := GenS4{F0: "foo bar", F1: "foo bar", F2: 255}
                                _ = v19
                            default:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            }
                            fmt.Println(v10, v10)
                            v20 := GenS5{F0: 127, F1: int64(v15), F2: 891}
                            _ = v20
                        } else {
                            if (!((2147483647 == 128) || false)) {
                                var v21 int8 = (v3 % ((v3) | 1))
                                _ = v21
                            } else {
                                switch ("a" + ("a" + "a")) {
                                case "a":
                                    var v22 uint64 = 0
                                    v23 := GenS2{F0: v15}
                                    v10 -= float64(-2.25)
                                    _ = v22
                                    _ = v23
                                case "\n\t":
                                    var v24 uint32 = 2147483647
                                    _ = v24
                                default:
                                    switch fi16 {
                                    case 554:
                                        delete(v2, 42)
                                    case 436:
                                        v2[416] = (734 != -1)
                                    case -9223372036854775808:
                                        if (65535 > 2147483647) {
                                            switch -9223372036854775808 {
                                            case 920:
                                                var v25 string = "foo bar"
                                                _ = v25
                                            }
                                            switch fi16 {
                                            case -1:
                                                var v26 uint8 = 127
                                                _ = v26
                                            case 1:
                                                v27 := GenS6{F0: -332}
                                                _ = v27
                                            }
                                        } else {
                                            switch -1 {
                                            case 478:
                                                var v28 int = (fi14 << ((747) & 63))
                                                _ = v28
                                            default:
                                                _ = len(v13)
                                            }
                                        }
                                    default:
                                        var v29 uint32 = 65535
                                        _ = v29
                                    }
                                }
                                switch ((("hello" + "foo bar") + "\n\t") + "a") {
                                case "":
                                    if ((v15 << ((2) & 63)) == v15) {
                                        if ("foo bar" != "\n\t") {
                                            v30 := GenS3{F0: 4294967295, F1: (fi14 + (-(fi16)))}
                                            var v31 uint = 1
                                            fmt.Println(v3)
                                            _ = v30
                                            _ = v31
                                        } else {
                                            switch "a" {
                                            case "":
                                                var v32 int16 = -1
                                                _ = v32
                                            case "a":
                                                var v33 uint32 = 42
                                                _ = v33
                                            case "\n\t":
                                                if false {
                                                    fmt.Println(v3, v1)
                                                } else {
                                                    fmt.Println(v15)
                                                    v10 += v10
                                                }
                                            }
                                            if true {
                                                var v34 uint = 968
                                                v10 -= float64((float32(v10) - 0.0))
                                                _ = v34
                                            }
                                            v1 -= 3.14159
                                        }
                                        _ = len(v13)
                                    } else {
                                        v35 := GenS1{F0: (^(v3)), F1: 2}
                                        switch (fi16 * genFn1()) {
                                        case 2147483647:
                                            v13 = append(v13, (^(fi14)))
                                        case -950:
                                            v36 := GenS4{F0: "", F1: "", F2: v15}
                                            _ = v36
                                        case 9223372036854775807:
                                            if (!false) {
                                                v3 = v3
                                            }
                                        }
                                        fmt.Println(v15, fi16)
                                        _ = v35
                                    }
                                case "\n\t":
                                    v37 := GenS6{F0: 42}
                                    _ = v37
                                case "foo bar":
                                    v10 = 0.1
                                }
                                v38 := GenS7{F0: 2, F1: 42, F2: 0}
                                v4[0] = (-2.25 == (float32(float64((v1 * -2.25))) / -2.25))
                                _ = v38
                            }
                            fmt.Println(v3, fi14)
                            if false {
                                if true {
                                    v13 = append(v13, (fi14 << ((18446744073709551615) & 63)))
                                } else {
                                    var v39 int16 = 32767
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v1)
                                    case "\n\t":
                                        var v40 float64 = float64(0.0)
                                        _ = v40
                                    case "foo bar":
                                        switch "" {
                                        case "":
                                            v41 := GenS4{F0: "", F1: ("foo bar" + ""), F2: 1}
                                            _ = v41
                                        case "foo bar":
                                            _ = len(v2)
                                        case "\n\t":
                                            switch (fi14 ^ 4294967295) {
                                            case -184:
                                                if true {
                                                    var v42 uint = 799
                                                    _ = v42
                                                } else {
                                                    if (false || (("" + "foo bar") >= "a")) {
                                                        switch fi16 {
                                                        case -569:
                                                            v10 -= (v10 + v10)
                                                        case 1:
                                                            v3 = 44
                                                            var v43 uint64 = 2
                                                            _ = v43
                                                        }
                                                    }
                                                    v44 := GenS8{F0: 9223372036854775807, F1: 426, F2: 0}
                                                    v45 := GenS9{F0: 4294967295}
                                                    _ = v44
                                                    _ = v45
                                                }
                                            case 2147483647:
                                                var v46 int64 = 0
                                                _ = v46
                                            case 127:
                                                var v47 int8 = (v3 ^ (v3 % (((v3 | v3)) | 1)))
                                                _ = v47
                                            }
                                        default:
                                            v1 = v1
                                        }
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            var v48 int32 = 42
                                            _ = v48
                                        case "hello":
                                            if false {
                                                v15 -= (v15 | (^(v15)))
                                                var v49 uint16 = genFn0(v3, 297, fi14)
                                                v1 = v1
                                                _ = v49
                                            } else {
                                                _ = len(v4)
                                                var v50 uint8 = 1
                                                v50 = v50
                                                _ = v50
                                            }
                                        case "a":
                                            switch genFn1() {
                                            case -743:
                                                v3 += 100
                                                if true {
                                                    v13 = append(v13, fi14)
                                                }
                                                v1 = (((v1 / v1) - ((v1 / v1) / float32((v10 - ((3.14159 + v10) / (float64(float32(v10)) * float64((v1 / v1)))))))) + v1)
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                    }
                                    if (869 == 9223372036854775807) {
                                        v51 := GenS9{F0: 39}
                                        _ = v51
                                    } else {
                                        v52 := GenS3{F0: 2147483647, F1: (fi16 + (fi16 % ((genFn1()) | 1)))}
                                        _ = v52
                                    }
                                    _ = v39
                                }
                                v53 := GenS10{F0: 0, F1: v3, F2: v15}
                                v10 = v10
                                _ = v53
                            } else {
                                v2[1] = (!false)
                                v3 = (-(v3))
                            }
                        }
                        if len(v13) > 0 {
                            v13[0] = -9223372036854775808
                        }
                        fmt.Println(fi14, v3, fi14)
                        _ = fi16
                    }
                    if true {
                        for fi54 := 0; fi54 < 17; fi54++ {
                            if len(v13) > 0 {
                                v13[0] = genFn1()
                            }
                            v4[1] = true
                            v10 += ((3.14159 + 0.0) + (v10 + (float64(((0.0 * (v1 * (((v1 / (v1 * v1)) + 0.1) * -2.25))) + 1.5)) + (v10 / (1e10 / v10)))))
                            _ = fi54
                        }
                    } else {
                        if true {
                            v15 -= (v15 % ((1) | 1))
                            var v55 int = (fi14 - (-(fi14)))
                            v56 := GenS11{F0: 42, F1: 2147483647}
                            _ = v55
                            _ = v56
                        } else {
                            v1 -= 0.0
                        }
                        v3 = v3
                    }
                } else {
                    v57 := GenS8{F0: 128, F1: 1, F2: 0}
                    v58 := GenS12{F0: 0, F1: -36, F2: -1}
                    _ = v57
                    _ = v58
                }
                for fi59 := 0; fi59 < 12; fi59++ {
                    if true {
                        if false {
                            switch "hello" {
                            case "a":
                                v15 = (v15 % ((v15) | 1))
                            case "foo bar":
                                v15 -= (^(v15))
                            case "\n\t":
                                if (725 != (^(v15))) {
                                    fmt.Println(v3, fi59)
                                    var v60 bool = (65535 >= 157)
                                    v60 = (355 != 0)
                                    _ = v60
                                } else {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                }
                            }
                            v10 -= (1.5 * v10)
                            v13 = append(v13, (fi59 >> ((1) & 63)))
                            v3 += v3
                        } else {
                            v61 := GenS13{F0: 1, F1: 122}
                            switch (fi14 / ((fi14) | 1)) {
                            case -603:
                                var v62 bool = (!false)
                                _ = v62
                            case -1:
                                var v63 uint16 = (^(v15))
                                _ = v63
                            case -152:
                                if false {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                }
                            }
                            var v64 int8 = (v3 * v3)
                            _ = v61
                            _ = v64
                        }
                    } else {
                        switch ("hello" + "a") {
                        case "":
                            var v65 int32 = 0
                            _ = v65
                        default:
                            if (255 != 128) {
                                v66 := GenS14{F0: 2147483647, F1: (-(v3))}
                                _ = v66
                            } else {
                                _ = len(v13)
                                fmt.Println(v1)
                            }
                        }
                        var v67 float32 = v1
                        _ = v67
                    }
                    _ = fi59
                }
                _ = v15
            }
        }
        switch fi14 {
        case 9223372036854775807:
            for fi68 := 0; fi68 < 20; fi68++ {
                v10 += 3.14159
                var v69 uint32 = 2147483647
                _ = fi68
                _ = v69
            }
        case 127:
            v3 += v3
        case 858:
            var v70 string = ("\n\t" + "foo bar")
            _ = v70
        default:
            v71 := v2[2]
            _ = v71
        }
        _ = fi14
    }
    _ = len(v13)
    v4[1] = true
    fmt.Println(v1)
    var v72 int64 = 42
    _ = len(v2)
    v73 := GenS2{F0: genFn0(42, 4294967295, genFn1())}
    v74 := GenS3{F0: 701, F1: -1}
    var v75 uint64 = 147
    var v76 float32 = 3.14159
    if false {
        v77 := GenS6{F0: 0}
        var v78 int8 = (v3 | v3)
        _ = v77
        _ = v78
    }
    var v79 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v72:", v72)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v79
    _ = v9
}
