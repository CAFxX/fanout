package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 0
    var v2 string = "\n\t"
    var v3 uint32 = 2147483648
    v4 := make(map[string]uint32)
    var v5 float32 = 3.14159
    v6 := GenS1{F0: 290}
    v5 = (1e10 + v5)
    v2 = (v2 + v2)
    var v7 uint64 = 18446744073709551615
    switch -1 {
    case 42:
        var v8 int = 127
        _ = v8
    default:
        fmt.Println(v1, v7, v2)
    }
    v5 = 1.5
    switch (v2 + (v2 + v2)) {
    case "\n\t":
        v5 = float32(float64(v5))
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            var v10 bool = (!(true || (-1 == 2147483647)))
            _ = v10
        case "a":
            var v11 int = 128
            _ = v11
        case "hello":
            var v12 uint64 = v7
            _ = v12
        default:
            v13 := GenS1{F0: 65535}
            _ = v13
        }
        _ = fi9
    }
    v14 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
