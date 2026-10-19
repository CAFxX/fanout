package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 int64 = 0
    var v5 float64 = (0.1 * 0.1)
    switch -433 {
    case 0:
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int64 = (^(v4))
            _ = fi6
            _ = v7
        }
    case 42:
        switch "a" {
        case "hello":
            v8 := GenS1{F0: 127, F1: 315}
            _ = v8
        case "":
            var v9 uint64 = 773
            _ = v9
        case "a":
            fmt.Println(v4, v1)
        }
    }
    v10 := []int{}
    v11 := GenS1{F0: -67, F1: 757}
    if len(v10) > 0 {
        v10[0] = 9223372036854775807
    }
    var v12 int8 = 42
    v4 -= v4
    _ = len(v10)
    fmt.Println(v1, v12, v4)
    switch "" {
    case "hello":
        fmt.Println(v5)
    case "":
        if false {
            if true {
                v12 = v12
            }
            v1 -= v1
        } else {
            _ = len(v3)
            var v13 string = ("a" + ("" + ("" + "a")))
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 uint = 518
                v4 -= 42
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    }
    fmt.Println(v12, v4, v12)
    v2[-378] = ""
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint8 = 255
        fmt.Println(v4, v4)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
