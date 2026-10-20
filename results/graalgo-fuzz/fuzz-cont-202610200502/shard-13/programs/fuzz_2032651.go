package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint8; F1 int8; F2 string }
type GenS7 struct { F0 int16; F1 bool; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint }
func genFn0(p0 int16, p1 string) int8 {
    if (false && ((-1 < 0) || true)) {
        fmt.Println(p0)
        v7 := GenS1{F0: 4294967295, F1: 0, F2: 1}
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{(fi8 << ((230) & 63)), 32768, (-(fi8)), (^(fi8))}
            var v10 uint32 = 159
            v11 := GenS2{F0: p1, F1: 42, F2: -156}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = v7
    } else {
        fmt.Println(p1, p0)
        var v12 int8 = -1
        fmt.Println(p1, p1)
        _ = v12
    }
    v13 := GenS3{F0: 9223372036854775807}
    switch -9223372036854775808 {
    case 0:
        switch p1 {
        case "hello":
            p0 += (p0 + -1)
        default:
            switch ("" + p1) {
            case "":
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := []int{(fi14 << ((86) & 63))}
                    if ((!(289 < 0)) && true) {
                        var v16 float64 = 1e10
                        for fi17 := 0; fi17 < 3; fi17++ {
                            var v18 int32 = 42
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
            default:
                v19 := []int{-711, -1, 753}
                _ = v19
            }
        }
    case 128:
        var v20 uint16 = 1
        _ = v20
    case -1:
        p1 = p1
    }
    v21 := GenS3{F0: 42}
    _ = p0
    _ = p1
    _ = v13
    _ = v21
    return 0
}
func genFn1(p0 int16) bool {
    p0 += 1
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 bool = (false && ((0 > 179) || true))
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        if (!true) {
            var v25 uint64 = 0
            v25 -= v25
            _ = v25
        }
        if (!false) {
            if (true || (65535 <= 558)) {
                var v26 int8 = -126
                var v27 int64 = 0
                _ = v26
                _ = v27
            }
            if false {
                if false {
                    v28 := GenS2{F0: "\n\t", F1: 100, F2: (^(fi24))}
                    fmt.Println(p0)
                    switch fi24 {
                    case 43:
                        for fi29 := 0; fi29 < 14; fi29++ {
                            var v30 float32 = (0.0 * (3.14159 - -2.25))
                            _ = fi29
                            _ = v30
                        }
                    default:
                        if (0 != 4294967295) {
                            var v31 float32 = 1e10
                            _ = v31
                        } else {
                            var v32 string = "a"
                            switch fi24 {
                            case 4294967295:
                                var v33 bool = true
                                _ = v33
                            }
                            fmt.Println(v32)
                            _ = v32
                        }
                    }
                    _ = v28
                }
                var v34 uint16 = 32767
                var v35 uint64 = 0
                _ = v34
                _ = v35
            } else {
                var v36 float32 = ((1e10 * float32((3.14159 / (float64(-2.25) - ((0.1 + -2.25) - -2.25))))) * -2.25)
                p0 = (p0 - (p0 / ((-547) | 1)))
                _ = v36
            }
        }
        _ = fi24
    }
    if (0 != 0) {
        switch "\n\t" {
        case "hello":
            v37 := []int{1}
            _ = v37
        case "":
            v38 := GenS4{F0: 2147483647, F1: -432, F2: 561}
            _ = v38
        case "foo bar":
            var v39 uint8 = 255
            _ = v39
        default:
            var v40 int8 = 47
            _ = v40
        }
    } else {
        var v41 float32 = 1e10
        if false {
            var v42 float64 = 0.0
            var v43 int16 = (p0 + (p0 % ((p0) | 1)))
            _ = v42
            _ = v43
        } else {
            v44 := GenS5{F0: 65535, F1: 42}
            _ = v44
        }
        v45 := GenS6{F0: 192, F1: genFn0(p0, ""), F2: "foo bar"}
        _ = v41
        _ = v45
    }
    _ = p0
    return (699 < 255)
}

func main() {
    var v1 string = (("hello" + "") + (("foo bar" + ("a" + "a")) + ""))
    var v2 uint64 = 4294967295
    var v3 uint16 = 777
    var v4 uint32 = 610
    var v5 uint16 = 42
    var v6 uint64 = (v2 + (v2 - v2))
    switch "\n\t" {
    case "a":
        fmt.Println(v6, v1, v2)
    case "hello":
        switch -874 {
        case 32767:
            v46 := GenS7{F0: 156, F1: true, F2: (v5 ^ (v5 + v5))}
            _ = v46
        }
    case "foo bar":
        if ((1 >= 42) || (!((!(true || true)) && true))) {
            v47 := GenS3{F0: 59}
            v48 := GenS1{F0: v2, F1: 2147483647, F2: (v6 / ((v2) | 1))}
            v49 := []int{32767, 4294967295}
            _ = v47
            _ = v48
            _ = v49
        }
    }
    v50 := []int{-1, -246, -2147483648, -2147483648}
    if true {
        switch v1 {
        case "foo bar":
            for fi51 := 0; fi51 < 16; fi51++ {
                var v52 uint8 = 64
                switch ("" + "hello") {
                case "\n\t":
                    v5 += (v5 % ((v5) | 1))
                case "a":
                    v53 := GenS8{F0: false, F1: 4294967295}
                    _ = v53
                default:
                    v54 := GenS5{F0: (v3 << ((uint(v52)) & 63)), F1: -444}
                    _ = v54
                }
                fmt.Println(v3)
                _ = fi51
                _ = v52
            }
        case "a":
            v55 := GenS5{F0: (v5 | (^(v3))), F1: 30}
            _ = v55
        case "\n\t":
            if true {
                v50 = append(v50, -460)
            } else {
                for fi56 := 0; fi56 < 5; fi56++ {
                    for fi57 := 0; fi57 < 16; fi57++ {
                        fmt.Println(v3)
                        _ = fi57
                    }
                    var v58 float64 = (-2.25 + (3.14159 + 0.0))
                    v59 := GenS5{F0: 1, F1: -39}
                    _ = fi56
                    _ = v58
                    _ = v59
                }
            }
        }
        v6 = v2
    }
    _ = len(v50)
    v3 = v3
    for fi60 := 0; fi60 < 11; fi60++ {
        v50 = append(v50, fi60)
        var v61 float32 = 0.1
        _ = fi60
        _ = v61
    }
    fmt.Println(v4, v6, v6)
    for fi62 := 0; fi62 < 13; fi62++ {
        v6 = (v2 ^ (v6 >> ((4294967295) & 63)))
        _ = len(v50)
        _ = fi62
    }
    var v63 uint16 = (v5 ^ v5)
    var v64 int8 = -1
    if len(v50) >= 2 {
        v50 = v50[1:len(v50)]
    }
    if false {
        var v65 int16 = 42
        _ = v65
    } else {
        fmt.Println(v2)
        v1 = "hello"
        v2 -= 0
    }
    switch -134 {
    case -992:
        if true {
            v63 = (v63 & (^(v5)))
            v1 = v1
            var v66 uint16 = (v3 % ((v63) | 1))
            _ = v66
        }
    }
    v67 := GenS6{F0: 0, F1: (v64 + 42), F2: (v1 + (((v1 + v1) + ((v1 + ("hello" + v1)) + v1)) + ("foo bar" + (v1 + ""))))}
    _ = len(v50)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v63
    _ = v64
    _ = v67
}
