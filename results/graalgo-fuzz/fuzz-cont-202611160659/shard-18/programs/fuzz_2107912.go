package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = 1000
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 uint32 = 2
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: -2147483648, F1: 510, F2: ""}
        var v7 uint16 = 32767
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: fi8, F1: -32768, F2: "\n\t"}
            _ = fi8
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "a":
        v12 := GenS1{F0: -9223372036854775808, F1: (v1 << ((255) & 63)), F2: ("hello" + "a")}
        _ = v12
    case "":
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(v2)
                fmt.Println(fi14, fi14)
                _ = fi14
            }
            _ = fi13
        }
    }
    var v15 float64 = (0.0 + (3.14159 * (float64(-2.25) + 3.14159)))
    fmt.Println(v15, v15)
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS1{F0: (^(fi16)), F1: 546, F2: ((((("" + ("hello" + "")) + ("foo bar" + "")) + "\n\t") + "a") + "")}
        v4 -= v4
        var v18 int64 = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint8 = 40
    for fi20 := 0; fi20 < 13; fi20++ {
        v19 = 252
        _ = fi20
    }
    var v21 int = int(v19)
    v2 = v2
    v22 := GenS1{F0: v21, F1: (^(v1)), F2: "a"}
    var v23 bool = (155 <= (v19 | v19))
    v19 = v19
    if v23 {
        var v24 int64 = 9223372036854775807
        fmt.Println(v21, v4, v4)
        v25 := GenS2{F0: v23}
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: (!true)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
