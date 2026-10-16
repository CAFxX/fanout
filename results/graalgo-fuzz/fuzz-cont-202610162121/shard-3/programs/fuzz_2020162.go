package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    var v1 int64 = 146
    v2 := []string{("" + ""), ("a" + "a")}
    var v3 float64 = 1.5
    v4 := make(map[string]string)
    switch "foo bar" {
    case "":
        var v5 int = 9223372036854775807
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int8 = 18
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float64 = v3
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int64 = v1
        var v12 float32 = 1e10
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: "", F1: 42}
    var v14 float64 = v3
    fmt.Println(v14)
    var v15 float32 = -2.25
    v14 = (v3 * (3.14159 - v3))
    fmt.Println(v3, v1)
    fmt.Println(v15, v15)
    var v16 int = 127
    v17 := GenS1{F0: "hello", F1: 4294967295}
    var v18 int8 = -128
    fmt.Println(v1, v18, v14)
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int16 = 851
        v21 := GenS1{F0: "\n\t", F1: 556}
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
