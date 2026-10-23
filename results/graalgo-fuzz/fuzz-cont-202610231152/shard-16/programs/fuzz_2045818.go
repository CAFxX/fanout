package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int8 = 100
    var v2 float64 = -2.25
    var v3 float64 = v2
    v4 := []int{0, 1, 128}
    if (false || (42 >= 65535)) {
        v5 := GenS1{F0: -1}
        var v6 string = "a"
        var v7 float64 = 0.0
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 float64 = v3
        v9 := GenS1{F0: -1}
        fmt.Println(v1, v3)
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -586}
    fmt.Println(v1)
    var v11 string = "\n\t"
    switch "\n\t" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = len(v4)
    var v12 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
