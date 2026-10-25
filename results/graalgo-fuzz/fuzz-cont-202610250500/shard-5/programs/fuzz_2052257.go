package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 uint32; F2 bool }
type GenS9 struct { F0 int; F1 int; F2 uint64 }
type GenS10 struct { F0 uint32; F1 uint16 }
type GenS11 struct { F0 uint64; F1 int32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint) int16 {
    var v4 int32 = -2147483648
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 float32 = 1e10
        var v7 uint16 = 32767
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{32768, (^(fi8)), -500, int(p0)}
        if (((p0 / ((42) | 1)) == 42) && true) {
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint32 = 2147483647
                _ = len(v9)
                switch fi10 {
                case 0:
                    if ((-2.25 / 1.5) != (0.1 / 1.0)) {
                        v11 = 372
                    } else {
                        var v12 float64 = 1.5
                        fmt.Println(v11, v12)
                        _ = v12
                    }
                case 32768:
                    if ((-(p0)) == p0) {
                        v13 := GenS1{F0: 117, F1: 0, F2: (v4 & (v4 >> ((p1) & 63)))}
                        _ = v13
                    }
                case 69:
                    switch fi10 {
                    case -9223372036854775808:
                        v14 := GenS1{F0: 10, F1: 100, F2: (v4 >> ((42) & 63))}
                        _ = v14
                    case 32767:
                        switch "foo bar" {
                        case "a":
                            fmt.Println(fi8, p0)
                        case "\n\t":
                            v15 := GenS1{F0: 128, F1: 0, F2: (v4 - 0)}
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            _ = v15
                        }
                    default:
                        fmt.Println(p0, v4, p1)
                    }
                    var v16 int16 = 0
                    if len(v9) > 0 {
                        v9[0] = (fi8 * int(v16))
                    }
                    _ = v16
                }
                _ = fi10
                _ = v11
            }
            _ = len(v9)
            p0 -= p0
        }
        _ = fi8
        _ = v9
    }
    v17 := GenS2{F0: false}
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    return 32767
}

