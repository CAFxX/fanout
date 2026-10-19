package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int16; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint64; F2 int8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint8 }
func genFn0() int16 {
    if true {
        v7 := GenS1{F0: (((("foo bar" + "a") + ("a" + "hello")) + ((("a" + ("" + "")) + "") + "\n\t")) + "\n\t"), F1: 313, F2: 935}
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint64 = 1
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        var v10 int64 = 908
        _ = v10
    }
    v11 := GenS1{F0: ("a" + "hello"), F1: 2, F2: 0}
    v12 := GenS1{F0: "\n\t", F1: 1, F2: 1}
    var v13 uint64 = 698
    v14 := GenS1{F0: "foo bar", F1: v13, F2: (v13 | 1)}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 42
}
func genFn1(p0 int64, p1 uint8) int8 {
    if (false && (-2147483648 != 2147483647)) {
        for fi15 := 0; fi15 < 17; fi15++ {
            p0 += p0
            var v16 int = 127
            _ = fi15
            _ = v16
        }
    }
    var v17 float32 = float32(3.14159)
    _ = p0
    _ = p1
    _ = v17
    return 0
}

func main() {
    v1 := []string{("\n\t" + ((("a" + "a") + ("\n\t" + "\n\t")) + (((("a" + "") + "a") + "hello") + ("foo bar" + "a")))), ("hello" + "a"), ("foo bar" + "a")}
    var v2 bool = true
    var v3 int64 = 0
    var v4 uint16 = 0
    var v5 string = "a"
    var v6 uint16 = 0
    var v18 string = (v5 + (((("a" + ((((v5 + v5) + ("foo bar" + (((v5 + (v5 + "hello")) + (v5 + v5)) + "hello"))) + v5) + "a")) + "\n\t") + "foo bar") + v5))
    if (!false) {
        if (0.1 > 0.0) {
            v18 = v18
            v3 = 1
        }
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v20 float64 = float64(3.14159)
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            if true {
                v2 = v2
                if false {
                    if true {
                        _ = len(v1)
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v3 -= v3
                            _ = fi22
                        }
                    }
                }
                switch -25 {
                case -793:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 uint32 = 2
                        var v25 int32 = -2147483648
                        v26 := GenS1{F0: v18, F1: 18446744073709551615, F2: 2}
                        var v27 uint8 = 76
                        if v2 {
                            var v28 uint64 = 18446744073709551615
                            v29 := GenS1{F0: "hello", F1: v28, F2: (v28 / ((443) | 1))}
                            _ = v28
                            _ = v29
                        } else {
                            var v30 float64 = ((0.1 / 1.5) / 1.0)
                            _ = v30
                        }
                        v3 = (v3 + v3)
                        fmt.Println(v5, v25)
                        if (v24 < v24) {
                            fmt.Println(v27, v2, v18)
                            var v31 bool = v2
                            var v32 uint8 = v27
                            _ = v31
                            _ = v32
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
            } else {
                switch fi21 {
                case 796:
                    v33 := GenS2{F0: int32(v3), F1: 129}
                    _ = v33
                default:
                    var v34 uint64 = 790
                    _ = v34
                }
                v35 := GenS1{F0: (v18 + v18), F1: 4294967295, F2: 4294967295}
                var v36 int64 = v3
                _ = v35
                _ = v36
            }
            v37 := GenS1{F0: v18, F1: 0, F2: 2}
            v2 = v2
            _ = fi21
            _ = v37
        }
        v6 += 32767
    }
    if v2 {
        v5 = ""
        var v38 int = 1
        if v2 {
            var v39 bool = (!v2)
            switch "" {
            case "hello":
                fmt.Println(v5, v4)
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v40 float32 = 3.14159
                _ = v40
            }
            _ = v39
        }
        _ = v38
    } else {
        v41 := GenS3{F0: 1}
        _ = v41
    }
    v5 = ("a" + "hello")
    v5 = ((((v18 + v5) + v5) + v18) + v18)
    var v42 uint32 = 2
    if (!v2) {
        if v2 {
            v2 = v2
            v6 -= (^(v6))
            v43 := GenS1{F0: "hello", F1: 0, F2: 477}
            _ = v43
        }
        v44 := GenS1{F0: ("a" + (((v18 + (v5 + "hello")) + (("foo bar" + ("\n\t" + v18)) + v5)) + ("\n\t" + (v18 + (((("" + v18) + v18) + (v18 + (("hello" + ("hello" + "")) + "hello"))) + "a"))))), F1: 1, F2: 4294967295}
        for fi45 := 0; fi45 < 2; fi45++ {
            v5 = ("hello" + "a")
            _ = fi45
        }
        _ = v44
    } else {
        if v2 {
            v46 := GenS4{F0: (-842 < genFn0()), F1: 191, F2: 693}
            v47 := GenS1{F0: v5, F1: 438, F2: 789}
            switch 9223372036854775807 {
            case -2147483648:
                for fi48 := 0; fi48 < 13; fi48++ {
                    v18 = "foo bar"
                    if v2 {
                        switch v5 {
                        case "a":
                            if ((fi48 % ((fi48) | 1)) > (fi48 | 128)) {
                                for fi49 := 0; fi49 < 7; fi49++ {
                                    var v50 int = (fi49 ^ (fi49 / (((fi48 - (^(fi49)))) | 1)))
                                    if (-430 == -1) {
                                        var v51 string = ((("foo bar" + ((v18 + ((v18 + "") + v5)) + "a")) + v5) + v18)
                                        _ = v51
                                    }
                                    fmt.Println(v5, v6, fi48)
                                    var v52 bool = v2
                                    var v53 int = 32767
                                    _ = fi49
                                    _ = v50
                                    _ = v52
                                    _ = v53
                                }
                            } else {
                                for fi54 := 0; fi54 < 14; fi54++ {
                                    var v55 int32 = 0
                                    var v56 float64 = 0.1
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                }
                            }
                        }
                        fmt.Println(v18, v5)
                    } else {
                        for fi57 := 0; fi57 < 20; fi57++ {
                            switch (-(fi48)) {
                            case 32768:
                                if false {
                                    fmt.Println(v42, fi48, fi48)
                                    v3 = v3
                                    if ((1 >= 0) && v2) {
                                        v4 -= v4
                                        var v58 int8 = 1
                                        _ = v58
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v3, fi48, v3)
                                    case "":
                                        var v59 float32 = 1e10
                                        fmt.Println(fi48)
                                        _ = v59
                                    default:
                                        v4 -= v6
                                    }
                                    if (true || (42 != 1)) {
                                        v18 = ("foo bar" + v5)
                                        var v60 int = 42
                                        fmt.Println(v5, v5)
                                        _ = v60
                                    }
                                    var v61 int32 = -1
                                    _ = v61
                                }
                            default:
                                if v2 {
                                    fmt.Println(v3, v18)
                                    var v62 uint64 = 4294967295
                                    fmt.Println(fi57)
                                    var v63 bool = v2
                                    if v63 {
                                        switch v18 {
                                        case "a":
                                            if (v42 > 2147483647) {
                                                fmt.Println(v62, v5)
                                                v64 := GenS5{F0: -1, F1: (v62 << ((1) & 63)), F2: 113}
                                                _ = v64
                                            }
                                        case "hello":
                                            v65 := GenS3{F0: fi48}
                                            _ = v65
                                        default:
                                            var v66 uint32 = v42
                                            _ = v66
                                        }
                                        switch (fi48 | fi48) {
                                        case 9223372036854775807:
                                            if (v42 == v42) {
                                                v42 = (v42 ^ (v42 / ((v42) | 1)))
                                                fmt.Println(fi48, v6, fi48)
                                                var v67 int = (fi48 % ((1) | 1))
                                                v68 := GenS6{F0: v3, F1: (v42 % (((v42 / (((v42 >> ((1) & 63))) | 1))) | 1))}
                                                _ = v67
                                                _ = v68
                                            } else {
                                                v69 := GenS5{F0: 0, F1: (v62 + (v62 & (v62 & (v62 - v62)))), F2: 100}
                                                v1 = append(v1, (v5 + v18))
                                                _ = v69
                                            }
                                        }
                                        fmt.Println(v18, v6)
                                    }
                                    v5 = ""
                                    switch -1 {
                                    case -309:
                                        if v2 {
                                            _ = len(v1)
                                        } else {
                                            var v70 int = fi57
                                            _ = v70
                                        }
                                    case -2147483648:
                                        var v71 uint64 = v62
                                        _ = v71
                                    case 42:
                                        var v72 int16 = 1000
                                        _ = v72
                                    }
                                    _ = v62
                                    _ = v63
                                } else {
                                    var v73 int16 = -32768
                                    if v2 {
                                        v6 = v6
                                    }
                                    _ = v73
                                }
                            }
                            switch (fi57 % ((0) | 1)) {
                            case -2147483648:
                                switch ((v18 + "\n\t") + ((v5 + (v18 + "a")) + (v18 + v18))) {
                                case "hello":
                                    var v74 int = (fi48 % ((fi48) | 1))
                                    _ = v74
                                }
                            case -717:
                                if v2 {
                                    switch -1 {
                                    case -1:
                                        v18 = (("foo bar" + ("hello" + v5)) + v5)
                                        v75 := GenS7{F0: 857}
                                        _ = v75
                                    }
                                    var v76 uint64 = 328
                                    _ = v76
                                } else {
                                    var v77 int16 = genFn0()
                                    switch (fi48 + (fi57 % ((fi57) | 1))) {
                                    case -2147483648:
                                        v78 := GenS8{F0: 1, F1: v42}
                                        _ = v78
                                    case 0:
                                        var v79 float32 = (1.5 - 1.5)
                                        _ = v79
                                    }
                                    _ = v77
                                }
                                var v80 uint8 = 2
                                var v81 uint8 = (v80 >> ((255) & 63))
                                _ = v80
                                _ = v81
                            default:
                                if false {
                                    var v82 uint16 = (v4 >> ((42) & 63))
                                    _ = v82
                                }
                            }
                            var v83 float64 = 1e10
                            _ = fi57
                            _ = v83
                        }
                        v2 = v2
                        v1 = append(v1, ("" + ("foo bar" + v18)))
                    }
                    for fi84 := 0; fi84 < 11; fi84++ {
                        _ = len(v1)
                        if len(v1) > 0 {
                            v1[0] = v18
                        }
                        _ = fi84
                    }
                    _ = fi48
                }
            }
            _ = v46
            _ = v47
        } else {
            _ = len(v1)
            if true {
                v5 = (("" + v5) + v5)
            } else {
                v85 := GenS9{F0: 0}
                _ = v85
            }
        }
        v86 := GenS7{F0: uint(v42)}
        _ = v86
    }
    var v87 string = (v5 + "foo bar")
    fmt.Println(v2, v87)
    var v88 int32 = 397
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v87
    _ = v88
}
