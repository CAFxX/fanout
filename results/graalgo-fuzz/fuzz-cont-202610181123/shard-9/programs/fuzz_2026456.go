package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint8, p1 uint32) int16 {
    var v6 uint16 = 493
    switch (("\n\t" + ("\n\t" + "foo bar")) + "") {
    case "\n\t":
        var v7 uint = 13
        _ = v7
    case "hello":
        switch "\n\t" {
        case "hello":
            fmt.Println(p1, p1, v6)
        case "a":
            p1 -= (^(p1))
        case "foo bar":
            fmt.Println(p0)
        default:
            var v8 uint16 = (^(v6))
            _ = v8
        }
    }
    v6 = (^(v6))
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() uint16 {
    var v9 int8 = 97
    v9 = v9
    v10 := []int{128, 127, 429, 127}
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v9, fi11)
        _ = fi11
    }
    _ = v10
    _ = v9
    return 584
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 uint32 = 2147483647
    var v5 uint = 0
    for fi12 := 0; fi12 < 18; fi12++ {
        v5 = (^(v5))
        _ = fi12
    }
    v4 += (^(v4))
    if len(v3) > 0 {
        v3[0] = ("\n\t" + ("foo bar" + "\n\t"))
    }
    v13 := GenS1{F0: (!(!true)), F1: 42}
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 float64 = 0.1
        _ = fi14
        _ = v15
    }
    var v16 int32 = 1
    var v17 uint8 = 127
    v18 := GenS1{F0: (("\n\t" + ("foo bar" + "hello")) >= "foo bar"), F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
