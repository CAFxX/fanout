package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int64; F2 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 666
    var v3 uint16 = 65535
    var v4 float32 = -2.25
    if true {
        fmt.Println(v4, v2, v4)
    }
    v5 := []int{32768, -2147483648, 2147483647, 0}
    if (42 == 224) {
        fmt.Println(v4, v1, v2)
        for fi6 := 0; fi6 < 4; fi6++ {
            v2 += (v2 % ((v2) | 1))
            _ = fi6
        }
        var v7 string = ""
        _ = v7
    }
    fmt.Println(v4, v3)
    if true {
        v4 += (((1e10 - (3.14159 + 1.5)) - v4) - 1e10)
        switch "a" {
        case "foo bar":
            fmt.Println(v3, v2)
        case "hello":
            if true {
                v3 = v3
                v4 += v4
            } else {
                v5 = append(v5, 0)
            }
        case "\n\t":
            v8 := GenS1{F0: "a"}
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: ""}
            _ = fi9
            _ = v10
        }
    }
    var v11 float64 = float64(0.1)
    switch 127 {
    case 128:
        if false {
            var v12 int = -997
            _ = v12
        }
    case -1:
        v13 := GenS1{F0: ""}
        _ = v13
    case 127:
        v3 -= 65535
    default:
        v5 = append(v5, 4294967295)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v4 += ((float32(float64(v4)) + (3.14159 - (v4 + float32(v11)))) + v4)
        _ = fi14
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = v2
    fmt.Println(v2, v3, v3)
    v15 := GenS2{F0: 42, F1: (v1 | v1), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
