package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 bool; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 int32 = -42
    var v5 int32 = 601
    var v6 uint8 = 128
    v6 += 127
    if ((true || false) && ("" <= (("foo bar" + (("a" + "hello") + "foo bar")) + "a"))) {
        v5 = 1
        if false {
            v5 = v5
            if ((-58 != -1) && (22 <= 0)) {
                var v7 uint64 = 18446744073709551615
                _ = v7
            } else {
                var v8 uint16 = 65535
                if false {
                    switch 32768 {
                    case 127:
                        v9 := GenS1{F0: (!(!false)), F1: 128, F2: 822}
                        _ = v9
                    case -1:
                        var v10 float32 = float32(0.1)
                        _ = v10
                    default:
                        var v11 int32 = -2147483648
                        _ = v11
                    }
                } else {
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 int8 = -128
                        var v14 uint16 = v8
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(v5, v4)
                    var v15 float32 = 1e10
                    _ = v15
                }
                _ = v8
            }
        }
        fmt.Println(v6)
    } else {
        var v16 uint = 18446744073709551615
        _ = v16
    }
    var v17 int = 42
    v4 -= (v5 >> ((4294967295) & 63))
    v18 := GenS1{F0: (0.0 < ((1.5 / ((1e10 - (3.14159 - (0.0 * (((-2.25 * 1e10) / 0.1) - 1e10)))) + (float32((((1e10 / 1.5) / 0.1) * 0.0)) - 3.14159))) / 1.0)), F1: v17, F2: 308}
    var v19 uint32 = 42
    v6 = v6
    if true {
        if false {
            v5 -= -393
            if false {
                var v20 int = (v17 - v17)
                v19 -= uint32(v6)
                _ = v20
            }
        }
    }
    switch "a" {
    case "":
        var v21 float32 = 3.14159
        _ = v21
    default:
        switch 220 {
        case 32767:
            v22 := GenS2{F0: "", F1: 2}
            _ = v22
        case 4294967295:
            fmt.Println(v5)
        case 9223372036854775807:
            var v23 uint16 = 177
            _ = v23
        default:
            if true {
                fmt.Println(v6)
                fmt.Println(v4)
            } else {
                var v24 int32 = (^(v5))
                _ = v24
            }
        }
    }
    var v25 uint32 = (v19 | v19)
    switch "a" {
    case "\n\t":
        v17 = (v17 + -2147483648)
    case "a":
        var v26 int32 = -2147483648
        _ = v26
    }
    v27 := GenS3{F0: (104 != 1), F1: ("" + "\n\t"), F2: (^(v6))}
    var v28 string = "hello"
    var v29 uint8 = (v6 % ((v6) | 1))
    v30 := GenS4{F0: (v29 - v6), F1: 127}
    var v31 int32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
