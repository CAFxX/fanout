package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint8; F2 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 uint32; F2 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 int64; F1 bool }
type GenS13 struct { F0 uint32; F1 uint8 }
type GenS14 struct { F0 int16; F1 int16; F2 string }
type GenS15 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 uint64) uint64 {
    p0 += p0
    v7 := []int{32767, 32768, -1, -375}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint64 = 883
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 bool = (134 > 2)
    var v5 int8 = 100
    v6 := make(map[string]string)
    v5 += (v5 / (((v5 >> ((65535) & 63))) | 1))
    if (255 >= 328) {
        v1 -= 793
    } else {
        var v8 int = -766
        _ = v8
    }
    if (v4 || (1 >= -2147483648)) {
        if ((964 <= 2147483648) || (v1 > genFn0(503, 722, genFn0(2147483647, 822, 42)))) {
            var v9 int = 4294967295
            v5 = (v5 | (v5 >> ((92) & 63)))
            v4 = v4
            _ = v9
        }
        v10 := GenS1{F0: 1, F1: 42}
        _ = v10
    } else {
        var v11 int8 = 0
        v12 := GenS2{F0: (v5 / ((v5) | 1)), F1: 255}
        var v13 uint = uint(v1)
        _ = v11
        _ = v12
        _ = v13
    }
    switch (("foo bar" + "a") + "\n\t") {
    case "\n\t":
        if (!(v5 != (v5 * -128))) {
            v14 := []int{32768, -2147483648}
            var v15 int32 = int32(v1)
            var v16 int8 = 127
            _ = v14
            _ = v15
            _ = v16
        }
    case "hello":
        v1 -= 2
    default:
        var v17 float64 = 3.14159
        _ = v17
    }
    delete(v6, "")
    switch ("hello" + "foo bar") {
    case "":
        switch "" {
        case "foo bar":
            var v18 int = 9223372036854775807
            _ = v18
        case "":
            _ = len(v3)
        }
    case "hello":
        if v4 {
            v4 = v4
            v19 := GenS1{F0: 793, F1: 1}
            var v20 uint16 = 101
            _ = v19
            _ = v20
        } else {
            var v21 uint8 = 1
            if (false || v4) {
                if (false || (v4 || v4)) {
                    switch 32768 {
                    case 32768:
                        if true {
                            if false {
                                v22 := GenS2{F0: (v5 % (((v5 % ((v5) | 1))) | 1)), F1: 65535}
                                var v23 uint16 = 0
                                _ = v22
                                _ = v23
                            }
                        }
                    }
                } else {
                    v4 = false
                    fmt.Println(v1)
                    var v24 float64 = ((3.14159 * (float64((3.14159 / 0.1)) - 0.1)) - 1.5)
                    _ = v24
                }
                v21 = v21
            }
            _ = v21
        }
    case "a":
        var v25 int = 4294967295
        _ = v25
    }
    var v26 int8 = 0
    v27 := GenS3{F0: -32768, F1: -983}
    if (-130 <= -995) {
        switch 32768 {
        case 4294967295:
            if (!(1000 != -1)) {
                var v28 int = 32768
                v4 = true
                fmt.Println(v28, v26, v1)
                _ = v28
            } else {
                var v29 int64 = 9223372036854775807
                if (!v4) {
                    var v30 uint64 = (v1 / (((^(v1))) | 1))
                    if false {
                        var v31 string = "foo bar"
                        _ = v31
                    } else {
                        if (!((!((49 == v5) && true)) && v4)) {
                            delete(v6, "hello")
                        }
                        var v32 uint = 189
                        _ = v32
                    }
                    _ = v30
                } else {
                    if (42 != -1) {
                        for fi33 := 0; fi33 < 11; fi33++ {
                            v34 := GenS4{F0: 42}
                            v35 := []int{}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        var v36 int16 = int16(v29)
                        v1 = (v1 >> ((42) & 63))
                        _ = v36
                    } else {
                        v37 := GenS5{F0: 65535}
                        if (1 != 4294967295) {
                            var v38 int = -1
                            _ = len(v3)
                            if (false && (42 > 2147483647)) {
                                var v39 int = v38
                                var v40 float32 = 0.1
                                _ = v39
                                _ = v40
                            } else {
                                v41 := GenS2{F0: v26, F1: 65535}
                                var v42 uint16 = 42
                                if (64 != 110) {
                                    switch -60 {
                                    case 127:
                                        var v43 uint32 = 42
                                        _ = v43
                                    default:
                                        for fi44 := 0; fi44 < 17; fi44++ {
                                            for fi45 := 0; fi45 < 2; fi45++ {
                                                v46 := GenS1{F0: (fi45 << ((0) & 63)), F1: (^(v42))}
                                                _ = fi45
                                                _ = v46
                                            }
                                            v47 := []int{}
                                            if false {
                                                if ((float64(0.1) * 0.1) > 1e10) {
                                                    fmt.Println(v29, v38)
                                                    var v48 int32 = 456
                                                    var v49 int = -1
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    if false {
                                                        if v4 {
                                                            var v50 int8 = (v5 * (v26 & 42))
                                                            _ = v50
                                                        }
                                                        var v51 float32 = ((0.1 / ((1.5 * (float32(0.0) * float32(-2.25))) + (((1.5 + 0.1) - 1e10) * -2.25))) / 1.0)
                                                        fmt.Println(v51, v38)
                                                        _ = v51
                                                    }
                                                    if v4 {
                                                        fmt.Println(fi44)
                                                    } else {
                                                        v1 = genFn0(42, (v42 ^ (v42 >> ((610) & 63))), v1)
                                                        fmt.Println(v29, v4, v4)
                                                        var v52 float64 = (0.1 - 1e10)
                                                        _ = v52
                                                    }
                                                    if (("" + ("hello" + ("hello" + "hello"))) > "foo bar") {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v53 uint = 255
                                                            _ = v53
                                                        case "\n\t":
                                                            _ = len(v47)
                                                        case "foo bar":
                                                            var v54 int8 = (v5 - v26)
                                                            _ = v54
                                                        }
                                                        for fi55 := 0; fi55 < 6; fi55++ {
                                                            var v56 bool = ((v5 - (v5 >> ((4294967295) & 63))) > (v5 << ((4294967295) & 63)))
                                                            if (v42 > v42) {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v57 uint8 = 65
                                                                    _ = v57
                                                                case "\n\t":
                                                                    var v58 float32 = (0.1 / -2.25)
                                                                    _ = v58
                                                                }
                                                                var v59 float32 = -2.25
                                                                _ = v59
                                                            } else {
                                                                var v60 uint64 = 42
                                                                if (((v5 | v26) == (v5 * (v26 / (((v26 + 127)) | 1)))) && (255 > 128)) {
                                                                    var v61 int8 = (^(v26))
                                                                    _ = v61
                                                                } else {
                                                                    if (!false) {
                                                                        if (true || v56) {
                                                                            v1 -= v60
                                                                            if (127 != 1) {
                                                                                if (1e10 > 3.14159) {
                                                                                    if true {
                                                                                        var v62 bool = (203 > 2147483647)
                                                                                        _ = v62
                                                                                    } else {
                                                                                        v63 := GenS6{F0: 535, F1: v60}
                                                                                        v64 := GenS6{F0: 1000, F1: v1}
                                                                                        _ = len(v2)
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    }
                                                                                    fmt.Println(fi44)
                                                                                }
                                                                                v38 = v38
                                                                            } else {
                                                                                switch 32767 {
                                                                                case 128:
                                                                                    fmt.Println(v5)
                                                                                default:
                                                                                    if v4 {
                                                                                        v47 = append(v47, (fi55 / ((fi55) | 1)))
                                                                                        if v4 {
                                                                                            v26 = (-(v26))
                                                                                            if v4 {
                                                                                                v65 := GenS2{F0: (v5 << ((2) & 63)), F1: v42}
                                                                                                v66 := GenS7{F0: v5, F1: 2, F2: 1}
                                                                                                v42 = v42
                                                                                                _ = v65
                                                                                                _ = v66
                                                                                            }
                                                                                        } else {
                                                                                            if ((!false) || v56) {
                                                                                                if ((0 >= 424) || v56) {
                                                                                                    v4 = v4
                                                                                                    v42 -= v42
                                                                                                    var v67 float32 = 0.1
                                                                                                    _ = v67
                                                                                                } else {
                                                                                                    if (1 <= 42) {
                                                                                                        switch "a" {
                                                                                                        case "foo bar":
                                                                                                            if v4 {
                                                                                                                var v68 uint16 = v42
                                                                                                                if v56 {
                                                                                                                    switch -1 {
                                                                                                                    case 70:
                                                                                                                        var v69 uint8 = 255
                                                                                                                        _ = v69
                                                                                                                    case 32768:
                                                                                                                        if len(v47) >= 2 {
                                                                                                                            v47 = v47[1:len(v47)]
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v70 int8 = 100
                                                                                                                        _ = v70
                                                                                                                    }
                                                                                                                    var v71 float32 = -2.25
                                                                                                                    var v72 uint8 = 128
                                                                                                                    _ = v71
                                                                                                                    _ = v72
                                                                                                                } else {
                                                                                                                    v47 = append(v47, (fi55 / ((fi55) | 1)))
                                                                                                                    var v73 uint16 = (v68 >> ((18446744073709551615) & 63))
                                                                                                                    _ = v73
                                                                                                                }
                                                                                                                v74 := v3[("\n\t" + "a")]
                                                                                                                _ = v68
                                                                                                                _ = v74
                                                                                                            }
                                                                                                            v5 -= int8(fi55)
                                                                                                            switch "hello" {
                                                                                                            case "foo bar":
                                                                                                                delete(v6, "hello")
                                                                                                            case "":
                                                                                                                switch fi44 {
                                                                                                                case -189:
                                                                                                                    var v75 int64 = (v29 + v29)
                                                                                                                    _ = v75
                                                                                                                case -2147483648:
                                                                                                                    v26 += (v5 >> ((18446744073709551615) & 63))
                                                                                                                case 32767:
                                                                                                                    switch (("\n\t" + ("" + "foo bar")) + ("hello" + "hello")) {
                                                                                                                    case "a":
                                                                                                                        if (127 == uint8(v60)) {
                                                                                                                            v4 = v56
                                                                                                                            v29 += (v29 >> ((2) & 63))
                                                                                                                        } else {
                                                                                                                            v76 := GenS8{F0: -1}
                                                                                                                            v77 := v3["\n\t"]
                                                                                                                            _ = v76
                                                                                                                            _ = v77
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        v26 = (v26 << ((1) & 63))
                                                                                                                    case "":
                                                                                                                        var v78 bool = ((1e10 < (((-2.25 - 1e10) - 0.1) / 1e10)) && (!(1.5 >= (1e10 / ((0.1 * float64((((3.14159 + (3.14159 - 1.5)) + 1e10) / 1.5))) + (float64(0.0) / (0.0 - (3.14159 * 1e10))))))))
                                                                                                                        _ = v78
                                                                                                                    default:
                                                                                                                        v38 = fi44
                                                                                                                    }
                                                                                                                }
                                                                                                                switch "\n\t" {
                                                                                                                case "\n\t":
                                                                                                                    v60 = v1
                                                                                                                default:
                                                                                                                    v79 := GenS9{F0: 128, F1: 42, F2: "a"}
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v42 = v42
                                                                                                            if v4 {
                                                                                                                v56 = (int32(v26) >= 2147483647)
                                                                                                                var v80 int64 = v29
                                                                                                                v60 = v1
                                                                                                                _ = len(v6)
                                                                                                                _ = v80
                                                                                                            }
                                                                                                            if len(v47) >= 2 {
                                                                                                                v47 = v47[1:len(v47)]
                                                                                                            }
                                                                                                            if len(v47) >= 2 {
                                                                                                                v47 = v47[1:len(v47)]
                                                                                                            }
                                                                                                        case "":
                                                                                                            var v81 uint8 = 42
                                                                                                            _ = v81
                                                                                                        default:
                                                                                                            v82 := GenS7{F0: (v26 - (v26 / (((v26 | 0)) | 1))), F1: 70, F2: 1000}
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        fmt.Println(v4)
                                                                                                    }
                                                                                                    v56 = (true && (-871 <= 0))
                                                                                                }
                                                                                                if len(v47) >= 2 {
                                                                                                    v47 = v47[1:len(v47)]
                                                                                                }
                                                                                                v60 += (^(v1))
                                                                                                var v83 int8 = -32
                                                                                                if len(v47) >= 2 {
                                                                                                    v47 = v47[1:len(v47)]
                                                                                                }
                                                                                                _ = v83
                                                                                            } else {
                                                                                                v84 := GenS2{F0: v26, F1: 65535}
                                                                                                _ = v84
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if (2147483647 == 636) {
                                                                                            v1 = 0
                                                                                        } else {
                                                                                            if v4 {
                                                                                                v85 := GenS8{F0: 1}
                                                                                                _ = v85
                                                                                            } else {
                                                                                                if v56 {
                                                                                                    v5 += (v5 >> ((554) & 63))
                                                                                                    v86 := GenS10{F0: 2}
                                                                                                    _ = v86
                                                                                                } else {
                                                                                                    v60 = v60
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v38)
                                                                                    }
                                                                                }
                                                                                var v87 int16 = 1
                                                                                if true {
                                                                                    var v88 uint16 = 0
                                                                                    _ = v88
                                                                                } else {
                                                                                    if true {
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            switch v38 {
                                                                                            case -1:
                                                                                                fmt.Println(v56)
                                                                                            case 9223372036854775807:
                                                                                                v89 := v2[2147483648]
                                                                                                _ = v89
                                                                                            case -647:
                                                                                                if (!(("" + (("foo bar" + "hello") + (("foo bar" + ("a" + "a")) + ("a" + (("foo bar" + ("\n\t" + "hello")) + "foo bar"))))) != "a")) {
                                                                                                    var v90 int = 9223372036854775807
                                                                                                    _ = v90
                                                                                                } else {
                                                                                                    var v91 float32 = (float32((float64((1e10 - 3.14159)) + 1e10)) / 1.0)
                                                                                                    _ = v91
                                                                                                }
                                                                                            default:
                                                                                                if (true && (!(((v29 + (v29 & (-(v29)))) == (v29 * v29)) || true))) {
                                                                                                    var v92 int16 = 1
                                                                                                    _ = v92
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            var v93 int8 = (v26 - (v26 ^ (v26 % ((-119) | 1))))
                                                                                            _ = v93
                                                                                        default:
                                                                                            switch (fi55 / ((32767) | 1)) {
                                                                                            case 9223372036854775807:
                                                                                                v42 += (v42 + v42)
                                                                                            case -1:
                                                                                                var v94 int16 = v87
                                                                                                _ = v94
                                                                                            default:
                                                                                                v95 := GenS6{F0: 1, F1: v60}
                                                                                                _ = v95
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v96 uint = 388
                                                                                        _ = v96
                                                                                    }
                                                                                    var v97 int = (fi55 | (fi55 * 32767))
                                                                                    _ = v97
                                                                                }
                                                                                v4 = v4
                                                                                _ = v87
                                                                            }
                                                                        } else {
                                                                            v47 = append(v47, (fi44 << ((255) & 63)))
                                                                            v29 += v29
                                                                        }
                                                                        if (722 < 872) {
                                                                            fmt.Println(v42, v42, v29)
                                                                            v98 := GenS3{F0: 0, F1: (fi44 | v38)}
                                                                            v99 := v2[65535]
                                                                            _ = v98
                                                                            _ = v99
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                _ = len(v2)
                                                                            case "a":
                                                                                v100 := GenS10{F0: 42}
                                                                                _ = v100
                                                                            default:
                                                                                v47 = append(v47, fi55)
                                                                            }
                                                                            var v101 uint8 = 1
                                                                            v1 = genFn0(241, 2, v60)
                                                                            var v102 int16 = -32768
                                                                            if (!false) {
                                                                                if v56 {
                                                                                    fmt.Println(v56)
                                                                                    var v103 uint16 = uint16(v101)
                                                                                    _ = v103
                                                                                }
                                                                                var v104 bool = (0.1 == ((float32(1e10) * float32(0.1)) / -2.25))
                                                                                v26 = v26
                                                                                _ = v104
                                                                            } else {
                                                                                if (!false) {
                                                                                    if true {
                                                                                        if v4 {
                                                                                            v4 = v56
                                                                                            var v105 int32 = 1
                                                                                            v106 := GenS7{F0: (v26 - (v26 - v26)), F1: v101, F2: v102}
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        } else {
                                                                                            var v107 uint64 = v1
                                                                                            v108 := GenS8{F0: (v102 + v102)}
                                                                                            v1 = v60
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                        var v109 uint = uint(v101)
                                                                                        v110 := GenS11{F0: ("foo bar" + "foo bar"), F1: (v109 & v109)}
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                    }
                                                                                } else {
                                                                                    var v111 uint8 = v101
                                                                                    fmt.Println(v60, v4)
                                                                                    fmt.Println(v60, v101, v26)
                                                                                    _ = v111
                                                                                }
                                                                            }
                                                                            _ = v101
                                                                            _ = v102
                                                                        }
                                                                    } else {
                                                                        var v112 int = -2147483648
                                                                        _ = v112
                                                                    }
                                                                    v2[0] = 2147483648
                                                                    v60 = 42
                                                                    v113 := GenS12{F0: (v29 % (((v29 | v29)) | 1)), F1: v4}
                                                                    if (2 > 1) {
                                                                        switch "" {
                                                                        case "":
                                                                            var v114 uint32 = 42
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                    v3["a"] = 2147483647
                                                                    _ = v113
                                                                }
                                                                _ = v60
                                                            }
                                                            var v115 float64 = 1.5
                                                            if (-79 < 32767) {
                                                                fmt.Println(v1)
                                                                v116 := GenS13{F0: 0, F1: 127}
                                                                _ = len(v47)
                                                                _ = v116
                                                            }
                                                            _ = fi55
                                                            _ = v115
                                                            _ = v56
                                                        }
                                                    } else {
                                                        switch (("hello" + ("a" + ("foo bar" + "a"))) + "\n\t") {
                                                        case "\n\t":
                                                            v117 := GenS2{F0: v26, F1: v42}
                                                            _ = v117
                                                        case "foo bar":
                                                            v42 = 209
                                                        }
                                                        if true {
                                                            var v118 int64 = (v29 * int64(fi44))
                                                            switch "foo bar" {
                                                            case "":
                                                                v5 = (v5 + (v26 * -109))
                                                            default:
                                                                var v119 string = ""
                                                                _ = v119
                                                            }
                                                            _ = v118
                                                        } else {
                                                            v120 := GenS7{F0: (v26 ^ v26), F1: 255, F2: -1}
                                                            if len(v47) >= 2 {
                                                                v47 = v47[1:len(v47)]
                                                            }
                                                            _ = v120
                                                        }
                                                    }
                                                }
                                                var v121 uint16 = 282
                                                v38 = v38
                                                _ = v121
                                            } else {
                                                if v4 {
                                                    var v122 bool = false
                                                    switch "" {
                                                    case "a":
                                                        v38 -= -9223372036854775808
                                                    case "":
                                                        v1 = v1
                                                    }
                                                    v123 := GenS13{F0: 833, F1: 127}
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    if len(v47) >= 2 {
                                                        v47 = v47[1:len(v47)]
                                                    }
                                                }
                                                v26 = -86
                                            }
                                            _ = fi44
                                            _ = v47
                                        }
                                    }
                                    for fi124 := 0; fi124 < 14; fi124++ {
                                        v4 = (65535 < 1)
                                        v125 := GenS4{F0: 2}
                                        _ = fi124
                                        _ = v125
                                    }
                                    for fi126 := 0; fi126 < 12; fi126++ {
                                        for fi127 := 0; fi127 < 18; fi127++ {
                                            var v128 int16 = -1
                                            _ = fi127
                                            _ = v128
                                        }
                                        v1 = 0
                                        for fi129 := 0; fi129 < 7; fi129++ {
                                            if true {
                                                v26 += (v26 & v5)
                                                fmt.Println(v29, fi126, v38)
                                                var v130 uint64 = (^(v1))
                                                _ = v130
                                            } else {
                                                v3[("a" + "hello")] = 2147483648
                                                var v131 float64 = (3.14159 - 0.0)
                                                _ = v131
                                            }
                                            v4 = (int8(fi129) > 42)
                                            _ = fi129
                                        }
                                        _ = fi126
                                    }
                                }
                                _ = v41
                                _ = v42
                            }
                            _ = v38
                        }
                        switch 0 {
                        case 2147483647:
                            switch 575 {
                            case 127:
                                switch "foo bar" {
                                case "a":
                                    switch 9223372036854775807 {
                                    case 42:
                                        v26 -= v5
                                    case -2147483648:
                                        fmt.Println(v1, v1)
                                    }
                                case "hello":
                                    for fi132 := 0; fi132 < 14; fi132++ {
                                        v5 = 100
                                        _ = fi132
                                    }
                                default:
                                    v5 -= (-(v26))
                                }
                            case 0:
                                fmt.Println(v26, v29, v4)
                            case 32768:
                                delete(v2, 65535)
                            default:
                                var v133 int = 0
                                _ = v133
                            }
                        case -1:
                            if (-128 < v26) {
                                _ = len(v3)
                                switch -434 {
                                case 48:
                                    _ = len(v2)
                                case 1:
                                    var v134 float64 = (0.0 * 1.5)
                                    _ = v134
                                case -9223372036854775808:
                                    v26 -= v5
                                default:
                                    switch -1 {
                                    case -703:
                                        if (-67 == 0) {
                                            for fi135 := 0; fi135 < 6; fi135++ {
                                                switch 32768 {
                                                case 128:
                                                    if true {
                                                        v136 := GenS14{F0: 1, F1: 93, F2: ("a" + "\n\t")}
                                                        _ = v136
                                                    } else {
                                                        v3["foo bar"] = 2147483648
                                                        var v137 float64 = ((0.1 + -2.25) / 1.0)
                                                        _ = v137
                                                    }
                                                case 32768:
                                                    delete(v6, "\n\t")
                                                default:
                                                    if true {
                                                        switch ("\n\t" + ("" + ("foo bar" + "foo bar"))) {
                                                        case "\n\t":
                                                            v4 = (4294967295 < 42)
                                                        }
                                                    }
                                                }
                                                _ = fi135
                                            }
                                            for fi138 := 0; fi138 < 4; fi138++ {
                                                v139 := GenS13{F0: 601, F1: 42}
                                                _ = fi138
                                                _ = v139
                                            }
                                            if ("foo bar" <= "\n\t") {
                                                v29 = v29
                                                if v4 {
                                                    var v140 float32 = 1e10
                                                    delete(v6, ("a" + ("\n\t" + ("\n\t" + "foo bar"))))
                                                    _ = v140
                                                }
                                                if v4 {
                                                    var v141 int16 = 982
                                                    var v142 uint32 = 238
                                                    var v143 bool = (1 >= -1)
                                                    _ = v141
                                                    _ = v142
                                                    _ = v143
                                                } else {
                                                    var v144 int = 127
                                                    v145 := []int{-9223372036854775808, 9223372036854775807, 1}
                                                    _ = v144
                                                    _ = v145
                                                }
                                            }
                                        }
                                    case 128:
                                        var v146 string = ("\n\t" + "\n\t")
                                        _ = v146
                                    }
                                }
                            }
                        default:
                            var v147 uint16 = 728
                            _ = v147
                        }
                        _ = v37
                    }
                }
                _ = v29
            }
        default:
            v4 = v4
        }
        v1 -= (v1 ^ (^(v1)))
    } else {
        v5 = (^(v5))
    }
    fmt.Println(v4, v4)
    for fi148 := 0; fi148 < 7; fi148++ {
        var v149 uint16 = 2
        var v150 string = "\n\t"
        _ = fi148
        _ = v149
        _ = v150
    }
    v151 := GenS11{F0: "foo bar", F1: 18446744073709551615}
    for fi152 := 0; fi152 < 6; fi152++ {
        v153 := GenS4{F0: 2}
        for fi154 := 0; fi154 < 13; fi154++ {
            var v155 float64 = 0.1
            v156 := []int{fi154}
            _ = fi154
            _ = v155
            _ = v156
        }
        for fi157 := 0; fi157 < 5; fi157++ {
            v158 := GenS15{F0: 1, F1: 0, F2: (v1 & genFn0(2147483648, 255, genFn0(65535, 732, v1)))}
            v159 := []int{32768}
            var v160 int = fi152
            var v161 float64 = 1.5
            _ = len(v159)
            var v162 bool = (true && v4)
            _ = fi157
            _ = v158
            _ = v159
            _ = v160
            _ = v161
            _ = v162
        }
        _ = fi152
        _ = v153
    }
    v4 = (!v4)
    fmt.Println(v26, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v151
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
