package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int16 = 1000
    var v2 float32 = float32(3.14159)
    var v3 bool = true
    v4 := make(map[uint32]string)
    v3 = (2 < 1)
    v2 = v2
    fmt.Println(v2, v2, v1)
    fmt.Println(v1)
    if (292 == -1) {
        v3 = true
        v1 -= (^(v1))
        v5 := GenS1{F0: ("hello" + "foo bar")}
        _ = v5
    }
    switch "\n\t" {
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        v6 := GenS2{F0: (^(v1))}
        _ = v6
    case "a":
        var v7 float32 = -2.25
        _ = v7
    }
    var v8 int32 = 1
    fmt.Println(v8)
    var v9 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
