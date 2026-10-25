package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -1
    var v3 int8 = 1
    switch "\n\t" {
    case "a":
        v4 := GenS1{F0: 1}
        _ = v4
    default:
        var v5 string = (("\n\t" + "foo bar") + "\n\t")
        _ = v5
    }
    if false {
        _ = len(v1)
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 string = ("" + "\n\t")
            _ = fi6
            _ = v7
        }
        var v8 uint16 = 65535
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        var v9 uint = 255
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint64 = 1
            _ = fi10
            _ = v11
        }
    case "foo bar":
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS2{F0: -918, F1: fi12}
            var v14 uint64 = 42
            var v15 bool = (88 <= 42)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        v3 = v3
    }
    v16 := GenS1{F0: 0}
    fmt.Println(v3, v2)
    var v17 uint32 = 0
    v18 := GenS2{F0: 9223372036854775807, F1: 0}
    var v19 int8 = (v3 % (((^(v3))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
