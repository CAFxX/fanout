package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint; F2 int32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 string; F1 int8; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint16; F2 bool }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint16, p1 int, p2 int8) uint32 {
    v7 := []int{p1, -573, int(p0), -9223372036854775808}
    var v8 uint64 = 725
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint64 {
    var v9 int = 128
    var v10 float32 = -2.25
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint32 = 2147483647
            _ = fi11
            _ = v12
        }
    } else {
        v9 -= (v9 - v9)
        if true {
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 string = ("\n\t" + "")
                fmt.Println(v14)
                fmt.Println(v9, v10, v14)
                _ = fi13
                _ = v14
            }
        }
        switch v9 {
        case 0:
            fmt.Println(v10, v10)
        case 1:
            fmt.Println(v9)
        case -2147483648:
            var v15 float32 = (-2.25 / v10)
            _ = v15
        }
    }
    switch "\n\t" {
    case "":
        v16 := []int{(v9 >> ((255) & 63))}
        _ = v16
    case "hello":
        v9 += (-(v9))
    case "a":
        for fi17 := 0; fi17 < 15; fi17++ {
            if (4294967295 < 974) {
                v18 := GenS1{F0: 610, F1: (v9 * (fi17 % ((42) | 1))), F2: 2}
                v19 := []int{0}
                switch ("a" + "\n\t") {
                case "foo bar":
                    switch ("hello" + "hello") {
                    case "a":
                        switch (fi17 << ((1) & 63)) {
                        case 407:
                            v9 = fi17
                        case 399:
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v9 = fi20
                                _ = fi20
                            }
                        case 4294967295:
                            v10 += v10
                        }
                    }
                case "\n\t":
                    v21 := GenS2{F0: ("foo bar" + "\n\t")}
                    _ = v21
                case "":
                    var v22 uint = 1
                    _ = v22
                default:
                    if (!true) {
                        var v23 uint64 = 0
                        fmt.Println(v23, v10)
                        _ = v23
                    }
                }
                _ = v18
                _ = v19
            } else {
                v10 += v10
                v24 := GenS1{F0: 65535, F1: (fi17 / (((v9 + (fi17 & (v9 * fi17)))) | 1)), F2: 2}
                _ = v24
            }
            var v25 float32 = ((float32(0.1) / v10) + v10)
            var v26 int64 = -934
            _ = fi17
            _ = v25
            _ = v26
        }
    }
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 33
    var v2 uint64 = 1
    var v3 uint64 = v2
    v4 := make(map[string]string)
    var v5 int16 = 111
    var v6 uint8 = 209
    v27 := GenS3{F0: (v1 - v1), F1: 18446744073709551615, F2: -818}
    if true {
        v28 := []int{}
        _ = v28
    }
    var v29 int = 378
    var v30 uint64 = 0
    if (!true) {
        v3 = v30
        var v31 int = (v29 / (((v29 & v29)) | 1))
        _ = v31
    }
    v30 -= v3
    fmt.Println(v1)
    if (3.14159 == 0.1) {
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 12; fi33++ {
                if ((-2.25 / (float32(((0.0 + float64((1.5 / 1e10))) / (-2.25 * ((-2.25 / (0.1 + (0.0 + (((-2.25 - ((((1e10 * 3.14159) * -2.25) - -2.25) - float64(((0.1 / (3.14159 / 1.5)) + 1e10)))) - float64(0.1)) + (3.14159 - (1e10 + 0.1)))))) + (0.1 + 0.0))))) / 1.0)) >= ((float32(1e10) / -2.25) + -2.25)) {
                    fmt.Println(v3, v2, fi32)
                    v34 := GenS1{F0: (v1 | (v1 >> ((1) & 63))), F1: (-(v29)), F2: genFn1()}
                    _ = v34
                }
                var v35 uint16 = 2
                v5 = v5
                switch (fi32 % (((fi32 - fi32)) | 1)) {
                case 128:
                    v30 += (v3 % ((42) | 1))
                case -9223372036854775808:
                    v36 := GenS4{F0: (v29 + 290), F1: 100}
                    _ = v36
                }
                v37 := []int{}
                switch "a" {
                case "":
                    var v38 int = (fi33 / (((^(fi32))) | 1))
                    _ = v38
                case "a":
                    if false {
                        if true {
                            fmt.Println(fi32, v29)
                        }
                    } else {
                        var v39 float64 = (0.1 - 3.14159)
                        var v40 int32 = -1
                        _ = v39
                        _ = v40
                    }
                }
                _ = len(v37)
                _ = fi33
                _ = v35
                _ = v37
            }
            v41 := []int{}
            var v42 float64 = 1e10
            _ = fi32
            _ = v41
            _ = v42
        }
        v43 := GenS1{F0: 631, F1: v29, F2: (v2 % (((v2 / (((v3 & v3)) | 1))) | 1))}
        _ = v43
    } else {
        v6 += (v6 + v6)
    }
    fmt.Println(v2, v2, v29)
    if (("\n\t" + ("hello" + "a")) <= ("\n\t" + "a")) {
        switch v29 {
        case 887:
            _ = len(v4)
        default:
            v6 -= (v6 >> ((421) & 63))
        }
    } else {
        v44 := v4["foo bar"]
        for fi45 := 0; fi45 < 16; fi45++ {
            if true {
                if (!true) {
                    v46 := GenS3{F0: 663, F1: 0, F2: -1}
                    _ = v46
                }
                var v47 float64 = float64(1.5)
                for fi48 := 0; fi48 < 4; fi48++ {
                    v49 := GenS5{F0: v44, F1: 100, F2: 42}
                    v50 := GenS2{F0: "hello"}
                    var v51 uint8 = 42
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = v47
            }
            v30 += (v30 / ((v30) | 1))
            if true {
                v3 -= v3
                if true {
                    v44 = (((("" + (((("\n\t" + "") + "foo bar") + (v44 + ("" + v44))) + "a")) + v44) + (v44 + v44)) + v44)
                    v2 = (v2 << ((4294967295) & 63))
                    var v52 int32 = 0
                    _ = v52
                } else {
                    v2 = (^(v3))
                    v53 := GenS6{F0: 100}
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v55 := GenS2{F0: "foo bar"}
                        _ = fi54
                        _ = v55
                    }
                    _ = v53
                }
                v29 = fi45
            } else {
                var v56 uint64 = v2
                var v57 int = (fi45 ^ (-(v29)))
                var v58 int = (v29 - 1)
                _ = v56
                _ = v57
                _ = v58
            }
            _ = fi45
        }
        if (true || (-5 == 0)) {
            if true {
                for fi59 := 0; fi59 < 12; fi59++ {
                    var v60 uint = 18446744073709551615
                    var v61 int8 = 42
                    var v62 float32 = -2.25
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                }
                v63 := []int{(v29 * v29), 32767}
                v3 = (v3 >> ((uint(v1)) & 63))
                _ = v63
            }
            v5 -= (v5 / ((32767) | 1))
        }
        _ = v44
    }
    if false {
        fmt.Println(v3)
        v64 := []int{(v29 - (v29 | (^(v29)))), v29, (v29 | v29), -563}
        _ = v64
    } else {
        fmt.Println(v5)
    }
    fmt.Println(v6, v29, v30)
    if true {
        v5 = (v5 * 32767)
    }
    v65 := []int{}
    if true {
        for fi66 := 0; fi66 < 5; fi66++ {
            if len(v65) > 0 {
                v65[0] = int(v5)
            }
            _ = fi66
        }
        v29 += (v29 & v29)
        v67 := GenS7{F0: "\n\t", F1: 1, F2: (1 < 403)}
        _ = v67
    }
    v68 := GenS8{F0: 446}
    var v69 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v68
    _ = v69
}
