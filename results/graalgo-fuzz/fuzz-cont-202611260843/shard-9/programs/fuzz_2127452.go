package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0() bool {
    var v4 string = "foo bar"
    if (60 == -79) {
        switch 701 {
        case 42:
            v5 := []int{567, 4294967295, 2147483647}
            _ = v5
        case 127:
            v6 := GenS1{F0: 42, F1: 1, F2: 1}
            _ = v6
        case 377:
            var v7 string = v4
            _ = v7
        }
    }
    v4 = "a"
    var v8 float64 = (3.14159 - float64(((0.1 / float32((0.0 + -2.25))) + (-2.25 + float32(3.14159)))))
    var v9 int8 = -52
    _ = v4
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]int)
    var v3 string = "foo bar"
    v10 := GenS2{F0: 60, F1: 2147483647, F2: true}
    var v11 uint = 18446744073709551615
    var v12 bool = true
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := []int{(fi13 + fi13)}
        _ = fi13
        _ = v14
    }
    var v15 int64 = 9223372036854775807
    switch v3 {
    case "hello":
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := GenS3{F0: uint32(v11), F1: 42, F2: 455}
                if (0.1 != 0.1) {
                    v19 := GenS4{F0: 18446744073709551615, F1: (v11 << (((^(v11))) & 63)), F2: 0}
                    if v12 {
                        fmt.Println(fi17, v12, v3)
                        v15 -= (-(v15))
                        v20 := []int{(fi16 ^ fi16)}
                        _ = v20
                    }
                    v21 := GenS5{F0: 1, F1: (v11 & (v11 | v11)), F2: v1}
                    v22 := GenS5{F0: 669, F1: (v11 | v11), F2: (v1 % ((v1) | 1))}
                    var v23 uint = (v11 & (v11 * (v11 * (v11 & 446))))
                    if (v12 && ((-(v15)) >= v15)) {
                        if ((v1 * 353) <= v1) {
                            fmt.Println(v11)
                            v24 := []int{fi17, fi17, (fi17 | fi16)}
                            _ = v24
                        } else {
                            if true {
                                if (-2147483648 >= -625) {
                                    fmt.Println(v1)
                                    var v25 int32 = 42
                                    _ = v25
                                }
                                var v26 int32 = 2147483647
                                v1 = (v1 >> ((v23) & 63))
                                _ = v26
                            }
                        }
                        switch fi16 {
                        case -352:
                            var v27 int16 = 1
                            _ = v27
                        }
                    }
                    _ = v19
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v28 string = v3
                _ = fi17
                _ = v18
                _ = v28
            }
            _ = fi16
        }
    case "":
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint = 943
            fmt.Println(fi29, v3)
            for fi31 := 0; fi31 < 3; fi31++ {
                if (!(false || (!((631 == 18446744073709551615) && v12)))) {
                    fmt.Println(fi31)
                }
                delete(v2, fi29)
                switch (fi29 * 0) {
                case 127:
                    var v32 float32 = 1e10
                    _ = v32
                default:
                    if v12 {
                        if v12 {
                            v33 := GenS4{F0: 0, F1: uint(v1), F2: 2147483647}
                            v3 = v3
                            if (((false && v12) && true) || (v12 && (1e10 == 0.1))) {
                                var v34 string = "a"
                                _ = v34
                            }
                            _ = len(v2)
                            _ = v33
                        } else {
                            if (v12 || (!v12)) {
                                v30 += uint(v1)
                                if (127 > -33) {
                                    var v35 int32 = -2147483648
                                    _ = v35
                                }
                                v2[(^(fi31))] = 127
                            } else {
                                v12 = (v12 || v12)
                                if (!true) {
                                    if true {
                                        var v36 int32 = -201
                                        var v37 int16 = -32768
                                        _ = v36
                                        _ = v37
                                    } else {
                                        if (false && (3.14159 > 1e10)) {
                                            var v38 int = (^(fi31))
                                            fmt.Println(fi29, v15)
                                            _ = v38
                                        }
                                    }
                                    v3 = v3
                                    if v12 {
                                        v1 = v1
                                        var v39 int16 = 32767
                                        var v40 bool = (v12 || ((fi29 >= 127) && (false && v12)))
                                        _ = v39
                                        _ = v40
                                    }
                                }
                                v12 = v12
                                v12 = (!v12)
                            }
                            var v41 uint64 = 42
                            var v42 bool = (!(((v30 < v30) && (0 != -523)) && (-2147483648 < -1)))
                            _ = v41
                            _ = v42
                        }
                        var v43 uint64 = 1
                        _ = v43
                    }
                }
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
    default:
        delete(v2, 32767)
    }
    v44 := []int{127, -388, -726, 42}
    v15 -= (v15 ^ v15)
    v12 = (v12 && v12)
    switch ("hello" + "") {
    case "":
        var v45 int8 = 45
        _ = v45
    }
    switch v3 {
    case "a":
        var v46 uint64 = 298
        _ = v46
    case "\n\t":
        if (false && ((!((0.0 < 1e10) && v12)) && ((!(1 > -433)) || false))) {
            v47 := GenS3{F0: 1, F1: 0, F2: 4294967295}
            fmt.Println(v3)
            if true {
                fmt.Println(v11, v3)
                delete(v2, 858)
                if (1 <= 0) {
                    v48 := GenS5{F0: 277, F1: (v11 ^ 145), F2: v1}
                    v49 := GenS4{F0: v11, F1: v11, F2: -72}
                    v50 := GenS5{F0: 869, F1: (v11 * 808), F2: (v1 >> ((1) & 63))}
                    _ = v48
                    _ = v49
                    _ = v50
                } else {
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                    _ = len(v44)
                }
            } else {
                _ = len(v2)
                v11 -= (^(v11))
            }
            _ = v47
        } else {
            for fi51 := 0; fi51 < 1; fi51++ {
                var v52 string = "hello"
                _ = fi51
                _ = v52
            }
            v53 := GenS3{F0: 2147483648, F1: 232, F2: 418}
            _ = v53
        }
    case "":
        fmt.Println(v1)
    default:
        if v12 {
            for fi54 := 0; fi54 < 8; fi54++ {
                delete(v2, (fi54 ^ fi54))
                v15 = 663
                _ = fi54
            }
            v15 = (v15 >> ((763) & 63))
            var v55 uint64 = 4294967295
            _ = v55
        }
    }
    v15 = v15
    _ = len(v44)
    for fi56 := 0; fi56 < 17; fi56++ {
        if (v12 || false) {
            for fi57 := 0; fi57 < 4; fi57++ {
                v11 = 42
                switch v3 {
                case "":
                    fmt.Println(v12, fi57)
                case "foo bar":
                    switch "hello" {
                    case "\n\t":
                        var v58 uint32 = 2147483647
                        _ = v58
                    }
                case "a":
                    if ((!(2 < uint8(v1))) && (!(!(0.0 <= (float32(0.1) + 0.0))))) {
                        var v59 uint64 = 115
                        var v60 uint8 = 1
                        _ = v59
                        _ = v60
                    }
                }
                v61 := GenS6{F0: -1, F1: v11}
                fmt.Println(v15, v11)
                var v62 int64 = (v15 & (v15 >> ((34) & 63)))
                _ = fi57
                _ = v61
                _ = v62
            }
        } else {
            if (float64(1.5) <= 3.14159) {
                fmt.Println(v3, v15, v15)
            }
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v44
}
