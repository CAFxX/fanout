package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }

func main() {
    var v1 uint16 = 1
    var v2 int16 = 852
    v3 := []int{}
    v4 := []uint32{2147483647, 261}
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    if len(v4) > 0 {
        v4[0] = 65535
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    switch ("hello" + "a") {
    case "":
        for fi5 := 0; fi5 < 6; fi5++ {
            v2 = (v2 ^ v2)
            _ = len(v3)
            _ = fi5
        }
    case "foo bar":
        fmt.Println(v1)
    default:
        if len(v4) > 0 {
            v4[0] = 487
        }
    }
    var v6 uint16 = 32767
    fmt.Println(v1)
    fmt.Println(v2, v1)
    if false {
        v7 := GenS1{F0: 1, F1: (42 <= 176), F2: v6}
        _ = v7
    }
    var v8 uint8 = 1
    var v9 string = ("a" + "foo bar")
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int64 = 1
        v3 = append(v3, (fi10 + (fi10 + fi10)))
        _ = fi10
        _ = v11
    }
    if false {
        var v12 int32 = 42
        _ = v12
    }
    fmt.Println(v8, v2)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
