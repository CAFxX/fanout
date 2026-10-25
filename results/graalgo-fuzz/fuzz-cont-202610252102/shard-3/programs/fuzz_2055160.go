package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 int16 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 int; F1 uint64; F2 int16 }
func genFn0(p0 int8) int8 {
    switch 42 {
    case -9223372036854775808:
        switch 42 {
        case 32768:
            p0 += p0
        case 0:
            v6 := []int{-1, 2147483647, -9223372036854775808, 2147483647}
            _ = v6
        default:
            v7 := GenS1{F0: true, F1: p0}
            _ = v7
        }
    default:
        var v8 uint64 = 999
        _ = v8
    }
    var v9 string = ("\n\t" + "\n\t")
    var v10 uint64 = 0
    if ((v10 >> ((614) & 63)) >= v10) {
        if true {
            if (false && (0 > 255)) {
                v11 := []int{128, -823}
                _ = len(v11)
                if len(v11) > 0 {
                    v11[0] = 42
                }
                _ = v11
            }
            fmt.Println(v10, p0, v10)
            v12 := GenS2{F0: (v10 - v10), F1: (v10 & (v10 & (^(v10)))), F2: (v9 + v9)}
            _ = v12
        } else {
            fmt.Println(v10)
        }
        var v13 int16 = 0
        switch 20 {
        case -9223372036854775808:
            fmt.Println(v10)
        case 2147483647:
            if true {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := GenS3{F0: 255, F1: 676, F2: 1}
                    _ = fi14
                    _ = v15
                }
                v10 = v10
                v13 = (-(v13))
            } else {
                var v16 uint16 = 891
                v17 := GenS4{F0: "hello", F1: 415, F2: 1000}
                for fi18 := 0; fi18 < 6; fi18++ {
                    fmt.Println(v13)
                    v9 = v9
                    var v19 uint = 255
                    _ = fi18
                    _ = v19
                }
                _ = v16
                _ = v17
            }
        default:
            v20 := []int{529, 127, 4294967295}
            _ = v20
        }
        _ = v13
    } else {
        v21 := GenS4{F0: v9, F1: 255, F2: 1}
        _ = v21
    }
    switch v9 {
    case "\n\t":
        v10 += v10
    case "hello":
        v22 := GenS5{F0: 0, F1: p0}
        _ = v22
    case "a":
        if false {
            v23 := GenS4{F0: v9, F1: 255, F2: 0}
            v24 := GenS2{F0: (v10 << ((929) & 63)), F1: v10, F2: v9}
            v25 := GenS6{F0: 9223372036854775807, F1: (v10 / ((2) | 1)), F2: 42}
            _ = v23
            _ = v24
            _ = v25
        }
    default:
        v26 := []int{}
        _ = v26
    }
    _ = p0
    _ = v10
    _ = v9
    return p0
}
func genFn1(p0 uint8) int {
    p0 += (p0 & (p0 ^ p0))
    var v27 uint = 42
    var v28 uint16 = 65535
    v29 := []int{}
    v27 -= 18446744073709551615
    _ = p0
    _ = v27
    _ = v28
    _ = v29
    return 128
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -1
    var v3 uint64 = 396
    var v4 float64 = (0.1 * 3.14159)
    var v5 uint = 42
    if (1 == 100) {
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 int32 = -594
            v32 := []int{(fi30 & (fi30 ^ (fi30 + 584))), fi30}
            _ = fi30
            _ = v31
            _ = v32
        }
        switch "\n\t" {
        case "hello":
            v2 -= (v2 + 2147483647)
        case "":
            switch 128 {
            case -60:
                fmt.Println(v1)
            case -2147483648:
                if true {
                    switch genFn1(255) {
                    case -794:
                        if (32767 == -1) {
                            v33 := []int{127, 32768, 577}
                            _ = v33
                        } else {
                            fmt.Println(v5, v4, v3)
                            fmt.Println(v4, v1, v5)
                        }
                    }
                } else {
                    v4 += (((1.5 + (1e10 * v4)) * v4) + float64((3.14159 + 3.14159)))
                    var v34 uint64 = uint64(v5)
                    switch ("foo bar" + (("a" + "foo bar") + "hello")) {
                    case "foo bar":
                        v1 = ((v1 * ((0.0 - v1) * 1e10)) / -2.25)
                    }
                    _ = v34
                }
            case -1:
                v35 := GenS6{F0: -831, F1: v3, F2: -507}
                _ = v35
            default:
                if true {
                    v5 -= (v5 / ((v5) | 1))
                    var v36 uint16 = 2
                    if true {
                        v37 := GenS3{F0: (^(v36)), F1: (v5 * uint(v36)), F2: 127}
                        _ = v37
                    }
                    _ = v36
                }
            }
        }
        var v38 uint32 = 2
        _ = v38
    } else {
        v39 := GenS1{F0: (false || true), F1: -1}
        v5 = (v5 - (v5 | v5))
        fmt.Println(v3, v2, v3)
        _ = v39
    }
    v3 -= 18446744073709551615
    switch "hello" {
    case "\n\t":
        for fi40 := 0; fi40 < 3; fi40++ {
            v41 := []int{fi40, fi40, -318}
            if (v3 != v3) {
                v41 = append(v41, genFn1(1))
                for fi42 := 0; fi42 < 14; fi42++ {
                    var v43 string = "foo bar"
                    switch fi40 {
                    case 128:
                        v44 := GenS5{F0: -1, F1: genFn0(100)}
                        _ = v44
                    }
                    _ = fi42
                    _ = v43
                }
            }
            _ = fi40
            _ = v41
        }
    default:
        switch "" {
        case "hello":
            fmt.Println(v1, v5, v4)
        }
    }
    v45 := []int{2147483647}
    var v46 uint64 = (^(v3))
    fmt.Println(v4, v46)
    v3 += v3
    for fi47 := 0; fi47 < 15; fi47++ {
        if len(v45) >= 2 {
            v45 = v45[1:len(v45)]
        }
        _ = fi47
    }
    var v48 float64 = v4
    v46 = v3
    _ = len(v45)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v5
}
