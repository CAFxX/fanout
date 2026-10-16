package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }

func main() {
    var v1 int32 = 86
    v2 := make(map[string]bool)
    var v3 string = ""
    fmt.Println(v1, v1)
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 9223372036854775807, F1: (fi4 | 42), F2: "hello"}
        _ = fi4
        _ = v5
    }
    fmt.Println(v3)
    v6 := GenS2{F0: 1}
    fmt.Println(v3, v1)
    v7 := []int{-598, -57, 271, 665}
    var v8 int64 = -664
    if (0 != 4294967295) {
        v9 := v2[v3]
        var v10 float64 = -2.25
        v3 = "hello"
        _ = v10
        _ = v9
    } else {
        switch "foo bar" {
        case "a":
            if false {
                var v11 int = -9223372036854775808
                switch v11 {
                case 1:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                case 0:
                    var v12 int16 = 32767
                    _ = v12
                case -1:
                    var v13 string = v3
                    _ = v13
                default:
                    v14 := GenS3{F0: "", F1: 105, F2: 963}
                    _ = v14
                }
                _ = v11
            }
        case "foo bar":
            if (v8 == (v8 | (v8 + v8))) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 int16 = 1000
                    if (18446744073709551615 > 1) {
                        v17 := GenS2{F0: 0}
                        _ = v17
                    } else {
                        fmt.Println(v1, v16)
                        fmt.Println(v16, v16)
                    }
                    v16 = 125
                    _ = fi15
                    _ = v16
                }
            }
        case "\n\t":
            v18 := GenS3{F0: ("" + v3), F1: 127, F2: 567}
            _ = v18
        }
        delete(v2, v3)
    }
    var v19 uint32 = 42
    if (false && true) {
        var v20 uint64 = 544
        var v21 float64 = (0.0 + 0.1)
        switch 572 {
        case -452:
            if true {
                v19 -= (^(v19))
                _ = len(v2)
            }
        case 9223372036854775807:
            v20 -= (v20 ^ (v20 / ((v20) | 1)))
        case 2147483647:
            var v22 float64 = 1.5
            _ = v22
        }
        _ = v20
        _ = v21
    }
    v7 = append(v7, 4294967295)
    var v23 uint16 = 42
    v23 += 171
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
