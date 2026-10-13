package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }

func main() {
    var v1 int64 = -129
    var v2 uint8 = 127
    var v3 float32 = 0.0
    var v4 int8 = -128
    var v5 int32 = -2147483648
    v6 := []uint32{2, 588}
    v7 := GenS1{F0: ("\n\t" + ("" + ("\n\t" + "a"))), F1: 873}
    if len(v6) > 0 {
        v6[0] = 42
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        v6 = append(v6, 0)
        fmt.Println(v2)
        _ = fi8
    }
    switch "a" {
    case "foo bar":
        v9 := GenS1{F0: ("foo bar" + (("\n\t" + ("hello" + "hello")) + "")), F1: 954}
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int64 = 9223372036854775807
            _ = fi10
            _ = v11
        }
    case "\n\t":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        v1 -= (v1 | v1)
    }
    var v12 float64 = -2.25
    v13 := GenS1{F0: "a", F1: 643}
    if len(v6) > 0 {
        v6[0] = 1
    }
    v3 = ((((v3 - 1e10) - 0.0) * v3) - (1.5 / 3.14159))
    switch "a" {
    case "":
        v12 = (v12 - 0.1)
    default:
        v14 := GenS1{F0: "a", F1: 18446744073709551615}
        _ = v14
    }
    v15 := GenS1{F0: "", F1: 416}
    var v16 uint8 = v2
    v6 = append(v6, 1)
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 20; fi18++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v19 int16 = -32768
        _ = fi18
        _ = v19
    }
    v2 = v2
    switch "hello" {
    case "foo bar":
        var v20 float32 = (v3 * v3)
        _ = v20
    default:
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
