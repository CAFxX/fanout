package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint8 = 147
    v2 := []uint32{742, 340}
    var v3 string = ("hello" + "\n\t")
    fmt.Println(v3, v3)
    v1 = v1
    var v4 uint = 42
    v5 := GenS1{F0: v1, F1: -9223372036854775808, F2: 2147483647}
    v3 = v3
    if false {
        v6 := GenS2{F0: v1}
        if (!false) {
            v1 = 2
        } else {
            v3 = v3
        }
        v3 = (v3 + v3)
        _ = v6
    } else {
        var v7 int16 = 1
        var v8 string = "hello"
        _ = v7
        _ = v8
    }
    var v9 bool = false
    var v10 float64 = 3.14159
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 string = ""
        _ = fi11
        _ = v12
    }
    v1 = (v1 - v1)
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint64 = 572
        v10 = ((1e10 + v10) * 0.1)
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 1
    v3 = ""
    switch ((("\n\t" + v3) + (v3 + v3)) + ("hello" + ((v3 + v3) + (((v3 + v3) + v3) + v3)))) {
    case "a":
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v9, v15, fi16)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi16
        }
    case "foo bar":
        switch 127 {
        case 0:
            v10 = -2.25
        case 128:
            fmt.Println(v10, v4, v15)
        }
    default:
        var v17 uint16 = uint16(v1)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
