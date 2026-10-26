package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 string }
type GenS11 struct { F0 uint }

func main() {
    var v1 int8 = 1
    var v2 uint16 = 65535
    var v3 string = ("a" + "\n\t")
    v4 := make(map[int]int)
    v1 = (v1 - v1)
    fmt.Println(v1)
    delete(v4, 981)
    if true {
        v4[127] = -799
        var v5 int = 32767
        fmt.Println(v3, v3, v2)
        _ = v5
    } else {
        switch v3 {
        case "a":
            if false {
                v6 := GenS1{F0: 65535, F1: 42}
                var v7 int8 = 42
                _ = v6
                _ = v7
            }
        case "hello":
            var v8 int32 = 0
            _ = v8
        case "foo bar":
            var v9 uint = 934
            _ = v9
        default:
            v3 = (((("foo bar" + "\n\t") + ((v3 + "foo bar") + ((v3 + v3) + "\n\t"))) + ("hello" + v3)) + "a")
        }
    }
    switch "hello" {
    case "a":
        v10 := []int{}
        _ = v10
    default:
        v11 := GenS2{F0: v2, F1: 127}
        _ = v11
    }
    if false {
        switch -934 {
        case 1:
            v1 = (v1 / ((-1) | 1))
        case 128:
            v3 = "\n\t"
        case 32767:
            fmt.Println(v3, v2, v3)
        }
        if ((!false) || false) {
            switch (v3 + v3) {
            case "a":
                switch "hello" {
                case "":
                    v12 := GenS1{F0: 143, F1: 47}
                    _ = v12
                case "a":
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v14 := GenS1{F0: 4294967295, F1: 42}
                        _ = fi13
                        _ = v14
                    }
                case "foo bar":
                    switch (((("hello" + "") + "a") + v3) + ((v3 + ("foo bar" + ("hello" + v3))) + v3)) {
                    case "":
                        switch "hello" {
                        case "hello":
                            switch (v3 + v3) {
                            case "hello":
                                v15 := []int{-252, 1, 0}
                                _ = v15
                            case "":
                                v16 := GenS2{F0: (v2 >> ((42) & 63)), F1: 254}
                                _ = v16
                            case "\n\t":
                                v17 := []int{}
                                _ = v17
                            }
                        case "a":
                            if (!(!(!true))) {
                                if false {
                                    switch 1 {
                                    case 2147483647:
                                        var v18 int32 = 0
                                        _ = v18
                                    case 9223372036854775807:
                                        v19 := v4[-595]
                                        _ = v19
                                    case -119:
                                        var v20 uint32 = 32
                                        _ = v20
                                    }
                                    var v21 int8 = (^(v1))
                                    v2 = (v2 + 424)
                                    _ = v21
                                } else {
                                    if true {
                                        if (true || false) {
                                            var v22 int32 = 42
                                            var v23 string = v3
                                            _ = v22
                                            _ = v23
                                        } else {
                                            for fi24 := 0; fi24 < 6; fi24++ {
                                                var v25 float32 = float32(((1e10 + 0.1) + 0.0))
                                                v25 -= v25
                                                _ = fi24
                                                _ = v25
                                            }
                                            var v26 string = (v3 + ("a" + "hello"))
                                            _ = v26
                                        }
                                        v3 = "\n\t"
                                    }
                                }
                                var v27 int64 = 814
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    for fi29 := 0; fi29 < 6; fi29++ {
                                        if true {
                                            var v30 uint = 65535
                                            v30 -= (v30 + (v30 % ((v30) | 1)))
                                            _ = v30
                                        }
                                        _ = fi29
                                    }
                                    if false {
                                        if true {
                                            for fi31 := 0; fi31 < 19; fi31++ {
                                                _ = len(v4)
                                                v32 := GenS2{F0: v2, F1: 199}
                                                _ = fi31
                                                _ = v32
                                            }
                                            if (477 >= 18446744073709551615) {
                                                var v33 float64 = (3.14159 - 1e10)
                                                _ = v33
                                            } else {
                                                var v34 float64 = 3.14159
                                                var v35 float64 = v34
                                                var v36 int8 = (v1 + -6)
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                    }
                                    if (1 < 0) {
                                        v37 := []int{(fi28 % (((-(fi28))) | 1)), fi28}
                                        fmt.Println(v3, v27)
                                        _ = v37
                                    } else {
                                        switch 0 {
                                        case 0:
                                            for fi38 := 0; fi38 < 3; fi38++ {
                                                v1 += int8(v2)
                                                _ = fi38
                                            }
                                        }
                                    }
                                    _ = fi28
                                }
                                _ = v27
                            }
                        }
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 127:
                            v39 := v4[-137]
                            _ = v39
                        case 1:
                            v1 = (^(v1))
                        case -1:
                            fmt.Println(v3, v1)
                        default:
                            var v40 int = 34
                            _ = v40
                        }
                    case "hello":
                        v1 = -69
                    }
                }
            default:
                v1 += v1
            }
            for fi41 := 0; fi41 < 13; fi41++ {
                if true {
                    v42 := []int{(fi41 & (fi41 & 1))}
                    var v43 int16 = -517
                    for fi44 := 0; fi44 < 7; fi44++ {
                        v45 := GenS2{F0: v2, F1: 147}
                        _ = fi44
                        _ = v45
                    }
                    _ = v42
                    _ = v43
                }
                switch "hello" {
                case "\n\t":
                    fmt.Println(fi41, v3, v1)
                case "foo bar":
                    var v46 int8 = 69
                    _ = v46
                }
                fmt.Println(fi41, v1)
                _ = fi41
            }
            if ((v2 >> ((4294967295) & 63)) != (v2 >> ((573) & 63))) {
                switch 42 {
                case 1:
                    for fi47 := 0; fi47 < 13; fi47++ {
                        var v48 string = "a"
                        var v49 int = (fi47 * fi47)
                        switch -2147483648 {
                        case 2147483647:
                            var v50 uint64 = 42
                            _ = v50
                        case 169:
                            var v51 uint = 42
                            _ = v51
                        }
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                case -2147483648:
                    var v52 int64 = 990
                    _ = v52
                case 128:
                    var v53 uint16 = v2
                    _ = v53
                default:
                    v54 := v4[1]
                    _ = v54
                }
            }
        } else {
            fmt.Println(v1, v3)
            v1 = v1
            if ((v1 > v1) || false) {
                if false {
                    for fi55 := 0; fi55 < 11; fi55++ {
                        v3 = "\n\t"
                        _ = fi55
                    }
                } else {
                    var v56 uint16 = 255
                    _ = v56
                }
                if false {
                    if (1 == 86) {
                        v3 = v3
                        var v57 int64 = 786
                        _ = v57
                    } else {
                        v2 = v2
                        v2 += 558
                        var v58 float64 = 0.1
                        _ = v58
                    }
                }
                v59 := GenS3{F0: 532, F1: 1}
                _ = v59
            } else {
                switch "" {
                case "":
                    for fi60 := 0; fi60 < 1; fi60++ {
                        var v61 int16 = 1000
                        _ = fi60
                        _ = v61
                    }
                case "\n\t":
                    for fi62 := 0; fi62 < 1; fi62++ {
                        for fi63 := 0; fi63 < 8; fi63++ {
                            v64 := v4[235]
                            fmt.Println(fi63, v2, v3)
                            v65 := []int{}
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 18; fi66++ {
                            v3 = (((v3 + v3) + ("hello" + ("a" + ("foo bar" + v3)))) + v3)
                            if false {
                                v3 = v3
                                if (fi62 < fi62) {
                                    if true {
                                        v1 = (v1 - 127)
                                        var v67 bool = false
                                        var v68 bool = v67
                                        _ = v67
                                        _ = v68
                                    } else {
                                        var v69 uint8 = 105
                                        v69 -= 2
                                        v70 := GenS4{F0: 0}
                                        var v71 int8 = (v1 & 1)
                                        var v72 bool = true
                                        var v73 int64 = -1
                                        v71 -= v71
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    }
                                    var v74 uint8 = 1
                                    _ = v74
                                } else {
                                    v2 = v2
                                    v2 -= (v2 + (v2 + 0))
                                    var v75 int8 = v1
                                    _ = v75
                                }
                                fmt.Println(fi62)
                                delete(v4, fi66)
                            } else {
                                v76 := GenS3{F0: 4294967295, F1: -2147483648}
                                v77 := GenS5{F0: 86}
                                v1 = 127
                                v2 = v2
                                var v78 uint32 = 42
                                _ = v76
                                _ = v77
                                _ = v78
                            }
                            _ = fi66
                        }
                        _ = fi62
                    }
                }
                v79 := []int{-2147483648}
                var v80 string = ("a" + "hello")
                _ = v79
                _ = v80
            }
        }
    } else {
        v81 := []int{9223372036854775807, -105, 4294967295, 351}
        var v82 int = -9223372036854775808
        for fi83 := 0; fi83 < 8; fi83++ {
            var v84 uint16 = (^(v2))
            if (v1 > int8(fi83)) {
                v85 := GenS1{F0: 42, F1: 0}
                if false {
                    var v86 float32 = -2.25
                    v87 := GenS2{F0: v2, F1: 238}
                    _ = v86
                    _ = v87
                } else {
                    if (v1 <= (v1 + v1)) {
                        if (-835 != v82) {
                            v84 += (v2 << ((847) & 63))
                            if ((^(fi83)) < 32768) {
                                switch -9223372036854775808 {
                                case 46:
                                    v88 := GenS6{F0: v1}
                                    _ = v88
                                default:
                                    v81 = append(v81, fi83)
                                }
                            }
                            switch (v3 + "") {
                            case "foo bar":
                                v84 = (v84 << ((448) & 63))
                            case "\n\t":
                                if false {
                                    v82 -= v82
                                    switch fi83 {
                                    case 127:
                                        v89 := GenS7{F0: 0, F1: 1}
                                        _ = v89
                                    default:
                                        var v90 uint32 = 1
                                        _ = v90
                                    }
                                    fmt.Println(v84, v3, v2)
                                }
                            case "a":
                                var v91 uint8 = 127
                                _ = v91
                            }
                        }
                    } else {
                        for fi92 := 0; fi92 < 8; fi92++ {
                            switch (v3 + v3) {
                            case "\n\t":
                                v82 = fi92
                            case "foo bar":
                                if false {
                                    v93 := GenS8{F0: 42}
                                    _ = v93
                                } else {
                                    v94 := GenS9{F0: 2147483647}
                                    var v95 string = (v3 + v3)
                                    switch (v3 + "\n\t") {
                                    case "":
                                        if len(v81) >= 2 {
                                            v81 = v81[1:len(v81)]
                                        }
                                        v2 += v84
                                        delete(v4, 9223372036854775807)
                                        v84 -= 304
                                        v82 = -11
                                        if false {
                                            if false {
                                                var v96 float32 = 0.0
                                                _ = v96
                                            } else {
                                                fmt.Println(v3, v95)
                                                v1 = v1
                                            }
                                        } else {
                                            fmt.Println(fi92)
                                        }
                                        fmt.Println(v3)
                                    case "foo bar":
                                        v97 := GenS10{F0: 0, F1: v3}
                                        if ("hello" > v3) {
                                            v98 := GenS11{F0: 162}
                                            _ = v98
                                        } else {
                                            v1 -= v1
                                            switch v95 {
                                            case "\n\t":
                                                var v99 int16 = 0
                                                _ = v99
                                            case "":
                                                fmt.Println(v2, v95)
                                            case "foo bar":
                                                v3 = v3
                                            }
                                            v3 = ("foo bar" + "")
                                        }
                                        fmt.Println(v2, fi83)
                                        v100 := GenS9{F0: 4294967295}
                                        _ = v100
                                        _ = v97
                                    case "\n\t":
                                        if (((-2.25 + 0.1) / 1.0) != 1.5) {
                                            _ = len(v81)
                                            var v101 int16 = 577
                                            _ = v101
                                        } else {
                                            fmt.Println(fi92, fi83)
                                            var v102 uint64 = 18446744073709551615
                                            _ = v102
                                        }
                                    }
                                    _ = v94
                                    _ = v95
                                }
                            default:
                                fmt.Println(v3, fi92)
                            }
                            switch (fi92 % ((0) | 1)) {
                            case 9223372036854775807:
                                v82 += (fi92 << ((255) & 63))
                            case 42:
                                if len(v81) >= 2 {
                                    v81 = v81[1:len(v81)]
                                }
                            case 128:
                                v81 = append(v81, -731)
                            default:
                                switch (v82 * fi83) {
                                case 640:
                                    var v103 float32 = 1e10
                                    switch (^(fi92)) {
                                    case -516:
                                        switch "" {
                                        case "":
                                            var v104 uint8 = 42
                                            _ = v104
                                        case "a":
                                            v4[fi92] = (fi83 | v82)
                                        case "hello":
                                            fmt.Println(fi92, fi92, v84)
                                        default:
                                            var v105 float64 = -2.25
                                            _ = v105
                                        }
                                    case 2147483647:
                                        var v106 uint = 2
                                        _ = v106
                                    case 592:
                                        v1 += (-(v1))
                                    default:
                                        fmt.Println(v82)
                                    }
                                    _ = v103
                                default:
                                    if false {
                                        if false {
                                            v1 -= 0
                                            v1 += (v1 >> ((1) & 63))
                                        }
                                        var v107 int16 = int16(v2)
                                        switch v82 {
                                        case -1:
                                            fmt.Println(fi92, v82)
                                        }
                                        _ = v107
                                    } else {
                                        if (0 == 0) {
                                            v84 -= 42
                                            v84 += 255
                                        }
                                        v108 := v4[(-(fi83))]
                                        _ = v108
                                    }
                                }
                            }
                            _ = fi92
                        }
                    }
                    v4[32768] = fi83
                }
                _ = v85
            } else {
                for fi109 := 0; fi109 < 11; fi109++ {
                    v84 -= 255
                    _ = fi109
                }
            }
            _ = fi83
            _ = v84
        }
        _ = v81
        _ = v82
    }
    if true {
        v1 = (^(v1))
        var v110 uint32 = 2147483647
        switch 0 {
        case 32768:
            v2 = 42
        }
        _ = v110
    }
    v4[32767] = 1
    v111 := v4[4294967295]
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v4
}
