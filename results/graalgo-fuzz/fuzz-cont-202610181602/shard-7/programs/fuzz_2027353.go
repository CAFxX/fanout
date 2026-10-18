package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 uint; F2 uint16 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 int8; F1 int; F2 uint }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 int; F1 int64 }
type GenS12 struct { F0 int; F1 int16; F2 uint64 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 float64 = (3.14159 - -2.25)
    var v5 int16 = -29
    var v6 int32 = 42
    fmt.Println(v5)
    v5 = (v5 + v5)
    switch -206 {
    case 9223372036854775807:
        v7 := []int{42, 9223372036854775807, 2147483647, 1}
        _ = v7
    case 32768:
        v4 = (v4 + float64(0.0))
    default:
        fmt.Println(v5, v6)
    }
    var v8 bool = true
    _ = len(v3)
    for fi9 := 0; fi9 < 2; fi9++ {
        if v8 {
            if false {
                if (765 < 613) {
                    if false {
                        if (459 != 682) {
                            for fi10 := 0; fi10 < 10; fi10++ {
                                v2[(fi9 | int(v5))] = 42
                                if (-506 > 35) {
                                    switch "" {
                                    case "":
                                        switch (^(fi10)) {
                                        case 128:
                                            var v11 int16 = (v5 / (((v5 << ((922) & 63))) | 1))
                                            _ = v11
                                        case 903:
                                            var v12 uint32 = 1
                                            v6 = v6
                                            _ = v12
                                        case 42:
                                            v13 := GenS1{F0: 100}
                                            v8 = (true || ((v4 == v4) || v8))
                                            var v14 string = ("\n\t" + ("\n\t" + "hello"))
                                            _ = v13
                                            _ = v14
                                        }
                                    case "foo bar":
                                        fmt.Println(v8, v8)
                                    case "\n\t":
                                        v15 := []int{}
                                        _ = v15
                                    default:
                                        v1 -= v1
                                    }
                                    var v16 int8 = 1
                                    v17 := []int{fi9}
                                    v18 := GenS1{F0: v16}
                                    var v19 string = "a"
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 int = (fi9 << ((uint(v1)) & 63))
                                    v21 := GenS2{F0: 18446744073709551615, F1: ((v6 >> ((2) & 63)) <= 2147483647), F2: 2}
                                    _ = v20
                                    _ = v21
                                }
                                v22 := v3[568]
                                v5 = (v5 / (((v5 ^ (v5 % ((v5) | 1)))) | 1))
                                _ = fi10
                                _ = v22
                            }
                            if (!v8) {
                                if v8 {
                                    v6 = (v6 % ((v6) | 1))
                                    v23 := GenS3{F0: v5}
                                    v1 -= 563
                                    _ = v23
                                }
                            } else {
                                if (v8 || (((int8(v6) >= 100) && false) || (v1 <= v1))) {
                                    for fi24 := 0; fi24 < 4; fi24++ {
                                        var v25 bool = v8
                                        v26 := GenS4{F0: 6, F1: 18446744073709551615}
                                        var v27 int64 = 0
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    for fi28 := 0; fi28 < 14; fi28++ {
                                        v29 := GenS5{F0: -128}
                                        v30 := []int{(-(fi9)), fi28, (-(fi28))}
                                        var v31 uint64 = 4294967295
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    if v8 {
                                        fmt.Println(fi9, v8, v6)
                                        var v32 float32 = (0.1 * 1.5)
                                        _ = v32
                                    } else {
                                        v6 = -2147483648
                                    }
                                } else {
                                    v3[1] = "hello"
                                    var v33 int = fi9
                                    _ = v33
                                }
                                v2[fi9] = 65535
                            }
                        } else {
                            if (!(42 == 796)) {
                                if (!false) {
                                    for fi34 := 0; fi34 < 11; fi34++ {
                                        var v35 string = "\n\t"
                                        fmt.Println(v5)
                                        var v36 uint64 = 42
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                    var v37 float32 = 3.14159
                                    switch fi9 {
                                    case -1:
                                        switch "" {
                                        case "foo bar":
                                            v38 := GenS4{F0: 1, F1: 4294967295}
                                            _ = v38
                                        case "":
                                            if ((-8 == 42) || v8) {
                                                fmt.Println(v4, v6)
                                            }
                                        case "hello":
                                            v8 = (1 < 220)
                                        default:
                                            if v8 {
                                                v2[fi9] = uint32(v1)
                                                var v39 uint32 = 0
                                                v40 := []int{32767, (fi9 & -2147483648)}
                                                _ = v39
                                                _ = v40
                                            } else {
                                                v2[(^(fi9))] = 4294967295
                                                v5 = (v5 * -1)
                                                v41 := []int{(fi9 + fi9)}
                                                _ = v41
                                            }
                                        }
                                    case -475:
                                        var v42 float64 = 0.0
                                        _ = v42
                                    default:
                                        if v8 {
                                            var v43 uint64 = 941
                                            if (!(65535 <= 42)) {
                                                var v44 float64 = v4
                                                if (!(333 < (v1 + v1))) {
                                                    v45 := GenS4{F0: 128, F1: 0}
                                                    if v8 {
                                                        v46 := []int{}
                                                        var v47 uint = 942
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        var v48 string = "hello"
                                                        for fi49 := 0; fi49 < 19; fi49++ {
                                                            v50 := GenS1{F0: 127}
                                                            _ = fi49
                                                            _ = v50
                                                        }
                                                        v48 = (v48 + ((v48 + "") + v48))
                                                        _ = v48
                                                    }
                                                    var v51 int16 = (-(v5))
                                                    _ = v45
                                                    _ = v51
                                                } else {
                                                    _ = len(v2)
                                                }
                                                _ = v44
                                            }
                                            _ = v43
                                        } else {
                                            for fi52 := 0; fi52 < 7; fi52++ {
                                                switch ("foo bar" + ("" + "")) {
                                                case "":
                                                    var v53 int16 = (v5 | v5)
                                                    _ = v53
                                                }
                                                if ((-(v6)) < (v6 * v6)) {
                                                    fmt.Println(fi9, v37, v37)
                                                    if (1.5 > v37) {
                                                        v6 = (v6 + (v6 | -2147483648))
                                                        v6 = v6
                                                        var v54 int32 = v6
                                                        _ = v54
                                                    } else {
                                                        fmt.Println(v8, v5)
                                                        v55 := GenS6{F0: 127, F1: 65535, F2: v1}
                                                        _ = v55
                                                    }
                                                } else {
                                                    var v56 int32 = (v6 + int32(v5))
                                                    _ = v56
                                                }
                                                if false {
                                                    fmt.Println(v8, v5)
                                                    if (100 != 115) {
                                                        var v57 uint8 = 128
                                                        _ = v57
                                                    }
                                                }
                                                v5 = (^(v5))
                                                _ = fi52
                                            }
                                            switch (fi9 ^ (fi9 << ((1) & 63))) {
                                            case 128:
                                                var v58 int32 = (v6 % ((v6) | 1))
                                                _ = v58
                                            default:
                                                switch "" {
                                                case "a":
                                                    fmt.Println(v1)
                                                case "foo bar":
                                                    var v59 uint32 = 2
                                                    _ = v59
                                                case "hello":
                                                    delete(v3, 42)
                                                default:
                                                    v1 = v1
                                                }
                                            }
                                        }
                                    }
                                    _ = v37
                                } else {
                                    if (v4 <= 0.1) {
                                        var v60 int32 = v6
                                        for fi61 := 0; fi61 < 1; fi61++ {
                                            var v62 uint8 = 125
                                            v8 = false
                                            var v63 int32 = 0
                                            v6 = (v60 + v6)
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                        fmt.Println(v6, v4)
                                        _ = v60
                                    }
                                    var v64 uint8 = 2
                                    _ = v64
                                }
                                switch fi9 {
                                case -9223372036854775808:
                                    for fi65 := 0; fi65 < 9; fi65++ {
                                        var v66 float64 = (v4 / float64((-2.25 / float32(float64(float32(float64(-2.25)))))))
                                        v8 = (65535 < 4294967295)
                                        if (579 == 4294967295) {
                                            switch "a" {
                                            case "\n\t":
                                                if v8 {
                                                    switch 127 {
                                                    case 4294967295:
                                                        var v67 uint = 213
                                                        _ = v67
                                                    default:
                                                        if v8 {
                                                            v68 := GenS7{F0: false, F1: ""}
                                                            fmt.Println(v6)
                                                            switch "a" {
                                                            case "":
                                                                var v69 uint8 = 128
                                                                _ = v69
                                                            }
                                                            _ = v68
                                                        }
                                                    }
                                                    v70 := []int{fi9, 42}
                                                    _ = len(v2)
                                                    _ = v70
                                                }
                                            case "a":
                                                switch fi65 {
                                                case 427:
                                                    if false {
                                                        switch (("foo bar" + "hello") + "hello") {
                                                        case "hello":
                                                            if false {
                                                                v71 := GenS3{F0: v5}
                                                                if v8 {
                                                                    v4 -= v4
                                                                } else {
                                                                    fmt.Println(v6, v66, fi9)
                                                                    var v72 string = ""
                                                                    _ = v72
                                                                }
                                                                var v73 float64 = v66
                                                                _ = v71
                                                                _ = v73
                                                            } else {
                                                                var v74 uint64 = 2
                                                                switch ("a" + ("hello" + "\n\t")) {
                                                                case "hello":
                                                                    switch (fi9 >> ((605) & 63)) {
                                                                    case -598:
                                                                        fmt.Println(v6, fi9, v6)
                                                                    case 985:
                                                                        if v8 {
                                                                            switch "a" {
                                                                            case "a":
                                                                                v75 := GenS8{F0: 224, F1: (0 >= 960)}
                                                                                _ = v75
                                                                            case "\n\t":
                                                                                if (v8 && (v8 && (uint16(v74) < (v1 & v1)))) {
                                                                                    if v8 {
                                                                                        switch (fi9 >> ((255) & 63)) {
                                                                                        case -9223372036854775808:
                                                                                            if (v8 || true) {
                                                                                                v5 -= 22
                                                                                                _ = len(v3)
                                                                                            }
                                                                                        }
                                                                                        v76 := GenS1{F0: 72}
                                                                                        v1 = 42
                                                                                        _ = v76
                                                                                    }
                                                                                    v77 := GenS1{F0: 127}
                                                                                    if (0 >= 4294967295) {
                                                                                        _ = len(v3)
                                                                                    } else {
                                                                                        if (!false) {
                                                                                            v78 := GenS3{F0: v5}
                                                                                            if true {
                                                                                                v5 -= (v5 | v5)
                                                                                                var v79 int32 = -2147483648
                                                                                                _ = v79
                                                                                            }
                                                                                            var v80 uint32 = 604
                                                                                            _ = v78
                                                                                            _ = v80
                                                                                        }
                                                                                        delete(v2, 734)
                                                                                        var v81 uint = 18446744073709551615
                                                                                        _ = v81
                                                                                    }
                                                                                    _ = v77
                                                                                }
                                                                            default:
                                                                                v82 := GenS1{F0: 126}
                                                                                _ = v82
                                                                            }
                                                                            v83 := GenS7{F0: (v6 == (v6 + v6)), F1: "hello"}
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                case "a":
                                                                    switch ("hello" + ("hello" + "hello")) {
                                                                    case "":
                                                                        var v84 int32 = (v6 ^ v6)
                                                                        _ = v84
                                                                    default:
                                                                        v6 -= v6
                                                                    }
                                                                case "\n\t":
                                                                    var v85 uint = 42
                                                                    _ = v85
                                                                default:
                                                                    var v86 bool = (-128 == -128)
                                                                    _ = v86
                                                                }
                                                                if true {
                                                                    v87 := GenS9{F0: 1, F1: fi65, F2: 65535}
                                                                    v88 := GenS9{F0: 119, F1: 4294967295, F2: 804}
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                                _ = v74
                                                            }
                                                            if v8 {
                                                                if false {
                                                                    v89 := GenS3{F0: -1}
                                                                    _ = v89
                                                                } else {
                                                                    v90 := GenS7{F0: v8, F1: ""}
                                                                    v5 = (v5 | v5)
                                                                    fmt.Println(fi65, v8)
                                                                    _ = v90
                                                                }
                                                                var v91 int8 = 23
                                                                _ = v91
                                                            }
                                                        case "a":
                                                            v4 -= (v66 + (v4 - v66))
                                                        case "foo bar":
                                                            v6 -= (v6 - v6)
                                                        }
                                                    } else {
                                                        var v92 string = "\n\t"
                                                        v93 := GenS10{F0: 449, F1: ((-211 <= 9223372036854775807) || false)}
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                }
                                            default:
                                                v4 += v66
                                            }
                                        }
                                        v94 := GenS9{F0: -16, F1: (fi9 - (fi9 ^ (fi65 << ((2) & 63)))), F2: 820}
                                        _ = fi65
                                        _ = v66
                                        _ = v94
                                    }
                                case -2147483648:
                                    if (!(("\n\t" + ("\n\t" + "\n\t")) == ("foo bar" + "\n\t"))) {
                                        v95 := GenS3{F0: (v5 & 32767)}
                                        v96 := GenS11{F0: (fi9 >> ((537) & 63)), F1: 1}
                                        if v8 {
                                            var v97 uint8 = 0
                                            v5 = (v5 - 1000)
                                            _ = v97
                                        }
                                        _ = v95
                                        _ = v96
                                    } else {
                                        fmt.Println(v6, v5)
                                    }
                                case 32768:
                                    v8 = true
                                }
                            }
                        }
                        for fi98 := 0; fi98 < 17; fi98++ {
                            if v8 {
                                if v8 {
                                    var v99 string = ""
                                    var v100 bool = false
                                    _ = v100
                                    _ = v99
                                } else {
                                    delete(v3, 42)
                                    v101 := []int{fi9, fi98}
                                    var v102 uint = 154
                                    _ = v101
                                    _ = v102
                                }
                            } else {
                                var v103 uint = 706
                                if v8 {
                                    switch -9223372036854775808 {
                                    case 2147483647:
                                        v4 -= v4
                                    }
                                    v4 = v4
                                    var v104 uint = v103
                                    _ = v104
                                } else {
                                    v105 := []int{fi9, (fi98 / (((fi9 * (fi9 | fi98))) | 1)), fi9}
                                    v103 += v103
                                    _ = v105
                                }
                                v106 := []int{fi9, fi98}
                                _ = v103
                                _ = v106
                            }
                            _ = fi98
                        }
                    } else {
                        fmt.Println(v4)
                        v5 = (v5 - (^(v5)))
                    }
                    var v107 uint16 = v1
                    if v8 {
                        var v108 int64 = 42
                        var v109 uint64 = 18446744073709551615
                        _ = v108
                        _ = v109
                    } else {
                        v110 := GenS3{F0: v5}
                        fmt.Println(v4, v4)
                        _ = v110
                    }
                    _ = v107
                }
            } else {
                v111 := GenS12{F0: fi9, F1: -1, F2: 2}
                _ = v111
            }
            var v112 int32 = (v6 - 0)
            v4 += v4
            _ = v112
        } else {
            var v113 uint32 = 2147483647
            fmt.Println(v1, v113, fi9)
            _ = v113
        }
        _ = fi9
    }
    var v114 uint16 = (v1 / (((v1 << ((1) & 63))) | 1))
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
