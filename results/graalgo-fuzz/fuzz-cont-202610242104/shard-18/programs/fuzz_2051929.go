package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint8; F1 int32; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 int64; F2 int }
type GenS8 struct { F0 int32; F1 uint8 }
func genFn0(p0 int, p1 uint) uint64 {
    var v6 int8 = 1
    v7 := GenS1{F0: 1}
    var v8 uint = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32) uint8 {
    p0 = 2147483647
    v9 := GenS1{F0: -1}
    for fi10 := 0; fi10 < 18; fi10++ {
        switch (fi10 | (^(fi10))) {
        case 42:
            v11 := []int{fi10}
            _ = v11
        case -9223372036854775808:
            p0 = (p0 & p0)
        case 128:
            var v12 uint8 = 0
            _ = v12
        }
        _ = fi10
    }
    var v13 bool = true
    for fi14 := 0; fi14 < 17; fi14++ {
        if (false || (100 != 100)) {
            fmt.Println(v13, v13)
            for fi15 := 0; fi15 < 3; fi15++ {
                v13 = v13
                if (float32(3.14159) != 3.14159) {
                    p0 = p0
                } else {
                    p0 = 0
                    fmt.Println(v13, v13, fi15)
                    fmt.Println(fi15, fi14, fi15)
                }
                if (fi14 < (fi14 | -662)) {
                    v16 := GenS1{F0: -9223372036854775808}
                    p0 = p0
                    _ = v16
                }
                _ = fi15
            }
        } else {
            var v17 int = fi14
            if ((0.1 * 3.14159) > (0.1 + (1e10 * (0.0 * -2.25)))) {
                v18 := []int{2147483647, 4294967295, 9223372036854775807, 127}
                v13 = (false || true)
                _ = len(v18)
                _ = v18
            }
            _ = v17
        }
        var v19 float64 = 0.1
        _ = fi14
        _ = v19
    }
    _ = p0
    _ = v13
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{0}
    var v3 bool = false
    v4 := []string{"a", "foo bar", "a", "hello"}
    v5 := []uint32{2147483647, 692, 1, 2}
    switch ("a" + "hello") {
    case "hello":
        var v20 uint = 4294967295
        _ = v20
    case "foo bar":
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                switch "a" {
                case "\n\t":
                    if ((false && ((((32767 > 737) || (("foo bar" + "foo bar") != ("a" + "a"))) && (2 > 1)) || v3)) || (!(v3 && (434 > 4294967295)))) {
                        v23 := v1["hello"]
                        _ = v23
                    }
                case "hello":
                    var v24 uint32 = 588
                    _ = v24
                }
                _ = len(v4)
                _ = fi22
            }
            v25 := GenS1{F0: 42}
            _ = fi21
            _ = v25
        }
    default:
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 uint8 = 0
            _ = fi26
            _ = v27
        }
    }
    v3 = v3
    if (27 != -1) {
        if (2147483648 > 42) {
            for fi28 := 0; fi28 < 13; fi28++ {
                if (-1 < -9223372036854775808) {
                    if (!v3) {
                        var v29 float32 = (3.14159 / (0.1 + -2.25))
                        _ = v29
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (-32768 != 42) {
                            if v3 {
                                fmt.Println(fi28, fi28)
                                if ((255 >= 1) || v3) {
                                    var v30 uint = 4294967295
                                    v3 = v3
                                    v31 := GenS2{F0: v30, F1: 233}
                                    _ = v30
                                    _ = v31
                                } else {
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        if (!(v3 && v3)) {
                                            var v33 uint32 = 331
                                            v34 := GenS1{F0: 119}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v33
                                            _ = v34
                                        }
                                        fmt.Println(v3, fi32, v3)
                                        _ = fi32
                                    }
                                    switch "hello" {
                                    case "hello":
                                        for fi35 := 0; fi35 < 19; fi35++ {
                                            var v36 int = (fi28 << ((255) & 63))
                                            var v37 int32 = 1
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                        }
                                    case "":
                                        for fi38 := 0; fi38 < 19; fi38++ {
                                            if v3 {
                                                if true {
                                                    delete(v1, ("\n\t" + ""))
                                                    switch (fi28 << ((4294967295) & 63)) {
                                                    case 420:
                                                        var v39 bool = false
                                                        if v39 {
                                                            v39 = v39
                                                            v39 = false
                                                        }
                                                        _ = v39
                                                    case 32768:
                                                        v40 := v1["hello"]
                                                        _ = v40
                                                    }
                                                    var v41 int = (fi28 >> ((219) & 63))
                                                    _ = v41
                                                } else {
                                                    if v3 {
                                                        switch fi28 {
                                                        case 687:
                                                            var v42 bool = ((42 <= 100) && true)
                                                            _ = v42
                                                        case -9223372036854775808:
                                                            fmt.Println(v3, fi38)
                                                        }
                                                        var v43 int16 = 32767
                                                        var v44 string = "foo bar"
                                                        v45 := GenS2{F0: 604, F1: -1}
                                                        v43 = 32767
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        var v46 bool = (2 != 821)
                                                        var v47 uint64 = 0
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    v3 = true
                                                }
                                                v48 := GenS1{F0: -560}
                                                var v49 int = (fi38 ^ -1)
                                                _ = v48
                                                _ = v49
                                            } else {
                                                _ = len(v1)
                                                if false {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v50 := GenS2{F0: 1, F1: 353}
                                                    _ = v50
                                                } else {
                                                    var v51 uint8 = genFn1(723)
                                                    _ = v51
                                                }
                                            }
                                            v52 := GenS3{F0: 265, F1: 173}
                                            _ = fi38
                                            _ = v52
                                        }
                                    }
                                    v4 = append(v4, (("hello" + ("" + "hello")) + ""))
                                }
                            }
                            for fi53 := 0; fi53 < 2; fi53++ {
                                if true {
                                    v3 = (v3 || ((2 > 1) || v3))
                                    var v54 uint64 = 42
                                    fmt.Println(v54, fi28)
                                    _ = v54
                                }
                                if (42 < (fi28 % ((fi28) | 1))) {
                                    v55 := v1["hello"]
                                    v3 = false
                                    fmt.Println(fi28, v55)
                                    v56 := v1[v55]
                                    _ = v55
                                    _ = v56
                                } else {
                                    switch fi53 {
                                    case 9223372036854775807:
                                        v3 = v3
                                    default:
                                        if ((fi53 >= (fi53 & (fi28 + fi53))) || (v3 && (100 >= 127))) {
                                            if ((true && (131 <= 2)) || false) {
                                                v3 = (!v3)
                                                v4 = append(v4, ("a" + "a"))
                                            }
                                            if v3 {
                                                if v3 {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v3 = ("" >= "a")
                                                    if true {
                                                        var v57 float32 = (-2.25 + 0.1)
                                                        v3 = (v3 || false)
                                                        _ = v57
                                                    }
                                                } else {
                                                    v3 = v3
                                                    v58 := GenS2{F0: 4294967295, F1: 9223372036854775807}
                                                    _ = v58
                                                }
                                            }
                                            var v59 float32 = (1.5 + float32(1e10))
                                            _ = v59
                                        } else {
                                            var v60 uint16 = 255
                                            fmt.Println(v60, v3, v60)
                                            _ = v60
                                        }
                                    }
                                }
                                var v61 uint8 = 1
                                _ = fi53
                                _ = v61
                            }
                        }
                        for fi62 := 0; fi62 < 8; fi62++ {
                            switch -1 {
                            case 32768:
                                if v3 {
                                    if v3 {
                                        v1[("a" + "a")] = "a"
                                        if (255 != 68) {
                                            if v3 {
                                                v63 := GenS4{F0: 255, F1: 760}
                                                switch (fi62 + fi28) {
                                                case -552:
                                                    v3 = v3
                                                case -9223372036854775808:
                                                    v64 := GenS3{F0: 102, F1: 42}
                                                    _ = v64
                                                case 2147483647:
                                                    if v3 {
                                                        v3 = (2 >= 4294967295)
                                                    }
                                                default:
                                                    fmt.Println(v3, fi28)
                                                    if (0 == 4294967295) {
                                                        fmt.Println(fi62)
                                                        var v65 int32 = 42
                                                        _ = v65
                                                    } else {
                                                        if true {
                                                            if false {
                                                                v66 := GenS5{F0: 42, F1: 0, F2: 2147483647}
                                                                var v67 uint32 = 2
                                                                if ("hello" <= "foo bar") {
                                                                    fmt.Println(v3, v3, v67)
                                                                    var v68 float64 = 1.5
                                                                    _ = v68
                                                                } else {
                                                                    if (0 != 645) {
                                                                        if (v3 || (!false)) {
                                                                            if v3 {
                                                                                v67 -= v67
                                                                                switch "" {
                                                                                case "a":
                                                                                    v3 = v3
                                                                                case "foo bar":
                                                                                    var v69 bool = (v3 || v3)
                                                                                    _ = v69
                                                                                default:
                                                                                    _ = len(v1)
                                                                                }
                                                                            }
                                                                            var v70 bool = (0 < 42)
                                                                            fmt.Println(fi28, v70, v3)
                                                                            _ = v70
                                                                        } else {
                                                                            var v71 uint8 = 127
                                                                            fmt.Println(v3)
                                                                            _ = v71
                                                                        }
                                                                        var v72 int = (fi28 + -510)
                                                                        v73 := GenS4{F0: 0, F1: 255}
                                                                        v1[""] = "a"
                                                                        v74 := GenS6{F0: 42}
                                                                        _ = v72
                                                                        _ = v73
                                                                        _ = v74
                                                                    } else {
                                                                        v67 += (v67 + (v67 + (v67 ^ 0)))
                                                                        fmt.Println(v3)
                                                                    }
                                                                    v3 = (v3 && v3)
                                                                }
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                v3 = v3
                                                                if v3 {
                                                                    if v3 {
                                                                        if true {
                                                                            if v3 {
                                                                                var v75 float64 = 0.1
                                                                                var v76 bool = (!true)
                                                                                var v77 uint64 = 4294967295
                                                                                v78 := GenS3{F0: 42, F1: 42}
                                                                                var v79 int8 = 1
                                                                                _ = v75
                                                                                _ = v76
                                                                                _ = v77
                                                                                _ = v78
                                                                                _ = v79
                                                                            } else {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                fmt.Println(fi28, v3, fi28)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v80 := GenS7{F0: 42, F1: 9223372036854775807, F2: (fi28 - (fi62 / ((fi62) | 1)))}
                                                                        _ = v80
                                                                    }
                                                                }
                                                                _ = len(v1)
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            var v81 uint8 = 1
                                                            _ = v81
                                                        }
                                                    }
                                                }
                                                switch fi28 {
                                                case 268:
                                                    var v82 string = ("\n\t" + "\n\t")
                                                    _ = v82
                                                case 127:
                                                    delete(v1, "foo bar")
                                                case -2147483648:
                                                    if (9223372036854775807 < 27) {
                                                        v83 := GenS2{F0: 576, F1: 1}
                                                        fmt.Println(fi28, v3, fi28)
                                                        v84 := v1["\n\t"]
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                default:
                                                    v3 = true
                                                }
                                                _ = v63
                                            } else {
                                                switch ("" + "") {
                                                case "a":
                                                    if v3 {
                                                        v3 = (1e10 <= (0.0 * 3.14159))
                                                    }
                                                    v3 = (!((-128 > -103) && (!true)))
                                                    if (!(!v3)) {
                                                        v85 := GenS7{F0: 127, F1: 0, F2: (-(fi62))}
                                                        var v86 float32 = ((1.5 + ((1.5 - 0.0) * 3.14159)) / 1.0)
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                case "\n\t":
                                                    fmt.Println(fi28, v3)
                                                case "foo bar":
                                                    var v87 float32 = -2.25
                                                    _ = v87
                                                }
                                            }
                                            switch -9223372036854775808 {
                                            case 128:
                                                fmt.Println(v3, fi62)
                                            case -2147483648:
                                                v3 = (((fi28 >> ((382) & 63)) < 853) && (0.0 > 1e10))
                                            case 32768:
                                                var v88 uint8 = 102
                                                _ = v88
                                            }
                                        } else {
                                            fmt.Println(v3)
                                        }
                                    }
                                    fmt.Println(fi62)
                                    v3 = v3
                                }
                            case -9223372036854775808:
                                if (true && (128 > 128)) {
                                    fmt.Println(fi62, fi28)
                                    var v89 uint64 = 4294967295
                                    v89 += v89
                                    _ = v89
                                }
                            case -1:
                                var v90 float64 = -2.25
                                _ = v90
                            }
                            _ = fi62
                        }
                    }
                    delete(v1, "hello")
                    var v91 uint8 = 255
                    _ = v91
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi28
            }
        }
        v92 := GenS4{F0: 1, F1: 0}
        _ = v92
    } else {
        v3 = (true || (1 == 9223372036854775807))
        v93 := GenS6{F0: -129}
        _ = v93
    }
    fmt.Println(v3)
    if v3 {
        v94 := GenS8{F0: -1, F1: 255}
        switch "hello" {
        case "":
            v3 = v3
        case "foo bar":
            v95 := GenS1{F0: 973}
            _ = v95
        }
        _ = v94
    } else {
        var v96 float64 = 0.0
        var v97 uint8 = 2
        _ = v96
        _ = v97
    }
    v3 = (32768 <= 2147483647)
    if (genFn1(382) >= 42) {
        v3 = false
        switch 0 {
        case -807:
            v1[("foo bar" + "hello")] = ("a" + ("a" + ""))
        case -871:
            v3 = (-1 == 137)
        case 75:
            var v98 float64 = (1.5 / -2.25)
            _ = v98
        }
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
