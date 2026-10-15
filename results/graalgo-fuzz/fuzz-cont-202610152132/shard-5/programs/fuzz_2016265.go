package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    var v3 bool = false
    var v4 float32 = 1e10
    v5 := []uint32{2147483648, 4294967295}
    v6 := GenS1{F0: 0, F1: 1}
    if false {
        v4 += (3.14159 - v4)
    } else {
        v1 -= -37
        if false {
            var v7 int = 42
            _ = v7
        }
        v1 -= v1
    }
    switch "" {
    case "":
        fmt.Println(v4)
    case "hello":
        var v8 uint32 = 42
        _ = v8
    }
    fmt.Println(v4, v1, v1)
    if (890 == 127) {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 string = "\n\t"
            if (318 < 2) {
                var v11 bool = true
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 uint = 703
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        switch "foo bar" {
        case "hello":
            fmt.Println(v1, v4, v1)
        case "a":
            var v14 int32 = 42
            _ = v14
        }
    }
    if true {
        v3 = v3
    }
    v15 := v2[2147483647]
    var v16 uint16 = 333
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
