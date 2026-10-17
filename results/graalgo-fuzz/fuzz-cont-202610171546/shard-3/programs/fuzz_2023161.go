package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }

func main() {
    var v1 uint8 = 167
    var v2 float32 = 3.14159
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 uint8 = v1
    if (v1 <= v1) {
        _ = len(v4)
        var v6 uint32 = 2147483648
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: v6}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    v9 := []int{}
    var v10 int8 = 42
    var v11 int = 127
    v10 += 0
    fmt.Println(v5)
    var v12 int32 = -70
    fmt.Println(v11)
    v13 := GenS2{F0: (v3 && (!(255 <= 0))), F1: (-(v10)), F2: 42}
    for fi14 := 0; fi14 < 1; fi14++ {
        v4[(("" + "hello") + "")] = (v11 >> ((65535) & 63))
        if false {
            var v15 float32 = float32(1e10)
            _ = v15
        } else {
            var v16 float32 = v2
            _ = v16
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
