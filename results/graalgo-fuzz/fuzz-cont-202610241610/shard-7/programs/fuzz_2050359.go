package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint; F1 int32; F2 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 uint8; F1 int8; F2 bool }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint32; F1 uint32 }
func genFn0() int32 {
    var v7 uint8 = 42
    v8 := GenS1{F0: -2147483648, F1: ""}
    var v9 uint = 18446744073709551615
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 1
    v2 := []int{370}
    var v3 uint = (v1 >> (((v1 | 696)) & 63))
    v4 := make(map[string]int)
    var v5 uint16 = 65535
    v6 := make(map[uint32]uint32)
    v5 = 37
    fmt.Println(v5, v1, v3)
    fmt.Println(v1, v5, v5)
    var v10 uint64 = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: 266, F1: ("foo bar" + "hello")}
        var v13 uint16 = 1
        _ = fi11
        _ = v12
        _ = v13
    }
    v1 = v3
    for fi14 := 0; fi14 < 4; fi14++ {
        if true {
            for fi15 := 0; fi15 < 4; fi15++ {
                delete(v6, 780)
                switch 0 {
                case 127:
                    var v16 int = 2147483647
                    _ = v16
                case -308:
                    v5 = (v5 * 65535)
                default:
                    v1 = (v3 << ((18446744073709551615) & 63))
                }
                v2 = append(v2, (fi14 >> ((v3) & 63)))
                _ = fi15
            }
            v17 := GenS1{F0: genFn0(), F1: ("hello" + ("" + "\n\t"))}
            v18 := GenS2{F0: 127}
            _ = v17
            _ = v18
        } else {
            if false {
                if true {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if (false && (false && (0.0 > 1e10))) {
                            if ("\n\t" >= ("\n\t" + ("" + ""))) {
                                v20 := GenS3{F0: v1, F1: -99}
                                var v21 bool = (1 > 9223372036854775807)
                                var v22 int16 = 978
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        v23 := GenS4{F0: 4294967295, F1: true, F2: (v10 % (((v10 & uint64(v5))) | 1))}
                        switch "\n\t" {
                        case "":
                            v24 := GenS2{F0: 1}
                            _ = v24
                        case "a":
                            if true {
                                switch 4294967295 {
                                case 9223372036854775807:
                                    var v25 int64 = int64(v5)
                                    _ = v25
                                case 32767:
                                    fmt.Println(v10, fi14, v3)
                                default:
                                    switch "foo bar" {
                                    case "":
                                        var v26 uint8 = 128
                                        switch ("hello" + "foo bar") {
                                        case "":
                                            v27 := GenS5{F0: 9223372036854775807, F1: (^(v3))}
                                            _ = v27
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v28 uint8 = (v26 << ((2) & 63))
                                                _ = v28
                                            case "a":
                                                v3 += v3
                                            default:
                                                if false {
                                                    var v29 uint32 = uint32(v1)
                                                    var v30 int8 = 8
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                        case "\n\t":
                                            switch (("hello" + ("foo bar" + "foo bar")) + ("\n\t" + "\n\t")) {
                                            case "a":
                                                var v31 int = -288
                                                _ = v31
                                            default:
                                                if (0 > -1) {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v32 uint8 = v26
                                                        _ = v32
                                                    case "a":
                                                        var v33 uint64 = (^(v10))
                                                        _ = v33
                                                    case "foo bar":
                                                        var v34 float64 = -2.25
                                                        _ = v34
                                                    }
                                                    if (!(fi19 == 720)) {
                                                        v35 := GenS6{F0: v1, F1: genFn0(), F2: (v1 + v1)}
                                                        if true {
                                                            var v36 uint = (v1 >> ((v3) & 63))
                                                            var v37 uint32 = 1
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi14 / ((int(v37)) | 1))
                                                            }
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        if (((0.1 == -2.25) && (true || false)) || ((!(1 != -1)) || false)) {
                                                            v38 := GenS7{F0: v3}
                                                            v39 := GenS8{F0: 124, F1: uint16(v26)}
                                                            switch int(v10) {
                                                            case -970:
                                                                switch fi19 {
                                                                case -378:
                                                                    var v40 int16 = 989
                                                                    _ = v40
                                                                case 2147483647:
                                                                    switch fi19 {
                                                                    case 128:
                                                                        var v41 string = ("a" + ((("" + "hello") + "\n\t") + ("a" + ("hello" + ("foo bar" + "foo bar")))))
                                                                        _ = v41
                                                                    default:
                                                                        var v42 uint64 = (^(v10))
                                                                        _ = v42
                                                                    }
                                                                default:
                                                                    v5 += v5
                                                                }
                                                            case 127:
                                                                var v43 int = (fi19 - fi14)
                                                                _ = v43
                                                            case -166:
                                                                v44 := GenS9{F0: 1, F1: "foo bar"}
                                                                _ = v44
                                                            default:
                                                                var v45 uint64 = (v10 << ((v3) & 63))
                                                                _ = v45
                                                            }
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v46 bool = ((false || false) || false)
                                                                _ = v46
                                                            case "a":
                                                                v47 := GenS7{F0: (v1 * (v3 - v1))}
                                                                switch (-(fi14)) {
                                                                case 534:
                                                                    delete(v6, uint32(v5))
                                                                case -9223372036854775808:
                                                                    v48 := GenS10{F0: 1, F1: 0}
                                                                    _ = v48
                                                                case 32768:
                                                                    switch fi14 {
                                                                    case 192:
                                                                        if true {
                                                                            v5 += (v5 - (^(v5)))
                                                                            if false {
                                                                                switch 32767 {
                                                                                case 127:
                                                                                    v49 := GenS8{F0: 127, F1: (^(v5))}
                                                                                    _ = v49
                                                                                case 9223372036854775807:
                                                                                    var v50 uint32 = 524
                                                                                    _ = v50
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        if (!true) {
                                                                                            v26 = (^(v26))
                                                                                        } else {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = fi14
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v3 += uint(v26)
                                                                                    }
                                                                                }
                                                                                switch 42 {
                                                                                case 58:
                                                                                    v51 := GenS5{F0: 0, F1: 42}
                                                                                    _ = v51
                                                                                default:
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (fi19 | (-(fi14)))
                                                                                    }
                                                                                }
                                                                                _ = len(v6)
                                                                            }
                                                                            var v52 int16 = 42
                                                                            v26 -= v26
                                                                            _ = v52
                                                                        } else {
                                                                            v53 := GenS8{F0: 100, F1: (v5 + v5)}
                                                                            _ = v53
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            if false {
                                                                                var v54 uint64 = v10
                                                                                v26 = 255
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    switch (fi14 << (((v3 | 2)) & 63)) {
                                                                                    case 960:
                                                                                        _ = len(v2)
                                                                                    case 9223372036854775807:
                                                                                        switch fi19 {
                                                                                        case 4294967295:
                                                                                            fmt.Println(fi19, v54)
                                                                                        case -124:
                                                                                            v54 = (^(v10))
                                                                                        }
                                                                                    case 128:
                                                                                        if (42 <= (v1 ^ (v3 % (((v1 >> (((v1 - v3)) & 63))) | 1)))) {
                                                                                            if (false && false) {
                                                                                                var v55 float64 = ((-2.25 + 1e10) / (((1.5 + ((-2.25 + (1e10 / (0.1 - 3.14159))) - float64(0.1))) / 3.14159) + float64((float32(((float64((3.14159 * ((1e10 / ((-2.25 / (1e10 * (float32(0.0) + 1.5))) - float32(-2.25))) + 0.1))) * (0.1 - 1.5)) * 0.0)) - (3.14159 + 3.14159)))))
                                                                                                v56 := GenS4{F0: 65535, F1: (true || (v55 < float64(((0.0 / (0.0 + float32(v55))) + 0.0)))), F2: v10}
                                                                                                _ = v55
                                                                                                _ = v56
                                                                                            } else {
                                                                                                fmt.Println(v1, v26, v1)
                                                                                                fmt.Println(v3, fi19)
                                                                                            }
                                                                                            var v57 float32 = 0.1
                                                                                            if (true && false) {
                                                                                                v58 := GenS5{F0: int64(v26), F1: v1}
                                                                                                var v59 float64 = 0.0
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    var v60 string = "\n\t"
                                                                                                    _ = v60
                                                                                                }
                                                                                                _ = v58
                                                                                                _ = v59
                                                                                            } else {
                                                                                                switch fi19 {
                                                                                                case 4294967295:
                                                                                                    var v61 uint16 = (v5 - v5)
                                                                                                    _ = v61
                                                                                                case 32767:
                                                                                                    fmt.Println(v57, v1, v1)
                                                                                                case -219:
                                                                                                    v62 := GenS5{F0: -9223372036854775808, F1: v1}
                                                                                                    _ = v62
                                                                                                }
                                                                                                fmt.Println(v54, v54)
                                                                                                v6[2147483647] = 494
                                                                                                v10 += uint64(v3)
                                                                                                v1 = (^(v3))
                                                                                            }
                                                                                            _ = v57
                                                                                        } else {
                                                                                            v54 = (^(v54))
                                                                                            var v63 float32 = 0.1
                                                                                            _ = v63
                                                                                        }
                                                                                    default:
                                                                                        v3 = (^(v3))
                                                                                    }
                                                                                case "\n\t":
                                                                                    fmt.Println(v1, v10)
                                                                                case "hello":
                                                                                    fmt.Println(fi14, v54, v10)
                                                                                default:
                                                                                    v4[("hello" + "hello")] = (fi19 + fi14)
                                                                                }
                                                                                _ = v54
                                                                            } else {
                                                                                var v64 uint16 = 42
                                                                                v64 = (v5 - uint16(v3))
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v65 string = (("hello" + "foo bar") + ((("a" + ("hello" + ((("" + "") + "hello") + ""))) + ("hello" + ("foo bar" + ""))) + "hello"))
                                                                                v64 = (v5 >> (((^(v1))) & 63))
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                            switch (-(fi14)) {
                                                                            case 2147483647:
                                                                                v66 := GenS9{F0: (v10 + v10), F1: (("foo bar" + "a") + ("foo bar" + "hello"))}
                                                                                _ = v66
                                                                            }
                                                                        } else {
                                                                            switch "a" {
                                                                            case "":
                                                                                _ = len(v2)
                                                                            }
                                                                            v67 := GenS10{F0: 201, F1: 1}
                                                                            var v68 string = "a"
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    v69 := GenS11{F0: v26, F1: 100, F2: (v3 <= 65535)}
                                                                    _ = v69
                                                                default:
                                                                    switch fi19 {
                                                                    case 32767:
                                                                        switch -674 {
                                                                        case 927:
                                                                            if true {
                                                                                fmt.Println(v5, v5)
                                                                                v10 = v10
                                                                                var v70 uint8 = (v26 ^ 2)
                                                                                _ = v70
                                                                            }
                                                                        }
                                                                    case -311:
                                                                        v3 += 65535
                                                                    case 2147483647:
                                                                        if (true && false) {
                                                                            v5 += 402
                                                                            if (2 < (v26 & (v26 + (v26 + 255)))) {
                                                                                v71 := GenS12{F0: (v3 % ((v3) | 1))}
                                                                                var v72 int8 = -128
                                                                                _ = v71
                                                                                _ = v72
                                                                            }
                                                                        } else {
                                                                            v2 = append(v2, (fi14 * 4294967295))
                                                                            v26 = v26
                                                                            switch "a" {
                                                                            case "":
                                                                                if true {
                                                                                    v26 = v26
                                                                                    var v73 int8 = 127
                                                                                    _ = v73
                                                                                } else {
                                                                                    var v74 uint16 = v5
                                                                                    _ = v74
                                                                                }
                                                                            case "foo bar":
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                    default:
                                                                        v26 = (^(v26))
                                                                    }
                                                                    switch (fi14 / ((fi19) | 1)) {
                                                                    case 0:
                                                                        _ = len(v2)
                                                                    }
                                                                    v5 = (v5 << ((v3) & 63))
                                                                }
                                                                _ = v47
                                                            case "hello":
                                                                var v75 uint8 = (v26 & 161)
                                                                _ = v75
                                                            default:
                                                                if true {
                                                                    fmt.Println(v26, v10)
                                                                    var v76 float32 = 0.1
                                                                    v3 += v1
                                                                    _ = v76
                                                                } else {
                                                                    v77 := GenS3{F0: (^(v1)), F1: int16(v26)}
                                                                    var v78 uint32 = 179
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                            }
                                                            v1 = v1
                                                            v10 = 0
                                                            switch 32767 {
                                                            case 557:
                                                                var v79 int32 = 0
                                                                fmt.Println(fi19, v26, v79)
                                                                v80 := GenS3{F0: (v1 + (v3 << ((v1) & 63))), F1: 32767}
                                                                v81 := GenS13{F0: 2147483647, F1: 1}
                                                                var v82 int8 = 100
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                                _ = v82
                                                            default:
                                                                var v83 uint8 = (v26 / ((v26) | 1))
                                                                _ = v83
                                                            }
                                                            fmt.Println(fi19, fi19, fi14)
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            fmt.Println(v10)
                                                            if false {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v3 = (v3 * (v1 * 2))
                                                                v5 += 0
                                                            }
                                                        }
                                                        _ = v35
                                                    }
                                                    var v84 string = "hello"
                                                    _ = v84
                                                } else {
                                                    v85 := GenS4{F0: 4294967295, F1: ((v1 * v3) == v3), F2: (^(v10))}
                                                    var v86 uint = v3
                                                    v87 := GenS10{F0: (v5 << ((v1) & 63)), F1: 100}
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                }
                                            }
                                        }
                                        _ = v26
                                    case "hello":
                                        var v88 uint16 = (v5 % (((v5 >> ((42) & 63))) | 1))
                                        _ = v88
                                    case "foo bar":
                                        fmt.Println(v5, v5, v5)
                                    }
                                }
                            } else {
                                var v89 int8 = -128
                                var v90 uint16 = v5
                                v91 := GenS5{F0: 9223372036854775807, F1: (^(v3))}
                                _ = v89
                                _ = v90
                                _ = v91
                            }
                        case "\n\t":
                            if true {
                                if true {
                                    var v92 bool = (false || true)
                                    _ = v92
                                }
                                if len(v2) > 0 {
                                    v2[0] = fi19
                                }
                            } else {
                                if true {
                                    switch (fi14 << ((v1) & 63)) {
                                    case 4294967295:
                                        v93 := GenS7{F0: v1}
                                        var v94 int32 = -2147483648
                                        v10 = v10
                                        _ = v93
                                        _ = v94
                                    }
                                }
                                var v95 int64 = 1
                                v96 := GenS1{F0: 2147483647, F1: "foo bar"}
                                _ = v95
                                _ = v96
                            }
                        default:
                            fmt.Println(v3)
                        }
                        _ = fi19
                        _ = v23
                    }
                    var v97 int32 = -175
                    _ = v97
                }
                if false {
                    v1 -= 18446744073709551615
                }
                v3 = (v3 & 4294967295)
            }
        }
        _ = fi14
    }
    for fi98 := 0; fi98 < 18; fi98++ {
        fmt.Println(v3, fi98)
        v3 = 0
        var v99 uint = 42
        _ = fi98
        _ = v99
    }
    fmt.Println(v5)
    v100 := GenS7{F0: 2}
    fmt.Println(v1)
    v101 := GenS6{F0: (v3 ^ 0), F1: 42, F2: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
