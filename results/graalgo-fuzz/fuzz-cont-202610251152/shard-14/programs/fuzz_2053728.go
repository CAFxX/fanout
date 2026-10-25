package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 bool; F1 int16; F2 int }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint }
func genFn0() int64 {
    switch 32767 {
    case 2147483647:
        if false {
            var v4 int32 = 1
            _ = v4
        } else {
            var v5 int32 = -1
            var v6 int32 = v5
            _ = v5
            _ = v6
        }
    case 9223372036854775807:
        v7 := GenS1{F0: 128}
        _ = v7
    case 42:
        if true {
            var v8 uint = 1
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi9, v8)
                v10 := GenS2{F0: 127}
                _ = fi9
                _ = v10
            }
            v11 := []int{}
            _ = v11
            _ = v8
        }
    default:
        switch (("a" + ("hello" + "")) + "\n\t") {
        case "\n\t":
            var v12 int = 1
            _ = v12
        case "foo bar":
            switch "hello" {
            case "hello":
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := []int{32768}
                    var v15 int16 = 32767
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "":
                if (-1 == -1) {
                    var v16 int32 = 695
                    v17 := []int{-9223372036854775808, 32767, 2147483647, 42}
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v19 := GenS2{F0: 127}
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v16
                    _ = v17
                } else {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 bool = true
                        var v22 float64 = 3.14159
                        v22 += v22
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            case "\n\t":
                var v23 float64 = (-2.25 + 0.0)
                _ = v23
            default:
                var v24 int8 = 0
                _ = v24
            }
        }
    }
    v25 := []int{1, -2147483648, 32767, 32767}
    if true {
        v26 := GenS3{F0: 18446744073709551615, F1: 325}
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := GenS3{F0: 0, F1: 42}
        var v29 uint16 = 638
        v30 := GenS3{F0: 2, F1: 2}
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    _ = v25
    return -9223372036854775808
}
func genFn1(p0 uint32) bool {
    switch -9223372036854775808 {
    case 127:
        fmt.Println(p0)
    }
    switch ("" + "foo bar") {
    case "a":
        var v31 int32 = 1
        _ = v31
    default:
        var v32 uint64 = 41
        _ = v32
    }
    v33 := GenS1{F0: 192}
    v34 := []int{1, -792, -33}
    if false {
        for fi35 := 0; fi35 < 13; fi35++ {
            v36 := GenS1{F0: 1}
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 uint = 18446744073709551615
                p0 -= (p0 | 548)
                _ = fi37
                _ = v38
            }
            _ = fi35
            _ = v36
        }
    } else {
        v39 := GenS4{F0: 32767, F1: -751, F2: 251}
        switch 42 {
        case 1:
            p0 = (p0 ^ 545)
        case -138:
            var v40 bool = true
            _ = v40
        case 4294967295:
            var v41 float64 = 0.0
            _ = v41
        default:
            switch ("hello" + "") {
            case "a":
                for fi42 := 0; fi42 < 1; fi42++ {
                    var v43 float32 = -2.25
                    v44 := GenS5{F0: -1}
                    for fi45 := 0; fi45 < 16; fi45++ {
                        if false {
                            p0 -= p0
                            v46 := GenS5{F0: 0}
                            v43 = 1e10
                            _ = v46
                        } else {
                            v43 = float32(0.0)
                            var v47 string = "hello"
                            switch fi45 {
                            case -375:
                                var v48 int8 = 72
                                _ = v48
                            }
                            _ = v47
                        }
                        v43 += v43
                        _ = fi45
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            case "\n\t":
                switch "hello" {
                case "a":
                    for fi49 := 0; fi49 < 10; fi49++ {
                        var v50 float32 = (0.0 * 0.1)
                        p0 += (p0 + p0)
                        v50 = (((v50 * (((((float32((((0.0 * 1.5) - 1e10) / 1e10)) / float32(((-2.25 / 1.0) - 3.14159))) + (float32(0.1) + 3.14159)) * 0.0) / 1.0) / (-2.25 / 1.0))) - v50) * v50)
                        _ = fi49
                        _ = v50
                    }
                case "foo bar":
                    var v51 int16 = -1
                    _ = v51
                case "":
                    p0 -= p0
                }
            case "":
                if true {
                    if len(v34) > 0 {
                        v34[0] = -256
                    }
                    var v52 float32 = -2.25
                    _ = v52
                } else {
                    if len(v34) >= 2 {
                        v34 = v34[1:len(v34)]
                    }
                }
            }
        }
        v53 := GenS6{F0: 127, F1: p0}
        _ = v39
        _ = v53
    }
    _ = p0
    _ = v33
    _ = v34
    return true
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]int)
    v3 := []int{}
    for fi54 := 0; fi54 < 12; fi54++ {
        if (!false) {
            var v55 bool = (108 == int8(v1))
            _ = v55
        }
        for fi56 := 0; fi56 < 4; fi56++ {
            var v57 int32 = 42
            if false {
                switch "" {
                case "":
                    v57 -= -2147483648
                    switch 32768 {
                    case 628:
                        if (!false) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                default:
                    var v58 uint64 = v1
                    _ = v58
                }
                var v59 int64 = 42
                _ = v59
            }
            _ = fi56
            _ = v57
        }
        _ = fi54
    }
    v60 := GenS7{F0: false, F1: 565, F2: 128}
    fmt.Println(v1)
    if true {
        v1 -= v1
    } else {
        v61 := v2[("" + "foo bar")]
        _ = v61
    }
    fmt.Println(v1)
    var v62 string = "a"
    switch "foo bar" {
    case "\n\t":
        if (24 != 1) {
            fmt.Println(v62, v1)
            fmt.Println(v62, v1)
            v1 += 273
        } else {
            var v63 float64 = 1e10
            _ = v63
        }
    case "a":
        v64 := GenS8{F0: 1, F1: v1, F2: 2}
        _ = v64
    case "":
        v65 := GenS6{F0: -128, F1: 4294967295}
        _ = v65
    default:
        var v66 uint64 = v1
        _ = v66
    }
    v62 = (v62 + "hello")
    switch 128 {
    case 128:
        v67 := GenS1{F0: 115}
        _ = v67
    case 4294967295:
        v68 := GenS7{F0: (v1 >= 42), F1: 42, F2: 2147483647}
        _ = v68
    case 736:
        v69 := GenS6{F0: 42, F1: 0}
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v62
}
