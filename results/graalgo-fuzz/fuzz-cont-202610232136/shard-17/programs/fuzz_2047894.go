package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint16; F1 int; F2 string }
type GenS6 struct { F0 uint16; F1 int; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int; F2 uint64 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 string }

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    var v5 uint32 = 2
    var v6 int64 = -1
    v6 = v6
    if true {
        switch -379 {
        case -435:
            var v7 uint64 = (^(v1))
            _ = v7
        case 42:
            var v8 uint64 = 0
            _ = v8
        case 1:
            v6 = 178
        }
    } else {
        v9 := v2[-1]
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint16 = 824
        _ = fi10
        _ = v11
    }
    if false {
        if ("\n\t" != "foo bar") {
            if true {
                v12 := GenS1{F0: 42, F1: -825, F2: 122}
                if ((1e10 + float32(float64(3.14159))) == (0.1 / 1.0)) {
                    v13 := []int{127, 32767}
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            if true {
                                var v14 int = 82
                                var v15 int = (v14 / ((v14) | 1))
                                _ = v14
                                _ = v15
                            } else {
                                switch 128 {
                                case 2147483647:
                                    fmt.Println(v1, v5)
                                case -323:
                                    v5 = (v5 * v5)
                                }
                                if ((0 != 1) && true) {
                                    delete(v4, 688)
                                    fmt.Println(v6, v1, v6)
                                } else {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        for fi17 := 0; fi17 < 19; fi17++ {
                                            var v18 float64 = (float64((0.1 / (0.1 / 3.14159))) * 0.0)
                                            var v19 float64 = v18
                                            fmt.Println(fi17, v5, v18)
                                            _ = fi17
                                            _ = v18
                                            _ = v19
                                        }
                                        v1 += (v1 & v1)
                                        v1 = v1
                                        _ = fi16
                                    }
                                    for fi20 := 0; fi20 < 4; fi20++ {
                                        v1 = v1
                                        var v21 int16 = 1
                                        _ = len(v4)
                                        _ = fi20
                                        _ = v21
                                    }
                                }
                                v6 = v6
                            }
                        } else {
                            var v22 float32 = 1.5
                            _ = v22
                        }
                    }
                    _ = v13
                } else {
                    fmt.Println(v5, v6)
                }
                var v23 int32 = -157
                _ = v12
                _ = v23
            }
            switch "" {
            case "":
                var v24 bool = (false || false)
                _ = v24
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                v25 := GenS2{F0: -1, F1: -128, F2: (v1 + (^(v1)))}
                _ = v25
            }
            var v26 uint64 = (v1 / (((v1 - 42)) | 1))
            v1 -= (v26 - 1)
            _ = v26
        }
        delete(v3, ((("foo bar" + (("a" + "a") + "foo bar")) + ("hello" + "a")) + ("\n\t" + "\n\t")))
    } else {
        var v27 int64 = v6
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 float32 = (-2.25 / 3.14159)
            fmt.Println(fi28)
            _ = fi28
            _ = v29
        }
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 11; fi31++ {
                v32 := GenS2{F0: 127, F1: -70, F2: 4294967295}
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                if false {
                    v27 = v27
                    if (!true) {
                        if (1.5 == (-2.25 - 3.14159)) {
                            v27 = v6
                            fmt.Println(v6)
                            v34 := v3[""]
                            _ = v34
                        } else {
                            switch "hello" {
                            case "a":
                                var v35 float64 = (3.14159 / (0.0 - float64((-2.25 * ((float32(1.5) - 3.14159) - 3.14159)))))
                                _ = v35
                            case "":
                                v36 := GenS3{F0: (v5 >> ((42) & 63)), F1: v6}
                                _ = v36
                            case "\n\t":
                                var v37 uint = 18
                                _ = v37
                            }
                        }
                        v1 = v1
                    }
                    var v38 int8 = -1
                    _ = v38
                } else {
                    var v39 float32 = (3.14159 * 1.5)
                    fmt.Println(v39)
                    v40 := GenS2{F0: -114, F1: -1, F2: (^(v1))}
                    if (65535 == 65535) {
                        fmt.Println(v39, v5, v27)
                        fmt.Println(fi33, v1)
                        var v41 int = fi33
                        var v42 int16 = -929
                        v6 = -9223372036854775808
                        _ = v41
                        _ = v42
                    }
                    var v43 int32 = 42
                    _ = v39
                    _ = v40
                    _ = v43
                }
                var v44 int16 = -32768
                _ = fi33
                _ = v44
            }
            _ = fi30
        }
        _ = v27
    }
    switch 538 {
    case 32768:
        if (false || true) {
            v6 -= v6
            fmt.Println(v1, v6)
        } else {
            switch 335 {
            case 32767:
                for fi45 := 0; fi45 < 15; fi45++ {
                    if (!false) {
                        v1 = (v1 | v1)
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v1 -= v1
                            _ = fi46
                        }
                        if (!false) {
                            var v47 uint = 4294967295
                            var v48 uint16 = 32767
                            var v49 int32 = 2147483647
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 20; fi50++ {
                                fmt.Println(fi45)
                                v5 += v5
                                fmt.Println(v6, v6)
                                _ = fi50
                            }
                        }
                    }
                    for fi51 := 0; fi51 < 15; fi51++ {
                        v6 -= (v6 * (-(v6)))
                        _ = fi51
                    }
                    _ = fi45
                }
            case -56:
                v6 -= (v6 >> ((2) & 63))
            case -2147483648:
                delete(v4, 1)
            default:
                if true {
                    var v52 uint8 = 42
                    var v53 int16 = 1000
                    switch "a" {
                    case "":
                        delete(v2, -507)
                    case "hello":
                        switch 1 {
                        case 175:
                            v54 := v4[-814]
                            _ = v54
                        case -2147483648:
                            var v55 int16 = (v53 << ((189) & 63))
                            _ = v55
                        }
                    default:
                        fmt.Println(v52, v5)
                    }
                    _ = v52
                    _ = v53
                } else {
                    _ = len(v4)
                }
            }
            v5 = 42
        }
    case 0:
        v56 := GenS4{F0: 2, F1: 1}
        _ = v56
    case -469:
        if false {
            v6 = (v6 << ((0) & 63))
            v57 := GenS4{F0: 42, F1: 42}
            _ = v57
        } else {
            if ((149 < 1) && (!((v1 | v1) < v1))) {
                if (!true) {
                    var v58 int = 1
                    fmt.Println(v1, v5)
                    _ = v58
                }
            }
            v59 := GenS5{F0: 65535, F1: 21, F2: "a"}
            var v60 string = ""
            _ = v59
            _ = v60
        }
    default:
        v61 := GenS3{F0: v5, F1: -9223372036854775808}
        _ = v61
    }
    fmt.Println(v6, v5, v1)
    var v62 int = 32767
    for fi63 := 0; fi63 < 9; fi63++ {
        v1 = (v1 + (v1 << ((1) & 63)))
        var v64 int16 = -32768
        _ = fi63
        _ = v64
    }
    var v65 int16 = -1
    v66 := []int{v62, (-(v62)), v62, 128}
    switch v62 {
    case 9223372036854775807:
        switch "hello" {
        case "foo bar":
            var v67 uint16 = 42
            _ = v67
        }
    case -1:
        for fi68 := 0; fi68 < 9; fi68++ {
            var v69 uint32 = (v5 * uint32(v1))
            v6 += v6
            _ = fi68
            _ = v69
        }
    case -629:
        if (!false) {
            var v70 uint = 4294967295
            if false {
                var v71 int32 = 42
                v72 := GenS3{F0: 1, F1: 149}
                switch ("hello" + ((("a" + "\n\t") + "") + "a")) {
                case "\n\t":
                    for fi73 := 0; fi73 < 13; fi73++ {
                        switch "hello" {
                        case "":
                            if false {
                                v62 -= v62
                                v71 -= (v71 >> ((830) & 63))
                                delete(v3, "a")
                            } else {
                                if false {
                                    v74 := GenS6{F0: 1, F1: v62, F2: 4294967295}
                                    var v75 int = fi73
                                    _ = v74
                                    _ = v75
                                } else {
                                    v76 := GenS7{F0: 128, F1: v62, F2: 4294967295}
                                    _ = v76
                                }
                                var v77 int16 = (v65 ^ v65)
                                var v78 float64 = float64(float32(float64((0.0 - (1e10 / float32(((1e10 - (-2.25 * ((1e10 + ((0.0 + (-2.25 + (float64(3.14159) + -2.25))) + -2.25)) - 1.5))) / 3.14159)))))))
                                _ = v77
                                _ = v78
                            }
                        }
                        _ = fi73
                    }
                case "a":
                    var v79 uint = (v70 - (v70 | (^(v70))))
                    _ = v79
                case "hello":
                    switch "a" {
                    case "hello":
                        delete(v3, ("foo bar" + ("" + "hello")))
                    case "a":
                        v80 := GenS2{F0: -128, F1: 100, F2: v1}
                        _ = v80
                    }
                default:
                    switch (v62 ^ (v62 - (v62 - v62))) {
                    case 1:
                        v5 -= 42
                    case 42:
                        if (!false) {
                            for fi81 := 0; fi81 < 18; fi81++ {
                                var v82 int64 = v6
                                v83 := GenS2{F0: 127, F1: 42, F2: v1}
                                switch 9223372036854775807 {
                                case 4294967295:
                                    v4[127] = 32767
                                case 194:
                                    var v84 uint8 = 1
                                    _ = v84
                                default:
                                    v85 := GenS1{F0: uint32(v70), F1: (fi81 % ((fi81) | 1)), F2: -893}
                                    _ = v85
                                }
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                            v86 := GenS8{F0: ((("" + "hello") + "\n\t") + "a"), F1: 2}
                            for fi87 := 0; fi87 < 13; fi87++ {
                                v88 := GenS2{F0: 0, F1: 100, F2: 18446744073709551615}
                                _ = fi87
                                _ = v88
                            }
                            _ = v86
                        }
                    case 32768:
                        var v89 uint16 = 420
                        _ = v89
                    default:
                        var v90 uint8 = 2
                        _ = v90
                    }
                }
                _ = v71
                _ = v72
            } else {
                var v91 uint64 = v1
                var v92 int = v62
                _ = v91
                _ = v92
            }
            _ = v70
        } else {
            v65 = v65
            v62 += v62
            v93 := GenS1{F0: (v5 / ((580) | 1)), F1: v62, F2: v62}
            _ = v93
        }
    }
    v1 += 2
    v94 := GenS3{F0: (v5 + 42), F1: (-(v6))}
    if (true && false) {
        var v95 int32 = -91
        v62 = (v62 >> ((2) & 63))
        _ = v95
    } else {
        switch ("hello" + ("hello" + "")) {
        case "\n\t":
            fmt.Println(v1, v6, v1)
        case "":
            var v96 int8 = 127
            _ = v96
        default:
            var v97 bool = true
            _ = v97
        }
        switch (^(v62)) {
        case 4294967295:
            var v98 int8 = -105
            _ = v98
        default:
            v99 := GenS4{F0: 220, F1: 2147483647}
            _ = v99
        }
        var v100 uint16 = 32767
        _ = v100
    }
    var v101 bool = (-32 == -27)
    for fi102 := 0; fi102 < 6; fi102++ {
        switch "\n\t" {
        case "hello":
            switch "" {
            case "foo bar":
                _ = len(v66)
            case "":
                for fi103 := 0; fi103 < 5; fi103++ {
                    delete(v3, "a")
                    _ = fi103
                }
            }
        case "":
            fmt.Println(v1, fi102, v65)
        case "\n\t":
            if v101 {
                var v104 int16 = v65
                var v105 int16 = 32767
                if (-2.25 <= 1e10) {
                    for fi106 := 0; fi106 < 3; fi106++ {
                        if (((v5 >> ((243) & 63)) != (v5 / ((4294967295) | 1))) || v101) {
                            v107 := GenS6{F0: 42, F1: int(v5), F2: v1}
                            var v108 uint32 = (v5 * v5)
                            _ = v107
                            _ = v108
                        }
                        var v109 uint32 = v5
                        _ = fi106
                        _ = v109
                    }
                    v5 += v5
                    v5 = v5
                }
                _ = v104
                _ = v105
            } else {
                v62 = (v62 - (fi102 & v62))
                v6 -= (v6 | v6)
            }
        }
        switch "hello" {
        case "":
            v65 += (v65 ^ v65)
        case "a":
            fmt.Println(v5, v62, v1)
        }
        if (!(32767 != (v65 * v65))) {
            v110 := GenS9{F0: "a"}
            var v111 int32 = -2147483648
            v101 = true
            _ = v110
            _ = v111
        }
        _ = fi102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v65
    _ = v66
    _ = v94
}
