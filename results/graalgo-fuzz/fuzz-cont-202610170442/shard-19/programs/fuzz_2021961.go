package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }

func main() {
    v1 := []string{("" + (("\n\t" + ("\n\t" + "a")) + "\n\t")), ((("" + "foo bar") + ("a" + "a")) + "foo bar")}
    var v2 uint16 = 145
    var v3 float32 = float32(3.14159)
    var v4 int16 = -928
    v5 := []int{32767, 42, 372}
    var v6 int64 = -928
    for fi7 := 0; fi7 < 18; fi7++ {
        v4 = v4
        _ = fi7
    }
    fmt.Println(v6, v2, v3)
    switch "foo bar" {
    case "foo bar":
        var v8 string = "hello"
        _ = v8
    }
    v6 = -976
    v4 = (v4 - v4)
    switch "a" {
    case "":
        var v9 bool = true
        _ = v9
    case "\n\t":
        v6 -= 0
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch "\n\t" {
    case "foo bar":
        var v10 int8 = 0
        _ = v10
    default:
        fmt.Println(v4, v4, v6)
    }
    v2 = v2
    fmt.Println(v3, v6)
    var v11 int64 = v6
    v12 := GenS1{F0: 2, F1: ("\n\t" + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
