package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 uint16 = 590
    var v2 bool = false
    var v3 uint32 = 2
    var v4 uint = 255
    var v5 bool = false
    var v6 int = 0
    fmt.Println(v4)
    fmt.Println(v6, v4)
    if v5 {
        v7 := GenS1{F0: 1, F1: "foo bar"}
        fmt.Println(v5, v2)
        for fi8 := 0; fi8 < 7; fi8++ {
            v1 -= v1
            _ = fi8
        }
        _ = v7
    }
    var v9 bool = v2
    var v10 int64 = 1
    var v11 uint32 = v3
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint64 = 18446744073709551615
        v14 := GenS1{F0: 0, F1: ("a" + (("\n\t" + "hello") + ("\n\t" + "foo bar")))}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := []int{v6, (v6 & (v6 ^ v6))}
    var v16 int16 = 1
    switch "hello" {
    case "":
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 float32 = 0.1
            _ = fi17
            _ = v18
        }
    default:
        var v19 int16 = v16
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
