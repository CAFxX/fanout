package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 string; F1 string; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int8; F1 string; F2 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint32; F1 uint16 }
type GenS15 struct { F0 int; F1 uint8 }
type GenS16 struct { F0 uint16; F1 int16; F2 int16 }
type GenS17 struct { F0 int64; F1 uint; F2 uint16 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint16; F1 uint; F2 int32 }
type GenS20 struct { F0 int; F1 int }
type GenS21 struct { F0 int32; F1 int32 }
type GenS22 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS23 struct { F0 uint8; F1 int32 }
type GenS24 struct { F0 uint64; F1 int16; F2 int32 }
type GenS25 struct { F0 uint; F1 int }
type GenS26 struct { F0 int8; F1 uint16 }
type GenS27 struct { F0 int; F1 int32; F2 int32 }
func genFn0(p0 int32) uint8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint32 = 42
        for fi7 := 0; fi7 < 10; fi7++ {
            switch fi5 {
            case 4294967295:
                fmt.Println(fi7, fi7, v6)
            }
            _ = fi7
        }
        v8 := []int{fi5, fi5, 32767}
        _ = fi5
        _ = v6
        _ = v8
    }
    p0 = (-(p0))
    v9 := GenS1{F0: 76}
    var v10 string = "hello"
    var v11 float64 = 1.5
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 2
}
func genFn1(p0 uint8, p1 uint64, p2 uint8) uint8 {
    p0 = p0
    v12 := []int{}
    fmt.Println(p2, p1, p0)
    switch (("\n\t" + "\n\t") + "hello") {
    case "foo bar":
        fmt.Println(p2, p2)
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (^(p0))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 255
    var v3 int8 = -128
    v4 := []int{-2147483648, 128}
    if (-1 <= 1) {
        v3 = 1
    } else {
        v3 = (v3 | (v3 * v3))
        v4 = append(v4, 32768)
    }
    v2 += v2
    v2 = 1
    var v13 int16 = 42
    switch ((("foo bar" + "foo bar") + ("" + (("hello" + "hello") + ("\n\t" + "foo bar")))) + "hello") {
    case "":
        for fi14 := 0; fi14 < 9; fi14++ {
            switch "foo bar" {
            case "hello":
                switch "hello" {
                case "":
                    v3 = v3
                case "hello":
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        v2 += (v2 * (v2 * v2))
                    default:
                        var v15 int = (fi14 - fi14)
                        _ = v15
                    }
                default:
                    v16 := GenS2{F0: "a", F1: -1}
                    _ = v16
                }
            default:
                v2 += 762
            }
            var v17 int64 = 1
            v3 = 1
            _ = fi14
            _ = v17
        }
    case "\n\t":
        delete(v1, "\n\t")
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        switch fi18 {
        case 679:
            var v19 uint64 = 0
            _ = v19
        case 0:
            var v20 uint32 = 2147483648
            _ = v20
        default:
            fmt.Println(v3, v13, v2)
        }
        var v21 uint32 = 1
        switch "a" {
        case "\n\t":
            if true {
                v3 = v3
                switch fi18 {
                case 42:
                    var v22 int64 = 438
                    _ = v22
                }
                var v23 string = ""
                _ = v23
            } else {
                if false {
                    fmt.Println(v2, fi18)
                    v2 += v2
                } else {
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v2, v2)
                    case "":
                        var v24 int8 = (v3 % ((105) | 1))
                        _ = v24
                    case "foo bar":
                        fmt.Println(fi18, v2, fi18)
                    }
                    if true {
                        var v25 int8 = (v3 & (v3 ^ (v3 << ((18446744073709551615) & 63))))
                        var v26 int16 = (v13 * (v13 * (v13 % ((-58) | 1))))
                        _ = v25
                        _ = v26
                    }
                    var v27 uint = 65535
                    _ = v27
                }
                if ((0.0 - -2.25) != float64(float32(3.14159))) {
                    v28 := GenS2{F0: ("\n\t" + "foo bar"), F1: -9223372036854775808}
                    if ((v2 + v2) != (v2 ^ 18446744073709551615)) {
                        switch (fi18 & int(v3)) {
                        case -219:
                            v29 := GenS2{F0: ("hello" + ""), F1: -84}
                            _ = v29
                        case 431:
                            if (false && true) {
                                v2 -= 4294967295
                            } else {
                                v30 := GenS3{F0: 1}
                                _ = v30
                            }
                        case -833:
                            switch 1 {
                            case -1:
                                var v31 uint16 = 2
                                _ = v31
                            case -463:
                                switch "hello" {
                                case "a":
                                    if ((fi18 & fi18) > fi18) {
                                        var v32 uint32 = (v21 >> ((v2) & 63))
                                        _ = v32
                                    } else {
                                        if (((v3 ^ v3) >= 127) && (false && true)) {
                                            var v33 int16 = (v13 - v13)
                                            _ = v33
                                        }
                                    }
                                case "\n\t":
                                    var v34 int32 = -280
                                    _ = v34
                                }
                            }
                        }
                        if (genFn1(128, 4294967295, 128) < 2) {
                            var v35 string = ((("a" + (((("" + ((("hello" + "\n\t") + ("" + "hello")) + ("\n\t" + "hello"))) + ("a" + "hello")) + ("\n\t" + "")) + "")) + "") + "a")
                            switch v35 {
                            case "a":
                                v36 := GenS4{F0: true}
                                _ = v36
                            case "":
                                var v37 int = 0
                                _ = v37
                            }
                            var v38 float64 = ((0.1 / 0.1) - (float64(3.14159) - (float64(((((3.14159 / 1e10) + 3.14159) - 1.5) - 1e10)) * 0.0)))
                            _ = v35
                            _ = v38
                        }
                    }
                    var v39 uint64 = 2
                    _ = v28
                    _ = v39
                } else {
                    if (v2 <= 2) {
                        if false {
                            for fi40 := 0; fi40 < 20; fi40++ {
                                v13 -= (v13 | (v13 + v13))
                                var v41 float64 = (-2.25 / (0.0 + 1e10))
                                if true {
                                    fmt.Println(v21)
                                    v42 := GenS4{F0: true}
                                    switch ("hello" + (("\n\t" + "\n\t") + "foo bar")) {
                                    case "foo bar":
                                        var v43 uint32 = (^(v21))
                                        if (v3 == 59) {
                                            var v44 int = (fi40 - fi18)
                                            fmt.Println(v21)
                                            if false {
                                                v43 = 4294967295
                                                switch "hello" {
                                                case "foo bar":
                                                    if (true || (-2147483648 <= 1)) {
                                                        if false {
                                                            if (!true) {
                                                                v45 := GenS5{F0: 1, F1: 0}
                                                                _ = len(v1)
                                                                switch fi18 {
                                                                case -9223372036854775808:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                case -918:
                                                                    var v46 int16 = (v13 % ((v13) | 1))
                                                                    _ = v46
                                                                case 1:
                                                                    if false {
                                                                        var v47 string = ("foo bar" + ("" + "\n\t"))
                                                                        _ = v47
                                                                    } else {
                                                                        fmt.Println(v21, v41, v44)
                                                                        delete(v1, ((("\n\t" + "hello") + "a") + ("a" + (("" + "hello") + ((("\n\t" + "\n\t") + "a") + ("hello" + "a"))))))
                                                                    }
                                                                    if false {
                                                                        switch (fi18 % ((v44) | 1)) {
                                                                        case 624:
                                                                            v4 = append(v4, v44)
                                                                        case -894:
                                                                            v48 := GenS6{F0: ((((("\n\t" + "foo bar") + (("a" + "a") + ("a" + "\n\t"))) + "hello") + (("a" + "foo bar") + "foo bar")) + "hello"), F1: "", F2: -9223372036854775808}
                                                                            _ = v48
                                                                        }
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v13 -= v13
                                                                    } else {
                                                                        if (v21 == (v21 >> (((v2 - v2)) & 63))) {
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                if (4294967295 == 747) {
                                                                                    switch (-(fi18)) {
                                                                                    case -360:
                                                                                        var v49 int = (v44 * fi18)
                                                                                        v21 = v21
                                                                                        _ = v49
                                                                                    case 29:
                                                                                        if ((!(false && true)) || true) {
                                                                                            v50 := GenS5{F0: 0, F1: 255}
                                                                                            if false {
                                                                                                fmt.Println(v43, v2, v3)
                                                                                                v51 := v1[("a" + "foo bar")]
                                                                                                _ = v51
                                                                                            }
                                                                                            _ = v50
                                                                                        } else {
                                                                                            v1["a"] = "\n\t"
                                                                                            v44 += (fi40 << (((v2 >> ((uint(v21)) & 63))) & 63))
                                                                                            if true {
                                                                                                v1[("foo bar" + ("" + ("a" + ((("hello" + ("a" + ("\n\t" + "foo bar"))) + "\n\t") + ("\n\t" + "")))))] = (("\n\t" + "\n\t") + "hello")
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v52 float32 = 0.0
                                                                                                    if true {
                                                                                                        var v53 bool = false
                                                                                                        v41 += v41
                                                                                                        _ = v53
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            var v54 int8 = v3
                                                                                                            v55 := GenS7{F0: 18446744073709551615, F1: 1}
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            _ = v54
                                                                                                            _ = v55
                                                                                                        }
                                                                                                        v3 = 25
                                                                                                        if ((fi40 * v44) < int(v2)) {
                                                                                                            var v56 int8 = (v3 % ((int8(v44)) | 1))
                                                                                                            if ((v44 << (((v2 * v2)) & 63)) > (fi18 * 32767)) {
                                                                                                                _ = len(v1)
                                                                                                            } else {
                                                                                                                switch "" {
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v13, fi40)
                                                                                                                case "":
                                                                                                                    if (1 <= 788) {
                                                                                                                        if false {
                                                                                                                            v3 += (v56 | 100)
                                                                                                                            var v57 int32 = 144
                                                                                                                            var v58 int16 = 1000
                                                                                                                            _ = v57
                                                                                                                            _ = v58
                                                                                                                        } else {
                                                                                                                            if false {
                                                                                                                                v59 := GenS3{F0: 0}
                                                                                                                                switch "\n\t" {
                                                                                                                                case "hello":
                                                                                                                                    var v60 int = (fi18 + (-(fi40)))
                                                                                                                                    _ = v60
                                                                                                                                case "":
                                                                                                                                    var v61 string = (("a" + "\n\t") + "")
                                                                                                                                    _ = v61
                                                                                                                                }
                                                                                                                                v62 := GenS5{F0: 236, F1: 255}
                                                                                                                                _ = v59
                                                                                                                                _ = v62
                                                                                                                            }
                                                                                                                            v3 = 59
                                                                                                                            v4 = append(v4, fi40)
                                                                                                                        }
                                                                                                                        v63 := v1[("\n\t" + (("\n\t" + "a") + "a"))]
                                                                                                                        _ = v63
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if true {
                                                                                                                        v64 := GenS8{F0: v56, F1: "\n\t", F2: ""}
                                                                                                                        _ = v64
                                                                                                                    } else {
                                                                                                                        var v65 string = "hello"
                                                                                                                        _ = v65
                                                                                                                    }
                                                                                                                }
                                                                                                                v1["\n\t"] = ("hello" + "hello")
                                                                                                            }
                                                                                                            if (v56 >= int8(fi18)) {
                                                                                                                if false {
                                                                                                                    switch v44 {
                                                                                                                    case -2147483648:
                                                                                                                        if (!true) {
                                                                                                                            v4 = append(v4, (fi18 - (v44 | (v44 / (((v44 / ((v44) | 1))) | 1)))))
                                                                                                                        } else {
                                                                                                                            _ = len(v4)
                                                                                                                            v21 = v43
                                                                                                                        }
                                                                                                                    case 32767:
                                                                                                                        v2 -= (v2 + v2)
                                                                                                                    }
                                                                                                                    var v66 int = (fi18 ^ (^(v44)))
                                                                                                                    var v67 float32 = v52
                                                                                                                    _ = v66
                                                                                                                    _ = v67
                                                                                                                } else {
                                                                                                                    var v68 int32 = int32(v13)
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                                if false {
                                                                                                                    v69 := GenS3{F0: 0}
                                                                                                                    if true {
                                                                                                                        switch (fi40 / (((v44 + v44)) | 1)) {
                                                                                                                        case 32767:
                                                                                                                            v2 -= 65535
                                                                                                                        default:
                                                                                                                            var v70 int16 = (v13 + (v13 + v13))
                                                                                                                            _ = v70
                                                                                                                        }
                                                                                                                        var v71 int64 = 0
                                                                                                                        _ = v71
                                                                                                                    } else {
                                                                                                                        delete(v1, "")
                                                                                                                        if true {
                                                                                                                            v52 += 1.5
                                                                                                                        } else {
                                                                                                                            fmt.Println(v56, v21)
                                                                                                                            switch v44 {
                                                                                                                            case 754:
                                                                                                                                var v72 uint64 = 2
                                                                                                                                _ = v72
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v73 int16 = -1
                                                                                                                        _ = v73
                                                                                                                    }
                                                                                                                    _ = v69
                                                                                                                } else {
                                                                                                                    var v74 int16 = v13
                                                                                                                    fmt.Println(v43)
                                                                                                                    _ = v74
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v56
                                                                                                        } else {
                                                                                                            _ = len(v1)
                                                                                                            v13 += (v13 & 32767)
                                                                                                            v75 := v1["hello"]
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v52
                                                                                                } else {
                                                                                                    var v76 int8 = v3
                                                                                                    _ = v76
                                                                                                }
                                                                                                switch -416 {
                                                                                                case 921:
                                                                                                    v77 := GenS9{F0: (^(v2))}
                                                                                                    _ = v77
                                                                                                case 128:
                                                                                                    var v78 uint64 = 263
                                                                                                    _ = v78
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case 2147483647:
                                                                                        delete(v1, "a")
                                                                                    }
                                                                                    v43 -= 104
                                                                                    var v79 uint8 = 1
                                                                                    _ = v79
                                                                                }
                                                                            case "a":
                                                                                _ = len(v1)
                                                                            case "":
                                                                                v2 -= v2
                                                                            }
                                                                            v80 := GenS10{F0: ("" + ("" + ("\n\t" + "\n\t"))), F1: 255}
                                                                            var v81 float32 = 3.14159
                                                                            _ = v80
                                                                            _ = v81
                                                                        }
                                                                        fmt.Println(v3)
                                                                        var v82 int64 = -133
                                                                        _ = v82
                                                                    }
                                                                    v83 := GenS11{F0: 4294967295}
                                                                    var v84 int = fi40
                                                                    _ = v83
                                                                    _ = v84
                                                                default:
                                                                    if true {
                                                                        if ((v3 * -128) > v3) {
                                                                            v43 += v21
                                                                            var v85 int64 = 42
                                                                            switch fi40 {
                                                                            case 42:
                                                                                if false {
                                                                                    fmt.Println(v2, fi40, v43)
                                                                                    if (true || false) {
                                                                                        switch -935 {
                                                                                        case 32768:
                                                                                            var v86 uint16 = 0
                                                                                            _ = v86
                                                                                        case -451:
                                                                                            if true {
                                                                                                switch 32767 {
                                                                                                case 32768:
                                                                                                    v44 = fi18
                                                                                                case 1:
                                                                                                    var v87 int = (v44 ^ 1)
                                                                                                    _ = v87
                                                                                                case -292:
                                                                                                    if (185 == 42) {
                                                                                                        fmt.Println(v13, fi40, v44)
                                                                                                        switch ("\n\t" + "foo bar") {
                                                                                                        case "":
                                                                                                            v43 = (^(v43))
                                                                                                            v21 -= (v21 / ((v21) | 1))
                                                                                                        default:
                                                                                                            var v88 bool = (!(false && true))
                                                                                                            _ = v88
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v89 := GenS12{F0: 65535}
                                                                                                _ = v89
                                                                                            }
                                                                                        case 1:
                                                                                            v2 += (v2 % (((^(v2))) | 1))
                                                                                        }
                                                                                        switch (^(fi18)) {
                                                                                        case -129:
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                if (v43 <= v43) {
                                                                                                    if true {
                                                                                                        if (v44 != 978) {
                                                                                                            v44 += fi40
                                                                                                        } else {
                                                                                                            var v90 int = 4294967295
                                                                                                            v2 -= 255
                                                                                                            _ = v90
                                                                                                        }
                                                                                                        v3 += v3
                                                                                                        fmt.Println(v85, fi40)
                                                                                                    }
                                                                                                    fmt.Println(v41, v3, v2)
                                                                                                    v91 := GenS10{F0: "foo bar", F1: 198}
                                                                                                    _ = v91
                                                                                                }
                                                                                            }
                                                                                            v43 = 2147483648
                                                                                            switch -1 {
                                                                                            case 60:
                                                                                                if (v85 > v85) {
                                                                                                    v3 += v3
                                                                                                }
                                                                                            default:
                                                                                                v92 := GenS13{F0: 1}
                                                                                                _ = v92
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            v41 = -2.25
                                                                                            switch (fi18 % (((-(fi40))) | 1)) {
                                                                                            case 499:
                                                                                                switch 2147483647 {
                                                                                                case -361:
                                                                                                    fmt.Println(v85)
                                                                                                case 138:
                                                                                                    v43 += (v21 + v21)
                                                                                                }
                                                                                            case 128:
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                            case 127:
                                                                                                if (false || false) {
                                                                                                    switch fi40 {
                                                                                                    case 1:
                                                                                                        v93 := GenS14{F0: v21, F1: 0}
                                                                                                        _ = v93
                                                                                                    case -1:
                                                                                                        v94 := GenS4{F0: (!(1e10 == 0.1))}
                                                                                                        _ = v94
                                                                                                    case 2147483647:
                                                                                                        v95 := GenS1{F0: 128}
                                                                                                        _ = v95
                                                                                                    }
                                                                                                    v96 := GenS5{F0: 53, F1: 127}
                                                                                                    fmt.Println(fi40)
                                                                                                    v3 += v3
                                                                                                    _ = v96
                                                                                                } else {
                                                                                                    if (true && (!true)) {
                                                                                                        var v97 int16 = (v13 / (((v13 + 32767)) | 1))
                                                                                                        if true {
                                                                                                            v13 = (^(v97))
                                                                                                            v97 += v97
                                                                                                            switch 128 {
                                                                                                            case -314:
                                                                                                                fmt.Println(fi18, v21)
                                                                                                            default:
                                                                                                                var v98 int16 = (-(v13))
                                                                                                                v21 -= 4294967295
                                                                                                                fmt.Println(v97)
                                                                                                                _ = v98
                                                                                                            }
                                                                                                            v2 += 65535
                                                                                                        } else {
                                                                                                            v97 = v97
                                                                                                        }
                                                                                                        v3 += v3
                                                                                                        if true {
                                                                                                            var v99 uint8 = 29
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            if ((v41 * 1e10) <= (v41 + (0.0 / 1.0))) {
                                                                                                                fmt.Println(v97)
                                                                                                                var v100 int = (v44 * (fi40 ^ fi40))
                                                                                                                if (int16(v43) == (v97 ^ (^(v13)))) {
                                                                                                                    v97 -= (-(v97))
                                                                                                                } else {
                                                                                                                    var v101 uint16 = 981
                                                                                                                    switch "\n\t" {
                                                                                                                    case "a":
                                                                                                                        fmt.Println(v97)
                                                                                                                    case "\n\t":
                                                                                                                        _ = len(v1)
                                                                                                                    }
                                                                                                                    var v102 int8 = int8(fi18)
                                                                                                                    _ = v101
                                                                                                                    _ = v102
                                                                                                                }
                                                                                                                _ = v100
                                                                                                            } else {
                                                                                                                switch v44 {
                                                                                                                case -372:
                                                                                                                    v44 = (v44 * fi40)
                                                                                                                case 9223372036854775807:
                                                                                                                    if (((v2 >> ((1) & 63)) >= 2) && ((255 <= 255) || true)) {
                                                                                                                        var v103 uint64 = 96
                                                                                                                        _ = v103
                                                                                                                    } else {
                                                                                                                        v41 -= v41
                                                                                                                    }
                                                                                                                }
                                                                                                                v104 := GenS7{F0: 18446744073709551615, F1: 32767}
                                                                                                                _ = v104
                                                                                                            }
                                                                                                            if false {
                                                                                                                fmt.Println(v44, v21)
                                                                                                            } else {
                                                                                                                var v105 float32 = 1e10
                                                                                                                fmt.Println(v85)
                                                                                                                _ = v105
                                                                                                            }
                                                                                                        }
                                                                                                        v2 += 65535
                                                                                                        _ = v97
                                                                                                    } else {
                                                                                                        switch "hello" {
                                                                                                        case "hello":
                                                                                                            v106 := GenS15{F0: fi40, F1: genFn1(2, 1, 128)}
                                                                                                            _ = v106
                                                                                                        case "\n\t":
                                                                                                            if false {
                                                                                                                if true {
                                                                                                                    v2 += (v2 % ((v2) | 1))
                                                                                                                    var v107 uint64 = uint64(v2)
                                                                                                                    _ = v107
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            _ = len(v4)
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v108 uint = v2
                                                                                                            if true {
                                                                                                                var v109 uint8 = 39
                                                                                                                if (555 > (^(v108))) {
                                                                                                                    delete(v1, "\n\t")
                                                                                                                    var v110 int16 = v13
                                                                                                                    var v111 uint32 = v43
                                                                                                                    var v112 uint16 = 272
                                                                                                                    fmt.Println(v2, v21, v85)
                                                                                                                    if true {
                                                                                                                        v113 := GenS16{F0: v112, F1: v13, F2: v13}
                                                                                                                        switch "\n\t" {
                                                                                                                        case "":
                                                                                                                            var v114 uint32 = uint32(v112)
                                                                                                                            _ = v114
                                                                                                                        case "\n\t":
                                                                                                                            v115 := GenS17{F0: v85, F1: 181, F2: (v112 << (((v2 / ((v108) | 1))) & 63))}
                                                                                                                            _ = v115
                                                                                                                        }
                                                                                                                        v116 := GenS9{F0: 882}
                                                                                                                        _ = v113
                                                                                                                        _ = v116
                                                                                                                    } else {
                                                                                                                        v2 += (v108 * 42)
                                                                                                                        v108 = 1
                                                                                                                    }
                                                                                                                    _ = v110
                                                                                                                    _ = v111
                                                                                                                    _ = v112
                                                                                                                } else {
                                                                                                                    v21 = v21
                                                                                                                    v117 := GenS14{F0: (v43 + uint32(v2)), F1: 65535}
                                                                                                                    var v118 float32 = (0.0 + float32(0.0))
                                                                                                                    _ = v117
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                                _ = v109
                                                                                                            } else {
                                                                                                                v119 := GenS1{F0: genFn1(1, 2, 2)}
                                                                                                                v41 += v41
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            fmt.Println(v44, fi18)
                                                                                                            var v120 uint64 = 25
                                                                                                            v44 = fi18
                                                                                                            _ = v108
                                                                                                            _ = v120
                                                                                                        } else {
                                                                                                            v85 += 0
                                                                                                            if ((1.5 / v41) >= -2.25) {
                                                                                                                fmt.Println(v43, fi18, v3)
                                                                                                            } else {
                                                                                                                fmt.Println(v2, v2)
                                                                                                                v21 += v21
                                                                                                            }
                                                                                                        }
                                                                                                        var v121 int = fi40
                                                                                                        _ = v121
                                                                                                    }
                                                                                                    var v122 bool = true
                                                                                                    _ = v122
                                                                                                }
                                                                                            default:
                                                                                                var v123 float64 = 1e10
                                                                                                _ = v123
                                                                                            }
                                                                                        }
                                                                                        v1["foo bar"] = (((("a" + "foo bar") + ("" + "")) + "hello") + "\n\t")
                                                                                        if (v3 <= v3) {
                                                                                            _ = len(v1)
                                                                                        } else {
                                                                                            var v124 float32 = 3.14159
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v124
                                                                                        }
                                                                                    }
                                                                                    var v125 uint = v2
                                                                                    var v126 int32 = int32(v21)
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                            case 0:
                                                                                var v127 uint64 = 4294967295
                                                                                _ = v127
                                                                            }
                                                                            _ = v85
                                                                        }
                                                                        if false {
                                                                            var v128 uint16 = 490
                                                                            v13 -= 42
                                                                            _ = len(v4)
                                                                            _ = v128
                                                                        }
                                                                    }
                                                                }
                                                                _ = v45
                                                            } else {
                                                                v129 := GenS4{F0: (false && (4294967295 > 0))}
                                                                _ = v129
                                                            }
                                                            var v130 uint16 = 32767
                                                            switch (fi40 & fi18) {
                                                            case 720:
                                                                v131 := GenS18{F0: "a"}
                                                                _ = v131
                                                            case 928:
                                                                var v132 uint64 = 688
                                                                _ = v132
                                                            default:
                                                                if false {
                                                                    v44 += fi18
                                                                } else {
                                                                    v133 := GenS11{F0: 1}
                                                                    _ = v133
                                                                }
                                                            }
                                                            _ = v130
                                                        } else {
                                                            var v134 uint64 = 18446744073709551615
                                                            switch (fi18 * 127) {
                                                            case 127:
                                                                var v135 uint64 = (v134 ^ 1)
                                                                _ = v135
                                                            case 0:
                                                                v136 := GenS6{F0: ("foo bar" + (("\n\t" + "\n\t") + ("a" + "a"))), F1: (("" + ("" + "hello")) + "hello"), F2: 9223372036854775807}
                                                                _ = v136
                                                            case 4294967295:
                                                                _ = len(v4)
                                                            }
                                                            v137 := GenS9{F0: (v2 >> ((v2) & 63))}
                                                            _ = v134
                                                            _ = v137
                                                        }
                                                        v4 = append(v4, v44)
                                                        v21 += v43
                                                    } else {
                                                        v44 = fi40
                                                    }
                                                case "\n\t":
                                                    v138 := GenS3{F0: 160}
                                                    _ = v138
                                                }
                                            }
                                            _ = v44
                                        } else {
                                            var v139 bool = (true || false)
                                            var v140 int64 = 0
                                            v141 := GenS17{F0: v140, F1: v2, F2: 971}
                                            if v139 {
                                                var v142 int64 = v140
                                                switch "" {
                                                case "hello":
                                                    var v143 float32 = (((0.1 / 0.1) / 1.0) * (0.0 * 1.5))
                                                    _ = v143
                                                case "a":
                                                    v144 := GenS19{F0: 145, F1: 65535, F2: -262}
                                                    _ = v144
                                                case "foo bar":
                                                    delete(v1, "\n\t")
                                                default:
                                                    switch ("" + "\n\t") {
                                                    case "a":
                                                        v139 = ((v139 && (v139 || v139)) || (2147483648 < v43))
                                                    case "\n\t":
                                                        v3 += 42
                                                    default:
                                                        switch fi40 {
                                                        case 0:
                                                            v3 = v3
                                                        case -1:
                                                            v3 = (v3 + v3)
                                                        case 1:
                                                            var v145 uint16 = 32767
                                                            _ = v145
                                                        }
                                                    }
                                                }
                                                _ = v142
                                            } else {
                                                fmt.Println(v140, fi18)
                                            }
                                            v146 := GenS1{F0: genFn1(127, 2, 42)}
                                            fmt.Println(v43, fi18, v13)
                                            switch fi40 {
                                            case 32767:
                                                if (v139 || v139) {
                                                    var v147 int64 = 392
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v43 -= v43
                                                    _ = v147
                                                }
                                            case -1:
                                                switch "a" {
                                                case "\n\t":
                                                    if v139 {
                                                        fmt.Println(v2)
                                                        v2 -= (v2 + v2)
                                                        v140 = v140
                                                    } else {
                                                        switch (fi18 | (fi40 & fi18)) {
                                                        case 42:
                                                            v148 := v1["a"]
                                                            _ = v148
                                                        }
                                                        v13 = v13
                                                        v149 := GenS2{F0: (("" + "hello") + ""), F1: (-(v140))}
                                                        v150 := GenS3{F0: 131}
                                                        _ = v149
                                                        _ = v150
                                                    }
                                                case "":
                                                    if (459 >= 18446744073709551615) {
                                                        delete(v1, "hello")
                                                        var v151 int16 = v13
                                                        _ = v151
                                                    } else {
                                                        var v152 uint32 = (v21 | 161)
                                                        var v153 int16 = 1
                                                        v13 = -32768
                                                        _ = v152
                                                        _ = v153
                                                    }
                                                case "a":
                                                    v154 := GenS17{F0: int64(v21), F1: 836, F2: 32767}
                                                    _ = v154
                                                }
                                            default:
                                                if len(v4) > 0 {
                                                    v4[0] = fi40
                                                }
                                                _ = len(v1)
                                            }
                                            _ = len(v1)
                                            if v139 {
                                                v21 += (v43 & v43)
                                                v13 = (v13 | -1)
                                            } else {
                                                var v155 uint32 = v43
                                                v140 = -1
                                                var v156 uint32 = (v43 & 2147483647)
                                                _ = v155
                                                _ = v156
                                            }
                                            _ = v139
                                            _ = v140
                                            _ = v141
                                            _ = v146
                                        }
                                        _ = v43
                                    }
                                    _ = v42
                                } else {
                                    switch fi18 {
                                    case 621:
                                        var v157 int32 = 1
                                        _ = v157
                                    case 32768:
                                        if false {
                                            v158 := GenS1{F0: genFn0(42)}
                                            if (true || (((v2 != (v2 / ((v2) | 1))) && (false || (true || true))) || false)) {
                                                if (1e10 >= ((3.14159 + (1.5 * 1e10)) * (-2.25 / 1.0))) {
                                                    v13 += (v13 + v13)
                                                    var v159 float64 = float64(float32(v41))
                                                    var v160 uint32 = (^(v21))
                                                    _ = v159
                                                    _ = v160
                                                } else {
                                                    v3 = 1
                                                    var v161 int = (fi40 * (fi18 & fi18))
                                                    fmt.Println(v2, fi18)
                                                    _ = v161
                                                }
                                                v41 = v41
                                                v41 += float64(-2.25)
                                            } else {
                                                var v162 uint8 = 255
                                                _ = v162
                                            }
                                            _ = v158
                                        }
                                    default:
                                        v2 += (v2 + (v2 + (v2 ^ (v2 & 18446744073709551615))))
                                    }
                                    _ = len(v1)
                                }
                                var v163 uint8 = uint8(v21)
                                if true {
                                    fmt.Println(v2, v13, v13)
                                    if true {
                                        if false {
                                            var v164 int64 = 966
                                            _ = v164
                                        }
                                        var v165 bool = (((2 > 801) || ("a" < "hello")) || (0.0 <= (((1e10 / ((0.0 + (0.1 + (3.14159 - 1e10))) + ((1e10 - ((1.5 / 1.0) * 0.1)) - (float32(float64(-2.25)) + (1e10 / (-2.25 - 1e10)))))) / 1.0) - 3.14159)))
                                        v166 := GenS20{F0: 788, F1: int(v21)}
                                        _ = v165
                                        _ = v166
                                    } else {
                                        v167 := GenS21{F0: -159, F1: 2147483647}
                                        var v168 int16 = (^(v13))
                                        _ = v167
                                        _ = v168
                                    }
                                    v169 := GenS22{F0: v163, F1: (v21 >> ((v2) & 63)), F2: v3}
                                    _ = v169
                                }
                                _ = fi40
                                _ = v163
                                _ = v41
                            }
                            var v170 int64 = 1
                            for fi171 := 0; fi171 < 15; fi171++ {
                                fmt.Println(v13, v3, fi18)
                                var v172 int = 32768
                                _ = fi171
                                _ = v172
                            }
                            _ = v170
                        }
                        v173 := GenS4{F0: (58 < 0)}
                        v3 += (v3 / (((v3 ^ -38)) | 1))
                        _ = v173
                    }
                    v3 -= (v3 / ((v3) | 1))
                    for fi174 := 0; fi174 < 16; fi174++ {
                        var v175 uint16 = 0
                        v21 = (v21 | (v21 >> (((^(v2))) & 63)))
                        v1[("a" + ("a" + ""))] = "hello"
                        _ = fi174
                        _ = v175
                    }
                }
                for fi176 := 0; fi176 < 7; fi176++ {
                    v3 = int8(v13)
                    if true {
                        fmt.Println(fi18)
                        var v177 float32 = 1.5
                        var v178 int64 = -1
                        v177 -= ((((v177 + -2.25) * 1.5) * (v177 * v177)) + 0.0)
                        v177 += v177
                        if true {
                            var v179 int = fi176
                            var v180 uint32 = v21
                            _ = v179
                            _ = v180
                        } else {
                            v181 := GenS23{F0: 128, F1: -282}
                            _ = v181
                        }
                        if (uint16(v2) < 446) {
                            if (false && (("hello" + "foo bar") == "hello")) {
                                var v182 uint8 = genFn1(159, 18446744073709551615, genFn0(0))
                                if false {
                                    v183 := GenS6{F0: "\n\t", F1: (("hello" + "hello") + ""), F2: v178}
                                    switch ("a" + "") {
                                    case "":
                                        var v184 uint = v2
                                        _ = v184
                                    case "foo bar":
                                        if true {
                                            if (int(v21) != (fi176 >> ((4294967295) & 63))) {
                                                v185 := GenS7{F0: 1, F1: 2}
                                                v186 := GenS24{F0: 18446744073709551615, F1: -419, F2: 42}
                                                v187 := GenS9{F0: v2}
                                                _ = v185
                                                _ = v186
                                                _ = v187
                                            }
                                            if ((-(fi18)) < (-(fi18))) {
                                                switch ("foo bar" + "\n\t") {
                                                case "hello":
                                                    if (!(false || ((v13 != v13) && true))) {
                                                        v182 += uint8(v2)
                                                        v3 -= v3
                                                    } else {
                                                        var v188 uint = (v2 + (v2 ^ 0))
                                                        v177 -= 1.5
                                                        switch ("a" + "foo bar") {
                                                        case "foo bar":
                                                            delete(v1, ((("foo bar" + ("\n\t" + (("a" + "hello") + ("\n\t" + ("foo bar" + ""))))) + "hello") + ("" + ("\n\t" + ""))))
                                                        case "a":
                                                            v177 = v177
                                                        case "hello":
                                                            v189 := GenS25{F0: v2, F1: int(v182)}
                                                            _ = v189
                                                        default:
                                                            v1["hello"] = ""
                                                        }
                                                        _ = v188
                                                    }
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        v13 -= (-(v13))
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = fi18
                                    }
                                    if (true && true) {
                                        var v190 int8 = (v3 | (v3 << ((603) & 63)))
                                        _ = v190
                                    }
                                    v191 := GenS5{F0: 255, F1: (v182 ^ (^(v182)))}
                                    _ = v183
                                    _ = v191
                                } else {
                                    var v192 int8 = v3
                                    if false {
                                        fmt.Println(v182, v21)
                                    } else {
                                        var v193 bool = false
                                        _ = len(v1)
                                        v192 -= v3
                                        var v194 uint64 = 2
                                        v178 += int64(v3)
                                        _ = v193
                                        _ = v194
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        var v195 uint16 = 65535
                                        _ = v195
                                    default:
                                        if (!true) {
                                            fmt.Println(v2, fi18, v13)
                                        } else {
                                            v196 := GenS5{F0: v182, F1: v182}
                                            if false {
                                                if (fi18 < int(v3)) {
                                                    var v197 uint8 = (v182 - (v182 << ((1) & 63)))
                                                    v197 = 0
                                                    v198 := GenS26{F0: v192, F1: 255}
                                                    _ = v197
                                                    _ = v198
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v21 += 1
                                                    v199 := GenS27{F0: fi18, F1: 2147483647, F2: -713}
                                                    _ = v199
                                                }
                                                switch (fi18 * fi176) {
                                                case 2147483647:
                                                    var v200 uint16 = 930
                                                    _ = v200
                                                default:
                                                    delete(v1, "a")
                                                }
                                            }
                                            v177 += -2.25
                                            _ = v196
                                        }
                                    }
                                    _ = v192
                                }
                                switch 32767 {
                                case -1:
                                    v13 += (v13 * -32768)
                                    if len(v4) > 0 {
                                        v4[0] = int(v2)
                                    }
                                    var v201 uint32 = (^(v21))
                                    _ = v201
                                default:
                                    _ = len(v1)
                                }
                                _ = v182
                            }
                            v202 := GenS21{F0: -1, F1: 42}
                            _ = v202
                        } else {
                            var v203 float64 = (3.14159 / 1.5)
                            _ = v203
                        }
                        _ = v177
                        _ = v178
                    }
                    switch 2147483647 {
                    case 1:
                        var v204 string = "a"
                        _ = v204
                    }
                    _ = fi176
                }
            }
        case "hello":
            fmt.Println(v3, v13, v3)
        }
        _ = fi18
        _ = v21
    }
    fmt.Println(v2, v3)
    v13 += (v13 ^ -32768)
    var v205 uint32 = 65535
    switch -326 {
    case 173:
        var v206 int64 = 9223372036854775807
        _ = v206
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v13 = -99
    fmt.Println(v3, v205, v205)
    v2 -= v2
    v207 := v1["foo bar"]
    var v208 uint = (v2 ^ v2)
    v2 = 527
    if (v13 > -140) {
        var v209 float32 = float32(0.1)
        if true {
            switch v207 {
            case "hello":
                var v210 int64 = -1
                _ = v210
            }
        }
        _ = v209
    }
    var v211 uint32 = v205
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v205:", v205)
    fmt.Println("v207:", v207)
    fmt.Println("v208:", v208)
    fmt.Println("v211:", v211)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v205
    _ = v207
    _ = v208
    _ = v211
    _ = v3
    _ = v4
}
