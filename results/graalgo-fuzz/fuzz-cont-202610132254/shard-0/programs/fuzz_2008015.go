package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint32 = 2
    v2 := []int{32767, 127, 32768, 4294967295}
    var v3 string = ("foo bar" + "foo bar")
    var v4 float32 = (3.14159 / 3.14159)
    v5 := GenS1{F0: (^(v1)), F1: 1}
    var v6 uint = 65535
    if (0.0 <= -2.25) {
        _ = len(v2)
        v6 += v6
    } else {
        v4 -= (0.1 - 0.1)
    }
    var v7 int = -861
    v8 := GenS1{F0: (^(v1)), F1: 287}
    fmt.Println(v3, v7)
    v4 = 0.1
    v6 = uint(v1)
    var v9 uint8 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        if len(v2) > 0 {
            v2[0] = 32767
        }
        var v11 uint64 = 1
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