func main() {
    var v1 uint8 = 157
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    if (!false) {
        for fi18 := 0; fi18 < 18; fi18++ {
            v1 = v1
            for fi19 := 0; fi19 < 12; fi19++ {
                fmt.Println(fi19, v1)
                if (946 <= 0) {
                    if false {
                        if ((v2 >> ((401) & 63)) < 1) {
                            var v20 float32 = -2.25
                            if (v20 >= v20) {
                                if (!true) {
                                    fmt.Println(v2)
                                    v21 := GenS2{F0: true}
                                    v1 -= v1
                                    _ = v21
                                } else {
                                    var v22 int8 = -128
                                    if true {
                                        v23 := []int{fi19}
                                        if false {
                                            v20 = 1.5
                                            fmt.Println(fi18)
                                            v1 += 127
                                        }
                                        v24 := GenS1{F0: (^(v1)), F1: 127, F2: -2147483648}
                                        var v25 int32 = -626
                                        _ = len(v3)
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 uint16 = 42
                                    _ = v22
                                    _ = v26
                                }
                                if true {
                                    fmt.Println(v1, fi19)
                                    var v27 int32 = 42
                                    _ = v27
                                } else {
                                    if ((("foo bar" + (("a" + "a") + "foo bar")) + "hello") == "foo bar") {
                                        v1 += 128
                                        var v28 uint8 = (v1 >> ((1) & 63))
                                        v20 += v20
                                        _ = v28
                                    }
                                    if true {
                                        var v29 int32 = 42
                                        _ = v29
                                    }
                                    v2 += (v2 | 1)
                                }
                                v20 = 1.5
                            } else {
                                var v30 uint64 = v2
                                _ = v30
                            }
                            delete(v3, 907)
                            switch "a" {
                            case "foo bar":
                                if false {
                                    v2 -= (v2 & v2)
                                    var v31 string = "hello"
                                    _ = v31
                                }
                            default:
                                v2 = 762
                            }
                            _ = v20
                        } else {
                            v1 = 128
                            v1 += 42
                            if ((0.1 * 3.14159) >= (1.5 / -2.25)) {
                                v2 = v2
                                var v32 int32 = 311
                                var v33 uint = 0
                                v2 -= 0
                                if (true || true) {
                                    var v34 bool = (!((3.14159 + 3.14159) >= (((float32(0.1) - (0.0 + 0.0)) + (0.0 * -2.25)) + 0.0)))
                                    var v35 bool = ((fi19 | (fi19 * fi18)) == (fi18 >> (((v33 ^ (v33 << ((v33) & 63)))) & 63)))
                                    var v36 string = ""
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                } else {
                                    v32 = v32
                                    v3[(fi18 | int(v33))] = ("hello" + ("\n\t" + ("hello" + "foo bar")))
                                    v37 := []int{(fi18 & (fi19 % (((fi19 / ((fi18) | 1))) | 1)))}
                                    fmt.Println(fi18, v2, fi19)
                                    _ = v37
                                }
                                _ = v32
                                _ = v33
                            } else {
                                v38 := GenS1{F0: 139, F1: -1, F2: 496}
                                v3[fi18] = "a"
                                _ = v38
                            }
                        }
                        v1 -= (v1 - v1)
                    } else {
                        v39 := GenS2{F0: (((("hello" + "foo bar") + ("foo bar" + "a")) + "\n\t") < "")}
                        var v40 float32 = 3.14159
                        _ = v39
                        _ = v40
                    }
                    v2 = 42
                    fmt.Println(fi19, fi19)
                    v2 = (v2 - (v2 << ((2) & 63)))
                }
                fmt.Println(v1, fi18, fi18)
                _ = fi19
            }
            _ = fi18
        }
    } else {
        if true {
            for fi41 := 0; fi41 < 16; fi41++ {
                for fi42 := 0; fi42 < 13; fi42++ {
                    v2 = 1
                    v43 := []int{fi42, (-(fi41)), fi42, 42}
                    switch 32767 {
                    case 4294967295:
                        if len(v43) > 0 {
                            v43[0] = 0
                        }
                        switch (fi42 * fi42) {
                        case -552:
                            if (4294967295 <= v2) {
                                var v44 int = 128
                                var v45 float32 = (3.14159 * 0.0)
                                _ = v44
                                _ = v45
                            } else {
                                v46 := GenS3{F0: 0, F1: 4294967295, F2: 42}
                                _ = v46
                            }
                        case 127:
                            var v47 string = "hello"
                            _ = v47
                        case 0:
                            v48 := GenS3{F0: uint(v1), F1: 2147483648, F2: 65535}
                            v1 += v1
                            var v49 float32 = (0.0 + 1e10)
                            _ = v48
                            _ = v49
                        }
                        v1 += v1
                    case 949:
                        if ((1.5 < float64((3.14159 + 1.5))) && (!(1 > 42))) {
                            v2 = v2
                            var v50 int32 = -2147483648
                            _ = v50
                        } else {
                            var v51 bool = (!false)
                            if (v51 && v51) {
                                var v52 uint = 0
                                if v51 {
                                    if true {
                                        var v53 int64 = 753
                                        _ = v53
                                    } else {
                                        switch "hello" {
                                        case "":
                                            var v54 float32 = 1e10
                                            _ = v54
                                        }
                                        var v55 string = "foo bar"
                                        var v56 float64 = ((0.1 + 0.1) - 1.5)
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 int32 = 2147483647
                                    v2 = 1
                                    if (65535 <= (^(v52))) {
                                        var v58 string = ""
                                        if v51 {
                                            fmt.Println(v52)
                                        } else {
                                            if (v51 && v51) {
                                                switch fi42 {
                                                case 42:
                                                    if len(v43) > 0 {
                                                        v43[0] = fi42
                                                    }
                                                default:
                                                    var v59 float64 = -2.25
                                                    if ((true && ("hello" != (((("" + ((v58 + ((v58 + (("" + v58) + "")) + (((("hello" + v58) + "") + ((("" + v58) + v58) + v58)) + v58))) + "\n\t")) + (v58 + "hello")) + v58) + "\n\t"))) || (v51 || false)) {
                                                        var v60 uint32 = 65535
                                                        if ("\n\t" != "foo bar") {
                                                            var v61 int = (fi41 % (((fi41 << (((^(v52))) & 63))) | 1))
                                                            v62 := GenS3{F0: (v52 + 946), F1: 2147483647, F2: uint(v2)}
                                                            fmt.Println(fi42, v58)
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        v2 -= v2
                                                        _ = v60
                                                    }
                                                    _ = v59
                                                }
                                                delete(v3, fi41)
                                                v57 += (v57 % (((v57 / ((1) | 1))) | 1))
                                            }
                                            var v63 bool = v51
                                            var v64 float64 = 0.1
                                            switch (v58 + (v58 + v58)) {
                                            case "":
                                                switch fi42 {
                                                case 2147483647:
                                                    switch (fi41 ^ 127) {
                                                    case -1:
                                                        switch 418 {
                                                        case 42:
                                                            if len(v43) >= 2 {
                                                                v43 = v43[1:len(v43)]
                                                            }
                                                        case 9223372036854775807:
                                                            var v65 int16 = 1000
                                                            _ = v65
                                                        default:
                                                            var v66 int64 = 42
                                                            _ = v66
                                                        }
                                                        if v63 {
                                                            v58 = (v58 + (v58 + ((v58 + v58) + (v58 + "foo bar"))))
                                                            _ = len(v43)
                                                            v67 := GenS4{F0: (v1 / ((v1) | 1)), F1: 496}
                                                            _ = v67
                                                        } else {
                                                            v2 -= (v2 ^ uint64(v1))
                                                            if v63 {
                                                                if true {
                                                                    v68 := GenS5{F0: v57}
                                                                    _ = len(v43)
                                                                    v2 = (v2 & v2)
                                                                    _ = v68
                                                                } else {
                                                                    if v63 {
                                                                        if len(v43) >= 2 {
                                                                            v43 = v43[1:len(v43)]
                                                                        }
                                                                        var v69 int8 = -128
                                                                        _ = v69
                                                                    } else {
                                                                        v70 := GenS6{F0: (!true)}
                                                                        fmt.Println(v1)
                                                                        if v63 {
                                                                            if ((float64(-2.25) * 0.1) >= v64) {
                                                                                v71 := GenS7{F0: (1 >= genFn0(int64(v1), (v52 % (((v52 << ((v52) & 63))) | 1))))}
                                                                                if (943 <= (v57 % ((v57) | 1))) {
                                                                                    v72 := GenS5{F0: (v57 | (v57 % (((^(v57))) | 1)))}
                                                                                    _ = v72
                                                                                } else {
                                                                                    switch (fi42 << ((1) & 63)) {
                                                                                    case -24:
                                                                                        v64 = v64
                                                                                    default:
                                                                                        fmt.Println(v57, v58, v2)
                                                                                    }
                                                                                    v73 := GenS8{F0: 722, F1: 2147483648, F2: v63}
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        v74 := GenS7{F0: v63}
                                                                                        _ = v74
                                                                                    case "foo bar":
                                                                                        v75 := GenS5{F0: (v57 / (((v57 * v57)) | 1))}
                                                                                        _ = v75
                                                                                    default:
                                                                                        if (!(1 < -776)) {
                                                                                            var v76 int64 = -9223372036854775808
                                                                                            if v63 {
                                                                                                v76 += (-(v76))
                                                                                                if (((v51 && true) && (((v52 << ((v52) & 63)) > uint(v2)) || v63)) || true) {
                                                                                                    var v77 int32 = (^(v57))
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        var v78 int8 = 127
                                                                                                        _ = v78
                                                                                                    }
                                                                                                    _ = v77
                                                                                                }
                                                                                            } else {
                                                                                                var v79 int16 = genFn0((v76 | -773), (v52 - (v52 >> (((v52 * 0)) & 63))))
                                                                                                delete(v3, -1)
                                                                                                if len(v43) > 0 {
                                                                                                    v43[0] = fi41
                                                                                                }
                                                                                                _ = v79
                                                                                            }
                                                                                            var v80 uint8 = v1
                                                                                            _ = v76
                                                                                            _ = v80
                                                                                        } else {
                                                                                            fmt.Println(fi41)
                                                                                        }
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                                _ = v71
                                                                            }
                                                                            switch (("foo bar" + (v58 + "\n\t")) + (v58 + v58)) {
                                                                            case "a":
                                                                                v81 := GenS9{F0: fi41, F1: (-(fi42)), F2: (^(v2))}
                                                                                _ = v81
                                                                            default:
                                                                                var v82 string = ((v58 + (v58 + v58)) + "a")
                                                                                _ = v82
                                                                            }
                                                                            v51 = (true && v63)
                                                                            var v83 int16 = -186
                                                                            v51 = (((v58 + v58) == "\n\t") && (v63 || ((v52 % ((42) | 1)) >= (^(v52)))))
                                                                            _ = v83
                                                                        } else {
                                                                            v84 := GenS1{F0: (^(v1)), F1: 42, F2: (v57 | (v57 >> (((v52 + 1)) & 63)))}
                                                                            v85 := GenS1{F0: (^(v1)), F1: 100, F2: 2147483647}
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                        _ = v70
                                                                    }
                                                                    v51 = ((fi41 + 143) <= fi41)
                                                                    if len(v43) > 0 {
                                                                        v43[0] = fi42
                                                                    }
                                                                }
                                                                if false {
                                                                    v86 := GenS10{F0: 2, F1: 1}
                                                                    if (float64(-2.25) < v64) {
                                                                        if false {
                                                                            var v87 string = "\n\t"
                                                                            if len(v43) > 0 {
                                                                                v43[0] = fi42
                                                                            }
                                                                            switch fi42 {
                                                                            case -708:
                                                                                var v88 uint = v52
                                                                                _ = v88
                                                                            case 0:
                                                                                delete(v3, fi41)
                                                                            case 1:
                                                                                v63 = v51
                                                                            }
                                                                            _ = v87
                                                                        } else {
                                                                            if len(v43) >= 2 {
                                                                                v43 = v43[1:len(v43)]
                                                                            }
                                                                            var v89 string = (v58 + v58)
                                                                            v57 = v57
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                    v64 = float64(1e10)
                                                                    _ = v86
                                                                } else {
                                                                    v90 := GenS11{F0: (v2 - (v2 / ((v2) | 1))), F1: v57}
                                                                    v91 := GenS12{F0: v57}
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                            }
                                                        }
                                                    case -828:
                                                        var v92 string = ""
                                                        var v93 bool = false
                                                        var v94 int64 = -9223372036854775808
                                                        var v95 int16 = -32768
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                        _ = v95
                                                    case -224:
                                                        switch v58 {
                                                        case "\n\t":
                                                            fmt.Println(v64)
                                                        case "hello":
                                                            _ = len(v3)
                                                        case "":
                                                            fmt.Println(v63, v1, v57)
                                                        }
                                                    }
                                                    v96 := GenS13{F0: 311}
                                                    delete(v3, (fi41 << (((v52 + (^(v52)))) & 63)))
                                                    _ = v96
                                                case -110:
                                                    var v97 uint32 = 42
                                                    var v98 int8 = 16
                                                    v99 := GenS12{F0: (^(v57))}
                                                    _ = v97
                                                    _ = v98
                                                    _ = v99
                                                case 267:
                                                    var v100 uint8 = 255
                                                    if ((fi42 > fi41) && ((!(0 != 1)) || ((((v58 != (((v58 + v58) + "hello") + v58)) || v51) && v51) || v51))) {
                                                        switch (fi41 & fi41) {
                                                        case 999:
                                                            v101 := GenS14{F0: (v2 | uint64(v52))}
                                                            _ = v101
                                                        case 2147483647:
                                                            var v102 uint16 = 2
                                                            _ = v102
                                                        case -9223372036854775808:
                                                            if len(v43) >= 2 {
                                                                v43 = v43[1:len(v43)]
                                                            }
                                                        default:
                                                            v64 = (v64 * (((float64(1e10) + v64) / 0.1) * v64))
                                                        }
                                                        var v103 uint8 = v1
                                                        _ = v103
                                                    }
                                                    v63 = v51
                                                    _ = v100
                                                }
                                            case "\n\t":
                                                var v104 uint32 = 44
                                                _ = v104
                                            case "foo bar":
                                                fmt.Println(v64, v52, v64)
                                            default:
                                                v58 = v58
                                            }
                                            var v105 uint = (v52 & (^(v52)))
                                            _ = v105
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = v58
                                    } else {
                                        if (!v51) {
                                            v52 = (^(v52))
                                            if (false || v51) {
                                                v52 = 725
                                            }
                                            v51 = v51
                                        }
                                        if v51 {
                                            var v106 uint8 = 42
                                            v1 -= (v1 | (v1 << ((v52) & 63)))
                                            v2 = v2
                                            _ = v106
                                        }
                                    }
                                    v107 := GenS5{F0: 0}
                                    _ = len(v3)
                                    v108 := GenS15{F0: 262}
                                    _ = v107
                                    _ = v108
                                    _ = v57
                                } else {
                                    v1 = uint8(v52)
                                    if (-9223372036854775808 >= 42) {
                                        v52 += 0
                                    } else {
                                        var v109 uint8 = (^(v1))
                                        v109 += uint8(v2)
                                        v2 = (v2 % ((v2) | 1))
                                        v110 := GenS15{F0: -32768}
                                        _ = v109
                                        _ = v110
                                    }
                                    if v51 {
                                        fmt.Println(v2, v52)
                                        var v111 uint64 = v2
                                        v3[(^(fi41))] = ""
                                        _ = v111
                                    }
                                    v52 = v52
                                }
                                var v112 uint16 = 255
                                _ = v112
                                _ = v52
                            } else {
                                v2 = (v2 & 1)
                                v113 := GenS8{F0: 933, F1: 2, F2: ((("a" + "hello") > "") || v51)}
                                _ = v113
                            }
                            _ = v51
                        }
                    case 524:
                        var v114 float32 = float32(((float64(0.0) * 0.1) * 0.0))
                        _ = v114
                    default:
                        var v115 float64 = (((3.14159 * (-2.25 + 0.1)) / (0.1 / 3.14159)) + float64(((0.0 * 0.0) * float32(0.0))))
                        _ = v115
                    }
                    v116 := GenS16{F0: 647}
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    v117 := GenS17{F0: v1, F1: ""}
                    _ = fi42
                    _ = v116
                    _ = v117
                    _ = v43
                }
                v118 := GenS16{F0: 2}
                _ = fi41
                _ = v118
            }
        }
        var v119 uint32 = 1
        v120 := GenS18{F0: v1}
        _ = v119
        _ = v120
    }
    v2 -= v2
    var v121 uint64 = (v2 << ((4294967295) & 63))
    v121 += v2
    v121 -= 4294967295
    v1 = v1
    if false {
        var v122 uint64 = v121
        _ = v122
    } else {
        v121 = (^(v121))
    }
    fmt.Println(v121)
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v121
    _ = v2
    _ = v3
}
