package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 int64 = 283
    v5 := GenS1{F0: 2, F1: 0, F2: 255}
    var v6 int = 2147483647
    var v7 int64 = 795
    v6 -= 0
    if true {
        fmt.Println(v6, v4)
    }
    fmt.Println(v6, v4, v4)
    var v8 int64 = (v4 & v4)
    v9 := GenS1{F0: 255, F1: 1, F2: 65535}
    switch "\n\t" {
    case "hello":
        var v10 int8 = 1
        _ = v10
    case "\n\t":
        var v11 string = ("a" + "a")
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int64 = v8
            switch (^(v6)) {
            case 42:
                var v14 string = ("hello" + "a")
                _ = v14
            case 2147483647:
                v8 = v4
            default:
                fmt.Println(v6)
            }
            _ = fi12
            _ = v13
        }
    }
    v4 -= -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
