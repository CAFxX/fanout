package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint32; F1 bool; F2 int32 }
type GenS7 struct { F0 uint; F1 bool; F2 bool }
type GenS8 struct { F0 int16; F1 int8; F2 uint16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16 }

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 42
    v3 := make(map[string]int)
    var v4 int64 = -9223372036854775808
    var v5 uint = 2
    if (true || false) {
        for fi6 := 0; fi6 < 11; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                v2 = v2
                _ = fi7
            }
            v5 = v5
            fmt.Println(v2)
            _ = fi6
        }
    }
    var v8 float32 = 0.0
    switch (("hello" + "") + "\n\t") {
    case "\n\t":
        if true {
            for fi9 := 0; fi9 < 12; fi9++ {
                v2 = 18446744073709551615
                v4 = v4
                _ = fi9
            }
        }
    case "a":
        v8 += 1e10
    default:
        if (v1 != (^(v1))) {
            v10 := []int{2147483647}
            if (127 != (v1 ^ (^(v1)))) {
                _ = len(v3)
            } else {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS1{F0: 0, F1: 1, F2: 0}
                    fmt.Println(v8, v4)
                    _ = fi11
                    _ = v12
                }
            }
            _ = v10
        } else {
            v13 := GenS2{F0: int64(v5), F1: 0, F2: v2}
            v14 := GenS1{F0: 100, F1: (v1 ^ 127), F2: 127}
            v15 := []int{0, 9223372036854775807, 332}
            _ = v13
            _ = v14
            _ = v15
        }
    }
    v16 := GenS1{F0: 23, F1: v1, F2: 42}
    if ((!(((("a" + ((("a" + "a") + "a") + (("\n\t" + (("foo bar" + "\n\t") + "\n\t")) + (("" + ("foo bar" + "\n\t")) + "hello")))) + ("\n\t" + "\n\t")) == ("a" + "a")) || (!false))) && (313 != 0)) {
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS1{F0: 42, F1: (v1 << ((v5) & 63)), F2: -1}
            _ = len(v3)
            _ = len(v3)
            _ = fi17
            _ = v18
        }
        switch -127 {
        case 128:
            var v19 uint16 = 65535
            _ = v19
        case 9223372036854775807:
            for fi20 := 0; fi20 < 8; fi20++ {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 int64 = (-(v4))
                    v23 := []int{fi21}
                    v24 := GenS1{F0: 0, F1: uint8(v2), F2: -123}
                    var v25 uint = (v5 * v5)
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                var v26 int64 = v4
                _ = fi20
                _ = v26
            }
        case 4294967295:
            v27 := GenS1{F0: 42, F1: 127, F2: 49}
            _ = v27
        }
    } else {
        var v28 uint32 = 512
        _ = v28
    }
    if false {
        var v29 int16 = -619
        _ = v29
    } else {
        v30 := GenS1{F0: 127, F1: (^(v1)), F2: -57}
        if false {
            fmt.Println(v5)
            fmt.Println(v4)
            if ((0 < 2) || (-932 <= -1)) {
                var v31 uint = v5
                if true {
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v33 := GenS3{F0: int64(v2), F1: (("hello" + ("hello" + ("" + "foo bar"))) + "\n\t"), F2: 100}
                        v3["a"] = (fi32 ^ (-(fi32)))
                        v31 += (v31 / ((2) | 1))
                        _ = fi32
                        _ = v33
                    }
                    v1 = (v1 ^ v1)
                } else {
                    if true {
                        var v34 uint = 213
                        for fi35 := 0; fi35 < 7; fi35++ {
                            v5 += uint(v2)
                            v36 := GenS4{F0: -1}
                            _ = fi35
                            _ = v36
                        }
                        if true {
                            fmt.Println(v34, v4)
                            for fi37 := 0; fi37 < 5; fi37++ {
                                v38 := []int{-228, 127, (fi37 << ((v34) & 63)), fi37}
                                v39 := GenS5{F0: true, F1: 42}
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        } else {
                            v40 := GenS6{F0: 877, F1: (!true), F2: -817}
                            v41 := GenS7{F0: (v5 & 267), F1: ((v31 ^ v34) <= (v31 / ((v5) | 1))), F2: true}
                            if (false && (float64(v8) == 0.0)) {
                                v31 = 2
                                if ((v4 << ((v31) & 63)) == (v4 / ((0) | 1))) {
                                    var v42 uint8 = v1
                                    var v43 string = ("foo bar" + "hello")
                                    v44 := []int{128, int(v4), int(v31), 9223372036854775807}
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                } else {
                                    v45 := GenS8{F0: 32767, F1: -28, F2: 1}
                                    var v46 float32 = 1.5
                                    _ = v45
                                    _ = v46
                                }
                                v31 = v5
                            }
                            _ = v40
                            _ = v41
                        }
                        _ = v34
                    }
                }
                fmt.Println(v4)
                _ = v31
            }
        }
        _ = v30
    }
    switch ("hello" + "") {
    case "":
        if (false || (1e10 < (v8 / (v8 / (1.5 / 1e10))))) {
            switch "hello" {
            case "a":
                for fi47 := 0; fi47 < 7; fi47++ {
                    v48 := []int{(fi47 ^ 0)}
                    _ = fi47
                    _ = v48
                }
            }
        } else {
            switch 42 {
            case 9223372036854775807:
                var v49 uint32 = 42
                _ = v49
            case 296:
                _ = len(v3)
            default:
                for fi50 := 0; fi50 < 13; fi50++ {
                    var v51 int16 = 1
                    _ = fi50
                    _ = v51
                }
            }
            if (false && (!true)) {
                switch 554 {
                case 1:
                    switch 968 {
                    case -9223372036854775808:
                        var v52 bool = true
                        _ = v52
                    }
                case -1:
                    for fi53 := 0; fi53 < 19; fi53++ {
                        var v54 uint32 = 65535
                        var v55 uint16 = 852
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                case 2147483647:
                    fmt.Println(v2, v5)
                default:
                    switch -1 {
                    case 750:
                        v5 -= (v5 >> (((v5 ^ v5)) & 63))
                    case 867:
                        var v56 float64 = 0.0
                        _ = v56
                    case 1:
                        fmt.Println(v2, v8, v8)
                    default:
                        v3["a"] = -588
                    }
                }
                var v57 int64 = (v4 << (((v5 & (^(v5)))) & 63))
                v8 = ((v8 * (1e10 + 0.0)) - v8)
                _ = v57
            }
            switch 32768 {
            case 1:
                for fi58 := 0; fi58 < 12; fi58++ {
                    var v59 int8 = 127
                    var v60 int = 689
                    v61 := []int{v60, fi58}
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                }
            }
        }
    case "hello":
        v62 := v3[("" + ("" + "hello"))]
        _ = v62
    case "a":
        fmt.Println(v8, v8)
    default:
        if (0.0 < 0.1) {
            var v63 uint32 = 2147483647
            _ = v63
        }
    }
    v64 := GenS9{F0: -128}
    fmt.Println(v1)
    for fi65 := 0; fi65 < 7; fi65++ {
        v4 = -541
        switch "hello" {
        case "a":
            switch (fi65 % ((127) | 1)) {
            case 184:
                for fi66 := 0; fi66 < 15; fi66++ {
                    v67 := GenS10{F0: 531}
                    switch (fi65 << ((v5) & 63)) {
                    case -584:
                        v1 += 255
                    case 128:
                        fmt.Println(v5, v1, v4)
                    default:
                        v2 -= v2
                    }
                    fmt.Println(v5, v4)
                    _ = fi66
                    _ = v67
                }
            case -9223372036854775808:
                for fi68 := 0; fi68 < 6; fi68++ {
                    var v69 float32 = 3.14159
                    _ = fi68
                    _ = v69
                }
            }
        case "foo bar":
            var v70 uint8 = 0
            _ = v70
        case "":
            v71 := []int{(fi65 * fi65), (fi65 ^ (fi65 * fi65)), (fi65 << (((v5 - (^(v5)))) & 63)), (fi65 + fi65)}
            _ = v71
        }
        _ = fi65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v8
}
