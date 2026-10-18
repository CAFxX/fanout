package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 uint16) int8 {
    switch "hello" {
    case "foo bar":
        switch "" {
        case "a":
            for fi4 := 0; fi4 < 7; fi4++ {
                v5 := []int{(fi4 % ((-9223372036854775808) | 1))}
                v6 := GenS1{F0: 1, F1: int32(fi4), F2: 2}
                fmt.Println(p0)
                _ = fi4
                _ = v5
                _ = v6
            }
        case "\n\t":
            switch ("foo bar" + "hello") {
            case "\n\t":
                var v7 uint8 = 255
                _ = v7
            default:
                v8 := []int{-785, 2147483647, 2147483647, 42}
                _ = v8
            }
        case "hello":
            v9 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: 255}
            _ = v9
        }
    case "\n\t":
        p0 = p0
    case "hello":
        switch "a" {
        case "a":
            p0 -= p0
        default:
            if true {
                switch -2147483648 {
                case -2147483648:
                    fmt.Println(p1)
                case 0:
                    var v10 bool = false
                    _ = v10
                default:
                    if false {
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v12 := GenS1{F0: 65, F1: 42, F2: 127}
                            v13 := []int{}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        switch -9223372036854775808 {
                        case 42:
                            v14 := []int{}
                            _ = v14
                        case -1:
                            fmt.Println(p1, p0)
                        default:
                            for fi15 := 0; fi15 < 14; fi15++ {
                                v16 := []int{127}
                                switch "" {
                                case "\n\t":
                                    var v17 float64 = -2.25
                                    _ = v17
                                }
                                _ = len(v16)
                                _ = fi15
                                _ = v16
                            }
                        }
                        var v18 float64 = 1e10
                        _ = v18
                    }
                }
                var v19 uint64 = 699
                _ = v19
            }
        }
    default:
        p0 = (p0 / ((2) | 1))
    }
    p0 = (^(p1))
    _ = p0
    _ = p1
    return 6
}
func genFn1() uint32 {
    if (1 > 864) {
        var v20 float32 = float32(1e10)
        v21 := []int{1, 42, 32768, 128}
        v22 := GenS1{F0: 4294967295, F1: 1, F2: 128}
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 uint16 = 252
    v23 += 42
    v24 := GenS2{F0: false, F1: "hello"}
    _ = v23
    _ = v24
    return 42
}

func main() {
    var v1 uint = 0
    var v2 uint64 = 4294967295
    var v3 uint16 = 586
    v25 := GenS1{F0: v2, F1: -776, F2: 128}
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(fi26)
        for fi27 := 0; fi27 < 5; fi27++ {
            if (!true) {
                v3 -= v3
            }
            var v28 string = ("a" + (("" + ("foo bar" + "")) + "hello"))
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 int16 = 0
    v3 = 255
    v30 := GenS3{F0: "\n\t"}
    v3 -= (v3 ^ 871)
    switch -1 {
    case -1:
        var v31 int16 = (v29 * v29)
        _ = v31
    case 32768:
        var v32 int = 32768
        _ = v32
    case -336:
        var v33 int32 = -2147483648
        _ = v33
    default:
        for fi34 := 0; fi34 < 1; fi34++ {
            fmt.Println(v2, fi34)
            _ = fi34
        }
    }
    v2 -= 0
    if (!false) {
        fmt.Println(v29, v2)
    } else {
        for fi35 := 0; fi35 < 13; fi35++ {
            v3 -= v3
            var v36 float64 = float64(float32(float64(0.1)))
            _ = fi35
            _ = v36
        }
    }
    var v37 uint = v1
    switch "hello" {
    case "hello":
        switch 0 {
        case -1:
            v1 = (v1 / (((v37 - (v1 / (((v37 - 1)) | 1)))) | 1))
        default:
            fmt.Println(v3, v29)
        }
    case "foo bar":
        var v38 string = ("hello" + ("a" + ""))
        _ = v38
    default:
        var v39 uint32 = 2147483647
        _ = v39
    }
    var v40 int32 = -2147483648
    v41 := []int{2147483647, 127, -449, 0}
    fmt.Println(v2, v29, v1)
    var v42 uint32 = 2
    var v43 uint16 = v3
    for fi44 := 0; fi44 < 19; fi44++ {
        for fi45 := 0; fi45 < 14; fi45++ {
            v29 -= (v29 + 42)
            if true {
                v3 -= (v43 * 0)
            }
            _ = fi45
        }
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
        for fi46 := 0; fi46 < 4; fi46++ {
            fmt.Println(v1, v37)
            _ = fi46
        }
        _ = fi44
    }
    var v47 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v47
}
