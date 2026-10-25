package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8; F1 uint }
type GenS12 struct { F0 uint32; F1 uint8; F2 bool }
type GenS13 struct { F0 int32; F1 uint8; F2 int8 }
type GenS14 struct { F0 uint16; F1 uint16 }
type GenS15 struct { F0 uint8; F1 uint; F2 uint }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint; F1 int64 }
func genFn0(p0 int64, p1 int64) uint16 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{32767}
        _ = fi7
        _ = v8
    }
    p0 -= -9223372036854775808
    v9 := GenS1{F0: 2}
    var v10 float64 = 0.0
    for fi11 := 0; fi11 < 3; fi11++ {
        if (true || ((1.5 >= v10) && (!(-790 < (fi11 / ((965) | 1)))))) {
            switch (fi11 + fi11) {
            case 999:
                v10 += -2.25
            }
        } else {
            p1 = p1
            var v12 float64 = v10
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 int16 = 42
    var v2 uint = 65535
    var v3 int64 = 0
    var v4 uint32 = 0
    var v5 int32 = int32(v2)
    v6 := make(map[int]bool)
    v1 = (^(v1))
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch ("foo bar" + (("hello" + (("" + "hello") + "")) + ("" + ""))) {
            case "a":
                for fi14 := 0; fi14 < 8; fi14++ {
                    v4 = 817
                    _ = fi14
                }
            case "foo bar":
                switch ("\n\t" + "hello") {
                case "foo bar":
                    v15 := GenS1{F0: 2}
                    _ = v15
                default:
                    var v16 uint32 = (v4 ^ v4)
                    _ = v16
                }
            }
            if false {
                switch fi13 {
                case -1:
                    v2 = v2
                case 32768:
                    v17 := []int{-746, 206, 9223372036854775807, (fi13 >> (((v2 << (((v2 << ((v2) & 63))) & 63))) & 63))}
                    _ = v17
                case 128:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 uint64 = 81
                        switch 2147483647 {
                        case 9223372036854775807:
                            v20 := GenS2{F0: v19}
                            _ = v20
                        case -729:
                            switch ("\n\t" + ("\n\t" + "hello")) {
                            case "":
                                if true {
                                    if false {
                                        v21 := []int{}
                                        _ = v21
                                    } else {
                                        v3 += v3
                                    }
                                    if true {
                                        v22 := GenS3{F0: "\n\t", F1: v19, F2: fi13}
                                        switch "" {
                                        case "hello":
                                            if false {
                                                if (false && (0.0 == 1e10)) {
                                                    v6[-1] = ((!(true || ((1 >= 9223372036854775807) && (int16(v4) <= -32768)))) || true)
                                                    v23 := GenS2{F0: (^(v19))}
                                                    if true {
                                                        v1 = -675
                                                    } else {
                                                        switch (("" + ("foo bar" + ("foo bar" + "\n\t"))) + "foo bar") {
                                                        case "\n\t":
                                                            _ = len(v6)
                                                        case "a":
                                                            v24 := GenS3{F0: "hello", F1: 0, F2: (fi18 & (-(fi13)))}
                                                            _ = v24
                                                        case "hello":
                                                            var v25 int32 = (v5 << ((42) & 63))
                                                            _ = v25
                                                        default:
                                                            v26 := GenS3{F0: (("" + "foo bar") + "\n\t"), F1: v19, F2: 358}
                                                            _ = v26
                                                        }
                                                    }
                                                    _ = v23
                                                } else {
                                                    v3 = v3
                                                }
                                                v27 := GenS2{F0: uint64(v2)}
                                                var v28 bool = false
                                                _ = v27
                                                _ = v28
                                            } else {
                                                v19 = 2
                                            }
                                        case "":
                                            if false {
                                                v29 := GenS4{F0: fi13, F1: 1}
                                                _ = v29
                                            } else {
                                                var v30 uint16 = 61
                                                v30 = v30
                                                v31 := GenS3{F0: ("hello" + ("" + "a")), F1: v19, F2: 9223372036854775807}
                                                switch (fi13 << ((v2) & 63)) {
                                                case 32767:
                                                    v5 = 312
                                                }
                                                v19 -= v19
                                                fmt.Println(v2, v5)
                                                if true {
                                                    v4 = (v4 & 4294967295)
                                                    v6[fi13] = true
                                                    v1 += (v1 & 32767)
                                                    v32 := GenS4{F0: 2147483647, F1: 127}
                                                    if true {
                                                        if (false || (float64(0.1) >= 0.0)) {
                                                            fmt.Println(fi18, fi18, v19)
                                                            fmt.Println(v3, v19)
                                                            v4 = 0
                                                        } else {
                                                            switch ("" + (("" + "a") + ("hello" + "foo bar"))) {
                                                            case "foo bar":
                                                                var v33 string = ""
                                                                _ = v33
                                                            case "":
                                                                fmt.Println(v2, v5)
                                                            case "a":
                                                                v34 := []int{fi13}
                                                                _ = v34
                                                            default:
                                                                switch fi18 {
                                                                case 0:
                                                                    v35 := []int{-2147483648, fi18}
                                                                    _ = v35
                                                                case -861:
                                                                    v36 := v6[2147483647]
                                                                    _ = v36
                                                                default:
                                                                    var v37 int32 = v5
                                                                    _ = v37
                                                                }
                                                                var v38 int16 = (v1 / (((v1 / ((int16(v4)) | 1))) | 1))
                                                                var v39 uint8 = 42
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            var v40 int16 = (-(v1))
                                                            var v41 uint8 = 0
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        v4 = uint32(v19)
                                                    }
                                                    _ = v32
                                                }
                                                v42 := GenS1{F0: 17}
                                                _ = v30
                                                _ = v31
                                                _ = v42
                                            }
                                        }
                                        if false {
                                            v43 := GenS5{F0: 42, F1: (v19 | v19)}
                                            v5 -= v5
                                            v44 := GenS6{F0: -763}
                                            _ = v43
                                            _ = v44
                                        }
                                        v45 := GenS3{F0: "a", F1: (^(v19)), F2: fi13}
                                        v46 := GenS7{F0: 9223372036854775807, F1: 32767}
                                        _ = v22
                                        _ = v45
                                        _ = v46
                                    }
                                    var v47 uint = v2
                                    _ = v47
                                } else {
                                    var v48 float64 = 0.0
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        var v49 uint8 = 42
                                        _ = v49
                                    case -2147483648:
                                        fmt.Println(fi18, v4, v19)
                                    }
                                    switch ("foo bar" + "a") {
                                    case "":
                                        fmt.Println(v4, v2, v2)
                                    case "\n\t":
                                        v1 = (v1 >> (((v2 & v2)) & 63))
                                    default:
                                        v50 := GenS8{F0: (-(fi18)), F1: v3}
                                        _ = v50
                                    }
                                    v51 := GenS2{F0: v19}
                                    switch "a" {
                                    case "hello":
                                        switch (-(fi18)) {
                                        case 2147483647:
                                            fmt.Println(v19, v2)
                                        }
                                    case "\n\t":
                                        _ = len(v6)
                                    default:
                                        v52 := GenS4{F0: (fi18 | fi13), F1: -128}
                                        _ = v52
                                    }
                                    if ((1.5 * 0.0) == 1.5) {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v53 := GenS9{F0: "foo bar", F1: 255}
                                            _ = v53
                                        case "hello":
                                            switch "" {
                                            case "":
                                                if ((v48 >= 0.1) || false) {
                                                    if false {
                                                        var v54 int32 = v5
                                                        _ = v54
                                                    } else {
                                                        v55 := []int{(fi18 >> (((v2 ^ v2)) & 63))}
                                                        v2 = (v2 - 4294967295)
                                                        _ = v55
                                                    }
                                                } else {
                                                    v2 += v2
                                                    switch "\n\t" {
                                                    case "":
                                                        var v56 float32 = (0.1 * (1e10 / 0.1))
                                                        _ = v56
                                                    }
                                                }
                                            case "\n\t":
                                                fmt.Println(v4, v2, v3)
                                            default:
                                                if (32767 >= 676) {
                                                    if true {
                                                        fmt.Println(v19, v48)
                                                        fmt.Println(fi13, fi18)
                                                        if true {
                                                            switch ("hello" + "") {
                                                            case "":
                                                                v57 := GenS7{F0: (fi13 + fi18), F1: 127}
                                                                _ = v57
                                                            }
                                                        } else {
                                                            v58 := []int{}
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                            var v59 int16 = v1
                                                            v60 := GenS2{F0: 2}
                                                            var v61 int64 = v3
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    } else {
                                                        var v62 uint16 = 167
                                                        if false {
                                                            v63 := GenS9{F0: "a", F1: 255}
                                                            _ = v63
                                                        }
                                                        v2 -= 255
                                                        _ = v62
                                                    }
                                                    if false {
                                                        v48 += float64((1.5 + 0.1))
                                                    }
                                                    delete(v6, (fi18 & -1))
                                                }
                                            }
                                        case "foo bar":
                                            switch fi18 {
                                            case -54:
                                                var v64 uint = (v2 | 65535)
                                                _ = v64
                                            case 42:
                                                v65 := GenS10{F0: 42}
                                                _ = v65
                                            case 905:
                                                v66 := GenS8{F0: fi18, F1: (^(v3))}
                                                _ = v66
                                            default:
                                                v6[128] = false
                                                var v67 string = ("foo bar" + ("\n\t" + "\n\t"))
                                                _ = v67
                                            }
                                        default:
                                            if true {
                                                v4 -= v4
                                                v68 := []int{int(v19)}
                                                var v69 int16 = v1
                                                switch "foo bar" {
                                                case "":
                                                    if false {
                                                        v70 := GenS11{F0: 1, F1: (v2 << (((v2 & v2)) & 63))}
                                                        _ = v70
                                                    } else {
                                                        var v71 float32 = (0.0 - 0.1)
                                                        _ = v71
                                                    }
                                                case "\n\t":
                                                    switch fi18 {
                                                    case -1:
                                                        if (1.5 != float32(v48)) {
                                                            v48 = v48
                                                            v3 = (^(v3))
                                                            if true {
                                                                v48 = float64((0.1 - 0.1))
                                                            } else {
                                                                if false {
                                                                    _ = len(v6)
                                                                    v6[42] = (-104 != 100)
                                                                }
                                                                delete(v6, (fi13 ^ fi18))
                                                                var v72 int32 = v5
                                                                _ = v72
                                                            }
                                                        }
                                                    case 32767:
                                                        v3 -= (v3 - int64(v2))
                                                    case 42:
                                                        v6[466] = false
                                                    default:
                                                        v19 += v19
                                                    }
                                                default:
                                                    switch (fi13 | fi18) {
                                                    case 0:
                                                        var v73 uint = v2
                                                        _ = v73
                                                    case -9223372036854775808:
                                                        v74 := GenS12{F0: (v4 << ((42) & 63)), F1: 2, F2: true}
                                                        _ = v74
                                                    default:
                                                        v75 := GenS4{F0: (fi13 / (((fi18 ^ fi13)) | 1)), F1: -119}
                                                        _ = v75
                                                    }
                                                }
                                                _ = v68
                                                _ = v69
                                            }
                                        }
                                    } else {
                                        if (!true) {
                                            fmt.Println(v3, v3, v5)
                                            v76 := v6[fi18]
                                            _ = v76
                                        } else {
                                            if false {
                                                v2 = v2
                                                if (v2 > (v2 >> ((v2) & 63))) {
                                                    fmt.Println(v4, v3, fi13)
                                                    fmt.Println(v48, v2, v1)
                                                    v77 := GenS13{F0: 968, F1: 43, F2: 127}
                                                    _ = v77
                                                }
                                                var v78 float32 = float32(1e10)
                                                _ = v78
                                            }
                                            v5 -= -76
                                            var v79 uint16 = genFn0(v3, v3)
                                            _ = v79
                                        }
                                        v80 := GenS4{F0: fi13, F1: 100}
                                        var v81 float32 = -2.25
                                        _ = v80
                                        _ = v81
                                    }
                                    fmt.Println(v48)
                                    _ = v48
                                    _ = v51
                                }
                            }
                        case -2147483648:
                            var v82 uint32 = v4
                            _ = v82
                        default:
                            v83 := GenS13{F0: (v5 - (v5 & (v5 << ((974) & 63)))), F1: 239, F2: 65}
                            _ = v83
                        }
                        _ = fi18
                        _ = v19
                    }
                default:
                    var v84 int64 = -913
                    _ = v84
                }
                fmt.Println(v3, v3, fi13)
                if true {
                    var v85 uint = (^(v2))
                    _ = v85
                } else {
                    v86 := []int{(^(fi13)), -397, fi13, fi13}
                    _ = v86
                }
            }
            _ = fi13
        }
        if (!((((("a" + ("" + "a")) + "hello") + ("" + ("" + ""))) == "hello") && true)) {
            v87 := GenS12{F0: v4, F1: 127, F2: (-2147483648 >= 488)}
            if true {
                for fi88 := 0; fi88 < 11; fi88++ {
                    var v89 bool = (v5 != -2147483648)
                    if v89 {
                        fmt.Println(v89, fi88)
                    }
                    var v90 float32 = ((3.14159 + 1e10) / (3.14159 / -2.25))
                    _ = fi88
                    _ = v89
                    _ = v90
                }
                if false {
                    fmt.Println(v2, v3)
                } else {
                    if false {
                        for fi91 := 0; fi91 < 10; fi91++ {
                            delete(v6, (fi91 ^ 9223372036854775807))
                            v92 := GenS14{F0: 292, F1: 32767}
                            _ = fi91
                            _ = v92
                        }
                        var v93 int16 = (-(v1))
                        _ = v93
                    } else {
                        var v94 string = ""
                        v94 = (("hello" + "\n\t") + v94)
                        if (false && false) {
                            v95 := []int{127, 42, -2147483648, 706}
                            _ = v95
                        } else {
                            switch -2147483648 {
                            case 2147483647:
                                v4 -= v4
                            default:
                                var v96 string = v94
                                _ = v96
                            }
                            var v97 uint16 = 65535
                            _ = v97
                        }
                        _ = v94
                    }
                }
                v1 += int16(v4)
            }
            _ = v87
        } else {
            switch "a" {
            case "hello":
                var v98 int64 = v3
                _ = v98
            }
        }
        fmt.Println(v3, v1, v5)
    }
    _ = len(v6)
    v99 := GenS15{F0: 1, F1: v2, F2: 42}
    if (100 <= 85) {
        v1 += (-(v1))
    } else {
        var v100 float64 = 0.0
        v5 += (v5 >> (((v2 % ((v2) | 1))) & 63))
        _ = v100
    }
    v2 = (^(v2))
    v3 += v3
    v101 := v6[-1]
    v102 := GenS5{F0: -92, F1: uint64(v4)}
    if v101 {
        if (1.5 != (3.14159 / 1.5)) {
            v103 := GenS16{F0: 2}
            _ = v103
        } else {
            fmt.Println(v2, v3, v4)
            var v104 uint64 = 805
            _ = v104
        }
        var v105 int64 = v3
        _ = v105
    }
    _ = len(v6)
    v106 := GenS17{F0: v2, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v101
    _ = v102
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
