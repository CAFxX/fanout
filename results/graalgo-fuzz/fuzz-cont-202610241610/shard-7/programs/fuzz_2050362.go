package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint16 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{}
    var v3 int8 = -47
    var v4 int = 32768
    var v5 uint = 1
    var v6 uint16 = 887
    var v7 uint8 = 91
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: -268, F1: "\n\t"}
        _ = fi8
        _ = v9
    }
    fmt.Println(v5, v6)
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 0, F1: ("foo bar" + "\n\t")}
        _ = fi10
        _ = v11
    }
    fmt.Println(v7)
    var v12 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int = (^(v4))
        switch "\n\t" {
        case "\n\t":
            v15 := GenS1{F0: v12, F1: (("hello" + "hello") + "hello")}
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v1, v12, v6)
    switch "hello" {
    case "":
        v16 := GenS1{F0: v12, F1: ("foo bar" + "")}
        _ = v16
    case "foo bar":
        var v17 float32 = 1e10
        _ = v17
    }
    switch "a" {
    case "foo bar":
        v18 := GenS2{F0: 4294967295, F1: 0, F2: 560}
        _ = v18
    }
    _ = len(v2)
    var v19 string = "\n\t"
    v6 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
