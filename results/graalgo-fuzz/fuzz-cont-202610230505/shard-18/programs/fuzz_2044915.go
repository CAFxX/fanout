package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 int8; F1 uint32; F2 int64 }
type GenS10 struct { F0 uint64 }

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 1
    var v3 int8 = -100
    v3 += v3
    v4 := GenS1{F0: 0, F1: v2}
    v2 = 42
    var v5 int64 = int64(v1)
    v1 = (v1 / (((v1 | (v1 / ((255) | 1)))) | 1))
    v2 -= (^(v2))
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v2, v1)
        var v7 int32 = -2147483648
        fmt.Println(fi6, v2)
        _ = fi6
        _ = v7
    }
    var v8 string = "a"
    var v9 bool = (4294967295 <= 4294967295)
    var v10 uint64 = (v2 >> ((650) & 63))
    v11 := []int{4294967295, 443, 19, -246}
    fmt.Println(v9, v9)
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(v3)
        var v13 bool = (552 < -2147483648)
        var v14 int64 = (v5 | v5)
        _ = fi12
        _ = v13
        _ = v14
    }
    v5 = v5
    if (true || (v9 && v9)) {
        v15 := GenS1{F0: 2, F1: (v2 | (v2 & (^(v10))))}
        for fi16 := 0; fi16 < 7; fi16++ {
            v2 = (v10 % (((^(v2))) | 1))
            v3 = v3
            for fi17 := 0; fi17 < 7; fi17++ {
                if (1000 > 42) {
                    if (!v9) {
                        v10 -= v10
                    }
                    var v18 int64 = 333
                    var v19 uint64 = (v2 >> ((v1) & 63))
                    _ = v18
                    _ = v19
                }
                fmt.Println(v1)
                v20 := GenS1{F0: 0, F1: v10}
                if v9 {
                    _ = len(v11)
                    v11 = append(v11, (fi17 ^ 225))
                    var v21 int64 = -847
                    _ = v21
                } else {
                    switch "" {
                    case "hello":
                        v22 := GenS2{F0: (^(v1)), F1: v1}
                        _ = v22
                    case "foo bar":
                        v23 := GenS2{F0: 2, F1: (v1 | v1)}
                        _ = v23
                    case "\n\t":
                        if ((!(v5 > (v5 % (((v5 % (((v5 & v5)) | 1))) | 1)))) || false) {
                            var v24 float32 = -2.25
                            _ = v24
                        } else {
                            v2 = (v2 & uint64(v1))
                        }
                    }
                }
                _ = fi17
                _ = v20
            }
            _ = fi16
        }
        v11 = append(v11, 42)
        _ = v15
    }
    var v25 string = v8
    if v9 {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        if false {
            switch ("foo bar" + "") {
            case "foo bar":
                fmt.Println(v5)
            default:
                if (!(false && v9)) {
                    if v9 {
                        v2 = 18446744073709551615
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            for fi27 := 0; fi27 < 5; fi27++ {
                                v28 := GenS2{F0: (v1 ^ v1), F1: 65535}
                                switch fi26 {
                                case 127:
                                    if v9 {
                                        switch "hello" {
                                        case "hello":
                                            v25 = v25
                                        case "\n\t":
                                            v9 = (!true)
                                        case "":
                                            switch 32768 {
                                            case 0:
                                                switch v8 {
                                                case "foo bar":
                                                    v3 -= v3
                                                case "\n\t":
                                                    var v29 uint8 = 1
                                                    _ = v29
                                                case "":
                                                    v8 = (v8 + v8)
                                                }
                                            case 604:
                                                switch fi26 {
                                                case -1:
                                                    fmt.Println(fi27)
                                                case 2147483647:
                                                    v8 = v25
                                                default:
                                                    fmt.Println(v9, v5)
                                                }
                                            case -9223372036854775808:
                                                fmt.Println(v2)
                                            default:
                                                var v30 uint32 = 0
                                                _ = v30
                                            }
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                        default:
                                            var v31 int8 = (^(v3))
                                            _ = v31
                                        }
                                        switch (v25 + v25) {
                                        case "a":
                                            if (true || (-69 != -112)) {
                                                var v32 uint64 = 967
                                                _ = v32
                                            } else {
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                if true {
                                                    v33 := GenS1{F0: 1, F1: v10}
                                                    _ = v33
                                                } else {
                                                    v1 = (v1 & (v1 << ((v1) & 63)))
                                                    switch 127 {
                                                    case 42:
                                                        if v9 {
                                                            var v34 uint8 = 1
                                                            if len(v11) >= 2 {
                                                                v11 = v11[1:len(v11)]
                                                            }
                                                            _ = v34
                                                        } else {
                                                            v35 := GenS2{F0: v1, F1: (v1 << (((v1 / ((v1) | 1))) & 63))}
                                                            v5 += (-(v5))
                                                            _ = v35
                                                        }
                                                    default:
                                                        fmt.Println(v9, v3, v1)
                                                    }
                                                }
                                                v36 := GenS1{F0: 127, F1: v2}
                                                _ = v36
                                            }
                                        case "":
                                            var v37 bool = v9
                                            _ = v37
                                        case "\n\t":
                                            if false {
                                                var v38 int = 241
                                                if (!v9) {
                                                    var v39 int8 = (v3 << (((v1 & 2)) & 63))
                                                    var v40 uint = (v1 + (v1 + (v1 / ((968) | 1))))
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v41 := GenS3{F0: 1, F1: fi27}
                                                    _ = v41
                                                }
                                                var v42 uint16 = 32767
                                                _ = v38
                                                _ = v42
                                            } else {
                                                var v43 int32 = -2147483648
                                                _ = len(v11)
                                                _ = v43
                                            }
                                        }
                                        v5 = (v5 ^ v5)
                                    }
                                case 503:
                                    if len(v11) > 0 {
                                        v11[0] = fi27
                                    }
                                default:
                                    v2 += v10
                                }
                                _ = fi27
                                _ = v28
                            }
                        case "a":
                            fmt.Println(v25)
                        }
                    }
                    v10 = (v2 << ((v1) & 63))
                    v9 = false
                }
            }
            if v9 {
                for fi44 := 0; fi44 < 20; fi44++ {
                    if (v9 || v9) {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        if (0.0 <= (float64(0.1) / 3.14159)) {
                            var v45 int = 32768
                            v2 = v10
                            switch 4294967295 {
                            case 32767:
                                v46 := GenS4{F0: v25}
                                _ = v46
                            case 32768:
                                if len(v11) > 0 {
                                    v11[0] = fi44
                                }
                            case 218:
                                v47 := GenS1{F0: 0, F1: (v10 ^ 18446744073709551615)}
                                v9 = (0 != 32767)
                                var v48 uint64 = 121
                                _ = v47
                                _ = v48
                            }
                            _ = v45
                        } else {
                            if len(v11) > 0 {
                                v11[0] = fi44
                            }
                            switch v8 {
                            case "hello":
                                var v49 uint = (^(v1))
                                _ = v49
                            case "":
                                v8 = (("hello" + v8) + ((("" + "\n\t") + v25) + ("a" + "\n\t")))
                                v10 = (v10 | v2)
                            case "foo bar":
                                var v50 int8 = 0
                                _ = v50
                            }
                            v11 = append(v11, (fi44 * (fi26 ^ 42)))
                            if (!v9) {
                                v51 := GenS2{F0: 255, F1: (v1 | (^(v1)))}
                                _ = v51
                            } else {
                                if v9 {
                                    var v52 uint16 = 1
                                    v53 := GenS5{F0: 752, F1: uint16(v2)}
                                    _ = v52
                                    _ = v53
                                }
                                v11 = append(v11, (fi44 % ((-259) | 1)))
                                v5 += v5
                            }
                            v11 = append(v11, fi44)
                            var v54 float32 = (-2.25 - (3.14159 / 1.5))
                            _ = v54
                        }
                    }
                    _ = fi44
                }
                if v9 {
                    var v55 uint = v1
                    v11 = append(v11, (-(fi26)))
                    v56 := GenS2{F0: 576, F1: (v1 - v55)}
                    _ = v55
                    _ = v56
                } else {
                    if v9 {
                        switch (v25 + v8) {
                        case "foo bar":
                            for fi57 := 0; fi57 < 19; fi57++ {
                                if v9 {
                                    var v58 uint64 = (v2 % ((v10) | 1))
                                    _ = v58
                                } else {
                                    if (fi57 != (fi26 ^ fi26)) {
                                        if (898 == 469) {
                                            if len(v11) > 0 {
                                                v11[0] = (fi57 % ((fi26) | 1))
                                            }
                                            if (true && v9) {
                                                v59 := GenS6{F0: (^(fi57))}
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                if false {
                                                    v60 := GenS7{F0: fi26, F1: uint8(v2)}
                                                    _ = v60
                                                }
                                                _ = v59
                                            }
                                        } else {
                                            v1 -= (v1 * 1)
                                            fmt.Println(v5, v25, v3)
                                            switch fi26 {
                                            case -998:
                                                fmt.Println(v5)
                                            }
                                        }
                                        var v61 uint32 = 2
                                        _ = v61
                                    } else {
                                        v62 := GenS1{F0: 1, F1: v2}
                                        if false {
                                            var v63 int8 = (v3 & (-(v3)))
                                            _ = v63
                                        }
                                        _ = v62
                                    }
                                    fmt.Println(fi57, v2, v25)
                                }
                                var v64 int64 = 1
                                v5 += v5
                                _ = fi57
                                _ = v64
                            }
                        default:
                            for fi65 := 0; fi65 < 11; fi65++ {
                                if v9 {
                                    var v66 int16 = 1
                                    var v67 uint32 = 2147483647
                                    if v9 {
                                        v66 = (v66 % ((v66) | 1))
                                        v66 += (v66 % ((-32768) | 1))
                                        var v68 uint16 = 72
                                        switch v25 {
                                        case "\n\t":
                                            var v69 uint = 694
                                            _ = v69
                                        }
                                        if false {
                                            if len(v11) > 0 {
                                                v11[0] = -1
                                            }
                                            fmt.Println(v10, v2)
                                            var v70 float64 = (0.0 - -2.25)
                                            _ = v70
                                        } else {
                                            v67 = 329
                                            if ((uint64(v68) < (v10 >> ((1) & 63))) && v9) {
                                                v11 = append(v11, fi26)
                                                v71 := GenS7{F0: (fi26 << (((^(v1))) & 63)), F1: 29}
                                                v68 += v68
                                                _ = v71
                                            } else {
                                                fmt.Println(fi65, fi26, v10)
                                                v11 = append(v11, fi65)
                                            }
                                        }
                                        _ = len(v11)
                                        _ = v68
                                    }
                                    _ = v66
                                    _ = v67
                                }
                                if (2 > v10) {
                                    v72 := GenS8{F0: 960, F1: (fi65 | 42)}
                                    var v73 int = (fi26 >> ((v1) & 63))
                                    _ = v72
                                    _ = v73
                                } else {
                                    var v74 int8 = int8(v5)
                                    _ = v74
                                }
                                v75 := GenS1{F0: 131, F1: v10}
                                v76 := GenS5{F0: -1, F1: 0}
                                _ = fi65
                                _ = v75
                                _ = v76
                            }
                        }
                        switch (v8 + ("hello" + ((v25 + v25) + v8))) {
                        case "\n\t":
                            var v77 bool = false
                            _ = v77
                        default:
                            for fi78 := 0; fi78 < 6; fi78++ {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                var v79 int32 = 2147483647
                                v79 += int32(v3)
                                _ = fi78
                                _ = v79
                            }
                        }
                    }
                    v11 = append(v11, 889)
                    var v80 uint8 = 255
                    _ = v80
                }
                var v81 int8 = v3
                _ = v81
            }
        } else {
            if true {
                if (304 >= -2147483648) {
                    switch (v8 + "a") {
                    case "foo bar":
                        v82 := GenS2{F0: v1, F1: 146}
                        _ = v82
                    case "a":
                        var v83 float64 = 3.14159
                        _ = v83
                    case "":
                        for fi84 := 0; fi84 < 6; fi84++ {
                            v11 = append(v11, fi84)
                            fmt.Println(v1)
                            _ = fi84
                        }
                    }
                    if v9 {
                        if (!(!((v3 | (^(v3))) == (v3 >> ((v1) & 63))))) {
                            for fi85 := 0; fi85 < 16; fi85++ {
                                if true {
                                    switch v25 {
                                    case "foo bar":
                                        switch (((("" + "foo bar") + (v25 + "hello")) + v25) + (v25 + "a")) {
                                        case "a":
                                            fmt.Println(v1)
                                        case "":
                                            if true {
                                                var v86 int32 = 42
                                                var v87 uint = v1
                                                var v88 uint64 = v10
                                                _ = v86
                                                _ = v87
                                                _ = v88
                                            }
                                        case "\n\t":
                                            switch (fi26 % (((fi26 | (fi85 % ((int(v3)) | 1)))) | 1)) {
                                            case 616:
                                                v10 += v10
                                                fmt.Println(v10, v5, v8)
                                                var v89 int64 = (v5 | v5)
                                                _ = v89
                                            }
                                        }
                                    case "hello":
                                        v25 = v25
                                    }
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    v90 := GenS9{F0: v3, F1: 919, F2: v5}
                                    _ = v90
                                }
                                v91 := GenS8{F0: 0, F1: (fi26 % ((fi85) | 1))}
                                var v92 int64 = v5
                                _ = fi85
                                _ = v91
                                _ = v92
                            }
                        } else {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            var v93 float64 = -2.25
                            _ = v93
                        }
                        v9 = (v9 || true)
                    }
                    fmt.Println(v25, v25)
                } else {
                    v94 := GenS8{F0: 42, F1: (fi26 / (((fi26 >> ((1) & 63))) | 1))}
                    for fi95 := 0; fi95 < 4; fi95++ {
                        switch fi95 {
                        case 0:
                            fmt.Println(v3, v2)
                        case -940:
                            v2 = (v2 * v2)
                        }
                        fmt.Println(v5, v3)
                        _ = fi95
                    }
                    var v96 string = ("foo bar" + "hello")
                    _ = v94
                    _ = v96
                }
                for fi97 := 0; fi97 < 1; fi97++ {
                    switch (fi26 | (fi26 ^ fi26)) {
                    case -696:
                        v98 := GenS10{F0: 42}
                        _ = v98
                    case -2147483648:
                        v25 = ("a" + v8)
                    case 786:
                        if len(v11) > 0 {
                            v11[0] = (fi97 ^ fi97)
                        }
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    var v99 bool = v9
                    _ = fi97
                    _ = v99
                }
                var v100 uint64 = 18446744073709551615
                _ = v100
            }
        }
        var v101 bool = true
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi26
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
