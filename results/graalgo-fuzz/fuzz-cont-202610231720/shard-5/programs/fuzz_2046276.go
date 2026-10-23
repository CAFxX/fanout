package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int8; F1 int8; F2 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 uint; F2 uint8 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint; F1 int8 }
type GenS13 struct { F0 string; F1 uint8; F2 int16 }
type GenS14 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint64 = 195
    v2 := []string{"hello", "\n\t"}
    var v3 int16 = -1
    switch "\n\t" {
    case "a":
        for fi4 := 0; fi4 < 8; fi4++ {
            v1 += v1
            switch "" {
            case "\n\t":
                var v5 uint = 4294967295
                _ = v5
            case "a":
                var v6 int8 = int8(v1)
                _ = v6
            case "foo bar":
                if (true || (!false)) {
                    var v7 bool = (172 >= 255)
                    v7 = (v7 && false)
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v3 = 32767
                        fmt.Println(v1, v1, fi4)
                        var v9 uint32 = 65535
                        _ = fi8
                        _ = v9
                    }
                    var v10 uint = 42
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := GenS1{F0: 65535}
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                }
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                v3 += (v3 % ((1000) | 1))
                _ = fi13
            }
            _ = fi4
        }
    case "hello":
        var v14 int16 = 42
        _ = v14
    case "\n\t":
        v1 = v1
    default:
        if (42 > int32(v3)) {
            if true {
                var v15 float64 = (-2.25 * (-2.25 * 0.0))
                var v16 uint32 = 2147483647
                _ = v15
                _ = v16
            }
        }
    }
    var v17 uint64 = v1
    for fi18 := 0; fi18 < 2; fi18++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 -= (v3 & v3)
        v17 = (v1 + v1)
        _ = fi18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int32 = -663
        _ = fi19
        _ = v20
    }
    v1 += v1
    var v21 uint64 = v1
    switch (("" + "") + "hello") {
    case "foo bar":
        var v22 float32 = -2.25
        _ = v22
    case "hello":
        fmt.Println(v17, v1, v1)
    case "":
        v1 = (^(v21))
    default:
        var v23 uint8 = 128
        _ = v23
    }
    if (!(!(!(v3 >= 42)))) {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        switch -2147483648 {
        case 2147483647:
            for fi24 := 0; fi24 < 9; fi24++ {
                _ = len(v2)
                v21 = (v21 & v17)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi24
            }
        case 128:
            if true {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v17 = (v17 >> ((658) & 63))
                    _ = fi25
                }
                v21 = 406
            } else {
                for fi26 := 0; fi26 < 4; fi26++ {
                    if true {
                        var v27 uint32 = 180
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(v17, v27, v1)
                            v29 := GenS2{F0: (v21 % ((v21) | 1)), F1: "foo bar", F2: -300}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    } else {
                        v3 = (-(v3))
                        var v30 uint64 = 42
                        var v31 uint = 18446744073709551615
                        _ = v30
                        _ = v31
                    }
                    _ = fi26
                }
                v1 -= 663
                var v32 int = -1
                _ = v32
            }
        case 0:
            if (false || false) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 float32 = (-2.25 + ((-2.25 - ((((1e10 + -2.25) - 0.1) / float32((0.1 * 1e10))) * (3.14159 / -2.25))) - (-2.25 * (((0.1 - 1.5) / 1e10) / -2.25))))
                    if len(v2) > 0 {
                        v2[0] = ("hello" + "hello")
                    }
                    if false {
                        if true {
                            for fi35 := 0; fi35 < 17; fi35++ {
                                if false {
                                    v2 = append(v2, "")
                                    v34 += 0.0
                                } else {
                                    v21 = (v1 + v21)
                                    var v36 float64 = float64(1e10)
                                    _ = v36
                                }
                                v21 += (v17 % (((^(v17))) | 1))
                                switch ("foo bar" + "hello") {
                                case "":
                                    var v37 uint8 = 42
                                    _ = v37
                                case "hello":
                                    v1 = v21
                                }
                                _ = len(v2)
                                v38 := GenS3{F0: 153, F1: v21}
                                _ = fi35
                                _ = v38
                            }
                        } else {
                            switch ((((("a" + ("hello" + ("a" + ("hello" + "\n\t")))) + "foo bar") + ("hello" + "\n\t")) + "hello") + ("\n\t" + (("\n\t" + "hello") + "\n\t"))) {
                            case "a":
                                var v39 int = (fi33 >> ((160) & 63))
                                _ = v39
                            default:
                                v40 := GenS4{F0: false}
                                _ = v40
                            }
                            var v41 int16 = 1000
                            for fi42 := 0; fi42 < 16; fi42++ {
                                var v43 uint32 = 2147483647
                                switch "" {
                                case "hello":
                                    if (false && true) {
                                        v44 := GenS3{F0: 288, F1: v21}
                                        _ = v44
                                    }
                                case "\n\t":
                                    var v45 int16 = 369
                                    _ = len(v2)
                                    v46 := GenS5{F0: 9223372036854775807}
                                    if (true || false) {
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        fmt.Println(v45, v34)
                                    } else {
                                        v47 := GenS5{F0: -9223372036854775808}
                                        var v48 float64 = 0.0
                                        var v49 uint64 = (v17 >> ((623) & 63))
                                        var v50 uint16 = 255
                                        v51 := GenS1{F0: v43}
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v45
                                    _ = v46
                                case "a":
                                    var v52 bool = true
                                    _ = v52
                                }
                                switch "" {
                                case "":
                                    v43 = (^(v43))
                                case "a":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    v1 = (v1 << ((315) & 63))
                                }
                                _ = fi42
                                _ = v43
                            }
                            _ = v41
                        }
                        fmt.Println(v34)
                        switch "hello" {
                        case "a":
                            _ = len(v2)
                        case "foo bar":
                            for fi53 := 0; fi53 < 14; fi53++ {
                                if true {
                                    v54 := GenS2{F0: v17, F1: "", F2: 1}
                                    v2 = append(v2, (("a" + ("" + "hello")) + ("foo bar" + "a")))
                                    if true {
                                        switch (^(fi53)) {
                                        case -1:
                                            v17 = (v17 % (((v17 / ((v17) | 1))) | 1))
                                        default:
                                            v55 := GenS6{F0: fi33, F1: 9223372036854775807}
                                            _ = v55
                                        }
                                        var v56 uint16 = 1
                                        var v57 uint32 = 811
                                        _ = v56
                                        _ = v57
                                    } else {
                                        v58 := GenS6{F0: (fi53 >> ((4294967295) & 63)), F1: 194}
                                        fmt.Println(fi53, v3, fi33)
                                        var v59 int64 = -9223372036854775808
                                        fmt.Println(v17)
                                        var v60 int8 = 84
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = v54
                                } else {
                                    v61 := GenS1{F0: 1}
                                    _ = v61
                                }
                                v1 -= 18446744073709551615
                                v3 = -829
                                _ = fi53
                            }
                        default:
                            var v62 string = "\n\t"
                            _ = v62
                        }
                    } else {
                        for fi63 := 0; fi63 < 16; fi63++ {
                            if false {
                                if true {
                                    _ = len(v2)
                                    v64 := GenS4{F0: true}
                                    if len(v2) > 0 {
                                        v2[0] = ("foo bar" + "\n\t")
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ("\n\t" + "")
                                    }
                                    _ = v64
                                } else {
                                    if false {
                                        if ((!false) && ((-424 > 1) && true)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        var v65 uint = 1
                                        _ = v65
                                    } else {
                                        fmt.Println(v1, v34)
                                        _ = len(v2)
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                    }
                                    switch "" {
                                    case "":
                                        v66 := GenS2{F0: v21, F1: "a", F2: 42}
                                        _ = v66
                                    case "\n\t":
                                        if true {
                                            var v67 int32 = -2147483648
                                            v68 := GenS3{F0: 878, F1: (v1 << ((18446744073709551615) & 63))}
                                            _ = v67
                                            _ = v68
                                        } else {
                                            v69 := GenS7{F0: -128, F1: -128, F2: 2147483647}
                                            _ = v69
                                        }
                                    default:
                                        v70 := GenS4{F0: (false || (!(2 <= 0)))}
                                        _ = v70
                                    }
                                }
                                v17 += (^(v21))
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(fi63)
                                v71 := GenS8{F0: int(v3)}
                                var v72 bool = true
                                var v73 uint = 255
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            _ = fi63
                        }
                        if (true || (v34 == 1.5)) {
                            switch "" {
                            case "\n\t":
                                v74 := GenS6{F0: fi33, F1: -853}
                                _ = v74
                            case "hello":
                                for fi75 := 0; fi75 < 13; fi75++ {
                                    v76 := GenS1{F0: 787}
                                    _ = fi75
                                    _ = v76
                                }
                            case "a":
                                if ((v3 != (v3 * (v3 >> ((175) & 63)))) && false) {
                                    v77 := GenS9{F0: int16(v17), F1: 4294967295, F2: 128}
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    _ = v77
                                }
                            }
                            var v78 int16 = v3
                            _ = v78
                        } else {
                            fmt.Println(fi33)
                            var v79 int16 = v3
                            v80 := GenS10{F0: 42, F1: 128}
                            _ = v79
                            _ = v80
                        }
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                for fi81 := 0; fi81 < 17; fi81++ {
                    var v82 int8 = 100
                    var v83 float64 = ((0.0 / ((0.1 / 1.0) - 1e10)) + 1.5)
                    _ = fi81
                    _ = v82
                    _ = v83
                }
            }
        default:
            switch 2147483647 {
            case -491:
                v84 := GenS11{F0: 326}
                _ = v84
            default:
                v3 -= 32767
            }
        }
        switch "foo bar" {
        case "hello":
            for fi85 := 0; fi85 < 16; fi85++ {
                v21 = v1
                _ = fi85
            }
        }
    } else {
        if (true || (127 >= 0)) {
            for fi86 := 0; fi86 < 19; fi86++ {
                if (42 < 0) {
                    var v87 uint = 418
                    if (1e10 >= float32(float64((-2.25 - (-2.25 - (0.1 / -2.25)))))) {
                        for fi88 := 0; fi88 < 14; fi88++ {
                            var v89 int64 = 44
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                var v90 uint64 = v21
                                v90 += 1
                                _ = v90
                            }
                            _ = fi88
                            _ = v89
                        }
                        _ = len(v2)
                        v1 = (^(v21))
                    } else {
                        var v91 uint64 = v17
                        v92 := GenS12{F0: 42, F1: 42}
                        _ = v91
                        _ = v92
                    }
                    _ = v87
                } else {
                    var v93 uint = 255
                    var v94 uint64 = 0
                    var v95 float64 = float64(0.0)
                    _ = v93
                    _ = v94
                    _ = v95
                }
                _ = fi86
            }
            var v96 float64 = (3.14159 / 1.0)
            fmt.Println(v1, v96)
            _ = v96
        }
    }
    switch ("hello" + "") {
    case "hello":
        var v97 int64 = -1
        _ = v97
    case "foo bar":
        v21 += (v21 % (((v17 + v21)) | 1))
    case "":
        var v98 int8 = 0
        _ = v98
    default:
        v99 := GenS7{F0: 1, F1: 42, F2: 1}
        _ = v99
    }
    switch -1 {
    case 1:
        var v100 int64 = 42
        _ = v100
    case 4294967295:
        var v101 uint64 = v1
        _ = v101
    case -1:
        v17 = (v21 + 18446744073709551615)
    }
    var v102 float64 = ((0.1 + (1e10 / 1.0)) / 3.14159)
    for fi103 := 0; fi103 < 13; fi103++ {
        if true {
            for fi104 := 0; fi104 < 2; fi104++ {
                v2 = append(v2, "hello")
                v105 := GenS8{F0: fi104}
                v106 := GenS13{F0: (("\n\t" + "a") + (("\n\t" + "hello") + "")), F1: 255, F2: v3}
                _ = fi104
                _ = v105
                _ = v106
            }
        } else {
            _ = len(v2)
            if (false || false) {
                var v107 int8 = 127
                fmt.Println(v1, v102, v107)
                _ = v107
            } else {
                v1 = (v1 ^ v1)
            }
            fmt.Println(fi103, v102)
        }
        _ = fi103
    }
    if (42 == 0) {
        v17 -= 18446744073709551615
        var v108 uint8 = 99
        for fi109 := 0; fi109 < 14; fi109++ {
            v3 = -232
            _ = fi109
        }
        _ = v108
    } else {
        if (1 <= 0) {
            v110 := GenS1{F0: 1}
            var v111 float64 = v102
            _ = v110
            _ = v111
        }
        var v112 int = 32767
        _ = v112
    }
    v113 := GenS14{F0: 42, F1: 65535}
    v114 := GenS4{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v102
    _ = v113
    _ = v114
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
