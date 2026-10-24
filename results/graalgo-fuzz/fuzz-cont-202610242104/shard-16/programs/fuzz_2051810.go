package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    var v3 uint64 = v2
    var v4 int32 = 31
    v5 := make(map[int]bool)
    var v6 int8 = 100
    v7 := GenS1{F0: -261, F1: "a", F2: (!false)}
    v8 := GenS1{F0: 32768, F1: ("\n\t" + ("\n\t" + "")), F2: (false && false)}
    v2 -= v3
    var v9 float32 = 0.1
    switch "hello" {
    case "":
        v3 = 2
    case "foo bar":
        var v10 int8 = v6
        _ = v10
    case "\n\t":
        var v11 uint32 = 42
        _ = v11
    }
    if false {
        var v12 float64 = float64(3.14159)
        var v13 string = ("" + "hello")
        _ = v12
        _ = v13
    } else {
        var v14 bool = (249 != 1)
        if true {
            v14 = true
            v6 = -101
        } else {
            if (float64(0.0) == 3.14159) {
                v6 = v6
            }
            v2 = (v2 ^ (v3 / ((v2) | 1)))
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(v2, v9)
            _ = fi15
        }
        _ = v14
    }
    v2 = v3
    if true {
        v6 = v6
        var v16 uint64 = v2
        _ = v16
    } else {
        var v17 uint32 = 0
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
