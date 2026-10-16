package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 int16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 1
    v3 := make(map[int]bool)
    var v4 int32 = 2147483647
    var v5 uint8 = 146
    v6 := make(map[string]bool)
    switch 1 {
    case 32767:
        if true {
            fmt.Println(v5, v2, v2)
            v4 = (^(v4))
        }
    case 128:
        if (!true) {
            v4 = v4
            v2 += -1
        }
    case 4294967295:
        var v7 uint32 = 2
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 32767, F2: false}
    switch ("\n\t" + "\n\t") {
    case "a":
        if (1.5 != -2.25) {
            switch 127 {
            case -774:
                if (((true && true) && false) || (!(!(42 == 495)))) {
                    switch 932 {
                    case 127:
                        switch ("a" + "a") {
                        case "\n\t":
                            v9 := []int{32767, 2147483647, -1, 127}
                            _ = v9
                        default:
                            v10 := []int{-641, 4294967295, 128}
                            _ = v10
                        }
                    default:
                        v4 = v4
                    }
                    v11 := GenS2{F0: -1}
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            var v14 float32 = (1.5 / (3.14159 + (1.5 / 1.5)))
                            var v15 uint32 = 1
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        var v16 uint8 = v5
                        v17 := []int{fi12, 0}
                        _ = fi12
                        _ = v16
                        _ = v17
                    }
                    var v18 float64 = ((-2.25 / -2.25) * (1e10 - 1.5))
                    _ = v18
                }
            case -360:
                v4 -= v4
            }
        } else {
            v19 := GenS1{F0: 32767, F1: 42, F2: true}
            if (v2 < -993) {
                switch "\n\t" {
                case "foo bar":
                    switch 32768 {
                    case -771:
                        fmt.Println(v4, v2, v2)
                    case 127:
                        if (("hello" <= "foo bar") && (0.1 < 1.5)) {
                            var v20 int32 = v4
                            _ = v20
                        }
                    case 162:
                        v2 += -1
                    }
                case "\n\t":
                    v5 -= (v5 - v5)
                case "":
                    v5 = (v5 ^ 1)
                default:
                    var v21 int8 = 1
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch fi22 {
                    case -9223372036854775808:
                        for fi23 := 0; fi23 < 7; fi23++ {
                            var v24 uint16 = 42
                            if (true || true) {
                                var v25 uint32 = uint32(v5)
                                fmt.Println(v25, v25)
                                var v26 int16 = int16(v5)
                                _ = v25
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                    case 339:
                        var v27 string = "foo bar"
                        _ = v27
                    case 843:
                        switch (fi22 % ((fi22) | 1)) {
                        case -342:
                            for fi28 := 0; fi28 < 14; fi28++ {
                                if false {
                                    v29 := GenS1{F0: 860, F1: 0, F2: (true && (-1 == int16(v4)))}
                                    if false {
                                        fmt.Println(fi28, fi28, fi22)
                                    } else {
                                        fmt.Println(v5, fi28)
                                        if false {
                                            var v30 bool = true
                                            _ = v30
                                        } else {
                                            v4 += -2147483648
                                            fmt.Println(v5)
                                        }
                                    }
                                    _ = v29
                                }
                                _ = fi28
                            }
                        default:
                            var v31 uint16 = 0
                            _ = v31
                        }
                    }
                    _ = fi22
                }
                v32 := []int{1, 9223372036854775807, 128}
                if false {
                    v2 -= (v2 ^ v2)
                    v33 := GenS3{F0: uint32(v5)}
                    _ = v33
                }
                _ = v32
            }
            _ = v19
        }
    case "":
        v34 := []int{1, -9223372036854775808, 42}
        _ = v34
    default:
        var v35 uint = 141
        _ = v35
    }
    var v36 uint = 1
    if false {
        v37 := GenS4{F0: v36, F1: 32767, F2: 1000}
        _ = v37
    }
    var v38 int = 1
    if (!true) {
        v4 -= v4
        for fi39 := 0; fi39 < 8; fi39++ {
            v40 := []int{int(v5), -1, v38, v38}
            v2 = (v2 * (v2 * (v2 >> (((v36 / ((v36) | 1))) & 63))))
            fmt.Println(v4, v4)
            _ = fi39
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 5; fi41++ {
            v4 -= 42
            switch -2147483648 {
            case 580:
                if (!false) {
                    var v42 uint64 = 887
                    v43 := []int{fi41, 42}
                    _ = v42
                    _ = v43
                } else {
                    v4 = 42
                    for fi44 := 0; fi44 < 4; fi44++ {
                        v45 := []int{2147483647, fi44, (^(v38))}
                        _ = fi44
                        _ = v45
                    }
                }
            case 127:
                v1[2147483648] = 1
            default:
                v5 = v5
            }
            _ = fi41
        }
        if true {
            v4 += v4
            switch ("a" + (("hello" + "\n\t") + "\n\t")) {
            case "\n\t":
                if (false || true) {
                    v36 = (v36 << (((v36 ^ (v36 % (((^(v36))) | 1)))) & 63))
                }
            default:
                v36 += v36
            }
            if (false || (!(true && true))) {
                v38 -= 32768
            } else {
                v2 = -1
                if (1 != 4294967295) {
                    switch "a" {
                    case "a":
                        v46 := v6["hello"]
                        _ = v46
                    default:
                        v6[(("\n\t" + "hello") + "foo bar")] = false
                    }
                    if (-2.25 < (1e10 / float64(float32(-2.25)))) {
                        fmt.Println(v2, v38)
                        var v47 bool = true
                        _ = v47
                    } else {
                        for fi48 := 0; fi48 < 3; fi48++ {
                            for fi49 := 0; fi49 < 8; fi49++ {
                                if (("a" + "a") <= ("\n\t" + "")) {
                                    switch (^(v38)) {
                                    case 599:
                                        if (false && false) {
                                            v4 -= -790
                                            _ = len(v3)
                                            if (0 <= 4294967295) {
                                                v36 = 65535
                                            }
                                        } else {
                                            switch (v38 * v38) {
                                            case 9223372036854775807:
                                                var v50 float64 = 0.0
                                                _ = v50
                                            case -956:
                                                switch "a" {
                                                case "hello":
                                                    v5 = v5
                                                default:
                                                    var v51 string = "a"
                                                    _ = v51
                                                }
                                                v52 := []int{(-(fi49)), fi49, int(v5)}
                                                v53 := GenS1{F0: 65535, F1: 65535, F2: true}
                                                var v54 int32 = (v4 % ((735) | 1))
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            case 1:
                                                v38 = fi49
                                                if false {
                                                    var v55 uint8 = 0
                                                    v55 += (^(v55))
                                                    _ = v55
                                                }
                                            }
                                            v56 := GenS5{F0: 1, F1: fi49}
                                            if false {
                                                v4 += (v4 / ((v4) | 1))
                                                v57 := []int{(v38 ^ v38)}
                                                if (v5 != v5) {
                                                    if (v5 >= (^(v5))) {
                                                        fmt.Println(v5)
                                                        if true {
                                                            switch ("\n\t" + "\n\t") {
                                                            case "hello":
                                                                fmt.Println(v38, v2)
                                                                v58 := GenS4{F0: uint(v5), F1: -835, F2: 0}
                                                                _ = v58
                                                            case "a":
                                                                v36 += (v36 | (v36 % ((v36) | 1)))
                                                            default:
                                                                var v59 int8 = int8(fi48)
                                                                _ = v59
                                                            }
                                                        } else {
                                                            v38 = fi49
                                                            var v60 int64 = v2
                                                            _ = v60
                                                        }
                                                        v61 := v3[v38]
                                                        switch (("" + ("\n\t" + "foo bar")) + "a") {
                                                        case "hello":
                                                            var v62 uint32 = 2
                                                            _ = v62
                                                        case "\n\t":
                                                            v5 = v5
                                                        case "foo bar":
                                                            v5 = (v5 & (v5 * v5))
                                                        }
                                                        v38 += v38
                                                        _ = v61
                                                    } else {
                                                        if len(v57) > 0 {
                                                            v57[0] = fi49
                                                        }
                                                        v5 -= v5
                                                        v63 := GenS6{F0: 248, F1: -2147483648}
                                                        _ = v63
                                                    }
                                                    v4 = 2147483647
                                                    var v64 string = "foo bar"
                                                    _ = v64
                                                } else {
                                                    if (true || ((("\n\t" + "hello") + "a") <= "foo bar")) {
                                                        fmt.Println(v36, v38)
                                                        var v65 uint = (v36 % (((v36 * v36)) | 1))
                                                        v2 -= v2
                                                        _ = v65
                                                    } else {
                                                        var v66 uint8 = v5
                                                        var v67 int16 = -305
                                                        fmt.Println(v36, v66, fi48)
                                                        _ = len(v6)
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                                _ = v57
                                            } else {
                                                v36 -= (v36 | 761)
                                                fmt.Println(v38, v36, v2)
                                            }
                                            _ = v56
                                        }
                                        if false {
                                            v68 := GenS6{F0: 790, F1: 0}
                                            _ = v68
                                        } else {
                                            var v69 int16 = 210
                                            if (true && false) {
                                                var v70 int64 = (v2 * v2)
                                                v71 := GenS7{F0: 18446744073709551615}
                                                switch fi49 {
                                                case 32767:
                                                    v72 := []int{-801, (fi49 / (((fi48 / ((4294967295) | 1))) | 1))}
                                                    _ = v72
                                                case 4294967295:
                                                    fmt.Println(v36, v5)
                                                case 32768:
                                                    v2 += v2
                                                }
                                                _ = v70
                                                _ = v71
                                            }
                                            if true {
                                                fmt.Println(fi48, v69, fi48)
                                                v36 = 766
                                                fmt.Println(fi49, v69, v36)
                                                fmt.Println(v38, v36)
                                                if (-1 <= -116) {
                                                    v73 := GenS8{F0: fi48, F1: (fi49 & (fi48 ^ -820))}
                                                    if false {
                                                        var v74 int8 = -1
                                                        _ = v74
                                                    }
                                                    _ = v73
                                                }
                                            }
                                            _ = v69
                                        }
                                        var v75 int32 = v4
                                        _ = v75
                                    case 9223372036854775807:
                                        switch ("hello" + ("hello" + "hello")) {
                                        case "a":
                                            _ = len(v3)
                                        case "hello":
                                            switch "" {
                                            case "hello":
                                                switch (v38 ^ v38) {
                                                case 555:
                                                    var v76 uint64 = 2
                                                    _ = v76
                                                case -1:
                                                    var v77 uint = v36
                                                    _ = v77
                                                }
                                            case "foo bar":
                                                var v78 uint16 = 0
                                                _ = v78
                                            }
                                        default:
                                            if (-1 > -128) {
                                                v4 = (v4 * int32(v2))
                                            }
                                        }
                                    case 2147483647:
                                        var v79 uint32 = 2147483648
                                        _ = v79
                                    default:
                                        var v80 string = ""
                                        _ = v80
                                    }
                                    var v81 int8 = 42
                                    fmt.Println(v2, v81)
                                    _ = v81
                                }
                                v36 -= (v36 >> ((42) & 63))
                                var v82 uint16 = 1
                                _ = fi49
                                _ = v82
                            }
                            _ = fi48
                        }
                        var v83 bool = false
                        v84 := GenS6{F0: 18446744073709551615, F1: (v4 << (((v36 % (((v36 + (v36 / ((v36) | 1)))) | 1))) & 63))}
                        _ = v83
                        _ = v84
                    }
                } else {
                    var v85 int8 = 127
                    var v86 int32 = -945
                    _ = v85
                    _ = v86
                }
            }
        }
    }
    fmt.Println(v4, v4, v2)
    fmt.Println(v2, v2)
    if (100 <= 42) {
        switch ("a" + ((("foo bar" + ("" + "\n\t")) + (("\n\t" + "\n\t") + "a")) + "foo bar")) {
        case "foo bar":
            if false {
                v87 := GenS1{F0: 1, F1: 1, F2: (1.5 == 0.0)}
                _ = v87
            } else {
                v88 := GenS2{F0: 0}
                switch v38 {
                case 2147483647:
                    fmt.Println(v5, v2, v38)
                case 42:
                    delete(v6, "hello")
                case -541:
                    v5 += v5
                }
                if (965 > 32767) {
                    v89 := []int{}
                    _ = v89
                } else {
                    for fi90 := 0; fi90 < 15; fi90++ {
                        v5 -= v5
                        var v91 uint16 = 2
                        _ = fi90
                        _ = v91
                    }
                    for fi92 := 0; fi92 < 8; fi92++ {
                        v5 += 127
                        _ = fi92
                    }
                }
                _ = v88
            }
        }
        v38 -= (v38 + v38)
        for fi93 := 0; fi93 < 20; fi93++ {
            v38 = fi93
            _ = fi93
        }
    }
    v4 += v4
    var v94 int64 = int64(v4)
    v2 -= int64(v38)
    v95 := GenS9{F0: v36}
    v96 := []int{v38, v38, (v38 + 9223372036854775807), v38}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v94
    _ = v95
    _ = v96
}
