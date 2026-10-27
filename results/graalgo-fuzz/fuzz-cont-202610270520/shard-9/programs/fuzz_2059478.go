package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS9 struct { F0 int16; F1 int8; F2 int8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 int16; F2 uint32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint8; F1 int64; F2 int32 }

func main() {
    var v1 float64 = float64((3.14159 + 3.14159))
    var v2 int64 = 0
    v3 := make(map[string]int)
    var v4 bool = (255 >= 127)
    v5 := GenS1{F0: 127}
    var v6 float64 = (0.0 * -2.25)
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            if false {
                if ((9223372036854775807 >= (fi7 & fi7)) || (0 != 999)) {
                    fmt.Println(v6, v4)
                    if (4294967295 < 42) {
                        v2 -= 1
                        var v8 bool = false
                        fmt.Println(v4, v8)
                        _ = v8
                    }
                } else {
                    _ = len(v3)
                    for fi9 := 0; fi9 < 6; fi9++ {
                        var v10 float64 = 1e10
                        switch 1 {
                        case -938:
                            fmt.Println(v2)
                        case 32768:
                            if true {
                                v10 = (v10 * (float64(3.14159) * 1e10))
                                var v11 float32 = float32(v6)
                                switch ((("hello" + "foo bar") + "\n\t") + ("" + "a")) {
                                case "":
                                    var v12 uint32 = 484
                                    _ = v12
                                default:
                                    v13 := GenS2{F0: 482}
                                    v14 := []int{(fi9 % ((fi7) | 1))}
                                    _ = v13
                                    _ = v14
                                }
                                var v15 uint64 = 18446744073709551615
                                v16 := GenS3{F0: 2147483648, F1: (fi9 == (-(fi9)))}
                                if v4 {
                                    v17 := GenS4{F0: -1}
                                    v6 = (float64(v11) + 1e10)
                                    _ = v17
                                } else {
                                    v18 := []int{fi9}
                                    switch "\n\t" {
                                    case "\n\t":
                                        if false {
                                            if (-63 < 42) {
                                                v19 := GenS5{F0: v4}
                                                var v20 int16 = 398
                                                _ = v19
                                                _ = v20
                                            }
                                            v2 = v2
                                            v18 = append(v18, 759)
                                            if false {
                                                var v21 uint = 2
                                                _ = v21
                                            } else {
                                                switch (fi7 | (^(fi7))) {
                                                case 32768:
                                                    var v22 int32 = -567
                                                    _ = v22
                                                case -652:
                                                    var v23 int8 = int8(fi7)
                                                    _ = v23
                                                }
                                            }
                                            var v24 int32 = -2147483648
                                            _ = v24
                                        } else {
                                            switch (("hello" + (("foo bar" + "hello") + "")) + ("" + "foo bar")) {
                                            case "":
                                                if v4 {
                                                    v15 -= v15
                                                } else {
                                                    switch ("hello" + "hello") {
                                                    case "foo bar":
                                                        v25 := GenS6{F0: v15}
                                                        _ = v25
                                                    case "\n\t":
                                                        if ((fi9 / ((-781) | 1)) >= (fi9 << ((396) & 63))) {
                                                            if true {
                                                                var v26 string = "foo bar"
                                                                _ = v26
                                                            }
                                                        }
                                                    }
                                                    if ((^(v15)) == (v15 + 290)) {
                                                        var v27 int32 = -1
                                                        _ = v27
                                                    }
                                                    var v28 string = "foo bar"
                                                    _ = v28
                                                }
                                            default:
                                                switch "a" {
                                                case "hello":
                                                    var v29 bool = true
                                                    _ = v29
                                                case "":
                                                    if (((((0.1 + v1) * v6) / float64(v11)) + float64(-2.25)) > (((float64(0.1) + (3.14159 * 0.1)) / 1.0) - ((v6 / v10) - (1e10 * float64(((1e10 + v11) * 1e10)))))) {
                                                        var v30 uint = 2
                                                        var v31 uint32 = 167
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    if (false && (101 != uint8(v15))) {
                                                        var v32 uint8 = 9
                                                        var v33 uint32 = 137
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                }
                                            }
                                            switch ((("hello" + "") + "hello") + "") {
                                            case "a":
                                                fmt.Println(v2, v10)
                                            case "":
                                                _ = len(v18)
                                            default:
                                                v34 := v3["foo bar"]
                                                _ = v34
                                            }
                                            v1 -= ((v6 + (((v6 / (v1 * (((1e10 - (v1 + float64(float32(3.14159)))) * v6) / 0.1))) + ((((((1.5 + ((v10 * float64(v11)) + v1)) + v6) - v10) * (3.14159 - -2.25)) + (v1 - (v1 + 0.1))) + v1)) / 3.14159)) + ((v6 + -2.25) * v1))
                                            if ("\n\t" <= "foo bar") {
                                                var v35 int16 = 42
                                                _ = v35
                                            } else {
                                                var v36 uint32 = 439
                                                _ = v36
                                            }
                                        }
                                    case "a":
                                        if false {
                                            switch "" {
                                            case "foo bar":
                                                _ = len(v18)
                                            default:
                                                var v37 string = "foo bar"
                                                _ = v37
                                            }
                                            v38 := GenS2{F0: 679}
                                            _ = v38
                                        } else {
                                            if (1 == 42) {
                                                v39 := GenS7{F0: true}
                                                _ = v39
                                            } else {
                                                v10 = -2.25
                                            }
                                            switch (fi9 * fi7) {
                                            case -2147483648:
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                            case 1:
                                                v18 = append(v18, (fi7 * (fi9 + fi7)))
                                            }
                                        }
                                    default:
                                        if ("a" >= "a") {
                                            if v4 {
                                                if (-631 < 755) {
                                                    switch "" {
                                                    case "\n\t":
                                                        v6 += (float64(v11) / 1.5)
                                                    case "foo bar":
                                                        if v4 {
                                                            v40 := GenS7{F0: v4}
                                                            v6 += v1
                                                            _ = v40
                                                        }
                                                    case "":
                                                        var v41 uint8 = 78
                                                        _ = v41
                                                    }
                                                } else {
                                                    if (((((2 >= 12) && v4) && true) && (!(42 >= 32767))) && v4) {
                                                        v42 := GenS2{F0: 42}
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v43 uint32 = 78
                                                            _ = v43
                                                        case "hello":
                                                            var v44 uint = 0
                                                            _ = v44
                                                        case "\n\t":
                                                            fmt.Println(v2, v11, v10)
                                                        }
                                                        _ = v42
                                                    } else {
                                                        fmt.Println(v6, v11)
                                                        if false {
                                                            fmt.Println(v1)
                                                            fmt.Println(fi7, fi9)
                                                            if (v11 < 1.5) {
                                                                _ = len(v18)
                                                                var v45 uint8 = 128
                                                                _ = v45
                                                            }
                                                        }
                                                    }
                                                }
                                                v10 = v6
                                            }
                                            v18 = append(v18, fi9)
                                        }
                                    }
                                    var v46 int32 = -1
                                    _ = v18
                                    _ = v46
                                }
                                v47 := GenS4{F0: (v2 ^ v2)}
                                switch "" {
                                case "":
                                    switch (("hello" + "") + "hello") {
                                    case "a":
                                        v4 = v4
                                    case "hello":
                                        if v4 {
                                            fmt.Println(fi7, v2, v15)
                                            if v4 {
                                                if v4 {
                                                    var v48 int = fi9
                                                    fmt.Println(fi7, v4, v15)
                                                    _ = v48
                                                } else {
                                                    switch "a" {
                                                    case "hello":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v6 -= (((3.14159 + (v6 + ((v10 + v1) * v1))) * (v10 - v10)) / (1.5 * v10))
                                                        case "hello":
                                                            v49 := GenS3{F0: 4294967295, F1: v4}
                                                            _ = v49
                                                        default:
                                                            delete(v3, "a")
                                                        }
                                                    }
                                                    v15 += (^(v15))
                                                    var v50 uint = 42
                                                    _ = v50
                                                }
                                                v51 := []int{}
                                                _ = v51
                                            } else {
                                                delete(v3, "a")
                                            }
                                        } else {
                                            if (1 < 676) {
                                                _ = len(v3)
                                                v4 = false
                                                if false {
                                                    var v52 int32 = 1
                                                    _ = v52
                                                } else {
                                                    var v53 bool = v4
                                                    v2 -= v2
                                                    switch (fi7 ^ fi9) {
                                                    case 2147483647:
                                                        var v54 bool = ((22 <= 255) || (((^(v2)) != v2) || v4))
                                                        _ = v54
                                                    case -1:
                                                        fmt.Println(fi7, v1, v4)
                                                    case -649:
                                                        var v55 bool = (!false)
                                                        _ = v55
                                                    default:
                                                        var v56 bool = true
                                                        _ = v56
                                                    }
                                                    _ = v53
                                                }
                                                if ((!((2 <= 241) || v4)) && ((v15 << ((255) & 63)) <= (v15 ^ v15))) {
                                                    v10 = float64((3.14159 / (((v11 * (float32((3.14159 - (-2.25 + 1.5))) - v11)) + v11) - v11)))
                                                    switch "" {
                                                    case "a":
                                                        _ = len(v3)
                                                    case "hello":
                                                        v11 += v11
                                                    case "\n\t":
                                                        fmt.Println(v10, v6, v11)
                                                    default:
                                                        v57 := GenS4{F0: v2}
                                                        _ = v57
                                                    }
                                                } else {
                                                    v2 = (v2 << ((1) & 63))
                                                    var v58 uint16 = 0
                                                    switch ("hello" + "") {
                                                    case "a":
                                                        v59 := GenS8{F0: 4294967295, F1: 0, F2: (^(v15))}
                                                        _ = v59
                                                    case "foo bar":
                                                        v60 := GenS9{F0: 1, F1: 1, F2: int8(v15)}
                                                        _ = v60
                                                    case "\n\t":
                                                        v11 -= (v11 - v11)
                                                    }
                                                    var v61 int8 = 92
                                                    var v62 bool = v4
                                                    fmt.Println(v61)
                                                    _ = v58
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 uint64 = 42
                                                _ = v63
                                            } else {
                                                var v64 int32 = -864
                                                fmt.Println(v2, v15, v2)
                                                _ = v64
                                            }
                                            if (32767 != 42) {
                                                fmt.Println(fi9, v4)
                                                v65 := []int{32767, fi7, fi7}
                                                _ = len(v3)
                                                _ = v65
                                            } else {
                                                v6 = v6
                                            }
                                        }
                                    default:
                                        v66 := []int{32768, fi9, 127, 4294967295}
                                        _ = v66
                                    }
                                case "hello":
                                    v1 = 0.1
                                }
                                _ = v11
                                _ = v15
                                _ = v16
                                _ = v47
                            } else {
                                var v67 float32 = (1.5 - 0.1)
                                var v68 int64 = (v2 + v2)
                                v4 = true
                                _ = v67
                                _ = v68
                            }
                        case -19:
                            v10 -= ((((1.5 / (0.1 - float64(((((0.0 - -2.25) - (3.14159 + 0.0)) / (1e10 / 1e10)) / 1.0)))) - float64(0.0)) + -2.25) + v10)
                        default:
                            if v4 {
                                if (127 > -2147483648) {
                                    var v69 int32 = 42
                                    v70 := GenS6{F0: 440}
                                    v69 -= v69
                                    _ = v69
                                    _ = v70
                                } else {
                                    v6 = (v1 * (0.1 + (0.0 / (v6 * float64(float32(float64(0.1)))))))
                                    _ = len(v3)
                                }
                                if v4 {
                                    fmt.Println(v1, v2, v6)
                                    if false {
                                        if false {
                                            switch "\n\t" {
                                            case "":
                                                var v71 uint64 = 18446744073709551615
                                                _ = v71
                                            case "foo bar":
                                                v72 := GenS1{F0: 30}
                                                _ = v72
                                            case "hello":
                                                var v73 int32 = 42
                                                _ = v73
                                            }
                                            v1 = 1.5
                                            v74 := []int{(fi7 << ((4294967295) & 63))}
                                            _ = v74
                                        } else {
                                            v10 = v6
                                            v75 := []int{(fi9 | 4294967295), fi7, (fi7 >> ((1) & 63)), int(v2)}
                                            _ = v75
                                        }
                                    }
                                    v2 += (v2 + v2)
                                } else {
                                    v76 := []int{4294967295, (fi9 & (fi9 | fi9)), fi7, (fi7 ^ -894)}
                                    var v77 uint8 = 255
                                    _ = v76
                                    _ = v77
                                }
                                if v4 {
                                    if (v10 < v6) {
                                        fmt.Println(v4, fi9, v2)
                                        delete(v3, "")
                                    }
                                    if v4 {
                                        var v78 int = (-(fi7))
                                        v79 := GenS10{F0: 0}
                                        _ = v78
                                        _ = v79
                                    }
                                }
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                    v80 := []int{4294967295, 629, (fi7 << ((161) & 63))}
                    _ = v80
                }
                var v81 float32 = float32((v6 + v1))
                _ = v81
            } else {
                switch (fi7 + (-(fi7))) {
                case -206:
                    var v82 uint16 = 1
                    _ = v82
                case 9223372036854775807:
                    v83 := GenS11{F0: 2, F1: -32768, F2: 4294967295}
                    _ = v83
                case -1:
                    v84 := GenS6{F0: 0}
                    _ = v84
                }
            }
            var v85 uint64 = 286
            _ = v85
        } else {
            v86 := GenS12{F0: v2}
            _ = v86
        }
        v1 = v1
        v6 = (0.0 / v1)
        _ = fi7
    }
    switch 1 {
    case 1:
        delete(v3, "")
    case 9223372036854775807:
        if (!v4) {
            var v87 int16 = -246
            v3["foo bar"] = -9223372036854775808
            var v88 int64 = v2
            _ = v87
            _ = v88
        } else {
            if v4 {
                var v89 bool = (!true)
                _ = v89
            } else {
                switch ("" + "hello") {
                case "\n\t":
                    v3["\n\t"] = -9223372036854775808
                }
                switch 42 {
                case 4294967295:
                    v4 = v4
                default:
                    var v90 uint64 = 1
                    _ = v90
                }
                for fi91 := 0; fi91 < 17; fi91++ {
                    for fi92 := 0; fi92 < 12; fi92++ {
                        _ = len(v3)
                        _ = fi92
                    }
                    v93 := []int{2147483647}
                    _ = fi91
                    _ = v93
                }
            }
        }
    case -707:
        if (!(!v4)) {
            var v94 int16 = -752
            _ = v94
        }
    }
    for fi95 := 0; fi95 < 8; fi95++ {
        switch -9223372036854775808 {
        case 128:
            var v96 string = ("" + "foo bar")
            _ = v96
        case 768:
            v97 := GenS4{F0: 9223372036854775807}
            _ = v97
        default:
            var v98 uint16 = 598
            _ = v98
        }
        v99 := GenS13{F0: 50, F1: v2, F2: 669}
        if v4 {
            for fi100 := 0; fi100 < 16; fi100++ {
                v2 += (-(v2))
                v101 := GenS3{F0: 1, F1: (true || (1 >= 127))}
                v4 = (1.5 >= v6)
                _ = fi100
                _ = v101
            }
            var v102 int = -1
            _ = v102
        }
        _ = fi95
        _ = v99
    }
    var v103 bool = (2147483648 < 849)
    var v104 int32 = -1
    v2 += v2
    v4 = v103
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
