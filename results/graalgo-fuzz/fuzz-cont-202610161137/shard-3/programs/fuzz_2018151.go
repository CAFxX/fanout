package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 1
    var v2 int = -2147483648
    var v3 float32 = float32(-2.25)
    v4 := []uint32{42}
    var v5 uint8 = 255
    v6 := GenS1{F0: -2147483648}
    var v7 int8 = 127
    v8 := GenS1{F0: 329}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (("hello" + "\n\t") + "a") {
    case "foo bar":
        var v9 int8 = 1
        _ = v9
    case "":
        v5 -= v5
    default:
        v2 += (-(v2))
    }
    v10 := GenS2{F0: v1}
    var v11 int64 = 1
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v11, v11, v5)
        var v13 bool = ((fi12 - fi12) != int(v5))
        for fi14 := 0; fi14 < 3; fi14++ {
            fmt.Println(fi12, v13)
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
