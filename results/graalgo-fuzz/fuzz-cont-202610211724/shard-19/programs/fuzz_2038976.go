package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS7 struct { F0 int; F1 bool; F2 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 uint8; F2 int8 }
type GenS10 struct { F0 string; F1 int64; F2 uint }
type GenS11 struct { F0 int32 }
func genFn0(p0 uint8, p1 int64, p2 uint16) uint {
    p1 = -9223372036854775808
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 1}
        v7 := []int{fi5}
        p2 += p2
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 4294967295}
    p1 -= p1
    var v9 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint8, p1 int32) int64 {
    switch 0 {
    case 32767:
        var v10 bool = false
        _ = v10
    case 4294967295:
        p1 -= p1
    case 675:
        v11 := []int{}
        _ = v11
    }
    p1 -= p1
    switch "foo bar" {
    case "foo bar":
        v12 := GenS2{F0: 4294967295, F1: -840}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 1; fi13++ {
            switch (("foo bar" + "") + "foo bar") {
            case "hello":
                if false {
                    v14 := GenS3{F0: 42, F1: 1, F2: 0}
                    switch 32767 {
                    case 128:
                        fmt.Println(fi13)
                    }
                    v15 := GenS4{F0: 0, F1: "\n\t"}
                    _ = v14
                    _ = v15
                }
            }
            if (true && true) {
                if (true || (true && (-1 != -43))) {
                    fmt.Println(p1)
                    if false {
                        v16 := GenS5{F0: 0, F1: 503}
                        v17 := GenS6{F0: 1, F1: p0, F2: 965}
                        _ = v16
                        _ = v17
                    }
                } else {
                    var v18 bool = false
                    v19 := []int{-840, (fi13 >> ((18446744073709551615) & 63)), fi13}
                    _ = v18
                    _ = v19
                }
                p0 -= (p0 >> ((18446744073709551615) & 63))
            } else {
                v20 := []int{128, (fi13 >> ((genFn0(p0, 9223372036854775807, 0)) & 63))}
                p1 = (^(p1))
                _ = v20
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                p1 = (p1 + p1)
                _ = fi21
            }
            _ = fi13
        }
    case "hello":
        v22 := GenS2{F0: -1, F1: 1000}
        _ = v22
    default:
        if false {
            if (0.1 < 3.14159) {
                if true {
                    fmt.Println(p0, p1)
                } else {
                    p0 -= 21
                }
            }
            fmt.Println(p1, p1)
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := []int{fi23, fi23, fi23, fi23}
                fmt.Println(fi23, p1)
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := []int{fi25, fi25, fi25, fi25}
                for fi27 := 0; fi27 < 15; fi27++ {
                    _ = len(v26)
                    p0 += p0
                    _ = fi27
                }
                _ = fi25
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint8 = 127
    var v4 int16 = 1000
    fmt.Println(v3, v3, v4)
    if (v4 >= 334) {
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := GenS4{F0: 255, F1: "hello"}
            var v30 uint8 = (v3 ^ v3)
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    var v31 bool = v1
    if (!v1) {
        switch "a" {
        case "\n\t":
            if (float32((3.14159 - 1e10)) <= 3.14159) {
                v4 = v4
                v32 := GenS7{F0: 0, F1: (!true), F2: 805}
                v31 = ((!(128 > v3)) && v31)
                _ = v32
            } else {
                fmt.Println(v1)
                v1 = false
                v1 = v1
            }
        case "hello":
            for fi33 := 0; fi33 < 5; fi33++ {
                if (!false) {
                    v31 = ((!(!true)) || v31)
                }
                _ = fi33
            }
        case "":
            var v34 uint32 = 65535
            _ = v34
        default:
            for fi35 := 0; fi35 < 15; fi35++ {
                switch ("foo bar" + "a") {
                case "":
                    if true {
                        var v36 uint32 = 2
                        _ = v36
                    } else {
                        var v37 uint16 = 1
                        for fi38 := 0; fi38 < 7; fi38++ {
                            v1 = (((!(!(42 <= (v37 + (v37 % ((v37) | 1)))))) && true) || false)
                            v1 = true
                            v39 := GenS8{F0: 127}
                            _ = fi38
                            _ = v39
                        }
                        fmt.Println(v37)
                        _ = v37
                    }
                case "\n\t":
                    switch ("foo bar" + "") {
                    case "\n\t":
                        if (42 > 368) {
                            var v40 uint64 = 0
                            _ = v40
                        }
                    case "a":
                        for fi41 := 0; fi41 < 12; fi41++ {
                            if v1 {
                                if v1 {
                                    var v42 string = (("" + "") + "\n\t")
                                    _ = v42
                                }
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    v31 = true
                                case "a":
                                    var v43 int64 = 318
                                    _ = v43
                                case "hello":
                                    if (true && (((100 > 72) || ((18446744073709551615 >= 2) && v1)) || true)) {
                                        if ((!(-2.25 == 0.1)) || (255 <= 255)) {
                                            var v44 uint64 = 4294967295
                                            var v45 uint16 = 658
                                            _ = v44
                                            _ = v45
                                        } else {
                                            if true {
                                                v2["\n\t"] = (fi41 & 32768)
                                                if v1 {
                                                    var v46 int16 = v4
                                                    v3 = v3
                                                    v47 := GenS6{F0: -110, F1: v3, F2: 618}
                                                    _ = v46
                                                    _ = v47
                                                }
                                            }
                                            var v48 uint64 = 2
                                            _ = v48
                                        }
                                        if v1 {
                                            fmt.Println(fi41, fi41)
                                            v1 = v1
                                            var v49 uint8 = v3
                                            _ = v49
                                        } else {
                                            v50 := GenS5{F0: 42, F1: (fi35 - fi35)}
                                            _ = v50
                                        }
                                        var v51 uint16 = 1
                                        var v52 uint16 = (v51 ^ 0)
                                        fmt.Println(v31, v52, v4)
                                        _ = v51
                                        _ = v52
                                    } else {
                                        v53 := GenS2{F0: (fi41 ^ fi35), F1: v4}
                                        _ = v53
                                    }
                                default:
                                    v54 := v2[("\n\t" + ("foo bar" + "a"))]
                                    _ = v54
                                }
                                v55 := GenS9{F0: (-(v4)), F1: v3, F2: -116}
                                fmt.Println(v31, fi41, fi41)
                                _ = v55
                            }
                            _ = fi41
                        }
                    case "hello":
                        if v31 {
                            v56 := []int{fi35, (fi35 * fi35), (fi35 ^ fi35), -2147483648}
                            var v57 int16 = (v4 ^ 361)
                            fmt.Println(v31, v57, v31)
                            _ = v56
                            _ = v57
                        } else {
                            v3 = v3
                            v31 = (("a" + "") > ("foo bar" + "\n\t"))
                            if (2 < 42) {
                                if v1 {
                                    switch (fi35 ^ (^(fi35))) {
                                    case -404:
                                        if (1 <= 255) {
                                            v4 = v4
                                            v2["a"] = -9223372036854775808
                                        }
                                    }
                                    if v31 {
                                        v31 = false
                                        var v58 int64 = 569
                                        _ = v58
                                    } else {
                                        var v59 string = "a"
                                        if (42 <= genFn0((v3 ^ (v3 % ((v3) | 1))), genFn1(v3, -1), 2)) {
                                            var v60 int64 = 42
                                            fmt.Println(v4, v1)
                                            _ = v60
                                        }
                                        fmt.Println(v59, v1, v4)
                                        _ = v59
                                    }
                                    var v61 float64 = 0.1
                                    _ = v61
                                } else {
                                    for fi62 := 0; fi62 < 6; fi62++ {
                                        v31 = (1 > 0)
                                        _ = fi62
                                    }
                                    if v31 {
                                        v31 = (v1 && v31)
                                        if true {
                                            v63 := GenS10{F0: ("foo bar" + "hello"), F1: 1, F2: 1}
                                            v64 := GenS3{F0: (v4 << ((634) & 63)), F1: genFn0(107, 830, 42), F2: 42}
                                            var v65 uint8 = (v3 << ((genFn0(42, 506, 65535)) & 63))
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                    } else {
                                        v66 := v2["hello"]
                                        _ = v66
                                    }
                                }
                                var v67 int16 = v4
                                v68 := GenS5{F0: 18446744073709551615, F1: 32767}
                                _ = v67
                                _ = v68
                            } else {
                                v69 := GenS11{F0: -2147483648}
                                _ = v69
                            }
                        }
                    default:
                        v70 := v2["foo bar"]
                        _ = v70
                    }
                default:
                    var v71 uint16 = 32767
                    _ = v71
                }
                v1 = (128 > v3)
                _ = fi35
            }
        }
    }
    v3 = 216
    var v72 uint16 = 218
    for fi73 := 0; fi73 < 7; fi73++ {
        var v74 uint64 = 42
        var v75 float32 = 3.14159
        if v31 {
            fmt.Println(fi73)
        } else {
            switch 9223372036854775807 {
            case 1:
                var v76 uint64 = 292
                _ = v76
            case 0:
                switch "a" {
                case "":
                    if v31 {
                        fmt.Println(v75)
                        var v77 string = ("" + ("a" + "foo bar"))
                        if v1 {
                            var v78 int64 = -917
                            v79 := GenS2{F0: fi73, F1: v4}
                            v4 += (v4 >> ((258) & 63))
                            _ = v78
                            _ = v79
                        } else {
                            if true {
                                for fi80 := 0; fi80 < 2; fi80++ {
                                    var v81 int16 = (v4 / ((32767) | 1))
                                    v82 := []int{fi80, (fi73 % ((42) | 1))}
                                    _ = fi80
                                    _ = v81
                                    _ = v82
                                }
                            } else {
                                v72 += v72
                                v74 = (v74 << ((285) & 63))
                            }
                        }
                        _ = v77
                    }
                case "\n\t":
                    var v83 uint16 = 0
                    _ = v83
                case "hello":
                    switch 4294967295 {
                    case -2147483648:
                        var v84 uint16 = (v72 ^ (v72 ^ 65535))
                        _ = v84
                    case 2147483647:
                        v75 = 0.0
                    }
                default:
                    v31 = ((fi73 << ((42) & 63)) <= fi73)
                }
            case 9223372036854775807:
                fmt.Println(fi73)
            default:
                v1 = (2 < 1)
            }
            var v85 int16 = -299
            var v86 uint64 = v74
            _ = v85
            _ = v86
        }
        _ = fi73
        _ = v74
        _ = v75
    }
    v87 := GenS7{F0: -9223372036854775808, F1: v1, F2: 0}
    fmt.Println(v4)
    if ((v72 & v72) != (v72 % (((v72 % ((65535) | 1))) | 1))) {
        var v88 int = 42
        if true {
            var v89 uint8 = v3
            if (814 > -104) {
                fmt.Println(v31)
            }
            v90 := GenS7{F0: -9223372036854775808, F1: (!((-1 <= 102) || ((13 == (v88 - 0)) && v31))), F2: 2147483647}
            _ = v89
            _ = v90
        } else {
            var v91 uint64 = 515
            var v92 uint16 = v72
            _ = v91
            _ = v92
        }
        switch v88 {
        case 757:
            fmt.Println(v4, v72, v4)
        }
        _ = v88
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v3)
    case "hello":
        v4 = v4
    case "":
        switch 548 {
        case 127:
            var v93 int = -9223372036854775808
            _ = v93
        }
    default:
        v72 += v72
    }
    var v94 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v72
    _ = v87
    _ = v94
}
