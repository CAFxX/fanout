package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint64 = 0
    var v2 float32 = float32((0.0 / 1.5))
    var v3 int16 = 1
    var v4 bool = false
    v4 = false
    v5 := GenS1{F0: 263, F1: (v3 / ((0) | 1)), F2: 42}
    switch -781 {
    case -716:
        fmt.Println(v2, v2)
    default:
        var v6 string = ("hello" + "\n\t")
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        switch "foo bar" {
        case "a":
            v8 := []int{fi7, (fi7 * 9223372036854775807)}
            _ = v8
        case "hello":
            fmt.Println(v1)
        case "foo bar":
            var v9 uint = 665
            _ = v9
        default:
            v10 := []int{0}
            _ = v10
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            var v12 uint16 = 65535
            switch "\n\t" {
            case "hello":
                var v13 int8 = 42
                _ = v13
            case "foo bar":
                v4 = v4
            case "":
                v14 := GenS1{F0: v1, F1: -32768, F2: 357}
                _ = v14
            default:
                v15 := []int{(fi11 - fi11), fi11, -406, -9223372036854775808}
                _ = v15
            }
            _ = v12
        }
        v16 := GenS2{F0: 42}
        _ = fi11
        _ = v16
    }
    v2 = v2
    fmt.Println(v2)
    v17 := GenS1{F0: (v1 ^ v1), F1: -96, F2: 65535}
    v18 := []int{-184, -578}
    v1 = (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
