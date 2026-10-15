package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
func genFn0(p0 int) int8 {
    v6 := GenS1{F0: (2 > 0), F1: -9223372036854775808, F2: "foo bar"}
    switch "" {
    case "\n\t":
        if ((^(p0)) > p0) {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS2{F0: -32768}
                v9 := []int{fi7}
                if true {
                    switch "\n\t" {
                    case "a":
                        if (true || (627 > 42)) {
                            var v10 string = "hello"
                            if false {
                                v10 = v10
                                fmt.Println(p0, p0)
                                switch ("foo bar" + v10) {
                                case "\n\t":
                                    v11 := GenS3{F0: -128, F1: v10, F2: 42}
                                    _ = v11
                                case "hello":
                                    p0 += fi7
                                default:
                                    if len(v9) > 0 {
                                        v9[0] = (p0 | (-(fi7)))
                                    }
                                }
                            } else {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            for fi12 := 0; fi12 < 7; fi12++ {
                                v13 := GenS3{F0: 1, F1: v10, F2: -128}
                                v10 = "hello"
                                _ = fi12
                                _ = v13
                            }
                            _ = v10
                        } else {
                            if len(v9) > 0 {
                                v9[0] = fi7
                            }
                        }
                    default:
                        for fi14 := 0; fi14 < 18; fi14++ {
                            if (4294967295 == 4294967295) {
                                if len(v9) > 0 {
                                    v9[0] = p0
                                }
                            } else {
                                fmt.Println(fi14, fi14, fi7)
                            }
                            p0 += 32767
                            switch ("hello" + "a") {
                            case "a":
                                switch "a" {
                                case "":
                                    v15 := GenS4{F0: (!(true || true)), F1: -2147483648, F2: 2147483647}
                                    _ = v15
                                case "a":
                                    var v16 uint16 = 1
                                    _ = v16
                                }
                            default:
                                switch p0 {
                                case -161:
                                    var v17 string = (("hello" + "\n\t") + "foo bar")
                                    _ = v17
                                case 1:
                                    if len(v9) > 0 {
                                        v9[0] = -704
                                    }
                                case -1:
                                    v18 := GenS5{F0: "hello"}
                                    _ = v18
                                }
                            }
                            _ = fi14
                        }
                    }
                    _ = len(v9)
                    if true {
                        p0 += -1
                        fmt.Println(fi7, fi7)
                    }
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            p0 -= 128
            p0 = p0
        }
    case "hello":
        var v19 string = "a"
        _ = v19
    case "a":
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    return 100
}
func genFn1(p0 int16) int8 {
    var v20 int = 42
    switch v20 {
    case -992:
        v21 := []int{128, v20, v20}
        _ = v21
    case 365:
        switch 827 {
        case 497:
            if false {
                v20 = v20
                var v22 float32 = 1e10
                var v23 bool = true
                _ = v22
                _ = v23
            }
        case -9223372036854775808:
            var v24 float32 = (((0.0 + 1e10) * (float32(0.1) - 0.0)) - -2.25)
            _ = v24
        case 1:
            if (0 > 0) {
                v25 := GenS3{F0: int8(v20), F1: "a", F2: -1}
                for fi26 := 0; fi26 < 16; fi26++ {
                    p0 = 1000
                    v27 := GenS4{F0: false, F1: 803, F2: 4294967295}
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        }
    }
    _ = p0
    _ = v20
    return 97
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1
    var v3 uint16 = 2
    v4 := make(map[int]uint32)
    var v5 int16 = (-(v2))
    var v28 int64 = 1
    var v29 int64 = (v28 + v28)
    var v30 int32 = 42
    for fi31 := 0; fi31 < 16; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 uint8 = 128
            v34 := GenS6{F0: (v2 >> ((217) & 63))}
            _ = fi32
            _ = v33
            _ = v34
        }
        if true {
            fmt.Println(fi31, v2)
            v35 := GenS1{F0: (-1 >= 1000), F1: (v28 * v29), F2: "foo bar"}
            for fi36 := 0; fi36 < 1; fi36++ {
                v5 = v5
                _ = fi36
            }
            _ = v35
        } else {
            fmt.Println(v2)
            if true {
                v37 := GenS6{F0: v2}
                _ = v37
            } else {
                v38 := GenS4{F0: true, F1: 1, F2: 42}
                var v39 float64 = ((3.14159 + (1e10 + (float64(3.14159) * 1e10))) * (1e10 - 3.14159))
                _ = v38
                _ = v39
            }
        }
        switch fi31 {
        case -721:
            var v40 bool = true
            _ = v40
        case -468:
            v41 := []int{-865, 127}
            _ = v41
        default:
            v29 = (v29 | v28)
        }
        _ = fi31
    }
    v29 -= v28
    for fi42 := 0; fi42 < 17; fi42++ {
        v43 := GenS3{F0: 0, F1: "foo bar", F2: -128}
        delete(v4, fi42)
        if ((fi42 - (fi42 | fi42)) > -2147483648) {
            v44 := GenS6{F0: (v2 << ((255) & 63))}
            fmt.Println(fi42, fi42)
            _ = v44
        } else {
            v45 := []int{fi42, (fi42 / (((-(fi42))) | 1))}
            v29 = v29
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    fmt.Println(v2)
    v2 -= v2
    fmt.Println(v30, v2)
    v46 := []int{}
    var v47 int16 = v5
    v5 = v2
    fmt.Println(v47)
    v3 = (v3 - 177)
    if true {
        var v48 uint16 = 87
        if true {
            v48 += v3
        }
        for fi49 := 0; fi49 < 6; fi49++ {
            fmt.Println(fi49)
            _ = fi49
        }
        _ = v48
    }
    v4[-232] = 250
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
