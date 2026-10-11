package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 string = "foo bar"
    var v2 int = 0
    var v3 string = "foo bar"
    v4 := []int{(v2 << ((1) & 63)), 128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 uint8 = 228
    var v6 int32 = -764
    v3 = "foo bar"
    v7 := GenS1{F0: 65535}
    v8 := GenS1{F0: 597}
    v1 = ("" + (("\n\t" + v3) + v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
