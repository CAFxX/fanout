package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 uint64; F2 int }
type GenS8 struct { F0 int32; F1 int32; F2 string }
func genFn0() int {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint64 = 2
        _ = fi5
        _ = v6
    }
    if (!true) {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint = 248
            _ = fi7
            _ = v8
        }
    }
    v9 := []int{127}
    var v10 uint64 = 2
    var v11 bool = false
    _ = v10
    _ = v11
    _ = v9
    return 32767
}
func genFn1(p0 uint) bool {
    p0 = p0
    var v12 int = -185
    switch (("foo bar" + (("\n\t" + "hello") + "foo bar")) + "a") {
    case "a":
        switch (v12 ^ (v12 / (((v12 | genFn0())) | 1))) {
        case -250:
            if true {
                switch "foo bar" {
                case "\n\t":
                    var v13 uint16 = 2
                    _ = v13
                case "foo bar":
                    v14 := GenS1{F0: (("a" + "") + "\n\t"), F1: 65535}
                    _ = v14
                case "a":
                    switch v12 {
                    case 9223372036854775807:
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v12)
                        }
                    case 2147483647:
                        v15 := GenS2{F0: 826, F1: 2147483647, F2: 1}
                        _ = v15
                    }
                }
                p0 = 661
                v16 := GenS3{F0: genFn0(), F1: 1, F2: 65535}
                _ = v16
            }
        case 9223372036854775807:
            v17 := GenS4{F0: 1, F1: 42}
            _ = v17
        case 42:
            if false {
                fmt.Println(p0)
                if false {
                    switch v12 {
                    case 328:
                        v18 := []int{(^(v12)), 167, (v12 | (v12 - v12))}
                        _ = v18
                    }
                    v19 := []int{(v12 | (v12 ^ 127)), (v12 - v12)}
                    v12 = 32768
                    _ = v19
                } else {
                    fmt.Println(v12, p0)
                }
            } else {
                v20 := GenS1{F0: "hello", F1: 26}
                if false {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        var v22 uint64 = 256
                        _ = fi21
                        _ = v22
                    }
                    v23 := []int{genFn0(), v12, genFn0(), v12}
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    switch "foo bar" {
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            if false {
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    p0 = 2
                                    if (42 < 1000) {
                                        fmt.Println(fi24, fi24)
                                        var v26 float32 = 0.1
                                        var v27 uint16 = 2
                                        _ = v26
                                        _ = v27
                                    }
                                    var v28 uint = 26
                                    _ = fi25
                                    _ = v28
                                }
                                if (-365 >= 0) {
                                    switch genFn0() {
                                    case 1:
                                        var v29 float32 = ((0.1 + 0.1) - -2.25)
                                        _ = v29
                                    }
                                }
                                v12 += v12
                            }
                        case "":
                            v12 = (fi24 * 32767)
                        default:
                            fmt.Println(v12)
                        }
                    case "\n\t":
                        switch genFn0() {
                        case 2147483647:
                            var v30 int8 = -14
                            _ = v30
                        case -769:
                            switch v12 {
                            case -2147483648:
                                var v31 int8 = 0
                                _ = v31
                            }
                        case -2147483648:
                            if false {
                                for fi32 := 0; fi32 < 14; fi32++ {
                                    var v33 int16 = 314
                                    v12 += 127
                                    fmt.Println(fi24)
                                    _ = fi32
                                    _ = v33
                                }
                                v34 := []int{263}
                                if (2147483647 != 884) {
                                    v12 -= v12
                                    if len(v34) > 0 {
                                        v34[0] = v12
                                    }
                                } else {
                                    if len(v34) > 0 {
                                        v34[0] = (fi24 | 128)
                                    }
                                }
                                _ = v34
                            }
                        }
                    }
                    _ = fi24
                }
                _ = v20
            }
        default:
            var v35 string = "foo bar"
            _ = v35
        }
    case "\n\t":
        var v36 uint32 = 2
        _ = v36
    default:
        v12 = (v12 ^ -2147483648)
    }
    p0 -= 2
    if (-2.25 >= 3.14159) {
        v12 = v12
        v37 := GenS5{F0: false, F1: 0}
        if false {
            var v38 bool = ((float64(float32((1e10 + (float64((1e10 * 0.0)) * 3.14159)))) / 0.1) > -2.25)
            for fi39 := 0; fi39 < 20; fi39++ {
                p0 -= (p0 & p0)
                fmt.Println(v12, p0, v12)
                v12 = genFn0()
                _ = fi39
            }
            p0 = (p0 | (p0 & (p0 / (((p0 / ((368) | 1))) | 1))))
            _ = v38
        }
        _ = v37
    } else {
        fmt.Println(v12)
    }
    _ = p0
    _ = v12
    return false
}

