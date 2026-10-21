package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint32; F2 int32 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int; F1 int; F2 int64 }
type GenS8 struct { F0 uint16; F1 int32; F2 int16 }
func genFn0(p0 bool, p1 uint32) uint32 {
    v7 := GenS1{F0: 42, F1: 0, F2: 160}
    v8 := []int{194, 9223372036854775807, -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 26
}
func genFn1(p0 int64, p1 uint8, p2 bool) int8 {
    var v9 uint64 = 315
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 0
    v2 := []string{("" + "hello"), "a", "hello", ("hello" + "")}
    v3 := []string{"\n\t"}
    var v4 int64 = 42
    var v5 uint64 = v1
    v6 := make(map[uint32]string)
    v1 -= (^(v1))
    var v10 uint8 = 251
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 17; fi12++ {
        v5 += v1
        _ = fi12
    }
    switch -2147483648 {
    case 42:
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint8 = (^(v10))
            switch 4294967295 {
            case -1:
                v1 = 0
            case 0:
                for fi15 := 0; fi15 < 13; fi15++ {
                    if ((fi15 + 42) < 32767) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v11 = 289
                    }
                    if (660 <= -2147483648) {
                        v2 = append(v2, "foo bar")
                    } else {
                        v16 := v6[2147483647]
                        var v17 uint64 = (^(v5))
                        _ = v16
                        _ = v17
                    }
                    _ = fi15
                }
            default:
                var v18 uint64 = 42
                _ = v18
            }
            fmt.Println(v4, fi13)
            _ = fi13
            _ = v14
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v5, v1)
        _ = fi19
    }
    if false {
        v20 := GenS2{F0: (v5 / ((uint64(v10)) | 1)), F1: -25, F2: 2147483648}
        v11 = v4
        _ = v20
    } else {
        v21 := GenS1{F0: v10, F1: -2147483648, F2: v10}
        _ = v21
    }
    v22 := v6[0]
    for fi23 := 0; fi23 < 17; fi23++ {
        v6[2] = v22
        _ = fi23
    }
    v24 := GenS3{F0: 65535}
    if false {
        var v25 bool = (!true)
        for fi26 := 0; fi26 < 5; fi26++ {
            if (-284 <= v11) {
                var v27 uint32 = 903
                _ = v27
            } else {
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 uint16 = 255
                    if (genFn1(v4, v10, ((!((!(((!v25) && (4294967295 != 1)) && (-912 <= -1))) && v25)) || v25)) == -128) {
                        fmt.Println(v25, v10, fi26)
                        fmt.Println(v25, v5, v5)
                        v30 := GenS2{F0: 150, F1: 65, F2: 4294967295}
                        _ = v30
                    } else {
                        if false {
                            var v31 uint64 = (v1 & v5)
                            _ = v31
                        } else {
                            switch fi28 {
                            case 0:
                                v10 = v10
                            case 42:
                                v1 = v5
                            }
                            var v32 float32 = 0.1
                            _ = v32
                        }
                        fmt.Println(v5)
                        switch 128 {
                        case 0:
                            if ((^(v29)) == (v29 << ((255) & 63))) {
                                _ = len(v3)
                            } else {
                                v33 := GenS4{F0: fi26, F1: 42, F2: v29}
                                switch v22 {
                                case "":
                                    fmt.Println(v10, fi28, v1)
                                default:
                                    var v34 uint32 = genFn0(true, uint32(v10))
                                    _ = v34
                                }
                                switch (v22 + "\n\t") {
                                case "\n\t":
                                    switch (fi28 | fi26) {
                                    case 32:
                                        var v35 bool = v25
                                        _ = v35
                                    case -65:
                                        var v36 int16 = -1
                                        v1 += v1
                                        _ = v36
                                    }
                                    fmt.Println(v25, v10)
                                case "a":
                                    if v25 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = len(v2)
                                    }
                                }
                                _ = v33
                            }
                        case 1:
                            _ = len(v6)
                        default:
                            switch -2147483648 {
                            case 32768:
                                v6[786] = ((v22 + ("\n\t" + ((((("" + v22) + "") + v22) + (v22 + v22)) + v22))) + (v22 + v22))
                            case -87:
                                delete(v6, 445)
                            case 9223372036854775807:
                                v37 := GenS1{F0: v10, F1: 42, F2: (v10 << ((1) & 63))}
                                _ = v37
                            default:
                                if (!((v22 + "") > v22)) {
                                    v25 = v25
                                }
                                v38 := GenS3{F0: 2}
                                v39 := GenS5{F0: (v10 >= v10), F1: genFn0(v25, 1), F2: -982}
                                _ = v38
                                _ = v39
                            }
                        }
                        var v40 uint32 = genFn0(((v25 && (v22 == v22)) || true), 823)
                        _ = len(v6)
                        _ = v40
                    }
                    v41 := GenS4{F0: (fi26 & (fi28 + 32768)), F1: uint(v10), F2: (v29 / ((1) | 1))}
                    _ = fi28
                    _ = v29
                    _ = v41
                }
            }
            v42 := GenS5{F0: ((3.14159 / 1e10) == 3.14159), F1: 2147483647, F2: 523}
            if false {
                if true {
                    for fi43 := 0; fi43 < 3; fi43++ {
                        _ = len(v6)
                        _ = fi43
                    }
                    fmt.Println(v10, fi26)
                    switch v22 {
                    case "foo bar":
                        var v44 uint16 = 2
                        _ = v44
                    case "hello":
                        var v45 int16 = 1
                        _ = v45
                    default:
                        v22 = (((v22 + "a") + "\n\t") + ("foo bar" + (v22 + "a")))
                    }
                } else {
                    for fi46 := 0; fi46 < 4; fi46++ {
                        if len(v2) > 0 {
                            v2[0] = v22
                        }
                        v47 := GenS5{F0: v25, F1: 0, F2: 861}
                        v22 = ("foo bar" + "\n\t")
                        _ = fi46
                        _ = v47
                    }
                    v3 = append(v3, (("a" + v22) + v22))
                    var v48 uint32 = 0
                    _ = v48
                }
            } else {
                for fi49 := 0; fi49 < 9; fi49++ {
                    switch (fi26 << ((599) & 63)) {
                    case 38:
                        var v50 uint32 = 4294967295
                        _ = v50
                    default:
                        var v51 int8 = 0
                        _ = v51
                    }
                    _ = fi49
                }
                var v52 int32 = 1
                _ = v52
            }
            _ = fi26
            _ = v42
        }
        _ = v25
    } else {
        for fi53 := 0; fi53 < 16; fi53++ {
            switch (((("foo bar" + "a") + ((v22 + v22) + (v22 + v22))) + (v22 + v22)) + v22) {
            case "foo bar":
                fmt.Println(v10)
            case "hello":
                for fi54 := 0; fi54 < 2; fi54++ {
                    switch ("foo bar" + (("hello" + v22) + (v22 + v22))) {
                    case "foo bar":
                        switch v22 {
                        case "a":
                            if true {
                                switch 42 {
                                case -2147483648:
                                    v55 := GenS6{F0: (^(v11)), F1: 42}
                                    _ = v55
                                case 127:
                                    if false {
                                        switch (fi53 - fi53) {
                                        case -397:
                                            v6[uint32(v10)] = "hello"
                                        case 4294967295:
                                            var v56 float32 = (((0.0 / 1e10) + 1e10) + 3.14159)
                                            _ = v56
                                        case 9223372036854775807:
                                            var v57 int16 = 42
                                            _ = v57
                                        default:
                                            var v58 string = v22
                                            _ = v58
                                        }
                                    }
                                case 42:
                                    v11 = (v11 + v11)
                                }
                            }
                        case "foo bar":
                            if true {
                                switch fi53 {
                                case 9223372036854775807:
                                    fmt.Println(v22)
                                case 127:
                                    var v59 int32 = 0
                                    _ = v59
                                }
                            } else {
                                fmt.Println(v1, v1)
                                v60 := GenS1{F0: v10, F1: -239, F2: v10}
                                fmt.Println(v11, v1, fi53)
                                _ = v60
                            }
                        }
                    }
                    var v61 uint8 = v10
                    fmt.Println(v10, v5)
                    var v62 int32 = 0
                    _ = fi54
                    _ = v61
                    _ = v62
                }
            case "\n\t":
                var v63 uint64 = (v5 - v1)
                _ = v63
            }
            _ = fi53
        }
        var v64 uint16 = 255
        _ = v64
    }
    for fi65 := 0; fi65 < 17; fi65++ {
        switch (fi65 / (((fi65 * (fi65 * fi65))) | 1)) {
        case 9223372036854775807:
            if (!(!(("hello" + (v22 + "hello")) == v22))) {
                v66 := GenS7{F0: 32767, F1: fi65, F2: int64(fi65)}
                _ = v66
            }
        default:
            v10 = v10
        }
        var v67 uint64 = (v5 * (v1 & (v1 >> ((uint(v1)) & 63))))
        v68 := GenS8{F0: 0, F1: 0, F2: 1000}
        _ = fi65
        _ = v67
        _ = v68
    }
    if (515 <= 2147483647) {
        for fi69 := 0; fi69 < 17; fi69++ {
            switch "\n\t" {
            case "foo bar":
                var v70 uint = 4294967295
                _ = v70
            case "hello":
                v6[240] = v22
            default:
                var v71 uint32 = 865
                _ = v71
            }
            var v72 uint = 255
            v11 = (v4 / (((v4 + (v4 | (v4 & v4)))) | 1))
            _ = fi69
            _ = v72
        }
        var v73 uint16 = 65535
        var v74 uint32 = 65535
        _ = v73
        _ = v74
    }
    fmt.Println(v10)
    v5 = 1
    v4 += (v4 ^ v11)
    for fi75 := 0; fi75 < 19; fi75++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (2147483647 == 63) {
            v11 += (^(v11))
            for fi76 := 0; fi76 < 1; fi76++ {
                var v77 uint32 = 2
                var v78 float64 = 0.0
                switch v22 {
                case "hello":
                    v78 += 3.14159
                case "foo bar":
                    v79 := GenS1{F0: v10, F1: 719, F2: (^(v10))}
                    _ = v79
                }
                _ = fi76
                _ = v77
                _ = v78
            }
            v1 += 523
        }
        fmt.Println(v1, v11)
        _ = fi75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
