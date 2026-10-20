package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []int{127}
    var v2 string = ""
    var v3 uint8 = 42
    v4 := []int{-585, -134, -154, -9223372036854775808}
    var v5 float32 = 1e10
    switch 9223372036854775807 {
    case 0:
        v6 := GenS1{F0: -1}
        _ = v6
    case -423:
        var v7 int8 = -128
        _ = v7
    default:
        _ = len(v4)
    }
    v5 = v5
    v5 = 0.0
    var v8 uint64 = 740
    for fi9 := 0; fi9 < 11; fi9++ {
        if false {
            for fi10 := 0; fi10 < 11; fi10++ {
                v2 = ("a" + v2)
                _ = fi10
            }
        }
        _ = fi9
    }
    v5 = v5
    var v11 uint = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
