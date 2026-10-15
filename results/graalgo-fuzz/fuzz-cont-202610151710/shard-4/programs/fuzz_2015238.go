package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int64; F2 int64 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint; F1 uint64; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint; F1 int64 }
type GenS14 struct { F0 uint; F1 int8; F2 int }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int32; F1 int8; F2 uint16 }
type GenS18 struct { F0 int64; F1 uint8; F2 int64 }

func main() {
    var v1 float64 = (0.1 / (1e10 - (((0.0 * 0.0) + 3.14159) - 1.5)))
    var v2 bool = ((v1 * v1) != 3.14159)
    v3 := []string{"", "\n\t", ("foo bar" + "foo bar"), ""}
    var v4 uint32 = 42
    var v5 int64 = 1
    v5 += -9223372036854775808
    v3 = append(v3, ("foo bar" + ("" + ((("a" + "\n\t") + "a") + ("\n\t" + "a")))))
    v4 = 2147483648
    v6 := GenS1{F0: 39}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint = 73
        v9 := GenS1{F0: 42}
        var v10 bool = false
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(v4, v4)
    v11 := GenS2{F0: v5, F1: 48, F2: 42}
    _ = len(v3)
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v1)
        _ = fi12
    }
    if (255 <= 255) {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 2
            _ = fi13
            _ = v14
        }
    }
    if true {
        if v2 {
            if v2 {
                for fi15 := 0; fi15 < 7; fi15++ {
                    if (("\n\t" + "hello") >= "foo bar") {
                        switch "a" {
                        case "a":
                            v2 = false
                        case "hello":
                            for fi16 := 0; fi16 < 3; fi16++ {
                                v1 -= v1
                                switch fi15 {
                                case 241:
                                    fmt.Println(fi16, v2, fi15)
                                }
                                var v17 uint32 = (v4 % ((v4) | 1))
                                fmt.Println(v1)
                                _ = fi16
                                _ = v17
                            }
                        case "":
                            var v18 uint16 = 2
                            _ = v18
                        default:
                            for fi19 := 0; fi19 < 17; fi19++ {
                                v3 = append(v3, ("foo bar" + "foo bar"))
                                _ = fi19
                            }
                        }
                        if (v2 || v2) {
                            var v20 uint = 1
                            fmt.Println(v1, v1, v4)
                            switch fi15 {
                            case 0:
                                var v21 int32 = 289
                                _ = v21
                            }
                            _ = v20
                        }
                    }
                    var v22 uint32 = (v4 - 966)
                    switch (fi15 & 2147483647) {
                    case 1:
                        if v2 {
                            v22 -= v4
                            if (v22 == (^(v22))) {
                                switch (fi15 | (fi15 << ((65535) & 63))) {
                                case 4294967295:
                                    for fi23 := 0; fi23 < 2; fi23++ {
                                        v24 := GenS2{F0: v5, F1: 2, F2: 1000}
                                        v1 -= v1
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                                switch "" {
                                case "a":
                                    v2 = v2
                                case "hello":
                                    if (true && (-430 == 42)) {
                                        switch 9223372036854775807 {
                                        case -1:
                                            if v2 {
                                                for fi25 := 0; fi25 < 19; fi25++ {
                                                    switch fi15 {
                                                    case 0:
                                                        fmt.Println(fi25, fi25, v4)
                                                    case 73:
                                                        switch (fi15 << ((18446744073709551615) & 63)) {
                                                        case 9223372036854775807:
                                                            v26 := GenS3{F0: 1, F1: 532}
                                                            _ = v26
                                                        }
                                                    default:
                                                        var v27 uint64 = 1
                                                        _ = len(v3)
                                                        _ = v27
                                                    }
                                                    _ = fi25
                                                }
                                            }
                                        case 1:
                                            switch fi15 {
                                            case 128:
                                                if (477 >= 1) {
                                                    var v28 uint64 = 199
                                                    var v29 int16 = -369
                                                    v2 = (0.1 != (0.0 * 1.5))
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = (((("foo bar" + "") + "a") + "") + "a")
                                                    }
                                                }
                                            case 42:
                                                if (v2 || v2) {
                                                    var v30 string = "foo bar"
                                                    _ = v30
                                                } else {
                                                    switch ("\n\t" + ("hello" + "\n\t")) {
                                                    case "a":
                                                        var v31 float32 = 1.5
                                                        _ = v31
                                                    case "foo bar":
                                                        var v32 uint64 = 2
                                                        _ = v32
                                                    }
                                                    for fi33 := 0; fi33 < 2; fi33++ {
                                                        if (-997 < (fi33 | fi33)) {
                                                            fmt.Println(v2)
                                                        } else {
                                                            _ = len(v3)
                                                            v34 := GenS3{F0: 316, F1: 1}
                                                            if (v2 && (false || v2)) {
                                                                v22 = v4
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v35 string = "a"
                                                                    _ = v35
                                                                }
                                                                var v36 uint16 = 32767
                                                                _ = v36
                                                            }
                                                            _ = v34
                                                        }
                                                        fmt.Println(fi15, fi15)
                                                        _ = fi33
                                                    }
                                                }
                                            case 127:
                                                var v37 int8 = 9
                                                _ = v37
                                            default:
                                                v1 = (((((0.1 / -2.25) + (1.5 + (v1 / v1))) - float64(0.0)) + 3.14159) * v1)
                                            }
                                        case 463:
                                            for fi38 := 0; fi38 < 10; fi38++ {
                                                if (1 > 870) {
                                                    var v39 uint32 = (v4 * v4)
                                                    _ = v39
                                                }
                                                v1 -= 1e10
                                                var v40 uint = 42
                                                if (2 > (v4 & v4)) {
                                                    switch fi38 {
                                                    case 9223372036854775807:
                                                        v41 := GenS1{F0: 0}
                                                        _ = v41
                                                    default:
                                                        var v42 int64 = v5
                                                        if false {
                                                            if false {
                                                                var v43 int16 = -229
                                                                _ = v43
                                                            }
                                                            v5 = v5
                                                            v5 = v5
                                                        } else {
                                                            var v44 int16 = -32768
                                                            v3 = append(v3, "a")
                                                            v45 := GenS4{F0: (v4 % ((v22) | 1)), F1: (v42 << ((v40) & 63)), F2: -9223372036854775808}
                                                            v46 := GenS3{F0: 2, F1: 4294967295}
                                                            _ = v44
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        _ = v42
                                                    }
                                                    fmt.Println(v40)
                                                    fmt.Println(v5, v40, fi15)
                                                }
                                                v4 = v4
                                                _ = fi38
                                                _ = v40
                                            }
                                        default:
                                            var v47 int64 = v5
                                            _ = v47
                                        }
                                        switch int(v4) {
                                        case 55:
                                            _ = len(v3)
                                        case -9223372036854775808:
                                            fmt.Println(v1)
                                        default:
                                            v2 = (!true)
                                        }
                                    } else {
                                        var v48 float32 = float32(v1)
                                        _ = v48
                                    }
                                case "":
                                    var v49 uint32 = (v22 - (v4 % ((0) | 1)))
                                    _ = v49
                                }
                                v4 = (v22 ^ v4)
                            }
                        } else {
                            for fi50 := 0; fi50 < 11; fi50++ {
                                v5 = (v5 + 731)
                                if v2 {
                                    if v2 {
                                        fmt.Println(v5, v4)
                                        switch (fi50 - (fi15 ^ (fi50 & (-(fi50))))) {
                                        case -1:
                                            fmt.Println(v4, v5)
                                        case 32768:
                                            fmt.Println(v5, v22, fi15)
                                        default:
                                            if ((float64(0.0) - float64(float32(v1))) >= v1) {
                                                if v2 {
                                                    var v51 bool = true
                                                    switch fi15 {
                                                    case -9223372036854775808:
                                                        v22 -= (v22 % (((v22 + v22)) | 1))
                                                    }
                                                    _ = v51
                                                }
                                            } else {
                                                if (-128 == 1) {
                                                    v4 += (^(v4))
                                                    v2 = (((v5 * -9223372036854775808) == v5) || false)
                                                    var v52 float32 = 0.0
                                                    var v53 uint16 = 779
                                                    fmt.Println(fi50, fi50)
                                                    v52 -= v52
                                                    switch fi50 {
                                                    case 703:
                                                        fmt.Println(v52)
                                                    }
                                                    _ = v52
                                                    _ = v53
                                                }
                                                var v54 bool = v2
                                                v4 += (^(v4))
                                                _ = v54
                                            }
                                        }
                                        var v55 float64 = 1e10
                                        v56 := GenS2{F0: v5, F1: 127, F2: 1}
                                        switch (fi50 ^ (^(fi50))) {
                                        case 408:
                                            v4 += (v22 << ((109) & 63))
                                        case -320:
                                            var v57 uint = 65535
                                            fmt.Println(fi50)
                                            _ = v57
                                        default:
                                            switch 32768 {
                                            case 32768:
                                                fmt.Println(v1, v2)
                                                v3 = append(v3, "foo bar")
                                                switch "foo bar" {
                                                case "a":
                                                    var v58 uint8 = 37
                                                    _ = v58
                                                case "hello":
                                                    switch ("hello" + "a") {
                                                    case "":
                                                        if len(v3) > 0 {
                                                            v3[0] = ""
                                                        }
                                                    case "foo bar":
                                                        if (v1 < -2.25) {
                                                            var v59 float32 = 0.0
                                                            fmt.Println(v59, v2, v2)
                                                            var v60 int = (fi50 | (fi15 / ((fi50) | 1)))
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            switch 776 {
                                                            case 923:
                                                                v5 -= (v5 ^ v5)
                                                            default:
                                                                var v61 int64 = 507
                                                                _ = v61
                                                            }
                                                        }
                                                    default:
                                                        if v2 {
                                                            switch fi15 {
                                                            case 736:
                                                                var v62 int16 = 1000
                                                                _ = v62
                                                            case 127:
                                                                if (!v2) {
                                                                    v63 := GenS5{F0: 18, F1: -70, F2: 599}
                                                                    _ = v63
                                                                }
                                                            default:
                                                                v1 -= v1
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                v2 = (!(true && ("hello" > ("" + ("hello" + "\n\t")))))
                                                            case "hello":
                                                                if v2 {
                                                                    var v64 int8 = 127
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ((("hello" + "hello") + "foo bar") + ((("a" + "\n\t") + ("foo bar" + "a")) + "\n\t"))
                                                                    }
                                                                    _ = v64
                                                                }
                                                            case "\n\t":
                                                                v1 = 0.0
                                                                v65 := GenS6{F0: 255, F1: -266}
                                                                var v66 int32 = 42
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            v55 = 1e10
                                                        } else {
                                                            var v67 float32 = (0.1 * (1.5 - (float32(float64(3.14159)) + 0.1)))
                                                            var v68 bool = (243 > 2)
                                                            v22 = (v22 | 2147483647)
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    }
                                                default:
                                                    switch ("foo bar" + ("" + "hello")) {
                                                    case "hello":
                                                        var v69 uint = 761
                                                        _ = v69
                                                    case "foo bar":
                                                        if (true || (true || (1 != 2147483647))) {
                                                            if v2 {
                                                                v4 += (^(v22))
                                                            } else {
                                                                switch ("foo bar" + "") {
                                                                case "foo bar":
                                                                    v70 := GenS7{F0: 4294967295, F1: 34, F2: (fi15 / (((fi50 + fi15)) | 1))}
                                                                    _ = v70
                                                                case "\n\t":
                                                                    if v2 {
                                                                        var v71 uint8 = 128
                                                                        fmt.Println(v2, fi15, v55)
                                                                        switch fi15 {
                                                                        case 519:
                                                                            v71 = 139
                                                                        case 37:
                                                                            var v72 float64 = 1e10
                                                                            _ = v72
                                                                        case -2147483648:
                                                                            fmt.Println(v1, v2, v2)
                                                                            v73 := GenS8{F0: 42}
                                                                            _ = v73
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        if (("foo bar" + "foo bar") == "\n\t") {
                                                                            switch 128 {
                                                                            case -2147483648:
                                                                                var v74 int32 = 1
                                                                                _ = v74
                                                                            case -774:
                                                                                fmt.Println(v1, v1, v55)
                                                                            default:
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v75 := GenS9{F0: 127}
                                                                                    _ = v75
                                                                                case "":
                                                                                    v76 := GenS6{F0: 63, F1: 2147483647}
                                                                                    var v77 bool = v2
                                                                                    var v78 uint = 18446744073709551615
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                case "\n\t":
                                                                                    if v2 {
                                                                                        var v79 uint = 2
                                                                                        if false {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            var v80 uint = (v79 % (((v79 * 0)) | 1))
                                                                                            _ = v80
                                                                                        }
                                                                                        switch (("" + (((((("hello" + ("a" + "foo bar")) + "") + (("foo bar" + (("foo bar" + "a") + "")) + ((("a" + "\n\t") + "foo bar") + "foo bar"))) + "a") + "foo bar") + "a")) + (("" + "a") + "\n\t")) {
                                                                                        case "\n\t":
                                                                                            var v81 uint8 = uint8(v4)
                                                                                            _ = v81
                                                                                        }
                                                                                        switch -424 {
                                                                                        case 128:
                                                                                            fmt.Println(v5)
                                                                                        default:
                                                                                            v82 := GenS2{F0: -1, F1: 1, F2: -1}
                                                                                            _ = v82
                                                                                        }
                                                                                        var v83 int32 = 2147483647
                                                                                        var v84 uint8 = 255
                                                                                        v85 := GenS10{F0: 127}
                                                                                        _ = v79
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    } else {
                                                                                        v5 = (v5 & v5)
                                                                                        v1 = v1
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch (fi50 >> ((4294967295) & 63)) {
                                                                            case 890:
                                                                                switch ("a" + "") {
                                                                                case "hello":
                                                                                    var v86 uint8 = 7
                                                                                    _ = v86
                                                                                case "":
                                                                                    fmt.Println(fi15, v4)
                                                                                }
                                                                            case 661:
                                                                                fmt.Println(fi15, fi15)
                                                                            }
                                                                            if true {
                                                                                var v87 int32 = -2147483648
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        switch (fi50 - (fi50 % ((9223372036854775807) | 1))) {
                                                                                        case 2147483647:
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                var v88 int8 = 1
                                                                                                _ = v88
                                                                                            case "a":
                                                                                                var v89 uint = 65
                                                                                                _ = v89
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v55, fi50)
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v90 uint16 = 2
                                                                                        _ = v90
                                                                                    default:
                                                                                        fmt.Println(v1, v87)
                                                                                    }
                                                                                case "a":
                                                                                    v87 += v87
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ((("hello" + ("" + ((("foo bar" + ((((("" + (("\n\t" + "\n\t") + "a")) + "a") + "foo bar") + "\n\t") + "foo bar")) + "\n\t") + "a"))) + (((("a" + "\n\t") + "hello") + ("a" + "foo bar")) + ("foo bar" + ""))) + "foo bar")
                                                                                        }
                                                                                    case "hello":
                                                                                        switch fi50 {
                                                                                        case 42:
                                                                                            v2 = false
                                                                                        case -864:
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        if v2 {
                                                                                            v3 = append(v3, "a")
                                                                                            v3 = append(v3, "hello")
                                                                                        }
                                                                                    }
                                                                                    switch fi50 {
                                                                                    case 162:
                                                                                        if (v87 <= 110) {
                                                                                            var v91 int16 = int16(v87)
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            v3 = append(v3, "a")
                                                                                            switch (fi50 - fi15) {
                                                                                            case 0:
                                                                                                switch ((("\n\t" + "\n\t") + "a") + ("a" + ("" + "\n\t"))) {
                                                                                                case "":
                                                                                                    v92 := GenS11{F0: 834, F1: 0, F2: 0}
                                                                                                    _ = v92
                                                                                                case "\n\t":
                                                                                                    switch (fi15 | fi50) {
                                                                                                    case 1:
                                                                                                        if true {
                                                                                                            v91 = (v91 * 0)
                                                                                                            v1 = (v55 + 1.5)
                                                                                                            v91 = 0
                                                                                                        }
                                                                                                    case 127:
                                                                                                        var v93 uint16 = 255
                                                                                                        v87 -= v87
                                                                                                        _ = v93
                                                                                                    case 4294967295:
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            v55 -= 1e10
                                                                                                            if (!(v2 && false)) {
                                                                                                                var v94 int32 = (^(v87))
                                                                                                                _ = v94
                                                                                                            } else {
                                                                                                                var v95 uint8 = 42
                                                                                                                _ = v95
                                                                                                            }
                                                                                                        case "a":
                                                                                                            var v96 float64 = 1.5
                                                                                                            _ = v96
                                                                                                        }
                                                                                                    }
                                                                                                case "a":
                                                                                                    var v97 int8 = 0
                                                                                                    _ = v97
                                                                                                }
                                                                                            case -9223372036854775808:
                                                                                                switch "\n\t" {
                                                                                                case "foo bar":
                                                                                                    if v2 {
                                                                                                        fmt.Println(v2)
                                                                                                    } else {
                                                                                                        if ((v1 - (-2.25 + ((v55 + 0.0) + (float64(1.5) - ((0.1 / (1.5 - -2.25)) * 1.5))))) != v55) {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            v5 += 9223372036854775807
                                                                                                        } else {
                                                                                                            fmt.Println(v4)
                                                                                                            switch (fi15 / ((fi15) | 1)) {
                                                                                                            case -9223372036854775808:
                                                                                                                var v98 uint8 = 127
                                                                                                                _ = v98
                                                                                                            case 469:
                                                                                                                v87 = v87
                                                                                                            default:
                                                                                                                var v99 int8 = -1
                                                                                                                _ = v99
                                                                                                            }
                                                                                                            if false {
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                        }
                                                                                                        var v100 int = (fi15 ^ fi15)
                                                                                                        var v101 int32 = 2147483647
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    }
                                                                                                default:
                                                                                                    v102 := GenS12{F0: 32767}
                                                                                                    _ = v102
                                                                                                }
                                                                                            case 2147483647:
                                                                                                v103 := GenS13{F0: 1, F1: -1}
                                                                                                _ = v103
                                                                                            default:
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    fmt.Println(v2)
                                                                                                default:
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = ((("a" + "foo bar") + "") + ("foo bar" + "foo bar"))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v5)
                                                                                            _ = v91
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    v2 = (42 <= 1)
                                                                                }
                                                                                v2 = (!(v4 < (v4 << ((20) & 63))))
                                                                                if (42 <= 42) {
                                                                                    v55 -= float64((0.0 * 0.1))
                                                                                    fmt.Println(v22, v2, v5)
                                                                                } else {
                                                                                    var v104 uint8 = 2
                                                                                    _ = v104
                                                                                }
                                                                                var v105 float32 = 1.5
                                                                                fmt.Println(v1, v105, v4)
                                                                                _ = v105
                                                                                _ = v87
                                                                            }
                                                                        }
                                                                        switch ((("hello" + ("hello" + ("hello" + "a"))) + (("" + ("foo bar" + "\n\t")) + "a")) + "\n\t") {
                                                                        case "\n\t":
                                                                            v106 := GenS14{F0: 622, F1: 42, F2: 32767}
                                                                            _ = v106
                                                                        case "a":
                                                                            v55 += float64(float32((((v1 + v1) / (v55 + ((v55 + v55) * (1e10 * 1.5)))) * ((-2.25 + v55) * (float64(((1e10 - 3.14159) + 3.14159)) + float64(3.14159))))))
                                                                        default:
                                                                            v1 = ((v55 + 1e10) - v55)
                                                                        }
                                                                        fmt.Println(fi15)
                                                                        fmt.Println(fi15)
                                                                        var v107 int64 = v5
                                                                        _ = v107
                                                                    }
                                                                case "":
                                                                    var v108 uint16 = 267
                                                                    _ = v108
                                                                default:
                                                                    var v109 uint16 = 42
                                                                    _ = v109
                                                                }
                                                                if (42 >= -1) {
                                                                    var v110 uint = 1
                                                                    _ = v110
                                                                } else {
                                                                    if ((fi15 * (fi15 / (((^(fi15))) | 1))) == int(v22)) {
                                                                        var v111 string = ("a" + "hello")
                                                                        fmt.Println(v5, v111)
                                                                        var v112 int64 = -9223372036854775808
                                                                        _ = v111
                                                                        _ = v112
                                                                    }
                                                                }
                                                            }
                                                            var v113 string = "\n\t"
                                                            v114 := GenS2{F0: v5, F1: 0, F2: 1000}
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                    case "":
                                                        v115 := GenS13{F0: 656, F1: v5}
                                                        _ = v115
                                                    }
                                                }
                                            case 42:
                                                if (v22 == v4) {
                                                    switch (^(fi50)) {
                                                    case 42:
                                                        v1 -= float64(3.14159)
                                                    case 4294967295:
                                                        fmt.Println(v2, fi15, v5)
                                                    case 128:
                                                        v55 -= 1.5
                                                    default:
                                                        v116 := GenS15{F0: -32768}
                                                        v3 = append(v3, ("hello" + (("\n\t" + ("hello" + "")) + "\n\t")))
                                                        v22 = 1
                                                        _ = v116
                                                    }
                                                } else {
                                                    if false {
                                                        fmt.Println(v55, v4)
                                                        v22 += v4
                                                        v55 += v1
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                }
                                            }
                                        }
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    if (1.5 >= 0.0) {
                                        v5 += 1
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            var v117 int32 = -2147483648
                                            _ = v117
                                        case "\n\t":
                                            v22 += 2
                                        default:
                                            v118 := GenS9{F0: fi50}
                                            _ = v118
                                        }
                                        v1 = float64(1.5)
                                    }
                                    v22 += (v22 | v4)
                                    var v119 int8 = 127
                                    _ = v119
                                }
                                v4 = v22
                                if v2 {
                                    var v120 bool = true
                                    _ = v120
                                }
                                v121 := GenS14{F0: 129, F1: -121, F2: (fi15 % ((fi15) | 1))}
                                v22 = v4
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v122 := GenS15{F0: 968}
                                _ = fi50
                                _ = v121
                                _ = v122
                            }
                            v2 = (!(4294967295 >= 813))
                        }
                    case -9223372036854775808:
                        fmt.Println(v22, v22)
                    default:
                        var v123 bool = v2
                        _ = v123
                    }
                    _ = fi15
                    _ = v22
                }
                v1 -= v1
                switch -421 {
                case 127:
                    var v124 uint8 = 128
                    _ = v124
                case -9223372036854775808:
                    var v125 float64 = (v1 - -2.25)
                    _ = v125
                case 840:
                    var v126 uint64 = 1
                    _ = v126
                default:
                    for fi127 := 0; fi127 < 3; fi127++ {
                        fmt.Println(v2, v5)
                        v128 := GenS6{F0: 216, F1: 292}
                        _ = fi127
                        _ = v128
                    }
                }
            } else {
                if v2 {
                    var v129 int8 = 0
                    v5 = -9223372036854775808
                    for fi130 := 0; fi130 < 9; fi130++ {
                        fmt.Println(v5, fi130)
                        _ = fi130
                    }
                    _ = v129
                }
            }
            switch "a" {
            case "foo bar":
                v131 := GenS16{F0: -32768}
                _ = v131
            }
            fmt.Println(v1, v5, v4)
        }
        v1 = -2.25
    } else {
        v5 = v5
        v1 += 0.0
    }
    if ((469 >= 855) || v2) {
        var v132 int = 1
        var v133 int32 = 0
        _ = v132
        _ = v133
    }
    var v134 int = -14
    v4 += v4
    if (!v2) {
        for fi135 := 0; fi135 < 3; fi135++ {
            switch (v134 + v134) {
            case 9223372036854775807:
                if v2 {
                    v136 := GenS1{F0: 0}
                    fmt.Println(v1, v134)
                    if len(v3) > 0 {
                        v3[0] = ((("a" + "") + "") + "hello")
                    }
                    _ = v136
                }
            default:
                v134 -= 4294967295
            }
            for fi137 := 0; fi137 < 13; fi137++ {
                if true {
                    if (-1 < 32767) {
                        v138 := GenS13{F0: 4294967295, F1: (^(v5))}
                        _ = len(v3)
                        _ = v138
                    }
                    var v139 bool = v2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v139 = v2
                    v140 := GenS17{F0: 1, F1: -128, F2: 255}
                    _ = v139
                    _ = v140
                } else {
                    v134 += (fi135 & fi137)
                }
                v1 = (1e10 / (0.0 + 1e10))
                v141 := GenS18{F0: (v5 & (v5 << ((65535) & 63))), F1: 164, F2: (v5 % ((885) | 1))}
                _ = fi137
                _ = v141
            }
            switch fi135 {
            case 4294967295:
                v4 += v4
            case 679:
                fmt.Println(fi135, v1, v1)
            case -9223372036854775808:
                var v142 int = 565
                _ = v142
            }
            _ = fi135
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if v2 {
            var v143 uint = 65535
            _ = v143
        }
        v144 := GenS10{F0: 1}
        _ = v144
    }
    var v145 int32 = -2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v134
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
