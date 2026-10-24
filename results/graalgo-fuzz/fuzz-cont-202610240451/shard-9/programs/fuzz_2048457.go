package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 string; F1 uint8; F2 string }
type GenS4 struct { F0 int8; F1 uint64; F2 string }
type GenS5 struct { F0 uint; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint32; F1 bool; F2 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 uint8; F2 int16 }
func genFn0(p0 uint8, p1 uint8) string {
    var v4 uint8 = p1
    var v5 uint32 = 368
    if true {
        v6 := GenS1{F0: 4294967295, F1: false}
        _ = v6
    } else {
        v7 := []int{127, -9223372036854775808, -9223372036854775808}
        _ = v7
    }
    if (-507 != 2147483647) {
        p1 = (v4 * 128)
        if false {
            var v8 uint64 = 408
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint16 = 2
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := []int{374}
                    if true {
                        v13 := GenS2{F0: (v8 >> ((828) & 63)), F1: uint64(v10), F2: (1e10 <= -2.25)}
                        _ = v13
                    }
                    var v14 uint16 = (^(v10))
                    _ = fi11
                    _ = v12
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        switch ("foo bar" + "") {
        case "":
            var v15 uint8 = (p0 ^ 128)
            _ = v15
        case "\n\t":
            v5 -= v5
        default:
            if true {
                v4 -= 255
            }
        }
    }
    v16 := []int{}
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    _ = v5
    return ""
}
func genFn1() uint8 {
    for fi17 := 0; fi17 < 12; fi17++ {
        switch ("hello" + "") {
        case "\n\t":
            v18 := []int{(fi17 - 1)}
            _ = v18
        case "":
            var v19 int32 = 920
            _ = v19
        }
        switch ("\n\t" + "hello") {
        case "":
            if (4294967295 != 22) {
                switch ("hello" + "") {
                case "a":
                    v20 := []int{fi17, 127}
                    _ = v20
                case "hello":
                    v21 := []int{fi17, (-(fi17)), fi17}
                    _ = v21
                case "foo bar":
                    v22 := []int{fi17}
                    _ = v22
                }
            }
        default:
            for fi23 := 0; fi23 < 1; fi23++ {
                if (0 >= 42) {
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        v24 := []int{(fi23 >> ((255) & 63)), fi17, fi17, -2147483648}
                        _ = v24
                    case "":
                        fmt.Println(fi23)
                    default:
                        v25 := []int{(^(fi17)), fi23}
                        _ = v25
                    }
                    var v26 float64 = 3.14159
                    _ = v26
                } else {
                    var v27 string = ""
                    _ = v27
                }
                _ = fi23
            }
        }
        switch (^(fi17)) {
        case 42:
            fmt.Println(fi17)
        default:
            if (true || false) {
                var v28 int64 = -9223372036854775808
                _ = v28
            } else {
                v29 := GenS2{F0: 42, F1: 1, F2: true}
                v30 := []int{fi17, fi17, -9223372036854775808}
                if true {
                    v31 := GenS3{F0: ("a" + "foo bar"), F1: 255, F2: (("a" + "") + ("hello" + "hello"))}
                    switch "hello" {
                    case "foo bar":
                        v32 := GenS2{F0: 0, F1: 2, F2: false}
                        _ = v32
                    case "a":
                        for fi33 := 0; fi33 < 5; fi33++ {
                            fmt.Println(fi33, fi17)
                            _ = fi33
                        }
                    case "\n\t":
                        var v34 int64 = -9223372036854775808
                        _ = v34
                    }
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    _ = v31
                }
                _ = v29
                _ = v30
            }
        }
        _ = fi17
    }
    var v35 int16 = 1000
    _ = v35
    return 255
}

