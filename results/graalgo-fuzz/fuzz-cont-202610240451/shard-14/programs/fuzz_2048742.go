package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint }
type GenS6 struct { F0 int8; F1 int8; F2 uint64 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 int16; F2 int8 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 string; F1 uint16; F2 int8 }
type GenS16 struct { F0 string; F1 uint8; F2 uint32 }
type GenS17 struct { F0 int16; F1 uint32; F2 int32 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 bool; F1 uint8 }
type GenS22 struct { F0 int64; F1 uint32 }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 int8; F1 bool; F2 int16 }
type GenS25 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS26 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = ((false || (127 == 740)) && ((0.1 / float64(-2.25)) < (1.5 * 0.1)))
    var v3 string = "\n\t"
    v4 := make(map[string]uint32)
    v5 := []uint32{2147483647, 65535, 42}
    switch v3 {
    case "foo bar":
        v6 := GenS1{F0: 519, F1: 1}
        _ = v6
    default:
        for fi7 := 0; fi7 < 8; fi7++ {
            if (2147483647 <= 2147483647) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                    var v9 uint64 = 4294967295
                    _ = fi8
                    _ = v9
                }
                var v10 bool = (991 != 2147483647)
                v2 = (402 <= 103)
                _ = v10
            } else {
                for fi11 := 0; fi11 < 13; fi11++ {
                    v3 = v3
                    v3 = v3
                    _ = len(v4)
                    _ = fi11
                }
            }
            _ = fi7
        }
    }
    if (0 != 2147483647) {
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(v3, v3)
            fmt.Println(v1)
            for fi13 := 0; fi13 < 15; fi13++ {
                if (!(!v2)) {
                    switch "a" {
                    case "\n\t":
                        if len(v5) > 0 {
                            v5[0] = 65535
                        }
                    case "":
                        var v14 int64 = 42
                        _ = v14
                    case "hello":
                        fmt.Println(v3, v1)
                    }
                    if true {
                        if false {
                            v3 = ("" + ("foo bar" + v3))
                        } else {
                            var v15 uint64 = 4294967295
                            v5 = append(v5, 0)
                            v16 := GenS1{F0: v15, F1: 37}
                            var v17 int64 = 0
                            var v18 int = (fi12 - (fi12 + 9223372036854775807))
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        if ((!(!(!(127 != 100)))) && v2) {
                            v19 := v4["hello"]
                            _ = v19
                        } else {
                            fmt.Println(v2)
                        }
                    } else {
                        var v20 string = "a"
                        _ = v20
                    }
                    fmt.Println(v2)
                }
                v1 += (v1 * v1)
                _ = fi13
            }
            _ = fi12
        }
        var v21 int16 = -1
        _ = v21
    }
    v3 = "hello"
    var v22 int8 = 100
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 6; fi24++ {
            switch (((v3 + "\n\t") + "hello") + (v3 + (v3 + (("a" + "\n\t") + ((v3 + (("foo bar" + "foo bar") + ("a" + v3))) + "hello"))))) {
            case "":
                fmt.Println(v2, v3, v3)
            case "hello":
                var v25 uint8 = 255
                _ = v25
            case "\n\t":
                fmt.Println(fi23)
            default:
                fmt.Println(fi24, v2)
            }
            _ = fi24
        }
        var v26 int8 = v22
        _ = fi23
        _ = v26
    }
    var v27 int64 = 9223372036854775807
    var v28 int8 = v22
    v29 := GenS1{F0: 2, F1: -1}
    v1 = (v1 - (v1 + v1))
    v27 = (v27 >> ((65535) & 63))
    var v30 uint = 42
    var v31 string = "foo bar"
    if true {
        if (!(v22 > -11)) {
            switch 913 {
            case 127:
                switch -2147483648 {
                case 9223372036854775807:
                    _ = len(v5)
                default:
                    v1 = v30
                }
            default:
                if true {
                    var v32 uint = (v30 & v30)
                    switch -1 {
                    case 2147483647:
                        v22 -= (v28 ^ (-(v28)))
                    case -131:
                        v33 := GenS2{F0: 1, F1: "a", F2: 32767}
                        _ = v33
                    }
                    _ = v32
                } else {
                    var v34 uint8 = 12
                    v1 = (v30 & (^(v1)))
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 16; fi35++ {
                switch fi35 {
                case 4294967295:
                    v36 := GenS3{F0: v2, F1: 128}
                    _ = v36
                case 42:
                    fmt.Println(v31, v30, v3)
                case 32768:
                    v1 = 599
                default:
                    fmt.Println(v31, fi35, v2)
                }
                for fi37 := 0; fi37 < 18; fi37++ {
                    switch "hello" {
                    case "\n\t":
                        var v38 float32 = 0.0
                        if ((32767 == 255) || ((v30 + 413) <= (v30 / (((v1 * (v1 * (v30 | 234)))) | 1)))) {
                            v31 = v3
                            if v2 {
                                var v39 int8 = (^(v22))
                                v40 := GenS2{F0: 255, F1: "\n\t", F2: 372}
                                _ = v39
                                _ = v40
                            } else {
                                if false {
                                    switch (fi35 * -182) {
                                    case -9223372036854775808:
                                        var v41 float32 = float32(1.5)
                                        _ = v41
                                    case 262:
                                        v38 += (v38 * v38)
                                    }
                                    fmt.Println(fi37, v27)
                                    _ = len(v5)
                                    v30 += v30
                                } else {
                                    v42 := GenS4{F0: 0, F1: true}
                                    var v43 string = v3
                                    _ = v42
                                    _ = v43
                                }
                                var v44 int8 = -128
                                var v45 int64 = (v27 ^ (v27 << ((255) & 63)))
                                v1 -= v30
                                if v2 {
                                    v46 := GenS5{F0: 1, F1: 4294967295, F2: 548}
                                    if false {
                                        var v47 uint64 = 18446744073709551615
                                        switch fi35 {
                                        case -842:
                                            if v2 {
                                                v2 = true
                                                var v48 uint64 = (^(v47))
                                                if len(v5) > 0 {
                                                    v5[0] = 4294967295
                                                }
                                                _ = v48
                                            } else {
                                                _ = len(v5)
                                            }
                                        }
                                        switch v31 {
                                        case "foo bar":
                                            v3 = v3
                                        }
                                        _ = v47
                                    }
                                    v49 := GenS6{F0: (v44 * (v22 & (v44 ^ v22))), F1: v44, F2: 42}
                                    _ = v46
                                    _ = v49
                                } else {
                                    switch (v3 + v31) {
                                    case "a":
                                        v1 = v1
                                    case "hello":
                                        var v50 int16 = 0
                                        _ = v50
                                    }
                                    v4[v3] = uint32(v30)
                                }
                                _ = v44
                                _ = v45
                            }
                            var v51 int = fi35
                            var v52 float32 = (v38 / (0.0 * ((v38 * v38) * ((v38 + v38) * 0.1))))
                            if v2 {
                                v53 := GenS2{F0: 128, F1: v31, F2: 531}
                                v28 -= (v22 + -24)
                                v52 -= ((float32(float64(v52)) + 3.14159) - 0.1)
                                _ = v53
                            } else {
                                v22 += v28
                                v51 -= fi35
                                if (v52 < (0.1 / (v52 / 1.0))) {
                                    fmt.Println(v22, v28, fi35)
                                    v27 = (v27 ^ v27)
                                    fmt.Println(v38)
                                } else {
                                    var v54 int = fi35
                                    v52 = (1e10 * 0.1)
                                    switch 283 {
                                    case 4294967295:
                                        v5 = append(v5, 704)
                                    case 128:
                                        v55 := GenS2{F0: 128, F1: v31, F2: int16(v28)}
                                        switch (fi37 * fi35) {
                                        case 4294967295:
                                            var v56 int64 = 1
                                            var v57 uint8 = 0
                                            v5 = append(v5, 65535)
                                            _ = v56
                                            _ = v57
                                        case 128:
                                            v58 := v4[""]
                                            _ = v58
                                        case 127:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        var v59 float64 = 0.0
                                        _ = v55
                                        _ = v59
                                    case 32767:
                                        fmt.Println(fi37)
                                    default:
                                        v60 := GenS3{F0: false, F1: 96}
                                        _ = v60
                                    }
                                    _ = v54
                                }
                            }
                            _ = v51
                            _ = v52
                        }
                        _ = v38
                    case "hello":
                        v1 -= (v1 >> (((v1 * v1)) & 63))
                    default:
                        switch fi35 {
                        case -2147483648:
                            switch (fi35 | (fi35 + (fi35 * 32767))) {
                            case -79:
                                fmt.Println(fi35)
                            case 617:
                                var v61 int16 = -472
                                _ = v61
                            }
                        case 9223372036854775807:
                            if len(v5) > 0 {
                                v5[0] = 2
                            }
                        case 42:
                            switch "" {
                            case "a":
                                var v62 float32 = ((1e10 + 1.5) / 1.0)
                                _ = v62
                            case "":
                                v63 := GenS5{F0: 155, F1: 0, F2: v30}
                                _ = v63
                            case "hello":
                                v4[("a" + "hello")] = 82
                            }
                        }
                    }
                    fmt.Println(fi37, v22, v1)
                    _ = fi37
                }
                for fi64 := 0; fi64 < 2; fi64++ {
                    var v65 uint32 = 42
                    v27 = 0
                    if v2 {
                        var v66 int = (fi64 << ((v1) & 63))
                        switch 1 {
                        case -1:
                            v67 := GenS5{F0: 42, F1: 1, F2: 2}
                            _ = v67
                        case -950:
                            var v68 uint32 = (v65 << (((v1 | 107)) & 63))
                            _ = v68
                        case 823:
                            if ("a" != "hello") {
                                v69 := GenS7{F0: v28, F1: (v27 / (((^(v27))) | 1))}
                                _ = v69
                            }
                        default:
                            if (!false) {
                                if (v2 || v2) {
                                    switch v3 {
                                    case "\n\t":
                                        var v70 float32 = float32(3.14159)
                                        _ = v70
                                    case "hello":
                                        if v2 {
                                            switch (v66 & 32767) {
                                            case -804:
                                                var v71 int64 = int64(v65)
                                                _ = v71
                                            case 127:
                                                var v72 int64 = (^(v27))
                                                switch (fi35 << ((47) & 63)) {
                                                case 378:
                                                    if (true && v2) {
                                                        if true {
                                                            delete(v4, "hello")
                                                            v30 = (^(v30))
                                                            if ((-(v27)) == int64(v28)) {
                                                                var v73 float32 = 3.14159
                                                                if (!v2) {
                                                                    fmt.Println(fi35, v3, v2)
                                                                    switch (fi64 >> ((v30) & 63)) {
                                                                    case -1:
                                                                        switch v31 {
                                                                        case "a":
                                                                            var v74 int64 = (v27 * v27)
                                                                            _ = v74
                                                                        case "\n\t":
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        v75 := GenS1{F0: 18446744073709551615, F1: -1}
                                                                        _ = v75
                                                                    }
                                                                }
                                                                if (v2 && v2) {
                                                                    v76 := GenS4{F0: 4294967295, F1: v2}
                                                                    if v2 {
                                                                        v28 -= v28
                                                                    } else {
                                                                        var v77 uint = 18446744073709551615
                                                                        v22 = v28
                                                                        if (!(false || (((2 < 249) && ((0 > 0) && v2)) || false))) {
                                                                            var v78 int8 = (-(v22))
                                                                            _ = v78
                                                                        } else {
                                                                            v79 := GenS3{F0: (fi64 <= (v66 - fi64)), F1: 127}
                                                                            v80 := GenS8{F0: (v72 ^ (v72 - 9223372036854775807))}
                                                                            _ = len(v4)
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        _ = len(v5)
                                                                        _ = v77
                                                                    }
                                                                    v65 = (^(v65))
                                                                    _ = v76
                                                                }
                                                                _ = v73
                                                            } else {
                                                                if v2 {
                                                                    switch ("" + v31) {
                                                                    case "a":
                                                                        fmt.Println(v28, v31)
                                                                    case "foo bar":
                                                                        var v81 uint8 = 182
                                                                        _ = v81
                                                                    default:
                                                                        switch v31 {
                                                                        case "hello":
                                                                            var v82 float32 = 0.1
                                                                            _ = v82
                                                                        case "":
                                                                            switch v3 {
                                                                            case "hello":
                                                                                v28 = v22
                                                                                v4[((v3 + "\n\t") + v31)] = 647
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (v65 >> (((v30 / (((^(v30))) | 1))) & 63))
                                                                                }
                                                                                if v2 {
                                                                                    switch (^(fi64)) {
                                                                                    case 42:
                                                                                        switch v3 {
                                                                                        case "":
                                                                                            switch (fi64 | v66) {
                                                                                            case -1:
                                                                                                var v83 int64 = (v27 * v72)
                                                                                                _ = v83
                                                                                            case -2147483648:
                                                                                                v84 := v4[(v31 + ("\n\t" + "hello"))]
                                                                                                _ = v84
                                                                                            case -204:
                                                                                                if len(v5) >= 2 {
                                                                                                    v5 = v5[1:len(v5)]
                                                                                                }
                                                                                            default:
                                                                                                v85 := GenS9{F0: 771, F1: -1, F2: (v28 & v22)}
                                                                                                _ = v85
                                                                                            }
                                                                                        case "a":
                                                                                            v1 = v30
                                                                                        case "hello":
                                                                                            fmt.Println(v3, v72)
                                                                                        default:
                                                                                            v86 := GenS10{F0: v65, F1: (255 == 937), F2: v30}
                                                                                            _ = v86
                                                                                        }
                                                                                    case 876:
                                                                                        switch v31 {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v30, v28)
                                                                                        }
                                                                                    default:
                                                                                        v72 += (-(v72))
                                                                                    }
                                                                                    var v87 uint = v30
                                                                                    v2 = ((fi64 & 1) == (fi64 >> ((4294967295) & 63)))
                                                                                    _ = v87
                                                                                } else {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = uint32(v30)
                                                                                    }
                                                                                    v30 = (v1 - 223)
                                                                                    v88 := GenS11{F0: 1, F1: 183}
                                                                                    _ = v88
                                                                                }
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (^(v65))
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            var v89 uint64 = 991
                                                                            _ = v89
                                                                        default:
                                                                            if false {
                                                                                var v90 uint16 = 2
                                                                                if true {
                                                                                    v4[(v31 + v31)] = v65
                                                                                    switch fi35 {
                                                                                    case 906:
                                                                                        var v91 string = ("a" + v31)
                                                                                        _ = v91
                                                                                    case -9223372036854775808:
                                                                                        var v92 uint16 = 255
                                                                                        _ = v92
                                                                                    case -2147483648:
                                                                                        fmt.Println(v28, v30)
                                                                                    default:
                                                                                        if v2 {
                                                                                            var v93 uint8 = 203
                                                                                            var v94 uint32 = (v65 >> (((^(v1))) & 63))
                                                                                            _ = v93
                                                                                            _ = v94
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v65 = (v65 + v65)
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        var v95 uint16 = v90
                                                                                        _ = v95
                                                                                    case "a":
                                                                                        if (42 < v22) {
                                                                                            if v2 {
                                                                                                v3 = "a"
                                                                                                switch v31 {
                                                                                                case "a":
                                                                                                    if false {
                                                                                                        v27 = v27
                                                                                                        var v96 uint16 = v90
                                                                                                        v97 := GenS11{F0: -65, F1: 128}
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                    } else {
                                                                                                        if (false && false) {
                                                                                                            fmt.Println(v28, v72, v27)
                                                                                                            fmt.Println(v1)
                                                                                                            fmt.Println(v1)
                                                                                                        } else {
                                                                                                            fmt.Println(v66, v72)
                                                                                                            v98 := GenS8{F0: v27}
                                                                                                            v65 = v65
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        v3 = v3
                                                                                                        var v99 uint8 = 42
                                                                                                        v4[("" + v31)] = v65
                                                                                                        switch (fi35 >> (((v1 - (v30 + 42))) & 63)) {
                                                                                                        case 0:
                                                                                                            v4[v31] = v65
                                                                                                        case 32767:
                                                                                                            switch (^(v66)) {
                                                                                                            case 32768:
                                                                                                                var v100 float32 = -2.25
                                                                                                                _ = v100
                                                                                                            case 1:
                                                                                                                fmt.Println(v30, fi64)
                                                                                                                v101 := v4["foo bar"]
                                                                                                                v22 = v22
                                                                                                                _ = v101
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v65)
                                                                                                        }
                                                                                                        _ = v99
                                                                                                    }
                                                                                                    var v102 uint = v1
                                                                                                    _ = v102
                                                                                                case "\n\t":
                                                                                                    v30 = 0
                                                                                                case "":
                                                                                                    v65 = (v65 & v65)
                                                                                                }
                                                                                                v103 := GenS11{F0: int16(v1), F1: 83}
                                                                                                switch 4294967295 {
                                                                                                case -1:
                                                                                                    switch v66 {
                                                                                                    case 128:
                                                                                                        v22 += (v28 | v28)
                                                                                                    default:
                                                                                                        if false {
                                                                                                            switch (v66 << ((v30) & 63)) {
                                                                                                            case -29:
                                                                                                                fmt.Println(v65)
                                                                                                            case 2147483647:
                                                                                                                if ((^(v65)) > uint32(v30)) {
                                                                                                                    switch fi64 {
                                                                                                                    case 4294967295:
                                                                                                                        v31 = v3
                                                                                                                    case 2147483647:
                                                                                                                        var v104 int8 = 0
                                                                                                                        _ = v104
                                                                                                                    case 0:
                                                                                                                        fmt.Println(v65, v66, v31)
                                                                                                                    default:
                                                                                                                        switch ((("\n\t" + "a") + "foo bar") + ((("foo bar" + "\n\t") + (("\n\t" + (v3 + v3)) + v31)) + v3)) {
                                                                                                                        case "foo bar":
                                                                                                                            var v105 uint32 = (v65 + (v65 ^ v65))
                                                                                                                            _ = v105
                                                                                                                        default:
                                                                                                                            _ = len(v5)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v106 int32 = int32(v1)
                                                                                                                    v1 += uint(v65)
                                                                                                                    _ = v106
                                                                                                                }
                                                                                                            default:
                                                                                                                var v107 float32 = 0.1
                                                                                                                fmt.Println(v72)
                                                                                                                _ = v107
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case 178:
                                                                                                    _ = len(v4)
                                                                                                default:
                                                                                                    v5 = append(v5, (v65 * 1))
                                                                                                }
                                                                                                _ = v103
                                                                                            } else {
                                                                                                if (0.0 > (((3.14159 / 1.5) + (0.0 * (float64(0.0) + ((3.14159 + 1e10) + ((float64(3.14159) - (1.5 / -2.25)) * 3.14159))))) / 1.0)) {
                                                                                                    v72 = v27
                                                                                                    var v108 int16 = -807
                                                                                                    v22 = 1
                                                                                                    _ = v108
                                                                                                }
                                                                                                v22 = v22
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(fi35)
                                                                                        }
                                                                                    case "hello":
                                                                                        if (!(-2.25 <= float64((0.1 * 0.1)))) {
                                                                                            var v109 uint32 = v65
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            _ = v109
                                                                                        } else {
                                                                                            var v110 uint8 = 255
                                                                                            _ = v110
                                                                                        }
                                                                                    default:
                                                                                        if v2 {
                                                                                            var v111 uint64 = 366
                                                                                            _ = v111
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v90
                                                                            } else {
                                                                                if v2 {
                                                                                    var v112 uint = 947
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        var v113 float32 = (1e10 * (0.0 / ((((3.14159 + (1e10 * (0.1 * 1e10))) + float32(1e10)) + ((1.5 / 1e10) / -2.25)) - (0.0 / 0.1))))
                                                                                        _ = v113
                                                                                    case "\n\t":
                                                                                        var v114 uint16 = 0
                                                                                        _ = v114
                                                                                    case "hello":
                                                                                        var v115 uint64 = 4294967295
                                                                                        _ = v115
                                                                                    }
                                                                                    var v116 int16 = -568
                                                                                    _ = v112
                                                                                    _ = v116
                                                                                }
                                                                                v3 = v31
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v117 := GenS2{F0: 128, F1: v31, F2: 0}
                                                                                    _ = v117
                                                                                case "\n\t":
                                                                                    v22 += v28
                                                                                }
                                                                                if false {
                                                                                    if v2 {
                                                                                        v1 -= v30
                                                                                        var v118 uint64 = 42
                                                                                        _ = v118
                                                                                    } else {
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        switch v66 {
                                                                                        case 902:
                                                                                            var v119 float64 = 1.5
                                                                                            _ = v119
                                                                                        }
                                                                                        var v120 int16 = -1
                                                                                        _ = v120
                                                                                    }
                                                                                }
                                                                                v121 := GenS11{F0: int16(v1), F1: 2}
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                    }
                                                                    switch fi64 {
                                                                    case 406:
                                                                        v122 := v4[v31]
                                                                        _ = v122
                                                                    }
                                                                    if (false || v2) {
                                                                        if v2 {
                                                                            switch (v31 + v31) {
                                                                            case "a":
                                                                                if (!(0.0 == (3.14159 * 1.5))) {
                                                                                    v5 = append(v5, v65)
                                                                                    var v123 int = 32767
                                                                                    _ = v123
                                                                                } else {
                                                                                    v1 = v30
                                                                                }
                                                                            }
                                                                            var v124 string = "a"
                                                                            v2 = false
                                                                            v125 := v4[v124]
                                                                            _ = v124
                                                                            _ = v125
                                                                        } else {
                                                                            if ((1e10 + (3.14159 * 1.5)) < (1.5 / (3.14159 / 1.0))) {
                                                                                v126 := GenS12{F0: v30}
                                                                                switch 32767 {
                                                                                case 1:
                                                                                    switch v31 {
                                                                                    case "":
                                                                                        if (true && v2) {
                                                                                            var v127 uint = 0
                                                                                            fmt.Println(fi64, v30)
                                                                                            _ = v127
                                                                                        }
                                                                                    case "hello":
                                                                                        var v128 int = fi35
                                                                                        _ = v128
                                                                                    case "\n\t":
                                                                                        fmt.Println(v27, v27, fi35)
                                                                                    }
                                                                                case 42:
                                                                                    switch (v31 + (v3 + v3)) {
                                                                                    case "a":
                                                                                        var v129 int = (-(fi64))
                                                                                        _ = v129
                                                                                    case "hello":
                                                                                        var v130 int64 = v27
                                                                                        v30 += 564
                                                                                        delete(v4, v3)
                                                                                        fmt.Println(v22, v2, v28)
                                                                                        _ = v130
                                                                                    case "":
                                                                                        v31 = ("" + v31)
                                                                                    }
                                                                                    var v131 bool = v2
                                                                                    if true {
                                                                                        v22 = (v22 / ((v22) | 1))
                                                                                        v132 := v4[((v31 + "foo bar") + (v31 + ((v31 + v3) + v31)))]
                                                                                        v72 += (v72 << (((v30 + v30)) & 63))
                                                                                        var v133 bool = (true || (v65 < (v65 << (((v1 * (v30 - 4294967295))) & 63))))
                                                                                        v131 = v131
                                                                                        _ = v132
                                                                                        _ = v133
                                                                                    }
                                                                                    var v134 string = v31
                                                                                    _ = v131
                                                                                    _ = v134
                                                                                default:
                                                                                    v65 = 785
                                                                                }
                                                                                var v135 uint32 = v65
                                                                                var v136 string = v3
                                                                                _ = v126
                                                                                _ = v135
                                                                                _ = v136
                                                                            }
                                                                            var v137 uint32 = (v65 << (((^(v30))) & 63))
                                                                            v5 = append(v5, v137)
                                                                            _ = v137
                                                                        }
                                                                        if (505 >= 964) {
                                                                            v138 := GenS11{F0: 662, F1: 2}
                                                                            fmt.Println(v28, v28)
                                                                            _ = v138
                                                                        }
                                                                    } else {
                                                                        var v139 bool = (0 > 388)
                                                                        switch v66 {
                                                                        case 128:
                                                                            switch v66 {
                                                                            case 128:
                                                                                v140 := GenS4{F0: 1, F1: v139}
                                                                                _ = v140
                                                                            case 4294967295:
                                                                                var v141 bool = false
                                                                                _ = v141
                                                                            case 42:
                                                                                v65 = 907
                                                                            default:
                                                                                switch v31 {
                                                                                case "hello":
                                                                                    v142 := GenS3{F0: (65535 <= v65), F1: 0}
                                                                                    _ = v142
                                                                                case "a":
                                                                                    if (!v2) {
                                                                                        var v143 int8 = (v28 ^ v22)
                                                                                        v144 := GenS6{F0: v22, F1: v22, F2: 508}
                                                                                        v5 = append(v5, v65)
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    } else {
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = v65
                                                                                        }
                                                                                        switch v31 {
                                                                                        case "":
                                                                                            fmt.Println(v31, v27, v2)
                                                                                        default:
                                                                                            v145 := GenS12{F0: (v1 & uint(v65))}
                                                                                            _ = v145
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v139
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            switch 1 {
                                                            case 4294967295:
                                                                if (0 != 42) {
                                                                    v30 += v1
                                                                    fmt.Println(v30)
                                                                    v28 += v28
                                                                }
                                                            case 394:
                                                                v146 := GenS13{F0: v66}
                                                                _ = v146
                                                            case 32768:
                                                                var v147 float32 = 1e10
                                                                _ = v147
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                            if ((v30 | 255) < v30) {
                                                                v22 -= v22
                                                                if len(v5) > 0 {
                                                                    v5[0] = 4294967295
                                                                }
                                                                var v148 int32 = 1
                                                                if (v1 > v1) {
                                                                    if (v2 && v2) {
                                                                        var v149 bool = (1e10 <= 1.5)
                                                                        _ = v149
                                                                    } else {
                                                                        var v150 uint32 = v65
                                                                        switch ((v31 + ((v3 + ("foo bar" + v31)) + (v3 + ("foo bar" + v3)))) + ("a" + ("" + ("a" + v31)))) {
                                                                        case "foo bar":
                                                                            v30 += (^(v1))
                                                                        case "hello":
                                                                            switch 128 {
                                                                            case 32768:
                                                                                if true {
                                                                                    v148 += -2147483648
                                                                                    var v151 uint = (^(v1))
                                                                                    _ = v151
                                                                                } else {
                                                                                    var v152 int8 = (v22 | v28)
                                                                                    _ = v152
                                                                                }
                                                                            case 9223372036854775807:
                                                                                var v153 bool = (32767 < v66)
                                                                                _ = v153
                                                                            default:
                                                                                v154 := GenS14{F0: 42}
                                                                                _ = v154
                                                                            }
                                                                        }
                                                                        switch 276 {
                                                                        case 0:
                                                                            switch (v66 / ((-969) | 1)) {
                                                                            case 32768:
                                                                                var v155 int16 = 1
                                                                                _ = v155
                                                                            default:
                                                                                fmt.Println(v148, v3, fi64)
                                                                            }
                                                                        case 127:
                                                                            if ((-2.25 >= 3.14159) || (v2 && (!((((((0.1 * 3.14159) * (0.0 + 1.5)) * 1e10) - 0.0) / 0.1) != 1e10)))) {
                                                                                switch fi35 {
                                                                                case 4294967295:
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                default:
                                                                                    v156 := GenS8{F0: (v27 & v72)}
                                                                                    _ = v156
                                                                                }
                                                                                if v2 {
                                                                                    fmt.Println(v28, v72)
                                                                                } else {
                                                                                    if v2 {
                                                                                        v157 := GenS15{F0: (((v31 + "a") + "hello") + v3), F1: 581, F2: v28}
                                                                                        if v2 {
                                                                                            fmt.Println(v31, v27, v22)
                                                                                            if (v22 == v22) {
                                                                                                v150 = v150
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    v28 = (-(v22))
                                                                                                }
                                                                                                var v158 int16 = -1
                                                                                                _ = v158
                                                                                            } else {
                                                                                                v1 = 344
                                                                                            }
                                                                                            v159 := GenS6{F0: 125, F1: (v28 >> ((v30) & 63)), F2: 42}
                                                                                            delete(v4, v31)
                                                                                            _ = v159
                                                                                        }
                                                                                        var v160 float32 = -2.25
                                                                                        fmt.Println(v31, v160)
                                                                                        _ = v157
                                                                                        _ = v160
                                                                                    }
                                                                                    v66 += 128
                                                                                    v1 = (v30 - v1)
                                                                                }
                                                                                v161 := GenS16{F0: v31, F1: 205, F2: (v65 * 1000)}
                                                                                fmt.Println(v3, v148)
                                                                                _ = v161
                                                                            }
                                                                        case 879:
                                                                            v150 = (v65 << ((255) & 63))
                                                                        default:
                                                                            var v162 int32 = v148
                                                                            _ = v162
                                                                        }
                                                                        _ = v150
                                                                    }
                                                                    var v163 int64 = 42
                                                                    _ = v163
                                                                } else {
                                                                    var v164 uint16 = 0
                                                                    _ = v164
                                                                }
                                                                var v165 int = (fi64 * fi35)
                                                                var v166 float32 = (((0.1 / -2.25) / (0.0 + 1.5)) / ((0.1 / 1e10) - -2.25))
                                                                _ = v148
                                                                _ = v165
                                                                _ = v166
                                                            } else {
                                                                var v167 uint64 = 1
                                                                var v168 int32 = 1
                                                                _ = v167
                                                                _ = v168
                                                            }
                                                        }
                                                        v1 = (^(v30))
                                                        if false {
                                                            switch v31 {
                                                            case "":
                                                                if ((-9223372036854775808 >= 128) || v2) {
                                                                    var v169 int = fi35
                                                                    var v170 bool = (!true)
                                                                    v171 := GenS9{F0: (v72 & (-(v27))), F1: -549, F2: (v22 / ((int8(v65)) | 1))}
                                                                    _ = v169
                                                                    _ = v170
                                                                    _ = v171
                                                                }
                                                            }
                                                        } else {
                                                            if v2 {
                                                                v172 := GenS5{F0: 128, F1: 42, F2: 255}
                                                                v173 := GenS4{F0: 18446744073709551615, F1: v2}
                                                                _ = v172
                                                                _ = v173
                                                            }
                                                        }
                                                    } else {
                                                        var v174 int32 = 42
                                                        v72 = v27
                                                        _ = len(v4)
                                                        var v175 bool = v2
                                                        v3 = (((v31 + ("\n\t" + v31)) + "") + "")
                                                        _ = len(v4)
                                                        _ = v174
                                                        _ = v175
                                                    }
                                                case 32768:
                                                    var v176 uint32 = v65
                                                    _ = v176
                                                case -2147483648:
                                                    v28 = (v22 - v28)
                                                default:
                                                    if v2 {
                                                        var v177 int16 = 42
                                                        switch (^(v66)) {
                                                        case 42:
                                                            v178 := GenS17{F0: v177, F1: (^(v65)), F2: 1}
                                                            _ = v178
                                                        case -976:
                                                            fmt.Println(fi35, v27)
                                                        }
                                                        _ = v177
                                                    }
                                                }
                                                _ = v72
                                            }
                                            var v179 string = (v3 + v31)
                                            _ = v179
                                        }
                                    case "":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    default:
                                        if v2 {
                                            var v180 int = fi64
                                            _ = len(v5)
                                            var v181 int32 = -18
                                            delete(v4, ((("\n\t" + ((v3 + "hello") + ((v3 + v3) + v3))) + ((((v31 + v3) + (v3 + (v3 + "hello"))) + v3) + v3)) + ("hello" + v3)))
                                            var v182 float32 = -2.25
                                            _ = v180
                                            _ = v181
                                            _ = v182
                                        } else {
                                            switch v3 {
                                            case "a":
                                                var v183 bool = v2
                                                _ = v183
                                            case "":
                                                fmt.Println(v65, v3)
                                            case "hello":
                                                if (((0.1 + ((0.1 + float32(0.1)) / ((1.5 * float32(0.1)) * (float32((1e10 - (-2.25 * 3.14159))) * (3.14159 - (-2.25 * (0.1 + 1e10))))))) == 3.14159) && v2) {
                                                    switch v66 {
                                                    case -1:
                                                        v184 := GenS18{F0: 36}
                                                        _ = v184
                                                    case 9223372036854775807:
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            switch v3 {
                                                            case "hello":
                                                                if ((1.5 / ((float64(((0.1 / 1.0) - (-2.25 + float32((3.14159 / 1e10))))) / (3.14159 / (1e10 - float64(-2.25)))) * (3.14159 - (3.14159 + float64((float32((3.14159 * float64(-2.25))) + (1.5 / 3.14159))))))) >= float64((1e10 - ((-2.25 - 1e10) / 1.5)))) {
                                                                    _ = len(v5)
                                                                    var v185 float32 = -2.25
                                                                    switch (fi64 >> ((v1) & 63)) {
                                                                    case 236:
                                                                        switch (fi35 + 127) {
                                                                        case 127:
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                var v186 float32 = (v185 + (v185 - (v185 * v185)))
                                                                                _ = v186
                                                                            default:
                                                                                v187 := GenS12{F0: v30}
                                                                                _ = v187
                                                                            }
                                                                        case 484:
                                                                            var v188 uint32 = 2
                                                                            _ = v188
                                                                        case -913:
                                                                            if v2 {
                                                                                v65 = v65
                                                                            }
                                                                        }
                                                                    case 32767:
                                                                        v189 := GenS19{F0: -2147483648}
                                                                        _ = v189
                                                                    default:
                                                                        v5 = append(v5, v65)
                                                                    }
                                                                    if len(v5) > 0 {
                                                                        v5[0] = v65
                                                                    }
                                                                    _ = v185
                                                                }
                                                            }
                                                        case "a":
                                                            v190 := GenS20{F0: 42}
                                                            _ = v190
                                                        }
                                                    }
                                                    if (700 >= (v65 % ((v65) | 1))) {
                                                        fmt.Println(v3, v3, fi64)
                                                    } else {
                                                        switch v3 {
                                                        case "":
                                                            fmt.Println(fi64)
                                                        case "hello":
                                                            v191 := GenS21{F0: ((fi64 | 42) > (fi64 - (fi35 % ((v66) | 1)))), F1: 201}
                                                            _ = v191
                                                        case "foo bar":
                                                            var v192 bool = false
                                                            _ = v192
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                        fmt.Println(v65, v28, v31)
                                                        switch "foo bar" {
                                                        case "hello":
                                                            fmt.Println(v2)
                                                        case "\n\t":
                                                            if (!(v2 || false)) {
                                                                var v193 uint = v1
                                                                v194 := GenS15{F0: ("a" + v3), F1: 0, F2: -128}
                                                                _ = v193
                                                                _ = v194
                                                            }
                                                        default:
                                                            fmt.Println(fi35, v30, v66)
                                                        }
                                                        v195 := GenS4{F0: 376, F1: v2}
                                                        switch v3 {
                                                        case "\n\t":
                                                            switch (fi64 & (fi35 / (((fi64 << ((0) & 63))) | 1))) {
                                                            case 869:
                                                                var v196 uint = v1
                                                                _ = v196
                                                            case -742:
                                                                switch (v3 + "a") {
                                                                case "hello":
                                                                    v3 = "\n\t"
                                                                default:
                                                                    var v197 uint = 65535
                                                                    _ = v197
                                                                }
                                                            }
                                                        case "foo bar":
                                                            switch int(v30) {
                                                            case -343:
                                                                var v198 int = 973
                                                                _ = v198
                                                            }
                                                        }
                                                        v28 = (v28 + -1)
                                                        _ = v195
                                                    }
                                                } else {
                                                    var v199 float64 = 3.14159
                                                    switch ("" + v31) {
                                                    case "hello":
                                                        if true {
                                                            var v200 uint32 = (v65 ^ 2147483647)
                                                            var v201 int16 = 855
                                                            _ = v200
                                                            _ = v201
                                                        }
                                                    default:
                                                        v199 = -2.25
                                                    }
                                                    _ = v199
                                                }
                                            default:
                                                v27 = v27
                                            }
                                        }
                                    }
                                    var v202 bool = (63 <= -2147483648)
                                    fmt.Println(v28, v2)
                                    _ = v202
                                } else {
                                    v4[v31] = v65
                                    v203 := v4[(v31 + v3)]
                                    _ = v203
                                }
                                var v204 uint = (v1 & (v1 << ((65535) & 63)))
                                _ = v204
                            }
                        }
                        _ = len(v4)
                        _ = v66
                    }
                    _ = fi64
                    _ = v65
                }
                _ = fi35
            }
            var v205 string = "\n\t"
            _ = v205
        } else {
            v31 = v3
            for fi206 := 0; fi206 < 14; fi206++ {
                v1 = (v1 & v1)
                if false {
                    if (124 <= 255) {
                        v3 = (("a" + v3) + "\n\t")
                        if (0 < 1) {
                            var v207 bool = (v2 && (-2.25 > -2.25))
                            v3 = v31
                            _ = v207
                        }
                    } else {
                        _ = len(v4)
                        for fi208 := 0; fi208 < 19; fi208++ {
                            delete(v4, ("\n\t" + v31))
                            var v209 string = (v3 + (v31 + (v3 + (v3 + (("" + "") + v31)))))
                            var v210 uint32 = 73
                            if true {
                                var v211 int16 = -125
                                _ = v211
                            } else {
                                v30 += (v30 % (((v1 ^ v1)) | 1))
                                if ((v27 >> ((v30) & 63)) >= (v27 & v27)) {
                                    if true {
                                        var v212 int16 = -1
                                        switch int(v22) {
                                        case -519:
                                            var v213 float32 = 0.1
                                            _ = v213
                                        case 2147483647:
                                            v209 = v3
                                        }
                                        if false {
                                            v3 = "\n\t"
                                            if (true || (v2 || (182 == 32767))) {
                                                var v214 int32 = 283
                                                _ = v214
                                            } else {
                                                v210 -= uint32(v30)
                                                if false {
                                                    v209 = v31
                                                    v2 = v2
                                                    v27 = (v27 - -508)
                                                    fmt.Println(fi208)
                                                    v2 = (v2 || (v2 && false))
                                                } else {
                                                    var v215 float32 = 0.0
                                                    fmt.Println(v215)
                                                    var v216 int32 = int32(v210)
                                                    v212 -= 1
                                                    _ = v215
                                                    _ = v216
                                                }
                                                _ = len(v5)
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        _ = v212
                                    } else {
                                        var v217 int32 = 2147483647
                                        switch "\n\t" {
                                        case "hello":
                                            v3 = v31
                                        case "foo bar":
                                            v2 = v2
                                        case "":
                                            var v218 int32 = v217
                                            _ = v218
                                        }
                                        _ = v217
                                    }
                                    v27 += -474
                                    v219 := GenS22{F0: v27, F1: (^(v210))}
                                    _ = v219
                                } else {
                                    var v220 string = v209
                                    v30 += 255
                                    v209 = ""
                                    _ = v220
                                }
                                var v221 uint16 = 1
                                _ = v221
                            }
                            _ = fi208
                            _ = v209
                            _ = v210
                        }
                    }
                    if (v27 > -641) {
                        if v2 {
                            fmt.Println(v3)
                            fmt.Println(v27)
                            if v2 {
                                if ((v1 * (^(v1))) == (v30 / ((18446744073709551615) | 1))) {
                                    for fi222 := 0; fi222 < 4; fi222++ {
                                        var v223 int16 = -58
                                        v30 = v1
                                        _ = fi222
                                        _ = v223
                                    }
                                    var v224 int32 = 0
                                    if true {
                                        v225 := GenS16{F0: "", F1: 0, F2: 2147483647}
                                        fmt.Println(v1, fi206)
                                        _ = v225
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    _ = v224
                                }
                                var v226 int = fi206
                                v227 := GenS3{F0: (!v2), F1: 42}
                                _ = v226
                                _ = v227
                            } else {
                                var v228 float64 = 0.0
                                if (v2 || v2) {
                                    switch (-(fi206)) {
                                    case 2147483647:
                                        var v229 int64 = -1
                                        _ = v229
                                    case -2147483648:
                                        var v230 float32 = float32(-2.25)
                                        _ = v230
                                    }
                                    for fi231 := 0; fi231 < 6; fi231++ {
                                        if true {
                                            var v232 int = (^(fi206))
                                            v233 := GenS11{F0: 128, F1: 127}
                                            switch fi206 {
                                            case 135:
                                                v28 = 100
                                            }
                                            _ = v232
                                            _ = v233
                                        } else {
                                            fmt.Println(fi206)
                                            var v234 uint32 = 557
                                            _ = v234
                                        }
                                        v235 := GenS12{F0: (^(v30))}
                                        v28 += v28
                                        v236 := GenS23{F0: v22}
                                        var v237 int64 = (v27 + v27)
                                        _ = fi231
                                        _ = v235
                                        _ = v236
                                        _ = v237
                                    }
                                }
                                _ = v228
                            }
                        } else {
                            var v238 int16 = 42
                            _ = v238
                        }
                    }
                } else {
                    fmt.Println(v28, v1)
                }
                switch (v3 + (v3 + v31)) {
                case "\n\t":
                    v3 = v3
                case "hello":
                    v5 = append(v5, 2147483648)
                case "":
                    switch (fi206 - (fi206 - (fi206 | fi206))) {
                    case 42:
                        v2 = v2
                    case 32767:
                        var v239 int = (fi206 - int(v1))
                        _ = v239
                    default:
                        v240 := GenS24{F0: v22, F1: v2, F2: 143}
                        _ = v240
                    }
                }
                _ = fi206
            }
            for fi241 := 0; fi241 < 11; fi241++ {
                if v2 {
                    var v242 uint8 = 2
                    fmt.Println(v27, v2, v3)
                    for fi243 := 0; fi243 < 10; fi243++ {
                        v244 := GenS25{F0: v242, F1: 2, F2: 2147483647}
                        v3 = (((v3 + (v3 + v3)) + "hello") + "foo bar")
                        _ = fi243
                        _ = v244
                    }
                    _ = v242
                } else {
                    switch v31 {
                    case "":
                        if false {
                            var v245 string = v3
                            if v2 {
                                var v246 int16 = 42
                                for fi247 := 0; fi247 < 14; fi247++ {
                                    var v248 uint32 = 42
                                    _ = fi247
                                    _ = v248
                                }
                                _ = v246
                            }
                            for fi249 := 0; fi249 < 14; fi249++ {
                                var v250 int64 = 42
                                v251 := GenS10{F0: 2147483648, F1: v2, F2: v1}
                                _ = fi249
                                _ = v250
                                _ = v251
                            }
                            _ = v245
                        }
                    case "\n\t":
                        v252 := GenS26{F0: (v1 | 553), F1: 817}
                        _ = v252
                    default:
                        if len(v5) > 0 {
                            v5[0] = 2147483648
                        }
                    }
                    v22 = (^(v28))
                    v30 = 4294967295
                }
                _ = fi241
            }
        }
    }
    v1 -= (v1 ^ 1)
    fmt.Println(v22, v31)
    var v253 int32 = 628
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v253:", v253)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v253
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
