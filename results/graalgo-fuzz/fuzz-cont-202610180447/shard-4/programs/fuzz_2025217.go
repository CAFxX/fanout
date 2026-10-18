package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32; F1 uint64; F2 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint32; F1 int16; F2 int }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int; F1 int8 }
func genFn0(p0 int32) uint32 {
    fmt.Println(p0)
    if (((-2.25 / 1.0) + (1.5 * float64(3.14159))) < (0.0 - 0.1)) {
        v6 := GenS1{F0: 379, F1: "\n\t"}
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int8 = 127
            var v9 int64 = 9223372036854775807
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        switch "foo bar" {
        case "a":
            for fi10 := 0; fi10 < 20; fi10++ {
                if false {
                    var v11 int16 = -755
                    _ = v11
                }
                v12 := GenS2{F0: (p0 % ((2147483647) | 1)), F1: 18446744073709551615, F2: "a"}
                _ = fi10
                _ = v12
            }
        case "\n\t":
            p0 = -38
        default:
            if true {
                v13 := GenS3{F0: 188, F1: -2147483648}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS4{F0: 735, F1: 0}
                    if (!(!false)) {
                        if true {
                            switch "hello" {
                            case "hello":
                                switch ("" + ("foo bar" + "a")) {
                                case "a":
                                    fmt.Println(p0, fi14)
                                default:
                                    var v16 uint64 = 372
                                    _ = v16
                                }
                            case "foo bar":
                                v17 := GenS5{F0: 1, F1: (42 <= 86)}
                                _ = v17
                            default:
                                var v18 uint64 = 941
                                _ = v18
                            }
                            switch ("a" + ("foo bar" + "\n\t")) {
                            case "foo bar":
                                var v19 float64 = 3.14159
                                _ = v19
                            default:
                                var v20 int64 = int64(p0)
                                _ = v20
                            }
                            v21 := GenS6{F0: 503, F1: 0, F2: (-(fi14))}
                            _ = v21
                        }
                        var v22 bool = (!false)
                        v23 := []int{(fi14 | (fi14 >> ((255) & 63))), (fi14 & -909), -2147483648, -137}
                        _ = v22
                        _ = v23
                    } else {
                        if (false || (2147483647 >= 2147483648)) {
                            p0 += 42
                        }
                        var v24 int8 = 127
                        _ = v24
                    }
                    v25 := GenS7{F0: 439, F1: ""}
                    _ = fi14
                    _ = v15
                    _ = v25
                }
                fmt.Println(p0)
            }
        }
    }
    var v26 int = 592
    _ = p0
    _ = v26
    return 42
}
func genFn1(p0 uint16, p1 uint, p2 int32) int {
    var v27 int8 = 0
    var v28 float64 = 1e10
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := []int{-9223372036854775808, (fi29 & (fi29 | fi29)), fi29}
        switch "a" {
        case "\n\t":
            switch -2147483648 {
            case 128:
                v31 := GenS8{F0: 4294967295}
                _ = v31
            }
        case "foo bar":
            var v32 uint8 = 1
            _ = v32
        default:
            var v33 float64 = float64(-2.25)
            _ = v33
        }
        fmt.Println(p2, p1, p0)
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return 0
}

func main() {
    v1 := []int{1, 32768, 2147483647}
    var v2 uint8 = 203
    var v3 string = "\n\t"
    v4 := []string{("a" + (v3 + "\n\t")), "\n\t"}
    var v5 int32 = 0
    for fi34 := 0; fi34 < 3; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint64 = 18446744073709551615
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v3 = v3
    switch (v3 + v3) {
    case "\n\t":
        var v37 int8 = 1
        _ = v37
    case "hello":
        v38 := GenS9{F0: 42, F1: (v2 / ((v2) | 1)), F2: 65535}
        _ = v38
    }
    fmt.Println(v3, v5)
    if (false && false) {
        fmt.Println(v5, v5)
    }
    if (!true) {
        var v39 int = 95
        _ = v39
    } else {
        for fi40 := 0; fi40 < 11; fi40++ {
            v41 := GenS10{F0: genFn1(540, 241, 2147483647), F1: 0}
            var v42 int16 = 1
            _ = fi40
            _ = v41
            _ = v42
        }
        if true {
            for fi43 := 0; fi43 < 17; fi43++ {
                _ = len(v4)
                switch v3 {
                case "hello":
                    switch (v3 + "foo bar") {
                    case "foo bar":
                        var v44 int = (fi43 % ((536) | 1))
                        _ = v44
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    for fi45 := 0; fi45 < 1; fi45++ {
                        v5 = (v5 | v5)
                        fmt.Println(v2, fi45)
                        _ = fi45
                    }
                }
                v46 := GenS7{F0: 748, F1: (v3 + v3)}
                _ = fi43
                _ = v46
            }
            for fi47 := 0; fi47 < 17; fi47++ {
                var v48 int8 = -1
                var v49 float32 = 1e10
                _ = fi47
                _ = v48
                _ = v49
            }
            switch v3 {
            case "a":
                if len(v4) > 0 {
                    v4[0] = (v3 + (v3 + ((v3 + "hello") + v3)))
                }
            case "hello":
                v3 = v3
            case "":
                for fi50 := 0; fi50 < 4; fi50++ {
                    v3 = v3
                    _ = len(v4)
                    if (!false) {
                        var v51 int32 = (v5 & v5)
                        _ = v51
                    }
                    _ = fi50
                }
            default:
                v2 -= v2
            }
        }
    }
    for fi52 := 0; fi52 < 5; fi52++ {
        fmt.Println(v3, v3, fi52)
        fmt.Println(v5, fi52)
        _ = fi52
    }
    var v53 int8 = 29
    for fi54 := 0; fi54 < 15; fi54++ {
        var v55 int8 = (v53 - v53)
        v5 -= (v5 & (v5 + (v5 << ((499) & 63))))
        _ = fi54
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
