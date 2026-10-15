package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 int; F1 int32; F2 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int; F1 uint16 }

func main() {
    var v1 int64 = 42
    var v2 int = -9223372036854775808
    var v3 float64 = (3.14159 * 0.0)
    var v4 uint64 = 1
    var v5 int32 = 0
    var v6 uint64 = 1
    v7 := GenS1{F0: true, F1: 2147483647}
    v8 := GenS1{F0: (true && ((!((v6 >= 4294967295) || (("hello" + ("foo bar" + ("a" + (("hello" + ("foo bar" + (("hello" + "a") + "foo bar"))) + "")))) > "\n\t"))) && (2 >= (v6 * (v4 - 222))))), F1: 136}
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: v5, F1: 96, F2: v6}
        v6 += v6
        _ = fi9
        _ = v10
    }
    if false {
        v11 := []int{(-(v2)), v2, (v2 << ((519) & 63)), (v2 / ((v2) | 1))}
        v12 := GenS1{F0: false, F1: 0}
        _ = v11
        _ = v12
    } else {
        v13 := GenS3{F0: (v1 & (v1 | -500))}
        _ = v13
    }
    var v14 int = (v2 - v2)
    fmt.Println(v2)
    v2 = 127
    var v15 uint16 = 65535
    v2 -= v2
    if ("foo bar" <= "") {
        var v16 int = v2
        var v17 float32 = 0.1
        if (v17 != float32(float64(((((-2.25 / (((v17 / v17) / v17) / v17)) * v17) + 3.14159) * v17)))) {
            v18 := GenS4{F0: -1}
            if ("" > "\n\t") {
                var v19 uint = 898
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 int32 = v5
                    _ = fi20
                    _ = v21
                }
                v2 -= (v16 | -2147483648)
                _ = v19
            }
            _ = v18
        } else {
            var v22 uint = 255
            switch (v2 | -2147483648) {
            case 0:
                v4 = uint64(v15)
            default:
                switch (("foo bar" + "\n\t") + "foo bar") {
                case "":
                    switch ("hello" + "hello") {
                    case "\n\t":
                        if false {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                v16 -= v16
                                if false {
                                    for fi24 := 0; fi24 < 18; fi24++ {
                                        v25 := GenS5{F0: 1}
                                        fmt.Println(v3, v17)
                                        v5 = v5
                                        switch (v14 * int(v4)) {
                                        case 128:
                                            fmt.Println(v6)
                                        case 1:
                                            fmt.Println(v15, v5, v1)
                                        case 9223372036854775807:
                                            v26 := GenS4{F0: 0}
                                            _ = v26
                                        default:
                                            if (255 < 1) {
                                                var v27 float64 = (v3 / 0.1)
                                                v28 := GenS6{F0: 1, F1: (v5 << ((234) & 63))}
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := []int{v16, 826, v2, -550}
                                            var v30 string = "hello"
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = fi24
                                        _ = v25
                                    }
                                    switch "a" {
                                    case "":
                                        fmt.Println(v17, fi23, v3)
                                    default:
                                        switch fi23 {
                                        case 4294967295:
                                            if false {
                                                var v31 int16 = 1
                                                _ = v31
                                            }
                                        }
                                    }
                                    var v32 int64 = v1
                                    _ = v32
                                } else {
                                    switch "a" {
                                    case "hello":
                                        switch "" {
                                        case "hello":
                                            switch (fi23 - (v16 ^ 9223372036854775807)) {
                                            case -1:
                                                if true {
                                                    var v33 uint8 = 255
                                                    _ = v33
                                                } else {
                                                    v17 = ((1e10 * v17) + 1e10)
                                                    var v34 int16 = 1
                                                    v35 := []int{v2, (v14 - fi23), v16, -179}
                                                    _ = v34
                                                    _ = v35
                                                }
                                            case 1:
                                                var v36 uint64 = (v6 >> ((255) & 63))
                                                _ = v36
                                            default:
                                                var v37 float64 = v3
                                                _ = v37
                                            }
                                        default:
                                            var v38 float64 = 0.0
                                            _ = v38
                                        }
                                    }
                                    if false {
                                        fmt.Println(v1, v17, v5)
                                    }
                                }
                                _ = fi23
                            }
                            switch (v16 >> (((v22 - v22)) & 63)) {
                            case 907:
                                var v39 uint64 = 2
                                _ = v39
                            case 0:
                                v40 := GenS7{F0: (v1 * (^(v1))), F1: (v4 >> ((v22) & 63))}
                                _ = v40
                            case -1:
                                if true {
                                    var v41 bool = (3.14159 == (v17 * (1.5 / 0.1)))
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v3, v14, v6)
                                    }
                                    v4 = (v6 & v6)
                                    _ = v41
                                } else {
                                    v42 := []int{(v2 - v16)}
                                    fmt.Println(v5)
                                    if false {
                                        var v43 int16 = 624
                                        var v44 uint32 = 239
                                        for fi45 := 0; fi45 < 15; fi45++ {
                                            for fi46 := 0; fi46 < 13; fi46++ {
                                                switch "" {
                                                case "":
                                                    var v47 uint = v22
                                                    _ = v47
                                                case "foo bar":
                                                    var v48 bool = false
                                                    _ = v48
                                                case "\n\t":
                                                    v42 = append(v42, 42)
                                                }
                                                fmt.Println(v43)
                                                if (v43 <= (v43 / (((-(v43))) | 1))) {
                                                    v2 = (-(v16))
                                                    var v49 uint32 = 1
                                                    _ = v49
                                                } else {
                                                    if ("foo bar" > "a") {
                                                        v4 += v4
                                                        v50 := GenS8{F0: v22, F1: 475}
                                                        v16 = v2
                                                        _ = v50
                                                    } else {
                                                        v51 := GenS9{F0: v2, F1: 0, F2: (v1 & -249)}
                                                        _ = v51
                                                    }
                                                    if true {
                                                        v22 = 42
                                                        if len(v42) >= 2 {
                                                            v42 = v42[1:len(v42)]
                                                        }
                                                        v52 := GenS10{F0: v5}
                                                        _ = v52
                                                    } else {
                                                        _ = len(v42)
                                                        v1 += -474
                                                    }
                                                    v17 = ((1e10 * v17) / float32((float64((1e10 - v17)) / v3)))
                                                }
                                                _ = fi46
                                            }
                                            if false {
                                                v17 += 1e10
                                                switch (("" + ("a" + "\n\t")) + "hello") {
                                                case "\n\t":
                                                    var v53 string = "hello"
                                                    _ = v53
                                                case "":
                                                    if (28 == 1) {
                                                        v54 := GenS10{F0: v5}
                                                        v55 := GenS11{F0: -128, F1: false}
                                                        for fi56 := 0; fi56 < 18; fi56++ {
                                                            if (v3 <= 0.0) {
                                                                _ = len(v42)
                                                                var v57 int32 = v5
                                                                _ = v57
                                                            }
                                                            _ = fi56
                                                        }
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        v1 = v1
                                                        v58 := GenS10{F0: -1}
                                                        var v59 float32 = v17
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                case "hello":
                                                    var v60 int8 = -1
                                                    _ = v60
                                                default:
                                                    v2 = -989
                                                }
                                                if true {
                                                    var v61 int8 = 1
                                                    if len(v42) >= 2 {
                                                        v42 = v42[1:len(v42)]
                                                    }
                                                    if len(v42) >= 2 {
                                                        v42 = v42[1:len(v42)]
                                                    }
                                                    _ = v61
                                                } else {
                                                    if ("\n\t" == "") {
                                                        switch "hello" {
                                                        case "":
                                                            if true {
                                                                var v62 int8 = -122
                                                                _ = v62
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v22, v1, v44)
                                                        case "foo bar":
                                                            switch (("hello" + ("" + ("hello" + (("\n\t" + "hello") + ("\n\t" + "hello"))))) + "") {
                                                            case "\n\t":
                                                                v3 = (float64(((v17 / v17) - v17)) + v3)
                                                            }
                                                        }
                                                        for fi63 := 0; fi63 < 5; fi63++ {
                                                            if len(v42) > 0 {
                                                                v42[0] = v14
                                                            }
                                                            var v64 int = (v16 + v14)
                                                            var v65 uint64 = v4
                                                            _ = fi63
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        v22 = v22
                                                    } else {
                                                        fmt.Println(v15, v15)
                                                    }
                                                }
                                            }
                                            v66 := GenS12{F0: (v4 ^ 2)}
                                            _ = fi45
                                            _ = v66
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v42
                                }
                            }
                            v67 := []int{(v14 - (-(v2))), (v16 << (((v22 / ((v22) | 1))) & 63)), (v2 | (v2 ^ int(v15)))}
                            _ = v67
                        } else {
                            var v68 uint64 = v4
                            var v69 uint16 = (v15 >> (((v22 | (^(v22)))) & 63))
                            _ = v68
                            _ = v69
                        }
                    case "":
                        for fi70 := 0; fi70 < 7; fi70++ {
                            if (255 == 255) {
                                v71 := []int{fi70, v14}
                                _ = v71
                            }
                            v72 := []int{}
                            _ = fi70
                            _ = v72
                        }
                    case "foo bar":
                        if false {
                            for fi73 := 0; fi73 < 11; fi73++ {
                                v14 = (fi73 * v16)
                                for fi74 := 0; fi74 < 18; fi74++ {
                                    var v75 uint64 = v6
                                    v6 -= 155
                                    _ = fi74
                                    _ = v75
                                }
                                var v76 int16 = -1
                                _ = fi73
                                _ = v76
                            }
                        }
                    }
                case "a":
                    var v77 uint = v22
                    _ = v77
                case "hello":
                    v78 := GenS13{F0: v16, F1: v15}
                    _ = v78
                default:
                    v15 -= (v15 ^ v15)
                }
            }
            if (42 >= 255) {
                if ((!((false && (2 != uint32(v4))) && (1 == -54))) && (true && (v3 != (v3 * 0.0)))) {
                    var v79 float32 = float32((0.0 / -2.25))
                    v80 := []int{v2, v14, v14, v2}
                    _ = v79
                    _ = v80
                }
            } else {
                if true {
                    var v81 uint32 = 44
                    _ = v81
                } else {
                    var v82 uint64 = 18446744073709551615
                    for fi83 := 0; fi83 < 7; fi83++ {
                        v22 += 18446744073709551615
                        var v84 uint8 = 42
                        var v85 string = ""
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                    _ = v82
                }
                v86 := []int{(v14 % (((-(v2))) | 1)), 127, v16, (v14 / (((v14 & (v16 % ((4294967295) | 1)))) | 1))}
                fmt.Println(v1, v6, v14)
                _ = v86
            }
            _ = v22
        }
        _ = v16
        _ = v17
    } else {
        var v87 uint16 = (v15 % (((v15 - (v15 / (((v15 - (v15 | v15))) | 1)))) | 1))
        _ = v87
    }
    v88 := []int{v14, v2}
    for fi89 := 0; fi89 < 5; fi89++ {
        v5 -= int32(v15)
        v1 = v1
        _ = fi89
    }
    if len(v88) >= 2 {
        v88 = v88[1:len(v88)]
    }
    var v90 float64 = (float64(3.14159) * v3)
    v91 := GenS7{F0: -622, F1: 6}
    if false {
        v92 := GenS5{F0: 2}
        _ = v92
    }
    var v93 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v88:", len(v88), v88)
    fmt.Println("v90:", v90)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v88
    _ = v90
    _ = v91
    _ = v93
}
