package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 bool }
func genFn0(p0 uint16, p1 int) bool {
    v7 := GenS1{F0: -968}
    p0 += p0
    p1 = p1
    var v8 int = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 uint64, p1 uint64) uint8 {
    v9 := []int{0}
    fmt.Println(p0, p1)
    fmt.Println(p0, p0)
    fmt.Println(p1)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    var v1 string = ""
    v2 := []uint32{2147483648, 2147483647, 392, 2147483648}
    var v3 float64 = (1e10 - float64(0.1))
    var v4 int32 = 1
    v5 := []uint32{2147483647}
    var v6 int16 = 32767
    if len(v5) > 0 {
        v5[0] = 42
    }
    fmt.Println(v3, v1)
    fmt.Println(v6, v3)
    switch "foo bar" {
    case "hello":
        if true {
            var v10 int8 = -128
            for fi11 := 0; fi11 < 13; fi11++ {
                v2 = append(v2, 4294967295)
                if (v1 <= v1) {
                    v10 -= v10
                    var v12 int8 = v10
                    v13 := GenS2{F0: 0, F1: 415, F2: 65535}
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
            _ = v10
        }
    case "a":
        for fi14 := 0; fi14 < 9; fi14++ {
            if true {
                v15 := GenS3{F0: 2, F1: 42, F2: 596}
                _ = v15
            } else {
                var v16 uint16 = 1
                _ = v16
            }
            switch 1 {
            case 128:
                if (-505 > -1) {
                    if (!false) {
                        switch fi14 {
                        case 1:
                            for fi17 := 0; fi17 < 12; fi17++ {
                                var v18 int32 = 2147483647
                                var v19 uint16 = 42
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        default:
                            v20 := GenS3{F0: 72, F1: 42, F2: 255}
                            _ = v20
                        }
                        if (0 <= 1) {
                            v21 := GenS4{F0: true, F1: v4}
                            if len(v2) > 0 {
                                v2[0] = 226
                            }
                            switch 9223372036854775807 {
                            case 1:
                                if (!(255 >= 0)) {
                                    var v22 string = "foo bar"
                                    var v23 string = ("hello" + v22)
                                    _ = v22
                                    _ = v23
                                }
                            case -9223372036854775808:
                                if true {
                                    v3 = v3
                                    var v24 float64 = float64((0.1 / (-2.25 * 0.1)))
                                    var v25 string = v1
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                switch ((((((v1 + ("hello" + ((v1 + v1) + (v1 + v1)))) + v1) + (v1 + "\n\t")) + v1) + ((("\n\t" + ("\n\t" + "hello")) + v1) + "hello")) + ("hello" + (((v1 + ((v1 + ("" + v1)) + ((v1 + (v1 + (v1 + v1))) + "a"))) + v1) + "foo bar"))) {
                                case "\n\t":
                                    if ((true || (v3 < (1e10 * (v3 / v3)))) || true) {
                                        var v26 bool = ((1.5 / ((0.1 + -2.25) / 1.5)) == 0.0)
                                        v6 = 32767
                                        _ = v26
                                    } else {
                                        var v27 bool = true
                                        fmt.Println(fi14, v3)
                                        _ = v27
                                    }
                                }
                            }
                            _ = v21
                        } else {
                            for fi28 := 0; fi28 < 12; fi28++ {
                                v6 = v6
                                _ = fi28
                            }
                            var v29 float64 = v3
                            _ = v29
                        }
                    } else {
                        v30 := GenS1{F0: (v6 * 32767)}
                        _ = v30
                    }
                    v31 := GenS5{F0: 1, F1: 109, F2: 42}
                    _ = v31
                }
            default:
                v32 := GenS1{F0: (-(v6))}
                _ = v32
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 float64 = (0.0 / float64(1.5))
                _ = fi33
                _ = v34
            }
            _ = fi14
        }
    default:
        for fi35 := 0; fi35 < 14; fi35++ {
            if (65535 <= 4294967295) {
                var v36 uint16 = 388
                v5 = append(v5, 2)
                _ = v36
            } else {
                if false {
                    v37 := GenS1{F0: -1}
                    _ = len(v5)
                    if (!(true || false)) {
                        v5 = append(v5, 0)
                    }
                    _ = v37
                } else {
                    v1 = ("a" + v1)
                    _ = len(v5)
                    for fi38 := 0; fi38 < 14; fi38++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v39 int = 127
                        _ = fi38
                        _ = v39
                    }
                }
                switch fi35 {
                case 2147483647:
                    v5 = append(v5, 238)
                case 32767:
                    v3 = v3
                case 128:
                    for fi40 := 0; fi40 < 15; fi40++ {
                        v41 := GenS3{F0: genFn1(231, 1), F1: 0, F2: 110}
                        switch ("hello" + v1) {
                        case "foo bar":
                            v1 = ""
                            fmt.Println(fi35, v1, v3)
                        }
                        _ = fi40
                        _ = v41
                    }
                }
            }
            _ = fi35
        }
    }
    fmt.Println(v6, v3, v3)
    v1 = v1
    v42 := GenS6{F0: (0 >= 1)}
    v1 = v1
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 uint32 = 2147483647
        v44 = 0
        v45 := GenS7{F0: (v4 | (v4 + 2147483647))}
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println(v6, v6)
    var v46 uint64 = 18446744073709551615
    fmt.Println(v46, v46, v46)
    v1 = v1
    _ = len(v2)
    var v47 int8 = 127
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 uint32 = 2
        var v50 bool = (false && ((true && true) || false))
        var v51 int16 = v6
        _ = fi48
        _ = v49
        _ = v50
        _ = v51
    }
    if false {
        if (true || (0 == 196)) {
            var v52 float32 = (1.5 * 0.0)
            switch (("hello" + v1) + v1) {
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
            }
            v2 = append(v2, 65535)
            _ = v52
        } else {
            var v53 int8 = 83
            _ = v53
        }
        for fi54 := 0; fi54 < 6; fi54++ {
            v46 -= (v46 * 4)
            var v55 float32 = -2.25
            switch (fi54 / ((fi54) | 1)) {
            case -2147483648:
                var v56 uint16 = 223
                _ = v56
            }
            _ = fi54
            _ = v55
        }
    } else {
        var v57 uint8 = 86
        switch "foo bar" {
        case "\n\t":
            switch v1 {
            case "\n\t":
                v6 = v6
            default:
                v58 := GenS5{F0: 255, F1: v47, F2: 65535}
                _ = v58
            }
        default:
            var v59 string = v1
            _ = v59
        }
        var v60 int64 = 955
        _ = v57
        _ = v60
    }
    if (2147483648 != 525) {
        if true {
            if (true && true) {
                v61 := GenS8{F0: v6, F1: false}
                v5 = append(v5, 521)
                _ = v61
            } else {
                fmt.Println(v3)
                for fi62 := 0; fi62 < 7; fi62++ {
                    var v63 uint16 = 255
                    if false {
                        v47 = -51
                    }
                    _ = fi62
                    _ = v63
                }
            }
            fmt.Println(v3)
        }
        var v64 int16 = 32767
        var v65 bool = false
        _ = v64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
