package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint64; F1 int32; F2 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16 }
func genFn0(p0 uint, p1 string) uint32 {
    fmt.Println(p0, p0)
    switch 686 {
    case -68:
        switch "foo bar" {
        case "hello":
            p0 -= 65535
        case "a":
            switch 42 {
            case -9223372036854775808:
                switch p1 {
                case "\n\t":
                    v5 := GenS1{F0: true, F1: 42}
                    _ = v5
                case "hello":
                    var v6 uint64 = 18446744073709551615
                    _ = v6
                case "foo bar":
                    switch p1 {
                    case "\n\t":
                        fmt.Println(p1)
                    case "a":
                        var v7 int = -308
                        _ = v7
                    }
                default:
                    var v8 float64 = 1.5
                    _ = v8
                }
            case -1:
                for fi9 := 0; fi9 < 14; fi9++ {
                    fmt.Println(p1, p0, fi9)
                    v10 := []int{}
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v12 := GenS1{F0: false, F1: 1}
                        _ = fi11
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
            }
        }
    case -1:
        v13 := GenS2{F0: "\n\t"}
        _ = v13
    case -2147483648:
        var v14 int16 = 32767
        _ = v14
    default:
        var v15 int8 = -128
        _ = v15
    }
    if ((0 == 350) || false) {
        fmt.Println(p0)
        fmt.Println(p0, p0)
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(p0, p1, p1)
            _ = fi16
        }
    } else {
        v17 := GenS1{F0: (255 > 32767), F1: 0}
        fmt.Println(p0, p0)
        _ = v17
    }
    var v18 int32 = 407
    var v19 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 2147483647
}
func genFn1(p0 string, p1 uint) bool {
    fmt.Println(p1)
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int64 = 9223372036854775807
        v21 = 1
        p0 = ""
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]int)
    var v3 bool = (0 <= 255)
    var v4 float64 = 0.1
    if (708 != 0) {
        var v22 uint32 = 2
        v3 = false
        _ = v22
    }
    if (0 < 1) {
        v23 := GenS1{F0: true, F1: -128}
        fmt.Println(v3, v3)
        _ = v23
    } else {
        var v24 float32 = 3.14159
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS1{F0: true, F1: -1}
            if v3 {
                v4 -= float64(v24)
                for fi27 := 0; fi27 < 3; fi27++ {
                    v28 := GenS1{F0: v3, F1: 127}
                    v29 := GenS3{F0: 42, F1: 2, F2: 2}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                fmt.Println(v1, v3, v3)
            } else {
                fmt.Println(v3)
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    var v30 int16 = 32767
    if ((v3 || v3) || v3) {
        var v31 int32 = 1
        if v3 {
            v31 = (v31 << ((427) & 63))
        } else {
            for fi32 := 0; fi32 < 7; fi32++ {
                if v3 {
                    if (568 < 553) {
                        var v33 int16 = 853
                        var v34 float64 = 0.0
                        _ = v33
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 20; fi35++ {
                            var v36 int = 997
                            _ = fi35
                            _ = v36
                        }
                        var v37 uint16 = 255
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 13; fi38++ {
                        switch (("" + "hello") + "") {
                        case "foo bar":
                            v1 = (v1 % ((fi32) | 1))
                        case "\n\t":
                            var v39 uint64 = 2
                            _ = v39
                        }
                        _ = fi38
                    }
                } else {
                    switch v1 {
                    case 9223372036854775807:
                        v40 := GenS4{F0: 1}
                        _ = v40
                    case 668:
                        if (0 != -3) {
                            v31 = 556
                            for fi41 := 0; fi41 < 7; fi41++ {
                                v42 := GenS5{F0: 68}
                                v43 := GenS3{F0: (v30 % ((1000) | 1)), F1: 144, F2: 4294967295}
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            v31 += (-(v31))
                        }
                    }
                }
                v44 := GenS6{F0: -920, F1: v3, F2: 42}
                v30 = (v30 * v30)
                _ = fi32
                _ = v44
            }
            v45 := v2[genFn0(18446744073709551615, ("" + ("a" + "\n\t")))]
            _ = v45
        }
        _ = v31
    }
    if v3 {
        if v3 {
            v46 := GenS4{F0: 1}
            _ = v46
        } else {
            var v47 uint64 = 148
            _ = v47
        }
    }
    switch (v1 - 783) {
    case 1:
        var v48 uint16 = 0
        _ = v48
    default:
        v49 := GenS7{F0: 914, F1: (("" + "a") + ("" + ""))}
        _ = v49
    }
    var v50 uint = 0
    if (true || (32767 < 697)) {
        if (528 >= (v1 % ((4294967295) | 1))) {
            v1 = v1
            if v3 {
                for fi51 := 0; fi51 < 17; fi51++ {
                    for fi52 := 0; fi52 < 12; fi52++ {
                        if v3 {
                            if v3 {
                                var v53 int64 = -9223372036854775808
                                var v54 uint8 = 0
                                if (!v3) {
                                    v50 -= v50
                                    if (v3 && (false && v3)) {
                                        v1 = fi52
                                    } else {
                                        var v55 uint64 = 4294967295
                                        _ = v55
                                    }
                                    v1 = fi51
                                } else {
                                    var v56 float32 = 1e10
                                    _ = v56
                                }
                                _ = v53
                                _ = v54
                            }
                            v57 := v2[944]
                            _ = v57
                        } else {
                            if v3 {
                                fmt.Println(v1, v1, v30)
                                v58 := GenS8{F0: (fi52 >> ((v50) & 63)), F1: "\n\t"}
                                switch 32767 {
                                case 920:
                                    if true {
                                        var v59 uint16 = 42
                                        var v60 uint64 = 0
                                        v61 := GenS4{F0: 598}
                                        v62 := GenS8{F0: (v1 ^ fi51), F1: "foo bar"}
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        _ = len(v2)
                                        v2[genFn0((v50 % (((v50 | (v50 - 699))) | 1)), ((("hello" + "hello") + ("foo bar" + "\n\t")) + ""))] = (fi51 / ((v1) | 1))
                                        var v63 uint8 = 255
                                        _ = v63
                                    }
                                case 721:
                                    v64 := []int{int(v50), 32768, v1, -1}
                                    _ = v64
                                case -1:
                                    var v65 int16 = v30
                                    _ = v65
                                }
                                if (uint32(v50) > 42) {
                                    var v66 string = ((("hello" + "a") + ((("" + "a") + "\n\t") + "a")) + "a")
                                    if v3 {
                                        var v67 string = v66
                                        _ = v67
                                    } else {
                                        var v68 uint8 = 169
                                        v66 = ""
                                        _ = v68
                                    }
                                    delete(v2, genFn0(v50, v66))
                                    _ = v66
                                }
                                _ = v58
                            } else {
                                var v69 uint = v50
                                v69 = (v69 << ((v69) & 63))
                                v30 -= v30
                                _ = v69
                            }
                        }
                        if ((0.1 * v4) <= 1e10) {
                            v70 := GenS9{F0: 625, F1: 42, F2: 127}
                            v3 = (!v3)
                            var v71 float32 = 1e10
                            _ = v70
                            _ = v71
                        }
                        var v72 uint = 65535
                        var v73 uint16 = 776
                        _ = fi52
                        _ = v72
                        _ = v73
                    }
                    v74 := []int{(fi51 & (v1 % ((v1) | 1))), 127, (fi51 >> ((42) & 63)), (v1 + fi51)}
                    _ = fi51
                    _ = v74
                }
                if (41 > 128) {
                    for fi75 := 0; fi75 < 4; fi75++ {
                        delete(v2, 1)
                        fmt.Println(v50, v4)
                        _ = fi75
                    }
                    var v76 int8 = -128
                    _ = v76
                } else {
                    for fi77 := 0; fi77 < 7; fi77++ {
                        var v78 int16 = (-(v30))
                        if v3 {
                            if (-42 >= 1) {
                                for fi79 := 0; fi79 < 13; fi79++ {
                                    v80 := GenS2{F0: "a"}
                                    _ = fi79
                                    _ = v80
                                }
                                switch "a" {
                                case "\n\t":
                                    var v81 float64 = v4
                                    _ = v81
                                case "hello":
                                    for fi82 := 0; fi82 < 19; fi82++ {
                                        if (v30 <= 32767) {
                                            v1 -= (fi77 * fi82)
                                            v83 := GenS3{F0: v78, F1: 245, F2: (v50 | (v50 << ((v50) & 63)))}
                                            _ = v83
                                        }
                                        v84 := GenS10{F0: v50}
                                        _ = fi82
                                        _ = v84
                                    }
                                }
                            }
                            v85 := []int{127, (-(v1)), v1, 209}
                            if true {
                                v86 := GenS8{F0: v1, F1: "hello"}
                                var v87 int8 = 100
                                if len(v85) >= 2 {
                                    v85 = v85[1:len(v85)]
                                }
                                _ = v86
                                _ = v87
                            }
                            _ = v85
                        } else {
                            var v88 float64 = ((0.1 + ((v4 - v4) / v4)) + v4)
                            var v89 float64 = 0.0
                            _ = v88
                            _ = v89
                        }
                        _ = fi77
                        _ = v78
                    }
                    v90 := []int{}
                    _ = v90
                }
                for fi91 := 0; fi91 < 1; fi91++ {
                    for fi92 := 0; fi92 < 7; fi92++ {
                        var v93 uint64 = 1
                        switch fi92 {
                        case 1:
                            fmt.Println(v93, v3, v30)
                        default:
                            v94 := GenS11{F0: 2}
                            _ = v94
                        }
                        var v95 uint = v50
                        _ = fi92
                        _ = v93
                        _ = v95
                    }
                    var v96 int16 = -32768
                    _ = fi91
                    _ = v96
                }
            } else {
                var v97 uint = (v50 << ((v50) & 63))
                v98 := GenS9{F0: 42, F1: -1, F2: 127}
                v50 += (v97 / ((v97) | 1))
                _ = v97
                _ = v98
            }
            v99 := []int{v1, -863, -2147483648}
            _ = v99
        } else {
            v30 = v30
        }
        var v100 int16 = v30
        v50 = (v50 + (v50 | v50))
        _ = v100
    }
    v50 = v50
    v4 = (v4 - (v4 / ((v4 * float64((-2.25 + -2.25))) * 1.5)))
    delete(v2, genFn0(v50, "a"))
    v101 := GenS6{F0: -626, F1: (true && ((true && (0.1 > (-2.25 * 1e10))) || true)), F2: 347}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v50
}
