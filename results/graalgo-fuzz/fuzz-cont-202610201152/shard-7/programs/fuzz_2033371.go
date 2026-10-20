package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{-1}
    var v3 int16 = 1
    v4 := make(map[uint32]string)
    var v5 int = 127
    for fi6 := 0; fi6 < 5; fi6++ {
        v3 = (v3 % ((-1) | 1))
        v7 := GenS1{F0: 0, F1: -1}
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, v5)
    var v8 uint32 = 749
    v2 = append(v2, v5)
    var v9 int8 = 100
    switch v5 {
    case 42:
        var v10 uint64 = 208
        _ = v10
    case 2147483647:
        var v11 int16 = (v3 >> ((65535) & 63))
        _ = v11
    case 95:
        v3 += v3
    default:
        v5 += (^(v5))
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint8 = 2
        switch "" {
        case "\n\t":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int8 = 42
                if false {
                    if (18446744073709551615 <= 18446744073709551615) {
                        v4[v8] = ("" + "a")
                        v13 = (^(v13))
                    } else {
                        v2 = append(v2, (v5 + -346))
                        switch ("hello" + "a") {
                        case "\n\t":
                            var v16 string = "a"
                            _ = v16
                        case "hello":
                            switch "hello" {
                            case "hello":
                                var v17 uint64 = 18446744073709551615
                                _ = v17
                            }
                            var v18 uint64 = 18446744073709551615
                            if false {
                                fmt.Println(v3, fi12)
                                var v19 int = fi12
                                var v20 uint = 42
                                _ = v19
                                _ = v20
                            } else {
                                v21 := GenS1{F0: (^(v8)), F1: 0}
                                v18 += v18
                                _ = v21
                            }
                            _ = v18
                        default:
                            fmt.Println(v5, v8)
                        }
                        var v22 float32 = 1e10
                        _ = v22
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi12 | fi14)
                    }
                }
                _ = fi14
                _ = v15
            }
        case "a":
            if ("foo bar" <= ("a" + ("foo bar" + "a"))) {
                v23 := GenS2{F0: 32767, F1: (v13 == (v13 * 2))}
                _ = v23
            } else {
                v24 := GenS1{F0: v8, F1: 42}
                v25 := GenS2{F0: int16(v9), F1: false}
                _ = v24
                _ = v25
            }
        }
        fmt.Println(v8, v8, v9)
        _ = fi12
        _ = v13
    }
    var v26 int16 = v3
    if (false && false) {
        var v27 int64 = 9223372036854775807
        if ((2 >= 4294967295) && (v8 < 2)) {
            v28 := GenS2{F0: (v26 ^ 1000), F1: (false && ("foo bar" <= "\n\t"))}
            var v29 float32 = float32(3.14159)
            _ = v28
            _ = v29
        }
        fmt.Println(v27, v8, v5)
        _ = v27
    } else {
        fmt.Println(v9, v9, v26)
        v26 = 32767
        v3 -= (v3 >> ((0) & 63))
    }
    var v30 int8 = v9
    var v31 uint8 = 115
    var v32 uint16 = 255
    fmt.Println(v30, v32)
    var v33 int16 = (v3 >> ((673) & 63))
    v31 = 128
    switch 1 {
    case 32768:
        delete(v1, (v5 & (v5 & int(v31))))
    }
    if (v3 < int16(v32)) {
        v32 -= uint16(v8)
    }
    v34 := GenS3{F0: 202, F1: (v8 >> ((42) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
