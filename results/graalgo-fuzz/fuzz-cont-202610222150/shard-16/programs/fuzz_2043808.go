package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 int = 0
    var v2 uint = 18446744073709551615
    v3 := []uint32{}
    var v4 string = ("a" + "hello")
    var v5 uint = (^(v2))
    var v6 float32 = ((((float32(3.14159) / 1.0) + ((float32(3.14159) - 1e10) / 3.14159)) * 3.14159) * (float32(1.5) * 0.1))
    var v7 string = (v4 + v4)
    v8 := GenS1{F0: 32767, F1: 219, F2: 225}
    switch "a" {
    case "":
        v9 := GenS2{F0: 127, F1: 4294967295}
        _ = v9
    case "hello":
        fmt.Println(v7)
    case "foo bar":
        for fi10 := 0; fi10 < 16; fi10++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi10
        }
    }
    v3 = append(v3, 4294967295)
    switch v1 {
    case 32768:
        v1 = 451
    }
    switch v7 {
    case "\n\t":
        var v11 int64 = 1
        _ = v11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
