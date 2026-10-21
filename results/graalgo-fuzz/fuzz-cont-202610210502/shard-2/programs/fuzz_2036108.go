package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 int64 = 1
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint64 = 2
        switch ("foo bar" + "") {
        case "a":
            for fi6 := 0; fi6 < 3; fi6++ {
                v5 = v5
                _ = fi6
            }
        default:
            v7 := GenS1{F0: 0, F1: -1}
            _ = v7
        }
        var v8 uint32 = 42
        _ = fi4
        _ = v5
        _ = v8
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        switch 0 {
        case 1:
            v9 := []int{128, 32768, 128, -9223372036854775808}
            _ = v9
        case 128:
            fmt.Println(v2, v2)
        case 32768:
            switch "hello" {
            case "a":
                if false {
                    v2 -= v2
                }
            }
        }
    default:
        v2 -= (v2 & v2)
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v3 -= (v3 * 195)
        _ = fi10
    }
    var v11 int = -288
    var v12 uint16 = 255
    var v13 uint32 = 2147483647
    fmt.Println(v12)
    v14 := GenS1{F0: 42, F1: 0}
    var v15 uint16 = (v12 / ((v12) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
