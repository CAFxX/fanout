package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int8; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int64; F2 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint16; F1 string; F2 string }
func genFn0(p0 int8) uint32 {
    fmt.Println(p0)
    var v5 uint16 = 18
    var v6 uint16 = (v5 % (((v5 % (((^(v5))) | 1))) | 1))
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = (p0 >> ((809) & 63))
        _ = fi7
    }
    if (false && false) {
        switch "foo bar" {
        case "":
            v8 := GenS1{F0: 42}
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: 42}
            _ = v9
        case "hello":
            switch 523 {
            case 758:
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := GenS2{F0: (-(fi10))}
                    var v12 float32 = ((-2.25 / 1.0) / -2.25)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                if true {
                    if false {
                        if (698 < 2147483647) {
                            v5 += v6
                        }
                    } else {
                        switch "foo bar" {
                        case "hello":
                            var v13 int8 = p0
                            _ = v13
                        }
                        var v14 float64 = -2.25
                        _ = v14
                    }
                    fmt.Println(v6)
                    v5 = (v6 | (v6 ^ 32767))
                } else {
                    v15 := GenS2{F0: -1}
                    for fi16 := 0; fi16 < 17; fi16++ {
                        var v17 uint = 2
                        if true {
                            switch (("\n\t" + "\n\t") + "foo bar") {
                            case "foo bar":
                                v18 := GenS1{F0: 0}
                                _ = v18
                            default:
                                var v19 string = ""
                                _ = v19
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                    fmt.Println(v6)
                    _ = v15
                }
            }
        }
        v20 := GenS3{F0: -287}
        switch "\n\t" {
        case "foo bar":
            v21 := []int{}
            _ = v21
        case "a":
            v22 := GenS4{F0: 238, F1: p0, F2: "hello"}
            _ = v22
        }
        _ = v20
    }
    _ = p0
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 bool, p1 int16) uint {
    p0 = p0
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 bool = (365 < 162)
        _ = fi23
        _ = v24
    }
    if ("" <= "") {
        switch -1000 {
        case -2147483648:
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 int64 = 929
                v27 := GenS3{F0: (fi25 ^ (fi25 | (^(fi25))))}
                _ = fi25
                _ = v26
                _ = v27
            }
        case -9223372036854775808:
            v28 := []int{61, 9223372036854775807, 949, 191}
            _ = v28
        default:
            v29 := GenS4{F0: 127, F1: 1, F2: ("foo bar" + "\n\t")}
            _ = v29
        }
        p1 = p1
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := []int{42, (fi30 << ((4294967295) & 63))}
            switch fi30 {
            case -869:
                p0 = ((true || (p0 || (368 <= 929))) && ((((fi30 + fi30) == (fi30 ^ 312)) || false) || p0))
            case 484:
                p0 = p0
            case 4294967295:
                for fi32 := 0; fi32 < 15; fi32++ {
                    if p0 {
                        if p0 {
                            _ = len(v31)
                        }
                    }
                    p0 = ((324 >= 571) && ("foo bar" <= ("\n\t" + "hello")))
                    p1 += (p1 % ((-32768) | 1))
                    _ = fi32
                }
            }
            v33 := GenS5{F0: genFn0(0)}
            _ = fi30
            _ = v31
            _ = v33
        }
    } else {
        fmt.Println(p0)
        var v34 int8 = -8
        _ = v34
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 42
    var v3 uint8 = 255
    var v4 uint16 = 32767
    switch 42 {
    case 556:
        for fi35 := 0; fi35 < 10; fi35++ {
            v1[fi35] = ("" + "foo bar")
            var v36 bool = false
            var v37 uint16 = v4
            _ = fi35
            _ = v36
            _ = v37
        }
    case -2147483648:
        if true {
            v38 := []int{535, 221, 32768}
            v1[127] = "foo bar"
            v1[1] = "hello"
            _ = v38
        }
    }
    switch 683 {
    case 9223372036854775807:
        fmt.Println(v4, v2)
    case 0:
        switch -327 {
        case 128:
            v3 += v3
        default:
            var v39 uint64 = 648
            _ = v39
        }
    default:
        if (true && true) {
            fmt.Println(v4, v3)
            fmt.Println(v2, v4, v3)
            var v40 int64 = 9223372036854775807
            _ = v40
        } else {
            var v41 bool = (1 > int8(v4))
            v42 := []int{-2147483648, -1, 781, 128}
            var v43 int64 = -459
            _ = v41
            _ = v42
            _ = v43
        }
    }
    v44 := v1[-9223372036854775808]
    v44 = (v44 + v44)
    v45 := v1[496]
    var v46 string = (((v45 + v44) + "a") + ((((v44 + "\n\t") + (v44 + v44)) + "a") + "a"))
    v47 := GenS5{F0: genFn0(-7)}
    v45 = "hello"
    v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
    v4 = v4
    v48 := []int{}
    for fi49 := 0; fi49 < 15; fi49++ {
        if ((1 >= 5) && ((1 == 9223372036854775807) || (880 <= 42))) {
            v50 := GenS5{F0: genFn0(-55)}
            _ = v50
        } else {
            for fi51 := 0; fi51 < 12; fi51++ {
                _ = len(v1)
                _ = fi51
            }
        }
        _ = fi49
    }
    var v52 uint32 = 1
    if (!(v3 > (^(v3)))) {
        for fi53 := 0; fi53 < 16; fi53++ {
            var v54 int32 = (^(v2))
            _ = fi53
            _ = v54
        }
        if (-767 != -1) {
            v44 = (v45 + "a")
        } else {
            var v55 float64 = (1e10 - 0.1)
            if len(v48) > 0 {
                v48[0] = 2147483647
            }
            v46 = ("" + ("foo bar" + ((v44 + "hello") + v45)))
            _ = v55
        }
    } else {
        v56 := GenS2{F0: 0}
        var v57 uint64 = 1
        v58 := GenS3{F0: -2147483648}
        _ = v56
        _ = v57
        _ = v58
    }
    switch 328 {
    case 9223372036854775807:
        fmt.Println(v4, v3, v3)
    default:
        if (((!((v2 - (-(v2))) == (v2 ^ v2))) || ((v2 % ((v2) | 1)) >= v2)) || ((v3 - v3) == 128)) {
            for fi59 := 0; fi59 < 20; fi59++ {
                fmt.Println(fi59, v44, v44)
                v52 = 724
                _ = fi59
            }
            var v60 uint64 = 73
            var v61 int8 = 1
            _ = v60
            _ = v61
        } else {
            if false {
                v3 += 217
                var v62 float64 = ((1e10 + -2.25) - (1e10 * 1e10))
                if true {
                    v63 := GenS2{F0: 2147483647}
                    var v64 int16 = -391
                    for fi65 := 0; fi65 < 5; fi65++ {
                        v3 = v3
                        _ = fi65
                    }
                    _ = v63
                    _ = v64
                } else {
                    v66 := GenS5{F0: (v52 % ((v52) | 1))}
                    _ = len(v48)
                    if ((v45 + (v46 + v46)) < "") {
                        switch -2147483648 {
                        case -1:
                            switch int(v2) {
                            case 835:
                                var v67 string = v46
                                _ = v67
                            case -2147483648:
                                switch 32768 {
                                case -9223372036854775808:
                                    for fi68 := 0; fi68 < 17; fi68++ {
                                        if (v3 == (v3 + (v3 % (((v3 >> ((255) & 63))) | 1)))) {
                                            if (4294967295 < uint64(v52)) {
                                                for fi69 := 0; fi69 < 17; fi69++ {
                                                    switch "\n\t" {
                                                    case "":
                                                        if (true && (!(false || true))) {
                                                            if (-107 != 706) {
                                                                v62 = 0.0
                                                                if len(v48) > 0 {
                                                                    v48[0] = fi69
                                                                }
                                                                var v70 int8 = -59
                                                                var v71 int8 = v70
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            v72 := GenS3{F0: (fi69 % (((fi68 | 32767)) | 1))}
                                                            if len(v48) >= 2 {
                                                                v48 = v48[1:len(v48)]
                                                            }
                                                            _ = v72
                                                        }
                                                    default:
                                                        switch 127 {
                                                        case 2147483647:
                                                            switch (-(fi68)) {
                                                            case 4294967295:
                                                                var v73 uint64 = 2
                                                                _ = v73
                                                            case -440:
                                                                var v74 int64 = 9223372036854775807
                                                                _ = v74
                                                            case -2147483648:
                                                                var v75 uint8 = v3
                                                                _ = v75
                                                            }
                                                        case -781:
                                                            v2 += 42
                                                            var v76 uint32 = v52
                                                            v3 = (v3 / ((v3) | 1))
                                                            _ = v76
                                                        case 921:
                                                            v77 := GenS6{F0: true}
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = fi69
                                                }
                                                switch fi68 {
                                                case 0:
                                                    var v78 uint16 = 273
                                                    _ = v78
                                                case 1:
                                                    fmt.Println(v45)
                                                case -129:
                                                    v1[-1] = ""
                                                default:
                                                    v79 := GenS7{F0: -1, F1: v4, F2: v3}
                                                    _ = v79
                                                }
                                                var v80 float32 = (0.1 + 0.0)
                                                _ = v80
                                            } else {
                                                if ((!((v52 >= 2147483647) && (genFn0(1) < 4294967295))) || (-696 == -881)) {
                                                    v44 = v45
                                                    if true {
                                                        for fi81 := 0; fi81 < 18; fi81++ {
                                                            v45 = ((v44 + v46) + v44)
                                                            fmt.Println(v4)
                                                            _ = fi81
                                                        }
                                                        v44 = ((v45 + ("foo bar" + ("hello" + (v44 + "\n\t")))) + v45)
                                                        v82 := GenS8{F0: 0, F1: -450, F2: ("" + (v44 + v45))}
                                                        _ = v82
                                                    }
                                                } else {
                                                    switch (v46 + "a") {
                                                    case "foo bar":
                                                        v3 -= 70
                                                    case "":
                                                        v52 -= (^(v52))
                                                    default:
                                                        if true {
                                                            v52 -= 2
                                                            v1[-2147483648] = ("foo bar" + v46)
                                                            switch "foo bar" {
                                                            case "":
                                                                _ = len(v48)
                                                            case "foo bar":
                                                                if true {
                                                                    v45 = (((v44 + "") + v46) + "")
                                                                }
                                                            default:
                                                                v1[fi68] = v46
                                                            }
                                                        } else {
                                                            v45 = "hello"
                                                            var v83 uint64 = 358
                                                            _ = v83
                                                        }
                                                    }
                                                    if len(v48) >= 2 {
                                                        v48 = v48[1:len(v48)]
                                                    }
                                                    fmt.Println(v44, v45, v46)
                                                }
                                                v45 = v45
                                                for fi84 := 0; fi84 < 10; fi84++ {
                                                    v85 := GenS7{F0: 1, F1: v4, F2: (v3 ^ 227)}
                                                    v86 := GenS9{F0: fi84}
                                                    _ = fi84
                                                    _ = v85
                                                    _ = v86
                                                }
                                            }
                                            if true {
                                                v44 = v45
                                            }
                                            if (true || false) {
                                                v87 := GenS2{F0: fi68}
                                                for fi88 := 0; fi88 < 7; fi88++ {
                                                    var v89 uint32 = 445
                                                    if true {
                                                        v4 -= (v4 & v4)
                                                        v46 = (v46 + v44)
                                                        if (false || true) {
                                                            fmt.Println(v3)
                                                            if len(v48) >= 2 {
                                                                v48 = v48[1:len(v48)]
                                                            }
                                                        } else {
                                                            switch v44 {
                                                            case "foo bar":
                                                                switch v45 {
                                                                case "":
                                                                    v45 = "hello"
                                                                case "a":
                                                                    var v90 float32 = (0.1 * ((1.5 / (1.5 + (1.5 - (3.14159 * -2.25)))) / float32(v62)))
                                                                    _ = v90
                                                                case "\n\t":
                                                                    v4 -= (v4 & v4)
                                                                }
                                                            case "a":
                                                                v91 := GenS2{F0: (fi88 | (fi88 | (fi68 << ((65535) & 63))))}
                                                                _ = v91
                                                            case "":
                                                                var v92 uint = 646
                                                                _ = v92
                                                            }
                                                        }
                                                        switch v46 {
                                                        case "foo bar":
                                                            if true {
                                                                var v93 uint64 = 42
                                                                if len(v48) > 0 {
                                                                    v48[0] = 1
                                                                }
                                                                if len(v48) > 0 {
                                                                    v48[0] = (fi68 | fi88)
                                                                }
                                                                _ = v93
                                                            } else {
                                                                if (!false) {
                                                                    fmt.Println(v52, v4, fi68)
                                                                    v62 = 1e10
                                                                    v94 := GenS10{F0: v89, F1: 2}
                                                                    _ = v94
                                                                }
                                                                v95 := GenS11{F0: 18446744073709551615, F1: -32768}
                                                                var v96 int = fi88
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                        case "a":
                                                            v4 -= v4
                                                            fmt.Println(fi88, v4)
                                                            v89 = v52
                                                        case "\n\t":
                                                            v97 := GenS9{F0: fi68}
                                                            _ = v97
                                                        default:
                                                            v4 += (v4 * v4)
                                                        }
                                                    } else {
                                                        v48 = append(v48, fi88)
                                                        v89 = genFn0(42)
                                                        switch (fi88 | fi68) {
                                                        case -1:
                                                            v44 = "a"
                                                        case 2147483647:
                                                            if len(v48) > 0 {
                                                                v48[0] = (fi88 * (fi68 / ((fi88) | 1)))
                                                            }
                                                        case 1:
                                                            var v98 int8 = 100
                                                            _ = v98
                                                        default:
                                                            var v99 int64 = -654
                                                            _ = v99
                                                        }
                                                    }
                                                    var v100 int8 = 42
                                                    _ = fi88
                                                    _ = v100
                                                    _ = v89
                                                }
                                                v44 = "foo bar"
                                                _ = v87
                                            } else {
                                                fmt.Println(v62, v62, v62)
                                            }
                                        } else {
                                            v101 := GenS12{F0: uint64(v52)}
                                            _ = v101
                                        }
                                        fmt.Println(v46, v3)
                                        _ = fi68
                                    }
                                }
                            }
                        }
                        v102 := v1[127]
                        switch -2147483648 {
                        case 820:
                            for fi103 := 0; fi103 < 9; fi103++ {
                                v104 := GenS3{F0: fi103}
                                v105 := GenS13{F0: v4, F1: v46, F2: v45}
                                _ = fi103
                                _ = v104
                                _ = v105
                            }
                        default:
                            if (!false) {
                                v106 := GenS12{F0: 202}
                                var v107 string = (v102 + v102)
                                _ = v106
                                _ = v107
                            }
                        }
                        _ = v102
                    } else {
                        delete(v1, 2147483647)
                    }
                    _ = v66
                }
                _ = v62
            } else {
                v108 := GenS11{F0: 691, F1: 0}
                for fi109 := 0; fi109 < 3; fi109++ {
                    v110 := GenS12{F0: 1}
                    _ = fi109
                    _ = v110
                }
                _ = v108
            }
            for fi111 := 0; fi111 < 12; fi111++ {
                switch fi111 {
                case -1:
                    v45 = "a"
                default:
                    _ = len(v48)
                }
                _ = fi111
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v52
}