func main() {
    var v1 int = 0
    var v2 float64 = ((1e10 + (-2.25 + (1.5 / (1e10 - 0.1)))) / 1.5)
    var v3 int = v1
    var v4 uint = 614
    v2 -= float64(((1e10 / 1.0) + 1e10))
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(v3)
        _ = fi40
    }
    if true {
        v41 := GenS5{F0: true, F1: 434}
        var v42 uint = (^(v4))
        _ = v41
        _ = v42
    } else {
        var v43 string = ("" + ("hello" + "hello"))
        v44 := GenS6{F0: true}
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        if (-1 >= 1) {
            var v46 int = (fi45 + (^(v1)))
            v2 -= v2
            _ = v46
        } else {
            v2 -= float64((1e10 - ((1.5 / ((0.1 + (((0.0 + (0.0 - 1e10)) + 1.5) + -2.25)) / ((-2.25 + float32(1e10)) * float32(((v2 / 1.0) / v2))))) + 0.0)))
            for fi47 := 0; fi47 < 14; fi47++ {
                if (!((float32((v2 - v2)) * 3.14159) <= ((1.5 - 1e10) * float32(float64(1e10))))) {
                    switch v1 {
                    case -1:
                        var v48 uint = (v4 ^ v4)
                        _ = v48
                    }
                }
                _ = fi47
            }
            var v49 int32 = -2147483648
            _ = v49
        }
        _ = fi45
    }
    switch "foo bar" {
    case "a":
        var v50 string = ("foo bar" + "")
        _ = v50
    case "":
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 int16 = -32768
            if (!(!(!(true || true)))) {
                var v53 string = (("a" + ("" + "a")) + ("\n\t" + (("foo bar" + "foo bar") + "")))
                _ = v53
            } else {
                v3 -= (v3 | v3)
                fmt.Println(fi51, fi51)
            }
            _ = fi51
            _ = v52
        }
    case "foo bar":
        if false {
            v54 := GenS7{F0: 4294967295, F1: 118, F2: (v1 * -747)}
            _ = v54
        }
    }
    switch v3 {
    case -1:
        fmt.Println(v2)
    case 867:
        if (370 == 147) {
            v55 := []int{}
            var v56 uint = 1
            _ = v55
            _ = v56
        }
    case 32767:
        for fi57 := 0; fi57 < 14; fi57++ {
            var v58 float32 = 3.14159
            for fi59 := 0; fi59 < 5; fi59++ {
                v60 := []int{v3, fi59, (fi59 + fi57), v1}
                fmt.Println(v2, fi59, v4)
                v58 -= ((float32((float64(0.1) + float64(0.1))) * 0.1) / v58)
                _ = fi59
                _ = v60
            }
            _ = fi57
            _ = v58
        }
    }
    fmt.Println(v3, v1, v3)
    var v61 uint32 = 65535
    var v62 bool = true
    v63 := []int{}
    v64 := GenS8{F0: -2147483648, F1: 1, F2: "a"}
    switch -265 {
    case 32768:
        v1 += (v1 - 0)
    case 0:
        var v65 uint = v4
        _ = v65
    default:
        fmt.Println(v2)
    }
    v4 = 1
    for fi66 := 0; fi66 < 2; fi66++ {
        if true {
            v67 := GenS5{F0: false, F1: (v61 / ((981) | 1))}
            _ = v67
        }
        _ = fi66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
