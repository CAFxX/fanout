package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
type GenS4 struct { F0 int16; F1 int16; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 bool }
func genFn0(p0 bool, p1 int8) int32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float64 = 1e10
        var v6 uint16 = 694
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p1, p1)
    p0 = p0
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 1, F1: fi7, F2: 65535}
        switch ("hello" + "") {
        case "hello":
            var v9 float32 = -2.25
            _ = v9
        }
        p1 = p1
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: 298, F1: 128, F2: 65535}
    _ = p0
    _ = p1
    _ = v10
    return -604
}
func genFn1(p0 uint16, p1 uint8) int {
    p0 -= p0
    var v11 int64 = 0
    if true {
        v11 = v11
        fmt.Println(v11, v11, v11)
        fmt.Println(p0, p1, p0)
    } else {
        p0 = p0
        var v12 uint64 = 0
        _ = v12
    }
    v13 := GenS2{F0: 0}
    v14 := []int{1, 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    _ = v14
    return 830
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 int8 = 42
    v3 -= v3
    v3 -= -128
    v15 := GenS1{F0: -1, F1: genFn1(65535, 255), F2: 277}
    var v16 string = (("hello" + "\n\t") + "hello")
    fmt.Println(v16, v3)
    if false {
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS1{F0: 42, F1: (fi17 * genFn1(0, 0)), F2: 201}
            var v19 string = (v16 + "foo bar")
            var v20 uint64 = 4294967295
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        if false {
            v3 = 103
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 bool = (true || (!(false && false)))
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    } else {
        _ = len(v1)
        if true {
            if (0 != 65535) {
                v24 := v1[4294967295]
                _ = v24
            }
            var v25 float32 = (float32(-2.25) + 3.14159)
            _ = v25
        }
        var v26 int = 769
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 bool = false
        var v29 string = v16
        delete(v1, 1)
        _ = fi27
        _ = v28
        _ = v29
    }
    if true {
        v30 := GenS3{F0: true, F1: 65535, F2: 65535}
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := GenS1{F0: -9223372036854775808, F1: 127, F2: 42}
            v33 := GenS4{F0: 1, F1: 32767, F2: (910 < 255)}
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 18; fi34++ {
            if false {
                var v35 float32 = float32(0.0)
                var v36 int16 = 0
                _ = v35
                _ = v36
            }
            switch fi34 {
            case 42:
                if (("a" + "") > "") {
                    if true {
                        v37 := GenS4{F0: 1, F1: 1000, F2: (true || (!true))}
                        for fi38 := 0; fi38 < 16; fi38++ {
                            if len(v2) > 0 {
                                v2[0] = v16
                            }
                            _ = fi38
                        }
                        if true {
                            v3 = v3
                            v16 = (v16 + ((v16 + (v16 + ("hello" + (("\n\t" + v16) + v16)))) + v16))
                            v3 -= (v3 & v3)
                        }
                        _ = v37
                    } else {
                        _ = len(v1)
                        v16 = ""
                        switch fi34 {
                        case -309:
                            v3 = v3
                        case -2147483648:
                            v3 += (v3 / ((v3) | 1))
                        default:
                            if len(v2) > 0 {
                                v2[0] = v16
                            }
                        }
                    }
                }
            case 9223372036854775807:
                for fi39 := 0; fi39 < 14; fi39++ {
                    var v40 uint64 = 42
                    _ = fi39
                    _ = v40
                }
            default:
                if (1 > 98) {
                    for fi41 := 0; fi41 < 19; fi41++ {
                        v3 += v3
                        _ = fi41
                    }
                    v3 -= 127
                    for fi42 := 0; fi42 < 1; fi42++ {
                        if true {
                            _ = len(v2)
                            v43 := GenS1{F0: 369, F1: (fi34 << ((491) & 63)), F2: 65535}
                            var v44 int32 = 0
                            _ = v43
                            _ = v44
                        }
                        fmt.Println(v3, fi34, v3)
                        _ = fi42
                    }
                } else {
                    v45 := GenS5{F0: (v3 & (v3 % ((42) | 1)))}
                    switch (v16 + ((v16 + (v16 + "\n\t")) + v16)) {
                    case "a":
                        if false {
                            if (0 >= 1) {
                                for fi46 := 0; fi46 < 15; fi46++ {
                                    v47 := GenS4{F0: 264, F1: 32767, F2: false}
                                    _ = fi46
                                    _ = v47
                                }
                                var v48 float64 = 0.1
                                for fi49 := 0; fi49 < 7; fi49++ {
                                    v50 := v1[2147483648]
                                    _ = fi49
                                    _ = v50
                                }
                                _ = v48
                            }
                            if true {
                                v51 := GenS6{F0: 42}
                                if ((v3 + (-(v3))) > (v3 % (((^(v3))) | 1))) {
                                    var v52 int8 = (v3 + (v3 % ((127) | 1)))
                                    var v53 bool = (-1 == -1)
                                    _ = v52
                                    _ = v53
                                }
                                fmt.Println(v16)
                                _ = v51
                            }
                            fmt.Println(v16)
                        } else {
                            if (!(!true)) {
                                if (230 == 808) {
                                    var v54 int = 127
                                    _ = len(v2)
                                    _ = v54
                                }
                                _ = len(v1)
                                v16 = (v16 + "")
                            } else {
                                v1[0] = 2
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    v56 := GenS2{F0: -2147483648}
                                    v57 := GenS5{F0: (v3 % ((126) | 1))}
                                    fmt.Println(v3, v16)
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                            }
                            for fi58 := 0; fi58 < 2; fi58++ {
                                var v59 uint64 = 2
                                v16 = "foo bar"
                                if (2 < 128) {
                                    var v60 int32 = -519
                                    _ = v60
                                } else {
                                    v61 := GenS2{F0: -1}
                                    _ = len(v2)
                                    _ = len(v1)
                                    _ = v61
                                }
                                _ = fi58
                                _ = v59
                            }
                        }
                    }
                    _ = v45
                }
            }
            for fi62 := 0; fi62 < 8; fi62++ {
                v1[2] = 1
                _ = fi62
            }
            _ = fi34
        }
        _ = v30
    } else {
        fmt.Println(v16, v3)
        v16 = ((v16 + (("\n\t" + ("\n\t" + v16)) + v16)) + v16)
    }
    for fi63 := 0; fi63 < 8; fi63++ {
        for fi64 := 0; fi64 < 5; fi64++ {
            var v65 int32 = -399
            v65 = genFn0(true, 100)
            fmt.Println(fi63, fi64, v65)
            _ = fi64
            _ = v65
        }
        v66 := GenS7{F0: 127}
        v67 := GenS6{F0: 65535}
        _ = fi63
        _ = v66
        _ = v67
    }
    fmt.Println(v3)
    var v68 float64 = (1.5 + 1.5)
    switch 42 {
    case -861:
        if true {
            if false {
                var v69 uint8 = 128
                var v70 string = v16
                v70 = v70
                _ = v69
                _ = v70
            }
        }
    case 0:
        v3 = 0
    default:
        var v71 int = int(v3)
        _ = v71
    }
    var v72 float32 = -2.25
    for fi73 := 0; fi73 < 18; fi73++ {
        if len(v2) > 0 {
            v2[0] = (v16 + ("hello" + v16))
        }
        for fi74 := 0; fi74 < 7; fi74++ {
            switch "\n\t" {
            case "":
                if false {
                    fmt.Println(v68, v68, v3)
                } else {
                    switch (("foo bar" + ((v16 + v16) + "foo bar")) + "a") {
                    case "a":
                        var v75 bool = true
                        _ = v75
                    case "foo bar":
                        v3 -= (v3 << ((1) & 63))
                    }
                    v16 = ("\n\t" + v16)
                    var v76 uint8 = 180
                    _ = v76
                }
            }
            fmt.Println(v16)
            fmt.Println(v16, v3, v3)
            var v77 uint8 = 109
            _ = fi74
            _ = v77
        }
        for fi78 := 0; fi78 < 6; fi78++ {
            switch fi78 {
            case 9223372036854775807:
                if (128 <= 2) {
                    switch 42 {
                    case 666:
                        v68 -= 0.0
                    }
                }
            case 4294967295:
                if (false && false) {
                    v79 := GenS8{F0: 924, F1: ("foo bar" + "")}
                    switch (fi78 / ((genFn1(697, 127)) | 1)) {
                    case -2147483648:
                        if false {
                            v16 = v16
                            if (!((-2.25 / (0.1 * (v72 + (1.5 / v72)))) <= v72)) {
                                v3 += 42
                            }
                        } else {
                            v3 = v3
                            switch (fi73 / (((fi78 / ((fi73) | 1))) | 1)) {
                            case 1:
                                if (!false) {
                                    v1[2] = 0
                                    var v80 int = (fi78 ^ genFn1(2, 255))
                                    _ = v80
                                } else {
                                    switch ("a" + v16) {
                                    case "":
                                        if (((int64(fi78) <= 767) && ((-478 == -1) && true)) || (v3 <= v3)) {
                                            var v81 string = (v16 + v16)
                                            if true {
                                                if (0.1 != v68) {
                                                    v82 := GenS1{F0: -886, F1: genFn1(32767, 248), F2: 2147483648}
                                                    v3 += (v3 + v3)
                                                    _ = v82
                                                }
                                                v83 := GenS2{F0: 0}
                                                var v84 float64 = v68
                                                _ = v83
                                                _ = v84
                                            } else {
                                                fmt.Println(fi73, fi78)
                                                fmt.Println(fi73, v81, v3)
                                            }
                                            _ = v81
                                        }
                                    case "foo bar":
                                        v16 = v16
                                    }
                                    fmt.Println(fi78)
                                }
                                var v85 int64 = -9223372036854775808
                                _ = v85
                            default:
                                switch "a" {
                                case "":
                                    var v86 uint16 = 65535
                                    _ = v86
                                }
                            }
                        }
                    }
                    v2 = append(v2, ((v16 + v16) + v16))
                    _ = v79
                } else {
                    fmt.Println(v68, v3)
                }
            case 42:
                v1[4294967295] = 329
            default:
                var v87 float64 = v68
                _ = v87
            }
            if (1 == 332) {
                if false {
                    v2 = append(v2, (v16 + v16))
                    if ((fi73 - fi73) != fi78) {
                        var v88 uint64 = 0
                        var v89 uint8 = 128
                        v90 := v1[uint32(v89)]
                        _ = v88
                        _ = v89
                        _ = v90
                    } else {
                        v91 := GenS6{F0: 479}
                        if true {
                            v3 -= (^(v3))
                        }
                        delete(v1, 2147483647)
                        _ = v91
                    }
                    if (390 < 32767) {
                        _ = len(v2)
                        v72 += v72
                        switch "\n\t" {
                        case "\n\t":
                            var v92 uint16 = 572
                            _ = v92
                        }
                    }
                    var v93 uint = 42
                    if (2147483647 != -2147483648) {
                        var v94 uint = v93
                        v3 = v3
                        v95 := GenS3{F0: false, F1: (^(v94)), F2: uint32(v93)}
                        _ = v94
                        _ = v95
                    } else {
                        var v96 uint8 = uint8(v93)
                        _ = v96
                    }
                    _ = v93
                } else {
                    fmt.Println(v72)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v97 bool = true
                    if v97 {
                        if false {
                            if v97 {
                                delete(v1, 0)
                                v3 = (v3 << ((0) & 63))
                                if len(v2) > 0 {
                                    v2[0] = ("a" + ("hello" + v16))
                                }
                            }
                            v72 = 0.0
                            if ((1 != 1) || true) {
                                var v98 bool = ((fi73 << ((65535) & 63)) <= (fi78 & (fi73 >> ((255) & 63))))
                                v99 := v1[908]
                                var v100 int = (fi73 + fi78)
                                if v97 {
                                    if (308 <= 568) {
                                        var v101 uint16 = 954
                                        _ = v101
                                    } else {
                                        v68 = 0.1
                                    }
                                    var v102 uint8 = 42
                                    _ = v102
                                } else {
                                    if v98 {
                                        var v103 int8 = (v3 << ((1) & 63))
                                        switch "" {
                                        case "":
                                            delete(v1, v99)
                                        default:
                                            switch fi73 {
                                            case 9223372036854775807:
                                                fmt.Println(fi78, v16, fi73)
                                            case 897:
                                                v104 := GenS9{F0: 0, F1: uint(v99)}
                                                _ = v104
                                            case 687:
                                                fmt.Println(v68)
                                            default:
                                                fmt.Println(fi73)
                                            }
                                        }
                                        var v105 float64 = float64(v72)
                                        _ = v103
                                        _ = v105
                                    } else {
                                        v99 = 65535
                                    }
                                    if v97 {
                                        v100 -= genFn1(42, 255)
                                        var v106 string = (("a" + (v16 + v16)) + v16)
                                        var v107 uint32 = (^(v99))
                                        _ = v106
                                        _ = v107
                                    }
                                }
                                _ = v100
                                _ = v98
                                _ = v99
                            }
                            var v108 uint32 = 0
                            _ = v108
                        } else {
                            switch fi73 {
                            case 1:
                                var v109 int = (^(fi78))
                                _ = v109
                            default:
                                fmt.Println(v68, fi78, v72)
                            }
                            if false {
                                v110 := GenS7{F0: 223}
                                var v111 uint = 0
                                if false {
                                    switch v16 {
                                    case "hello":
                                        v16 = (v16 + v16)
                                        if (v97 || (v97 && (v97 && (v97 && (!(0 <= int16(fi73))))))) {
                                            fmt.Println(v68)
                                        } else {
                                            fmt.Println(v97)
                                            fmt.Println(v97, v97, v3)
                                            switch v16 {
                                            case "foo bar":
                                                v97 = true
                                            case "hello":
                                                var v112 uint64 = 0
                                                _ = v112
                                            default:
                                                _ = len(v2)
                                            }
                                        }
                                        switch v16 {
                                        case "foo bar":
                                            v72 -= 1.5
                                        case "a":
                                            v68 = 1.5
                                        }
                                    default:
                                        var v113 int8 = (v3 << ((v111) & 63))
                                        _ = v113
                                    }
                                    if v97 {
                                        v72 = v72
                                        var v114 bool = v97
                                        v115 := GenS10{F0: (2147483648 == 42)}
                                        v16 = ((v16 + (v16 + "hello")) + v16)
                                        switch fi73 {
                                        case 4294967295:
                                            if (!v114) {
                                                var v116 int32 = 886
                                                fmt.Println(fi78)
                                                _ = v116
                                            }
                                        }
                                        _ = v114
                                        _ = v115
                                    }
                                }
                                _ = v110
                                _ = v111
                            } else {
                                if v97 {
                                    v72 = ((float32(v68) * float32(0.0)) * v72)
                                }
                                v68 = v68
                            }
                        }
                        if (4294967295 >= 2) {
                            var v117 int64 = 1
                            v118 := GenS1{F0: (-(v117)), F1: fi78, F2: 2}
                            _ = v117
                            _ = v118
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if v97 {
                        v68 = 1.5
                        if v97 {
                            _ = len(v2)
                            v16 = ""
                        }
                        fmt.Println(v72, v3, v97)
                    }
                    if (!((fi78 & fi78) < fi78)) {
                        if len(v2) > 0 {
                            v2[0] = v16
                        }
                        v119 := GenS3{F0: v97, F1: 255, F2: 42}
                        _ = v119
                    }
                    v68 = float64((v72 + v72))
                    var v120 int32 = genFn0(v97, (v3 << ((1) & 63)))
                    if false {
                        v97 = (v97 || (v97 && (-557 == 1000)))
                        var v121 uint = 42
                        _ = v121
                    }
                    var v122 int32 = 1
                    _ = v120
                    _ = v122
                    _ = v97
                }
                fmt.Println(fi73, v3)
                v68 -= 3.14159
            } else {
                v123 := GenS7{F0: 255}
                _ = v123
            }
            v68 = v68
            _ = fi78
        }
        _ = fi73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v68:", v68)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v68
    _ = v72
}
