package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []string{"hello", "hello", "\n\t"}
    var v2 int16 = 32767
    var v3 int16 = v2
    var v4 uint = 18446744073709551615
    var v5 bool = true
    v3 += (v3 ^ (^(v2)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "" {
    case "foo bar":
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 0, F1: 0, F2: 555}
            _ = fi6
            _ = v7
        }
    case "\n\t":
        var v8 string = "\n\t"
        _ = v8
    case "":
        v2 -= -32768
    }
    v1 = append(v1, "\n\t")
    v9 := GenS2{F0: v4}
    v5 = (128 == 32768)
    v5 = (0 < 4294967295)
    var v10 uint = 2
    if (((-2.25 * (1e10 * 3.14159)) * 1e10) < 0.1) {
        fmt.Println(v4, v3, v10)
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v5 = v5
        v5 = false
        var v12 uint8 = 1
        _ = fi11
        _ = v12
    }
    v3 = v3
    var v13 bool = v5
    for fi14 := 0; fi14 < 4; fi14++ {
        switch "hello" {
        case "":
            v15 := GenS1{F0: 2, F1: 1, F2: -2147483648}
            _ = v15
        case "hello":
            v16 := GenS3{F0: 635}
            _ = v16
        default:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        _ = fi14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = (65535 != v4)
    v17 := GenS2{F0: 1}
    v10 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
