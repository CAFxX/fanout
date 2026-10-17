package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 int64; F2 int64 }
type GenS9 struct { F0 bool }
func genFn0() int8 {
    var v5 uint16 = 255
    switch "foo bar" {
    case "\n\t":
        var v6 uint8 = 2
        _ = v6
    }
    if (-1 >= 1) {
        v5 = v5
    }
    if false {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch "hello" {
            case "":
                v8 := []int{}
                _ = v8
            case "foo bar":
                v9 := []int{}
                _ = v9
            default:
                switch fi7 {
                case 128:
                    fmt.Println(fi7)
                case -701:
                    v5 += 0
                default:
                    v10 := GenS1{F0: 0, F1: int8(fi7), F2: 2}
                    _ = v10
                }
            }
            _ = fi7
        }
        var v11 int32 = -459
        fmt.Println(v5, v5)
        _ = v11
    }
    _ = v5
    return -128
}
func genFn1(p0 int32, p1 uint64) int {
    fmt.Println(p0)
    var v12 int32 = 0
    switch 42 {
    case 235:
        if true {
            p1 = p1
            p1 = 42
            if true {
                v13 := []int{-703, -9223372036854775808, 32767}
                _ = v13
            } else {
                v14 := []int{4294967295, -964, 32767}
                _ = v14
            }
        } else {
            p0 = v12
            fmt.Println(p1, p1)
            var v15 int16 = 597
            _ = v15
        }
    case -820:
        var v16 float64 = ((-2.25 - -2.25) - (1.5 / -2.25))
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS1{F0: p0, F1: 127, F2: 255}
        var v19 bool = (true || true)
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := []int{-887, 880, 9223372036854775807, -218}
    _ = p0
    _ = p1
    _ = v12
    _ = v20
    return 597
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    v3 := []string{}
    v4 := make(map[int]string)
    var v21 uint32 = 36
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS2{F0: "a"}
        if false {
            for fi24 := 0; fi24 < 16; fi24++ {
                v21 += (v21 >> ((924) & 63))
                if true {
                    var v25 uint8 = 206
                    var v26 uint16 = 42
                    v25 = 255
                    _ = v25
                    _ = v26
                }
                v21 -= v21
                fmt.Println(fi22, fi24)
                _ = fi24
            }
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 string = "hello"
                _ = fi27
                _ = v28
            }
        }
        _ = fi22
        _ = v23
    }
    var v29 uint8 = 126
    if true {
        v21 = (v21 | 2147483648)
    } else {
        if false {
            var v30 uint8 = 1
            for fi31 := 0; fi31 < 3; fi31++ {
                fmt.Println(v29, fi31, v29)
                switch fi31 {
                case 1:
                    v29 = (v30 - v30)
                case -1:
                    v32 := GenS1{F0: 42, F1: 127, F2: 295}
                    _ = v32
                }
                if false {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi31)
                    default:
                        var v33 bool = (!false)
                        _ = v33
                    }
                    v34 := GenS1{F0: 692, F1: -1, F2: 643}
                    v30 += v30
                    _ = v34
                }
                _ = fi31
            }
            _ = v30
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi35 := 0; fi35 < 16; fi35++ {
            fmt.Println(fi35)
            for fi36 := 0; fi36 < 20; fi36++ {
                var v37 uint16 = 65535
                fmt.Println(fi36, fi35, fi36)
                _ = fi36
                _ = v37
            }
            if true {
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 uint8 = v29
                    if (!(false || true)) {
                        v40 := GenS1{F0: -1, F1: -1, F2: 383}
                        if (v21 <= (v21 % (((v21 * (v21 << ((255) & 63)))) | 1))) {
                            switch ("hello" + ((((("\n\t" + ("" + "\n\t")) + "hello") + ("foo bar" + "a")) + (("a" + "foo bar") + "hello")) + ("\n\t" + ("hello" + ("foo bar" + ""))))) {
                            case "foo bar":
                                switch ("foo bar" + ("" + ("\n\t" + "a"))) {
                                case "hello":
                                    v41 := GenS3{F0: -9223372036854775808, F1: 487, F2: 536}
                                    var v42 int = (fi38 << ((2) & 63))
                                    switch fi38 {
                                    case 0:
                                        var v43 int8 = 42
                                        _ = v43
                                    default:
                                        v44 := GenS2{F0: "hello"}
                                        _ = v44
                                    }
                                    _ = v41
                                    _ = v42
                                case "foo bar":
                                    v45 := GenS4{F0: 42, F1: 802, F2: 1}
                                    _ = v45
                                case "a":
                                    var v46 int32 = -1
                                    _ = v46
                                }
                            }
                            v39 = (v29 % (((v29 % ((v39) | 1))) | 1))
                            v47 := GenS1{F0: 0, F1: 1, F2: 18446744073709551615}
                            _ = v47
                        } else {
                            switch (fi35 << ((18446744073709551615) & 63)) {
                            case -2147483648:
                                v48 := GenS5{F0: int8(v29)}
                                _ = v48
                            case 1:
                                var v49 string = "\n\t"
                                _ = v49
                            case -1:
                                var v50 bool = ((v21 * (v21 % ((v21) | 1))) < 362)
                                _ = v50
                            }
                            var v51 uint32 = 2147483648
                            var v52 int = fi38
                            _ = v51
                            _ = v52
                        }
                        v53 := GenS6{F0: 65535, F1: 100}
                        _ = v40
                        _ = v53
                    }
                    v21 = v21
                    _ = fi38
                    _ = v39
                }
            }
            _ = fi35
        }
    default:
        var v54 uint64 = 2
        _ = v54
    }
    v55 := GenS3{F0: 0, F1: 462, F2: 255}
    v56 := GenS7{F0: 255}
    v57 := GenS8{F0: 401, F1: -485, F2: -46}
    for fi58 := 0; fi58 < 19; fi58++ {
        v59 := GenS9{F0: ((1e10 / 3.14159) != 0.0)}
        v60 := GenS8{F0: 0, F1: 9223372036854775807, F2: 210}
        var v61 string = "\n\t"
        _ = fi58
        _ = v59
        _ = v60
        _ = v61
    }
    v29 = v29
    switch "foo bar" {
    case "a":
        var v62 int8 = 100
        _ = v62
    case "foo bar":
        v21 = v21
    case "hello":
        var v63 uint = 2
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
}
