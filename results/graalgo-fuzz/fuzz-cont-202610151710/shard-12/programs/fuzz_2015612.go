package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int8; F2 uint64 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint; F1 uint32; F2 string }
type GenS11 struct { F0 int64 }

func main() {
    var v1 uint = 42
    var v2 uint16 = 65535
    var v3 int = 4294967295
    var v4 uint32 = 2147483648
    var v5 uint16 = v2
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = "a"
            if true {
                if (int16(fi7) >= 1000) {
                    switch v8 {
                    case "foo bar":
                        v1 += (v1 * v1)
                    case "\n\t":
                        if true {
                            fmt.Println(v8, v2)
                        } else {
                            switch "\n\t" {
                            case "a":
                                if false {
                                    if false {
                                        var v9 uint32 = 2
                                        var v10 int32 = 4
                                        switch (v8 + v8) {
                                        case "foo bar":
                                            if (0.0 >= 3.14159) {
                                                v11 := GenS1{F0: 128}
                                                v12 := []int{(-(v3)), fi6, fi6, int(v10)}
                                                _ = v11
                                                _ = v12
                                            } else {
                                                fmt.Println(v10)
                                                v13 := GenS2{F0: "\n\t", F1: 234, F2: 1}
                                                if (!(false && false)) {
                                                    v14 := []int{fi7, 9223372036854775807, fi7}
                                                    var v15 int16 = 42
                                                    v2 = v5
                                                    _ = v14
                                                    _ = v15
                                                }
                                                _ = v13
                                            }
                                        }
                                        _ = v10
                                        _ = v9
                                    }
                                } else {
                                    switch "" {
                                    case "hello":
                                        v16 := GenS2{F0: v8, F1: 255, F2: (v2 ^ (^(v5)))}
                                        _ = v16
                                    case "\n\t":
                                        fmt.Println(v1, v4, v1)
                                    }
                                    switch "a" {
                                    case "a":
                                        var v17 uint32 = (v4 - 65535)
                                        _ = v17
                                    case "foo bar":
                                        v18 := []int{952, (v3 + (fi7 + -185))}
                                        _ = v18
                                    case "":
                                        v19 := GenS1{F0: 127}
                                        _ = v19
                                    }
                                }
                            case "":
                                v20 := []int{(fi6 << ((v1) & 63))}
                                _ = v20
                            }
                            v21 := []int{}
                            _ = v21
                        }
                    }
                }
                switch (fi7 / (((fi7 | 2147483647)) | 1)) {
                case 1:
                    var v22 float32 = 0.0
                    _ = v22
                case 128:
                    v23 := GenS3{F0: v3, F1: 18446744073709551615}
                    _ = v23
                }
            } else {
                v5 -= (v2 >> (((v1 | v1)) & 63))
                switch fi6 {
                case 4294967295:
                    v2 -= v5
                case 2147483647:
                    var v24 float32 = 1.5
                    _ = v24
                }
                v2 = 2
                v8 = v8
            }
            v25 := GenS2{F0: v8, F1: (^(v5)), F2: v2}
            _ = fi7
            _ = v25
            _ = v8
        }
        _ = fi6
    }
    v26 := GenS1{F0: 179}
    for fi27 := 0; fi27 < 4; fi27++ {
        if true {
            if false {
                if ((v4 >> ((v1) & 63)) == (^(v4))) {
                    v5 -= v5
                } else {
                    if true {
                        for fi28 := 0; fi28 < 12; fi28++ {
                            var v29 bool = (255 < (v1 % (((v1 / (((^(v1))) | 1))) | 1)))
                            if v29 {
                                v30 := GenS3{F0: 32768, F1: 18446744073709551615}
                                _ = v30
                            }
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        v31 := GenS4{F0: 151, F1: 1000}
                        switch -825 {
                        case 411:
                            v4 -= 2147483647
                        case -138:
                            v32 := []int{v3}
                            _ = v32
                        default:
                            if false {
                                v33 := GenS5{F0: true}
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    if false {
                                        v4 -= v4
                                    } else {
                                        var v35 uint32 = (v4 - v4)
                                        v3 = 462
                                        _ = v35
                                    }
                                    var v36 uint32 = v4
                                    v37 := []int{}
                                    v38 := GenS3{F0: (fi34 & 128), F1: 0}
                                    v39 := GenS1{F0: 128}
                                    _ = fi34
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                v2 += 106
                                _ = v33
                            }
                        }
                        v40 := GenS6{F0: -1}
                        _ = v31
                        _ = v40
                    }
                }
                v41 := []int{}
                var v42 uint32 = (^(v4))
                _ = v41
                _ = v42
            }
            for fi43 := 0; fi43 < 8; fi43++ {
                if false {
                    var v44 int = fi43
                    v45 := GenS2{F0: (("foo bar" + ("foo bar" + ((("foo bar" + "") + "foo bar") + "\n\t"))) + "foo bar"), F1: (v2 ^ v2), F2: v5}
                    _ = v44
                    _ = v45
                }
                _ = fi43
            }
        } else {
            v4 = v4
            for fi46 := 0; fi46 < 13; fi46++ {
                fmt.Println(fi46)
                if (112 <= 1) {
                    v2 = (v5 & v2)
                } else {
                    var v47 int = 32768
                    _ = v47
                }
                switch v3 {
                case 32767:
                    fmt.Println(fi46, fi46)
                case -9223372036854775808:
                    if false {
                        fmt.Println(v1, v5, fi27)
                    } else {
                        if false {
                            if (!(((v1 / ((255) | 1)) == v1) && true)) {
                                fmt.Println(v4)
                                if true {
                                    var v48 bool = ((true && false) || (true && ((0 <= 697) && ((^(v1)) >= (v1 << ((uint(v2)) & 63))))))
                                    v49 := GenS1{F0: 1}
                                    var v50 int = (fi46 << (((^(v1))) & 63))
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                }
                                switch 32767 {
                                case 127:
                                    fmt.Println(v2)
                                case -1:
                                    v2 = 2
                                    if ((v3 + v3) == (-(fi46))) {
                                        v4 = v4
                                    }
                                    v51 := GenS7{F0: -1, F1: 0, F2: 1}
                                    _ = v51
                                case -277:
                                    fmt.Println(v4)
                                default:
                                    v1 -= 861
                                    var v52 uint8 = 127
                                    v53 := []int{fi27, fi27, (fi27 | (-(fi27)))}
                                    _ = v52
                                    _ = v53
                                }
                                switch "hello" {
                                case "a":
                                    v54 := []int{72, fi46, v3}
                                    _ = v54
                                }
                            } else {
                                v4 = v4
                                if true {
                                    v5 = (v5 + (v2 | (v2 % ((v5) | 1))))
                                } else {
                                    v55 := []int{fi46, (v3 / ((fi46) | 1)), v3}
                                    _ = v55
                                }
                                v56 := GenS6{F0: 0}
                                if true {
                                    var v57 int16 = 1
                                    v58 := GenS6{F0: v57}
                                    _ = v57
                                    _ = v58
                                }
                                v3 += fi27
                                _ = v56
                            }
                            v59 := GenS2{F0: "a", F1: (v2 % ((2) | 1)), F2: (v5 * 0)}
                            v1 = uint(v2)
                            _ = v59
                        } else {
                            v2 = 1
                            switch (fi46 | 42) {
                            case 32768:
                                v4 = 723
                            case 9223372036854775807:
                                v3 = (-(fi27))
                            }
                            fmt.Println(v2, v2)
                            fmt.Println(fi46, v1, v3)
                            if true {
                                fmt.Println(v1, v3, v3)
                                v60 := GenS7{F0: -792, F1: -14, F2: 0}
                                _ = v60
                            } else {
                                var v61 bool = true
                                v62 := []int{}
                                v62 = append(v62, (fi27 + (^(fi27))))
                                _ = v61
                                _ = v62
                            }
                        }
                        var v63 uint64 = 917
                        v64 := []int{}
                        _ = v63
                        _ = v64
                    }
                case 2147483647:
                    fmt.Println(fi27, v5, v1)
                    var v65 int64 = -1
                    if false {
                        v3 += fi46
                        fmt.Println(fi27, v4)
                    } else {
                        v66 := GenS8{F0: 1, F1: ""}
                        v67 := GenS7{F0: -9223372036854775808, F1: -1, F2: 4294967295}
                        if false {
                            v4 += (v4 - v4)
                            if (false || ((42 == 128) || true)) {
                                if (0.1 < (0.0 / 0.1)) {
                                    v68 := []int{v3}
                                    if (((1.5 / float64(1e10)) - 1.5) >= float64((1e10 / 0.1))) {
                                        switch 2147483647 {
                                        case 0:
                                            v65 -= v65
                                        case 1:
                                            var v69 int64 = v65
                                            _ = v69
                                        case 32768:
                                            var v70 float64 = 0.1
                                            _ = v70
                                        default:
                                            v3 += (fi27 * fi27)
                                        }
                                        v71 := GenS6{F0: 32767}
                                        _ = v71
                                    }
                                    _ = v68
                                }
                                v72 := GenS6{F0: 0}
                                _ = v72
                            } else {
                                if false {
                                    var v73 bool = true
                                    if (v4 == v4) {
                                        var v74 uint16 = 1
                                        fmt.Println(v5, v3, v4)
                                        _ = v74
                                    } else {
                                        var v75 uint8 = 0
                                        v3 -= 42
                                        v76 := GenS3{F0: (fi27 & (fi27 * fi46)), F1: 31}
                                        v77 := GenS9{F0: 774}
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    }
                                    if ("hello" > "\n\t") {
                                        var v78 uint16 = v2
                                        _ = v78
                                    }
                                    _ = v73
                                }
                                v79 := []int{(-(fi27)), fi27, (fi46 + 32767), (v3 << (((v1 | v1)) & 63))}
                                _ = v79
                            }
                            switch (fi27 | -9223372036854775808) {
                            case 127:
                                fmt.Println(v5)
                            case 449:
                                var v80 float64 = float64((0.0 - (((((1e10 + 0.1) + 1e10) / 1e10) + 1.5) - 1e10)))
                                _ = v80
                            default:
                                v2 += (v5 >> ((0) & 63))
                            }
                        }
                        _ = v66
                        _ = v67
                    }
                    if true {
                        v81 := []int{}
                        var v82 int16 = 614
                        if (fi27 > v3) {
                            fmt.Println(v4, v2, v1)
                            fmt.Println(v82, v82, v2)
                        } else {
                            fmt.Println(v5, v2, v3)
                        }
                        if ((2 <= 42) && true) {
                            v1 = v1
                        }
                        _ = len(v81)
                        _ = v81
                        _ = v82
                    } else {
                        v83 := GenS10{F0: v1, F1: uint32(v5), F2: "\n\t"}
                        _ = v83
                    }
                    _ = v65
                }
                v84 := []int{(fi27 * (v3 << ((v1) & 63))), (fi27 & 1), 0, v3}
                switch 128 {
                case 1:
                    if (1e10 <= 0.1) {
                        if ((v1 | v1) <= 0) {
                            v85 := GenS11{F0: 42}
                            var v86 bool = (161 <= 18446744073709551615)
                            if len(v84) > 0 {
                                v84[0] = (^(fi46))
                            }
                            _ = v85
                            _ = v86
                        } else {
                            v2 = 1
                            fmt.Println(v1, v3)
                        }
                        var v87 bool = false
                        if v87 {
                            if len(v84) >= 2 {
                                v84 = v84[1:len(v84)]
                            }
                            fmt.Println(v87, v3, v3)
                        }
                        _ = v87
                    }
                case -2147483648:
                    var v88 uint8 = 1
                    _ = v88
                default:
                    if true {
                        v5 = (^(v5))
                    }
                }
                if true {
                    if len(v84) > 0 {
                        v84[0] = fi46
                    }
                    var v89 bool = (!(32767 >= -1))
                    v84 = append(v84, 128)
                    _ = v89
                }
                _ = fi46
                _ = v84
            }
        }
        _ = fi27
    }
    fmt.Println(v3)
    v90 := GenS8{F0: -1, F1: "a"}
    var v91 int = (v3 ^ v3)
    v2 += (^(v2))
    v92 := GenS3{F0: (v3 & 42), F1: 546}
    v93 := []int{v91, (v91 % (((v3 | (v3 >> ((42) & 63)))) | 1)), v91, v91}
    var v94 int = v3
    if true {
        v3 = v3
    } else {
        var v95 string = ("foo bar" + ("hello" + "foo bar"))
        v91 += v91
        _ = v95
    }
    v91 = 1
    if true {
        v93 = append(v93, (v91 - 679))
    }
    switch ("\n\t" + ("" + "\n\t")) {
    case "\n\t":
        if true {
            v96 := GenS10{F0: 255, F1: 0, F2: ("foo bar" + "a")}
            if len(v93) >= 2 {
                v93 = v93[1:len(v93)]
            }
            _ = v96
        } else {
            v97 := GenS5{F0: (-128 <= 0)}
            _ = v97
        }
    }
    var v98 uint64 = 1
    v2 += v2
    var v99 float32 = (-2.25 * float32(0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    fmt.Println("slicelen v93:", len(v93), v93)
    fmt.Println("v94:", v94)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v90
    _ = v91
    _ = v92
    _ = v93
    _ = v94
    _ = v98
    _ = v99
}
