package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8; F1 int32; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint32; F1 bool; F2 int32 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 bool }
func genFn0(p0 string, p1 int16, p2 int) uint64 {
    switch p0 {
    case "hello":
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{32768, 2147483647, fi7}
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 0}
    var v10 int32 = 0
    fmt.Println(p0, v10, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 360
}

func main() {
    var v1 string = ((("a" + "") + "") + "")
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := make(map[uint32]string)
    v6 := []int{0, 127}
    switch 35 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: 2}
            fmt.Println(v1, v1)
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS1{F0: 255}
                v1 = v1
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    case 127:
        for fi15 := 0; fi15 < 13; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                v1 = v1
                _ = fi16
            }
            v1 = "a"
            v1 = "\n\t"
            _ = fi15
        }
    case -2147483648:
        if true {
            var v17 bool = ((("a" > "a") || false) && false)
            switch ((v1 + v1) + v1) {
            case "foo bar":
                for fi18 := 0; fi18 < 20; fi18++ {
                    v17 = true
                    _ = fi18
                }
            default:
                var v19 int16 = 328
                _ = v19
            }
            _ = v17
        } else {
            fmt.Println(v1)
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(fi20, v1)
                var v21 float64 = 3.14159
                for fi22 := 0; fi22 < 18; fi22++ {
                    v21 = v21
                    _ = fi22
                }
                _ = fi20
                _ = v21
            }
        }
    }
    switch ((v1 + (v1 + (v1 + "a"))) + ("\n\t" + (("hello" + v1) + (v1 + v1)))) {
    case "\n\t":
        delete(v2, 906)
    case "a":
        for fi23 := 0; fi23 < 15; fi23++ {
            v1 = "a"
            _ = fi23
        }
    case "foo bar":
        if (-1 == -32768) {
            fmt.Println(v1)
            v1 = "foo bar"
        }
    default:
        v1 = (v1 + v1)
    }
    _ = len(v6)
    v1 = ""
    switch ("hello" + (v1 + v1)) {
    case "":
        v24 := GenS2{F0: 369, F1: 42}
        _ = v24
    case "a":
        for fi25 := 0; fi25 < 1; fi25++ {
            switch v1 {
            case "a":
                fmt.Println(v1, fi25)
            case "hello":
                switch (-(fi25)) {
                case 944:
                    if true {
                        var v26 uint = 2
                        if len(v6) > 0 {
                            v6[0] = (fi25 | (fi25 / ((fi25) | 1)))
                        }
                        _ = v26
                    }
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            case "\n\t":
                for fi27 := 0; fi27 < 4; fi27++ {
                    v28 := GenS1{F0: 1}
                    var v29 float64 = float64(float32(3.14159))
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi25
        }
    case "foo bar":
        v1 = (v1 + "foo bar")
    }
    v4[(v1 + v1)] = v1
    if len(v6) > 0 {
        v6[0] = 127
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint8 = 42
            v5[0] = v1
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            if true {
                var v34 int32 = 1
                if (false && (-966 <= (^(fi30)))) {
                    switch v1 {
                    case "":
                        var v35 uint = 42
                        _ = v35
                    case "foo bar":
                        v34 -= v34
                    case "\n\t":
                        switch (("foo bar" + (((v1 + ((v1 + (v1 + v1)) + ((((v1 + v1) + v1) + v1) + ((v1 + (("a" + "hello") + (v1 + (v1 + v1)))) + "foo bar")))) + (v1 + "foo bar")) + (v1 + v1))) + v1) {
                        case "":
                            var v36 int = (fi33 * (fi30 * fi33))
                            _ = v36
                        case "hello":
                            fmt.Println(fi30, fi30, fi30)
                        default:
                            var v37 int = fi33
                            _ = v37
                        }
                    }
                    v1 = v1
                    v38 := GenS1{F0: 255}
                    _ = v38
                }
                _ = v34
            }
            v39 := GenS3{F0: 18446744073709551615, F1: 135, F2: 840}
            if true {
                if false {
                    if len(v6) > 0 {
                        v6[0] = 128
                    }
                    if len(v6) > 0 {
                        v6[0] = (fi30 | (fi30 - 32767))
                    }
                }
                if ("hello" < v1) {
                    fmt.Println(fi33, v1)
                    if (0 > 1) {
                        if false {
                            if (931 > -1) {
                                v1 = (v1 + v1)
                                v1 = "foo bar"
                            }
                        }
                        v1 = v1
                        if (true && false) {
                            var v40 string = v1
                            v1 = ("\n\t" + "hello")
                            _ = v40
                        }
                        v1 = (v1 + v1)
                    }
                    var v41 float64 = 1.5
                    v42 := GenS4{F0: 255, F1: 573, F2: 42}
                    if len(v6) > 0 {
                        v6[0] = -2147483648
                    }
                    _ = v41
                    _ = v42
                }
            }
            var v43 float32 = 1e10
            _ = fi33
            _ = v39
            _ = v43
        }
        _ = len(v6)
        _ = fi30
    }
    _ = len(v6)
    v1 = (v1 + v1)
    v1 = ""
    switch 486 {
    case 4294967295:
        fmt.Println(v1)
    case 2147483647:
        for fi44 := 0; fi44 < 7; fi44++ {
            switch "" {
            case "foo bar":
                fmt.Println(v1, fi44)
            default:
                for fi45 := 0; fi45 < 11; fi45++ {
                    switch v1 {
                    case "\n\t":
                        v1 = v1
                    case "foo bar":
                        v1 = ((("" + "") + (v1 + "hello")) + (v1 + ("a" + ("hello" + v1))))
                        v1 = (v1 + ("\n\t" + ((v1 + ((v1 + "hello") + ("foo bar" + v1))) + v1)))
                        if (((-2.25 + (((1.5 * (float64(1.5) / 1.0)) / (3.14159 * ((1.5 - -2.25) - float64((float32(0.1) - 0.1))))) * (3.14159 + ((-2.25 / -2.25) * 1e10)))) * (1e10 * float64(0.0))) < 0.1) {
                            v46 := GenS5{F0: 892}
                            v1 = "a"
                            _ = v46
                        }
                    case "":
                        var v47 int64 = 267
                        _ = v47
                    default:
                        if false {
                            v1 = (v1 + v1)
                            switch fi45 {
                            case 9223372036854775807:
                                if false {
                                    v1 = "a"
                                    v1 = "foo bar"
                                    v1 = v1
                                    switch (fi45 ^ -683) {
                                    case -579:
                                        switch (fi45 & fi45) {
                                        case -9223372036854775808:
                                            v1 = v1
                                        case -518:
                                            if len(v6) > 0 {
                                                v6[0] = (fi44 * -9223372036854775808)
                                            }
                                        case -2147483648:
                                            if false {
                                                switch 0 {
                                                case -9223372036854775808:
                                                    var v48 float64 = (1e10 - (0.1 * 1.5))
                                                    _ = v48
                                                case 145:
                                                    if false {
                                                        if true {
                                                            switch fi44 {
                                                            case 127:
                                                                v1 = (((v1 + v1) + "\n\t") + ((v1 + ("hello" + v1)) + v1))
                                                            case 138:
                                                                v1 = "\n\t"
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            var v49 bool = (2 > 42)
                                                            _ = v49
                                                        }
                                                        fmt.Println(fi44, fi44, fi44)
                                                        v1 = (v1 + v1)
                                                    }
                                                case 9223372036854775807:
                                                    fmt.Println(v1, fi44)
                                                }
                                                var v50 float64 = -2.25
                                                _ = v50
                                            }
                                        }
                                    default:
                                        if len(v6) > 0 {
                                            v6[0] = (fi45 - (^(fi45)))
                                        }
                                    }
                                    v51 := GenS6{F0: "a", F1: 100, F2: 42}
                                    if (297 == 42) {
                                        if (true || true) {
                                            fmt.Println(fi44)
                                            var v52 float32 = 0.1
                                            v53 := GenS7{F0: 65535, F1: ((1000 >= -1) || false), F2: 0}
                                            _ = v52
                                            _ = v53
                                        }
                                        var v54 int8 = 86
                                        fmt.Println(v1, v1)
                                        v55 := GenS8{F0: 0, F1: 582}
                                        switch (fi45 * (fi44 << ((641) & 63))) {
                                        case -2147483648:
                                            v1 = (v1 + "a")
                                        case 711:
                                            var v56 int = (fi44 >> ((1) & 63))
                                            _ = v56
                                        case 128:
                                            _ = len(v6)
                                        default:
                                            var v57 float32 = float32(0.0)
                                            _ = v57
                                        }
                                        _ = v54
                                        _ = v55
                                    } else {
                                        v1 = ""
                                        if false {
                                            v1 = ((v1 + (v1 + ((v1 + v1) + v1))) + (v1 + "a"))
                                            var v58 int = (fi45 - (^(fi44)))
                                            v58 += -2147483648
                                            var v59 uint8 = 51
                                            _ = v58
                                            _ = v59
                                        }
                                        switch (fi44 ^ fi44) {
                                        case 128:
                                            v1 = "a"
                                        }
                                    }
                                    _ = v51
                                }
                            case 32767:
                                if false {
                                    var v60 int = fi45
                                    if true {
                                        var v61 uint32 = 1
                                        _ = v61
                                    } else {
                                        v1 = "foo bar"
                                        var v62 int8 = 108
                                        _ = v62
                                    }
                                    _ = v60
                                } else {
                                    switch (fi45 % (((fi45 / ((fi45) | 1))) | 1)) {
                                    case 961:
                                        var v63 uint32 = 4294967295
                                        _ = v63
                                    case -9223372036854775808:
                                        switch v1 {
                                        case "":
                                            if false {
                                                switch (fi45 % (((^(fi45))) | 1)) {
                                                case 42:
                                                    if (false || ((((v1 + v1) + ("hello" + ((v1 + ((((v1 + "foo bar") + v1) + (("foo bar" + "") + v1)) + v1)) + v1))) + v1) == "hello")) {
                                                        v1 = v1
                                                        switch 127 {
                                                        case 32768:
                                                            fmt.Println(v1, fi45, fi44)
                                                        case 70:
                                                            v64 := GenS9{F0: 42, F1: (false || false)}
                                                            _ = v64
                                                        case -9223372036854775808:
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        default:
                                                            if true {
                                                                v65 := GenS3{F0: 2, F1: 2, F2: 255}
                                                                _ = v65
                                                            } else {
                                                                v1 = (v1 + "hello")
                                                                switch fi45 {
                                                                case 2147483647:
                                                                    var v66 int8 = -128
                                                                    _ = v66
                                                                case 127:
                                                                    v67 := GenS1{F0: 127}
                                                                    _ = v67
                                                                case -2147483648:
                                                                    var v68 float64 = (0.1 / 0.1)
                                                                    _ = v68
                                                                }
                                                                v69 := GenS9{F0: 2, F1: (42 == 0)}
                                                                if (false || true) {
                                                                    v1 = (((v1 + "\n\t") + ((v1 + "") + "a")) + ((v1 + ("foo bar" + "")) + (("" + ((v1 + (v1 + (v1 + v1))) + "\n\t")) + v1)))
                                                                } else {
                                                                    var v70 int32 = -2147483648
                                                                    fmt.Println(fi45, v1, v70)
                                                                    fmt.Println(fi45, fi45)
                                                                    if (true || (!((3.14159 < 1e10) && false))) {
                                                                        _ = len(v2)
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            _ = len(v4)
                                                                        case "a":
                                                                            v3[658] = v1
                                                                        case "":
                                                                            var v71 uint64 = 991
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                    var v72 int64 = 42
                                                                    _ = v70
                                                                    _ = v72
                                                                }
                                                                _ = v69
                                                            }
                                                        }
                                                    } else {
                                                        var v73 uint32 = 615
                                                        var v74 uint = 0
                                                        v75 := GenS5{F0: fi45}
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                }
                                                delete(v3, 21)
                                                v6 = append(v6, (-(fi44)))
                                            }
                                        }
                                    case 42:
                                        v1 = v1
                                    default:
                                        if false {
                                            switch 1 {
                                            case 1000:
                                                var v76 uint = 0
                                                _ = v76
                                            }
                                        } else {
                                            var v77 uint16 = 827
                                            switch -2147483648 {
                                            case 0:
                                                switch fi44 {
                                                case 1:
                                                    _ = len(v6)
                                                case 532:
                                                    v1 = v1
                                                case 32768:
                                                    if true {
                                                        var v78 int16 = 1000
                                                        _ = v78
                                                    }
                                                default:
                                                    v79 := GenS5{F0: (fi44 * fi44)}
                                                    _ = v79
                                                }
                                            case 128:
                                                var v80 bool = (false && false)
                                                _ = v80
                                            case -9223372036854775808:
                                                var v81 float32 = 0.0
                                                _ = v81
                                            }
                                            _ = v77
                                        }
                                    }
                                    if false {
                                        fmt.Println(fi44, v1)
                                        switch fi44 {
                                        case -9223372036854775808:
                                            if (-2.25 > (((((0.1 - 0.0) * 1.5) * (0.1 + (0.0 + 0.0))) / 1.0) + ((float32((3.14159 / 1e10)) - 0.0) + (-2.25 + -2.25)))) {
                                                v1 = ("hello" + v1)
                                                v1 = v1
                                            } else {
                                                v4[(("foo bar" + v1) + ("foo bar" + ((v1 + v1) + ((("" + (v1 + ((("" + v1) + ("foo bar" + v1)) + ("" + v1)))) + v1) + (((v1 + v1) + v1) + "foo bar")))))] = v1
                                                v82 := GenS1{F0: 255}
                                                _ = v82
                                            }
                                        case -391:
                                            var v83 uint8 = 42
                                            _ = v83
                                        case 127:
                                            if (!false) {
                                                fmt.Println(fi45, fi44, fi44)
                                                v1 = "\n\t"
                                                var v84 int64 = -41
                                                _ = v84
                                            } else {
                                                fmt.Println(fi44, v1, fi45)
                                                var v85 uint64 = 1
                                                if true {
                                                    _ = len(v5)
                                                    v2[515] = fi44
                                                }
                                                _ = v85
                                            }
                                        }
                                        var v86 uint16 = 707
                                        if true {
                                            var v87 int16 = -1
                                            v88 := v4[(v1 + "")]
                                            v1 = "hello"
                                            _ = v87
                                            _ = v88
                                        }
                                        _ = v86
                                    } else {
                                        v1 = (((v1 + v1) + v1) + "")
                                    }
                                }
                            }
                        }
                    }
                    _ = fi45
                }
            }
            var v89 uint64 = 4294967295
            fmt.Println(v89, fi44, fi44)
            _ = fi44
            _ = v89
        }
    default:
        var v90 uint32 = 65535
        _ = v90
    }
    var v91 int = 42
    if ((-2.25 * 3.14159) > float32(3.14159)) {
        if (!false) {
            var v92 uint16 = 32767
            for fi93 := 0; fi93 < 19; fi93++ {
                v94 := GenS10{F0: (!true)}
                _ = fi93
                _ = v94
            }
            _ = v92
        }
    } else {
        var v95 int = (v91 + 32767)
        _ = v95
    }
    for fi96 := 0; fi96 < 14; fi96++ {
        v1 = (v1 + "a")
        _ = fi96
    }
    v91 = v91
    switch v91 {
    case 4294967295:
        v1 = "a"
    }
    var v97 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v91:", v91)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v97
}
