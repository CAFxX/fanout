package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint32; F2 int64 }
type GenS8 struct { F0 int; F1 string; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint32; F2 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int16; F1 int16; F2 uint16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS17 struct { F0 int8; F1 uint16; F2 string }
type GenS18 struct { F0 uint16; F1 uint16 }
type GenS19 struct { F0 int16; F1 bool; F2 int16 }
type GenS20 struct { F0 bool; F1 int64 }
type GenS21 struct { F0 uint16; F1 int16 }
type GenS22 struct { F0 uint64; F1 uint64 }
type GenS23 struct { F0 int32; F1 int16; F2 int16 }
type GenS24 struct { F0 uint32; F1 bool; F2 int32 }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int8 }
type GenS28 struct { F0 int64; F1 uint32; F2 int32 }
type GenS29 struct { F0 uint16; F1 uint8 }
type GenS30 struct { F0 uint32 }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 int8; F1 int64; F2 string }
type GenS33 struct { F0 uint64 }
type GenS34 struct { F0 int16; F1 int8; F2 uint }
type GenS35 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS36 struct { F0 uint32; F1 uint8; F2 uint }
type GenS37 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint, p1 int16) uint16 {
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            var v8 string = "a"
            v9 := GenS1{F0: 42, F1: fi7, F2: false}
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 float32 = 1e10
                var v12 bool = true
                var v13 int = fi7
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = v8
            _ = v9
        } else {
            var v14 float32 = 3.14159
            _ = v14
        }
        switch "hello" {
        case "hello":
            v15 := []int{(^(fi7)), (fi7 & (fi7 / ((fi7) | 1)))}
            _ = v15
        case "\n\t":
            var v16 float32 = 1.5
            _ = v16
        case "foo bar":
            v17 := GenS2{F0: 42, F1: (-(p1))}
            _ = v17
        default:
            for fi18 := 0; fi18 < 14; fi18++ {
                p0 = (p0 << (((^(p0))) & 63))
                if ((-2.25 <= (3.14159 * 0.0)) && ((!true) || (97 > 1))) {
                    if (2 <= 153) {
                        v19 := GenS2{F0: 0, F1: p1}
                        switch (("\n\t" + "foo bar") + "\n\t") {
                        case "\n\t":
                            p0 = 65535
                        case "a":
                            v20 := GenS1{F0: 2147483647, F1: fi7, F2: false}
                            _ = v20
                        case "hello":
                            fmt.Println(p0, p1)
                        }
                        _ = v19
                    }
                }
                fmt.Println(p0)
                var v21 float64 = (1.5 / 0.1)
                _ = fi18
                _ = v21
            }
        }
        _ = fi7
    }
    if ("a" == "") {
        v22 := []int{-30, 2147483647, 889}
        p1 = -263
        var v23 uint8 = 127
        _ = v22
        _ = v23
    } else {
        var v24 int = -470
        p0 -= p0
        _ = v24
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = ((-611 != 9223372036854775807) || (0.0 > 1.5))
    v3 := []string{}
    var v4 bool = true
    var v5 string = ""
    var v6 int64 = 9223372036854775807
    if true {
        var v25 int16 = -103
        var v26 int8 = 1
        if (-2.25 == -2.25) {
            v27 := GenS3{F0: 2}
            for fi28 := 0; fi28 < 9; fi28++ {
                switch "hello" {
                case "hello":
                    if (fi28 == (fi28 * (^(fi28)))) {
                        for fi29 := 0; fi29 < 16; fi29++ {
                            var v30 int = fi29
                            if v4 {
                                if v2 {
                                    var v31 uint64 = 9
                                    _ = v31
                                }
                                v26 = (v26 + v26)
                                v4 = v2
                            } else {
                                var v32 int = fi29
                                if v4 {
                                    v2 = (!(true || v2))
                                } else {
                                    v3 = append(v3, "foo bar")
                                }
                                var v33 uint = 837
                                _ = v32
                                _ = v33
                            }
                            switch -868 {
                            case -371:
                                var v34 int32 = 2147483647
                                _ = v34
                            case 32768:
                                v26 += v26
                            default:
                                v3 = append(v3, "\n\t")
                                if (127 <= 0) {
                                    if v2 {
                                        var v35 float32 = float32(3.14159)
                                        if false {
                                            if (true && true) {
                                                var v36 float32 = float32((((0.0 - 3.14159) + -2.25) - 1.5))
                                                _ = v36
                                            } else {
                                                switch v30 {
                                                case 42:
                                                    var v37 bool = ((-(v6)) == int64(v25))
                                                    _ = v37
                                                case -1:
                                                    var v38 uint8 = 2
                                                    _ = v38
                                                case 128:
                                                    v30 += 854
                                                default:
                                                    if ((v25 >> ((165) & 63)) < 1) {
                                                        if len(v3) > 0 {
                                                            v3[0] = v5
                                                        }
                                                        var v39 bool = v4
                                                        delete(v1, (fi28 | (v30 - 42)))
                                                        _ = v39
                                                    }
                                                }
                                            }
                                            v25 = v25
                                            var v40 int16 = (v25 - v25)
                                            v41 := GenS4{F0: 408, F1: 247, F2: 2}
                                            var v42 int = v30
                                            v4 = (0 == 2)
                                            switch (v5 + (v5 + ("hello" + v5))) {
                                            case "hello":
                                                if (v40 > (v25 ^ (v40 * -644))) {
                                                    v43 := v1[(fi29 - v42)]
                                                    var v44 uint16 = 42
                                                    if v2 {
                                                        v45 := GenS5{F0: 892}
                                                        _ = len(v1)
                                                        _ = v45
                                                    } else {
                                                        var v46 bool = true
                                                        v47 := GenS5{F0: 648}
                                                        _ = len(v3)
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    _ = len(v1)
                                                    var v48 int64 = v6
                                                    if (v2 || ((100 != (^(v26))) && v2)) {
                                                        switch (fi29 >> ((4294967295) & 63)) {
                                                        case -9223372036854775808:
                                                            if (42 != 128) {
                                                                if v4 {
                                                                    fmt.Println(v2, v6, v48)
                                                                    var v49 uint64 = 18446744073709551615
                                                                    _ = v49
                                                                } else {
                                                                    var v50 float32 = 3.14159
                                                                    if (-2147483648 < 1) {
                                                                        v35 = v35
                                                                    }
                                                                    if false {
                                                                        v2 = v4
                                                                    }
                                                                    _ = v50
                                                                }
                                                                v42 = v30
                                                            } else {
                                                                v48 += 0
                                                                var v51 int = int(v26)
                                                                v42 -= (-(fi28))
                                                                _ = v51
                                                            }
                                                            v30 = fi29
                                                            var v52 string = v5
                                                            _ = v52
                                                        case -306:
                                                            fmt.Println(fi29, v30, v6)
                                                        case 0:
                                                            fmt.Println(v2, v2)
                                                        }
                                                        if false {
                                                            if (!(!(2 < 155))) {
                                                                v53 := GenS6{F0: v25}
                                                                v54 := v1[v42]
                                                                _ = v53
                                                                _ = v54
                                                            } else {
                                                                var v55 int16 = (v25 | int16(v26))
                                                                var v56 float32 = 1.5
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                            if true {
                                                                delete(v1, (-(fi29)))
                                                            } else {
                                                                if v4 {
                                                                    if ((v26 << ((2) & 63)) <= (v26 / (((v26 ^ (^(v26)))) | 1))) {
                                                                        v5 = v5
                                                                        var v57 uint32 = 2147483647
                                                                        _ = len(v1)
                                                                        _ = v57
                                                                    }
                                                                    if true {
                                                                        v35 = v35
                                                                    } else {
                                                                        var v58 uint = 2
                                                                        _ = v58
                                                                    }
                                                                    v42 += fi29
                                                                } else {
                                                                    if (!(false && (!v2))) {
                                                                        fmt.Println(v5, v6, v30)
                                                                        switch (v42 * 127) {
                                                                        case -9223372036854775808:
                                                                            var v59 uint8 = 0
                                                                            _ = v59
                                                                        case -751:
                                                                            v60 := GenS7{F0: v5, F1: 65535, F2: (v48 ^ v6)}
                                                                            var v61 uint = 42
                                                                            v61 = 0
                                                                            v35 -= v35
                                                                            _ = v60
                                                                            _ = v61
                                                                        case 0:
                                                                            v30 = v30
                                                                        default:
                                                                            if v4 {
                                                                                if v2 {
                                                                                    v30 = fi29
                                                                                }
                                                                                v62 := GenS8{F0: -998, F1: v5, F2: 42}
                                                                                _ = v62
                                                                            } else {
                                                                                if (v48 >= (v6 | 9223372036854775807)) {
                                                                                    if ((1e10 + v35) < v35) {
                                                                                        fmt.Println(v42)
                                                                                        var v63 bool = (-909 >= 2147483647)
                                                                                        v64 := GenS9{F0: 0, F1: 2, F2: (65535 > 4294967295)}
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    }
                                                                                } else {
                                                                                    var v65 int64 = v6
                                                                                    var v66 int64 = v48
                                                                                    var v67 float64 = 1.5
                                                                                    v2 = ((42 != 1) && (0 == 42))
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                                var v68 int = (-(v30))
                                                                                v4 = v4
                                                                                v2 = ((-1 == 1) && (1 != 2147483647))
                                                                                switch (-(v68)) {
                                                                                case -2147483648:
                                                                                    var v69 uint = 4294967295
                                                                                    _ = v69
                                                                                case 9223372036854775807:
                                                                                    if v2 {
                                                                                        fmt.Println(fi28, fi28)
                                                                                        v70 := GenS10{F0: 1}
                                                                                        _ = v70
                                                                                    }
                                                                                case -28:
                                                                                    switch (v5 + v5) {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v5, v2)
                                                                                    case "":
                                                                                        var v71 uint = 374
                                                                                        _ = v71
                                                                                    case "\n\t":
                                                                                        fmt.Println(v26, v5)
                                                                                    default:
                                                                                        var v72 int16 = 0
                                                                                        _ = v72
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        v73 := GenS3{F0: 71}
                                                                                        v6 += 9223372036854775807
                                                                                        _ = v73
                                                                                    }
                                                                                }
                                                                                v25 += (v40 | (v40 >> ((289) & 63)))
                                                                                v3 = append(v3, "")
                                                                                if false {
                                                                                    v74 := GenS11{F0: 42}
                                                                                    v2 = v2
                                                                                    fmt.Println(fi29, v6)
                                                                                    var v75 string = v5
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                } else {
                                                                                    fmt.Println(v26)
                                                                                    v48 = (v6 + (-(v6)))
                                                                                }
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v3 = append(v3, (v5 + ""))
                                                                        if v4 {
                                                                            fmt.Println(v26)
                                                                            fmt.Println(v40, v42)
                                                                        }
                                                                        switch -9223372036854775808 {
                                                                        case 970:
                                                                            v25 += v25
                                                                        case 920:
                                                                            switch "hello" {
                                                                            case "a":
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = "\n\t"
                                                                                }
                                                                            case "":
                                                                                var v76 uint32 = 65535
                                                                                _ = v76
                                                                            }
                                                                        case 128:
                                                                            switch v5 {
                                                                            case "foo bar":
                                                                                switch "hello" {
                                                                                case "":
                                                                                    var v77 uint32 = 2
                                                                                    _ = v77
                                                                                case "foo bar":
                                                                                    var v78 int32 = -602
                                                                                    _ = v78
                                                                                }
                                                                                var v79 string = v5
                                                                                _ = v79
                                                                            case "":
                                                                                var v80 uint = 422
                                                                                _ = v80
                                                                            }
                                                                        default:
                                                                            v1[(v42 - (fi29 & fi29))] = "foo bar"
                                                                        }
                                                                        var v81 uint = 2
                                                                        _ = v81
                                                                    }
                                                                    v82 := GenS9{F0: 0, F1: 2, F2: v4}
                                                                    var v83 uint = 42
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                                if false {
                                                                    v84 := GenS12{F0: (v48 + -9223372036854775808), F1: 1000}
                                                                    _ = v84
                                                                } else {
                                                                    if v2 {
                                                                        fmt.Println(v35, v6, v6)
                                                                        if (-222 == -873) {
                                                                            switch (v5 + ((("foo bar" + v5) + (v5 + (v5 + v5))) + v5)) {
                                                                            case "":
                                                                                v85 := GenS13{F0: (^(v25))}
                                                                                v30 -= (v42 >> ((65535) & 63))
                                                                                _ = v85
                                                                            case "hello":
                                                                                var v86 uint64 = 4294967295
                                                                                _ = v86
                                                                            }
                                                                            var v87 int32 = -610
                                                                            _ = v87
                                                                        } else {
                                                                            v5 = v5
                                                                        }
                                                                        _ = len(v3)
                                                                        var v88 uint16 = 2
                                                                        _ = v88
                                                                    } else {
                                                                        var v89 int8 = v26
                                                                        if v4 {
                                                                            var v90 int16 = 1000
                                                                            if ((1 <= 250) || ((fi29 / (((fi29 + v42)) | 1)) >= (-(fi29)))) {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    var v91 uint64 = 421
                                                                                    _ = v91
                                                                                case "":
                                                                                    v48 += 1
                                                                                case "\n\t":
                                                                                    if v4 {
                                                                                        v42 += (fi29 + fi29)
                                                                                        if false {
                                                                                            var v92 float64 = 0.1
                                                                                            _ = v92
                                                                                        } else {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            v93 := GenS3{F0: 207}
                                                                                            v42 = 86
                                                                                            _ = v93
                                                                                        }
                                                                                        switch -1 {
                                                                                        case 4294967295:
                                                                                            if (false && true) {
                                                                                                if v4 {
                                                                                                    var v94 int32 = 0
                                                                                                    _ = v94
                                                                                                }
                                                                                                if v4 {
                                                                                                    v1[v42] = v5
                                                                                                } else {
                                                                                                    if (false || true) {
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = (v5 + (((v5 + v5) + (v5 + ((v5 + v5) + "\n\t"))) + v5))
                                                                                                        }
                                                                                                        if (float64(3.14159) >= (3.14159 / (0.0 * float64(((1.5 / v35) - float32((float64((v35 - float32(-2.25))) + 0.0))))))) {
                                                                                                            v5 = (v5 + (v5 + "hello"))
                                                                                                        } else {
                                                                                                            switch v5 {
                                                                                                            case "a":
                                                                                                                v25 = (v40 * (v40 << ((0) & 63)))
                                                                                                            default:
                                                                                                                v95 := GenS14{F0: (-(v90)), F1: v40, F2: 32767}
                                                                                                                switch (fi28 - 9223372036854775807) {
                                                                                                                case 1:
                                                                                                                    if false {
                                                                                                                        v89 += v26
                                                                                                                        v96 := GenS1{F0: 4294967295, F1: 32767, F2: (v4 && (!true))}
                                                                                                                        _ = v96
                                                                                                                    } else {
                                                                                                                        v97 := GenS15{F0: (v89 >> ((42) & 63))}
                                                                                                                        v5 = v5
                                                                                                                        if v4 {
                                                                                                                            fmt.Println(v89, v2, v30)
                                                                                                                            var v98 int16 = 239
                                                                                                                            if len(v3) >= 2 {
                                                                                                                                v3 = v3[1:len(v3)]
                                                                                                                            }
                                                                                                                            _ = v98
                                                                                                                        }
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                case -1:
                                                                                                                    var v99 uint16 = 65535
                                                                                                                    _ = v99
                                                                                                                default:
                                                                                                                    v5 = "\n\t"
                                                                                                                }
                                                                                                                fmt.Println(v26)
                                                                                                                _ = v95
                                                                                                            }
                                                                                                            v35 = 0.1
                                                                                                            switch 423 {
                                                                                                            case 247:
                                                                                                                switch v5 {
                                                                                                                case "foo bar":
                                                                                                                    var v100 float32 = (0.0 - 0.1)
                                                                                                                    _ = v100
                                                                                                                default:
                                                                                                                    if v2 {
                                                                                                                        v4 = (!(!(v90 > (v90 << ((282) & 63)))))
                                                                                                                    }
                                                                                                                }
                                                                                                            case 2147483647:
                                                                                                                if true {
                                                                                                                    v48 -= v6
                                                                                                                    v3 = append(v3, (v5 + v5))
                                                                                                                } else {
                                                                                                                    v1[(fi29 % ((fi28) | 1))] = "foo bar"
                                                                                                                    if false {
                                                                                                                        v40 += (-(v25))
                                                                                                                    } else {
                                                                                                                        fmt.Println(v4, fi28, v26)
                                                                                                                        v101 := GenS2{F0: 877, F1: v40}
                                                                                                                        _ = v101
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v102 := GenS16{F0: 42, F1: 508, F2: 824}
                                                                                                                _ = v102
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v30 = -644
                                                                                                }
                                                                                            }
                                                                                        case 52:
                                                                                            v103 := GenS2{F0: 42, F1: -1}
                                                                                            if true {
                                                                                                v104 := GenS12{F0: v48, F1: v25}
                                                                                                v105 := GenS8{F0: fi28, F1: v5, F2: int32(v48)}
                                                                                                v106 := GenS17{F0: v89, F1: 740, F2: "a"}
                                                                                                if (!v2) {
                                                                                                    v4 = (!(v35 != v35))
                                                                                                    if (2 != 276) {
                                                                                                        if v4 {
                                                                                                            var v107 uint16 = 2
                                                                                                            var v108 float32 = ((v35 - (v35 - 0.0)) / 1.0)
                                                                                                            _ = v107
                                                                                                            _ = v108
                                                                                                        } else {
                                                                                                            v6 += (v48 % ((v6) | 1))
                                                                                                            v42 = 219
                                                                                                            v1[(-(v42))] = ((v5 + v5) + "foo bar")
                                                                                                        }
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = "\n\t"
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    v109 := GenS15{F0: (-(v26))}
                                                                                                    switch 32768 {
                                                                                                    case 32767:
                                                                                                        if v4 {
                                                                                                            var v110 int64 = 129
                                                                                                            _ = v110
                                                                                                        } else {
                                                                                                            if v2 {
                                                                                                                if ((v42 / ((-128) | 1)) <= -354) {
                                                                                                                    var v111 float32 = v35
                                                                                                                    v112 := GenS16{F0: 255, F1: 980, F2: 391}
                                                                                                                    _ = v111
                                                                                                                    _ = v112
                                                                                                                } else {
                                                                                                                    var v113 int64 = 1
                                                                                                                    if false {
                                                                                                                        switch (("\n\t" + (v5 + v5)) + v5) {
                                                                                                                        case "hello":
                                                                                                                            var v114 uint64 = 1
                                                                                                                            _ = v114
                                                                                                                        case "":
                                                                                                                            switch v42 {
                                                                                                                            case -1:
                                                                                                                                v115 := GenS18{F0: 0, F1: genFn0(568, -32768)}
                                                                                                                                _ = v115
                                                                                                                            }
                                                                                                                            var v116 bool = false
                                                                                                                            var v117 uint64 = 2
                                                                                                                            fmt.Println(v48)
                                                                                                                            var v118 float64 = (1.5 / 1.0)
                                                                                                                            v26 = v26
                                                                                                                            switch ("" + (v5 + v5)) {
                                                                                                                            case "\n\t":
                                                                                                                                var v119 uint64 = (v117 + (^(v117)))
                                                                                                                                _ = v119
                                                                                                                            default:
                                                                                                                                if false {
                                                                                                                                    v6 = (v48 >> ((4294967295) & 63))
                                                                                                                                    v120 := GenS19{F0: (v40 % ((int16(v117)) | 1)), F1: v4, F2: v25}
                                                                                                                                    var v121 uint8 = 42
                                                                                                                                    _ = len(v3)
                                                                                                                                    _ = v120
                                                                                                                                    _ = v121
                                                                                                                                }
                                                                                                                            }
                                                                                                                            switch ((v5 + (v5 + v5)) + "a") {
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(v48, v40, v35)
                                                                                                                            }
                                                                                                                            var v122 int32 = 0
                                                                                                                            v89 = 66
                                                                                                                            var v123 int32 = (v122 + 0)
                                                                                                                            _ = v116
                                                                                                                            _ = v117
                                                                                                                            _ = v118
                                                                                                                            _ = v122
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                        if v2 {
                                                                                                                            if v4 {
                                                                                                                                fmt.Println(v48, fi28, v26)
                                                                                                                            } else {
                                                                                                                                fmt.Println(v6)
                                                                                                                                var v124 uint = 4294967295
                                                                                                                                v6 = int64(v124)
                                                                                                                                if ((!true) && (127 >= 127)) {
                                                                                                                                    var v125 float32 = ((float32(float64(((float32(-2.25) + (float32(float64(v35)) - 0.1)) + -2.25))) - ((-2.25 + v35) - 3.14159)) + (1e10 * v35))
                                                                                                                                    if (-1 <= int32(v89)) {
                                                                                                                                        if v2 {
                                                                                                                                            fmt.Println(v42, v25)
                                                                                                                                            switch 9223372036854775807 {
                                                                                                                                            case 42:
                                                                                                                                                var v126 float64 = ((0.1 * 1.5) / 1.0)
                                                                                                                                                _ = v126
                                                                                                                                            case 905:
                                                                                                                                                var v127 int8 = (v26 & (v89 * (v89 >> ((v124) & 63))))
                                                                                                                                                _ = v127
                                                                                                                                            default:
                                                                                                                                                var v128 string = (("\n\t" + v5) + (v5 + (v5 + v5)))
                                                                                                                                                _ = v128
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        if (!false) {
                                                                                                                                            fmt.Println(v113, v40)
                                                                                                                                            v129 := GenS20{F0: v2, F1: v6}
                                                                                                                                            _ = v129
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v125
                                                                                                                                } else {
                                                                                                                                    v130 := GenS21{F0: uint16(v124), F1: (v90 & (v90 + 1000))}
                                                                                                                                    _ = v130
                                                                                                                                }
                                                                                                                                v131 := GenS18{F0: 2, F1: 749}
                                                                                                                                _ = v124
                                                                                                                                _ = v131
                                                                                                                            }
                                                                                                                            v132 := GenS18{F0: 1, F1: 65535}
                                                                                                                            if len(v3) >= 2 {
                                                                                                                                v3 = v3[1:len(v3)]
                                                                                                                            }
                                                                                                                            _ = v132
                                                                                                                        } else {
                                                                                                                            v133 := GenS17{F0: (v89 + 117), F1: 42, F2: v5}
                                                                                                                            _ = v133
                                                                                                                        }
                                                                                                                        v134 := GenS8{F0: 9223372036854775807, F1: ("" + v5), F2: 42}
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                    _ = v113
                                                                                                                }
                                                                                                                var v135 uint16 = 32767
                                                                                                                v48 += v6
                                                                                                                _ = v135
                                                                                                            }
                                                                                                            if v2 {
                                                                                                                v136 := GenS22{F0: 273, F1: 208}
                                                                                                                if ((0 != 1) || (true || (!(!((!false) || false))))) {
                                                                                                                    var v137 float32 = v35
                                                                                                                    delete(v1, int(v40))
                                                                                                                    _ = v137
                                                                                                                } else {
                                                                                                                    if ((((0.1 - (1.5 - 0.0)) / float64(float32(1e10))) != 3.14159) || (1 == 0)) {
                                                                                                                        v138 := GenS23{F0: 2147483647, F1: v25, F2: (v40 - (-(v25)))}
                                                                                                                        v139 := GenS24{F0: 54, F1: true, F2: 2147483647}
                                                                                                                        v26 -= (^(v26))
                                                                                                                        _ = v138
                                                                                                                        _ = v139
                                                                                                                    }
                                                                                                                    if v4 {
                                                                                                                        v89 = v89
                                                                                                                    }
                                                                                                                    var v140 uint16 = 375
                                                                                                                    _ = v140
                                                                                                                }
                                                                                                                v3 = append(v3, (("\n\t" + "foo bar") + v5))
                                                                                                                _ = v136
                                                                                                            } else {
                                                                                                                v42 += fi28
                                                                                                                v3 = append(v3, v5)
                                                                                                            }
                                                                                                            if (682 > 2147483647) {
                                                                                                                var v141 uint32 = 42
                                                                                                                var v142 float64 = (0.1 / (((1e10 - -2.25) + (3.14159 * 0.0)) / (0.0 - float64(1e10))))
                                                                                                                _ = v141
                                                                                                                _ = v142
                                                                                                            }
                                                                                                        }
                                                                                                    case -2147483648:
                                                                                                        fmt.Println(v35)
                                                                                                    }
                                                                                                    var v143 uint = 65535
                                                                                                    _ = v109
                                                                                                    _ = v143
                                                                                                }
                                                                                                _ = v104
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            }
                                                                                            _ = v103
                                                                                        default:
                                                                                            v144 := GenS8{F0: 2147483647, F1: v5, F2: -453}
                                                                                            _ = len(v1)
                                                                                            _ = v144
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch (fi29 * v42) {
                                                                                case 921:
                                                                                    var v145 float32 = 0.0
                                                                                    _ = v145
                                                                                case 0:
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        v6 -= v48
                                                                                    case "\n\t":
                                                                                        delete(v1, (fi29 % ((v30) | 1)))
                                                                                    default:
                                                                                        v146 := v1[v30]
                                                                                        _ = v146
                                                                                    }
                                                                                default:
                                                                                    var v147 float64 = 1e10
                                                                                    _ = v147
                                                                                }
                                                                                v148 := GenS25{F0: int8(fi29)}
                                                                                _ = v148
                                                                            }
                                                                            _ = v90
                                                                        } else {
                                                                            v149 := GenS26{F0: (v26 * (v26 >> ((4294967295) & 63)))}
                                                                            if (2147483648 >= 604) {
                                                                                fmt.Println(fi28, v30, v89)
                                                                                v150 := v1[(fi28 * (v30 + (-(fi29))))]
                                                                                if ((-2.25 * v35) == ((0.1 + v35) / ((1e10 + 0.1) + (v35 / v35)))) {
                                                                                    v3 = append(v3, v5)
                                                                                    v1[fi29] = ((v5 + v150) + (v150 + "hello"))
                                                                                    v25 -= 1000
                                                                                }
                                                                                _ = v150
                                                                            } else {
                                                                                if ((4294967295 > 0) && ((v42 & (fi28 | (v30 >> ((0) & 63)))) <= 128)) {
                                                                                    var v151 int8 = 100
                                                                                    v35 -= ((0.0 / 1.0) * v35)
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = v5
                                                                                    }
                                                                                    _ = v151
                                                                                }
                                                                            }
                                                                            v152 := GenS5{F0: 0}
                                                                            _ = v149
                                                                            _ = v152
                                                                        }
                                                                        _ = v89
                                                                    }
                                                                    v1[(v42 ^ -2147483648)] = "foo bar"
                                                                    v40 = (v25 + int16(v30))
                                                                    if v2 {
                                                                        fmt.Println(v25, v42, v2)
                                                                        v6 = (^(v48))
                                                                        fmt.Println(fi28, fi28, v5)
                                                                    } else {
                                                                        var v153 float32 = 1e10
                                                                        if true {
                                                                            if v2 {
                                                                                v2 = v2
                                                                                v5 = ("hello" + (("foo bar" + "") + ("\n\t" + ("" + "foo bar"))))
                                                                                v1[fi28] = "a"
                                                                            }
                                                                            var v154 int64 = -9223372036854775808
                                                                            _ = len(v1)
                                                                            _ = v154
                                                                        } else {
                                                                            if v4 {
                                                                                if ((18446744073709551615 <= 4294967295) && (!false)) {
                                                                                    v30 = fi28
                                                                                } else {
                                                                                    v26 += (v26 - v26)
                                                                                    v155 := GenS27{F0: v26}
                                                                                    _ = len(v3)
                                                                                    _ = v155
                                                                                }
                                                                                if (!((v40 >= (v25 - (^(v40)))) || (v4 && false))) {
                                                                                    var v156 int64 = v6
                                                                                    _ = v156
                                                                                } else {
                                                                                    v30 = (^(v42))
                                                                                    if v4 {
                                                                                        v5 = (v5 + (v5 + (("foo bar" + v5) + "foo bar")))
                                                                                        switch 2147483647 {
                                                                                        case 127:
                                                                                            delete(v1, 0)
                                                                                        }
                                                                                        var v157 uint = 1
                                                                                        _ = v157
                                                                                    }
                                                                                }
                                                                                v2 = ((v25 + 0) > (v25 & 810))
                                                                            }
                                                                            if (-1 < 42) {
                                                                                if true {
                                                                                    switch (^(fi28)) {
                                                                                    case -9223372036854775808:
                                                                                        v2 = (!((v6 + (-(v48))) != v48))
                                                                                    }
                                                                                    v158 := GenS6{F0: (^(v40))}
                                                                                    _ = v158
                                                                                }
                                                                            } else {
                                                                                if v4 {
                                                                                    if v4 {
                                                                                        v40 += (v40 * v25)
                                                                                        fmt.Println(v35, v2)
                                                                                    } else {
                                                                                        v159 := GenS21{F0: genFn0(0, v40), F1: 0}
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (v5 + v5)
                                                                                        }
                                                                                        _ = v159
                                                                                    }
                                                                                    if (1 == 591) {
                                                                                        delete(v1, v42)
                                                                                    }
                                                                                }
                                                                            }
                                                                            v160 := GenS13{F0: v25}
                                                                            fmt.Println(v2)
                                                                            v161 := GenS2{F0: 18446744073709551615, F1: v25}
                                                                            _ = v160
                                                                            _ = v161
                                                                        }
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            var v162 uint32 = 65535
                                                                            _ = v162
                                                                        case "":
                                                                            switch (-(fi28)) {
                                                                            case -201:
                                                                                var v163 int32 = 1
                                                                                _ = v163
                                                                            default:
                                                                                v6 = v48
                                                                            }
                                                                        case "hello":
                                                                            var v164 float32 = (((1e10 - (v35 / float32(1.5))) / v35) - -2.25)
                                                                            if true {
                                                                                v165 := GenS3{F0: 0}
                                                                                _ = v165
                                                                            } else {
                                                                                var v166 string = v5
                                                                                if (!(!(((v2 || true) && v4) && v2))) {
                                                                                    v3 = append(v3, v166)
                                                                                }
                                                                                _ = v166
                                                                            }
                                                                            _ = v164
                                                                        default:
                                                                            v42 += -1
                                                                        }
                                                                        _ = v153
                                                                    }
                                                                    v1[(v42 / ((fi28) | 1))] = "a"
                                                                }
                                                            }
                                                            if true {
                                                                if (v4 || (4294967295 == 0)) {
                                                                    v48 += v48
                                                                    if (fi29 >= 42) {
                                                                        if (!(99 < 42)) {
                                                                            v167 := GenS28{F0: (-(v6)), F1: 42, F2: 0}
                                                                            _ = v167
                                                                        } else {
                                                                            if (!(!(v2 || v2))) {
                                                                                if v2 {
                                                                                    var v168 uint32 = 0
                                                                                    v169 := GenS23{F0: 42, F1: 32767, F2: (^(v25))}
                                                                                    _ = v168
                                                                                    _ = v169
                                                                                }
                                                                            }
                                                                        }
                                                                        var v170 int32 = -131
                                                                        switch v5 {
                                                                        case "hello":
                                                                            var v171 uint = 42
                                                                            _ = v171
                                                                        case "a":
                                                                            v172 := GenS20{F0: (((0 >= (v40 >> ((1) & 63))) && true) || (1e10 == 0.1)), F1: (-(v48))}
                                                                            _ = v172
                                                                        case "":
                                                                            _ = len(v1)
                                                                        default:
                                                                            if ((1 < 2) && v2) {
                                                                                var v173 int16 = v25
                                                                                v174 := GenS29{F0: 0, F1: 46}
                                                                                switch (^(fi29)) {
                                                                                case 1:
                                                                                    v175 := GenS10{F0: (v48 + v48)}
                                                                                    _ = v175
                                                                                }
                                                                                switch (v42 | (v42 | fi28)) {
                                                                                case 0:
                                                                                    _ = len(v3)
                                                                                }
                                                                                v176 := GenS22{F0: 4294967295, F1: 721}
                                                                                _ = v173
                                                                                _ = v174
                                                                                _ = v176
                                                                            } else {
                                                                                delete(v1, -2147483648)
                                                                                switch (v30 ^ v30) {
                                                                                case 127:
                                                                                    v48 = (v6 * (v6 / (((v48 / ((v48) | 1))) | 1)))
                                                                                case 2147483647:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                case 68:
                                                                                    v3 = append(v3, (("\n\t" + ("hello" + v5)) + (v5 + v5)))
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v170
                                                                    } else {
                                                                        v2 = v2
                                                                        var v177 bool = v2
                                                                        if (4294967295 < 2) {
                                                                            if v2 {
                                                                                _ = len(v1)
                                                                                v178 := GenS30{F0: 137}
                                                                                _ = v178
                                                                            } else {
                                                                                if v177 {
                                                                                    v6 = 1
                                                                                }
                                                                            }
                                                                            fmt.Println(v2, v177)
                                                                            delete(v1, (v30 >> ((255) & 63)))
                                                                            v177 = ((v177 || (821 == 882)) || v177)
                                                                        }
                                                                        _ = v177
                                                                    }
                                                                }
                                                                if (((v5 > "foo bar") || v2) && v4) {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            } else {
                                                                v179 := GenS31{F0: 1}
                                                                v180 := GenS18{F0: 65535, F1: 65535}
                                                                var v181 uint8 = 128
                                                                _ = v179
                                                                _ = v180
                                                                _ = v181
                                                            }
                                                        }
                                                        v182 := GenS32{F0: (v26 << ((0) & 63)), F1: 1, F2: (v5 + ("a" + v5))}
                                                        _ = v182
                                                    }
                                                    _ = v48
                                                }
                                            case "a":
                                                fmt.Println(fi29)
                                                v30 = 127
                                            case "foo bar":
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(v30, v42, v40)
                                                }
                                            default:
                                                v183 := GenS3{F0: 42}
                                                v3 = append(v3, "foo bar")
                                                _ = v183
                                            }
                                            if ((!(0 >= 0)) || v4) {
                                                var v184 uint16 = 1
                                                _ = v184
                                            }
                                            fmt.Println(v6, v4)
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        var v185 string = "hello"
                                        v2 = false
                                        switch 32768 {
                                        case -253:
                                            if v4 {
                                                v185 = "foo bar"
                                                v35 = (v35 - (v35 + 0.0))
                                            }
                                        case 32767:
                                            var v186 int32 = -2147483648
                                            _ = v186
                                        case 520:
                                            if true {
                                                v26 = v26
                                                v187 := GenS33{F0: 1}
                                                _ = v187
                                            } else {
                                                fmt.Println(v185)
                                                v188 := GenS34{F0: (v25 * 48), F1: (v26 / ((v26) | 1)), F2: 921}
                                                var v189 int = 2147483647
                                                if (255 < 979) {
                                                    var v190 uint8 = 1
                                                    if v2 {
                                                        v190 = (^(v190))
                                                        v191 := v1[fi29]
                                                        var v192 float64 = 0.1
                                                        _ = v191
                                                        _ = v192
                                                    } else {
                                                        v189 += (fi29 - 132)
                                                        var v193 int64 = (v6 / (((v6 & (^(v6)))) | 1))
                                                        _ = v193
                                                    }
                                                    _ = v190
                                                }
                                                _ = v188
                                                _ = v189
                                            }
                                        default:
                                            var v194 float32 = 1e10
                                            _ = v194
                                        }
                                        _ = v185
                                        _ = v35
                                    }
                                    switch (fi29 / (((fi28 * (fi29 + v30))) | 1)) {
                                    case -823:
                                        switch v5 {
                                        case "":
                                            var v195 int16 = v25
                                            _ = v195
                                        case "a":
                                            v196 := GenS9{F0: 2, F1: 257, F2: (255 != 2)}
                                            _ = v196
                                        }
                                    case 2147483647:
                                        var v197 uint32 = 999
                                        _ = v197
                                    case 4294967295:
                                        var v198 int32 = 0
                                        _ = v198
                                    default:
                                        v3 = append(v3, v5)
                                    }
                                } else {
                                    if v2 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            fmt.Println(v2, fi28)
                                            v199 := v1[fi29]
                                            _ = v199
                                        }
                                    }
                                    delete(v1, fi29)
                                }
                            }
                            _ = fi29
                            _ = v30
                        }
                        var v200 int64 = (v6 << ((1) & 63))
                        v26 -= (v26 ^ v26)
                        _ = v200
                    } else {
                        switch ("hello" + v5) {
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "":
                            for fi201 := 0; fi201 < 16; fi201++ {
                                switch 32768 {
                                case 9223372036854775807:
                                    delete(v1, fi201)
                                    if (!v4) {
                                        v25 -= -32768
                                        var v202 uint32 = 65535
                                        v202 = v202
                                        v203 := v1[9223372036854775807]
                                        _ = v202
                                        _ = v203
                                    } else {
                                        switch fi201 {
                                        case -198:
                                            var v204 uint32 = 2147483647
                                            _ = v204
                                        case 1:
                                            v205 := GenS35{F0: (18446744073709551615 <= 560), F1: 2, F2: 18446744073709551615}
                                            _ = v205
                                        case -9223372036854775808:
                                            if (!v4) {
                                                _ = len(v1)
                                                v4 = false
                                                v2 = false
                                            } else {
                                                if false {
                                                    v206 := GenS15{F0: (v26 | v26)}
                                                    v207 := GenS33{F0: 432}
                                                    var v208 uint = 2
                                                    _ = v206
                                                    _ = v207
                                                    _ = v208
                                                }
                                                v2 = v2
                                            }
                                        default:
                                            v6 = (v6 * v6)
                                            v2 = (int32(v26) < -992)
                                            v209 := GenS4{F0: 42, F1: 42, F2: 2}
                                            var v210 string = (v5 + v5)
                                            _ = v209
                                            _ = v210
                                        }
                                        v26 = v26
                                    }
                                case -866:
                                    if (false || (false && v4)) {
                                        if v2 {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v4 = (v2 || true)
                                            if true {
                                                if false {
                                                    v5 = "hello"
                                                } else {
                                                    fmt.Println(v26, v25, fi28)
                                                    var v211 int16 = v25
                                                    var v212 uint8 = 2
                                                    if v4 {
                                                        var v213 string = "hello"
                                                        _ = v213
                                                    }
                                                    _ = v211
                                                    _ = v212
                                                }
                                            }
                                        }
                                        if v2 {
                                            var v214 uint16 = genFn0(2, (v25 << ((724) & 63)))
                                            v25 = -32768
                                            _ = v214
                                        }
                                        if (v4 || (v2 && true)) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            fmt.Println(v25, v2)
                                            delete(v1, (fi28 << ((20) & 63)))
                                            fmt.Println(fi201, v2)
                                        }
                                    }
                                case -830:
                                    v215 := GenS30{F0: 2147483648}
                                    _ = v215
                                }
                                if (255 <= 65535) {
                                    if (1 != 0) {
                                        switch fi201 {
                                        case 9223372036854775807:
                                            fmt.Println(v25, v26, v2)
                                        }
                                    }
                                    var v216 int = (fi28 | fi201)
                                    switch v5 {
                                    case "hello":
                                        v216 = fi28
                                    default:
                                        v2 = (false || v2)
                                        v5 = v5
                                    }
                                    _ = v216
                                } else {
                                    v5 = v5
                                }
                                _ = fi201
                            }
                        case "foo bar":
                            for fi217 := 0; fi217 < 16; fi217++ {
                                v1[(fi28 % (((fi28 ^ (^(fi217)))) | 1))] = ("a" + v5)
                                if v4 {
                                    v6 -= v6
                                }
                                if v2 {
                                    var v218 float32 = (1e10 / 3.14159)
                                    _ = v218
                                } else {
                                    var v219 uint = 434
                                    if v4 {
                                        v5 = v5
                                        switch "foo bar" {
                                        case "hello":
                                            var v220 uint64 = 42
                                            _ = v220
                                        default:
                                            var v221 uint = (v219 << ((42) & 63))
                                            _ = v221
                                        }
                                        v3 = append(v3, (v5 + "hello"))
                                        v4 = v2
                                        v5 = (("" + (((v5 + (("\n\t" + "") + v5)) + (v5 + (v5 + v5))) + "hello")) + "foo bar")
                                    }
                                    var v222 uint8 = 157
                                    if (1.5 != 1e10) {
                                        var v223 uint16 = 255
                                        if (!false) {
                                            v26 = v26
                                        }
                                        _ = v223
                                    }
                                    var v224 int16 = (v25 + v25)
                                    v224 -= (v224 / ((v25) | 1))
                                    _ = v219
                                    _ = v222
                                    _ = v224
                                }
                                _ = fi217
                            }
                        default:
                            var v225 uint32 = 845
                            _ = v225
                        }
                        if (v5 > "\n\t") {
                            if (-842 > 1) {
                                if len(v3) > 0 {
                                    v3[0] = (("a" + ((v5 + "") + "hello")) + v5)
                                }
                                var v226 uint8 = 2
                                _ = v226
                            }
                            v25 -= (v25 << ((65535) & 63))
                            if len(v3) > 0 {
                                v3[0] = v5
                            }
                        } else {
                            v25 += (-(v25))
                            var v227 int32 = 2147483647
                            _ = v227
                        }
                        fmt.Println(v6)
                    }
                case "foo bar":
                    for fi228 := 0; fi228 < 17; fi228++ {
                        var v229 string = (("hello" + "\n\t") + (("" + (v5 + v5)) + "\n\t"))
                        if (1 >= 42) {
                            var v230 uint8 = 1
                            fmt.Println(fi228, v5, v4)
                            _ = v230
                        }
                        _ = fi228
                        _ = v229
                    }
                case "\n\t":
                    v231 := GenS33{F0: 42}
                    _ = v231
                default:
                    fmt.Println(v5)
                }
                if v2 {
                    for fi232 := 0; fi232 < 16; fi232++ {
                        var v233 uint8 = 94
                        if false {
                            if (true || (0 >= 0)) {
                                var v234 int64 = (v6 % (((-(v6))) | 1))
                                v26 += 0
                                v3 = append(v3, "a")
                                _ = v234
                            } else {
                                v6 = v6
                            }
                        }
                        _ = fi232
                        _ = v233
                    }
                    fmt.Println(v4, fi28)
                    fmt.Println(v26, v25)
                } else {
                    v2 = v4
                    v235 := GenS36{F0: 2147483647, F1: 0, F2: 42}
                    _ = v235
                }
                _ = fi28
            }
            _ = v27
        } else {
            var v236 int64 = v6
            v2 = (!(!(v4 || (0 >= 0))))
            var v237 int = 127
            _ = v236
            _ = v237
        }
        _ = v25
        _ = v26
    }
    fmt.Println(v2, v6)
    var v238 uint16 = 2
    delete(v1, 128)
    var v239 bool = (4294967295 >= 0)
    v2 = (v2 || false)
    for fi240 := 0; fi240 < 7; fi240++ {
        v241 := GenS16{F0: 0, F1: 18446744073709551615, F2: 2147483648}
        _ = fi240
        _ = v241
    }
    v4 = false
    v242 := GenS15{F0: 0}
    if v4 {
        v5 = ((((v5 + v5) + v5) + "foo bar") + v5)
        var v243 float64 = 0.0
        _ = v243
    } else {
        if (!false) {
            for fi244 := 0; fi244 < 4; fi244++ {
                for fi245 := 0; fi245 < 1; fi245++ {
                    v2 = (0 != 0)
                    _ = fi245
                }
                if (32767 <= 685) {
                    v246 := GenS10{F0: (-(v6))}
                    if true {
                        v247 := GenS37{F0: 27, F1: 127}
                        v248 := GenS26{F0: 42}
                        for fi249 := 0; fi249 < 1; fi249++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (v239 || v4) {
                                _ = len(v1)
                                var v250 float32 = 0.0
                                var v251 int64 = (v6 >> ((1) & 63))
                                _ = v250
                                _ = v251
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi249
                        }
                        _ = v247
                        _ = v248
                    } else {
                        _ = len(v3)
                    }
                    var v252 uint64 = 203
                    _ = v246
                    _ = v252
                } else {
                    var v253 float32 = -2.25
                    _ = v253
                }
                _ = fi244
            }
            v238 -= genFn0(1, 42)
            v2 = false
        }
        v238 = 65535
    }
    v254 := GenS10{F0: v6}
    var v255 uint8 = 167
    var v256 uint = 431
    for fi257 := 0; fi257 < 8; fi257++ {
        v6 = (-(v6))
        for fi258 := 0; fi258 < 18; fi258++ {
            v239 = (true && (v4 || v4))
            _ = fi258
        }
        var v259 int = 0
        _ = fi257
        _ = v259
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v238:", v238)
    fmt.Println("v239:", v239)
    fmt.Println("v255:", v255)
    fmt.Println("v256:", v256)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v238
    _ = v239
    _ = v242
    _ = v254
    _ = v255
    _ = v256
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
