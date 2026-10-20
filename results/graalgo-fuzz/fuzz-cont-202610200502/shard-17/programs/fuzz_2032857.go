package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int8; F1 int64; F2 bool }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 bool; F1 int32; F2 uint32 }

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]string)
    var v3 uint8 = 127
    v4 := GenS1{F0: 2, F1: (("foo bar" + "hello") + "")}
    var v5 string = ("" + ("foo bar" + ("hello" + ("" + ("" + "a")))))
    v2[42] = v5
    v6 := []int{}
    for fi7 := 0; fi7 < 6; fi7++ {
        v5 = ((v5 + (v5 + v5)) + v5)
        v5 = ((v5 + ("" + (v5 + "hello"))) + v5)
        v3 = 2
        _ = fi7
    }
    v2[1] = v5
    v3 = (v3 + (v3 & (v3 / ((v3) | 1))))
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 0
        switch ("hello" + v5) {
        case "hello":
            var v10 string = v5
            _ = v10
        case "foo bar":
            if ((true || (uint64(v1) > v9)) && ((-575 == 1) && (!true))) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    switch (fi8 << (((v1 % ((v1) | 1))) & 63)) {
                    case 128:
                        v9 += (v9 % ((uint64(v3)) | 1))
                    case 4294967295:
                        switch v5 {
                        case "foo bar":
                            v12 := GenS2{F0: (v9 ^ v9)}
                            _ = v12
                        default:
                            var v13 bool = true
                            _ = v13
                        }
                    case 0:
                        if (!((-366 >= 1) && false)) {
                            var v14 float32 = -2.25
                            if true {
                                v5 = v5
                                if (2147483647 <= 686) {
                                    switch (fi11 % ((fi11) | 1)) {
                                    case 4294967295:
                                        _ = len(v6)
                                    case 128:
                                        fmt.Println(v5, v3)
                                    }
                                    var v15 int = fi11
                                    _ = v15
                                }
                                v14 -= float32(0.1)
                            }
                            _ = v14
                        } else {
                            if ((-814 <= 0) || true) {
                                var v16 int64 = -9223372036854775808
                                v1 = (v1 & 0)
                                v17 := v2[983]
                                _ = v16
                                _ = v17
                            } else {
                                v9 = 42
                                fmt.Println(fi8)
                                if (1 == -9223372036854775808) {
                                    switch "" {
                                    case "\n\t":
                                        var v18 bool = true
                                        _ = v18
                                    case "":
                                        if (!false) {
                                            if false {
                                                v9 = (v9 | (v9 & v9))
                                            } else {
                                                if ((v5 + v5) == v5) {
                                                    var v19 uint16 = 32767
                                                    fmt.Println(v19, v3)
                                                    var v20 int8 = 1
                                                    _ = v19
                                                    _ = v20
                                                }
                                                switch "hello" {
                                                case "a":
                                                    fmt.Println(fi8)
                                                default:
                                                    var v21 uint32 = 640
                                                    _ = v21
                                                }
                                                if false {
                                                    var v22 uint64 = (v9 / ((v9) | 1))
                                                    _ = v22
                                                } else {
                                                    var v23 uint8 = v3
                                                    _ = v23
                                                }
                                                v3 -= (v3 | v3)
                                                if len(v6) > 0 {
                                                    v6[0] = (fi11 << (((v1 ^ (^(v1)))) & 63))
                                                }
                                                var v24 uint32 = 42
                                                var v25 float32 = 0.1
                                                v5 = v5
                                                if true {
                                                    var v26 bool = (42 > 0)
                                                    if true {
                                                        v1 = v1
                                                        _ = len(v6)
                                                    } else {
                                                        var v27 uint64 = (^(v9))
                                                        fmt.Println(v26, v25, v5)
                                                        fmt.Println(v5, v25, fi11)
                                                        _ = v27
                                                    }
                                                    _ = v26
                                                } else {
                                                    if (false || ((v3 ^ (v3 << (((v1 | (v1 / (((v1 * (v1 & v1))) | 1)))) & 63))) <= (^(v3)))) {
                                                        var v28 uint16 = 65535
                                                        var v29 uint32 = v24
                                                        v3 -= (v3 * 255)
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        v30 := GenS2{F0: 4294967295}
                                                        v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
                                                        _ = v30
                                                    }
                                                    switch v5 {
                                                    case "a":
                                                        var v31 string = ""
                                                        _ = v31
                                                    case "hello":
                                                        v6 = append(v6, fi11)
                                                    case "":
                                                        if ((-505 > -9223372036854775808) || (!(true && (false && false)))) {
                                                            if len(v6) > 0 {
                                                                v6[0] = int(v24)
                                                            }
                                                            v32 := GenS1{F0: 294, F1: "foo bar"}
                                                            _ = v32
                                                        }
                                                    default:
                                                        v3 = (v3 | v3)
                                                    }
                                                    var v33 bool = true
                                                    _ = v33
                                                }
                                                _ = v24
                                                _ = v25
                                            }
                                            switch v5 {
                                            case "a":
                                                switch (fi11 & fi8) {
                                                case 9223372036854775807:
                                                    v5 = (v5 + (v5 + (("" + v5) + v5)))
                                                case -335:
                                                    v34 := GenS1{F0: 140, F1: ""}
                                                    _ = v34
                                                case 0:
                                                    v35 := GenS1{F0: 111, F1: (((("hello" + v5) + (v5 + v5)) + (v5 + ("\n\t" + "a"))) + ((v5 + v5) + v5))}
                                                    _ = v35
                                                }
                                            default:
                                                if false {
                                                    var v36 float32 = -2.25
                                                    var v37 int64 = 42
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                        } else {
                                            v38 := GenS3{F0: -2147483648, F1: 28, F2: 0}
                                            v5 = v5
                                            v39 := GenS2{F0: (v9 / ((v9) | 1))}
                                            _ = v38
                                            _ = v39
                                        }
                                    case "a":
                                        if (-756 > fi11) {
                                            _ = len(v6)
                                            if true {
                                                fmt.Println(v1)
                                                v40 := v2[2]
                                                if false {
                                                    var v41 uint64 = 18446744073709551615
                                                    _ = v41
                                                }
                                                _ = v40
                                            } else {
                                                v42 := GenS4{F0: 0, F1: (^(v1)), F2: -32768}
                                                v43 := GenS5{F0: false}
                                                v1 = uint(v3)
                                                _ = v42
                                                _ = v43
                                            }
                                        } else {
                                            if false {
                                                var v44 uint8 = v3
                                                if ((fi8 <= (fi11 - (-(fi11)))) || false) {
                                                    fmt.Println(v5, v1, fi8)
                                                }
                                                _ = v44
                                            } else {
                                                if (!(1 <= v3)) {
                                                    switch v5 {
                                                    case "a":
                                                        switch (("a" + (v5 + "")) + (v5 + "\n\t")) {
                                                        case "\n\t":
                                                            v9 -= (v9 ^ 1)
                                                        default:
                                                            v45 := GenS6{F0: 474, F1: (v9 ^ 236)}
                                                            _ = v45
                                                        }
                                                    case "\n\t":
                                                        v5 = (v5 + (v5 + v5))
                                                        v46 := GenS7{F0: 127, F1: -1, F2: ((false && (1.5 <= float64(3.14159))) && false)}
                                                        var v47 float32 = (-2.25 + 3.14159)
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    var v48 int32 = 1
                                                    _ = v48
                                                } else {
                                                    var v49 int64 = -26
                                                    var v50 uint32 = 42
                                                    _ = v49
                                                    _ = v50
                                                }
                                                if true {
                                                    if (1 > -32768) {
                                                        var v51 int32 = -1
                                                        _ = v51
                                                    } else {
                                                        _ = len(v2)
                                                        var v52 uint32 = 144
                                                        fmt.Println(fi11, v3)
                                                        fmt.Println(fi8, v9, v3)
                                                        var v53 int64 = -210
                                                        v3 -= (v3 | 2)
                                                        if (0 != -1) {
                                                            if (!((2147483647 > 0) || (!(1.5 != 0.1)))) {
                                                                if false {
                                                                    var v54 int32 = -2147483648
                                                                    var v55 uint8 = (v3 - uint8(v9))
                                                                    _ = v54
                                                                    _ = v55
                                                                } else {
                                                                    _ = len(v2)
                                                                    v52 -= 2147483647
                                                                    if true {
                                                                        var v56 string = v5
                                                                        v57 := GenS5{F0: ((42 == 744) && (65535 >= (v52 / (((^(v52))) | 1))))}
                                                                        v58 := GenS5{F0: (195 != 1)}
                                                                        _ = v56
                                                                        _ = v57
                                                                        _ = v58
                                                                    } else {
                                                                        v3 = (v3 * v3)
                                                                        var v59 uint32 = 671
                                                                        v1 -= v1
                                                                        _ = v59
                                                                    }
                                                                }
                                                            }
                                                            var v60 uint = 663
                                                            v61 := v2[(v52 & v52)]
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            var v62 uint8 = uint8(v9)
                                                            _ = v62
                                                        }
                                                        if false {
                                                            v1 = uint(v52)
                                                            v53 = (v53 % (((v53 & int64(v1))) | 1))
                                                            v63 := GenS8{F0: 42, F1: 1000}
                                                            _ = v63
                                                        } else {
                                                            if true {
                                                                switch -823 {
                                                                case -9223372036854775808:
                                                                    switch (-(fi8)) {
                                                                    case 1000:
                                                                        v5 = (("foo bar" + v5) + v5)
                                                                    case 9223372036854775807:
                                                                        var v64 uint64 = v9
                                                                        _ = v64
                                                                    default:
                                                                        _ = len(v6)
                                                                    }
                                                                case 32768:
                                                                    v65 := GenS4{F0: (v3 & v3), F1: v1, F2: -918}
                                                                    _ = v65
                                                                default:
                                                                    var v66 int16 = -32768
                                                                    _ = v66
                                                                }
                                                                if ((255 < 647) || true) {
                                                                    v3 = v3
                                                                    if true {
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi8
                                                                        }
                                                                    } else {
                                                                        v67 := GenS2{F0: v9}
                                                                        _ = v67
                                                                    }
                                                                }
                                                                v9 = v9
                                                                var v68 uint = v1
                                                                _ = v68
                                                            }
                                                        }
                                                        v69 := GenS6{F0: 1, F1: uint64(v1)}
                                                        v70 := GenS4{F0: 58, F1: (v1 - (v1 ^ v1)), F2: 0}
                                                        var v71 int16 = 247
                                                        _ = v52
                                                        _ = v53
                                                        _ = v69
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                                fmt.Println(v9)
                                            }
                                        }
                                    }
                                    switch fi11 {
                                    case 0:
                                        _ = len(v6)
                                    default:
                                        if (v1 < 2) {
                                            v1 = (v1 * (^(v1)))
                                            v1 = 65535
                                            fmt.Println(fi11)
                                        } else {
                                            var v72 float64 = ((3.14159 + (0.0 / -2.25)) / 1.0)
                                            if false {
                                                switch (fi11 / ((-1) | 1)) {
                                                case 9223372036854775807:
                                                    _ = len(v2)
                                                case -497:
                                                    var v73 int16 = 1000
                                                    _ = v73
                                                case -457:
                                                    v74 := GenS8{F0: 461, F1: 42}
                                                    _ = v74
                                                }
                                            } else {
                                                v5 = "hello"
                                                switch fi8 {
                                                case -723:
                                                    v75 := GenS4{F0: 2, F1: v1, F2: 244}
                                                    _ = v75
                                                case -174:
                                                    var v76 float64 = v72
                                                    _ = v76
                                                case 42:
                                                    v77 := GenS9{F0: (fi11 % ((-74) | 1))}
                                                    _ = v77
                                                default:
                                                    v3 -= v3
                                                }
                                                v72 -= (-2.25 / float64(float32(v72)))
                                            }
                                            var v78 float64 = float64(-2.25)
                                            _ = v72
                                            _ = v78
                                        }
                                    }
                                }
                            }
                            if (!((v9 * (v9 - (v9 * (^(v9))))) < 2)) {
                                if true {
                                    v1 += 663
                                    switch fi11 {
                                    case 0:
                                        v3 += v3
                                    }
                                } else {
                                    fmt.Println(v9, fi11, v1)
                                    v3 = (^(v3))
                                    v79 := GenS10{F0: v1, F1: 127, F2: 65535}
                                    _ = v79
                                }
                            } else {
                                fmt.Println(fi8, fi11)
                                var v80 int32 = -464
                                v3 += 1
                                _ = v80
                            }
                            fmt.Println(v3)
                        }
                    }
                    var v81 int8 = 54
                    _ = fi11
                    _ = v81
                }
            } else {
                if false {
                    for fi82 := 0; fi82 < 20; fi82++ {
                        if false {
                            switch (v5 + v5) {
                            case "a":
                                if (2 == 42) {
                                    var v83 uint16 = 255
                                    v84 := GenS11{F0: (127 < 42), F1: "foo bar"}
                                    _ = v83
                                    _ = v84
                                } else {
                                    var v85 int16 = 0
                                    v86 := GenS2{F0: (v9 >> ((v1) & 63))}
                                    v87 := GenS12{F0: v3}
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                            }
                            v5 = (((v5 + v5) + v5) + ("" + v5))
                            var v88 uint64 = 595
                            _ = v88
                        }
                        var v89 int32 = 0
                        _ = fi82
                        _ = v89
                    }
                } else {
                    v3 -= 42
                }
                if (!(100 == 100)) {
                    v90 := GenS1{F0: 333, F1: v5}
                    switch fi8 {
                    case -38:
                        v5 = v5
                    default:
                        v9 = 1
                    }
                    switch "hello" {
                    case "\n\t":
                        for fi91 := 0; fi91 < 9; fi91++ {
                            fmt.Println(fi8, fi8)
                            _ = len(v2)
                            _ = fi91
                        }
                    default:
                        v3 -= (v3 << (((v1 ^ 2)) & 63))
                    }
                    _ = v90
                }
                var v92 int32 = -63
                _ = v92
            }
        case "":
            fmt.Println(v3, fi8, v5)
        default:
            v5 = v5
        }
        _ = fi8
        _ = v9
    }
    v5 = (("foo bar" + v5) + v5)
    switch ("hello" + "") {
    case "foo bar":
        var v93 float32 = (((3.14159 / 1.5) - 0.0) - 1.5)
        _ = v93
    case "":
        if (!("a" == ((v5 + (("hello" + "hello") + v5)) + ("foo bar" + v5)))) {
            var v94 float64 = 3.14159
            _ = v94
        } else {
            var v95 uint = 837
            v3 -= v3
            _ = v95
        }
    case "hello":
        v96 := GenS13{F0: v5, F1: -9223372036854775808}
        _ = v96
    }
    switch v5 {
    case "":
        v1 = (v1 & 18446744073709551615)
    case "foo bar":
        var v97 int = -931
        _ = v97
    }
    v98 := GenS14{F0: v3}
    v99 := GenS15{F0: false, F1: -945, F2: 2147483648}
    for fi100 := 0; fi100 < 8; fi100++ {
        v101 := GenS14{F0: v3}
        _ = fi100
        _ = v101
    }
    for fi102 := 0; fi102 < 8; fi102++ {
        for fi103 := 0; fi103 < 18; fi103++ {
            var v104 int16 = -1
            var v105 int16 = -32768
            v104 = (v104 >> (((v1 | (v1 << ((v1) & 63)))) & 63))
            _ = fi103
            _ = v104
            _ = v105
        }
        _ = fi102
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
    _ = v99
}
