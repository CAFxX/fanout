package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS6 struct { F0 bool }

func main() {
    var v1 int16 = 32767
    var v2 uint = 1
    var v3 uint64 = 2
    var v4 int16 = 1000
    v5 := GenS1{F0: 25, F1: -9223372036854775808, F2: 42}
    v6 := GenS2{F0: -2147483648, F1: v4}
    if true {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{(fi7 & (fi7 % (((-(fi7))) | 1)))}
            v9 := GenS2{F0: 1, F1: v1}
            _ = fi7
            _ = v8
            _ = v9
        }
        v1 = (v4 * 1000)
        v10 := GenS3{F0: -1, F1: v2}
        _ = v10
    }
    fmt.Println(v1, v2, v4)
    if false {
        fmt.Println(v4, v3)
        fmt.Println(v4)
    }
    v11 := []int{4294967295, 630}
    v3 = (v3 % (((v3 | 18446744073709551615)) | 1))
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if (true || false) {
        if false {
            var v12 int8 = -5
            _ = v12
        } else {
            if (v1 != 628) {
                var v13 int = 127
                var v14 uint8 = 127
                _ = v13
                _ = v14
            } else {
                v2 -= 4294967295
            }
            v1 = 42
        }
        v1 = v1
    }
    v15 := GenS2{F0: 2147483647, F1: v1}
    switch "" {
    case "":
        if (!true) {
            v16 := GenS4{F0: 65535}
            _ = len(v11)
            if false {
                fmt.Println(v3, v4, v4)
            } else {
                if len(v11) > 0 {
                    v11[0] = -909
                }
                if ("a" >= "hello") {
                    var v17 bool = (true || true)
                    v18 := GenS5{F0: uint64(v2), F1: 337, F2: uint8(v3)}
                    v19 := GenS6{F0: (9223372036854775807 >= -705)}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    var v20 int = 128
                    var v21 uint64 = v3
                    fmt.Println(v1)
                    _ = v20
                    _ = v21
                }
            }
            _ = v16
        }
    case "a":
        var v22 int8 = 100
        _ = v22
    default:
        v1 = v4
    }
    fmt.Println(v1)
    switch ("\n\t" + "") {
    case "foo bar":
        v23 := GenS3{F0: (v1 << ((255) & 63)), F1: v2}
        _ = v23
    case "hello":
        v1 -= 1000
    case "a":
        switch 9223372036854775807 {
        case 32767:
            switch "a" {
            case "foo bar":
                if false {
                    fmt.Println(v2)
                } else {
                    v1 = v1
                    v4 += v4
                }
            }
        case 42:
            if false {
                var v24 uint64 = v3
                if len(v11) > 0 {
                    v11[0] = 764
                }
                var v25 string = "hello"
                _ = v24
                _ = v25
            } else {
                v26 := GenS2{F0: 62, F1: v1}
                var v27 float32 = 0.0
                _ = v26
                _ = v27
            }
        case -9223372036854775808:
            switch 346 {
            case 1:
                v1 += int16(v3)
            case -9223372036854775808:
                v1 += v1
            }
        default:
            var v28 float64 = float64((1.5 - 1.5))
            _ = v28
        }
    default:
        var v29 uint = (^(v2))
        _ = v29
    }
    v11 = append(v11, -576)
    if (-2.25 <= 0.0) {
        var v30 int32 = -531
        v3 = 997
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
