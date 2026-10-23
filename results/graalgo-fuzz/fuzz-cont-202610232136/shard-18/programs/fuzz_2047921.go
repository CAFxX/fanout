package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 int; F1 int16 }
type GenS11 struct { F0 int64; F1 string; F2 uint64 }

func main() {
    var v1 uint = 255
    var v2 int8 = -1
    v3 := make(map[uint32]int)
    v2 += (v2 | -1)
    var v4 bool = false
    v1 = v1
    v5 := GenS1{F0: 0, F1: -2147483648}
    for fi6 := 0; fi6 < 18; fi6++ {
        if v4 {
            _ = len(v3)
        }
        v7 := GenS2{F0: 115}
        var v8 uint8 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 0
    if true {
        for fi10 := 0; fi10 < 19; fi10++ {
            if v4 {
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (127 <= v2) {
                        var v12 uint8 = (^(v9))
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 2, F1: 696}
                        v1 += (v1 + 255)
                        switch fi11 {
                        case 9223372036854775807:
                            switch (fi11 ^ (fi10 / ((4294967295) | 1))) {
                            case 72:
                                v14 := []int{fi11}
                                _ = v14
                            case -754:
                                var v15 int = int(v1)
                                _ = v15
                            case -712:
                                switch "a" {
                                case "a":
                                    fmt.Println(v2, v4, fi11)
                                }
                            }
                        case 32767:
                            if v4 {
                                v1 += v1
                                var v16 int8 = (v2 << ((v1) & 63))
                                fmt.Println(v16)
                                var v17 int8 = (v2 ^ v16)
                                if false {
                                    fmt.Println(v4, v2)
                                    v17 = (v2 + (v2 - (^(v2))))
                                    switch -1 {
                                    case 4294967295:
                                        v18 := []int{(fi11 << ((v1) & 63)), 711, (^(fi11)), -945}
                                        _ = v18
                                    case -977:
                                        v19 := []int{fi10, 9223372036854775807, (fi10 & 718), fi10}
                                        _ = v19
                                    case -9223372036854775808:
                                        if v4 {
                                            v16 += (v2 + (v17 * (v17 << ((785) & 63))))
                                        } else {
                                            v20 := GenS2{F0: 32767}
                                            _ = v20
                                        }
                                    default:
                                        fmt.Println(v2, v16, v1)
                                        v21 := GenS3{F0: 32767, F1: v1, F2: 67}
                                        delete(v3, 688)
                                        _ = v21
                                    }
                                }
                                _ = v16
                                _ = v17
                            } else {
                                v9 = v9
                                v22 := []int{}
                                _ = v22
                            }
                        }
                        v23 := GenS2{F0: 2}
                        if ("" >= "hello") {
                            v2 = 100
                            v24 := GenS3{F0: 65535, F1: (v1 & v1), F2: 841}
                            var v25 uint16 = 32767
                            fmt.Println(fi11, v25)
                            _ = v24
                            _ = v25
                        }
                        var v26 uint64 = uint64(v1)
                        _ = v13
                        _ = v23
                        _ = v26
                    }
                    v27 := GenS3{F0: 428, F1: (v1 ^ 255), F2: uint16(v1)}
                    var v28 string = "hello"
                    _ = fi11
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 13; fi29++ {
                    if ((998 != 32767) || v4) {
                        switch "foo bar" {
                        case "a":
                            var v30 int64 = 1
                            _ = v30
                        case "hello":
                            if (((0.0 != ((0.1 / 3.14159) - float64(1.5))) && v4) || v4) {
                                v1 -= v1
                                v4 = (false && (0 == -1))
                                var v31 uint16 = 153
                                switch (fi10 * fi29) {
                                case 2147483647:
                                    fmt.Println(v9)
                                case 9223372036854775807:
                                    switch (fi29 + (fi10 | fi29)) {
                                    case 666:
                                        var v32 bool = (true || v4)
                                        _ = v32
                                    default:
                                        var v33 string = ("foo bar" + "foo bar")
                                        _ = v33
                                    }
                                    var v34 uint64 = 18446744073709551615
                                    fmt.Println(v34)
                                    _ = v34
                                case 127:
                                    if ((v9 << ((v1) & 63)) >= v9) {
                                        switch (fi29 | fi29) {
                                        case 998:
                                            var v35 float64 = 0.1
                                            _ = v35
                                        case 465:
                                            var v36 float64 = (((0.1 / (0.1 - (0.0 - (3.14159 * 1.5)))) + 1e10) + -2.25)
                                            _ = v36
                                        case 1:
                                            v1 = (v1 * (v1 ^ (^(v1))))
                                        default:
                                            v37 := GenS4{F0: (v31 << (((^(v1))) & 63)), F1: v1}
                                            _ = v37
                                        }
                                    }
                                    var v38 int16 = 0
                                    _ = v38
                                }
                                _ = v31
                            }
                        case "foo bar":
                            fmt.Println(v2)
                        default:
                            v2 += v2
                            fmt.Println(fi10)
                            if v4 {
                                v39 := GenS5{F0: -851, F1: (fi10 ^ fi10)}
                                _ = v39
                            } else {
                                v2 += (v2 - v2)
                                v40 := []int{(-(fi29))}
                                _ = v40
                            }
                        }
                        v41 := v3[2]
                        _ = v41
                    }
                    var v42 int8 = v2
                    _ = fi29
                    _ = v42
                }
                v43 := GenS6{F0: -9223372036854775808, F1: 703}
                _ = v43
            }
            switch "foo bar" {
            case "foo bar":
                if ((v1 * 104) <= v1) {
                    var v44 uint64 = 1
                    if ((v44 - v44) < (^(v44))) {
                        v4 = false
                        v2 -= (v2 | v2)
                    } else {
                        if (!v4) {
                            var v45 uint16 = 65535
                            fmt.Println(v4, v1)
                            if ((((3.14159 / -2.25) + float32(0.1)) >= 0.0) || ((v2 % ((v2) | 1)) == -114)) {
                                fmt.Println(v1, fi10, fi10)
                            } else {
                                var v46 uint = v1
                                for fi47 := 0; fi47 < 14; fi47++ {
                                    fmt.Println(v1)
                                    switch ("hello" + "") {
                                    case "a":
                                        var v48 string = ""
                                        _ = v48
                                    case "\n\t":
                                        v9 = 2
                                    case "foo bar":
                                        fmt.Println(v2)
                                    default:
                                        if v4 {
                                            if (-1 > 630) {
                                                switch "" {
                                                case "a":
                                                    v49 := GenS5{F0: 0, F1: fi10}
                                                    if v4 {
                                                        v50 := GenS6{F0: 894, F1: (^(v44))}
                                                        var v51 int64 = -9223372036854775808
                                                        v52 := v3[65535]
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    _ = v49
                                                case "":
                                                    if true {
                                                        if (0 == 496) {
                                                            v53 := GenS6{F0: -73, F1: v44}
                                                            delete(v3, 774)
                                                            _ = v53
                                                        } else {
                                                            var v54 uint = v1
                                                            _ = v54
                                                        }
                                                        switch fi10 {
                                                        case -432:
                                                            switch fi10 {
                                                            case 4294967295:
                                                                v44 += 18446744073709551615
                                                            case 649:
                                                                v3[65535] = fi47
                                                                if v4 {
                                                                    fmt.Println(v46, v4, v46)
                                                                    var v55 uint64 = v44
                                                                    if v4 {
                                                                        var v56 int8 = int8(v45)
                                                                        _ = v56
                                                                    } else {
                                                                        fmt.Println(fi47, fi10)
                                                                    }
                                                                    _ = v55
                                                                }
                                                                v1 = 4294967295
                                                                fmt.Println(v2, v2)
                                                            default:
                                                                if (true && (-1 < -90)) {
                                                                    v57 := GenS7{F0: true}
                                                                    v58 := GenS1{F0: v9, F1: 1}
                                                                    v46 = v1
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    var v59 uint8 = v9
                                                                    if v4 {
                                                                        var v60 uint = v46
                                                                        _ = v60
                                                                    } else {
                                                                        switch (("hello" + ("a" + "a")) + "foo bar") {
                                                                        case "hello":
                                                                            fmt.Println(v4)
                                                                        default:
                                                                            if v4 {
                                                                                _ = len(v3)
                                                                                fmt.Println(v1, fi10, v9)
                                                                                _ = len(v3)
                                                                            } else {
                                                                                v61 := GenS8{F0: false, F1: v46}
                                                                                v4 = (false && true)
                                                                                v62 := []int{-1, (fi10 >> ((65535) & 63)), fi47}
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                        }
                                                                    }
                                                                    if (-824 == -1) {
                                                                        if (4294967295 > uint(v44)) {
                                                                            var v63 int8 = v2
                                                                            v45 = 255
                                                                            v63 -= (v2 + 0)
                                                                            var v64 int16 = -32768
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v65 uint = 980
                                                                                _ = v65
                                                                            default:
                                                                                switch fi47 {
                                                                                case 127:
                                                                                    if v4 {
                                                                                        v66 := GenS3{F0: 24, F1: v46, F2: (v45 << ((65535) & 63))}
                                                                                        _ = v66
                                                                                    } else {
                                                                                        fmt.Println(fi47)
                                                                                        var v67 uint8 = v9
                                                                                        var v68 int64 = 1
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                    }
                                                                                default:
                                                                                    v4 = v4
                                                                                }
                                                                                if v4 {
                                                                                    var v69 string = "a"
                                                                                    v1 = 525
                                                                                    if v4 {
                                                                                        v45 = (v45 - (v45 % ((v45) | 1)))
                                                                                        var v70 uint32 = 2147483648
                                                                                        if (uint(v70) < (^(v46))) {
                                                                                            fmt.Println(v45, v59)
                                                                                            v9 -= (^(v59))
                                                                                        }
                                                                                        _ = v70
                                                                                    }
                                                                                    _ = v69
                                                                                }
                                                                            }
                                                                            var v71 float64 = 1.5
                                                                            _ = v71
                                                                        }
                                                                        fmt.Println(fi47, v46, fi10)
                                                                    } else {
                                                                        v46 += 0
                                                                        v3[4294967295] = fi47
                                                                        v3[2] = fi47
                                                                        fmt.Println(v9, v9, v46)
                                                                    }
                                                                    if (false || v4) {
                                                                        var v72 int64 = -692
                                                                        switch fi47 {
                                                                        case 789:
                                                                            v73 := []int{fi47}
                                                                            _ = v73
                                                                        case 32768:
                                                                            v4 = v4
                                                                        default:
                                                                            var v74 uint32 = 4294967295
                                                                            _ = v74
                                                                        }
                                                                        v72 = v72
                                                                        _ = v72
                                                                    }
                                                                    _ = v59
                                                                }
                                                            }
                                                        case -124:
                                                            v44 = (v44 >> (((v1 | 2)) & 63))
                                                        default:
                                                            v1 -= (v46 << ((v1) & 63))
                                                        }
                                                    }
                                                }
                                                v75 := []int{(fi47 & fi47), fi10, fi47}
                                                v76 := GenS6{F0: 9223372036854775807, F1: (v44 << (((v1 + v1)) & 63))}
                                                var v77 float32 = -2.25
                                                fmt.Println(v45, v1)
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            } else {
                                                var v78 int = 822
                                                switch (fi47 ^ v78) {
                                                case 124:
                                                    if v4 {
                                                        v45 = v45
                                                        var v79 float32 = (((((1.5 + (float32(float64(0.1)) - 3.14159)) * float32(-2.25)) + 3.14159) - 1.5) * float32((1e10 + -2.25)))
                                                        _ = v79
                                                    }
                                                case -615:
                                                    switch fi10 {
                                                    case 32768:
                                                        var v80 float32 = 1.5
                                                        _ = v80
                                                    }
                                                }
                                                v81 := []int{(v78 >> ((uint(v44)) & 63)), (^(v78))}
                                                _ = v78
                                                _ = v81
                                            }
                                        }
                                    }
                                    v82 := []int{128, (fi47 + 4294967295), (fi47 * 0), fi47}
                                    _ = fi47
                                    _ = v82
                                }
                                var v83 float64 = 1e10
                                _ = v46
                                _ = v83
                            }
                            _ = v45
                        } else {
                            v84 := GenS9{F0: 2, F1: ""}
                            _ = v84
                        }
                    }
                    _ = v44
                }
            }
            switch (fi10 * fi10) {
            case 9223372036854775807:
                v85 := []int{-621}
                _ = v85
            case -694:
                var v86 bool = (((^(fi10)) == 127) || v4)
                _ = v86
            case 32767:
                var v87 uint64 = 42
                _ = v87
            default:
                var v88 uint16 = 2
                _ = v88
            }
            _ = fi10
        }
    }
    var v89 int8 = v2
    switch ("" + ("hello" + ("hello" + ("" + "a")))) {
    case "":
        v90 := []int{9223372036854775807, 4294967295}
        _ = v90
    case "hello":
        v3[2147483648] = 124
    case "a":
        for fi91 := 0; fi91 < 2; fi91++ {
            for fi92 := 0; fi92 < 11; fi92++ {
                var v93 int32 = 183
                v94 := GenS8{F0: (!((fi91 % (((-(fi91))) | 1)) <= fi92)), F1: 4294967295}
                _ = fi92
                _ = v93
                _ = v94
            }
            v2 += v2
            v95 := []int{}
            _ = fi91
            _ = v95
        }
    default:
        var v96 uint = (v1 & (v1 - 42))
        _ = v96
    }
    for fi97 := 0; fi97 < 13; fi97++ {
        switch (fi97 << (((^(v1))) & 63)) {
        case 226:
            for fi98 := 0; fi98 < 2; fi98++ {
                var v99 int64 = int64(v1)
                _ = fi98
                _ = v99
            }
        }
        var v100 bool = (!(!true))
        for fi101 := 0; fi101 < 10; fi101++ {
            if false {
                var v102 int8 = 100
                v103 := GenS10{F0: int(v1), F1: 32767}
                v104 := GenS11{F0: int64(v1), F1: "foo bar", F2: 18446744073709551615}
                _ = v102
                _ = v103
                _ = v104
            }
            _ = fi101
        }
        _ = fi97
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v9
}
