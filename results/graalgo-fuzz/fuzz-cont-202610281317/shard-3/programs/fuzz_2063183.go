package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint64 = 2
    var v2 uint = 18446744073709551615
    v3 := make(map[int]string)
    var v4 uint8 = 62
    v3[128] = ""
    var v5 int16 = -32768
    switch ("\n\t" + ("\n\t" + "a")) {
    case "hello":
        v4 = 245
    case "":
        v1 += uint64(v4)
    }
    var v6 uint8 = v4
    v5 -= (v5 + v5)
    fmt.Println(v2, v1)
    fmt.Println(v2)
    var v7 uint32 = 2147483647
    v7 = (v7 / ((710) | 1))
    v8 := GenS1{F0: -186}
    fmt.Println(v7)
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v5)
        _ = fi9
    }
    v5 = v5
    fmt.Println(v7, v4, v4)
    var v10 int = 656
    v11 := GenS1{F0: int(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
