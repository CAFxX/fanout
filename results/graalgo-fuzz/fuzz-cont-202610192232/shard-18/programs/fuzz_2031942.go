package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 255
    v3 := []int{4294967295, 42, -797}
    fmt.Println(v2)
    var v4 uint = 255
    var v5 int = 0
    v6 := GenS1{F0: ("hello" + ("foo bar" + (("" + ((("" + ("" + "")) + ("\n\t" + "foo bar")) + "hello")) + "")))}
    v5 -= v5
    fmt.Println(v5, v2)
    v7 := v1[869]
    var v8 float64 = float64(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
