package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    var v3 string = ("" + "")
    var v4 int8 = 100
    v5 := v2[490]
    v5 = (65535 == 33)
    var v6 bool = v5
    v7 := []int{42}
    var v8 float32 = -2.25
    _ = len(v7)
    if (v1 < (v1 - 789)) {
        fmt.Println(v5, v5, v5)
    } else {
        var v9 int16 = 395
        var v10 float64 = 0.1
        v3 = (((v3 + "") + "foo bar") + v3)
        _ = v10
        _ = v9
    }
    v4 -= 127
    v1 += 1
    if ((!false) || ((42 == 0) || (true && true))) {
        var v11 uint16 = 2
        _ = v11
    } else {
        if (42 < 988) {
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int64 = -1
                _ = fi12
                _ = v13
            }
            v5 = ((733 < 255) || true)
            var v14 float32 = 1.5
            _ = v14
        }
        v15 := GenS1{F0: 550}
        v16 := GenS1{F0: 1}
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
