package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint = 65535
    var v2 int16 = 1
    v3 := make(map[uint32]int)
    v4 := []uint32{2147483648}
    var v5 float64 = ((3.14159 / 3.14159) / ((1e10 / (3.14159 / 3.14159)) - (float64(((1.5 / 1.0) / (1e10 - float32(0.1)))) + 0.1)))
    var v6 bool = (!false)
    if (!v6) {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int = fi7
            _ = fi7
            _ = v8
        }
    }
    if v6 {
        v2 = (v2 >> ((v1) & 63))
        v9 := GenS1{F0: 100, F1: 286}
        var v10 bool = (((v2 >> ((v1) & 63)) < (v2 | (v2 ^ 32767))) && true)
        _ = v10
        _ = v9
    }
    fmt.Println(v6, v6, v1)
    switch "a" {
    case "hello":
        v11 := GenS1{F0: -79, F1: 65535}
        _ = v11
    default:
        v2 -= (v2 & 42)
    }
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int8 = 127
        v2 = -32768
        _ = len(v3)
        _ = fi13
        _ = v14
    }
    switch 1 {
    case -543:
        var v15 float64 = float64(1e10)
        _ = v15
    }
    var v16 float64 = 3.14159
    if v6 {
        if (-1 != 0) {
            var v17 float64 = (v16 / v16)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v17
        }
        fmt.Println(v6, v2, v12)
    }
    var v18 int8 = 100
    var v19 int16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
