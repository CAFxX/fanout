package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int; F1 bool }
func genFn0() int8 {
    switch "hello" {
    case "":
        if false {
            v6 := []int{933, 127}
            var v7 uint32 = 966
            var v8 float64 = 1e10
            _ = v6
            _ = v7
            _ = v8
        } else {
            var v9 int = 32768
            fmt.Println(v9)
            switch "" {
            case "a":
                v10 := GenS1{F0: (v9 + v9), F1: true}
                _ = v10
            case "foo bar":
                v11 := []int{(v9 | v9), 32768, 0}
                _ = v11
            case "hello":
                v9 = (v9 % ((v9) | 1))
            }
            _ = v9
        }
    default:
        if true {
            var v12 int = 755
            if (42 < 263) {
                v13 := GenS1{F0: (v12 << ((0) & 63)), F1: true}
                v12 += v12
                _ = v13
            }
            _ = v12
        } else {
            v14 := []int{-444}
            v15 := GenS1{F0: 0, F1: ((false || false) && (32767 < 2))}
            _ = v14
            _ = v15
        }
    }
    var v16 uint32 = 4294967295
    var v17 string = ("foo bar" + "a")
    v18 := GenS1{F0: 4294967295, F1: false}
    _ = v16
    _ = v17
    _ = v18
    return 0
}
func genFn1(p0 int16, p1 bool, p2 uint8) uint16 {
    p1 = (-1 < -371)
    p1 = false
    var v19 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 0
}

