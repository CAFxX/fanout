package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int64; F1 int16; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 uint8; F2 uint }
func genFn0(p0 int) int {
    p0 += (p0 & p0)
    fmt.Println(p0)
    _ = p0
    return (p0 & p0)
}
func genFn1(p0 int32) int8 {
    var v7 float64 = (1e10 / (0.0 + (1.5 + 1e10)))
    v7 -= 0.0
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = -1
    v3 := []int{128, -1, 128}
    v4 := []uint32{65535, v1}
    var v5 uint16 = 255
    v6 := make(map[string]int)
    var v8 int16 = -32768
    v2 -= v2
    v9 := GenS1{F0: 896}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 string = "foo bar"
        if false {
            fmt.Println(v2)
            switch v11 {
            case "\n\t":
                fmt.Println(v2, v1, v2)
            case "":
                var v12 int32 = -1
                _ = v12
            case "a":
                if ((!true) && (100 > 100)) {
                    v2 += 42
                }
            }
            if true {
                var v13 int = fi10
                _ = v13
            } else {
                var v14 int32 = -2147483648
                _ = v14
            }
        }
        _ = fi10
        _ = v11
    }
    var v15 uint = 255
    if true {
        switch 445 {
        case -822:
            fmt.Println(v5, v5, v2)
        default:
            switch 0 {
            case 32767:
                var v16 bool = true
                _ = v16
            case 9223372036854775807:
                var v17 int = 2147483647
                _ = v17
            default:
                switch "foo bar" {
                case "hello":
                    var v18 int8 = -128
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v21 := GenS1{F0: 18446744073709551615}
                            if ((((3.14159 > -2.25) && (false || (42 < 1))) && (uint32(v15) > (v1 << (((v15 ^ 4294967295)) & 63)))) || false) {
                                v5 = v5
                                _ = len(v3)
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = len(v6)
                        if false {
                            var v22 int = fi19
                            for fi23 := 0; fi23 < 16; fi23++ {
                                var v24 int32 = 2147483647
                                v25 := GenS1{F0: 4294967295}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(v8, fi19)
                            _ = v22
                        } else {
                            var v26 float64 = 1e10
                            _ = v26
                        }
                        _ = fi19
                    }
                }
            }
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            v2 = v2
            if (!(v1 <= v1)) {
                switch fi27 {
                case 9223372036854775807:
                    for fi28 := 0; fi28 < 17; fi28++ {
                        _ = len(v4)
                        var v29 uint8 = 255
                        var v30 uint64 = 489
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                default:
                    for fi31 := 0; fi31 < 11; fi31++ {
                        var v32 int16 = v8
                        v2 -= v2
                        _ = fi31
                        _ = v32
                    }
                }
                switch 32768 {
                case 127:
                    switch 476 {
                    case 32767:
                        var v33 int8 = 127
                        _ = v33
                    case 466:
                        if false {
                            v2 += 1
                            var v34 int16 = (v8 | 456)
                            _ = v34
                        } else {
                            v5 += v5
                            var v35 uint16 = v5
                            _ = v35
                        }
                    case 32768:
                        v8 -= (v8 / (((v8 + v8)) | 1))
                    }
                case 32767:
                    fmt.Println(fi27, v8, fi27)
                case -9223372036854775808:
                    fmt.Println(v8, v1, v1)
                default:
                    v1 -= (v1 << ((556) & 63))
                }
                switch ("hello" + "foo bar") {
                case "":
                    v36 := GenS1{F0: 2}
                    _ = v36
                case "\n\t":
                    v15 = v15
                case "foo bar":
                    if ((float32(1e10) > float32(0.1)) || (uint(v5) > v15)) {
                        fmt.Println(v1, v5, v5)
                    } else {
                        for fi37 := 0; fi37 < 18; fi37++ {
                            fmt.Println(v1, fi37, v15)
                            if (uint16(v1) > v5) {
                                v1 += (v1 % (((v1 | 1)) | 1))
                                switch (("a" + "foo bar") + "foo bar") {
                                case "":
                                    if (float64((0.1 / (-2.25 - (1e10 + (((-2.25 * ((0.1 / (0.1 * 3.14159)) / -2.25)) - (((((1.5 * 1e10) + 3.14159) * (1e10 + ((1e10 / 1.5) * 0.1))) / (3.14159 + float32(1.5))) + 0.1)) + ((1e10 - (1e10 / ((-2.25 - float32(1e10)) - 3.14159))) / -2.25)))))) < 1e10) {
                                        switch (fi27 - (-(fi27))) {
                                        case 114:
                                            fmt.Println(v5, v8, v5)
                                        case -9223372036854775808:
                                            v8 = (v8 % ((v8) | 1))
                                        case -488:
                                            var v38 float64 = -2.25
                                            _ = v38
                                        }
                                        switch fi27 {
                                        case 32768:
                                            var v39 float32 = 3.14159
                                            _ = v39
                                        case -723:
                                            v40 := GenS2{F0: -128, F1: (^(v1))}
                                            _ = v40
                                        }
                                        var v41 uint32 = 1
                                        _ = v41
                                    }
                                case "a":
                                    fmt.Println(v2)
                                default:
                                    var v42 uint8 = uint8(v15)
                                    v8 -= 32767
                                    _ = v42
                                }
                                v1 += 2
                            } else {
                                v8 -= (v8 ^ v8)
                                v43 := GenS1{F0: 611}
                                var v44 int16 = 32767
                                _ = v43
                                _ = v44
                            }
                            if (!true) {
                                v6[("hello" + "hello")] = (fi27 - (fi27 >> ((v15) & 63)))
                                if (-2.25 < 1e10) {
                                    _ = len(v6)
                                    switch 128 {
                                    case 32767:
                                        var v45 uint8 = 128
                                        _ = v45
                                    case 4294967295:
                                        switch "hello" {
                                        case "":
                                            switch fi37 {
                                            case 145:
                                                fmt.Println(v5)
                                                v46 := GenS2{F0: -1, F1: (v1 % ((uint32(v5)) | 1))}
                                                _ = v46
                                            }
                                        case "hello":
                                            v2 -= (v2 / ((-1) | 1))
                                        case "foo bar":
                                            v47 := GenS3{F0: (v8 < v8)}
                                            _ = v47
                                        }
                                    }
                                    var v48 uint8 = 172
                                    v49 := GenS1{F0: 0}
                                    v8 -= (v8 << (((v15 >> (((v15 >> ((v15) & 63))) & 63))) & 63))
                                    _ = v48
                                    _ = v49
                                }
                                switch (-(fi37)) {
                                case 32768:
                                    v1 += v1
                                case 128:
                                    switch "foo bar" {
                                    case "foo bar":
                                        v50 := GenS4{F0: -1, F1: false}
                                        _ = v50
                                    case "a":
                                        var v51 int = (fi27 ^ (fi27 | 2147483647))
                                        _ = v51
                                    default:
                                        v2 = (v2 << ((42) & 63))
                                    }
                                case -9223372036854775808:
                                    switch "\n\t" {
                                    case "hello":
                                        var v52 string = "hello"
                                        _ = v52
                                    case "\n\t":
                                        if ((!("\n\t" <= "")) && (((0.0 * ((1.5 + 0.1) + float64((-2.25 - ((((-2.25 - 0.1) - 3.14159) - (-2.25 + 1.5)) * (3.14159 / 0.1)))))) * (float64(((0.1 * 1e10) + 0.0)) + -2.25)) != 0.1)) {
                                            fmt.Println(v1)
                                        } else {
                                            var v53 uint8 = 1
                                            v1 = 2147483648
                                            _ = v53
                                        }
                                    case "foo bar":
                                        fmt.Println(v2, fi37, v8)
                                    default:
                                        v6["hello"] = (fi37 ^ (fi37 - (fi37 << ((v15) & 63))))
                                    }
                                }
                            }
                            _ = fi37
                        }
                    }
                default:
                    var v54 float64 = ((3.14159 - -2.25) + 3.14159)
                    _ = v54
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (1e10 != (3.14159 * 1e10)) {
                    if (((v1 << (((v15 % ((v15) | 1))) & 63)) != (v1 >> ((v15) & 63))) && (-2147483648 > -942)) {
                        if ((!(true || (true && ((185 < 39) && false)))) && (((((1e10 / -2.25) * (0.0 - float64((((float32(((0.0 * (0.1 + 0.1)) - -2.25)) / ((1.5 / (3.14159 - (1e10 / 3.14159))) / 1e10)) / (((3.14159 + (3.14159 * 1e10)) / 3.14159) - 0.0)) - 0.1)))) - 1e10) * ((1.5 / 1.0) / -2.25)) >= 0.1)) {
                            if true {
                                v55 := GenS3{F0: (false && false)}
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 17; fi56++ {
                                    if false {
                                        if true {
                                            v6["\n\t"] = (fi27 - 812)
                                        } else {
                                            v1 -= v1
                                            _ = len(v6)
                                            _ = len(v6)
                                        }
                                        fmt.Println(fi27)
                                        if false {
                                            fmt.Println(fi56, fi56, v1)
                                            if (v5 == 669) {
                                                delete(v6, ("a" + ("" + ("" + "\n\t"))))
                                                var v57 bool = (!((!false) && true))
                                                v58 := GenS5{F0: 42, F1: int16(v5), F2: (v8 | -1)}
                                                _ = v57
                                                _ = v58
                                            } else {
                                                v1 += (v1 >> ((1) & 63))
                                                var v59 uint = 18446744073709551615
                                                var v60 float64 = (1e10 * 1.5)
                                                _ = v59
                                                _ = v60
                                            }
                                            var v61 float64 = (0.0 + (1e10 * 1.5))
                                            if false {
                                                v5 = 0
                                                switch fi56 {
                                                case 4294967295:
                                                    var v62 int8 = genFn1(0)
                                                    _ = v62
                                                default:
                                                    var v63 int64 = -9223372036854775808
                                                    _ = v63
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = fi27
                                                }
                                            } else {
                                                fmt.Println(v61, v2, fi27)
                                                v64 := GenS6{F0: (v8 | v8)}
                                                v65 := GenS7{F0: v15}
                                                _ = v64
                                                _ = v65
                                            }
                                            v15 -= (v15 % (((^(v15))) | 1))
                                            _ = v61
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v2, v8, v2)
                                    _ = fi56
                                }
                            }
                        } else {
                            var v66 uint8 = uint8(v15)
                            var v67 int8 = -110
                            _ = v66
                            _ = v67
                        }
                        fmt.Println(v15)
                    }
                    var v68 uint64 = 951
                    _ = v68
                }
            }
            v3 = append(v3, -9223372036854775808)
            _ = fi27
        }
    } else {
        v4 = append(v4, (v1 % ((2147483648) | 1)))
        var v69 uint = (v15 | v15)
        v1 = (v1 << ((v69) & 63))
        _ = v69
    }
    var v70 int16 = v8
    var v71 int64 = (v2 | (v2 | (v2 << (((v15 | v15)) & 63))))
    v8 = v70
    switch -959 {
    case 4294967295:
        v6["hello"] = genFn0(0)
    default:
        var v72 string = (("\n\t" + "") + "foo bar")
        _ = v72
    }
    if true {
        var v73 int = -844
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v73
    }
    switch "hello" {
    case "hello":
        for fi74 := 0; fi74 < 18; fi74++ {
            var v75 uint16 = (v5 * (v5 * (v5 * v5)))
            var v76 int32 = 42
            v2 = (^(v71))
            _ = fi74
            _ = v75
            _ = v76
        }
    case "a":
        v77 := GenS2{F0: 18, F1: v1}
        _ = v77
    case "\n\t":
        v3 = append(v3, -845)
    default:
        v78 := GenS5{F0: 9223372036854775807, F1: -1, F2: v70}
        _ = v78
    }
    v1 += 2
    if true {
        for fi79 := 0; fi79 < 8; fi79++ {
            var v80 int16 = (v8 - v70)
            if (false || (!true)) {
                v81 := GenS1{F0: 18446744073709551615}
                v82 := GenS3{F0: false}
                _ = len(v3)
                _ = v81
                _ = v82
            }
            _ = fi79
            _ = v80
        }
    } else {
        if false {
            for fi83 := 0; fi83 < 19; fi83++ {
                fmt.Println(v70, fi83, v5)
                v84 := v6[""]
                if len(v4) > 0 {
                    v4[0] = uint32(v5)
                }
                _ = fi83
                _ = v84
            }
            v8 = 42
            _ = len(v4)
        } else {
            var v85 string = (("\n\t" + "hello") + "\n\t")
            if false {
                v86 := GenS6{F0: 32767}
                v2 = v2
                _ = v86
            }
            if true {
                v70 = -156
                for fi87 := 0; fi87 < 5; fi87++ {
                    delete(v6, "a")
                    if false {
                        if ((v1 == (v1 % (((v1 ^ (v1 / ((1) | 1)))) | 1))) && true) {
                            fmt.Println(v5)
                            switch "foo bar" {
                            case "\n\t":
                                var v88 bool = false
                                _ = v88
                            case "hello":
                                var v89 float64 = 1e10
                                _ = v89
                            case "":
                                fmt.Println(v2, v85)
                            default:
                                v8 += v70
                            }
                        }
                    } else {
                        v90 := v6[(("hello" + ("hello" + v85)) + "\n\t")]
                        v91 := GenS3{F0: false}
                        _ = v90
                        _ = v91
                    }
                    _ = fi87
                }
            }
            _ = v85
        }
        var v92 int = -469
        _ = v92
    }
    var v93 int = genFn0(-430)
    for fi94 := 0; fi94 < 19; fi94++ {
        if true {
            v95 := GenS3{F0: true}
            _ = v95
        }
        if true {
            for fi96 := 0; fi96 < 17; fi96++ {
                switch (fi94 | int(v5)) {
                case 9223372036854775807:
                    if (true && (!false)) {
                        v70 = v70
                        _ = len(v4)
                    } else {
                        var v97 int32 = 42
                        delete(v6, "hello")
                        switch ("\n\t" + ("hello" + "\n\t")) {
                        case "\n\t":
                            v8 = (v70 >> ((v15) & 63))
                            var v98 float32 = 1.5
                            var v99 uint = 18446744073709551615
                            v100 := GenS8{F0: (v70 / (((-(v70))) | 1)), F1: 255, F2: 712}
                            var v101 uint32 = 2
                            _ = v100
                            _ = v101
                            _ = v98
                            _ = v99
                        default:
                            fmt.Println(v97, fi94)
                        }
                        _ = v97
                    }
                case 2147483647:
                    var v102 string = "a"
                    _ = v102
                default:
                    var v103 uint32 = (v1 / ((v1) | 1))
                    _ = v103
                }
                _ = fi96
            }
            v104 := GenS5{F0: v71, F1: (v70 * -1), F2: 0}
            _ = v104
        } else {
            switch genFn0(v93) {
            case 32768:
                switch (fi94 * -398) {
                case 32767:
                    fmt.Println(v71, v15, v2)
                case -9223372036854775808:
                    v6[("foo bar" + "a")] = (v93 | fi94)
                case 0:
                    v5 -= v5
                }
            case 688:
                if (18446744073709551615 > 18446744073709551615) {
                    var v105 int = (v93 % ((fi94) | 1))
                    if true {
                        var v106 int8 = -128
                        v1 = v1
                        _ = v106
                    } else {
                        switch 32767 {
                        case -9223372036854775808:
                            v70 = -32768
                        }
                        fmt.Println(v105)
                        switch genFn0((v93 % ((fi94) | 1))) {
                        case 127:
                            v70 += (v8 * v70)
                        case 42:
                            delete(v6, ("foo bar" + ("foo bar" + "hello")))
                        default:
                            switch "a" {
                            case "a":
                                var v107 uint = (^(v15))
                                _ = v107
                            }
                        }
                    }
                    _ = v105
                } else {
                    v6["a"] = genFn0(genFn0(genFn0(v93)))
                }
            default:
                v70 = (v8 & (v8 + v8))
            }
            delete(v6, ("foo bar" + "foo bar"))
            fmt.Println(v15, v70)
        }
        switch "" {
        case "foo bar":
            var v108 uint16 = 255
            _ = v108
        case "":
            v6["foo bar"] = (v93 % (((fi94 + (fi94 >> (((v15 ^ (v15 << ((v15) & 63)))) & 63)))) | 1))
        default:
            var v109 uint8 = 0
            _ = v109
        }
        _ = fi94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v8:", v8)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v8
    _ = v9
    _ = v93
}
