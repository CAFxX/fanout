package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int; F1 string; F2 uint64 }
type GenS7 struct { F0 int16; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint64; F1 int8; F2 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 string }
func genFn0(p0 int, p1 int64) uint16 {
    fmt.Println(p1, p1)
    v5 := GenS1{F0: 42, F1: 577}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1() int16 {
    if (255 != 4294967295) {
        if (-2.25 < 1.5) {
            if true {
                var v6 uint = 42
                v7 := GenS2{F0: 501}
                v6 = 4294967295
                _ = v6
                _ = v7
            }
            var v8 float64 = 0.0
            switch -1 {
            case 216:
                for fi9 := 0; fi9 < 5; fi9++ {
                    v8 += v8
                    v10 := []int{(^(fi9)), fi9, -827, fi9}
                    v8 -= v8
                    _ = fi9
                    _ = v10
                }
            case 32767:
                var v11 uint16 = 255
                _ = v11
            case -2147483648:
                var v12 uint64 = 499
                _ = v12
            }
            _ = v8
        } else {
            switch -20 {
            case -178:
                var v13 uint32 = 2147483648
                _ = v13
            case 2147483647:
                v14 := GenS1{F0: 1, F1: 0}
                _ = v14
            }
            var v15 int = 9223372036854775807
            v15 += (v15 - -232)
            _ = v15
        }
        var v16 int = 127
        v17 := []int{(-(v16)), (v16 + v16), (v16 >> ((1) & 63))}
        _ = v16
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 string = ""
            v19 = v19
            if (true || false) {
                for fi20 := 0; fi20 < 4; fi20++ {
                    switch fi20 {
                    case 42:
                        if (18446744073709551615 > 4294967295) {
                            var v21 uint32 = 4294967295
                            if true {
                                fmt.Println(v21, v21)
                                v21 = v21
                            } else {
                                fmt.Println(v19, fi20)
                                var v22 uint8 = 127
                                var v23 uint32 = (v21 << ((4294967295) & 63))
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        } else {
                            var v24 string = (v19 + "")
                            var v25 uint32 = 1
                            switch (^(fi18)) {
                            case -569:
                                var v26 float64 = 1.5
                                var v27 int = fi20
                                v28 := []int{0, v27, 32767, fi18}
                                _ = v26
                                _ = v27
                                _ = v28
                            default:
                                var v29 float64 = (1.5 / -2.25)
                                _ = v29
                            }
                            _ = v24
                            _ = v25
                        }
                    case 144:
                        if false {
                            var v30 uint16 = 2
                            var v31 uint64 = uint64(v30)
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint = 0
                            _ = v32
                        }
                    case -1:
                        if true {
                            switch fi20 {
                            case 952:
                                v33 := GenS1{F0: 212, F1: 65535}
                                _ = v33
                            case 765:
                                fmt.Println(fi18, fi18, fi20)
                            }
                            v34 := []int{}
                            fmt.Println(fi18, v19, v19)
                            _ = v34
                        } else {
                            if (true && false) {
                                v35 := []int{(fi20 - (fi20 << ((42) & 63)))}
                                if ((("hello" + v19) + (v19 + (v19 + (v19 + "a")))) == ("" + (("" + (v19 + "\n\t")) + v19))) {
                                    if len(v35) > 0 {
                                        v35[0] = 127
                                    }
                                    fmt.Println(fi18, fi20)
                                    v35 = append(v35, (fi20 & fi20))
                                }
                                fmt.Println(v19)
                                _ = v35
                            } else {
                                v36 := []int{(fi18 + (fi20 | (fi20 + (^(fi20))))), (fi20 + (fi18 % ((fi18) | 1)))}
                                switch (((v19 + v19) + v19) + "foo bar") {
                                case "foo bar":
                                    v19 = ((v19 + "a") + ("\n\t" + (("foo bar" + ("a" + ((v19 + ("\n\t" + "a")) + (((v19 + "") + v19) + v19)))) + "\n\t")))
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        if false {
                                            if len(v36) >= 2 {
                                                v36 = v36[1:len(v36)]
                                            }
                                            var v37 float32 = (1.5 / (1e10 + 1.5))
                                            _ = v37
                                        } else {
                                            fmt.Println(v19, fi20)
                                        }
                                    case "":
                                        v38 := GenS2{F0: 2}
                                        _ = v38
                                    case "a":
                                        var v39 uint32 = 2
                                        _ = v39
                                    default:
                                        v19 = v19
                                    }
                                }
                                switch fi18 {
                                case 32768:
                                    v19 = v19
                                default:
                                    fmt.Println(fi20, fi18)
                                }
                                _ = v36
                            }
                            if (1 < 32767) {
                                v19 = (("a" + ("foo bar" + (v19 + "\n\t"))) + v19)
                                fmt.Println(fi18, fi20)
                                v40 := GenS3{F0: 42}
                                _ = v40
                            } else {
                                var v41 bool = false
                                v42 := []int{362, 2147483647, (fi20 >> ((0) & 63)), fi18}
                                var v43 uint8 = 120
                                v19 = v19
                                var v44 float64 = (float64(0.1) + (1.5 + 0.0))
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            v19 = v19
                            switch fi20 {
                            case 2147483647:
                                v45 := []int{fi18, fi20}
                                _ = v45
                            case 4294967295:
                                if (!false) {
                                    v19 = v19
                                    v19 = ("" + (("foo bar" + (v19 + v19)) + (v19 + v19)))
                                    fmt.Println(fi20, fi20, v19)
                                }
                            default:
                                var v46 bool = (genFn0((^(fi20)), 178) > 250)
                                _ = v46
                            }
                        }
                    default:
                        var v47 int32 = 1
                        _ = v47
                    }
                    v48 := GenS2{F0: 255}
                    switch fi18 {
                    case 42:
                        v19 = "foo bar"
                    case -681:
                        var v49 int = (fi20 * fi20)
                        _ = v49
                    case 9223372036854775807:
                        v50 := []int{}
                        _ = v50
                    }
                    v51 := GenS4{F0: 0, F1: 42}
                    _ = fi20
                    _ = v48
                    _ = v51
                }
                if (-32768 >= 42) {
                    fmt.Println(fi18)
                    switch "\n\t" {
                    case "hello":
                        v19 = ((v19 + (v19 + (v19 + v19))) + ((v19 + "foo bar") + ("\n\t" + v19)))
                    case "a":
                        if (-9223372036854775808 > -32) {
                            if false {
                                var v52 string = v19
                                v53 := []int{-9223372036854775808, fi18, (fi18 | fi18), 0}
                                _ = v52
                                _ = v53
                            } else {
                                fmt.Println(fi18, v19)
                                v19 = "hello"
                                v54 := GenS2{F0: 1}
                                _ = v54
                            }
                            v55 := []int{}
                            fmt.Println(fi18)
                            _ = v55
                        } else {
                            fmt.Println(fi18)
                        }
                    case "\n\t":
                        var v56 int16 = 775
                        _ = v56
                    default:
                        var v57 uint16 = 0
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 15; fi58++ {
                        fmt.Println(fi58, fi58, fi18)
                        v59 := GenS4{F0: -112, F1: -1}
                        switch (("\n\t" + "foo bar") + v19) {
                        case "":
                            v60 := GenS5{F0: -32768, F1: 62}
                            _ = v60
                        case "\n\t":
                            fmt.Println(fi18, fi18)
                        }
                        _ = fi58
                        _ = v59
                    }
                } else {
                    switch fi18 {
                    case 620:
                        switch "hello" {
                        case "hello":
                            fmt.Println(v19)
                        default:
                            fmt.Println(v19, v19)
                        }
                    case -1:
                        v19 = (v19 + v19)
                    case 725:
                        var v61 int16 = -93
                        _ = v61
                    default:
                        var v62 bool = true
                        _ = v62
                    }
                    fmt.Println(fi18, v19)
                }
                v19 = v19
            }
            _ = fi18
            _ = v19
        }
        var v63 string = "\n\t"
        var v64 int32 = 2147483647
        _ = v63
        _ = v64
    } else {
        var v65 float64 = (-2.25 * 1e10)
        switch 973 {
        case 61:
            fmt.Println(v65)
        case 2147483647:
            var v66 uint64 = 393
            _ = v66
        default:
            for fi67 := 0; fi67 < 19; fi67++ {
                fmt.Println(fi67)
                v68 := []int{}
                switch "\n\t" {
                case "a":
                    v65 = v65
                default:
                    if (true && false) {
                        fmt.Println(fi67, fi67)
                        var v69 int64 = 42
                        var v70 uint16 = 262
                        _ = v69
                        _ = v70
                    }
                }
                _ = fi67
                _ = v68
            }
        }
        var v71 int16 = 294
        _ = v65
        _ = v71
    }
    v72 := GenS6{F0: 95, F1: "hello", F2: 2}
    _ = v72
    return 42
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := []uint32{58, 922, 1, 0}
    v1 = ("\n\t" == "foo bar")
    v73 := GenS2{F0: 42}
    switch 128 {
    case 99:
        var v74 int8 = 41
        _ = v74
    case -1:
        for fi75 := 0; fi75 < 16; fi75++ {
            v76 := GenS7{F0: 0, F1: 255, F2: 1000}
            _ = len(v4)
            _ = fi75
            _ = v76
        }
    default:
        v1 = true
    }
    v77 := GenS4{F0: -128, F1: -1}
    v78 := GenS8{F0: 65535, F1: 100}
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v79 int = -458
        _ = v79
    }
    v2[590] = true
    var v80 int8 = -42
    for fi81 := 0; fi81 < 9; fi81++ {
        switch "hello" {
        case "hello":
            if (v1 || v1) {
                v3[2147483648] = ("" + "\n\t")
                v82 := GenS9{F0: 4294967295, F1: (v80 + (v80 % ((v80) | 1))), F2: -138}
                v80 = v80
                _ = v82
            } else {
                delete(v3, 65535)
                if true {
                    fmt.Println(v1)
                    v1 = v1
                } else {
                    var v83 uint8 = 1
                    if true {
                        var v84 uint = 4294967295
                        v85 := GenS10{F0: 0}
                        var v86 int16 = 615
                        _ = v84
                        _ = v85
                        _ = v86
                    } else {
                        switch fi81 {
                        case 32767:
                            for fi87 := 0; fi87 < 17; fi87++ {
                                switch (-(fi87)) {
                                case -942:
                                    v83 -= v83
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if (2 >= 2) {
                                    switch "a" {
                                    case "":
                                        _ = len(v2)
                                    default:
                                        v1 = (!true)
                                    }
                                } else {
                                    var v88 bool = (v1 && (!v1))
                                    var v89 uint64 = 2
                                    v1 = ((false && v88) && (false && ((0.1 - 3.14159) >= 0.1)))
                                    v4 = append(v4, 562)
                                    _ = v88
                                    _ = v89
                                }
                                var v90 uint8 = (v83 >> ((87) & 63))
                                var v91 uint32 = 2
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v90, v83)
                                    v1 = ((837 <= 2) || true)
                                case "hello":
                                    var v92 uint8 = (v90 + (v90 & v83))
                                    _ = v92
                                case "":
                                    var v93 bool = (2 < 118)
                                    _ = v93
                                }
                                _ = fi87
                                _ = v90
                                _ = v91
                            }
                        default:
                            if (32767 >= 32767) {
                                var v94 float32 = (1e10 * ((3.14159 + (1.5 * 0.0)) / 1e10))
                                for fi95 := 0; fi95 < 16; fi95++ {
                                    var v96 uint16 = 65535
                                    var v97 bool = (!(v83 <= v83))
                                    var v98 int32 = 42
                                    v97 = v97
                                    var v99 string = "hello"
                                    _ = fi95
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                    _ = v99
                                }
                                for fi100 := 0; fi100 < 19; fi100++ {
                                    v83 = 42
                                    _ = fi100
                                }
                                _ = v94
                            }
                        }
                        fmt.Println(v1, fi81, v1)
                    }
                    _ = v83
                }
            }
        case "a":
            if false {
                v101 := GenS11{F0: -1, F1: 255}
                v102 := GenS7{F0: 32767, F1: 451, F2: -240}
                switch "a" {
                case "\n\t":
                    v103 := GenS12{F0: "hello"}
                    _ = v103
                default:
                    if ((!(v1 || v1)) || true) {
                        var v104 uint = 976
                        _ = len(v2)
                        _ = v104
                    }
                }
                _ = v101
                _ = v102
            } else {
                var v105 string = "foo bar"
                v80 = v80
                _ = v105
            }
        case "foo bar":
            fmt.Println(v80)
        default:
            v1 = (true && (("hello" + ("hello" + "a")) == "hello"))
        }
        _ = fi81
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v1)
    fmt.Println(v80, v1)
    var v106 uint64 = 228
    v107 := GenS5{F0: -528, F1: 1000}
    v3[65535] = ("a" + "\n\t")
    fmt.Println(v1, v80, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v77
    _ = v78
    _ = v80
}
