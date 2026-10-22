package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 string }
type GenS4 struct { F0 string; F1 int; F2 int16 }
type GenS5 struct { F0 int; F1 uint32; F2 uint }
func genFn0(p0 int, p1 int) int8 {
    var v6 int64 = 299
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 bool = true
        v8 = (2147483648 > 2147483647)
        if (float64(1.5) < 0.1) {
            if (("a" < "a") || (3.14159 != 0.1)) {
                var v9 int16 = 32767
                _ = v9
            }
            v10 := GenS1{F0: 483}
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if (-261 > -431) {
        v6 = v6
        switch "hello" {
        case "\n\t":
            v11 := GenS1{F0: 1}
            _ = v11
        default:
            v12 := GenS1{F0: 707}
            _ = v12
        }
        fmt.Println(p0, v6, p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 65535
    var v3 int16 = 21
    var v4 int16 = v3
    var v5 float64 = (((0.0 / (0.0 - 1.5)) * ((3.14159 - 0.1) + -2.25)) * float64(3.14159))
    v13 := GenS1{F0: 18446744073709551615}
    v14 := GenS2{F0: 217, F1: (v2 | (v2 + v2)), F2: (1 < 2)}
    v15 := GenS1{F0: 427}
    v16 := GenS2{F0: 255, F1: (v2 & v2), F2: false}
    v17 := []int{598}
    v4 = v3
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            v1[""] = "hello"
            fmt.Println(fi19)
            _ = fi19
        }
        switch (("" + "") + ((("a" + "foo bar") + ("\n\t" + "\n\t")) + "")) {
        case "":
            var v20 uint64 = 1
            _ = v20
        default:
            if len(v17) > 0 {
                v17[0] = fi18
            }
        }
        _ = fi18
    }
    var v21 int16 = v3
    fmt.Println(v21)
    v22 := GenS1{F0: 0}
    fmt.Println(v5, v3)
    if true {
        if false {
            fmt.Println(v5, v21)
        } else {
            if true {
                if (((3.14159 - 0.1) <= 1.5) || false) {
                    if (!((v2 << ((255) & 63)) < 0)) {
                        var v23 string = "hello"
                        switch -387 {
                        case 4294967295:
                            fmt.Println(v3)
                        case 128:
                            var v24 float64 = 0.1
                            _ = v24
                        }
                        _ = v23
                    } else {
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v26 := GenS2{F0: 42, F1: 65535, F2: (313 >= 2147483647)}
                            fmt.Println(fi25)
                            var v27 float64 = float64(-2.25)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                }
                fmt.Println(v2, v5)
                if true {
                    v5 -= 3.14159
                } else {
                    if false {
                        for fi28 := 0; fi28 < 12; fi28++ {
                            var v29 int8 = 42
                            for fi30 := 0; fi30 < 5; fi30++ {
                                switch (fi30 + (fi28 ^ (fi28 / ((32768) | 1)))) {
                                case 946:
                                    if true {
                                        v5 = float64((1.5 * -2.25))
                                        if (170 > -9223372036854775808) {
                                            v31 := GenS1{F0: 811}
                                            _ = v31
                                        }
                                        fmt.Println(fi28, v21, v4)
                                    } else {
                                        v4 = v3
                                        switch fi28 {
                                        case 158:
                                            fmt.Println(v2, v21)
                                            v32 := GenS2{F0: 105, F1: 2, F2: (2147483648 <= 1)}
                                            var v33 int8 = v29
                                            if (!false) {
                                                switch fi28 {
                                                case -518:
                                                    var v34 uint8 = 2
                                                    _ = v34
                                                case 4294967295:
                                                    if false {
                                                        fmt.Println(v4)
                                                        v33 = 92
                                                        v35 := GenS1{F0: 0}
                                                        _ = v35
                                                    } else {
                                                        v33 -= (v33 - (v29 & (v29 << ((uint(v2)) & 63))))
                                                        v3 = v3
                                                        if (149 != 128) {
                                                            var v36 float32 = 0.1
                                                            _ = v36
                                                        }
                                                    }
                                                default:
                                                    v17 = append(v17, 569)
                                                }
                                                var v37 bool = ("foo bar" < "a")
                                                v2 = (v2 ^ 255)
                                                _ = v37
                                            } else {
                                                v38 := GenS3{F0: 42, F1: 4294967295, F2: ("" + ("" + "a"))}
                                                _ = v38
                                            }
                                            var v39 uint64 = 754
                                            _ = v32
                                            _ = v33
                                            _ = v39
                                        case 1:
                                            var v40 int32 = int32(v29)
                                            _ = v40
                                        case 127:
                                            if (3.14159 <= (0.0 - (0.1 + -2.25))) {
                                                switch "" {
                                                case "foo bar":
                                                    var v41 uint = 4294967295
                                                    _ = v41
                                                case "\n\t":
                                                    var v42 string = "hello"
                                                    _ = v42
                                                }
                                            } else {
                                                v5 -= (float64(0.0) * 1e10)
                                                var v43 string = ("" + (("foo bar" + "a") + ""))
                                                fmt.Println(v43)
                                                _ = v43
                                            }
                                        }
                                    }
                                case 570:
                                    _ = len(v17)
                                }
                                _ = fi30
                            }
                            _ = fi28
                            _ = v29
                        }
                        if true {
                            if len(v17) > 0 {
                                v17[0] = 2147483647
                            }
                            v44 := GenS2{F0: 141, F1: (v2 / ((v2) | 1)), F2: true}
                            var v45 float64 = float64(0.0)
                            _ = v44
                            _ = v45
                        } else {
                            if false {
                                v46 := GenS4{F0: "a", F1: -793, F2: (v4 | (^(v3)))}
                                v47 := GenS1{F0: 2}
                                _ = v46
                                _ = v47
                            }
                        }
                        var v48 int = -9223372036854775808
                        _ = v48
                    } else {
                        v21 -= (v3 >> ((2) & 63))
                        for fi49 := 0; fi49 < 3; fi49++ {
                            var v50 uint8 = 1
                            fmt.Println(v4)
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            _ = fi49
                            _ = v50
                        }
                    }
                    if ("" > "\n\t") {
                        var v51 int8 = 68
                        var v52 uint16 = (v2 % (((v2 * (v2 + (v2 << ((1) & 63))))) | 1))
                        for fi53 := 0; fi53 < 8; fi53++ {
                            v54 := GenS5{F0: (fi53 * fi53), F1: 42, F2: uint(v52)}
                            var v55 uint16 = (v52 % (((v2 >> ((255) & 63))) | 1))
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        _ = v51
                        _ = v52
                    } else {
                        v1["hello"] = ("\n\t" + (("hello" + "hello") + "\n\t"))
                        switch int(v2) {
                        case 140:
                            switch ("a" + ("a" + (("a" + "\n\t") + ("\n\t" + ("" + ""))))) {
                            case "a":
                                if (!((!(!(51 >= 128))) && false)) {
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(v2, v3, v5)
                                    case "a":
                                        for fi56 := 0; fi56 < 12; fi56++ {
                                            switch "hello" {
                                            case "\n\t":
                                                var v57 int32 = -620
                                                _ = v57
                                            case "hello":
                                                fmt.Println(v3, v2, v4)
                                            case "a":
                                                if true {
                                                    var v58 int64 = 0
                                                    switch fi56 {
                                                    case 32768:
                                                        fmt.Println(v5)
                                                    case -881:
                                                        if false {
                                                            var v59 bool = false
                                                            switch (fi56 * (fi56 - (fi56 >> ((777) & 63)))) {
                                                            case 4294967295:
                                                                var v60 int = fi56
                                                                _ = v60
                                                            default:
                                                                var v61 bool = (v59 || (v59 && true))
                                                                _ = v61
                                                            }
                                                            switch 32767 {
                                                            case -624:
                                                                var v62 uint16 = (v2 * (^(v2)))
                                                                _ = v62
                                                            }
                                                            _ = v59
                                                        } else {
                                                            var v63 int64 = v58
                                                            for fi64 := 0; fi64 < 17; fi64++ {
                                                                v65 := GenS3{F0: 100, F1: 1, F2: ""}
                                                                _ = fi64
                                                                _ = v65
                                                            }
                                                            if len(v17) > 0 {
                                                                v17[0] = int(v2)
                                                            }
                                                            _ = v63
                                                        }
                                                    }
                                                    _ = v58
                                                }
                                            }
                                            var v66 int = (fi56 << ((42) & 63))
                                            for fi67 := 0; fi67 < 18; fi67++ {
                                                var v68 string = ""
                                                var v69 int32 = 1
                                                v70 := GenS5{F0: v66, F1: 291, F2: 2}
                                                _ = fi67
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            }
                                            _ = fi56
                                            _ = v66
                                        }
                                    case "foo bar":
                                        if true {
                                            for fi71 := 0; fi71 < 10; fi71++ {
                                                if false {
                                                    v4 = v21
                                                    var v72 float64 = (v5 - v5)
                                                    _ = v72
                                                }
                                                _ = fi71
                                            }
                                            for fi73 := 0; fi73 < 6; fi73++ {
                                                var v74 float64 = (v5 / v5)
                                                switch "" {
                                                case "a":
                                                    v2 = (v2 >> ((828) & 63))
                                                }
                                                var v75 int16 = 1000
                                                _ = fi73
                                                _ = v74
                                                _ = v75
                                            }
                                            if true {
                                                if (!true) {
                                                    var v76 bool = (-128 < -128)
                                                    var v77 uint64 = 0
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    fmt.Println(v5)
                                                    switch "\n\t" {
                                                    case "":
                                                        for fi78 := 0; fi78 < 16; fi78++ {
                                                            if len(v17) >= 2 {
                                                                v17 = v17[1:len(v17)]
                                                            }
                                                            v21 = (v21 >> ((42) & 63))
                                                            v5 = 1.5
                                                            _ = fi78
                                                        }
                                                    case "foo bar":
                                                        switch 127 {
                                                        case 32768:
                                                            var v79 float32 = 0.1
                                                            _ = v79
                                                        case 9223372036854775807:
                                                            v3 -= v4
                                                        case 225:
                                                            var v80 uint16 = v2
                                                            _ = v80
                                                        default:
                                                            v81 := v1[("foo bar" + "")]
                                                            _ = v81
                                                        }
                                                    default:
                                                        if len(v17) > 0 {
                                                            v17[0] = -2147483648
                                                        }
                                                    }
                                                }
                                            } else {
                                                _ = len(v1)
                                                if false {
                                                    switch 505 {
                                                    case 32768:
                                                        var v82 int = -773
                                                        _ = v82
                                                    case 0:
                                                        var v83 int8 = 42
                                                        _ = v83
                                                    }
                                                    var v84 string = ("a" + "hello")
                                                    var v85 int8 = 38
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    if true {
                                                        fmt.Println(v21)
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v4 = (v4 ^ 809)
                                                        default:
                                                            v86 := v1[("a" + "hello")]
                                                            _ = v86
                                                        }
                                                    }
                                                    var v87 int64 = -801
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    _ = v87
                                                }
                                            }
                                        } else {
                                            fmt.Println(v2, v2, v3)
                                        }
                                    }
                                    switch ("a" + "\n\t") {
                                    case "":
                                        v88 := GenS2{F0: 42, F1: v2, F2: (!(739 >= 880))}
                                        _ = v88
                                    default:
                                        v89 := GenS4{F0: "\n\t", F1: -132, F2: -32768}
                                        _ = v89
                                    }
                                    for fi90 := 0; fi90 < 8; fi90++ {
                                        fmt.Println(fi90)
                                        if true {
                                            var v91 int = fi90
                                            switch v91 {
                                            case -9223372036854775808:
                                                v1["\n\t"] = ((("" + ("a" + "foo bar")) + "foo bar") + "")
                                            case 42:
                                                v3 = -1
                                            case -1:
                                                switch "" {
                                                case "a":
                                                    if false {
                                                        v3 = -228
                                                        switch (fi90 << ((uint(v2)) & 63)) {
                                                        case -611:
                                                            v91 = v91
                                                        }
                                                        var v92 int8 = -12
                                                        _ = v92
                                                    }
                                                case "foo bar":
                                                    switch 32767 {
                                                    case 4294967295:
                                                        switch (fi90 / ((v91) | 1)) {
                                                        case -9223372036854775808:
                                                            v3 -= (v4 | v3)
                                                        }
                                                    }
                                                default:
                                                    v21 -= (-(v21))
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            _ = v91
                                        }
                                        v17 = append(v17, fi90)
                                        _ = fi90
                                    }
                                }
                            }
                        case 4294967295:
                            for fi93 := 0; fi93 < 2; fi93++ {
                                for fi94 := 0; fi94 < 9; fi94++ {
                                    if false {
                                        var v95 uint = 752
                                        var v96 float64 = 3.14159
                                        var v97 uint8 = 107
                                        _ = v95
                                        _ = v96
                                        _ = v97
                                    } else {
                                        var v98 int32 = -2147483648
                                        v2 -= v2
                                        _ = v98
                                    }
                                    var v99 uint32 = 42
                                    var v100 uint8 = 2
                                    _ = fi94
                                    _ = v100
                                    _ = v99
                                }
                                fmt.Println(v5, v2, v2)
                                _ = fi93
                            }
                        default:
                            var v101 int = -32
                            _ = v101
                        }
                    }
                    v2 = v2
                }
            } else {
                v2 += 255
            }
        }
    }
    if true {
        var v102 bool = ((false || (61 <= 0)) && true)
        var v103 int16 = v21
        var v104 uint16 = v2
        _ = v102
        _ = v103
        _ = v104
    }
    v105 := GenS5{F0: -1, F1: 187, F2: 65535}
    var v106 uint64 = 1
    if (!(v5 >= v5)) {
        v5 -= float64(float32(v5))
        var v107 int8 = genFn0(2147483647, -1)
        if true {
            v106 -= 938
        }
        _ = v107
    }
    var v108 int8 = int8(v106)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v105
    _ = v106
    _ = v108
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
