package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 bool; F1 int64; F2 uint }
type GenS12 struct { F0 int8; F1 string }
type GenS13 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS14 struct { F0 uint32; F1 int8; F2 int8 }
type GenS15 struct { F0 string }
func genFn0(p0 int64, p1 int8, p2 uint) uint16 {
    var v4 float64 = (0.0 * (3.14159 / 1e10))
    if false {
        var v5 string = "foo bar"
        fmt.Println(v4)
        fmt.Println(p0, v4)
        _ = v5
    } else {
        v6 := GenS1{F0: 90, F1: 940, F2: 42}
        p0 = (p0 | p0)
        var v7 uint16 = 721
        _ = v6
        _ = v7
    }
    p2 = 255
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 2
}
func genFn1(p0 uint8) bool {
    switch -2147483648 {
    case -9223372036854775808:
        v9 := GenS1{F0: 127, F1: 539, F2: p0}
        _ = v9
    }
    fmt.Println(p0)
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: 32768, F1: 2147483647}
        switch fi10 {
        case 678:
            var v12 int32 = 1
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    if (42 != 0) {
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(fi13, p0)
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS3{F0: 988, F1: "a"}
                switch 1 {
                case 4294967295:
                    p0 = p0
                case -671:
                    switch fi13 {
                    case 4294967295:
                        var v16 uint8 = p0
                        _ = v16
                    }
                default:
                    v17 := []int{(-(fi13)), fi13, (-(fi14)), (fi13 - (fi13 ^ (fi13 << ((18446744073709551615) & 63))))}
                    _ = v17
                }
                p0 -= 2
                _ = fi14
                _ = v15
            }
            switch "hello" {
            case "hello":
                if (false || (((2147483647 >= 65535) && (2 > 65535)) || false)) {
                    var v18 uint16 = 1
                    switch "" {
                    case "\n\t":
                        switch "" {
                        case "foo bar":
                            var v19 int64 = -1
                            _ = v19
                        case "a":
                            if true {
                                v20 := GenS4{F0: 42}
                                switch (fi13 % ((fi13) | 1)) {
                                case 32767:
                                    for fi21 := 0; fi21 < 3; fi21++ {
                                        var v22 string = "\n\t"
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                                _ = v20
                            }
                        case "\n\t":
                            for fi23 := 0; fi23 < 5; fi23++ {
                                var v24 uint64 = 2
                                v25 := []int{(fi13 * int(v18)), (-(fi13)), (fi23 & (fi13 << ((42) & 63)))}
                                if (0 == 0) {
                                    v26 := GenS5{F0: false, F1: uint16(p0), F2: 42}
                                    fmt.Println(fi13, v18)
                                    var v27 uint32 = 859
                                    if (p0 < p0) {
                                        if (!true) {
                                            v28 := GenS5{F0: (-1 > -1), F1: (v18 >> ((uint(v27)) & 63)), F2: 752}
                                            var v29 uint = 4294967295
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 uint8 = p0
                                            v27 -= (v27 + (^(v27)))
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            v18 = (v18 ^ (v18 & 0))
                                            v31 := GenS6{F0: 2147483647}
                                            _ = v30
                                            _ = v31
                                        }
                                        if (!false) {
                                            v27 -= (v27 & (v27 / ((v27) | 1)))
                                            v25 = append(v25, fi23)
                                        } else {
                                            fmt.Println(fi13, v24)
                                            v32 := GenS7{F0: 42}
                                            var v33 int = (fi13 % ((fi23) | 1))
                                            _ = v32
                                            _ = v33
                                        }
                                        switch (("foo bar" + ("hello" + (("hello" + "hello") + ("a" + "\n\t")))) + "a") {
                                        case "\n\t":
                                            v24 += 18446744073709551615
                                        case "a":
                                            var v34 int64 = -109
                                            _ = v34
                                        case "":
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(fi23, v18, p0)
                                            }
                                            var v35 uint16 = (v18 >> ((785) & 63))
                                            var v36 bool = ((735 <= v18) || false)
                                            _ = v35
                                            _ = v36
                                        }
                                        p0 += (^(p0))
                                    } else {
                                        if true {
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            if true {
                                                var v37 int8 = -128
                                                if (v27 <= v27) {
                                                    v24 = v24
                                                    var v38 float32 = float32((-2.25 / ((-2.25 - 0.0) * 1e10)))
                                                    _ = v38
                                                } else {
                                                    if ((-2.25 / 1.5) != 3.14159) {
                                                        if true {
                                                            var v39 bool = false
                                                            v40 := GenS8{F0: ("a" + "a")}
                                                            _ = len(v25)
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                        fmt.Println(v27, v24, v18)
                                                        v41 := GenS9{F0: (v27 % ((v27) | 1))}
                                                        fmt.Println(fi23, fi13, fi13)
                                                        v37 = int8(p0)
                                                        _ = v41
                                                    }
                                                    p0 += (p0 >> ((0) & 63))
                                                }
                                                fmt.Println(v37, v18)
                                                _ = v37
                                            }
                                        } else {
                                            var v42 int16 = -1
                                            _ = v42
                                        }
                                        v43 := GenS10{F0: (v27 << ((42) & 63)), F1: "foo bar"}
                                        switch 32768 {
                                        case -1:
                                            var v44 float64 = 0.0
                                            _ = v44
                                        case 127:
                                            var v45 string = ""
                                            _ = v45
                                        }
                                        _ = v43
                                    }
                                    var v46 int16 = 1
                                    fmt.Println(v27, fi13, v46)
                                    _ = v26
                                    _ = v27
                                    _ = v46
                                }
                                var v47 int = fi13
                                if len(v25) > 0 {
                                    v25[0] = 128
                                }
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v47
                            }
                        }
                    case "a":
                        if false {
                            for fi48 := 0; fi48 < 9; fi48++ {
                                var v49 uint64 = 200
                                _ = fi48
                                _ = v49
                            }
                            if (("\n\t" != "\n\t") || false) {
                                for fi50 := 0; fi50 < 1; fi50++ {
                                    v51 := []int{409, (-(fi50))}
                                    if len(v51) > 0 {
                                        v51[0] = -2147483648
                                    }
                                    if len(v51) > 0 {
                                        v51[0] = 42
                                    }
                                    _ = fi50
                                    _ = v51
                                }
                                switch "\n\t" {
                                case "":
                                    fmt.Println(p0, fi13, p0)
                                case "a":
                                    var v52 uint32 = uint32(p0)
                                    _ = v52
                                case "hello":
                                    for fi53 := 0; fi53 < 8; fi53++ {
                                        if true {
                                            p0 = p0
                                        } else {
                                            p0 = 46
                                            v54 := GenS3{F0: 4294967295, F1: "foo bar"}
                                            v18 = (v18 % (((v18 >> ((158) & 63))) | 1))
                                            _ = v54
                                        }
                                        v55 := []int{fi53, (^(fi53)), -830, 2147483647}
                                        switch 1 {
                                        case 32767:
                                            if false {
                                                p0 += 1
                                                if len(v55) > 0 {
                                                    v55[0] = (fi53 - 2147483647)
                                                }
                                            }
                                        case -251:
                                            v56 := GenS7{F0: 42}
                                            _ = v56
                                        default:
                                            v57 := GenS11{F0: (!true), F1: int64(v18), F2: 388}
                                            _ = v57
                                        }
                                        _ = fi53
                                        _ = v55
                                    }
                                }
                            } else {
                                p0 += (p0 & (p0 >> ((749) & 63)))
                                v58 := GenS12{F0: 42, F1: ""}
                                v59 := []int{128, (fi13 ^ (fi13 ^ fi13)), (fi13 * (fi13 << ((883) & 63)))}
                                _ = v58
                                _ = v59
                            }
                        }
                    default:
                        if true {
                            v60 := []int{}
                            _ = v60
                        } else {
                            v18 += (v18 * v18)
                            v61 := []int{(-(fi13)), (-(fi13)), fi13, (fi13 * 4294967295)}
                            _ = v61
                        }
                    }
                    v62 := []int{fi13, (fi13 >> ((0) & 63))}
                    _ = v18
                    _ = v62
                }
            }
            _ = fi13
        }
        v63 := []int{128, 9223372036854775807, 2147483647}
        v64 := GenS13{F0: (!false), F1: p0, F2: 0}
        _ = v63
        _ = v64
    } else {
        v65 := []int{-2147483648, 518, 0, 128}
        _ = v65
    }
    v66 := []int{775, 263, -181}
    _ = p0
    _ = v66
    return false
}

