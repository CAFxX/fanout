package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := make(map[int]bool)
    v4 := []string{(("\n\t" + "a") + ("\n\t" + "foo bar")), ("foo bar" + "a"), ((("hello" + "hello") + ("a" + ("hello" + "foo bar"))) + "hello")}
    v5 := GenS1{F0: 351}
    v6 := v1[-87]
    v7 := GenS2{F0: 255, F1: 828}
    fmt.Println(v6)
    _ = len(v4)
    switch (((((("hello" + "") + "\n\t") + "\n\t") + ("a" + "\n\t")) + "hello") + "hello") {
    case "\n\t":
        fmt.Println(v6)
    case "":
        v6 = v6
    case "foo bar":
        fmt.Println(v6)
    default:
        v2 = append(v2, -805)
    }
    var v8 int8 = 92
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: 600}
            v12 := GenS1{F0: 32767}
            _ = fi10
            _ = v11
            _ = v12
        }
        v6 = v6
        var v13 uint16 = 255
        _ = fi9
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v6 = v6
            v16 := GenS1{F0: 32767}
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            _ = fi15
            _ = v16
        }
        fmt.Println(v8)
        _ = fi14
    }
    switch "hello" {
    case "hello":
        v17 := GenS2{F0: 1, F1: 128}
        _ = v17
    case "a":
        if v6 {
            fmt.Println(v6)
            if true {
                v8 += (v8 & v8)
            }
            fmt.Println(v8, v6)
        } else {
            v18 := GenS2{F0: 681, F1: 185}
            v6 = (0 != 227)
            v8 += (v8 / ((v8) | 1))
            _ = v18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
