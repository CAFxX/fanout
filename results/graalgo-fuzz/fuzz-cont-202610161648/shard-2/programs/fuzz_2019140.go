package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int }

func main() {
    var v1 int32 = 42
    var v2 uint = 42
    var v3 bool = false
    v4 := []string{"foo bar", ("foo bar" + "\n\t"), "foo bar", "hello"}
    fmt.Println(v2, v1, v3)
    if true {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    } else {
        v3 = v3
        v3 = v3
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(v3)
        var v6 float64 = 1.5
        _ = fi5
        _ = v6
    }
    if (4294967295 < 348) {
        var v7 int8 = 0
        v2 = v2
        _ = v7
    }
    v8 := GenS1{F0: "\n\t", F1: 128}
    for fi9 := 0; fi9 < 14; fi9++ {
        if (!(((1 >= 255) && true) || (1 == 1000))) {
            fmt.Println(v1, v1)
        }
        var v10 uint16 = 255
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    v3 = (v2 < v2)
    var v11 float32 = 0.0
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: (fi12 | (fi12 ^ fi12))}
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
