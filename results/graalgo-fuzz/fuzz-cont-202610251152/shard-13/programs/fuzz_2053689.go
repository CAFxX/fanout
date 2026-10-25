package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    v1 := []int{0, 4294967295, 4294967295, 9223372036854775807}
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    v5 := []string{"foo bar"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int16 = 1
        var v9 int32 = 2147483647
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        var v10 int = 0
        var v11 int32 = 0
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint = 2
            v14 := GenS1{F0: 1000, F1: 798}
            v6 = (v6 % ((949) | 1))
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 uint32 = 42
        _ = v15
    }
    var v16 float32 = 1.5
    v6 += (^(v6))
    if ((!(42 < 1)) && true) {
        if (255 != 2) {
            v16 = (v16 - float32((1.5 - -2.25)))
        } else {
            var v17 float64 = 1.5
            fmt.Println(v17)
            _ = v17
        }
    }
    v16 -= (v16 * v16)
    var v18 float32 = ((((v16 + v16) + 1e10) + v16) - v16)
    v16 = 1.5
    _ = len(v5)
    fmt.Println(v18)
    if true {
        v18 = v16
    } else {
        switch "" {
        case "a":
            fmt.Println(v6)
        case "":
            v19 := GenS1{F0: 660, F1: -130}
            _ = v19
        case "foo bar":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    var v20 string = ("foo bar" + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
