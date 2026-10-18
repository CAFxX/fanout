package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 uint8; F1 uint8; F2 string }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]string)
    var v3 bool = (false || (1e10 >= float32(1e10)))
    var v4 bool = false
    fmt.Println(v4)
    v5 := GenS1{F0: 173, F1: 301, F2: 2147483647}
    switch 182 {
    case 0:
        v6 := []int{32768, 42}
        _ = v6
    case 905:
        var v7 uint = 2
        _ = v7
    case 2147483647:
        v3 = v4
    default:
        var v8 uint = 786
        _ = v8
    }
    if ((81 == 42) || v4) {
        switch ((("hello" + "a") + "hello") + "foo bar") {
        case "hello":
            v9 := GenS1{F0: 42, F1: 42, F2: 42}
            _ = v9
        case "":
            var v10 uint64 = 829
            _ = v10
        case "foo bar":
            v11 := GenS2{F0: 255, F1: -439, F2: "hello"}
            _ = v11
        }
        var v12 int32 = 0
        _ = v12
    } else {
        if (2 >= 620) {
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 bool = (1 > 14)
                _ = fi13
                _ = v14
            }
            if (v4 && v4) {
                var v15 int8 = -93
                v4 = (65535 < 0)
                switch 2147483647 {
                case -515:
                    if (("foo bar" < "hello") || v4) {
                        fmt.Println(v4, v3, v4)
                    }
                case 594:
                    if (v1 == v1) {
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            for fi16 := 0; fi16 < 19; fi16++ {
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    fmt.Println(fi17, fi16, v4)
                                    if (2147483647 <= 1) {
                                        var v18 int8 = v15
                                        _ = v18
                                    } else {
                                        var v19 uint8 = 131
                                        var v20 string = "a"
                                        var v21 bool = (-32768 < 190)
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                    v15 = 127
                                    _ = fi17
                                }
                                v22 := GenS2{F0: 101, F1: -9223372036854775808, F2: ("\n\t" + "")}
                                var v23 uint16 = 65535
                                _ = fi16
                                _ = v22
                                _ = v23
                            }
                        case "foo bar":
                            var v24 uint32 = 2147483647
                            _ = v24
                        case "":
                            for fi25 := 0; fi25 < 20; fi25++ {
                                if true {
                                    for fi26 := 0; fi26 < 18; fi26++ {
                                        v4 = (false && false)
                                        fmt.Println(v4, v3, fi25)
                                        if (0 == 42) {
                                            var v27 float64 = float64(-2.25)
                                            v1 = (float64((3.14159 + 1.5)) * v1)
                                            v28 := GenS2{F0: 0, F1: -9223372036854775808, F2: ("a" + ((("\n\t" + ("\n\t" + "\n\t")) + ("a" + "foo bar")) + "\n\t"))}
                                            _ = v27
                                            _ = v28
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                var v29 float32 = float32(v1)
                                                _ = v29
                                            case "foo bar":
                                                v1 = 1e10
                                            case "hello":
                                                var v30 int16 = 335
                                                var v31 uint = 0
                                                _ = v30
                                                _ = v31
                                            }
                                            if ((349 == 32767) && (2 >= 0)) {
                                                v32 := GenS1{F0: 42, F1: 826, F2: fi26}
                                                fmt.Println(v15, fi26)
                                                _ = v32
                                            } else {
                                                var v33 uint8 = 65
                                                var v34 float64 = float64((0.1 * 1e10))
                                                if (false && (!((0.0 >= 3.14159) && (!(v4 && ((330 <= 433) || v4)))))) {
                                                    var v35 int16 = 42
                                                    if v4 {
                                                        v36 := GenS3{F0: 42}
                                                        _ = v36
                                                    }
                                                    _ = v35
                                                } else {
                                                    v1 = float64(3.14159)
                                                }
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        v1 += v1
                                        switch fi25 {
                                        case -9223372036854775808:
                                            v15 = v15
                                        case 0:
                                            var v37 uint64 = 758
                                            _ = v37
                                        }
                                        switch 42 {
                                        case -43:
                                            var v38 int = (fi26 + (-(fi26)))
                                            _ = v38
                                        case 9223372036854775807:
                                            if (0 == 2) {
                                                v15 = v15
                                                var v39 float64 = ((((v1 / v1) - 0.0) * v1) * v1)
                                                _ = v39
                                            } else {
                                                var v40 bool = v3
                                                switch "" {
                                                case "":
                                                    var v41 int64 = 1
                                                    switch (fi25 >> ((624) & 63)) {
                                                    case -656:
                                                        v4 = ("\n\t" > "\n\t")
                                                    case 0:
                                                        v42 := []int{fi25}
                                                        _ = v42
                                                    case 1:
                                                        var v43 uint64 = 18446744073709551615
                                                        _ = v43
                                                    }
                                                    switch (("a" + ("a" + "hello")) + "a") {
                                                    case "":
                                                        if v4 {
                                                            var v44 int8 = v15
                                                            v4 = false
                                                            v45 := GenS1{F0: 0, F1: v41, F2: (fi25 + (fi26 / ((fi25) | 1)))}
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    case "foo bar":
                                                        var v46 uint16 = 0
                                                        _ = v46
                                                    case "hello":
                                                        v47 := GenS4{F0: (v4 || ((v41 / (((^(v41))) | 1)) < (v41 * 385)))}
                                                        _ = v47
                                                    default:
                                                        v48 := GenS5{F0: (fi26 | (fi26 >> ((1) & 63))), F1: fi26}
                                                        _ = v48
                                                    }
                                                    _ = v41
                                                case "foo bar":
                                                    v15 += (v15 * (v15 & (v15 * (v15 - (v15 & 42)))))
                                                case "hello":
                                                    if true {
                                                        v1 -= (v1 / v1)
                                                    } else {
                                                        if ("foo bar" <= (("foo bar" + "a") + "a")) {
                                                            v49 := []int{4294967295, fi26}
                                                            switch fi25 {
                                                            case 560:
                                                                var v50 int64 = 9223372036854775807
                                                                _ = v50
                                                            default:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    fmt.Println(fi26, v3, v15)
                                                                }
                                                            }
                                                            _ = v49
                                                        } else {
                                                            var v51 uint = 10
                                                            _ = v51
                                                        }
                                                        var v52 int = 128
                                                        fmt.Println(v15, v4, v3)
                                                        _ = v52
                                                    }
                                                default:
                                                    v1 = v1
                                                }
                                                _ = v40
                                            }
                                        case 32768:
                                            switch ("a" + ("foo bar" + "hello")) {
                                            case "a":
                                                v53 := []int{}
                                                _ = v53
                                            default:
                                                fmt.Println(v4)
                                            }
                                        default:
                                            v54 := GenS4{F0: v4}
                                            _ = v54
                                        }
                                        var v55 int16 = 32767
                                        _ = fi26
                                        _ = v55
                                    }
                                    var v56 int = (fi25 << ((2) & 63))
                                    _ = v56
                                } else {
                                    fmt.Println(v1, fi25)
                                }
                                v15 -= (v15 ^ (v15 - v15))
                                _ = fi25
                            }
                        default:
                            var v57 int = -861
                            _ = v57
                        }
                        v4 = (!true)
                        if (9223372036854775807 < 2147483647) {
                            v58 := GenS6{F0: int64(v15), F1: 0}
                            fmt.Println(v3, v3, v4)
                            fmt.Println(v4)
                            _ = v58
                        } else {
                            for fi59 := 0; fi59 < 11; fi59++ {
                                v3 = v3
                                v1 = float64(3.14159)
                                if (0 >= 2) {
                                    var v60 string = ""
                                    _ = v60
                                } else {
                                    var v61 int = (fi59 >> ((2) & 63))
                                    _ = v61
                                }
                                _ = fi59
                            }
                        }
                    }
                case 9223372036854775807:
                    switch "a" {
                    case "hello":
                        v62 := GenS3{F0: 2}
                        _ = v62
                    case "a":
                        v63 := v2[1]
                        _ = v63
                    default:
                        v3 = (v4 || v3)
                    }
                default:
                    for fi64 := 0; fi64 < 14; fi64++ {
                        for fi65 := 0; fi65 < 4; fi65++ {
                            v66 := GenS6{F0: -9223372036854775808, F1: (fi64 - fi65)}
                            fmt.Println(v15, v1, v1)
                            _ = fi65
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 15; fi67++ {
                            v68 := []int{fi64, -1, fi64, 42}
                            if (((fi64 ^ (fi64 - (fi64 | (fi64 << ((853) & 63))))) > fi64) && false) {
                                v69 := GenS6{F0: 498, F1: (fi64 - (-(fi67)))}
                                _ = v69
                            }
                            var v70 uint32 = 2
                            v3 = (255 == 820)
                            _ = fi67
                            _ = v68
                            _ = v70
                        }
                        if (961 != int64(fi64)) {
                            v71 := GenS1{F0: 18446744073709551615, F1: -1, F2: 911}
                            _ = v71
                        }
                        _ = fi64
                    }
                }
                _ = v15
            }
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 string = "foo bar"
                if (-27 == 42) {
                    if false {
                        switch (v73 + v73) {
                        case "foo bar":
                            for fi74 := 0; fi74 < 9; fi74++ {
                                var v75 uint16 = 297
                                v76 := []int{fi72, 4294967295}
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                        case "a":
                            v77 := GenS7{F0: 18446744073709551615, F1: -1}
                            _ = v77
                        case "":
                            var v78 uint64 = 0
                            _ = v78
                        }
                    }
                    var v79 uint32 = 2
                    if false {
                        for fi80 := 0; fi80 < 10; fi80++ {
                            v4 = (!(!true))
                            if (570 == 1) {
                                delete(v2, v79)
                                var v81 int = fi80
                                if (-2147483648 >= -2147483648) {
                                    switch "\n\t" {
                                    case "a":
                                        var v82 int = (fi72 & 32768)
                                        _ = v82
                                    case "\n\t":
                                        v79 = (v79 + v79)
                                    case "foo bar":
                                        var v83 float32 = 1.5
                                        _ = v83
                                    default:
                                        if false {
                                            v79 -= v79
                                        } else {
                                            var v84 int = (v81 + fi72)
                                            v4 = (v4 && true)
                                            fmt.Println(v81, v1, v81)
                                            _ = v84
                                        }
                                    }
                                }
                                _ = v81
                            } else {
                                switch (v73 + v73) {
                                case "a":
                                    var v85 uint64 = 42
                                    _ = v85
                                }
                                v1 = (0.0 / v1)
                                if (!(65535 > 0)) {
                                    v73 = v73
                                } else {
                                    v3 = v3
                                    var v86 int16 = 1
                                    _ = v86
                                }
                            }
                            var v87 string = v73
                            _ = fi80
                            _ = v87
                        }
                        var v88 int16 = -32768
                        switch (fi72 / ((fi72) | 1)) {
                        case -844:
                            v4 = v4
                        case 2147483647:
                            if false {
                                if (128 != 1) {
                                    switch (("hello" + v73) + (((v73 + "\n\t") + ("" + "\n\t")) + v73)) {
                                    case "\n\t":
                                        if v4 {
                                            v79 -= 285
                                            for fi89 := 0; fi89 < 13; fi89++ {
                                                var v90 uint8 = 1
                                                _ = fi89
                                                _ = v90
                                            }
                                        }
                                    default:
                                        for fi91 := 0; fi91 < 6; fi91++ {
                                            switch (-(fi72)) {
                                            case 899:
                                                var v92 uint64 = 558
                                                if v3 {
                                                    fmt.Println(v1, v4)
                                                } else {
                                                    var v93 uint8 = 127
                                                    switch "hello" {
                                                    case "a":
                                                        v92 = 1
                                                        var v94 uint = 607
                                                        v3 = false
                                                        _ = v94
                                                    case "foo bar":
                                                        var v95 float32 = (3.14159 + (-2.25 / 1.5))
                                                        _ = v95
                                                    case "\n\t":
                                                        _ = len(v2)
                                                    default:
                                                        v73 = (((v73 + v73) + "foo bar") + ((v73 + v73) + ((v73 + ("foo bar" + (v73 + v73))) + ((v73 + v73) + ""))))
                                                    }
                                                    _ = v93
                                                }
                                                v1 += v1
                                                _ = v92
                                            default:
                                                v79 += v79
                                            }
                                            v96 := []int{}
                                            _ = fi91
                                            _ = v96
                                        }
                                    }
                                }
                                v97 := GenS8{F0: (v79 * (v79 << ((2) & 63))), F1: uint64(v79)}
                                _ = v97
                            } else {
                                for fi98 := 0; fi98 < 9; fi98++ {
                                    if v4 {
                                        v99 := v2[v79]
                                        _ = v99
                                    } else {
                                        var v100 uint64 = 1
                                        _ = v100
                                    }
                                    _ = fi98
                                }
                                switch ((("foo bar" + ((v73 + ((("" + "") + "foo bar") + "a")) + v73)) + ((("foo bar" + v73) + (("a" + (v73 + "a")) + v73)) + v73)) + v73) {
                                case "a":
                                    v88 = (v88 << ((42) & 63))
                                case "foo bar":
                                    v79 = 2147483647
                                default:
                                    var v101 uint64 = 2
                                    _ = v101
                                }
                            }
                        }
                        _ = v88
                    }
                    _ = v79
                } else {
                    var v102 uint8 = 70
                    v103 := GenS9{F0: v102, F1: (v102 * (v102 >> ((130) & 63))), F2: v73}
                    _ = v102
                    _ = v103
                }
                _ = fi72
                _ = v73
            }
        } else {
            var v104 int32 = -284
            var v105 int16 = 32767
            v3 = ((2 == 65535) && (("" + "hello") != ""))
            _ = v104
            _ = v105
        }
    }
    v4 = v4
    fmt.Println(v3)
    var v106 int = 128
    var v107 uint64 = 403
    var v108 int8 = -110
    for fi109 := 0; fi109 < 6; fi109++ {
        v110 := GenS8{F0: 990, F1: 2}
        for fi111 := 0; fi111 < 12; fi111++ {
            var v112 uint8 = 142
            fmt.Println(fi111, v107, v4)
            v113 := []int{v106, (fi109 << ((65535) & 63)), 32767, fi111}
            v1 = float64(3.14159)
            _ = fi111
            _ = v112
            _ = v113
        }
        switch "foo bar" {
        case "hello":
            switch fi109 {
            case -2147483648:
                v114 := []int{fi109, 9223372036854775807, fi109, (^(fi109))}
                _ = v114
            case 42:
                var v115 int16 = 32767
                _ = v115
            }
        }
        _ = fi109
        _ = v110
    }
    var v116 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v107
    _ = v108
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
