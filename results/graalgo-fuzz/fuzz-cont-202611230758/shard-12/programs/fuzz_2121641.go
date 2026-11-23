package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("\n\t" + ("" + ""))
    v3 := []int{42, 32768, 1, -425}
    v4 := []string{}
    v2 = v2
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 0, F1: 42}
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 32767
    switch "foo bar" {
    case "\n\t":
        var v8 float64 = -2.25
        _ = v8
    default:
        fmt.Println(v7, v2)
    }
    var v9 bool = (true || (!true))
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint = 18446744073709551615
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (v7 << ((50) & 63)), F1: 1}
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
