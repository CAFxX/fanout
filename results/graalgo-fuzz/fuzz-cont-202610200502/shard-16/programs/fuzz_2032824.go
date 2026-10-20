package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }

func main() {
    var v1 int8 = 127
    var v2 bool = (606 != 42)
    v3 := []uint32{4294967295, 4294967295, 42, 2}
    var v4 bool = v2
    v5 := []string{(("hello" + ("foo bar" + (("hello" + "\n\t") + ("" + "a")))) + (("a" + "") + ("a" + (("" + ("" + ("\n\t" + ""))) + "\n\t")))), "\n\t", "hello"}
    fmt.Println(v4)
    fmt.Println(v1)
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(v1)
        fmt.Println(v2)
        _ = fi6
    }
    v2 = v2
    if v4 {
        fmt.Println(v2)
    }
    v7 := GenS1{F0: 100, F1: 2, F2: 0}
    var v8 uint64 = 0
    v1 = -128
    fmt.Println(v8, v1, v1)
    v8 -= (v8 | v8)
    var v9 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
