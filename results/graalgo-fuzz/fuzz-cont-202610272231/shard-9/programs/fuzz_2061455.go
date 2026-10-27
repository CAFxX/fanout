package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 uint8 = 44
    v4 := make(map[uint32]int)
    var v5 int8 = 85
    var v6 int64 = 0
    if (!false) {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(v5)
            fmt.Println(v6)
            for fi8 := 0; fi8 < 8; fi8++ {
                if (!false) {
                    var v9 uint32 = 4294967295
                    if false {
                        switch (fi7 << ((2) & 63)) {
                        case 998:
                            var v10 float32 = -2.25
                            _ = v10
                        default:
                            if false {
                                if (!true) {
                                    switch ("" + "hello") {
                                    case "hello":
                                        _ = len(v2)
                                    case "":
                                        fmt.Println(v9)
                                    }
                                }
                                switch ("a" + "foo bar") {
                                case "hello":
                                    var v11 float64 = -2.25
                                    _ = v11
                                case "":
                                    var v12 float32 = 3.14159
                                    _ = v12
                                default:
                                    switch "foo bar" {
                                    case "foo bar":
                                        v3 = 255
                                    case "a":
                                        v6 -= v6
                                    case "":
                                        switch 1 {
                                        case 1:
                                            v13 := GenS1{F0: ("a" + "a")}
                                            _ = v13
                                        case 32767:
                                            v1[(fi7 % (((fi8 << ((65535) & 63))) | 1))] = (-1 == 616)
                                            v9 = v9
                                        case 9223372036854775807:
                                            v3 = (v3 / (((v3 / ((42) | 1))) | 1))
                                        default:
                                            var v14 uint64 = 748
                                            _ = v14
                                        }
                                    }
                                    delete(v2, (v9 & (v9 + 42)))
                                    v5 -= (v5 % ((42) | 1))
                                }
                                if false {
                                    v5 += (v5 % ((-128) | 1))
                                }
                            } else {
                                if true {
                                    v15 := []int{(fi8 >> ((157) & 63)), fi7, (fi8 & (^(fi8))), (fi8 ^ (fi8 | int(v9)))}
                                    var v16 uint = 18446744073709551615
                                    _ = v15
                                    _ = v16
                                }
                                if false {
                                    var v17 string = "hello"
                                    var v18 uint16 = 65535
                                    _ = v17
                                    _ = v18
                                }
                            }
                        }
                        var v19 bool = false
                        _ = v19
                    } else {
                        var v20 string = ("foo bar" + "foo bar")
                        var v21 string = ("\n\t" + "foo bar")
                        var v22 int = (fi7 % ((1) | 1))
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := v1[fi7]
                    v24 := GenS2{F0: (v6 | int64(v9))}
                    _ = v23
                    _ = v24
                    _ = v9
                } else {
                    v5 -= (v5 + v5)
                    var v25 int8 = (v5 * v5)
                    _ = v25
                }
                _ = fi8
            }
            _ = fi7
        }
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi26 := 0; fi26 < 4; fi26++ {
                for fi27 := 0; fi27 < 4; fi27++ {
                    if false {
                        if (((v5 | (v5 * (v5 % ((0) | 1)))) <= 1) || ((3.14159 <= -2.25) && false)) {
                            delete(v1, fi26)
                            switch 0 {
                            case 153:
                                fmt.Println(v6, v5)
                                var v28 uint8 = (v3 * v3)
                                v3 += 225
                                _ = v28
                            case 42:
                                v3 = v3
                            default:
                                v29 := GenS1{F0: "hello"}
                                _ = v29
                            }
                            switch int(v3) {
                            case -1:
                                v30 := GenS3{F0: (v5 | (-(v5))), F1: "hello"}
                                _ = v30
                            case 1:
                                v6 += v6
                            default:
                                v31 := []int{(fi27 & (fi27 & fi26)), (fi27 * (fi26 - (fi27 - fi27))), (fi26 / ((fi27) | 1))}
                                _ = v31
                            }
                            v32 := []int{fi26}
                            if true {
                                v32 = append(v32, fi26)
                                if true {
                                    v33 := GenS3{F0: -74, F1: ("foo bar" + (("foo bar" + "a") + ""))}
                                    fmt.Println(v6)
                                    if false {
                                        v5 = (v5 * v5)
                                        switch ("a" + "") {
                                        case "":
                                            v2[2147483647] = (-(fi27))
                                        case "a":
                                            var v34 uint32 = 937
                                            if (("foo bar" != (("a" + "") + "hello")) && (v34 <= 195)) {
                                                v34 += v34
                                                v35 := GenS3{F0: (v5 + v5), F1: (("a" + "foo bar") + "hello")}
                                                v5 = (v5 / ((v5) | 1))
                                                _ = v35
                                            } else {
                                                v34 = (^(v34))
                                                v5 -= v5
                                            }
                                            _ = v34
                                        default:
                                            switch "" {
                                            case "":
                                                v36 := GenS1{F0: ("hello" + "")}
                                                _ = v36
                                            default:
                                                v32 = append(v32, fi26)
                                            }
                                        }
                                        fmt.Println(fi27)
                                    }
                                    _ = v33
                                }
                                if (((0 == 42) || true) && (false || false)) {
                                    var v37 int64 = (v6 - (-(v6)))
                                    _ = v37
                                } else {
                                    var v38 string = "\n\t"
                                    _ = v38
                                }
                            }
                            var v39 int8 = (v5 + v5)
                            _ = v32
                            _ = v39
                        } else {
                            v40 := []int{fi27, fi27}
                            switch "a" {
                            case "a":
                                v5 = (^(v5))
                            case "hello":
                                v6 -= 174
                            case "":
                                if ((true && (-32768 != 1000)) && false) {
                                    v3 += v3
                                    v41 := GenS1{F0: "foo bar"}
                                    v6 = (v6 >> ((125) & 63))
                                    _ = v41
                                }
                            }
                            var v42 bool = false
                            _ = v40
                            _ = v42
                        }
                        var v43 uint16 = 65535
                        _ = v43
                    }
                    fmt.Println(fi26, v6, v6)
                    fmt.Println(v5)
                    v44 := GenS2{F0: (v6 | (-(v6)))}
                    v45 := []int{fi27, fi27}
                    switch fi26 {
                    case 1:
                        v46 := GenS3{F0: v5, F1: "hello"}
                        _ = v46
                    case 497:
                        switch -9223372036854775808 {
                        case 464:
                            switch "hello" {
                            case "\n\t":
                                if (true || (!(true || (2147483647 == fi27)))) {
                                    v3 -= (^(v3))
                                    switch -401 {
                                    case -9223372036854775808:
                                        fmt.Println(v6, v6)
                                    case 2147483647:
                                        _ = len(v45)
                                        switch ("hello" + (("" + "a") + "\n\t")) {
                                        case "\n\t":
                                            if false {
                                                switch fi26 {
                                                case 128:
                                                    if false {
                                                        v5 = 42
                                                    } else {
                                                        var v47 float32 = -2.25
                                                        _ = v47
                                                    }
                                                case 442:
                                                    if (true || (float32(3.14159) != 0.1)) {
                                                        var v48 int8 = int8(v3)
                                                        v49 := GenS1{F0: ""}
                                                        var v50 string = ("a" + "")
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                default:
                                                    if true {
                                                        v6 += (v6 + (v6 * v6))
                                                        v3 -= v3
                                                        var v51 int = fi27
                                                        if true {
                                                            var v52 int32 = 0
                                                            var v53 int64 = (v6 | v6)
                                                            _ = v52
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS1{F0: "\n\t"}
                                                            _ = v54
                                                        }
                                                        _ = v51
                                                    }
                                                }
                                            }
                                        case "":
                                            if false {
                                                var v55 int = fi26
                                                if ("foo bar" >= "\n\t") {
                                                    delete(v1, v55)
                                                    var v56 uint32 = 1
                                                    _ = v56
                                                } else {
                                                    v57 := GenS1{F0: "foo bar"}
                                                    _ = v57
                                                }
                                                v58 := GenS4{F0: (v3 - (^(v3))), F1: 0, F2: v3}
                                                _ = v55
                                                _ = v58
                                            } else {
                                                v3 = v3
                                                if (0 > int32(v6)) {
                                                    fmt.Println(v6, fi27, v3)
                                                    if len(v45) >= 2 {
                                                        v45 = v45[1:len(v45)]
                                                    }
                                                    switch "" {
                                                    case "a":
                                                        var v59 uint8 = (v3 >> ((42) & 63))
                                                        _ = v59
                                                    case "\n\t":
                                                        var v60 int = (fi27 >> ((1) & 63))
                                                        _ = v60
                                                    case "hello":
                                                        var v61 string = "a"
                                                        if true {
                                                            _ = len(v45)
                                                            var v62 bool = false
                                                            _ = len(v1)
                                                            if false {
                                                                var v63 float64 = (0.0 + 1.5)
                                                                _ = v63
                                                            }
                                                            v5 -= 1
                                                            _ = v62
                                                        }
                                                        _ = v61
                                                    }
                                                    v64 := GenS5{F0: v3, F1: 255}
                                                    _ = v64
                                                }
                                                if (42 > -32768) {
                                                    var v65 string = "foo bar"
                                                    v66 := GenS6{F0: true}
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    var v67 uint8 = (v3 / ((255) | 1))
                                                    _ = v67
                                                }
                                            }
                                        default:
                                            fmt.Println(v6, v3)
                                        }
                                    case -2147483648:
                                        switch (fi26 ^ -1) {
                                        case -671:
                                            var v68 uint = 2
                                            _ = v68
                                        case -2147483648:
                                            fmt.Println(fi27)
                                        default:
                                            var v69 uint16 = 192
                                            _ = v69
                                        }
                                    }
                                    _ = len(v4)
                                }
                            default:
                                fmt.Println(fi27, v5, fi26)
                            }
                        }
                    default:
                        var v70 uint = 458
                        _ = v70
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v71 := GenS4{F0: (v3 - (v3 & v3)), F1: 4294967295, F2: (v3 & v3)}
                        _ = v71
                    case "hello":
                        v5 -= (v5 | 100)
                    case "":
                        v5 += v5
                    default:
                        var v72 int16 = 1
                        _ = v72
                    }
                    _ = fi27
                    _ = v44
                    _ = v45
                }
                var v73 bool = false
                _ = fi26
                _ = v73
            }
        case 1:
            if false {
                for fi74 := 0; fi74 < 18; fi74++ {
                    var v75 uint = 4294967295
                    _ = fi74
                    _ = v75
                }
            } else {
                var v76 int32 = int32(v6)
                var v77 int8 = (v5 | 1)
                if (-2147483648 > 0) {
                    v3 = (v3 / (((v3 - (v3 * 255))) | 1))
                    if (((!("hello" >= "foo bar")) || ((v3 > v3) || true)) && (3.14159 <= 0.1)) {
                        v78 := GenS1{F0: "a"}
                        var v79 uint16 = 620
                        v6 += 42
                        _ = v78
                        _ = v79
                    } else {
                        fmt.Println(v5, v5, v5)
                    }
                    v80 := []int{-253}
                    _ = v80
                } else {
                    if ((!(42 <= 42)) || false) {
                        v81 := GenS7{F0: 2, F1: 2, F2: (v3 * 1)}
                        _ = v81
                    } else {
                        v6 = v6
                    }
                }
                _ = v76
                _ = v77
            }
        case 913:
            for fi82 := 0; fi82 < 15; fi82++ {
                var v83 int64 = 987
                var v84 uint16 = 1
                _ = fi82
                _ = v83
                _ = v84
            }
        default:
            v85 := GenS6{F0: false}
            _ = v85
        }
        if (!(true && ((v3 >= v3) || false))) {
            for fi86 := 0; fi86 < 4; fi86++ {
                var v87 uint8 = v3
                _ = fi86
                _ = v87
            }
            v88 := []int{-486}
            var v89 float64 = 1e10
            _ = v88
            _ = v89
        }
        v3 = v3
    }
    v5 -= 100
    var v90 uint16 = 65535
    var v91 uint = 550
    var v92 uint32 = 42
    if false {
        var v93 string = ("" + "foo bar")
        switch "hello" {
        case "":
            v94 := GenS8{F0: v90, F1: (-922 > 354), F2: (v92 - (^(v92)))}
            _ = v94
        }
        _ = v93
    } else {
        v92 = v92
        fmt.Println(v5)
    }
    v3 = 103
    if false {
        fmt.Println(v90, v3, v90)
        for fi95 := 0; fi95 < 1; fi95++ {
            for fi96 := 0; fi96 < 20; fi96++ {
                if (!(true || true)) {
                    v97 := []int{(fi95 % ((-660) | 1))}
                    v92 -= 1
                    v91 -= (v91 * (v91 * 18446744073709551615))
                    _ = v97
                }
                _ = fi96
            }
            var v98 float32 = 3.14159
            _ = fi95
            _ = v98
        }
    }
    v90 = (v90 + v90)
    switch 0 {
    case 42:
        v3 += (v3 + (v3 / (((v3 << (((v91 ^ v91)) & 63))) | 1)))
    case 426:
        v99 := GenS9{F0: 137, F1: (v6 & v6)}
        _ = v99
    case 0:
        v100 := GenS8{F0: (v90 & v90), F1: (18446744073709551615 != (v91 % (((v91 % ((v91) | 1))) | 1))), F2: (^(v92))}
        _ = v100
    }
    v101 := GenS10{F0: 42, F1: (v6 / ((v6) | 1))}
    v90 -= (v90 ^ (v90 * 65535))
    fmt.Println(v90, v90, v3)
    if ((float32(-2.25) / 1.0) < 1e10) {
        fmt.Println(v6, v6, v5)
        var v102 int8 = v5
        var v103 bool = (0 >= v91)
        _ = v102
        _ = v103
    }
    for fi104 := 0; fi104 < 15; fi104++ {
        v1[-793] = true
        fmt.Println(v3, v5)
        v105 := v1[fi104]
        _ = fi104
        _ = v105
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
    _ = v92
}
