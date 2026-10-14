package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = 0
    var v2 string = "foo bar"
    var v3 int64 = (-(v1))
    var v4 int16 = 32767
    var v5 int = -9223372036854775808
    v2 = ("a" + v2)
    fmt.Println(v2, v4)
    v6 := GenS1{F0: 128}
    for fi7 := 0; fi7 < 10; fi7++ {
        v3 = v1
        var v8 uint = 42
        _ = fi7
        _ = v8
    }
    v2 = ("" + ((v2 + v2) + v2))
    v1 += v1
    v9 := []int{v5, 42, -2147483648}
    fmt.Println(v3, v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
