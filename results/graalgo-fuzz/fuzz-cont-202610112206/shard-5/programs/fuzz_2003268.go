package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := []string{"\n\t"}
    var v2 uint = 12
    var v3 uint16 = 1
    fmt.Println(v2, v2)
    fmt.Println(v2)
    fmt.Println(v2, v2)
    v2 += v2
    v4 := GenS1{F0: 0, F1: 128}
    v3 -= (v3 % ((v3) | 1))
    for fi5 := 0; fi5 < 2; fi5++ {
        switch "\n\t" {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v6 uint64 = 0
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int8 = 1
        _ = fi7
        _ = v8
    }
    v2 = v2
    var v9 int64 = -183
    fmt.Println(v2, v2, v2)
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint32 = 2
        fmt.Println(v11, v2)
        _ = fi10
        _ = v11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
