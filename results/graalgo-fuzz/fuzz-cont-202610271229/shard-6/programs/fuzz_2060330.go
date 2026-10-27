package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 string = "hello"
    var v5 int64 = 42
    v6 := []uint32{}
    v7 := GenS1{F0: 317}
    var v8 bool = (false || (!true))
    v9 := GenS1{F0: 67}
    fmt.Println(v1, v4, v4)
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int8 = 0
            _ = fi12
            _ = v13
        }
        if ((613 >= 255) || v8) {
            var v14 uint16 = 817
            v1 = v1
            v2[4294967295] = (v14 != v14)
            _ = v14
        } else {
            v5 -= 1
        }
        switch fi11 {
        case -64:
            var v15 int64 = v5
            _ = v15
        case 815:
            if (v8 && false) {
                fmt.Println(v5, v4, v8)
            } else {
                v4 = ((v4 + v4) + "hello")
                delete(v3, 4294967295)
            }
        case -949:
            var v16 int64 = (v5 ^ v5)
            _ = v16
        }
        _ = fi11
    }
    var v17 bool = (v8 && v8)
    v5 += (v5 + 0)
    v10 = (v10 * (v10 * ((v10 / v10) - v10)))
    var v18 bool = v8
    fmt.Println(v17, v8, v4)
    if (v8 || true) {
        var v19 uint32 = 2147483647
        v19 += (v19 * (^(v19)))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
