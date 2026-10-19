package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 bool; F1 int32; F2 int }
type GenS10 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS11 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS12 struct { F0 uint64; F1 bool; F2 int16 }
type GenS13 struct { F0 uint32; F1 int32; F2 int8 }
type GenS14 struct { F0 string; F1 int; F2 bool }
type GenS15 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool; F1 uint32; F2 string }
type GenS18 struct { F0 int8; F1 int64; F2 int64 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint; F1 int64 }
type GenS22 struct { F0 uint16; F1 uint }
type GenS23 struct { F0 int; F1 uint16; F2 int8 }
type GenS24 struct { F0 int8 }

func main() {
    var v1 uint = 10
    v2 := make(map[string]bool)
    v3 := []int{127, 978, -1, 42}
    v4 := []uint32{2147483647}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int16 = 1
        v4 = append(v4, 974)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint64 = 1
        _ = fi7
        _ = v8
    }
    v1 += (v1 % ((v1) | 1))
    v1 = 65535
    v1 -= (v1 / (((v1 % ((694) | 1))) | 1))
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float64 = -2.25
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := v2["foo bar"]
            if v12 {
                v13 := v2[("a" + "")]
                if (!(2 > 127)) {
                    if ((100 <= 5) || v12) {
                        if ((127 != uint8(v1)) && (!((v1 + v1) <= (v1 << ((v1) & 63))))) {
                            v14 := GenS1{F0: 255, F1: 1}
                            fmt.Println(v10, v12, v10)
                            _ = v14
                        }
                        v13 = v13
                    }
                    var v15 float32 = (0.0 + ((((3.14159 - ((3.14159 * (((0.1 / 3.14159) * ((-2.25 * 1e10) * 3.14159)) - 1e10)) - 3.14159)) / (-2.25 * 3.14159)) + float32((((v10 * v10) - (v10 / v10)) / v10))) + -2.25))
                    if v13 {
                        fmt.Println(v12, v10)
                        switch (("foo bar" + "") + "hello") {
                        case "":
                            v15 -= 3.14159
                        }
                    } else {
                        v1 -= v1
                        v16 := GenS2{F0: 1}
                        v15 += (-2.25 - 0.1)
                        _ = v16
                    }
                    _ = v15
                }
                var v17 int8 = -75
                _ = v13
                _ = v17
            } else {
                if true {
                    v18 := GenS3{F0: 42, F1: 255, F2: v1}
                    _ = v18
                }
            }
            _ = fi11
            _ = v12
        }
        switch (("a" + "a") + "hello") {
        case "":
            if ((((1.5 / (1.5 + -2.25)) > 3.14159) && false) && ("hello" <= (("\n\t" + "hello") + "foo bar"))) {
                var v19 uint = v1
                _ = v19
            } else {
                var v20 uint32 = 4294967295
                var v21 uint = (v1 >> ((v1) & 63))
                switch (fi9 & (fi9 * fi9)) {
                case 1:
                    v22 := GenS4{F0: 1}
                    _ = v22
                case 4294967295:
                    v4 = append(v4, v20)
                case -139:
                    if (true && (261 > -2147483648)) {
                        v2["hello"] = (669 <= 42)
                        switch fi9 {
                        case 785:
                            v23 := GenS5{F0: -899, F1: 1000}
                            _ = v23
                        default:
                            _ = len(v2)
                        }
                        var v24 uint8 = 128
                        _ = v24
                    } else {
                        if ("a" >= ("\n\t" + ("" + ("\n\t" + ("\n\t" + "foo bar"))))) {
                            v25 := GenS4{F0: 9223372036854775807}
                            for fi26 := 0; fi26 < 15; fi26++ {
                                v27 := GenS2{F0: 100}
                                if false {
                                    var v28 float64 = 0.1
                                    switch fi9 {
                                    case -9223372036854775808:
                                        var v29 bool = true
                                        _ = v29
                                    case -2147483648:
                                        v10 += 1.5
                                    case 460:
                                        var v30 uint32 = 255
                                        _ = v30
                                    default:
                                        if ((18446744073709551615 != 948) || (((((("\n\t" + (("\n\t" + "hello") + "foo bar")) + ("a" + ("hello" + ""))) + "") + ("foo bar" + "a")) + "foo bar") <= ("" + "foo bar"))) {
                                            switch "a" {
                                            case "hello":
                                                switch ("hello" + ("a" + ("" + "hello"))) {
                                                case "\n\t":
                                                    var v31 int32 = 2147483647
                                                    _ = v31
                                                }
                                            case "a":
                                                var v32 uint64 = 18446744073709551615
                                                _ = v32
                                            default:
                                                switch (("hello" + ("a" + ("foo bar" + ""))) + ("\n\t" + "a")) {
                                                case "foo bar":
                                                    v21 -= (^(v1))
                                                    v28 -= v28
                                                    fmt.Println(v21, v28)
                                                default:
                                                    switch "foo bar" {
                                                    case "":
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(v1, v21, v1)
                                                        default:
                                                            if (((0 > v20) || false) || (false || (false || (!false)))) {
                                                                var v33 int = (fi9 << (((v1 - (^(v1)))) & 63))
                                                                _ = v33
                                                            }
                                                            switch (fi9 | -9223372036854775808) {
                                                            case 188:
                                                                var v34 int16 = -586
                                                                _ = v34
                                                            case 32768:
                                                                v20 = v20
                                                            }
                                                            v3 = append(v3, fi26)
                                                        }
                                                    case "a":
                                                        if (18446744073709551615 >= 307) {
                                                            if true {
                                                                if true {
                                                                    v28 += v28
                                                                    var v35 int64 = 9223372036854775807
                                                                    if false {
                                                                        var v36 uint16 = 32767
                                                                        if (int(v1) <= fi9) {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            if (-12 < -128) {
                                                                                v35 += (v35 | (v35 * -521))
                                                                            }
                                                                        }
                                                                        _ = v36
                                                                    }
                                                                    var v37 string = "a"
                                                                    if true {
                                                                        v35 = v35
                                                                    }
                                                                    _ = v35
                                                                    _ = v37
                                                                }
                                                                v1 = v21
                                                                fmt.Println(v1)
                                                            }
                                                            v28 -= ((v28 * v10) / v28)
                                                            fmt.Println(v10, fi26)
                                                            var v38 float64 = v28
                                                            if true {
                                                                v39 := GenS4{F0: 42}
                                                                var v40 uint64 = 1
                                                                v1 = (v1 - 4294967295)
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                            if false {
                                                                v41 := GenS5{F0: 0, F1: 42}
                                                                v42 := GenS4{F0: -509}
                                                                v38 += v10
                                                                fmt.Println(v38)
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            _ = v38
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v43 bool = true
                                                            _ = v43
                                                        }
                                                    default:
                                                        var v44 string = "hello"
                                                        _ = v44
                                                    }
                                                }
                                            }
                                            var v45 string = "hello"
                                            fmt.Println(fi9, v28, v45)
                                            _ = v45
                                        }
                                        _ = len(v2)
                                        v21 += 42
                                    }
                                    _ = v28
                                } else {
                                    switch fi26 {
                                    case -166:
                                        if (!true) {
                                            var v46 int16 = 1000
                                            _ = v46
                                        }
                                    case 0:
                                        fmt.Println(v21, v20, v21)
                                    default:
                                        var v47 int16 = -649
                                        _ = v47
                                    }
                                }
                                var v48 uint8 = 0
                                var v49 int8 = 94
                                if (true && (!true)) {
                                    fmt.Println(v20)
                                    _ = len(v4)
                                } else {
                                    var v50 int64 = int64(v20)
                                    var v51 int = fi26
                                    var v52 uint16 = 32767
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                }
                                _ = fi26
                                _ = v27
                                _ = v48
                                _ = v49
                            }
                            var v53 float64 = float64(float32(float64(float32(0.0))))
                            _ = v25
                            _ = v53
                        } else {
                            v21 += (v1 * v1)
                            switch (fi9 - fi9) {
                            case 9223372036854775807:
                                if (3.14159 >= v10) {
                                    var v54 int = 1
                                    v55 := v2["hello"]
                                    switch ("hello" + "foo bar") {
                                    case "\n\t":
                                        v54 = 42
                                    case "foo bar":
                                        for fi56 := 0; fi56 < 7; fi56++ {
                                            switch (^(fi9)) {
                                            case 1:
                                                if (v1 > v1) {
                                                    if (v55 || v55) {
                                                        switch fi9 {
                                                        case 2147483647:
                                                            var v57 uint = v21
                                                            _ = v57
                                                        case 415:
                                                            fmt.Println(v1, v1)
                                                        case -1:
                                                            v2["a"] = ((fi56 >> (((v1 ^ v1)) & 63)) > 32768)
                                                        default:
                                                            v58 := GenS1{F0: 255, F1: 127}
                                                            _ = v58
                                                        }
                                                    } else {
                                                        v2["a"] = ((!(128 > 34)) || false)
                                                        v59 := GenS6{F0: -925}
                                                        _ = v59
                                                    }
                                                    if v55 {
                                                        if (413 <= 0) {
                                                            v60 := GenS4{F0: 1}
                                                            v61 := GenS7{F0: 32767}
                                                            var v62 float64 = ((v10 + (3.14159 + 0.1)) / 1.0)
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            if (754 != -1) {
                                                                v63 := GenS8{F0: 0, F1: -260}
                                                                _ = v63
                                                            }
                                                            v54 = fi56
                                                            var v64 uint8 = 11
                                                            v65 := GenS9{F0: v55, F1: 2147483647, F2: (v54 << ((v1) & 63))}
                                                            if (((v10 / v10) / v10) >= v10) {
                                                                var v66 int8 = 0
                                                                if (uint64(v20) > 4294967295) {
                                                                    if v55 {
                                                                        if (-2.25 != (1e10 - float32(v10))) {
                                                                            var v67 uint8 = 57
                                                                            _ = v67
                                                                        } else {
                                                                            if v55 {
                                                                                v68 := GenS10{F0: 42, F1: 1, F2: v64}
                                                                                switch v54 {
                                                                                case -9223372036854775808:
                                                                                    v3 = append(v3, 9223372036854775807)
                                                                                case 127:
                                                                                    switch "" {
                                                                                    case "":
                                                                                        if ((v66 <= v66) || false) {
                                                                                            if (1.5 >= (0.1 - 0.1)) {
                                                                                                if (0 > 30) {
                                                                                                    fmt.Println(v10, fi56)
                                                                                                }
                                                                                                switch (-(fi9)) {
                                                                                                case 1:
                                                                                                    v64 -= (v64 >> ((v21) & 63))
                                                                                                case 9223372036854775807:
                                                                                                    v10 += (1e10 * v10)
                                                                                                default:
                                                                                                    v66 -= -127
                                                                                                }
                                                                                            } else {
                                                                                                var v69 int = fi56
                                                                                                _ = v69
                                                                                            }
                                                                                            var v70 int16 = -996
                                                                                            _ = v70
                                                                                        } else {
                                                                                            var v71 uint8 = (v64 - uint8(v1))
                                                                                            var v72 uint64 = 2
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = v54
                                                                                            }
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                        }
                                                                                    default:
                                                                                        var v73 float64 = ((((-2.25 - ((v10 - 1e10) / (((3.14159 + (0.0 - 0.0)) - (v10 - float64(3.14159))) + 3.14159))) - float64(3.14159)) / ((v10 + v10) * 1.5)) + v10)
                                                                                        fmt.Println(fi9, v64, v21)
                                                                                        _ = v73
                                                                                    }
                                                                                }
                                                                                v10 += 1.5
                                                                                _ = v68
                                                                            } else {
                                                                                fmt.Println(v66)
                                                                            }
                                                                            if v55 {
                                                                                fmt.Println(v10, fi56)
                                                                                fmt.Println(v20, fi56, v66)
                                                                            } else {
                                                                                var v74 uint8 = v64
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                        switch (("" + ("foo bar" + (((((("" + "hello") + "foo bar") + "") + "\n\t") + ("foo bar" + "\n\t")) + ("" + "a")))) + "foo bar") {
                                                                        case "foo bar":
                                                                            if v55 {
                                                                                v10 = v10
                                                                                _ = len(v2)
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 966
                                                                            }
                                                                            switch v54 {
                                                                            case 1:
                                                                                if v55 {
                                                                                    v75 := GenS10{F0: 2, F1: 670, F2: (^(v64))}
                                                                                    _ = v75
                                                                                }
                                                                            case 0:
                                                                                if (v55 || v55) {
                                                                                    v76 := GenS2{F0: (v66 + v66)}
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v20 = 1
                                                                        case "\n\t":
                                                                            var v77 uint32 = 42
                                                                            _ = v77
                                                                        }
                                                                        if v55 {
                                                                            fmt.Println(v21, v1)
                                                                            if (!v55) {
                                                                                switch int(v64) {
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v21, v20)
                                                                                case 366:
                                                                                    v21 += (v1 + v1)
                                                                                case 127:
                                                                                    v54 = int(v20)
                                                                                    if (float64(((float32((((1.5 + v10) * (v10 * 1.5)) + (v10 - (3.14159 * v10)))) + (3.14159 * 0.1)) * 0.1)) == 1e10) {
                                                                                        v10 = (v10 - (((v10 - (v10 / v10)) / v10) + 1e10))
                                                                                        delete(v2, "foo bar")
                                                                                        switch ("hello" + "hello") {
                                                                                        case "foo bar":
                                                                                            fmt.Println(fi56)
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            var v78 float64 = v10
                                                                                            _ = v78
                                                                                        case "\n\t":
                                                                                            switch (v54 >> ((42) & 63)) {
                                                                                            case 32768:
                                                                                                v55 = (true && (!v55))
                                                                                            case 4294967295:
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = 1
                                                                                                }
                                                                                                var v79 float32 = 1e10
                                                                                                v80 := v2[""]
                                                                                                _ = v79
                                                                                                _ = v80
                                                                                            case 32767:
                                                                                                var v81 uint32 = v20
                                                                                                _ = v81
                                                                                            default:
                                                                                                delete(v2, "\n\t")
                                                                                            }
                                                                                        case "hello":
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    } else {
                                                                                        v82 := GenS4{F0: -248}
                                                                                        v83 := GenS11{F0: v64, F1: (^(v20)), F2: (^(v20))}
                                                                                        _ = v82
                                                                                        _ = v83
                                                                                    }
                                                                                default:
                                                                                    var v84 int64 = 42
                                                                                    _ = v84
                                                                                }
                                                                                v54 += -2147483648
                                                                            }
                                                                            v1 += (^(v21))
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v21, fi56, v20)
                                                                        if false {
                                                                            fmt.Println(v66, v20)
                                                                            _ = len(v3)
                                                                            fmt.Println(v64, fi9)
                                                                            if (((1 >= 42) && false) && v55) {
                                                                                fmt.Println(v55)
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                            var v85 float32 = (1.5 * 1e10)
                                                                            v86 := GenS9{F0: false, F1: -726, F2: fi56}
                                                                            if v55 {
                                                                                v20 = v20
                                                                            } else {
                                                                                v10 += float64(0.0)
                                                                                if false {
                                                                                    switch fi56 {
                                                                                    case 128:
                                                                                        var v87 float64 = 1.5
                                                                                        _ = v87
                                                                                    case 9223372036854775807:
                                                                                        if true {
                                                                                            _ = len(v2)
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                var v88 uint = (v1 | v21)
                                                                                                _ = v88
                                                                                            case "a":
                                                                                                var v89 uint64 = 937
                                                                                                _ = v89
                                                                                            case "hello":
                                                                                                v1 = (v21 << ((34) & 63))
                                                                                            default:
                                                                                                v64 -= v64
                                                                                            }
                                                                                        } else {
                                                                                            switch v54 {
                                                                                            case -2147483648:
                                                                                                v10 -= v10
                                                                                            case 127:
                                                                                                v10 = (((1.5 / float64(((v85 * (0.0 / (1.5 - float32(1e10)))) / (3.14159 / (1.5 * 1.5))))) + float64(-2.25)) * ((v10 - ((float64(3.14159) - -2.25) / v10)) - float64(1.5)))
                                                                                            }
                                                                                        }
                                                                                    case 32767:
                                                                                        var v90 uint64 = 42
                                                                                        _ = v90
                                                                                    }
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            fmt.Println(fi56, v64)
                                                                                        }
                                                                                        v54 = (v54 * (-(fi9)))
                                                                                        if ("hello" > (("a" + "hello") + ("hello" + ""))) {
                                                                                            v91 := GenS8{F0: 42, F1: -1}
                                                                                            _ = v91
                                                                                        }
                                                                                    }
                                                                                    var v92 int8 = v66
                                                                                    v1 = 0
                                                                                    _ = v92
                                                                                } else {
                                                                                    if (v10 < (float64(((float32(v10) + (0.0 * ((-2.25 * 3.14159) * 1e10))) * (float32(0.0) - v85))) * v10)) {
                                                                                        if true {
                                                                                            v54 = 9223372036854775807
                                                                                            v93 := GenS12{F0: 2, F1: v55, F2: 1}
                                                                                            switch v54 {
                                                                                            case 9223372036854775807:
                                                                                                if ((v1 >> (((v21 * (^(v1)))) & 63)) == uint(v64)) {
                                                                                                    var v94 uint8 = (v64 >> (((^(v21))) & 63))
                                                                                                    switch "\n\t" {
                                                                                                    case "a":
                                                                                                        var v95 uint32 = 2
                                                                                                        _ = v95
                                                                                                    case "\n\t":
                                                                                                        if (v55 || (!false)) {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            v96 := GenS13{F0: v20, F1: -1, F2: (v66 >> ((v1) & 63))}
                                                                                                            switch ("a" + "\n\t") {
                                                                                                            case "\n\t":
                                                                                                                switch fi56 {
                                                                                                                case 87:
                                                                                                                    var v97 uint = v1
                                                                                                                    _ = v97
                                                                                                                default:
                                                                                                                    v55 = ((("" + "hello") + "\n\t") < "a")
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                if v55 {
                                                                                                                    var v98 int8 = v66
                                                                                                                    var v99 int16 = -800
                                                                                                                    v98 = v98
                                                                                                                    fmt.Println(v20)
                                                                                                                    _ = v98
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                            case "":
                                                                                                                var v100 float64 = (v10 / v10)
                                                                                                                _ = v100
                                                                                                            default:
                                                                                                                var v101 float32 = 0.1
                                                                                                                _ = v101
                                                                                                            }
                                                                                                            _ = v96
                                                                                                        } else {
                                                                                                            v2["hello"] = (!false)
                                                                                                            var v102 int64 = -723
                                                                                                            switch "a" {
                                                                                                            case "a":
                                                                                                                var v103 uint64 = uint64(v94)
                                                                                                                _ = v103
                                                                                                            case "foo bar":
                                                                                                                v104 := GenS3{F0: fi9, F1: v94, F2: (^(v21))}
                                                                                                                _ = v104
                                                                                                            }
                                                                                                            v1 = v1
                                                                                                            if (!(((1 != 42) || v55) && (((^(v94)) >= (v94 | (v64 - (v64 << (((^(v1))) & 63))))) || v55))) {
                                                                                                                _ = len(v2)
                                                                                                                var v105 float64 = 3.14159
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                                _ = v105
                                                                                                            }
                                                                                                            _ = v102
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v94
                                                                                                }
                                                                                            }
                                                                                            v106 := GenS12{F0: 4294967295, F1: v55, F2: 1000}
                                                                                            v85 -= v85
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (^(fi9))
                                                                                            }
                                                                                            _ = v106
                                                                                            _ = v93
                                                                                        }
                                                                                    }
                                                                                    v21 = 42
                                                                                }
                                                                                v21 += v1
                                                                                v55 = v55
                                                                                _ = len(v2)
                                                                                var v107 uint64 = 455
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    if (v107 < (^(v107))) {
                                                                                        switch -1 {
                                                                                        case 4294967295:
                                                                                            var v108 uint = (v21 - v1)
                                                                                            _ = v108
                                                                                        default:
                                                                                            switch v54 {
                                                                                            case 479:
                                                                                                v107 -= 4294967295
                                                                                            case -1:
                                                                                                var v109 uint = v1
                                                                                                _ = v109
                                                                                            case -9223372036854775808:
                                                                                                v110 := GenS6{F0: 0}
                                                                                                _ = v110
                                                                                            default:
                                                                                                v21 = v1
                                                                                            }
                                                                                            v111 := GenS6{F0: 1000}
                                                                                            v85 = v85
                                                                                            _ = v111
                                                                                        }
                                                                                        var v112 uint32 = (v20 % ((v20) | 1))
                                                                                        var v113 int = (^(fi56))
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                default:
                                                                                    if (v55 || (!false)) {
                                                                                        v114 := GenS1{F0: (v64 >> ((v21) & 63)), F1: (v64 * v64)}
                                                                                        _ = v114
                                                                                    } else {
                                                                                        fmt.Println(fi56, v10, v55)
                                                                                        var v115 uint64 = v107
                                                                                        v116 := GenS14{F0: ("a" + "\n\t"), F1: fi9, F2: (v55 || (v64 == uint8(v1)))}
                                                                                        var v117 float64 = v10
                                                                                        var v118 int16 = 1000
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        fmt.Println(v20, v54)
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                    }
                                                                                }
                                                                                _ = v107
                                                                            }
                                                                            _ = v85
                                                                            _ = v86
                                                                        } else {
                                                                            var v119 uint = (v1 / ((v21) | 1))
                                                                            v21 -= (v119 % ((425) | 1))
                                                                            _ = v119
                                                                        }
                                                                        var v120 uint = v1
                                                                        fmt.Println(v66, v1, v10)
                                                                        _ = v120
                                                                    }
                                                                    if v55 {
                                                                        v21 = (^(v1))
                                                                        if false {
                                                                            var v121 float64 = v10
                                                                            v66 = v66
                                                                            if (!(!(4294967295 >= 654))) {
                                                                                fmt.Println(v10, v20)
                                                                                v20 = (v20 - v20)
                                                                                if true {
                                                                                    v121 += (0.1 + 1e10)
                                                                                    var v122 int = (fi56 % ((32767) | 1))
                                                                                    _ = v122
                                                                                } else {
                                                                                    if v55 {
                                                                                        v54 -= v54
                                                                                        _ = len(v3)
                                                                                    } else {
                                                                                        if ((!((!((v64 | (v64 << (((v21 / ((uint(v64)) | 1))) & 63))) == 1)) || v55)) || (v10 <= v10)) {
                                                                                            var v123 int8 = (v66 % ((v66) | 1))
                                                                                            if (uint8(v21) < 1) {
                                                                                                delete(v2, (("a" + ((("" + "a") + "\n\t") + (("hello" + ("hello" + "a")) + ("a" + "foo bar")))) + ("\n\t" + "foo bar")))
                                                                                            }
                                                                                            _ = len(v3)
                                                                                            _ = v123
                                                                                        } else {
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                v124 := GenS10{F0: 65535, F1: 1, F2: 2}
                                                                                                _ = v124
                                                                                            case "":
                                                                                                v125 := GenS15{F0: 1, F1: -552, F2: 0}
                                                                                                _ = v125
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v121
                                                                        }
                                                                    } else {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            v3 = append(v3, (^(fi9)))
                                                                        case "":
                                                                            fmt.Println(v66, v66)
                                                                        case "\n\t":
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                delete(v2, "")
                                                                            }
                                                                        }
                                                                    }
                                                                    v126 := GenS12{F0: 42, F1: (v55 || (false || true)), F2: 1000}
                                                                    _ = v126
                                                                }
                                                                _ = v66
                                                            } else {
                                                                fmt.Println(v20, v20)
                                                                fmt.Println(fi9)
                                                            }
                                                            v127 := GenS16{F0: 127}
                                                            _ = v127
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        var v128 int64 = 0
                                                        if (255 == 1) {
                                                            if v55 {
                                                                if v55 {
                                                                    var v129 string = ""
                                                                    v20 -= (^(v20))
                                                                    _ = len(v2)
                                                                    if (-2147483648 >= -2147483648) {
                                                                        v130 := GenS7{F0: 42}
                                                                        var v131 uint8 = 255
                                                                        _ = v130
                                                                        _ = v131
                                                                    }
                                                                    _ = v129
                                                                } else {
                                                                    fmt.Println(fi56, v20)
                                                                    if (float64(float32(0.0)) > 0.1) {
                                                                        var v132 int64 = 42
                                                                        if v55 {
                                                                            var v133 int8 = 100
                                                                            _ = v133
                                                                        }
                                                                        _ = v132
                                                                    } else {
                                                                        v55 = false
                                                                        v2["foo bar"] = (true || v55)
                                                                    }
                                                                    var v134 int32 = 969
                                                                    _ = v134
                                                                }
                                                                v10 -= v10
                                                                v135 := GenS17{F0: ((^(v20)) >= (v20 + v20)), F1: (v20 + v20), F2: "hello"}
                                                                _ = v135
                                                            } else {
                                                                v54 += v54
                                                                if v55 {
                                                                    v21 = (^(v1))
                                                                    if ((1 > (v21 * 18446744073709551615)) || false) {
                                                                        var v136 uint16 = 937
                                                                        var v137 int8 = -1
                                                                        v136 -= v136
                                                                        _ = v136
                                                                        _ = v137
                                                                    } else {
                                                                        v128 += v128
                                                                    }
                                                                    if v55 {
                                                                        var v138 uint16 = 42
                                                                        fmt.Println(v21)
                                                                        _ = v138
                                                                    }
                                                                }
                                                                if (-2.25 <= 0.0) {
                                                                    fmt.Println(v20, v21, fi9)
                                                                    var v139 float32 = (-2.25 / -2.25)
                                                                    switch 9223372036854775807 {
                                                                    case 9223372036854775807:
                                                                        v140 := v2[(("foo bar" + "") + "hello")]
                                                                        _ = v140
                                                                    case -763:
                                                                        fmt.Println(v128, v1, v55)
                                                                        v141 := GenS14{F0: "hello", F1: (-(v54)), F2: v55}
                                                                        switch fi9 {
                                                                        case -611:
                                                                            v142 := GenS18{F0: 127, F1: (v128 % ((-562) | 1)), F2: 42}
                                                                            _ = v142
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v55, v10, v21)
                                                                        case -717:
                                                                            v20 = v20
                                                                        }
                                                                        _ = v141
                                                                    default:
                                                                        switch 32767 {
                                                                        case 32768:
                                                                            var v143 uint8 = 0
                                                                            _ = v143
                                                                        case -330:
                                                                            if v55 {
                                                                                var v144 int32 = 485
                                                                                v21 = v21
                                                                                _ = len(v4)
                                                                                fmt.Println(v1, v10, v10)
                                                                                _ = len(v4)
                                                                                _ = v144
                                                                            } else {
                                                                                if (128 > 255) {
                                                                                    if (!v55) {
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            v145 := GenS5{F0: 0, F1: 1}
                                                                                            if (65535 == 42) {
                                                                                                v146 := GenS1{F0: 0, F1: 1}
                                                                                                _ = v146
                                                                                            }
                                                                                            _ = len(v4)
                                                                                            _ = v145
                                                                                        case "\n\t":
                                                                                            var v147 int = -452
                                                                                            _ = v147
                                                                                        default:
                                                                                            v139 = float32(3.14159)
                                                                                            if false {
                                                                                                if (417 >= 42) {
                                                                                                    v20 += (v20 * v20)
                                                                                                    _ = len(v2)
                                                                                                    v20 += v20
                                                                                                } else {
                                                                                                    v21 = (v1 * (v1 - v1))
                                                                                                    var v148 int16 = -1
                                                                                                    var v149 bool = false
                                                                                                    v150 := GenS11{F0: 0, F1: v20, F2: v20}
                                                                                                    v20 -= v20
                                                                                                    _ = v148
                                                                                                    _ = v149
                                                                                                    _ = v150
                                                                                                }
                                                                                                switch (fi9 | fi9) {
                                                                                                case -9223372036854775808:
                                                                                                    if v55 {
                                                                                                        v139 = float32(1.5)
                                                                                                        if (!v55) {
                                                                                                            if (0 > 32767) {
                                                                                                                v54 = fi9
                                                                                                            }
                                                                                                            var v151 uint16 = 382
                                                                                                            _ = v151
                                                                                                        }
                                                                                                        var v152 string = "foo bar"
                                                                                                        _ = v152
                                                                                                    } else {
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        delete(v2, "")
                                                                                                        v153 := GenS5{F0: 453, F1: -32768}
                                                                                                        _ = v153
                                                                                                    }
                                                                                                case -1:
                                                                                                    if false {
                                                                                                        v1 = 502
                                                                                                        v154 := GenS19{F0: v55}
                                                                                                        var v155 bool = v55
                                                                                                        _ = v154
                                                                                                        _ = v155
                                                                                                    } else {
                                                                                                        var v156 int64 = v128
                                                                                                        _ = v156
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v54 += 4294967295
                                                                                        }
                                                                                        var v157 bool = (494 <= 167)
                                                                                        _ = v157
                                                                                    }
                                                                                    v139 += float32(v10)
                                                                                    var v158 int = (v54 | v54)
                                                                                    _ = v158
                                                                                } else {
                                                                                    var v159 string = ""
                                                                                    v160 := GenS13{F0: (v20 ^ (v20 & (v20 * v20))), F1: 1, F2: -51}
                                                                                    v20 -= (v20 % ((v20) | 1))
                                                                                    v161 := GenS6{F0: 32767}
                                                                                    _ = v159
                                                                                    _ = v160
                                                                                    _ = v161
                                                                                }
                                                                                fmt.Println(v1, v21, v128)
                                                                                fmt.Println(v139, fi9, v128)
                                                                            }
                                                                            fmt.Println(v10, fi9, fi9)
                                                                            v21 = 65535
                                                                        default:
                                                                            if (248 != 255) {
                                                                                v162 := GenS15{F0: 1, F1: (v128 >> (((v21 << ((0) & 63))) & 63)), F2: 289}
                                                                                fmt.Println(fi56, v10, v139)
                                                                                v54 = 42
                                                                                _ = v162
                                                                            } else {
                                                                                v163 := GenS15{F0: 32767, F1: (^(v128)), F2: 2}
                                                                                _ = v163
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v139
                                                                }
                                                            }
                                                        }
                                                        _ = v128
                                                    } else {
                                                        if v55 {
                                                            v55 = v55
                                                            if true {
                                                                var v164 uint16 = 65535
                                                                if v55 {
                                                                    if v55 {
                                                                        switch fi56 {
                                                                        case 4294967295:
                                                                            if true {
                                                                                _ = len(v2)
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    fmt.Println(v1)
                                                                                case "\n\t":
                                                                                    v165 := GenS5{F0: int64(fi9), F1: 0}
                                                                                    _ = v165
                                                                                case "":
                                                                                    v1 = v1
                                                                                }
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    v166 := GenS19{F0: ((v164 * 583) <= uint16(v1))}
                                                                                    var v167 int16 = 42
                                                                                    _ = v166
                                                                                    _ = v167
                                                                                case "\n\t":
                                                                                    v168 := GenS20{F0: v54}
                                                                                    _ = v168
                                                                                default:
                                                                                    v20 += uint32(v164)
                                                                                }
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    if false {
                                                                                        switch -9223372036854775808 {
                                                                                        case -1:
                                                                                            var v169 uint32 = v20
                                                                                            _ = v169
                                                                                        case 754:
                                                                                            var v170 int32 = 0
                                                                                            _ = v170
                                                                                        default:
                                                                                            fmt.Println(v20, v20)
                                                                                        }
                                                                                        var v171 float64 = ((0.0 * 1e10) + 0.1)
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            v172 := GenS11{F0: 42, F1: (^(v20)), F2: v20}
                                                                                            _ = v172
                                                                                        default:
                                                                                            if ((127 < 128) || true) {
                                                                                                v54 = 128
                                                                                            }
                                                                                        }
                                                                                        _ = v171
                                                                                    } else {
                                                                                        v20 = v20
                                                                                    }
                                                                                default:
                                                                                    v55 = (0 >= 100)
                                                                                    v164 -= 65535
                                                                                    switch (fi9 + v54) {
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(fi56)
                                                                                    case 0:
                                                                                        fmt.Println(v55, fi56)
                                                                                    default:
                                                                                        fmt.Println(v54, fi56, v20)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v21)
                                                                                v173 := GenS6{F0: 1}
                                                                                _ = v173
                                                                            }
                                                                        default:
                                                                            switch fi56 {
                                                                            case 9223372036854775807:
                                                                                v174 := GenS7{F0: (v164 % ((v164) | 1))}
                                                                                _ = v174
                                                                            case -415:
                                                                                v175 := GenS21{F0: (v1 / (((v1 / (((v1 << ((v21) & 63))) | 1))) | 1)), F1: 42}
                                                                                _ = v175
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v55, v21, v20)
                                                                    var v176 uint16 = (v164 ^ 42)
                                                                    _ = v176
                                                                } else {
                                                                    v4 = append(v4, 321)
                                                                    var v177 float64 = v10
                                                                    v164 += (v164 ^ (v164 << (((^(v1))) & 63)))
                                                                    _ = v177
                                                                }
                                                                v20 = 1
                                                                _ = v164
                                                            }
                                                            switch (v54 >> (((v1 * v21)) & 63)) {
                                                            case 9223372036854775807:
                                                                var v178 float64 = float64(0.0)
                                                                _ = v178
                                                            case 127:
                                                                switch (("\n\t" + "a") + ("" + ("\n\t" + "a"))) {
                                                                case "\n\t":
                                                                    var v179 int16 = -32768
                                                                    _ = v179
                                                                default:
                                                                    switch fi9 {
                                                                    case 0:
                                                                        fmt.Println(fi56)
                                                                    case -1:
                                                                        v21 = 4294967295
                                                                    case 4294967295:
                                                                        var v180 uint8 = 42
                                                                        _ = v180
                                                                    default:
                                                                        v181 := GenS5{F0: int64(v21), F1: 32767}
                                                                        var v182 float64 = 0.1
                                                                        v20 += (v20 << ((v1) & 63))
                                                                        _ = v181
                                                                        _ = v182
                                                                    }
                                                                }
                                                            case 32768:
                                                                var v183 float32 = -2.25
                                                                _ = v183
                                                            default:
                                                                var v184 float64 = v10
                                                                _ = v184
                                                            }
                                                        } else {
                                                            v54 += -509
                                                            v54 = (v54 & fi9)
                                                            _ = len(v4)
                                                            v10 = v10
                                                            if (42 < -2147483648) {
                                                                if (2 != 18446744073709551615) {
                                                                    var v185 uint64 = 0
                                                                    _ = v185
                                                                }
                                                                v186 := GenS9{F0: v55, F1: -2147483648, F2: fi56}
                                                                v21 = 0
                                                                _ = v186
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                if v55 {
                                                    if v55 {
                                                        var v187 int16 = 1000
                                                        fmt.Println(v1, v21)
                                                        _ = v187
                                                    } else {
                                                        v188 := GenS22{F0: 255, F1: v21}
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v1 -= 5
                                                        case "hello":
                                                            _ = len(v4)
                                                        }
                                                        var v189 bool = false
                                                        _ = v188
                                                        _ = v189
                                                    }
                                                    var v190 int32 = -2147483648
                                                    _ = v190
                                                } else {
                                                    if (18446744073709551615 == 18446744073709551615) {
                                                        v191 := GenS12{F0: 605, F1: v55, F2: 1000}
                                                        v20 = v20
                                                        var v192 int8 = 1
                                                        if v55 {
                                                            v21 += 65535
                                                            v192 = (^(v192))
                                                        }
                                                        if ((v21 >> ((2) & 63)) > (v1 / ((4294967295) | 1))) {
                                                            fmt.Println(v55, fi9)
                                                            v55 = (0.0 < 1.5)
                                                            _ = len(v4)
                                                        } else {
                                                            if (1 < 1) {
                                                                v193 := v2["a"]
                                                                _ = v193
                                                            } else {
                                                                v194 := GenS14{F0: ("a" + ("hello" + "a")), F1: (fi56 >> ((v1) & 63)), F2: (127 >= 100)}
                                                                _ = v194
                                                            }
                                                        }
                                                        _ = v191
                                                        _ = v192
                                                    } else {
                                                        v195 := GenS4{F0: 0}
                                                        _ = v195
                                                    }
                                                    v196 := GenS21{F0: 255, F1: -9223372036854775808}
                                                    _ = v196
                                                }
                                            }
                                            if (42 > 659) {
                                                v21 = v1
                                                v10 += (float64((-2.25 + 3.14159)) + v10)
                                                fmt.Println(fi56)
                                            }
                                            if (!true) {
                                                var v197 uint16 = uint16(v20)
                                                _ = v197
                                            } else {
                                                fmt.Println(v1, fi9)
                                            }
                                            _ = fi56
                                        }
                                    }
                                    _ = v54
                                    _ = v55
                                } else {
                                    v10 += 0.1
                                }
                            case 42:
                                var v198 uint32 = v20
                                _ = v198
                            default:
                                _ = len(v3)
                            }
                            if true {
                                if false {
                                    v3 = append(v3, fi9)
                                }
                            } else {
                                var v199 string = "\n\t"
                                _ = v199
                            }
                        }
                        v21 = 4294967295
                        if (-406 < 978) {
                            var v200 float64 = 1.5
                            v201 := GenS11{F0: 255, F1: (v20 * v20), F2: v20}
                            _ = v200
                            _ = v201
                        } else {
                            var v202 int = (fi9 | fi9)
                            _ = v202
                        }
                    }
                }
                _ = v20
                _ = v21
            }
        case "\n\t":
            var v203 uint16 = 1
            _ = v203
        case "hello":
            v1 = (v1 / (((v1 ^ (v1 | 1))) | 1))
        default:
            switch fi9 {
            case 127:
                fmt.Println(fi9, v10)
            case -1:
                var v204 int64 = -441
                _ = v204
            case 0:
                v10 = ((v10 - ((1.5 - (v10 + float64(1.5))) - v10)) / 3.14159)
            }
        }
        _ = fi9
        _ = v10
    }
    v1 -= 554
    var v205 int = 245
    v205 = (v205 + (v205 - -52))
    v206 := GenS17{F0: (!true), F1: 4294967295, F2: "foo bar"}
    delete(v2, ((("a" + ("foo bar" + "foo bar")) + ("\n\t" + "")) + (("foo bar" + (("foo bar" + "a") + "")) + ("" + "foo bar"))))
    var v207 int16 = 42
    v208 := GenS4{F0: 362}
    v205 = v205
    switch "foo bar" {
    case "hello":
        var v209 int16 = v207
        _ = v209
    case "\n\t":
        fmt.Println(v207)
    case "a":
        _ = len(v2)
    default:
        var v210 float64 = -2.25
        _ = v210
    }
    v1 += (v1 & 18446744073709551615)
    for fi211 := 0; fi211 < 9; fi211++ {
        var v212 int = (fi211 * (-(fi211)))
        for fi213 := 0; fi213 < 12; fi213++ {
            v207 -= (v207 - (-(v207)))
            if false {
                if (319 != -9223372036854775808) {
                    var v214 int = (v212 ^ v212)
                    _ = v214
                } else {
                    var v215 uint64 = 345
                    v216 := GenS23{F0: -42, F1: 65535, F2: 1}
                    if (94 <= 117) {
                        if false {
                            var v217 string = ("\n\t" + "")
                            if false {
                                v218 := GenS24{F0: -18}
                                v219 := GenS15{F0: 1, F1: 627, F2: 2}
                                _ = v218
                                _ = v219
                            } else {
                                v1 -= (v1 << (((^(v1))) & 63))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v220 int32 = 52
                                _ = v220
                            }
                            _ = v217
                        } else {
                            fmt.Println(v205)
                            var v221 string = ("foo bar" + "foo bar")
                            _ = v221
                        }
                        v222 := v2["\n\t"]
                        v212 -= fi213
                        _ = v222
                    } else {
                        v223 := GenS18{F0: 1, F1: -899, F2: -9223372036854775808}
                        v224 := v2["a"]
                        var v225 float32 = float32((-2.25 * 0.1))
                        _ = v223
                        _ = v224
                        _ = v225
                    }
                    _ = v215
                    _ = v216
                }
                if true {
                    fmt.Println(v205)
                }
                v226 := GenS22{F0: 255, F1: v1}
                _ = v226
            } else {
                var v227 uint16 = 597
                _ = v227
            }
            v228 := GenS13{F0: 2147483648, F1: -2147483648, F2: int8(v1)}
            _ = fi213
            _ = v228
        }
        _ = fi211
        _ = v212
    }
    var v229 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v205:", v205)
    fmt.Println("v207:", v207)
    fmt.Println("v229:", v229)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v205
    _ = v206
    _ = v207
    _ = v208
    _ = v229
    _ = v3
    _ = v4
}
