package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 string }
func genFn0(p0 uint16, p1 uint) bool {
    v6 := []int{2, 972, 42, 128}
    var v7 int8 = -1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (18446744073709551615 >= 1)
}
func genFn1() int {
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint32 = 1
            v11 := []int{fi8, fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            if (!(fi12 == (fi12 << ((65535) & 63)))) {
                v13 := GenS1{F0: (1 != 42)}
                _ = v13
            }
            switch fi8 {
            case 0:
                var v14 uint32 = 65535
                _ = v14
            case -491:
                var v15 int16 = -1
                _ = v15
            case -2147483648:
                var v16 int16 = 509
                _ = v16
            default:
                switch (fi12 % (((fi8 + fi8)) | 1)) {
                case -540:
                    v17 := GenS2{F0: true, F1: fi8, F2: 2}
                    _ = v17
                case 0:
                    var v18 int32 = -928
                    _ = v18
                case -2147483648:
                    switch fi12 {
                    case 9223372036854775807:
                        v19 := []int{4294967295, 32768, fi8, 2147483647}
                        _ = v19
                    default:
                        switch 127 {
                        case 128:
                            v20 := GenS2{F0: false, F1: fi8, F2: 2147483648}
                            _ = v20
                        default:
                            var v21 uint64 = 0
                            _ = v21
                        }
                    }
                default:
                    switch 1 {
                    case 9223372036854775807:
                        var v22 uint64 = 1
                        switch (fi8 / ((fi8) | 1)) {
                        case 2147483647:
                            var v23 bool = false
                            _ = v23
                        case -853:
                            v24 := []int{(fi8 ^ (^(fi8))), fi8}
                            _ = v24
                        default:
                            v25 := []int{fi8, 128}
                            _ = v25
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                v26 := []int{1, (fi12 ^ fi8)}
                                if ((255 == 255) || false) {
                                    var v27 uint = 0
                                    v28 := GenS2{F0: ((v27 == 0) && (1000 == 89)), F1: fi12, F2: 0}
                                    switch fi12 {
                                    case -9223372036854775808:
                                        v27 -= 18446744073709551615
                                    }
                                    _ = v27
                                    _ = v28
                                } else {
                                    v29 := GenS1{F0: (float32(1e10) <= float32(1e10))}
                                    v30 := GenS3{F0: (32767 == 73), F1: v22}
                                    v22 += v22
                                    var v31 int32 = 1
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v26
                            } else {
                                if false {
                                    var v32 int16 = 1000
                                    var v33 int64 = 1
                                    v34 := GenS1{F0: false}
                                    v33 -= -1
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 float32 = 3.14159
                                    _ = v35
                                }
                                v22 = 42
                            }
                        }
                        var v36 int64 = -938
                        _ = v22
                        _ = v36
                    case -822:
                        if (42 >= 9223372036854775807) {
                            var v37 bool = false
                            if (-2.25 != -2.25) {
                                v37 = v37
                            }
                            var v38 int16 = 1
                            _ = v37
                            _ = v38
                        }
                    case 779:
                        var v39 int16 = 1
                        _ = v39
                    }
                }
            }
            var v40 uint64 = 24
            _ = fi12
            _ = v40
        }
        _ = fi8
    }
    v41 := GenS1{F0: (88 >= 18446744073709551615)}
    _ = v41
    return 127
}

