package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int32; F1 string; F2 int16 }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 int8; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS10 struct { F0 uint16; F1 int16; F2 int8 }
type GenS11 struct { F0 int16; F1 bool; F2 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 int }
type GenS14 struct { F0 bool; F1 uint16; F2 bool }
type GenS15 struct { F0 int32; F1 int16; F2 uint8 }
type GenS16 struct { F0 int64 }
func genFn0(p0 bool, p1 bool, p2 int) uint {
    switch (p2 + 711) {
    case 9223372036854775807:
        var v4 int8 = 42
        _ = v4
    case 42:
        if (!((((0.0 + 1e10) / 1.0) == (3.14159 * 0.1)) || p0)) {
            if true {
                var v5 int32 = 537
                _ = v5
            }
        } else {
            p1 = p1
            if ((3.14159 * ((3.14159 * 1.5) + ((-2.25 + 1e10) / 1e10))) < 3.14159) {
                p2 += p2
            }
            v6 := GenS1{F0: p2, F1: (650 < 33)}
            _ = v6
        }
    case 132:
        fmt.Println(p0, p1)
    default:
        switch (p2 % (((p2 % ((32768) | 1))) | 1)) {
        case 32768:
            var v7 float64 = 0.1
            _ = v7
        case 2147483647:
            v8 := []int{p2, p2, p2, 42}
            _ = v8
        }
    }
    v9 := []int{p2, (p2 & p2), p2, -2147483648}
    v10 := GenS2{F0: false, F1: 65535, F2: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 uint = 453
    var v2 float64 = (0.1 + float64(1e10))
    var v3 int = 538
    var v11 bool = false
    fmt.Println(v11, v3)
    var v12 uint16 = 32767
    switch "foo bar" {
    case "hello":
        fmt.Println(v3)
    case "foo bar":
        var v13 string = ("a" + "a")
        _ = v13
    default:
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(v11, v12, v3)
            for fi15 := 0; fi15 < 16; fi15++ {
                v1 -= (v1 ^ 0)
                v16 := []int{128, fi15, fi15, (fi14 - -2147483648)}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    switch "\n\t" {
    case "":
        fmt.Println(v12)
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS3{F0: 946}
        var v19 uint8 = 128
        _ = fi17
        _ = v18
        _ = v19
    }
    v11 = false
    v20 := GenS4{F0: v11}
    for fi21 := 0; fi21 < 8; fi21++ {
        switch (fi21 / (((v3 >> ((genFn0(v11, v11, fi21)) & 63))) | 1)) {
        case 127:
            v3 += (fi21 * fi21)
        }
        _ = fi21
    }
    var v22 string = ("hello" + (("" + "hello") + "a"))
    switch v22 {
    case "\n\t":
        v23 := []int{(v3 >> (((v1 / ((2) | 1))) & 63)), 475}
        _ = v23
    case "a":
        if v11 {
            fmt.Println(v2, v22, v22)
            for fi24 := 0; fi24 < 16; fi24++ {
                switch v22 {
                case "\n\t":
                    if v11 {
                        fmt.Println(v12, v1)
                        v3 -= fi24
                    }
                case "":
                    if ((42 <= -1) || (12 != 2)) {
                        if (!v11) {
                            v12 = (^(v12))
                            switch v22 {
                            case "a":
                                v2 = v2
                            case "hello":
                                fmt.Println(v1, fi24)
                            case "":
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v2 = (3.14159 - v2)
                                    v26 := GenS4{F0: (-781 == 0)}
                                    if v11 {
                                        v27 := []int{-1, (fi24 >> ((genFn0((42 <= 32768), false, 42)) & 63))}
                                        v28 := GenS4{F0: (v12 > (v12 % (((v12 >> (((v1 & (v1 ^ v1))) & 63))) | 1)))}
                                        var v29 float64 = (1.5 / v2)
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            var v30 bool = false
                            _ = v30
                        } else {
                            fmt.Println(fi24, v12, v22)
                        }
                        v12 -= v12
                        if (536 >= 18446744073709551615) {
                            v3 -= fi24
                            v11 = (true && (1 <= 42))
                            if false {
                                switch (v22 + v22) {
                                case "hello":
                                    if (!(v11 && false)) {
                                        v22 = ("" + ("" + ("\n\t" + (v22 + ("hello" + v22)))))
                                        v31 := GenS4{F0: v11}
                                        _ = v31
                                    } else {
                                        var v32 int64 = 0
                                        v11 = v11
                                        if (1 > (^(v32))) {
                                            v2 -= ((1.5 - float64(0.1)) * (v2 - ((-2.25 / v2) / (((v2 - ((float64(0.1) - ((v2 * v2) / float64(0.1))) + (v2 / (float64(0.0) * (v2 * float64(0.1)))))) * (3.14159 + ((((((v2 - v2) - 1.5) * v2) * 1.5) - float64(float32(-2.25))) / (v2 * v2)))) + float64((((((0.1 - 0.0) - 0.0) + float32(v2)) * float32(((v2 + v2) + float64(3.14159)))) + (1e10 * 3.14159)))))))
                                            v33 := GenS5{F0: (v3 & v3), F1: 574, F2: 42}
                                            _ = v33
                                        }
                                        _ = v32
                                    }
                                default:
                                    for fi34 := 0; fi34 < 11; fi34++ {
                                        v12 -= (^(v12))
                                        var v35 int32 = 42
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                                fmt.Println(v1, v3)
                                v12 = v12
                            }
                        }
                    }
                case "a":
                    v36 := []int{-1, 2147483647, (v3 - fi24)}
                    _ = v36
                default:
                    v1 += v1
                }
                var v37 uint32 = 139
                v38 := GenS6{F0: -1, F1: v22, F2: 1000}
                _ = fi24
                _ = v37
                _ = v38
            }
            v22 = "hello"
        }
    case "":
        v39 := []int{v3, (v3 * (v3 - (-(v3)))), v3}
        _ = v39
    }
    switch v3 {
    case -125:
        switch ((v22 + v22) + ("foo bar" + (v22 + v22))) {
        case "foo bar":
            v40 := []int{v3, 424, (v3 - v3), 128}
            _ = v40
        default:
            switch "" {
            case "foo bar":
                if (806 != 65535) {
                    v1 = (v1 - (v1 - 42))
                    v1 += genFn0(v11, (v1 > v1), v3)
                } else {
                    if (v3 < v3) {
                        switch v22 {
                        case "":
                            for fi41 := 0; fi41 < 10; fi41++ {
                                var v42 bool = v11
                                _ = fi41
                                _ = v42
                            }
                        case "\n\t":
                            switch (-(v3)) {
                            case 7:
                                if (!(!(-2.25 == (v2 + v2)))) {
                                    var v43 int64 = 0
                                    _ = v43
                                }
                            case 32768:
                                var v44 uint8 = 42
                                _ = v44
                            case -9223372036854775808:
                                v45 := []int{(v3 << (((v1 | v1)) & 63)), 775}
                                _ = v45
                            default:
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    v47 := GenS1{F0: int(v12), F1: (457 >= 1000)}
                                    if (-1 < 0) {
                                        if true {
                                            var v48 float64 = v2
                                            _ = v48
                                        } else {
                                            if v11 {
                                                fmt.Println(v1, fi46, v3)
                                                v49 := GenS7{F0: uint32(v12), F1: -1}
                                                v50 := GenS8{F0: -128, F1: 2147483648, F2: 111}
                                                _ = v49
                                                _ = v50
                                            } else {
                                                for fi51 := 0; fi51 < 14; fi51++ {
                                                    var v52 float32 = ((1e10 / ((float32(v2) * 1e10) + 3.14159)) * 3.14159)
                                                    var v53 string = "hello"
                                                    _ = fi51
                                                    _ = v52
                                                    _ = v53
                                                }
                                                switch v3 {
                                                case 128:
                                                    v54 := GenS9{F0: 255, F1: -32768, F2: 30}
                                                    _ = v54
                                                case 32768:
                                                    if v11 {
                                                        v11 = v11
                                                        v11 = (-2147483648 >= 183)
                                                    } else {
                                                        for fi55 := 0; fi55 < 16; fi55++ {
                                                            v56 := []int{9223372036854775807, fi55, (v3 >> ((uint(v12)) & 63)), v3}
                                                            _ = fi55
                                                            _ = v56
                                                        }
                                                        var v57 int32 = 923
                                                        v22 = "foo bar"
                                                        _ = v57
                                                    }
                                                case 32767:
                                                    var v58 string = (v22 + v22)
                                                    _ = v58
                                                }
                                                if v11 {
                                                    switch 32768 {
                                                    case -397:
                                                        switch v22 {
                                                        case "":
                                                            var v59 int = (-(fi46))
                                                            _ = v59
                                                        }
                                                    case -968:
                                                        v22 = (v22 + (v22 + v22))
                                                    case 32767:
                                                        switch v22 {
                                                        case "a":
                                                            v3 = 32768
                                                        }
                                                    default:
                                                        v60 := GenS10{F0: 955, F1: -485, F2: -80}
                                                        _ = v60
                                                    }
                                                    for fi61 := 0; fi61 < 20; fi61++ {
                                                        fmt.Println(v12, v12)
                                                        if v11 {
                                                            if v11 {
                                                                v11 = (-2.25 == ((float64(3.14159) / v2) - ((3.14159 - float64((float32(1e10) - (float32(v2) * 0.1)))) - (3.14159 + v2))))
                                                            } else {
                                                                v2 = 1.5
                                                                var v62 uint16 = v12
                                                                if v11 {
                                                                    if (!true) {
                                                                        v3 += fi46
                                                                        var v63 uint8 = 213
                                                                        v64 := GenS1{F0: (fi46 ^ fi46), F1: (v12 > (v62 % ((v12) | 1)))}
                                                                        _ = v63
                                                                        _ = v64
                                                                    } else {
                                                                        v3 = (fi46 ^ (v3 >> ((337) & 63)))
                                                                    }
                                                                    var v65 float64 = v2
                                                                    var v66 uint16 = (^(v12))
                                                                    _ = v65
                                                                    _ = v66
                                                                } else {
                                                                    if (!true) {
                                                                        fmt.Println(v3, fi46)
                                                                        v67 := GenS11{F0: -558, F1: v11, F2: 326}
                                                                        var v68 uint32 = 4294967295
                                                                        _ = v67
                                                                        _ = v68
                                                                    } else {
                                                                        var v69 int32 = 0
                                                                        v11 = (42 >= -32768)
                                                                        _ = v69
                                                                    }
                                                                    v70 := GenS3{F0: 2147483647}
                                                                    v71 := []int{127, 957}
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                _ = v62
                                                            }
                                                            fmt.Println(v11, v2)
                                                        }
                                                        switch v22 {
                                                        case "foo bar":
                                                            v1 = 2
                                                        default:
                                                            switch v22 {
                                                            case "foo bar":
                                                                var v72 float64 = v2
                                                                _ = v72
                                                            default:
                                                                v11 = true
                                                            }
                                                            if true {
                                                                fmt.Println(v11, v2, fi46)
                                                                var v73 uint8 = uint8(v1)
                                                                var v74 uint64 = 18446744073709551615
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                v75 := GenS12{F0: 0}
                                                                _ = v75
                                                            }
                                                            var v76 float32 = 1.5
                                                            _ = v76
                                                        }
                                                        v22 = v22
                                                        _ = fi61
                                                    }
                                                }
                                            }
                                            v77 := []int{v3, fi46, (v3 >> (((v1 % (((^(v1))) | 1))) & 63))}
                                            fmt.Println(v3, v11)
                                            _ = v77
                                        }
                                        var v78 int16 = 1000
                                        _ = v78
                                    } else {
                                        fmt.Println(fi46, v11, v11)
                                        switch v22 {
                                        case "":
                                            if (-408 >= 1) {
                                                var v79 int8 = 127
                                                v80 := []int{127, (v3 << ((genFn0(true, true, v3)) & 63))}
                                                v11 = (v79 == v79)
                                                _ = v79
                                                _ = v80
                                            }
                                        case "a":
                                            switch v22 {
                                            case "foo bar":
                                                switch v22 {
                                                case "a":
                                                    if true {
                                                        if ((fi46 % (((-(v3))) | 1)) == v3) {
                                                            v3 -= (^(fi46))
                                                        }
                                                        v81 := GenS12{F0: 2147483647}
                                                        _ = v81
                                                    } else {
                                                        v11 = (!(!v11))
                                                    }
                                                default:
                                                    v3 = -840
                                                }
                                            default:
                                                var v82 uint32 = 65535
                                                _ = v82
                                            }
                                        case "\n\t":
                                            if (!(!(((0 <= (v1 >> ((v1) & 63))) || (0 <= 902)) && false))) {
                                                v12 = 1
                                            } else {
                                                switch 0 {
                                                case 937:
                                                    var v83 uint16 = v12
                                                    _ = v83
                                                }
                                            }
                                        default:
                                            v84 := GenS13{F0: -301, F1: (v3 * v3)}
                                            _ = v84
                                        }
                                        switch (v3 - (^(fi46))) {
                                        case 2147483647:
                                            if v11 {
                                                v85 := GenS14{F0: (true && v11), F1: v12, F2: v11}
                                                _ = v85
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    v86 := GenS9{F0: 1, F1: -32768, F2: 0}
                                                    _ = v86
                                                default:
                                                    v87 := GenS4{F0: v11}
                                                    _ = v87
                                                }
                                            }
                                        case 895:
                                            v88 := GenS11{F0: 42, F1: (((1e10 + (1.5 + float32(-2.25))) / ((((1e10 + 0.1) + -2.25) * 1e10) - 0.0)) <= 0.1), F2: fi46}
                                            _ = v88
                                        }
                                    }
                                    if v11 {
                                        v89 := []int{(fi46 / (((fi46 & -2147483648)) | 1)), 127, (-(v3)), (v3 - -494)}
                                        _ = v89
                                    }
                                    _ = fi46
                                    _ = v47
                                }
                            }
                        default:
                            v12 += 177
                        }
                        for fi90 := 0; fi90 < 18; fi90++ {
                            var v91 int16 = -32768
                            if (int32(v91) == 1) {
                                if false {
                                    switch (-(v3)) {
                                    case 4294967295:
                                        if v11 {
                                            if (v12 > (v12 % ((v12) | 1))) {
                                                var v92 uint8 = uint8(v12)
                                                v12 = v12
                                                _ = v92
                                            } else {
                                                for fi93 := 0; fi93 < 15; fi93++ {
                                                    v94 := GenS15{F0: int32(v3), F1: v91, F2: 0}
                                                    fmt.Println(fi93, v2, fi93)
                                                    if false {
                                                        v11 = (false || true)
                                                    }
                                                    _ = fi93
                                                    _ = v94
                                                }
                                            }
                                            var v95 int16 = v91
                                            v12 -= v12
                                            _ = v95
                                        } else {
                                            if (v3 == (fi90 / ((v3) | 1))) {
                                                if ("\n\t" < ((v22 + ("a" + v22)) + v22)) {
                                                    v12 = v12
                                                    switch v22 {
                                                    case "foo bar":
                                                        for fi96 := 0; fi96 < 15; fi96++ {
                                                            var v97 int32 = 0
                                                            _ = fi96
                                                            _ = v97
                                                        }
                                                    case "":
                                                        v2 = 3.14159
                                                    }
                                                } else {
                                                    v98 := []int{849, -9223372036854775808, (v3 | (v3 << ((v1) & 63)))}
                                                    _ = v98
                                                }
                                                v12 = (v12 / ((32767) | 1))
                                            } else {
                                                for fi99 := 0; fi99 < 11; fi99++ {
                                                    v100 := GenS5{F0: (fi90 + 127), F1: 2147483648, F2: 128}
                                                    _ = fi99
                                                    _ = v100
                                                }
                                            }
                                        }
                                    default:
                                        v3 += (-(v3))
                                    }
                                    fmt.Println(v91, v2, v91)
                                    for fi101 := 0; fi101 < 5; fi101++ {
                                        v2 = (v2 / (v2 - (float64(1e10) * 1.5)))
                                        _ = fi101
                                    }
                                } else {
                                    var v102 int16 = 0
                                    var v103 string = ("a" + ("foo bar" + ("\n\t" + v22)))
                                    _ = v102
                                    _ = v103
                                }
                                v104 := GenS16{F0: -9223372036854775808}
                                _ = v104
                            } else {
                                if (-32768 >= 42) {
                                    if v11 {
                                        v3 = (fi90 << (((v1 << (((v1 >> ((v1) & 63))) & 63))) & 63))
                                    }
                                    v11 = v11
                                    var v105 uint = 936
                                    _ = v105
                                }
                            }
                            _ = fi90
                            _ = v91
                        }
                        var v106 uint64 = 18446744073709551615
                        _ = v106
                    }
                    var v107 uint = v1
                    _ = v107
                }
            default:
                switch v22 {
                case "":
                    for fi108 := 0; fi108 < 7; fi108++ {
                        v109 := GenS6{F0: 42, F1: v22, F2: 0}
                        _ = fi108
                        _ = v109
                    }
                case "foo bar":
                    for fi110 := 0; fi110 < 3; fi110++ {
                        var v111 bool = v11
                        var v112 float32 = 1e10
                        fmt.Println(v112, v22, v2)
                        _ = fi110
                        _ = v111
                        _ = v112
                    }
                default:
                    var v113 float32 = 3.14159
                    _ = v113
                }
            }
        }
    case -172:
        v114 := GenS2{F0: false, F1: v1, F2: (v22 + ((v22 + "\n\t") + v22))}
        _ = v114
    default:
        var v115 int16 = 42
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
