package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int16; F1 uint16; F2 int }
type GenS7 struct { F0 int; F1 string; F2 int16 }
type GenS8 struct { F0 int8; F1 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64; F1 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 int8; F2 uint16 }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8; F1 string; F2 int8 }
type GenS16 struct { F0 string; F1 string; F2 int32 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 bool; F1 int8 }
type GenS22 struct { F0 uint; F1 string; F2 int16 }
func genFn0(p0 int8, p1 string) bool {
    p0 -= p0
    fmt.Println(p1)
    _ = p0
    _ = p1
    return (42 != 65535)
}
func genFn1() bool {
    v4 := []int{2147483647, 209, -620, 381}
    var v5 bool = (!true)
    v5 = (!(22 <= 47))
    if (v5 && v5) {
        fmt.Println(v5)
        v6 := GenS1{F0: 42, F1: 9223372036854775807, F2: 100}
        if len(v4) > 0 {
            v4[0] = 128
        }
        _ = v6
    }
    v7 := GenS2{F0: 1, F1: 0, F2: "hello"}
    _ = v4
    _ = v5
    _ = v7
    return v5
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = ((float64(3.14159) - 1.5) / 1.0)
    v3 := make(map[int]string)
    v2 += (float64((3.14159 / float32(v2))) - float64(0.1))
    var v8 uint16 = 875
    v9 := GenS3{F0: v8}
    var v10 uint16 = 42
    if true {
        if true {
            if (!true) {
                v2 -= v2
                v10 += v10
                for fi11 := 0; fi11 < 17; fi11++ {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        if (!false) {
                            var v13 int = 128
                            _ = v13
                        } else {
                            fmt.Println(fi11, v10)
                            if (float32(1e10) < 0.1) {
                                switch "hello" {
                                case "":
                                    v10 -= v8
                                case "\n\t":
                                    v2 += (3.14159 + (v2 / float64(-2.25)))
                                case "hello":
                                    v14 := GenS4{F0: 2147483647}
                                    _ = v14
                                default:
                                    if (false || true) {
                                        v8 += v8
                                        v8 = 663
                                        if true {
                                            var v15 uint = 1
                                            if (25 < 127) {
                                                v8 = 32767
                                            } else {
                                                v16 := []int{}
                                                var v17 uint8 = 42
                                                _ = v16
                                                _ = v17
                                            }
                                            v2 = (float64(-2.25) - float64(0.1))
                                            v10 += v8
                                            _ = v15
                                        } else {
                                            fmt.Println(v10)
                                        }
                                        v10 += v10
                                    }
                                }
                            }
                        }
                        switch "" {
                        case "\n\t":
                            if (947 == 1) {
                                if (9223372036854775807 < -1) {
                                    v18 := GenS5{F0: 65535, F1: 2}
                                    switch fi12 {
                                    case 32768:
                                        fmt.Println(v8, v10)
                                    case 127:
                                        v8 += v10
                                    case 128:
                                        v10 -= (v8 & (v8 ^ v10))
                                    default:
                                        var v19 float64 = (float64((0.1 * 0.0)) + float64(float32(0.1)))
                                        _ = v19
                                    }
                                    v20 := GenS4{F0: 1}
                                    _ = v18
                                    _ = v20
                                }
                            } else {
                                _ = len(v1)
                                if (0 == 2147483648) {
                                    v21 := []int{fi11, (fi12 | fi12)}
                                    var v22 float64 = (v2 * v2)
                                    if (0 > 18446744073709551615) {
                                        v23 := v1[fi12]
                                        var v24 float64 = (0.0 * 3.14159)
                                        v24 = float64((3.14159 / 1.0))
                                        switch (-(fi12)) {
                                        case 2147483647:
                                            v21 = append(v21, fi12)
                                        case 42:
                                            v25 := GenS6{F0: 1, F1: v8, F2: (fi12 >> ((65535) & 63))}
                                            _ = v25
                                        case 9223372036854775807:
                                            if true {
                                                v8 -= 42
                                                if false {
                                                    if false {
                                                        v23 = 2147483647
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v8)
                                                    case "":
                                                        v22 -= (((1e10 - -2.25) - float64(3.14159)) * ((0.0 - v22) * (3.14159 + v2)))
                                                    default:
                                                        if ((true && (v2 >= -2.25)) || true) {
                                                            if (((115 <= 100) || false) || (0 == 1)) {
                                                                if true {
                                                                    v26 := GenS2{F0: v10, F1: 24, F2: "foo bar"}
                                                                    var v27 int32 = 242
                                                                    if true {
                                                                        if (0 != -585) {
                                                                            v28 := GenS5{F0: (v23 | 2147483648), F1: 255}
                                                                            v3[(-(fi12))] = ("\n\t" + "")
                                                                            fmt.Println(v24, v22)
                                                                            _ = v28
                                                                        } else {
                                                                            var v29 float64 = (float64((1e10 + (3.14159 * float32((v22 / v24))))) / (v22 + 0.1))
                                                                            delete(v3, 9223372036854775807)
                                                                            v8 -= 32767
                                                                            _ = v29
                                                                        }
                                                                        v10 -= v10
                                                                    } else {
                                                                        if (false && false) {
                                                                            v24 = (1e10 / (-2.25 / v24))
                                                                        } else {
                                                                            if true {
                                                                                var v30 uint64 = 2
                                                                                fmt.Println(v22, v23, fi12)
                                                                                _ = v30
                                                                            }
                                                                            var v31 int32 = v27
                                                                            v2 -= v24
                                                                            _ = v31
                                                                        }
                                                                        v27 += (v27 & (v27 & 0))
                                                                        var v32 uint8 = 255
                                                                        if len(v21) > 0 {
                                                                            v21[0] = (fi12 | fi12)
                                                                        }
                                                                        _ = v32
                                                                    }
                                                                    v22 += v22
                                                                    if false {
                                                                        var v33 float32 = 3.14159
                                                                        v34 := GenS7{F0: 4294967295, F1: "a", F2: 1}
                                                                        _ = v33
                                                                        _ = v34
                                                                    }
                                                                    _ = v26
                                                                    _ = v27
                                                                } else {
                                                                    var v35 int64 = 1
                                                                    var v36 uint8 = 43
                                                                    if (((false && (v35 <= (v35 << ((849) & 63)))) || (!false)) || (v35 <= v35)) {
                                                                        if false {
                                                                            var v37 float64 = v22
                                                                            v38 := GenS1{F0: -709, F1: v35, F2: uint8(v23)}
                                                                            switch fi11 {
                                                                            case 4294967295:
                                                                                var v39 uint64 = 380
                                                                                _ = v39
                                                                            case -9223372036854775808:
                                                                                v22 = v37
                                                                            }
                                                                            _ = v37
                                                                            _ = v38
                                                                        } else {
                                                                            var v40 int16 = 32767
                                                                            _ = v40
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v8, v23, fi11)
                                                                            if false {
                                                                                v23 = v23
                                                                                v41 := GenS6{F0: 42, F1: v10, F2: (fi12 + fi12)}
                                                                                fmt.Println(v10, v22, v24)
                                                                                _ = v41
                                                                            } else {
                                                                                var v42 bool = true
                                                                                fmt.Println(v10, v36)
                                                                                _ = v42
                                                                            }
                                                                            v35 -= (v35 - (^(v35)))
                                                                        } else {
                                                                            fmt.Println(v2, v10)
                                                                            if len(v21) >= 2 {
                                                                                v21 = v21[1:len(v21)]
                                                                            }
                                                                        }
                                                                        v23 += (v23 * (v23 >> ((65535) & 63)))
                                                                        v2 += (float64((((0.0 - 0.1) + ((1.5 - ((float32(3.14159) / 1e10) * -2.25)) / -2.25)) + 0.0)) - v24)
                                                                    } else {
                                                                        switch (fi11 << ((uint(v23)) & 63)) {
                                                                        case 127:
                                                                            fmt.Println(v24)
                                                                        case 42:
                                                                            fmt.Println(fi12, v22, v24)
                                                                        case 32768:
                                                                            var v43 uint64 = 549
                                                                            _ = v43
                                                                        default:
                                                                            var v44 int8 = 1
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                    _ = v35
                                                                    _ = v36
                                                                }
                                                                v23 = v23
                                                            } else {
                                                                if (false || (!true)) {
                                                                    var v45 bool = true
                                                                    if v45 {
                                                                        v45 = (((0.0 + ((float32((v24 / 1.0)) * 3.14159) + (3.14159 * (((0.1 * 0.0) - (((0.0 * 0.1) * float32(v22)) - 0.0)) + (float32(0.1) + 0.0))))) - -2.25) <= ((3.14159 * float32((-2.25 * (v24 * (v22 / 1.5))))) * 0.1))
                                                                        v8 = (^(v10))
                                                                        switch -975 {
                                                                        case 533:
                                                                            var v46 uint8 = 128
                                                                            _ = v46
                                                                        }
                                                                    } else {
                                                                        var v47 string = ("a" + "foo bar")
                                                                        var v48 uint8 = 2
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    _ = v45
                                                                }
                                                            }
                                                            if ((293 >= 0) && true) {
                                                                var v49 uint = 1
                                                                var v50 uint16 = (^(v10))
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                if true {
                                                                    if ((!true) || false) {
                                                                        fmt.Println(v2, v2, v2)
                                                                        v23 = v23
                                                                    }
                                                                }
                                                            }
                                                            v51 := GenS3{F0: v8}
                                                            if (42 < 42) {
                                                                var v52 int = 32768
                                                                v52 = int(v23)
                                                                var v53 float32 = -2.25
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            var v54 float32 = 1e10
                                                            fmt.Println(v8, v10, v54)
                                                            _ = v51
                                                            _ = v54
                                                        } else {
                                                            var v55 int8 = -1
                                                            fmt.Println(v8, v24)
                                                            v56 := GenS2{F0: v8, F1: 10, F2: "hello"}
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    if true {
                                                        if (!(!false)) {
                                                            v23 += v23
                                                            if true {
                                                                switch "\n\t" {
                                                                case "":
                                                                    v23 = v23
                                                                }
                                                                if false {
                                                                    var v57 bool = false
                                                                    _ = v57
                                                                } else {
                                                                    if true {
                                                                        switch fi11 {
                                                                        case 32767:
                                                                            var v58 int64 = -9223372036854775808
                                                                            var v59 uint64 = 18446744073709551615
                                                                            fmt.Println(v23, v10)
                                                                            _ = v58
                                                                            _ = v59
                                                                        case 1:
                                                                            if false {
                                                                                fmt.Println(fi12, v2, v8)
                                                                                var v60 uint = 332
                                                                                var v61 int64 = 0
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                        case -2147483648:
                                                                            v8 = (^(v10))
                                                                        default:
                                                                            v22 += 0.1
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if false {
                                                                                _ = len(v21)
                                                                                v24 += (v22 - ((v2 - v24) - 0.0))
                                                                                if (942 >= 165) {
                                                                                    v22 += v24
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        var v62 uint64 = 1
                                                                                        _ = v62
                                                                                    case "foo bar":
                                                                                        var v63 uint32 = v23
                                                                                        _ = v63
                                                                                    }
                                                                                    if true {
                                                                                        if (false && false) {
                                                                                            v2 = v24
                                                                                        } else {
                                                                                            v64 := GenS1{F0: 1, F1: 631, F2: 255}
                                                                                            switch ("hello" + ("a" + "")) {
                                                                                            case "a":
                                                                                                v2 += ((v22 / v2) * v24)
                                                                                            case "\n\t":
                                                                                                v8 += (^(v8))
                                                                                            case "foo bar":
                                                                                                _ = len(v3)
                                                                                            default:
                                                                                                v65 := GenS4{F0: 2147483647}
                                                                                                _ = v65
                                                                                            }
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v66 uint8 = 33
                                                                                v10 = 287
                                                                                v10 = (v10 ^ (^(v10)))
                                                                                var v67 uint64 = 2
                                                                                _ = v66
                                                                                _ = v67
                                                                            } else {
                                                                                v8 += v10
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v2, v10)
                                                                        case "\n\t":
                                                                            var v68 int64 = -9223372036854775808
                                                                            _ = v68
                                                                        default:
                                                                            v2 -= 1.5
                                                                        }
                                                                    } else {
                                                                        switch fi11 {
                                                                        case 0:
                                                                            if (255 != 128) {
                                                                                if (2147483647 != -203) {
                                                                                    v10 = (^(v8))
                                                                                    switch fi12 {
                                                                                    case 32768:
                                                                                        var v69 int = (-(fi11))
                                                                                        _ = v69
                                                                                    case -845:
                                                                                        if true {
                                                                                            v10 = 65535
                                                                                            v70 := GenS5{F0: (^(v23)), F1: 42}
                                                                                            if ((!(((-2.25 - 0.1) - (1.5 - (float32(3.14159) - 1.5))) != (0.0 / (1e10 - (1.5 + (1e10 * 1.5)))))) && true) {
                                                                                                fmt.Println(v24, fi12)
                                                                                                fmt.Println(fi12)
                                                                                            } else {
                                                                                                var v71 float64 = 1e10
                                                                                                v22 += v24
                                                                                                if len(v21) > 0 {
                                                                                                    v21[0] = fi11
                                                                                                }
                                                                                                _ = v71
                                                                                            }
                                                                                            _ = v70
                                                                                        } else {
                                                                                            var v72 int16 = 32767
                                                                                            _ = v72
                                                                                        }
                                                                                        if len(v21) > 0 {
                                                                                            v21[0] = -398
                                                                                        }
                                                                                        v1[fi11] = (v23 | v23)
                                                                                    case -2147483648:
                                                                                        if false {
                                                                                            v73 := v3[128]
                                                                                            var v74 int64 = 283
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        } else {
                                                                                            var v75 int64 = 42
                                                                                            _ = v75
                                                                                        }
                                                                                    default:
                                                                                        if len(v21) > 0 {
                                                                                            v21[0] = (^(fi11))
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v2, fi11)
                                                                            }
                                                                        case -254:
                                                                            var v76 float32 = ((1.5 / 3.14159) / 1.0)
                                                                            _ = v76
                                                                        default:
                                                                            v77 := GenS8{F0: 4, F1: fi12}
                                                                            _ = v77
                                                                        }
                                                                        var v78 bool = true
                                                                        _ = v78
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "foo bar":
                                                                    _ = len(v3)
                                                                case "a":
                                                                    v23 = v23
                                                                case "\n\t":
                                                                    switch 23 {
                                                                    case -1:
                                                                        v79 := v1[fi12]
                                                                        v23 += (v23 << ((1) & 63))
                                                                        v80 := GenS6{F0: 32767, F1: (v10 & 32767), F2: -1}
                                                                        _ = v79
                                                                        _ = v80
                                                                    case 2147483647:
                                                                        v81 := GenS9{F0: 127}
                                                                        _ = v81
                                                                    case 4294967295:
                                                                        var v82 uint16 = (^(v10))
                                                                        _ = v82
                                                                    default:
                                                                        var v83 int8 = 0
                                                                        _ = v83
                                                                    }
                                                                }
                                                                var v84 uint16 = (v8 + 255)
                                                                _ = v84
                                                            }
                                                        } else {
                                                            v85 := GenS10{F0: 42, F1: 65}
                                                            fmt.Println(v24)
                                                            v24 = ((3.14159 / (-2.25 / v22)) / (v22 + (v24 / 3.14159)))
                                                            _ = v85
                                                        }
                                                    }
                                                case "a":
                                                    fmt.Println(v23, v8, v23)
                                                default:
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    switch fi11 {
                                                    case 4294967295:
                                                        v8 += (v10 >> ((440) & 63))
                                                    }
                                                case "hello":
                                                    if ((1 <= 42) && (369 <= 42)) {
                                                        fmt.Println(v8, v24, v22)
                                                    }
                                                case "":
                                                    var v86 uint16 = 32767
                                                    _ = v86
                                                }
                                                v23 -= v23
                                            } else {
                                                var v87 int8 = 1
                                                fmt.Println(v22)
                                                _ = v87
                                            }
                                        default:
                                            v3[fi11] = "foo bar"
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v21
                                    _ = v22
                                }
                            }
                        case "foo bar":
                            var v88 int16 = -32768
                            _ = v88
                        case "a":
                            var v89 float64 = ((v2 + v2) * v2)
                            _ = v89
                        default:
                            if (true || ((("foo bar" + "hello") != "a") || (0 >= 458))) {
                                var v90 int32 = 779
                                v8 = 32767
                                _ = v90
                            }
                        }
                        _ = fi12
                    }
                    switch (fi11 + (fi11 / (((fi11 / ((fi11) | 1))) | 1))) {
                    case 32768:
                        v91 := GenS4{F0: 42}
                        _ = v91
                    case 1:
                        for fi92 := 0; fi92 < 12; fi92++ {
                            switch (fi92 * (fi11 + -981)) {
                            case 4294967295:
                                v93 := GenS9{F0: 226}
                                _ = v93
                            }
                            if (((false || false) || true) && (!(v2 == ((v2 / (((1.5 - v2) / (((v2 + v2) * (-2.25 - (1.5 - v2))) - (v2 + float64(float32(((v2 * ((v2 * v2) - v2)) * 0.1)))))) - v2)) / v2)))) {
                                var v94 uint16 = v10
                                fmt.Println(v10)
                                _ = v94
                            }
                            if (true && (v2 == (v2 + v2))) {
                                v95 := []int{127}
                                if true {
                                    if (-2.25 >= float32(v2)) {
                                        v96 := v1[(fi92 << ((0) & 63))]
                                        _ = v96
                                    }
                                    v10 = 1
                                } else {
                                    switch (fi92 + 128) {
                                    case 127:
                                        v10 += (^(v8))
                                    case -1:
                                        v97 := GenS3{F0: 42}
                                        _ = v97
                                    }
                                    var v98 int8 = 42
                                    _ = v98
                                }
                                _ = v95
                            } else {
                                switch (("" + ("a" + "a")) + "") {
                                case "foo bar":
                                    v10 = v10
                                case "hello":
                                    var v99 uint32 = 612
                                    _ = v99
                                default:
                                    v100 := GenS11{F0: -258}
                                    _ = v100
                                }
                                v101 := GenS12{F0: 1, F1: 0, F2: 1}
                                _ = v101
                            }
                            v102 := []int{(fi11 ^ 2147483647), fi92, (fi92 << ((2) & 63)), fi92}
                            if true {
                                var v103 float32 = (3.14159 + 1e10)
                                switch fi92 {
                                case 127:
                                    switch (-(fi11)) {
                                    case 32767:
                                        switch (fi11 << ((4294967295) & 63)) {
                                        case 42:
                                            v103 = (((0.1 - (3.14159 - 3.14159)) + (((v103 * (v103 / v103)) / 1.5) - float32(v2))) / (v103 / -2.25))
                                        default:
                                            v103 = 0.1
                                        }
                                    default:
                                        var v104 bool = true
                                        _ = v104
                                    }
                                }
                                v2 = (0.0 / v2)
                                _ = v103
                            } else {
                                var v105 int64 = -9223372036854775808
                                if (v10 >= (v8 ^ 2)) {
                                    var v106 int16 = -1
                                    _ = v106
                                } else {
                                    var v107 int32 = 0
                                    if true {
                                        v2 += v2
                                        v108 := GenS12{F0: -113, F1: 57, F2: (v8 - 32767)}
                                        v105 = (-(v105))
                                        _ = v108
                                    }
                                    var v109 int8 = 100
                                    var v110 int = fi92
                                    _ = v107
                                    _ = v109
                                    _ = v110
                                }
                                var v111 bool = true
                                _ = v105
                                _ = v111
                            }
                            _ = fi92
                            _ = v102
                        }
                    default:
                        switch "" {
                        case "a":
                            v112 := GenS8{F0: 1, F1: fi11}
                            _ = v112
                        default:
                            var v113 uint16 = (v8 / (((v10 * (^(v8)))) | 1))
                            _ = v113
                        }
                    }
                    for fi114 := 0; fi114 < 4; fi114++ {
                        if false {
                            switch 0 {
                            case 32767:
                                if (!true) {
                                    var v115 string = ("hello" + "foo bar")
                                    _ = v115
                                } else {
                                    var v116 string = "foo bar"
                                    if (!true) {
                                        v117 := v1[(^(fi11))]
                                        v116 = "foo bar"
                                        switch (-(fi114)) {
                                        case 9223372036854775807:
                                            switch fi114 {
                                            case 42:
                                                fmt.Println(fi11)
                                            default:
                                                v10 = v10
                                            }
                                            var v118 uint32 = (v117 & (v117 % (((v117 & v117)) | 1)))
                                            _ = v118
                                        }
                                        _ = v117
                                    } else {
                                        fmt.Println(v2)
                                        var v119 float32 = float32((0.0 - v2))
                                        if false {
                                            v10 -= v10
                                            fmt.Println(v2)
                                            v120 := GenS13{F0: (v8 - (v10 ^ (^(v8)))), F1: -128}
                                            _ = v120
                                        }
                                        _ = v119
                                    }
                                    if false {
                                        _ = len(v3)
                                    }
                                    _ = v116
                                }
                            case 369:
                                switch (("foo bar" + ("\n\t" + "foo bar")) + ("hello" + "hello")) {
                                case "\n\t":
                                    var v121 int = fi114
                                    _ = v121
                                default:
                                    v8 = 1
                                }
                            case 9223372036854775807:
                                var v122 int8 = 42
                                _ = v122
                            default:
                                v2 += v2
                            }
                            var v123 bool = (!((!(!((^(v10)) != v10))) || false))
                            switch "\n\t" {
                            case "":
                                var v124 int16 = 912
                                _ = v124
                            case "\n\t":
                                var v125 uint8 = 2
                                if (4294967295 < 357) {
                                    var v126 int16 = 1000
                                    delete(v3, fi114)
                                    _ = v126
                                }
                                v123 = true
                                _ = v125
                            case "foo bar":
                                var v127 string = ""
                                _ = v127
                            default:
                                var v128 int8 = 57
                                _ = v128
                            }
                            _ = v123
                        }
                        fmt.Println(v2, v10, fi11)
                        _ = fi114
                    }
                    _ = fi11
                }
            } else {
                for fi129 := 0; fi129 < 16; fi129++ {
                    switch ("foo bar" + ((("" + "foo bar") + "") + "\n\t")) {
                    case "a":
                        if (42 < 28) {
                            var v130 int = fi129
                            if (42 <= 156) {
                                var v131 uint32 = 2147483647
                                for fi132 := 0; fi132 < 19; fi132++ {
                                    var v133 int = 127
                                    var v134 float32 = -2.25
                                    switch ("" + (("hello" + "hello") + "foo bar")) {
                                    case "\n\t":
                                        v135 := []int{}
                                        _ = v135
                                    default:
                                        switch (("\n\t" + (("a" + "foo bar") + "\n\t")) + "hello") {
                                        case "\n\t":
                                            v136 := GenS7{F0: (fi132 % ((v130) | 1)), F1: "\n\t", F2: 1000}
                                            _ = v136
                                        default:
                                            switch "\n\t" {
                                            case "hello":
                                                switch fi129 {
                                                case 4294967295:
                                                    if true {
                                                        v137 := v1[(-(fi129))]
                                                        var v138 uint = 1
                                                        _ = v137
                                                        _ = v138
                                                    } else {
                                                        fmt.Println(v134, fi132)
                                                        var v139 int16 = 42
                                                        _ = v139
                                                    }
                                                case -759:
                                                    switch "" {
                                                    case "foo bar":
                                                        if (127 != 1) {
                                                            v140 := GenS14{F0: -32768}
                                                            fmt.Println(v2, fi132, v10)
                                                            v141 := GenS15{F0: -1, F1: (("\n\t" + (("" + (("hello" + ((("hello" + ("a" + "")) + "foo bar") + "a")) + "")) + "foo bar")) + (((("\n\t" + "a") + "foo bar") + "\n\t") + "\n\t")), F2: 1}
                                                            v130 = 42
                                                            _ = v140
                                                            _ = v141
                                                        } else {
                                                            var v142 int64 = 9223372036854775807
                                                            v143 := GenS3{F0: v8}
                                                            _ = v142
                                                            _ = v143
                                                        }
                                                    case "a":
                                                        if ((-559 <= -641) && false) {
                                                            if false {
                                                                fmt.Println(v134, fi132)
                                                                var v144 int16 = 291
                                                                _ = v144
                                                            }
                                                            fmt.Println(v10, fi132, v133)
                                                        } else {
                                                            if true {
                                                                v133 += (fi132 & fi132)
                                                                if (!(((float64(0.1) / ((v2 - (1e10 - (v2 * v2))) * float64((v134 * 0.0)))) * ((float64(v134) * v2) - float64(float32(v2)))) == 1e10)) {
                                                                    v134 = 3.14159
                                                                    switch fi129 {
                                                                    case -600:
                                                                        switch (^(v133)) {
                                                                        case 32767:
                                                                            if true {
                                                                                fmt.Println(v8, v134)
                                                                                v2 = float64(1e10)
                                                                            }
                                                                        case 0:
                                                                            fmt.Println(v10)
                                                                        default:
                                                                            var v145 float32 = (v134 / (v134 / (v134 + v134)))
                                                                            _ = v145
                                                                        }
                                                                    default:
                                                                        var v146 string = ""
                                                                        _ = v146
                                                                    }
                                                                    if (((float32((v2 * float64(-2.25))) == v134) || false) && false) {
                                                                        v147 := GenS16{F0: (("hello" + "foo bar") + ((("\n\t" + "hello") + (("" + ("" + "\n\t")) + "foo bar")) + "\n\t")), F1: "", F2: 42}
                                                                        v148 := []int{fi132, fi129, fi129}
                                                                        _ = v147
                                                                        _ = v148
                                                                    }
                                                                }
                                                                var v149 string = "a"
                                                                _ = v149
                                                            } else {
                                                                v150 := GenS12{F0: int8(v8), F1: 42, F2: (v8 ^ 0)}
                                                                v151 := GenS17{F0: 2147483647}
                                                                var v152 uint64 = 1
                                                                _ = v150
                                                                _ = v151
                                                                _ = v152
                                                            }
                                                            var v153 int = v133
                                                            if true {
                                                                var v154 string = ("a" + (("a" + "hello") + "a"))
                                                                var v155 bool = (false || (int16(v8) == 0))
                                                                _ = v154
                                                                _ = v155
                                                            }
                                                            _ = v153
                                                        }
                                                    case "hello":
                                                        if (true && false) {
                                                            v156 := GenS9{F0: 0}
                                                            v130 = (fi129 | 32768)
                                                            var v157 uint64 = 18446744073709551615
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                        v133 = 22
                                                        switch "a" {
                                                        case "\n\t":
                                                            v158 := []int{fi129, 128}
                                                            _ = v158
                                                        case "hello":
                                                            var v159 bool = true
                                                            _ = v159
                                                        case "a":
                                                            fmt.Println(v2)
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    fmt.Println(fi132, v133, v134)
                                                default:
                                                    var v160 int8 = 0
                                                    _ = v160
                                                }
                                            }
                                        }
                                        var v161 int32 = 429
                                        v134 = v134
                                        _ = v161
                                    }
                                    delete(v1, (^(fi132)))
                                    _ = fi132
                                    _ = v133
                                    _ = v134
                                }
                                _ = len(v1)
                                _ = v131
                            } else {
                                var v162 uint = 2
                                v163 := v3[(v130 & fi129)]
                                _ = v162
                                _ = v163
                            }
                            _ = v130
                        }
                    case "foo bar":
                        switch "a" {
                        case "a":
                            if true {
                                var v164 uint64 = 902
                                var v165 uint64 = (v164 & 18446744073709551615)
                                if true {
                                    var v166 int64 = -1
                                    _ = v166
                                } else {
                                    if ("" >= "\n\t") {
                                        if false {
                                            v10 = 42
                                            if true {
                                                fmt.Println(v164, fi129, fi129)
                                                for fi167 := 0; fi167 < 13; fi167++ {
                                                    v168 := GenS18{F0: v164, F1: 42, F2: v8}
                                                    v165 -= v165
                                                    _ = fi167
                                                    _ = v168
                                                }
                                            } else {
                                                delete(v1, (-(fi129)))
                                            }
                                            v165 -= (v165 >> ((0) & 63))
                                        } else {
                                            v169 := GenS19{F0: 605}
                                            v170 := []int{(fi129 | 2147483647), fi129, 557}
                                            var v171 float32 = (-2.25 * 1.5)
                                            _ = v169
                                            _ = v170
                                            _ = v171
                                        }
                                        if false {
                                            v172 := GenS20{F0: 328}
                                            v164 = v165
                                            var v173 int32 = 0
                                            _ = v172
                                            _ = v173
                                        } else {
                                            v10 = v8
                                            var v174 int16 = -1
                                            if (-2147483648 <= -183) {
                                                v175 := []int{(fi129 >> ((374) & 63))}
                                                var v176 int = 532
                                                _ = v175
                                                _ = v176
                                            } else {
                                                if (2147483647 > 0) {
                                                    v174 += (v174 + v174)
                                                    var v177 bool = (true || (int64(v174) > 9223372036854775807))
                                                    if false {
                                                        var v178 int64 = -1
                                                        _ = v178
                                                    }
                                                    _ = v177
                                                }
                                                var v179 int64 = 298
                                                fmt.Println(v10)
                                                _ = v179
                                            }
                                            _ = v174
                                        }
                                        v180 := v1[(fi129 / ((0) | 1))]
                                        _ = v180
                                    }
                                    v181 := []int{fi129}
                                    _ = v181
                                }
                                _ = v164
                                _ = v165
                            } else {
                                v182 := []int{fi129, (fi129 | (fi129 / ((-2147483648) | 1))), fi129, fi129}
                                _ = v182
                            }
                        case "":
                            v183 := GenS14{F0: -32768}
                            _ = v183
                        }
                    default:
                        v8 = 764
                    }
                    if (true && (!(false || (!true)))) {
                        var v184 uint64 = 42
                        var v185 int8 = -1
                        _ = v184
                        _ = v185
                    } else {
                        if false {
                            var v186 float32 = 3.14159
                            _ = v186
                        }
                        var v187 int64 = -1
                        var v188 float32 = 0.0
                        _ = v187
                        _ = v188
                    }
                    _ = fi129
                }
            }
            v8 += (v10 * 2)
        }
        v8 -= 348
    } else {
        v8 = 210
    }
    if true {
        fmt.Println(v2, v2)
        switch 2147483647 {
        case 32767:
            fmt.Println(v10, v2)
        default:
            var v189 int16 = -32768
            _ = v189
        }
        var v190 uint = 2
        _ = v190
    }
    v2 -= v2
    if false {
        v10 -= v10
        v2 = v2
        v2 = (((((v2 - float64(((3.14159 - 3.14159) + -2.25))) - (v2 + (((v2 + ((1e10 / (v2 * (float64((0.1 - 0.1)) * (3.14159 - 1e10)))) - v2)) / v2) + float64(float32((v2 + 0.0)))))) + v2) / v2) / float64(((0.0 - -2.25) / 0.1)))
    }
    v191 := GenS21{F0: (true || (v8 <= v10)), F1: 127}
    var v192 uint16 = 32767
    v8 += (v10 * 65535)
    switch "hello" {
    case "a":
        switch "\n\t" {
        case "a":
            switch "\n\t" {
            case "\n\t":
                if true {
                    if (false && (true || true)) {
                        v1[-985] = 2147483648
                        if true {
                            switch "a" {
                            case "hello":
                                switch ("a" + "") {
                                case "foo bar":
                                    if true {
                                        var v193 int8 = 0
                                        switch 32767 {
                                        case 1:
                                            if true {
                                                v2 += float64(3.14159)
                                                var v194 int = 127
                                                _ = v194
                                            }
                                        case 4294967295:
                                            var v195 int64 = 0
                                            _ = v195
                                        case 0:
                                            var v196 float32 = float32(v2)
                                            _ = v196
                                        default:
                                            var v197 uint64 = 42
                                            _ = v197
                                        }
                                        v198 := GenS15{F0: 24, F1: ("a" + "foo bar"), F2: v193}
                                        _ = v193
                                        _ = v198
                                    }
                                case "":
                                    for fi199 := 0; fi199 < 20; fi199++ {
                                        for fi200 := 0; fi200 < 18; fi200++ {
                                            v201 := []int{(^(fi200)), -464}
                                            _ = fi200
                                            _ = v201
                                        }
                                        v202 := GenS9{F0: 0}
                                        _ = fi199
                                        _ = v202
                                    }
                                case "hello":
                                    for fi203 := 0; fi203 < 13; fi203++ {
                                        v204 := GenS18{F0: 29, F1: 255, F2: v10}
                                        if (uint(v192) < 4294967295) {
                                            v205 := []int{(^(fi203))}
                                            _ = v205
                                        }
                                        if (v2 < (1.5 + 3.14159)) {
                                            v2 -= (1.5 - v2)
                                        } else {
                                            switch 128 {
                                            case 127:
                                                fmt.Println(v8)
                                            case -359:
                                                var v206 int8 = 0
                                                _ = v206
                                            case 2147483647:
                                                fmt.Println(v10, v10)
                                            default:
                                                if (878 < 2) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if (v192 >= (v8 * (v8 << ((227) & 63)))) {
                                                            v2 = 0.1
                                                            v8 += 0
                                                            for fi207 := 0; fi207 < 14; fi207++ {
                                                                if (1 > 18446744073709551615) {
                                                                    if false {
                                                                        if false {
                                                                            var v208 uint16 = 65535
                                                                            v192 = 43
                                                                            switch ("" + (("hello" + "") + (("a" + "a") + ("foo bar" + "\n\t")))) {
                                                                            case "\n\t":
                                                                                v8 -= (v192 - 436)
                                                                            }
                                                                            var v209 uint8 = 173
                                                                            var v210 uint64 = 4294967295
                                                                            _ = v208
                                                                            _ = v209
                                                                            _ = v210
                                                                        } else {
                                                                            var v211 uint64 = 765
                                                                            _ = v211
                                                                        }
                                                                        fmt.Println(v2, v10, v10)
                                                                    }
                                                                    if true {
                                                                        if true {
                                                                            v10 += v8
                                                                            var v212 int8 = -64
                                                                            _ = v212
                                                                        }
                                                                    } else {
                                                                        v2 += 1e10
                                                                        switch "hello" {
                                                                        case "":
                                                                            if true {
                                                                                v8 = 970
                                                                                var v213 uint64 = 2
                                                                                _ = v213
                                                                            }
                                                                        case "a":
                                                                            var v214 bool = true
                                                                            _ = v214
                                                                        default:
                                                                            var v215 int = (fi203 & (fi203 / ((fi203) | 1)))
                                                                            _ = v215
                                                                        }
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            fmt.Println(v192, v8, v2)
                                                                        case "hello":
                                                                            var v216 uint16 = (v10 / (((v8 % (((v192 ^ (v10 % (((v8 << ((255) & 63))) | 1)))) | 1))) | 1))
                                                                            _ = v216
                                                                        default:
                                                                            _ = len(v3)
                                                                        }
                                                                        v217 := []int{(fi203 | 128), fi203}
                                                                        _ = v217
                                                                    }
                                                                    v2 = (v2 - 1.5)
                                                                } else {
                                                                    var v218 int16 = 464
                                                                    _ = v218
                                                                }
                                                                switch "hello" {
                                                                case "":
                                                                    v2 = v2
                                                                default:
                                                                    fmt.Println(v10, fi203, v2)
                                                                }
                                                                v219 := GenS15{F0: 1, F1: "", F2: -128}
                                                                if true {
                                                                    fmt.Println(v8, v8)
                                                                    v220 := GenS20{F0: 0}
                                                                    _ = v220
                                                                }
                                                                _ = fi207
                                                                _ = v219
                                                            }
                                                        } else {
                                                            var v221 int = fi203
                                                            var v222 uint64 = 0
                                                            _ = v221
                                                            _ = v222
                                                        }
                                                    case "a":
                                                        v223 := GenS11{F0: 42}
                                                        _ = v223
                                                    case "":
                                                        var v224 int16 = 215
                                                        _ = v224
                                                    }
                                                    v225 := GenS10{F0: 42, F1: 65535}
                                                    _ = v225
                                                } else {
                                                    v10 += (v192 - 65535)
                                                    var v226 bool = true
                                                    var v227 int = 883
                                                    _ = v226
                                                    _ = v227
                                                }
                                            }
                                        }
                                        _ = fi203
                                        _ = v204
                                    }
                                default:
                                    var v228 uint8 = 1
                                    _ = v228
                                }
                            }
                            v192 -= (v8 % ((v192) | 1))
                        }
                    } else {
                        if true {
                            v229 := GenS22{F0: 4294967295, F1: "hello", F2: 42}
                            _ = v229
                        }
                    }
                } else {
                    var v230 int16 = 1
                    _ = v230
                }
            default:
                v231 := GenS17{F0: 42}
                _ = v231
            }
        default:
            for fi232 := 0; fi232 < 3; fi232++ {
                v192 = 42
                fmt.Println(v2, v10)
                _ = fi232
            }
        }
    case "foo bar":
        v233 := []int{-2147483648, 363, 127}
        _ = v233
    default:
        v8 = (^(v192))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v192:", v192)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v191
    _ = v192
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
