package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint64 = 42
    var v2 uint64 = (v1 << ((4294967295) & 63))
    var v3 float64 = -2.25
    v4 := []string{("\n\t" + ""), "", "hello"}
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint16 = 42
        var v8 int64 = 905
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3)
    v9 := GenS1{F0: 2}
    var v10 int32 = 2147483647
    v11 := GenS2{F0: (!(127 > 36)), F1: -128}
    fmt.Println(v1, v10, v2)
    if false {
        fmt.Println(v1, v2, v1)
    }
    var v12 int64 = -212
    v4 = append(v4, ("foo bar" + "hello"))
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v3, v2)
    case "foo bar":
        for fi13 := 0; fi13 < 14; fi13++ {
            v2 = v1
            _ = fi13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
