package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int8; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int; F1 int64 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(fi5)
            _ = fi5
        }
        _ = fi4
    }
    if (0 > 674) {
        var v6 int8 = -128
        v7 := []int{}
        var v8 string = "\n\t"
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: 65535, F1: 0, F2: 1}
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: -1, F1: -9223372036854775808}
        _ = fi9
        _ = v12
    }
    return 42
}
func genFn1(p0 bool, p1 int8) uint32 {
    v13 := GenS1{F0: 65535, F1: 811, F2: 4294967295}
    var v14 uint16 = 65535
    if false {
        v15 := []int{1, 9223372036854775807, 4294967295}
        var v16 int32 = -1
        _ = v15
        _ = v16
    } else {
        switch 127 {
        case -58:
            v17 := []int{-2147483648}
            _ = v17
        case 1:
            v18 := []int{}
            _ = v18
        default:
            switch "\n\t" {
            case "\n\t":
                var v19 uint8 = 2
                _ = v19
            case "a":
                var v20 string = "a"
                _ = v20
            default:
                var v21 uint32 = 42
                _ = v21
            }
        }
        fmt.Println(v14)
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint32 = 0
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return uint32(v14)
}

func main() {
    var v1 uint8 = 42
    var v2 int = 128
    var v3 uint32 = 2147483648
    v2 = 0
    if (2 <= 255) {
        v25 := []int{}
        var v26 bool = true
        _ = v25
        _ = v26
    }
    var v27 int64 = 600
    if (!true) {
        for fi28 := 0; fi28 < 2; fi28++ {
            switch "\n\t" {
            case "\n\t":
                if ((^(v1)) == v1) {
                    v29 := GenS1{F0: (v3 + (v3 | 2147483648)), F1: 65535, F2: (^(v3))}
                    _ = v29
                }
            case "hello":
                v27 -= (v27 % ((v27) | 1))
            case "":
                v30 := GenS2{F0: 259, F1: -1}
                _ = v30
            }
            _ = fi28
        }
        v27 = (v27 * v27)
    } else {
        switch (("\n\t" + "\n\t") + "foo bar") {
        case "":
            v31 := []int{(v2 - v2), (^(v2)), 274, v2}
            _ = v31
        case "foo bar":
            var v32 uint64 = 18446744073709551615
            _ = v32
        case "\n\t":
            var v33 uint8 = 1
            _ = v33
        default:
            if true {
                fmt.Println(v3)
                v2 = (v2 ^ (-(v2)))
            } else {
                v3 -= (v3 + 2)
            }
        }
        var v34 uint64 = 2
        switch ("hello" + ("" + "")) {
        case "":
            for fi35 := 0; fi35 < 3; fi35++ {
                switch "foo bar" {
                case "foo bar":
                    for fi36 := 0; fi36 < 6; fi36++ {
                        switch "" {
                        case "foo bar":
                            v37 := []int{int(v27), (^(fi35)), 1, (fi35 % (((-(fi35))) | 1))}
                            _ = v37
                        case "":
                            v38 := []int{(fi35 << ((910) & 63)), 128, (fi35 & -350), (fi36 / ((fi36) | 1))}
                            if (int16(fi36) == 32767) {
                                fmt.Println(v2, v34)
                                var v39 uint64 = 368
                                _ = v39
                            }
                            _ = v38
                        case "a":
                            var v40 float64 = 1e10
                            _ = v40
                        default:
                            v3 -= (^(v3))
                        }
                        if (0.1 == (0.0 + 0.1)) {
                            var v41 uint64 = v34
                            v3 += (^(v3))
                            v1 += (v1 >> ((704) & 63))
                            _ = v41
                        }
                        if false {
                            v3 = v3
                            if (23 != v27) {
                                var v42 uint64 = v34
                                v2 += (-(fi35))
                                _ = v42
                            } else {
                                var v43 uint64 = (v34 & v34)
                                _ = v43
                            }
                            switch "foo bar" {
                            case "a":
                                fmt.Println(fi35, v3, v34)
                            case "hello":
                                fmt.Println(v3, v3)
                            default:
                                var v44 float32 = -2.25
                                _ = v44
                            }
                        } else {
                            if false {
                                if (-1 < 2147483647) {
                                    v2 = v2
                                    var v45 bool = true
                                    v27 += 42
                                    v34 = (v34 % ((genFn0()) | 1))
                                    fmt.Println(v1, v2, v2)
                                    _ = v45
                                }
                                v2 = fi35
                                v27 = (v27 | 888)
                            } else {
                                if (42 != -291) {
                                    v3 = genFn1(false, 1)
                                    v46 := GenS3{F0: 1000, F1: -32768}
                                    v47 := GenS4{F0: 1, F1: 306, F2: v27}
                                    _ = v46
                                    _ = v47
                                }
                            }
                        }
                        _ = fi36
                    }
                case "a":
                    for fi48 := 0; fi48 < 15; fi48++ {
                        fmt.Println(v1, v3)
                        _ = fi48
                    }
                }
                var v49 uint8 = (^(v1))
                _ = fi35
                _ = v49
            }
        case "hello":
            var v50 uint64 = (v34 - v34)
            _ = v50
        default:
            v27 = (v27 % (((v27 - (v27 | int64(v34)))) | 1))
        }
        _ = v34
    }
    v51 := []int{int(v27), 32767}
    var v52 int8 = 100
    if (!(0.0 < (0.1 + float64(float32(((0.1 * 3.14159) / float64(1.5))))))) {
        var v53 uint64 = 899
        _ = v53
    }
    var v54 bool = false
    v1 += 2
    switch "hello" {
    case "hello":
        var v55 float32 = (-2.25 - (1.5 + 0.1))
        _ = v55
    case "foo bar":
        for fi56 := 0; fi56 < 11; fi56++ {
            v2 -= (fi56 + 38)
            _ = fi56
        }
    case "":
        v57 := GenS5{F0: 0, F1: 0, F2: genFn1(((630 > 1) && v54), (v52 ^ v52))}
        _ = v57
    }
    v27 = (-(v27))
    switch "foo bar" {
    case "a":
        v54 = v54
    default:
        v58 := GenS6{F0: v2, F1: (v27 << ((0) & 63))}
        _ = v58
    }
    v59 := GenS6{F0: 0, F1: (v27 + int64(v2))}
    v60 := GenS3{F0: 719, F1: -884}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v51
    _ = v52
    _ = v54
    _ = v59
    _ = v60
}
