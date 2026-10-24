package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 uint; F2 bool }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0() int64 {
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: v6}
        var v9 int16 = -32768
        v10 := []int{(fi7 / ((fi7) | 1)), (fi7 | (fi7 - fi7)), 726}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := []int{2147483647, fi11, 42}
        _ = fi11
        _ = v12
    }
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 int8, p1 uint16, p2 int) int16 {
    fmt.Println(p0, p2, p2)
    if true {
        if false {
            if true {
                var v13 int16 = -32768
                _ = v13
            } else {
                var v14 string = (("" + "hello") + "")
                if (!((p2 << ((18446744073709551615) & 63)) < (p2 * (p2 * 9223372036854775807)))) {
                    var v15 int32 = 1
                    _ = v15
                }
                p0 -= (p0 - p0)
                _ = v14
            }
            p2 -= -1
            v16 := []int{}
            _ = v16
        } else {
            v17 := GenS1{F0: 1}
            v18 := GenS1{F0: 18446744073709551615}
            _ = v17
            _ = v18
        }
        var v19 int = p2
        _ = v19
    } else {
        p0 = 100
        var v20 uint16 = (p1 - p1)
        _ = v20
    }
    v21 := GenS1{F0: uint64(p1)}
    fmt.Println(p2, p0)
    for fi22 := 0; fi22 < 9; fi22++ {
        p2 += fi22
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := []int{(fi22 & fi23), (-(p2)), (fi22 << ((991) & 63))}
            if (112 <= 2) {
                v25 := GenS1{F0: 1}
                p1 = p1
                if true {
                    var v26 float64 = ((1e10 * float64(1e10)) - -2.25)
                    p1 -= (p1 + 32767)
                    _ = v26
                } else {
                    var v27 int = 32767
                    fmt.Println(fi23)
                    _ = v27
                }
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 32767
}

func main() {
    v1 := []uint32{4294967295, 2147483648}
    v2 := make(map[string]int)
    var v3 uint16 = 2
    var v4 int32 = 1
    var v5 int8 = int8(v4)
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 4; fi29++ {
            switch ("a" + "") {
            case "a":
                if true {
                    v30 := v2["\n\t"]
                    v4 = v4
                    v30 -= 0
                    var v31 bool = (v3 == v3)
                    v32 := GenS1{F0: uint64(v3)}
                    _ = v30
                    _ = v31
                    _ = v32
                }
            case "\n\t":
                var v33 uint = 42
                _ = v33
            case "hello":
                v34 := GenS1{F0: 0}
                _ = v34
            default:
                var v35 uint8 = 127
                _ = v35
            }
            _ = fi29
        }
        switch ("hello" + ("\n\t" + (("hello" + ("" + "")) + "foo bar"))) {
        case "\n\t":
            var v36 uint16 = v3
            _ = v36
        case "hello":
            v5 = 100
        default:
            _ = len(v1)
        }
        _ = fi28
    }
    if (1e10 < (1e10 - 1e10)) {
        fmt.Println(v3)
        fmt.Println(v4)
    } else {
        fmt.Println(v4, v5)
    }
    if (!(((-2.25 * 1.5) / ((float64((3.14159 * 1.5)) / float64(-2.25)) + 0.1)) > 3.14159)) {
        if len(v1) > 0 {
            v1[0] = 0
        }
        if false {
            var v37 uint = 0
            v37 -= v37
            _ = v37
        }
        switch 128 {
        case -9223372036854775808:
            if false {
                var v38 string = "hello"
                v38 = (v38 + ((v38 + ("a" + "\n\t")) + v38))
                _ = v38
            } else {
                v4 = v4
                if true {
                    v39 := GenS2{F0: uint32(v3)}
                    _ = len(v1)
                    _ = v39
                } else {
                    var v40 int32 = int32(v5)
                    for fi41 := 0; fi41 < 5; fi41++ {
                        for fi42 := 0; fi42 < 10; fi42++ {
                            v1 = append(v1, 2147483648)
                            _ = fi42
                        }
                        _ = fi41
                    }
                    v5 += v5
                    _ = v40
                }
                v5 = (v5 << ((2) & 63))
            }
        }
    } else {
        if (1 == 2) {
            v2["hello"] = 0
        } else {
            v4 = int32(v5)
            v43 := GenS1{F0: 764}
            v5 += v5
            _ = v43
        }
        var v44 int8 = v5
        _ = v44
    }
    v45 := GenS1{F0: 1}
    var v46 uint = 468
    if (!(1 < -970)) {
        fmt.Println(v4, v4)
        _ = len(v2)
    }
    var v47 uint16 = (v3 >> ((v46) & 63))
    if false {
        for fi48 := 0; fi48 < 16; fi48++ {
            if (42 == 948) {
                for fi49 := 0; fi49 < 6; fi49++ {
                    v50 := GenS3{F0: "\n\t"}
                    _ = fi49
                    _ = v50
                }
                _ = len(v1)
                if false {
                    if (!(-9223372036854775808 == 42)) {
                        var v51 int8 = v5
                        v46 = (v46 & 42)
                        for fi52 := 0; fi52 < 11; fi52++ {
                            v46 += (v46 / ((v46) | 1))
                            var v53 uint16 = v47
                            var v54 int64 = 462
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        _ = v51
                    } else {
                        _ = len(v1)
                    }
                }
            }
            v55 := v2["foo bar"]
            _ = fi48
            _ = v55
        }
        v56 := GenS1{F0: 18446744073709551615}
        _ = v56
    } else {
        if true {
            fmt.Println(v46, v3)
            v57 := v2["foo bar"]
            _ = v57
        } else {
            var v58 uint16 = 1
            for fi59 := 0; fi59 < 20; fi59++ {
                fmt.Println(v3, v47, fi59)
                _ = fi59
            }
            var v60 int8 = (v5 % (((-(v5))) | 1))
            _ = v58
            _ = v60
        }
        var v61 float32 = float32(1e10)
        var v62 int = 2147483647
        _ = v61
        _ = v62
    }
    v47 -= v3
    switch "\n\t" {
    case "foo bar":
        for fi63 := 0; fi63 < 2; fi63++ {
            var v64 uint64 = 4294967295
            v47 = (^(v3))
            _ = fi63
            _ = v64
        }
    case "":
        v65 := GenS2{F0: 2147483648}
        _ = v65
    }
    var v66 int8 = int8(v46)
    var v67 bool = (-930 != -852)
    switch 42 {
    case 300:
        v68 := GenS4{F0: ("\n\t" + "a"), F1: ""}
        _ = v68
    case 1:
        v47 -= 255
    case -806:
        for fi69 := 0; fi69 < 10; fi69++ {
            for fi70 := 0; fi70 < 20; fi70++ {
                v71 := v2["hello"]
                switch "\n\t" {
                case "":
                    if ((1.5 + -2.25) < (float32(1e10) + 3.14159)) {
                        if false {
                            var v72 float64 = (3.14159 + (-2.25 * (1e10 / -2.25)))
                            if v67 {
                                if false {
                                    var v73 string = "foo bar"
                                    v71 = fi70
                                    var v74 float64 = 1e10
                                    _ = v73
                                    _ = v74
                                } else {
                                    v67 = v67
                                    if true {
                                        var v75 int32 = v4
                                        v5 -= (v5 + v5)
                                        _ = v75
                                    } else {
                                        var v76 int32 = (v4 & v4)
                                        _ = v76
                                    }
                                    var v77 uint32 = 2147483647
                                    if (genFn0() >= 9223372036854775807) {
                                        v78 := GenS1{F0: uint64(v46)}
                                        _ = v78
                                    } else {
                                        var v79 float64 = float64(1e10)
                                        _ = v79
                                    }
                                    _ = v77
                                }
                                if (((^(v5)) < (^(v5))) && v67) {
                                    v46 = (v46 ^ v46)
                                }
                                v80 := GenS5{F0: ("" + "a")}
                                _ = v80
                            }
                            var v81 int8 = 42
                            _ = v72
                            _ = v81
                        }
                        v82 := GenS3{F0: "hello"}
                        var v83 int8 = (v5 - v5)
                        _ = v82
                        _ = v83
                    } else {
                        if (!v67) {
                            v1 = append(v1, 0)
                            v84 := GenS3{F0: ("a" + ("\n\t" + "a"))}
                            _ = v84
                        } else {
                            v3 -= 32767
                        }
                        var v85 uint64 = 42
                        _ = v85
                    }
                case "\n\t":
                    v67 = (v67 && false)
                default:
                    var v86 float32 = (3.14159 / -2.25)
                    _ = v86
                }
                var v87 string = ("foo bar" + ((("" + (("\n\t" + "") + ("" + ("\n\t" + "")))) + "") + ""))
                _ = fi70
                _ = v71
                _ = v87
            }
            for fi88 := 0; fi88 < 8; fi88++ {
                v89 := GenS6{F0: ("hello" + (("\n\t" + "\n\t") + ("" + ("" + "foo bar")))), F1: v46, F2: (false || false)}
                v3 = (v3 << ((uint(v47)) & 63))
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                _ = fi88
                _ = v89
            }
            v66 -= v5
            _ = fi69
        }
    }
    v90 := GenS7{F0: 127, F1: (v4 / (((v4 * (v4 * (v4 - 0)))) | 1))}
    v91 := v2["a"]
    for fi92 := 0; fi92 < 2; fi92++ {
        if false {
            for fi93 := 0; fi93 < 11; fi93++ {
                v91 += fi93
                var v94 float32 = -2.25
                switch (^(v91)) {
                case -2147483648:
                    if false {
                        if v67 {
                            var v95 float64 = 0.0
                            _ = v95
                        }
                        var v96 uint = 18446744073709551615
                        _ = v96
                    } else {
                        fmt.Println(v46, v66)
                    }
                case 1:
                    switch v91 {
                    case -1:
                        var v97 bool = (v67 && (!((!v67) && true)))
                        _ = v97
                    }
                case 2147483647:
                    v67 = (!(9223372036854775807 >= 9223372036854775807))
                }
                v98 := GenS8{F0: 127}
                if (false || false) {
                    v67 = ((!v67) && v67)
                    if ((-2.25 / float64((v94 - 0.0))) > 1.5) {
                        switch "foo bar" {
                        case "a":
                            switch v91 {
                            case -9223372036854775808:
                                if v67 {
                                    v99 := GenS9{F0: ((-417 != fi93) && v67), F1: v66, F2: 128}
                                    _ = v99
                                } else {
                                    if (!v67) {
                                        _ = len(v2)
                                        v3 = v3
                                        var v100 int32 = v4
                                        _ = v100
                                    }
                                    delete(v2, "foo bar")
                                    _ = len(v2)
                                }
                            case 4294967295:
                                var v101 int8 = v5
                                _ = v101
                            case 1:
                                fmt.Println(v3, v5)
                            }
                        case "\n\t":
                            v5 = (^(v5))
                        default:
                            var v102 float32 = (v94 - float32((((1.5 * (float64(v94) / ((3.14159 - (1.5 + (0.1 + 1.5))) - -2.25))) * 1e10) * 1.5)))
                            _ = v102
                        }
                        v91 += fi93
                        switch "\n\t" {
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                        default:
                            v91 = 32767
                        }
                    } else {
                        var v103 uint16 = (v3 + 1)
                        v91 -= (v91 << ((v46) & 63))
                        _ = v103
                    }
                    v47 -= v3
                } else {
                    v3 -= v3
                    if (((0.1 / 1.5) / 3.14159) > 0.1) {
                        fmt.Println(v66, v94, fi93)
                        fmt.Println(v67)
                    } else {
                        var v104 uint16 = 2
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        _ = v104
                    }
                    v5 = v5
                }
                _ = fi93
                _ = v94
                _ = v98
            }
            var v105 int8 = int8(v91)
            var v106 int = fi92
            _ = v105
            _ = v106
        } else {
            v91 = (v91 | v91)
            v107 := v2["a"]
            v66 += 45
            _ = v107
        }
        fmt.Println(v47, v5, v66)
        _ = fi92
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v66
    _ = v67
    _ = v90
    _ = v91
}