func main() {
    var v1 int16 = 0
    var v2 int32 = -2147483648
    v3 := []int{}
    if ((true || false) || (true && true)) {
        fmt.Println(v2, v2)
        var v36 float64 = 3.14159
        for fi37 := 0; fi37 < 8; fi37++ {
            switch ("foo bar" + "\n\t") {
            case "hello":
                v3 = append(v3, (fi37 * 42))
            }
            _ = fi37
        }
        _ = v36
    }
    v3 = append(v3, -375)
    if false {
        v1 = v1
        v2 += -1
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if (-9223372036854775808 == 0) {
        if (false || false) {
            for fi38 := 0; fi38 < 15; fi38++ {
                if (0 > 4294967295) {
                    for fi39 := 0; fi39 < 8; fi39++ {
                        if (1e10 != (3.14159 / 0.1)) {
                            fmt.Println(fi38)
                            switch (("hello" + "hello") + "") {
                            case "\n\t":
                                v40 := GenS2{F0: 1, F1: 1, F2: true}
                                _ = v40
                            case "foo bar":
                                v41 := GenS2{F0: 549, F1: 606, F2: (!("hello" <= "hello"))}
                                var v42 float64 = -2.25
                                _ = v41
                                _ = v42
                            case "":
                                fmt.Println(fi38, fi38, v1)
                            }
                            v43 := GenS4{F0: -41, F1: 42, F2: ("\n\t" + "a")}
                            switch fi38 {
                            case 127:
                                var v44 bool = (true || true)
                                _ = v44
                            }
                            v2 = v2
                            var v45 int8 = -93
                            _ = v43
                            _ = v45
                        }
                        if len(v3) > 0 {
                            v3[0] = (fi39 & 9223372036854775807)
                        }
                        if len(v3) > 0 {
                            v3[0] = fi38
                        }
                        if (!true) {
                            v1 = -908
                        } else {
                            var v46 int32 = (-(v2))
                            _ = v46
                        }
                        v1 += v1
                        _ = fi39
                    }
                } else {
                    if true {
                        var v47 uint64 = 2
                        v47 = v47
                        _ = v47
                    }
                    v48 := GenS5{F0: 18446744073709551615, F1: v1, F2: 2}
                    for fi49 := 0; fi49 < 2; fi49++ {
                        v1 -= (v1 | (v1 % ((32767) | 1)))
                        _ = fi49
                    }
                    _ = v48
                }
                for fi50 := 0; fi50 < 20; fi50++ {
                    if false {
                        var v51 int8 = int8(v2)
                        _ = v51
                    } else {
                        switch (fi38 & (fi38 / ((fi50) | 1))) {
                        case 0:
                            var v52 int16 = (v1 + v1)
                            _ = v52
                        default:
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v2 = -505
                                switch "hello" {
                                case "":
                                    fmt.Println(fi50, fi50)
                                case "\n\t":
                                    if (887 > v2) {
                                        v1 -= 380
                                        var v53 string = ""
                                        var v54 int = fi50
                                        _ = v53
                                        _ = v54
                                    }
                                case "a":
                                    fmt.Println(fi38, v2)
                                }
                            }
                        }
                        v2 = v2
                        if (!true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (fi50 != (fi38 - 32767)) {
                                var v55 uint = 0
                                _ = v55
                            } else {
                                var v56 float32 = 0.0
                                switch fi38 {
                                case 42:
                                    v56 -= (-2.25 / ((float32((-2.25 * 0.1)) - (1e10 - v56)) / v56))
                                case -165:
                                    var v57 int = -2147483648
                                    switch ("\n\t" + "hello") {
                                    case "\n\t":
                                        v58 := GenS6{F0: 1, F1: 639}
                                        _ = v58
                                    case "hello":
                                        var v59 uint64 = 0
                                        v60 := GenS2{F0: v59, F1: v59, F2: (0.0 > (1e10 / (-2.25 + float64(((((float32(3.14159) + v56) / (0.1 - v56)) * v56) / (-2.25 - v56))))))}
                                        if (255 < 2) {
                                            if (-589 != v2) {
                                                switch ("\n\t" + (((((("" + "") + "\n\t") + ("foo bar" + "hello")) + "hello") + "hello") + (((((("hello" + "") + "") + ("hello" + ("" + "hello"))) + "a") + ("hello" + "\n\t")) + "a"))) {
                                                case "":
                                                    switch (("\n\t" + "a") + (("foo bar" + "") + ("a" + ((("\n\t" + ("\n\t" + ("hello" + "foo bar"))) + "\n\t") + "")))) {
                                                    case "hello":
                                                        v1 = v1
                                                    }
                                                case "foo bar":
                                                    if true {
                                                        var v61 uint = 1
                                                        _ = v61
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi38 | (fi50 << ((2) & 63)))
                                                        }
                                                    }
                                                }
                                                v57 = -750
                                                var v62 float32 = v56
                                                if false {
                                                    v57 -= fi38
                                                    var v63 string = (("hello" + ("" + "hello")) + "a")
                                                    fmt.Println(fi50, fi50)
                                                    _ = v63
                                                } else {
                                                    v2 += (v2 - -2147483648)
                                                }
                                                if false {
                                                    _ = len(v3)
                                                    if (1 >= 0) {
                                                        if false {
                                                            var v64 int = fi38
                                                            var v65 int8 = -128
                                                            _ = len(v3)
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            v57 += (v57 & int(v1))
                                                            switch (("foo bar" + ("" + "hello")) + ("a" + "\n\t")) {
                                                            case "foo bar":
                                                                if (false && ((v1 + (v1 ^ int16(v59))) != (v1 ^ -1))) {
                                                                    v57 = v57
                                                                    fmt.Println(v2, fi50)
                                                                    v2 = v2
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (-(v57))
                                                                    }
                                                                    v2 = v2
                                                                }
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        v57 = (fi38 - fi50)
                                                        fmt.Println(v62, fi38, fi38)
                                                    } else {
                                                        v56 = v56
                                                    }
                                                } else {
                                                    v1 = v1
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v62
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                v66 := GenS7{F0: 2, F1: (1.5 <= ((0.0 * 0.0) + (3.14159 / (1e10 / -2.25)))), F2: 0}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v66
                                            case "":
                                                switch ("" + "a") {
                                                case "":
                                                    if true {
                                                        _ = len(v3)
                                                    }
                                                default:
                                                    switch v57 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v57, v56, v2)
                                                    case -96:
                                                        if false {
                                                            v59 += v59
                                                        } else {
                                                            var v67 uint16 = 42
                                                            _ = v67
                                                        }
                                                    default:
                                                        var v68 int8 = 42
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v68
                                                    }
                                                }
                                            case "\n\t":
                                                if (!false) {
                                                    v59 = 1
                                                    v57 += fi38
                                                    var v69 uint8 = genFn1()
                                                    _ = v69
                                                }
                                            default:
                                                var v70 int32 = v2
                                                _ = v70
                                            }
                                            v59 = (v59 % (((v59 + (v59 ^ v59))) | 1))
                                            if ((0 > 2) && false) {
                                                v56 -= ((((v56 - ((-2.25 * 1.5) * v56)) - v56) + -2.25) + (v56 / 1.0))
                                            }
                                        }
                                        _ = v59
                                        _ = v60
                                    default:
                                        v57 = fi50
                                    }
                                    fmt.Println(v2, fi50, v2)
                                    _ = v57
                                default:
                                    if true {
                                        var v71 int8 = -124
                                        var v72 int8 = (-(v71))
                                        if true {
                                            v72 = v72
                                            v71 = (v71 ^ int8(fi50))
                                            fmt.Println(v72, v1)
                                            if true {
                                                v56 += v56
                                                v71 = -128
                                                if false {
                                                    if false {
                                                        if false {
                                                            if true {
                                                                switch 32768 {
                                                                case -653:
                                                                    switch fi50 {
                                                                    case 42:
                                                                        fmt.Println(v56, v72)
                                                                    case 127:
                                                                        v71 = (v72 * v71)
                                                                    case 32768:
                                                                        fmt.Println(fi50, v2)
                                                                        if ((v72 > int8(v1)) && true) {
                                                                            var v73 float64 = 3.14159
                                                                            _ = v73
                                                                        }
                                                                    default:
                                                                        v3 = append(v3, 2147483647)
                                                                    }
                                                                case 128:
                                                                    if true {
                                                                        v56 = 1e10
                                                                        var v74 int8 = v71
                                                                        _ = v74
                                                                    } else {
                                                                        switch fi50 {
                                                                        case 1:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = fi50
                                                                            }
                                                                        case 401:
                                                                            v71 = 127
                                                                        default:
                                                                            switch fi50 {
                                                                            case 4294967295:
                                                                                v72 = (v71 * (v72 + (-(v71))))
                                                                            }
                                                                        }
                                                                        switch (("hello" + "") + "a") {
                                                                        case "a":
                                                                            var v75 int16 = (v1 / (((v1 << ((255) & 63))) | 1))
                                                                            _ = v75
                                                                        case "\n\t":
                                                                            v2 = v2
                                                                        }
                                                                        var v76 bool = (!false)
                                                                        _ = v76
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v1)
                                                        } else {
                                                            v72 += 1
                                                            v77 := GenS3{F0: "a", F1: 1, F2: ""}
                                                            var v78 int = -2147483648
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v79 int = (fi38 >> ((377) & 63))
                                                            var v80 uint = 255
                                                            v3 = append(v3, -1)
                                                            _ = v79
                                                            _ = v80
                                                        case "a":
                                                            var v81 int = (fi38 / ((127) | 1))
                                                            _ = v81
                                                        case "hello":
                                                            var v82 uint16 = 65535
                                                            _ = v82
                                                        default:
                                                            v72 -= (v71 / ((v71) | 1))
                                                        }
                                                    }
                                                    v71 = 42
                                                    var v83 uint16 = 443
                                                    _ = v83
                                                }
                                                v56 = v56
                                                var v84 float64 = 1e10
                                                _ = v84
                                            }
                                        } else {
                                            v72 -= v72
                                        }
                                        _ = v71
                                        _ = v72
                                    }
                                }
                                if (9223372036854775807 <= 9223372036854775807) {
                                    if (float32((1.5 - (((0.1 * (3.14159 + 3.14159)) / float64(1.5)) - 0.1))) > 1e10) {
                                        _ = len(v3)
                                        if false {
                                            v1 = (v1 ^ v1)
                                        }
                                        _ = len(v3)
                                    } else {
                                        if (true || ((v2 ^ v2) != 1)) {
                                            var v85 string = ("" + ("" + ""))
                                            var v86 uint = 255
                                            v85 = ""
                                            _ = v85
                                            _ = v86
                                        } else {
                                            fmt.Println(fi50)
                                        }
                                    }
                                    if (false && false) {
                                        if true {
                                            if (("hello" + "hello") <= (((("hello" + "a") + "foo bar") + "foo bar") + ("foo bar" + ""))) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v87 := GenS8{F0: 65535}
                                                    _ = v87
                                                default:
                                                    if ("a" > "hello") {
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi38 + fi38)
                                                        }
                                                    }
                                                }
                                            } else {
                                                v2 += v2
                                            }
                                            var v88 uint32 = 0
                                            _ = v88
                                        } else {
                                            if (65535 != 660) {
                                                _ = len(v3)
                                                if (!false) {
                                                    var v89 int = (fi38 >> ((255) & 63))
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v90 int64 = 0
                                                        _ = v90
                                                    }
                                                    v56 = (0.1 - (-2.25 + v56))
                                                    _ = v89
                                                } else {
                                                    var v91 bool = (!(false && true))
                                                    if false {
                                                        v2 = 146
                                                    }
                                                    v1 = 921
                                                    _ = v91
                                                }
                                                if true {
                                                    var v92 uint16 = 42
                                                    var v93 float64 = -2.25
                                                    v92 = 154
                                                    _ = v92
                                                    _ = v93
                                                }
                                            }
                                        }
                                    }
                                    var v94 int32 = (v2 * (-(v2)))
                                    _ = v94
                                } else {
                                    var v95 uint64 = 0
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if (false || true) {
                                        var v96 int = 4294967295
                                        _ = v96
                                    } else {
                                        if ((!(true && false)) || (0.0 <= 0.0)) {
                                            fmt.Println(v1, v56, v1)
                                            if (-128 != -100) {
                                                v2 = (^(v2))
                                                v3 = append(v3, fi50)
                                            }
                                            if ((v1 > 1000) && (-267 > (-(v1)))) {
                                                if true {
                                                    v2 += v2
                                                } else {
                                                    switch fi50 {
                                                    case -908:
                                                        var v97 uint64 = v95
                                                        _ = v97
                                                    }
                                                    var v98 int = int(v1)
                                                    fmt.Println(fi50, v1)
                                                    _ = v98
                                                }
                                                v95 -= v95
                                            }
                                            fmt.Println(fi50, v1, fi38)
                                            var v99 int8 = 42
                                            _ = v99
                                        }
                                    }
                                    _ = v95
                                }
                                _ = v56
                            }
                            var v100 float32 = 0.1
                            fmt.Println(fi50, fi38, v2)
                            if false {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v100)
                                }
                                var v101 uint8 = 255
                                var v102 float64 = -2.25
                                _ = v101
                                _ = v102
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v1 = v1
                            }
                            v103 := GenS5{F0: 522, F1: 32767, F2: 65535}
                            if ((-1 != 127) || false) {
                                var v104 int16 = v1
                                switch (fi38 * fi38) {
                                case 393:
                                    var v105 string = "hello"
                                    _ = v105
                                default:
                                    v104 = v1
                                }
                                _ = v104
                            } else {
                                fmt.Println(v2, fi50, fi50)
                            }
                            _ = v100
                            _ = v103
                        } else {
                            v2 += v2
                            v1 = v1
                            var v106 float64 = ((((3.14159 / ((-2.25 - (((0.0 + (0.1 / 1.0)) - 1.5) / (3.14159 + 1e10))) + 1e10)) * ((0.1 / 1.0) - 0.0)) / (0.1 / ((1e10 * 1e10) - float64(float32(0.0))))) / 1.0)
                            _ = v106
                        }
                    }
                    fmt.Println(v2, fi50)
                    if len(v3) > 0 {
                        v3[0] = fi50
                    }
                    _ = fi50
                }
                _ = len(v3)
                _ = fi38
            }
        }
    }
    v2 = (v2 % ((int32(v1)) | 1))
    v2 = (v2 | v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi107 := 0; fi107 < 1; fi107++ {
        v108 := GenS9{F0: ("foo bar" + "\n\t"), F1: 255, F2: v1}
        fmt.Println(fi107)
        _ = fi107
        _ = v108
    }
    var v109 int = 712
    var v110 bool = (1 > 240)
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v109
    _ = v110
    _ = v2
    _ = v3
}
