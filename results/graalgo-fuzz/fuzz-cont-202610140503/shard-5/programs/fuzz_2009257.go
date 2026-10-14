package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint64; F1 int }
func genFn0() int16 {
    switch "a" {
    case "hello":
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := GenS1{F0: ""}
            var v7 int16 = 142
            for fi8 := 0; fi8 < 14; fi8++ {
                v7 = -786
                if (1 < 255) {
                    if false {
                        if true {
                            v7 += (v7 % ((605) | 1))
                            fmt.Println(v7, fi8)
                            v9 := GenS1{F0: ("hello" + (("foo bar" + "") + "a"))}
                            var v10 float32 = 1.5
                            var v11 bool = true
                            _ = v10
                            _ = v11
                            _ = v9
                        } else {
                            var v12 string = ""
                            v13 := []int{}
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: ""}
                        _ = v14
                    }
                    v7 = (v7 * -32768)
                }
                _ = fi8
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    case "":
        v15 := GenS2{F0: 4294967295, F1: 2147483648}
        _ = v15
    case "\n\t":
        if false {
            var v16 float64 = 0.1
            switch "foo bar" {
            case "hello":
                for fi17 := 0; fi17 < 2; fi17++ {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        fmt.Println(fi18, fi18)
                        if true {
                            v19 := []int{}
                            _ = v19
                        } else {
                            var v20 int32 = 2147483647
                            _ = v20
                        }
                        if (-1 <= 0) {
                            v21 := GenS1{F0: ""}
                            if ((1e10 / 3.14159) > v16) {
                                fmt.Println(v16, v16, v16)
                                v16 = (float64(1.5) * 1.5)
                            }
                            _ = v21
                        } else {
                            switch (-(fi17)) {
                            case 2147483647:
                                v22 := []int{}
                                _ = v22
                            case 0:
                                v23 := GenS1{F0: ((("hello" + "\n\t") + "foo bar") + "\n\t")}
                                _ = v23
                            }
                            fmt.Println(fi17, fi18)
                            if true {
                                var v24 uint8 = 2
                                _ = v24
                            } else {
                                var v25 int32 = -1
                                if true {
                                    if (0.0 != (v16 - v16)) {
                                        var v26 uint16 = 8
                                        switch "" {
                                        case "\n\t":
                                            if (!(v26 != 867)) {
                                                v27 := GenS2{F0: 1, F1: 438}
                                                if false {
                                                    var v28 int8 = 127
                                                    _ = v28
                                                } else {
                                                    if true {
                                                        var v29 uint64 = 1
                                                        v30 := GenS2{F0: 4294967295, F1: 305}
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    v31 := GenS1{F0: "\n\t"}
                                                    _ = v31
                                                }
                                                _ = v27
                                            } else {
                                                v32 := GenS2{F0: 4294967295, F1: 4294967295}
                                                _ = v32
                                            }
                                        case "":
                                            var v33 float64 = ((((v16 + (v16 + v16)) * (3.14159 - v16)) * (float64(((-2.25 - 0.1) + ((-2.25 / (1.5 / 1.0)) - (3.14159 + 0.1)))) / v16)) / (0.0 / (0.1 * v16)))
                                            _ = v33
                                        case "hello":
                                            var v34 uint64 = uint64(v26)
                                            _ = v34
                                        }
                                        _ = v26
                                    }
                                } else {
                                    v35 := GenS3{F0: true}
                                    if (v25 > v25) {
                                        var v36 uint64 = 42
                                        var v37 uint32 = 65535
                                        var v38 uint64 = 380
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v39 := GenS4{F0: 127, F1: 1}
                                        v25 = (v25 << ((18446744073709551615) & 63))
                                        _ = v39
                                    }
                                    fmt.Println(fi17, v16)
                                    _ = v35
                                }
                                fmt.Println(v16)
                                _ = v25
                            }
                            v40 := GenS4{F0: 127, F1: 9223372036854775807}
                            _ = v40
                        }
                        v41 := []int{565, -9223372036854775808, (fi18 << ((499) & 63))}
                        _ = len(v41)
                        v42 := GenS1{F0: "hello"}
                        _ = fi18
                        _ = v41
                        _ = v42
                    }
                    var v43 bool = true
                    _ = fi17
                    _ = v43
                }
            case "a":
                if true {
                    var v44 uint = 65535
                    _ = v44
                }
            case "\n\t":
                var v45 uint16 = 641
                _ = v45
            }
            var v46 int32 = 2147483647
            _ = v16
            _ = v46
        } else {
            var v47 int64 = 42
            _ = v47
        }
    default:
        v48 := []int{32767}
        _ = v48
    }
    var v49 bool = (973 > 42)
    _ = v49
    return 32767
}
func genFn1(p0 int, p1 uint32, p2 int16) uint8 {
    p2 = (p2 >> ((65535) & 63))
    switch (("foo bar" + "") + "\n\t") {
    case "hello":
        p2 = (p2 / ((genFn0()) | 1))
    case "":
        p1 = (p1 >> ((355) & 63))
    case "a":
        v50 := GenS2{F0: 133, F1: p1}
        _ = v50
    default:
        switch (p0 << ((805) & 63)) {
        case 386:
            p2 -= p2
        case -191:
            p1 = p1
        }
    }
    v51 := []int{p0}
    switch p0 {
    case -96:
        var v52 string = ("" + "\n\t")
        _ = v52
    default:
        for fi53 := 0; fi53 < 11; fi53++ {
            var v54 string = ""
            var v55 uint8 = 129
            for fi56 := 0; fi56 < 8; fi56++ {
                p1 = (p1 + 699)
                v57 := GenS5{F0: v54, F1: 0}
                _ = len(v51)
                _ = fi56
                _ = v57
            }
            _ = fi53
            _ = v54
            _ = v55
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 bool = (982 > 66)
    v58 := GenS6{F0: 2147483648}
    _ = len(v3)
    v59 := v3[2147483648]
    var v60 uint32 = 2
    v61 := GenS7{F0: ((("foo bar" + "") + "foo bar") + "a"), F1: (v60 & v60), F2: v60}
    v1 += v1
    v62 := GenS6{F0: v60}
    if v4 {
        var v63 uint = 654
        var v64 uint16 = 255
        var v65 int = 154
        _ = v63
        _ = v64
        _ = v65
    } else {
        if (0 == 496) {
            var v66 uint32 = v60
            _ = v66
        }
        fmt.Println(v59, v1)
    }
    v67 := GenS2{F0: 235, F1: 0}
    var v68 float64 = (1e10 + 1e10)
    if v4 {
        fmt.Println(v59)
        var v69 int = v59
        _ = v69
    } else {
        for fi70 := 0; fi70 < 4; fi70++ {
            v3[(v60 << ((2) & 63))] = (fi70 << ((2) & 63))
            if true {
                var v71 int = fi70
                for fi72 := 0; fi72 < 3; fi72++ {
                    var v73 uint = 1
                    v73 = 2
                    v1 = (v1 & v1)
                    _ = fi72
                    _ = v73
                }
                var v74 int8 = 0
                _ = v71
                _ = v74
            }
            for fi75 := 0; fi75 < 11; fi75++ {
                var v76 string = "foo bar"
                switch v76 {
                case "hello":
                    var v77 bool = false
                    _ = v77
                case "a":
                    fmt.Println(v68)
                }
                if true {
                    var v78 int32 = -1
                    _ = v78
                }
                v79 := []int{(fi70 << ((1) & 63))}
                v80 := GenS5{F0: v76, F1: -128}
                var v81 uint64 = v1
                v4 = (!false)
                if ((v76 <= v76) && ((!(127 > -9223372036854775808)) && (uint8(v60) <= 255))) {
                    v68 += v68
                } else {
                    var v82 uint64 = 42
                    v79 = append(v79, -493)
                    delete(v3, v60)
                    switch "hello" {
                    case "\n\t":
                        var v83 float64 = v68
                        _ = v83
                    case "":
                        switch ("hello" + v76) {
                        case "\n\t":
                            var v84 uint = 18446744073709551615
                            _ = v84
                        case "a":
                            fmt.Println(v81, v68)
                        case "hello":
                            var v85 uint = 2
                            _ = v85
                        }
                    case "hello":
                        switch (v59 << ((545) & 63)) {
                        case 1:
                            var v86 uint8 = 0
                            _ = v86
                        default:
                            var v87 int = fi75
                            _ = v87
                        }
                    default:
                        _ = len(v2)
                    }
                    switch (v76 + v76) {
                    case "a":
                        if true {
                            delete(v2, v76)
                            var v88 int16 = 32767
                            _ = v88
                        }
                    case "":
                        switch fi75 {
                        case 260:
                            switch v76 {
                            case "foo bar":
                                v60 = (v60 - (^(v60)))
                                switch v76 {
                                case "a":
                                    fmt.Println(v4, v76, v1)
                                case "":
                                    v82 -= 694
                                    if v4 {
                                        v79 = append(v79, fi70)
                                    } else {
                                        v89 := GenS7{F0: ("hello" + (v76 + (v76 + v76))), F1: 2147483648, F2: v60}
                                        var v90 float32 = 1.5
                                        fmt.Println(v82, v60, v60)
                                        _ = v89
                                        _ = v90
                                    }
                                case "\n\t":
                                    var v91 uint64 = 591
                                    _ = v91
                                }
                                if false {
                                    var v92 uint64 = v82
                                    delete(v2, ((v76 + v76) + v76))
                                    _ = v92
                                } else {
                                    fmt.Println(v59, v81, fi70)
                                    if len(v79) >= 2 {
                                        v79 = v79[1:len(v79)]
                                    }
                                }
                            case "\n\t":
                                v79 = append(v79, fi70)
                            default:
                                var v93 string = v76
                                _ = v93
                            }
                        default:
                            var v94 int = fi75
                            _ = v94
                        }
                    case "\n\t":
                        fmt.Println(v60, v4, v60)
                    default:
                        if (255 < 65535) {
                            if v4 {
                                v3[v60] = fi70
                                v95 := GenS3{F0: (286 <= 246)}
                                fmt.Println(fi70, fi70)
                                if v4 {
                                    v96 := GenS8{F0: v60, F1: 0, F2: 2}
                                    v1 = uint64(v60)
                                    v97 := GenS7{F0: v76, F1: (^(v60)), F2: 1}
                                    _ = v96
                                    _ = v97
                                }
                                v79 = append(v79, (fi70 / (((fi75 % (((-(v59))) | 1))) | 1)))
                                _ = v95
                            } else {
                                switch -235 {
                                case 384:
                                    v98 := GenS9{F0: v82, F1: 9223372036854775807}
                                    _ = v98
                                default:
                                    fmt.Println(v68, fi75, v1)
                                }
                            }
                            v59 = (-(v59))
                            fmt.Println(v1)
                            v59 = 440
                        }
                    }
                    _ = v82
                }
                _ = len(v79)
                delete(v2, ((v76 + v76) + (v76 + v76)))
                _ = fi75
                _ = v76
                _ = v79
                _ = v80
                _ = v81
            }
            _ = fi70
        }
        var v99 uint64 = v1
        v100 := []int{v59}
        _ = v100
        _ = v99
    }
    if (65535 == 65535) {
        fmt.Println(v59)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v67
    _ = v68
}