func main() {
    v1 := []int{-9223372036854775808, -9223372036854775808, 9223372036854775807, 9223372036854775807}
    var v2 uint16 = 0
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v5 := []uint32{}
    if true {
        for fi42 := 0; fi42 < 20; fi42++ {
            if false {
                for fi43 := 0; fi43 < 15; fi43++ {
                    switch fi42 {
                    case -9223372036854775808:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -1:
                        if false {
                            v2 -= 0
                            v2 -= 206
                        } else {
                            v44 := v4[2147483648]
                            v45 := GenS3{F0: (4294967295 <= 42), F1: 18446744073709551615}
                            v46 := GenS4{F0: 0, F1: 919, F2: ((float64(1e10) <= (1e10 * 1e10)) && true)}
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                    default:
                        v2 = v2
                    }
                    var v47 uint8 = 0
                    v47 += (v47 | (v47 / ((v47) | 1)))
                    var v48 uint16 = v2
                    if (true && ((!((v47 + (v47 / (((v47 | 128)) | 1))) >= v47)) && (4294967295 <= 0))) {
                        if true {
                            if false {
                                var v49 bool = false
                                v50 := GenS4{F0: -493, F1: 18446744073709551615, F2: v49}
                                _ = v49
                                _ = v50
                            }
                        } else {
                            var v51 bool = (!(0 != -110))
                            _ = v51
                        }
                        switch "\n\t" {
                        case "hello":
                            v52 := GenS4{F0: 42, F1: 1, F2: (42 != 1)}
                            v48 += v48
                            if false {
                                if true {
                                    var v53 uint8 = 2
                                    _ = v53
                                } else {
                                    v54 := GenS3{F0: ((v47 < 109) && true), F1: 588}
                                    v48 += (^(v48))
                                    v48 += v2
                                    fmt.Println(v47)
                                    _ = v54
                                }
                                v2 = (^(v48))
                                var v55 int16 = 32767
                                switch fi43 {
                                case 32768:
                                    v56 := v3[65535]
                                    _ = v56
                                case 128:
                                    if true {
                                        v57 := GenS1{F0: (v55 >= 24)}
                                        _ = v57
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v58 uint32 = 0
                                        _ = v58
                                    }
                                    v59 := GenS5{F0: (("foo bar" + "foo bar") + ("\n\t" + "")), F1: 42}
                                    _ = v59
                                case -831:
                                    v1 = append(v1, (fi42 ^ (fi42 / (((fi43 + genFn1())) | 1))))
                                }
                                _ = v55
                            }
                            _ = v52
                        case "foo bar":
                            v60 := v3[4294967295]
                            _ = v60
                        default:
                            if true {
                                var v61 uint64 = 490
                                _ = v61
                            } else {
                                var v62 uint = 262
                                if true {
                                    var v63 int64 = -1
                                    switch (fi42 % ((127) | 1)) {
                                    case 32767:
                                        switch fi42 {
                                        case 32768:
                                            switch (fi42 * 949) {
                                            case -1:
                                                v62 = (v62 - v62)
                                            case -2147483648:
                                                var v64 int32 = 1
                                                _ = v64
                                            default:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        case -1:
                                            if len(v5) > 0 {
                                                v5[0] = 230
                                            }
                                        case -223:
                                            if (int64(v62) != v63) {
                                                var v65 uint8 = (^(v47))
                                                v63 += v63
                                                var v66 uint16 = v2
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                    case 96:
                                        v63 = v63
                                    }
                                    delete(v3, 977)
                                    _ = v63
                                }
                                var v67 uint64 = 18446744073709551615
                                _ = v62
                                _ = v67
                            }
                        }
                    }
                    if ((302 != 0) || true) {
                        delete(v4, 376)
                        var v68 string = "hello"
                        _ = v68
                    }
                    _ = fi43
                    _ = v47
                    _ = v48
                }
                var v69 int = 0
                if (true && (2147483647 >= 2147483648)) {
                    v3[90] = "foo bar"
                } else {
                    v1 = append(v1, (v69 << ((65535) & 63)))
                    fmt.Println(v69, v69, fi42)
                    var v70 uint = 42
                    _ = v70
                }
                _ = v69
            }
            var v71 int = fi42
            _ = fi42
            _ = v71
        }
        for fi72 := 0; fi72 < 20; fi72++ {
            if true {
                for fi73 := 0; fi73 < 20; fi73++ {
                    fmt.Println(fi72, fi73)
                    v2 = (^(v2))
                    v2 -= 115
                    _ = fi73
                }
            } else {
                var v74 float64 = 0.0
                fmt.Println(fi72, v2)
                var v75 uint16 = 611
                _ = v74
                _ = v75
            }
            v1 = append(v1, (^(fi72)))
            _ = fi72
        }
    }
    for fi76 := 0; fi76 < 17; fi76++ {
        var v77 int32 = 375
        _ = fi76
        _ = v77
    }
    if false {
        if true {
            if true {
                for fi78 := 0; fi78 < 10; fi78++ {
                    var v79 uint32 = 304
                    var v80 int = fi78
                    for fi81 := 0; fi81 < 2; fi81++ {
                        if false {
                            v79 = (^(v79))
                        } else {
                            switch genFn1() {
                            case 32768:
                                var v82 string = "foo bar"
                                _ = v82
                            }
                            fmt.Println(fi78)
                            var v83 int16 = 32767
                            _ = v83
                        }
                        v2 = (v2 | v2)
                        var v84 float64 = 0.0
                        _ = fi81
                        _ = v84
                    }
                    _ = fi78
                    _ = v79
                    _ = v80
                }
            } else {
                v2 = 42
                if (!(2 <= 65535)) {
                    var v85 uint64 = 71
                    _ = v85
                } else {
                    v2 = v2
                    switch ("" + "") {
                    case "hello":
                        v2 += 65535
                    case "\n\t":
                        v2 += 531
                    case "":
                        v2 -= v2
                    }
                    v86 := GenS5{F0: "", F1: -9223372036854775808}
                    _ = v86
                }
                if true {
                    for fi87 := 0; fi87 < 10; fi87++ {
                        switch ("foo bar" + "") {
                        case "foo bar":
                            if false {
                                var v88 float32 = 3.14159
                                _ = v88
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi87
                    }
                    fmt.Println(v2)
                    if (0.1 > -2.25) {
                        var v89 float32 = 0.0
                        _ = v89
                    }
                } else {
                    if true {
                        for fi90 := 0; fi90 < 1; fi90++ {
                            fmt.Println(v2, v2)
                            for fi91 := 0; fi91 < 12; fi91++ {
                                v2 = 255
                                var v92 bool = true
                                var v93 uint64 = 1
                                _ = fi91
                                _ = v92
                                _ = v93
                            }
                            var v94 uint16 = (v2 - v2)
                            _ = fi90
                            _ = v94
                        }
                        var v95 bool = false
                        _ = v95
                    }
                    if true {
                        var v96 uint64 = 18446744073709551615
                        if true {
                            v97 := GenS5{F0: "a", F1: 42}
                            _ = v97
                        } else {
                            fmt.Println(v96, v96)
                        }
                        v98 := GenS6{F0: ""}
                        _ = v96
                        _ = v98
                    }
                }
            }
            if true {
                v2 += (v2 % (((v2 | v2)) | 1))
                _ = len(v5)
                if (!false) {
                    for fi99 := 0; fi99 < 9; fi99++ {
                        fmt.Println(v2, v2)
                        if (-105 > 0) {
                            v2 = 1
                            fmt.Println(fi99, v2)
                        } else {
                            fmt.Println(fi99, fi99)
                            v2 = (v2 - 65535)
                        }
                        _ = fi99
                    }
                    for fi100 := 0; fi100 < 7; fi100++ {
                        if (!true) {
                            v2 += (v2 * 729)
                            switch "\n\t" {
                            case "foo bar":
                                v2 -= 299
                            case "hello":
                                if (255 == 127) {
                                    v2 -= (v2 | v2)
                                }
                            case "\n\t":
                                v2 -= 0
                            default:
                                var v101 int16 = 42
                                _ = v101
                            }
                            v102 := GenS6{F0: "a"}
                            _ = v102
                        } else {
                            fmt.Println(v2, v2)
                        }
                        v2 = v2
                        _ = fi100
                    }
                }
            }
        }
        for fi103 := 0; fi103 < 11; fi103++ {
            v104 := GenS2{F0: (false && (fi103 <= (fi103 << ((1) & 63)))), F1: 381, F2: 42}
            _ = fi103
            _ = v104
        }
        switch 128 {
        case 127:
            for fi105 := 0; fi105 < 12; fi105++ {
                v106 := GenS2{F0: false, F1: genFn1(), F2: 2147483647}
                _ = fi105
                _ = v106
            }
        default:
            v4[65535] = ("hello" + ("hello" + "\n\t"))
        }
    } else {
        fmt.Println(v2)
        fmt.Println(v2)
    }
    v2 += 65535
    fmt.Println(v2)
    v2 -= (v2 % ((v2) | 1))
    v107 := GenS5{F0: ("foo bar" + "hello"), F1: 0}
    var v108 uint = 255
    v3[4294967295] = "a"
    for fi109 := 0; fi109 < 8; fi109++ {
        var v110 uint = v108
        _ = fi109
        _ = v110
    }
    fmt.Println(v2, v2)
    _ = len(v1)
    var v111 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v108:", v108)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v107
    _ = v108
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
