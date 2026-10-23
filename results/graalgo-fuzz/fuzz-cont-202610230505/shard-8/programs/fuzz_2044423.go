package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 int8; F2 int }
func genFn0() int {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint64 = 42
        if false {
            v9 := GenS1{F0: 42, F1: -1, F2: "foo bar"}
            v10 := GenS2{F0: 57, F1: false}
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: 42, F1: (!false)}
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS3{F0: 1}
    switch 0 {
    case 4294967295:
        var v13 int8 = 0
        _ = v13
    case 127:
        switch "\n\t" {
        case "hello":
            var v14 bool = (false || (9223372036854775807 < 459))
            _ = v14
        case "":
            var v15 int16 = 42
            _ = v15
        case "foo bar":
            v16 := []int{746}
            _ = v16
        }
    case 9223372036854775807:
        var v17 bool = false
        _ = v17
    default:
        switch -2147483648 {
        case 207:
            if (!true) {
                var v18 uint64 = 0
                switch "a" {
                case "a":
                    fmt.Println(v18)
                }
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 bool = (-2147483648 == 956)
                    var v21 int64 = -25
                    switch "a" {
                    case "a":
                        v20 = false
                    case "hello":
                        var v22 float64 = -2.25
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
        }
    }
    v23 := GenS4{F0: -32768}
    var v24 uint32 = 42
    _ = v12
    _ = v23
    _ = v24
    return -476
}
func genFn1(p0 bool) uint8 {
    if ("hello" >= "a") {
        v25 := []int{genFn0(), 42, 784, 42}
        var v26 int8 = -1
        _ = v25
        _ = v26
    }
    if (145 >= 255) {
        p0 = (true || (127 < 5))
    } else {
        v27 := []int{0, genFn0(), -877}
        _ = v27
    }
    p0 = (!(false && p0))
    _ = p0
    return 1
}

func main() {
    var v1 int8 = 0
    var v2 float64 = 3.14159
    var v3 uint16 = 65535
    v4 := make(map[uint32]int)
    var v5 int8 = (v1 % (((v1 | (v1 - v1))) | 1))
    var v6 uint64 = 42
    if true {
        var v28 bool = (false || false)
        v29 := v4[2147483648]
        if true {
            v3 = 212
            v30 := GenS1{F0: 927, F1: v1, F2: ((("hello" + ("a" + ("\n\t" + (("\n\t" + (("foo bar" + ("hello" + "")) + "")) + (("a" + ("\n\t" + "a")) + "hello"))))) + "foo bar") + "a")}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 int8 = (v1 - v5)
                _ = fi31
                _ = v32
            }
            v33 := GenS5{F0: "hello", F1: 950}
            if v28 {
                v3 -= (^(v3))
            }
            _ = v33
        }
        _ = v28
        _ = v29
    }
    var v34 uint8 = 134
    v2 = float64((1.5 - 3.14159))
    switch 9223372036854775807 {
    case 2147483647:
        fmt.Println(v6)
    case 128:
        v35 := GenS3{F0: 2147483648}
        _ = v35
    case 32768:
        var v36 uint = 2
        _ = v36
    default:
        v37 := []int{}
        _ = v37
    }
    var v38 int64 = -9223372036854775808
    if false {
        _ = len(v4)
    } else {
        fmt.Println(v5)
        if (false && true) {
            v3 += uint16(v34)
        } else {
            v39 := GenS4{F0: 9}
            switch "\n\t" {
            case "hello":
                for fi40 := 0; fi40 < 1; fi40++ {
                    switch (((("hello" + (("hello" + ("\n\t" + ((("" + "\n\t") + (("a" + "") + "hello")) + "a"))) + ("a" + ("\n\t" + ("foo bar" + "hello"))))) + ("hello" + ("foo bar" + ""))) + ("hello" + "")) + "") {
                    case "foo bar":
                        for fi41 := 0; fi41 < 4; fi41++ {
                            var v42 float32 = (0.1 - (0.0 + 3.14159))
                            fmt.Println(v1, v6, v38)
                            v6 -= v6
                            _ = fi41
                            _ = v42
                        }
                    case "\n\t":
                        switch fi40 {
                        case 4294967295:
                            if true {
                                var v43 int32 = -807
                                switch (("foo bar" + "") + "") {
                                case "":
                                    v44 := []int{fi40, (fi40 * (fi40 / (((fi40 + (^(fi40)))) | 1))), 128, (-(fi40))}
                                    _ = v44
                                }
                                _ = v43
                            } else {
                                var v45 int8 = v1
                                v3 = v3
                                _ = v45
                            }
                        case -646:
                            v46 := GenS6{F0: (fi40 & fi40)}
                            _ = v46
                        case 9223372036854775807:
                            switch -2147483648 {
                            case 0:
                                var v47 uint32 = 2147483648
                                _ = v47
                            default:
                                v48 := []int{(fi40 << ((18446744073709551615) & 63))}
                                _ = v48
                            }
                        }
                    }
                    var v49 int64 = 9223372036854775807
                    _ = fi40
                    _ = v49
                }
            case "a":
                v50 := []int{}
                _ = v50
            default:
                if true {
                    v51 := GenS7{F0: 169, F1: 0}
                    _ = v51
                } else {
                    if (!true) {
                        for fi52 := 0; fi52 < 11; fi52++ {
                            if true {
                                for fi53 := 0; fi53 < 10; fi53++ {
                                    v38 += (v38 * -9223372036854775808)
                                    fmt.Println(v3)
                                    _ = fi53
                                }
                                var v54 int64 = 360
                                if false {
                                    for fi55 := 0; fi55 < 11; fi55++ {
                                        fmt.Println(v6)
                                        _ = fi55
                                    }
                                    v56 := GenS7{F0: (v6 * v6), F1: 255}
                                    delete(v4, uint32(v3))
                                    _ = v56
                                } else {
                                    v57 := GenS6{F0: (fi52 >> ((298) & 63))}
                                    _ = v57
                                }
                                _ = v54
                            }
                            var v58 uint8 = (v34 + (v34 << ((742) & 63)))
                            v59 := v4[2147483648]
                            _ = fi52
                            _ = v58
                            _ = v59
                        }
                        delete(v4, 0)
                    }
                }
            }
            _ = v39
        }
        fmt.Println(v34, v38)
    }
    var v60 uint64 = v6
    switch ("" + "foo bar") {
    case "":
        switch -2147483648 {
        case 1:
            v61 := []int{-214, 4294967295, 42, 2147483647}
            _ = v61
        case 130:
            if false {
                v38 += int64(v1)
                var v62 uint32 = 2147483648
                v60 = v60
                _ = v62
            } else {
                var v63 uint32 = uint32(v34)
                _ = v63
            }
        case -1:
            v64 := GenS8{F0: "hello", F1: (v1 - (v1 ^ 113))}
            _ = v64
        }
    case "a":
        v6 += v6
    case "hello":
        var v65 uint16 = (v3 >> ((0) & 63))
        _ = v65
    default:
        var v66 int16 = 1000
        _ = v66
    }
    fmt.Println(v60, v60, v3)
    switch "a" {
    case "foo bar":
        var v67 uint8 = (v34 >> ((4294967295) & 63))
        _ = v67
    case "a":
        if (2 <= 255) {
            v34 -= (v34 % (((v34 - (v34 * v34))) | 1))
            if (!("" <= "")) {
                var v68 float64 = 1.5
                v5 = v5
                switch ("a" + "") {
                case "":
                    var v69 uint8 = (v34 << ((65535) & 63))
                    _ = v69
                case "\n\t":
                    v70 := GenS9{F0: (v34 - (^(v34)))}
                    _ = v70
                case "hello":
                    for fi71 := 0; fi71 < 4; fi71++ {
                        var v72 int32 = 1
                        for fi73 := 0; fi73 < 6; fi73++ {
                            v74 := []int{genFn0()}
                            if (false || true) {
                                var v75 int8 = v5
                                _ = v75
                            } else {
                                v76 := GenS1{F0: -32768, F1: -24, F2: ""}
                                _ = v76
                            }
                            var v77 int16 = 0
                            _ = fi73
                            _ = v74
                            _ = v77
                        }
                        fmt.Println(v34, v6, v72)
                        _ = fi71
                        _ = v72
                    }
                }
                _ = v68
            } else {
                for fi78 := 0; fi78 < 14; fi78++ {
                    if false {
                        v5 -= 100
                        var v79 int32 = 0
                        for fi80 := 0; fi80 < 12; fi80++ {
                            var v81 int16 = 42
                            v81 = v81
                            v1 -= v5
                            _ = fi80
                            _ = v81
                        }
                        _ = v79
                    } else {
                        v82 := []int{}
                        _ = v82
                    }
                    if ((v60 ^ 176) > (v6 << ((65535) & 63))) {
                        switch fi78 {
                        case 268:
                            switch fi78 {
                            case -1:
                                v83 := []int{(^(fi78))}
                                _ = v83
                            case 32767:
                                fmt.Println(v5, v6)
                            default:
                                v84 := []int{fi78, (fi78 % ((fi78) | 1)), (^(fi78))}
                                _ = v84
                            }
                        default:
                            v85 := GenS8{F0: ("\n\t" + "a"), F1: (v5 << ((545) & 63))}
                            _ = v85
                        }
                    }
                    _ = fi78
                }
                v86 := GenS10{F0: (("a" + "foo bar") != (("" + "foo bar") + "foo bar"))}
                if false {
                    var v87 uint32 = 42
                    _ = v87
                }
                _ = v86
            }
            v4[4294967295] = 42
        } else {
            if ((v5 * v5) <= v5) {
                fmt.Println(v3, v34, v2)
                var v88 int64 = 0
                switch 737 {
                case 421:
                    switch 2147483647 {
                    case 0:
                        var v89 int8 = (v1 << ((1) & 63))
                        _ = v89
                    case 32768:
                        v1 = v5
                    default:
                        switch "a" {
                        case "":
                            v6 = (v60 & (v6 + v6))
                        case "\n\t":
                            var v90 float32 = (1e10 + (((0.1 / 1.0) - 0.1) - 1e10))
                            _ = v90
                        }
                    }
                case -819:
                    var v91 uint32 = 0
                    _ = v91
                case 32768:
                    switch "hello" {
                    case "a":
                        v92 := GenS3{F0: 4294967295}
                        _ = v92
                    }
                }
                _ = v88
            } else {
                if (!true) {
                    var v93 string = "\n\t"
                    _ = v93
                } else {
                    for fi94 := 0; fi94 < 20; fi94++ {
                        v95 := []int{fi94}
                        delete(v4, 2147483647)
                        if len(v95) > 0 {
                            v95[0] = genFn0()
                        }
                        _ = fi94
                        _ = v95
                    }
                    v96 := GenS11{F0: true, F1: int8(v3), F2: 2147483647}
                    for fi97 := 0; fi97 < 18; fi97++ {
                        if (!((false && (v2 != (v2 - -2.25))) || (1 >= 450))) {
                            if false {
                                v38 -= (v38 ^ 42)
                            } else {
                                fmt.Println(v3, v5)
                            }
                        }
                        v98 := []int{}
                        _ = fi97
                        _ = v98
                    }
                    _ = v96
                }
                v99 := v4[457]
                v1 = (v5 | (-(v5)))
                _ = v99
            }
        }
    default:
        v5 = (-(v5))
    }
    for fi100 := 0; fi100 < 2; fi100++ {
        v3 += (^(v3))
        for fi101 := 0; fi101 < 2; fi101++ {
            var v102 int8 = v1
            delete(v4, 0)
            v103 := []int{fi101, 1}
            _ = fi101
            _ = v102
            _ = v103
        }
        _ = fi100
    }
    for fi104 := 0; fi104 < 19; fi104++ {
        fmt.Println(v60, v6)
        fmt.Println(v6)
        _ = fi104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
