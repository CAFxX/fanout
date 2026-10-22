package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 string; F1 int8; F2 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    v3 := make(map[string]uint32)
    v4 := []string{"hello"}
    var v5 string = ((("hello" + ("foo bar" + "\n\t")) + "foo bar") + "hello")
    var v6 uint8 = 165
    delete(v3, "")
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int32 = 42
        v8 += (v8 ^ (v8 | v8))
        _ = fi7
        _ = v8
    }
    switch ("" + "foo bar") {
    case "":
        var v9 int64 = 9223372036854775807
        _ = v9
    case "hello":
        var v10 int = 4294967295
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 bool = (true || (false && (v5 == "foo bar")))
            _ = fi11
            _ = v12
        }
    default:
        var v13 float32 = 1.5
        _ = v13
    }
    if ((v5 + v5) == v5) {
        v5 = "\n\t"
        if false {
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: int32(v6), F1: (0 >= (^(v2))), F2: v5}
                _ = fi14
                _ = v15
            }
            fmt.Println(v2, v6)
            var v16 float32 = -2.25
            _ = v16
        } else {
            if false {
                v17 := GenS1{F0: -140, F1: (42 > 0), F2: v5}
                delete(v3, v5)
                _ = v17
            } else {
                var v18 uint64 = 18446744073709551615
                if ((^(v18)) > v18) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                if len(v4) > 0 {
                    v4[0] = v5
                }
                v6 += v6
                _ = fi19
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = len(v3)
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS1{F0: 2147483647, F1: false, F2: v5}
        v5 = ("hello" + v5)
        var v22 float32 = 1e10
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        v6 = (v6 / ((1) | 1))
        for fi23 := 0; fi23 < 4; fi23++ {
            _ = len(v1)
            _ = fi23
        }
        fmt.Println(v5, v2, v6)
    }
    if false {
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                if true {
                    var v26 int8 = -128
                    v1[""] = 196
                    if false {
                        if false {
                            if ((float64(0.0) - (float64(0.0) + ((0.0 * 1.5) + 1.5))) >= -2.25) {
                                fmt.Println(fi24)
                                var v27 int32 = 633
                                v4 = append(v4, ((v5 + (v5 + "foo bar")) + (("" + (("" + v5) + "")) + "\n\t")))
                                v28 := GenS1{F0: 1, F1: false, F2: "a"}
                                _ = v27
                                _ = v28
                            }
                            v4 = append(v4, "foo bar")
                        }
                        v29 := GenS1{F0: int32(fi25), F1: (0.1 < (float64(float32((1.5 - ((0.0 / -2.25) + (3.14159 - (3.14159 * 0.1)))))) * -2.25)), F2: v5}
                        _ = v29
                    }
                    v5 = ((v5 + "foo bar") + v5)
                    _ = v26
                }
                v6 = v6
                switch v5 {
                case "a":
                    fmt.Println(v6, v2)
                default:
                    v6 += (v6 ^ 59)
                }
                var v30 int64 = 1
                _ = fi25
                _ = v30
            }
            if len(v4) > 0 {
                v4[0] = v5
            }
            _ = fi24
        }
        v6 = 128
    } else {
        fmt.Println(v6)
        if true {
            for fi31 := 0; fi31 < 14; fi31++ {
                var v32 uint = 726
                v33 := v3[v5]
                v34 := GenS2{F0: (false || ((fi31 ^ 4294967295) < 9223372036854775807))}
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            switch 1 {
            case -1:
                switch 127 {
                case 4294967295:
                    v35 := GenS1{F0: -1, F1: true, F2: ("a" + (v5 + ("" + "foo bar")))}
                    _ = v35
                case 32767:
                    switch v5 {
                    case "hello":
                        fmt.Println(v5, v6)
                    case "\n\t":
                        v5 = ((v5 + v5) + v5)
                    default:
                        v6 = (v6 >> ((4294967295) & 63))
                    }
                case 308:
                    fmt.Println(v6)
                }
            case -512:
                if false {
                    v5 = ("foo bar" + v5)
                    v36 := v1["hello"]
                    _ = v36
                }
            case 1:
                for fi37 := 0; fi37 < 7; fi37++ {
                    var v38 float64 = 3.14159
                    _ = fi37
                    _ = v38
                }
            default:
                var v39 uint64 = 773
                _ = v39
            }
        }
    }
    var v40 int64 = 9223372036854775807
    _ = len(v1)
    for fi41 := 0; fi41 < 3; fi41++ {
        v5 = v5
        v2 = (^(v2))
        if (-128 > -128) {
            switch fi41 {
            case -185:
                for fi42 := 0; fi42 < 18; fi42++ {
                    var v43 float32 = -2.25
                    v43 += v43
                    switch fi41 {
                    case 42:
                        v2 += (v2 >> ((2) & 63))
                        v3[v5] = 2147483647
                        v44 := GenS3{F0: 2}
                        _ = v44
                    case 0:
                        var v45 int8 = -128
                        _ = v45
                    default:
                        switch "\n\t" {
                        case "":
                            _ = len(v4)
                        default:
                            switch ("" + v5) {
                            case "\n\t":
                                switch ("\n\t" + v5) {
                                case "":
                                    var v46 string = "hello"
                                    _ = v46
                                case "foo bar":
                                    if true {
                                        v5 = (v5 + (v5 + v5))
                                        if false {
                                            v43 = 0.1
                                            v47 := GenS3{F0: 2}
                                            if (0 > 18446744073709551615) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v48 int64 = (v40 % (((v40 - v40)) | 1))
                                                    _ = v48
                                                case "foo bar":
                                                    switch (v5 + (("\n\t" + (v5 + v5)) + v5)) {
                                                    case "":
                                                        v6 = (v6 | (v6 >> ((18446744073709551615) & 63)))
                                                    case "\n\t":
                                                        switch v5 {
                                                        case "a":
                                                            fmt.Println(v5, fi41)
                                                        case "":
                                                            var v49 float32 = v43
                                                            var v50 int64 = v40
                                                            var v51 int32 = -1
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        default:
                                                            v52 := GenS4{F0: -731, F1: -49}
                                                            _ = v52
                                                        }
                                                    }
                                                case "":
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v53 uint = 149
                                                        _ = v53
                                                    }
                                                }
                                                if false {
                                                    v54 := GenS3{F0: 2147483648}
                                                    if ((true && (false || true)) && false) {
                                                        delete(v3, v5)
                                                        if (float32(1e10) < (v43 + 0.1)) {
                                                            var v55 float32 = v43
                                                            var v56 int8 = 4
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            v43 = v43
                                                        }
                                                        v2 = (v2 % ((v2) | 1))
                                                    } else {
                                                        var v57 int8 = 18
                                                        var v58 bool = false
                                                        var v59 int32 = -784
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    _ = v54
                                                } else {
                                                    var v60 uint32 = 774
                                                    var v61 float32 = (0.0 + 1.5)
                                                    _ = v60
                                                    _ = v61
                                                }
                                                v2 -= (v2 & (v2 & (^(v2))))
                                                if false {
                                                    fmt.Println(v2, v2, v5)
                                                    var v62 bool = (42 != 42)
                                                    _ = v62
                                                } else {
                                                    v5 = v5
                                                    v63 := GenS4{F0: fi41, F1: 56}
                                                    v5 = v5
                                                    _ = v63
                                                }
                                                v64 := v1["hello"]
                                                var v65 float64 = 1e10
                                                _ = v64
                                                _ = v65
                                            } else {
                                                if false {
                                                    var v66 float32 = v43
                                                    v40 -= 0
                                                    _ = len(v4)
                                                    _ = v66
                                                }
                                                v5 = ""
                                            }
                                            _ = v47
                                        }
                                    } else {
                                        var v67 int8 = 1
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v68 int8 = 1
                                            _ = v68
                                        case "a":
                                            switch (^(fi42)) {
                                            case -1:
                                                var v69 uint32 = 2147483647
                                                _ = v69
                                            case 68:
                                                _ = len(v1)
                                                var v70 int64 = (v40 >> ((247) & 63))
                                                var v71 uint16 = v2
                                                _ = v70
                                                _ = v71
                                            }
                                        case "\n\t":
                                            fmt.Println(fi41, v40, v2)
                                        default:
                                            v40 = (v40 | v40)
                                        }
                                        var v72 uint = 65535
                                        v73 := GenS3{F0: 0}
                                        _ = v67
                                        _ = v72
                                        _ = v73
                                    }
                                    if true {
                                        var v74 uint32 = 4294967295
                                        var v75 int32 = 42
                                        var v76 uint32 = 4294967295
                                        _ = v74
                                        _ = v75
                                        _ = v76
                                    } else {
                                        var v77 bool = (!(1e10 <= 0.0))
                                        v4 = append(v4, v5)
                                        _ = v77
                                    }
                                    if ((v40 <= (v40 + (v40 - (^(v40))))) && (525 > 42)) {
                                        v40 = v40
                                    }
                                }
                            case "a":
                                var v78 float32 = (v43 * ((1e10 * v43) - v43))
                                _ = v78
                            default:
                                switch (fi41 % (((fi41 ^ -2147483648)) | 1)) {
                                case -9223372036854775808:
                                    fmt.Println(v43, fi42, v40)
                                }
                            }
                        }
                    }
                    _ = fi42
                    _ = v43
                }
            case 42:
                v2 -= (v2 % (((v2 | 69)) | 1))
            case 1:
                v6 = v6
            default:
                fmt.Println(v2, v6)
            }
            v6 += (v6 / (((v6 ^ (v6 >> ((2) & 63)))) | 1))
        } else {
            v79 := GenS1{F0: -390, F1: true, F2: "foo bar"}
            _ = v79
        }
        _ = fi41
    }
    var v80 string = (("hello" + "\n\t") + ((((v5 + v5) + ("" + "hello")) + v5) + v5))
    if (!false) {
        switch "\n\t" {
        case "\n\t":
            if true {
                for fi81 := 0; fi81 < 19; fi81++ {
                    v82 := GenS5{F0: (v2 / (((v2 * 32767)) | 1)), F1: 1}
                    var v83 int32 = -392
                    _ = fi81
                    _ = v82
                    _ = v83
                }
            }
        case "hello":
            v40 = 42
        case "":
            if (((0.0 >= float64((((1e10 + 3.14159) / 3.14159) * 0.1))) || ((903 > 418) && (!false))) && false) {
                var v84 uint8 = v6
                _ = v84
            }
        default:
            for fi85 := 0; fi85 < 3; fi85++ {
                v86 := GenS6{F0: "\n\t", F1: 42, F2: (-(fi85))}
                _ = len(v1)
                _ = fi85
                _ = v86
            }
        }
        if (!false) {
            v80 = v80
            if true {
                v2 = v2
                for fi87 := 0; fi87 < 5; fi87++ {
                    var v88 float32 = 0.1
                    fmt.Println(v6, v5, v2)
                    if true {
                        v89 := GenS7{F0: 255}
                        fmt.Println(v2, v40)
                        for fi90 := 0; fi90 < 2; fi90++ {
                            var v91 int16 = 0
                            v40 = 1
                            _ = fi90
                            _ = v91
                        }
                        _ = v89
                    }
                    _ = fi87
                    _ = v88
                }
            }
            v2 = v2
        } else {
            if false {
                v5 = v80
            }
            if (!(0 != -32768)) {
                for fi92 := 0; fi92 < 13; fi92++ {
                    v93 := GenS8{F0: 117, F1: -925}
                    var v94 float64 = -2.25
                    var v95 uint16 = v2
                    _ = fi92
                    _ = v93
                    _ = v94
                    _ = v95
                }
            } else {
                fmt.Println(v40, v6)
                for fi96 := 0; fi96 < 2; fi96++ {
                    for fi97 := 0; fi97 < 1; fi97++ {
                        if ((v40 / ((1) | 1)) != v40) {
                            fmt.Println(v5, v5, v2)
                            v98 := GenS3{F0: 391}
                            if true {
                                if true {
                                    var v99 float32 = (float32(3.14159) * 0.1)
                                    v5 = v5
                                    v6 = 2
                                    _ = v99
                                }
                                var v100 int32 = -624
                                _ = v100
                            } else {
                                switch (fi97 ^ -1) {
                                case 42:
                                    fmt.Println(v40, fi96)
                                case -1:
                                    v5 = (v5 + "foo bar")
                                case 924:
                                    v101 := GenS8{F0: -1, F1: 0}
                                    _ = v101
                                }
                            }
                            var v102 int = fi96
                            _ = v102
                            _ = v98
                        }
                        if ((0.0 + float32(0.1)) == 0.1) {
                            var v103 uint32 = 0
                            var v104 string = v5
                            var v105 int64 = (v40 - v40)
                            _ = v103
                            _ = v104
                            _ = v105
                        }
                        _ = fi97
                    }
                    v106 := GenS5{F0: (v2 + (v2 << ((956) & 63))), F1: 127}
                    _ = fi96
                    _ = v106
                }
            }
            switch v80 {
            case "":
                if true {
                    switch 42 {
                    case -2147483648:
                        _ = len(v3)
                    case 1:
                        v107 := GenS4{F0: -2147483648, F1: 1}
                        _ = v107
                    case 957:
                        var v108 uint16 = v2
                        _ = v108
                    }
                    if (42 == 42) {
                        switch 9223372036854775807 {
                        case 42:
                            if len(v4) > 0 {
                                v4[0] = v5
                            }
                        case 934:
                            v109 := v1["foo bar"]
                            _ = v109
                        case 32768:
                            var v110 uint32 = 2
                            _ = v110
                        }
                    }
                }
            }
        }
    }
    v5 = "\n\t"
    v6 = (^(v6))
    if ((!true) || (1 > 597)) {
        switch (v80 + v80) {
        case "a":
            v6 = (v6 + v6)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v80
}
