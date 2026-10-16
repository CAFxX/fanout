package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := []string{"foo bar", ("\n\t" + "foo bar"), ("a" + "hello")}
    v2 := make(map[string]bool)
    v3 := []uint32{2147483648}
    var v4 int8 = -128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi5 := 0; fi5 < 14; fi5++ {
        v4 = 52
        _ = fi5
    }
    fmt.Println(v4)
    v6 := GenS1{F0: 229, F1: 4294967295}
    switch ("hello" + "") {
    case "a":
        var v7 float64 = 1.5
        _ = v7
    }
    if true {
        v4 = (v4 / ((v4) | 1))
    }
    var v8 int64 = -9223372036854775808
    var v9 int64 = v8
    var v10 int = -1
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 float32 = float32(0.1)
        _ = fi11
        _ = v12
    }
    v9 = v9
    fmt.Println(v10)
    var v13 uint32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
