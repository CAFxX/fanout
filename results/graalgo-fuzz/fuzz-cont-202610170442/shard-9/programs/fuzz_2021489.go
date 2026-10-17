package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    var v1 bool = false
    var v2 string = "\n\t"
    var v3 uint8 = 1
    var v4 int64 = 578
    v4 = 0
    var v5 string = ("" + v2)
    v5 = ((v2 + ("foo bar" + ((v2 + ("foo bar" + "a")) + v2))) + v5)
    v6 := []int{32767}
    var v7 uint16 = 42
    v3 = (v3 >> ((4294967295) & 63))
    var v8 float32 = 0.0
    var v9 int = int(v4)
    if (v9 > (v9 % ((v9) | 1))) {
        if (v4 == v4) {
            var v10 float64 = 3.14159
            _ = v10
        }
        if (!true) {
            fmt.Println(v3)
            fmt.Println(v3)
        }
    }
    v11 := GenS1{F0: (v9 % ((853) | 1)), F1: (v3 - (^(v3)))}
    var v12 uint8 = (v3 << ((764) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
