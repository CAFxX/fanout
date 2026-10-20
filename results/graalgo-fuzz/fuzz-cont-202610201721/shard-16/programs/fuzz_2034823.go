package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 42
    var v3 int32 = -894
    v4 := make(map[int]int)
    var v5 uint16 = 0
    v6 := GenS1{F0: 1}
    v2 = (v2 - (^(v2)))
    for fi7 := 0; fi7 < 16; fi7++ {
        v1 = (-(v1))
        _ = fi7
    }
    v8 := []int{}
    switch ("\n\t" + "") {
    case "\n\t":
        fmt.Println(v5)
    default:
        v9 := GenS1{F0: 567}
        _ = v9
    }
    var v10 bool = true
    var v11 string = ""
    v2 = (v2 * (^(v2)))
    if v10 {
        var v12 float64 = -2.25
        v13 := GenS1{F0: 841}
        _ = v12
        _ = v13
    } else {
        var v14 bool = v10
        if false {
            _ = len(v4)
        } else {
            switch "a" {
            case "a":
                _ = len(v8)
            }
        }
        _ = v14
    }
    var v15 int = 32768
    v5 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
