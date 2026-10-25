package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 bool; F1 uint64; F2 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 uint; F2 string }
type GenS9 struct { F0 uint16; F1 int32; F2 uint }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 int32; F1 string }
type GenS12 struct { F0 int64 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", "\n\t", "\n\t", ""}
    var v4 float32 = ((0.0 * 0.1) + (0.1 * 1.5))
    v5 := []string{"hello", ("hello" + ""), ("hello" + "\n\t")}
    v6 := make(map[int]uint32)
    switch -954 {
    case 93:
        var v7 uint64 = 4294967295
        _ = v7
    case 2147483647:
        if false {
            if ("\n\t" <= "hello") {
                v4 = 1.5
                for fi8 := 0; fi8 < 19; fi8++ {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        fmt.Println(v1, fi9)
                        if (42 == -1) {
                            v10 := GenS1{F0: 42, F1: (false && (!false)), F2: 65535}
                            var v11 int16 = 768
                            if (int64(fi9) > 9223372036854775807) {
                                delete(v2, (((("foo bar" + "foo bar") + ("a" + "hello")) + "hello") + "hello"))
                            } else {
                                _ = len(v6)
                                v12 := GenS1{F0: v11, F1: ("a" <= "a"), F2: 32767}
                                v1 += 1
                                _ = v12
                            }
                            _ = v10
                            _ = v11
                        } else {
                            var v13 uint = 42
                            v3 = append(v3, "foo bar")
                            if ((-1 <= 0) || (1 < -128)) {
                                var v14 float64 = -2.25
                                _ = v14
                            } else {
                                if (((v1 << ((v13) & 63)) != 42) || true) {
                                    if (0 > 2) {
                                        if (false || false) {
                                            var v15 uint64 = 308
                                            var v16 uint16 = 32767
                                            var v17 uint32 = uint32(v16)
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                        }
                                    }
                                    var v18 int = 0
                                    _ = v18
                                }
                                fmt.Println(fi9, v4)
                                if (495 > 18446744073709551615) {
                                    v13 -= v13
                                    v4 = (v4 + (v4 - v4))
                                }
                            }
                            _ = v13
                        }
                        v19 := GenS1{F0: 1000, F1: true, F2: 2}
                        _ = fi9
                        _ = v19
                    }
                    _ = fi8
                }
                switch "a" {
                case "\n\t":
                    v20 := GenS2{F0: ("" + ("foo bar" + ("a" + "\n\t"))), F1: 32767, F2: 32767}
                    _ = v20
                case "hello":
                    v4 = (-2.25 * float32(1.5))
                case "":
                    v1 = v1
                default:
                    var v21 int64 = -1
                    _ = v21
                }
            } else {
                v4 = 1.5
                var v22 uint8 = 193
                _ = v22
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                v4 -= 3.14159
                switch fi23 {
                case -53:
                    _ = len(v5)
                case 128:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v4 += v4
                        v2["\n\t"] = 432
                        v25 := GenS2{F0: "hello", F1: 24, F2: 42}
                        _ = fi24
                        _ = v25
                    }
                case 42:
                    v26 := v6[fi23]
                    _ = v26
                }
                _ = fi23
            }
        } else {
            var v27 int8 = -52
            var v28 int32 = 2147483647
            _ = v27
            _ = v28
        }
    case 4294967295:
        v1 = (-(v1))
    default:
        var v29 uint64 = 0
        _ = v29
    }
    v30 := v2["foo bar"]
    v31 := GenS2{F0: ("" + "hello"), F1: -32768, F2: 688}
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 int8 = 120
            _ = fi33
            _ = v34
        }
        v3 = append(v3, "a")
        for fi35 := 0; fi35 < 7; fi35++ {
            if false {
                if true {
                    v4 = (((float32((1.5 / 3.14159)) + 1.5) / 1.0) - 3.14159)
                    v30 -= 65535
                    v36 := GenS1{F0: -32768, F1: (v4 > (((v4 * 0.0) - (((v4 - v4) / ((v4 - float32((1.5 - 1.5))) / 1e10)) - 3.14159)) * 1e10)), F2: 0}
                    fmt.Println(fi32, fi32)
                    if false {
                        var v37 int32 = v1
                        var v38 uint8 = 1
                        _ = len(v3)
                        _ = v37
                        _ = v38
                    } else {
                        _ = len(v6)
                        var v39 int32 = (v1 * (v1 | (-(v1))))
                        _ = v39
                    }
                    v30 = 0
                    var v40 int = fi35
                    switch fi32 {
                    case 127:
                        switch (("foo bar" + "\n\t") + (("foo bar" + "\n\t") + "\n\t")) {
                        case "a":
                            switch ("foo bar" + (("foo bar" + "a") + ("hello" + ("foo bar" + "foo bar")))) {
                            case "":
                                var v41 uint64 = 0
                                _ = v41
                            case "a":
                                if false {
                                    v5 = append(v5, "a")
                                }
                            }
                        default:
                            var v42 uint8 = 135
                            _ = v42
                        }
                    default:
                        v43 := GenS3{F0: ((1.5 - 3.14159) >= 0.0), F1: 32767}
                        _ = v43
                    }
                    _ = v36
                    _ = v40
                } else {
                    switch (-(fi35)) {
                    case -723:
                        v30 = v30
                    case 127:
                        var v44 float64 = (float64(v4) * 3.14159)
                        _ = v44
                    }
                }
                v30 = v30
                if (v1 == v1) {
                    _ = len(v5)
                    fmt.Println(fi32, v4)
                }
                var v45 uint16 = 968
                _ = v45
            }
            var v46 uint64 = 125
            _ = fi35
            _ = v46
        }
        _ = fi32
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        if false {
            v48 := GenS4{F0: 65535, F1: 0}
            var v49 uint64 = 42
            _ = v48
            _ = v49
        } else {
            for fi50 := 0; fi50 < 18; fi50++ {
                fmt.Println(fi50, v30)
                if false {
                    v51 := GenS5{F0: ((v30 >> ((1) & 63)) == v30), F1: 4294967295, F2: 9223372036854775807}
                    if ((0.0 / 1.0) < (3.14159 / ((0.1 - 0.1) - (-2.25 + 3.14159)))) {
                        switch fi50 {
                        case 32:
                            var v52 int8 = -66
                            _ = v52
                        }
                        fmt.Println(v1)
                        v6[fi50] = v30
                        if (!(true || ((3.14159 - -2.25) > v4))) {
                            v1 = (^(v1))
                            v4 += 0.1
                        } else {
                            switch "\n\t" {
                            case "":
                                if false {
                                    var v53 uint32 = (v30 * v30)
                                    _ = v53
                                } else {
                                    if (v4 == (((0.0 - 0.0) * v4) * 0.0)) {
                                        v54 := GenS6{F0: -1}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v55 := GenS7{F0: int32(v30)}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v56 float32 = float32(1.5)
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 bool = (false || true)
                                    var v58 uint = 18446744073709551615
                                    _ = v57
                                    _ = v58
                                }
                            case "foo bar":
                                v59 := GenS5{F0: ((fi47 >= (fi50 + fi47)) || (475 >= 362)), F1: 892, F2: 0}
                                _ = v59
                            default:
                                if false {
                                    fmt.Println(fi47, v30, fi50)
                                    var v60 uint64 = 18446744073709551615
                                    _ = v60
                                }
                            }
                            fmt.Println(v30)
                        }
                        if (42 == -1) {
                            var v61 int8 = -13
                            _ = v61
                        }
                        v62 := v6[(^(fi50))]
                        _ = v62
                    }
                    var v63 uint = 0
                    _ = v51
                    _ = v63
                } else {
                    var v64 int = fi50
                    _ = v64
                }
                v65 := GenS1{F0: 0, F1: false, F2: 307}
                if false {
                    var v66 uint32 = v30
                    if false {
                        v67 := GenS5{F0: (uint8(v66) == 182), F1: 42, F2: -728}
                        _ = v67
                    } else {
                        var v68 int = (fi47 % (((fi47 - fi50)) | 1))
                        var v69 uint = 743
                        _ = v68
                        _ = v69
                    }
                    v1 -= v1
                    var v70 uint32 = v66
                    switch (fi50 & fi50) {
                    case 128:
                        _ = len(v2)
                    case 4294967295:
                        _ = len(v6)
                    }
                    fmt.Println(v1)
                    var v71 int16 = int16(v30)
                    if ("\n\t" == (("\n\t" + "hello") + "hello")) {
                        var v72 string = ""
                        _ = v72
                    } else {
                        if false {
                            fmt.Println(v70, fi47)
                            switch 4294967295 {
                            case 4294967295:
                                fmt.Println(fi47, v30, fi50)
                            }
                            var v73 uint64 = 42
                            _ = v73
                        } else {
                            if false {
                                switch fi50 {
                                case 32767:
                                    var v74 float64 = 3.14159
                                    _ = v74
                                }
                                v75 := GenS8{F0: 159, F1: 715, F2: "foo bar"}
                                v76 := GenS9{F0: 0, F1: (v1 ^ 166), F2: 2}
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "hello":
                                    fmt.Println(fi47, v30, v70)
                                    fmt.Println(v66)
                                }
                                _ = v75
                                _ = v76
                            }
                            var v77 int64 = 42
                            _ = v77
                        }
                        switch fi50 {
                        case 32768:
                            if (3.14159 >= 1e10) {
                                if (2 > 255) {
                                    v78 := GenS7{F0: v1}
                                    _ = v78
                                } else {
                                    fmt.Println(fi50, fi50, fi47)
                                }
                            } else {
                                v4 -= float32(1.5)
                            }
                        case 2147483647:
                            v79 := GenS10{F0: "", F1: "hello"}
                            fmt.Println(v71, v70, v70)
                            switch 32767 {
                            case 127:
                                fmt.Println(v71, v71, v66)
                            }
                            _ = v79
                        }
                        v66 += (v70 ^ (v70 & (v66 ^ (v30 + v66))))
                    }
                    var v80 int8 = -128
                    v1 -= v1
                    _ = v66
                    _ = v70
                    _ = v71
                    _ = v80
                }
                _ = fi50
                _ = v65
            }
            var v81 int32 = 42
            _ = v81
        }
        _ = fi47
    }
    for fi82 := 0; fi82 < 8; fi82++ {
        v30 += (v30 & 42)
        var v83 uint = 4294967295
        if false {
            var v84 float64 = (-2.25 - 0.0)
            v83 = 707
            _ = v84
        }
        _ = fi82
        _ = v83
    }
    if (0 != -1) {
        v30 -= v30
        switch "foo bar" {
        case "foo bar":
            v30 -= (v30 + 2147483647)
        case "hello":
            for fi85 := 0; fi85 < 19; fi85++ {
                for fi86 := 0; fi86 < 1; fi86++ {
                    v87 := GenS11{F0: v1, F1: ("" + ((("" + "\n\t") + "") + ("hello" + "a")))}
                    _ = fi86
                    _ = v87
                }
                var v88 int32 = (v1 >> ((42) & 63))
                for fi89 := 0; fi89 < 12; fi89++ {
                    if false {
                        var v90 uint = 18446744073709551615
                        v91 := GenS12{F0: 0}
                        v1 = 303
                        if false {
                            var v92 int = (fi85 - fi89)
                            _ = v92
                        }
                        _ = v90
                        _ = v91
                    }
                    var v93 int = fi85
                    _ = fi89
                    _ = v93
                }
                _ = fi85
                _ = v88
            }
        }
    }
    _ = len(v3)
    var v94 int = 42
    v95 := GenS11{F0: v1, F1: "\n\t"}
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v96 float64 = float64(v4)
    v1 += (v1 + v1)
    switch "foo bar" {
    case "":
        if (((42 <= 63) && (true && (32767 == 32767))) || true) {
            v97 := v6[v94]
            for fi98 := 0; fi98 < 12; fi98++ {
                var v99 float32 = (v4 - v4)
                v100 := GenS5{F0: (2147483647 < (v1 >> ((0) & 63))), F1: 0, F2: 9223372036854775807}
                v6[-1] = (v30 + v97)
                _ = fi98
                _ = v100
                _ = v99
            }
            switch v94 {
            case -1:
                if (!((false || false) && ((v30 + (v30 << ((255) & 63))) == 9))) {
                    switch v94 {
                    case 128:
                        fmt.Println(v96)
                    }
                }
            default:
                v97 = v97
            }
            _ = v97
        }
    case "a":
        v101 := GenS5{F0: (v1 != v1), F1: 2, F2: -34}
        _ = v101
    case "foo bar":
        if (!(!(106 > 50))) {
            var v102 uint8 = 199
            _ = v102
        }
    default:
        var v103 uint16 = 42
        _ = v103
    }
    fmt.Println(v94, v4)
    for fi104 := 0; fi104 < 16; fi104++ {
        v3 = append(v3, "\n\t")
        var v105 uint32 = (v30 - 556)
        var v106 float32 = v4
        _ = fi104
        _ = v105
        _ = v106
    }
    var v107 uint = 65535
    v108 := GenS5{F0: false, F1: 864, F2: int64(v94)}
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v94
    _ = v95
    _ = v96
}
