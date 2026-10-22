package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = ("" + "hello")
    var v2 int32 = 42
    var v3 string = v1
    var v4 uint32 = 2
    var v5 uint64 = 142
    v6 := []string{v3, v1}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v4, v5)
    case "a":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int16 = int16(v2)
            switch fi7 {
            case 0:
                var v9 uint32 = (v4 + (^(v4)))
                _ = v9
            case -255:
                _ = len(v6)
            case 32767:
                switch ((v3 + v3) + ((v1 + "") + "")) {
                case "\n\t":
                    fmt.Println(v3)
                case "hello":
                    v1 = ""
                default:
                    for fi10 := 0; fi10 < 18; fi10++ {
                        _ = len(v6)
                        switch (v1 + v3) {
                        case "foo bar":
                            v11 := GenS1{F0: "\n\t"}
                            _ = v11
                        case "":
                            v8 = 1
                        }
                        _ = fi10
                    }
                }
            default:
                var v12 uint8 = 0
                _ = v12
            }
            v13 := GenS1{F0: ("foo bar" + v3)}
            _ = fi7
            _ = v13
            _ = v8
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v5 = (v5 | v5)
        v4 = v4
        _ = fi14
    }
    v15 := GenS2{F0: 448, F1: (^(v5))}
    v16 := GenS1{F0: v3}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2)
    v1 = (v3 + (("foo bar" + "a") + v1))
    v5 += (v5 ^ v5)
    if (v5 >= v5) {
        _ = len(v6)
        v17 := GenS1{F0: v1}
        _ = v17
    }
    var v18 int16 = -32768
    v19 := GenS3{F0: true, F1: v4, F2: v2}
    if (!(127 != 0)) {
        if (v4 < (^(v4))) {
            var v20 bool = true
            for fi21 := 0; fi21 < 3; fi21++ {
                fmt.Println(v5)
                v2 -= int32(v5)
                var v22 int16 = (v18 << ((665) & 63))
                _ = fi21
                _ = v22
            }
            fmt.Println(v4)
            _ = v20
        } else {
            switch -861 {
            case -466:
                v23 := GenS3{F0: (((false || (true || ((!false) || false))) || false) || (v2 > 1)), F1: (v4 + 2147483648), F2: v2}
                _ = v23
            case 9223372036854775807:
                v2 -= v2
            case 42:
                for fi24 := 0; fi24 < 19; fi24++ {
                    if (fi24 != fi24) {
                        var v25 int16 = v18
                        _ = v25
                    }
                    fmt.Println(v3, v18)
                    _ = fi24
                }
            default:
                fmt.Println(v2)
            }
            v2 += -91
            var v26 int8 = 125
            _ = v26
        }
    } else {
        v2 -= v2
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 int32 = (v2 + 42)
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 uint8 = 128
                switch fi27 {
                case 42:
                    var v31 bool = true
                    _ = v31
                case 32768:
                    switch v3 {
                    case "a":
                        switch -167 {
                        case 589:
                            v4 = (^(v4))
                        }
                    case "hello":
                        var v32 int8 = 1
                        _ = v32
                    case "\n\t":
                        v2 = v2
                    }
                default:
                    v33 := GenS1{F0: "a"}
                    _ = v33
                }
                var v34 float32 = 1e10
                _ = fi29
                _ = v30
                _ = v34
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 string = "hello"
                if false {
                    v3 = "a"
                    if false {
                        v37 := GenS3{F0: false, F1: (v4 | v4), F2: v2}
                        v3 = v3
                        if true {
                            switch "a" {
                            case "hello":
                                switch v1 {
                                case "\n\t":
                                    v3 = ((v36 + ("\n\t" + ("hello" + "a"))) + "a")
                                    var v38 uint = 0
                                    v39 := GenS4{F0: 127, F1: 0}
                                    if (0 != (v18 | (^(v18)))) {
                                        v36 = ""
                                        var v40 uint64 = (^(v5))
                                        v1 = v3
                                        switch "" {
                                        case "foo bar":
                                            var v41 uint64 = v5
                                            _ = v41
                                        }
                                        _ = v40
                                    }
                                    _ = len(v6)
                                    _ = v38
                                    _ = v39
                                default:
                                    if (!true) {
                                        fmt.Println(v36, v28)
                                    } else {
                                        fmt.Println(v1, v36, v18)
                                    }
                                }
                            default:
                                v3 = (("a" + "hello") + v1)
                            }
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                fmt.Println(v1, v18, v36)
                                if (v3 >= (v3 + "a")) {
                                    var v42 string = "a"
                                    v1 = "a"
                                    v43 := GenS4{F0: 1, F1: 245}
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v44 int8 = 127
                                    var v45 int64 = 9223372036854775807
                                    if len(v6) > 0 {
                                        v6[0] = ("hello" + v36)
                                    }
                                    _ = v44
                                    _ = v45
                                }
                                v36 = "a"
                                var v46 uint = 2
                                if (2147483647 < v4) {
                                    var v47 int32 = (v2 % (((v28 & (v2 & v28))) | 1))
                                    switch (-(fi35)) {
                                    case 32768:
                                        switch -1 {
                                        case 138:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case 2147483647:
                                            if ((!false) && false) {
                                                var v48 bool = false
                                                var v49 int16 = (v18 << ((v46) & 63))
                                                v50 := GenS5{F0: v49, F1: v1}
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            } else {
                                                var v51 uint32 = (v4 + (v4 - (^(v4))))
                                                var v52 int64 = int64(v51)
                                                v5 += 18446744073709551615
                                                _ = v51
                                                _ = v52
                                            }
                                        case -2147483648:
                                            switch (v1 + (v36 + v1)) {
                                            case "foo bar":
                                                fmt.Println(v18, v28, v46)
                                            case "a":
                                                var v53 uint16 = 255
                                                _ = v53
                                            case "\n\t":
                                                v3 = v36
                                                if false {
                                                    v4 = 65535
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = ("hello" + v36)
                                                }
                                            default:
                                                v28 = v28
                                            }
                                        }
                                    case 610:
                                        var v54 int16 = 32767
                                        _ = v54
                                    }
                                    _ = v47
                                } else {
                                    if false {
                                        var v55 int32 = v2
                                        _ = len(v6)
                                        _ = v55
                                    }
                                }
                                _ = v46
                            }
                            v36 = ((("foo bar" + v3) + v3) + "foo bar")
                        }
                        _ = v37
                    } else {
                        fmt.Println(v18, v18, v28)
                        v56 := GenS1{F0: v36}
                        var v57 int32 = v28
                        fmt.Println(fi27)
                        var v58 int16 = (^(v18))
                        switch fi35 {
                        case 32768:
                            if (v2 <= (v28 | v2)) {
                                var v59 uint64 = (v5 - 18446744073709551615)
                                _ = v59
                            } else {
                                var v60 string = v3
                                var v61 string = v1
                                _ = v60
                                _ = v61
                            }
                        case -9223372036854775808:
                            if (!false) {
                                if false {
                                    var v62 int8 = 11
                                    switch (fi35 & fi27) {
                                    case 42:
                                        var v63 uint32 = 363
                                        _ = v63
                                    case -506:
                                        v2 += v57
                                    case -1:
                                        if (284 < uint(v5)) {
                                            if false {
                                                switch -1 {
                                                case -2147483648:
                                                    v36 = v3
                                                }
                                                var v64 uint16 = 65535
                                                switch (v3 + "hello") {
                                                case "hello":
                                                    if false {
                                                        if true {
                                                            var v65 bool = true
                                                            if v65 {
                                                                var v66 int32 = v57
                                                                _ = v66
                                                            }
                                                            fmt.Println(fi35, v3, v57)
                                                            _ = v65
                                                        } else {
                                                            v67 := GenS6{F0: 65535, F1: (v4 * v4)}
                                                            if true {
                                                                v6 = append(v6, v1)
                                                                v68 := GenS4{F0: 42, F1: 160}
                                                                var v69 int = (-(fi35))
                                                                fmt.Println(v28)
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = v67
                                                        }
                                                        var v70 uint8 = 14
                                                        v18 -= -476
                                                        v36 = v3
                                                        _ = v70
                                                    }
                                                case "a":
                                                    var v71 uint64 = 0
                                                    if (108 > 156) {
                                                        if true {
                                                            var v72 int32 = 2147483647
                                                            var v73 int8 = v62
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        var v74 int64 = 323
                                                        switch (-(fi35)) {
                                                        case -1:
                                                            v64 -= v64
                                                        default:
                                                            switch "" {
                                                            case "a":
                                                                fmt.Println(fi27, v57, v58)
                                                            }
                                                        }
                                                        _ = v74
                                                    } else {
                                                        v4 += (v4 | (v4 & (v4 / ((832) | 1))))
                                                        if (!false) {
                                                            var v75 int32 = v28
                                                            _ = v75
                                                        } else {
                                                            var v76 bool = true
                                                            var v77 int32 = -1
                                                            v78 := GenS1{F0: (v36 + v3)}
                                                            _ = v76
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        fmt.Println(v28, v18)
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v79 uint32 = (v4 ^ (v4 % ((v4) | 1)))
                                                        _ = v79
                                                    case "foo bar":
                                                        v80 := GenS7{F0: 0}
                                                        _ = v80
                                                    default:
                                                        v81 := GenS1{F0: v1}
                                                        v6 = append(v6, "\n\t")
                                                        _ = v81
                                                    }
                                                    _ = v71
                                                }
                                                if (v57 >= 948) {
                                                    var v82 int16 = v58
                                                    _ = v82
                                                } else {
                                                    if true {
                                                        var v83 int64 = 0
                                                        var v84 uint64 = (v5 >> ((832) & 63))
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        var v85 float32 = 1e10
                                                        fmt.Println(v58, v2)
                                                        switch (fi35 / (((fi27 >> ((18446744073709551615) & 63))) | 1)) {
                                                        case 9223372036854775807:
                                                            if true {
                                                                v2 = 2147483647
                                                            }
                                                        case 128:
                                                            var v86 uint64 = 874
                                                            _ = v86
                                                        case -164:
                                                            var v87 float64 = 1.5
                                                            _ = v87
                                                        }
                                                        _ = v85
                                                    }
                                                    switch v3 {
                                                    case "\n\t":
                                                        v88 := GenS7{F0: (v62 >> ((190) & 63))}
                                                        _ = v88
                                                    case "":
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case "a":
                                                        v5 -= (v5 << ((444) & 63))
                                                    }
                                                }
                                                _ = v64
                                            } else {
                                                switch 128 {
                                                case 549:
                                                    fmt.Println(v4)
                                                case -1:
                                                    v89 := GenS8{F0: v18, F1: v57}
                                                    _ = v89
                                                case 42:
                                                    if len(v6) > 0 {
                                                        v6[0] = ((((v36 + (v1 + v1)) + ("foo bar" + v1)) + (v36 + ((v36 + "a") + v3))) + v36)
                                                    }
                                                }
                                                v90 := GenS3{F0: false, F1: 86, F2: v57}
                                                _ = v90
                                            }
                                            v6 = append(v6, v36)
                                        } else {
                                            switch fi27 {
                                            case -891:
                                                v62 += 41
                                            default:
                                                var v91 int16 = (-(v18))
                                                _ = v91
                                            }
                                        }
                                    }
                                    switch (fi27 >> ((18446744073709551615) & 63)) {
                                    case 127:
                                        if true {
                                            v62 += (v62 - 42)
                                            v57 -= (^(v2))
                                            if (v58 != -32768) {
                                                switch v3 {
                                                case "hello":
                                                    fmt.Println(v18, v2, v4)
                                                }
                                            }
                                        } else {
                                            v4 -= (v4 << ((2) & 63))
                                        }
                                    case -963:
                                        v28 = -280
                                    case 153:
                                        var v92 bool = false
                                        _ = v92
                                    default:
                                        fmt.Println(v57, v57)
                                    }
                                    _ = v62
                                } else {
                                    v93 := GenS8{F0: -32768, F1: 42}
                                    fmt.Println(v57, v5, fi35)
                                    fmt.Println(v18)
                                    if false {
                                        switch v1 {
                                        case "hello":
                                            var v94 uint8 = 42
                                            _ = v94
                                        case "a":
                                            v58 -= v58
                                        case "\n\t":
                                            if (true && false) {
                                                if false {
                                                    if false {
                                                        fmt.Println(fi27, v1)
                                                        var v95 int = (fi27 / ((int(v5)) | 1))
                                                        _ = v95
                                                    } else {
                                                        v28 += 1
                                                    }
                                                }
                                                v6 = append(v6, (v1 + ((v36 + v3) + (v36 + v1))))
                                                v57 = 350
                                                switch fi27 {
                                                case 32767:
                                                    if false {
                                                        v28 += (v57 ^ (v2 - (v28 * (v2 >> ((4294967295) & 63)))))
                                                    }
                                                case -743:
                                                    var v96 int32 = v2
                                                    _ = v96
                                                case 2147483647:
                                                    v3 = "a"
                                                default:
                                                    v4 += v4
                                                    v6 = append(v6, (v1 + (v3 + v1)))
                                                }
                                            } else {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            v6 = append(v6, "a")
                                        }
                                        var v97 uint32 = 2
                                        v18 = 867
                                        _ = v97
                                    }
                                    v98 := GenS5{F0: v58, F1: ("foo bar" + (v36 + v1))}
                                    _ = v93
                                    _ = v98
                                }
                            } else {
                                v99 := GenS6{F0: 65535, F1: v4}
                                var v100 int8 = int8(v28)
                                _ = v100
                                _ = v99
                            }
                        case 127:
                            switch fi35 {
                            case 1:
                                v36 = v3
                                fmt.Println(v5, v18, fi27)
                            case 127:
                                var v101 int16 = (v58 + -1)
                                _ = v101
                            default:
                                v28 = int32(v5)
                            }
                        }
                        v4 += v4
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    fmt.Println(v1, v18, v18)
                }
                _ = fi35
                _ = v36
            }
            _ = fi27
            _ = v28
        }
        v3 = v3
    }
    switch v3 {
    case "foo bar":
        var v102 int = 875
        _ = v102
    case "":
        var v103 float64 = 0.0
        _ = v103
    case "\n\t":
        var v104 int8 = int8(v5)
        _ = v104
    }
    _ = len(v6)
    var v105 int32 = v2
    for fi106 := 0; fi106 < 18; fi106++ {
        v18 = (v18 * int16(fi106))
        v107 := GenS4{F0: 190, F1: 127}
        if len(v6) > 0 {
            v6[0] = ("a" + v3)
        }
        _ = fi106
        _ = v107
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v105)
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v105
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
