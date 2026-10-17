package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{uint32(v2), 65535, 42, 42}
    v1 = (v1 % (((v1 | (v1 << ((2) & 63)))) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 int16 = 1
        _ = fi4
        _ = v5
    }
    v1 -= (v1 - v1)
    switch -200 {
    case -2147483648:
        v1 -= v1
    case -9223372036854775808:
        var v6 int16 = 1
        _ = v6
    }
    v7 := GenS1{F0: 65535}
    var v8 uint8 = 62
    v2 += (v2 - v2)
    if ((v2 & 0) != v2) {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch (fi9 % ((fi9) | 1)) {
            case -9223372036854775808:
                v1 += 100
            default:
                v2 = 301
            }
            _ = fi9
        }
        v2 -= 42
    }
    switch 942 {
    case -2147483648:
        v10 := GenS1{F0: 2147483648}
        _ = v10
    default:
        for fi11 := 0; fi11 < 2; fi11++ {
            _ = len(v3)
            var v12 string = "a"
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int = fi14
            switch fi13 {
            case -256:
                v2 = (^(v2))
            case 836:
                if true {
                    v16 := GenS1{F0: 2147483648}
                    _ = v16
                } else {
                    v2 += uint64(v8)
                    v2 -= v2
                }
            default:
                var v17 float64 = 0.0
                _ = v17
            }
            switch "hello" {
            case "hello":
                var v18 uint8 = (v8 << ((255) & 63))
                _ = v18
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 0
                }
            default:
                v1 -= (v1 ^ v1)
            }
            if (true && (-9223372036854775808 != 9223372036854775807)) {
                var v19 bool = true
                var v20 float32 = 0.1
                if false {
                    if (v19 || (!v19)) {
                        if v19 {
                            fmt.Println(fi14)
                            var v21 uint64 = 42
                            v22 := GenS2{F0: v15, F1: -1, F2: 18446744073709551615}
                            _ = v21
                            _ = v22
                        }
                        var v23 int16 = 42
                        if (!v19) {
                            var v24 float32 = v20
                            var v25 uint64 = v2
                            v26 := GenS3{F0: 200, F1: (v25 / ((807) | 1)), F2: v19}
                            switch 1 {
                            case 128:
                                fmt.Println(v20, v1, v1)
                            default:
                                v25 = (^(v25))
                                fmt.Println(v1)
                                if (-1 != 518) {
                                    if ("\n\t" < ("a" + "\n\t")) {
                                        var v27 int8 = (v1 | (-(v1)))
                                        v25 -= (v2 >> ((65535) & 63))
                                        _ = v27
                                    }
                                } else {
                                    v28 := GenS3{F0: -1, F1: (v2 % (((v2 | (v2 * (v25 >> ((913) & 63))))) | 1)), F2: ((v15 / ((128) | 1)) <= (v15 | v15))}
                                    _ = v28
                                }
                            }
                            switch (fi13 ^ 699) {
                            case 9223372036854775807:
                                v15 = (fi14 % ((-2147483648) | 1))
                            case 217:
                                v29 := GenS2{F0: (^(fi14)), F1: 854, F2: 0}
                                _ = v29
                            }
                            v1 = 120
                            if (!(2147483647 >= 0)) {
                                var v30 bool = false
                                v31 := GenS2{F0: v15, F1: 1, F2: 255}
                                _ = v30
                                _ = v31
                            }
                            var v32 uint64 = v25
                            if ((((3.14159 + 1.5) / 0.1) <= 0.0) && v19) {
                                v24 -= -2.25
                                switch 32767 {
                                case -404:
                                    _ = len(v3)
                                case 128:
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                case 32768:
                                    v1 -= 51
                                default:
                                    switch (v15 * fi14) {
                                    case 42:
                                        v33 := GenS2{F0: fi14, F1: -545, F2: 1}
                                        _ = v33
                                    case -1:
                                        v8 -= 1
                                    }
                                }
                                v24 = v24
                            }
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v32
                        } else {
                            if (!v19) {
                                if true {
                                    v8 = v8
                                    v34 := GenS4{F0: (("a" + "a") + ""), F1: 4294967295}
                                    v35 := GenS5{F0: v23, F1: 2}
                                    v19 = (42 > 1)
                                    v8 += 221
                                    _ = v34
                                    _ = v35
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v36 int = fi14
                                switch (-(fi13)) {
                                case 42:
                                    if false {
                                        v36 = fi13
                                        v36 = -9223372036854775808
                                        switch "a" {
                                        case "a":
                                            if v19 {
                                                v2 = (v2 | (v2 * 42))
                                            } else {
                                                v8 = v8
                                            }
                                        case "hello":
                                            v37 := GenS4{F0: ("hello" + "foo bar"), F1: (v36 | 32767)}
                                            _ = v37
                                        default:
                                            var v38 int32 = -2147483648
                                            _ = v38
                                        }
                                    }
                                case 739:
                                    fmt.Println(v2, v1)
                                }
                                _ = v36
                            } else {
                                if true {
                                    v39 := GenS3{F0: 1, F1: (v2 / ((0) | 1)), F2: v19}
                                    if false {
                                        fmt.Println(v2, v23)
                                    } else {
                                        if true {
                                            fmt.Println(v19, v20)
                                            v1 += (v1 * v1)
                                        }
                                    }
                                    v20 -= 1.5
                                    _ = v39
                                } else {
                                    switch ((((("" + "foo bar") + "\n\t") + ("foo bar" + "foo bar")) + "\n\t") + "") {
                                    case "foo bar":
                                        v20 = (0.1 / float32(1.5))
                                    case "\n\t":
                                        v40 := GenS4{F0: "", F1: 2147483647}
                                        _ = v40
                                    case "a":
                                        switch (fi14 | 127) {
                                        case 42:
                                            var v41 int8 = v1
                                            var v42 uint8 = (v8 % ((v8) | 1))
                                            v23 = (v23 + v23)
                                            _ = v41
                                            _ = v42
                                        }
                                    default:
                                        if (!true) {
                                            v43 := GenS2{F0: v15, F1: -617, F2: 2}
                                            if len(v3) > 0 {
                                                v3[0] = 4294967295
                                            }
                                            var v44 uint32 = 42
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                }
                                var v45 int16 = v23
                                _ = v45
                            }
                            var v46 uint16 = 42
                            _ = v46
                        }
                        var v47 uint8 = (v8 | 22)
                        _ = v23
                        _ = v47
                    } else {
                        var v48 float32 = 0.1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v48
                    }
                    fmt.Println(fi13, v19, v2)
                    var v49 int16 = 758
                    if (-908 == 2147483647) {
                        v50 := GenS3{F0: -1, F1: 42, F2: v19}
                        fmt.Println(v15, v2, v8)
                        v49 -= 1000
                        _ = v50
                    }
                    _ = v49
                } else {
                    v15 += fi13
                    var v51 int16 = -797
                    if (!(118 >= 649)) {
                        v51 -= (v51 - (v51 >> ((255) & 63)))
                    } else {
                        if v19 {
                            v19 = v19
                            var v52 uint = 65535
                            _ = v52
                        }
                        _ = len(v3)
                        switch 2147483647 {
                        case -779:
                            var v53 uint8 = v8
                            _ = v53
                        default:
                            if true {
                                if v19 {
                                    v54 := GenS2{F0: int(v51), F1: -92, F2: 1}
                                    var v55 bool = v19
                                    v56 := GenS4{F0: "foo bar", F1: (v15 ^ fi13)}
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                } else {
                                    v8 += v8
                                    if v19 {
                                        var v57 int32 = 0
                                        var v58 string = ("foo bar" + "")
                                        _ = v57
                                        _ = v58
                                    }
                                }
                                v59 := GenS6{F0: -9223372036854775808, F1: 637}
                                fmt.Println(v2, v15)
                                _ = v59
                            } else {
                                v60 := GenS6{F0: int64(v1), F1: 18446744073709551615}
                                fmt.Println(v2, v19)
                                _ = v60
                            }
                        }
                        var v61 uint16 = 517
                        v8 += v8
                        _ = v61
                    }
                    _ = v51
                }
                _ = v19
                _ = v20
            } else {
                v62 := GenS5{F0: 42, F1: 0}
                if ((v8 * 213) != v8) {
                    var v63 uint64 = (^(v2))
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        if (v63 > v2) {
                            var v64 int = v15
                            v3 = append(v3, 48)
                            var v65 int8 = v1
                            v15 = fi13
                            v63 = 18446744073709551615
                            v66 := GenS3{F0: 1, F1: (v63 ^ (v2 % ((18446744073709551615) | 1))), F2: (!true)}
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                        if (true || (!(false || (!(false && (uint(v63) != 0)))))) {
                            if true {
                                if (true || false) {
                                    switch fi14 {
                                    case -1:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = len(v3)
                                        v1 = (v1 % ((123) | 1))
                                    default:
                                        var v67 int16 = 1
                                        var v68 int16 = (v67 << ((uint(v8)) & 63))
                                        if false {
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v15)
                                            case "a":
                                                var v69 uint = 255
                                                _ = v69
                                            default:
                                                v68 -= v68
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if ((v15 | v15) >= v15) {
                                                if (((float64(0.1) - -2.25) + ((0.1 * float64((float32(-2.25) * -2.25))) / (3.14159 / 0.1))) == 3.14159) {
                                                    if false {
                                                        var v70 uint16 = 42
                                                        _ = v70
                                                    } else {
                                                        if false {
                                                            var v71 uint = 480
                                                            _ = v71
                                                        }
                                                        var v72 uint = 313
                                                        if false {
                                                            if false {
                                                                var v73 uint64 = (v63 + v2)
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if true {
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "":
                                                                            v3 = append(v3, 0)
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    _ = len(v3)
                                                                default:
                                                                    var v74 uint = 461
                                                                    _ = v74
                                                                }
                                                                var v75 uint16 = 65535
                                                                v76 := GenS7{F0: v75}
                                                                _ = v73
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            var v77 float64 = 0.1
                                                            v78 := GenS8{F0: 1, F1: (^(v2)), F2: 9223372036854775807}
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            v79 := GenS6{F0: 42, F1: (v72 & (^(v72)))}
                                                            v80 := GenS1{F0: 145}
                                                            fmt.Println(fi13, v72)
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        var v81 int32 = int32(v8)
                                                        _ = v72
                                                        _ = v81
                                                    }
                                                    v67 += v68
                                                    var v82 uint = 42
                                                    _ = v82
                                                }
                                            }
                                        }
                                        _ = v67
                                        _ = v68
                                    }
                                }
                                switch fi13 {
                                case 821:
                                    var v83 int16 = 589
                                    _ = v83
                                default:
                                    v2 -= v63
                                }
                                v84 := GenS9{F0: 541, F1: (^(v1))}
                                v85 := GenS10{F0: (-(v1)), F1: (fi14 << ((511) & 63))}
                                _ = v84
                                _ = v85
                            } else {
                                fmt.Println(fi13)
                                var v86 bool = false
                                if (v8 >= v8) {
                                    switch fi13 {
                                    case 892:
                                        v2 = (v63 << ((18446744073709551615) & 63))
                                    case -385:
                                        fmt.Println(v8, v63, fi13)
                                    default:
                                        v87 := GenS1{F0: 168}
                                        _ = v87
                                    }
                                    _ = len(v3)
                                    var v88 float64 = (((3.14159 + 1.5) * 0.0) + 0.0)
                                    _ = v88
                                } else {
                                    switch (("\n\t" + "\n\t") + "hello") {
                                    case "":
                                        fmt.Println(v63, v8, fi13)
                                    }
                                    v8 += 179
                                }
                                _ = v86
                            }
                            v3 = append(v3, 65535)
                            v15 = (fi14 | (-(fi13)))
                            var v89 uint64 = (v63 % ((48) | 1))
                            _ = v89
                        }
                    } else {
                        if (2147483647 != 2147483647) {
                            var v90 int8 = v1
                            var v91 uint32 = 2147483647
                            var v92 uint64 = (v63 / ((2) | 1))
                            _ = v90
                            _ = v91
                            _ = v92
                        }
                        v2 = 2
                    }
                    v93 := GenS11{F0: v1, F1: (0 != -415)}
                    _ = v63
                    _ = v93
                }
                _ = v62
            }
            _ = fi14
            _ = v15
        }
        v2 += (^(v2))
        _ = fi13
    }
    v1 -= -1
    v94 := GenS12{F0: 65535, F1: v8}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v94
}
