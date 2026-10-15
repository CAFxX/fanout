package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint16; F2 string }
type GenS5 struct { F0 int16; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint32; F2 bool }
type GenS7 struct { F0 bool; F1 int; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 uint, p1 int16) uint64 {
    v6 := GenS1{F0: 108}
    p0 += p0
    if true {
        var v7 int64 = 42
        var v8 uint64 = 18446744073709551615
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (fi9 * (fi9 - (fi9 / (((^(fi9))) | 1)))) {
            case 2147483647:
                var v10 uint64 = v8
                _ = v10
            case 4294967295:
                var v11 string = ("a" + "a")
                _ = v11
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}
func genFn1() uint64 {
    var v12 bool = ((("a" + "hello") + "") != ("a" + ""))
    var v13 uint16 = 905
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []int{349, 32767}
    var v2 uint64 = 4294967295
    var v3 int64 = -9223372036854775808
    var v4 bool = ((("\n\t" + (("" + "foo bar") + "")) < "hello") || ((-71 != 0) && (2 >= 414)))
    v5 := []string{""}
    var v14 uint32 = 384
    switch -802 {
    case 19:
        v3 = 0
    }
    v14 = v14
    if (1 != 65535) {
        v15 := GenS2{F0: 2147483647, F1: 2}
        v4 = false
        v2 = (v2 ^ genFn0(4294967295, 1000))
        _ = v15
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                switch (((("hello" + "hello") + "a") + ("" + ("hello" + "foo bar"))) + "foo bar") {
                case "foo bar":
                    var v18 int64 = (v3 << ((255) & 63))
                    _ = v18
                case "\n\t":
                    v19 := GenS2{F0: -301, F1: 1}
                    _ = v19
                case "":
                    v2 = genFn1()
                default:
                    var v20 float32 = (((((1e10 - (1.5 * 0.0)) - float32(0.0)) + ((3.14159 * -2.25) - float32((1.5 / float64(-2.25))))) * float32(float64(float32(float64(-2.25))))) - float32(0.0))
                    var v21 uint = 0
                    v1 = append(v1, fi16)
                    _ = v20
                    _ = v21
                }
                if v4 {
                    v14 -= (^(v14))
                } else {
                    if ((0 >= 2) || (fi17 <= (fi17 & (fi16 >> ((42) & 63))))) {
                        if v4 {
                            if (!(601 < 32767)) {
                                if true {
                                    if ((v14 >> ((199) & 63)) == v14) {
                                        var v22 bool = (!true)
                                        switch "hello" {
                                        case "a":
                                            v23 := GenS1{F0: 128}
                                            _ = v23
                                        case "hello":
                                            var v24 string = (("" + (("hello" + "a") + "a")) + "")
                                            _ = v24
                                        case "\n\t":
                                            if (v22 || (1 > 237)) {
                                                _ = len(v1)
                                                v14 = (v14 | (v14 - v14))
                                                var v25 float64 = -2.25
                                                var v26 int32 = 2147483647
                                                var v27 float64 = (float64(0.1) * 0.1)
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = ""
                                        }
                                        _ = v22
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = (fi17 / (((fi16 % ((fi16) | 1))) | 1))
                                        }
                                    }
                                    if v4 {
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(v2)
                                            v14 = (^(v14))
                                        case "a":
                                            v1 = append(v1, 725)
                                        case "foo bar":
                                            fmt.Println(fi17)
                                        default:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        switch ("foo bar" + ("foo bar" + ((("" + "") + ("\n\t" + "")) + ("" + "")))) {
                                        case "":
                                            v28 := GenS3{F0: 255}
                                            _ = v28
                                        }
                                        if (v14 == v14) {
                                            v14 = 750
                                        } else {
                                            var v29 bool = v4
                                            var v30 int32 = 0
                                            if (!(v14 != (^(v14)))) {
                                                v29 = v29
                                                v31 := GenS4{F0: v14, F1: 255, F2: ("" + "a")}
                                                var v32 int8 = 67
                                                _ = v31
                                                _ = v32
                                            } else {
                                                if true {
                                                    if ((!((v3 - -661) != 42)) && (v30 >= v30)) {
                                                        var v33 float32 = -2.25
                                                        fmt.Println(v3)
                                                        _ = v33
                                                    }
                                                    switch ("hello" + ("\n\t" + "hello")) {
                                                    case "\n\t":
                                                        v5 = append(v5, ("a" + ((("\n\t" + ("a" + ("hello" + "a"))) + "hello") + ("" + "foo bar"))))
                                                    case "":
                                                        var v34 bool = (273 >= 471)
                                                        _ = v34
                                                    case "a":
                                                        if v4 {
                                                            v14 -= (v14 ^ (v14 - 197))
                                                            var v35 uint8 = 1
                                                            var v36 uint32 = (^(v14))
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    default:
                                                        var v37 uint8 = 118
                                                        _ = v37
                                                    }
                                                    var v38 bool = ((v3 - (^(v3))) >= v3)
                                                    var v39 int64 = (-(v3))
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    case "":
                                                        switch 42 {
                                                        case -381:
                                                            v3 = -1
                                                        }
                                                    case "\n\t":
                                                        v40 := GenS1{F0: 255}
                                                        _ = v40
                                                    }
                                                    v41 := GenS2{F0: 1, F1: 65535}
                                                    v42 := GenS3{F0: 4294967295}
                                                    _ = v38
                                                    _ = v39
                                                    _ = v41
                                                    _ = v42
                                                }
                                                v43 := GenS3{F0: 302}
                                                if v4 {
                                                    v4 = (fi17 != -343)
                                                    var v44 int32 = (v30 ^ v30)
                                                    fmt.Println(fi16)
                                                    v3 += v3
                                                    _ = v44
                                                }
                                                var v45 int16 = -651
                                                _ = v43
                                                _ = v45
                                            }
                                            var v46 string = ("hello" + "\n\t")
                                            switch v46 {
                                            case "":
                                                var v47 bool = (!v4)
                                                _ = v47
                                            case "foo bar":
                                                if v4 {
                                                    if (!(v4 && v29)) {
                                                        v30 = (-(v30))
                                                    }
                                                    var v48 float64 = (0.1 * ((0.1 + (1.5 - 0.1)) / 1.0))
                                                    _ = v48
                                                } else {
                                                    v4 = (v29 || v29)
                                                }
                                            default:
                                                var v49 int8 = -88
                                                _ = v49
                                            }
                                            if v4 {
                                                if (float64((0.1 + 1.5)) != 0.1) {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if (-128 > -128) {
                                                            var v50 int8 = 127
                                                            fmt.Println(v29, v46)
                                                            v51 := GenS4{F0: (v14 & v14), F1: 32767, F2: "a"}
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            v29 = (-96 == int8(v3))
                                                            _ = len(v1)
                                                            v52 := GenS2{F0: (v30 + v30), F1: 255}
                                                            _ = v52
                                                        }
                                                    case "a":
                                                        v5 = append(v5, "foo bar")
                                                    }
                                                }
                                                v53 := GenS5{F0: -33, F1: 100, F2: (v14 >> ((1) & 63))}
                                                fmt.Println(fi17, v3, v4)
                                                v29 = (32767 != 42)
                                                v54 := GenS2{F0: (^(v30)), F1: 0}
                                                _ = v53
                                                _ = v54
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = (^(fi16))
                                                }
                                                var v55 int32 = v30
                                                _ = v55
                                            }
                                            v56 := GenS3{F0: 255}
                                            v14 = 1
                                            var v57 int = fi16
                                            v29 = (true && (42 != 1000))
                                            var v58 string = "hello"
                                            v4 = v29
                                            _ = v29
                                            _ = v30
                                            _ = v46
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                        }
                                    } else {
                                        if v4 {
                                            v14 = 490
                                        }
                                        v59 := GenS4{F0: v14, F1: 575, F2: ""}
                                        _ = v59
                                    }
                                    var v60 int16 = 1000
                                    _ = v60
                                }
                            }
                        } else {
                            if v4 {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    var v61 int16 = 42
                                    v61 -= v61
                                    v14 += (v14 & 2147483647)
                                    _ = v61
                                } else {
                                    switch 32768 {
                                    case 32767:
                                        fmt.Println(v2)
                                    case -2147483648:
                                        switch ("\n\t" + "") {
                                        case "":
                                            v4 = (!(!true))
                                        case "hello":
                                            if ((v14 < (v14 * v14)) || v4) {
                                                var v62 uint64 = (v2 / ((v2) | 1))
                                                _ = v62
                                            }
                                        }
                                    }
                                    fmt.Println(fi17, v2, v4)
                                    switch fi17 {
                                    case 42:
                                        v2 = v2
                                    }
                                }
                            } else {
                                var v63 uint8 = 2
                                v4 = (false && (v4 && (v4 && (float32((-2.25 - 1.5)) > 3.14159))))
                                if ((v14 % ((v14) | 1)) == 2) {
                                    fmt.Println(v2)
                                    v3 = (v3 & 42)
                                    if false {
                                        var v64 int32 = 1
                                        _ = v64
                                    } else {
                                        var v65 int32 = 2147483647
                                        var v66 int8 = 1
                                        if v4 {
                                            if (255 < 42) {
                                                var v67 string = "\n\t"
                                                var v68 string = (v67 + "foo bar")
                                                _ = v67
                                                _ = v68
                                            }
                                            v63 = 85
                                            var v69 uint8 = (v63 + (v63 ^ v63))
                                            if (65535 >= 65535) {
                                                fmt.Println(v65, v65, fi16)
                                                if v4 {
                                                    var v70 int8 = v66
                                                    v66 += 71
                                                    var v71 int8 = (v70 | 1)
                                                    _ = v70
                                                    _ = v71
                                                }
                                                if (1000 == 1000) {
                                                    if true {
                                                        if len(v5) > 0 {
                                                            v5[0] = "\n\t"
                                                        }
                                                        var v72 uint64 = (v2 / ((4294967295) | 1))
                                                        if v4 {
                                                            switch 32767 {
                                                            case 127:
                                                                var v73 string = "foo bar"
                                                                _ = v73
                                                            case 128:
                                                                var v74 string = ("foo bar" + ("a" + "hello"))
                                                                _ = v74
                                                            case 2147483647:
                                                                v2 -= genFn0(42, 888)
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v75 string = "hello"
                                                            _ = v75
                                                        } else {
                                                            var v76 int32 = (v65 | v65)
                                                            _ = len(v5)
                                                            _ = v76
                                                        }
                                                        _ = v72
                                                    }
                                                    v77 := GenS3{F0: 42}
                                                    fmt.Println(v2, v4)
                                                    var v78 uint16 = 685
                                                    if (v4 && ((v4 || false) || (174 <= v63))) {
                                                        if ((v4 || v4) && (int16(fi16) > -1)) {
                                                            v78 = (v78 / (((v78 | 2)) | 1))
                                                            v79 := GenS2{F0: (v65 / (((v65 - v65)) | 1)), F1: v78}
                                                            fmt.Println(v63, v65)
                                                            var v80 int64 = 0
                                                            _ = v79
                                                            _ = v80
                                                        } else {
                                                            v81 := GenS6{F0: 44, F1: (v14 ^ (v14 ^ 2147483647)), F2: (!v4)}
                                                            v66 = v66
                                                            if true {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (-(fi16))
                                                                }
                                                                var v82 uint = 519
                                                                if v4 {
                                                                    var v83 uint16 = (v78 & 490)
                                                                    var v84 uint32 = uint32(v2)
                                                                    if v4 {
                                                                        var v85 uint8 = v63
                                                                        v3 -= -1
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        _ = v85
                                                                    }
                                                                    var v86 int64 = v3
                                                                    _ = v83
                                                                    _ = v84
                                                                    _ = v86
                                                                } else {
                                                                    if (int16(v65) == 1000) {
                                                                        var v87 float64 = 0.1
                                                                        if false {
                                                                            v88 := GenS3{F0: (v82 << ((255) & 63))}
                                                                            if v4 {
                                                                                var v89 uint8 = v69
                                                                                v82 -= v82
                                                                                _ = v89
                                                                            } else {
                                                                                v5 = append(v5, (("" + ("hello" + "")) + "hello"))
                                                                                v90 := GenS1{F0: (v63 + (v69 * v69))}
                                                                                v87 = v87
                                                                                _ = v90
                                                                            }
                                                                            var v91 int64 = v3
                                                                            var v92 uint32 = (v14 & (^(v14)))
                                                                            fmt.Println(v69, v2)
                                                                            _ = v88
                                                                            _ = v91
                                                                            _ = v92
                                                                        }
                                                                        if (!v4) {
                                                                            v93 := GenS6{F0: v82, F1: v14, F2: (v4 || ((fi17 == fi16) && (!v4)))}
                                                                            _ = len(v5)
                                                                            if len(v5) > 0 {
                                                                                v5[0] = ("hello" + "a")
                                                                            }
                                                                            _ = v93
                                                                        }
                                                                        _ = v87
                                                                    } else {
                                                                        v94 := GenS3{F0: (v82 | (v82 / ((329) | 1)))}
                                                                        _ = v94
                                                                    }
                                                                    var v95 int32 = v65
                                                                    v14 = (v14 | v14)
                                                                    _ = v95
                                                                }
                                                                _ = v82
                                                            } else {
                                                                var v96 bool = v4
                                                                _ = v96
                                                            }
                                                            fmt.Println(v3, v4)
                                                            _ = v81
                                                        }
                                                        v97 := GenS6{F0: 4294967295, F1: (v14 | (v14 * v14)), F2: v4}
                                                        _ = v97
                                                    }
                                                    v98 := GenS1{F0: (v63 - uint8(v78))}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = ("hello" + "")
                                                    }
                                                    _ = v77
                                                    _ = v78
                                                    _ = v98
                                                } else {
                                                    var v99 float32 = 1.5
                                                    v100 := GenS7{F0: (int16(v66) >= 42), F1: fi16, F2: 100}
                                                    _ = v100
                                                    _ = v99
                                                }
                                                if v4 {
                                                    v101 := GenS3{F0: 1}
                                                    switch ("" + ("foo bar" + ("a" + "hello"))) {
                                                    case "a":
                                                        _ = len(v5)
                                                    default:
                                                        v1 = append(v1, (fi17 % (((fi17 - 32767)) | 1)))
                                                    }
                                                    if (!v4) {
                                                        var v102 uint64 = genFn1()
                                                        var v103 float32 = 0.0
                                                        v69 = v69
                                                        if len(v5) > 0 {
                                                            v5[0] = "foo bar"
                                                        }
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi16 * 2147483647)
                                                    }
                                                    fmt.Println(v14, v3)
                                                    _ = v101
                                                } else {
                                                    v104 := GenS3{F0: 0}
                                                    v65 -= v65
                                                    _ = v104
                                                }
                                            } else {
                                                v14 = 2
                                                v105 := GenS5{F0: -32768, F1: (v66 / ((-128) | 1)), F2: (v14 >> ((469) & 63))}
                                                _ = v105
                                            }
                                            _ = v69
                                        }
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    v106 := GenS2{F0: int32(v2), F1: 0}
                                    if v4 {
                                        var v107 uint16 = 798
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(v3, v14, v14)
                                        case "hello":
                                            fmt.Println(v2, v14, v3)
                                        case "":
                                            if (-2147483648 > 369) {
                                                v2 = v2
                                                v4 = v4
                                                v108 := GenS8{F0: v14, F1: v2}
                                                var v109 uint16 = (v107 | (v107 % ((v107) | 1)))
                                                _ = v108
                                                _ = v109
                                            } else {
                                                v63 += v63
                                            }
                                        default:
                                            switch ("hello" + "\n\t") {
                                            case "\n\t":
                                                _ = len(v5)
                                            case "hello":
                                                var v110 int = fi16
                                                _ = v110
                                            case "a":
                                                if false {
                                                    if ((v14 + uint32(v2)) > (v14 ^ 2147483648)) {
                                                        var v111 float64 = ((1e10 + 0.1) / 1.0)
                                                        v111 += -2.25
                                                        _ = v111
                                                    }
                                                    fmt.Println(v107)
                                                    var v112 uint64 = v2
                                                    _ = v112
                                                }
                                            }
                                            if (v2 > 505) {
                                                var v113 float32 = 0.1
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                switch (fi16 | 344) {
                                                case 2147483647:
                                                    fmt.Println(v107, v107, fi17)
                                                case 0:
                                                    if (!(968 != 2147483647)) {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v2 = (v2 / ((88) | 1))
                                                        _ = len(v1)
                                                    }
                                                default:
                                                    fmt.Println(v63)
                                                }
                                                _ = v113
                                            } else {
                                                var v114 float32 = 3.14159
                                                _ = v114
                                            }
                                            var v115 int32 = 0
                                            _ = v115
                                        }
                                        v107 -= 65535
                                        _ = v107
                                    }
                                    v116 := GenS9{F0: v14, F1: 774}
                                    _ = v106
                                    _ = v116
                                }
                                _ = v63
                            }
                            switch fi16 {
                            case 1:
                                v117 := GenS8{F0: 2147483648, F1: (v2 >> ((983) & 63))}
                                _ = v117
                            case -1:
                                if v4 {
                                    v5 = append(v5, ("" + ("hello" + "foo bar")))
                                    v2 = v2
                                    v118 := GenS1{F0: 128}
                                    _ = v118
                                } else {
                                    v2 += 2
                                    v2 = 823
                                }
                            }
                            v2 += (v2 % (((^(v2))) | 1))
                        }
                        var v119 int64 = -780
                        v3 += 440
                        v119 = -1
                        switch -1 {
                        case 128:
                            fmt.Println(v119)
                        case 32767:
                            switch "" {
                            case "foo bar":
                                fmt.Println(v3, v14)
                            case "":
                                if (0.1 >= 0.0) {
                                    if len(v5) > 0 {
                                        v5[0] = "\n\t"
                                    }
                                    var v120 uint16 = 65535
                                    var v121 uint = uint(v120)
                                    fmt.Println(v119, v4)
                                    _ = v120
                                    _ = v121
                                } else {
                                    v2 = genFn1()
                                }
                            }
                            var v122 int64 = (v119 | 9223372036854775807)
                            if ("hello" >= ("\n\t" + "")) {
                                var v123 uint16 = 415
                                v1 = append(v1, 4294967295)
                                _ = v123
                            }
                            _ = v122
                        }
                        v124 := GenS6{F0: 42, F1: v14, F2: v4}
                        var v125 int64 = 1
                        _ = v119
                        _ = v124
                        _ = v125
                    } else {
                        if (127 == -125) {
                            v3 = v3
                            var v126 int64 = v3
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v126
                        } else {
                            v5 = append(v5, "a")
                        }
                        v14 = 42
                        fmt.Println(v4, v3, fi16)
                    }
                    v127 := GenS9{F0: v14, F1: (v14 % ((v14) | 1))}
                    v128 := GenS10{F0: -1}
                    _ = v127
                    _ = v128
                }
                _ = fi17
            }
            _ = fi16
        }
    }
    v129 := GenS11{F0: ((("a" + (("hello" + "foo bar") + "")) + ("" + "\n\t")) + (("hello" + ((("foo bar" + "foo bar") + (("foo bar" + ("a" + "\n\t")) + "\n\t")) + "")) + "hello"))}
    v4 = ((197 != 2) || ("foo bar" == "foo bar"))
    _ = len(v1)
    for fi130 := 0; fi130 < 7; fi130++ {
        var v131 uint64 = (v2 % ((v2) | 1))
        _ = fi130
        _ = v131
    }
    for fi132 := 0; fi132 < 2; fi132++ {
        if true {
            v133 := GenS12{F0: (v3 + (v3 | (v3 << ((65535) & 63)))), F1: 232}
            v134 := GenS4{F0: (v14 << ((42) & 63)), F1: uint16(v14), F2: "foo bar"}
            switch "" {
            case "":
                v2 = (v2 << ((65535) & 63))
            default:
                for fi135 := 0; fi135 < 13; fi135++ {
                    fmt.Println(fi135, fi132)
                    v2 += (^(v2))
                    _ = fi135
                }
            }
            _ = v133
            _ = v134
        }
        v3 += (v3 ^ 629)
        _ = fi132
    }
    v2 = 0
    var v136 int8 = -94
    if (-454 <= -2147483648) {
        for fi137 := 0; fi137 < 8; fi137++ {
            fmt.Println(v136)
            v14 = 88
            _ = fi137
        }
        switch "\n\t" {
        case "foo bar":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            v136 = v136
        }
        for fi138 := 0; fi138 < 9; fi138++ {
            if len(v1) > 0 {
                v1[0] = int(v3)
            }
            v139 := GenS13{F0: (v2 - v2), F1: (^(v14)), F2: (v3 + v3)}
            v140 := GenS6{F0: 18446744073709551615, F1: (v14 & (v14 % (((v14 & (v14 ^ (v14 * v14)))) | 1))), F2: v4}
            _ = fi138
            _ = v139
            _ = v140
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v136:", v136)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v129
    _ = v136
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