func main() {
    var v1 float64 = ((0.0 - (1e10 * -2.25)) * ((1e10 + 0.0) - 1.5))
    var v2 uint32 = 903
    var v3 int = 32768
    var v4 int16 = 0
    var v5 int32 = 2147483647
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 uint64 = 0
        v1 = 3.14159
        var v22 int32 = (v5 | 1)
        _ = fi20
        _ = v21
        _ = v22
    }
    v5 = (-(v5))
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint8 = 55
        v25 := []int{(v3 - -256), 27, 32767, v3}
        var v26 bool = (-1 > 127)
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    if true {
        for fi27 := 0; fi27 < 13; fi27++ {
            fmt.Println(v4, v5, fi27)
            v28 := GenS1{F0: (v3 | v3), F1: false}
            _ = fi27
            _ = v28
        }
    }
    switch v3 {
    case 9223372036854775807:
        if true {
            fmt.Println(v3, v5, v1)
            var v29 bool = (v5 > (v5 | (v5 - v5)))
            _ = v29
        } else {
            var v30 float32 = ((0.1 + 0.0) / 3.14159)
            v2 = 771
            _ = v30
        }
    case 42:
        var v31 uint16 = 1
        _ = v31
    default:
        var v32 uint16 = 583
        _ = v32
    }
    if (true || (!true)) {
        v33 := GenS1{F0: 127, F1: (((0.0 - 3.14159) < (0.0 / (3.14159 / (float32(v1) - (((-2.25 - 3.14159) - 0.0) / 1.0))))) || false)}
        var v34 uint32 = v2
        var v35 float32 = 0.1
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := GenS1{F0: (v3 & (v3 >> ((496) & 63))), F1: true}
    v37 := []int{}
    if ((-2.25 * 0.0) < 0.0) {
        if len(v37) > 0 {
            v37[0] = v3
        }
        var v38 bool = false
        _ = v38
    } else {
        switch (v3 & v3) {
        case 128:
            switch 127 {
            case -6:
                if true {
                    v4 -= v4
                } else {
                    v3 = (-(v3))
                }
            case 32767:
                var v39 uint16 = 65535
                _ = v39
            case -9223372036854775808:
                switch (-(v3)) {
                case 1:
                    v40 := GenS1{F0: (v3 - v3), F1: (!((!(!true)) && true))}
                    _ = v40
                case 0:
                    fmt.Println(v1)
                default:
                    v41 := GenS2{F0: 2, F1: v2}
                    _ = v41
                }
            }
        case 9223372036854775807:
            v2 -= (^(v2))
        case -1:
            if ("" == ("\n\t" + "\n\t")) {
                v1 = float64(0.1)
                for fi42 := 0; fi42 < 6; fi42++ {
                    v1 += v1
                    var v43 uint64 = 469
                    var v44 int16 = (v4 ^ -18)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            } else {
                v37 = append(v37, 128)
                fmt.Println(v4, v1, v1)
            }
        default:
            v4 = -1
        }
        if len(v37) > 0 {
            v37[0] = v3
        }
    }
    v45 := GenS3{F0: 9223372036854775807, F1: (v4 ^ v4)}
    for fi46 := 0; fi46 < 9; fi46++ {
        v3 += fi46
        _ = fi46
    }
    if (false || true) {
        switch "\n\t" {
        case "\n\t":
            var v47 uint16 = genFn1(v4, false, 166)
            _ = v47
        case "hello":
            if true {
                v5 = (v5 ^ v5)
            } else {
                if true {
                    if (1 >= 42) {
                        if (1.5 >= ((1.5 + v1) + ((0.1 * (v1 * (0.1 / v1))) * (float64(0.1) - ((v1 + ((3.14159 / 1.0) + v1)) * float64(((((((float32(0.1) * 3.14159) - (0.0 / ((3.14159 + (float32(0.0) / ((1.5 + 0.0) * 1.5))) + -2.25))) - 1.5) - -2.25) * -2.25) - 1.5))))))) {
                            v5 -= -34
                        } else {
                            v48 := GenS4{F0: "hello"}
                            v49 := GenS2{F0: 2, F1: (v2 >> ((42) & 63))}
                            v50 := GenS1{F0: (v3 >> ((711) & 63)), F1: (-9223372036854775808 != 0)}
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        v3 = v3
                    } else {
                        fmt.Println(v2)
                        v2 -= (v2 * (v2 & (v2 - v2)))
                    }
                    v51 := GenS1{F0: 2147483647, F1: (-1 >= 1)}
                    _ = v51
                } else {
                    v3 = v3
                    if false {
                        var v52 float32 = 3.14159
                        _ = v52
                    }
                }
                var v53 float32 = (1e10 / -2.25)
                _ = v53
            }
        }
    }
    v1 = v1
    for fi54 := 0; fi54 < 18; fi54++ {
        _ = len(v37)
        var v55 int32 = -958
        _ = fi54
        _ = v55
    }
    switch "hello" {
    case "\n\t":
        v56 := GenS4{F0: (("\n\t" + "") + "\n\t")}
        _ = v56
    }
    fmt.Println(v2, v3)
    if (4294967295 >= (v3 ^ -9223372036854775808)) {
        if (!((0.1 + 0.1) <= 0.0)) {
            v1 += v1
        } else {
            switch v3 {
            case -2147483648:
                for fi57 := 0; fi57 < 18; fi57++ {
                    v3 = fi57
                    var v58 int = 0
                    _ = fi57
                    _ = v58
                }
            case -82:
                v59 := GenS3{F0: 1, F1: -32768}
                _ = v59
            default:
                for fi60 := 0; fi60 < 6; fi60++ {
                    if true {
                        v61 := GenS1{F0: (fi60 - (v3 & (fi60 * fi60))), F1: ((32767 >= 0) && false)}
                        if (0 >= (v4 & v4)) {
                            var v62 uint = 18446744073709551615
                            _ = v62
                        }
                        v3 += -36
                        _ = v61
                    }
                    _ = fi60
                }
            }
            var v63 uint = 4294967295
            v37 = append(v37, 1)
            _ = v63
        }
    } else {
        if false {
            switch (v3 - (v3 * v3)) {
            case -492:
                if (!(true || true)) {
                    fmt.Println(v4)
                    var v64 float64 = 0.0
                    _ = v64
                } else {
                    var v65 float32 = ((1.5 - ((float32(v1) + 3.14159) * 3.14159)) / 1.0)
                    fmt.Println(v1)
                    v37 = append(v37, 802)
                    _ = v65
                }
            case 127:
                var v66 int = v3
                _ = v66
            }
            var v67 int8 = -77
            _ = v67
        }
        for fi68 := 0; fi68 < 10; fi68++ {
            for fi69 := 0; fi69 < 14; fi69++ {
                v70 := GenS5{F0: 127}
                v37 = append(v37, v3)
                v71 := GenS6{F0: 0, F1: "foo bar"}
                _ = fi69
                _ = v70
                _ = v71
            }
            v72 := GenS7{F0: (v3 - (^(fi68))), F1: true}
            var v73 int8 = -128
            _ = fi68
            _ = v72
            _ = v73
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v5
}
