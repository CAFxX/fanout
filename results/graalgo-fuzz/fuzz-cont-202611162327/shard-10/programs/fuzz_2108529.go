package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint8; F1 uint64 }
func genFn0() bool {
    var v5 string = ("" + "hello")
    var v6 uint32 = 2147483648
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 int) uint {
    v7 := GenS1{F0: 4294967295, F1: 892, F2: (127 < 42)}
    v8 := GenS1{F0: 386, F1: 2147483647, F2: false}
    var v9 uint16 = 32767
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []int{18}
    v2 := make(map[string]bool)
    var v3 uint32 = 8
    var v4 uint16 = 0
    fmt.Println(v4, v3)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint16 = (v4 & (v4 >> ((2) & 63)))
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: genFn1(-518), F1: -1, F2: (1 >= -2147483648)}
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 16; fi13++ {
        if (!true) {
            if false {
                if false {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if (true || ((float32(((1.5 * ((0.0 * (float64(0.1) - float64(0.0))) / 1.5)) / 1.0)) * 1e10) == 1e10)) {
                            v15 := GenS2{F0: 727}
                            v16 := GenS1{F0: 0, F1: 2147483647, F2: false}
                            _ = v15
                            _ = v16
                        } else {
                            v17 := GenS3{F0: v3, F1: 304, F2: 0}
                            _ = v17
                        }
                        if true {
                            var v18 int8 = 0
                            var v19 int16 = 273
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int64 = 9223372036854775807
                            _ = v20
                        }
                        v3 -= 1
                        _ = fi14
                    }
                }
                var v21 int8 = 0
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS1{F0: 4294967295, F1: 42, F2: (7 <= 246)}
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                v24 := v2[("\n\t" + ("foo bar" + "a"))]
                v25 := GenS1{F0: 42, F1: 0, F2: (18446744073709551615 < genFn1(-758))}
                v24 = v24
                _ = v24
                _ = v25
            }
            if true {
                fmt.Println(v3, v4)
            } else {
                if false {
                    v26 := GenS2{F0: 42}
                    _ = v26
                }
                if true {
                    v27 := GenS2{F0: 67}
                    var v28 int16 = 1000
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi13)
                    case "a":
                        var v29 string = ("foo bar" + "")
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                } else {
                    switch fi13 {
                    case -2147483648:
                        v4 = (v4 % (((v4 << ((2) & 63))) | 1))
                    case 9223372036854775807:
                        for fi30 := 0; fi30 < 3; fi30++ {
                            switch 32768 {
                            case 127:
                                var v31 string = "\n\t"
                                _ = v31
                            case 566:
                                var v32 uint = 76
                                _ = v32
                            default:
                                var v33 int16 = 0
                                _ = v33
                            }
                            _ = fi30
                        }
                    case -178:
                        fmt.Println(fi13)
                    }
                    switch -1 {
                    case 0:
                        v2["a"] = (!(!(313 >= 648)))
                    case 67:
                        v4 -= (v4 / ((255) | 1))
                    }
                    fmt.Println(v3)
                }
                v34 := GenS1{F0: 873, F1: 111, F2: (!((3.14159 - 1e10) < (((1e10 - -2.25) - (float32(0.0) - 1e10)) - 1.5)))}
                _ = v34
            }
            if len(v1) > 0 {
                v1[0] = (fi13 * (fi13 >> ((1) & 63)))
            }
        }
        if true {
            if true {
                v35 := GenS4{F0: 2147483647, F1: -705}
                v36 := v2["a"]
                _ = v35
                _ = v36
            } else {
                var v37 int64 = -9223372036854775808
                if false {
                    v3 += v3
                }
                _ = v37
            }
            var v38 uint16 = v4
            if (127 >= 128) {
                v39 := GenS4{F0: 733, F1: 686}
                _ = v39
            }
            _ = v38
        }
        _ = fi13
    }
    if ((255 >= 127) && (!(!(0 <= -1)))) {
        fmt.Println(v4, v4)
        delete(v2, ("a" + ("\n\t" + "\n\t")))
        var v40 uint16 = v4
        _ = v40
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 float64 = float64((((0.1 + -2.25) + ((0.1 + float32(-2.25)) / 3.14159)) - (1e10 - 1e10)))
        var v43 int8 = 127
        _ = fi41
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        for fi45 := 0; fi45 < 20; fi45++ {
            switch "" {
            case "hello":
                var v46 int32 = 42
                _ = v46
            case "a":
                var v47 int32 = 386
                _ = v47
            }
            if (-373 < 9223372036854775807) {
                v48 := GenS1{F0: genFn1(0), F1: -94, F2: false}
                v3 += 1
                fmt.Println(v3, v3)
                v49 := GenS3{F0: v3, F1: 820, F2: -1}
                _ = v48
                _ = v49
            } else {
                if false {
                    v3 -= (v3 % ((v3) | 1))
                    switch ("foo bar" + ("foo bar" + "a")) {
                    case "a":
                        v3 = (v3 + 2147483647)
                    case "\n\t":
                        if ((v3 + v3) > (v3 / ((821) | 1))) {
                            fmt.Println(fi44)
                            v50 := GenS3{F0: v3, F1: 0, F2: -441}
                            v4 = v4
                            _ = v50
                        } else {
                            v4 = 0
                            var v51 uint16 = 767
                            if (("\n\t" + "a") < "") {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                var v52 int = (fi44 * fi44)
                                if (v3 == (v3 >> ((18446744073709551615) & 63))) {
                                    v53 := GenS5{F0: -685}
                                    _ = v53
                                } else {
                                    if false {
                                        if ((!true) && (!(false && (400 == -675)))) {
                                            v54 := GenS1{F0: 0, F1: 112, F2: false}
                                            if (4294967295 == 255) {
                                                if len(v1) > 0 {
                                                    v1[0] = 9223372036854775807
                                                }
                                                if false {
                                                    v55 := v2["\n\t"]
                                                    _ = v55
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v51 = 1
                                                        case "":
                                                            v51 = (^(v4))
                                                        default:
                                                            fmt.Println(v51, v51)
                                                        }
                                                    case "\n\t":
                                                        v51 += (v4 | v51)
                                                    case "hello":
                                                        switch (v52 >> ((genFn1((-(v52)))) & 63)) {
                                                        case 514:
                                                            switch ("" + ("a" + ("a" + ("foo bar" + "")))) {
                                                            case "a":
                                                                v52 -= fi44
                                                            case "hello":
                                                                v52 = (fi45 >> ((18446744073709551615) & 63))
                                                            case "":
                                                                if true {
                                                                    v52 -= fi44
                                                                    if false {
                                                                        v56 := GenS3{F0: (^(v3)), F1: 42, F2: 132}
                                                                        _ = v56
                                                                    } else {
                                                                        fmt.Println(v51, fi45, fi45)
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v51 += 0
                                                                        case "foo bar":
                                                                            v3 -= (v3 | 65535)
                                                                        case "\n\t":
                                                                            v52 -= v52
                                                                        }
                                                                        if true {
                                                                            fmt.Println(fi44)
                                                                        } else {
                                                                            var v57 uint16 = (^(v51))
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                    fmt.Println(v3, v51, fi45)
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            var v58 uint8 = 200
                                                            _ = v58
                                                        case 530:
                                                            if ((!false) && false) {
                                                                var v59 bool = (!true)
                                                                delete(v2, "\n\t")
                                                                _ = v59
                                                            }
                                                        default:
                                                            v51 = 1
                                                        }
                                                    default:
                                                        if (1 >= 2) {
                                                            v2[(("hello" + "foo bar") + (("\n\t" + "foo bar") + "hello"))] = (!true)
                                                            v2[("foo bar" + "")] = true
                                                            var v60 int = (^(v52))
                                                            delete(v2, "foo bar")
                                                            _ = v60
                                                        } else {
                                                            switch (fi44 >> ((genFn1(365)) & 63)) {
                                                            case 9223372036854775807:
                                                                var v61 float64 = 1e10
                                                                _ = v61
                                                            case 42:
                                                                var v62 int64 = -102
                                                                _ = v62
                                                            case 32767:
                                                                switch v52 {
                                                                case 4294967295:
                                                                    v3 = (v3 << ((uint(v4)) & 63))
                                                                case 1:
                                                                    v1 = append(v1, v52)
                                                                default:
                                                                    var v63 int64 = 0
                                                                    _ = v63
                                                                }
                                                            }
                                                            if true {
                                                                v64 := GenS6{F0: (v4 << ((0) & 63)), F1: 137, F2: 255}
                                                                var v65 float64 = 1e10
                                                                var v66 int64 = 1
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                switch ("hello" + "") {
                                                                case "":
                                                                    v51 += (v4 | (v51 & 65535))
                                                                case "hello":
                                                                    v3 = (v3 * (^(v3)))
                                                                    v51 += 129
                                                                    v3 = (v3 ^ (v3 % (((v3 | 270)) | 1)))
                                                                case "a":
                                                                    v67 := GenS5{F0: 1}
                                                                    _ = v67
                                                                default:
                                                                    v52 = (v52 >> ((42) & 63))
                                                                }
                                                            }
                                                            if ((0.0 * 0.0) <= 1e10) {
                                                                fmt.Println(v52, v51, v52)
                                                                v68 := GenS5{F0: 42}
                                                                fmt.Println(v52, v52, v52)
                                                                _ = v68
                                                            }
                                                        }
                                                    }
                                                    v69 := GenS5{F0: -811}
                                                    _ = v69
                                                }
                                            }
                                            v1 = append(v1, 32767)
                                            _ = v54
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v51, fi44)
                                    }
                                    var v70 float32 = (0.0 * 3.14159)
                                    _ = v70
                                }
                                var v71 int16 = 1
                                _ = v52
                                _ = v71
                            }
                            v4 -= (v4 | 65535)
                            if (77 != 565) {
                                delete(v2, ("hello" + "\n\t"))
                                var v72 float32 = ((0.0 * 0.1) / 1.0)
                                _ = v72
                            }
                            _ = v51
                        }
                    default:
                        if true {
                            var v73 uint8 = 1
                            v74 := GenS7{F0: v4, F1: 875, F2: (v73 + 0)}
                            v73 = v73
                            var v75 bool = false
                            switch ("\n\t" + "") {
                            case "hello":
                                _ = len(v2)
                            case "a":
                                if ((^(v73)) > v73) {
                                    fmt.Println(v3)
                                    var v76 uint16 = v4
                                    v77 := GenS2{F0: -636}
                                    if (65535 >= uint16(v73)) {
                                        fmt.Println(fi45, v73, v4)
                                    } else {
                                        var v78 int32 = 0
                                        fmt.Println(v73, v4, v3)
                                        var v79 int16 = 1000
                                        _ = v78
                                        _ = v79
                                    }
                                    v3 = v3
                                    v75 = v75
                                    _ = v76
                                    _ = v77
                                }
                            }
                            _ = v73
                            _ = v74
                            _ = v75
                        } else {
                            var v80 int8 = -76
                            var v81 int16 = 595
                            var v82 int = (fi45 / ((-246) | 1))
                            _ = v80
                            _ = v81
                            _ = v82
                        }
                        var v83 int = fi44
                        v84 := GenS3{F0: (v3 % (((v3 << ((genFn1(fi44)) & 63))) | 1)), F1: 18446744073709551615, F2: -281}
                        _ = v83
                        _ = v84
                    }
                    var v85 uint8 = 0
                    _ = v85
                }
                var v86 float64 = 0.1
                _ = v86
            }
            _ = fi45
        }
        for fi87 := 0; fi87 < 8; fi87++ {
            v88 := GenS5{F0: -1}
            if false {
                var v89 bool = true
                v89 = (!v89)
                switch 128 {
                case 32767:
                    v90 := GenS8{F0: 127, F1: 18446744073709551615}
                    _ = v90
                }
                _ = v89
            }
            _ = fi87
            _ = v88
        }
        _ = fi44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
