package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint16; F1 uint; F2 int16 }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 int64) int32 {
    p0 += -9223372036854775808
    var v5 uint64 = 0
    _ = p0
    _ = v5
    return -2147483648
}
func genFn1() uint32 {
    if (0.0 <= 0.1) {
        var v6 int = 0
        fmt.Println(v6)
        var v7 uint8 = 255
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float64 = 1e10
            if (20 == 42) {
                var v10 bool = (false || true)
                _ = v10
            } else {
                var v11 bool = (!(-2.25 <= 1e10))
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS1{F0: 18446744073709551615, F1: -1}
                    if (true || (false && (v11 && (-128 <= -128)))) {
                        v14 := GenS1{F0: 255, F1: 1}
                        if (42 > 18446744073709551615) {
                            v15 := []int{0, (fi12 / ((fi8) | 1))}
                            if ((1 >= -1) && v11) {
                                var v16 uint32 = 63
                                fmt.Println(v16, v9)
                                v17 := GenS1{F0: 65535, F1: -82}
                                _ = v16
                                _ = v17
                            } else {
                                if v11 {
                                    var v18 string = ("" + "hello")
                                    _ = v18
                                } else {
                                    v9 += ((((v9 / v9) * v9) * 0.0) - -2.25)
                                    var v19 int64 = -1
                                    if (!v11) {
                                        v20 := GenS2{F0: v11, F1: 2, F2: fi8}
                                        _ = v20
                                    }
                                    v21 := GenS3{F0: fi8}
                                    _ = v19
                                    _ = v21
                                }
                                v22 := GenS4{F0: 0, F1: 2147483648, F2: 149}
                                var v23 string = ((("foo bar" + "") + "\n\t") + "a")
                                _ = v22
                                _ = v23
                            }
                            v11 = v11
                            v11 = (2147483647 <= 471)
                            _ = v15
                        }
                        _ = v14
                    } else {
                        if ((v9 + v9) > (0.0 + v9)) {
                            v24 := []int{(fi12 % (((fi8 + 767)) | 1)), -714}
                            var v25 string = "foo bar"
                            fmt.Println(fi8, fi8, fi12)
                            _ = v24
                            _ = v25
                        } else {
                            v26 := GenS4{F0: 0, F1: 460, F2: 6}
                            v27 := GenS4{F0: 2147483647, F1: 87, F2: 1}
                            _ = v26
                            _ = v27
                        }
                        v11 = (42 < 2)
                    }
                    if true {
                        v28 := []int{fi8, 42, 1, fi8}
                        v29 := GenS3{F0: fi8}
                        fmt.Println(v9, v9)
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS1{F0: 274, F1: -128}
                        _ = v30
                    }
                    _ = fi12
                    _ = v13
                }
                v31 := []int{fi8, 0}
                _ = v11
                _ = v31
            }
            _ = fi8
            _ = v9
        }
    }
    if true {
        var v32 float32 = -2.25
        _ = v32
    } else {
        v33 := []int{-1, 127}
        _ = v33
    }
    v34 := []int{-2147483648}
    if true {
        v35 := GenS5{F0: 1}
        var v36 uint16 = 277
        _ = v35
        _ = v36
    }
    var v37 float32 = (1.5 / -2.25)
    _ = v34
    _ = v37
    return 387
}

