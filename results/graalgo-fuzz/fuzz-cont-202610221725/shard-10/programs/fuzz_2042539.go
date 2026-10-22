package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 int16 = 1000
    var v2 bool = true
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    var v6 float64 = 1.5
    if (-61 >= -1) {
        v7 := v3[-600]
        v8 := GenS1{F0: 42}
        var v9 int16 = v1
        _ = v7
        _ = v8
        _ = v9
    } else {
        fmt.Println(v1)
        switch 32768 {
        case 1:
            if (287 != 0) {
                if v2 {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        var v11 uint32 = 42
                        _ = fi10
                        _ = v11
                    }
                }
                var v12 bool = v2
                v13 := []int{127, -223}
                _ = v12
                _ = v13
            }
        case 2147483647:
            v2 = v2
        }
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        switch (("foo bar" + "hello") + "foo bar") {
        case "\n\t":
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := GenS1{F0: 0}
                if v2 {
                    fmt.Println(v2, v6)
                    v17 := GenS1{F0: 252}
                    _ = v17
                }
                v2 = v2
                if (!(!(!(v2 && (127 != 1))))) {
                    if (false && v2) {
                        var v18 int8 = 127
                        var v19 int32 = 1
                        v20 := []int{(fi15 + fi15)}
                        v18 += (v18 << ((150) & 63))
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint64 = 591
                        if (2 <= 42) {
                            v6 -= v6
                            fmt.Println(v21, v1, v1)
                            v22 := GenS1{F0: 255}
                            _ = v22
                        }
                        v2 = v2
                        fmt.Println(v21, v1, fi14)
                        if v2 {
                            v1 = 42
                        } else {
                            var v23 bool = v2
                            _ = v23
                        }
                        _ = v21
                    }
                    if ((9223372036854775807 < 1) && v2) {
                        v1 = -32768
                        v1 = 0
                    }
                    if v2 {
                        switch "foo bar" {
                        case "foo bar":
                            switch (("hello" + (("\n\t" + "") + "")) + "") {
                            case "a":
                                var v24 int16 = v1
                                _ = v24
                            case "hello":
                                v25 := GenS1{F0: 50}
                                _ = v25
                            case "":
                                var v26 uint16 = 11
                                _ = v26
                            default:
                                var v27 float64 = v6
                                _ = v27
                            }
                        }
                        delete(v5, 959)
                    }
                    v28 := []int{32768, (^(fi14))}
                    var v29 int32 = -2147483648
                    v6 -= (1.5 + (((((v6 - v6) + ((v6 / 3.14159) / -2.25)) / (((v6 / v6) / ((1e10 / float64(3.14159)) + (float64(((0.1 - float32(3.14159)) * -2.25)) * (v6 + (float64(0.0) - 3.14159))))) * 3.14159)) - 1e10) * v6))
                    v30 := GenS1{F0: 101}
                    var v31 int64 = 0
                    if (true && true) {
                        v32 := GenS1{F0: 0}
                        v29 -= v29
                        _ = v32
                    }
                    var v33 float32 = (1.5 + 1e10)
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v33
                }
                fmt.Println(fi14, v6, fi14)
                _ = fi15
                _ = v16
            }
        case "hello":
            if v2 {
                if v2 {
                    fmt.Println(v1)
                    v6 = float64(-2.25)
                } else {
                    if (v6 == ((1.5 * v6) + (float64(-2.25) / v6))) {
                        var v34 int32 = -488
                        _ = v34
                    } else {
                        v35 := []int{(fi14 % (((fi14 << ((1) & 63))) | 1))}
                        _ = v35
                    }
                    var v36 int8 = -59
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v1)
                    }
                    _ = v36
                }
                var v37 uint64 = 710
                var v38 uint16 = 1
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 1; fi39++ {
                    fmt.Println(fi14)
                    v40 := []int{(fi39 + fi14), fi39, -193, fi39}
                    _ = fi39
                    _ = v40
                }
            }
        default:
            var v41 uint64 = 553
            _ = v41
        }
        var v42 float64 = v6
        v43 := GenS1{F0: 42}
        _ = fi14
        _ = v42
        _ = v43
    }
    v6 = 3.14159
    v44 := GenS2{F0: 127}
    v45 := GenS1{F0: 21}
    switch 2147483647 {
    case 32768:
        if (!v2) {
            v46 := []int{-1}
            if (v2 && ((!v2) || v2)) {
                v47 := GenS3{F0: "a"}
                v6 = ((-2.25 / (v6 / 1e10)) * v6)
                if (!true) {
                    v2 = ((796 < 767) || ((!v2) || (216 != 719)))
                    for fi48 := 0; fi48 < 5; fi48++ {
                        v49 := GenS4{F0: 42}
                        for fi50 := 0; fi50 < 20; fi50++ {
                            fmt.Println(fi50, fi50)
                            _ = fi50
                        }
                        _ = fi48
                        _ = v49
                    }
                }
                _ = v47
            }
            for fi51 := 0; fi51 < 16; fi51++ {
                v52 := GenS1{F0: 255}
                for fi53 := 0; fi53 < 6; fi53++ {
                    var v54 int64 = 9223372036854775807
                    v2 = ((v2 && v2) && true)
                    var v55 uint32 = 2147483648
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                _ = fi51
                _ = v52
            }
            _ = v46
        } else {
            fmt.Println(v2)
            var v56 bool = false
            _ = v56
        }
    case 32767:
        v57 := GenS3{F0: ""}
        _ = v57
    default:
        var v58 int = 728
        _ = v58
    }
    v6 += ((v6 * float64(0.1)) / float64(3.14159))
    var v59 int = -773
    v2 = v2
    var v60 uint = 446
    fmt.Println(v6)
    v61 := GenS4{F0: 42}
    v62 := GenS5{F0: 846, F1: (v1 * (v1 % ((1) | 1)))}
    if (!v2) {
        v63 := GenS6{F0: 1}
        _ = v63
    } else {
        if false {
            switch "\n\t" {
            case "":
                v64 := GenS1{F0: 1}
                _ = v64
            }
            if (!v2) {
                fmt.Println(v2)
                fmt.Println(v2, v59, v60)
                v6 = 3.14159
            } else {
                v60 -= 810
            }
        } else {
            v60 -= 687
            var v65 int32 = 1
            v2 = (v2 && (true && (!(42 <= 1))))
            _ = v65
        }
        for fi66 := 0; fi66 < 19; fi66++ {
            for fi67 := 0; fi67 < 2; fi67++ {
                v1 = v1
                _ = fi67
            }
            v68 := v5[2147483648]
            _ = fi66
            _ = v68
        }
        if (((v60 >= v60) || v2) && (0.0 <= ((0.1 + -2.25) - 0.0))) {
            v6 = (((v6 + 3.14159) * v6) + 0.1)
        } else {
            switch (v59 + (v59 - (^(v59)))) {
            case 9223372036854775807:
                if v2 {
                    fmt.Println(v6)
                    var v69 int = v59
                    _ = v69
                } else {
                    if true {
                        for fi70 := 0; fi70 < 17; fi70++ {
                            v71 := []int{(-(v59)), fi70, (v59 - (-(v59)))}
                            delete(v3, (v59 ^ fi70))
                            _ = fi70
                            _ = v71
                        }
                        v1 = v1
                    } else {
                        fmt.Println(v59, v59, v60)
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                var v72 int8 = 127
                                var v73 uint16 = 255
                                _ = v72
                                _ = v73
                            }
                        case "\n\t":
                            fmt.Println(v6, v2, v59)
                        case "a":
                            if (v6 != 1e10) {
                                if ((836 >= -2147483648) || v2) {
                                    v2 = false
                                    if v2 {
                                        var v74 uint8 = 0
                                        for fi75 := 0; fi75 < 7; fi75++ {
                                            v76 := []int{(v59 & fi75), 9223372036854775807, fi75}
                                            _ = fi75
                                            _ = v76
                                        }
                                        _ = v74
                                    } else {
                                        if false {
                                            v1 = (v1 | 32767)
                                            if v2 {
                                                v59 = (v59 ^ -924)
                                            } else {
                                                v77 := GenS2{F0: v59}
                                                delete(v5, 639)
                                                _ = v77
                                            }
                                            v78 := []int{(v59 << (((v60 + (^(v60)))) & 63)), (v59 ^ (v59 ^ 28))}
                                            _ = v78
                                        }
                                        v79 := []int{-9223372036854775808, 9223372036854775807}
                                        _ = v79
                                    }
                                    var v80 float32 = 1e10
                                    _ = v80
                                }
                                for fi81 := 0; fi81 < 12; fi81++ {
                                    v1 = v1
                                    if v2 {
                                        if ("a" < ("hello" + "hello")) {
                                            var v82 float32 = (0.0 + float32(float64(0.0)))
                                            v1 += (^(v1))
                                            for fi83 := 0; fi83 < 12; fi83++ {
                                                switch (fi83 << ((v60) & 63)) {
                                                case 32768:
                                                    fmt.Println(fi81, fi81)
                                                case -212:
                                                    v84 := []int{fi83, (v59 ^ (^(v59))), (fi81 + (^(fi81))), v59}
                                                    _ = v84
                                                case -369:
                                                    v2 = v2
                                                default:
                                                    fmt.Println(fi83)
                                                }
                                                _ = fi83
                                            }
                                            _ = v82
                                        }
                                    } else {
                                        v85 := []int{(fi81 ^ (-(v59))), v59, 728, (fi81 << ((v60) & 63))}
                                        var v86 int = (v59 % (((^(fi81))) | 1))
                                        _ = v85
                                        _ = v86
                                    }
                                    switch -929 {
                                    case 1:
                                        v87 := v4[42]
                                        _ = v87
                                    }
                                    _ = fi81
                                }
                            }
                        default:
                            v6 = ((float64(((-2.25 + 1.5) - 0.1)) - (3.14159 - v6)) * (0.1 / 1.0))
                        }
                    }
                }
            default:
                delete(v3, (v59 >> (((v60 >> ((138) & 63))) & 63)))
            }
            v6 += v6
            v88 := []int{(v59 << (((v60 & v60)) & 63)), (v59 & (-(v59)))}
            _ = v88
        }
    }
    v89 := GenS5{F0: 356, F1: int16(v59)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v89
}
