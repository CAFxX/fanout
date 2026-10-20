package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int8; F1 string; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 uint32 }

func main() {
    var v1 int = -636
    var v2 string = ""
    var v3 int8 = 94
    v4 := make(map[string]uint32)
    var v5 int = 127
    fmt.Println(v3, v5, v2)
    if (!true) {
        var v6 string = (v2 + v2)
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint = 2
            if (-1 == 0) {
                if (!true) {
                    if (-833 != -872) {
                        fmt.Println(fi7)
                        fmt.Println(v2, v6)
                        var v9 float64 = float64(float32(3.14159))
                        _ = v9
                    } else {
                        var v10 uint16 = 1
                        v1 = fi7
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 int8 = (v3 % ((v3) | 1))
                            var v13 bool = false
                            var v14 float64 = -2.25
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = v10
                    }
                    var v15 uint = (v8 - v8)
                    _ = v15
                }
                v3 = v3
                for fi16 := 0; fi16 < 13; fi16++ {
                    v2 = v2
                    _ = fi16
                }
            } else {
                v1 = (-(v5))
            }
            var v17 string = v2
            _ = fi7
            _ = v17
            _ = v8
        }
        switch v2 {
        case "\n\t":
            v18 := GenS1{F0: 18446744073709551615, F1: 0}
            _ = v18
        case "a":
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int32 = -2147483648
                _ = fi19
                _ = v20
            }
        case "hello":
            v21 := GenS1{F0: 4294967295, F1: -2147483648}
            _ = v21
        }
        _ = v6
    } else {
        switch "foo bar" {
        case "":
            v3 -= v3
        case "hello":
            v1 = (^(v5))
        case "a":
            v5 = (v1 + v5)
        default:
            v22 := GenS2{F0: 1}
            _ = v22
        }
        var v23 uint = 267
        if (!((v2 + v2) >= "hello")) {
            var v24 uint = (^(v23))
            var v25 uint64 = 926
            _ = v24
            _ = v25
        } else {
            var v26 int8 = 100
            _ = v26
        }
        _ = v23
    }
    fmt.Println(v5, v5, v3)
    var v27 int = 1
    if false {
        v28 := GenS3{F0: -1, F1: (^(v5))}
        fmt.Println(v5, v3, v2)
        fmt.Println(v2, v27)
        _ = v28
    }
    if (551 >= 65535) {
        switch 0 {
        case -171:
            if false {
                for fi29 := 0; fi29 < 6; fi29++ {
                    v30 := GenS4{F0: (v3 * 74), F1: (v2 + (v2 + "\n\t")), F2: 1}
                    if true {
                        if (0 == 42) {
                            var v31 int32 = -270
                            _ = v31
                        }
                        v32 := GenS4{F0: 0, F1: v2, F2: 1}
                        _ = v32
                    }
                    v3 -= v3
                    _ = fi29
                    _ = v30
                }
            }
        case 42:
            v5 += -901
        case -1:
            var v33 uint32 = 1
            _ = v33
        default:
            v27 = (-(v1))
        }
    } else {
        for fi34 := 0; fi34 < 15; fi34++ {
            switch v2 {
            case "a":
                if true {
                    if true {
                        if (42 != 18446744073709551615) {
                            v3 = (-(v3))
                        } else {
                            v35 := GenS3{F0: 1, F1: fi34}
                            switch v1 {
                            case -623:
                                switch (v5 << ((65535) & 63)) {
                                case -78:
                                    var v36 float32 = 1e10
                                    _ = v36
                                default:
                                    switch fi34 {
                                    case -595:
                                        var v37 uint64 = 18446744073709551615
                                        _ = v37
                                    case -843:
                                        var v38 int64 = 42
                                        _ = v38
                                    case -392:
                                        v1 = fi34
                                    }
                                }
                            case 733:
                                v5 -= v5
                            case 32767:
                                var v39 float32 = (((1.5 / 0.1) - (3.14159 - -2.25)) - float32(1.5))
                                _ = v39
                            default:
                                v40 := GenS5{F0: 193, F1: 2, F2: -763}
                                _ = v40
                            }
                            _ = v35
                        }
                        var v41 uint = 566
                        v5 += v27
                        _ = v41
                    }
                } else {
                    v27 -= fi34
                    v3 = 42
                    v5 = v5
                }
            case "":
                switch "foo bar" {
                case "\n\t":
                    for fi42 := 0; fi42 < 8; fi42++ {
                        v43 := GenS4{F0: (v3 + 100), F1: v2, F2: 42}
                        v1 += v27
                        _ = fi42
                        _ = v43
                    }
                case "a":
                    v44 := v4["foo bar"]
                    _ = v44
                default:
                    _ = len(v4)
                }
            default:
                var v45 string = "foo bar"
                _ = v45
            }
            fmt.Println(fi34)
            if false {
                v46 := GenS2{F0: 18446744073709551615}
                if (((3.14159 <= ((float32(3.14159) / 0.1) / 1.5)) && (false || (!(1 <= 42)))) || false) {
                    v47 := GenS6{F0: -3}
                    _ = v47
                } else {
                    v2 = "a"
                    v5 = v5
                    fmt.Println(v1)
                }
                var v48 uint8 = 0
                _ = v46
                _ = v48
            }
            _ = fi34
        }
        v49 := []int{42}
        fmt.Println(v5, v27)
        _ = v49
    }
    fmt.Println(v27)
    var v50 int64 = -240
    var v51 int32 = 42
    if (97 == 281) {
        for fi52 := 0; fi52 < 20; fi52++ {
            v1 = 1
            _ = fi52
        }
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        v54 := GenS7{F0: v2, F1: (!true)}
        _ = fi53
        _ = v54
    }
    v55 := []int{}
    switch (v27 >> ((255) & 63)) {
    case 9223372036854775807:
        for fi56 := 0; fi56 < 1; fi56++ {
            for fi57 := 0; fi57 < 5; fi57++ {
                v50 = (v50 / ((42) | 1))
                v58 := GenS2{F0: 70}
                _ = fi57
                _ = v58
            }
            fmt.Println(v2, v27)
            if (!((!true) || (128 <= 127))) {
                if true {
                    v59 := GenS7{F0: v2, F1: (!true)}
                    delete(v4, v2)
                    _ = v59
                }
                v50 = (v50 * 1)
                var v60 uint = 42
                _ = v60
            }
            _ = fi56
        }
    case -2147483648:
        var v61 float32 = -2.25
        _ = v61
    case 2147483647:
        var v62 int64 = (v50 << ((0) & 63))
        _ = v62
    }
    var v63 bool = true
    var v64 int = (v27 / ((v27) | 1))
    var v65 bool = false
    if ((v65 && (v51 == (v51 | int32(v1)))) && (((^(v3)) < 100) || v63)) {
        v27 -= 367
    }
    if v63 {
        if true {
            switch v2 {
            case "":
                switch ("\n\t" + v2) {
                case "foo bar":
                    var v66 uint64 = 2
                    _ = v66
                case "\n\t":
                    if (18446744073709551615 <= 137) {
                        if true {
                            if len(v55) > 0 {
                                v55[0] = 32767
                            }
                        } else {
                            var v67 float64 = ((1e10 / 0.1) - (-2.25 * -2.25))
                            if (4294967295 > 42) {
                                var v68 uint64 = 42
                                fmt.Println(v2)
                                _ = v68
                            }
                            v50 = (-(v50))
                            _ = v67
                        }
                        if v63 {
                            fmt.Println(v5)
                        }
                        if v63 {
                            v65 = ((-(v51)) == v51)
                        }
                    } else {
                        v63 = v65
                        _ = len(v4)
                    }
                case "hello":
                    for fi69 := 0; fi69 < 17; fi69++ {
                        if len(v55) >= 2 {
                            v55 = v55[1:len(v55)]
                        }
                        v3 = 127
                        var v70 int32 = v51
                        _ = fi69
                        _ = v70
                    }
                }
            case "a":
                var v71 int32 = (v51 | -2147483648)
                _ = v71
            case "\n\t":
                if false {
                    v72 := GenS8{F0: -32768}
                    v63 = ((!((v2 + v2) == v2)) && ((-(v51)) >= v51))
                    _ = v72
                } else {
                    for fi73 := 0; fi73 < 16; fi73++ {
                        if false {
                            var v74 float32 = (1e10 * 1.5)
                            switch fi73 {
                            case -577:
                                var v75 bool = (-1 <= -1)
                                _ = v75
                            default:
                                var v76 string = (v2 + (v2 + ""))
                                _ = v76
                            }
                            for fi77 := 0; fi77 < 17; fi77++ {
                                if (!(true && (false && v65))) {
                                    fmt.Println(v5)
                                } else {
                                    if (!(2147483647 >= 631)) {
                                        if (v64 != (^(v27))) {
                                            v50 = v50
                                            if v65 {
                                                var v78 uint8 = 1
                                                var v79 float64 = (0.1 + 1.5)
                                                if (2 > v78) {
                                                    fmt.Println(v79, fi77)
                                                    fmt.Println(v79, v27)
                                                    if (1000 >= int16(v78)) {
                                                        var v80 int = (^(fi77))
                                                        _ = v80
                                                    }
                                                } else {
                                                    var v81 int64 = 260
                                                    _ = v81
                                                }
                                                if (!((20 != 255) || false)) {
                                                    v2 = "a"
                                                    var v82 int = (v27 | (v64 & fi77))
                                                    _ = v82
                                                } else {
                                                    var v83 string = "a"
                                                    _ = v83
                                                }
                                                v84 := GenS8{F0: -596}
                                                _ = v78
                                                _ = v79
                                                _ = v84
                                            }
                                            if (255 == 65535) {
                                                var v85 float64 = ((1.5 * 0.0) * (3.14159 / ((1e10 - (1e10 * 0.0)) + float64(v74))))
                                                _ = v85
                                            }
                                        } else {
                                            v86 := GenS9{F0: 1, F1: ""}
                                            v64 -= fi73
                                            _ = v86
                                        }
                                        var v87 int32 = (v51 - (-(v51)))
                                        if len(v55) > 0 {
                                            v55[0] = v1
                                        }
                                        _ = v87
                                    }
                                    switch v64 {
                                    case 4294967295:
                                        v55 = append(v55, (v64 & (^(fi73))))
                                    }
                                    v88 := GenS2{F0: 0}
                                    v51 = 1
                                    _ = v88
                                }
                                var v89 int = 32767
                                var v90 uint8 = 179
                                _ = fi77
                                _ = v89
                                _ = v90
                            }
                            _ = v74
                        } else {
                            switch v2 {
                            case "\n\t":
                                _ = len(v4)
                            case "hello":
                                switch (("a" + v2) + v2) {
                                case "a":
                                    _ = len(v55)
                                }
                            }
                            var v91 uint64 = 18446744073709551615
                            _ = v91
                        }
                        _ = fi73
                    }
                    for fi92 := 0; fi92 < 14; fi92++ {
                        v5 = (v1 * 32767)
                        v65 = v63
                        switch 9223372036854775807 {
                        case 2147483647:
                            v93 := GenS10{F0: 2147483647}
                            _ = v93
                        }
                        _ = fi92
                    }
                    var v94 uint32 = 4294967295
                    _ = v94
                }
            default:
                var v95 int64 = 1
                _ = v95
            }
            v5 = (v64 & 0)
            var v96 uint64 = 4294967295
            _ = v96
        }
        for fi97 := 0; fi97 < 19; fi97++ {
            if v65 {
                var v98 float32 = 1e10
                var v99 uint8 = 127
                _ = v98
                _ = v99
            }
            if (!((v51 >= (v51 / (((v51 / (((v51 * v51)) | 1))) | 1))) && v63)) {
                v3 = v3
                v64 = -368
                if (42 != 1) {
                    fmt.Println(v5, v5)
                    var v100 int32 = (v51 >> ((0) & 63))
                    _ = v100
                }
            } else {
                var v101 float64 = 1e10
                v5 -= v1
                switch "a" {
                case "":
                    switch 32768 {
                    case 128:
                        switch v2 {
                        case "foo bar":
                            v1 += v5
                        default:
                            v102 := GenS1{F0: 4294967295, F1: -1}
                            _ = v102
                        }
                    case 42:
                        var v103 uint8 = 2
                        _ = v103
                    case 1:
                        var v104 uint8 = 2
                        _ = v104
                    }
                case "a":
                    v51 += v51
                }
                _ = v101
            }
            _ = fi97
        }
        v105 := GenS2{F0: 18446744073709551615}
        _ = v105
    } else {
        var v106 int8 = v3
        v107 := GenS1{F0: 1, F1: v51}
        var v108 int32 = (v51 * v51)
        _ = v106
        _ = v107
        _ = v108
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v55
    _ = v63
    _ = v64
    _ = v65
}
