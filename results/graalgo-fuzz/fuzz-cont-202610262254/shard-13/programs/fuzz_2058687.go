package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0(p0 uint) string {
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 2147483647
    var v3 uint64 = 0
    if true {
        fmt.Println(v2, v2)
    }
    if true {
        v2 -= v2
        for fi4 := 0; fi4 < 13; fi4++ {
            v2 = 128
            var v5 string = "foo bar"
            fmt.Println(v2, v3, v5)
            _ = fi4
            _ = v5
        }
        var v6 int = -689
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            v3 -= 42
            v2 -= v2
            v2 = 562
            _ = fi8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v1["foo bar"] = "a"
        _ = fi9
    }
    v10 := GenS1{F0: -327, F1: v3, F2: 628}
    fmt.Println(v2, v2)
    v11 := []int{}
    for fi12 := 0; fi12 < 8; fi12++ {
        switch ("hello" + ("" + "foo bar")) {
        case "foo bar":
            v13 := GenS2{F0: 9223372036854775807, F1: 2, F2: -927}
            _ = v13
        }
        var v14 uint8 = 136
        _ = fi12
        _ = v14
    }
    fmt.Println(v3)
    v3 -= 372
    v2 = (-(v2))
    var v15 int8 = -91
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}
