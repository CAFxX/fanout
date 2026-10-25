package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 uint64; F2 int32 }
type GenS8 struct { F0 bool; F1 bool; F2 uint }
type GenS9 struct { F0 bool }
func genFn0(p0 uint, p1 uint32) bool {
    v7 := []int{}
    p0 = p0
    p1 += 2147483647
    _ = p0
    _ = p1
    _ = v7
    return false
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(fi8 / ((fi8) | 1)), (fi8 >> ((948) & 63)), 127}
        switch 9223372036854775807 {
        case 32767:
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                if ("foo bar" != ("a" + "\n\t")) {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        var v11 int32 = 42
                        v11 = (v11 ^ v11)
                        _ = fi10
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 int32 = -2147483648
                        v13 += (v13 >> ((18446744073709551615) & 63))
                        var v14 uint32 = 42
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                } else {
                    if (true && false) {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            v16 := GenS1{F0: 2}
                            fmt.Println(fi8, fi8)
                            _ = fi15
                            _ = v16
                        }
                        fmt.Println(fi8)
                    } else {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            _ = len(v9)
                            _ = fi17
                        }
                    }
                    var v18 uint16 = 1
                    v9 = append(v9, fi8)
                    _ = v18
                }
            }
        case 2147483647:
            for fi19 := 0; fi19 < 13; fi19++ {
                if (!(227 == 65535)) {
                    var v20 uint16 = 0
                    if false {
                        var v21 int = fi8
                        _ = v21
                    }
                    if (-1 == -1) {
                        if (18446744073709551615 > 4294967295) {
                            v22 := GenS1{F0: 42}
                            _ = v22
                        } else {
                            var v23 float32 = 1.5
                            v23 -= ((v23 / v23) - (1.5 * (((v23 * v23) - (v23 + 1e10)) - -2.25)))
                            if ("hello" <= "hello") {
                                if false {
                                    switch int(v20) {
                                    case 9223372036854775807:
                                        if (!(!true)) {
                                            v23 += 0.0
                                        }
                                    default:
                                        var v24 float64 = 1e10
                                        _ = v24
                                    }
                                } else {
                                    v23 -= 0.0
                                }
                                v25 := GenS2{F0: (v20 / ((322) | 1)), F1: 656, F2: 970}
                                v26 := GenS2{F0: 32767, F1: 1, F2: -32768}
                                fmt.Println(fi19)
                                _ = v25
                                _ = v26
                            }
                            _ = v23
                        }
                        fmt.Println(fi19, fi19)
                        v20 -= v20
                        var v27 int64 = 42
                        _ = v27
                    }
                    _ = v20
                } else {
                    var v28 string = "hello"
                    switch v28 {
                    case "hello":
                        _ = len(v9)
                    case "\n\t":
                        if true {
                            _ = len(v9)
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v29 float32 = 0.1
                            _ = v29
                        }
                    case "a":
                        _ = len(v9)
                    default:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    var v30 int16 = 0
                    _ = v28
                    _ = v30
                }
                switch "foo bar" {
                case "hello":
                    if ((1.5 * 0.1) >= 0.0) {
                        v31 := GenS2{F0: 42, F1: 0, F2: -963}
                        fmt.Println(fi8, fi8)
                        var v32 uint = 447
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS2{F0: 255, F1: 537, F2: 1}
                    var v34 int16 = 42
                    _ = v33
                    _ = v34
                case "a":
                    fmt.Println(fi8)
                case "\n\t":
                    var v35 int64 = 42
                    _ = v35
                default:
                    var v36 uint32 = 0
                    _ = v36
                }
                fmt.Println(fi19)
                _ = fi19
            }
        case 42:
            var v37 uint32 = 42
            _ = v37
        }
        for fi38 := 0; fi38 < 4; fi38++ {
            var v39 int8 = -1
            _ = fi38
            _ = v39
        }
        _ = fi8
        _ = v9
    }
    var v40 float64 = (((1e10 * 3.14159) - 1e10) + ((((0.1 * (3.14159 / 0.1)) * 1.5) / 1e10) * 0.0))
    _ = v40
    return 42
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    var v3 float64 = (0.1 / 1.0)
    var v4 bool = false
    v5 := make(map[uint32]string)
    var v6 string = ("foo bar" + ("\n\t" + "hello"))
    v1 = 1000
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 7; fi42++ {
            var v43 uint64 = 291
            if v4 {
                v44 := []int{-369, (fi41 * fi41)}
                _ = v44
            }
            var v45 int32 = -2147483648
            _ = fi42
            _ = v43
            _ = v45
        }
        _ = fi41
    }
    if (2 >= 365) {
        var v46 int64 = -610
        _ = v46
    }
    if false {
        v3 += (v3 * 0.0)
        fmt.Println(v6)
        v4 = v4
    } else {
        for fi47 := 0; fi47 < 4; fi47++ {
            v1 = v1
            _ = fi47
        }
        var v48 string = ""
        _ = v48
    }
    var v49 bool = v4
    var v50 uint32 = 98
    switch 42 {
    case 42:
        var v51 uint8 = 111
        _ = v51
    default:
        if (!false) {
            if (v49 || (!(-128 < 0))) {
                v1 += v1
            }
            for fi52 := 0; fi52 < 19; fi52++ {
                v6 = v6
                if (!true) {
                    v53 := GenS3{F0: 985}
                    _ = v53
                }
                _ = fi52
            }
        } else {
            v54 := GenS4{F0: 2147483648}
            _ = v54
        }
    }
    v55 := GenS5{F0: "foo bar", F1: 42}
    v3 = float64(0.0)
    fmt.Println(v6, v4, v50)
    fmt.Println(v49, v6, v50)
    v1 -= v1
    var v56 float64 = v3
    if (v6 <= v6) {
        var v57 uint16 = 255
        v58 := []int{42}
        _ = v57
        _ = v58
    }
    switch (v6 + v6) {
    case "":
        for fi59 := 0; fi59 < 19; fi59++ {
            for fi60 := 0; fi60 < 13; fi60++ {
                v61 := GenS4{F0: (v50 ^ 161)}
                fmt.Println(fi60, v1, v49)
                var v62 uint = 2
                _ = fi60
                _ = v61
                _ = v62
            }
            switch (fi59 / ((fi59) | 1)) {
            case 32768:
                switch fi59 {
                case -2147483648:
                    v6 = v6
                }
            case 32767:
                fmt.Println(v6, v1)
            }
            _ = fi59
        }
    default:
        if (561 >= 774) {
            v49 = (v6 == v6)
            var v63 int64 = -509
            for fi64 := 0; fi64 < 19; fi64++ {
                if (v4 && (v50 == 284)) {
                    for fi65 := 0; fi65 < 5; fi65++ {
                        if ("a" == ((v6 + v6) + "")) {
                            var v66 uint32 = v50
                            v67 := GenS3{F0: 2147483647}
                            if v4 {
                                var v68 int64 = 261
                                _ = v68
                            }
                            _ = v66
                            _ = v67
                        }
                        switch fi64 {
                        case 32767:
                            switch (v6 + ("hello" + "hello")) {
                            case "hello":
                                v56 = (1e10 * 1e10)
                            case "":
                                if v49 {
                                    if v49 {
                                        if ((v63 >= (v63 / ((v63) | 1))) || (!(true || v4))) {
                                            fmt.Println(v63)
                                            fmt.Println(v6, v6, v49)
                                            v69 := GenS3{F0: (fi64 / ((127) | 1))}
                                            if (v50 != 65535) {
                                                var v70 uint = 18446744073709551615
                                                if (1e10 <= 1e10) {
                                                    switch fi65 {
                                                    case 42:
                                                        fmt.Println(v70, v50)
                                                    case 127:
                                                        if v4 {
                                                            v2[v50] = (^(v50))
                                                        } else {
                                                            v63 = 1
                                                            var v71 string = v6
                                                            _ = v71
                                                        }
                                                    default:
                                                        switch fi64 {
                                                        case -1:
                                                            v56 -= ((v3 - (3.14159 + -2.25)) * ((v3 * 1e10) / v3))
                                                        case 624:
                                                            var v72 int32 = 1
                                                            _ = v72
                                                        case 322:
                                                            if v4 {
                                                                v73 := GenS1{F0: 740}
                                                                v74 := []int{int(v63), (fi64 ^ (fi64 % ((fi65) | 1))), (fi64 - fi65)}
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                    }
                                                    var v75 uint = 275
                                                    _ = v75
                                                }
                                                _ = v70
                                            }
                                            var v76 float32 = 1e10
                                            _ = v69
                                            _ = v76
                                        }
                                    } else {
                                        var v77 float32 = ((1.5 * 0.1) * float32((v3 / (0.1 - 1e10))))
                                        _ = v77
                                    }
                                    v78 := v5[(v50 % ((v50) | 1))]
                                    var v79 uint8 = genFn1()
                                    _ = v78
                                    _ = v79
                                } else {
                                    v49 = (!false)
                                    v80 := GenS1{F0: 1}
                                    var v81 int64 = -720
                                    v3 -= 1e10
                                    switch v6 {
                                    case "a":
                                        fmt.Println(v56)
                                    case "foo bar":
                                        var v82 uint = 185
                                        _ = v82
                                    case "\n\t":
                                        var v83 float64 = ((0.1 - 0.0) + (((0.0 + v3) / 1e10) - float64((0.0 * 0.1))))
                                        _ = v83
                                    default:
                                        v84 := GenS6{F0: 4294967295}
                                        _ = v84
                                    }
                                    var v85 int8 = 42
                                    fmt.Println(v85, v63)
                                    if true {
                                        v4 = (42 != 127)
                                        delete(v5, (v50 * v50))
                                        _ = len(v2)
                                    } else {
                                        switch v6 {
                                        case "":
                                            if v49 {
                                                v4 = true
                                                v49 = (((!(0 > genFn1())) && true) || ("foo bar" < ("" + v6)))
                                                var v86 uint16 = uint16(v50)
                                                _ = v86
                                            } else {
                                                var v87 bool = false
                                                v88 := GenS1{F0: 463}
                                                var v89 int16 = v1
                                                _ = v87
                                                _ = v88
                                                _ = v89
                                            }
                                        case "\n\t":
                                            if false {
                                                v85 = 0
                                                var v90 uint64 = 0
                                                _ = v90
                                            } else {
                                                v91 := []int{}
                                                _ = v91
                                            }
                                        case "hello":
                                            var v92 int32 = -1
                                            _ = v92
                                        }
                                        var v93 int64 = v63
                                        _ = v93
                                    }
                                    _ = v80
                                    _ = v81
                                    _ = v85
                                }
                            case "\n\t":
                                v94 := GenS5{F0: v6, F1: 136}
                                v56 = (float64((((0.0 + 3.14159) - (1.5 - (1e10 + 0.1))) * -2.25)) + (-2.25 + (1e10 - v3)))
                                _ = v94
                            }
                        }
                        if (!(!true)) {
                            v1 = v1
                            v95 := GenS7{F0: (v63 >> ((0) & 63)), F1: 1, F2: 42}
                            _ = v95
                        }
                        _ = fi65
                    }
                } else {
                    v96 := GenS8{F0: (int16(v50) >= v1), F1: v4, F2: 18446744073709551615}
                    _ = v96
                }
                _ = fi64
            }
            _ = v63
        }
    }
    v97 := GenS9{F0: (genFn1() >= 2)}
    fmt.Println(v4, v1, v4)
    var v98 string = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v55
    _ = v56
    _ = v6
    _ = v97
    _ = v98
}
