package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{0, 2147483648, 2147483648, 527}
    v3 := make(map[string]bool)
    var v4 int16 = 681
    switch "\n\t" {
    case "":
        for fi5 := 0; fi5 < 15; fi5++ {
            if true {
                _ = len(v1)
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v6 bool = false
                _ = v6
            }
            _ = fi5
        }
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        v4 = v4
    }
    var v7 uint = 0
    fmt.Println(v4, v4)
    if ((!false) && true) {
        fmt.Println(v4)
        var v8 int64 = -1
        v9 := GenS1{F0: v8, F1: 32767}
        _ = v8
        _ = v9
    } else {
        fmt.Println(v4, v4)
        v7 = (v7 - (v7 >> ((v7) & 63)))
    }
    fmt.Println(v7)
    var v10 uint8 = 209
    fmt.Println(v4, v10, v10)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
