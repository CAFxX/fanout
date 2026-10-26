package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint; F1 int64 }
func genFn0(p0 int16) int32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(p0, fi5)
        _ = fi5
    }
    var v6 int16 = p0
    v7 := []int{127, 32768, 247, 685}
    v8 := GenS1{F0: 1}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() int16 {
    var v9 uint = 46
    var v10 uint8 = 198
    switch -823 {
    case -9223372036854775808:
        v11 := GenS1{F0: 2}
        _ = v11
    case 42:
        v12 := []int{}
        _ = v12
    case 32767:
        var v13 string = ("a" + "foo bar")
        _ = v13
    default:
        var v14 uint = (v9 >> (((v9 << ((4) & 63))) & 63))
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int = fi15
            for fi17 := 0; fi17 < 17; fi17++ {
                switch ("" + ("foo bar" + ((("" + (("foo bar" + "foo bar") + (("\n\t" + "\n\t") + "foo bar"))) + (("a" + "foo bar") + "")) + ("\n\t" + "\n\t")))) {
                case "foo bar":
                    switch (("a" + ("\n\t" + (((("hello" + "foo bar") + "\n\t") + "a") + "a"))) + "\n\t") {
                    case "hello":
                        switch ("foo bar" + ("" + "foo bar")) {
                        case "foo bar":
                            if (!(true && (1 >= 499))) {
                                if (!false) {
                                    var v18 int64 = 1
                                    var v19 float64 = 0.1
                                    var v20 string = ""
                                    if ((574 != 2147483647) || true) {
                                        v19 = v19
                                    } else {
                                        v21 := GenS1{F0: 65535}
                                        v22 := GenS1{F0: 255}
                                        v10 = 128
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 int8 = -128
                                    v24 := []int{fi15, 127, v16}
                                    v19 = 0.1
                                    v18 += (v18 << ((v9) & 63))
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                    _ = v23
                                    _ = v24
                                }
                                var v25 int8 = 60
                                if (v9 >= 858) {
                                    v26 := GenS1{F0: 0}
                                    _ = v26
                                } else {
                                    v27 := GenS2{F0: (v25 & (v25 | (v25 ^ v25))), F1: 32767, F2: false}
                                    switch (("" + "") + "") {
                                    case "a":
                                        v28 := []int{(-(v16)), fi15}
                                        _ = v28
                                    default:
                                        v29 := []int{32768, fi17, (fi17 % (((v16 << (((v9 ^ v9)) & 63))) | 1))}
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        if true {
                                            switch v16 {
                                            case 9223372036854775807:
                                                v16 += fi17
                                            case -331:
                                                if len(v29) > 0 {
                                                    v29[0] = int(v9)
                                                }
                                            }
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            var v30 int = fi17
                                            fmt.Println(fi17, v25)
                                            if true {
                                                var v31 int64 = -1
                                                if (true || (4294967295 == v9)) {
                                                    if len(v29) > 0 {
                                                        v29[0] = v16
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v32 uint32 = 2147483648
                                                        _ = v32
                                                    }
                                                    var v33 bool = (!(65535 > 32767))
                                                    _ = v33
                                                }
                                                v31 += int64(v16)
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v16 = v16
                                                case "a":
                                                    _ = len(v29)
                                                }
                                                fmt.Println(v30, v9, v10)
                                                v16 += fi17
                                                _ = v31
                                            } else {
                                                fmt.Println(v30, v10, v30)
                                                var v34 uint16 = 65535
                                                _ = v34
                                            }
                                            _ = v30
                                        }
                                        _ = v29
                                    }
                                    v9 = v9
                                    switch v16 {
                                    case -943:
                                        if false {
                                            v35 := GenS2{F0: v25, F1: 35, F2: (114 != 9223372036854775807)}
                                            if true {
                                                v16 = (fi17 >> ((v9) & 63))
                                                fmt.Println(v16, v10)
                                                v36 := []int{v16, (fi17 + fi17), (v16 + (v16 >> ((4294967295) & 63)))}
                                                if (!false) {
                                                    if len(v36) > 0 {
                                                        v36[0] = v16
                                                    }
                                                    v9 += v9
                                                } else {
                                                    v37 := GenS3{F0: 0}
                                                    var v38 bool = false
                                                    _ = v37
                                                    _ = v38
                                                }
                                                var v39 uint = (v9 & v9)
                                                _ = v36
                                                _ = v39
                                            }
                                            v40 := []int{fi17, (fi15 | fi15), (-(fi15)), (v16 & 42)}
                                            _ = v35
                                            _ = v40
                                        } else {
                                            if true {
                                                if (false && (1e10 == 0.0)) {
                                                    v41 := []int{-423}
                                                    if false {
                                                        if false {
                                                            switch "foo bar" {
                                                            case "":
                                                                v42 := GenS2{F0: -127, F1: 744, F2: (615 >= 848)}
                                                                _ = v42
                                                            case "a":
                                                                switch ("foo bar" + "a") {
                                                                case "":
                                                                    v43 := GenS1{F0: uint16(v9)}
                                                                    _ = v43
                                                                case "hello":
                                                                    fmt.Println(fi15, v25)
                                                                }
                                                            case "foo bar":
                                                                switch (fi17 & 42) {
                                                                case 685:
                                                                    v16 = fi17
                                                                case 9223372036854775807:
                                                                    fmt.Println(fi17, fi17)
                                                                case 0:
                                                                    fmt.Println(fi17, v16, v25)
                                                                default:
                                                                    _ = len(v41)
                                                                }
                                                            default:
                                                                fmt.Println(v10, v25)
                                                            }
                                                            if len(v41) >= 2 {
                                                                v41 = v41[1:len(v41)]
                                                            }
                                                        }
                                                        if (1 >= 2147483647) {
                                                            if (!true) {
                                                                var v44 uint16 = 32767
                                                                v44 = 0
                                                                _ = v44
                                                            }
                                                        } else {
                                                            var v45 int = fi17
                                                            var v46 float64 = 0.1
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                    var v47 bool = (true && true)
                                                    _ = v41
                                                    _ = v47
                                                }
                                                fmt.Println(fi17, fi15)
                                            }
                                        }
                                    case 582:
                                        if (42 != 32767) {
                                            var v48 uint8 = (v10 >> ((65535) & 63))
                                            v49 := GenS1{F0: 2}
                                            _ = v48
                                            _ = v49
                                        }
                                    case 127:
                                        v16 = fi15
                                    }
                                    _ = v27
                                }
                                _ = v25
                            }
                        case "":
                            v9 = (^(v9))
                        case "\n\t":
                            fmt.Println(v10, v16)
                        default:
                            v10 -= (v10 % (((v10 * v10)) | 1))
                        }
                    default:
                        switch (("hello" + "") + "a") {
                        case "":
                            var v50 float32 = (float32(3.14159) - (((float32((1e10 / 0.1)) + 1e10) + ((1e10 * 1.5) - float32(((1.5 - (1.5 - ((3.14159 - 0.0) + 1e10))) / 3.14159)))) / 1e10))
                            _ = v50
                        }
                    }
                case "\n\t":
                    v16 = (fi17 | fi15)
                    v16 -= (^(v16))
                    v51 := GenS4{F0: 0}
                    fmt.Println(v16, v10, v10)
                    v9 = 65535
                    var v52 uint = (v9 % (((v9 / (((v9 | (v9 ^ (v9 ^ v9)))) | 1))) | 1))
                    _ = v51
                    _ = v52
                case "":
                    var v53 float64 = 1.5
                    _ = v53
                }
                var v54 uint8 = v10
                _ = fi17
                _ = v54
            }
            _ = fi15
            _ = v16
        }
        for fi55 := 0; fi55 < 16; fi55++ {
            for fi56 := 0; fi56 < 2; fi56++ {
                v10 += 84
                v9 -= v9
                if (false && (!(((!false) && true) || (false && ("\n\t" >= ("a" + "a")))))) {
                    var v57 int32 = -271
                    _ = v57
                } else {
                    v58 := GenS5{F0: 65535, F1: 890}
                    v10 = v10
                    _ = v58
                }
                v9 -= v9
                if (9223372036854775807 != 208) {
                    v59 := []int{32768, fi55}
                    v10 += v10
                    v10 = v10
                    _ = v59
                } else {
                    var v60 int16 = 640
                    var v61 string = ("\n\t" + (("foo bar" + "a") + ""))
                    _ = v60
                    _ = v61
                }
                _ = fi56
            }
            if true {
                if (false && (-902 != 0)) {
                    fmt.Println(v10, fi55)
                }
                var v62 int32 = 1
                v9 = (v9 >> ((v9) & 63))
                _ = v62
            } else {
                v63 := []int{0, fi55}
                if ((!true) && true) {
                    for fi64 := 0; fi64 < 20; fi64++ {
                        v9 = (v9 + v9)
                        fmt.Println(fi55)
                        var v65 uint32 = 2
                        _ = fi64
                        _ = v65
                    }
                } else {
                    var v66 int = (fi55 & (^(fi55)))
                    fmt.Println(v66, fi55, v10)
                    _ = v66
                }
                if (4294967295 < 4294967295) {
                    for fi67 := 0; fi67 < 17; fi67++ {
                        v10 += v10
                        if len(v63) > 0 {
                            v63[0] = -757
                        }
                        _ = fi67
                    }
                    v68 := GenS6{F0: "a", F1: -2147483648}
                    _ = v68
                }
                _ = v63
            }
            _ = fi55
        }
    } else {
        if ((32767 > 150) || false) {
            v69 := []int{18}
            _ = v69
        }
        if true {
            for fi70 := 0; fi70 < 2; fi70++ {
                var v71 int = fi70
                var v72 uint16 = 328
                v73 := GenS3{F0: 20}
                _ = fi70
                _ = v71
                _ = v72
                _ = v73
            }
        } else {
            var v74 float64 = 0.1
            if true {
                switch ("" + "a") {
                case "":
                    if true {
                        var v75 float32 = 1e10
                        _ = v75
                    } else {
                        v76 := GenS6{F0: "\n\t", F1: 842}
                        _ = v76
                    }
                case "foo bar":
                    fmt.Println(v10, v10, v10)
                case "a":
                    switch 42 {
                    case 42:
                        v74 -= 1e10
                    case -9223372036854775808:
                        if false {
                            var v77 int = 42
                            v78 := []int{}
                            _ = v77
                            _ = v78
                        }
                    default:
                        v79 := []int{-164}
                        _ = v79
                    }
                }
                var v80 bool = (false && false)
                _ = v80
            } else {
                var v81 int = 32767
                var v82 int32 = -516
                _ = v81
                _ = v82
            }
            var v83 uint8 = (v10 | (^(v10)))
            _ = v74
            _ = v83
        }
    }
    _ = v10
    _ = v9
    return 602
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 98
    v3 := make(map[uint32]bool)
    var v4 float32 = (float32((float64((3.14159 + float32(0.0))) * float64((3.14159 / 1.5)))) / 1.5)
    v2 = (^(v2))
    v1 = genFn1()
    v84 := []int{int(v1), 4294967295, 127}
    var v85 uint16 = uint16(v2)
    if (!false) {
        if (float64((v4 - float32((1.5 * 3.14159)))) <= ((3.14159 - 1.5) - (float64((float32(1e10) + float32((0.0 - 0.0)))) + 1.5))) {
            for fi86 := 0; fi86 < 9; fi86++ {
                v4 -= (float32((1.5 + 0.1)) - v4)
                v2 -= v2
                _ = fi86
            }
        } else {
            v87 := GenS6{F0: "hello", F1: genFn0(-781)}
            switch 318 {
            case -283:
                var v88 float32 = float32(3.14159)
                _ = v88
            case 4294967295:
                if (127 < 255) {
                    v4 = v4
                    if false {
                        _ = len(v84)
                    }
                    v2 -= (v2 * (v2 - (v2 % ((v2) | 1))))
                }
            case 662:
                v84 = append(v84, 128)
            }
            _ = v87
        }
        if len(v84) >= 2 {
            v84 = v84[1:len(v84)]
        }
        var v89 int64 = -925
        _ = v89
    } else {
        switch "foo bar" {
        case "foo bar":
            v85 = (^(v85))
        case "\n\t":
            fmt.Println(v85)
        }
        v90 := GenS2{F0: 100, F1: 1, F2: false}
        _ = v90
    }
    v91 := GenS6{F0: ("foo bar" + "foo bar"), F1: genFn0(v1)}
    if (!(-2.25 >= (float32(1e10) - (1.5 - v4)))) {
        _ = len(v3)
    } else {
        v92 := GenS5{F0: 1, F1: v1}
        switch "hello" {
        case "foo bar":
            for fi93 := 0; fi93 < 10; fi93++ {
                v4 -= v4
                _ = fi93
            }
        case "\n\t":
            v1 = -32768
        }
        if true {
            fmt.Println(v4)
        } else {
            v4 = v4
            v1 -= (v1 << ((259) & 63))
        }
        _ = v92
    }
    for fi94 := 0; fi94 < 6; fi94++ {
        switch "foo bar" {
        case "foo bar":
            v95 := GenS6{F0: "foo bar", F1: 2147483647}
            _ = v95
        case "":
            v96 := GenS2{F0: 127, F1: v1, F2: (true || (!(569 == 42)))}
            _ = v96
        case "\n\t":
            v97 := GenS7{F0: v85, F1: v1}
            _ = v97
        }
        for fi98 := 0; fi98 < 7; fi98++ {
            v85 = 286
            _ = fi98
        }
        v99 := GenS3{F0: -1}
        _ = fi94
        _ = v99
    }
    if len(v84) > 0 {
        v84[0] = 961
    }
    fmt.Println(v2, v4)
    if ((661 <= 190) || false) {
        if true {
            fmt.Println(v1)
        } else {
            var v100 int32 = -1
            _ = v100
        }
    }
    v101 := GenS5{F0: 793, F1: v1}
    v4 += ((v4 + float32(0.0)) - v4)
    fmt.Println(v4, v1)
    for fi102 := 0; fi102 < 1; fi102++ {
        fmt.Println(v85)
        switch (fi102 | (fi102 << ((255) & 63))) {
        case 221:
            var v103 int64 = 1
            _ = v103
        case 2147483647:
            if false {
                v104 := GenS4{F0: (v2 / ((v2) | 1))}
                for fi105 := 0; fi105 < 18; fi105++ {
                    v2 = v2
                    if (fi105 < (fi102 & (^(fi102)))) {
                        switch (fi105 + int(v2)) {
                        case 32768:
                            v106 := GenS1{F0: (v85 + 152)}
                            _ = v106
                        case 580:
                            var v107 uint16 = v85
                            _ = v107
                        case 952:
                            v4 = 1.5
                        }
                    }
                    if false {
                        v4 += (v4 * float32(((0.1 / 1.5) + -2.25)))
                        fmt.Println(fi105)
                        v3[2147483647] = (1 > -238)
                    }
                    var v108 uint = 255
                    _ = fi105
                    _ = v108
                }
                _ = v104
            }
        case -522:
            v4 = 0.0
        default:
            switch "\n\t" {
            case "a":
                v109 := GenS8{F0: 42, F1: -745}
                _ = v109
            }
        }
        for fi110 := 0; fi110 < 18; fi110++ {
            fmt.Println(fi102)
            switch -440 {
            case 127:
                var v111 uint8 = 42
                _ = v111
            case 32767:
                var v112 uint = 255
                _ = v112
            case -295:
                v113 := GenS3{F0: 127}
                _ = v113
            }
            v4 = v4
            v85 = v85
            _ = fi110
        }
        _ = fi102
    }
    v85 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v84
    _ = v85
    _ = v91
}
