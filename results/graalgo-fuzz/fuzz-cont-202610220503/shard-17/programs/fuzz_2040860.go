package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int; F2 uint16 }
func genFn0(p0 int64, p1 string) int32 {
    var v7 uint16 = 32767
    v8 := GenS1{F0: -892}
    v9 := GenS1{F0: -2147483648}
    var v10 uint16 = v7
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if (true && (2147483648 == 42)) {
                v13 := GenS1{F0: 2147483647}
                _ = v13
            } else {
                var v14 uint16 = 683
                v15 := []int{0}
                switch p1 {
                case "":
                    switch (p1 + p1) {
                    case "a":
                        v16 := GenS2{F0: 162}
                        _ = v16
                    case "\n\t":
                        if (!true) {
                            fmt.Println(fi11, p1)
                            switch ("a" + p1) {
                            case "a":
                                if len(v15) > 0 {
                                    v15[0] = fi11
                                }
                            default:
                                switch 599 {
                                case 128:
                                    if len(v15) > 0 {
                                        v15[0] = fi12
                                    }
                                }
                            }
                            _ = len(v15)
                            var v17 uint64 = 42
                            v18 := GenS2{F0: 127}
                            _ = len(v15)
                            if (float64(1.5) >= 1e10) {
                                p1 = "a"
                                var v19 int8 = 42
                                if (2147483647 >= -143) {
                                    if (717 != 1) {
                                        p1 = p1
                                    } else {
                                        var v20 int64 = 9223372036854775807
                                        v17 = v17
                                        switch "foo bar" {
                                        case "a":
                                            var v21 int8 = (v19 - (v19 & v19))
                                            _ = v21
                                        default:
                                            switch 869 {
                                            case 32768:
                                                var v22 int16 = 429
                                                _ = v22
                                            }
                                        }
                                        _ = v20
                                    }
                                    var v23 uint8 = 127
                                    _ = v23
                                }
                                _ = v19
                            }
                            _ = v17
                            _ = v18
                        }
                    case "foo bar":
                        if false {
                            v15 = append(v15, fi12)
                            var v24 float32 = (0.0 - 1.5)
                            _ = v24
                        } else {
                            if true {
                                var v25 float32 = (3.14159 / (1e10 / 1.0))
                                if (false && true) {
                                    fmt.Println(fi12)
                                }
                                var v26 float32 = 1e10
                                fmt.Println(fi12, v7, p1)
                                _ = v25
                                _ = v26
                            }
                            if false {
                                v27 := GenS2{F0: 255}
                                fmt.Println(v7)
                                var v28 float32 = 1.5
                                v29 := GenS1{F0: -1}
                                v10 += v7
                                var v30 int64 = (p0 / (((p0 >> ((42) & 63))) | 1))
                                var v31 uint32 = 4294967295
                                if (1e10 > 3.14159) {
                                    var v32 int16 = -1
                                    fmt.Println(v10, v31)
                                    var v33 int = fi12
                                    _ = v32
                                    _ = v33
                                }
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            var v34 int64 = 0
                            if len(v15) > 0 {
                                v15[0] = (fi12 + fi11)
                            }
                            v34 = (^(p0))
                            _ = v34
                        }
                    default:
                        _ = len(v15)
                    }
                default:
                    v10 = (^(v10))
                }
                _ = v14
                _ = v15
            }
            fmt.Println(fi11)
            _ = fi12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 152
    var v2 uint8 = v1
    var v3 uint64 = 349
    var v4 int8 = -128
    v5 := []int{127, -9223372036854775808, -2147483648, 2147483647}
    var v6 bool = (!true)
    v2 += 0
    if (("foo bar" + "") == "a") {
        for fi35 := 0; fi35 < 3; fi35++ {
            switch "hello" {
            case "":
                v5 = append(v5, (fi35 / ((-1) | 1)))
            }
            var v36 int32 = 42
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 5; fi37++ {
            if false {
                switch (("a" + "a") + "foo bar") {
                case "hello":
                    v38 := GenS1{F0: -446}
                    _ = v38
                case "a":
                    v1 = 0
                case "foo bar":
                    var v39 string = ""
                    _ = v39
                default:
                    v40 := GenS3{F0: 315, F1: -2147483648}
                    _ = v40
                }
            } else {
                for fi41 := 0; fi41 < 5; fi41++ {
                    if true {
                        var v42 int = fi41
                        if v6 {
                            switch (fi41 | (fi37 | (v42 | 0))) {
                            case 86:
                                if len(v5) > 0 {
                                    v5[0] = (^(v42))
                                }
                            default:
                                if v6 {
                                    if (!(("foo bar" + "") <= ("" + (("hello" + ((((("\n\t" + "\n\t") + "a") + "hello") + ("\n\t" + "\n\t")) + "hello")) + "a")))) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        fmt.Println(v4, v4, v42)
                                        v43 := GenS1{F0: 1}
                                        _ = v43
                                    }
                                    var v44 uint8 = 0
                                    _ = v44
                                } else {
                                    var v45 int16 = 1000
                                    _ = v45
                                }
                            }
                            var v46 string = "hello"
                            v2 -= (^(v1))
                            _ = v46
                        }
                        var v47 int = fi37
                        _ = v42
                        _ = v47
                    }
                    _ = fi41
                }
                fmt.Println(v3, v4, v2)
                v2 += 1
            }
            _ = fi37
        }
        for fi48 := 0; fi48 < 9; fi48++ {
            fmt.Println(v1)
            if (v4 > (v4 % (((v4 >> ((1) & 63))) | 1))) {
                if len(v5) > 0 {
                    v5[0] = (^(fi48))
                }
                v49 := GenS3{F0: 42, F1: fi48}
                _ = v49
            }
            if v6 {
                v4 -= 127
                var v50 uint32 = 42
                var v51 uint = 18446744073709551615
                _ = v50
                _ = v51
            }
            _ = fi48
        }
    } else {
        fmt.Println(v6, v1, v1)
        var v52 uint16 = 2
        if true {
            var v53 uint = 166
            _ = v53
        }
        _ = v52
    }
    if (422 > -9223372036854775808) {
        fmt.Println(v1)
        var v54 string = (("\n\t" + "") + "")
        _ = v54
    } else {
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        fmt.Println(v4, v1)
        if len(v5) > 0 {
            v5[0] = -590
        }
    }
    var v55 int32 = int32(v4)
    v56 := GenS4{F0: v1, F1: v2, F2: 2147483647}
    var v57 uint8 = v1
    v58 := GenS1{F0: -1}
    if (!v6) {
        v59 := GenS1{F0: (v55 % (((v55 + (v55 % ((-2147483648) | 1)))) | 1))}
        var v60 bool = (v3 <= (v3 + 4294967295))
        v61 := GenS5{F0: genFn0(-170, "a"), F1: 42, F2: v1}
        _ = v59
        _ = v60
        _ = v61
    }
    var v62 uint8 = (v1 ^ (^(v2)))
    for fi63 := 0; fi63 < 17; fi63++ {
        if (!true) {
            var v64 uint8 = v2
            for fi65 := 0; fi65 < 7; fi65++ {
                var v66 float64 = float64(-2.25)
                var v67 string = "hello"
                switch v67 {
                case "\n\t":
                    if len(v5) > 0 {
                        v5[0] = int(v55)
                    }
                case "hello":
                    v67 = v67
                case "":
                    if v6 {
                        if (50 != 253) {
                            if (1 == 4294967295) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v68 uint32 = 65535
                                if (!false) {
                                    v69 := GenS6{F0: -1}
                                    _ = v69
                                } else {
                                    var v70 float64 = float64(3.14159)
                                    v3 -= (v3 - v3)
                                    v5 = append(v5, (fi63 % ((fi65) | 1)))
                                    _ = v70
                                }
                                v57 += (^(v62))
                                switch v67 {
                                case "a":
                                    if (v6 || (!true)) {
                                        v64 -= v2
                                        if (false && (!v6)) {
                                            v57 += 127
                                            if len(v5) > 0 {
                                                v5[0] = fi65
                                            }
                                            switch v67 {
                                            case "":
                                                var v71 uint32 = (^(v68))
                                                if v6 {
                                                    v3 += v3
                                                    fmt.Println(v62, v57, fi65)
                                                }
                                                v5 = append(v5, (^(fi65)))
                                                _ = v71
                                            }
                                        } else {
                                            v1 = (v1 | uint8(v3))
                                        }
                                    }
                                }
                                var v72 bool = v6
                                var v73 uint32 = (v68 * 187)
                                _ = v68
                                _ = v72
                                _ = v73
                            } else {
                                v74 := GenS2{F0: v57}
                                fmt.Println(v64, v2, v66)
                                v75 := GenS2{F0: 255}
                                _ = v74
                                _ = v75
                            }
                            switch fi65 {
                            case -625:
                                v2 += 183
                            }
                            fmt.Println(fi65, v67, v2)
                        } else {
                            _ = len(v5)
                            var v76 int32 = v55
                            _ = v76
                        }
                        if (v1 == v2) {
                            v6 = ((!(v4 == (v4 ^ (-(v4))))) || ((1e10 / 1e10) != 3.14159))
                            v77 := GenS7{F0: (fi63 + (fi63 | (fi65 * fi63))), F1: fi65, F2: 372}
                            switch (fi63 / ((fi63) | 1)) {
                            case -2147483648:
                                var v78 int16 = int16(v3)
                                _ = v78
                            default:
                                fmt.Println(v2)
                            }
                            _ = v77
                        }
                        v79 := GenS5{F0: genFn0(0, v67), F1: 0, F2: v2}
                        fmt.Println(v66, fi65, fi65)
                        var v80 uint32 = 2
                        _ = v79
                        _ = v80
                    }
                default:
                    fmt.Println(v2)
                }
                v81 := GenS2{F0: 127}
                _ = fi65
                _ = v66
                _ = v67
                _ = v81
            }
            v82 := GenS2{F0: 156}
            _ = v64
            _ = v82
        } else {
            v83 := GenS5{F0: genFn0(-802, "foo bar"), F1: 0, F2: v62}
            _ = v83
        }
        v2 -= v57
        _ = fi63
    }
    v84 := GenS4{F0: v57, F1: (v1 | v57), F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v62
    _ = v84
}
