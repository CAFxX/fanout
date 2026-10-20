package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint16; F1 int; F2 bool }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int16; F1 int; F2 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int) uint32 {
    v6 := []int{p0, p0, (p0 * p0)}
    p0 += p0
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: "hello"}
        var v9 string = "hello"
        if (!false) {
            for fi10 := 0; fi10 < 15; fi10++ {
                _ = len(v6)
                _ = fi10
            }
        } else {
            var v11 int32 = -2147483648
            var v12 float64 = 0.0
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v13 uint = 255
    if ("\n\t" < (("" + "hello") + "")) {
        var v14 uint16 = 2
        if (1 == 42) {
            _ = len(v6)
            switch "" {
            case "hello":
                switch (("" + (("" + "\n\t") + "hello")) + ("" + "a")) {
                case "\n\t":
                    v6 = append(v6, (-(p0)))
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            case "":
                v13 = 274
            default:
                for fi15 := 0; fi15 < 19; fi15++ {
                    if len(v6) > 0 {
                        v6[0] = p0
                    }
                    fmt.Println(p0, p0, p0)
                    _ = fi15
                }
            }
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 float64 = (((1.5 + -2.25) * 0.0) * 3.14159)
                v14 -= (v14 / ((1) | 1))
                var v18 uint64 = uint64(v13)
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        p0 -= (p0 & 128)
        _ = v14
    }
    _ = p0
    _ = v13
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -128
    var v3 int16 = -678
    var v4 int8 = -26
    var v5 float32 = (float32(-2.25) * (1e10 - ((((((-2.25 - -2.25) * (0.1 + 0.1)) * 1.5) + 1.5) - 3.14159) - (-2.25 - (0.0 / (((-2.25 + 1e10) - (1e10 - 1e10)) - 3.14159))))))
    v2 += (v4 >> ((169) & 63))
    fmt.Println(v2, v5, v3)
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v4, v2, v5)
        if (!(!(!(!(42 == 0))))) {
            fmt.Println(v5, fi19)
            if ((7 > 128) && false) {
                v3 = v3
            } else {
                if true {
                    v20 := []int{127}
                    _ = v20
                } else {
                    if true {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            if true {
                                v22 := []int{fi19, -1, -805, (fi21 & fi19)}
                                if true {
                                    v23 := GenS1{F0: "foo bar"}
                                    _ = v23
                                }
                                _ = v22
                            }
                            var v24 bool = ((("foo bar" + "a") + "foo bar") < ("\n\t" + "foo bar"))
                            v25 := GenS1{F0: ""}
                            fmt.Println(fi19)
                            _ = fi21
                            _ = v24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 4; fi26++ {
                            v27 := []int{}
                            var v28 uint16 = 711
                            switch (("a" + ("foo bar" + "")) + "foo bar") {
                            case "foo bar":
                                var v29 int32 = 2147483647
                                _ = v29
                            case "hello":
                                if (false && ((fi19 % ((fi26) | 1)) <= 0)) {
                                    if len(v27) > 0 {
                                        v27[0] = (fi26 | -1)
                                    }
                                    v4 += v2
                                    if false {
                                        switch fi19 {
                                        case 32767:
                                            var v30 string = "foo bar"
                                            _ = v30
                                        case 2147483647:
                                            v31 := GenS1{F0: ""}
                                            _ = v31
                                        default:
                                            switch 0 {
                                            case 32768:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "a":
                                                        if true {
                                                            if (253 == 364) {
                                                                v5 -= v5
                                                                v32 := GenS2{F0: ("a" + ("" + "")), F1: 193}
                                                                _ = v32
                                                            }
                                                            fmt.Println(v28, fi19)
                                                            v33 := GenS3{F0: (v28 % ((v28) | 1)), F1: (-(fi26)), F2: ((v4 | v4) >= v2)}
                                                            _ = v33
                                                        } else {
                                                            switch (fi26 * (fi19 * (fi19 + fi26))) {
                                                            case 127:
                                                                fmt.Println(v4, fi19)
                                                            case -2147483648:
                                                                v3 += v3
                                                            default:
                                                                var v34 int64 = 42
                                                                var v35 float32 = v5
                                                                switch fi19 {
                                                                case 127:
                                                                    v2 = (v2 / (((v4 - v2)) | 1))
                                                                case 32768:
                                                                    v36 := GenS4{F0: 0, F1: (v34 << ((661) & 63))}
                                                                    _ = v36
                                                                case 0:
                                                                    v34 = (^(v34))
                                                                }
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v28 = 255
                                                    default:
                                                        switch (fi19 + fi26) {
                                                        case -420:
                                                            var v37 uint8 = 128
                                                            _ = v37
                                                        case 0:
                                                            if false {
                                                                v38 := GenS3{F0: 162, F1: 0, F2: false}
                                                                v39 := v1[42]
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                        }
                                                    }
                                                    switch 42 {
                                                    case 9223372036854775807:
                                                        var v40 bool = ((-22 >= v2) || true)
                                                        _ = v40
                                                    case 257:
                                                        var v41 uint = 190
                                                        _ = v41
                                                    case 4294967295:
                                                        var v42 int32 = 1
                                                        var v43 int = (fi19 & fi19)
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    var v44 int8 = (v2 / ((127) | 1))
                                                    _ = v44
                                                }
                                            case -1:
                                                v4 -= v4
                                            case 745:
                                                fmt.Println(v5, v5, fi19)
                                                fmt.Println(v3, v4)
                                            }
                                        }
                                    } else {
                                        if false {
                                            v3 -= 488
                                            var v45 bool = false
                                            if true {
                                                switch "foo bar" {
                                                case "a":
                                                    var v46 int = (fi19 + 127)
                                                    _ = v46
                                                }
                                                fmt.Println(v4, v45)
                                                v27 = append(v27, 0)
                                                var v47 int8 = 100
                                                v5 += (v5 - v5)
                                                _ = v47
                                            } else {
                                                var v48 uint64 = 2
                                                v49 := GenS3{F0: (v28 | v28), F1: (fi19 << ((139) & 63)), F2: v45}
                                                _ = v48
                                                _ = v49
                                            }
                                            _ = v45
                                        } else {
                                            v3 = (v3 | -629)
                                            delete(v1, 65535)
                                            if ((!false) && false) {
                                                v5 += (float32(-2.25) - 1.5)
                                            }
                                            v50 := GenS5{F0: (^(v3)), F1: fi19, F2: 4294967295}
                                            var v51 float64 = (1e10 - (float64(v5) - 1.5))
                                            _ = len(v27)
                                            _ = v50
                                            _ = v51
                                        }
                                        v28 = (^(v28))
                                        v4 -= v4
                                    }
                                } else {
                                    if len(v27) > 0 {
                                        v27[0] = 57
                                    }
                                    var v52 float32 = v5
                                    _ = v52
                                }
                            }
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v4 += 100
                    } else {
                        for fi53 := 0; fi53 < 10; fi53++ {
                            var v54 bool = (((-2147483648 < -990) && false) || false)
                            _ = fi53
                            _ = v54
                        }
                        var v55 string = ("hello" + "")
                        _ = v55
                    }
                }
            }
        }
        _ = fi19
    }
    var v56 float64 = 1.5
    for fi57 := 0; fi57 < 15; fi57++ {
        var v58 float64 = v56
        fmt.Println(v4)
        _ = fi57
        _ = v58
    }
    if true {
        switch 86 {
        case -9223372036854775808:
            for fi59 := 0; fi59 < 10; fi59++ {
                v60 := GenS2{F0: "\n\t", F1: 1}
                delete(v1, 4294967295)
                if (!false) {
                    v5 = v5
                    if (-902 == 2147483647) {
                        switch (("a" + "foo bar") + "foo bar") {
                        case "hello":
                            if true {
                                v61 := []int{fi59, 61}
                                _ = len(v1)
                                _ = v61
                            } else {
                                v5 -= ((v5 / float32((1e10 - (v56 + 0.0)))) * v5)
                                switch (("\n\t" + "hello") + ("a" + "")) {
                                case "foo bar":
                                    v2 += 0
                                }
                            }
                        case "":
                            switch fi59 {
                            case 830:
                                for fi62 := 0; fi62 < 12; fi62++ {
                                    var v63 uint16 = 263
                                    var v64 int16 = (v3 << ((2) & 63))
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                            }
                        }
                        if false {
                            var v65 uint = 786
                            v2 = (v4 | (v2 & v4))
                            _ = v65
                        }
                        fmt.Println(v4, v3)
                    } else {
                        delete(v1, 4294967295)
                        for fi66 := 0; fi66 < 16; fi66++ {
                            v67 := GenS1{F0: "hello"}
                            _ = fi66
                            _ = v67
                        }
                        v1[297] = 65535
                    }
                } else {
                    v68 := []int{}
                    fmt.Println(v2, v5)
                    _ = v68
                }
                _ = fi59
                _ = v60
            }
        case 406:
            v69 := []int{32767, 9223372036854775807, 3}
            _ = v69
        case -766:
            v4 -= v4
        default:
            switch 1 {
            case 1:
                var v70 uint32 = 0
                _ = v70
            }
        }
        if (((255 != 1) || true) && false) {
            v3 = -32768
            _ = len(v1)
            fmt.Println(v4, v3, v4)
        }
    } else {
        fmt.Println(v4)
    }
    v1[382] = 65535
    switch "foo bar" {
    case "foo bar":
        var v71 float32 = (1e10 + (v5 - 3.14159))
        _ = v71
    case "a":
        var v72 string = ((("foo bar" + "foo bar") + ("a" + ("a" + ""))) + "\n\t")
        _ = v72
    case "hello":
        fmt.Println(v56)
    default:
        for fi73 := 0; fi73 < 2; fi73++ {
            if false {
                var v74 string = ("" + "a")
                _ = v74
            }
            fmt.Println(v3, v56, v56)
            _ = fi73
        }
    }
    if true {
        v2 -= v4
        if (4294967295 < 162) {
            fmt.Println(v5, v2, v3)
            v75 := []int{2147483647}
            _ = v75
        }
        v5 = (-2.25 + 0.1)
    } else {
        fmt.Println(v2)
        if (2147483647 > genFn0(0)) {
            fmt.Println(v5, v56)
            v3 -= v3
        } else {
            v56 = v56
            v76 := []int{}
            _ = v76
        }
    }
    v77 := []int{1, 128}
    v5 -= v5
    fmt.Println(v5)
    v78 := GenS5{F0: (v3 ^ int16(v2)), F1: 42, F2: genFn0(-1)}
    v56 = v56
    if (((v5 / -2.25) + v5) < ((((-2.25 + ((v5 - v5) + v5)) - -2.25) / (v5 - v5)) * v5)) {
        for fi79 := 0; fi79 < 12; fi79++ {
            if false {
                switch "" {
                case "foo bar":
                    var v80 string = ("foo bar" + "\n\t")
                    _ = v80
                case "hello":
                    for fi81 := 0; fi81 < 1; fi81++ {
                        delete(v1, 129)
                        fmt.Println(v5)
                        if len(v77) > 0 {
                            v77[0] = (fi79 * -398)
                        }
                        _ = fi81
                    }
                case "a":
                    var v82 uint = 18446744073709551615
                    _ = v82
                default:
                    v83 := GenS6{F0: -9223372036854775808}
                    _ = v83
                }
                v1[42] = 0
            } else {
                var v84 float32 = 1e10
                if true {
                    v2 = -128
                }
                var v85 float64 = 1.5
                _ = v84
                _ = v85
            }
            _ = fi79
        }
        fmt.Println(v5)
        if true {
            var v86 float64 = (float64((v5 / (1.5 / v5))) - 3.14159)
            var v87 float64 = 0.0
            var v88 int = -1
            _ = v86
            _ = v87
            _ = v88
        }
    } else {
        var v89 uint16 = 65535
        var v90 string = "foo bar"
        _ = v89
        _ = v90
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v77:", len(v77), v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v77
    _ = v78
}
