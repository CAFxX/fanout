package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 bool; F1 string; F2 int32 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int16; F1 int16; F2 uint }
type GenS9 struct { F0 bool }
func genFn0(p0 bool) bool {
    if ((789 == 32767) || p0) {
        v7 := GenS1{F0: 4294967295, F1: (float64((float32(1.5) * -2.25)) < 0.1)}
        _ = v7
    } else {
        v8 := GenS2{F0: 2147483648}
        p0 = (!p0)
        var v9 int32 = -2147483648
        _ = v8
        _ = v9
    }
    switch 32767 {
    case 128:
        var v10 string = "foo bar"
        _ = v10
    case 1:
        v11 := GenS2{F0: 128}
        _ = v11
    case 9223372036854775807:
        if (!("" > "\n\t")) {
            var v12 int64 = 0
            var v13 string = ("a" + "foo bar")
            _ = v12
            _ = v13
        } else {
            fmt.Println(p0)
            for fi14 := 0; fi14 < 12; fi14++ {
                if (18446744073709551615 > 64) {
                    var v15 int16 = 1000
                    _ = v15
                } else {
                    var v16 uint32 = 4294967295
                    _ = v16
                }
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := []int{(-(fi17)), fi17, (fi17 << ((2) & 63))}
                    var v19 int8 = 42
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi14
            }
            if p0 {
                for fi20 := 0; fi20 < 13; fi20++ {
                    p0 = ((!(p0 || p0)) || p0)
                    v21 := GenS3{F0: (!(((38 != 127) && true) || p0))}
                    _ = fi20
                    _ = v21
                }
                switch 127 {
                case 9223372036854775807:
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := GenS4{F0: 0, F1: -1}
                        switch ("a" + ("a" + "")) {
                        case "":
                            for fi24 := 0; fi24 < 19; fi24++ {
                                p0 = false
                                if p0 {
                                    v25 := []int{fi24, (fi24 ^ 9223372036854775807), fi22, fi22}
                                    p0 = true
                                    _ = v25
                                } else {
                                    v26 := GenS1{F0: 255, F1: p0}
                                    p0 = (2147483647 < 4294967295)
                                    v27 := []int{fi22}
                                    _ = v26
                                    _ = v27
                                }
                                if p0 {
                                    if p0 {
                                        v28 := []int{(fi24 % ((1) | 1)), fi22, 69}
                                        if (p0 || p0) {
                                            v29 := GenS5{F0: 2, F1: 83, F2: 1}
                                            fmt.Println(fi22, fi24)
                                            var v30 int64 = -997
                                            v30 -= v30
                                            _ = v29
                                            _ = v30
                                        }
                                        fmt.Println(fi24, p0)
                                        _ = v28
                                    } else {
                                        fmt.Println(fi24, fi22)
                                        switch "hello" {
                                        case "":
                                            v31 := []int{353, (fi24 & -9223372036854775808)}
                                            _ = v31
                                        case "hello":
                                            v32 := []int{-723}
                                            _ = v32
                                        case "foo bar":
                                            var v33 uint16 = 0
                                            _ = v33
                                        }
                                    }
                                    if ((-128 < 127) && true) {
                                        if (p0 && false) {
                                            v34 := []int{-569, (fi24 - (fi22 + (fi24 << ((14) & 63))))}
                                            p0 = p0
                                            _ = v34
                                        } else {
                                            switch fi24 {
                                            case 4294967295:
                                                p0 = p0
                                            case -342:
                                                p0 = p0
                                            }
                                            fmt.Println(p0)
                                        }
                                        v35 := GenS1{F0: 65535, F1: true}
                                        _ = v35
                                    } else {
                                        v36 := []int{}
                                        v37 := GenS6{F0: true, F1: ("foo bar" + "\n\t"), F2: 2147483647}
                                        _ = v36
                                        _ = v37
                                    }
                                } else {
                                    var v38 uint32 = 266
                                    _ = v38
                                }
                                _ = fi24
                            }
                        case "\n\t":
                            if (!p0) {
                                if p0 {
                                    var v39 uint64 = 689
                                    _ = v39
                                } else {
                                    p0 = p0
                                }
                                p0 = p0
                                p0 = true
                            }
                        case "a":
                            p0 = p0
                        }
                        _ = fi22
                        _ = v23
                    }
                case 447:
                    if (0.0 < (1.5 - 3.14159)) {
                        p0 = (537 > 0)
                    }
                case -2147483648:
                    switch -1 {
                    case -806:
                        var v40 int32 = 285
                        _ = v40
                    case -119:
                        fmt.Println(p0)
                    case 23:
                        fmt.Println(p0)
                    default:
                        var v41 bool = false
                        _ = v41
                    }
                }
            }
        }
    default:
        v42 := GenS5{F0: 42, F1: 2147483648, F2: 1}
        _ = v42
    }
    if ((!p0) || (-961 >= -432)) {
        if false {
            var v43 uint16 = 255
            _ = v43
        } else {
            switch "foo bar" {
            case "hello":
                switch 456 {
                case 32767:
                    for fi44 := 0; fi44 < 10; fi44++ {
                        if (257 == 338) {
                            v45 := GenS4{F0: -128, F1: -1}
                            v46 := []int{(fi44 % ((4294967295) | 1))}
                            fmt.Println(fi44)
                            _ = v45
                            _ = v46
                        }
                        var v47 uint8 = 223
                        v48 := []int{fi44, fi44, (fi44 * fi44)}
                        _ = fi44
                        _ = v47
                        _ = v48
                    }
                case 128:
                    p0 = (-532 < 263)
                case 4294967295:
                    var v49 bool = (0 < -1)
                    _ = v49
                }
            case "foo bar":
                fmt.Println(p0)
            default:
                v50 := GenS7{F0: -51, F1: -32768}
                _ = v50
            }
            p0 = true
        }
    }
    _ = p0
    return (127 > 1)
}
func genFn1(p0 uint8, p1 uint64, p2 int8) int {
    p2 = (-(p2))
    fmt.Println(p0, p1)
    v51 := GenS3{F0: true}
    p1 += (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    return 32768
}

func main() {
    var v1 float32 = float32((1.5 - 1e10))
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 float32 = v1
    v5 := []int{}
    v6 := make(map[int]uint32)
    if len(v5) > 0 {
        v5[0] = 128
    }
    v1 = v4
    fmt.Println(v4, v1)
    v4 = ((3.14159 - 3.14159) - v1)
    fmt.Println(v4, v4)
    v52 := GenS8{F0: 32767, F1: 1, F2: 18446744073709551615}
    var v53 int8 = -67
    delete(v3, int(v53))
    var v54 bool = false
    var v55 uint16 = 2
    if ((v53 & (v53 * (v53 % ((127) | 1)))) > 0) {
        v4 = v1
        v56 := GenS9{F0: v54}
        _ = v56
    } else {
        var v57 uint = 4294967295
        _ = v57
    }
    delete(v6, -970)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
