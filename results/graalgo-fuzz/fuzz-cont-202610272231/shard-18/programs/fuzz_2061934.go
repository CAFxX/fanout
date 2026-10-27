package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint }
type GenS6 struct { F0 int64; F1 string; F2 uint }
type GenS7 struct { F0 uint; F1 string }
func genFn0(p0 uint16, p1 int8) uint32 {
    v7 := []int{}
    if (false || false) {
        var v8 uint = 255
        p1 -= -1
        for fi9 := 0; fi9 < 1; fi9++ {
            if true {
                switch fi9 {
                case 32768:
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 bool = ((9223372036854775807 != 719) || (149 > 255))
                        var v12 float32 = -2.25
                        if (0.0 != 3.14159) {
                            switch (fi10 - -2147483648) {
                            case 0:
                                fmt.Println(p1, v11, fi10)
                            case 127:
                                p1 = (^(p1))
                            case -2147483648:
                                if len(v7) > 0 {
                                    v7[0] = fi10
                                }
                            }
                            if v11 {
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    fmt.Println(fi9)
                                case "foo bar":
                                    if false {
                                        v11 = (((("" + "a") + "") + "a") < "hello")
                                        p1 = p1
                                        switch (("" + "a") + "") {
                                        case "foo bar":
                                            if (0.0 < 1.5) {
                                                if v11 {
                                                    v12 += 0.0
                                                    switch 128 {
                                                    case -403:
                                                        v13 := GenS1{F0: "a", F1: -1}
                                                        _ = v13
                                                    case -23:
                                                        fmt.Println(p1, fi10, v12)
                                                    case -477:
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                    }
                                                }
                                                v7 = append(v7, fi9)
                                            }
                                            v11 = (!false)
                                            var v14 uint32 = 129
                                            _ = v14
                                        case "\n\t":
                                            var v15 float64 = (0.0 * 0.1)
                                            _ = v15
                                        default:
                                            if false {
                                                v12 = 0.1
                                            }
                                        }
                                    } else {
                                        if false {
                                            v16 := GenS2{F0: -1}
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            p0 += p0
                                            var v17 uint16 = 50
                                            _ = v16
                                            _ = v17
                                        } else {
                                            var v18 bool = true
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            _ = v18
                                        }
                                        fmt.Println(p0, fi9)
                                        var v19 uint64 = 4294967295
                                        _ = v19
                                    }
                                case "hello":
                                    var v20 int8 = (p1 % (((-(p1))) | 1))
                                    _ = v20
                                default:
                                    _ = len(v7)
                                }
                                v21 := GenS3{F0: 32768}
                                _ = v21
                            } else {
                                v22 := GenS1{F0: ((("\n\t" + ("a" + "a")) + "") + ("" + ("hello" + ""))), F1: 1}
                                var v23 int = (fi10 % (((fi9 << ((18446744073709551615) & 63))) | 1))
                                _ = v22
                                _ = v23
                            }
                        } else {
                            var v24 float64 = 3.14159
                            var v25 uint16 = (p0 * 185)
                            _ = v24
                            _ = v25
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case -511:
                    v26 := GenS4{F0: (^(p1)), F1: (p0 | (^(p0))), F2: p0}
                    _ = v26
                case -742:
                    v7 = append(v7, 9223372036854775807)
                }
                v27 := GenS2{F0: -1}
                _ = v27
            }
            _ = fi9
        }
        _ = v8
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 int64 = -119
        switch "\n\t" {
        case "":
            v30 := GenS5{F0: 1, F1: 1, F2: 18446744073709551615}
            _ = v30
        case "hello":
            var v31 float32 = 1.5
            _ = v31
        default:
            _ = len(v7)
        }
        p1 = (p1 ^ 1)
        _ = fi28
        _ = v29
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := GenS1{F0: "", F1: -1}
        fmt.Println(p1, p0, fi32)
        _ = fi32
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 bool) bool {
    for fi34 := 0; fi34 < 14; fi34++ {
        p0 -= p0
        v35 := []int{128, fi34, fi34}
        if true {
            var v36 bool = (false || p1)
            _ = len(v35)
            _ = v36
        } else {
            if len(v35) >= 2 {
                v35 = v35[1:len(v35)]
            }
        }
        _ = fi34
        _ = v35
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 int = -470
    var v2 uint32 = 65535
    var v3 uint32 = v2
    var v4 int32 = 0
    v5 := make(map[int]uint32)
    var v6 int8 = -1
    v37 := GenS4{F0: v6, F1: 912, F2: 65535}
    v6 = v6
    v2 -= (v2 | v2)
    var v38 float32 = 3.14159
    switch v1 {
    case 4294967295:
        v2 = v2
    default:
        v39 := GenS5{F0: 673, F1: genFn0(2, (v6 | (v6 + v6))), F2: 2}
        _ = v39
    }
    if (-9223372036854775808 > -1) {
        switch v1 {
        case 32767:
            var v40 float32 = (v38 + v38)
            _ = v40
        case 1:
            v41 := GenS6{F0: -9223372036854775808, F1: "hello", F2: 260}
            _ = v41
        }
    } else {
        if (((("foo bar" + ("foo bar" + "")) + ("" + "foo bar")) + "a") != "") {
            if true {
                if (v38 != v38) {
                    v42 := v5[v1]
                    for fi43 := 0; fi43 < 16; fi43++ {
                        switch v1 {
                        case -9223372036854775808:
                            var v44 float32 = (-2.25 - float32(0.1))
                            _ = v44
                        }
                        _ = fi43
                    }
                    for fi45 := 0; fi45 < 17; fi45++ {
                        var v46 uint16 = uint16(v42)
                        fmt.Println(v6, fi45)
                        v1 -= (v1 & 130)
                        _ = fi45
                        _ = v46
                    }
                    _ = v42
                }
                if (-32768 != 0) {
                    v47 := GenS5{F0: 3, F1: v3, F2: 1}
                    _ = v47
                } else {
                    var v48 uint = 42
                    v3 = (^(v2))
                    _ = v48
                }
                var v49 int16 = 0
                _ = v49
            }
            for fi50 := 0; fi50 < 6; fi50++ {
                var v51 int16 = -1
                switch "foo bar" {
                case "a":
                    var v52 uint8 = 1
                    _ = v52
                default:
                    v53 := GenS1{F0: "hello", F1: -1}
                    _ = v53
                }
                var v54 bool = (false && (v51 < (^(v51))))
                _ = fi50
                _ = v51
                _ = v54
            }
        } else {
            v1 += 4294967295
            var v55 int16 = 0
            v56 := []int{v1, v1}
            _ = v55
            _ = v56
        }
        fmt.Println(v4, v4)
    }
    v57 := v5[(v1 | v1)]
    if true {
        v6 += v6
        v58 := GenS2{F0: (v4 & (^(v4)))}
        var v59 bool = true
        _ = v58
        _ = v59
    } else {
        for fi60 := 0; fi60 < 2; fi60++ {
            var v61 int8 = (v6 + v6)
            if (!((9223372036854775807 == (^(v1))) && ((v4 + 2147483647) != v4))) {
                var v62 int8 = v61
                var v63 uint8 = 217
                _ = v62
                _ = v63
            }
            var v64 uint8 = 9
            _ = fi60
            _ = v61
            _ = v64
        }
        var v65 uint64 = 986
        v1 = -281
        _ = v65
    }
    delete(v5, v1)
    v2 -= 2147483648
    v57 = v2
    for fi66 := 0; fi66 < 3; fi66++ {
        switch "a" {
        case "hello":
            v67 := GenS7{F0: 201, F1: (((("foo bar" + "\n\t") + "a") + "") + ("hello" + (("" + "\n\t") + "")))}
            _ = v67
        }
        for fi68 := 0; fi68 < 13; fi68++ {
            fmt.Println(v2)
            switch "a" {
            case "foo bar":
                v5[v1] = v57
            case "hello":
                v69 := GenS4{F0: v6, F1: 780, F2: 0}
                _ = v69
            case "\n\t":
                v6 = v6
            default:
                var v70 float64 = 1e10
                _ = v70
            }
            v2 += 695
            _ = fi68
        }
        _ = fi66
    }
    if ((((0.0 + v38) - 0.1) / ((((v38 - ((0.0 - 0.0) / v38)) / (v38 - ((0.1 * (v38 - v38)) + (1e10 * 0.0)))) * v38) / (float32((1.5 / 1.5)) * -2.25))) == (1.5 + (0.1 - v38))) {
        var v71 string = "hello"
        switch v71 {
        case "\n\t":
            for fi72 := 0; fi72 < 10; fi72++ {
                var v73 int = (^(fi72))
                fmt.Println(v38)
                _ = fi72
                _ = v73
            }
        default:
            var v74 uint = 18446744073709551615
            _ = v74
        }
        _ = v71
    } else {
        fmt.Println(v57, v4)
        v1 += 32768
        v2 = (^(v2))
    }
    switch "a" {
    case "\n\t":
        if ((1.5 + (3.14159 - 0.0)) >= 1.5) {
            if true {
                for fi75 := 0; fi75 < 4; fi75++ {
                    if ((v4 ^ v4) >= (-(v4))) {
                        var v76 int64 = int64(v4)
                        if true {
                            for fi77 := 0; fi77 < 20; fi77++ {
                                var v78 uint64 = 18446744073709551615
                                _ = fi77
                                _ = v78
                            }
                            if true {
                                var v79 uint8 = 2
                                _ = v79
                            }
                            v3 += (v57 << ((1) & 63))
                        } else {
                            v1 = v1
                            var v80 bool = (854 > 3)
                            _ = v80
                        }
                        _ = v76
                    }
                    switch 474 {
                    case 42:
                        if (0 >= -302) {
                            v81 := []int{(v1 ^ (v1 | -2147483648)), (-(fi75)), (fi75 & v1)}
                            var v82 uint8 = 0
                            _ = v81
                            _ = v82
                        }
                    case 2147483647:
                        for fi83 := 0; fi83 < 17; fi83++ {
                            if true {
                                if false {
                                    var v84 string = (("" + "\n\t") + (("a" + ("a" + "\n\t")) + "\n\t"))
                                    fmt.Println(v6)
                                    _ = v84
                                } else {
                                    if true {
                                        fmt.Println(v6)
                                        v85 := GenS6{F0: -9223372036854775808, F1: "\n\t", F2: 42}
                                        var v86 int16 = 437
                                        _ = v85
                                        _ = v86
                                    } else {
                                        var v87 bool = false
                                        fmt.Println(v38, v1, v3)
                                        v3 = v2
                                        _ = v87
                                    }
                                }
                                var v88 float64 = 1.5
                                v89 := GenS1{F0: "\n\t", F1: 42}
                                _ = v88
                                _ = v89
                            } else {
                                v2 -= v57
                                var v90 float64 = (1e10 - 1.5)
                                _ = v90
                            }
                            _ = fi83
                        }
                    case 748:
                        for fi91 := 0; fi91 < 6; fi91++ {
                            if (567 >= 2) {
                                v1 += fi75
                                var v92 float32 = -2.25
                                _ = v92
                            } else {
                                if (!(65535 == 4294967295)) {
                                    v38 -= v38
                                }
                                switch -2147483648 {
                                case -9223372036854775808:
                                    switch (fi91 / ((v1) | 1)) {
                                    case 9223372036854775807:
                                        v4 = v4
                                    case 42:
                                        v93 := GenS1{F0: "hello", F1: 9223372036854775807}
                                        _ = v93
                                    case -1:
                                        if true {
                                            var v94 uint8 = 127
                                            var v95 uint8 = 1
                                            _ = v94
                                            _ = v95
                                        }
                                    default:
                                        var v96 float64 = 1e10
                                        _ = v96
                                    }
                                case 195:
                                    var v97 int = 32767
                                    _ = v97
                                }
                                v38 = v38
                            }
                            _ = fi91
                        }
                    }
                    for fi98 := 0; fi98 < 13; fi98++ {
                        var v99 int16 = int16(v57)
                        var v100 uint = 2
                        v3 += v3
                        _ = fi98
                        _ = v100
                        _ = v99
                    }
                    _ = fi75
                }
                delete(v5, v1)
            } else {
                fmt.Println(v1)
                v101 := []int{v1}
                v3 += (v57 * 72)
                _ = v101
            }
            v102 := []int{(v1 / ((-880) | 1)), v1}
            _ = v102
        }
    default:
        v6 = 42
    }
    switch v1 {
    case 0:
        v5[-816] = v57
    case 42:
        for fi103 := 0; fi103 < 17; fi103++ {
            var v104 int = -1
            v104 -= v1
            _ = fi103
            _ = v104
        }
    case 32768:
        var v105 int16 = 0
        _ = v105
    default:
        var v106 uint32 = (v57 % ((246) | 1))
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
