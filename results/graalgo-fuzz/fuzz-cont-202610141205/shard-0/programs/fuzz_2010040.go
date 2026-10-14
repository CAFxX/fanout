package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -1
    var v2 string = "foo bar"
    var v3 float32 = -2.25
    var v4 bool = (4294967295 <= 1)
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 14; fi6++ {
        v3 -= ((0.1 - 3.14159) - v3)
        v4 = false
        _ = fi6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v1 -= v1
        _ = fi7
    }
    v8 := []int{32768, 324}
    if len(v8) > 0 {
        v8[0] = -104
    }
    v1 = v1
    v9 := GenS1{F0: 255}
    var v10 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
