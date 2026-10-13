package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0(p0 uint64) int16 {
    var v4 float64 = 1.5
    v4 = v4
    _ = p0
    _ = v4
    return 42
}

func main() {
    v1 := []string{("" + "foo bar"), (("a" + ("" + ("\n\t" + ""))) + "foo bar"), "", "hello"}
    var v2 int16 = 1
    v3 := make(map[uint32]string)
    v5 := GenS1{F0: 42, F1: 836}
    var v6 int = 42
    v2 -= genFn0(47)
    v6 -= (v6 ^ (v6 - (v6 + v6)))
    var v7 int64 = 9223372036854775807
    var v8 uint = 0
    for fi9 := 0; fi9 < 18; fi9++ {
        v6 += (v6 + fi9)
        _ = fi9
    }
    switch "a" {
    case "":
        v8 = v8
    case "\n\t":
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(v7, fi10)
            _ = fi10
        }
    default:
        var v11 uint = 494
        _ = v11
    }
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
