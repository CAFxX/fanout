package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 string = (("" + "\n\t") + "")
    var v2 string = (v1 + v1)
    var v3 uint8 = 0
    var v4 int64 = -1
    v5 := []string{}
    var v6 int = 0
    var v7 uint32 = 2147483648
    fmt.Println(v3, v1, v3)
    v6 = v6
    v7 += (v7 >> ((269) & 63))
    v6 += v6
    v5 = append(v5, ("a" + v1))
    _ = len(v5)
    v7 = (v7 >> ((42) & 63))
    v8 := GenS1{F0: 1, F1: 2}
    v4 = v4
    v2 = v1
    v9 := GenS1{F0: 1000, F1: 1}
    v6 -= -512
    var v10 int64 = (^(v4))
    if len(v5) > 0 {
        v5[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
