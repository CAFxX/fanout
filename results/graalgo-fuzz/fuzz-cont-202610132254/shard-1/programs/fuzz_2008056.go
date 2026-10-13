package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint16) uint8 {
    switch "hello" {
    case "a":
        v5 := []int{199, 9223372036854775807}
        _ = v5
    }
    var v6 float64 = 1.5
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            v6 = 1.5
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int64 = 0
                v9 += (v9 % (((v9 / ((v9) | 1))) | 1))
                fmt.Println(v6, fi7)
                fmt.Println(v6, p0)
                if ((-568 < 104) && ((p0 * p0) >= p0)) {
                    var v10 int = fi8
                    v11 := GenS1{F0: 170}
                    v12 := []int{-362}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    if ((2 > (p0 & (p0 << ((2) & 63)))) && (!(false && (false || (!true))))) {
                        var v13 int = 663
                        _ = v13
                    } else {
                        fmt.Println(p0, fi8, p0)
                    }
                    var v14 int16 = 0
                    _ = v14
                }
                _ = fi8
                _ = v9
            }
            p0 += p0
        } else {
            v15 := GenS2{F0: 0, F1: 125}
            _ = v15
        }
        fmt.Println(p0, v6, v6)
        _ = fi7
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 int8 = -1
    var v2 uint = 0
    var v3 int8 = v1
    v4 := []int{-1, -9223372036854775808, 2147483647}
    v16 := GenS2{F0: 2147483647, F1: 127}
    var v17 float64 = ((-2.25 * (1e10 + float64((0.1 * (0.0 - -2.25))))) * (1e10 * (0.1 * 0.1)))
    v18 := GenS1{F0: 1}
    v3 = (v1 | (v1 * 1))
    for fi19 := 0; fi19 < 5; fi19++ {
        switch (fi19 ^ fi19) {
        case 32768:
            v20 := GenS2{F0: 0, F1: 127}
            _ = v20
        case -1:
            var v21 bool = (!true)
            _ = v21
        }
        switch (fi19 >> (((v2 ^ v2)) & 63)) {
        case 4294967295:
            var v22 float32 = 0.1
            _ = v22
        case 32767:
            if false {
                v23 := GenS3{F0: 105}
                _ = v23
            } else {
                fmt.Println(v3, fi19, v3)
                if (1 != 0) {
                    v2 = v2
                    _ = len(v4)
                }
                var v24 int64 = 9223372036854775807
                _ = v24
            }
        }
        _ = fi19
    }
    var v25 float32 = -2.25
    var v26 uint64 = 4294967295
    switch "" {
    case "a":
        v27 := GenS3{F0: -2147483648}
        _ = v27
    }
    switch -9223372036854775808 {
    case 503:
        v3 -= v3
    case 174:
        v26 = v26
    case 0:
        v3 -= (v1 / ((-84) | 1))
    default:
        v28 := GenS2{F0: 1, F1: uint8(v26)}
        _ = v28
    }
    v25 = (float32(v17) + v25)
    v3 += (v1 >> ((462) & 63))
    if false {
        _ = len(v4)
        if (-1 != 128) {
            if (203 < 255) {
                var v29 bool = false
                _ = v29
            } else {
                var v30 float64 = (v17 / 3.14159)
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 20; fi31++ {
                switch (fi31 + fi31) {
                case 32767:
                    for fi32 := 0; fi32 < 2; fi32++ {
                        switch "\n\t" {
                        case "a":
                            v33 := GenS1{F0: 32767}
                            _ = v33
                        case "hello":
                            v1 = (v1 | (v1 & 1))
                        case "":
                            fmt.Println(fi32)
                        default:
                            v4 = append(v4, fi32)
                        }
                        _ = fi32
                    }
                }
                v34 := GenS4{F0: false}
                v26 -= v26
                _ = fi31
                _ = v34
            }
            switch ("a" + "foo bar") {
            case "\n\t":
                v25 = ((0.1 * -2.25) + (3.14159 - v25))
            case "hello":
                fmt.Println(v26, v1, v25)
            case "a":
                if true {
                    v35 := GenS5{F0: 255, F1: -2147483648}
                    _ = v35
                }
            }
        }
        if true {
            if true {
                switch ("foo bar" + "\n\t") {
                case "":
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 uint64 = (v26 | 18446744073709551615)
                        _ = fi36
                        _ = v37
                    }
                case "a":
                    var v38 string = ""
                    _ = v38
                default:
                    var v39 uint8 = 255
                    _ = v39
                }
            }
        } else {
            v40 := GenS3{F0: 120}
            v2 = v2
            _ = v40
        }
    }
    v26 -= (v26 & v26)
    var v41 bool = (v1 > (v3 / ((v3) | 1)))
    _ = len(v4)
    var v42 int32 = 2147483647
    v43 := GenS2{F0: (v42 * (v42 >> (((^(v2))) & 63))), F1: 79}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
