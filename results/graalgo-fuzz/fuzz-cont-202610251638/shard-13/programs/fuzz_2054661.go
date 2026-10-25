package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8, p1 bool) uint8 {
    if p1 {
        p0 = (^(p0))
        v5 := GenS1{F0: -2147483648, F1: 0}
        v6 := []int{0, -971, -9223372036854775808}
        _ = v5
        _ = v6
    }
    v7 := []int{-9223372036854775808, 4294967295}
    var v8 int16 = -32768
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 uint8, p1 int) int64 {
    p1 = (p1 / (((p1 % ((2147483647) | 1))) | 1))
    p1 -= p1
    v9 := []int{p1, 104, (p1 % ((p1) | 1)), p1}
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2}
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    if (!(!(-1 >= 128))) {
        switch 32768 {
        case 2147483647:
            if true {
                v10 := GenS2{F0: genFn1(0, 960), F1: 9223372036854775807}
                switch ("hello" + "") {
                case "foo bar":
                    var v11 uint8 = 46
                    _ = v11
                case "\n\t":
                    var v12 int32 = 2147483647
                    _ = v12
                }
                _ = v10
            }
        case 32767:
            var v13 int = 97
            _ = v13
        default:
            var v14 float64 = (3.14159 - float64((float32(3.14159) * 1e10)))
            _ = v14
        }
    }
    var v15 string = (("" + "hello") + "hello")
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(fi16, v15)
        fmt.Println(v15, fi16)
        var v17 int8 = 100
        _ = fi16
        _ = v17
    }
    var v18 bool = false
    switch "foo bar" {
    case "hello":
        var v19 uint16 = 835
        _ = v19
    case "a":
        var v20 uint = 562
        _ = v20
    }
    switch ("" + v15) {
    case "foo bar":
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 bool = v18
            for fi23 := 0; fi23 < 9; fi23++ {
                if (2147483647 > 42) {
                    fmt.Println(fi21, v15)
                    if ((48 > 594) && ((-32768 >= -1) && v22)) {
                        if v18 {
                            var v24 uint32 = 810
                            _ = v24
                        } else {
                            var v25 float32 = ((3.14159 + -2.25) + -2.25)
                            _ = len(v2)
                            v26 := GenS3{F0: 919}
                            _ = v25
                            _ = v26
                        }
                        if (genFn1(42, -1) == genFn1(42, (^(fi21)))) {
                            switch fi23 {
                            case 42:
                                if ((fi23 + (-(fi23))) != fi23) {
                                    v18 = true
                                    v18 = true
                                }
                            case -1:
                                if false {
                                    fmt.Println(v18, fi21, v15)
                                    switch ("" + v15) {
                                    case "hello":
                                        v18 = (v15 <= v15)
                                    default:
                                        v15 = v15
                                    }
                                    switch v15 {
                                    case "\n\t":
                                        delete(v4, fi23)
                                    case "a":
                                        v18 = (0 <= 32767)
                                    }
                                    var v27 uint64 = 1
                                    var v28 float32 = 0.0
                                    _ = v27
                                    _ = v28
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            v22 = (!v22)
                        }
                    } else {
                        var v29 uint16 = 2
                        if v18 {
                            v1[(("\n\t" + v15) + ((v15 + "hello") + ""))] = (fi21 & (fi23 * fi21))
                            v22 = v18
                            v30 := GenS3{F0: 465}
                            v15 = v15
                            _ = v30
                        }
                        _ = v29
                    }
                }
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    case "\n\t":
        for fi31 := 0; fi31 < 8; fi31++ {
            for fi32 := 0; fi32 < 2; fi32++ {
                var v33 int = 32768
                v34 := GenS1{F0: fi31, F1: -9223372036854775808}
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    }
    var v35 float64 = 0.0
    switch v15 {
    case "foo bar":
        v36 := GenS2{F0: 42, F1: -148}
        _ = v36
    case "\n\t":
        v37 := GenS1{F0: -9223372036854775808, F1: 959}
        _ = v37
    case "a":
        v18 = (-621 > 1)
    default:
        switch -2147483648 {
        case -2147483648:
            if (2 <= 209) {
                var v38 string = ("\n\t" + (("" + v15) + v15))
                _ = v38
            }
        default:
            v15 = v15
        }
    }
    var v39 int64 = 42
    v40 := GenS4{F0: 100, F1: 9223372036854775807}
    v39 -= (v39 + -9223372036854775808)
    if true {
        _ = len(v2)
        v41 := GenS4{F0: 127, F1: -332}
        v15 = "\n\t"
        _ = v41
    } else {
        v42 := GenS5{F0: 694}
        fmt.Println(v35)
        _ = v42
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v2 = append(v2, 65535)
    v3[(v15 + v15)] = 824
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
}
