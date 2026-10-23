package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{2, 65535, 433, 2}
    var v2 bool = false
    v3 := []string{("foo bar" + "a"), "foo bar", "foo bar", (("foo bar" + "hello") + "")}
    var v4 uint32 = 0
    var v5 string = "a"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 float32 = (-2.25 + 1.5)
    var v7 int = 42
    _ = len(v1)
    v3 = append(v3, v5)
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v5, v4, v4)
        _ = fi8
    }
    if len(v1) > 0 {
        v1[0] = (v4 ^ v4)
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v7 = 588
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
