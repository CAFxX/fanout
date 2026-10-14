package main

import (
    "fmt"
)


func main() {
    v1 := []string{}
    v2 := []string{}
    v3 := []string{"foo bar", ("" + "hello"), ""}
    var v4 float32 = float32((-2.25 / (-2.25 / 3.14159)))
    v4 = v4
    _ = len(v3)
    var v5 uint = 28
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v6 int16 = -1
    _ = len(v3)
    for fi7 := 0; fi7 < 16; fi7++ {
        v5 -= (v5 & v5)
        v5 += v5
        _ = fi7
    }
    v4 += 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
