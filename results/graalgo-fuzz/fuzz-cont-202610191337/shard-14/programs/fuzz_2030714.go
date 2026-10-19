package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []uint32{}
    var v2 int32 = -594
    var v3 float32 = 1.5
    v4 := make(map[uint32]bool)
    if (false && false) {
        v2 = v2
        v3 = v3
        fmt.Println(v3, v3)
    } else {
        _ = len(v1)
    }
    if false {
        v5 := GenS1{F0: (v2 * v2)}
        v6 := v4[2]
        fmt.Println(v2, v2)
        _ = v5
        _ = v6
    }
    fmt.Println(v2, v2)
    var v7 uint8 = 255
    var v8 string = (("a" + ("\n\t" + ("hello" + ((("" + "\n\t") + "") + "foo bar")))) + "\n\t")
    var v9 float32 = v3
    _ = len(v4)
    var v10 float64 = 1.5
    var v11 float32 = (3.14159 / (-2.25 + v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