func main() {
    var v1 uint32 = 970
    var v2 float64 = 1e10
    v3 := make(map[uint32]int)
    var v4 int32 = -58
    var v38 float64 = v2
    v1 -= (v1 - (v1 + (v1 - 2)))
    v39 := []int{2147483647, 32767}
    var v40 int64 = 614
    v41 := GenS4{F0: (v4 * 0), F1: genFn1(), F2: 2}
    var v42 uint8 = 2
    if (true && true) {
        var v43 uint32 = (v1 << ((2) & 63))
        var v44 int64 = 42
        if len(v39) > 0 {
            v39[0] = 32767
        }
        _ = v43
        _ = v44
    }
    switch 127 {
    case 590:
        for fi45 := 0; fi45 < 9; fi45++ {
            if true {
                v40 = (v40 / ((v40) | 1))
                if (!(1.5 == 1.5)) {
                    fmt.Println(v40, v1, v38)
                    v39 = append(v39, fi45)
                    _ = len(v3)
                } else {
                    v1 = v1
                }
                fmt.Println(v1)
            }
            v40 = (v40 + v40)
            _ = fi45
        }
    case 0:
        v46 := v3[genFn1()]
        _ = v46
    default:
        v4 = 1
    }
    v1 += 1
    v2 += v38
    if (true && (("foo bar" + "a") >= "hello")) {
        var v47 string = ""
        if false {
            var v48 int16 = 30
            v42 += v42
            _ = v48
        }
        if (false || true) {
            v1 = genFn1()
            if false {
                fmt.Println(v47)
                var v49 int16 = -32768
                switch 32767 {
                case 128:
                    if (v4 != (v4 * v4)) {
                        if (243 >= 408) {
                            v49 += (v49 >> ((42) & 63))
                        } else {
                            var v50 float32 = 0.0
                            var v51 uint8 = (^(v42))
                            v47 = ((v47 + v47) + v47)
                            _ = v50
                            _ = v51
                        }
                        if (v49 <= v49) {
                            if true {
                                v52 := GenS6{F0: 0, F1: v47}
                                v3[v1] = 4294967295
                                _ = v52
                            }
                            switch 2147483647 {
                            case 956:
                                switch 32768 {
                                case 128:
                                    v4 = v4
                                case 450:
                                    fmt.Println(v49, v42, v2)
                                case 9223372036854775807:
                                    var v53 int8 = -1
                                    _ = v53
                                default:
                                    if true {
                                        v54 := GenS7{F0: v49}
                                        var v55 uint64 = 0
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            case -887:
                                v56 := GenS3{F0: 32767}
                                _ = v56
                            case 4294967295:
                                if false {
                                    if (true && ((true || ((false || true) || true)) && ("\n\t" == ("foo bar" + ((v47 + "\n\t") + ("\n\t" + v47)))))) {
                                        switch "a" {
                                        case "a":
                                            v57 := GenS8{F0: ((v38 * float64(1e10)) <= v2), F1: 255}
                                            _ = v57
                                        default:
                                            v58 := GenS2{F0: (false && true), F1: (v42 & v42), F2: -877}
                                            _ = v58
                                        }
                                        v59 := GenS3{F0: 934}
                                        _ = v59
                                    }
                                    if true {
                                        v60 := GenS5{F0: 1}
                                        _ = v60
                                    } else {
                                        for fi61 := 0; fi61 < 1; fi61++ {
                                            var v62 uint8 = v42
                                            var v63 uint16 = 806
                                            v38 -= v38
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                        if false {
                                            for fi64 := 0; fi64 < 17; fi64++ {
                                                if false {
                                                    v65 := v3[(v1 + v1)]
                                                    _ = v65
                                                } else {
                                                    if (0.0 == 0.1) {
                                                        v49 += v49
                                                    } else {
                                                        if false {
                                                            for fi66 := 0; fi66 < 3; fi66++ {
                                                                v47 = (("" + (("hello" + (v47 + "hello")) + (v47 + v47))) + ((v47 + "") + v47))
                                                                _ = fi66
                                                            }
                                                        } else {
                                                            switch (v47 + "") {
                                                            case "":
                                                                v47 = v47
                                                            case "hello":
                                                                v67 := GenS4{F0: (^(v4)), F1: v1, F2: (v42 | (v42 / (((v42 * 130)) | 1)))}
                                                                _ = v67
                                                            case "foo bar":
                                                                var v68 int64 = -9223372036854775808
                                                                _ = v68
                                                            default:
                                                                fmt.Println(v38, v2, v49)
                                                            }
                                                            v38 = (float64(0.0) - -2.25)
                                                            v69 := GenS8{F0: (238 != (v40 ^ 1)), F1: 945}
                                                            _ = v69
                                                        }
                                                        if len(v39) >= 2 {
                                                            v39 = v39[1:len(v39)]
                                                        }
                                                    }
                                                    for fi70 := 0; fi70 < 1; fi70++ {
                                                        var v71 bool = (42 != -55)
                                                        v1 = v1
                                                        _ = fi70
                                                        _ = v71
                                                    }
                                                    v72 := GenS4{F0: v4, F1: genFn1(), F2: v42}
                                                    _ = v72
                                                }
                                                var v73 float64 = (0.0 - v2)
                                                _ = fi64
                                                _ = v73
                                            }
                                            var v74 float64 = ((((0.1 + float64((float32(-2.25) + ((-2.25 - 1.5) - 0.0)))) / v38) / 3.14159) * v2)
                                            _ = v74
                                        }
                                        switch v47 {
                                        case "":
                                            var v75 string = "\n\t"
                                            _ = v75
                                        case "hello":
                                            fmt.Println(v42, v2)
                                        case "foo bar":
                                            var v76 bool = (32767 == 42)
                                            _ = v76
                                        }
                                    }
                                } else {
                                    fmt.Println(v47, v1)
                                }
                            }
                            if false {
                                for fi77 := 0; fi77 < 1; fi77++ {
                                    var v78 uint64 = 42
                                    var v79 int8 = -76
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                }
                                var v80 uint64 = 18446744073709551615
                                _ = v80
                            } else {
                                var v81 int8 = 42
                                _ = v81
                            }
                        } else {
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                            v82 := GenS5{F0: 65535}
                            v4 += (v4 / ((v4) | 1))
                            _ = v82
                        }
                        var v83 uint64 = 0
                        _ = v83
                    } else {
                        v3[v1] = -138
                        if (true || (true && false)) {
                            if true {
                                v39 = append(v39, int(v40))
                                v42 -= v42
                                v1 = 65535
                            } else {
                                v2 = v2
                                v84 := GenS2{F0: (!(!(!((v1 | (v1 * genFn1())) <= v1)))), F1: v42, F2: 406}
                                v2 = v2
                                _ = v84
                            }
                            switch "a" {
                            case "\n\t":
                                v38 = (3.14159 * (v2 + (0.0 + (((-2.25 * (v38 / float64(1.5))) + float64(1e10)) - 0.0))))
                            case "a":
                                for fi85 := 0; fi85 < 2; fi85++ {
                                    if false {
                                        var v86 string = v47
                                        var v87 int32 = 0
                                        var v88 uint64 = 1
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    } else {
                                        var v89 string = (("hello" + "foo bar") + ("foo bar" + (v47 + v47)))
                                        var v90 float64 = v38
                                        _ = v89
                                        _ = v90
                                    }
                                    for fi91 := 0; fi91 < 11; fi91++ {
                                        var v92 int16 = (v49 >> ((65535) & 63))
                                        if true {
                                            if len(v39) > 0 {
                                                v39[0] = -2147483648
                                            }
                                            v49 -= (v49 | (v92 & v92))
                                        }
                                        _ = fi91
                                        _ = v92
                                    }
                                    _ = fi85
                                }
                            default:
                                for fi93 := 0; fi93 < 4; fi93++ {
                                    v42 = (v42 / (((v42 & (v42 % (((^(v42))) | 1)))) | 1))
                                    v94 := GenS9{F0: 65535, F1: 741, F2: (^(v49))}
                                    _ = fi93
                                    _ = v94
                                }
                            }
                        }
                    }
                }
                _ = v49
            } else {
                var v95 uint8 = v42
                _ = v95
            }
        }
        _ = v47
    }
    v96 := GenS10{F0: 0, F1: uint8(v1)}
    v42 = (v42 + v42)
    if (!false) {
        if ((-2.25 == -2.25) || (false || (776 == 42))) {
            if (true && true) {
                v4 = (-(v4))
            }
            v97 := GenS11{F0: "", F1: "hello", F2: -32768}
            var v98 uint = 255
            _ = v97
            _ = v98
        } else {
            var v99 uint32 = v1
            _ = v99
        }
        if false {
            v40 += 9223372036854775807
        }
    } else {
        var v100 int32 = v4
        var v101 uint8 = v42
        _ = len(v3)
        _ = v100
        _ = v101
    }
    var v102 uint64 = 0
    var v103 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v96
}
