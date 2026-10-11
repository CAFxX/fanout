package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 bool }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int; F1 uint8; F2 int }
func genFn0() uint {
    var v4 int16 = 400
    switch "" {
    case "":
        var v5 int16 = v4
        _ = v5
    case "a":
        v6 := GenS1{F0: 0, F1: 2147483647, F2: 65535}
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: -2147483648, F2: 65535}
    var v8 uint = 4294967295
    v4 += (v4 / (((v4 & v4)) | 1))
    _ = v4
    _ = v7
    _ = v8
    return (v8 % ((v8) | 1))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 469
    var v3 float64 = -2.25
    switch "\n\t" {
    case "":
        v9 := GenS2{F0: -2147483648}
        _ = v9
    case "foo bar":
        var v10 float64 = ((((1.5 * v3) + v3) + (v3 - 3.14159)) - v3)
        _ = v10
    default:
        fmt.Println(v3, v2)
    }
    var v11 float32 = 1.5
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 float32 = (((1.5 - (v11 + 0.0)) - 3.14159) * (v11 / (v11 * v11)))
        fmt.Println(fi12, v13)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(v11, v2, v11)
            v16 := GenS1{F0: 65535, F1: -299, F2: 2}
            v2 -= int16(fi14)
            _ = fi15
            _ = v16
        }
        v2 = (^(v2))
        v3 += float64(((((-2.25 - 3.14159) / (((v11 / -2.25) * v11) / 3.14159)) * 1.5) / v11))
        _ = fi14
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 float64 = (v3 - v3)
        var v19 uint = genFn0()
        v3 = (3.14159 + (v3 - 3.14159))
        _ = fi17
        _ = v18
        _ = v19
    }
    switch ("" + "a") {
    case "\n\t":
        if true {
            if (false && true) {
                v20 := GenS1{F0: 1, F1: -1, F2: 462}
                var v21 uint32 = 42
                var v22 uint8 = 128
                _ = v20
                _ = v21
                _ = v22
            } else {
                var v23 float64 = ((v3 * (3.14159 + (float64(((1.5 * 0.0) - float32(1.5))) - 0.0))) - 1.5)
                _ = v23
            }
            if ((1 >= 18446744073709551615) || (42 > 2)) {
                v3 = v3
                switch 127 {
                case 127:
                    v24 := GenS3{F0: 99, F1: 4294967295, F2: true}
                    _ = v24
                case 0:
                    v2 = -1
                case -775:
                    fmt.Println(v11, v3)
                default:
                    var v25 bool = (128 <= 55)
                    _ = v25
                }
                var v26 float32 = v11
                _ = v26
            } else {
                var v27 uint64 = 4294967295
                v28 := GenS3{F0: 9223372036854775807, F1: genFn0(), F2: false}
                _ = v27
                _ = v28
            }
            switch 2147483647 {
            case 597:
                if (true && (!(-74 != -96))) {
                    if false {
                        var v29 uint64 = 2
                        v2 += v2
                        _ = v29
                    }
                    v30 := GenS1{F0: 2147483647, F1: 42, F2: 176}
                    _ = v30
                } else {
                    if (v3 != v3) {
                        var v31 int64 = -98
                        if true {
                            v2 = 0
                        }
                        _ = v31
                    }
                    switch "" {
                    case "\n\t":
                        delete(v1, 566)
                    }
                    v32 := GenS3{F0: 1, F1: 133, F2: true}
                    _ = v32
                }
            case -2147483648:
                var v33 bool = false
                _ = v33
            }
        }
    case "hello":
        var v34 bool = (42 < 42)
        _ = v34
    }
    v11 = ((v11 * v11) / v11)
    v2 = (v2 % ((1) | 1))
    if (!(!(((-2.25 - -2.25) - v3) > (v3 + (3.14159 + (1e10 * float64(0.0))))))) {
        v2 += v2
        v35 := GenS3{F0: 42, F1: 65535, F2: false}
        v36 := GenS4{F0: v2, F1: 2}
        _ = v35
        _ = v36
    } else {
        v2 += v2
    }
    v37 := GenS5{F0: -9223372036854775808, F1: 42, F2: 51}
    switch ("hello" + "foo bar") {
    case "\n\t":
        fmt.Println(v2, v3)
    case "":
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 int16 = 1
            _ = fi38
            _ = v39
        }
    case "hello":
        v11 = 3.14159
    default:
        var v40 int16 = v2
        _ = v40
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
}
