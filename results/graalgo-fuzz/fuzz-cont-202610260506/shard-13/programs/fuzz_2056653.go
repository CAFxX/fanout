package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int8 = 42
    var v2 uint = 2
    var v3 uint = 999
    var v4 uint8 = 2
    switch -712 {
    case 42:
        var v5 int8 = (v1 << (((v2 + v2)) & 63))
        _ = v5
    case 32767:
        if (!(false || true)) {
            var v6 string = "foo bar"
            _ = v6
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v3 -= v2
        default:
            for fi7 := 0; fi7 < 8; fi7++ {
                if true {
                    if (!false) {
                        fmt.Println(v1, fi7, v4)
                        v2 -= (v2 | v3)
                    }
                    v8 := []int{2147483647}
                    if (!true) {
                        for fi9 := 0; fi9 < 11; fi9++ {
                            var v10 string = "\n\t"
                            if ((-1 <= 1) || (-2.25 >= 1.5)) {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                var v11 uint8 = 42
                                v12 := GenS1{F0: v1, F1: 0, F2: (v3 | v3)}
                                _ = v11
                                _ = v12
                            } else {
                                v1 -= v1
                                v8 = append(v8, 32768)
                                var v13 int = -2147483648
                                v2 += 255
                                _ = v13
                            }
                            _ = fi9
                            _ = v10
                        }
                        var v14 uint8 = v4
                        var v15 int32 = -666
                        _ = v14
                        _ = v15
                    }
                    _ = v8
                }
                _ = fi7
            }
        }
    }
    if false {
        v16 := []int{2147483647, -916, -842, 32767}
        v4 += v4
        _ = v16
    } else {
        if false {
            fmt.Println(v3)
        } else {
            v17 := []int{}
            _ = v17
        }
        v1 = (-(v1))
        v18 := GenS1{F0: (-(v1)), F1: 228, F2: 255}
        _ = v18
    }
    if (!true) {
        var v19 int8 = (v1 / ((v1) | 1))
        _ = v19
    }
    fmt.Println(v1, v2)
    for fi20 := 0; fi20 < 15; fi20++ {
        v1 += v1
        _ = fi20
    }
    v21 := GenS2{F0: 343, F1: 65535}
    var v22 int8 = v1
    v22 -= 127
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := []int{(fi23 - -1), (fi23 | fi23), (fi23 - fi23)}
        fmt.Println(v2)
        switch fi23 {
        case -2147483648:
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        default:
            v2 = 509
        }
        _ = fi23
        _ = v24
    }
    v22 -= 127
    if ((!(18446744073709551615 != 18446744073709551615)) && false) {
        for fi25 := 0; fi25 < 6; fi25++ {
            if (!false) {
                v1 = (v22 * v22)
                switch fi25 {
                case 32767:
                    var v26 int16 = 587
                    _ = v26
                case 1:
                    var v27 uint8 = 42
                    _ = v27
                default:
                    v28 := []int{fi25, fi25, (fi25 & fi25)}
                    _ = v28
                }
                v29 := []int{(fi25 / ((fi25) | 1))}
                _ = v29
            } else {
                v3 -= v3
                switch "" {
                case "a":
                    var v30 uint = (v3 >> (((v2 | (v3 << ((v2) & 63)))) & 63))
                    _ = v30
                }
                if (270 != 1) {
                    switch "" {
                    case "a":
                        fmt.Println(v22, v2, v2)
                    case "\n\t":
                        for fi31 := 0; fi31 < 12; fi31++ {
                            v22 = (v1 / (((v1 << ((v2) & 63))) | 1))
                            v1 -= v22
                            v32 := GenS2{F0: 890, F1: 301}
                            v1 -= (v1 / ((v22) | 1))
                            _ = fi31
                            _ = v32
                        }
                    default:
                        fmt.Println(v22)
                    }
                }
            }
            _ = fi25
        }
        var v33 int32 = 2147483647
        _ = v33
    }
    var v34 int16 = 1000
    v35 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
