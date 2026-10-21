package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint16 = 235
    var v2 uint = 42
    v3 := []string{"\n\t", "hello", "foo bar"}
    var v4 uint = v2
    var v5 uint = 159
    v4 = (v2 / ((1) | 1))
    switch 32768 {
    case 0:
        if false {
            fmt.Println(v5, v2)
            var v6 uint64 = 1
            fmt.Println(v2, v1)
            _ = v6
        }
    case 128:
        var v7 int16 = 42
        _ = v7
    case 127:
        v8 := GenS1{F0: 127}
        _ = v8
    }
    var v9 int8 = -128
    fmt.Println(v5)
    fmt.Println(v5, v5, v5)
    v2 -= v4
    var v10 int8 = v9
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint = (v5 / (((^(v5))) | 1))
        v5 += (v4 - v2)
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
