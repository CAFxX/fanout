package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64; F1 int8; F2 string }
type GenS9 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS10 struct { F0 bool; F1 bool; F2 uint }
type GenS11 struct { F0 string; F1 int32; F2 int32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 bool; F1 uint8 }
type GenS15 struct { F0 int32; F1 uint16 }
type GenS16 struct { F0 int; F1 string; F2 uint16 }
type GenS17 struct { F0 bool; F1 uint }
type GenS18 struct { F0 uint64; F1 uint64; F2 uint32 }
func genFn0(p0 int8, p1 int32, p2 string) int8 {
    var v6 uint = 65535
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 + (-(p0)))
}

func main() {
    var v1 string = ""
    var v2 float32 = 0.1
    var v3 uint16 = 125
    var v4 int64 = -727
    var v5 int32 = 0
    v5 -= v5
    v2 += ((((0.1 + v2) / (0.1 * (((v2 + v2) - (((float32(float64(-2.25)) + float32(0.0)) * (v2 * v2)) + v2)) - (((-2.25 - v2) + (v2 / v2)) - (((((v2 / 3.14159) / (v2 - ((0.0 - v2) + v2))) * 0.0) + (v2 - (v2 - 1e10))) / v2))))) * float32(-2.25)) * v2)
    fmt.Println(v1, v1)
    v8 := GenS1{F0: 188}
    v5 += v5
    v2 = v2
    v1 = v1
    v9 := []int{-2147483648, 127, 32767, 4294967295}
    if (42 == 0) {
        var v10 int8 = genFn0(genFn0(0, -98, v1), v5, (v1 + (v1 + v1)))
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        switch v1 {
        case "\n\t":
            if ((float64(v2) < 0.0) || (v2 >= (v2 + v2))) {
                if (float32(1.5) == float32(float64((v2 * v2)))) {
                    v2 += 1.5
                } else {
                    switch 1 {
                    case 518:
                        v9 = append(v9, 2147483647)
                    case 32767:
                        v5 = 1
                    }
                }
                fmt.Println(v4, v5, v5)
            }
        }
        _ = v10
    }
    if false {
        v1 = "foo bar"
    }
    v11 := GenS2{F0: v4}
    v9 = append(v9, 963)
    switch v1 {
    case "a":
        switch (v1 + (("hello" + ("foo bar" + v1)) + ("hello" + (v1 + v1)))) {
        case "":
            v3 -= (v3 | 472)
        }
    case "hello":
        v9 = append(v9, 855)
    default:
        v12 := GenS2{F0: (^(v4))}
        _ = v12
    }
    v2 = (-2.25 / v2)
    if (!false) {
        var v13 uint = 785
        _ = v13
    } else {
        var v14 int8 = 100
        if (v14 > 0) {
            v14 -= (v14 >> ((477) & 63))
            _ = len(v9)
            switch "foo bar" {
            case "foo bar":
                var v15 uint16 = v3
                _ = v15
            }
        } else {
            v16 := GenS1{F0: 699}
            if len(v9) > 0 {
                v9[0] = 42
            }
            fmt.Println(v5, v3, v3)
            _ = v16
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v14
    }
    if ((false && (1e10 <= ((float64(v2) - 0.1) + ((0.1 / (((1.5 / 3.14159) * 0.1) - (1e10 - (3.14159 * 1e10)))) - 3.14159)))) || (v3 >= v3)) {
        for fi17 := 0; fi17 < 15; fi17++ {
            if false {
                v5 += (v5 / ((-2147483648) | 1))
            } else {
                fmt.Println(v1, v4, v5)
                if true {
                    v4 -= -1
                    v18 := GenS3{F0: (v5 - (v5 - (-(v5))))}
                    var v19 uint8 = 247
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint32 = 65535
                    if (fi17 != -2147483648) {
                        var v21 uint8 = uint8(v20)
                        switch -9223372036854775808 {
                        case 1:
                            var v22 uint8 = (v21 % (((v21 % ((v21) | 1))) | 1))
                            _ = v22
                        }
                        _ = v21
                    }
                    v23 := GenS4{F0: "foo bar"}
                    _ = v20
                    _ = v23
                }
                var v24 uint8 = 2
                _ = v24
            }
            var v25 uint32 = 0
            _ = fi17
            _ = v25
        }
    } else {
        if (true && false) {
            for fi26 := 0; fi26 < 15; fi26++ {
                switch -1 {
                case 4294967295:
                    if len(v9) > 0 {
                        v9[0] = (-(fi26))
                    }
                case 127:
                    fmt.Println(v3, v3, v2)
                default:
                    var v27 float32 = (3.14159 / (0.0 / (v2 * float32(3.14159))))
                    _ = v27
                }
                _ = fi26
            }
        } else {
            var v28 int32 = (-(v5))
            if false {
                v1 = "a"
                v9 = append(v9, -660)
            } else {
                var v29 uint8 = 128
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v30 uint8 = v29
                _ = v29
                _ = v30
            }
            _ = v28
        }
        if (931 >= 4294967295) {
            _ = len(v9)
            v1 = "hello"
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            if (false || ((127 >= 42) && (true && (0 != 65535)))) {
                if true {
                    v32 := GenS5{F0: 18446744073709551615}
                    switch fi31 {
                    case 2147483647:
                        for fi33 := 0; fi33 < 8; fi33++ {
                            v34 := GenS3{F0: (-(v5))}
                            v35 := GenS5{F0: 952}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    case 613:
                        var v36 int32 = v5
                        _ = v36
                    case -166:
                        if false {
                            if (!true) {
                                for fi37 := 0; fi37 < 14; fi37++ {
                                    var v38 uint64 = 0
                                    fmt.Println(fi31, v3, v3)
                                    switch (v1 + v1) {
                                    case "hello":
                                        _ = len(v9)
                                    case "\n\t":
                                        fmt.Println(fi31, v2, fi37)
                                    case "":
                                        if (v3 <= 875) {
                                            var v39 uint = 4294967295
                                            v4 = (v4 ^ 9223372036854775807)
                                            _ = v39
                                        }
                                    default:
                                        if (false && true) {
                                            if false {
                                                v3 += v3
                                            }
                                        } else {
                                            if true {
                                                if len(v9) > 0 {
                                                    v9[0] = (^(fi37))
                                                }
                                                switch ((("a" + "a") + v1) + (v1 + "foo bar")) {
                                                case "":
                                                    v40 := GenS6{F0: (v5 & 2147483647)}
                                                    _ = v40
                                                default:
                                                    var v41 string = (("" + v1) + ((v1 + v1) + v1))
                                                    _ = v41
                                                }
                                            } else {
                                                var v42 int = fi31
                                                _ = v42
                                            }
                                            v9 = append(v9, fi31)
                                            v43 := GenS3{F0: v5}
                                            v38 = (v38 - (v38 - v38))
                                            _ = v43
                                        }
                                        v38 = (v38 - 0)
                                        v3 = v3
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    var v45 uint = 153
                                    if (!true) {
                                        switch "foo bar" {
                                        case "hello":
                                            v45 -= 4294967295
                                        case "foo bar":
                                            v3 = (v3 + v3)
                                        case "a":
                                            var v46 int8 = 42
                                            _ = v46
                                        }
                                        fmt.Println(v3)
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                if (!(true || (true || (v5 == v5)))) {
                                    for fi47 := 0; fi47 < 3; fi47++ {
                                        fmt.Println(fi31)
                                        v48 := GenS7{F0: 127}
                                        _ = fi47
                                        _ = v48
                                    }
                                    v5 = v5
                                    v2 = v2
                                }
                            } else {
                                v4 = (v4 >> ((2) & 63))
                                switch fi31 {
                                case 1:
                                    switch (v1 + ((v1 + v1) + v1)) {
                                    case "":
                                        if len(v9) > 0 {
                                            v9[0] = fi31
                                        }
                                    }
                                default:
                                    v49 := GenS3{F0: (-(v5))}
                                    _ = v49
                                }
                                switch ("" + (v1 + v1)) {
                                case "\n\t":
                                    v3 -= v3
                                default:
                                    fmt.Println(v1, v2, v4)
                                }
                            }
                            switch (fi31 & fi31) {
                            case -574:
                                for fi50 := 0; fi50 < 15; fi50++ {
                                    var v51 string = v1
                                    fmt.Println(fi50, v2)
                                    if true {
                                        var v52 uint64 = 4294967295
                                        fmt.Println(v5, v5)
                                        _ = v52
                                    } else {
                                        if false {
                                            v5 = (v5 & (v5 / (((v5 * v5)) | 1)))
                                        }
                                        if true {
                                            v1 = "foo bar"
                                            if ((!false) || (32767 >= 32767)) {
                                                if len(v9) > 0 {
                                                    v9[0] = (fi31 + 9223372036854775807)
                                                }
                                            }
                                            if (!(v1 >= v1)) {
                                                if (!(42 >= 18446744073709551615)) {
                                                    v5 -= (v5 - (v5 - (v5 & (v5 + (v5 + (v5 / ((v5) | 1)))))))
                                                    switch fi31 {
                                                    case -9223372036854775808:
                                                        if false {
                                                            fmt.Println(v51)
                                                        }
                                                    case 0:
                                                        if (203 != 0) {
                                                            if false {
                                                                if (1 > 1000) {
                                                                    v9 = append(v9, int(v4))
                                                                } else {
                                                                    fmt.Println(fi50)
                                                                    var v53 int64 = (-(v4))
                                                                    var v54 int16 = 727
                                                                    if (false || ((3.14159 * (3.14159 - 1e10)) != -2.25)) {
                                                                        v55 := GenS1{F0: 207}
                                                                        fmt.Println(v3, v53)
                                                                        _ = v55
                                                                    }
                                                                    var v56 int = fi31
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v56
                                                                }
                                                                var v57 float32 = v2
                                                                _ = v57
                                                            }
                                                            v4 = v4
                                                            v3 -= (v3 % ((2) | 1))
                                                        } else {
                                                            v58 := GenS2{F0: -1}
                                                            v59 := GenS4{F0: v51}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    case 819:
                                                        v60 := GenS7{F0: 2}
                                                        _ = v60
                                                    }
                                                } else {
                                                    v5 = (v5 ^ -2147483648)
                                                    _ = len(v9)
                                                }
                                                v1 = v51
                                            } else {
                                                if len(v9) > 0 {
                                                    v9[0] = 221
                                                }
                                                if len(v9) > 0 {
                                                    v9[0] = fi31
                                                }
                                                var v61 string = (("a" + (("\n\t" + "\n\t") + "hello")) + v51)
                                                _ = v61
                                            }
                                            var v62 uint32 = 4294967295
                                            _ = v62
                                        }
                                        v51 = v1
                                        var v63 uint16 = (v3 << ((255) & 63))
                                        if len(v9) > 0 {
                                            v9[0] = 42
                                        }
                                        _ = v63
                                    }
                                    _ = fi50
                                    _ = v51
                                }
                            case 4294967295:
                                if (!false) {
                                    var v64 int8 = 1
                                    fmt.Println(v5, v3)
                                    var v65 uint16 = v3
                                    _ = v64
                                    _ = v65
                                } else {
                                    v9 = append(v9, 42)
                                    if len(v9) > 0 {
                                        v9[0] = fi31
                                    }
                                }
                            case 0:
                                for fi66 := 0; fi66 < 10; fi66++ {
                                    switch v1 {
                                    case "":
                                        if (!true) {
                                            if (42 == 2) {
                                                fmt.Println(v4)
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    var v67 uint16 = 2
                                                    _ = v67
                                                case "\n\t":
                                                    switch (v1 + (v1 + v1)) {
                                                    case "foo bar":
                                                        v3 = (v3 * 65535)
                                                    default:
                                                        var v68 int16 = 1
                                                        _ = v68
                                                    }
                                                case "hello":
                                                    switch "a" {
                                                    case "\n\t":
                                                        if true {
                                                            v69 := GenS3{F0: (v5 | v5)}
                                                            _ = v69
                                                        }
                                                    case "a":
                                                        switch (fi66 >> ((18446744073709551615) & 63)) {
                                                        case 588:
                                                            v4 = -1
                                                        case 0:
                                                            switch "hello" {
                                                            case "hello":
                                                                var v70 int8 = genFn0(0, (v5 % (((v5 * v5)) | 1)), (v1 + "a"))
                                                                _ = v70
                                                            case "\n\t":
                                                                v2 -= v2
                                                            case "":
                                                                switch (v1 + (v1 + v1)) {
                                                                case "a":
                                                                    fmt.Println(v3, v1)
                                                                case "foo bar":
                                                                    v2 -= (0.0 * -2.25)
                                                                default:
                                                                    v9 = append(v9, fi66)
                                                                }
                                                            default:
                                                                v1 = (v1 + (("hello" + v1) + v1))
                                                            }
                                                        default:
                                                            v2 -= v2
                                                        }
                                                    case "foo bar":
                                                        if (false || (false || (1 >= 4294967295))) {
                                                            if true {
                                                                var v71 string = (((v1 + "") + "a") + v1)
                                                                v72 := GenS8{F0: 382, F1: -1, F2: v71}
                                                                _ = v71
                                                                _ = v72
                                                            } else {
                                                                var v73 uint = 65535
                                                                _ = v73
                                                            }
                                                        }
                                                    }
                                                default:
                                                    switch v1 {
                                                    case "hello":
                                                        if (32767 >= 995) {
                                                            if len(v9) > 0 {
                                                                v9[0] = 792
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v5, v3)
                                                    }
                                                }
                                                if false {
                                                    _ = len(v9)
                                                    if ((((false && true) && true) || (true || (!false))) && (42 < 57)) {
                                                        if ((false && (!(!false))) && (((false && (!true)) && (true && (!true))) && true)) {
                                                            v74 := GenS9{F0: 1, F1: 164, F2: 107}
                                                            var v75 int = (fi66 - fi31)
                                                            switch (v1 + (v1 + "a")) {
                                                            case "a":
                                                                v76 := GenS4{F0: ""}
                                                                _ = v76
                                                            case "":
                                                                v77 := GenS5{F0: 255}
                                                                _ = v77
                                                            }
                                                            var v78 float32 = (v2 + -2.25)
                                                            _ = v74
                                                            _ = v75
                                                            _ = v78
                                                        } else {
                                                            fmt.Println(fi31, v5)
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v79 uint8 = 127
                                                                _ = v79
                                                            case "hello":
                                                                v5 = v5
                                                            case "":
                                                                switch fi31 {
                                                                case 4294967295:
                                                                    v9 = append(v9, (fi31 ^ 32768))
                                                                case 127:
                                                                    fmt.Println(fi66)
                                                                }
                                                            default:
                                                                v1 = v1
                                                            }
                                                            var v80 bool = false
                                                            var v81 uint = 255
                                                            var v82 uint8 = 17
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    }
                                                    if true {
                                                        switch int(v5) {
                                                        case -201:
                                                            v83 := GenS10{F0: ((fi31 + 1000) < 676), F1: (-128 == -128), F2: 2}
                                                            _ = v83
                                                        }
                                                        fmt.Println(v1)
                                                        if true {
                                                            v84 := GenS11{F0: v1, F1: (v5 * v5), F2: (v5 >> ((2) & 63))}
                                                            var v85 int8 = 85
                                                            var v86 float32 = 3.14159
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            if (!(0 <= 605)) {
                                                                v3 -= 12
                                                                if (false || true) {
                                                                    _ = len(v9)
                                                                } else {
                                                                    var v87 int64 = (v4 | (v4 / ((v4) | 1)))
                                                                    var v88 uint8 = 0
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                            if (0 != 18446744073709551615) {
                                                                var v89 string = ("" + "foo bar")
                                                                v90 := GenS7{F0: 75}
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                if len(v9) > 0 {
                                                                    v9[0] = fi66
                                                                }
                                                                if len(v9) > 0 {
                                                                    v9[0] = 127
                                                                }
                                                                if (42 < 59) {
                                                                    var v91 float32 = v2
                                                                    var v92 uint32 = 0
                                                                    v93 := GenS5{F0: 42}
                                                                    v5 = v5
                                                                    fmt.Println(v4)
                                                                    v94 := GenS12{F0: v1}
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                            }
                                                            _ = len(v9)
                                                        }
                                                    } else {
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        if (!(!false)) {
                                                            switch (^(fi66)) {
                                                            case 2147483647:
                                                                switch ((v1 + "foo bar") + v1) {
                                                                case "":
                                                                    v95 := GenS13{F0: -128}
                                                                    _ = v95
                                                                case "foo bar":
                                                                    var v96 uint32 = 876
                                                                    var v97 bool = (!(!(9223372036854775807 == (v4 ^ (^(v4))))))
                                                                    _ = v96
                                                                    _ = v97
                                                                case "\n\t":
                                                                    v9 = append(v9, (fi31 % ((-780) | 1)))
                                                                }
                                                            case -174:
                                                                switch 9223372036854775807 {
                                                                case 89:
                                                                    if true {
                                                                        var v98 int32 = -181
                                                                        var v99 string = v1
                                                                        fmt.Println(v2)
                                                                        fmt.Println(v98)
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        v5 += (v5 >> ((2) & 63))
                                                                    }
                                                                case 42:
                                                                    v100 := GenS11{F0: "", F1: int32(v4), F2: 1}
                                                                    _ = v100
                                                                }
                                                            case 9223372036854775807:
                                                                var v101 uint16 = (v3 / (((v3 * (v3 ^ 65535))) | 1))
                                                                _ = v101
                                                            default:
                                                                v4 = (v4 - 692)
                                                            }
                                                        } else {
                                                            switch v1 {
                                                            case "a":
                                                                v9 = append(v9, (fi66 % ((int(v4)) | 1)))
                                                            case "\n\t":
                                                                _ = len(v9)
                                                            }
                                                        }
                                                        v2 = v2
                                                    }
                                                }
                                                switch fi66 {
                                                case -440:
                                                    v102 := GenS10{F0: true, F1: (v5 < 0), F2: 18446744073709551615}
                                                    _ = v102
                                                case 475:
                                                    v2 -= 1e10
                                                case 4294967295:
                                                    var v103 float64 = (float64(v2) * 0.1)
                                                    _ = v103
                                                }
                                                switch v1 {
                                                case "foo bar":
                                                    if (!false) {
                                                        v3 += v3
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        v104 := GenS14{F0: (!((v3 == (v3 ^ v3)) && (v3 == (^(v3))))), F1: 212}
                                                        _ = v104
                                                    }
                                                }
                                                switch v1 {
                                                case "a":
                                                    var v105 uint8 = 42
                                                    _ = v105
                                                case "foo bar":
                                                    fmt.Println(v4, v4)
                                                    v4 -= (v4 << ((65535) & 63))
                                                    var v106 uint32 = 539
                                                    _ = v106
                                                default:
                                                    var v107 float32 = ((v2 * 3.14159) + -2.25)
                                                    _ = v107
                                                }
                                            }
                                        } else {
                                            switch 70 {
                                            case 9223372036854775807:
                                                if (true && (4294967295 <= 2147483648)) {
                                                    v5 += v5
                                                    v1 = v1
                                                }
                                            case 128:
                                                if len(v9) > 0 {
                                                    v9[0] = (fi31 % ((824) | 1))
                                                }
                                            case 852:
                                                var v108 int64 = v4
                                                _ = v108
                                            }
                                            var v109 int64 = int64(v5)
                                            v110 := GenS4{F0: "a"}
                                            v5 = -2147483648
                                            _ = v109
                                            _ = v110
                                        }
                                    }
                                    _ = fi66
                                }
                            }
                            for fi111 := 0; fi111 < 10; fi111++ {
                                switch 127 {
                                case 32768:
                                    fmt.Println(v3, v3, fi111)
                                default:
                                    if true {
                                        if (127 <= 9) {
                                            var v112 int16 = 0
                                            v5 += (v5 | (-(v5)))
                                            var v113 int8 = 42
                                            _ = v112
                                            _ = v113
                                        }
                                    }
                                }
                                fmt.Println(fi111)
                                switch ((v1 + v1) + (v1 + v1)) {
                                case "hello":
                                    if len(v9) > 0 {
                                        v9[0] = 42
                                    }
                                case "a":
                                    v114 := GenS11{F0: v1, F1: -2147483648, F2: 1}
                                    _ = v114
                                case "foo bar":
                                    var v115 uint32 = 65535
                                    if true {
                                        switch (fi111 & -739) {
                                        case 1:
                                            if true {
                                                v1 = v1
                                                v115 += 42
                                            }
                                        default:
                                            v5 += 776
                                        }
                                        v116 := GenS15{F0: (v5 + v5), F1: (v3 - v3)}
                                        v117 := GenS12{F0: v1}
                                        _ = v116
                                        _ = v117
                                    } else {
                                        v115 -= (v115 - (v115 % ((2147483648) | 1)))
                                        fmt.Println(v4)
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    v118 := GenS10{F0: (-1 < -1), F1: (0 != -111), F2: 65535}
                                    _ = v115
                                    _ = v118
                                default:
                                    if false {
                                        v5 += -330
                                        if false {
                                            fmt.Println(v5, v2)
                                        }
                                        v1 = "hello"
                                        if (4294967295 >= 1) {
                                            v119 := GenS8{F0: 42, F1: -128, F2: v1}
                                            v3 = (v3 / ((v3) | 1))
                                            fmt.Println(v5)
                                            _ = v119
                                        } else {
                                            v120 := GenS16{F0: 42, F1: "a", F2: 42}
                                            v5 += v5
                                            var v121 int8 = -1
                                            switch (-(fi111)) {
                                            case 127:
                                                if len(v9) > 0 {
                                                    v9[0] = int(v5)
                                                }
                                            case -1:
                                                v4 -= v4
                                            }
                                            _ = len(v9)
                                            _ = v120
                                            _ = v121
                                        }
                                        var v122 int = 2147483647
                                        _ = v122
                                    }
                                }
                                var v123 bool = (true && false)
                                _ = fi111
                                _ = v123
                            }
                        }
                    default:
                        fmt.Println(v2, v1)
                    }
                    _ = v32
                } else {
                    v124 := GenS17{F0: (4294967295 <= 1), F1: uint(v3)}
                    _ = v124
                }
            } else {
                v125 := GenS18{F0: 862, F1: 1, F2: 2147483647}
                var v126 int16 = 1000
                _ = v125
                _ = v126
            }
            _ = fi31
        }
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
