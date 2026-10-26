package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint16; F2 uint }
type GenS6 struct { F0 int8; F1 bool }
func genFn0(p0 uint64, p1 uint) uint16 {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int8 = 127
        var v6 int32 = 1
        fmt.Println(v5)
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        switch fi7 {
        case 587:
            v8 := []int{fi7, 624, fi7}
            _ = v8
        case 32768:
            v9 := GenS1{F0: 42}
            _ = v9
        case 0:
            var v10 int = fi7
            _ = v10
        default:
            if (((!true) && true) && false) {
                p0 = (p0 << (((p1 | p1)) & 63))
                switch ((("foo bar" + "foo bar") + ("" + ((("" + "") + ("" + "hello")) + "a"))) + "") {
                case "":
                    if (true || (int8(p1) >= 0)) {
                        var v11 uint32 = 1
                        if true {
                            var v12 float32 = 3.14159
                            for fi13 := 0; fi13 < 5; fi13++ {
                                switch (fi13 & (fi13 >> ((p1) & 63))) {
                                case 42:
                                    var v14 bool = true
                                    _ = v14
                                }
                                _ = fi13
                            }
                            _ = v12
                        } else {
                            v11 += v11
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                switch fi7 {
                                case 0:
                                    if false {
                                        switch -1 {
                                        case 1:
                                            if (false || (1 <= p1)) {
                                                v15 := GenS1{F0: -1}
                                                p1 = p1
                                                v16 := []int{42, 32767}
                                                _ = v15
                                                _ = v16
                                            } else {
                                                fmt.Println(fi7)
                                                fmt.Println(fi7, p1, p1)
                                            }
                                        case 9223372036854775807:
                                            switch (fi7 & (fi7 ^ 1)) {
                                            case -2147483648:
                                                p0 = p0
                                            case 42:
                                                if (0 > 32767) {
                                                    if (0 < 127) {
                                                        v17 := GenS2{F0: 337, F1: 412, F2: 79}
                                                        var v18 float64 = 0.1
                                                        var v19 uint8 = uint8(p1)
                                                        _ = v17
                                                        _ = v18
                                                        _ = v19
                                                    } else {
                                                        p1 = 378
                                                        p0 += 0
                                                        v11 = 800
                                                    }
                                                }
                                            default:
                                                switch ("\n\t" + (("" + (("a" + "hello") + ("a" + "foo bar"))) + "a")) {
                                                case "a":
                                                    var v20 uint16 = 255
                                                    _ = v20
                                                default:
                                                    v21 := []int{fi7, fi7}
                                                    _ = v21
                                                }
                                            }
                                        case 722:
                                            v11 = (v11 << ((p1) & 63))
                                        }
                                        v22 := []int{(^(fi7)), fi7}
                                        fmt.Println(p1, p0)
                                        _ = v22
                                    } else {
                                        for fi23 := 0; fi23 < 16; fi23++ {
                                            var v24 int8 = -127
                                            var v25 uint = (p1 / ((p1) | 1))
                                            v26 := GenS1{F0: int8(fi7)}
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                        v27 := GenS2{F0: -800, F1: 0, F2: 42}
                                        _ = v27
                                    }
                                case 128:
                                    switch "hello" {
                                    case "":
                                        switch "\n\t" {
                                        case "foo bar":
                                            for fi28 := 0; fi28 < 17; fi28++ {
                                                v29 := GenS1{F0: 127}
                                                v11 += (v11 & v11)
                                                _ = fi28
                                                _ = v29
                                            }
                                        default:
                                            v30 := []int{-767, int(v11), (-(fi7)), 1}
                                            _ = v30
                                        }
                                    case "hello":
                                        for fi31 := 0; fi31 < 15; fi31++ {
                                            p1 = (p1 >> (((^(p1))) & 63))
                                            _ = fi31
                                        }
                                    default:
                                        v32 := GenS2{F0: 1, F1: -1, F2: -9223372036854775808}
                                        _ = v32
                                    }
                                case -9223372036854775808:
                                    fmt.Println(fi7, fi7)
                                }
                            case "hello":
                                v33 := GenS2{F0: -277, F1: 42, F2: 1}
                                _ = v33
                            }
                        }
                        for fi34 := 0; fi34 < 4; fi34++ {
                            v35 := GenS2{F0: -280, F1: -1, F2: -444}
                            v11 = 2
                            var v36 uint8 = 10
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        _ = v11
                    }
                case "a":
                    p1 = p1
                default:
                    v37 := []int{fi7}
                    _ = v37
                }
                var v38 int8 = -128
                _ = v38
            } else {
                fmt.Println(fi7, p1, p1)
            }
        }
        v39 := GenS1{F0: 18}
        _ = fi7
        _ = v39
    }
    _ = p0
    _ = p1
    return 676
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("hello" + "a")
    var v3 int16 = -32768
    v2 = v2
    if false {
        for fi40 := 0; fi40 < 14; fi40++ {
            v41 := GenS3{F0: -2147483648}
            _ = fi40
            _ = v41
        }
        var v42 int32 = 42
        if (true || true) {
            v43 := GenS1{F0: 115}
            var v44 int = 44
            v44 += v44
            _ = v43
            _ = v44
        }
        _ = v42
    } else {
        if false {
            v3 = (v3 * (v3 - v3))
        }
        var v45 uint8 = 127
        var v46 uint64 = 4294967295
        _ = v45
        _ = v46
    }
    v47 := GenS4{F0: (18446744073709551615 < 1)}
    switch v2 {
    case "":
        if false {
            for fi48 := 0; fi48 < 18; fi48++ {
                v49 := []int{327, fi48}
                var v50 float32 = 0.1
                _ = fi48
                _ = v49
                _ = v50
            }
            var v51 float32 = float32(float64(0.0))
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 int8 = 127
                _ = fi52
                _ = v53
            }
            _ = v51
        }
    default:
        if true {
            v2 = "a"
            v54 := GenS3{F0: 641}
            _ = v54
        }
    }
    if true {
        var v55 int8 = 42
        var v56 int64 = 0
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        var v58 int = fi57
        _ = fi57
        _ = v58
    }
    switch 32768 {
    case -157:
        if (159 >= 30) {
            switch ((v2 + ((v2 + v2) + v2)) + (v2 + "foo bar")) {
            case "\n\t":
                v59 := []int{-899}
                _ = v59
            case "foo bar":
                v3 += (v3 % ((v3) | 1))
            default:
                fmt.Println(v2, v2)
            }
        }
    case -708:
        for fi60 := 0; fi60 < 20; fi60++ {
            v3 += -32768
            switch ("a" + "hello") {
            case "":
                var v61 uint16 = 255
                _ = v61
            case "foo bar":
                var v62 uint16 = 65535
                _ = v62
            case "hello":
                for fi63 := 0; fi63 < 2; fi63++ {
                    var v64 bool = (-2.25 != 1.5)
                    var v65 int = (fi63 - fi63)
                    v65 -= (fi60 << ((65535) & 63))
                    v1[v65] = 2147483648
                    v64 = (!true)
                    if v64 {
                        v3 = (v3 | (-(v3)))
                        v66 := GenS5{F0: (fi63 & fi60), F1: 700, F2: 1}
                        _ = v66
                    } else {
                        var v67 int8 = 46
                        var v68 string = "foo bar"
                        if v64 {
                            switch (v65 | 1) {
                            case 9223372036854775807:
                                v64 = v64
                            case -570:
                                var v69 int64 = -9223372036854775808
                                _ = v69
                            default:
                                var v70 uint8 = 2
                                _ = v70
                            }
                            var v71 uint = 65535
                            v67 += (v67 % ((0) | 1))
                            if (!(!v64)) {
                                v72 := GenS6{F0: v67, F1: v64}
                                v73 := GenS5{F0: (v65 / ((fi60) | 1)), F1: uint16(v71), F2: 2}
                                _ = v72
                                _ = v73
                            }
                            _ = v71
                        }
                        v68 = v2
                        _ = v67
                        _ = v68
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            default:
                switch v2 {
                case "foo bar":
                    v74 := []int{int(v3), fi60, -2147483648, (fi60 & fi60)}
                    _ = v74
                case "":
                    var v75 uint32 = 0
                    _ = v75
                case "hello":
                    switch (fi60 ^ (fi60 >> ((255) & 63))) {
                    case -2147483648:
                        var v76 string = ("" + (v2 + v2))
                        _ = v76
                    case 787:
                        fmt.Println(fi60, v2)
                    default:
                        _ = len(v1)
                    }
                }
            }
            var v77 uint8 = 0
            _ = fi60
            _ = v77
        }
    case 4294967295:
        for fi78 := 0; fi78 < 14; fi78++ {
            v79 := []int{(fi78 / (((fi78 - fi78)) | 1))}
            var v80 uint = 4294967295
            switch fi78 {
            case 820:
                switch (fi78 & (fi78 | fi78)) {
                case 4294967295:
                    if ((v2 == v2) || true) {
                        var v81 string = v2
                        var v82 int64 = 9223372036854775807
                        _ = v81
                        _ = v82
                    } else {
                        v80 -= v80
                        for fi83 := 0; fi83 < 8; fi83++ {
                            switch (fi83 << ((v80) & 63)) {
                            case -2147483648:
                                v80 += (^(v80))
                            case 128:
                                var v84 uint8 = 255
                                _ = v84
                            default:
                                var v85 float32 = 1.5
                                _ = v85
                            }
                            _ = fi83
                        }
                    }
                }
            }
            _ = fi78
            _ = v79
            _ = v80
        }
    default:
        if (2147483647 < 1) {
            switch "a" {
            case "":
                v86 := GenS6{F0: 1, F1: false}
                _ = v86
            case "a":
                v87 := GenS2{F0: 42, F1: 42, F2: 42}
                _ = v87
            default:
                v88 := v1[-723]
                _ = v88
            }
        } else {
            v89 := GenS1{F0: -41}
            if (v2 != "\n\t") {
                v90 := GenS4{F0: (("a" + v2) == "a")}
                _ = v90
            }
            var v91 uint = 378
            _ = v89
            _ = v91
        }
    }
    if true {
        var v92 float32 = 0.1
        _ = v92
    } else {
        fmt.Println(v2, v2)
        for fi93 := 0; fi93 < 9; fi93++ {
            for fi94 := 0; fi94 < 19; fi94++ {
                fmt.Println(v3)
                _ = fi94
            }
            for fi95 := 0; fi95 < 1; fi95++ {
                var v96 string = v2
                switch 2147483647 {
                case 9223372036854775807:
                    var v97 int8 = int8(fi95)
                    _ = v97
                case 456:
                    _ = len(v1)
                }
                v98 := GenS6{F0: -1, F1: (!(true || true))}
                v3 = -32768
                switch fi93 {
                case 0:
                    fmt.Println(v2, v96)
                case -2147483648:
                    v99 := GenS1{F0: 42}
                    _ = v99
                default:
                    v96 = (v96 + ("\n\t" + "foo bar"))
                }
                _ = fi95
                _ = v96
                _ = v98
            }
            var v100 int64 = 654
            _ = fi93
            _ = v100
        }
        fmt.Println(v2, v2)
    }
    if (!(int(v3) < 679)) {
        v101 := GenS2{F0: -829, F1: -560, F2: -1}
        for fi102 := 0; fi102 < 8; fi102++ {
            fmt.Println(v2, v3)
            v2 = v2
            v1[fi102] = 374
            _ = fi102
        }
        var v103 int64 = -1
        _ = v101
        _ = v103
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
}
