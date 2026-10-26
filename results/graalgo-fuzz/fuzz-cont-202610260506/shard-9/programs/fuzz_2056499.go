package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    var v1 uint16 = 42
    var v2 string = (("foo bar" + "a") + "foo bar")
    var v3 bool = (-2.25 >= 0.1)
    v4 := make(map[int]bool)
    var v5 uint8 = uint8(v1)
    var v6 uint8 = v5
    switch -1 {
    case 685:
        fmt.Println(v3, v3)
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            v1 -= 672
            var v8 uint32 = 395
            v3 = (int(v1) <= -657)
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 42}
    v10 := []int{0}
    v3 = ((v3 && ((1 < 128) || (!v3))) && v3)
    var v11 float64 = ((1.5 + 3.14159) * (1.5 / 1.5))
    if (((!v3) && (1000 > -1)) && (!false)) {
        fmt.Println(v5, v1, v1)
    } else {
        _ = len(v10)
    }
    fmt.Println(v1)
    fmt.Println(v1)
    if (float32(v11) == -2.25) {
        var v12 int = 0
        _ = v12
    }
    fmt.Println(v6, v1)
    v13 := GenS2{F0: -1, F1: -1}
    _ = len(v10)
    v2 = ""
    v5 -= (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
