package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v4 := []string{"hello"}
    v5 := []uint32{}
    var v6 uint16 = 255
    fmt.Println(v1)
    v7 := v2[2147483647]
    if (255 == 2) {
        _ = len(v4)
        if false {
            v8 := GenS1{F0: 0, F1: (v6 << ((0) & 63)), F2: 0}
            fmt.Println(v7, v6, v1)
            var v9 uint8 = 215
            _ = v8
            _ = v9
        }
    } else {
        _ = len(v4)
        v7 = (v7 >> ((255) & 63))
    }
    fmt.Println(v6, v7)
    fmt.Println(v6, v7, v6)
    var v10 string = "a"
    if (!(!true)) {
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                var v12 uint8 = 0
                v1 -= v1
                v1 = (v1 * v1)
                _ = v12
            } else {
                v6 = v6
            }
            var v13 int8 = 42
            v4 = append(v4, "foo bar")
            _ = fi11
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := v2[(v7 & (v7 | (v7 & 2147483647)))]
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
