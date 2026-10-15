package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int8; F1 uint64; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int; F1 uint8; F2 uint32 }
type GenS13 struct { F0 bool; F1 int16 }
type GenS14 struct { F0 uint64; F1 int; F2 uint }
type GenS15 struct { F0 int16 }
func genFn0(p0 uint) uint16 {
    var v5 uint = (p0 - 42)
    var v6 uint = (p0 >> (((p0 - (v5 + (p0 >> (((p0 * 4294967295)) & 63))))) & 63))
    var v7 int64 = 467
    v5 = 676
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int = 127
    var v2 uint = 18446744073709551615
    v3 := []uint32{}
    v4 := []int{(v1 | v1)}
    if false {
        v2 = v2
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        switch "a" {
        case "\n\t":
            var v9 int16 = 665
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v2)
            _ = fi10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            switch ("foo bar" + "hello") {
            case "hello":
                v1 += (^(fi8))
            case "a":
                v2 -= v2
            default:
                v3 = append(v3, 2147483648)
            }
            _ = fi11
        }
        _ = fi8
    }
    var v12 string = "\n\t"
    switch (((v12 + "\n\t") + ("foo bar" + v12)) + (v12 + v12)) {
    case "\n\t":
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(fi13, fi13, v2)
            v14 := GenS1{F0: (!true)}
            _ = fi13
            _ = v14
        }
    case "a":
        fmt.Println(v1)
    default:
        if true {
            var v15 uint64 = 2
            v16 := GenS1{F0: (!true)}
            _ = v15
            _ = v16
        }
    }
    _ = len(v4)
    switch (v1 - (-(v1))) {
    case -2147483648:
        if (v12 > "a") {
            switch v1 {
            case 1:
                v17 := GenS2{F0: 127}
                _ = v17
            case 296:
                v1 += (v1 & (v1 - (v1 >> ((v2) & 63))))
            case -9223372036854775808:
                var v18 uint32 = 2147483648
                _ = v18
            default:
                if ("" >= v12) {
                    var v19 bool = false
                    if (!v19) {
                        for fi20 := 0; fi20 < 13; fi20++ {
                            v2 -= (^(v2))
                            _ = fi20
                        }
                        var v21 uint16 = 0
                        _ = v21
                    }
                    _ = v19
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        fmt.Println(v12, fi22, fi22)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v1, v2)
                        _ = fi22
                    }
                    _ = len(v3)
                    if (float32(float64(1e10)) != float32(0.1)) {
                        if ("a" == (((v12 + "foo bar") + "hello") + (v12 + v12))) {
                            v2 += 0
                        }
                    } else {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            var v24 uint32 = 0
                            _ = fi23
                            _ = v24
                        }
                        v25 := GenS2{F0: 128}
                        _ = v25
                    }
                }
            }
        } else {
            var v26 int32 = -979
            _ = v26
        }
    case 127:
        var v27 float64 = 0.0
        _ = v27
    case 4294967295:
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 uint8 = 127
                if false {
                    v31 := GenS1{F0: false}
                    v12 = "foo bar"
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 uint = v2
                _ = fi32
                _ = v33
            }
            var v34 int8 = -128
            _ = fi28
            _ = v34
        }
    default:
        fmt.Println(v12)
    }
    if (42 != -128) {
        if len(v4) > 0 {
            v4[0] = 1
        }
        var v35 uint = 18446744073709551615
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        for fi37 := 0; fi37 < 3; fi37++ {
            switch fi37 {
            case 2147483647:
                if false {
                    v12 = "foo bar"
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    var v38 int16 = -32768
                    _ = v38
                } else {
                    v39 := GenS1{F0: (v12 != v12)}
                    _ = v39
                }
            case -390:
                v2 = (v2 | (v2 + v2))
                v12 = (("a" + (v12 + v12)) + v12)
            }
            fmt.Println(fi36, v2)
            if false {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            } else {
                if true {
                    fmt.Println(fi36)
                    switch -2147483648 {
                    case 42:
                        fmt.Println(v1)
                    case -514:
                        if (-2147483648 >= -190) {
                            fmt.Println(fi36, v12, v2)
                            var v40 string = "foo bar"
                            _ = v40
                        }
                    case 1:
                        v41 := GenS3{F0: 42}
                        _ = v41
                    default:
                        v12 = (v12 + "foo bar")
                    }
                } else {
                    switch (v12 + "foo bar") {
                    case "":
                        if (!true) {
                            var v42 int16 = 709
                            v43 := GenS2{F0: 0}
                            if false {
                                v12 = v12
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                                switch (fi36 ^ v1) {
                                case 0:
                                    v42 = (v42 / (((v42 ^ -450)) | 1))
                                case -9223372036854775808:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    switch v12 {
                                    case "foo bar":
                                        v44 := GenS3{F0: 42}
                                        _ = v44
                                    case "hello":
                                        v45 := GenS4{F0: 1}
                                        _ = v45
                                    default:
                                        v2 += 42
                                    }
                                }
                                v2 += (v2 << ((42) & 63))
                                var v46 string = v12
                                _ = v46
                            } else {
                                v3 = append(v3, 0)
                                v2 -= 739
                            }
                            _ = v42
                            _ = v43
                        } else {
                            if false {
                                v12 = v12
                                v47 := GenS5{F0: "hello", F1: v2}
                                v12 = v12
                                fmt.Println(fi36, v12, fi37)
                                _ = v47
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v1 += (fi36 / ((fi36) | 1))
                            var v48 int16 = 1
                            _ = v48
                        }
                    case "foo bar":
                        var v49 int32 = -2147483648
                        _ = v49
                    default:
                        var v50 int64 = 9223372036854775807
                        _ = v50
                    }
                    _ = len(v3)
                }
                v1 = fi36
                var v51 uint8 = 255
                v52 := GenS3{F0: 193}
                switch 186 {
                case 128:
                    var v53 uint = (v2 ^ v2)
                    _ = v53
                case -1:
                    fmt.Println(v51, v12)
                default:
                    v12 = (v12 + v12)
                }
                _ = v51
                _ = v52
            }
            if ((^(v2)) == 887) {
                if ((0.1 / 0.1) <= ((-2.25 * (-2.25 * 3.14159)) / 1.0)) {
                    var v54 uint32 = 1
                    switch 4294967295 {
                    case 4294967295:
                        switch (("" + "hello") + (v12 + "hello")) {
                        case "":
                            var v55 int8 = -24
                            _ = v55
                        default:
                            switch (v12 + v12) {
                            case "a":
                                v2 = uint(v54)
                            case "":
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    v56 := GenS5{F0: v12, F1: 0}
                                    _ = v56
                                case "foo bar":
                                    if false {
                                        v2 = v2
                                    } else {
                                        v57 := GenS6{F0: 112, F1: uint64(v54), F2: 0}
                                        v12 = ("a" + v12)
                                        _ = v57
                                    }
                                    switch (v1 >> (((v2 % ((0) | 1))) & 63)) {
                                    case 42:
                                        v58 := GenS2{F0: 218}
                                        _ = v58
                                    default:
                                        if true {
                                            if (3.14159 >= 1.5) {
                                                v59 := GenS7{F0: true}
                                                _ = v59
                                            }
                                            v60 := GenS8{F0: (v2 - (v2 / (((^(v2))) | 1))), F1: 18446744073709551615, F2: 1}
                                            _ = v60
                                        } else {
                                            var v61 int16 = 1
                                            v62 := GenS8{F0: v2, F1: 733, F2: 720}
                                            if false {
                                                var v63 int32 = -2147483648
                                                if false {
                                                    var v64 bool = (!(100 >= int8(v1)))
                                                    if ((-2.25 + float32(1e10)) <= 0.0) {
                                                        var v65 uint = (v2 | (v2 * v2))
                                                        switch (v12 + "a") {
                                                        case "foo bar":
                                                            var v66 uint64 = 42
                                                            _ = v66
                                                        case "":
                                                            fmt.Println(v61, v2, v61)
                                                            var v67 float32 = 1.5
                                                            _ = v67
                                                        case "\n\t":
                                                            v54 = v54
                                                        }
                                                        if v64 {
                                                            if ((!((1e10 > 1e10) || false)) && (127 <= 234)) {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v68 int64 = 1
                                                                fmt.Println(v63, v2)
                                                                _ = v68
                                                            }
                                                            v65 = 0
                                                            var v69 uint32 = v54
                                                            _ = v69
                                                        } else {
                                                            _ = len(v4)
                                                        }
                                                        fmt.Println(v12, v54, v1)
                                                        _ = v65
                                                    } else {
                                                        var v70 int32 = 109
                                                        fmt.Println(fi37, v1, v61)
                                                        _ = v70
                                                    }
                                                    _ = v64
                                                } else {
                                                    switch (^(fi37)) {
                                                    case -648:
                                                        if (!(971 <= 18446744073709551615)) {
                                                            v4 = append(v4, 2147483647)
                                                            v71 := GenS9{F0: v12}
                                                            var v72 uint8 = 127
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            if (0.1 < (-2.25 + 0.1)) {
                                                                if (true || true) {
                                                                    switch (v12 + v12) {
                                                                    case "hello":
                                                                        switch v12 {
                                                                        case "a":
                                                                            switch v1 {
                                                                            case 32768:
                                                                                if false {
                                                                                    v12 = v12
                                                                                    _ = len(v3)
                                                                                    var v73 string = (v12 + "")
                                                                                    _ = v73
                                                                                }
                                                                            }
                                                                            v4 = append(v4, v1)
                                                                        case "foo bar":
                                                                            if (float32(float64(0.0)) > 0.1) {
                                                                                if (v12 >= (v12 + "foo bar")) {
                                                                                    var v74 string = v12
                                                                                    if ("hello" >= v74) {
                                                                                        var v75 int32 = v63
                                                                                        if ((v75 ^ (v75 >> (((v2 * 1)) & 63))) >= int32(v54)) {
                                                                                            var v76 int32 = v75
                                                                                            var v77 bool = (!true)
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                        }
                                                                                        _ = v75
                                                                                    }
                                                                                    v12 = v12
                                                                                    _ = v74
                                                                                } else {
                                                                                    if false {
                                                                                        var v78 float64 = 0.0
                                                                                        _ = v78
                                                                                    }
                                                                                    var v79 uint16 = genFn0(v2)
                                                                                    _ = v79
                                                                                }
                                                                                var v80 int32 = (v63 - (v63 * (v63 >> ((4294967295) & 63))))
                                                                                v81 := GenS10{F0: (^(v54)), F1: 842, F2: 127}
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                var v82 int64 = 160
                                                                                v1 = (fi36 % ((42) | 1))
                                                                                _ = v82
                                                                            }
                                                                        default:
                                                                            if (!(v12 < v12)) {
                                                                                v83 := GenS11{F0: ((-(v63)) >= 0), F1: -32}
                                                                                if (v63 <= -2147483648) {
                                                                                    var v84 int32 = v63
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = fi36
                                                                                    }
                                                                                    var v85 string = v12
                                                                                    v3 = append(v3, v54)
                                                                                    v63 -= (v84 >> (((v2 | (v2 + v2))) & 63))
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                } else {
                                                                                    switch ("" + "hello") {
                                                                                    case "":
                                                                                        if (!true) {
                                                                                            var v86 uint16 = 0
                                                                                            v3 = append(v3, v54)
                                                                                            fmt.Println(fi36, v12)
                                                                                            _ = v86
                                                                                        } else {
                                                                                            v2 = 255
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(v54, fi37)
                                                                                        var v87 string = (("hello" + v12) + v12)
                                                                                        _ = v87
                                                                                    case "foo bar":
                                                                                        var v88 string = ""
                                                                                        _ = v88
                                                                                    default:
                                                                                        if ((!true) && (!(!(!((42 >= -8) || true))))) {
                                                                                            v12 = (v12 + "\n\t")
                                                                                        }
                                                                                    }
                                                                                    v1 = fi36
                                                                                    v89 := GenS7{F0: true}
                                                                                    fmt.Println(v1, v54)
                                                                                    _ = v89
                                                                                }
                                                                                _ = v83
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v2, fi37, v12)
                                                }
                                                _ = v63
                                            } else {
                                                v12 = ("" + v12)
                                            }
                                            _ = v61
                                            _ = v62
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = v54
                                    }
                                }
                            case "foo bar":
                                if (("" + ((("a" + (v12 + v12)) + "foo bar") + (v12 + ""))) < ("a" + v12)) {
                                    v90 := GenS12{F0: 127, F1: 127, F2: v54}
                                    v91 := GenS13{F0: true, F1: 32767}
                                    _ = v90
                                    _ = v91
                                }
                            default:
                                if (18446744073709551615 != 29) {
                                    v54 -= (v54 + v54)
                                }
                            }
                        }
                    case 42:
                        v92 := GenS12{F0: (fi37 + 4294967295), F1: 1, F2: 67}
                        _ = v92
                    }
                    _ = v54
                }
                fmt.Println(fi36)
                var v93 uint64 = 29
                _ = v93
            } else {
                v94 := GenS14{F0: 2, F1: fi37, F2: v2}
                _ = v94
            }
            _ = fi37
        }
        if false {
            var v95 uint16 = 949
            var v96 uint8 = 1
            if len(v4) > 0 {
                v4[0] = v1
            }
            _ = v95
            _ = v96
        } else {
            v97 := GenS3{F0: 100}
            _ = v97
        }
        _ = fi36
    }
    v1 = v1
    v98 := GenS15{F0: 123}
    v4 = append(v4, 127)
    var v99 uint8 = 128
    if true {
        for fi100 := 0; fi100 < 8; fi100++ {
            var v101 string = v12
            var v102 string = v101
            switch (v12 + v12) {
            case "":
                v101 = "a"
            }
            _ = fi100
            _ = v101
            _ = v102
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v103 int8 = -19
        _ = v103
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v98
    _ = v99
}
