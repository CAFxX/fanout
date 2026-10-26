package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = -859
    v2 := make(map[int]bool)
    var v3 int64 = v1
    var v4 float64 = 1e10
    v5 := []int{-2147483648, 42, 32767}
    var v6 int8 = 42
    v7 := v2[0]
    var v8 float32 = 1.5
    var v9 string = ("a" + (("a" + "\n\t") + (((("" + ((("\n\t" + "\n\t") + (("" + "foo bar") + "hello")) + (("" + ("foo bar" + "a")) + "hello"))) + ("a" + "")) + "") + "a")))
    v7 = false
    var v10 uint8 = 255
    _ = len(v2)
    _ = len(v5)
    if v7 {
        var v11 uint16 = 32767
        _ = v11
    }
    fmt.Println(v3, v6)
    v9 = v9
    v8 = v8
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: v9}
        for fi14 := 0; fi14 < 8; fi14++ {
            v8 -= float32(1.5)
            if ((v6 / ((v6) | 1)) <= v6) {
                v9 = v9
                fmt.Println(fi14, v4)
                _ = len(v2)
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v15 := GenS1{F0: v9}
    var v16 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
