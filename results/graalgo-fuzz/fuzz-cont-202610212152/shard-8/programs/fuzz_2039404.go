package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 bool = (!(false || (0 > 42)))
    v2 := []int{652, 42, 0, 42}
    var v3 uint32 = 2147483648
    for fi4 := 0; fi4 < 18; fi4++ {
        switch "foo bar" {
        case "hello":
            var v5 uint8 = 1
            _ = v5
        case "\n\t":
            fmt.Println(v1)
        case "a":
            var v6 uint = 92
            _ = v6
        }
        var v7 uint64 = 42
        _ = fi4
        _ = v7
    }
    if (!v1) {
        v1 = (65535 != 962)
    } else {
        var v8 int32 = 865
        if (967 >= 4294967295) {
            var v9 int8 = 1
            v10 := GenS1{F0: 32767, F1: -2147483648}
            _ = v10
            _ = v9
        }
        var v11 uint8 = 0
        _ = v11
        _ = v8
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            v13 := GenS2{F0: (^(v3)), F1: 128, F2: 42}
            _ = v13
        }
        _ = fi12
    }
    fmt.Println(v1)
    v3 = v3
    var v14 int = 534
    switch "hello" {
    case "hello":
        switch "foo bar" {
        case "foo bar":
            v14 += (v14 & 2147483647)
        case "":
            var v15 int = (-(v14))
            _ = v15
        default:
            if (0 >= 1) {
                fmt.Println(v14, v3)
            } else {
                fmt.Println(v1, v1)
            }
        }
    default:
        var v16 uint8 = 128
        _ = v16
    }
    var v17 uint32 = (^(v3))
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            switch "a" {
            case "\n\t":
                v20 := GenS3{F0: "a"}
                _ = v20
            case "hello":
                var v21 float64 = 0.0
                _ = len(v2)
                var v22 uint8 = 206
                _ = v21
                _ = v22
            case "":
                if len(v2) > 0 {
                    v2[0] = -1
                }
            default:
                if (false && false) {
                    switch "" {
                    case "":
                        var v23 uint16 = 65535
                        _ = v23
                    case "foo bar":
                        fmt.Println(v17, v1, v14)
                    default:
                        v14 += (fi18 % ((v14) | 1))
                    }
                    if true {
                        if v1 {
                            v24 := GenS4{F0: 42, F1: 95}
                            var v25 int = v14
                            _ = v24
                            _ = v25
                        }
                        var v26 bool = (true && (18446744073709551615 > 18446744073709551615))
                        if (v26 || ((3.14159 / 3.14159) < 0.1)) {
                            var v27 float32 = 1.5
                            v1 = v26
                            if len(v2) > 0 {
                                v2[0] = -2147483648
                            }
                            _ = v27
                        }
                        _ = v26
                    } else {
                        v14 = -9223372036854775808
                        var v28 uint8 = 189
                        v29 := GenS2{F0: v3, F1: v28, F2: 0}
                        _ = v28
                        _ = v29
                    }
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                }
            }
            _ = fi19
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            v14 = fi31
            v2 = append(v2, int(v3))
            _ = fi31
        }
        _ = fi18
    }
    var v32 uint16 = 1
    var v33 int64 = -727
    for fi34 := 0; fi34 < 5; fi34++ {
        switch v14 {
        case -903:
            fmt.Println(v3, fi34)
        case 9223372036854775807:
            v32 = v32
        }
        v14 -= fi34
        _ = fi34
    }
    v32 -= v32
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
