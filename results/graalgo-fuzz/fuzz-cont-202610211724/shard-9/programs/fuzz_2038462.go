package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int32 }

func main() {
    var v1 uint16 = 749
    var v2 int16 = 32767
    var v3 int32 = 2147483647
    v4 := []int{127}
    if len(v4) > 0 {
        v4[0] = 1
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        switch "foo bar" {
        case "a":
            var v6 uint64 = 0
            _ = v6
        }
        _ = fi5
    }
    v3 = 951
    if false {
        var v7 uint8 = 255
        var v8 float64 = -2.25
        _ = v7
        _ = v8
    } else {
        v3 += (-(v3))
        var v9 uint16 = 1
        _ = v9
    }
    v1 = 490
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 -= v2
    var v10 int8 = 100
    if len(v4) > 0 {
        v4[0] = 1
    }
    v4 = append(v4, 386)
    switch int(v1) {
    case -9223372036854775808:
        v11 := GenS1{F0: (!((0.1 + (0.1 + (-2.25 + 1.5))) <= 0.0)), F1: (^(v1)), F2: (v3 * v3)}
        _ = v11
    case 32768:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
