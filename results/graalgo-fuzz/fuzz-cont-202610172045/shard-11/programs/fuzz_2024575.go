package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []string{"foo bar", ("foo bar" + "a"), ("a" + ("" + "hello"))}
    v2 := make(map[string]uint32)
    var v3 uint = 42
    var v4 float64 = 0.0
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 4294967295, F1: 65535}
    if (!false) {
        switch 32768 {
        case -9223372036854775808:
            v3 = v3
        default:
            fmt.Println(v3, v3)
        }
    } else {
        var v7 int = -81
        if false {
            var v8 float64 = -2.25
            _ = v8
        } else {
            fmt.Println(v7)
            _ = len(v5)
            _ = len(v1)
        }
        v5["\n\t"] = (!false)
        _ = v7
    }
    var v9 float32 = float32(v4)
    switch 192 {
    case 2147483647:
        fmt.Println(v4)
    }
    fmt.Println(v3, v9, v9)
    if false {
        switch 4294967295 {
        case 128:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        v10 := GenS1{F0: v3, F1: 2}
        v3 = (^(v3))
        _ = v10
    }
    v11 := v5[("hello" + ("a" + "foo bar"))]
    v3 = v3
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
