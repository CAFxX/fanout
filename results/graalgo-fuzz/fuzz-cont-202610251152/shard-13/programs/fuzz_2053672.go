package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64; F1 string }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint }
func genFn0() uint32 {
    var v5 int32 = 410
    v6 := []int{}
    fmt.Println(v5)
    if true {
        v7 := GenS1{F0: -72, F1: 1, F2: 1}
        _ = v7
    } else {
        v5 = (v5 << ((4294967295) & 63))
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int = 413
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            var v11 int8 = 100
            _ = v11
        } else {
            fmt.Println(v5)
            var v12 int32 = (v5 >> ((1) & 63))
            _ = v12
        }
        v13 := GenS2{F0: 32767, F1: 640}
        _ = fi10
        _ = v13
    }
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    v3 := []uint32{4294967295, 2}
    v4 := []string{"\n\t", ("" + ("" + ("" + ("a" + "a")))), (("a" + "hello") + ("foo bar" + ""))}
    if false {
        if true {
            if len(v3) > 0 {
                v3[0] = uint32(v2)
            }
        }
        if true {
            v4 = append(v4, "\n\t")
            v14 := GenS3{F0: -53}
            for fi15 := 0; fi15 < 18; fi15++ {
                switch "foo bar" {
                case "hello":
                    v2 = (v2 / ((286) | 1))
                case "\n\t":
                    delete(v1, (("foo bar" + "a") + "foo bar"))
                case "a":
                    switch (fi15 + fi15) {
                    case 1:
                        fmt.Println(fi15, v2)
                    case 32768:
                        var v16 int16 = -858
                        _ = v16
                    case -523:
                        fmt.Println(fi15)
                    }
                default:
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if ((v2 ^ (v2 % ((v2) | 1))) <= v2) {
                            if false {
                                switch "a" {
                                case "hello":
                                    if (!(((!false) && true) && (!false))) {
                                        var v18 int64 = -1
                                        fmt.Println(fi17, v18, fi17)
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(fi15, fi15)
                                            if true {
                                                v19 := GenS3{F0: (v18 % (((v18 | (v18 - 1))) | 1))}
                                                v20 := GenS4{F0: 50, F1: 42}
                                                var v21 string = ("" + "hello")
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            } else {
                                                if false {
                                                    v22 := GenS3{F0: -781}
                                                    if false {
                                                        v18 -= v18
                                                    }
                                                    if (v2 != (v2 * v2)) {
                                                        fmt.Println(fi17, v18, v18)
                                                        var v23 uint = (v2 >> ((v2) & 63))
                                                        _ = v23
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = 567
                                                        }
                                                        fmt.Println(fi15, v18)
                                                    }
                                                    var v24 uint = (v2 % ((v2) | 1))
                                                    v18 = (v18 | v18)
                                                    _ = v22
                                                    _ = v24
                                                } else {
                                                    if (550 == int16(fi17)) {
                                                        if (2 != (v2 - (v2 / ((4294967295) | 1)))) {
                                                            if (((1e10 == ((1.5 / ((1e10 / 0.1) / 1.0)) / (1.5 / 1.5))) || (false && true)) && true) {
                                                                var v25 bool = (0.0 == (3.14159 - float32(float64((-2.25 / (1.5 - ((0.0 * -2.25) + 0.1)))))))
                                                                var v26 uint32 = 1
                                                                _ = v25
                                                                _ = v26
                                                            }
                                                            var v27 uint8 = 37
                                                            v27 += v27
                                                            _ = v27
                                                        }
                                                        v2 -= 82
                                                        fmt.Println(fi15, v18, v18)
                                                    }
                                                }
                                                v2 += v2
                                                switch (fi15 / ((42) | 1)) {
                                                case 359:
                                                    delete(v1, (((("hello" + ((("hello" + (("hello" + "\n\t") + "foo bar")) + "a") + "\n\t")) + "foo bar") + "\n\t") + "hello"))
                                                case 32767:
                                                    fmt.Println(fi17, v18)
                                                case 187:
                                                    v2 += 255
                                                default:
                                                    if (!true) {
                                                        fmt.Println(fi15, v2, fi15)
                                                        fmt.Println(fi15, fi15, v2)
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = 573
                                                        }
                                                        var v28 int32 = 1
                                                        var v29 int32 = v28
                                                        switch "foo bar" {
                                                        case "":
                                                            v18 = (v18 >> (((v2 << ((v2) & 63))) & 63))
                                                            switch (("hello" + "foo bar") + ("hello" + (("a" + (("\n\t" + "") + (("hello" + ("hello" + ("\n\t" + ""))) + ("" + "foo bar")))) + "foo bar"))) {
                                                            case "hello":
                                                                var v30 uint64 = 4294967295
                                                                _ = v30
                                                            default:
                                                                if false {
                                                                    if (false || ((!false) || (42 > -9223372036854775808))) {
                                                                        var v31 bool = false
                                                                        var v32 uint = (^(v2))
                                                                        v32 = v2
                                                                        _ = v31
                                                                        _ = v32
                                                                    } else {
                                                                        if false {
                                                                            if true {
                                                                                v29 = (v28 - v29)
                                                                                if true {
                                                                                    v33 := GenS1{F0: (v18 >> ((v2) & 63)), F1: 909, F2: 18446744073709551615}
                                                                                    v1["a"] = fi17
                                                                                    v34 := GenS5{F0: 2, F1: (^(v2))}
                                                                                    _ = v33
                                                                                    _ = v34
                                                                                } else {
                                                                                    switch (-(fi17)) {
                                                                                    case 32767:
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            if (int8(fi15) != -108) {
                                                                                                switch "a" {
                                                                                                case "\n\t":
                                                                                                    if (true && (false && (!false))) {
                                                                                                        v2 += v2
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    fmt.Println(fi15, v29, v28)
                                                                                                case "a":
                                                                                                    if false {
                                                                                                        switch ("foo bar" + "") {
                                                                                                        case "hello":
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = genFn0()
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            var v35 int32 = -2147483648
                                                                                                            _ = v35
                                                                                                        }
                                                                                                        v18 -= v18
                                                                                                    } else {
                                                                                                        if ("a" >= "\n\t") {
                                                                                                            if false {
                                                                                                                var v36 int32 = (v28 | -1)
                                                                                                                var v37 int8 = int8(v2)
                                                                                                                _ = v36
                                                                                                                _ = v37
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = 912
                                                                                                        }
                                                                                                        v29 -= -2147483648
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            v38 := GenS4{F0: 0, F1: 18446744073709551615}
                                                                                            if (1 < (v2 - v2)) {
                                                                                                var v39 bool = (v2 >= 18446744073709551615)
                                                                                                _ = v39
                                                                                            } else {
                                                                                                if (fi17 <= fi17) {
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = "foo bar"
                                                                                                    }
                                                                                                    v40 := GenS4{F0: 1, F1: 18446744073709551615}
                                                                                                    _ = v40
                                                                                                } else {
                                                                                                    var v41 int64 = v18
                                                                                                    v42 := GenS2{F0: 42, F1: 18446744073709551615}
                                                                                                    if (!(42 == int16(v2))) {
                                                                                                        _ = len(v1)
                                                                                                        v2 = (v2 | v2)
                                                                                                        delete(v1, ("hello" + ""))
                                                                                                    }
                                                                                                    _ = v41
                                                                                                    _ = v42
                                                                                                }
                                                                                                v18 = 9223372036854775807
                                                                                            }
                                                                                            _ = v38
                                                                                        case "":
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = ("\n\t" + "a")
                                                                                            }
                                                                                        }
                                                                                    case 0:
                                                                                        v43 := GenS5{F0: 2, F1: (v2 - v2)}
                                                                                        var v44 float64 = 1e10
                                                                                        if true {
                                                                                            fmt.Println(v44, v29, v29)
                                                                                            fmt.Println(fi15)
                                                                                            v18 -= v18
                                                                                        }
                                                                                        _ = v43
                                                                                        _ = v44
                                                                                    case -2147483648:
                                                                                        v45 := GenS3{F0: (v18 ^ v18)}
                                                                                        _ = v45
                                                                                    }
                                                                                    var v46 uint = (^(v2))
                                                                                    v46 = (v2 * (^(v46)))
                                                                                    _ = v46
                                                                                }
                                                                            } else {
                                                                                if (!false) {
                                                                                    if (((0.1 / 1.0) * 0.0) <= float32(float64((float32((((0.0 / 3.14159) * 0.1) / 1.5)) / float32(0.1))))) {
                                                                                        v4 = append(v4, "foo bar")
                                                                                    }
                                                                                    if (((1 > 0) || (17 > 42)) && (!false)) {
                                                                                        fmt.Println(fi15)
                                                                                    }
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    v18 = -586
                                                                                }
                                                                                fmt.Println(v18, v28)
                                                                                fmt.Println(v29, v2)
                                                                            }
                                                                        }
                                                                        v2 -= v2
                                                                        v18 -= (v18 % ((v18) | 1))
                                                                        if true {
                                                                            var v47 string = "\n\t"
                                                                            delete(v1, v47)
                                                                            _ = v47
                                                                        } else {
                                                                            var v48 float64 = 3.14159
                                                                            _ = v48
                                                                        }
                                                                        v49 := GenS1{F0: v18, F1: 4294967295, F2: 78}
                                                                        _ = v49
                                                                    }
                                                                    v2 = 498
                                                                }
                                                            }
                                                            fmt.Println(fi15)
                                                        case "\n\t":
                                                            var v50 float32 = -2.25
                                                            _ = v50
                                                        default:
                                                            fmt.Println(v18, fi15)
                                                        }
                                                        var v51 float32 = 0.0
                                                        _ = v28
                                                        _ = v29
                                                        _ = v51
                                                    }
                                                }
                                            }
                                            v52 := GenS1{F0: (v18 & 9223372036854775807), F1: 169, F2: 2}
                                            var v53 int16 = 42
                                            _ = v52
                                            _ = v53
                                        case "hello":
                                            var v54 bool = ((!true) || false)
                                            _ = v54
                                        }
                                        _ = v18
                                    } else {
                                        v2 += 1
                                        if (116 < 0) {
                                            v2 = 1
                                            var v55 uint = v2
                                            _ = v55
                                        }
                                        v56 := GenS2{F0: 42, F1: 4294967295}
                                        v57 := GenS6{F0: -9223372036854775808}
                                        _ = v56
                                        _ = v57
                                    }
                                case "\n\t":
                                    delete(v1, "\n\t")
                                case "foo bar":
                                    var v58 int16 = -38
                                    _ = v58
                                default:
                                    if false {
                                        v1["\n\t"] = fi15
                                        var v59 uint8 = 2
                                        v59 = (v59 % ((v59) | 1))
                                        _ = v59
                                    } else {
                                        _ = len(v1)
                                        fmt.Println(fi15, fi17, fi15)
                                    }
                                    v2 = (v2 * (v2 + v2))
                                    switch "foo bar" {
                                    case "foo bar":
                                        v2 = v2
                                        v60 := GenS4{F0: 255, F1: 1}
                                        _ = v60
                                    case "\n\t":
                                        var v61 uint64 = 18446744073709551615
                                        _ = v61
                                    }
                                }
                            } else {
                                v62 := GenS7{F0: 8}
                                var v63 uint = v2
                                v63 -= 153
                                _ = v62
                                _ = v63
                            }
                        }
                        var v64 bool = true
                        var v65 int64 = 0
                        _ = fi17
                        _ = v64
                        _ = v65
                    }
                }
                var v66 int16 = -849
                _ = fi15
                _ = v66
            }
            _ = v14
        }
        delete(v1, "foo bar")
    } else {
        v67 := GenS8{F0: 640, F1: -2147483648}
        v68 := GenS9{F0: 1}
        fmt.Println(v2)
        _ = v67
        _ = v68
    }
    fmt.Println(v2)
    if true {
        v3 = append(v3, 47)
    } else {
        if ((!true) && (128 < 250)) {
            var v69 float64 = -2.25
            _ = len(v1)
            _ = v69
        }
        var v70 bool = (905 < v2)
        if (-2.25 > 0.1) {
            fmt.Println(v70, v70)
            v70 = ((4294967295 < 65535) || (((1e10 * 1.5) <= 0.1) && v70))
        } else {
            for fi71 := 0; fi71 < 14; fi71++ {
                fmt.Println(fi71, v2)
                if len(v3) > 0 {
                    v3[0] = 0
                }
                for fi72 := 0; fi72 < 3; fi72++ {
                    v73 := GenS7{F0: 127}
                    fmt.Println(v70, v70, fi72)
                    _ = fi72
                    _ = v73
                }
                _ = fi71
            }
        }
        _ = v70
    }
    if false {
        switch -9223372036854775808 {
        case 9223372036854775807:
            if len(v4) > 0 {
                v4[0] = ("" + "")
            }
        case 32768:
            for fi74 := 0; fi74 < 19; fi74++ {
                v75 := v1["a"]
                var v76 uint = v2
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi74
                _ = v75
                _ = v76
            }
        }
    } else {
        var v77 uint8 = 159
        var v78 uint16 = 820
        _ = v77
        _ = v78
    }
    if false {
        var v79 uint16 = 65535
        var v80 int64 = -352
        _ = v79
        _ = v80
    }
    if (!true) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi81 := 0; fi81 < 4; fi81++ {
            for fi82 := 0; fi82 < 12; fi82++ {
                var v83 int16 = -1
                switch -901 {
                case 820:
                    var v84 string = "\n\t"
                    _ = v84
                case 9223372036854775807:
                    var v85 int16 = (v83 & (v83 % ((v83) | 1)))
                    _ = v85
                }
                _ = fi82
                _ = v83
            }
            v2 -= (v2 >> ((18446744073709551615) & 63))
            _ = fi81
        }
    }
    switch 127 {
    case 0:
        v2 -= (v2 / ((v2) | 1))
    case 2147483647:
        for fi86 := 0; fi86 < 5; fi86++ {
            var v87 uint64 = 1
            var v88 int32 = -5
            switch (^(fi86)) {
            case -9223372036854775808:
                if false {
                    v89 := GenS10{F0: 1}
                    _ = v89
                } else {
                    v90 := GenS11{F0: 1, F1: ""}
                    switch (fi86 - fi86) {
                    case 0:
                        for fi91 := 0; fi91 < 20; fi91++ {
                            v88 = v88
                            _ = fi91
                        }
                    case 42:
                        for fi92 := 0; fi92 < 11; fi92++ {
                            v88 += (v88 - v88)
                            switch ((("a" + "a") + "\n\t") + "") {
                            case "hello":
                                if len(v4) > 0 {
                                    v4[0] = ("" + (("a" + ("a" + "a")) + "a"))
                                }
                            }
                            _ = len(v1)
                            var v93 float64 = (0.0 / -2.25)
                            _ = fi92
                            _ = v93
                        }
                    }
                    v1[((("\n\t" + ("foo bar" + ((("hello" + "\n\t") + "foo bar") + "a"))) + "foo bar") + "foo bar")] = (fi86 + fi86)
                    _ = v90
                }
            case 32767:
                var v94 int16 = 1000
                _ = v94
            case 42:
                fmt.Println(v87)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi86
            _ = v87
            _ = v88
        }
    default:
        switch 138 {
        case 2147483647:
            fmt.Println(v2)
        }
    }
    if (-2147483648 == 0) {
        v2 -= (^(v2))
    }
    v95 := GenS10{F0: 200}
    var v96 float64 = 1e10
    var v97 uint = 255
    for fi98 := 0; fi98 < 8; fi98++ {
        v96 = (-2.25 - (v96 * (((v96 - v96) / float64(float32(-2.25))) * v96)))
        _ = fi98
    }
    var v99 string = "a"
    for fi100 := 0; fi100 < 8; fi100++ {
        v101 := GenS12{F0: 2147483647}
        _ = fi100
        _ = v101
    }
    v3 = append(v3, 913)
    v96 = ((v96 / 1e10) + v96)
    for fi102 := 0; fi102 < 19; fi102++ {
        v103 := GenS13{F0: 42}
        for fi104 := 0; fi104 < 6; fi104++ {
            if true {
                v2 = 1
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                if (!true) {
                    v99 = (v99 + ((("" + v99) + v99) + "foo bar"))
                    var v105 float64 = -2.25
                    _ = v105
                } else {
                    _ = len(v4)
                    if false {
                        v106 := GenS1{F0: 0, F1: 725, F2: 2}
                        var v107 float32 = 1e10
                        fmt.Println(v107, v96)
                        switch fi104 {
                        case 1:
                            var v108 string = "foo bar"
                            var v109 int64 = 0
                            v99 = v108
                            _ = v108
                            _ = v109
                        case 2147483647:
                            fmt.Println(v99, v107)
                        case 128:
                            _ = len(v1)
                        }
                        v110 := GenS5{F0: 42, F1: 255}
                        _ = v106
                        _ = v107
                        _ = v110
                    } else {
                        var v111 uint64 = 1
                        fmt.Println(v2, fi102, v99)
                        var v112 float32 = (3.14159 + (0.0 / (3.14159 + float32(v96))))
                        _ = v111
                        _ = v112
                    }
                }
            }
            if true {
                delete(v1, "a")
                _ = len(v1)
            }
            _ = fi104
        }
        v113 := GenS6{F0: 9}
        _ = fi102
        _ = v103
        _ = v113
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v95
    _ = v96
    _ = v97
    _ = v99
}
