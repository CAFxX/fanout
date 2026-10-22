package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int8 = -63
    var v2 int = -744
    var v3 int8 = v1
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    v6 := make(map[string]uint32)
    var v7 int32 = -723
    v2 -= (v2 & (v2 << ((2) & 63)))
    v8 := GenS1{F0: 984, F1: 32767}
    v9 := GenS1{F0: 1, F1: 255}
    var v10 float32 = -2.25
    switch "hello" {
    case "\n\t":
        v11 := GenS1{F0: 65535, F1: 65535}
        _ = v11
    default:
        v1 += v3
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch "\n\t" {
        case "hello":
            switch "hello" {
            case "a":
                var v13 int16 = -32768
                _ = v13
            default:
                fmt.Println(v3)
            }
        case "foo bar":
            v14 := []int{(v2 * v2), v2, fi12}
            _ = v14
        default:
            v7 = -1
        }
        v15 := v6["\n\t"]
        fmt.Println(fi12, v15, fi12)
        _ = fi12
        _ = v15
    }
    v2 = v2
    if true {
        var v16 string = ""
        _ = v16
    } else {
        fmt.Println(v1, v7)
    }
    fmt.Println(v1, v7)
    v7 -= -774
    var v17 int32 = -305
    var v18 uint16 = 2
    v3 = v1
    v3 = -43
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
