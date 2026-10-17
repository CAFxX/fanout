package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint; F2 string }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int8; F1 uint; F2 int32 }
type GenS10 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0(p0 uint16, p1 uint8, p2 uint) int64 {
    var v4 uint8 = p1
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint32 = 2147483647
        var v7 uint64 = 4294967295
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -322, F1: 393, F2: 483}
    v9 := []int{}
    for fi10 := 0; fi10 < 8; fi10++ {
        if ((513 >= 0) || (2147483647 >= 0)) {
            p0 += (p0 | (p0 ^ p0))
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS1{F0: -279, F1: 42, F2: 0}
                _ = fi11
                _ = v12
            }
        }
        switch "hello" {
        case "":
            if (!false) {
                var v13 uint64 = 551
                var v14 uint8 = p1
                var v15 uint16 = (p0 / (((p0 - 0)) | 1))
                _ = v13
                _ = v14
                _ = v15
            }
        case "a":
            var v16 uint = p2
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int64 = 1
            if true {
                var v19 uint = 1
                v9 = append(v9, fi17)
                v4 = v4
                _ = v19
            }
            var v20 uint32 = 1
            var v21 uint64 = 4294967295
            var v22 string = ""
            var v23 float32 = (0.1 + float32(0.1))
            _ = fi17
            _ = v18
            _ = v20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 int64, p2 int8) uint {
    switch ("hello" + ("a" + "")) {
    case "hello":
        fmt.Println(p0, p1)
    case "\n\t":
        v24 := GenS1{F0: -1, F1: 2147483647, F2: 1000}
        _ = v24
    }
    if (p1 != p1) {
        p0 -= 1
        v25 := GenS2{F0: (^(p2))}
        _ = v25
    }
    v26 := []int{128, 307, 2147483647, 4294967295}
    switch 32768 {
    case -2147483648:
        p1 = (p1 >> ((1) & 63))
    default:
        v27 := GenS3{F0: 365}
        _ = v27
    }
    p2 = int8(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 906
}

func main() {
    v1 := []uint32{614, 4294967295, 4294967295}
    var v2 int64 = -1
    var v3 uint16 = 32767
    var v28 int64 = 42
    v29 := GenS2{F0: 0}
    var v30 bool = ((-481 > 42) && false)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -526 {
    case -187:
        if true {
            if ((-2147483648 < int32(v2)) && (-2147483648 <= -548)) {
                if len(v1) > 0 {
                    v1[0] = 521
                }
                v28 = genFn0((v3 << ((626) & 63)), 1, uint(v3))
            }
            var v31 string = ("" + "a")
            switch v31 {
            case "\n\t":
                v28 = v28
            case "hello":
                fmt.Println(v2, v28, v31)
            case "a":
                for fi32 := 0; fi32 < 13; fi32++ {
                    if true {
                        v2 += 831
                        v33 := GenS4{F0: 1, F1: 535, F2: 843}
                        if len(v1) > 0 {
                            v1[0] = 671
                        }
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 8; fi34++ {
                        if (!(!true)) {
                            fmt.Println(v28, fi34)
                            fmt.Println(v2, fi34, v30)
                            var v35 string = "\n\t"
                            _ = len(v1)
                            _ = v35
                        } else {
                            switch (fi34 ^ fi32) {
                            case 32768:
                                if false {
                                    if (116 != 65535) {
                                        var v36 int = 9223372036854775807
                                        var v37 int64 = 1000
                                        _ = v36
                                        _ = v37
                                    }
                                    switch v31 {
                                    case "":
                                        var v38 string = v31
                                        _ = v38
                                    case "foo bar":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = 42
                                        }
                                    default:
                                        if v30 {
                                            if (!v30) {
                                                if ((255 == 2) || v30) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    fmt.Println(v31)
                                                    fmt.Println(fi32)
                                                }
                                            } else {
                                                if false {
                                                    fmt.Println(fi34)
                                                    v31 = v31
                                                    var v39 int16 = 633
                                                    _ = v39
                                                } else {
                                                    switch v31 {
                                                    case "\n\t":
                                                        v3 = (^(v3))
                                                    case "":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case "hello":
                                                        if v30 {
                                                            v31 = ((v31 + ((v31 + v31) + "\n\t")) + v31)
                                                            v31 = v31
                                                        }
                                                    }
                                                }
                                            }
                                            if v30 {
                                                v2 += (v2 - 1)
                                                if (-2147483648 < 1) {
                                                    v31 = v31
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v40 := GenS5{F0: v31, F1: 4294967295, F2: 812}
                                                    _ = v40
                                                }
                                                v28 = v28
                                            }
                                            fmt.Println(v28, v3, fi34)
                                        }
                                    }
                                    var v41 int8 = -128
                                    _ = v41
                                }
                            case 154:
                                fmt.Println(v3)
                            default:
                                var v42 float32 = 3.14159
                                _ = v42
                            }
                            var v43 uint16 = v3
                            _ = v43
                        }
                        fmt.Println(fi34, v28)
                        if (!false) {
                            var v44 float32 = (-2.25 - ((((float32(3.14159) * 1.5) + 0.1) - (0.1 * 3.14159)) + float32(-2.25)))
                            switch (fi34 >> ((255) & 63)) {
                            case 42:
                                if v30 {
                                    v44 += float32(0.1)
                                } else {
                                    fmt.Println(v44, fi34, v28)
                                    if (-2147483648 < 312) {
                                        var v45 float32 = v44
                                        var v46 string = ((((v31 + v31) + v31) + "\n\t") + (v31 + ("foo bar" + (v31 + v31))))
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v47 uint = genFn1(128, -428, 0)
                                        if false {
                                            fmt.Println(v2, v44, v3)
                                        }
                                        _ = v47
                                    }
                                    if v30 {
                                        v48 := GenS6{F0: 140, F1: 749, F2: "foo bar"}
                                        var v49 uint = 48
                                        _ = len(v1)
                                        _ = v48
                                        _ = v49
                                    } else {
                                        _ = len(v1)
                                        var v50 int32 = 672
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v51 float64 = (0.0 - 0.0)
                                            _ = v51
                                        case "hello":
                                            v52 := GenS7{F0: 42, F1: fi34}
                                            _ = v52
                                        case "":
                                            switch (fi32 << ((4294967295) & 63)) {
                                            case -906:
                                                v2 = (v2 % ((v28) | 1))
                                            case 925:
                                                var v53 uint32 = 2147483647
                                                _ = v53
                                            }
                                        }
                                        _ = v50
                                    }
                                    v31 = "hello"
                                }
                            }
                            _ = v44
                        }
                        _ = fi34
                    }
                    switch 1 {
                    case -843:
                        v3 += (v3 / (((^(v3))) | 1))
                    case 140:
                        v3 = v3
                    case 32768:
                        for fi54 := 0; fi54 < 16; fi54++ {
                            if (0 <= 0) {
                                v2 += 189
                                v30 = (127 < 1)
                            }
                            switch (fi54 / ((fi32) | 1)) {
                            case 42:
                                var v55 int8 = 1
                                v56 := GenS3{F0: -267}
                                _ = v55
                                _ = v56
                            }
                            var v57 string = (v31 + "a")
                            _ = fi54
                            _ = v57
                        }
                    default:
                        for fi58 := 0; fi58 < 6; fi58++ {
                            if v30 {
                                if (!(!(v2 >= v2))) {
                                    var v59 uint32 = 2147483648
                                    v1 = append(v1, v59)
                                    _ = v59
                                } else {
                                    var v60 uint = 45
                                    _ = v60
                                }
                                var v61 float64 = 1.5
                                var v62 uint32 = 725
                                var v63 int32 = 1
                                v1 = append(v1, (v62 >> ((954) & 63)))
                                if (255 != 857) {
                                    fmt.Println(fi58, v3)
                                }
                                v64 := GenS6{F0: (v62 >> ((624) & 63)), F1: genFn1(uint8(v3), -1, -128), F2: v31}
                                _ = v61
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                            if v30 {
                                v65 := GenS8{F0: 1, F1: fi58}
                                _ = v65
                            }
                            _ = fi58
                        }
                    }
                    _ = fi32
                }
            }
            _ = v31
        }
    default:
        var v66 uint = 2
        _ = v66
    }
    v3 = (v3 % ((v3) | 1))
    var v67 uint64 = 619
    if false {
        if v30 {
            v67 = v67
            v30 = v30
            if true {
                v28 = v28
                for fi68 := 0; fi68 < 2; fi68++ {
                    for fi69 := 0; fi69 < 13; fi69++ {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        v3 = 65535
                        var v70 uint16 = v3
                        v2 = (-(v2))
                        var v71 int8 = 127
                        if (!true) {
                            v3 = v70
                            v30 = false
                            var v72 int = 2147483647
                            _ = v72
                        }
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    _ = fi68
                }
            }
        }
        _ = len(v1)
    } else {
        for fi73 := 0; fi73 < 5; fi73++ {
            for fi74 := 0; fi74 < 5; fi74++ {
                fmt.Println(fi73, v67)
                _ = fi74
            }
            if false {
                for fi75 := 0; fi75 < 20; fi75++ {
                    if true {
                        var v76 bool = ((!v30) && (1e10 >= ((0.0 * ((-2.25 + 0.0) + 3.14159)) - (1.5 - -2.25))))
                        _ = v76
                    } else {
                        var v77 int64 = (v28 % ((genFn0((v3 & 255), 127, 573)) | 1))
                        v30 = false
                        _ = v77
                    }
                    _ = fi75
                }
                if v30 {
                    if v30 {
                        v67 = v67
                        for fi78 := 0; fi78 < 17; fi78++ {
                            v79 := GenS4{F0: -1, F1: 42, F2: 849}
                            fmt.Println(v30)
                            var v80 int = fi73
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    }
                    if (genFn0(v3, 2, 42) > 9223372036854775807) {
                        switch "\n\t" {
                        case "foo bar":
                            if (!(198 > 2)) {
                                v28 -= -1
                            } else {
                                var v81 bool = false
                                _ = v81
                            }
                        case "a":
                            var v82 int = (^(fi73))
                            _ = v82
                        case "":
                            v83 := GenS5{F0: "hello", F1: 42, F2: 4294967295}
                            _ = v83
                        default:
                            if (((v2 * v2) == (v28 * (v2 >> ((255) & 63)))) && (true && v30)) {
                                v28 = (v28 | (v2 - v28))
                                var v84 string = ""
                                _ = v84
                            } else {
                                if (-2.25 == (float32(0.0) / float32(float64((float32(1.5) / (float32(float64(1.5)) / 1e10)))))) {
                                    var v85 int16 = 42
                                    if v30 {
                                        switch (fi73 / (((fi73 ^ fi73)) | 1)) {
                                        case 1:
                                            if false {
                                                var v86 int16 = v85
                                                var v87 uint32 = 2
                                                _ = v86
                                                _ = v87
                                            }
                                        case -9223372036854775808:
                                            v88 := GenS9{F0: 127, F1: 255, F2: -80}
                                            _ = v88
                                        case 4294967295:
                                            for fi89 := 0; fi89 < 4; fi89++ {
                                                v85 -= v85
                                                v90 := GenS1{F0: (v85 - (v85 * v85)), F1: 0, F2: (^(v85))}
                                                var v91 bool = v30
                                                _ = fi89
                                                _ = v90
                                                _ = v91
                                            }
                                        default:
                                            var v92 int32 = 2147483647
                                            _ = v92
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = v85
                                } else {
                                    var v93 uint32 = 2
                                    _ = v93
                                }
                                var v94 uint32 = 4294967295
                                v2 += v2
                                _ = v94
                            }
                        }
                        v95 := GenS10{F0: int8(fi73), F1: 135, F2: (v67 << ((18446744073709551615) & 63))}
                        _ = v95
                    } else {
                        v96 := GenS6{F0: 169, F1: 1, F2: ("a" + "foo bar")}
                        _ = v96
                    }
                    switch ("a" + "foo bar") {
                    case "a":
                        var v97 int8 = 0
                        _ = v97
                    default:
                        v1 = append(v1, 716)
                    }
                } else {
                    v2 += v28
                }
            }
            _ = fi73
        }
        if len(v1) > 0 {
            v1[0] = uint32(v3)
        }
    }
    if ((v30 || (3.14159 >= (-2.25 - -2.25))) && (v2 < genFn0(727, uint8(v67), 65535))) {
        var v98 int16 = -1
        _ = v98
    } else {
        v99 := GenS8{F0: (v67 * 2), F1: -2147483648}
        v30 = (!(genFn0((v3 << ((2) & 63)), 1, 4294967295) <= (v28 / ((0) | 1))))
        _ = v99
    }
    switch "\n\t" {
    case "\n\t":
        var v100 uint16 = 255
        _ = v100
    case "":
        for fi101 := 0; fi101 < 9; fi101++ {
            fmt.Println(v67, v30, v2)
            var v102 uint16 = 1
            v1 = append(v1, uint32(v102))
            _ = fi101
            _ = v102
        }
    case "foo bar":
        v28 -= (v2 * genFn0(v3, 2, 18446744073709551615))
    default:
        for fi103 := 0; fi103 < 19; fi103++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v67 = (v67 << ((42) & 63))
            for fi104 := 0; fi104 < 1; fi104++ {
                if v30 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v105 uint64 = v67
                    if (210 < 255) {
                        fmt.Println(v30)
                        var v106 uint32 = 238
                        _ = v106
                    }
                    _ = v105
                }
                v107 := GenS6{F0: 2147483647, F1: 663, F2: ("a" + "")}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi104
                _ = v107
            }
            _ = fi103
        }
    }
    v28 -= 42
    switch "" {
    case "":
        fmt.Println(v30)
    case "foo bar":
        var v108 int = -88
        _ = v108
    case "hello":
        if v30 {
            var v109 int16 = 1
            var v110 int64 = (-(v28))
            _ = v109
            _ = v110
        } else {
            fmt.Println(v28, v3, v3)
        }
    default:
        v2 -= genFn0(1, 127, 1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v67
}
