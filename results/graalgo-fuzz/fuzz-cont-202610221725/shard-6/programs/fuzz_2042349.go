package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint; F1 int32; F2 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int16; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 uint; F1 string; F2 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint8; F1 uint8; F2 bool }
type GenS14 struct { F0 uint64; F1 uint32 }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 string; F1 int32 }
type GenS17 struct { F0 int32 }
func genFn0(p0 int32, p1 uint16) int8 {
    v6 := []int{1}
    v7 := GenS1{F0: 4294967295, F1: 524}
    fmt.Println(p0, p1)
    p0 = (p0 - p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 69
}

func main() {
    var v1 float64 = -2.25
    var v2 int8 = 127
    var v3 int32 = 2147483647
    v4 := make(map[int]uint32)
    var v5 int = 32767
    for fi8 := 0; fi8 < 12; fi8++ {
        switch 2147483647 {
        case 81:
            v9 := GenS1{F0: 255, F1: 42}
            _ = v9
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v1, fi10)
        _ = fi10
    }
    v3 = 42
    v11 := GenS2{F0: 42}
    var v12 float64 = v1
    if true {
        var v13 uint16 = 65535
        _ = v13
    } else {
        if true {
            v3 = 1
            switch "foo bar" {
            case "a":
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint64 = 42
                    fmt.Println(v5)
                    _ = fi14
                    _ = v15
                }
            case "hello":
                fmt.Println(v2, v2, v1)
            case "\n\t":
                v5 = int(v3)
            default:
                v16 := []int{32768, v5}
                _ = v16
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := v4[(v5 << ((61) & 63))]
                v12 += (float64((((float32(((((v12 + (v1 + float64(((1.5 + 3.14159) + (((1e10 - 3.14159) * -2.25) * (3.14159 + 0.0)))))) - 1e10) * float64(-2.25)) - v1)) * 3.14159) - 1.5) + (0.0 + 0.1))) / (0.1 * float64(float32(1e10))))
                switch v5 {
                case 4294967295:
                    if (1 > 2) {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            if (((-(v3)) < v3) && true) {
                                var v20 bool = (0.1 > 1.5)
                                v3 = v3
                                if (((v12 - ((v1 - -2.25) * (3.14159 + float64((0.1 - 0.0))))) + v12) == v12) {
                                    if (!("a" != ("\n\t" + ("\n\t" + "foo bar")))) {
                                        if ((v3 & (-(v3))) >= (v3 | (v3 * v3))) {
                                            var v21 uint32 = 644
                                            var v22 bool = (v20 && (65535 < 434))
                                            v5 -= (v5 % ((-9223372036854775808) | 1))
                                            _ = v21
                                            _ = v22
                                        } else {
                                            switch 32767 {
                                            case -762:
                                                fmt.Println(v20, fi19)
                                            case 697:
                                                if true {
                                                    v23 := GenS3{F0: (^(v3)), F1: 18446744073709551615}
                                                    var v24 int16 = 32767
                                                    if ((("hello" + "foo bar") + ("" + ("\n\t" + "foo bar"))) >= "a") {
                                                        fmt.Println(v20, v2, v12)
                                                    } else {
                                                        var v25 uint8 = 42
                                                        v26 := GenS3{F0: (-(v3)), F1: 65535}
                                                        if true {
                                                            v27 := GenS3{F0: v3, F1: 2}
                                                            _ = v27
                                                        }
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case -2147483648:
                                                switch "a" {
                                                case "\n\t":
                                                    v3 = (v3 % (((v3 ^ (v3 & v3))) | 1))
                                                case "hello":
                                                    v3 = v3
                                                }
                                            default:
                                                v28 := []int{}
                                                _ = v28
                                            }
                                            v3 = (v3 * 932)
                                            var v29 string = "\n\t"
                                            _ = v29
                                        }
                                    } else {
                                        v30 := GenS4{F0: v5, F1: (v3 % ((0) | 1))}
                                        _ = len(v4)
                                        if (false || (("\n\t" + "a") > ("hello" + "\n\t"))) {
                                            v2 += genFn0(v3, 0)
                                            v31 := GenS5{F0: 1, F1: (v3 | (v3 | v3)), F2: -672}
                                            v2 += v2
                                            v32 := GenS6{F0: -1}
                                            _ = v31
                                            _ = v32
                                        } else {
                                            if v20 {
                                                var v33 float64 = (v12 + (float64(3.14159) + float64((float32(-2.25) - 1e10))))
                                                v34 := GenS7{F0: 18446744073709551615, F1: 42}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    v35 := []int{v5, fi17, int(v2)}
                                                    _ = v35
                                                default:
                                                    if v20 {
                                                        fmt.Println(v12)
                                                        if (!false) {
                                                            v5 -= v5
                                                            var v36 uint = 1
                                                            v2 += v2
                                                            _ = v36
                                                        }
                                                        v37 := []int{210, (fi17 >> ((56) & 63)), (fi19 << ((8) & 63)), -1}
                                                        _ = v37
                                                    }
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    var v38 string = ("\n\t" + "a")
                                                    _ = v38
                                                case "hello":
                                                    v20 = true
                                                case "a":
                                                    v39 := GenS6{F0: 42}
                                                    _ = v39
                                                default:
                                                    v12 = 1e10
                                                }
                                                delete(v4, (fi19 & 0))
                                                var v40 int = (^(fi19))
                                                fmt.Println(v20, v20, v2)
                                                var v41 bool = (v20 && true)
                                                _ = v40
                                                _ = v41
                                            }
                                            v1 += v12
                                            if (!v20) {
                                                v42 := GenS8{F0: -1, F1: 100, F2: (v18 / ((2147483648) | 1))}
                                                _ = v42
                                            } else {
                                                if v20 {
                                                    v43 := GenS4{F0: (fi19 | (-(v5))), F1: v3}
                                                    v1 = (((v1 + v1) / (v12 / v1)) - (v12 + 1.5))
                                                    var v44 int32 = (v3 << ((275) & 63))
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    v45 := []int{}
                                                    v4[(fi19 - (fi17 + (fi19 % (((fi19 / ((2147483647) | 1))) | 1))))] = v18
                                                    var v46 uint64 = 0
                                                    _ = v45
                                                    _ = v46
                                                }
                                                v3 = v3
                                            }
                                        }
                                        _ = v30
                                    }
                                }
                                _ = v20
                            } else {
                                var v47 float32 = 0.1
                                v47 = (v47 + v47)
                                if (int(v3) >= -2147483648) {
                                    v48 := GenS5{F0: 42, F1: (v3 << ((255) & 63)), F2: 231}
                                    _ = v48
                                } else {
                                    v1 = (float64((v47 + ((float32(v1) * v47) / (v47 - float32(v12))))) - 1.5)
                                    var v49 uint32 = (^(v18))
                                    if true {
                                        if ((fi19 % ((-776) | 1)) <= 128) {
                                            v3 += v3
                                            v50 := GenS1{F0: 1, F1: 42}
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                switch ("hello" + "hello") {
                                                case "foo bar":
                                                    v5 -= fi17
                                                case "hello":
                                                    if true {
                                                        if (!true) {
                                                            var v51 string = ("hello" + "hello")
                                                            fmt.Println(v3)
                                                            if true {
                                                                v52 := []int{}
                                                                if (!(1.5 != float32(v12))) {
                                                                    switch (("foo bar" + (("" + v51) + "a")) + ((v51 + v51) + v51)) {
                                                                    case "hello":
                                                                        v53 := GenS8{F0: 42, F1: v2, F2: 1}
                                                                        _ = v53
                                                                    case "foo bar":
                                                                        fmt.Println(v3)
                                                                    default:
                                                                        v54 := GenS1{F0: 505, F1: 4294967295}
                                                                        switch fi17 {
                                                                        case 599:
                                                                            switch (fi19 / ((int(v3)) | 1)) {
                                                                            case 128:
                                                                                switch v51 {
                                                                                case "hello":
                                                                                    v55 := GenS9{F0: 18446744073709551615, F1: -1}
                                                                                    _ = v55
                                                                                case "":
                                                                                    v56 := v4[(fi17 + (^(v5)))]
                                                                                    _ = v56
                                                                                case "\n\t":
                                                                                    if len(v52) >= 2 {
                                                                                        v52 = v52[1:len(v52)]
                                                                                    }
                                                                                default:
                                                                                    var v57 float32 = 1.5
                                                                                    _ = v57
                                                                                }
                                                                            }
                                                                        default:
                                                                            if false {
                                                                                v3 = 42
                                                                                v51 = ("" + v51)
                                                                            } else {
                                                                                fmt.Println(fi17)
                                                                                v52 = append(v52, (fi17 & (v5 << ((65535) & 63))))
                                                                                v58 := v4[32767]
                                                                                _ = v58
                                                                            }
                                                                        }
                                                                        v59 := GenS7{F0: 0, F1: 0}
                                                                        _ = v54
                                                                        _ = v59
                                                                    }
                                                                    v60 := GenS3{F0: (v3 & 1), F1: 42}
                                                                    _ = v60
                                                                }
                                                                _ = v52
                                                            } else {
                                                                v3 = (v3 + v3)
                                                                v1 += (v1 + v1)
                                                                v61 := v4[v5]
                                                                if false {
                                                                    v62 := []int{(fi19 << ((255) & 63)), (^(v5)), (^(fi17))}
                                                                    _ = v62
                                                                }
                                                                delete(v4, 0)
                                                                if ((v2 == genFn0(v3, 65535)) && (!(false && (2 <= 255)))) {
                                                                    if ((^(v3)) != (v3 - 1)) {
                                                                        v63 := GenS8{F0: 0, F1: 127, F2: v18}
                                                                        v64 := v4[(fi17 | int(v18))]
                                                                        v65 := GenS4{F0: 128, F1: v3}
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v65
                                                                    } else {
                                                                        var v66 int16 = 1000
                                                                        _ = v66
                                                                    }
                                                                    fmt.Println(v5, v12, v49)
                                                                } else {
                                                                    v67 := v4[fi17]
                                                                    v68 := GenS3{F0: v3, F1: 255}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        _ = len(v4)
                                                                        if false {
                                                                            fmt.Println(v51, fi17)
                                                                            switch "a" {
                                                                            case "":
                                                                                v69 := GenS3{F0: v3, F1: 793}
                                                                                _ = v69
                                                                            case "hello":
                                                                                switch ((v51 + (v51 + "foo bar")) + (v51 + v51)) {
                                                                                case "a":
                                                                                    fmt.Println(v1, v61)
                                                                                case "":
                                                                                    switch "" {
                                                                                    case "":
                                                                                        var v70 int = (fi17 / ((fi17) | 1))
                                                                                        _ = v70
                                                                                    }
                                                                                default:
                                                                                    v3 = v3
                                                                                }
                                                                                var v71 uint = 803
                                                                                var v72 int16 = 0
                                                                                _ = v71
                                                                                _ = v72
                                                                            case "a":
                                                                                var v73 float64 = (1e10 + float64(-2.25))
                                                                                _ = v73
                                                                            }
                                                                            v74 := GenS4{F0: fi19, F1: -2147483648}
                                                                            v2 = (v2 >> ((277) & 63))
                                                                            v4[336] = v18
                                                                            if (65535 != 2) {
                                                                                var v75 int64 = -9223372036854775808
                                                                                v51 = "\n\t"
                                                                                _ = v75
                                                                            }
                                                                            v61 += v18
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v76 := GenS10{F0: 2, F1: v51, F2: fi17}
                                                                                _ = v76
                                                                            case "a":
                                                                                v77 := v4[(fi17 << ((529) & 63))]
                                                                                _ = v77
                                                                            case "hello":
                                                                                v78 := GenS8{F0: 42, F1: (v2 | (v2 + v2)), F2: (^(v67))}
                                                                                _ = v78
                                                                            default:
                                                                                var v79 float32 = v47
                                                                                _ = v79
                                                                            }
                                                                            var v80 uint64 = 604
                                                                            _ = v74
                                                                            _ = v80
                                                                        } else {
                                                                            v4[201] = v67
                                                                        }
                                                                        if true {
                                                                            v5 = fi17
                                                                            v81 := GenS5{F0: 217, F1: v3, F2: 0}
                                                                            if (18446744073709551615 >= 1) {
                                                                                v12 += v12
                                                                                v67 += (v61 >> ((42) & 63))
                                                                            } else {
                                                                                switch 128 {
                                                                                case 2147483647:
                                                                                    switch v51 {
                                                                                    case "\n\t":
                                                                                        var v82 float32 = ((float32(float64((v47 + (float32(0.0) + v47)))) / v47) + ((v47 - float32((((((v1 * (float64(v47) * v12)) - 0.0) * (v12 / 1.0)) / v1) * ((3.14159 * (float64(v47) + float64((((0.1 + float32(0.1)) + float32((3.14159 * 1.5))) + v47)))) * -2.25)))) - ((v47 + v47) / 1.5)))
                                                                                        _ = v82
                                                                                    case "hello":
                                                                                        v18 -= 2147483647
                                                                                    default:
                                                                                        switch v51 {
                                                                                        case "\n\t":
                                                                                            var v83 int32 = (v3 << ((319) & 63))
                                                                                            _ = v83
                                                                                        case "foo bar":
                                                                                            v84 := GenS10{F0: 386, F1: v51, F2: (v5 | fi17)}
                                                                                            _ = v84
                                                                                        default:
                                                                                            v85 := []int{127, (^(v5))}
                                                                                            _ = v85
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v86 string = v51
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                            fmt.Println(v67, fi19)
                                                                            fmt.Println(v51)
                                                                            var v87 int8 = 0
                                                                            v88 := GenS6{F0: 1}
                                                                            v89 := GenS11{F0: int32(v87)}
                                                                            var v90 string = ("a" + v51)
                                                                            if (false && (true && (v87 > (v87 % ((1) | 1))))) {
                                                                                switch v90 {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        v18 = (v61 / (((^(v67))) | 1))
                                                                                    } else {
                                                                                        switch (v51 + ("foo bar" + v51)) {
                                                                                        case "hello":
                                                                                            switch (v51 + v51) {
                                                                                            case "hello":
                                                                                                switch (("\n\t" + (v51 + v90)) + ("\n\t" + v90)) {
                                                                                                case "\n\t":
                                                                                                    switch (v90 + v51) {
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            v90 = ("foo bar" + "a")
                                                                                                        } else {
                                                                                                            v4[fi17] = (v18 << ((0) & 63))
                                                                                                            v91 := GenS7{F0: 266, F1: 115}
                                                                                                            if true {
                                                                                                                if (688 > (v61 + 0)) {
                                                                                                                    switch ("foo bar" + (v90 + ("hello" + (v90 + (v51 + v90))))) {
                                                                                                                    case "a":
                                                                                                                        v18 = 2147483647
                                                                                                                    case "":
                                                                                                                        v4[-9223372036854775808] = (^(v49))
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v67, v5, v49)
                                                                                                                    default:
                                                                                                                        var v92 uint64 = 18446744073709551615
                                                                                                                        _ = v92
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v91
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        delete(v4, (-(fi17)))
                                                                                                    }
                                                                                                case "a":
                                                                                                    if true {
                                                                                                        switch fi17 {
                                                                                                        case 32768:
                                                                                                            v12 = (v1 - -2.25)
                                                                                                        case 786:
                                                                                                            switch (v5 + -9223372036854775808) {
                                                                                                            case 9223372036854775807:
                                                                                                                fmt.Println(v2, v2)
                                                                                                            }
                                                                                                        case 42:
                                                                                                            _ = len(v4)
                                                                                                        }
                                                                                                    } else {
                                                                                                        v3 -= (v3 + v3)
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v93 := []int{(-(fi19)), 2147483647, (fi19 & v5), v5}
                                                                                                _ = v93
                                                                                            }
                                                                                        }
                                                                                        v18 += v49
                                                                                    }
                                                                                case "":
                                                                                    v94 := v4[fi19]
                                                                                    _ = v94
                                                                                default:
                                                                                    fmt.Println(fi17, v87, v47)
                                                                                }
                                                                            } else {
                                                                                v3 += (v3 - v3)
                                                                                v95 := v4[0]
                                                                                v67 = 2147483648
                                                                                _ = v95
                                                                            }
                                                                            _ = v81
                                                                            _ = v87
                                                                            _ = v88
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                    case "":
                                                                        var v96 int8 = v2
                                                                        _ = v96
                                                                    case "a":
                                                                        var v97 uint32 = (^(v18))
                                                                        _ = v97
                                                                    default:
                                                                        fmt.Println(v12, v1, v2)
                                                                    }
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                fmt.Println(v12, v1, fi19)
                                                                _ = v61
                                                            }
                                                            _ = v51
                                                        } else {
                                                            v98 := []int{fi17}
                                                            v99 := GenS12{F0: 42}
                                                            var v100 uint32 = (^(v49))
                                                            _ = v100
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                        var v101 uint64 = 271
                                                        _ = v101
                                                    } else {
                                                        v102 := []int{(fi17 << ((42) & 63)), 9223372036854775807, v5, fi19}
                                                        _ = v102
                                                    }
                                                }
                                            case "a":
                                                v103 := GenS13{F0: 2, F1: 167, F2: (1 >= -446)}
                                                _ = v103
                                            case "":
                                                fmt.Println(fi17, v5)
                                            default:
                                                var v104 int16 = 0
                                                _ = v104
                                            }
                                            var v105 int8 = (v2 ^ (v2 % ((117) | 1)))
                                            _ = v105
                                            _ = v50
                                        } else {
                                            v47 -= v47
                                            switch "a" {
                                            case "":
                                                v3 += (-(v3))
                                            case "hello":
                                                fmt.Println(fi17)
                                            case "\n\t":
                                                if (true || false) {
                                                    fmt.Println(v49, v3, v1)
                                                    fmt.Println(v3, v12, v3)
                                                    if false {
                                                        v106 := v4[(^(v5))]
                                                        _ = v106
                                                    } else {
                                                        v12 = float64((0.1 - (3.14159 / -2.25)))
                                                        v12 = 0.0
                                                    }
                                                    v12 = (v1 / v1)
                                                } else {
                                                    var v107 int = 9223372036854775807
                                                    _ = v107
                                                }
                                            }
                                            var v108 uint64 = 4294967295
                                            v109 := GenS5{F0: 65535, F1: (-(v3)), F2: 408}
                                            _ = v108
                                            _ = v109
                                        }
                                        v110 := []int{}
                                        v111 := GenS5{F0: 18446744073709551615, F1: (v3 / (((v3 * (v3 | v3))) | 1)), F2: 614}
                                        _ = v110
                                        _ = v111
                                    }
                                    _ = v49
                                }
                                _ = v47
                            }
                            _ = fi19
                        }
                        if (!true) {
                            switch v5 {
                            case 42:
                                for fi112 := 0; fi112 < 9; fi112++ {
                                    v1 = -2.25
                                    switch (-(fi17)) {
                                    case 0:
                                        var v113 uint = 4294967295
                                        _ = v113
                                    case -2147483648:
                                        v114 := GenS14{F0: 0, F1: (v18 + 746)}
                                        _ = v114
                                    case 1:
                                        v115 := GenS15{F0: 137, F1: 128}
                                        _ = v115
                                    default:
                                        var v116 int = 32767
                                        _ = v116
                                    }
                                    _ = len(v4)
                                    _ = fi112
                                }
                            case 32768:
                                v117 := v4[(v5 | fi17)]
                                _ = v117
                            case 128:
                                var v118 uint = 65535
                                _ = v118
                            }
                            for fi119 := 0; fi119 < 8; fi119++ {
                                if false {
                                    var v120 uint64 = 42
                                    v121 := []int{(fi119 & (fi119 << ((255) & 63))), -2147483648, v5, 2147483647}
                                    if (fi17 == (fi119 * -364)) {
                                        switch 0 {
                                        case 897:
                                            v2 += -74
                                        case -2147483648:
                                            v2 = (v2 & (v2 | v2))
                                        case 32768:
                                            var v122 uint = 2
                                            _ = v122
                                        default:
                                            var v123 uint32 = 700
                                            _ = v123
                                        }
                                    } else {
                                        if (255 != 32767) {
                                            v120 -= (v120 / ((v120) | 1))
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            v124 := GenS16{F0: "\n\t", F1: (v3 >> ((4294967295) & 63))}
                                            _ = v124
                                        case "":
                                            v4[804] = v18
                                        case "\n\t":
                                            var v125 int64 = 42
                                            _ = v125
                                        }
                                    }
                                    v5 += 32767
                                    _ = v120
                                    _ = v121
                                } else {
                                    fmt.Println(v5, v18, fi119)
                                }
                                fmt.Println(v1, v5, v5)
                                _ = fi119
                            }
                        }
                        v126 := v4[v5]
                        _ = v126
                    }
                case 1:
                    for fi127 := 0; fi127 < 15; fi127++ {
                        var v128 float32 = (float32(v12) / 3.14159)
                        switch "hello" {
                        case "hello":
                            v129 := GenS3{F0: 42, F1: 688}
                            _ = v129
                        }
                        _ = fi127
                        _ = v128
                    }
                case 976:
                    _ = len(v4)
                default:
                    var v130 float32 = ((0.1 / (0.0 + (1.5 * ((3.14159 - -2.25) / -2.25)))) + float32(v12))
                    _ = v130
                }
                _ = fi17
                _ = v18
            }
        }
    }
    v1 = (0.0 - (v1 - (0.1 + float64((-2.25 + 1.5)))))
    v131 := GenS17{F0: (v3 << ((18446744073709551615) & 63))}
    var v132 int64 = 42
    fmt.Println(v132)
    var v133 int = (v5 - (v5 & int(v3)))
    if false {
        var v134 uint32 = 0
        fmt.Println(v133, v3, v12)
        var v135 uint32 = (v134 & (^(v134)))
        _ = v134
        _ = v135
    }
    var v136 uint8 = 0
    if (1 <= 65535) {
        if (18446744073709551615 > 710) {
            fmt.Println(v132, v3, v12)
        }
        var v137 int8 = (v2 >> ((585) & 63))
        _ = v137
    }
    v138 := []int{(v5 ^ (v133 ^ (v5 * -1))), v5}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v136:", v136)
    fmt.Println("slicelen v138:", len(v138), v138)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v131
    _ = v132
    _ = v133
    _ = v136
    _ = v138
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
