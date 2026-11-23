package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 int }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 bool; F1 uint64 }
func genFn0() string {
    if true {
        switch -9223372036854775808 {
        case 127:
            v4 := GenS1{F0: 389}
            _ = v4
        case 883:
            var v5 bool = ((778 <= -1) && (2147483647 > 1))
            _ = v5
        default:
            var v6 float32 = (1.5 * float32(-2.25))
            _ = v6
        }
    } else {
        var v7 bool = true
        var v8 string = "a"
        _ = v7
        _ = v8
    }
    var v9 int16 = -410
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint64) string {
    for fi10 := 0; fi10 < 4; fi10++ {
        if (true && true) {
            v11 := []int{fi10, 1}
            _ = v11
        } else {
            fmt.Println(p0, p0)
            var v12 float64 = 0.0
            _ = v12
        }
        fmt.Println(fi10, p0)
        _ = fi10
    }
    var v13 uint64 = p0
    _ = p0
    _ = v13
    return ""
}

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = (false && false)
    var v3 string = ("" + "")
    v2 = v2
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint8 = 0
            var v17 int16 = -596
            v3 = v3
            _ = fi15
            _ = v16
            _ = v17
        }
        v2 = false
        var v18 uint64 = (v1 | v1)
        _ = fi14
        _ = v18
    }
    var v19 int8 = 127
    var v20 float64 = float64(float32(0.1))
    v1 = (v1 * v1)
    switch "hello" {
    case "hello":
        v21 := []int{0}
        _ = v21
    default:
        for fi22 := 0; fi22 < 4; fi22++ {
            switch (fi22 >> ((1) & 63)) {
            case 9223372036854775807:
                if (3.14159 == ((1e10 + 3.14159) - 3.14159)) {
                    var v23 int16 = -802
                    var v24 uint32 = 65535
                    _ = v23
                    _ = v24
                }
            case -9223372036854775808:
                switch (fi22 | fi22) {
                case 42:
                    var v25 uint8 = 128
                    _ = v25
                case 128:
                    v2 = false
                }
            case 1:
                var v26 int8 = (v19 & (-(v19)))
                _ = v26
            default:
                v27 := GenS2{F0: -1}
                _ = v27
            }
            switch (fi22 & 128) {
            case 128:
                v28 := []int{32767}
                _ = v28
            case -207:
                v29 := []int{(fi22 + fi22)}
                _ = v29
            default:
                fmt.Println(v20, fi22, fi22)
            }
            _ = fi22
        }
    }
    v30 := []int{}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS3{F0: 42, F1: 803}
        if ((true || false) || (false || true)) {
            v33 := GenS1{F0: 606}
            v30 = append(v30, fi31)
            _ = v33
        } else {
            for fi34 := 0; fi34 < 7; fi34++ {
                var v35 uint64 = (v1 % ((0) | 1))
                _ = fi34
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 7; fi36++ {
            v19 = (v19 % (((v19 << ((690) & 63))) | 1))
            _ = fi36
        }
        _ = fi31
        _ = v32
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 string = (v3 + ("" + v3))
        for fi39 := 0; fi39 < 1; fi39++ {
            v2 = v2
            v40 := GenS3{F0: 9223372036854775807, F1: uint32(v1)}
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    v41 := GenS4{F0: 32767}
    if v2 {
        v20 = (v20 * (0.1 / v20))
        for fi42 := 0; fi42 < 14; fi42++ {
            var v43 uint = 255
            if (!v2) {
                if v2 {
                    v44 := GenS4{F0: 683}
                    fmt.Println(v19, v19, v43)
                    _ = v44
                } else {
                    var v45 int16 = 32767
                    _ = v45
                }
                fmt.Println(v2)
                var v46 string = ""
                _ = v46
            } else {
                for fi47 := 0; fi47 < 11; fi47++ {
                    var v48 uint8 = 21
                    _ = fi47
                    _ = v48
                }
            }
            var v49 uint16 = 1
            _ = fi42
            _ = v43
            _ = v49
        }
    }
    v30 = append(v30, 32768)
    if v2 {
        if v2 {
            if v2 {
                v50 := GenS4{F0: 566}
                for fi51 := 0; fi51 < 2; fi51++ {
                    v1 = 2
                    v3 = v3
                    switch v3 {
                    case "\n\t":
                        switch -9223372036854775808 {
                        case 338:
                            if v2 {
                                if (v20 == v20) {
                                    v1 -= (^(v1))
                                    if false {
                                        for fi52 := 0; fi52 < 12; fi52++ {
                                            v2 = (42 != 103)
                                            var v53 bool = (32767 > 32767)
                                            _ = fi52
                                            _ = v53
                                        }
                                    } else {
                                        for fi54 := 0; fi54 < 19; fi54++ {
                                            fmt.Println(v19)
                                            v1 = (v1 >> ((42) & 63))
                                            var v55 string = "hello"
                                            _ = fi54
                                            _ = v55
                                        }
                                        for fi56 := 0; fi56 < 15; fi56++ {
                                            var v57 uint16 = 65535
                                            _ = fi56
                                            _ = v57
                                        }
                                        v58 := GenS2{F0: 1}
                                        _ = v58
                                    }
                                }
                                switch v3 {
                                case "":
                                    fmt.Println(v2, v3)
                                }
                                v59 := GenS1{F0: 42}
                                _ = v59
                            }
                        case 32768:
                            fmt.Println(fi51, fi51)
                        }
                    case "":
                        if (65535 == 572) {
                            v60 := GenS2{F0: 9223372036854775807}
                            v61 := GenS2{F0: int64(v19)}
                            var v62 uint8 = 255
                            _ = v60
                            _ = v61
                            _ = v62
                        } else {
                            if ("" <= (v3 + "a")) {
                                for fi63 := 0; fi63 < 12; fi63++ {
                                    if v2 {
                                        v20 = float64(float32(0.0))
                                        var v64 float32 = float32(v20)
                                        _ = v64
                                    }
                                    fmt.Println(v2)
                                    if ((v19 ^ (v19 % ((v19) | 1))) != v19) {
                                        switch (v3 + ("" + v3)) {
                                        case "hello":
                                            if len(v30) > 0 {
                                                v30[0] = (^(fi63))
                                            }
                                        case "a":
                                            var v65 int8 = -9
                                            v65 += v65
                                            v65 = (v19 / ((v65) | 1))
                                            _ = v65
                                        case "\n\t":
                                            if (!(!(!(v2 && true)))) {
                                                _ = len(v30)
                                            } else {
                                                fmt.Println(v20)
                                                switch v3 {
                                                case "hello":
                                                    var v66 float64 = v20
                                                    _ = v66
                                                case "foo bar":
                                                    if (v2 && false) {
                                                        _ = len(v30)
                                                        v67 := GenS4{F0: 255}
                                                        _ = v67
                                                    }
                                                default:
                                                    switch fi51 {
                                                    case 128:
                                                        v19 -= v19
                                                    default:
                                                        if (391 == 970) {
                                                            var v68 uint32 = 298
                                                            var v69 string = ("\n\t" + ((v3 + (v3 + v3)) + (v3 + "foo bar")))
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            v70 := GenS5{F0: fi51, F1: 290, F2: 875}
                                                            _ = len(v30)
                                                            switch "\n\t" {
                                                            case "":
                                                                v2 = (2147483648 < 65535)
                                                            case "hello":
                                                                var v71 string = v3
                                                                _ = v71
                                                            default:
                                                                v19 -= v19
                                                            }
                                                            _ = v70
                                                        }
                                                        v72 := GenS2{F0: 42}
                                                        _ = v72
                                                    }
                                                }
                                            }
                                        default:
                                            if v2 {
                                                var v73 uint16 = 2
                                                _ = v73
                                            } else {
                                                if (fi63 <= 835) {
                                                    var v74 int8 = (-(v19))
                                                    var v75 uint32 = 898
                                                    v2 = (42 < 128)
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    if (true || (1e10 != 1.5)) {
                                                        v76 := GenS6{F0: 9223372036854775807, F1: 789}
                                                        v2 = false
                                                        v77 := GenS7{F0: (42 != fi63), F1: v1}
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        var v78 float32 = 0.0
                                                        _ = v78
                                                    }
                                                    fmt.Println(fi63, fi63, v2)
                                                    var v79 float64 = 3.14159
                                                    _ = v79
                                                }
                                                v2 = v2
                                            }
                                        }
                                        var v80 int16 = 0
                                        _ = v80
                                    }
                                    _ = fi63
                                }
                                if (42 != 255) {
                                    for fi81 := 0; fi81 < 12; fi81++ {
                                        fmt.Println(v2, v20)
                                        _ = fi81
                                    }
                                    v30 = append(v30, -516)
                                    if len(v30) >= 2 {
                                        v30 = v30[1:len(v30)]
                                    }
                                }
                            }
                        }
                    default:
                        switch ((v3 + (v3 + v3)) + v3) {
                        case "foo bar":
                            if v2 {
                                switch fi51 {
                                case 42:
                                    if (!false) {
                                        if (int16(fi51) <= -32768) {
                                            if (2 <= 807) {
                                                v30 = append(v30, fi51)
                                            }
                                            fmt.Println(v1, v20, v19)
                                        }
                                    }
                                case 32767:
                                    fmt.Println(v19)
                                default:
                                    if ((((1e10 + (3.14159 / ((1e10 - float32(((v20 + (0.1 - (0.1 - v20))) - v20))) + float32(1.5)))) - 1e10) > ((3.14159 + 1.5) * 3.14159)) && v2) {
                                        if v2 {
                                            v82 := GenS2{F0: -1}
                                            for fi83 := 0; fi83 < 17; fi83++ {
                                                fmt.Println(v3)
                                                v2 = (!false)
                                                _ = fi83
                                            }
                                            _ = v82
                                        } else {
                                            if len(v30) > 0 {
                                                v30[0] = (fi51 ^ (-(fi51)))
                                            }
                                            fmt.Println(fi51, v20, v3)
                                        }
                                        var v84 uint = 635
                                        var v85 uint = v84
                                        _ = v84
                                        _ = v85
                                    }
                                }
                                if len(v30) >= 2 {
                                    v30 = v30[1:len(v30)]
                                }
                            }
                        case "\n\t":
                            switch (fi51 & 426) {
                            case 4294967295:
                                fmt.Println(v20)
                            }
                        }
                    }
                    _ = fi51
                }
                switch (v3 + ((v3 + v3) + v3)) {
                case "":
                    fmt.Println(v2, v2, v19)
                case "\n\t":
                    switch v3 {
                    case "a":
                        for fi86 := 0; fi86 < 7; fi86++ {
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                            fmt.Println(v2, v1, fi86)
                            switch (v3 + v3) {
                            case "a":
                                var v87 int8 = -9
                                _ = v87
                            case "":
                                v19 = v19
                            case "hello":
                                for fi88 := 0; fi88 < 6; fi88++ {
                                    var v89 uint = 992
                                    _ = fi88
                                    _ = v89
                                }
                            }
                            _ = fi86
                        }
                    default:
                        switch v3 {
                        case "":
                            for fi90 := 0; fi90 < 16; fi90++ {
                                v91 := GenS4{F0: 1}
                                var v92 int = 1
                                v93 := GenS4{F0: 624}
                                _ = fi90
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                        case "hello":
                            if false {
                                v20 = v20
                                for fi94 := 0; fi94 < 11; fi94++ {
                                    v2 = true
                                    v30 = append(v30, fi94)
                                    v20 -= ((v20 / v20) + 3.14159)
                                    _ = fi94
                                }
                            }
                        case "a":
                            v30 = append(v30, 9223372036854775807)
                        default:
                            fmt.Println(v3, v3, v1)
                        }
                    }
                default:
                    v19 += 29
                }
                _ = v50
            }
        } else {
            for fi95 := 0; fi95 < 3; fi95++ {
                fmt.Println(v19, v1, v20)
                fmt.Println(v3)
                _ = fi95
            }
            switch ("" + v3) {
            case "hello":
                v2 = (!false)
            case "foo bar":
                var v96 string = "\n\t"
                _ = v96
            default:
                var v97 uint16 = 2
                _ = v97
            }
        }
        v3 = v3
        switch (((v3 + v3) + ("hello" + "foo bar")) + v3) {
        case "":
            v1 += 1
        }
    }
    var v98 uint16 = 32767
    var v99 bool = (!(v19 != -1))
    var v100 int16 = 1000
    v2 = (v100 > v100)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v41
    _ = v98
    _ = v99
}
