package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int8; F1 string; F2 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 uint32 }
func genFn0() bool {
    v4 := []int{-9223372036854775808, 42, -763}
    v5 := GenS1{F0: false}
    switch 1 {
    case -9223372036854775808:
        var v6 uint8 = 128
        _ = v6
    case 9223372036854775807:
        v7 := GenS1{F0: (((false || (true && (255 != 42))) || true) || false)}
        _ = v7
    case 4294967295:
        var v8 float32 = float32(0.0)
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = v4
    _ = v5
    return false
}
func genFn1() int {
    v9 := GenS2{F0: 65535}
    var v10 float32 = float32((0.0 / 1e10))
    _ = v10
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    if true {
        _ = len(v2)
    }
    var v11 uint64 = 828
    switch "" {
    case "a":
        switch -2147483648 {
        case 670:
            v1 += v1
        default:
            var v12 float32 = 1.5
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            switch fi13 {
            case -861:
                v14 := []int{genFn1(), fi13, (fi13 & -1)}
                _ = v14
            case 336:
                v15 := GenS1{F0: true}
                _ = v15
            case 127:
                if ((135 <= v11) && false) {
                    if (!(2 <= 2)) {
                        switch "a" {
                        case "":
                            if true {
                                if ((-1 != 42) && (!true)) {
                                    v1 = 0.0
                                }
                            }
                        default:
                            v3 += v3
                        }
                        if true {
                            fmt.Println(v1, v3, fi13)
                            fmt.Println(v11)
                        }
                    }
                    var v16 float64 = v1
                    _ = v16
                } else {
                    v1 = float64(0.0)
                }
            }
            var v17 int16 = 1
            v1 -= 1.5
            _ = fi13
            _ = v17
        }
    }
    if (0 == 750) {
        v1 = 1.5
        if (42 == v11) {
            switch genFn1() {
            case 0:
                var v18 int8 = (v3 & 75)
                _ = v18
            case -1:
                v19 := GenS3{F0: -881, F1: 1000}
                _ = v19
            case 903:
                if ((!true) && true) {
                    var v20 string = "hello"
                    switch -334 {
                    case 2147483647:
                        var v21 float32 = 0.0
                        _ = v21
                    case 781:
                        v1 += 0.1
                    default:
                        v22 := GenS3{F0: 958, F1: 300}
                        _ = v22
                    }
                    _ = v20
                } else {
                    v1 = ((0.1 * v1) + v1)
                    v23 := GenS1{F0: ((v11 & (v11 - (v11 << ((2) & 63)))) >= v11)}
                    _ = v23
                }
            }
        } else {
            for fi24 := 0; fi24 < 8; fi24++ {
                v1 = float64(0.1)
                fmt.Println(v1)
                _ = fi24
            }
        }
        var v25 uint64 = 517
        _ = v25
    } else {
        fmt.Println(v3, v11)
    }
    v2[533] = genFn1()
    var v26 uint = 255
    if false {
        if true {
            v27 := v2[267]
            v1 -= v1
            if false {
                v27 = v27
                if ((v1 / float64(-2.25)) != v1) {
                    v27 += v27
                    v28 := v2[1]
                    _ = v28
                } else {
                    var v29 uint16 = 0
                    var v30 int32 = 42
                    _ = v29
                    _ = v30
                }
            }
            _ = v27
        }
        if (2147483647 >= -975) {
            v11 += (^(v11))
            var v31 uint64 = (v11 % (((v11 + 0)) | 1))
            _ = v31
        } else {
            var v32 int = 1
            _ = v32
        }
        switch 1 {
        case 127:
            if (false && false) {
                var v33 string = ("foo bar" + "\n\t")
                v11 -= v11
                _ = v33
            } else {
                if (!(false || true)) {
                    v34 := GenS4{F0: 4294967295}
                    fmt.Println(v1)
                    _ = v34
                }
                v2[42] = genFn1()
                if false {
                    v35 := GenS3{F0: 42, F1: 32767}
                    _ = len(v2)
                    _ = v35
                }
            }
        case -2147483648:
            for fi36 := 0; fi36 < 14; fi36++ {
                for fi37 := 0; fi37 < 5; fi37++ {
                    switch (fi37 - (fi36 + genFn1())) {
                    case 32767:
                        v38 := GenS3{F0: genFn1(), F1: 32767}
                        _ = v38
                    case 2147483647:
                        var v39 uint = 88
                        _ = v39
                    case -354:
                        var v40 int16 = 1000
                        _ = v40
                    default:
                        var v41 float32 = -2.25
                        _ = v41
                    }
                    var v42 uint = uint(v11)
                    _ = fi37
                    _ = v42
                }
                switch fi36 {
                case 2147483647:
                    var v43 float64 = v1
                    _ = v43
                case 128:
                    for fi44 := 0; fi44 < 20; fi44++ {
                        v1 = v1
                        v45 := GenS5{F0: (967 == 993), F1: (^(v11))}
                        v46 := GenS1{F0: false}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                default:
                    if (("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) <= ("" + "foo bar")) {
                        fmt.Println(v1, v3, v1)
                        v11 += (^(v11))
                        if false {
                            v47 := GenS1{F0: ((2147483648 == 1) && false)}
                            for fi48 := 0; fi48 < 20; fi48++ {
                                v11 += v11
                                if (214 > 215) {
                                    var v49 float64 = v1
                                    _ = v49
                                }
                                if false {
                                    fmt.Println(v11)
                                    v50 := GenS2{F0: v26}
                                    v51 := GenS6{F0: 348, F1: -9223372036854775808, F2: v11}
                                    _ = v50
                                    _ = v51
                                } else {
                                    v1 -= ((3.14159 / float64(1.5)) + (v1 / v1))
                                    fmt.Println(v3)
                                }
                                _ = fi48
                            }
                            v52 := v2[uint32(v26)]
                            _ = v47
                            _ = v52
                        } else {
                            if true {
                                if true {
                                    if true {
                                        v11 = v11
                                        if false {
                                            v53 := GenS7{F0: (^(fi36)), F1: (fi36 >> (((v26 << (((v26 | 4294967295)) & 63))) & 63))}
                                            _ = v53
                                        } else {
                                            v26 = 42
                                            v11 -= v11
                                        }
                                        v54 := GenS8{F0: int8(v11), F1: "hello", F2: ""}
                                        _ = v54
                                    }
                                } else {
                                    switch (("foo bar" + (("hello" + "a") + "hello")) + "") {
                                    case "":
                                        if true {
                                            v55 := GenS9{F0: 94}
                                            _ = v55
                                        } else {
                                            if false {
                                                v3 -= (v3 / ((-128) | 1))
                                                _ = len(v2)
                                            } else {
                                                v3 = (v3 - v3)
                                            }
                                        }
                                    }
                                    var v56 float32 = -2.25
                                    v57 := v2[205]
                                    _ = v56
                                    _ = v57
                                }
                                var v58 int = genFn1()
                                for fi59 := 0; fi59 < 13; fi59++ {
                                    fmt.Println(fi59, v26)
                                    if false {
                                        v60 := GenS1{F0: false}
                                        if false {
                                            v61 := GenS4{F0: 0}
                                            _ = v61
                                        }
                                        _ = v60
                                    } else {
                                        switch v58 {
                                        case -9223372036854775808:
                                            var v62 float32 = (3.14159 + 1.5)
                                            _ = v62
                                        case 9223372036854775807:
                                            v1 = v1
                                        }
                                    }
                                    _ = fi59
                                }
                                _ = v58
                            } else {
                                var v63 int8 = v3
                                var v64 uint16 = 32767
                                _ = v63
                                _ = v64
                            }
                            switch "" {
                            case "hello":
                                if (!true) {
                                    v65 := GenS10{F0: 0, F1: 2147483647}
                                    _ = v65
                                }
                            case "a":
                                v1 = 1e10
                            case "foo bar":
                                switch "\n\t" {
                                case "":
                                    var v66 uint8 = 127
                                    _ = v66
                                case "foo bar":
                                    fmt.Println(fi36)
                                }
                            default:
                                var v67 float32 = (0.0 - ((3.14159 / 1.5) + 0.1))
                                _ = v67
                            }
                        }
                    }
                }
                _ = fi36
            }
        case 69:
            if false {
                v68 := []int{-9223372036854775808, genFn1(), 186}
                _ = v68
            } else {
                v69 := GenS10{F0: 42, F1: 2}
                if (0 == 0) {
                    v11 = 926
                } else {
                    var v70 bool = false
                    v71 := []int{128, 128}
                    _ = v70
                    _ = v71
                }
                _ = v69
            }
        default:
            if true {
                v72 := []int{1, 919}
                v11 = (v11 << ((v26) & 63))
                _ = v72
            } else {
                v3 = v3
                fmt.Println(v26, v1, v3)
            }
        }
    }
    for fi73 := 0; fi73 < 15; fi73++ {
        v74 := []int{254, genFn1(), (fi73 & (fi73 + (fi73 - fi73)))}
        _ = fi73
        _ = v74
    }
    v11 = (v11 % ((0) | 1))
    if (127 == v3) {
        v3 -= (v3 ^ (v3 & v3))
        fmt.Println(v1, v26, v3)
        for fi75 := 0; fi75 < 19; fi75++ {
            for fi76 := 0; fi76 < 17; fi76++ {
                switch fi76 {
                case -2147483648:
                    var v77 string = "\n\t"
                    _ = v77
                case 474:
                    var v78 uint8 = 255
                    _ = v78
                case 436:
                    v26 -= (v26 / ((v26) | 1))
                default:
                    v79 := []int{(fi75 ^ (fi75 * -2147483648)), (^(fi75)), 850}
                    _ = v79
                }
                if (65535 >= 400) {
                    v3 -= v3
                    switch -9223372036854775808 {
                    case 128:
                        var v80 uint8 = 127
                        _ = v80
                    case 4294967295:
                        if true {
                            delete(v2, 2147483647)
                        } else {
                            var v81 int32 = 719
                            v82 := GenS5{F0: true, F1: 2}
                            var v83 int16 = 843
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                    }
                } else {
                    v3 -= v3
                    delete(v2, 4)
                }
                _ = fi76
            }
            var v84 uint32 = 539
            _ = fi75
            _ = v84
        }
    } else {
        v3 = (v3 | (v3 / (((v3 * v3)) | 1)))
        v1 = v1
    }
    fmt.Println(v26, v11, v26)
    switch 32767 {
    case 9223372036854775807:
        v85 := []int{32767, genFn1(), 128}
        _ = v85
    default:
        switch ("foo bar" + (("foo bar" + ((("a" + (("\n\t" + "a") + "a")) + ("hello" + "a")) + "")) + "a")) {
        case "foo bar":
            if true {
                var v86 uint8 = 0
                if true {
                    var v87 int16 = 32767
                    fmt.Println(v1)
                    _ = v87
                }
                _ = v86
            }
        case "a":
            var v88 float64 = float64(-2.25)
            _ = v88
        default:
            v3 -= 67
        }
    }
    v89 := v2[4294967295]
    fmt.Println(v3)
    fmt.Println(v11, v3)
    if (1 != -1) {
        for fi90 := 0; fi90 < 15; fi90++ {
            switch (fi90 & 0) {
            case 32767:
                var v91 uint64 = 579
                _ = v91
            }
            var v92 int16 = 960
            var v93 uint = 0
            _ = fi90
            _ = v92
            _ = v93
        }
        v3 -= v3
    } else {
        if false {
            v89 += (v89 << ((v26) & 63))
        } else {
            if true {
                v94 := []int{}
                for fi95 := 0; fi95 < 10; fi95++ {
                    var v96 int = (-(fi95))
                    _ = fi95
                    _ = v96
                }
                _ = len(v94)
                _ = v94
            }
            v97 := []int{127, -650, v89, genFn1()}
            var v98 float64 = (1.5 - (v1 * (float64(0.1) / v1)))
            _ = v97
            _ = v98
        }
        switch (v89 * v89) {
        case 2147483647:
            v99 := []int{genFn1(), v89, (v89 & v89)}
            _ = v99
        }
        v100 := []int{}
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v89
}
