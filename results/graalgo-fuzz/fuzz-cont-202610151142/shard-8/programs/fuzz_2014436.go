package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint8; F2 int }
type GenS4 struct { F0 string; F1 string; F2 int8 }
type GenS5 struct { F0 int16; F1 int16; F2 uint64 }
func genFn0(p0 uint16, p1 uint64) bool {
    var v7 uint = 18446744073709551615
    var v8 uint16 = 42
    var v9 string = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 uint16, p1 uint8, p2 uint32) int {
    v10 := []int{-896, 211, 2147483647, -5}
    v11 := GenS1{F0: 42, F1: -1}
    v12 := GenS1{F0: 14, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 246
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -379
    v3 := []uint32{1, 42, 888}
    var v4 float32 = -2.25
    var v5 uint = 1
    v6 := make(map[uint32]uint32)
    v13 := GenS1{F0: -960, F1: -574}
    var v14 bool = false
    var v15 string = ("foo bar" + ("hello" + "hello"))
    v16 := GenS1{F0: 0, F1: 126}
    switch 1 {
    case 2147483647:
        if ((-(v2)) < (v2 - (v2 * (v2 - v2)))) {
            var v17 bool = (false && v14)
            fmt.Println(v15, v17)
            _ = v17
        } else {
            var v18 bool = v14
            for fi19 := 0; fi19 < 10; fi19++ {
                v3 = append(v3, 2)
                v14 = false
                if true {
                    v20 := GenS1{F0: 789, F1: 2147483647}
                    if (v14 || (42 <= 236)) {
                        v3 = append(v3, 1)
                        v6[65535] = 42
                        if (v18 || (v14 || false)) {
                            v5 += 4294967295
                            v21 := GenS1{F0: 0, F1: -689}
                            var v22 uint16 = 32767
                            _ = v21
                            _ = v22
                        } else {
                            if v14 {
                                var v23 float64 = 3.14159
                                v2 = v2
                                _ = v23
                            }
                        }
                    }
                    if (0 != uint8(v5)) {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            var v25 int16 = (v2 - v2)
                            v26 := GenS1{F0: 9223372036854775807, F1: 940}
                            v1[genFn1(547, 255, 690)] = 128
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        if v14 {
                            var v27 uint8 = 2
                            v15 = "a"
                            _ = v27
                        }
                        v28 := GenS2{F0: (fi19 + fi19), F1: 11}
                        _ = v28
                    } else {
                        fmt.Println(v4)
                        if true {
                            v29 := GenS1{F0: 0, F1: -1}
                            for fi30 := 0; fi30 < 14; fi30++ {
                                var v31 uint = v5
                                _ = fi30
                                _ = v31
                            }
                            _ = v29
                        }
                    }
                    _ = v20
                } else {
                    var v32 uint8 = 1
                    _ = v32
                }
                _ = fi19
            }
            _ = v18
        }
    case -2147483648:
        v33 := GenS2{F0: -73, F1: 127}
        _ = v33
    case 859:
        v2 -= (v2 / ((v2) | 1))
    default:
        fmt.Println(v5, v14, v5)
    }
    v34 := v1[1]
    var v35 int32 = 1
    if false {
        v34 = 0
        v36 := v1[genFn1(149, 1, 42)]
        _ = v36
    }
    v37 := GenS3{F0: int8(v34), F1: 241, F2: genFn1(32767, 127, 4294967295)}
    fmt.Println(v4)
    if v14 {
        if ((^(v2)) > (^(v2))) {
            var v38 int64 = -9223372036854775808
            v4 -= (v4 * ((v4 - (((1e10 * v4) - ((float32(3.14159) - (3.14159 + v4)) + (v4 + -2.25))) * v4)) + v4))
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 bool = ((^(v5)) <= 390)
                if ((true || (v35 > -112)) || v14) {
                    v41 := GenS2{F0: genFn1(170, 252, 117), F1: 2}
                    fmt.Println(v35, v2, v38)
                    v35 = int32(v5)
                    _ = v41
                }
                _ = fi39
                _ = v40
            }
            _ = v38
        } else {
            var v42 uint16 = 607
            var v43 uint = v5
            var v44 uint = (^(v5))
            _ = v42
            _ = v43
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        if true {
            v46 := GenS3{F0: 126, F1: 127, F2: (v34 - genFn1(32767, 255, 0))}
            var v47 int64 = 42
            if (v14 && ((^(v2)) == v2)) {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                v4 += v4
                v14 = v14
            } else {
                for fi48 := 0; fi48 < 8; fi48++ {
                    switch genFn1(32767, 128, 2147483648) {
                    case -377:
                        fmt.Println(v47, v14)
                    case 128:
                        _ = len(v1)
                    case 484:
                        switch (fi48 | -2147483648) {
                        case 42:
                            fmt.Println(fi45, v14)
                        case -387:
                            var v49 int = int(v47)
                            _ = v49
                        case -323:
                            var v50 int64 = v47
                            _ = v50
                        }
                    default:
                        var v51 int16 = v2
                        if ((42 >= 1) || true) {
                            v15 = v15
                            var v52 uint32 = 4294967295
                            _ = v52
                        } else {
                            if (!(v47 < v47)) {
                                if (!(2 <= 0)) {
                                    switch "a" {
                                    case "a":
                                        switch v15 {
                                        case "a":
                                            var v53 bool = false
                                            _ = v53
                                        case "":
                                            switch (("" + "\n\t") + v15) {
                                            case "foo bar":
                                                switch genFn1(543, 128, 42) {
                                                case 495:
                                                    if ((-1 == -53) && (42 >= 32767)) {
                                                        var v54 uint32 = 65535
                                                        fmt.Println(v4, v14)
                                                        var v55 uint16 = 0
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                default:
                                                    var v56 int32 = -311
                                                    _ = v56
                                                }
                                            default:
                                                v57 := GenS1{F0: v47, F1: -468}
                                                _ = v57
                                            }
                                        default:
                                            if false {
                                                var v58 bool = (3.14159 != v4)
                                                if (!v14) {
                                                    v59 := GenS4{F0: "hello", F1: v15, F2: 100}
                                                    v60 := GenS4{F0: (v15 + ("foo bar" + ("a" + ""))), F1: v15, F2: -128}
                                                    var v61 float32 = v4
                                                    fmt.Println(v14, v51)
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v34 = -1
                                                    var v62 int = v34
                                                    switch v15 {
                                                    case "\n\t":
                                                        v47 = v47
                                                    case "":
                                                        v5 = (v5 | 42)
                                                    }
                                                    _ = v62
                                                }
                                                _ = v58
                                            }
                                        }
                                    case "foo bar":
                                        v63 := GenS2{F0: (fi48 >> ((v5) & 63)), F1: 145}
                                        _ = v63
                                    case "hello":
                                        v3 = append(v3, 2147483647)
                                        if (1 != 185) {
                                            var v64 float64 = 0.0
                                            _ = v64
                                        } else {
                                            _ = len(v6)
                                            switch (fi45 + -936) {
                                            case 32768:
                                                var v65 bool = (!((v14 || v14) && (true && (!true))))
                                                _ = v65
                                            case -745:
                                                var v66 uint8 = 127
                                                _ = v66
                                            default:
                                                switch (fi48 * (fi48 - fi48)) {
                                                case 1:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case 2147483647:
                                                    var v67 float32 = (((float32(((0.0 + 0.0) + 0.1)) + (v4 * (v4 * v4))) + 0.1) * (v4 * (1e10 - (((1.5 / v4) * (v4 - (0.0 / (0.0 + (v4 * v4))))) + 3.14159))))
                                                    _ = v67
                                                default:
                                                    if true {
                                                        if (v2 == 1000) {
                                                            switch v15 {
                                                            case "a":
                                                                v14 = v14
                                                            case "foo bar":
                                                                if (((v14 || (128 >= 0)) && v14) && v14) {
                                                                    v4 -= 1e10
                                                                }
                                                            }
                                                        }
                                                        var v68 bool = (((v2 / ((-32768) | 1)) >= (v51 * -1)) && v14)
                                                        var v69 int = genFn1(0, 235, 414)
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    var v70 bool = v14
                                                    delete(v1, v34)
                                                    _ = v70
                                                }
                                            }
                                            v15 = v15
                                            fmt.Println(fi48, v35)
                                            delete(v6, 185)
                                            var v71 float64 = (0.0 + -2.25)
                                            _ = v71
                                        }
                                        var v72 uint16 = 766
                                        _ = v72
                                    default:
                                        var v73 int64 = (v47 >> ((42) & 63))
                                        _ = v73
                                    }
                                    v34 -= (fi48 / ((-177) | 1))
                                    fmt.Println(v4, fi45)
                                }
                            }
                        }
                        _ = v51
                    }
                    if (876 < 2) {
                        v34 += fi45
                        v74 := GenS4{F0: v15, F1: "foo bar", F2: 127}
                        _ = v74
                    }
                    var v75 int = (fi48 % ((v34) | 1))
                    var v76 uint16 = 42
                    var v77 bool = true
                    _ = fi48
                    _ = v75
                    _ = v76
                    _ = v77
                }
            }
            _ = v46
            _ = v47
        } else {
            var v78 int8 = -32
            v34 += v34
            v79 := GenS5{F0: v2, F1: v2, F2: 2}
            _ = v78
            _ = v79
        }
        _ = fi45
    }
    v80 := GenS3{F0: 100, F1: 239, F2: v34}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v80
}
