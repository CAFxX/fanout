package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 bool }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int; F2 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 int32; F1 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 bool }
func genFn0() uint64 {
    var v4 int64 = -9223372036854775808
    var v5 float64 = 3.14159
    if false {
        v6 := GenS1{F0: 1}
        v7 := GenS1{F0: -9223372036854775808}
        var v8 string = "\n\t"
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: 4294967295, F1: -41}
        v12 := GenS3{F0: 42, F1: 42, F2: false}
        if (!false) {
            var v13 uint64 = 1
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v4
    _ = v5
    _ = v9
    return 2
}
func genFn1(p0 uint32, p1 bool) int64 {
    p1 = false
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int32 = 1
            var v17 int8 = 42
            v16 = 0
            v18 := GenS3{F0: 0, F1: v16, F2: p1}
            switch "" {
            case "hello":
                var v19 float64 = 3.14159
                _ = v19
            }
            v20 := GenS3{F0: 128, F1: -532, F2: true}
            v21 := []int{fi14}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
            _ = v20
            _ = v21
        }
        _ = fi14
    }
    var v22 bool = true
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int64 = -9223372036854775808
        switch (("hello" + ("a" + ("a" + ""))) + (((("hello" + "a") + "hello") + (("\n\t" + "foo bar") + "a")) + "foo bar")) {
        case "a":
            p1 = (4294967295 <= 609)
        case "":
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS4{F0: 0, F1: -128}
                if (!true) {
                    v24 = (^(v24))
                    var v27 uint64 = genFn0()
                    _ = v27
                } else {
                    var v28 float32 = 0.1
                    v29 := []int{-352, fi23, (fi23 * fi25), fi25}
                    if p1 {
                        _ = len(v29)
                        var v30 uint = 4294967295
                        if (("" + "") > ((("a" + (((("foo bar" + "foo bar") + "a") + "") + (("hello" + "\n\t") + "a"))) + "foo bar") + "hello")) {
                            var v31 int8 = 1
                            switch "" {
                            case "\n\t":
                                var v32 string = ("a" + "\n\t")
                                _ = v32
                            case "foo bar":
                                switch 32767 {
                                case 32767:
                                    if v22 {
                                        v33 := GenS5{F0: -32768}
                                        fmt.Println(v22, v31)
                                        v22 = false
                                        _ = v33
                                    }
                                case 347:
                                    switch ("" + "\n\t") {
                                    case "":
                                        v22 = ((0.0 < 1.5) || (("\n\t" + "foo bar") == ("" + ((("a" + ("hello" + (("" + ((("foo bar" + "hello") + "foo bar") + "foo bar")) + "a"))) + ("" + "foo bar")) + ("foo bar" + "foo bar")))))
                                    default:
                                        fmt.Println(v30, p0, fi25)
                                    }
                                default:
                                    if false {
                                        switch fi25 {
                                        case -1:
                                            switch fi25 {
                                            case 32768:
                                                p0 -= (p0 >> ((65535) & 63))
                                            case 127:
                                                v29 = append(v29, 309)
                                            }
                                        case 0:
                                            fmt.Println(p1)
                                        }
                                    }
                                }
                            }
                            if ((1 < v24) || true) {
                                _ = len(v29)
                                v24 = -909
                                v34 := GenS6{F0: 42, F1: 60}
                                v35 := GenS3{F0: 2, F1: 1, F2: p1}
                                _ = v34
                                _ = v35
                            } else {
                                var v36 uint8 = 127
                                var v37 int32 = 1
                                _ = v36
                                _ = v37
                            }
                            _ = v31
                        }
                        switch ("foo bar" + (("a" + "hello") + "\n\t")) {
                        case "":
                            fmt.Println(v28, v22)
                        case "hello":
                            fmt.Println(p1, v28)
                        }
                        var v38 int16 = -20
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        _ = v30
                        _ = v38
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
        case "hello":
            v22 = p1
        }
        p0 += (p0 / ((p0) | 1))
        _ = fi23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v22
    return 42
}

func main() {
    v1 := []int{32768}
    var v2 int = 364
    var v3 string = ""
    v39 := GenS1{F0: (v2 | (v2 - (^(v2))))}
    v40 := GenS1{F0: (-(v2))}
    v3 = ("" + v3)
    if len(v1) > 0 {
        v1[0] = v2
    }
    var v41 int64 = -9223372036854775808
    v42 := GenS7{F0: -582}
    if false {
        if (2147483647 > 4294967295) {
            v41 = v41
            for fi43 := 0; fi43 < 19; fi43++ {
                v3 = "a"
                v2 = v2
                if ((v41 * v41) > -855) {
                    v41 = (v41 << ((18446744073709551615) & 63))
                }
                _ = fi43
            }
            for fi44 := 0; fi44 < 8; fi44++ {
                switch ((v3 + v3) + (v3 + ((v3 + (v3 + "")) + "\n\t"))) {
                case "":
                    v3 = v3
                case "hello":
                    if true {
                        for fi45 := 0; fi45 < 11; fi45++ {
                            v2 = (v2 << ((661) & 63))
                            var v46 uint8 = 204
                            _ = fi45
                            _ = v46
                        }
                        if true {
                            for fi47 := 0; fi47 < 2; fi47++ {
                                fmt.Println(v41, v3, v2)
                                if (true && (1 > 346)) {
                                    var v48 string = "a"
                                    v3 = v48
                                    _ = v48
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (fi47 & 4294967295)
                                    }
                                }
                                switch (v3 + (v3 + v3)) {
                                case "a":
                                    switch (fi44 & fi44) {
                                    case 42:
                                        if true {
                                            var v49 uint = 2
                                            _ = v49
                                        } else {
                                            if true {
                                                var v50 int64 = (v41 / (((^(v41))) | 1))
                                                var v51 string = v3
                                                _ = v50
                                                _ = v51
                                            } else {
                                                var v52 uint32 = 4294967295
                                                _ = v52
                                            }
                                            fmt.Println(fi47)
                                            v2 = 128
                                            if (v41 < genFn1(824, (!(535 != 555)))) {
                                                if true {
                                                    v53 := GenS7{F0: -1}
                                                    v54 := GenS1{F0: (fi44 % (((fi47 << ((255) & 63))) | 1))}
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    switch ("\n\t" + "") {
                                                    case "foo bar":
                                                        v3 = v3
                                                    }
                                                    v55 := GenS1{F0: 128}
                                                    _ = v55
                                                }
                                            } else {
                                                v1 = append(v1, 128)
                                                var v56 int16 = -1
                                                fmt.Println(fi44, v56, v3)
                                                if len(v1) > 0 {
                                                    v1[0] = (v2 % (((v2 >> ((770) & 63))) | 1))
                                                }
                                                v56 += v56
                                                if (float32(float64(-2.25)) == 1.5) {
                                                    fmt.Println(v2, v41)
                                                    fmt.Println(fi47)
                                                } else {
                                                    if true {
                                                        fmt.Println(v41)
                                                        if true {
                                                            var v57 int16 = (v56 % (((v56 >> ((255) & 63))) | 1))
                                                            var v58 float64 = -2.25
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            v59 := GenS6{F0: 2, F1: 1}
                                                            v2 += (-(fi47))
                                                            v60 := GenS3{F0: 0, F1: -1, F2: false}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        v1 = append(v1, v2)
                                                    }
                                                    fmt.Println(fi47)
                                                }
                                                _ = v56
                                            }
                                            fmt.Println(v2)
                                        }
                                    }
                                case "":
                                    v61 := GenS7{F0: fi44}
                                    _ = v61
                                case "foo bar":
                                    v3 = "hello"
                                }
                                _ = fi47
                            }
                            for fi62 := 0; fi62 < 12; fi62++ {
                                v2 += 32767
                                v3 = "a"
                                _ = fi62
                            }
                        } else {
                            for fi63 := 0; fi63 < 2; fi63++ {
                                switch -108 {
                                case 0:
                                    v2 -= 2147483647
                                case 2147483647:
                                    var v64 bool = false
                                    _ = v64
                                default:
                                    v3 = ((("\n\t" + v3) + ((("a" + v3) + v3) + v3)) + v3)
                                }
                                fmt.Println(fi63)
                                v3 = "\n\t"
                                _ = fi63
                            }
                            if false {
                                for fi65 := 0; fi65 < 19; fi65++ {
                                    if false {
                                        if (127 != 1) {
                                            var v66 uint16 = 1
                                            var v67 int64 = (^(v41))
                                            _ = v66
                                            _ = v67
                                        } else {
                                            var v68 string = v3
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = len(v1)
                                            _ = v68
                                        }
                                        v1 = append(v1, -2147483648)
                                        v3 = (v3 + v3)
                                    } else {
                                        if (1 < 6) {
                                            switch (fi65 + (fi44 % (((^(fi65))) | 1))) {
                                            case -1:
                                                if (false && (0.0 <= 1.5)) {
                                                    if (!true) {
                                                        v2 += (fi44 % ((32768) | 1))
                                                    } else {
                                                        if false {
                                                            v69 := GenS4{F0: 2147483647, F1: 127}
                                                            var v70 uint = 400
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        fmt.Println(v41, fi44)
                                                    }
                                                    if (-68 < -128) {
                                                        v41 = v41
                                                        var v71 uint = 255
                                                        v72 := GenS2{F0: uint64(v71), F1: 127}
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        switch fi65 {
                                                        case -212:
                                                            if ((fi44 * (fi44 << ((42) & 63))) != v2) {
                                                                v1 = append(v1, 2147483647)
                                                            } else {
                                                                var v73 uint8 = 162
                                                                v41 += (-(v41))
                                                                switch ((v3 + (v3 + ("\n\t" + "hello"))) + "foo bar") {
                                                                case "":
                                                                    switch (fi44 / ((v2) | 1)) {
                                                                    case 4294967295:
                                                                        fmt.Println(v73, fi65)
                                                                        if (2147483647 > 1) {
                                                                            if (!(!false)) {
                                                                                v3 = (("hello" + v3) + (((((v3 + v3) + (v3 + v3)) + (v3 + (v3 + "\n\t"))) + ((("" + v3) + v3) + "a")) + v3))
                                                                                _ = len(v1)
                                                                                fmt.Println(fi44, fi65)
                                                                            } else {
                                                                                v3 = ""
                                                                                if (-1 == -1) {
                                                                                    if (255 == v73) {
                                                                                        v41 -= (v41 << ((255) & 63))
                                                                                        fmt.Println(v2, fi44)
                                                                                    }
                                                                                    var v74 float32 = 3.14159
                                                                                    switch v3 {
                                                                                    case "hello":
                                                                                        if (false || (v2 <= (-(v2)))) {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            v75 := GenS8{F0: (v41 / ((1) | 1)), F1: (fi65 * 4294967295), F2: (^(v73))}
                                                                                            if (true && (v2 <= 674)) {
                                                                                                v41 += (-(v41))
                                                                                            } else {
                                                                                                var v76 uint = 255
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                _ = v76
                                                                                            }
                                                                                            if true {
                                                                                                v74 += -2.25
                                                                                                if false {
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = (v2 / (((-(fi65))) | 1))
                                                                                                    }
                                                                                                    fmt.Println(fi65, v74)
                                                                                                } else {
                                                                                                    v2 -= fi44
                                                                                                }
                                                                                                if false {
                                                                                                    v3 = v3
                                                                                                    if true {
                                                                                                        v74 = float32(((-2.25 + (3.14159 - 0.0)) * 0.0))
                                                                                                        fmt.Println(fi65, v3, v2)
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                v77 := GenS6{F0: 42, F1: 0}
                                                                                                                _ = v77
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    var v78 uint16 = 13
                                                                                                                    fmt.Println(v78, fi65)
                                                                                                                    _ = v78
                                                                                                                }
                                                                                                                v41 = (v41 - (v41 ^ 42))
                                                                                                            }
                                                                                                            v79 := GenS4{F0: 1, F1: 42}
                                                                                                            _ = v79
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (2147483648 > 65535) {
                                                                                                            var v80 int = 32768
                                                                                                            if false {
                                                                                                                fmt.Println(fi44, v73)
                                                                                                                v73 += v73
                                                                                                            }
                                                                                                            _ = v80
                                                                                                        }
                                                                                                        if (!true) {
                                                                                                            switch v3 {
                                                                                                            case "hello":
                                                                                                                v73 += (v73 & 128)
                                                                                                            }
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                        }
                                                                                                        var v81 string = v3
                                                                                                        var v82 uint32 = 78
                                                                                                        _ = v81
                                                                                                        _ = v82
                                                                                                    }
                                                                                                    switch fi65 {
                                                                                                    case -291:
                                                                                                        v3 = "\n\t"
                                                                                                    case 32767:
                                                                                                        var v83 bool = true
                                                                                                        _ = v83
                                                                                                    default:
                                                                                                        switch (v3 + "a") {
                                                                                                        case "foo bar":
                                                                                                            var v84 uint16 = 0
                                                                                                            _ = v84
                                                                                                        case "":
                                                                                                            v85 := GenS5{F0: 0}
                                                                                                            _ = v85
                                                                                                        }
                                                                                                    }
                                                                                                    var v86 uint64 = 2
                                                                                                    _ = v86
                                                                                                }
                                                                                            } else {
                                                                                                v2 += (v2 ^ (fi65 / ((127) | 1)))
                                                                                                v87 := GenS4{F0: 42, F1: -128}
                                                                                                var v88 uint64 = 4294967295
                                                                                                _ = v87
                                                                                                _ = v88
                                                                                            }
                                                                                            _ = v75
                                                                                        } else {
                                                                                            v89 := GenS1{F0: (-(fi44))}
                                                                                            switch v2 {
                                                                                            case -276:
                                                                                                var v90 int16 = -32768
                                                                                                _ = v90
                                                                                            case 782:
                                                                                                v41 = v41
                                                                                            default:
                                                                                                v3 = (("" + "") + "foo bar")
                                                                                            }
                                                                                            _ = v89
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v91 uint32 = 1
                                                                                        _ = v91
                                                                                    case "a":
                                                                                        v1 = append(v1, (v2 | -9223372036854775808))
                                                                                    }
                                                                                    _ = v74
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (v2 / (((v2 & (v2 / (((fi44 / (((-(fi44))) | 1))) | 1)))) | 1))
                                                                                }
                                                                            }
                                                                            v73 += 42
                                                                            var v92 uint16 = 1
                                                                            _ = v92
                                                                        }
                                                                        v93 := GenS3{F0: 255, F1: 473, F2: false}
                                                                        switch v3 {
                                                                        case "":
                                                                            var v94 int8 = 42
                                                                            _ = v94
                                                                        case "hello":
                                                                            var v95 uint16 = 0
                                                                            _ = v95
                                                                        case "\n\t":
                                                                            switch (v3 + ("" + (v3 + "a"))) {
                                                                            case "\n\t":
                                                                                fmt.Println(fi44, v3)
                                                                            case "hello":
                                                                                v1 = append(v1, -435)
                                                                            case "foo bar":
                                                                                v41 -= 1
                                                                            }
                                                                        }
                                                                        v96 := GenS9{F0: 1}
                                                                        _ = v93
                                                                        _ = v96
                                                                    case 127:
                                                                        v73 -= v73
                                                                    case -9223372036854775808:
                                                                        if true {
                                                                            v97 := GenS7{F0: v2}
                                                                            v98 := GenS5{F0: 1000}
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                    default:
                                                                        v73 -= (v73 + 255)
                                                                    }
                                                                case "hello":
                                                                    v73 = 0
                                                                default:
                                                                    if (fi44 <= (fi65 >> ((4294967295) & 63))) {
                                                                        v99 := GenS10{F0: -365}
                                                                        _ = v99
                                                                    } else {
                                                                        v3 = v3
                                                                        v3 = v3
                                                                        v41 += v41
                                                                        var v100 int16 = 42
                                                                        _ = v100
                                                                    }
                                                                    v2 -= fi65
                                                                    var v101 int16 = -495
                                                                    _ = v101
                                                                }
                                                                _ = v73
                                                            }
                                                        case 2147483647:
                                                            v3 = (v3 + ("\n\t" + (v3 + v3)))
                                                        case -886:
                                                            v102 := GenS2{F0: 4294967295, F1: 0}
                                                            _ = v102
                                                        }
                                                        v2 += (fi65 & (v2 ^ fi65))
                                                        if false {
                                                            v41 += genFn1(2147483648, ((!true) && true))
                                                            if false {
                                                                if (true && true) {
                                                                    switch fi44 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(fi44, v41)
                                                                    case -1:
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v2 += v2
                                                                    case -902:
                                                                        v103 := GenS11{F0: ((v3 + v3) + v3), F1: 255}
                                                                        _ = v103
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v104 := GenS12{F0: 1, F1: "hello"}
                                                    _ = v104
                                                }
                                            default:
                                                var v105 string = v3
                                                _ = v105
                                            }
                                            if true {
                                                switch v2 {
                                                case -9223372036854775808:
                                                    v41 = genFn1(42, (!(true && true)))
                                                }
                                            }
                                            fmt.Println(v41, v3, v2)
                                            switch v2 {
                                            case 2147483647:
                                                fmt.Println(v41, fi44)
                                            case -2147483648:
                                                v106 := GenS9{F0: 303}
                                                _ = v106
                                            default:
                                                var v107 int32 = int32(v41)
                                                _ = v107
                                            }
                                        } else {
                                            fmt.Println(fi44)
                                            if (!false) {
                                                v2 = v2
                                                if false {
                                                    fmt.Println(v2)
                                                } else {
                                                    fmt.Println(fi65, v41)
                                                    v2 = -1
                                                    v41 -= (v41 ^ v41)
                                                }
                                            }
                                            v108 := GenS13{F0: v41}
                                            v41 += -1
                                            var v109 bool = (v3 > ((v3 + "") + v3))
                                            _ = v108
                                            _ = v109
                                        }
                                    }
                                    v110 := GenS14{F0: 0}
                                    switch (v2 / (((fi65 - fi65)) | 1)) {
                                    case 42:
                                        switch v2 {
                                        case -9223372036854775808:
                                            if len(v1) > 0 {
                                                v1[0] = -521
                                            }
                                        default:
                                            v111 := GenS15{F0: (float32((1.5 * 1e10)) > -2.25)}
                                            _ = v111
                                        }
                                    }
                                    _ = fi65
                                    _ = v110
                                }
                                v2 += v2
                                var v112 uint32 = 318
                                _ = v112
                            } else {
                                v3 = (v3 + v3)
                            }
                            for fi113 := 0; fi113 < 2; fi113++ {
                                if true {
                                    if (242 != 1) {
                                        switch fi44 {
                                        case 128:
                                            v41 -= v41
                                        case 1:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        switch v3 {
                                        case "\n\t":
                                            var v114 uint64 = genFn0()
                                            _ = v114
                                        case "foo bar":
                                            v2 += (fi113 >> ((42) & 63))
                                        case "a":
                                            v115 := GenS2{F0: genFn0(), F1: 42}
                                            _ = v115
                                        }
                                        v2 += fi44
                                    } else {
                                        fmt.Println(v3, v2)
                                        var v116 uint16 = 255
                                        var v117 float32 = 1.5
                                        _ = v116
                                        _ = v117
                                    }
                                }
                                _ = fi113
                            }
                        }
                    }
                case "foo bar":
                    if (true || true) {
                        v3 = ((v3 + ((v3 + "foo bar") + v3)) + v3)
                        var v118 int32 = -1
                        var v119 float64 = float64(1e10)
                        _ = v118
                        _ = v119
                    }
                }
                _ = fi44
            }
        }
        v120 := GenS4{F0: -356, F1: -128}
        _ = v120
    }
    switch v3 {
    case "\n\t":
        v121 := GenS11{F0: "foo bar", F1: 255}
        _ = v121
    default:
        if false {
            for fi122 := 0; fi122 < 7; fi122++ {
                if true {
                    v2 += v2
                }
                _ = len(v1)
                var v123 int8 = 127
                _ = fi122
                _ = v123
            }
        } else {
            for fi124 := 0; fi124 < 19; fi124++ {
                fmt.Println(v3, v2, v3)
                _ = fi124
            }
            var v125 uint64 = 0
            _ = v125
        }
    }
    for fi126 := 0; fi126 < 6; fi126++ {
        for fi127 := 0; fi127 < 4; fi127++ {
            v3 = "a"
            var v128 int = (fi127 << ((1) & 63))
            if len(v1) > 0 {
                v1[0] = (^(fi127))
            }
            if (!false) {
                v128 += (fi127 >> ((4294967295) & 63))
                v1 = append(v1, (fi127 | v2))
            } else {
                if len(v1) > 0 {
                    v1[0] = (fi126 / ((v2) | 1))
                }
                if true {
                    v41 = v41
                    if (2 == 255) {
                        v1 = append(v1, 556)
                        var v129 uint32 = 65535
                        _ = v129
                    }
                    v3 = ("a" + "a")
                } else {
                    v130 := GenS10{F0: 42}
                    var v131 int = v128
                    _ = len(v1)
                    _ = v130
                    _ = v131
                }
            }
            var v132 int = 1
            v132 = 1
            fmt.Println(v128, v3)
            var v133 int32 = 0
            _ = fi127
            _ = v128
            _ = v132
            _ = v133
        }
        v41 = -9223372036854775808
        _ = fi126
    }
    var v134 uint16 = 2
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v134
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