func main() {
    v1 := []string{}
    var v2 float32 = (0.1 - 1.5)
    v3 := make(map[string]int)
    var v67 uint8 = 92
    v68 := GenS10{F0: 26, F1: "a"}
    v1 = append(v1, "hello")
    switch "\n\t" {
    case "foo bar":
        if true {
            if (-128 <= 42) {
                var v69 float32 = v2
                _ = v69
            }
        }
    case "a":
        v1 = append(v1, ("" + (("a" + "hello") + "a")))
    case "\n\t":
        for fi70 := 0; fi70 < 4; fi70++ {
            _ = len(v3)
            v2 = v2
            _ = fi70
        }
    }
    switch -538 {
    case 32768:
        for fi71 := 0; fi71 < 6; fi71++ {
            for fi72 := 0; fi72 < 19; fi72++ {
                if ((!("foo bar" == "\n\t")) && (true || (!(42 <= 42)))) {
                    v73 := GenS6{F0: 1}
                    v2 = v2
                    v2 = v2
                    if true {
                        v67 -= 84
                        var v74 int16 = -32768
                        if false {
                            if (4294967295 < 914) {
                                var v75 int64 = -9223372036854775808
                                _ = v75
                            } else {
                                v1 = append(v1, ("a" + "foo bar"))
                                v76 := GenS14{F0: 385, F1: -42, F2: 73}
                                var v77 float32 = (0.1 * ((float32(1.5) + v2) / v2))
                                _ = v76
                                _ = v77
                            }
                        }
                        _ = v74
                    }
                    if (32767 >= genFn0(1, -78, 1)) {
                        v2 += (((v2 - v2) * (0.0 + (v2 / (0.1 - 0.0)))) - (((v2 / 0.1) - v2) - v2))
                        var v78 int8 = -92
                        fmt.Println(fi71)
                        _ = v78
                    }
                    _ = v73
                }
                _ = fi72
            }
            var v79 int = fi71
            var v80 float64 = 1e10
            _ = fi71
            _ = v79
            _ = v80
        }
    case -752:
        v81 := GenS15{F0: ("foo bar" + "")}
        _ = v81
    case 1:
        v2 = v2
    }
    delete(v3, "a")
    v67 -= v67
    var v82 int = 32768
    switch (-(v82)) {
    case 2147483647:
        fmt.Println(v82, v2, v82)
    default:
        v1 = append(v1, ("a" + ("a" + "foo bar")))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v67:", v67)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v68
    _ = v82
}
