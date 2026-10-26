package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 uint }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 bool) uint64 {
    p0 = (p0 % (((p0 << ((18446744073709551615) & 63))) | 1))
    for fi4 := 0; fi4 < 1; fi4++ {
        p1 = (3.14159 >= 0.0)
        v5 := GenS1{F0: -128, F1: 127}
        _ = fi4
        _ = v5
    }
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: -1, F1: -124}
        v8 := []int{42}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 590
}
func genFn1(p0 uint64) uint {
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 += (^(p0))
        switch (((("foo bar" + "foo bar") + "foo bar") + "\n\t") + ("a" + "foo bar")) {
        case "foo bar":
            v10 := []int{(fi9 << ((1) & 63)), 4294967295}
            _ = v10
        default:
            var v11 int32 = 1
            _ = v11
        }
        _ = fi9
    }
    v12 := []int{0}
    p0 = (p0 << ((65535) & 63))
    _ = p0
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 int32 = 1
    var v13 uint32 = 2147483647
    v3 += (-(v3))
    var v14 int32 = (v3 * (^(v3)))
    v13 += 0
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{}
        v13 = v13
        if true {
            var v17 int8 = 92
            fmt.Println(v17, fi15, v3)
            v17 = -1
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    var v18 int32 = v14
    var v19 uint16 = 42
    for fi20 := 0; fi20 < 18; fi20++ {
        v14 += v3
        _ = fi20
    }
    if true {
        var v21 int8 = 42
        if false {
            v22 := GenS1{F0: v21, F1: (v21 / (((v21 % ((v21) | 1))) | 1))}
            fmt.Println(v13, v18)
            _ = v22
        } else {
            _ = len(v2)
            v1[206] = 976
            if (!(false || ((v19 >> ((0) & 63)) <= 0))) {
                v14 = v3
            }
        }
        _ = v21
    }
    var v23 bool = true
    if v23 {
        fmt.Println(v18)
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS1{F0: 32, F1: 1}
            for fi26 := 0; fi26 < 19; fi26++ {
                _ = len(v2)
                if (v23 && (-2.25 == 0.1)) {
                    var v27 float64 = 3.14159
                    switch (fi26 | fi26) {
                    case -886:
                        if ("\n\t" != ("" + "hello")) {
                            _ = len(v1)
                            v23 = false
                        }
                    }
                    var v28 uint32 = v13
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v23, v18)
                }
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
    } else {
        switch "\n\t" {
        case "":
            v23 = v23
        default:
            switch "hello" {
            case "foo bar":
                if (!v23) {
                    if (v19 != v19) {
                        switch "" {
                        case "":
                            v29 := GenS2{F0: -128, F1: 65535, F2: 42}
                            _ = v29
                        case "a":
                            if (95 > 128) {
                                var v30 uint = genFn1(2)
                                _ = v30
                            }
                        }
                    } else {
                        if v23 {
                            var v31 int64 = 0
                            _ = v31
                        }
                        if false {
                            switch -2147483648 {
                            case 32767:
                                v32 := []int{42, 413, 350, -2147483648}
                                _ = v32
                            case -2147483648:
                                var v33 string = "a"
                                _ = v33
                            case -1:
                                var v34 float64 = float64(float32(0.1))
                                _ = v34
                            }
                        }
                        switch "foo bar" {
                        case "":
                            fmt.Println(v23, v18, v14)
                        case "foo bar":
                            for fi35 := 0; fi35 < 10; fi35++ {
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    var v37 uint8 = 2
                                    var v38 int = (fi35 + 128)
                                    switch 369 {
                                    case -1:
                                        switch "a" {
                                        case "hello":
                                            v23 = v23
                                        case "foo bar":
                                            fmt.Println(v14)
                                        case "a":
                                            var v39 int8 = 115
                                            _ = v39
                                        default:
                                            v13 = v13
                                        }
                                    case 128:
                                        v40 := v1[65535]
                                        _ = v40
                                    case 2147483647:
                                        v19 = (v19 / ((v19) | 1))
                                    }
                                    var v41 string = "a"
                                    v1[0] = (fi36 & v38)
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                    _ = v41
                                }
                                v42 := GenS1{F0: 42, F1: -16}
                                delete(v1, 2147483647)
                                _ = fi35
                                _ = v42
                            }
                        case "\n\t":
                            v18 = (-(v18))
                        }
                    }
                    v43 := []int{-888, 32767, -912}
                    _ = v43
                } else {
                    if (112 >= 42) {
                        if false {
                            for fi44 := 0; fi44 < 6; fi44++ {
                                if false {
                                    v18 -= 0
                                    var v45 uint = uint(v13)
                                    if (!true) {
                                        for fi46 := 0; fi46 < 8; fi46++ {
                                            var v47 int32 = v14
                                            var v48 int16 = int16(v47)
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                        }
                                        v49 := GenS3{F0: -128, F1: 1000}
                                        v50 := GenS4{F0: 9223372036854775807, F1: (v45 * 23)}
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v51 := GenS5{F0: 127}
                                        for fi52 := 0; fi52 < 20; fi52++ {
                                            v53 := GenS3{F0: -74, F1: -32768}
                                            fmt.Println(v14, fi44, v14)
                                            v54 := []int{32768}
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = v51
                                    }
                                    _ = v45
                                }
                                _ = fi44
                            }
                            if (0 == 926) {
                                _ = len(v1)
                                var v55 uint32 = v13
                                var v56 uint64 = 4294967295
                                _ = v55
                                _ = v56
                            }
                            for fi57 := 0; fi57 < 2; fi57++ {
                                v13 = 4294967295
                                for fi58 := 0; fi58 < 1; fi58++ {
                                    fmt.Println(v18, fi57)
                                    v59 := v2[v13]
                                    _ = fi58
                                    _ = v59
                                }
                                _ = fi57
                            }
                        }
                    } else {
                        v60 := GenS1{F0: 127, F1: 127}
                        v13 -= 65535
                        fmt.Println(v18, v19)
                        _ = v60
                    }
                    fmt.Println(v19, v18)
                }
            }
        }
        v13 -= v13
        v19 += v19
    }
    switch "\n\t" {
    case "hello":
        var v61 uint16 = (v19 >> ((951) & 63))
        _ = v61
    }
    v62 := GenS6{F0: ("" + "a"), F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v62
}
