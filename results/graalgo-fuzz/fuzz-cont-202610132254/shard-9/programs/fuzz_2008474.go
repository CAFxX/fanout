package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 212
    var v3 int32 = -231
    var v4 uint16 = 42
    switch -1 {
    case 42:
        var v5 int = 4294967295
        _ = v5
    case 0:
        v3 = v3
    }
    v4 -= 255
    if (!true) {
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int = 9223372036854775807
                _ = fi7
                _ = v8
            }
            v4 = v4
            var v9 int32 = (v3 | -1)
            _ = fi6
            _ = v9
        }
        var v10 uint = 255
        _ = v10
    }
    var v11 int = 228
    v12 := GenS1{F0: 127, F1: (v2 | uint32(v4))}
    var v13 string = "foo bar"
    v14 := GenS1{F0: 45, F1: 1}
    delete(v1, v13)
    v15 := GenS2{F0: 42, F1: ""}
    v16 := GenS1{F0: 89, F1: v2}
    var v17 string = (v13 + "")
    v2 += v2
    v18 := []int{42}
    v11 += -9223372036854775808
    v17 = ("a" + v17)
    v17 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
