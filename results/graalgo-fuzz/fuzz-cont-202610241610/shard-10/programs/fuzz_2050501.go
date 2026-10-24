package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }
type GenS5 struct { F0 int }

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{2147483647, 2147483647, 127, 2147483647}
    var v3 int32 = 0
    var v4 uint64 = 18446744073709551615
    var v5 float64 = (float64((v1 + float32(0.1))) + (1e10 * 1.5))
    var v6 int16 = -824
    v3 += (v3 ^ (v3 << ((1) & 63)))
    v7 := GenS1{F0: 0}
    switch "" {
    case "hello":
        v3 = (-(v3))
    case "foo bar":
        var v8 uint64 = v4
        _ = v8
    default:
        v9 := GenS1{F0: 550}
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS2{F0: true, F1: 128, F2: uint32(v4)}
        v12 := GenS3{F0: uint8(v4), F1: 1, F2: 2}
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 += 42
    switch -1 {
    case -2147483648:
        v13 := GenS3{F0: 0, F1: -1, F2: 42}
        _ = v13
    default:
        fmt.Println(v1, v6, v1)
    }
    if ("foo bar" != "") {
        v14 := GenS1{F0: (v6 / ((v6) | 1))}
        v15 := GenS4{F0: 42, F1: 654, F2: 1}
        fmt.Println(v3)
        _ = v14
        _ = v15
    } else {
        var v16 uint8 = 127
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint8 = 145
            _ = fi17
            _ = v18
        }
        switch 42 {
        case 9223372036854775807:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v4 = (v4 - v4)
        }
    } else {
        if (!true) {
            fmt.Println(v4, v6)
            var v19 string = (("" + "") + "")
            var v20 float32 = v1
            _ = v19
            _ = v20
        } else {
            v2 = append(v2, 65535)
            switch 1 {
            case 620:
                for fi21 := 0; fi21 < 19; fi21++ {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        if false {
                            fmt.Println(v1)
                        }
                        var v23 uint64 = (v4 * v4)
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v4, v4, v4)
                    v4 -= 4294967295
                    _ = fi21
                }
            }
        }
        if ("foo bar" >= "foo bar") {
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS2{F0: (v1 > ((3.14159 + (v1 / (0.0 - (((v1 * (v1 / 3.14159)) * -2.25) - v1)))) + -2.25)), F1: 2, F2: 1}
                fmt.Println(v3)
                for fi26 := 0; fi26 < 9; fi26++ {
                    v4 -= v4
                    _ = fi26
                }
                _ = fi24
                _ = v25
            }
        }
    }
    var v27 float64 = v5
    if true {
        v3 -= (v3 + (v3 & 804))
        var v28 int16 = (v6 + -813)
        v5 = (1e10 + 1e10)
        _ = v28
    } else {
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int8 = 34
            for fi31 := 0; fi31 < 16; fi31++ {
                v2 = append(v2, 4294967295)
                v6 -= v6
                var v32 uint = 0
                var v33 int32 = 950
                v34 := GenS5{F0: (fi29 >> ((42) & 63))}
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi29
            _ = v30
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch 32768 {
        case 128:
            v6 += 923
        case 32768:
            v35 := GenS5{F0: -179}
            _ = v35
        case -9223372036854775808:
            v2 = append(v2, 1)
        default:
            var v36 bool = true
            _ = v36
        }
    }
    v2 = append(v2, 767)
    var v37 int32 = v3
    v27 += v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
