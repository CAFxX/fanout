package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint8 = 187
    v2 := make(map[string]string)
    var v3 int32 = 42
    var v4 uint = 255
    var v5 float32 = (float32(1.5) * 0.1)
    v6 := []int{-9223372036854775808, 9223372036854775807, -1}
    var v7 int = -75
    v8 := GenS1{F0: 461, F1: v7}
    v6 = append(v6, -713)
    fmt.Println(v1, v3)
    v5 = float32(0.1)
    var v9 int32 = v3
    v5 = v5
    for fi10 := 0; fi10 < 4; fi10++ {
        if ((!true) || ((1.5 + 3.14159) == float64(((1.5 * 1e10) * v5)))) {
            if (uint32(v1) >= 65535) {
                var v11 string = "foo bar"
                _ = v11
            } else {
                v1 -= v1
                var v12 string = (("" + "\n\t") + ("a" + (("foo bar" + "a") + "\n\t")))
                var v13 uint8 = (v1 & (v1 << ((42) & 63)))
                _ = v12
                _ = v13
            }
            v4 = (v4 - (v4 - 4294967295))
            v14 := GenS2{F0: true, F1: 49}
            _ = v14
        }
        fmt.Println(v7, v7)
        fmt.Println(v3)
        _ = fi10
    }
    if true {
        var v15 int16 = -590
        var v16 string = ((("\n\t" + ("" + "\n\t")) + "a") + (("hello" + (("a" + (("a" + "") + "foo bar")) + "")) + "foo bar"))
        var v17 uint16 = 0
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
