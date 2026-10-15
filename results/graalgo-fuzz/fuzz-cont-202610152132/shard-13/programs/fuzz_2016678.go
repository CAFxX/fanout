package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }

func main() {
    v1 := []int{-445, -650, 40}
    var v2 uint16 = 32767
    v3 := []string{"foo bar", "", "foo bar"}
    v4 := make(map[int]uint32)
    v5 := []string{("foo bar" + "\n\t"), "foo bar", ""}
    for fi6 := 0; fi6 < 10; fi6++ {
        v2 = (v2 + v2)
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint8 = 42
            _ = fi7
            _ = v8
        }
        var v9 float64 = ((((0.1 - 1e10) / 1.0) - 0.1) + 1e10)
        _ = fi6
        _ = v9
    }
    var v10 string = ("a" + ("hello" + ""))
    _ = len(v5)
    v11 := v4[-883]
    v11 = v11
    v12 := GenS1{F0: 42, F1: -92, F2: "a"}
    fmt.Println(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v10)
    var v13 uint = 42
    v14 := GenS1{F0: -83, F1: -499, F2: v10}
    fmt.Println(v13)
    v15 := GenS1{F0: -2147483648, F1: -1, F2: v10}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
