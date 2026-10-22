package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool; F1 int8; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint16; F1 int8; F2 int8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int32; F1 int; F2 uint16 }
type GenS14 struct { F0 string; F1 int64; F2 bool }
type GenS15 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint) int8 {
    v6 := []int{}
    p1 -= (p1 - (p1 << ((2) & 63)))
    p0 = p0
    if (!false) {
        for fi7 := 0; fi7 < 12; fi7++ {
            p1 = p1
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 100
}

func main() {
    var v1 bool = true
    var v2 int64 = 42
    var v3 int32 = 2147483647
    var v4 uint16 = 467
    var v5 int16 = 38
    v2 = (v2 / (((v2 ^ (^(v2)))) | 1))
    v8 := []int{}
    fmt.Println(v1, v2)
    v9 := GenS1{F0: 18446744073709551615, F1: true, F2: (0.0 <= float32(-2.25))}
    v1 = (v2 < (v2 | (^(v2))))
    v10 := GenS2{F0: 0, F1: "\n\t"}
    v4 = (v4 & (^(v4)))
    var v11 float32 = float32((((-2.25 / 1.0) - 0.0) - (-2.25 + 0.0)))
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch "a" {
    case "\n\t":
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v3)
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 float32 = (float32(0.0) + (v11 + (((v11 * float32(0.1)) + (v11 + ((float32((1.5 - 0.1)) + v11) + (1.5 - float32(-2.25))))) / (1.5 / (0.1 + 1e10)))))
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                fmt.Println(v11, v2, fi13)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    case "foo bar":
        var v15 int16 = v5
        _ = v15
    case "a":
        v8 = append(v8, -674)
    }
    fmt.Println(v2, v4, v4)
    if (v1 || (v1 || v1)) {
        if true {
            switch 0 {
            case 358:
                v1 = v1
            case -771:
                v16 := GenS3{F0: true, F1: genFn0(128, 255), F2: -1}
                _ = v16
            case 4294967295:
                var v17 uint = 255
                _ = v17
            }
        } else {
            if v1 {
                if (0 != 9223372036854775807) {
                    v18 := GenS2{F0: 365, F1: "foo bar"}
                    v19 := GenS2{F0: (v4 % ((65535) | 1)), F1: "foo bar"}
                    if true {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            for fi21 := 0; fi21 < 7; fi21++ {
                                var v22 uint64 = 2
                                v2 = (v2 & (v2 >> ((uint(v4)) & 63)))
                                switch "foo bar" {
                                case "a":
                                    var v23 uint8 = 2
                                    _ = v23
                                case "foo bar":
                                    var v24 uint32 = 613
                                    _ = v24
                                case "":
                                    v25 := GenS2{F0: uint16(v22), F1: ""}
                                    _ = v25
                                }
                                _ = fi21
                                _ = v22
                            }
                            if (true && false) {
                                var v26 int32 = (v3 % (((v3 ^ (v3 / (((v3 - (^(v3)))) | 1)))) | 1))
                                _ = v26
                            }
                            _ = fi20
                        }
                        if true {
                            if v1 {
                                if len(v8) > 0 {
                                    v8[0] = -485
                                }
                                v3 = (v3 >> ((4294967295) & 63))
                            }
                            if v1 {
                                _ = len(v8)
                                switch "foo bar" {
                                case "a":
                                    v4 = (v4 % ((255) | 1))
                                case "foo bar":
                                    _ = len(v8)
                                }
                            } else {
                                if (2 > 2) {
                                    if (2 != 1) {
                                        v2 = (^(v2))
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        fmt.Println(v5, v11, v2)
                                    }
                                } else {
                                    fmt.Println(v11, v11, v1)
                                    v3 -= (v3 << ((42) & 63))
                                }
                            }
                        } else {
                            var v27 int = -998
                            v4 += v4
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 12; fi28++ {
                            var v29 uint8 = 2
                            if (v29 < (v29 * (v29 * (v29 & (v29 * v29))))) {
                                if len(v8) > 0 {
                                    v8[0] = 987
                                }
                                if (989 != int32(fi28)) {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    v30 := GenS4{F0: -128}
                                    var v31 float64 = -2.25
                                    _ = v30
                                    _ = v31
                                } else {
                                    v11 -= 0.0
                                    for fi32 := 0; fi32 < 15; fi32++ {
                                        var v33 int16 = v5
                                        fmt.Println(fi28, fi32, v4)
                                        v3 = v3
                                        _ = fi32
                                        _ = v33
                                    }
                                    var v34 int = -2147483648
                                    _ = v34
                                }
                            }
                            var v35 uint32 = 909
                            _ = fi28
                            _ = v29
                            _ = v35
                        }
                    }
                    _ = v18
                    _ = v19
                } else {
                    if v1 {
                        var v36 int32 = -1
                        for fi37 := 0; fi37 < 2; fi37++ {
                            var v38 uint16 = 862
                            switch ("a" + "a") {
                            case "hello":
                                v36 -= 1
                            case "foo bar":
                                var v39 int16 = int16(v4)
                                _ = v39
                            case "\n\t":
                                for fi40 := 0; fi40 < 10; fi40++ {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    var v41 bool = false
                                    _ = fi40
                                    _ = v41
                                }
                            default:
                                v2 = v2
                            }
                            switch "a" {
                            case "hello":
                                for fi42 := 0; fi42 < 16; fi42++ {
                                    v4 = 445
                                    v3 -= v36
                                    var v43 string = "foo bar"
                                    switch (fi37 / (((fi37 ^ (fi42 - fi42))) | 1)) {
                                    case 414:
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        if v1 {
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            var v44 uint16 = (v38 | 255)
                                            fmt.Println(v2)
                                            _ = v44
                                        } else {
                                            var v45 string = v43
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            if v1 {
                                                v46 := GenS2{F0: v4, F1: ("\n\t" + (v45 + v43))}
                                                _ = v46
                                            }
                                            if false {
                                                if (((!(0 > 2)) || false) || v1) {
                                                    var v47 int = fi37
                                                    v2 = (v2 << ((2) & 63))
                                                    _ = v47
                                                }
                                            } else {
                                                var v48 float32 = 1e10
                                                var v49 int = (fi42 * fi37)
                                                fmt.Println(v43, v36)
                                                var v50 uint = 0
                                                if v1 {
                                                    v51 := GenS2{F0: 2, F1: ("\n\t" + v43)}
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    _ = v51
                                                }
                                                _ = len(v8)
                                                if ((18446744073709551615 != 4294967295) || v1) {
                                                    fmt.Println(v3, v43)
                                                    var v52 bool = v1
                                                    _ = v52
                                                }
                                                if ((!(v50 < 255)) || ((!(!(970 != 18446744073709551615))) || ((0 == 77) || ((v11 + 3.14159) < ((float32(-2.25) - (v11 - (v48 / v48))) / 0.1))))) {
                                                    v48 += (0.1 * float32(0.1))
                                                    fmt.Println(v49, v50)
                                                    var v53 uint = uint(v4)
                                                    _ = v53
                                                } else {
                                                    v54 := GenS5{F0: 255, F1: uint(v4)}
                                                    fmt.Println(v50, v48, v5)
                                                    fmt.Println(v11, v3, v38)
                                                    switch int(v3) {
                                                    case 6:
                                                        var v55 uint16 = (v38 & 0)
                                                        _ = v55
                                                    default:
                                                        if len(v8) > 0 {
                                                            v8[0] = fi37
                                                        }
                                                    }
                                                    if true {
                                                        v56 := GenS2{F0: 2, F1: v45}
                                                        v48 -= float32((((1.5 + 1e10) - (((1e10 * 3.14159) * float64(((v48 / 1.0) - float32(1.5)))) * (1e10 + 1e10))) / 1e10))
                                                        v57 := GenS6{F0: v3}
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    var v58 uint32 = 852
                                                    v59 := GenS4{F0: 1}
                                                    _ = v54
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                            var v60 int16 = v5
                                            fmt.Println(v36, v38, v1)
                                            v61 := GenS1{F0: 861, F1: (!true), F2: v1}
                                            v62 := GenS7{F0: 236, F1: 100, F2: int8(v38)}
                                            _ = v45
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                    case 4294967295:
                                        v38 = (^(v4))
                                    }
                                    v8 = append(v8, fi37)
                                    _ = fi42
                                    _ = v43
                                }
                            case "\n\t":
                                v5 = 42
                            case "a":
                                if false {
                                    if v1 {
                                        if len(v8) > 0 {
                                            v8[0] = fi37
                                        }
                                        v63 := GenS8{F0: (v4 - v38)}
                                        _ = v63
                                    } else {
                                        v4 -= v38
                                    }
                                    var v64 int = (fi37 << ((15) & 63))
                                    _ = v64
                                }
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = v36
                    }
                    for fi65 := 0; fi65 < 1; fi65++ {
                        v66 := GenS9{F0: v5}
                        for fi67 := 0; fi67 < 9; fi67++ {
                            if v1 {
                                v1 = v1
                                v68 := GenS10{F0: v4, F1: 4294967295}
                                var v69 int = -9223372036854775808
                                _ = v68
                                _ = v69
                            } else {
                                if (!(("hello" == "\n\t") || (!(v4 >= (v4 & 32767))))) {
                                    if true {
                                        _ = len(v8)
                                        v70 := GenS10{F0: (v4 << ((311) & 63)), F1: 255}
                                        v1 = v1
                                        _ = v70
                                    } else {
                                        var v71 float64 = (-2.25 - 1.5)
                                        _ = v71
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        v3 = (v3 * -2147483648)
                                    case "foo bar":
                                        var v72 int64 = (v2 | (v2 / (((v2 - v2)) | 1)))
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        _ = v72
                                    default:
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    }
                                }
                            }
                            _ = fi67
                        }
                        var v73 int8 = -1
                        _ = fi65
                        _ = v66
                        _ = v73
                    }
                    v3 -= (v3 | -52)
                }
                var v74 uint8 = 0
                _ = v74
            }
            v11 -= (float32(-2.25) * v11)
            v75 := GenS11{F0: 255, F1: 6, F2: -35}
            _ = v75
        }
        var v76 float64 = 1e10
        _ = v76
    }
    var v77 bool = (float32((-2.25 - (0.1 * float64(3.14159)))) != 1e10)
    v77 = (0 >= 2)
    v2 += int64(v5)
    fmt.Println(v2, v77)
    if (v3 == 2147483647) {
        var v78 uint16 = v4
        var v79 uint8 = 0
        v77 = v1
        _ = v78
        _ = v79
    } else {
        v3 += v3
        for fi80 := 0; fi80 < 10; fi80++ {
            v81 := GenS2{F0: (v4 / ((v4) | 1)), F1: ("a" + "\n\t")}
            switch 32768 {
            case -2147483648:
                v8 = append(v8, fi80)
            case 9223372036854775807:
                for fi82 := 0; fi82 < 3; fi82++ {
                    if (("foo bar" + "foo bar") >= "foo bar") {
                        switch "" {
                        case "foo bar":
                            if false {
                                switch fi80 {
                                case 411:
                                    fmt.Println(fi80, v1)
                                case 661:
                                    switch (-(fi82)) {
                                    case 774:
                                        v2 = (v2 + (v2 / (((^(v2))) | 1)))
                                    default:
                                        var v83 uint = 308
                                        _ = v83
                                    }
                                }
                                v5 = (v5 / ((-1) | 1))
                                if v1 {
                                    v8 = append(v8, 9223372036854775807)
                                    switch fi82 {
                                    case 294:
                                        if v77 {
                                            v84 := GenS10{F0: v4, F1: 482}
                                            v2 -= (v2 & v2)
                                            _ = v84
                                        }
                                    default:
                                        fmt.Println(fi80, fi80)
                                    }
                                    var v85 int = (fi80 ^ (fi80 & fi80))
                                    _ = v85
                                } else {
                                    v86 := GenS12{F0: (v3 ^ v3)}
                                    if v77 {
                                        fmt.Println(v3)
                                    } else {
                                        switch fi80 {
                                        case 32768:
                                            var v87 bool = v77
                                            _ = v87
                                        case -2147483648:
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        }
                                        if len(v8) > 0 {
                                            v8[0] = fi80
                                        }
                                        v88 := GenS7{F0: 1, F1: 2, F2: genFn0(1, 18446744073709551615)}
                                        _ = v88
                                    }
                                    _ = v86
                                }
                                v8 = append(v8, fi80)
                            }
                        }
                        v3 += (v3 * v3)
                        v77 = ("" != "a")
                        v1 = ((0 >= uint8(v4)) && (0.1 != (-2.25 * -2.25)))
                    } else {
                        if v1 {
                            var v89 int64 = 1
                            switch ((("foo bar" + (("foo bar" + "hello") + "")) + "hello") + "\n\t") {
                            case "\n\t":
                                switch "hello" {
                                case "":
                                    var v90 uint = 1
                                    if v1 {
                                        if v77 {
                                            v77 = v77
                                            v91 := GenS13{F0: -1, F1: 32768, F2: v4}
                                            _ = v91
                                        } else {
                                            v1 = true
                                            if (v77 || v1) {
                                                switch fi80 {
                                                case -212:
                                                    v11 = -2.25
                                                case 615:
                                                    v92 := GenS4{F0: 92}
                                                    _ = v92
                                                case 128:
                                                    v1 = true
                                                    var v93 int64 = v2
                                                    _ = v93
                                                default:
                                                    var v94 int8 = 99
                                                    _ = v94
                                                }
                                                if (true && v1) {
                                                    var v95 int8 = int8(v2)
                                                    var v96 int = fi80
                                                    _ = len(v8)
                                                    _ = v95
                                                    _ = v96
                                                }
                                                v89 = v89
                                                if ("hello" > ("" + "hello")) {
                                                    v97 := GenS7{F0: 4294967295, F1: 255, F2: 100}
                                                    _ = v97
                                                }
                                                fmt.Println(v5, v89, v90)
                                            } else {
                                                if (((0.0 * 1.5) * 1e10) < 3.14159) {
                                                    fmt.Println(v90, v3)
                                                    v8 = append(v8, 915)
                                                } else {
                                                    v98 := GenS14{F0: "", F1: int64(fi80), F2: (-128 > -72)}
                                                    _ = v98
                                                }
                                                v5 = (v5 / (((^(v5))) | 1))
                                                v89 = 42
                                                var v99 float64 = 0.0
                                                _ = len(v8)
                                                v8 = append(v8, fi80)
                                                var v100 int = (fi80 >> (((v90 >> ((v90) & 63))) & 63))
                                                _ = v100
                                                _ = v99
                                            }
                                        }
                                        v101 := GenS10{F0: v4, F1: (v90 / ((v90) | 1))}
                                        if (v5 <= v5) {
                                            v102 := GenS12{F0: v3}
                                            _ = v102
                                        } else {
                                            v89 -= v89
                                            fmt.Println(v89, v90, v89)
                                            v11 += v11
                                        }
                                        _ = v101
                                    } else {
                                        if (!false) {
                                            v103 := GenS15{F0: (v4 % (((^(v4))) | 1))}
                                            fmt.Println(v2)
                                            _ = v103
                                        }
                                    }
                                    v77 = (v77 || (true || (("\n\t" <= "hello") || (uint16(v90) < v4))))
                                    _ = v90
                                }
                            default:
                                if (1.5 != 0.1) {
                                    v2 -= 688
                                    var v104 int = (^(fi82))
                                    _ = v104
                                }
                            }
                            _ = v89
                        }
                        var v105 bool = false
                        v11 = v11
                        _ = v105
                    }
                    _ = fi82
                }
            default:
                v4 -= v4
            }
            _ = fi80
            _ = v81
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v8
    _ = v9
}
