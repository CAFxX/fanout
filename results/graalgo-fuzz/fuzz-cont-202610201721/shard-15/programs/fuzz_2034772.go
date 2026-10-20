package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int64; F1 int16; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint16; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 uint32 }

func main() {
    v1 := []int{9223372036854775807, -712}
    v2 := []string{("" + "a"), ("a" + "hello"), "foo bar"}
    var v3 bool = false
    var v4 uint64 = 0
    var v5 int = 4294967295
    if (v4 <= (^(v4))) {
        switch "foo bar" {
        case "":
            v3 = (1e10 >= float64(1.5))
        case "hello":
            for fi6 := 0; fi6 < 2; fi6++ {
                for fi7 := 0; fi7 < 9; fi7++ {
                    fmt.Println(v5)
                    _ = fi7
                }
                var v8 int32 = -2147483648
                var v9 uint8 = 93
                _ = fi6
                _ = v8
                _ = v9
            }
        default:
            var v10 string = ("" + ("foo bar" + "hello"))
            _ = v10
        }
        var v11 uint32 = 2
        _ = v11
    }
    if (v3 || v3) {
        var v12 uint32 = 0
        if v3 {
            if (!(42 < 1)) {
                if len(v2) > 0 {
                    v2[0] = ""
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch ("\n\t" + "") {
                case "\n\t":
                    if ((v5 << ((1) & 63)) != (v5 << ((18446744073709551615) & 63))) {
                        v5 = 0
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi13 := 0; fi13 < 4; fi13++ {
                            switch v5 {
                            case -2147483648:
                                var v14 uint32 = v12
                                _ = v14
                            case 897:
                                if (0 > 32767) {
                                    v15 := GenS1{F0: 868, F1: -61}
                                    switch fi13 {
                                    case 1:
                                        var v16 int8 = 10
                                        _ = v16
                                    case -674:
                                        v17 := GenS2{F0: 32767, F1: 1}
                                        _ = v17
                                    case 9223372036854775807:
                                        if true {
                                            v3 = true
                                        }
                                    default:
                                        var v18 int32 = -1
                                        _ = v18
                                    }
                                    v3 = v3
                                    _ = v15
                                } else {
                                    v4 = (v4 >> ((18446744073709551615) & 63))
                                }
                            }
                            var v19 int32 = 42
                            var v20 uint = 1
                            _ = fi13
                            _ = v19
                            _ = v20
                        }
                    } else {
                        var v21 int64 = 42
                        _ = v21
                    }
                case "foo bar":
                    var v22 string = "\n\t"
                    _ = v22
                case "hello":
                    var v23 int8 = 0
                    _ = v23
                default:
                    v2 = append(v2, ("\n\t" + "foo bar"))
                }
            } else {
                if true {
                    v24 := GenS1{F0: 965, F1: -87}
                    v25 := GenS2{F0: 0, F1: 42}
                    var v26 uint32 = uint32(v4)
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = v12
    }
    v4 = v4
    var v27 int = v5
    fmt.Println(v27)
    v5 -= v5
    if false {
        switch "hello" {
        case "hello":
            if (!v3) {
                fmt.Println(v27, v4)
                v5 += 32767
                if ((!(995 >= 2147483647)) || (!v3)) {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            if (!true) {
                                fmt.Println(v3)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v30 uint32 = 42
                                v27 += (^(v27))
                                _ = v30
                            }
                            switch fi29 {
                            case 32768:
                                if false {
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v27, v27, v4)
                                    default:
                                        fmt.Println(fi29, v27)
                                    }
                                    if (v3 || true) {
                                        fmt.Println(v3, v27, v3)
                                    } else {
                                        if v3 {
                                            switch ("" + ("a" + "foo bar")) {
                                            case "a":
                                                switch (((("" + "foo bar") + ("" + "\n\t")) + "hello") + "hello") {
                                                case "":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            default:
                                                switch (("" + "foo bar") + "") {
                                                case "hello":
                                                    if (!v3) {
                                                        var v31 int64 = 9223372036854775807
                                                        v27 -= -67
                                                        v27 = 1
                                                        _ = v31
                                                    } else {
                                                        var v32 int16 = -953
                                                        _ = v32
                                                    }
                                                }
                                            }
                                            v3 = true
                                            var v33 uint64 = 0
                                            if v3 {
                                                var v34 uint16 = 521
                                                switch "hello" {
                                                case "":
                                                    var v35 uint8 = 138
                                                    _ = v35
                                                }
                                                _ = v34
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = ("hello" + (("a" + "a") + ("\n\t" + "a")))
                                            }
                                            _ = v33
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                } else {
                                    if v3 {
                                        v3 = v3
                                        if (!(float64(float32((1.5 * -2.25))) <= 1.5)) {
                                            v36 := GenS2{F0: 30, F1: -471}
                                            v37 := GenS2{F0: 796, F1: -32768}
                                            _ = v36
                                            _ = v37
                                        }
                                        var v38 uint16 = 0
                                        _ = v38
                                    }
                                    v39 := GenS3{F0: 65535, F1: 100}
                                    _ = v39
                                }
                            default:
                                _ = len(v2)
                            }
                            var v40 uint = 1
                            _ = fi29
                            _ = v40
                        }
                        switch "" {
                        case "hello":
                            if (-2.25 != 0.1) {
                                if false {
                                    if (-389 == fi28) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (!true) {
                                                if v3 {
                                                    if (((255 != 43) && v3) && v3) {
                                                        for fi41 := 0; fi41 < 8; fi41++ {
                                                            switch "a" {
                                                            case "a":
                                                                v5 = (-(v27))
                                                            case "foo bar":
                                                                fmt.Println(v27, v4)
                                                            case "":
                                                                var v42 int64 = 9223372036854775807
                                                                _ = v42
                                                            default:
                                                                switch ("" + "\n\t") {
                                                                case "\n\t":
                                                                    v27 += fi28
                                                                }
                                                            }
                                                            if v3 {
                                                                if (!v3) {
                                                                    v1 = append(v1, v5)
                                                                    var v43 bool = ((float64(float32(((1e10 / 1e10) + (float64(0.1) / (1.5 * (0.1 / -2.25)))))) == 0.1) && (!(193 < -1)))
                                                                    var v44 float32 = 3.14159
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                var v45 uint8 = 127
                                                                _ = v45
                                                            } else {
                                                                v4 += v4
                                                                switch (("hello" + "") + "") {
                                                                case "a":
                                                                    if (255 == 128) {
                                                                        var v46 uint64 = 680
                                                                        var v47 float64 = 1.5
                                                                        v27 -= fi41
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                default:
                                                                    if true {
                                                                        v4 += (^(v4))
                                                                        switch (v27 - (v27 + (v5 & fi28))) {
                                                                        case -147:
                                                                            _ = len(v1)
                                                                        case 42:
                                                                            switch v27 {
                                                                            case 1:
                                                                                var v48 uint = 432
                                                                                _ = v48
                                                                            }
                                                                        case -993:
                                                                            v3 = ((!v3) && (!v3))
                                                                        default:
                                                                            switch (v5 & 4294967295) {
                                                                            case -894:
                                                                                _ = len(v1)
                                                                            case 32768:
                                                                                v49 := GenS3{F0: 552, F1: 1}
                                                                                _ = v49
                                                                            case 2147483647:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (fi41 % ((fi41) | 1))
                                                                                }
                                                                            default:
                                                                                switch "" {
                                                                                case "a":
                                                                                    v4 += 4294967295
                                                                                case "foo bar":
                                                                                    switch "" {
                                                                                    case "":
                                                                                        switch 4294967295 {
                                                                                        case -1:
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = fi28
                                                                                            }
                                                                                        case -2147483648:
                                                                                            var v50 int8 = 102
                                                                                            _ = v50
                                                                                        case 2147483647:
                                                                                            var v51 uint32 = 65535
                                                                                            _ = v51
                                                                                        }
                                                                                    case "hello":
                                                                                        v52 := GenS2{F0: 65535, F1: 590}
                                                                                        _ = v52
                                                                                    case "foo bar":
                                                                                        v3 = true
                                                                                    default:
                                                                                        switch (v5 >> ((4294967295) & 63)) {
                                                                                        case 128:
                                                                                            fmt.Println(fi41, fi41)
                                                                                        }
                                                                                        v4 = (v4 & 42)
                                                                                        var v53 int8 = 69
                                                                                        _ = v53
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v54 := GenS4{F0: -9223372036854775808, F1: 42, F2: 601}
                                                                        _ = v54
                                                                    } else {
                                                                        v5 -= (-(fi41))
                                                                        v27 = (-(fi28))
                                                                        if (-32768 != 42) {
                                                                            if (v3 || v3) {
                                                                                v55 := GenS5{F0: v3}
                                                                                v56 := GenS5{F0: (42 == 2)}
                                                                                var v57 int16 = 32767
                                                                                _ = v55
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                v5 = 0
                                                                            }
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v5, v5)
                                                                            var v58 uint64 = v4
                                                                            _ = v58
                                                                        }
                                                                        fmt.Println(fi28)
                                                                    }
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    _ = len(v1)
                                                                case "":
                                                                    var v59 int16 = 1000
                                                                    _ = v59
                                                                case "hello":
                                                                    var v60 uint16 = 959
                                                                    _ = v60
                                                                default:
                                                                    v27 = 1
                                                                }
                                                            }
                                                            _ = fi41
                                                        }
                                                    } else {
                                                        v27 = (v5 + -9223372036854775808)
                                                    }
                                                    for fi61 := 0; fi61 < 7; fi61++ {
                                                        v5 -= v27
                                                        if false {
                                                            if true {
                                                                if v3 {
                                                                    switch fi28 {
                                                                    case -2147483648:
                                                                        v62 := GenS6{F0: ("hello" + ("" + ("a" + (("" + "\n\t") + "\n\t")))), F1: 1, F2: (fi28 / (((v27 * fi61)) | 1))}
                                                                        _ = v62
                                                                    case 32768:
                                                                        v4 += (v4 % (((v4 / ((v4) | 1))) | 1))
                                                                    default:
                                                                        var v63 int = (fi28 << ((37) & 63))
                                                                        _ = v63
                                                                    }
                                                                    var v64 int16 = 1
                                                                    _ = v64
                                                                }
                                                                v3 = (0 <= -1)
                                                                fmt.Println(v27)
                                                                v5 -= (v27 << ((4294967295) & 63))
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = ("a" + ((("a" + ("hello" + "\n\t")) + "a") + ("foo bar" + "a")))
                                                                }
                                                                fmt.Println(fi61, v27, v3)
                                                            }
                                                        }
                                                        _ = fi61
                                                    }
                                                    v5 = (v5 / ((-9223372036854775808) | 1))
                                                }
                                                var v65 uint32 = 65535
                                                _ = v65
                                            } else {
                                                switch ((("" + "hello") + ((("foo bar" + "") + "a") + "a")) + "foo bar") {
                                                case "\n\t":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "a":
                                                    fmt.Println(fi28)
                                                }
                                            }
                                        case "":
                                            switch "\n\t" {
                                            case "":
                                                _ = len(v1)
                                            case "a":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v4 = v4
                                                case "a":
                                                    _ = len(v2)
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi28, v5, v27)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = v27
                                    }
                                    for fi66 := 0; fi66 < 8; fi66++ {
                                        _ = len(v1)
                                        fmt.Println(fi28, fi28)
                                        if (-128 > 1) {
                                            fmt.Println(v27)
                                            v27 = (v5 * fi66)
                                        }
                                        v67 := GenS7{F0: 1000}
                                        _ = fi66
                                        _ = v67
                                    }
                                } else {
                                    v27 = v5
                                    v2 = append(v2, ("hello" + ("" + ("foo bar" + "hello"))))
                                }
                                var v68 int32 = 110
                                _ = v68
                            } else {
                                v4 -= v4
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        _ = fi28
                    }
                    v69 := GenS8{F0: (1 < 0)}
                    _ = v69
                } else {
                    for fi70 := 0; fi70 < 5; fi70++ {
                        switch (v27 & 42) {
                        case 128:
                            var v71 uint64 = 594
                            _ = v71
                        }
                        v4 -= v4
                        _ = fi70
                    }
                    var v72 uint64 = 0
                    if (!true) {
                        if ((!false) || (("a" + "a") > "")) {
                            v73 := GenS7{F0: -1}
                            _ = v73
                        }
                        switch "\n\t" {
                        case "a":
                            var v74 uint8 = 98
                            _ = v74
                        default:
                            if true {
                                switch v27 {
                                case 1:
                                    switch (v27 << ((994) & 63)) {
                                    case 0:
                                        fmt.Println(v3)
                                    default:
                                        v75 := GenS5{F0: v3}
                                        _ = v75
                                    }
                                case 32768:
                                    fmt.Println(v5, v72)
                                default:
                                    var v76 bool = (2147483647 < -1)
                                    _ = v76
                                }
                            } else {
                                v3 = v3
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi77 := 0; fi77 < 14; fi77++ {
                                    for fi78 := 0; fi78 < 16; fi78++ {
                                        fmt.Println(v5, v5, v5)
                                        var v79 string = "a"
                                        _ = fi78
                                        _ = v79
                                    }
                                    var v80 int32 = 1
                                    if (!false) {
                                        var v81 float32 = (1.5 + 0.1)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v3 = false
                                        _ = v81
                                    } else {
                                        v5 += v5
                                        if v3 {
                                            if (!(!v3)) {
                                                for fi82 := 0; fi82 < 16; fi82++ {
                                                    var v83 uint = 255
                                                    var v84 uint64 = v72
                                                    _ = fi82
                                                    _ = v83
                                                    _ = v84
                                                }
                                                v27 = v5
                                            }
                                            fmt.Println(v80)
                                            var v85 uint = 1
                                            _ = v85
                                        } else {
                                            var v86 int16 = -1
                                            _ = v86
                                        }
                                    }
                                    _ = fi77
                                    _ = v80
                                }
                            }
                        }
                        for fi87 := 0; fi87 < 13; fi87++ {
                            if false {
                                switch (fi87 >> ((18446744073709551615) & 63)) {
                                case -9223372036854775808:
                                    v4 += 1
                                case -906:
                                    v88 := GenS5{F0: v3}
                                    _ = v88
                                case -2147483648:
                                    v89 := GenS5{F0: v3}
                                    _ = v89
                                }
                                if (v3 && (false && (-667 > 9223372036854775807))) {
                                    v72 = (^(v72))
                                } else {
                                    var v90 string = "hello"
                                    _ = v90
                                }
                                v27 += -563
                            } else {
                                if true {
                                    v91 := GenS9{F0: 0}
                                    v92 := GenS10{F0: -257, F1: 1}
                                    _ = v91
                                    _ = v92
                                }
                                fmt.Println(v72, v5, v27)
                            }
                            if true {
                                var v93 uint = 142
                                v3 = (42 >= 2147483647)
                                var v94 int16 = -1
                                _ = v93
                                _ = v94
                            }
                            if v3 {
                                for fi95 := 0; fi95 < 5; fi95++ {
                                    v96 := GenS11{F0: 782}
                                    _ = len(v2)
                                    _ = fi95
                                    _ = v96
                                }
                                v2 = append(v2, "\n\t")
                            }
                            _ = fi87
                        }
                    }
                    _ = v72
                }
            }
        }
        switch int(v4) {
        case 0:
            var v97 int = -2147483648
            _ = v97
        default:
            fmt.Println(v27)
        }
        _ = len(v2)
    } else {
        if len(v1) > 0 {
            v1[0] = (v27 / ((v5) | 1))
        }
    }
    if len(v2) > 0 {
        v2[0] = ((("a" + (("foo bar" + ("\n\t" + "")) + ("hello" + ""))) + "a") + "a")
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ("hello" > "foo bar") {
        for fi98 := 0; fi98 < 17; fi98++ {
            if false {
                v4 = 181
                var v99 uint64 = v4
                _ = v99
            } else {
                var v100 float32 = (((3.14159 / float32((0.0 - (3.14159 - (1e10 / (3.14159 / ((-2.25 * (3.14159 - 1e10)) - 0.1))))))) - 1e10) + 1.5)
                v101 := GenS10{F0: -1, F1: 378}
                _ = v100
                _ = v101
            }
            for fi102 := 0; fi102 < 6; fi102++ {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                if (-1 > -102) {
                    fmt.Println(fi98, v5, fi102)
                    var v103 uint32 = 65535
                    _ = v103
                }
                _ = fi102
            }
            _ = fi98
        }
    } else {
        if true {
            fmt.Println(v4)
        }
    }
    var v104 int = (^(v5))
    var v105 int32 = -2147483648
    v106 := GenS1{F0: 632, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
