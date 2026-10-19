package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    p0 = p0
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: "\n\t", F1: "\n\t"}
        _ = fi6
        _ = v7
    }
    switch -1 {
    case 127:
        fmt.Println(p0)
    case 42:
        p0 = p0
    }
    _ = p0
    return 588
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 float64 = float64(0.0)
    var v4 int32 = -1
    v5 := make(map[string]uint32)
    switch ("" + "foo bar") {
    case "a":
        v3 = 3.14159
    case "foo bar":
        fmt.Println(v3, v3)
    case "\n\t":
        var v8 float32 = 1e10
        _ = v8
    }
    switch -9223372036854775808 {
    case 919:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        fmt.Println(v3, v3)
    }
    v4 = v4
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v3, fi9, v4)
        v4 = (v4 / (((-(v4))) | 1))
        var v10 float64 = v3
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v3)
    fmt.Println(v4)
    v11 := GenS2{F0: 9223372036854775807}
    v12 := GenS1{F0: ("a" + "\n\t"), F1: ("\n\t" + ("a" + "hello"))}
    var v13 int = -505
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
