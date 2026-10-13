package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int64; F1 uint16; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 string; F2 uint32 }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 bool; F1 uint32; F2 bool }
type GenS13 struct { F0 uint8; F1 string; F2 int16 }
type GenS14 struct { F0 uint8; F1 int32; F2 uint }
type GenS15 struct { F0 int; F1 int16 }
func genFn0(p0 int16, p1 bool, p2 uint64) uint {
    p2 = (p2 | (p2 | (p2 % ((p2) | 1))))
    var v7 uint = 4294967295
    switch 0 {
    case 32768:
        p0 -= p0
    case 127:
        var v8 int16 = (p0 % (((p0 * -32768)) | 1))
        _ = v8
    default:
        p0 += 1000
    }
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}
func genFn1(p0 uint32, p1 uint32, p2 int16) uint {
    var v9 int32 = -1
    if false {
        fmt.Println(v9, v9, v9)
    } else {
        var v10 uint32 = (p1 << ((521) & 63))
        var v11 int = -9223372036854775808
        fmt.Println(p0)
        _ = v10
        _ = v11
    }
    var v12 uint32 = (^(p1))
    switch "a" {
    case "a":
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 158
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{0}
    var v3 float64 = 0.0
    var v4 uint8 = 245
    var v5 int64 = 736
    var v6 int8 = -83
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v13 int64 = 570
        v14 := GenS1{F0: 2147483647, F1: v4, F2: (v6 >> ((42) & 63))}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v5, v6)
        v16 := GenS1{F0: 947, F1: v4, F2: -45}
        switch (-(fi15)) {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 32768:
            switch 0 {
            case 4294967295:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if (v3 == v3) {
                    var v17 int8 = 42
                    v2 = append(v2, 798)
                    _ = v17
                }
            }
        case -9223372036854775808:
            switch (-(fi15)) {
            case 128:
                var v18 bool = (true || true)
                _ = v18
            }
        default:
            v19 := GenS1{F0: 997, F1: (^(v4)), F2: (v6 / (((v6 * -128)) | 1))}
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    v4 = (^(v4))
    if false {
        var v20 string = "\n\t"
        if (!((int(v1) < 9223372036854775807) || (0 >= v4))) {
            v21 := GenS1{F0: 4294967295, F1: v4, F2: v6}
            _ = v21
        }
        v1 = 1
        _ = v20
    }
    if (!true) {
        var v22 bool = false
        _ = v22
    } else {
        var v23 int8 = (v6 % (((v6 * (v6 + int8(v5)))) | 1))
        var v24 string = ("" + "\n\t")
        var v25 bool = false
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 float64 = v3
    var v27 bool = (!true)
    if ((v1 >> ((1) & 63)) > int32(v6)) {
        var v28 uint64 = 1
        v29 := GenS1{F0: 2147483647, F1: v4, F2: v6}
        _ = v28
        _ = v29
    }
    if (4294967295 < 843) {
        for fi30 := 0; fi30 < 15; fi30++ {
            v26 = (((float64((1.5 / 1.5)) - v26) + float64((((1.5 / 3.14159) + (1e10 + 0.1)) - (-2.25 / 1e10)))) - (0.1 * 1e10))
            for fi31 := 0; fi31 < 13; fi31++ {
                switch 128 {
                case 1000:
                    v32 := GenS2{F0: 0}
                    _ = v32
                case 4294967295:
                    fmt.Println(v1, v5)
                }
                v33 := GenS1{F0: 42, F1: 127, F2: 0}
                _ = fi31
                _ = v33
            }
            _ = fi30
        }
        if (65535 == 866) {
            if (v27 && v27) {
                if (true || (4294967295 < -62)) {
                    switch "foo bar" {
                    case "foo bar":
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 uint64 = 18446744073709551615
                            switch fi34 {
                            case 128:
                                var v36 int32 = (v1 >> ((0) & 63))
                                _ = v36
                            }
                            v35 -= 710
                            _ = fi34
                            _ = v35
                        }
                    case "\n\t":
                        var v37 float32 = 1.5
                        _ = v37
                    default:
                        v38 := GenS3{F0: (v1 & v1), F1: true}
                        _ = v38
                    }
                } else {
                    v6 = v6
                    for fi39 := 0; fi39 < 18; fi39++ {
                        if len(v2) > 0 {
                            v2[0] = 938
                        }
                        if (0.1 <= (0.0 / 1.0)) {
                            switch (fi39 * fi39) {
                            case 1:
                                if ((true || (v27 || ((v4 + (v4 + 2)) < 0))) && false) {
                                    var v40 uint16 = 0
                                    if true {
                                        if v27 {
                                            var v41 int32 = (v1 | v1)
                                            if false {
                                                v42 := GenS2{F0: (v5 / (((v5 % ((v5) | 1))) | 1))}
                                                _ = v42
                                            } else {
                                                var v43 uint32 = 462
                                                if false {
                                                    var v44 int8 = (v6 + 42)
                                                    var v45 uint32 = uint32(v40)
                                                    v44 = (v44 + (v6 / (((^(v44))) | 1)))
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    for fi46 := 0; fi46 < 14; fi46++ {
                                                        v40 -= (v40 % ((2) | 1))
                                                        _ = fi46
                                                    }
                                                }
                                                if v27 {
                                                    switch "hello" {
                                                    case "":
                                                        v4 -= (v4 ^ (v4 ^ (v4 + (v4 ^ 255))))
                                                    case "foo bar":
                                                        if len(v2) > 0 {
                                                            v2[0] = v43
                                                        }
                                                    case "hello":
                                                        var v47 int64 = (v5 % (((^(v5))) | 1))
                                                        _ = v47
                                                    }
                                                } else {
                                                    if true {
                                                        switch (((("\n\t" + "\n\t") + "a") + ("a" + "hello")) + "\n\t") {
                                                        case "":
                                                            for fi48 := 0; fi48 < 8; fi48++ {
                                                                var v49 bool = false
                                                                v50 := GenS1{F0: v43, F1: v4, F2: int8(v43)}
                                                                var v51 int8 = 1
                                                                _ = fi48
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        }
                                                    }
                                                    v1 += (v41 / ((943) | 1))
                                                    for fi52 := 0; fi52 < 12; fi52++ {
                                                        if false {
                                                            if len(v2) > 0 {
                                                                v2[0] = v43
                                                            }
                                                        } else {
                                                            var v53 int8 = (v6 + 42)
                                                            _ = v53
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = v43
                                                        }
                                                        switch ("a" + ("a" + "hello")) {
                                                        case "foo bar":
                                                            v54 := GenS1{F0: v43, F1: v4, F2: (v6 << ((4294967295) & 63))}
                                                            _ = v54
                                                        }
                                                        _ = fi52
                                                    }
                                                }
                                                _ = v43
                                            }
                                            v6 += -128
                                            _ = v41
                                        } else {
                                            v55 := GenS4{F0: genFn1(65535, 2, 1), F1: (v5 << ((1000) & 63))}
                                            v56 := GenS4{F0: uint(v4), F1: (v5 + -1)}
                                            switch ("" + "a") {
                                            case "":
                                                var v57 float32 = -2.25
                                                _ = v57
                                            case "foo bar":
                                                var v58 int32 = int32(v40)
                                                _ = v58
                                            case "\n\t":
                                                if (!(21 >= (v6 / ((v6) | 1)))) {
                                                    v27 = v27
                                                    var v59 int = 128
                                                    _ = v59
                                                }
                                            }
                                            _ = v55
                                            _ = v56
                                        }
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch (fi39 / ((fi39) | 1)) {
                                        case 4294967295:
                                            _ = len(v2)
                                        case 128:
                                            v1 = (v1 << ((2) & 63))
                                        case -140:
                                            for fi60 := 0; fi60 < 15; fi60++ {
                                                switch (-(fi60)) {
                                                case 42:
                                                    var v61 uint64 = 1
                                                    _ = v61
                                                case 127:
                                                    v62 := GenS5{F0: 475, F1: v40, F2: (fi39 >> ((genFn0(-1, false, 4294967295)) & 63))}
                                                    _ = v62
                                                case -665:
                                                    v4 += v4
                                                }
                                                _ = fi60
                                            }
                                        }
                                        var v63 uint32 = 663
                                        _ = v63
                                    }
                                    _ = v40
                                } else {
                                    for fi64 := 0; fi64 < 16; fi64++ {
                                        var v65 uint = 8
                                        var v66 float32 = 0.0
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            case 42:
                                v67 := GenS2{F0: v5}
                                _ = v67
                            }
                            var v68 int64 = v5
                            _ = v68
                        }
                        _ = fi39
                    }
                    switch "hello" {
                    case "":
                        v69 := GenS6{F0: v4}
                        _ = v69
                    case "foo bar":
                        if false {
                            fmt.Println(v1, v1)
                            v27 = (true && (!v27))
                            var v70 uint32 = 0
                            _ = v70
                        }
                    default:
                        if (!(2147483647 != 127)) {
                            v2 = append(v2, 897)
                            v1 = v1
                            v71 := GenS7{F0: 976}
                            _ = v71
                        } else {
                            switch 4294967295 {
                            case 1:
                                var v72 bool = ((-(v1)) != v1)
                                _ = v72
                            case -9223372036854775808:
                                if (482 != 0) {
                                    if ((((0.0 * (1.5 + v26)) == (-2.25 - v3)) || v27) || (v4 < (v4 >> ((uint(v4)) & 63)))) {
                                        v6 += v6
                                        if v27 {
                                            switch ("\n\t" + "hello") {
                                            case "":
                                                var v73 int16 = 42
                                                _ = v73
                                            case "hello":
                                                switch ("a" + "\n\t") {
                                                case "a":
                                                    v1 += 1
                                                case "foo bar":
                                                    switch 113 {
                                                    case 9223372036854775807:
                                                        if false {
                                                            v3 -= v26
                                                            v1 = v1
                                                        } else {
                                                            if (2 < 0) {
                                                                v74 := GenS8{F0: false, F1: 32768}
                                                                _ = v74
                                                            }
                                                        }
                                                    case 127:
                                                        if ((!(-1 < -666)) || (255 <= 958)) {
                                                            v75 := GenS9{F0: 216}
                                                            var v76 uint64 = 447
                                                            switch 128 {
                                                            case 0:
                                                                if false {
                                                                    _ = len(v2)
                                                                    if (v6 > v6) {
                                                                        v27 = v27
                                                                        var v77 int64 = 9223372036854775807
                                                                        if (!(v27 || v27)) {
                                                                            var v78 int8 = v6
                                                                            _ = v78
                                                                        } else {
                                                                            v79 := GenS10{F0: 104, F1: (("hello" + ((("a" + "a") + "hello") + "a")) + ""), F2: 65535}
                                                                            for fi80 := 0; fi80 < 8; fi80++ {
                                                                                for fi81 := 0; fi81 < 16; fi81++ {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 638
                                                                                    }
                                                                                    if false {
                                                                                        v82 := GenS5{F0: (v5 & (v77 & (v77 + (v5 % (((v5 | v77)) | 1))))), F1: 65535, F2: fi80}
                                                                                        _ = v82
                                                                                    }
                                                                                    fmt.Println(v3, v6, v26)
                                                                                    _ = fi81
                                                                                }
                                                                                fmt.Println(v4)
                                                                                _ = fi80
                                                                            }
                                                                            if false {
                                                                                switch "" {
                                                                                case "":
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        switch 1 {
                                                                                        case 127:
                                                                                            var v83 uint = 60
                                                                                            _ = v83
                                                                                        default:
                                                                                            v26 += (1e10 * 0.0)
                                                                                        }
                                                                                    case "a":
                                                                                        v4 -= (v4 >> ((719) & 63))
                                                                                    default:
                                                                                        switch (("hello" + "hello") + ("\n\t" + ("foo bar" + "a"))) {
                                                                                        case "a":
                                                                                            for fi84 := 0; fi84 < 11; fi84++ {
                                                                                                fmt.Println(v6)
                                                                                                _ = fi84
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v85 int = 4294967295
                                                                                            _ = v85
                                                                                        case "foo bar":
                                                                                            fmt.Println(v4, v5)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v3 += v3
                                                                            }
                                                                            _ = v79
                                                                        }
                                                                        _ = v77
                                                                    } else {
                                                                        fmt.Println(v76, v1)
                                                                        var v86 uint16 = 2
                                                                        v87 := GenS8{F0: v27, F1: 128}
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                    v4 = v4
                                                                }
                                                            case 42:
                                                                for fi88 := 0; fi88 < 1; fi88++ {
                                                                    v27 = (("foo bar" + "a") >= (("hello" + "foo bar") + (("foo bar" + "hello") + (("hello" + "\n\t") + "foo bar"))))
                                                                    _ = fi88
                                                                }
                                                            case 240:
                                                                for fi89 := 0; fi89 < 1; fi89++ {
                                                                    switch (fi89 * (fi89 & (^(fi89)))) {
                                                                    case 9223372036854775807:
                                                                        fmt.Println(v27, v76, v4)
                                                                    default:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    for fi90 := 0; fi90 < 12; fi90++ {
                                                                        if v27 {
                                                                            switch "" {
                                                                            case "":
                                                                                if (("a" + "") > ("foo bar" + "foo bar")) {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    fmt.Println(v6)
                                                                                } else {
                                                                                    v6 = 78
                                                                                    if (1 == 255) {
                                                                                        if false {
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                var v91 string = ("hello" + "\n\t")
                                                                                                switch fi89 {
                                                                                                case 128:
                                                                                                    var v92 int = 823
                                                                                                    _ = v92
                                                                                                case 42:
                                                                                                    fmt.Println(v76, fi90, v26)
                                                                                                case -803:
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = 150
                                                                                                    }
                                                                                                default:
                                                                                                    v93 := GenS8{F0: false, F1: fi90}
                                                                                                    _ = v93
                                                                                                }
                                                                                                if true {
                                                                                                    if v27 {
                                                                                                        fmt.Println(v91, v6)
                                                                                                        v5 = 9223372036854775807
                                                                                                        if v27 {
                                                                                                            v26 -= 1.5
                                                                                                            _ = len(v2)
                                                                                                        } else {
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = 65535
                                                                                                        }
                                                                                                        if (v5 == 0) {
                                                                                                            var v94 uint8 = 0
                                                                                                            _ = v94
                                                                                                        } else {
                                                                                                            v95 := GenS11{F0: (v1 % ((-1) | 1)), F1: 658}
                                                                                                            if ((v5 >= int64(fi90)) && false) {
                                                                                                                if (-2.25 == 3.14159) {
                                                                                                                    if (1 <= (v4 - v4)) {
                                                                                                                        fmt.Println(v3)
                                                                                                                        v1 += (v1 - v1)
                                                                                                                    }
                                                                                                                    v4 = 2
                                                                                                                } else {
                                                                                                                    var v96 float64 = (v26 / 1.0)
                                                                                                                    var v97 uint64 = (v76 % (((^(v76))) | 1))
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    _ = v96
                                                                                                                    _ = v97
                                                                                                                }
                                                                                                                v26 = (v3 / 1.5)
                                                                                                                v98 := GenS12{F0: (v76 != 145), F1: 2, F2: true}
                                                                                                                _ = v98
                                                                                                            }
                                                                                                            if v27 {
                                                                                                                var v99 float32 = (0.1 * 1.5)
                                                                                                                v100 := GenS13{F0: (v4 * (v4 >> ((1) & 63))), F1: (v91 + v91), F2: -638}
                                                                                                                v2 = append(v2, 345)
                                                                                                                _ = v100
                                                                                                                _ = v99
                                                                                                            } else {
                                                                                                                v4 += uint8(v76)
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v95
                                                                                                        }
                                                                                                    }
                                                                                                    switch v91 {
                                                                                                    case "foo bar":
                                                                                                        v26 = 1e10
                                                                                                    case "\n\t":
                                                                                                        v1 -= (-(v1))
                                                                                                    default:
                                                                                                        if true {
                                                                                                            var v101 float64 = v26
                                                                                                            v102 := GenS4{F0: 255, F1: int64(v4)}
                                                                                                            _ = v101
                                                                                                            _ = v102
                                                                                                        } else {
                                                                                                            var v103 int = 619
                                                                                                            var v104 uint8 = v4
                                                                                                            _ = v103
                                                                                                            _ = v104
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v105 int16 = 32767
                                                                                                    v91 = v91
                                                                                                    _ = v105
                                                                                                }
                                                                                                _ = v91
                                                                                            case "a":
                                                                                                var v106 uint = 84
                                                                                                _ = v106
                                                                                            case "\n\t":
                                                                                                var v107 int = fi90
                                                                                                fmt.Println(v27)
                                                                                                _ = v107
                                                                                            default:
                                                                                                var v108 uint64 = 4294967295
                                                                                                _ = v108
                                                                                            }
                                                                                            v109 := GenS6{F0: (^(v4))}
                                                                                            _ = v109
                                                                                        } else {
                                                                                            switch (("a" + ("\n\t" + (("hello" + ("foo bar" + "foo bar")) + "\n\t"))) + "") {
                                                                                            case "":
                                                                                                fmt.Println(fi90, v3, v26)
                                                                                            case "foo bar":
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 42
                                                                                                }
                                                                                                _ = len(v2)
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            case "\n\t":
                                                                                                fmt.Println(v5, fi90, v27)
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v4, v27)
                                                                                    }
                                                                                }
                                                                                var v110 float32 = 1e10
                                                                                _ = v110
                                                                            }
                                                                        }
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v111 := GenS13{F0: v4, F1: "", F2: -32768}
                                                                            fmt.Println(v5, v4, v26)
                                                                            var v112 int = (-(fi90))
                                                                            _ = v111
                                                                            _ = v112
                                                                        default:
                                                                            v2 = append(v2, 692)
                                                                        }
                                                                        _ = fi90
                                                                    }
                                                                    v1 -= (v1 >> ((18446744073709551615) & 63))
                                                                    _ = fi89
                                                                }
                                                            default:
                                                                var v113 uint32 = 712
                                                                _ = v113
                                                            }
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            var v114 int32 = (v1 ^ v1)
                                                            var v115 float32 = -2.25
                                                            v114 += -1
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                    default:
                                                        fmt.Println(v27)
                                                    }
                                                case "hello":
                                                    v2 = append(v2, 949)
                                                }
                                            case "a":
                                                if false {
                                                    v4 = v4
                                                    v116 := GenS14{F0: v4, F1: -425, F2: 42}
                                                    _ = v116
                                                } else {
                                                    if false {
                                                        for fi117 := 0; fi117 < 15; fi117++ {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v118 int64 = (v5 | (v5 * (v5 ^ (v5 + v5))))
                                                                _ = v118
                                                            }
                                                            var v119 uint16 = 2
                                                            v2 = append(v2, 42)
                                                            _ = fi117
                                                            _ = v119
                                                        }
                                                        var v120 string = ("foo bar" + "hello")
                                                        _ = v120
                                                    } else {
                                                        for fi121 := 0; fi121 < 9; fi121++ {
                                                            switch (fi121 - fi121) {
                                                            case -166:
                                                                var v122 int8 = v6
                                                                _ = v122
                                                            case -269:
                                                                var v123 int64 = (^(v5))
                                                                _ = v123
                                                            case 0:
                                                                v27 = (830 >= 18446744073709551615)
                                                            default:
                                                                v6 -= (v6 * -128)
                                                            }
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v124 uint32 = 929
                                                                _ = v124
                                                            }
                                                            _ = len(v2)
                                                            _ = fi121
                                                        }
                                                        v125 := GenS14{F0: (^(v4)), F1: (v1 + (v1 ^ (v1 & v1))), F2: 0}
                                                        var v126 int8 = v6
                                                        _ = v125
                                                        _ = v126
                                                    }
                                                    v127 := GenS15{F0: -374, F1: -1}
                                                    var v128 int16 = -580
                                                    _ = v127
                                                    _ = v128
                                                }
                                            }
                                            v6 = v6
                                            fmt.Println(v3, v5, v27)
                                        }
                                        for fi129 := 0; fi129 < 17; fi129++ {
                                            var v130 uint = 2
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v130)
                                            _ = fi129
                                            _ = v130
                                        }
                                    }
                                    for fi131 := 0; fi131 < 17; fi131++ {
                                        v4 += (v4 ^ v4)
                                        v4 = 255
                                        _ = fi131
                                    }
                                    v3 += (v26 + 1e10)
                                }
                            case 387:
                                var v132 uint64 = 99
                                _ = v132
                            default:
                                fmt.Println(v3, v27, v27)
                            }
                            _ = len(v2)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                }
                v133 := GenS14{F0: v4, F1: v1, F2: 2}
                var v134 int64 = (-(v5))
                _ = v133
                _ = v134
            }
            v135 := GenS2{F0: v5}
            _ = v135
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v136 := GenS7{F0: -149}
        _ = v136
    } else {
        var v137 int = 9223372036854775807
        for fi138 := 0; fi138 < 3; fi138++ {
            v139 := GenS6{F0: (v4 * v4)}
            v26 = v3
            fmt.Println(v137, v4)
            _ = fi138
            _ = v139
        }
        fmt.Println(v3)
        _ = v137
    }
    v6 -= (v6 * -95)
    if v27 {
        v140 := GenS13{F0: (v4 - (v4 / (((v4 << ((255) & 63))) | 1))), F1: "a", F2: -1}
        _ = v140
    }
    switch 4294967295 {
    case -895:
        v26 -= v3
    default:
        var v141 bool = true
        _ = v141
    }
    if (v1 != (v1 >> ((42) & 63))) {
        v2 = append(v2, 1)
        var v142 int = -9223372036854775808
        _ = v142
    } else {
        v6 += v6
    }
    v143 := GenS10{F0: 0, F1: "foo bar", F2: 610}
    for fi144 := 0; fi144 < 15; fi144++ {
        v26 -= v26
        v3 += v26
        if (!v27) {
            fmt.Println(v5)
        }
        _ = fi144
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v143
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
