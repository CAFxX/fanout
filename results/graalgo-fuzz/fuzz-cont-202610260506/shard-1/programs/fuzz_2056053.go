package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []int{128}
    var v2 string = ""
    v3 := []string{v2, (v2 + "hello"), "a"}
    v4 := []int{}
    v2 = v2
    v2 = ""
    v2 = (v2 + (("\n\t" + v2) + ((((("" + v2) + "\n\t") + "foo bar") + "hello") + v2)))
    _ = len(v4)
    v5 := GenS1{F0: 0}
    var v6 uint16 = 0
    v6 = (^(v6))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
