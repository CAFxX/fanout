package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 int16; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 128
    var v3 uint32 = 32
    v4 := make(map[uint32]uint32)
    v2 = (v2 >> ((42) & 63))
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: ("foo bar" + ""), F1: 2147483647, F2: true}
            _ = len(v4)
            if (fi6 < (fi5 + 615)) {
                if (true || false) {
                    v3 += (v3 | (v3 - (v3 + (v3 * v3))))
                    var v8 int8 = 100
                    _ = v8
                }
                fmt.Println(fi5)
                v9 := []int{}
                _ = v9
            } else {
                v2 += v2
                switch (fi6 & fi6) {
                case 0:
                    v10 := GenS1{F0: ("\n\t" + ("hello" + "hello")), F1: 1, F2: false}
                    _ = v10
                case -628:
                    v2 -= 255
                default:
                    switch fi6 {
                    case 9223372036854775807:
                        v2 = (v2 / ((223) | 1))
                    default:
                        switch fi6 {
                        case -327:
                            var v11 uint16 = 65535
                            _ = v11
                        case -1:
                            fmt.Println(fi5, v3, fi5)
                        case 2147483647:
                            fmt.Println(v3)
                        }
                    }
                }
                if ("hello" > "a") {
                    switch "hello" {
                    case "a":
                        delete(v4, (^(v3)))
                    case "foo bar":
                        if (0.1 > 1.5) {
                            v2 -= v2
                        } else {
                            if (!((0 < 1) && (!true))) {
                                v12 := GenS2{F0: 0, F1: 42, F2: 658}
                                _ = len(v1)
                                var v13 int64 = 1
                                _ = v12
                                _ = v13
                            } else {
                                v14 := []int{(^(fi6)), 13, (fi6 << ((282) & 63)), (fi5 % ((fi5) | 1))}
                                var v15 uint64 = 1
                                _ = v14
                                _ = v15
                            }
                        }
                        v16 := []int{fi5}
                        _ = v16
                    default:
                        if false {
                            var v17 int32 = 42
                            if true {
                                v3 += 2147483648
                                v18 := GenS1{F0: "foo bar", F1: (v17 ^ (v17 ^ 0)), F2: (((3.14159 * 0.1) * float64((float32(((1e10 / 0.1) - 1e10)) * 3.14159))) >= (((0.0 - 0.1) + (-2.25 + 1.5)) + (1.5 * -2.25)))}
                                _ = v18
                            }
                            v2 -= 130
                            switch "a" {
                            case "hello":
                                if true {
                                    v17 = v17
                                    if ("hello" <= "a") {
                                        switch fi6 {
                                        case -9223372036854775808:
                                            v19 := GenS1{F0: ("hello" + ("hello" + (("foo bar" + "hello") + "\n\t"))), F1: v17, F2: false}
                                            _ = v19
                                        case 0:
                                            var v20 float32 = (1.5 * float32(float64(1.5)))
                                            _ = v20
                                        }
                                        var v21 int64 = 9223372036854775807
                                        _ = v21
                                    } else {
                                        v17 += v17
                                        if (false || true) {
                                            var v22 int16 = 1000
                                            v3 += 305
                                            v2 = v2
                                            _ = v22
                                        } else {
                                            v3 = v3
                                            var v23 int = -493
                                            _ = v23
                                        }
                                    }
                                }
                            }
                            v24 := GenS1{F0: ("a" + "\n\t"), F1: -445, F2: false}
                            _ = v17
                            _ = v24
                        } else {
                            if (v2 < (v2 << ((370) & 63))) {
                                if true {
                                    if true {
                                        v2 += v2
                                        v25 := GenS1{F0: "a", F1: -2147483648, F2: false}
                                        _ = v25
                                    }
                                    v1[32768] = fi6
                                    v26 := GenS1{F0: (("a" + "\n\t") + "hello"), F1: -1, F2: ((1e10 / 3.14159) == 3.14159)}
                                    _ = v26
                                } else {
                                    if false {
                                        if false {
                                            var v27 bool = (673 < 4294967295)
                                            var v28 int16 = 32767
                                            var v29 uint16 = 42
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        var v30 uint64 = 0
                                        if true {
                                            v2 -= 128
                                        } else {
                                            if false {
                                                _ = len(v4)
                                                if false {
                                                    var v31 int64 = 0
                                                    fmt.Println(v3, fi6)
                                                    _ = v31
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v32 := GenS3{F0: v2, F1: 283, F2: false}
                                                        _ = v32
                                                    case "a":
                                                        switch (fi6 * fi6) {
                                                        case -1:
                                                            v33 := GenS2{F0: 889, F1: 108, F2: 0}
                                                            _ = v33
                                                        case 0:
                                                            fmt.Println(v30)
                                                        case 441:
                                                            switch fi5 {
                                                            case 0:
                                                                var v34 int = (fi6 % ((fi6) | 1))
                                                                _ = v34
                                                            case 1:
                                                                v1[fi5] = -483
                                                            default:
                                                                v35 := GenS4{F0: -32768, F1: (v3 >> ((648) & 63)), F2: -174}
                                                                _ = v35
                                                            }
                                                        }
                                                    case "":
                                                        v30 = uint64(v3)
                                                    }
                                                }
                                            }
                                            v30 = (v30 << ((18446744073709551615) & 63))
                                        }
                                        _ = v30
                                    } else {
                                        switch fi6 {
                                        case 32767:
                                            v2 -= v2
                                        case 9223372036854775807:
                                            var v36 uint = 65535
                                            _ = v36
                                        default:
                                            if (((1.5 + 1.5) >= float64(1.5)) || (!(false && (-64 >= 100)))) {
                                                var v37 int = -1
                                                fmt.Println(fi5, v3)
                                                fmt.Println(v37)
                                                _ = v37
                                            } else {
                                                v38 := v1[(-(fi6))]
                                                _ = v38
                                            }
                                        }
                                        v3 = 969
                                        var v39 bool = (int64(v3) >= -1)
                                        _ = v39
                                    }
                                    if false {
                                        var v40 string = "a"
                                        v2 = (v2 >> ((42) & 63))
                                        v41 := []int{fi5, 128}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v42 := []int{fi6, (^(fi5)), fi5, fi5}
                                        var v43 uint64 = 137
                                        if (0.1 == (1.5 / 1.0)) {
                                            var v44 bool = ((7 == -2147483648) || (1.5 >= 0.0))
                                            if ((v2 & v2) >= 1) {
                                                if v44 {
                                                    if (!(!v44)) {
                                                        switch (("a" + "\n\t") + "a") {
                                                        case "\n\t":
                                                            v45 := GenS4{F0: 1000, F1: (^(v3)), F2: -1}
                                                            _ = v45
                                                        case "hello":
                                                            if v44 {
                                                                var v46 float32 = float32(3.14159)
                                                                v47 := GenS2{F0: 0, F1: -2147483648, F2: (v43 << ((0) & 63))}
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        }
                                                        var v48 uint16 = 32767
                                                        switch fi5 {
                                                        case 0:
                                                            if ((^(v3)) >= v3) {
                                                                v48 -= v48
                                                                var v49 uint64 = (v43 - (v43 ^ 673))
                                                                _ = v49
                                                            } else {
                                                                v50 := GenS5{F0: (^(v2)), F1: 18446744073709551615}
                                                                _ = v50
                                                            }
                                                        case 32768:
                                                            var v51 int32 = 1
                                                            _ = v51
                                                        case 2147483647:
                                                            switch (^(fi6)) {
                                                            case 128:
                                                                v43 = (v43 + (v43 ^ 18446744073709551615))
                                                                v52 := v1[fi6]
                                                                v53 := GenS6{F0: "hello", F1: 18446744073709551615}
                                                                var v54 float32 = ((((-2.25 * (0.1 / ((0.1 * 0.0) + (((((1e10 / 1e10) - 1.5) * 1.5) + (1e10 - 0.1)) - -2.25)))) + float32((((((float64(float32(1e10)) + 0.0) * (0.1 - ((0.0 / 1.5) - 0.1))) - 1e10) / 0.1) / 1.5))) * (-2.25 / float32(3.14159))) / 3.14159)
                                                                switch 42 {
                                                                case 4294967295:
                                                                    v42 = append(v42, (fi6 - v52))
                                                                case 128:
                                                                    switch v52 {
                                                                    case -26:
                                                                        v52 += 42
                                                                    case -1:
                                                                        var v55 int64 = 1
                                                                        _ = v55
                                                                    default:
                                                                        if true {
                                                                            switch (fi6 + int(v2)) {
                                                                            case 128:
                                                                                var v56 bool = v44
                                                                                _ = v56
                                                                            case 32768:
                                                                                var v57 uint16 = (v48 + (v48 ^ 42))
                                                                                _ = v57
                                                                            default:
                                                                                var v58 uint32 = (v3 & v3)
                                                                                _ = v58
                                                                            }
                                                                            var v59 int32 = 42
                                                                            v60 := GenS5{F0: 255, F1: (v43 & v43)}
                                                                            var v61 uint32 = (^(v3))
                                                                            _ = v59
                                                                            _ = v60
                                                                            _ = v61
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                if len(v42) > 0 {
                                                                                    v42[0] = fi5
                                                                                }
                                                                                var v62 int8 = int8(v3)
                                                                                v42 = append(v42, (v52 >> ((42) & 63)))
                                                                                _ = v62
                                                                            }
                                                                            if len(v42) > 0 {
                                                                                v42[0] = (fi5 / ((fi5) | 1))
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                v52 = fi6
                                                                            case "a":
                                                                                delete(v4, (v3 >> ((352) & 63)))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v52
                                                                _ = v53
                                                                _ = v54
                                                            case 4294967295:
                                                                if false {
                                                                    if ((false && v44) || false) {
                                                                        if v44 {
                                                                            _ = len(v42)
                                                                            v43 -= v43
                                                                            switch (-(fi5)) {
                                                                            case 2147483647:
                                                                                var v63 int8 = 1
                                                                                _ = v63
                                                                            default:
                                                                                if (32768 != fi6) {
                                                                                    if false {
                                                                                        if v44 {
                                                                                            if len(v42) > 0 {
                                                                                                v42[0] = fi5
                                                                                            }
                                                                                        } else {
                                                                                            var v64 float32 = (1e10 / (3.14159 / 1e10))
                                                                                            _ = v64
                                                                                        }
                                                                                        v43 += v43
                                                                                        var v65 float64 = (0.0 + (0.0 / 1.0))
                                                                                        _ = v65
                                                                                    }
                                                                                    v43 += (^(v43))
                                                                                    v44 = false
                                                                                }
                                                                            }
                                                                        }
                                                                        v66 := GenS3{F0: (v2 >> ((867) & 63)), F1: 9223372036854775807, F2: v44}
                                                                        _ = v66
                                                                    }
                                                                    v3 = 2
                                                                    switch -9223372036854775808 {
                                                                    case 127:
                                                                        switch (fi5 + fi6) {
                                                                        case -513:
                                                                            v48 -= v48
                                                                        }
                                                                    default:
                                                                        if len(v42) >= 2 {
                                                                            v42 = v42[1:len(v42)]
                                                                        }
                                                                    }
                                                                }
                                                            case 127:
                                                                v1[fi5] = fi6
                                                            default:
                                                                v44 = v44
                                                            }
                                                        }
                                                        var v67 bool = v44
                                                        var v68 uint8 = uint8(v3)
                                                        _ = v48
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    var v69 int16 = 502
                                                    _ = v69
                                                }
                                            }
                                            switch (fi6 % ((fi5) | 1)) {
                                            case 32768:
                                                fmt.Println(v44, fi5, v3)
                                            default:
                                                v2 -= (v2 | (v2 * v2))
                                            }
                                            _ = v44
                                        } else {
                                            v70 := GenS7{F0: -66, F1: 698}
                                            _ = v70
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    if (false && ("\n\t" == ("hello" + "\n\t"))) {
                                        if true {
                                            v71 := v4[996]
                                            v71 += v3
                                            var v72 int8 = 1
                                            v73 := GenS8{F0: 479, F1: (v2 - v2)}
                                            v3 = (v71 << ((255) & 63))
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        }
                                        var v74 int8 = -93
                                        _ = v74
                                    }
                                }
                                var v75 bool = false
                                _ = v75
                            }
                        }
                        var v76 uint32 = 384
                        _ = v76
                    }
                    var v77 uint32 = 2147483648
                    _ = v77
                }
            }
            _ = fi6
            _ = v7
        }
        v78 := []int{(fi5 % ((589) | 1)), fi5}
        if len(v78) > 0 {
            v78[0] = (fi5 - fi5)
        }
        _ = fi5
        _ = v78
    }
    for fi79 := 0; fi79 < 1; fi79++ {
        v2 = (v2 | (v2 << ((558) & 63)))
        for fi80 := 0; fi80 < 9; fi80++ {
            v81 := GenS2{F0: -1, F1: 42, F2: 4294967295}
            v82 := []int{fi79}
            _ = fi80
            _ = v81
            _ = v82
        }
        _ = fi79
    }
    switch 42 {
    case 32768:
        if true {
            v83 := GenS5{F0: v2, F1: 4294967295}
            v84 := GenS4{F0: 0, F1: (v3 << ((413) & 63)), F2: -252}
            fmt.Println(v3)
            _ = v83
            _ = v84
        }
    case 706:
        v85 := v4[416]
        _ = v85
    case 0:
        var v86 float32 = ((1.5 + (1.5 * (0.1 * (-2.25 * float32((0.1 - 0.1)))))) * 0.0)
        _ = v86
    }
    for fi87 := 0; fi87 < 12; fi87++ {
        for fi88 := 0; fi88 < 9; fi88++ {
            fmt.Println(v2)
            var v89 int32 = 1
            v3 = 524
            _ = fi88
            _ = v89
        }
        v90 := GenS9{F0: (v3 >> ((4294967295) & 63)), F1: (fi87 + fi87)}
        v91 := GenS7{F0: -76, F1: -1}
        _ = fi87
        _ = v90
        _ = v91
    }
    v92 := v4[(v3 & (v3 >> ((255) & 63)))]
    fmt.Println(v2, v92, v3)
    v2 = (v2 * v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v92
}
