package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 float32 = float32(-2.25)
    var v4 int8 = 100
    if true {
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(v4, v3)
            var v6 int = fi5
            _ = fi5
            _ = v6
        }
        v3 -= 3.14159
        for fi7 := 0; fi7 < 19; fi7++ {
            if true {
                var v8 float32 = v3
                _ = v8
            } else {
                v4 = v4
                v4 = 34
                var v9 uint8 = 2
                _ = v9
            }
            var v10 float32 = v3
            _ = fi7
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(fi11, fi11)
            _ = fi11
        }
        v4 = v4
        var v12 uint16 = 562
        _ = v12
    }
    v3 = v3
    v13 := []int{127, 32767}
    v4 = (v4 % ((v4) | 1))
    fmt.Println(v4, v4)
    fmt.Println(v3)
    v3 += float32(1.5)
    v4 = v4
    v3 -= (v3 + (0.0 * v3))
    var v14 uint = 0
    v15 := GenS1{F0: "a", F1: 133, F2: ("a" + "foo bar")}
    v14 -= v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
