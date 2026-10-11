package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint32; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint64; F1 int8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint32; F1 string; F2 uint8 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint }
func genFn0() uint {
    var v7 uint8 = 200
    switch 32767 {
    case 127:
        var v8 int16 = -879
        _ = v8
    }
    _ = v7
    return 795
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    var v4 string = ("" + ("a" + ("\n\t" + ((("foo bar" + (("a" + "") + "\n\t")) + "hello") + "a"))))
    var v5 uint8 = 128
    var v6 bool = (42 > 811)
    v4 = ("a" + "hello")
    if v6 {
        fmt.Println(v3, v3)
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        if v6 {
            for fi10 := 0; fi10 < 10; fi10++ {
                if ((3.14159 * 3.14159) == 1e10) {
                    var v11 int16 = 42
                    switch v4 {
                    case "\n\t":
                        switch -2147483648 {
                        case 269:
                            fmt.Println(fi9)
                        case -1:
                            if (!(2147483647 <= 65535)) {
                                var v12 uint8 = 164
                                var v13 int64 = v3
                                var v14 uint = 42
                                if v6 {
                                    var v15 int64 = (^(v3))
                                    if v6 {
                                        switch (fi10 | fi9) {
                                        case 812:
                                            v16 := []int{42}
                                            _ = v16
                                        }
                                        v5 = v12
                                        if true {
                                            var v17 int64 = (v3 + v15)
                                            if (v1 >= 545) {
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v1)
                                                case "hello":
                                                    if v6 {
                                                        var v18 uint = v14
                                                        v19 := []int{fi9, (fi9 & fi10), fi10}
                                                        if v6 {
                                                            fmt.Println(v3)
                                                            v20 := GenS1{F0: v11}
                                                            _ = v20
                                                        }
                                                        _ = v18
                                                        _ = v19
                                                    }
                                                case "a":
                                                    fmt.Println(v11)
                                                default:
                                                    if (-128 == -9) {
                                                        switch v4 {
                                                        case "foo bar":
                                                            v21 := v2[2147483647]
                                                            _ = v21
                                                        }
                                                        if ((true || v6) || false) {
                                                            switch int(v1) {
                                                            case 9223372036854775807:
                                                                v5 -= 127
                                                            case -812:
                                                                switch (fi9 / ((1) | 1)) {
                                                                case -518:
                                                                    if false {
                                                                        v22 := GenS2{F0: 2}
                                                                        _ = v22
                                                                    }
                                                                    if v6 {
                                                                        switch fi9 {
                                                                        case 32768:
                                                                            fmt.Println(v14)
                                                                        case 851:
                                                                            if (v6 || (950 > 2)) {
                                                                                if (!("hello" >= v4)) {
                                                                                    var v23 uint32 = 930
                                                                                    v24 := GenS3{F0: 2, F1: 2}
                                                                                    if v6 {
                                                                                        var v25 float32 = -2.25
                                                                                        switch (fi10 * (fi9 + (fi10 - (fi9 * (-(fi10)))))) {
                                                                                        case 91:
                                                                                            switch (fi10 & fi9) {
                                                                                            case -851:
                                                                                                v12 = (v12 >> ((0) & 63))
                                                                                            case 42:
                                                                                                delete(v2, 2147483647)
                                                                                            case -2147483648:
                                                                                                if (((0.1 - -2.25) > (((float64(v25) * (((1e10 - 3.14159) - float64(v25)) + float64(v25))) * -2.25) * float64((1e10 / 0.1)))) && v6) {
                                                                                                    v25 = (3.14159 - -2.25)
                                                                                                    v14 = v14
                                                                                                    v26 := GenS4{F0: ("hello" + ((v4 + (("\n\t" + ("\n\t" + "\n\t")) + ("" + (v4 + (v4 + ("a" + v4)))))) + ("\n\t" + (v4 + "hello")))), F1: v23}
                                                                                                    _ = v26
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    fmt.Println(v17, v13)
                                                                                                    var v27 int32 = 1
                                                                                                    v28 := []int{fi9, 9223372036854775807, (-(fi9)), int(v14)}
                                                                                                    _ = v27
                                                                                                    _ = v28
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v4, v4)
                                                                                            v29 := GenS3{F0: 65535, F1: 2}
                                                                                            _ = v29
                                                                                        default:
                                                                                            switch (v4 + "a") {
                                                                                            case "\n\t":
                                                                                                v25 += ((v25 * (v25 / (v25 * (v25 * 1e10)))) + (float32((0.0 / 1.0)) / 3.14159))
                                                                                            case "":
                                                                                                if (0 == 1) {
                                                                                                    var v30 uint8 = (v5 & (v12 / ((v12) | 1)))
                                                                                                    _ = v30
                                                                                                }
                                                                                            case "a":
                                                                                                fmt.Println(v13, v5, v23)
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v14)
                                                                                        _ = v25
                                                                                    }
                                                                                    _ = v23
                                                                                    _ = v24
                                                                                } else {
                                                                                    _ = len(v2)
                                                                                }
                                                                            } else {
                                                                                if (fi9 != (fi9 * 4294967295)) {
                                                                                    fmt.Println(v13, v14)
                                                                                }
                                                                                v5 = (v5 >> (((v14 | v14)) & 63))
                                                                            }
                                                                        default:
                                                                            v31 := GenS1{F0: (v11 + v11)}
                                                                            if ((v17 <= 1) && v6) {
                                                                                fmt.Println(v1, v15)
                                                                            }
                                                                            if v6 {
                                                                                if true {
                                                                                    switch (v4 + (((v4 + (("foo bar" + v4) + v4)) + "hello") + "foo bar")) {
                                                                                    case "":
                                                                                        v4 = v4
                                                                                    case "a":
                                                                                        v13 = v17
                                                                                        switch v4 {
                                                                                        case "a":
                                                                                            v17 += v13
                                                                                            if (!v6) {
                                                                                                v32 := v2[620]
                                                                                                if v6 {
                                                                                                    var v33 bool = (1 < 42)
                                                                                                    _ = v33
                                                                                                }
                                                                                                _ = v32
                                                                                            } else {
                                                                                                v5 += 255
                                                                                                v15 += (v15 % ((v3) | 1))
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v5 = (^(v12))
                                                                                        case "":
                                                                                            if false {
                                                                                                v34 := GenS5{F0: int16(fi9), F1: 2147483647, F2: ((!(8 > 979)) && (!((v5 >= (v12 + (v5 & 153))) && ((v12 & (^(v5))) == 32))))}
                                                                                                v15 = 9
                                                                                                _ = v34
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v5, v15)
                                                                                        }
                                                                                    case "hello":
                                                                                        if ((v12 << (((^(v14))) & 63)) < v12) {
                                                                                            v14 += 873
                                                                                        }
                                                                                    default:
                                                                                        v35 := GenS2{F0: v14}
                                                                                        _ = v35
                                                                                    }
                                                                                } else {
                                                                                    _ = len(v2)
                                                                                    v36 := []int{881, fi10, fi9, fi9}
                                                                                    var v37 float32 = 0.1
                                                                                    _ = v36
                                                                                    _ = v37
                                                                                }
                                                                            }
                                                                            _ = v31
                                                                        }
                                                                        delete(v2, 42)
                                                                        v6 = (!false)
                                                                        if v6 {
                                                                            v38 := []int{fi10}
                                                                            _ = v38
                                                                        }
                                                                        var v39 float64 = float64(float32(1e10))
                                                                        _ = v39
                                                                    } else {
                                                                        if false {
                                                                            var v40 uint8 = 1
                                                                            switch fi10 {
                                                                            case -2147483648:
                                                                                switch v4 {
                                                                                case "\n\t":
                                                                                    delete(v2, 953)
                                                                                }
                                                                            case 32767:
                                                                                v41 := GenS5{F0: -32768, F1: 485, F2: true}
                                                                                _ = v41
                                                                            default:
                                                                                fmt.Println(v5)
                                                                            }
                                                                            _ = v40
                                                                        }
                                                                        var v42 float32 = 1.5
                                                                        v43 := GenS5{F0: v11, F1: 449, F2: (!(true && (true || (94 >= 9))))}
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                case 4294967295:
                                                                    if (true || v6) {
                                                                        v14 -= v14
                                                                        if true {
                                                                            v17 -= (v3 + v13)
                                                                            v6 = v6
                                                                        } else {
                                                                            switch (-(fi9)) {
                                                                            case 154:
                                                                                switch (v4 + (v4 + "\n\t")) {
                                                                                case "a":
                                                                                    var v44 float32 = (0.0 * (0.1 + 1e10))
                                                                                    _ = v44
                                                                                }
                                                                            case -625:
                                                                                v45 := GenS4{F0: (("foo bar" + v4) + (v4 + v4)), F1: 65535}
                                                                                _ = v45
                                                                            case 1:
                                                                                switch fi9 {
                                                                                case 32768:
                                                                                    var v46 float64 = ((float64(((((3.14159 - float32((-2.25 / 1.5))) / 1.0) + float32(((-2.25 + 3.14159) / 1.0))) + 3.14159)) - 1e10) + float64(-2.25))
                                                                                    _ = v46
                                                                                case 4294967295:
                                                                                    switch fi10 {
                                                                                    case 9223372036854775807:
                                                                                        var v47 uint32 = uint32(v14)
                                                                                        _ = v47
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v48 int16 = 1
                                                                                _ = v48
                                                                            }
                                                                            switch fi10 {
                                                                            case -847:
                                                                                v17 = (v17 % ((v15) | 1))
                                                                            default:
                                                                                fmt.Println(v3, v5, v3)
                                                                                var v49 uint = 2
                                                                                _ = v49
                                                                            }
                                                                            var v50 int8 = 1
                                                                            _ = v50
                                                                        }
                                                                    }
                                                                }
                                                            case -198:
                                                                var v51 float64 = 3.14159
                                                                _ = v51
                                                            default:
                                                                v11 -= v11
                                                            }
                                                            v52 := GenS1{F0: -930}
                                                            var v53 uint32 = uint32(v5)
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    } else {
                                                        var v54 uint32 = 2
                                                        var v55 uint16 = 127
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                }
                                                var v56 float32 = (0.1 / float32(1e10))
                                                var v57 bool = true
                                                _ = v56
                                                _ = v57
                                            } else {
                                                _ = len(v2)
                                                var v58 uint32 = 2147483648
                                                v2[(v58 % ((v58) | 1))] = (^(v58))
                                                _ = v58
                                            }
                                            v14 += (v14 / (((v14 | 65535)) | 1))
                                            _ = v17
                                        }
                                    }
                                    _ = v15
                                }
                                v59 := []int{(fi9 % ((406) | 1)), (fi10 - (-(fi9))), (fi10 * 2147483647), (fi10 * (fi10 + (fi9 << (((^(v14))) & 63))))}
                                v4 = "\n\t"
                                fmt.Println(v14)
                                _ = v12
                                _ = v13
                                _ = v14
                                _ = v59
                            } else {
                                fmt.Println(v6)
                                v60 := GenS4{F0: v4, F1: 2147483648}
                                v3 = v3
                                switch (fi10 | fi10) {
                                case 127:
                                    var v61 uint = 65535
                                    _ = v61
                                case 128:
                                    v62 := GenS1{F0: (v11 - v11)}
                                    _ = v62
                                }
                                _ = v60
                            }
                        case 32768:
                            fmt.Println(fi10, v11, v1)
                        }
                    case "a":
                        v1 = int32(v11)
                    default:
                        v63 := []int{}
                        _ = v63
                    }
                    v64 := v2[967]
                    fmt.Println(v6, fi10)
                    var v65 uint8 = (v5 & 56)
                    fmt.Println(fi9, v4, fi10)
                    fmt.Println(fi9, fi10, fi9)
                    _ = v11
                    _ = v64
                    _ = v65
                }
                v66 := GenS6{F0: 4294967295}
                _ = fi10
                _ = v66
            }
        } else {
            var v67 float32 = 1.5
            var v68 uint64 = 1
            if false {
                v69 := GenS6{F0: uint(v68)}
                if v6 {
                    var v70 uint8 = v5
                    var v71 float64 = (-2.25 * float64(1e10))
                    _ = v70
                    _ = v71
                }
                _ = v69
            }
            _ = v67
            _ = v68
        }
        _ = fi9
    }
    if (-834 <= 42) {
        for fi72 := 0; fi72 < 13; fi72++ {
            v3 += -94
            v73 := []int{fi72}
            var v74 int64 = (v3 * (v3 + (v3 * v3)))
            _ = fi72
            _ = v73
            _ = v74
        }
    }
    var v75 int8 = 42
    v2[1] = 58
    if (1.5 == 1e10) {
        var v76 uint8 = (v5 ^ (v5 * v5))
        _ = v76
    }
    v1 -= v1
    v75 -= -126
    v2[479] = 2147483648
    if false {
        if true {
            if (!(false && ((!false) && v6))) {
                var v77 int = 128
                for fi78 := 0; fi78 < 2; fi78++ {
                    fmt.Println(v4, v6)
                    switch (v4 + (v4 + v4)) {
                    case "foo bar":
                        v6 = ("\n\t" <= v4)
                    case "a":
                        v5 = v5
                    case "\n\t":
                        for fi79 := 0; fi79 < 16; fi79++ {
                            switch int(v5) {
                            case 32767:
                                v80 := GenS7{F0: 4294967295}
                                _ = v80
                            default:
                                if v6 {
                                    v81 := []int{}
                                    switch (v4 + v4) {
                                    case "foo bar":
                                        v82 := GenS3{F0: 65535, F1: 18446744073709551615}
                                        _ = v82
                                    default:
                                        v3 = v3
                                    }
                                    var v83 uint = 0
                                    fmt.Println(v1, v4, v77)
                                    _ = v81
                                    _ = v83
                                } else {
                                    fmt.Println(v5, v4)
                                    if v6 {
                                        if false {
                                            var v84 int16 = 0
                                            _ = len(v2)
                                            if v6 {
                                                v85 := []int{(fi78 & (v77 ^ 756)), (v77 / ((fi78) | 1)), (fi78 % (((fi79 / ((int(v84)) | 1))) | 1)), (fi79 ^ (v77 >> ((1) & 63)))}
                                                _ = v85
                                            }
                                            v75 -= v75
                                            _ = v84
                                        }
                                    }
                                    fmt.Println(v5)
                                    var v86 uint = 65535
                                    v87 := GenS8{F0: 333, F1: (v75 & v75)}
                                    v77 = (v77 - 930)
                                    fmt.Println(v3)
                                    _ = v86
                                    _ = v87
                                }
                            }
                            if (!(942 >= 1)) {
                                v4 = v4
                            } else {
                                v88 := v2[0]
                                var v89 int16 = 376
                                v90 := GenS9{F0: ("\n\t" <= "a")}
                                _ = v88
                                _ = v89
                                _ = v90
                            }
                            var v91 int32 = v1
                            _ = fi79
                            _ = v91
                        }
                    default:
                        for fi92 := 0; fi92 < 18; fi92++ {
                            _ = len(v2)
                            switch "foo bar" {
                            case "\n\t":
                                switch (fi78 / ((v77) | 1)) {
                                case -10:
                                    v93 := GenS8{F0: 65535, F1: v75}
                                    _ = v93
                                case -941:
                                    var v94 uint32 = 1
                                    _ = v94
                                case 32768:
                                    v95 := v2[87]
                                    _ = v95
                                default:
                                    v96 := v2[42]
                                    _ = v96
                                }
                            case "foo bar":
                                switch fi92 {
                                case -1:
                                    var v97 int64 = (v3 % (((v3 * (v3 | v3))) | 1))
                                    switch v77 {
                                    case -1:
                                        v5 = (v5 ^ 86)
                                        var v98 int64 = v97
                                        _ = v98
                                    case 4294967295:
                                        var v99 uint64 = 396
                                        _ = v99
                                    }
                                    _ = len(v2)
                                    _ = v97
                                case 570:
                                    if v6 {
                                        if true {
                                            if false {
                                                v6 = v6
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v5 = v5
                                                case "hello":
                                                    var v100 uint = 18446744073709551615
                                                    _ = v100
                                                case "foo bar":
                                                    if (true && (42 < 451)) {
                                                        v101 := []int{fi92, fi78, -415, fi78}
                                                        if v6 {
                                                            fmt.Println(v75, v1, v4)
                                                            var v102 int64 = v3
                                                            var v103 int = fi92
                                                            v1 -= (^(v1))
                                                            _ = v102
                                                            _ = v103
                                                        } else {
                                                            fmt.Println(v1, v1)
                                                        }
                                                        v6 = (v5 < (^(v5)))
                                                        _ = v101
                                                    }
                                                }
                                            }
                                            v5 = 0
                                            v104 := GenS3{F0: 65535, F1: 165}
                                            _ = v104
                                        } else {
                                            if (!false) {
                                                v105 := v2[0]
                                                v106 := GenS10{F0: true, F1: 833}
                                                _ = v105
                                                _ = v106
                                            }
                                            v107 := GenS11{F0: false, F1: 288}
                                            if true {
                                                v108 := GenS12{F0: (v75 & v75), F1: (v77 >> ((450) & 63))}
                                                _ = v108
                                            }
                                            _ = v107
                                        }
                                        var v109 uint16 = 65535
                                        if false {
                                            var v110 int16 = 32767
                                            _ = v110
                                        } else {
                                            v75 = (v75 << ((2) & 63))
                                            var v111 uint32 = 350
                                            _ = v111
                                        }
                                        _ = v109
                                    }
                                case 2147483647:
                                    v5 += (v5 >> ((18446744073709551615) & 63))
                                default:
                                    v77 -= fi78
                                }
                            default:
                                if (!v6) {
                                    if v6 {
                                        var v112 int16 = 42
                                        v113 := []int{(fi92 - fi92)}
                                        v75 += (v75 << ((67) & 63))
                                        _ = v112
                                        _ = v113
                                    } else {
                                        v114 := GenS9{F0: true}
                                        var v115 int16 = 0
                                        _ = v114
                                        _ = v115
                                    }
                                } else {
                                    if v6 {
                                        if (uint(v5) >= genFn0()) {
                                            if (!v6) {
                                                v3 += 9223372036854775807
                                                v116 := GenS5{F0: -32768, F1: 1, F2: (768 >= 2147483648)}
                                                _ = v116
                                            } else {
                                                v117 := v2[1]
                                                var v118 string = v4
                                                _ = v117
                                                _ = v118
                                            }
                                        }
                                    }
                                    v6 = false
                                    var v119 uint32 = 55
                                    var v120 int16 = 144
                                    _ = v119
                                    _ = v120
                                }
                            }
                            _ = fi92
                        }
                    }
                    for fi121 := 0; fi121 < 20; fi121++ {
                        if true {
                            switch 885 {
                            case 32767:
                                v122 := GenS4{F0: ((((("foo bar" + v4) + v4) + (v4 + v4)) + v4) + v4), F1: 2}
                                _ = v122
                            default:
                                if v6 {
                                    fmt.Println(v5, v1)
                                    switch 719 {
                                    case 4294967295:
                                        if (0.0 >= (float32(0.0) + 3.14159)) {
                                            v123 := v2[777]
                                            if v6 {
                                                switch (-(fi78)) {
                                                case 1:
                                                    if (v6 || false) {
                                                        if (!((^(v1)) <= v1)) {
                                                            v124 := GenS13{F0: 0}
                                                            v6 = true
                                                            _ = v124
                                                        }
                                                        var v125 uint = 733
                                                        var v126 int32 = v1
                                                        v75 = (v75 / (((v75 % ((73) | 1))) | 1))
                                                        _ = v125
                                                        _ = v126
                                                    } else {
                                                        var v127 uint64 = 42
                                                        v77 = 504
                                                        _ = v127
                                                    }
                                                case -647:
                                                    if ((0.1 - (3.14159 + (float64(float32(3.14159)) + 1e10))) < 1.5) {
                                                        v3 += (v3 & v3)
                                                        var v128 int32 = v1
                                                        _ = v128
                                                    } else {
                                                        if (false && (!v6)) {
                                                            fmt.Println(v4)
                                                            var v129 int = (v77 >> ((4294967295) & 63))
                                                            v130 := []int{(v77 & 32767), -404, 32767}
                                                            _ = v129
                                                            _ = v130
                                                        }
                                                        v77 = (fi121 + int(v123))
                                                        v1 += v1
                                                    }
                                                case 159:
                                                    v131 := GenS14{F0: v123}
                                                    _ = v131
                                                default:
                                                    if v6 {
                                                        var v132 int16 = 472
                                                        var v133 int8 = 1
                                                        v134 := []int{}
                                                        _ = len(v2)
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    } else {
                                                        v135 := []int{fi121, v77, -93}
                                                        v136 := GenS2{F0: 1}
                                                        v2[0] = v123
                                                        _ = v135
                                                        _ = v136
                                                    }
                                                }
                                                var v137 uint8 = 2
                                                v138 := []int{(v77 / ((-1) | 1)), (fi121 & fi121), 51, -708}
                                                _ = v137
                                                _ = v138
                                            }
                                            _ = v123
                                        } else {
                                            var v139 uint = 65535
                                            var v140 uint32 = 0
                                            _ = v139
                                            _ = v140
                                        }
                                    case 2147483647:
                                        if false {
                                            var v141 int16 = 1
                                            v142 := []int{(^(fi121)), fi78}
                                            var v143 uint8 = v5
                                            _ = v141
                                            _ = v142
                                            _ = v143
                                        } else {
                                            v1 = v1
                                        }
                                    default:
                                        fmt.Println(v6)
                                        var v144 uint16 = 32767
                                        var v145 int8 = (^(v75))
                                        _ = v144
                                        _ = v145
                                    }
                                    var v146 uint64 = 42
                                    _ = v146
                                } else {
                                    v4 = v4
                                    var v147 bool = true
                                    v148 := []int{}
                                    v149 := GenS15{F0: 1, F1: -128}
                                    _ = v147
                                    _ = v148
                                    _ = v149
                                }
                            }
                            switch (v4 + v4) {
                            case "foo bar":
                                if v6 {
                                    var v150 int32 = v1
                                    _ = v150
                                }
                            }
                            if false {
                                var v151 string = v4
                                var v152 uint64 = 313
                                _ = v151
                                _ = v152
                            }
                        }
                        var v153 uint8 = (v5 | v5)
                        var v154 uint8 = v5
                        v154 -= v153
                        if ((^(v75)) != 127) {
                            v6 = v6
                        }
                        var v155 float64 = ((0.1 - 1.5) / 1.5)
                        _ = fi121
                        _ = v153
                        _ = v154
                        _ = v155
                    }
                    _ = fi78
                }
                if v6 {
                    v75 += 42
                    var v156 int = (v77 ^ v77)
                    _ = v156
                } else {
                    var v157 int16 = 1000
                    v158 := GenS16{F0: (v1 & v1)}
                    _ = v157
                    _ = v158
                }
                _ = v77
            }
            if v6 {
                v159 := GenS17{F0: 2147483648, F1: (v4 + v4), F2: (v5 ^ (^(v5)))}
                v75 -= v75
                fmt.Println(v4, v4, v3)
                _ = v159
            }
        }
    } else {
        v160 := GenS18{F0: ("\n\t" + "\n\t")}
        v4 = (v4 + ((v4 + "a") + "a"))
        _ = v160
    }
    var v161 int = 4294967295
    delete(v2, 42)
    v162 := GenS10{F0: false, F1: 18446744073709551615}
    v4 = ((v4 + (((v4 + ("a" + ((v4 + "") + v4))) + (v4 + v4)) + (v4 + (("\n\t" + ("foo bar" + "\n\t")) + "")))) + (v4 + ("" + "a")))
    switch v4 {
    case "":
        for fi163 := 0; fi163 < 9; fi163++ {
            if (((3.14159 + 0.1) - (0.1 / (float64(0.1) / 1.5))) < 3.14159) {
                var v164 float32 = 3.14159
                _ = v164
            }
            _ = fi163
        }
    }
    v1 = -2147483648
    if true {
        fmt.Println(v3)
        v165 := GenS19{F0: 2}
        var v166 uint = 0
        _ = v165
        _ = v166
    }
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v161
    _ = v162
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
}
