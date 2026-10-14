package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 bool, p2 uint64) int8 {
    switch 1 {
    case 128:
        var v7 string = ("foo bar" + "\n\t")
        _ = v7
    case 2147483647:
        v8 := GenS1{F0: (32767 != 336)}
        _ = v8
    case 127:
        var v9 uint64 = 18446744073709551615
        _ = v9
    default:
        var v10 bool = (1.5 < 1e10)
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: p1}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if false {
            var v14 int8 = 35
            var v15 uint8 = 76
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    if false {
        v16 := []int{32768}
        if false {
            p1 = p1
            fmt.Println(p2, p0)
        }
        _ = v16
    } else {
        fmt.Println(p1, p1)
    }
    p0 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    return -128
}
func genFn1(p0 uint) uint32 {
    switch -2147483648 {
    case 4294967295:
        p0 -= p0
    }
    if ((0.0 * (0.1 - 0.0)) <= 1e10) {
        var v17 uint16 = 2
        if (42 < 627) {
            if false {
                if true {
                    switch 4294967295 {
                    case 971:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            switch 42 {
                            case 1:
                                var v19 uint8 = 2
                                _ = v19
                            case -9223372036854775808:
                                var v20 uint32 = uint32(p0)
                                _ = v20
                            case -661:
                                v21 := []int{}
                                _ = v21
                            }
                            v22 := []int{4294967295}
                            v23 := GenS2{F0: -143, F1: 42, F2: "\n\t"}
                            _ = fi18
                            _ = v22
                            _ = v23
                        }
                    case 4294967295:
                        var v24 int8 = 1
                        _ = v24
                    case 868:
                        p0 = p0
                    default:
                        fmt.Println(v17, v17)
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        p0 -= p0
                        v17 = (v17 % ((v17) | 1))
                        v26 := []int{32767, (^(fi25)), 2147483647}
                        _ = fi25
                        _ = v26
                    }
                    v27 := []int{}
                    _ = v27
                } else {
                    p0 = (^(p0))
                    switch -270 {
                    case 9223372036854775807:
                        if (32767 > int16(p0)) {
                            if false {
                                var v28 uint16 = v17
                                if ("foo bar" == "") {
                                    if (!true) {
                                        v28 += (^(v17))
                                        if (1.5 > 3.14159) {
                                            v29 := []int{-693, 9223372036854775807, 354}
                                            _ = v29
                                        } else {
                                            v30 := []int{233, -798, 127, 1}
                                            fmt.Println(p0, v17)
                                            _ = v30
                                        }
                                        v31 := []int{-373}
                                        _ = v31
                                    }
                                }
                                v32 := GenS3{F0: 57}
                                _ = v28
                                _ = v32
                            }
                        }
                    case -634:
                        var v33 int8 = 81
                        _ = v33
                    default:
                        v34 := GenS1{F0: (32767 >= 374)}
                        _ = v34
                    }
                }
            }
            v35 := []int{879, -2147483648, 431}
            v36 := GenS2{F0: -32768, F1: -353, F2: "hello"}
            _ = v35
            _ = v36
        }
        _ = v17
    }
    fmt.Println(p0)
    _ = p0
    return 2147483647
}

func main() {
    v1 := []string{"foo bar", ""}
    var v2 int64 = -388
    v3 := []uint32{100, 889, 1, 2147483648}
    var v4 int32 = -2147483648
    var v5 int64 = -555
    var v6 uint16 = 0
    v3 = append(v3, 185)
    switch "a" {
    case "":
        v37 := GenS2{F0: 519, F1: -1, F2: ("foo bar" + ("a" + ""))}
        _ = v37
    case "\n\t":
        var v38 int = 128
        _ = v38
    default:
        var v39 float32 = float32((((0.1 * 0.0) + -2.25) + 0.0))
        _ = v39
    }
    v40 := GenS3{F0: 1}
    v2 = (v2 % (((v2 << ((uint(v6)) & 63))) | 1))
    switch ("a" + ("" + ("a" + "hello"))) {
    case "foo bar":
        var v41 int32 = v4
        _ = v41
    case "":
        var v42 uint8 = 104
        _ = v42
    case "\n\t":
        v4 = v4
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v43 bool = (("" == "a") && (!((!(((3.14159 - float32(((-2.25 * -2.25) - -2.25))) <= 0.0) || false)) && (!true))))
    if (4294967295 == 492) {
        for fi44 := 0; fi44 < 11; fi44++ {
            fmt.Println(v6, v6, fi44)
            _ = fi44
        }
        var v45 int16 = -1
        v46 := GenS3{F0: 265}
        _ = v45
        _ = v46
    }
    v47 := GenS4{F0: v43, F1: (v6 - v6), F2: 42}
    v5 = (v5 / ((9223372036854775807) | 1))
    var v48 uint16 = v6
    if (v5 != (v5 << ((65535) & 63))) {
        fmt.Println(v5)
    } else {
        v49 := GenS2{F0: 0, F1: -358, F2: "a"}
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
