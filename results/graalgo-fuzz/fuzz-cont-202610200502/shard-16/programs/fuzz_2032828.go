package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int64 = -1
    v4 := []int{669, 42}
    v3 = (v3 + 0)
    var v5 uint64 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (v3 > (-(v3))) {
        fmt.Println(v5, v5)
    }
    v3 = v3
    for fi6 := 0; fi6 < 10; fi6++ {
        v5 = 0
        _ = fi6
    }
    v7 := GenS1{F0: 2, F1: 565, F2: 127}
    v5 = v5
    v3 = v3
    var v8 int16 = -201
    var v9 uint16 = 153
    if (true && false) {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS2{F0: v8, F1: -1}
            v9 = v9
            v2["a"] = "foo bar"
            _ = fi10
            _ = v11
        }
    }
    if true {
        for fi12 := 0; fi12 < 17; fi12++ {
            delete(v2, "")
            var v13 float64 = 0.1
            _ = fi12
            _ = v13
        }
        fmt.Println(v9, v9, v8)
        var v14 bool = false
        _ = v14
    } else {
        switch "hello" {
        case "hello":
            v9 = (v9 | v9)
        }
    }
    v3 -= v3
    var v15 uint8 = 255
    var v16 int64 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
