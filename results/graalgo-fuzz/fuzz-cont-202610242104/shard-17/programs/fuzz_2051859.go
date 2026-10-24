package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint8; F2 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int64; F1 bool; F2 int }
type GenS11 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS12 struct { F0 uint16; F1 uint; F2 int32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int; F1 int8; F2 int }
type GenS15 struct { F0 int64 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]string)
    var v3 int8 = 0
    v1 = ("a" + v1)
    var v4 float64 = float64(((1e10 * (3.14159 - float32(1.5))) - float32(float64((3.14159 / 3.14159)))))
    v5 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: 1}
    var v6 uint64 = 0
    var v7 uint8 = 41
    if (1 > -1) {
        var v8 uint8 = 189
        _ = v8
    }
    if (!false) {
        switch ((v1 + v1) + (v1 + v1)) {
        case "foo bar":
            v6 = 18446744073709551615
        case "":
            fmt.Println(v7)
        case "a":
            for fi9 := 0; fi9 < 1; fi9++ {
                switch (fi9 % (((^(fi9))) | 1)) {
                case 56:
                    v1 = (v1 + ("" + ("a" + "hello")))
                case 32768:
                    v10 := GenS2{F0: 32767}
                    _ = v10
                case 2147483647:
                    switch (v1 + (((v1 + ("foo bar" + v1)) + v1) + "\n\t")) {
                    case "a":
                        v2[161] = v1
                    case "":
                        var v11 int64 = -802
                        _ = v11
                    default:
                        v1 = v1
                    }
                }
                _ = fi9
            }
        default:
            var v12 uint = 332
            _ = v12
        }
        fmt.Println(v4, v4)
        if (255 < 911) {
            for fi13 := 0; fi13 < 5; fi13++ {
                delete(v2, 24)
                for fi14 := 0; fi14 < 12; fi14++ {
                    if (0 < 0) {
                        var v15 uint = 255
                        switch fi14 {
                        case 127:
                            v16 := v2[670]
                            _ = v16
                        case 1:
                            v4 = (1.5 / -2.25)
                        case 0:
                            var v17 int8 = -57
                            _ = v17
                        }
                        v15 += v15
                        var v18 bool = (v15 >= v15)
                        _ = v15
                        _ = v18
                    }
                    _ = fi14
                }
                fmt.Println(v3)
                _ = fi13
            }
            if (484 != 2147483647) {
                v1 = v1
                switch 9223372036854775807 {
                case 1:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        fmt.Println(v4)
                        switch int(v3) {
                        case -2147483648:
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v6, v7, v4)
                            case "\n\t":
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    v21 := GenS1{F0: 912, F1: -917, F2: -468}
                                    v4 = v4
                                    var v22 uint = 353
                                    switch (("a" + "") + v1) {
                                    case "a":
                                        switch (fi20 * 32768) {
                                        case -484:
                                            var v23 uint8 = v7
                                            _ = v23
                                        default:
                                            v24 := GenS2{F0: 0}
                                            _ = v24
                                        }
                                    }
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            case "a":
                                switch fi19 {
                                case 4294967295:
                                    for fi25 := 0; fi25 < 7; fi25++ {
                                        fmt.Println(v4)
                                        var v26 uint64 = 2
                                        var v27 int32 = -421
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                case 9223372036854775807:
                                    v28 := []int{fi19, (fi19 & (fi19 ^ 245)), fi19, fi19}
                                    _ = v28
                                }
                            default:
                                if (1 >= 2147483647) {
                                    v29 := []int{(fi19 << ((18446744073709551615) & 63)), fi19, fi19}
                                    v30 := GenS1{F0: v6, F1: -9223372036854775808, F2: 0}
                                    v6 += (v6 % ((v6) | 1))
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS2{F0: -959}
                                    if (25 >= (v3 * -128)) {
                                        v32 := []int{}
                                        v33 := GenS3{F0: (v3 % ((v3) | 1)), F1: 255, F2: 9223372036854775807}
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v31
                                }
                            }
                        case 4294967295:
                            delete(v2, 101)
                        case 1:
                            if ((0.0 + 1e10) > 1e10) {
                                v34 := GenS1{F0: (^(v6)), F1: -1, F2: 42}
                                _ = v34
                            } else {
                                v4 -= v4
                                v7 -= (v7 % (((^(v7))) | 1))
                                if true {
                                    if ((3.14159 != (v4 / -2.25)) && true) {
                                        v1 = "\n\t"
                                        v35 := GenS3{F0: 1, F1: uint(v6), F2: -9223372036854775808}
                                        _ = v35
                                    }
                                } else {
                                    var v36 uint32 = 0
                                    var v37 uint64 = v6
                                    _ = v36
                                    _ = v37
                                }
                            }
                        }
                        v38 := GenS4{F0: 807}
                        _ = fi19
                        _ = v38
                    }
                default:
                    fmt.Println(v4)
                }
            }
        } else {
            switch 128 {
            case 127:
                var v39 float32 = float32((float64(1.5) + (v4 + (v4 - (float64((1e10 - (((1e10 * float32(3.14159)) + 1e10) / float32(v4)))) + 0.0)))))
                _ = v39
            case 42:
                if (!(-82 >= 0)) {
                    v40 := v2[42]
                    _ = v40
                }
            case 4294967295:
                v1 = ("" + "\n\t")
            }
        }
    }
    if true {
        v3 = 26
    }
    fmt.Println(v7)
    for fi41 := 0; fi41 < 18; fi41++ {
        var v42 int32 = -1
        var v43 bool = (false && true)
        v44 := GenS2{F0: 1000}
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    var v45 bool = (false || ((v7 & (v7 | (v7 * 2))) != 1))
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 int8 = v3
        var v48 int16 = -32768
        switch ((v1 + v1) + v1) {
        case "\n\t":
            v49 := []int{4294967295, 128, -1, 42}
            _ = v49
        case "foo bar":
            v50 := []int{(fi46 | (fi46 << ((255) & 63))), 42, 128}
            _ = v50
        case "hello":
            if (!((v6 + uint64(v7)) >= v6)) {
                v51 := GenS5{F0: 255, F1: v7, F2: fi46}
                for fi52 := 0; fi52 < 1; fi52++ {
                    if v45 {
                        v1 = v1
                        v3 -= v3
                    }
                    if v45 {
                        v47 -= (v3 + -1)
                        delete(v2, 440)
                        fmt.Println(v7, v45, v47)
                    }
                    switch 42 {
                    case 32767:
                        v47 = (v3 >> ((1) & 63))
                    case 339:
                        var v53 string = v1
                        _ = v53
                    default:
                        var v54 string = "\n\t"
                        _ = v54
                    }
                    _ = fi52
                }
                _ = v51
            }
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    for fi55 := 0; fi55 < 5; fi55++ {
        switch (fi55 % (((-(fi55))) | 1)) {
        case 0:
            v56 := GenS4{F0: 0}
            _ = v56
        case 32768:
            if ((true || false) || (1 == -1)) {
                var v57 bool = (32767 <= 0)
                _ = v57
            } else {
                switch (v1 + v1) {
                case "":
                    if v45 {
                        switch 67 {
                        case 2147483647:
                            v58 := GenS3{F0: v3, F1: 18446744073709551615, F2: -1}
                            _ = v58
                        case 844:
                            v59 := GenS6{F0: v6}
                            _ = v59
                        case 128:
                            v60 := v2[2147483648]
                            _ = v60
                        }
                        var v61 uint32 = 2147483647
                        _ = v61
                    }
                case "\n\t":
                    delete(v2, 2)
                }
                var v62 float64 = v4
                var v63 int8 = (-(v3))
                _ = v62
                _ = v63
            }
        default:
            switch 42 {
            case 42:
                switch fi55 {
                case 127:
                    v7 -= v7
                }
            case 29:
                var v64 int32 = -1
                _ = v64
            case 1:
                v3 = 100
            }
        }
        _ = fi55
    }
    if v45 {
        var v65 string = "a"
        _ = v65
    }
    switch ("" + "") {
    case "foo bar":
        switch v1 {
        case "foo bar":
            v6 += (^(v6))
        case "\n\t":
            var v66 uint = 248
            _ = v66
        case "":
            for fi67 := 0; fi67 < 8; fi67++ {
                v7 += (^(v7))
                v7 = v7
                var v68 uint16 = 255
                _ = fi67
                _ = v68
            }
        default:
            if (-633 <= -1) {
                v3 = (v3 / ((v3) | 1))
                if (3.14159 < 3.14159) {
                    for fi69 := 0; fi69 < 17; fi69++ {
                        v4 = 1e10
                        _ = fi69
                    }
                }
            } else {
                if (v45 && true) {
                    var v70 uint16 = 42
                    v71 := []int{-639, 42, 128}
                    _ = v70
                    _ = v71
                } else {
                    v6 = uint64(v7)
                }
                if v45 {
                    fmt.Println(v7)
                    var v72 uint = 0
                    _ = v72
                } else {
                    v73 := GenS4{F0: 65535}
                    v74 := GenS1{F0: uint64(v7), F1: 9223372036854775807, F2: 42}
                    _ = v73
                    _ = v74
                }
                var v75 string = (v1 + v1)
                _ = v75
            }
        }
    default:
        v7 -= (v7 * (v7 + v7))
    }
    for fi76 := 0; fi76 < 2; fi76++ {
        if ((fi76 / (((fi76 / ((fi76) | 1))) | 1)) >= (fi76 & fi76)) {
            for fi77 := 0; fi77 < 11; fi77++ {
                v1 = v1
                v78 := GenS7{F0: 413}
                v79 := GenS3{F0: v3, F1: 65535, F2: -9223372036854775808}
                _ = fi77
                _ = v78
                _ = v79
            }
        } else {
            switch (-(fi76)) {
            case 1:
                for fi80 := 0; fi80 < 3; fi80++ {
                    v1 = ("" + (((v1 + v1) + (v1 + v1)) + v1))
                    v7 -= (v7 | v7)
                    v81 := []int{}
                    _ = fi80
                    _ = v81
                }
            }
            if true {
                if (2147483647 != -718) {
                    for fi82 := 0; fi82 < 20; fi82++ {
                        v83 := GenS8{F0: 4294967295}
                        _ = fi82
                        _ = v83
                    }
                } else {
                    if true {
                        v7 += 178
                        for fi84 := 0; fi84 < 19; fi84++ {
                            switch (^(fi84)) {
                            case -9223372036854775808:
                                var v85 uint8 = v7
                                _ = v85
                            case 2147483647:
                                v86 := []int{46, (fi84 - fi76)}
                                _ = v86
                            case 1:
                                switch ((v1 + "hello") + ("foo bar" + v1)) {
                                case "foo bar":
                                    var v87 uint16 = 1
                                    _ = v87
                                case "":
                                    v7 = (^(v7))
                                case "\n\t":
                                    if v45 {
                                        var v88 int32 = 305
                                        _ = v88
                                    } else {
                                        v45 = (!v45)
                                        if ((65535 <= 65535) || true) {
                                            if v45 {
                                                switch (v1 + (v1 + (v1 + v1))) {
                                                case "\n\t":
                                                    v6 = v6
                                                case "a":
                                                    if true {
                                                        v89 := GenS4{F0: 65535}
                                                        _ = v89
                                                    } else {
                                                        v90 := v2[2147483648]
                                                        v91 := GenS9{F0: 1, F1: int16(fi76)}
                                                        if v45 {
                                                            v7 -= 219
                                                        }
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                case "hello":
                                                    var v92 bool = v45
                                                    _ = v92
                                                default:
                                                    v45 = v45
                                                }
                                            } else {
                                                switch fi84 {
                                                case -404:
                                                    fmt.Println(v45)
                                                case -854:
                                                    v45 = (1 < (v6 << ((0) & 63)))
                                                case 2147483647:
                                                    v93 := GenS10{F0: 42, F1: v45, F2: fi76}
                                                    _ = v93
                                                }
                                                v7 = v7
                                                if v45 {
                                                    if false {
                                                        v4 = float64(0.1)
                                                    } else {
                                                        if (-1 != v3) {
                                                            var v94 int64 = -986
                                                            v94 -= v94
                                                            v3 -= v3
                                                            _ = v94
                                                        }
                                                        v95 := GenS11{F0: v7, F1: v7, F2: 0}
                                                        _ = v95
                                                    }
                                                }
                                                v4 = 0.1
                                            }
                                            v96 := GenS4{F0: 18446744073709551615}
                                            _ = v96
                                        }
                                        var v97 uint64 = (v6 - (v6 % (((v6 / ((216) | 1))) | 1)))
                                        _ = v97
                                    }
                                }
                            default:
                                var v98 uint8 = (^(v7))
                                _ = v98
                            }
                            v45 = true
                            _ = fi84
                        }
                    } else {
                        if (-1 != 1) {
                            var v99 int16 = 32767
                            _ = v99
                        }
                    }
                }
                var v100 uint16 = 32767
                for fi101 := 0; fi101 < 19; fi101++ {
                    switch v1 {
                    case "\n\t":
                        v102 := GenS12{F0: (^(v100)), F1: 2, F2: 42}
                        _ = v102
                    default:
                        v103 := GenS13{F0: 816}
                        _ = v103
                    }
                    v104 := GenS14{F0: fi101, F1: 7, F2: (fi76 + (fi101 / ((fi101) | 1)))}
                    _ = fi101
                    _ = v104
                }
                _ = v100
            } else {
                v4 = float64(float32(v4))
            }
            var v105 int64 = 42
            _ = v105
        }
        v106 := GenS1{F0: v6, F1: 205, F2: 0}
        v107 := GenS15{F0: -1}
        _ = fi76
        _ = v106
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
