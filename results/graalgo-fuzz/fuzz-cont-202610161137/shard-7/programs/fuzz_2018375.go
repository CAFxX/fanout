package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16; F1 string; F2 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int; F1 int64 }
type GenS12 struct { F0 string; F1 int8; F2 string }
type GenS13 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS14 struct { F0 int; F1 uint64 }
type GenS15 struct { F0 bool; F1 int16; F2 int16 }

func main() {
    v1 := []string{("\n\t" + "foo bar")}
    var v2 uint64 = 1
    var v3 int16 = -1
    var v4 int = 888
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            if true {
                var v7 int32 = int32(fi5)
                if true {
                    v2 = (v2 >> ((0) & 63))
                    v7 = (^(v7))
                }
                fmt.Println(v3, v2)
                _ = v7
            }
            _ = fi6
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi5
    }
    v2 = 0
    var v8 int = (v4 << ((266) & 63))
    fmt.Println(v3, v3, v4)
    v9 := GenS1{F0: (("foo bar" + "a") + "a")}
    v4 += (v8 | 127)
    var v10 uint = 255
    for fi11 := 0; fi11 < 15; fi11++ {
        if ((fi11 | (fi11 >> ((v10) & 63))) > v4) {
            if true {
                var v12 bool = false
                if v12 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        if ("foo bar" == "") {
                            v1 = append(v1, ("foo bar" + "a"))
                            for fi13 := 0; fi13 < 20; fi13++ {
                                v14 := GenS2{F0: uint32(v2)}
                                var v15 float32 = (0.1 * (0.1 + (3.14159 / 3.14159)))
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            var v16 float64 = float64(1e10)
                            _ = v16
                        }
                    default:
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 string = "\n\t"
                            fmt.Println(v8, fi17)
                            _ = fi17
                            _ = v18
                        }
                    }
                    if v12 {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        var v19 int8 = -128
                        _ = v19
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v20 float32 = (3.14159 + -2.25)
                    _ = v20
                }
                _ = v12
            } else {
                var v21 int32 = 0
                if true {
                    var v22 int32 = 1
                    if true {
                        v4 = int(v2)
                        v2 = (v2 >> (((^(v10))) & 63))
                        for fi23 := 0; fi23 < 3; fi23++ {
                            if ((v3 - -181) >= 1000) {
                                var v24 string = "\n\t"
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v25 := GenS1{F0: ""}
                                v26 := GenS1{F0: (v24 + (((v24 + (v24 + v24)) + "") + ((v24 + (v24 + "foo bar")) + v24)))}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int8 = -128
                                _ = v27
                            }
                            v28 := GenS3{F0: "\n\t", F1: 128}
                            v29 := GenS4{F0: int64(v2), F1: v21}
                            _ = fi23
                            _ = v28
                            _ = v29
                        }
                    } else {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            fmt.Println(fi30)
                            var v31 uint32 = 42
                            if ("foo bar" < "") {
                                if false {
                                    switch 128 {
                                    case 1:
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                    case 32768:
                                        var v32 int8 = 1
                                        _ = v32
                                    case 2147483647:
                                        switch ("foo bar" + ("foo bar" + "foo bar")) {
                                        case "foo bar":
                                            v8 -= int(v10)
                                        default:
                                            if true {
                                                if false {
                                                    fmt.Println(v22, v3)
                                                    switch 127 {
                                                    case -1:
                                                        fmt.Println(fi30, fi11)
                                                    case 127:
                                                        if false {
                                                            var v33 float32 = 1e10
                                                            v2 -= v2
                                                            if (1 <= 100) {
                                                                v34 := GenS3{F0: "hello", F1: 128}
                                                                _ = len(v1)
                                                                var v35 int16 = v3
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            switch ("a" + "") {
                                                            case "foo bar":
                                                                v36 := GenS5{F0: (^(v3)), F1: false}
                                                                _ = v36
                                                            default:
                                                                v10 = v10
                                                            }
                                                            v37 := GenS6{F0: v3, F1: ("hello" + "hello"), F2: 0}
                                                            _ = v33
                                                            _ = v37
                                                        }
                                                    }
                                                    var v38 uint16 = 77
                                                    v10 += v10
                                                    if (v31 <= v31) {
                                                        if len(v1) > 0 {
                                                            v1[0] = ("foo bar" + ("hello" + "foo bar"))
                                                        }
                                                    } else {
                                                        var v39 int32 = v21
                                                        if false {
                                                            fmt.Println(v39, v10, v10)
                                                            var v40 int = fi30
                                                            _ = v40
                                                        }
                                                        if (true && (!true)) {
                                                            if (-101 < 127) {
                                                                v41 := GenS7{F0: (v39 & 290)}
                                                                if true {
                                                                    v3 -= v3
                                                                    v42 := GenS7{F0: (-(v22))}
                                                                    var v43 int8 = 1
                                                                    var v44 uint64 = 42
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    v45 := GenS7{F0: v22}
                                                                    _ = v45
                                                                }
                                                                var v46 uint16 = (v38 + (v38 << ((v10) & 63)))
                                                                _ = v41
                                                                _ = v46
                                                            }
                                                            v8 += -256
                                                            v1 = append(v1, ("hello" + "foo bar"))
                                                        }
                                                        _ = v39
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if len(v1) > 0 {
                                                            v1[0] = ("a" + "\n\t")
                                                        }
                                                    case "":
                                                        fmt.Println(v31, v22, v10)
                                                    default:
                                                        var v47 string = ("foo bar" + "hello")
                                                        _ = v47
                                                    }
                                                    var v48 uint = (v10 | 255)
                                                    _ = v38
                                                    _ = v48
                                                } else {
                                                    v49 := GenS1{F0: "foo bar"}
                                                    v1 = append(v1, "\n\t")
                                                    v50 := GenS7{F0: (v21 & v21)}
                                                    _ = v49
                                                    _ = v50
                                                }
                                                v51 := GenS2{F0: v31}
                                                fmt.Println(fi11, v10)
                                                _ = v51
                                            } else {
                                                if (2147483648 <= 376) {
                                                    v52 := GenS1{F0: "hello"}
                                                    _ = v52
                                                }
                                            }
                                        }
                                    default:
                                        v53 := GenS3{F0: "", F1: 241}
                                        _ = v53
                                    }
                                    var v54 string = ""
                                    _ = v54
                                } else {
                                    var v55 int8 = -1
                                    if false {
                                        if true {
                                            var v56 uint8 = 217
                                            _ = len(v1)
                                            _ = v56
                                        }
                                        fmt.Println(fi30)
                                    }
                                    if (v2 != v2) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v57 uint64 = v2
                                        v58 := GenS6{F0: v3, F1: "foo bar", F2: 734}
                                        _ = v57
                                        _ = v58
                                    } else {
                                        if (true || true) {
                                            v8 -= 4294967295
                                        }
                                        switch -598 {
                                        case -188:
                                            _ = len(v1)
                                        case 128:
                                            v10 -= 4294967295
                                            v59 := GenS8{F0: 0}
                                            _ = v59
                                        default:
                                            if (false || false) {
                                                var v60 int = 32768
                                                if (!(!(int64(v2) > -544))) {
                                                    var v61 uint = 2
                                                    _ = v61
                                                }
                                                v62 := GenS9{F0: 42}
                                                if true {
                                                    v63 := GenS8{F0: int8(v10)}
                                                    var v64 int64 = -9223372036854775808
                                                    v8 += (v8 + (v60 * v4))
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    if (((^(v55)) < (v55 | (v55 - (v55 | v55)))) || (false || true)) {
                                                        v1 = append(v1, ("hello" + ("" + ("foo bar" + "hello"))))
                                                    } else {
                                                        switch v60 {
                                                        case 127:
                                                            v2 += v2
                                                        default:
                                                            v4 = (v60 / ((fi30) | 1))
                                                        }
                                                        if (!(false || (true || true))) {
                                                            switch v8 {
                                                            case -432:
                                                                switch "hello" {
                                                                case "a":
                                                                    v22 = v21
                                                                case "hello":
                                                                    fmt.Println(v3, v3)
                                                                }
                                                            case 42:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 1:
                                                                var v65 uint8 = 42
                                                                _ = v65
                                                            default:
                                                                v4 = v4
                                                            }
                                                            switch fi30 {
                                                            case 48:
                                                                var v66 uint64 = v2
                                                                v3 += (v3 % ((v3) | 1))
                                                                _ = v66
                                                            case -386:
                                                                v10 = (^(v10))
                                                            case 2147483647:
                                                                v67 := GenS7{F0: v22}
                                                                _ = v67
                                                            default:
                                                                if (29 == 27) {
                                                                    v60 -= -9223372036854775808
                                                                    v60 = -1
                                                                    switch (("foo bar" + "\n\t") + "\n\t") {
                                                                    case "foo bar":
                                                                        _ = len(v1)
                                                                        _ = len(v1)
                                                                        var v68 uint = (^(v10))
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v68
                                                                    case "\n\t":
                                                                        var v69 int32 = (v21 ^ 1)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v69
                                                                    }
                                                                } else {
                                                                    var v70 int64 = 420
                                                                    if (v70 > int64(fi11)) {
                                                                        if true {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (("" + "") + "")
                                                                            }
                                                                            if (1 >= 169) {
                                                                                v71 := GenS4{F0: (v70 ^ v70), F1: (^(v21))}
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                var v72 float64 = (-2.25 + 0.0)
                                                                                _ = v71
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                        var v73 int16 = (v3 * -32768)
                                                                        var v74 int8 = 42
                                                                        _ = v73
                                                                        _ = v74
                                                                    } else {
                                                                        if true {
                                                                            v75 := GenS10{F0: uint(v31)}
                                                                            fmt.Println(fi11)
                                                                            var v76 uint = (v10 | (v10 - v10))
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    _ = v70
                                                                }
                                                                v77 := GenS11{F0: (v8 | fi30), F1: -364}
                                                                _ = v77
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v60
                                                _ = v62
                                            }
                                        }
                                        switch ("foo bar" + "") {
                                        case "":
                                            switch v8 {
                                            case 128:
                                                fmt.Println(v21, v3)
                                            default:
                                                fmt.Println(v22)
                                            }
                                        }
                                        var v78 uint = (v10 << ((255) & 63))
                                        v55 = v55
                                        _ = v78
                                    }
                                    _ = v55
                                }
                                fmt.Println(v31, v31, v8)
                            } else {
                                var v79 uint = (v10 + v10)
                                _ = v79
                            }
                            _ = fi30
                            _ = v31
                        }
                        var v80 uint64 = v2
                        v8 -= 0
                        _ = v80
                    }
                    _ = v22
                } else {
                    if (false && true) {
                        v81 := GenS10{F0: (^(v10))}
                        _ = v81
                    } else {
                        v82 := GenS8{F0: int8(v2)}
                        var v83 int8 = -110
                        for fi84 := 0; fi84 < 3; fi84++ {
                            var v85 int16 = v3
                            _ = fi84
                            _ = v85
                        }
                        _ = v82
                        _ = v83
                    }
                }
                if (!false) {
                    for fi86 := 0; fi86 < 19; fi86++ {
                        if (2147483648 >= 65535) {
                            v3 -= 1000
                        } else {
                            var v87 string = "foo bar"
                            var v88 float64 = 1.5
                            _ = v87
                            _ = v88
                        }
                        var v89 uint8 = 255
                        _ = fi86
                        _ = v89
                    }
                    if false {
                        var v90 int = fi11
                        v10 -= v10
                        if (((v4 + v4) == -180) && (!((!((false && ((v10 * (v10 | (v10 << (((v10 & v10)) & 63)))) == v10)) && (!((v2 - (v2 ^ uint64(v10))) > v2)))) || (v8 < v4)))) {
                            fmt.Println(v10, v10, v3)
                            if false {
                                switch (v4 ^ (v4 << ((v10) & 63))) {
                                case -9223372036854775808:
                                    v90 += 378
                                default:
                                    v90 = (v8 - (v8 % ((9223372036854775807) | 1)))
                                }
                            } else {
                                for fi91 := 0; fi91 < 5; fi91++ {
                                    var v92 float32 = 0.1
                                    _ = len(v1)
                                    if (!true) {
                                        fmt.Println(v90, v21)
                                        v93 := GenS2{F0: 4294967295}
                                        v94 := GenS12{F0: "hello", F1: -1, F2: ("a" + "foo bar")}
                                        _ = v93
                                        _ = v94
                                    }
                                    _ = fi91
                                    _ = v92
                                }
                            }
                            if false {
                                var v95 int16 = 0
                                if ((-(v21)) >= (v21 / ((2147483647) | 1))) {
                                    for fi96 := 0; fi96 < 11; fi96++ {
                                        var v97 uint = (^(v10))
                                        v10 = (^(v97))
                                        var v98 int8 = 0
                                        _ = fi96
                                        _ = v97
                                        _ = v98
                                    }
                                }
                                fmt.Println(v8, v21, v95)
                                _ = v95
                            } else {
                                v2 -= v2
                            }
                        }
                        _ = v90
                    }
                    switch (-(fi11)) {
                    case 32768:
                        for fi99 := 0; fi99 < 20; fi99++ {
                            var v100 bool = (v10 == (v10 + v10))
                            _ = fi99
                            _ = v100
                        }
                    case -2147483648:
                        v101 := GenS12{F0: "foo bar", F1: 42, F2: ("hello" + (("a" + "") + "hello"))}
                        _ = v101
                    default:
                        v102 := GenS2{F0: 4294967295}
                        _ = v102
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v8, v3, v2)
                    var v103 uint8 = 128
                    _ = v103
                }
                _ = v21
            }
        } else {
            fmt.Println(v3, v8, v10)
        }
        _ = fi11
    }
    var v104 uint8 = 24
    if (false && true) {
        v105 := GenS5{F0: 1, F1: true}
        _ = v105
    } else {
        var v106 uint8 = v104
        if false {
            if (false || (!(false && (1 < (v4 % (((v8 << (((v10 * (v10 | v10))) & 63))) | 1)))))) {
                v107 := GenS11{F0: -291, F1: -9223372036854775808}
                v1 = append(v1, "")
                _ = v107
            }
            fmt.Println(v106)
            v2 = v2
        }
        v108 := GenS13{F0: 865, F1: 42, F2: 343}
        _ = v106
        _ = v108
    }
    for fi109 := 0; fi109 < 2; fi109++ {
        var v110 uint16 = 907
        if (("a" + "") < "a") {
            if ((0.0 / (float32(-2.25) - 1.5)) <= -2.25) {
                for fi111 := 0; fi111 < 20; fi111++ {
                    v8 += (v4 | (v8 % ((-9223372036854775808) | 1)))
                    _ = fi111
                }
                switch "hello" {
                case "\n\t":
                    switch "" {
                    case "a":
                        v112 := GenS9{F0: v10}
                        _ = v112
                    case "foo bar":
                        switch int(v110) {
                        case -752:
                            if false {
                                for fi113 := 0; fi113 < 6; fi113++ {
                                    var v114 bool = (62 < v104)
                                    v115 := GenS11{F0: (v4 * (^(v4))), F1: 982}
                                    v116 := GenS14{F0: v8, F1: 2}
                                    _ = fi113
                                    _ = v114
                                    _ = v115
                                    _ = v116
                                }
                                for fi117 := 0; fi117 < 17; fi117++ {
                                    v118 := GenS12{F0: "a", F1: -100, F2: (("\n\t" + "\n\t") + ("a" + "foo bar"))}
                                    _ = fi117
                                    _ = v118
                                }
                                var v119 int = v4
                                _ = v119
                            } else {
                                _ = len(v1)
                            }
                        }
                    case "hello":
                        switch (v4 ^ v4) {
                        case 736:
                            v120 := GenS5{F0: v3, F1: true}
                            _ = v120
                        default:
                            _ = len(v1)
                        }
                    }
                case "foo bar":
                    for fi121 := 0; fi121 < 4; fi121++ {
                        fmt.Println(v4, v110, v8)
                        _ = fi121
                    }
                }
            } else {
                fmt.Println(v4)
                var v122 int16 = (v3 >> ((v10) & 63))
                _ = len(v1)
                _ = v122
            }
            var v123 int = -1
            fmt.Println(v104, v8)
            _ = v123
        }
        _ = fi109
        _ = v110
    }
    if true {
        v124 := GenS9{F0: (v10 | 0)}
        v4 = 32768
        switch v8 {
        case 9223372036854775807:
            v10 -= v10
        case 1:
            if false {
                var v125 uint = v10
                _ = v125
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 32768:
            v2 = v2
        }
        _ = v124
    }
    var v126 int32 = 0
    v127 := GenS15{F0: true, F1: v3, F2: v3}
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v10 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v104:", v104)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v104
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
