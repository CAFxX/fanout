package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint = 1
    var v2 int = 1
    var v3 bool = true
    v4 := []int{}
    switch "\n\t" {
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        if (float32(3.14159) != 1e10) {
            fmt.Println(v3, v2)
        } else {
            v5 := GenS1{F0: -9223372036854775808, F1: 42, F2: 212}
            _ = v5
        }
    default:
        if len(v4) > 0 {
            v4[0] = v2
        }
    }
    var v6 uint = (v1 / ((v1) | 1))
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 float64 = 0.0
        v9 := GenS2{F0: 42, F1: 67}
        if len(v4) > 0 {
            v4[0] = fi7
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: ""}
    var v11 int8 = 42
    var v12 uint8 = 7
    for fi13 := 0; fi13 < 15; fi13++ {
        switch fi13 {
        case -1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            switch v2 {
            case 32768:
                v3 = v3
            default:
                v4 = append(v4, 127)
            }
        }
        _ = fi13
    }
    v3 = v3
    switch v2 {
    case 32768:
        var v14 uint8 = (v12 | (^(v12)))
        _ = v14
    }
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    var v15 uint32 = 904
    v16 := GenS4{F0: (18446744073709551615 != (^(v6)))}
    var v17 uint16 = 329
    fmt.Println(v1, v1, v15)
    if len(v4) > 0 {
        v4[0] = v2
    }
    switch ("foo bar" + "hello") {
    case "":
        var v18 int16 = 1000
        _ = v18
    case "a":
        switch ("" + "") {
        case "hello":
            var v19 uint32 = v15
            _ = v19
        case "\n\t":
            v15 -= v15
        default:
            var v20 string = ("foo bar" + "")
            _ = v20
        }
    }
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
