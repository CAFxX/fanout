package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int8; F1 int8; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int; F1 uint8; F2 uint64 }
type GenS8 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int32; F1 uint8; F2 string }
func genFn0(p0 uint) int16 {
    var v7 uint32 = 42
    fmt.Println(p0)
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(p0, fi8)
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int8 = -31
            var v11 uint16 = 255
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            for fi12 := 0; fi12 < 18; fi12++ {
                v7 -= (v7 & 975)
                v13 := []int{fi8, fi8, fi8}
                p0 = p0
                v7 += (^(v7))
                _ = fi12
                _ = v13
            }
            v7 -= uint32(p0)
        }
        _ = fi8
    }
    v14 := GenS1{F0: true, F1: 127, F2: 998}
    fmt.Println(p0, p0)
    _ = p0
    _ = v14
    _ = v7
    return -1
}
func genFn1(p0 uint32, p1 int, p2 uint64) uint8 {
    if false {
        p0 = 2147483647
    }
    v15 := GenS1{F0: false, F1: 79, F2: 935}
    p1 += (p1 | p1)
    if (!((1.5 - 1.5) > (1.5 + 3.14159))) {
        v16 := []int{(p1 | (p1 ^ (p1 & 1))), -9223372036854775808}
        _ = v16
    }
    switch "foo bar" {
    case "":
        v17 := []int{}
        _ = v17
    case "a":
        v18 := []int{2147483647, 9223372036854775807, (p1 ^ (p1 ^ p1))}
        _ = v18
    case "\n\t":
        fmt.Println(p0, p0, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    v3 := make(map[string]bool)
    var v4 bool = v2
    var v5 uint64 = 0
    v6 := []string{"\n\t"}
    delete(v1, (("foo bar" + ("foo bar" + "foo bar")) + ""))
    if (42 > 0) {
        if (!false) {
            if true {
                var v19 int32 = 30
                _ = v19
            }
        } else {
            switch 32767 {
            case 132:
                v6 = append(v6, "hello")
            case 4294967295:
                fmt.Println(v4, v2)
            default:
                var v20 float32 = -2.25
                _ = v20
            }
        }
        v4 = v4
        var v21 int = 37
        _ = v21
    } else {
        switch 0 {
        case 42:
            if v2 {
                v22 := GenS1{F0: ("hello" > ("foo bar" + ("" + "hello"))), F1: genFn1(2147483648, 2147483647, (v5 | v5)), F2: 255}
                switch "foo bar" {
                case "\n\t":
                    var v23 int = -2147483648
                    _ = v23
                default:
                    if false {
                        switch "" {
                        case "":
                            var v24 int16 = genFn0(18446744073709551615)
                            _ = v24
                        case "hello":
                            var v25 uint16 = 32767
                            _ = v25
                        case "\n\t":
                            v6 = append(v6, ("a" + "\n\t"))
                        }
                        if v4 {
                            var v26 uint8 = 128
                            for fi27 := 0; fi27 < 15; fi27++ {
                                var v28 string = "hello"
                                if (!v2) {
                                    switch fi27 {
                                    case 4294967295:
                                        _ = len(v1)
                                    case -2147483648:
                                        if (((!v4) && (v4 || false)) || ("" == "\n\t")) {
                                            var v29 uint32 = 2
                                            _ = v29
                                        } else {
                                            var v30 int = fi27
                                            _ = v30
                                        }
                                    case 127:
                                        for fi31 := 0; fi31 < 5; fi31++ {
                                            if (v2 || false) {
                                                var v32 float32 = float32(0.0)
                                                v33 := GenS2{F0: 42}
                                                var v34 string = v28
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            } else {
                                                var v35 int = fi31
                                                _ = v35
                                            }
                                            v36 := GenS3{F0: (v28 < ("" + v28))}
                                            v4 = (42 > 532)
                                            v4 = (-1 == -124)
                                            v5 += v5
                                            _ = fi31
                                            _ = v36
                                        }
                                    }
                                } else {
                                    v37 := GenS3{F0: v4}
                                    _ = v37
                                }
                                _ = fi27
                                _ = v28
                            }
                            _ = v26
                        }
                        v2 = v4
                    }
                }
                _ = v22
            }
        case 202:
            for fi38 := 0; fi38 < 14; fi38++ {
                var v39 bool = ((fi38 * fi38) <= (fi38 + (fi38 | 32768)))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                switch 573 {
                case -531:
                    v4 = (true || v2)
                case -107:
                    if (-841 >= -456) {
                        fmt.Println(v39, v4, v39)
                        var v40 float32 = ((1e10 * (((1e10 / 1.5) * (1.5 + 0.1)) + ((3.14159 * (0.0 - (1e10 + 1.5))) - 0.1))) + 1e10)
                        _ = v40
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v41 := GenS1{F0: true, F1: 127, F2: 255}
                        v42 := GenS1{F0: (2147483648 >= 452), F1: 255, F2: 131}
                        _ = v41
                        _ = v42
                    }
                case -169:
                    for fi43 := 0; fi43 < 11; fi43++ {
                        fmt.Println(v5, v5, v4)
                        v44 := GenS4{F0: v5, F1: 255}
                        _ = fi43
                        _ = v44
                    }
                default:
                    _ = len(v6)
                }
                _ = fi38
                _ = v39
            }
        case -9223372036854775808:
            if (1 <= 419) {
                for fi45 := 0; fi45 < 19; fi45++ {
                    var v46 int = fi45
                    for fi47 := 0; fi47 < 16; fi47++ {
                        if (fi47 == (fi45 & (fi47 << ((uint(v5)) & 63)))) {
                            v48 := GenS5{F0: 127, F1: -1, F2: 714}
                            fmt.Println(fi45, v4)
                            fmt.Println(v2)
                            _ = v48
                        }
                        _ = fi47
                    }
                    switch "" {
                    case "foo bar":
                        v2 = v2
                    }
                    _ = fi45
                    _ = v46
                }
                var v49 float64 = (0.1 - 1.5)
                _ = v49
            }
        default:
            var v50 int32 = -1
            _ = v50
        }
        v51 := GenS6{F0: 1, F1: v5}
        if v2 {
            if v2 {
                if v2 {
                    if v2 {
                        var v52 bool = ((-273 != -600) && (42 == 928))
                        var v53 int8 = -118
                        _ = v52
                        _ = v53
                    }
                    var v54 int16 = 1
                    _ = v54
                }
                var v55 uint = 1
                _ = v55
            }
            if false {
                if v4 {
                    fmt.Println(v2, v5, v5)
                    var v56 float32 = (float32(0.1) - ((1.5 + (float32((0.0 * -2.25)) / 1e10)) / (1.5 - float32((1.5 + (-2.25 + 1e10))))))
                    _ = v56
                }
                for fi57 := 0; fi57 < 10; fi57++ {
                    var v58 string = "\n\t"
                    fmt.Println(v58, v4)
                    _ = fi57
                    _ = v58
                }
                v2 = true
            }
        }
        _ = v51
    }
    if (65535 > 990) {
        v59 := v3[("" + "")]
        switch "a" {
        case "\n\t":
            if v4 {
                switch "foo bar" {
                case "a":
                    var v60 int = -354
                    _ = v60
                case "hello":
                    var v61 uint64 = (v5 % (((v5 << ((829) & 63))) | 1))
                    _ = v61
                default:
                    var v62 bool = v59
                    _ = v62
                }
            } else {
                _ = len(v6)
                v63 := GenS5{F0: 42, F1: 100, F2: 0}
                _ = v63
            }
        }
        var v64 string = "hello"
        _ = v59
        _ = v64
    }
    for fi65 := 0; fi65 < 19; fi65++ {
        v66 := GenS6{F0: 9223372036854775807, F1: (v5 - v5)}
        var v67 uint64 = v5
        _ = fi65
        _ = v66
        _ = v67
    }
    switch -472 {
    case -228:
        var v68 uint32 = 2
        _ = v68
    case 2147483647:
        var v69 string = "foo bar"
        _ = v69
    case 4294967295:
        for fi70 := 0; fi70 < 6; fi70++ {
            v6 = append(v6, ("\n\t" + (((("" + "") + "hello") + ("" + "hello")) + "a")))
            var v71 int16 = -792
            var v72 float64 = 0.1
            _ = fi70
            _ = v71
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 3; fi73++ {
        fmt.Println(v2, v2, v5)
        v4 = v2
        v74 := GenS5{F0: 0, F1: 42, F2: 2147483648}
        _ = fi73
        _ = v74
    }
    if v4 {
        if true {
            switch "a" {
            case "":
                v6 = append(v6, ("\n\t" + "a"))
            case "hello":
                v75 := v1[(("\n\t" + "hello") + "a")]
                _ = v75
            }
            var v76 int8 = -25
            _ = v76
        }
    } else {
        for fi77 := 0; fi77 < 9; fi77++ {
            if (42 <= -1) {
                if (false || (v5 != (v5 + (v5 | (v5 | (v5 | v5)))))) {
                    fmt.Println(v4)
                } else {
                    var v78 string = "\n\t"
                    if v2 {
                        v3[v78] = ((fi77 / ((127) | 1)) >= (fi77 ^ fi77))
                    }
                    var v79 int = fi77
                    _ = v78
                    _ = v79
                }
            }
            v4 = v2
            if ((v5 ^ 0) == v5) {
                v80 := GenS5{F0: 0, F1: 42, F2: 949}
                _ = v80
            } else {
                for fi81 := 0; fi81 < 9; fi81++ {
                    var v82 float64 = float64(-2.25)
                    v83 := GenS7{F0: (fi81 * fi77), F1: 146, F2: (^(v5))}
                    v84 := GenS4{F0: (v5 >> ((883) & 63)), F1: 255}
                    _ = fi81
                    _ = v82
                    _ = v83
                    _ = v84
                }
                v5 = (v5 >> ((0) & 63))
            }
            _ = fi77
        }
        v4 = (genFn0(0) <= -1)
    }
    v4 = (0 >= 255)
    delete(v3, "")
    for fi85 := 0; fi85 < 11; fi85++ {
        for fi86 := 0; fi86 < 8; fi86++ {
            var v87 float64 = (float64(3.14159) / 1.0)
            v88 := v3["a"]
            var v89 string = "\n\t"
            _ = fi86
            _ = v87
            _ = v88
            _ = v89
        }
        _ = fi85
    }
    for fi90 := 0; fi90 < 1; fi90++ {
        v5 += v5
        switch "\n\t" {
        case "hello":
            v2 = v4
        case "\n\t":
            for fi91 := 0; fi91 < 10; fi91++ {
                switch fi91 {
                case -9223372036854775808:
                    fmt.Println(v5, v2, v2)
                case 9223372036854775807:
                    v92 := GenS1{F0: v4, F1: 1, F2: 255}
                    _ = v92
                default:
                    var v93 uint16 = 414
                    v94 := GenS8{F0: 255, F1: v93, F2: 1}
                    _ = v93
                    _ = v94
                }
                _ = fi91
            }
        case "a":
            var v95 uint64 = 18446744073709551615
            _ = v95
        default:
            v4 = true
        }
        if v2 {
            v96 := v1[""]
            var v97 float64 = (float64(float32(1e10)) + 3.14159)
            var v98 int8 = 100
            _ = v96
            _ = v97
            _ = v98
        }
        _ = fi90
    }
    v99 := GenS9{F0: 362, F1: 255, F2: "foo bar"}
    var v100 int16 = 42
    fmt.Println(v2, v2)
    v101 := GenS9{F0: -2147483648, F1: 2, F2: "\n\t"}
    v2 = ((v5 & v5) <= (v5 + v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
