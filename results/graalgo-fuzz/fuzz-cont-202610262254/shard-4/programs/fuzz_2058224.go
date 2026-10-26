package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int8; F1 int; F2 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int = 2147483647
    v3 := []int{}
    var v4 float32 = (1.5 + -2.25)
    var v5 uint32 = 44
    if (1 > 42) {
        v2 += v2
        var v6 uint64 = 18446744073709551615
        if (-1 >= 333) {
            if true {
                v7 := GenS1{F0: v5, F1: ("a" + "foo bar"), F2: true}
                _ = v7
            }
            var v8 uint32 = v5
            _ = v8
        }
        _ = v6
    } else {
        switch "\n\t" {
        case "foo bar":
            if false {
                v9 := GenS1{F0: v5, F1: "", F2: true}
                _ = v9
            }
        case "hello":
            var v10 uint32 = 42
            _ = v10
        default:
            var v11 uint64 = 4294967295
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            v4 += 1e10
            _ = fi12
        }
    }
    if true {
        var v13 uint32 = (^(v5))
        fmt.Println(v5, v4)
        v4 += v4
        _ = v13
    } else {
        v5 = (v5 - v5)
    }
    var v14 int = (v2 + (v2 << ((65535) & 63)))
    if true {
        var v15 string = ("hello" + "a")
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS1{F0: v5, F1: (("\n\t" + "") + ("\n\t" + "foo bar")), F2: false}
            if true {
                v5 -= (v5 / ((v5) | 1))
                _ = len(v3)
                var v19 float64 = ((0.1 - 3.14159) / 1.5)
                v20 := GenS1{F0: (v5 % ((v5) | 1)), F1: ("a" + "a"), F2: true}
                var v21 float64 = 1e10
                v14 += int(v5)
                v21 -= 1.5
                _ = len(v1)
                _ = v19
                _ = v20
                _ = v21
            }
            if false {
                if (100 < 24) {
                    v2 = v14
                    if (((1.5 - ((0.1 / (((0.1 - (((v4 / v4) - float32(-2.25)) + (0.0 / (v4 * 3.14159)))) * (float32((0.0 - 1.5)) - v4)) * 0.1)) / (0.1 * v4))) / v4) < v4) {
                        v4 = v4
                    }
                    delete(v1, "a")
                    v22 := GenS2{F0: 1, F1: (738 >= 42), F2: (v5 ^ (v5 - 4294967295))}
                    if (!false) {
                        v5 -= (^(v5))
                        var v23 int = (v14 * (-(fi16)))
                        v14 -= (v14 + 4294967295)
                        _ = v23
                    } else {
                        v2 += (fi16 & v2)
                    }
                    if (3.14159 == 1e10) {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            var v24 uint32 = 859
                            _ = v24
                        case "foo bar":
                            switch ("\n\t" + ("foo bar" + "")) {
                            case "a":
                                if ((false && false) || (!((-2147483648 > -1) || (945 >= 544)))) {
                                    switch "foo bar" {
                                    case "hello":
                                        if false {
                                            v2 = (fi16 | (v2 << ((903) & 63)))
                                            if (!false) {
                                                var v25 uint = 2
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        if false {
                                                            if (42 == 42) {
                                                                var v26 int = v14
                                                                _ = v26
                                                            }
                                                            var v27 float64 = 3.14159
                                                            v28 := v1["hello"]
                                                            _ = v27
                                                            _ = v28
                                                        } else {
                                                            switch "a" {
                                                            case "foo bar":
                                                                v4 -= v4
                                                            default:
                                                                delete(v1, ("" + ""))
                                                            }
                                                            if false {
                                                                var v29 int32 = 0
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v29
                                                            }
                                                        }
                                                    case "foo bar":
                                                        switch (fi16 % (((v2 ^ -1)) | 1)) {
                                                        case 127:
                                                            v5 -= uint32(v25)
                                                        case 0:
                                                            var v30 uint16 = uint16(v5)
                                                            _ = v30
                                                        default:
                                                            v14 = (^(fi17))
                                                        }
                                                    case "\n\t":
                                                        v2 = int(v5)
                                                    }
                                                    var v31 string = "foo bar"
                                                    v32 := GenS2{F0: 128, F1: true, F2: v5}
                                                    var v33 uint32 = uint32(v25)
                                                    v2 -= (v2 ^ -334)
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v25
                                            } else {
                                                v4 = v4
                                                var v34 uint8 = 121
                                                _ = v34
                                            }
                                            switch v14 {
                                            case 253:
                                                switch "a" {
                                                case "hello":
                                                    v35 := GenS2{F0: 1, F1: false, F2: v5}
                                                    v36 := GenS2{F0: uint8(v5), F1: (2147483647 <= v14), F2: v5}
                                                    v5 = v5
                                                    v14 -= (-(fi17))
                                                    fmt.Println(v14, fi17, v5)
                                                    _ = len(v3)
                                                    _ = len(v1)
                                                    _ = v35
                                                    _ = v36
                                                case "a":
                                                    var v37 int8 = 0
                                                    _ = v37
                                                default:
                                                    v3 = append(v3, 753)
                                                }
                                            case -393:
                                                if (0 <= 2147483647) {
                                                    v14 = (fi16 | v2)
                                                    v14 = fi17
                                                    var v38 uint32 = (^(v5))
                                                    v5 = 1
                                                    v38 = (v38 << ((65535) & 63))
                                                    var v39 uint8 = 1
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    _ = len(v3)
                                                }
                                            default:
                                                var v40 int8 = -1
                                                _ = v40
                                            }
                                        }
                                    case "foo bar":
                                        if ("foo bar" <= ("" + "foo bar")) {
                                            v41 := GenS1{F0: (v5 & (^(v5))), F1: "foo bar", F2: (uint16(v5) < 411)}
                                            if (!(!(!(((42 != v5) && ("foo bar" != "hello")) || true)))) {
                                                v14 = (fi17 - v14)
                                                _ = len(v1)
                                                var v42 float32 = (v4 * v4)
                                                _ = v42
                                            } else {
                                                if false {
                                                    v1[("foo bar" + (("foo bar" + (("" + "a") + ("hello" + ("" + ("\n\t" + "hello"))))) + "a"))] = v2
                                                } else {
                                                    fmt.Println(fi17)
                                                }
                                            }
                                            _ = v41
                                        }
                                    case "a":
                                        var v43 uint64 = 569
                                        _ = v43
                                    }
                                    switch (("a" + "foo bar") + "hello") {
                                    case "hello":
                                        v44 := GenS3{F0: 42, F1: true, F2: fi17}
                                        _ = v44
                                    case "a":
                                        switch (v2 >> ((4294967295) & 63)) {
                                        case -1:
                                            v5 += (v5 * v5)
                                        default:
                                            v14 -= 42
                                        }
                                    case "foo bar":
                                        v45 := GenS3{F0: 42, F1: false, F2: (fi17 & (^(fi17)))}
                                        _ = v45
                                    default:
                                        v46 := GenS1{F0: 375, F1: (("a" + ("a" + "\n\t")) + ""), F2: true}
                                        _ = v46
                                    }
                                    v47 := GenS4{F0: 0, F1: 693}
                                    _ = v47
                                }
                            case "hello":
                                switch 32768 {
                                case -123:
                                    fmt.Println(fi17, v4)
                                default:
                                    fmt.Println(v4, v14)
                                }
                            case "\n\t":
                                v14 = (v14 >> ((2) & 63))
                            }
                        default:
                            v4 += v4
                        }
                    }
                    v48 := GenS3{F0: 9223372036854775807, F1: false, F2: (-(fi16))}
                    if false {
                        v49 := GenS5{F0: v5, F1: 65535}
                        v1["\n\t"] = fi16
                        _ = v49
                    } else {
                        var v50 int16 = -1
                        _ = v50
                    }
                    _ = v22
                    _ = v48
                }
                var v51 uint8 = 128
                var v52 bool = false
                _ = v51
                _ = v52
            }
            _ = fi17
            _ = v18
        }
        switch v2 {
        case 884:
            var v53 int = v14
            _ = v53
        case 2147483647:
            switch "a" {
            case "":
                var v54 int64 = -366
                _ = v54
            case "a":
                for fi55 := 0; fi55 < 20; fi55++ {
                    var v56 int64 = 315
                    if len(v3) > 0 {
                        v3[0] = fi55
                    }
                    _ = len(v1)
                    v57 := GenS2{F0: 127, F1: true, F2: (v5 * 42)}
                    v58 := GenS6{F0: -1, F1: -9223372036854775808, F2: "a"}
                    v4 += 0.0
                    var v59 int32 = -1
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v58
                    _ = v59
                }
            default:
                if true {
                    var v60 int64 = -317
                    _ = v60
                } else {
                    if (true && false) {
                        v61 := GenS1{F0: 42, F1: "hello", F2: (false || (v4 <= (v4 - (((v4 - v4) + 0.1) - v4))))}
                        var v62 int8 = 22
                        _ = v61
                        _ = v62
                    }
                }
            }
        }
        _ = fi16
    }
    var v63 int64 = 1
    v64 := GenS7{F0: (-1 < -1)}
    delete(v1, "hello")
    var v65 uint = 0
    _ = len(v1)
    for fi66 := 0; fi66 < 13; fi66++ {
        switch 9223372036854775807 {
        case 127:
            v67 := GenS3{F0: (v63 * (v63 | int64(v65))), F1: (42 == 100), F2: v14}
            _ = v67
        case 292:
            var v68 uint64 = 42
            _ = v68
        default:
            if (v4 > float32((3.14159 / 1e10))) {
                var v69 float32 = ((v4 + v4) / v4)
                if (42 != -2147483648) {
                    v5 += (v5 << (((v65 & (v65 - (v65 << ((0) & 63))))) & 63))
                    for fi70 := 0; fi70 < 4; fi70++ {
                        if (!(v5 > v5)) {
                            v69 -= (float32(((float64(v4) - (1.5 * 0.0)) + -2.25)) / 1e10)
                            v71 := GenS8{F0: "foo bar"}
                            v69 -= ((v69 + -2.25) - 3.14159)
                            _ = v71
                        } else {
                            var v72 float32 = 0.1
                            v73 := GenS8{F0: "\n\t"}
                            if (-1 == 248) {
                                v65 = 18446744073709551615
                                v74 := GenS7{F0: (0 != -32768)}
                                _ = v74
                            }
                            _ = v72
                            _ = v73
                        }
                        if (-2.25 < (0.0 * ((float64(v4) + float64(v69)) * -2.25))) {
                            if len(v3) > 0 {
                                v3[0] = -2147483648
                            }
                        }
                        v75 := GenS4{F0: 1, F1: 42}
                        v76 := GenS9{F0: (v65 ^ (v65 - (^(v65))))}
                        if true {
                            var v77 int = (fi70 & v14)
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v3)
                            default:
                                v69 = (float32(1.5) - (v4 / (v4 + (v69 - v69))))
                            }
                            _ = v77
                        } else {
                            v69 -= v69
                            var v78 string = "foo bar"
                            _ = v78
                        }
                        _ = fi70
                        _ = v75
                        _ = v76
                    }
                } else {
                    for fi79 := 0; fi79 < 17; fi79++ {
                        v5 -= 2
                        fmt.Println(v63, v5, fi66)
                        _ = fi79
                    }
                }
                var v80 bool = (42 <= (fi66 & (-(fi66))))
                _ = v69
                _ = v80
            }
        }
        if (2147483648 < (^(v5))) {
            v81 := GenS9{F0: (v65 | 1)}
            var v82 uint16 = 42
            _ = v81
            _ = v82
        } else {
            var v83 float64 = 1e10
            for fi84 := 0; fi84 < 14; fi84++ {
                v14 -= v2
                v1[("hello" + "a")] = (fi66 + -1)
                _ = fi84
            }
            v63 = (v63 ^ v63)
            _ = v83
        }
        _ = fi66
    }
    v85 := GenS3{F0: v63, F1: true, F2: (-(v2))}
    _ = len(v3)
    if (v4 <= v4) {
        switch (v2 % ((-652) | 1)) {
        case 4294967295:
            for fi86 := 0; fi86 < 8; fi86++ {
                v87 := GenS3{F0: 42, F1: (true && (!(v5 > (v5 % ((v5) | 1))))), F2: (v14 + (v2 / (((^(fi86))) | 1)))}
                fmt.Println(fi86, v14)
                var v88 uint32 = (v5 / ((v5) | 1))
                _ = fi86
                _ = v87
                _ = v88
            }
        }
        switch v14 {
        case -2147483648:
            v14 = 9223372036854775807
        case 1:
            if len(v3) > 0 {
                v3[0] = (v14 << (((^(v65))) & 63))
            }
        }
    } else {
        fmt.Println(v65)
        fmt.Println(v65)
    }
    if true {
        for fi89 := 0; fi89 < 9; fi89++ {
            if len(v3) > 0 {
                v3[0] = (v14 << ((v65) & 63))
            }
            v4 -= (1.5 / 3.14159)
            for fi90 := 0; fi90 < 13; fi90++ {
                v91 := GenS4{F0: 177, F1: 42}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v92 uint32 = (v5 ^ 65535)
                if false {
                    v65 = (^(v65))
                } else {
                    if true {
                        var v93 uint = v65
                        fmt.Println(fi89, v63, fi89)
                        switch "a" {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v94 := GenS2{F0: uint8(v93), F1: (!true), F2: 2147483648}
                                _ = v94
                            }
                        case "a":
                            v5 = 2147483647
                        case "hello":
                            if (v5 <= v5) {
                                if len(v3) > 0 {
                                    v3[0] = (fi89 >> ((v93) & 63))
                                }
                                var v95 string = ((("foo bar" + "\n\t") + "") + ((("a" + "\n\t") + "\n\t") + "hello"))
                                switch -1 {
                                case 42:
                                    fmt.Println(v14, v92)
                                    var v96 int64 = -251
                                    delete(v1, "foo bar")
                                    _ = v96
                                }
                                _ = v95
                            }
                        default:
                            var v97 int = (v2 ^ fi89)
                            v5 = uint32(v93)
                            var v98 float64 = float64((v4 * v4))
                            _ = v97
                            _ = v98
                        }
                        _ = v93
                    }
                    _ = len(v1)
                }
                v3 = append(v3, -9223372036854775808)
                v63 += (^(v63))
                _ = fi90
                _ = v91
                _ = v92
            }
            _ = fi89
        }
        var v99 uint32 = 2
        _ = v99
    }
    if false {
        v100 := v1["a"]
        v101 := GenS3{F0: 42, F1: (false || true), F2: -1}
        for fi102 := 0; fi102 < 14; fi102++ {
            _ = len(v3)
            _ = fi102
        }
        _ = v100
        _ = v101
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v85
}
