package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = -475
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 uint = 42
    var v5 uint64 = 42
    fmt.Println(v2, v4, v5)
    fmt.Println(v2, v1, v2)
    fmt.Println(v1)
    var v6 int8 = -128
    v5 = 459
    v7 := GenS1{F0: ""}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = v5
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 float64 = (0.0 * -2.25)
        _ = fi10
        _ = v11
    }
    v2 = v2
    v12 := v3[448]
    switch (^(v12)) {
    case 9223372036854775807:
        var v13 uint64 = (^(v5))
        _ = v13
    case 32767:
        v6 = -122
    }
    switch "a" {
    case "foo bar":
        var v14 float32 = 1e10
        _ = v14
    case "":
        v1 -= v1
    }
    var v15 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
