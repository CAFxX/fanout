package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float32 = -2.25
    var v2 int = 32767
    var v3 uint8 = 188
    var v4 int16 = -32768
    var v5 int16 = -32768
    var v6 bool = (35 <= 1)
    switch ("foo bar" + "hello") {
    case "foo bar":
        v1 = v1
    case "":
        for fi7 := 0; fi7 < 12; fi7++ {
            if true {
                switch fi7 {
                case 0:
                    var v8 uint = 18446744073709551615
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 19; fi9++ {
                        switch "" {
                        case "a":
                            var v10 uint32 = 2
                            var v11 int = 2147483647
                            _ = v10
                            _ = v11
                        case "hello":
                            fmt.Println(v4, v2)
                        case "foo bar":
                            if v6 {
                                fmt.Println(v1, v5)
                                v3 = (^(v3))
                                var v12 int64 = -9223372036854775808
                                _ = v12
                            }
                        }
                        v13 := []int{(fi7 ^ v2), fi7, (fi7 - (fi9 << ((1) & 63)))}
                        _ = fi9
                        _ = v13
                    }
                }
                v4 = (-(v5))
                v3 += (v3 + (v3 ^ 255))
            } else {
                v3 = (^(v3))
                var v14 int8 = 42
                _ = v14
            }
            _ = fi7
        }
    }
    v1 += v1
    v15 := []int{}
    for fi16 := 0; fi16 < 8; fi16++ {
        switch ("hello" + "\n\t") {
        case "":
            v17 := GenS1{F0: 2, F1: 2}
            _ = v17
        default:
            v2 = v2
        }
        _ = fi16
    }
    switch "hello" {
    case "a":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    default:
        if (v6 || v6) {
            var v18 float32 = ((v1 / v1) * ((((float32(0.1) - (0.0 * 0.0)) - 3.14159) - (0.1 - (v1 * 3.14159))) / (0.1 / (v1 * 3.14159))))
            var v19 uint16 = 0
            switch -1 {
            case 127:
                v18 = (1.5 * -2.25)
            case 4294967295:
                if (1 >= v5) {
                    var v20 string = "hello"
                    _ = v20
                }
            default:
                fmt.Println(v6, v19)
            }
            _ = v18
            _ = v19
        } else {
            v21 := GenS2{F0: v5}
            _ = v21
        }
    }
    v4 = -32768
    if true {
        var v22 uint8 = 255
        v23 := GenS2{F0: (^(v4))}
        var v24 uint = 348
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 float32 = float32(0.0)
        var v26 int64 = -9223372036854775808
        _ = v25
        _ = v26
    }
    var v27 uint64 = 4294967295
    v28 := GenS1{F0: (v27 >> ((0) & 63)), F1: (v27 << ((65535) & 63))}
    switch "\n\t" {
    case "hello":
        v29 := GenS3{F0: 1}
        _ = v29
    case "a":
        var v30 int8 = 62
        _ = v30
    default:
        v5 -= (v5 + v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
