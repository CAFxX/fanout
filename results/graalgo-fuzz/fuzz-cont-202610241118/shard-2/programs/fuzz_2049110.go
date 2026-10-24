package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    var v1 int32 = 42
    var v2 uint = 796
    var v3 float64 = 0.0
    v4 := make(map[string]uint32)
    var v5 uint32 = 2147483647
    var v6 bool = (!false)
    var v7 uint64 = 2
    var v8 uint16 = 255
    v9 := GenS1{F0: v1, F1: v8, F2: (-69 != -1)}
    v6 = v6
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint8 = 49
        v12 := v4[((((("" + "") + (((((("\n\t" + ("\n\t" + ("a" + (("hello" + "a") + "a")))) + "hello") + "\n\t") + ("" + "")) + ("\n\t" + "hello")) + "a")) + (("\n\t" + "foo bar") + "\n\t")) + ("foo bar" + "a")) + "a")]
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v5)
    fmt.Println(v5, v7, v6)
    v5 = 731
    var v13 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
