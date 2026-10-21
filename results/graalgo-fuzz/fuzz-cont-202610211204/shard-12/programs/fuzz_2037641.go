package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := []int{128, 1, -1, 32768}
    var v4 string = "a"
    var v5 uint = 4294967295
    var v6 int = 32767
    for fi7 := 0; fi7 < 8; fi7++ {
        if len(v3) > 0 {
            v3[0] = -742
        }
        v8 := GenS1{F0: 2, F1: (("\n\t" + (v4 + v4)) + (v4 + "foo bar"))}
        v5 = (v5 ^ v5)
        _ = fi7
        _ = v8
    }
    switch (v6 ^ v6) {
    case 2147483647:
        v9 := GenS1{F0: 1, F1: v4}
        _ = v9
    }
    var v10 int16 = 1000
    var v11 bool = (!(-2.25 <= -2.25))
    var v12 float32 = (-2.25 / 1.0)
    v6 = v6
    v13 := GenS1{F0: 42, F1: "hello"}
    v3 = append(v3, v6)
    fmt.Println(v4, v10, v12)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int = 0
        _ = fi14
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
