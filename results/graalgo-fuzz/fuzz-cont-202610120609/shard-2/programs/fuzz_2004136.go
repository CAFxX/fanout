package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint8) int32 {
    var v7 int32 = 0
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: ((1.5 + 0.1) != 0.0), F1: -123, F2: (fi8 >= 42)}
        if false {
            var v10 float64 = 1.5
            _ = v10
        }
        var v11 int32 = v7
        _ = fi8
        _ = v11
        _ = v9
    }
    p0 -= p1
    if true {
        p1 += (p1 & p0)
        for fi12 := 0; fi12 < 15; fi12++ {
            if (!(!false)) {
                v13 := GenS1{F0: true, F1: int8(v7), F2: (532 > 164)}
                switch (fi12 ^ 32768) {
                case -2147483648:
                    v14 := []int{42, 1}
                    _ = v14
                case 105:
                    fmt.Println(fi12, p1)
                default:
                    p1 -= (p0 - 42)
                }
                _ = v13
            } else {
                var v15 int = fi12
                _ = v15
            }
            fmt.Println(p1)
            if (179 != 578) {
                var v16 uint64 = 22
                _ = v16
            } else {
                switch fi12 {
                case 127:
                    if (3.14159 < 0.0) {
                        v17 := GenS1{F0: (false && (-1 != (v7 * 42))), F1: -128, F2: true}
                        if (-2.25 <= float32(float64(-2.25))) {
                            p0 += p0
                            switch (-(fi12)) {
                            case -9223372036854775808:
                                v18 := []int{(^(fi12)), (fi12 % ((2147483647) | 1))}
                                _ = v18
                            default:
                                v19 := GenS1{F0: false, F1: -1, F2: false}
                                _ = v19
                            }
                        } else {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                var v21 bool = true
                                var v22 int32 = (v7 & v7)
                                var v23 float64 = 3.14159
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            switch (fi12 >> ((2) & 63)) {
                            case 1:
                                v7 -= 257
                            default:
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    v25 := []int{fi12, fi24, 114}
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        }
                        _ = v17
                    }
                case 2147483647:
                    for fi26 := 0; fi26 < 8; fi26++ {
                        switch "foo bar" {
                        case "\n\t":
                            p1 = p0
                        case "foo bar":
                            if (738 >= 42) {
                                var v27 uint8 = (p1 | (^(p1)))
                                switch "a" {
                                case "a":
                                    v7 -= (v7 / ((v7) | 1))
                                default:
                                    p0 += 1
                                }
                                _ = v27
                            }
                        case "a":
                            p0 += p0
                        }
                        v28 := []int{fi26, (^(fi26)), fi26, fi26}
                        var v29 bool = false
                        _ = fi26
                        _ = v28
                        _ = v29
                    }
                }
            }
            _ = fi12
        }
        v30 := []int{1, 1, 2147483647}
        _ = v30
    } else {
        fmt.Println(p1)
        v31 := []int{-2147483648, 450, 9223372036854775807}
        switch 9223372036854775807 {
        case -64:
            var v32 uint8 = (p0 & (p0 << ((895) & 63)))
            _ = v32
        default:
            v33 := GenS1{F0: (12 > 127), F1: 42, F2: (!("\n\t" <= ((("" + "a") + "hello") + "")))}
            _ = v33
        }
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v7
    return (^(v7))
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = 2147483647
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 uint64 = 616
    var v6 bool = v3
    v5 = 18446744073709551615
    var v34 float64 = -2.25
    switch v1 {
    case "a":
        var v35 int8 = 100
        _ = v35
    case "\n\t":
        for fi36 := 0; fi36 < 20; fi36++ {
            for fi37 := 0; fi37 < 3; fi37++ {
                var v38 uint8 = 0
                v1 = (((v1 + "hello") + ("a" + (v1 + "\n\t"))) + v1)
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    case "foo bar":
        v2 = v2
    default:
        v39 := GenS2{F0: 42, F1: 930}
        _ = v39
    }
    switch 280 {
    case 9223372036854775807:
        if v6 {
            if (235 >= 4294967295) {
                v6 = (2 <= 42)
                var v40 float32 = 0.1
                _ = v40
            } else {
                var v41 float32 = (0.1 * 1e10)
                v4[-2147483648] = (!(75 > 127))
                v42 := GenS3{F0: v5}
                _ = v41
                _ = v42
            }
            v2 = 109
            if v6 {
                fmt.Println(v6, v5)
                v3 = (!false)
                v43 := []int{9223372036854775807}
                _ = v43
            }
        }
    default:
        var v44 float64 = 0.1
        _ = v44
    }
    var v45 uint16 = 255
    v46 := GenS1{F0: (1 <= 570), F1: 1, F2: ((v5 + v5) != v5)}
    fmt.Println(v5)
    v47 := GenS3{F0: (v5 / ((v5) | 1))}
    v4[128] = true
    if ((!(("hello" + v1) >= v1)) && (863 >= 18446744073709551615)) {
        v3 = v3
    } else {
        fmt.Println(v34, v3, v34)
        v48 := v4[930]
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
